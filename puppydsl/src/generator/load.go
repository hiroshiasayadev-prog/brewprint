package generator

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/puppydsl/src/native"
	"gopkg.in/yaml.v3"
)

func loadProgram(root string) (program, error) {
	absoluteRoot, err := filepath.Abs(root)
	if err != nil {
		return program{}, fmt.Errorf("resolve PuppyDSL root: %w", err)
	}
	info, err := os.Stat(absoluteRoot)
	if err != nil {
		return program{}, fmt.Errorf("stat PuppyDSL root: %w", err)
	}
	if !info.IsDir() {
		return program{}, fmt.Errorf("PuppyDSL root is not a directory: %s", absoluteRoot)
	}

	goNativeRoot := filepath.Join(absoluteRoot, "go_native")
	moduleRoot, modulePath, err := findGoModule(goNativeRoot)
	if err != nil {
		return program{}, err
	}
	outputDirectory := filepath.Join(goNativeRoot, "puppygen")
	outputRelative, err := filepath.Rel(moduleRoot, outputDirectory)
	if err != nil {
		return program{}, fmt.Errorf("resolve generated import path: %w", err)
	}
	if strings.HasPrefix(outputRelative, "..") {
		return program{}, fmt.Errorf("PuppyDSL root is outside Go module: %s", absoluteRoot)
	}

	sources, fingerprint, err := discoverSources(absoluteRoot)
	if err != nil {
		return program{}, err
	}
	manifests, manifestFingerprint, err := loadManifests(absoluteRoot)
	if err != nil {
		return program{}, err
	}
	loaded := program{
		Root:                absoluteRoot,
		ModuleRoot:          moduleRoot,
		ModulePath:          modulePath,
		OutputImportPath:    joinImportPath(modulePath, outputRelative),
		NativeImportPath:    reflect.TypeOf(native.Binding{}).PkgPath(),
		SourceFingerprint:   fingerprint,
		ManifestFingerprint: manifestFingerprint,
		Manifests:           manifests,
	}

	typeIDs := map[string]string{}
	functionIDs := map[string]string{}
	for _, source := range sources {
		if err := loadSource(&loaded, source, typeIDs, functionIDs); err != nil {
			return program{}, err
		}
	}

	sort.Slice(loaded.Types, func(left, right int) bool {
		return loaded.Types[left].ID < loaded.Types[right].ID
	})
	sort.Slice(loaded.Functions, func(left, right int) bool {
		return loaded.Functions[left].ID < loaded.Functions[right].ID
	})
	if err := validateProgram(loaded); err != nil {
		return program{}, err
	}
	return loaded, nil
}

func discoverSources(root string) ([]string, string, error) {
	var sources []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if path != root && isGeneratedPuppygenDirectory(entry.Name()) {
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

func loadSource(loaded *program, source string, typeIDs, functionIDs map[string]string) error {
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

	relative, err := filepath.Rel(loaded.Root, source)
	if err != nil {
		return fmt.Errorf("resolve source path %s: %w", source, err)
	}
	relative = filepath.ToSlash(relative)

	switch rootName {
	case "types":
		return loadTypes(loaded, declarations, relative, typeIDs)
	case "functions":
		return loadFunctions(loaded, declarations, relative, functionIDs)
	case "contexts", "events", "modules":
		return nil
	default:
		return fmt.Errorf("%s: unsupported declaration root %q", source, rootName)
	}
}

func loadTypes(loaded *program, declarations *yaml.Node, source string, seen map[string]string) error {
	for index := 0; index < len(declarations.Content); index += 2 {
		id := declarations.Content[index].Value
		body := declarations.Content[index+1]
		if previous, exists := seen[id]; exists {
			return fmt.Errorf("duplicate type %q in %s and %s", id, previous, source)
		}
		declaration, err := decodeType(id, body, source)
		if err != nil {
			return err
		}
		seen[id] = source
		loaded.Types = append(loaded.Types, declaration)
	}
	return nil
}

func decodeType(id string, body *yaml.Node, source string) (typeDeclaration, error) {
	if body.Kind != yaml.MappingNode {
		return typeDeclaration{}, fmt.Errorf("%s: type %s body must be a mapping", source, id)
	}
	declaration := typeDeclaration{ID: id, Source: source}
	kindCount := 0
	for index := 0; index < len(body.Content); index += 2 {
		key := body.Content[index].Value
		value := body.Content[index+1]
		switch key {
		case "detail":
			declaration.Detail = value.Value
		case "scalar":
			declaration.Kind = typeKindScalar
			declaration.Scalar = value.Value
			kindCount++
		case "enum":
			declaration.Kind = typeKindEnum
			members, err := scalarSequence(value)
			if err != nil {
				return typeDeclaration{}, fmt.Errorf("%s: enum %s: %w", source, id, err)
			}
			declaration.Members = members
			kindCount++
		case "record":
			declaration.Kind = typeKindRecord
			if value.Kind != yaml.MappingNode {
				return typeDeclaration{}, fmt.Errorf("%s: record %s fields must be a mapping", source, id)
			}
			for fieldIndex := 0; fieldIndex < len(value.Content); fieldIndex += 2 {
				fieldName := value.Content[fieldIndex].Value
				fieldType, err := parseTypeExpressionSource(value.Content[fieldIndex+1].Value)
				if err != nil {
					return typeDeclaration{}, fmt.Errorf("%s: record %s field %s: %w", source, id, fieldName, err)
				}
				declaration.Fields = append(declaration.Fields, recordField{Name: fieldName, Type: fieldType})
			}
			kindCount++
		case "union":
			declaration.Kind = typeKindUnion
			variants, err := scalarSequence(value)
			if err != nil {
				return typeDeclaration{}, fmt.Errorf("%s: union %s: %w", source, id, err)
			}
			declaration.Variants = variants
			kindCount++
		case "opaque":
			if value.Value != "true" {
				return typeDeclaration{}, fmt.Errorf("%s: opaque type %s must declare opaque: true", source, id)
			}
			declaration.Kind = typeKindOpaque
			kindCount++
		case "generic":
			return typeDeclaration{}, fmt.Errorf("%s: generic type generation is not implemented for %s", source, id)
		}
	}
	if declaration.Detail == "" {
		return typeDeclaration{}, fmt.Errorf("%s: type %s has no detail", source, id)
	}
	if kindCount != 1 {
		return typeDeclaration{}, fmt.Errorf("%s: type %s must declare exactly one type body", source, id)
	}
	return declaration, nil
}

func loadFunctions(loaded *program, declarations *yaml.Node, source string, seen map[string]string) error {
	for index := 0; index < len(declarations.Content); index += 2 {
		id := declarations.Content[index].Value
		body := declarations.Content[index+1]
		if previous, exists := seen[id]; exists {
			return fmt.Errorf("duplicate function %q in %s and %s", id, previous, source)
		}
		seen[id] = source
		if strings.Contains(id, "<") {
			// Go generic functions require concrete type arguments before they can
			// be used as function values. The initial generated facade catalog is
			// therefore limited to non-generic declarations.
			continue
		}
		declaration, err := decodeFunction(id, body, source)
		if err != nil {
			return err
		}
		loaded.Functions = append(loaded.Functions, declaration)
	}
	return nil
}

func decodeFunction(id string, body *yaml.Node, source string) (functionDeclaration, error) {
	if body.Kind != yaml.MappingNode {
		return functionDeclaration{}, fmt.Errorf("%s: function %s body must be a mapping", source, id)
	}
	if strings.Contains(id, "<") {
		return functionDeclaration{}, fmt.Errorf("%s: generic function generation is not implemented for %s", source, id)
	}
	declaration := functionDeclaration{ID: id, Source: source}
	for index := 0; index < len(body.Content); index += 2 {
		key := body.Content[index].Value
		value := body.Content[index+1]
		switch key {
		case "detail":
			declaration.Detail = value.Value
		case "signature":
			declaration.SignatureSource = value.Value
		case "implementation":
			kind, bindingID, err := decodeImplementation(value)
			if err != nil {
				return functionDeclaration{}, fmt.Errorf("%s: function %s: %w", source, id, err)
			}
			declaration.Implementation = kind
			declaration.NativeBindingID = bindingID
		}
	}
	if declaration.Detail == "" {
		return functionDeclaration{}, fmt.Errorf("%s: function %s has no detail", source, id)
	}
	if declaration.SignatureSource == "" {
		return functionDeclaration{}, fmt.Errorf("%s: function %s has no signature", source, id)
	}
	if declaration.Implementation == "" {
		return functionDeclaration{}, fmt.Errorf("%s: function %s has no implementation", source, id)
	}
	signature, err := parseFunctionSignature(declaration.SignatureSource)
	if err != nil {
		return functionDeclaration{}, fmt.Errorf("%s: function %s signature: %w", source, id, err)
	}
	declaration.Signature = signature

	if id == "main" {
		declaration.Terminal = "main"
		declaration.GoName = "Main"
		return declaration, nil
	}
	segments := strings.Split(id, ".")
	if len(segments) < 2 {
		return functionDeclaration{}, fmt.Errorf("%s: function %s must be qualified", source, id)
	}
	declaration.Namespace = append([]string(nil), segments[:len(segments)-1]...)
	declaration.Terminal = segments[len(segments)-1]
	declaration.GoName = goExportedName(strings.TrimPrefix(declaration.Terminal, "_"))
	if declaration.GoName == "" {
		return functionDeclaration{}, fmt.Errorf("%s: function %s has no Go symbol", source, id)
	}
	return declaration, nil
}

func decodeImplementation(node *yaml.Node) (implementationKind, string, error) {
	if node.Kind != yaml.MappingNode {
		return "", "", fmt.Errorf("implementation must be a mapping")
	}
	var kind implementationKind
	var bindingID string
	for index := 0; index < len(node.Content); index += 2 {
		key := node.Content[index].Value
		value := node.Content[index+1]
		switch key {
		case "native":
			if kind != "" {
				return "", "", fmt.Errorf("multiple implementation forms")
			}
			kind = implementationNative
			bindingID = value.Value
		case "steps":
			if kind != "" {
				return "", "", fmt.Errorf("multiple implementation forms")
			}
			kind = implementationSteps
		case "stub":
			if kind != "" {
				return "", "", fmt.Errorf("multiple implementation forms")
			}
			kind = implementationStub
		case "return":
			// The return key belongs to a steps implementation and is not a fourth form.
		}
	}
	if kind == implementationNative && bindingID == "" {
		return "", "", fmt.Errorf("native implementation has an empty binding ID")
	}
	return kind, bindingID, nil
}

func scalarSequence(node *yaml.Node) ([]string, error) {
	if node.Kind != yaml.SequenceNode {
		return nil, fmt.Errorf("expected sequence")
	}
	values := make([]string, 0, len(node.Content))
	for _, item := range node.Content {
		if item.Kind != yaml.ScalarNode {
			return nil, fmt.Errorf("expected scalar sequence item")
		}
		values = append(values, item.Value)
	}
	return values, nil
}

func findGoModule(start string) (string, string, error) {
	current := start
	for {
		goMod := filepath.Join(current, "go.mod")
		content, err := os.ReadFile(goMod)
		if err == nil {
			for _, line := range strings.Split(string(content), "\n") {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "module ") {
					modulePath := strings.TrimSpace(strings.TrimPrefix(line, "module "))
					if modulePath == "" {
						break
					}
					return current, modulePath, nil
				}
			}
			return "", "", fmt.Errorf("%s has no module directive", goMod)
		}
		if !os.IsNotExist(err) {
			return "", "", fmt.Errorf("read %s: %w", goMod, err)
		}
		parent := filepath.Dir(current)
		if parent == current {
			return "", "", fmt.Errorf("no go.mod found above %s", start)
		}
		current = parent
	}
}

func joinImportPath(modulePath, relative string) string {
	relative = filepath.ToSlash(relative)
	if relative == "." || relative == "" {
		return modulePath
	}
	return strings.TrimSuffix(modulePath, "/") + "/" + strings.TrimPrefix(relative, "/")
}
