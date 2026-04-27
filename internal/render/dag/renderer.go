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

	b.WriteString("```mermaid\n")
	b.WriteString("flowchart TD\n")
	writeBoundary(&b, "params", paramNames(task))
	writeBoundary(&b, "returns", returnNames(task))
	if len(task.Params) > 0 || task.Returns != nil {
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
		fmt.Fprintf(&b, "  %s --> %s\n", task.ID, task.Returns.Name)
	}
	fmt.Fprintf(&b, "  %s ==> _end([End])\n\n", task.ID)

	b.WriteString("  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff\n")
	if len(storeClassNames) > 0 {
		b.WriteString("  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff\n")
	}
	b.WriteString("  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff\n")
	if len(task.Params) > 0 || task.Returns != nil {
		b.WriteString("  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff\n")
	}
	fmt.Fprintf(&b, "  class %s taskNode\n", task.ID)
	if len(storeClassNames) > 0 {
		fmt.Fprintf(&b, "  class %s storeNode\n", strings.Join(storeClassNames, ","))
	}
	b.WriteString("  class _start,_end terminalNode\n")
	boundaryNames := append(paramNames(task), returnNames(task)...)
	if len(boundaryNames) > 0 {
		fmt.Fprintf(&b, "  class %s boundaryNode\n", strings.Join(boundaryNames, ","))
	}
	b.WriteString("```\n\n")

	writeTasksDetail(&b, task)
	return b.String(), nil
}

func writeBoundary(b *strings.Builder, name string, nodes []string) {
	if len(nodes) == 0 {
		return
	}
	fmt.Fprintf(b, "  subgraph %s\n", name)
	for _, node := range nodes {
		fmt.Fprintf(b, "    %s([%s])\n", node, node)
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
		b.WriteString("#### Returns\n\n")
		b.WriteString("| name | model |\n")
		b.WriteString("|---|---|\n")
		fmt.Fprintf(b, "| %s | %s |\n\n", task.Returns.Name, task.Returns.ModelName)
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

func returnNames(task *semantic.Task) []string {
	if task.Returns == nil {
		return nil
	}
	return []string{task.Returns.Name}
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
