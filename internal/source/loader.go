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

	if filepath.Base(filepath.Clean(root)) == "yaml" {
		renderIndexPath := filepath.Join(filepath.Dir(root), "render_index.yaml")
		if _, err := os.Stat(renderIndexPath); err == nil {
			file, err := loadFile(renderIndexPath, "render_index.yaml")
			if err != nil {
				return nil, err
			}
			files = append(files, file)
		} else if !os.IsNotExist(err) {
			return nil, err
		}
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
	if class.kind == rawyaml.FileKindRenderIndex {
		renderIndex, err := decodeRenderIndex(root)
		if err != nil {
			return rawyaml.File{}, fmt.Errorf("%s: %w", fileID, err)
		}
		file.RenderIndex = renderIndex
		return file, nil
	}
	if class.kind == rawyaml.FileKindView && class.viewAs == "sequence_diagram" {
		sequenceScenario, err := decodeSequenceScenario(root)
		if err != nil {
			return rawyaml.File{}, fmt.Errorf("%s: %w", fileID, err)
		}
		file.SequenceScenario = sequenceScenario
		return file, nil
	}
	if class.kind == rawyaml.FileKindView && class.viewAs == "er_diagram" {
		erView, err := decodeERView(root)
		if err != nil {
			return rawyaml.File{}, fmt.Errorf("%s: %w", fileID, err)
		}
		file.ERView = erView
		return file, nil
	}
	if class.kind == rawyaml.FileKindView && class.viewAs == "api_table" {
		apiView, err := decodeAPIView(root)
		if err != nil {
			return rawyaml.File{}, fmt.Errorf("%s: %w", fileID, err)
		}
		file.APIView = apiView
		return file, nil
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

func decodeRenderIndex(root *yaml.Node) (*rawyaml.RenderIndex, error) {
	var out rawyaml.RenderIndex
	if err := root.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func decodeSequenceScenario(root *yaml.Node) (*rawyaml.SequenceScenario, error) {
	var out rawyaml.SequenceScenario
	if err := root.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func decodeERView(root *yaml.Node) (*rawyaml.ERView, error) {
	var out rawyaml.ERView
	if err := root.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func decodeAPIView(root *yaml.Node) (*rawyaml.APIView, error) {
	var out rawyaml.APIView
	if err := root.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
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
		case "state":
			var state rawyaml.State
			if err := node.Decode(&state); err != nil {
				return nil, err
			}
			out.States = append(out.States, state)
		case "event":
			var event rawyaml.Event
			if err := node.Decode(&event); err != nil {
				return nil, err
			}
			out.Events = append(out.Events, event)
		case "branch":
			var branch rawyaml.ControlNode
			if err := node.Decode(&branch); err != nil {
				return nil, err
			}
			out.Branches = append(out.Branches, branch)
		case "fork":
			var fork rawyaml.ControlNode
			if err := node.Decode(&fork); err != nil {
				return nil, err
			}
			out.Forks = append(out.Forks, fork)
		case "join":
			var join rawyaml.ControlNode
			if err := node.Decode(&join); err != nil {
				return nil, err
			}
			out.Joins = append(out.Joins, join)
		default:
			out.Unsupported = append(out.Unsupported, common)
		}
	}
	if flowNode, ok := mappingValue(root, "flow"); ok {
		if flowNode.Kind != yaml.SequenceNode {
			return nil, fmt.Errorf("top-level flow must be a sequence")
		}
		if err := flowNode.Decode(&out.Flow); err != nil {
			return nil, err
		}
	}
	if transitionsNode, ok := mappingValue(root, "transitions"); ok {
		if transitionsNode.Kind != yaml.SequenceNode {
			return nil, fmt.Errorf("top-level transitions must be a sequence")
		}
		if err := transitionsNode.Decode(&out.Transitions); err != nil {
			return nil, err
		}
	}
	return out, nil
}
