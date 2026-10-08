package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateWritesRootRelativeFacadeTree(t *testing.T) {
	moduleRoot := t.TempDir()
	writeTestFile(t, filepath.Join(moduleRoot, "go.mod"), "module example.com/test\n\ngo 1.22\n")
	root := filepath.Join(moduleRoot, "app")
	writeTestFile(t, filepath.Join(root, "types.puppy.yaml"), `# responsibility: Defines test values.
# excludes: Other declarations.

types:
  Name:
    detail: Identifies one name.
    scalar: string

  Greeting:
    detail: Carries one greeting.
    record:
      name: Name

  DirectoryTree:
    detail: Preserves recursively nested directories by name.
    record:
      directories: dict<Name, DirectoryTree>
`)
	writeTestFile(t, filepath.Join(root, "functions.puppy.yaml"), `# responsibility: Defines test functions.
# excludes: Other declarations.

functions:
  demo.greet:
    detail: Produces one greeting.
    signature: >-
      (name: Name) -> Greeting
    implementation:
      native: demo.greet

  demo.notify:
    detail: Sends one notification.
    signature: >-
      () -> void
    implementation:
      stub: native

  demo.identity<T>:
    detail: Preserves one generic value.
    signature: >-
      (value: T) -> T
    implementation:
      stub: native
`)
	writeTestFile(t, filepath.Join(root, "manifests", "mcp", "tools.manifest.yaml"), `#bind: mcp_tools

shared: &shared
  enabled: true
tools:
  greet:
    detail: Greets one caller.
    defaults: *shared
`)

	if err := Generate(root); err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	typesSource := readTestFile(t, filepath.Join(root, "go_native", "puppygen", "types", "types.gen.go"))
	if !strings.Contains(typesSource, "type Name string") {
		t.Fatalf("generated types missing named scalar:\n%s", typesSource)
	}
	if !strings.Contains(typesSource, "type Greeting struct") {
		t.Fatalf("generated types missing record:\n%s", typesSource)
	}
	if !strings.Contains(typesSource, "Directories map[Name]DirectoryTree") {
		t.Fatalf("generated types missing recursive dictionary field:\n%s", typesSource)
	}

	functionSource := readTestFile(t, filepath.Join(root, "go_native", "puppygen", "demo", "functions.gen.go"))
	for _, expected := range []string{
		"func Greet(call native.Call, name types.Name) (types.Greeting, error)",
		"type GreetImplementation func(call native.Call, name types.Name) (types.Greeting, error)",
		"func BindGreet(implementation GreetImplementation) native.Binding",
		"if implementation == nil",
		"func(call native.Call, arguments []any) (any, error)",
		"native.ArgumentCountError(greetFunction, 1, len(arguments))",
		"input0, err := native.Input[types.Name](greetFunction, arguments, 0, \"name\")",
		"result, err := implementation(call, input0)",
		"func Notify(call native.Call) error",
	} {
		if !strings.Contains(functionSource, expected) {
			t.Fatalf("generated function source missing %q:\n%s", expected, functionSource)
		}
	}
	if strings.Contains(functionSource, "BindNotify") {
		t.Fatalf("stub function unexpectedly received binder:\n%s", functionSource)
	}

	metadata := readTestFile(t, filepath.Join(root, "go_native", "puppygen", "metadata.gen.go"))
	if !strings.Contains(metadata, "const SourceFingerprint =") {
		t.Fatalf("generated metadata missing fingerprint:\n%s", metadata)
	}

	catalog := readTestFile(t, filepath.Join(root, "go_native", "puppygen", "function_catalog.gen.go"))
	for _, expected := range []string{
		"var Functions = map[string]any",
		`"demo.greet"`,
		"functionpkg0.Greet",
		`"demo.notify"`,
		"functionpkg0.Notify",
	} {
		if !strings.Contains(catalog, expected) {
			t.Fatalf("generated function catalog missing %q:\n%s", expected, catalog)
		}
	}
	if strings.Index(catalog, `"demo.greet"`) > strings.Index(catalog, `"demo.notify"`) {
		t.Fatalf("generated function catalog is not ordered by FunctionID:\n%s", catalog)
	}
	if strings.Contains(catalog, `"demo.identity<T>"`) || strings.Contains(functionSource, "func Identity") {
		t.Fatalf("generic function unexpectedly entered the initial generated Go facade set:\n%s\n%s", catalog, functionSource)
	}

	manifestSource := readTestFile(t, filepath.Join(root, "go_native", "puppygen", "manifest", "manifest.gen.go"))
	for _, expected := range []string{
		"type Document struct",
		"const Fingerprint =",
		"var McpTools = Document",
		`BindName: "mcp_tools"`,
		`"manifests/mcp/tools.manifest.yaml"`,
		"func ByBindName(name string) (Document, bool)",
		".Alias = &nodes[",
	} {
		if !strings.Contains(manifestSource, expected) {
			t.Fatalf("generated manifest source missing %q:\n%s", expected, manifestSource)
		}
	}
}

func TestGenerateRejectsInvalidManifestEnvelopeAndDocument(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		wantText string
	}{
		{name: "missing header", content: "tools: {}\n", wantText: "physical line one"},
		{name: "blank before header", content: "\n#bind: mcp_tools\ntools: {}\n", wantText: "physical line one"},
		{name: "invalid bind name", content: "#bind: MCPTools\ntools: {}\n", wantText: "physical line one"},
		{name: "empty document", content: "#bind: mcp_tools\n", wantText: "contains no YAML document"},
		{name: "sequence root", content: "#bind: mcp_tools\n- tool\n", wantText: "root must be a mapping"},
		{name: "multiple documents", content: "#bind: mcp_tools\ntools: {}\n---\nother: {}\n", wantText: "exactly one YAML document"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			moduleRoot := t.TempDir()
			writeTestFile(t, filepath.Join(moduleRoot, "go.mod"), "module example.com/test\n\ngo 1.22\n")
			root := filepath.Join(moduleRoot, "app")
			writeMinimalProgram(t, root)
			writeTestFile(t, filepath.Join(root, "test.manifest.yaml"), test.content)

			err := Generate(root)
			if err == nil {
				t.Fatal("Generate() error = nil, want manifest validation error")
			}
			if !strings.Contains(err.Error(), test.wantText) {
				t.Fatalf("Generate() error = %q, want text %q", err, test.wantText)
			}
		})
	}
}

func TestGenerateRejectsManifestBindAndGoSymbolCollisions(t *testing.T) {
	tests := []struct {
		name     string
		first    string
		second   string
		wantText string
	}{
		{name: "duplicate bind", first: "#bind: mcp_tools\na: 1\n", second: "#bind: mcp_tools\nb: 2\n", wantText: "duplicate manifest bind name"},
		{name: "generated symbol", first: "#bind: foo_1\na: 1\n", second: "#bind: foo1\nb: 2\n", wantText: "collides as generated Go symbol Foo1"},
		{name: "reserved symbol", first: "#bind: document\na: 1\n", second: "", wantText: "collides as generated Go symbol Document"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			moduleRoot := t.TempDir()
			writeTestFile(t, filepath.Join(moduleRoot, "go.mod"), "module example.com/test\n\ngo 1.22\n")
			root := filepath.Join(moduleRoot, "app")
			writeMinimalProgram(t, root)
			writeTestFile(t, filepath.Join(root, "a.manifest.yaml"), test.first)
			if test.second != "" {
				writeTestFile(t, filepath.Join(root, "b.manifest.yaml"), test.second)
			}

			err := Generate(root)
			if err == nil {
				t.Fatal("Generate() error = nil, want manifest collision error")
			}
			if !strings.Contains(err.Error(), test.wantText) {
				t.Fatalf("Generate() error = %q, want text %q", err, test.wantText)
			}
		})
	}
}

func TestGenerateRejectsInvalidDictionaryKeyType(t *testing.T) {
	moduleRoot := t.TempDir()
	writeTestFile(t, filepath.Join(moduleRoot, "go.mod"), "module example.com/test\n\ngo 1.22\n")
	root := filepath.Join(moduleRoot, "app")
	writeTestFile(t, filepath.Join(root, "types.puppy.yaml"), `# responsibility: Defines invalid dictionary test types.
# excludes: Other declarations.

types:
  InvalidKey:
    detail: Provides one record that cannot be used as a dictionary key.
    record:
      value: string

  InvalidDictionary:
    detail: Attempts to use a record as a dictionary key.
    record:
      values: dict<InvalidKey, string>
`)

	err := Generate(root)
	if err == nil {
		t.Fatal("Generate() error = nil, want invalid dictionary-key error")
	}
	if !strings.Contains(err.Error(), "dict key") {
		t.Fatalf("Generate() error = %q, want dictionary-key diagnostic", err)
	}
}

func TestGenerateReplacesPreviousGeneratedTree(t *testing.T) {
	moduleRoot := t.TempDir()
	writeTestFile(t, filepath.Join(moduleRoot, "go.mod"), "module example.com/test\n\ngo 1.22\n")
	root := filepath.Join(moduleRoot, "app")
	writeTestFile(t, filepath.Join(root, "main.puppy.yaml"), `# responsibility: Defines the entry function.
# excludes: Other declarations.

functions:
  main:
    detail: Runs the application.
    signature: >-
      () -> void
    implementation:
      stub: native
`)
	writeTestFile(t, filepath.Join(root, "go_native", "puppygen", "stale.go"), "package puppygen\n")

	if err := Generate(root); err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "go_native", "puppygen", "stale.go")); !os.IsNotExist(err) {
		t.Fatalf("stale generated file still exists; stat error = %v", err)
	}
	mainSource := readTestFile(t, filepath.Join(root, "go_native", "puppygen", "functions.gen.go"))
	if !strings.Contains(mainSource, "func Main(call native.Call) error") {
		t.Fatalf("generated root facade missing Main:\n%s", mainSource)
	}
	catalog := readTestFile(t, filepath.Join(root, "go_native", "puppygen", "function_catalog.gen.go"))
	if !strings.Contains(catalog, `"main": Main`) {
		t.Fatalf("generated function catalog missing root Main facade:\n%s", catalog)
	}
}

func TestLoadCurrentDRMCPProgram(t *testing.T) {
	root := filepath.Join("..", "..", "..", "drmcp", "puppydsl")
	loaded, err := loadProgram(root)
	if err != nil {
		t.Fatalf("loadProgram(%s) error = %v", root, err)
	}
	if len(loaded.Types) == 0 || len(loaded.Functions) == 0 {
		t.Fatalf("current DRMCP program was not loaded: types=%d functions=%d", len(loaded.Types), len(loaded.Functions))
	}
	if len(loaded.Manifests) != 1 || loaded.Manifests[0].BindName != "mcp_tools" {
		t.Fatalf("current DRMCP manifests = %#v, want one mcp_tools binding", loaded.Manifests)
	}
	if err := renderProgram(loaded, t.TempDir()); err != nil {
		t.Fatalf("renderProgram(current DRMCP) error = %v", err)
	}
}

func TestParseNestedSignature(t *testing.T) {
	signature, err := parseFunctionSignature(`(
        values: list<optional<Name>>,
        lookup: dict<string, Name>
    ) -> list<Name>`)
	if err != nil {
		t.Fatalf("parseFunctionSignature() error = %v", err)
	}
	if len(signature.Inputs) != 2 || signature.Output.Name != "list" {
		t.Fatalf("unexpected signature: %#v", signature)
	}
}

func writeMinimalProgram(t *testing.T, root string) {
	t.Helper()
	writeTestFile(t, filepath.Join(root, "main.puppy.yaml"), `# responsibility: Defines one test entry function.
# excludes: Other declarations.

functions:
  main:
    detail: Runs the test application.
    signature: >-
      () -> void
    implementation:
      stub: native
`)
}

func writeTestFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll(%s) error = %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile(%s) error = %v", path, err)
	}
}

func readTestFile(t *testing.T, path string) string {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%s) error = %v", path, err)
	}
	return string(content)
}
