package generator

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

var bindHeaderPattern = regexp.MustCompile(`^#bind: ([a-z][a-z0-9]*(?:_[a-z0-9]+)*)\z`)

func loadManifests(root string) ([]manifestDeclaration, string, error) {
	sources, err := discoverManifestSources(root)
	if err != nil {
		return nil, "", err
	}

	hash := sha256.New()
	declarations := make([]manifestDeclaration, 0, len(sources))
	bindNames := map[string]string{}
	goSymbols := map[string]string{
		"ByBindName":  "generated lookup function",
		"Document":    "generated document type",
		"Fingerprint": "generated fingerprint constant",
	}

	for _, source := range sources {
		relative, err := filepath.Rel(root, source)
		if err != nil {
			return nil, "", fmt.Errorf("resolve manifest source path %s: %w", source, err)
		}
		relative = filepath.ToSlash(relative)

		content, err := os.ReadFile(source)
		if err != nil {
			return nil, "", fmt.Errorf("read manifest %s: %w", relative, err)
		}
		hash.Write([]byte(relative))
		hash.Write([]byte{0})
		hash.Write(content)
		hash.Write([]byte{0})

		declaration, err := decodeManifest(relative, content)
		if err != nil {
			return nil, "", err
		}
		if previous, exists := bindNames[declaration.BindName]; exists {
			return nil, "", fmt.Errorf(
				"duplicate manifest bind name %q in %s and %s",
				declaration.BindName,
				previous,
				relative,
			)
		}
		if previous, exists := goSymbols[declaration.GoName]; exists {
			return nil, "", fmt.Errorf(
				"manifest bind name %q in %s collides as generated Go symbol %s with %s",
				declaration.BindName,
				relative,
				declaration.GoName,
				previous,
			)
		}
		bindNames[declaration.BindName] = relative
		goSymbols[declaration.GoName] = fmt.Sprintf("manifest %q in %s", declaration.BindName, relative)
		declarations = append(declarations, declaration)
	}

	return declarations, hex.EncodeToString(hash.Sum(nil)), nil
}

func discoverManifestSources(root string) ([]string, error) {
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
		if strings.HasSuffix(entry.Name(), ".manifest.yaml") {
			sources = append(sources, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("discover bound manifests: %w", err)
	}
	sort.Strings(sources)
	return sources, nil
}

func isGeneratedPuppygenDirectory(name string) bool {
	return name == "puppygen" || strings.HasPrefix(name, ".puppygen-")
}

func decodeManifest(source string, content []byte) (manifestDeclaration, error) {
	lineEnd := bytes.IndexByte(content, '\n')
	var firstLine, body []byte
	if lineEnd < 0 {
		firstLine = content
	} else {
		firstLine = content[:lineEnd]
		body = content[lineEnd+1:]
	}
	firstLine = bytes.TrimSuffix(firstLine, []byte{'\r'})
	matches := bindHeaderPattern.FindSubmatch(firstLine)
	if matches == nil {
		return manifestDeclaration{}, fmt.Errorf(
			"%s: physical line one must be exactly #bind: <lower_snake_name>",
			source,
		)
	}

	bindName := string(matches[1])
	goName := goExportedName(bindName)
	if goName == "" {
		return manifestDeclaration{}, fmt.Errorf("%s: manifest bind name %q has no Go symbol", source, bindName)
	}

	decoder := yaml.NewDecoder(bytes.NewReader(body))
	var document yaml.Node
	if err := decoder.Decode(&document); err != nil {
		if err == io.EOF {
			return manifestDeclaration{}, fmt.Errorf("%s: manifest contains no YAML document", source)
		}
		return manifestDeclaration{}, fmt.Errorf("decode manifest %s: %w", source, err)
	}

	var additional yaml.Node
	if err := decoder.Decode(&additional); err == nil {
		return manifestDeclaration{}, fmt.Errorf("%s: manifest must contain exactly one YAML document", source)
	} else if err != io.EOF {
		return manifestDeclaration{}, fmt.Errorf("decode additional manifest document %s: %w", source, err)
	}

	if len(document.Content) != 1 || document.Content[0].Kind != yaml.MappingNode {
		return manifestDeclaration{}, fmt.Errorf("%s: manifest document root must be a mapping", source)
	}
	root := document.Content[0]
	shiftYAMLNodeLines(root, 1, map[*yaml.Node]struct{}{})

	return manifestDeclaration{
		BindName: bindName,
		GoName:   goName,
		Source:   source,
		Root:     root,
	}, nil
}

func shiftYAMLNodeLines(node *yaml.Node, offset int, seen map[*yaml.Node]struct{}) {
	if node == nil {
		return
	}
	if _, exists := seen[node]; exists {
		return
	}
	seen[node] = struct{}{}
	if node.Line > 0 {
		node.Line += offset
	}
	for _, child := range node.Content {
		shiftYAMLNodeLines(child, offset, seen)
	}
	shiftYAMLNodeLines(node.Alias, offset, seen)
}

func renderManifests(loaded program, outputDirectory string) error {
	directory := filepath.Join(outputDirectory, "manifest")
	var source bytes.Buffer
	source.WriteString(generatedHeader)
	source.WriteString("package manifest\n\n")
	source.WriteString("import yaml \"gopkg.in/yaml.v3\"\n\n")
	source.WriteString("// Document is one build-time-bound application manifest.\n")
	source.WriteString("type Document struct {\n")
	source.WriteString("\tBindName string\n")
	source.WriteString("\tSource string\n")
	source.WriteString("\tRoot yaml.Node\n")
	source.WriteString("}\n\n")
	source.WriteString("// Fingerprint identifies the exact sorted .manifest.yaml source set used for generation.\n")
	fmt.Fprintf(&source, "const Fingerprint = %s\n\n", strconv.Quote(loaded.ManifestFingerprint))

	for _, declaration := range loaded.Manifests {
		fmt.Fprintf(&source, "// %s is the manifest bound as %s.\n", declaration.GoName, strconv.Quote(declaration.BindName))
		fmt.Fprintf(&source, "var %s = Document{\n", declaration.GoName)
		fmt.Fprintf(&source, "\tBindName: %s,\n", strconv.Quote(declaration.BindName))
		fmt.Fprintf(&source, "\tSource: %s,\n", strconv.Quote(declaration.Source))
		fmt.Fprintf(&source, "\tRoot: build%sRoot(),\n", declaration.GoName)
		source.WriteString("}\n\n")
	}

	source.WriteString("var documentsByBindName = map[string]Document{\n")
	for _, declaration := range loaded.Manifests {
		fmt.Fprintf(&source, "\t%s: %s,\n", strconv.Quote(declaration.BindName), declaration.GoName)
	}
	source.WriteString("}\n\n")
	source.WriteString("// ByBindName resolves one generated manifest by its exact bind name.\n")
	source.WriteString("func ByBindName(name string) (Document, bool) {\n")
	source.WriteString("\tdocument, exists := documentsByBindName[name]\n")
	source.WriteString("\treturn document, exists\n")
	source.WriteString("}\n\n")

	for _, declaration := range loaded.Manifests {
		renderYAMLNodeBuilder(&source, declaration)
	}

	return writeFormatted(filepath.Join(directory, "manifest.gen.go"), source.Bytes())
}

func renderYAMLNodeBuilder(source *bytes.Buffer, declaration manifestDeclaration) {
	nodes, indexes := collectYAMLNodes(declaration.Root)
	fmt.Fprintf(source, "func build%sRoot() yaml.Node {\n", declaration.GoName)
	fmt.Fprintf(source, "\tnodes := make([]yaml.Node, %d)\n", len(nodes))
	for index, node := range nodes {
		fmt.Fprintf(source, "\tnodes[%d] = yaml.Node{\n", index)
		fmt.Fprintf(source, "\t\tKind: yaml.Kind(%d),\n", node.Kind)
		fmt.Fprintf(source, "\t\tStyle: yaml.Style(%d),\n", node.Style)
		fmt.Fprintf(source, "\t\tTag: %s,\n", strconv.Quote(node.Tag))
		fmt.Fprintf(source, "\t\tValue: %s,\n", strconv.Quote(node.Value))
		fmt.Fprintf(source, "\t\tAnchor: %s,\n", strconv.Quote(node.Anchor))
		fmt.Fprintf(source, "\t\tHeadComment: %s,\n", strconv.Quote(node.HeadComment))
		fmt.Fprintf(source, "\t\tLineComment: %s,\n", strconv.Quote(node.LineComment))
		fmt.Fprintf(source, "\t\tFootComment: %s,\n", strconv.Quote(node.FootComment))
		fmt.Fprintf(source, "\t\tLine: %d,\n", node.Line)
		fmt.Fprintf(source, "\t\tColumn: %d,\n", node.Column)
		source.WriteString("\t}\n")
	}
	for index, node := range nodes {
		if len(node.Content) > 0 {
			fmt.Fprintf(source, "\tnodes[%d].Content = []*yaml.Node{", index)
			for childIndex, child := range node.Content {
				if childIndex > 0 {
					source.WriteString(", ")
				}
				if child == nil {
					source.WriteString("nil")
					continue
				}
				fmt.Fprintf(source, "&nodes[%d]", indexes[child])
			}
			source.WriteString("}\n")
		}
		if node.Alias != nil {
			fmt.Fprintf(source, "\tnodes[%d].Alias = &nodes[%d]\n", index, indexes[node.Alias])
		}
	}
	source.WriteString("\treturn nodes[0]\n")
	source.WriteString("}\n\n")
}

func collectYAMLNodes(root *yaml.Node) ([]*yaml.Node, map[*yaml.Node]int) {
	indexes := map[*yaml.Node]int{}
	var nodes []*yaml.Node
	var visit func(*yaml.Node)
	visit = func(node *yaml.Node) {
		if node == nil {
			return
		}
		if _, exists := indexes[node]; exists {
			return
		}
		indexes[node] = len(nodes)
		nodes = append(nodes, node)
		for _, child := range node.Content {
			visit(child)
		}
		visit(node.Alias)
	}
	visit(root)
	return nodes, indexes
}
