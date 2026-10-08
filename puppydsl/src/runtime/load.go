package runtime

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/puppydsl/src/native"
	"gopkg.in/yaml.v3"
)

// Load discovers one PuppyDSL source set and prepares it for execution.
func Load(root string, options Options) (*Program, error) {
	absoluteRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolve PuppyDSL root: %w", err)
	}
	info, err := os.Stat(absoluteRoot)
	if err != nil {
		return nil, fmt.Errorf("stat PuppyDSL root: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("PuppyDSL root is not a directory: %s", absoluteRoot)
	}

	sources, fingerprint, err := discoverSources(absoluteRoot)
	if err != nil {
		return nil, err
	}
	if expected := options.ExpectedSourceFingerprint; expected != "" && expected != fingerprint {
		return nil, fmt.Errorf(
			"PuppyDSL source fingerprint mismatch: loaded %q; generated %q",
			fingerprint,
			expected,
		)
	}

	program := &Program{
		root:              absoluteRoot,
		sourceFingerprint: fingerprint,
		functions:         map[string]functionDeclaration{},
		contexts:          map[string]struct{}{},
		nativeRegistry:    options.NativeRegistry,
		contextManager:    options.ContextManager,
	}
	for _, source := range sources {
		if err := loadSource(program, source); err != nil {
			return nil, err
		}
	}
	if _, exists := program.functions["main"]; !exists {
		return nil, fmt.Errorf("PuppyDSL source set has no main function declared by main.puppy.yaml")
	}
	if err := program.validateBootstrap(); err != nil {
		return nil, err
	}
	return program, nil
}

func discoverSources(root string) ([]string, string, error) {
	var sources []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if path != root && (entry.Name() == "puppygen" || strings.HasPrefix(entry.Name(), ".puppygen-")) {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(entry.Name(), ".puppy.yaml") {
			sources = append(sources, path)
		}
		return nil
	})
	if err != nil {
		return nil, "", fmt.Errorf("discover PuppyDSL source: %w", err)
	}
	sort.Strings(sources)
	if len(sources) == 0 {
		return nil, "", fmt.Errorf("no .puppy.yaml source found beneath %s", root)
	}

	hash := sha256.New()
	for _, source := range sources {
		relative, err := filepath.Rel(root, source)
		if err != nil {
			return nil, "", fmt.Errorf("resolve source path %s: %w", source, err)
		}
		content, err := os.ReadFile(source)
		if err != nil {
			return nil, "", fmt.Errorf("read PuppyDSL source %s: %w", source, err)
		}
		hash.Write([]byte(filepath.ToSlash(relative)))
		hash.Write([]byte{0})
		hash.Write(content)
		hash.Write([]byte{0})
	}
	return sources, hex.EncodeToString(hash.Sum(nil)), nil
}

func loadSource(program *Program, source string) error {
	content, err := os.ReadFile(source)
	if err != nil {
		return fmt.Errorf("read %s: %w", source, err)
	}
	var document yaml.Node
	if err := yaml.Unmarshal(content, &document); err != nil {
		return fmt.Errorf("decode %s: %w", source, err)
	}
	if len(document.Content) != 1 || document.Content[0].Kind != yaml.MappingNode {
		return fmt.Errorf("%s: expected one YAML mapping document", source)
	}
	root := document.Content[0]
	if len(root.Content) != 2 {
		return fmt.Errorf("%s: expected exactly one declaration root", source)
	}
	rootName := root.Content[0].Value
	declarations := root.Content[1]
	if declarations.Kind != yaml.MappingNode {
		return fmt.Errorf("%s: declaration root %q must be a mapping", source, rootName)
	}

	relative, err := filepath.Rel(program.root, source)
	if err != nil {
		return fmt.Errorf("resolve source path %s: %w", source, err)
	}
	relative = filepath.ToSlash(relative)

	switch rootName {
	case "functions":
		return loadFunctions(program, declarations, relative)
	case "contexts":
		return loadContexts(program, declarations, relative)
	case "types", "events", "modules":
		return nil
	default:
		return fmt.Errorf("%s: unsupported declaration root %q", source, rootName)
	}
}

func loadFunctions(program *Program, declarations *yaml.Node, source string) error {
	for index := 0; index < len(declarations.Content); index += 2 {
		id := declarations.Content[index].Value
		body := declarations.Content[index+1]
		if previous, exists := program.functions[id]; exists {
			return fmt.Errorf("duplicate function %q in %s and %s", id, previous.Source, source)
		}
		declaration, err := decodeFunction(id, body, source)
		if err != nil {
			return err
		}
		if id == "main" && source != "main.puppy.yaml" {
			return fmt.Errorf("%s: main must be declared by main.puppy.yaml", source)
		}
		program.functions[id] = declaration
	}
	return nil
}

func decodeFunction(id string, body *yaml.Node, source string) (functionDeclaration, error) {
	if body.Kind != yaml.MappingNode {
		return functionDeclaration{}, fmt.Errorf("%s: function %s body must be a mapping", source, id)
	}
	declaration := functionDeclaration{Source: source}
	var signature string
	var implementation *yaml.Node
	for index := 0; index < len(body.Content); index += 2 {
		key := body.Content[index].Value
		value := body.Content[index+1]
		switch key {
		case "detail":
			declaration.Detail = strings.TrimSpace(value.Value)
		case "opens":
			opens, err := scalarSequence(value)
			if err != nil {
				return functionDeclaration{}, fmt.Errorf("%s: function %s opens: %w", source, id, err)
			}
			declaration.Opens = opens
		case "signature":
			signature = canonicalSignature(value.Value)
		case "implementation":
			implementation = value
		}
	}
	if declaration.Detail == "" {
		return functionDeclaration{}, fmt.Errorf("%s: function %s has no detail", source, id)
	}
	if signature == "" {
		return functionDeclaration{}, fmt.Errorf("%s: function %s has no signature", source, id)
	}
	if implementation == nil {
		return functionDeclaration{}, fmt.Errorf("%s: function %s has no implementation", source, id)
	}
	declaration.Function = native.Function{ID: id, Signature: signature}
	if err := decodeImplementation(&declaration, implementation); err != nil {
		return functionDeclaration{}, fmt.Errorf("%s: function %s: %w", source, id, err)
	}
	return declaration, nil
}

func decodeImplementation(declaration *functionDeclaration, node *yaml.Node) error {
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("implementation must be a mapping")
	}
	for index := 0; index < len(node.Content); index += 2 {
		key := node.Content[index].Value
		value := node.Content[index+1]
		switch key {
		case "native":
			if declaration.Implementation != "" {
				return fmt.Errorf("multiple implementation forms")
			}
			if value.Kind != yaml.ScalarNode || strings.TrimSpace(value.Value) == "" {
				return fmt.Errorf("native implementation has an empty binding ID")
			}
			declaration.Implementation = implementationNative
			declaration.NativeBindingID = value.Value
		case "steps":
			if declaration.Implementation != "" {
				return fmt.Errorf("multiple implementation forms")
			}
			if value.Kind != yaml.MappingNode || len(value.Content) == 0 {
				return fmt.Errorf("steps implementation must be a non-empty mapping")
			}
			declaration.Implementation = implementationSteps
			for stepIndex := 0; stepIndex < len(value.Content); stepIndex += 2 {
				declaration.Steps = append(declaration.Steps, stepDeclaration{
					Name:       value.Content[stepIndex].Value,
					Expression: value.Content[stepIndex+1],
				})
			}
		case "stub":
			if declaration.Implementation != "" {
				return fmt.Errorf("multiple implementation forms")
			}
			declaration.Implementation = implementationStub
		case "return":
			if value.Kind != yaml.ScalarNode {
				return fmt.Errorf("steps return must be a scalar binding name")
			}
			declaration.ReturnBinding = value.Value
		default:
			return fmt.Errorf("unsupported implementation field %q", key)
		}
	}
	if declaration.Implementation == "" {
		return fmt.Errorf("implementation has no form")
	}
	if declaration.Implementation != implementationSteps && declaration.ReturnBinding != "" {
		return fmt.Errorf("return is permitted only with steps")
	}
	return nil
}

func loadContexts(program *Program, declarations *yaml.Node, source string) error {
	for index := 0; index < len(declarations.Content); index += 2 {
		id := declarations.Content[index].Value
		if _, exists := program.contexts[id]; exists {
			return fmt.Errorf("duplicate context %q in %s", id, source)
		}
		program.contexts[id] = struct{}{}
	}
	return nil
}

func (program *Program) validateBootstrap() error {
	for _, declaration := range program.functions {
		seenContexts := map[string]struct{}{}
		for _, contextID := range declaration.Opens {
			if _, exists := program.contexts[contextID]; !exists {
				return fmt.Errorf(
					"%s: function %s opens unknown context %q",
					declaration.Source,
					declaration.Function.ID,
					contextID,
				)
			}
			if _, duplicate := seenContexts[contextID]; duplicate {
				return fmt.Errorf(
					"%s: function %s opens context %q more than once",
					declaration.Source,
					declaration.Function.ID,
					contextID,
				)
			}
			seenContexts[contextID] = struct{}{}
		}

		if declaration.Implementation != implementationNative {
			continue
		}
		if program.nativeRegistry == nil {
			return fmt.Errorf("function %q requires native binding %q but no native registry was supplied", declaration.Function.ID, declaration.NativeBindingID)
		}
		binding, exists := program.nativeRegistry.ByFunctionID(declaration.Function.ID)
		if !exists {
			return fmt.Errorf("function %q has no activated native binding", declaration.Function.ID)
		}
		if binding.BindingID != declaration.NativeBindingID {
			return fmt.Errorf(
				"function %q declares native binding %q; activated binding is %q",
				declaration.Function.ID,
				declaration.NativeBindingID,
				binding.BindingID,
			)
		}
		if binding.Function.Signature != declaration.Function.Signature {
			return fmt.Errorf(
				"function %q signature mismatch: source %q; activated binding %q",
				declaration.Function.ID,
				declaration.Function.Signature,
				binding.Function.Signature,
			)
		}
	}
	return nil
}

func scalarSequence(node *yaml.Node) ([]string, error) {
	if node.Kind != yaml.SequenceNode {
		return nil, fmt.Errorf("expected sequence")
	}
	values := make([]string, 0, len(node.Content))
	for _, item := range node.Content {
		if item.Kind != yaml.ScalarNode || strings.TrimSpace(item.Value) == "" {
			return nil, fmt.Errorf("expected non-empty scalar sequence item")
		}
		values = append(values, item.Value)
	}
	return values, nil
}

func canonicalSignature(value string) string {
	return strings.Join(strings.Fields(value), " ")
}
