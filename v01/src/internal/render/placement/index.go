package placement

import (
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

type DAGEntry struct {
	GroupID string
	TaskID  string
	Title   string
	Path    string
}

type ModelEntry struct {
	GroupID string
	ModelID string
	Title   string
	Path    string
}

func (r *Resolver) ModelEntries(project *semantic.Project) []ModelEntry {
	if r == nil || project == nil {
		return nil
	}
	entries := []ModelEntry{}
	for _, model := range project.ModelsByQID {
		if model.FilePrivate || !isModelFileID(model.FileID) {
			continue
		}
		groupID, ok := r.GroupForFile(model.FileID)
		if !ok || groupID == "" {
			continue
		}
		entries = append(entries, ModelEntry{
			GroupID: groupID,
			ModelID: model.ID,
			Title:   model.ID,
			Path:    groupID + "/model-" + model.ID + ".md",
		})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].GroupID != entries[j].GroupID {
			return entries[i].GroupID < entries[j].GroupID
		}
		return entries[i].ModelID < entries[j].ModelID
	})
	return entries
}

func (r *Resolver) DAGEntries(project *semantic.Project) []DAGEntry {
	if r == nil || project == nil {
		return nil
	}
	entries := []DAGEntry{}
	for _, task := range project.TasksByQID {
		if !task.Main {
			continue
		}
		path, err := r.DAGPath(task)
		if err != nil {
			continue
		}
		groupID, _ := r.GroupForFile(task.FileID)
		entries = append(entries, DAGEntry{GroupID: groupID, TaskID: task.ID, Title: task.ID, Path: path})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].GroupID != entries[j].GroupID {
			return entries[i].GroupID < entries[j].GroupID
		}
		return entries[i].TaskID < entries[j].TaskID
	})
	return entries
}

func (r *Resolver) MasterIndexMarkdown(projectName string, project *semantic.Project) string {
	if projectName == "" {
		projectName = "brewprint"
	}
	modelCounts := map[string]int{}
	for _, entry := range r.ModelEntries(project) {
		modelCounts[entry.GroupID]++
	}
	dagCounts := map[string]int{}
	for _, entry := range r.DAGEntries(project) {
		dagCounts[entry.GroupID]++
	}

	var b strings.Builder
	b.WriteString("# " + projectName + " render index\n\n")
	b.WriteString("| group | Model | DAG | State | Sequence | Wireframe | ER | API |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|---|---|\n")
	for _, group := range r.Groups {
		label := group.Label
		if label == "" {
			label = group.ID
		}
		modelCount := "-"
		if modelCounts[group.ID] > 0 {
			modelCount = intString(modelCounts[group.ID])
		}
		dagCount := "-"
		if dagCounts[group.ID] > 0 {
			dagCount = intString(dagCounts[group.ID])
		}
		b.WriteString("| [" + label + "](" + group.ID + "/index.md) | " + modelCount + " | " + dagCount + " | - | - | - | - | - |\n")
	}
	b.WriteString("| *(cross)* | - | - | - | - | - | [er](_cross/er.md) | [api](_cross/api.md) |\n")
	b.WriteString("| *(preview)* | - | - | - | - | [wireframe preview](_preview/wireframe.html) | - | - |\n")
	return b.String()
}

func (r *Resolver) GroupIndexMarkdown(groupID string, project *semantic.Project) string {
	var b strings.Builder
	b.WriteString("# " + groupID + " render index\n\n")
	b.WriteString("| kind | title | path |\n")
	b.WriteString("|---|---|---|\n")
	for _, entry := range r.ModelEntries(project) {
		if entry.GroupID != groupID {
			continue
		}
		fileName := strings.TrimPrefix(entry.Path, groupID+"/")
		b.WriteString("| Model | " + entry.Title + " | [" + fileName + "](" + fileName + ") |\n")
	}
	for _, entry := range r.DAGEntries(project) {
		if entry.GroupID != groupID {
			continue
		}
		fileName := strings.TrimPrefix(entry.Path, groupID+"/")
		b.WriteString("| DAG | " + entry.Title + " | [" + fileName + "](" + fileName + ") |\n")
	}
	return b.String()
}

func intString(n int) string {
	if n == 0 {
		return "0"
	}
	digits := []byte{}
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}

func isModelFileID(fileID semantic.FileID) bool {
	for _, part := range strings.Split(fileID.String(), "/") {
		if part == "model" {
			return true
		}
	}
	return false
}
