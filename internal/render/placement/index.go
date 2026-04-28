package placement

import (
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

type DAGEntry struct {
	GroupID string
	TaskID  string
	Title   string
	Path    string
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
	counts := map[string]int{}
	for _, entry := range r.DAGEntries(project) {
		counts[entry.GroupID]++
	}

	var b strings.Builder
	b.WriteString("# " + projectName + " render index\n\n")
	b.WriteString("| group | DAG | State | Sequence | Wireframe | ER | API |\n")
	b.WriteString("|---|---:|---:|---:|---:|---|---|\n")
	for _, group := range r.Groups {
		label := group.Label
		if label == "" {
			label = group.ID
		}
		dagCount := "-"
		if counts[group.ID] > 0 {
			dagCount = intString(counts[group.ID])
		}
		b.WriteString("| [" + label + "](" + group.ID + "/index.md) | " + dagCount + " | - | - | - | - | - |\n")
	}
	b.WriteString("| *(cross)* | - | - | - | - | [er](_cross/er.md) | [api](_cross/api.md) |\n")
	b.WriteString("| *(preview)* | - | - | - | [wireframe preview](_preview/wireframe.html) | - | - |\n")
	return b.String()
}

func (r *Resolver) GroupIndexMarkdown(groupID string, project *semantic.Project) string {
	var b strings.Builder
	b.WriteString("# " + groupID + " render index\n\n")
	b.WriteString("| kind | title | path |\n")
	b.WriteString("|---|---|---|\n")
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
