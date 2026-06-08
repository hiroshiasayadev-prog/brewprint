package api

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

type routeRow struct {
	TaskID  string
	Method  string
	Path    string
	Params  string
	Returns string
}

type section struct {
	Module string
	Rows   []routeRow
}

func RenderView(project *semantic.Project, viewID string) (string, error) {
	if project == nil {
		return "", fmt.Errorf("project is nil")
	}
	view := project.APIViewsByID[viewID]
	if view == nil {
		return "", fmt.Errorf("API view not found: %s", viewID)
	}
	return render(project, view), nil
}

func render(project *semantic.Project, view *semantic.APIView) string {
	sections := collectSections(project, view)

	var b strings.Builder
	b.WriteString("# " + view.ID + "\n\n")
	if strings.TrimSpace(view.Note) != "" {
		b.WriteString(strings.TrimSpace(view.Note) + "\n\n")
	}
	b.WriteString("## Routes\n\n")
	for _, section := range sections {
		b.WriteString("- [" + section.Module + "](#" + anchor(section.Module) + ")\n")
	}
	b.WriteString("\n")

	for _, section := range sections {
		b.WriteString("## " + section.Module + "\n\n")
		b.WriteString("| task id | method | path | params | returns |\n")
		b.WriteString("|---|---|---|---|---|\n")
		for _, row := range section.Rows {
			b.WriteString("| " + row.TaskID + " | " + row.Method + " | " + row.Path + " | " + row.Params + " | " + row.Returns + " |\n")
		}
		b.WriteString("\n")
	}
	return b.String()
}

func collectSections(project *semantic.Project, view *semantic.APIView) []section {
	var sections []section
	for _, viewModule := range view.Modules {
		rows := collectRows(project, view, viewModule)
		if len(rows) == 0 {
			continue
		}
		sections = append(sections, section{Module: viewModule.Module, Rows: rows})
	}
	return sections
}

func collectRows(project *semantic.Project, view *semantic.APIView, viewModule semantic.APIViewModule) []routeRow {
	var rows []routeRow
	for _, task := range project.TasksByQID {
		if !task.Endpoint {
			continue
		}
		module := moduleForFile(task.FileID)
		if !moduleIncluded(module, viewModule) {
			continue
		}
		taskID := taskDisplayID(viewModule.Module, module, task)
		routeID := taskRouteID(viewModule.Module, module, task)
		rows = append(rows, routeRow{
			TaskID:  taskID,
			Method:  absent(task.Method),
			Path:    fullPath(view.HTTPRootPath, routeID),
			Params:  paramsLabel(task),
			Returns: returnsLabel(task),
		})
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].TaskID < rows[j].TaskID })
	return rows
}

func moduleIncluded(module string, viewModule semantic.APIViewModule) bool {
	if module == viewModule.Module {
		return true
	}
	return viewModule.IncludeSubmodules && strings.HasPrefix(module, viewModule.Module+".")
}

func taskDisplayID(sectionModule, taskModule string, task *semantic.Task) string {
	return moduleRelativeID(sectionModule, taskModule, task.ID)
}

func taskRouteID(sectionModule, taskModule string, task *semantic.Task) string {
	leaf := task.Path
	if leaf == "" {
		leaf = task.ID
	}
	return moduleRelativeID(sectionModule, taskModule, leaf)
}

func moduleRelativeID(sectionModule, taskModule, leaf string) string {
	if taskModule == sectionModule {
		return leaf
	}
	prefix := strings.TrimPrefix(taskModule, sectionModule+".")
	if prefix != taskModule {
		return strings.ReplaceAll(prefix, ".", "/") + "/" + leaf
	}
	return strings.ReplaceAll(taskModule, ".", "/") + "/" + leaf
}

func fullPath(root, taskID string) string {
	root = strings.TrimRight(root, "/")
	if root == "" {
		root = "/"
	}
	taskID = strings.TrimLeft(taskID, "/")
	if root == "/" {
		return "/" + taskID
	}
	return root + "/" + taskID
}

func paramsLabel(task *semantic.Task) string {
	if len(task.Params) == 0 {
		return "-"
	}
	parts := make([]string, 0, len(task.Params))
	for _, param := range task.Params {
		if param.ModelName != "" {
			parts = append(parts, param.ModelName)
		} else {
			parts = append(parts, shortName(param.Model))
		}
	}
	return strings.Join(parts, "<br/>")
}

func returnsLabel(task *semantic.Task) string {
	if task.Returns == nil {
		return "-"
	}
	if task.Returns.ModelName != "" {
		return task.Returns.ModelName
	}
	return shortName(task.Returns.Model)
}

func absent(s string) string {
	if strings.TrimSpace(s) == "" {
		return "-"
	}
	return s
}

func shortName(qid semantic.QualifiedID) string {
	parts := strings.Split(qid.String(), ".")
	if len(parts) == 0 {
		return qid.String()
	}
	return parts[len(parts)-1]
}

func moduleForFile(fileID semantic.FileID) string {
	parts := strings.Split(fileID.String(), "/")
	if len(parts) == 0 {
		return ""
	}
	for i, part := range parts[:len(parts)-1] {
		if part == "model" || part == "store" || part == "task" || part == "state" || part == "event" || part == "branch" || part == "fork" || part == "join" {
			return strings.Join(parts[:i], ".")
		}
	}
	if len(parts) == 1 {
		return ""
	}
	return strings.Join(parts[:len(parts)-1], ".")
}

func anchor(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, " ", "-")
	return s
}
