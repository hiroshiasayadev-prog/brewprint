package source

import (
	"path/filepath"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/rawyaml"
	"gopkg.in/yaml.v3"
)

type classification struct {
	kind   rawyaml.FileKind
	viewAs string
}

func classifyFile(fileID string, root *yaml.Node) classification {
	if filepath.Base(filepath.FromSlash(fileID)) == "render_index.yaml" {
		return classification{kind: rawyaml.FileKindRenderIndex}
	}
	if value, ok := mappingValue(root, "as"); ok {
		return classification{kind: rawyaml.FileKindView, viewAs: scalarString(value)}
	}
	if _, ok := mappingValue(root, "nodes"); ok {
		return classification{kind: rawyaml.FileKindNode}
	}
	return classification{kind: rawyaml.FileKindUnsupported}
}

func mappingValue(node *yaml.Node, key string) (*yaml.Node, bool) {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil, false
	}
	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			return node.Content[i+1], true
		}
	}
	return nil, false
}

func scalarString(node *yaml.Node) string {
	if node == nil || node.Kind != yaml.ScalarNode {
		return ""
	}
	return node.Value
}
