package dag

import (
	"fmt"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func RenderFile(project *semantic.Project, fileID semantic.FileID) (string, error) {
	qid, ok := project.MainNodeByFile[fileID]
	if !ok {
		return "", fmt.Errorf("main node not found for file %s", fileID)
	}
	task, ok := project.TasksByQID[qid]
	if !ok {
		return "", fmt.Errorf("main node is not a task: %s", qid)
	}
	return RenderTask(project, task)
}

func RenderTask(project *semantic.Project, task *semantic.Task) (string, error) {
	if len(project.FlowByFile[task.FileID]) > 0 {
		return renderFlowTask(project, task), nil
	}

	var b strings.Builder

	fmt.Fprintf(&b, "# %s\n\n", task.ID)
	if task.Endpoint {
		method := task.Method
		leafPath := task.Path
		if leafPath == "" {
			leafPath = task.ID
		}
		leafPath = strings.TrimPrefix(leafPath, "/")
		fmt.Fprintf(&b, "**API**: [%s /api/%s](../_cross/api.md)\n\n", method, leafPath)
	}
	if note := strings.TrimRight(task.Note, "\n"); note != "" {
		b.WriteString(note)
		b.WriteString("\n\n")
	}

	var assetRefs []*semantic.TypeRef
	for _, param := range task.Params {
		assetRefs = append(assetRefs, param.TypeRef)
	}
	if task.Returns != nil {
		assetRefs = append(assetRefs, task.Returns.TypeRef)
	}
	ambiguous := computeAmbiguousHints(assetRefs)

	b.WriteString("```mermaid\n")
	b.WriteString("flowchart TD\n")
	writeBoundary(&b, "params", paramBoundaryNodes(task, ambiguous))
	if len(task.Initializes) > 0 {
		fmt.Fprintf(&b, "  subgraph initializes\n")
		for _, init := range task.Initializes {
			fmt.Fprintf(&b, "    %s[(%s)]\n", init.Name, init.Name)
		}
		fmt.Fprintf(&b, "  end\n")
	}
	if len(task.Params) > 0 || len(task.Initializes) > 0 {
		b.WriteString("\n")
	}

	fmt.Fprintf(&b, "  _start([Start]) ==> %s[%s]\n", task.ID, task.ID)
	for _, param := range task.Params {
		fmt.Fprintf(&b, "  %s --> %s\n", param.Name, task.ID)
	}

	storeClassNames := storeNames(task)
	if len(storeClassNames) > 0 {
		b.WriteString("\n")
		for _, read := range task.Reads {
			if hasWrite(task, read.Name) {
				continue
			}
			fmt.Fprintf(&b, "  %s[(%s)] -- \"read\" --> %s\n", read.Name, read.Name, task.ID)
		}
		for _, write := range task.Writes {
			if hasRead(task, write.Name) {
				fmt.Fprintf(&b, "  %s <-- \"read/write\" --> %s[(%s)]\n", task.ID, write.Name, write.Name)
				continue
			}
			fmt.Fprintf(&b, "  %s -- \"write\" --> %s[(%s)]\n", task.ID, write.Name, write.Name)
		}
	}

	if task.Returns != nil {
		b.WriteString("\n")
		returnsHint := calcAssetHint(task.Returns.TypeRef, ambiguous)
		if returnsHint != "" {
			fmt.Fprintf(&b, "  %s --> %s([%s: %s])\n", task.ID, task.Returns.Name, task.Returns.Name, returnsHint)
		} else {
			fmt.Fprintf(&b, "  %s --> %s([%s])\n", task.ID, task.Returns.Name, task.Returns.Name)
		}
		if task.Returns.Source != "" {
			source := task.Returns.Source
			if task.Returns.SourceRef.Kind == semantic.FlowSourceParam {
				source = task.Returns.SourceRef.ParamName
			}
			if task.Returns.SourceRef.Kind == semantic.FlowSourceInitialized {
				source = task.Returns.SourceRef.Raw
			}
			if task.Returns.SourceRef.Kind == semantic.FlowSourceNode && task.Returns.SourceRef.AssetName != "" {
				source = task.Returns.SourceRef.AssetName
			}
			fmt.Fprintf(&b, "  %s -- \"returns as %s\" --> _end\n", source, task.Returns.Name)
		}
	}
	fmt.Fprintf(&b, "  %s ==> _end([End])\n\n", task.ID)

	b.WriteString("  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff\n")
	if len(storeClassNames) > 0 {
		b.WriteString("  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff\n")
	}
	if len(task.Initializes) > 0 {
		b.WriteString("  classDef initStoreNode fill:#F0C674,stroke:#B07820,color:#000\n")
	}
	if task.Returns != nil {
		b.WriteString("  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff\n")
	}
	b.WriteString("  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff\n")
	if len(task.Params) > 0 || task.Returns != nil {
		b.WriteString("  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff\n")
	}
	fmt.Fprintf(&b, "  class %s taskNode\n", task.ID)
	if len(storeClassNames) > 0 {
		fmt.Fprintf(&b, "  class %s storeNode\n", strings.Join(storeClassNames, ","))
	}
	if len(task.Initializes) > 0 {
		var initNames []string
		for _, init := range task.Initializes {
			initNames = append(initNames, init.Name)
		}
		fmt.Fprintf(&b, "  class %s initStoreNode\n", strings.Join(initNames, ","))
	}
	if task.Returns != nil {
		fmt.Fprintf(&b, "  class %s assetNode\n", task.Returns.Name)
	}
	b.WriteString("  class _start,_end terminalNode\n")
	boundaryNames := paramNames(task)
	if len(boundaryNames) > 0 {
		fmt.Fprintf(&b, "  class %s boundaryNode\n", strings.Join(boundaryNames, ","))
	}
	b.WriteString("```\n\n")

	writeTasksDetail(&b, task)
	writePrivateModelsSection(&b, project, task.FileID)
	return b.String(), nil
}

type boundaryNode struct {
	name string
	hint string
}

func writeBoundary(b *strings.Builder, name string, nodes []boundaryNode) {
	if len(nodes) == 0 {
		return
	}
	fmt.Fprintf(b, "  subgraph %s\n", name)
	for _, node := range nodes {
		if node.hint != "" {
			fmt.Fprintf(b, "    %s([%s: %s])\n", node.name, node.name, node.hint)
		} else {
			fmt.Fprintf(b, "    %s([%s])\n", node.name, node.name)
		}
	}
	b.WriteString("  end\n")
}

func writeTasksDetail(b *strings.Builder, task *semantic.Task) {
	b.WriteString("## Tasks\n\n")
	fmt.Fprintf(b, "### %s\n\n", task.ID)
	if len(task.Params) > 0 {
		b.WriteString("#### Params\n\n")
		b.WriteString("| name | model | note |\n")
		b.WriteString("|---|---|---|\n")
		for _, param := range task.Params {
			note := param.Note
			if note == "" {
				note = "—"
			}
			fmt.Fprintf(b, "| %s | %s | %s |\n", param.Name, param.ModelName, note)
		}
		b.WriteString("\n")
	}
	if task.Returns != nil {
		writeReturnsSection(b, task.Returns)
	}
	accessRows := storeAccessRows(task)
	if len(accessRows) > 0 {
		b.WriteString("#### Store access\n\n")
		b.WriteString("| access | store |\n")
		b.WriteString("|---|---|\n")
		for _, row := range accessRows {
			fmt.Fprintf(b, "| %s | %s |\n", row.Access, row.Store)
		}
	}
}

func paramNames(task *semantic.Task) []string {
	out := make([]string, 0, len(task.Params))
	for _, param := range task.Params {
		out = append(out, param.Name)
	}
	return out
}

func paramBoundaryNodes(task *semantic.Task, ambiguous map[string]bool) []boundaryNode {
	out := make([]boundaryNode, 0, len(task.Params))
	for _, param := range task.Params {
		out = append(out, boundaryNode{
			name: param.Name,
			hint: calcAssetHint(param.TypeRef, ambiguous),
		})
	}
	return out
}

func storeNames(task *semantic.Task) []string {
	seen := map[string]bool{}
	var out []string
	for _, read := range task.Reads {
		if !seen[read.Name] {
			seen[read.Name] = true
			out = append(out, read.Name)
		}
	}
	for _, write := range task.Writes {
		if !seen[write.Name] {
			seen[write.Name] = true
			out = append(out, write.Name)
		}
	}
	return out
}

func hasRead(task *semantic.Task, name string) bool {
	for _, read := range task.Reads {
		if read.Name == name {
			return true
		}
	}
	return false
}

func hasWrite(task *semantic.Task, name string) bool {
	for _, write := range task.Writes {
		if write.Name == name {
			return true
		}
	}
	return false
}

func storeAccessRows(task *semantic.Task) []StoreAccessRow {
	read := map[string]bool{}
	write := map[string]bool{}
	var order []string
	seen := map[string]bool{}
	for _, item := range task.Reads {
		read[item.Name] = true
		if !seen[item.Name] {
			seen[item.Name] = true
			order = append(order, item.Name)
		}
	}
	for _, item := range task.Writes {
		write[item.Name] = true
		if !seen[item.Name] {
			seen[item.Name] = true
			order = append(order, item.Name)
		}
	}
	var rows []StoreAccessRow
	for _, name := range order {
		access := "read"
		switch {
		case read[name] && write[name]:
			access = "read/write"
		case write[name]:
			access = "write"
		}
		rows = append(rows, StoreAccessRow{Access: access, Store: name})
	}
	return rows
}
