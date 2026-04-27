package source

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	"gopkg.in/yaml.v3"
)

type Loader struct{}

func (Loader) Load(root string) (*rawyaml.Project, error) {
	var files []rawyaml.File

	if err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || !isYAMLPath(path) {
			return nil
		}

		fileID, err := normalizeFileID(root, path)
		if err != nil {
			return err
		}
		file, err := loadFile(path, fileID)
		if err != nil {
			return err
		}
		files = append(files, file)
		return nil
	}); err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool { return files[i].ID < files[j].ID })
	return &rawyaml.Project{Root: root, Files: files}, nil
}

func loadFile(path, fileID string) (rawyaml.File, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return rawyaml.File{}, err
	}

	var doc yaml.Node
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return rawyaml.File{}, fmt.Errorf("%s: %w", fileID, err)
	}
	root := documentRoot(&doc)
	class := classifyFile(fileID, root)

	file := rawyaml.File{
		ID:     fileID,
		Path:   path,
		Kind:   class.kind,
		ViewAs: class.viewAs,
	}
	if class.kind != rawyaml.FileKindNode {
		return file, nil
	}

	nodeFile, err := decodeNodeFile(root)
	if err != nil {
		return rawyaml.File{}, fmt.Errorf("%s: %w", fileID, err)
	}
	file.NodeFile = nodeFile
	return file, nil
}

func documentRoot(doc *yaml.Node) *yaml.Node {
	if doc == nil {
		return nil
	}
	if doc.Kind == yaml.DocumentNode && len(doc.Content) > 0 {
		return doc.Content[0]
	}
	return doc
}

func decodeNodeFile(root *yaml.Node) (*rawyaml.NodeFile, error) {
	nodesNode, ok := mappingValue(root, "nodes")
	if !ok || nodesNode.Kind != yaml.SequenceNode {
		return nil, fmt.Errorf("top-level nodes must be a sequence")
	}

	out := &rawyaml.NodeFile{}
	for _, node := range nodesNode.Content {
		var common rawyaml.CommonNode
		if err := node.Decode(&common); err != nil {
			return nil, err
		}
		switch common.Type {
		case "task":
			var task rawyaml.Task
			if err := node.Decode(&task); err != nil {
				return nil, err
			}
			out.Tasks = append(out.Tasks, task)
		case "model":
			var model rawyaml.Model
			if err := node.Decode(&model); err != nil {
				return nil, err
			}
			out.Models = append(out.Models, model)
		case "store":
			var store rawyaml.Store
			if err := node.Decode(&store); err != nil {
				return nil, err
			}
			out.Stores = append(out.Stores, store)
		case "actor":
			var actor rawyaml.Actor
			if err := node.Decode(&actor); err != nil {
				return nil, err
			}
			out.Actors = append(out.Actors, actor)
		default:
			out.Unsupported = append(out.Unsupported, common)
		}
	}
	return out, nil
}
