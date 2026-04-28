package project

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml"
	apirender "github.com/hiroshiasayadev-prog/brewprint/internal/render/api"
	dagrender "github.com/hiroshiasayadev-prog/brewprint/internal/render/dag"
	errender "github.com/hiroshiasayadev-prog/brewprint/internal/render/er"
	"github.com/hiroshiasayadev-prog/brewprint/internal/render/placement"
	sequencerender "github.com/hiroshiasayadev-prog/brewprint/internal/render/sequence"
	staterender "github.com/hiroshiasayadev-prog/brewprint/internal/render/state"
	wireframerender "github.com/hiroshiasayadev-prog/brewprint/internal/render/wireframe"
	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

type File struct {
	Path    string
	Content string
}

func Render(raw *rawyaml.Project, semanticProject *semantic.Project) ([]File, []placement.Diagnostic, error) {
	resolver, diagnostics := placement.NewResolver(raw, semanticProject)
	if countPlacementErrors(diagnostics) > 0 {
		return nil, diagnostics, fmt.Errorf("render placement failed: %d error(s)", countPlacementErrors(diagnostics))
	}

	files, err := renderFiles(raw, semanticProject, resolver)
	if err != nil {
		return nil, diagnostics, err
	}
	files = append(indexFiles(raw, resolver, files), files...)
	sort.Slice(files, func(i, j int) bool { return files[i].Path < files[j].Path })
	return files, diagnostics, nil
}

func Write(outRoot string, files []File) error {
	if outRoot == "" {
		return fmt.Errorf("out root is required")
	}
	for _, file := range files {
		if file.Path == "" {
			return fmt.Errorf("render output path is empty")
		}
		path := filepath.Join(outRoot, filepath.FromSlash(file.Path))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return fmt.Errorf("create render output directory %s: %w", filepath.Dir(path), err)
		}
		if err := os.WriteFile(path, []byte(file.Content), 0o644); err != nil {
			return fmt.Errorf("write render output %s: %w", path, err)
		}
	}
	return nil
}

func renderFiles(raw *rawyaml.Project, semanticProject *semantic.Project, resolver *placement.Resolver) ([]File, error) {
	var files []File
	for _, task := range sortedMainTasks(semanticProject) {
		path, err := resolver.DAGPath(task)
		if err != nil {
			return nil, err
		}
		content, err := dagrender.RenderTask(semanticProject, task)
		if err != nil {
			return nil, err
		}
		files = append(files, File{Path: path, Content: content})
	}

	for _, fileID := range sortedStateFileIDs(semanticProject) {
		groupID, ok := resolver.GroupForFile(fileID)
		if !ok || groupID == "" {
			return nil, fmt.Errorf("no render group for state file %s", fileID)
		}
		content, err := staterender.RenderFile(semanticProject, fileID)
		if err != nil {
			return nil, err
		}
		files = append(files, File{Path: groupID + "/state-" + fsmID(fileID) + ".md", Content: content})
	}

	for _, scenarioID := range sortedScenarioIDs(semanticProject) {
		scenario := semanticProject.ScenariosByID[scenarioID]
		groupID, ok := resolver.GroupForFile(scenario.StateFile)
		if !ok || groupID == "" {
			return nil, fmt.Errorf("no render group for sequence scenario %s", scenario.ID)
		}
		content, err := sequencerender.RenderScenario(semanticProject, scenario.ID)
		if err != nil {
			return nil, err
		}
		files = append(files, File{Path: groupID + "/seq-" + scenario.ID + ".md", Content: content})
	}

	erIDs := sortedERViewIDs(semanticProject)
	for _, viewID := range erIDs {
		content, err := errender.RenderView(semanticProject, viewID)
		if err != nil {
			return nil, err
		}
		files = append(files, File{Path: crossViewPath("er", erIDs, viewID), Content: content})
	}

	apiIDs := sortedAPIViewIDs(semanticProject)
	for _, viewID := range apiIDs {
		content, err := apirender.RenderView(semanticProject, viewID)
		if err != nil {
			return nil, err
		}
		files = append(files, File{Path: crossViewPath("api", apiIDs, viewID), Content: content})
	}

	wireframeStates := sortedWireframeStates(semanticProject)
	for _, state := range wireframeStates {
		groupID, ok := resolver.GroupForFile(state.FileID)
		if !ok || groupID == "" {
			return nil, fmt.Errorf("no render group for wireframe state %s", state.QID)
		}
		content, err := wireframerender.RenderState(semanticProject, state.QID)
		if err != nil {
			return nil, err
		}
		files = append(files, File{Path: groupID + "/wireframe-" + fsmID(state.FileID) + "-" + state.ID + ".html", Content: content})
	}

	if len(wireframeStates) > 0 {
		files = append(files, File{Path: "_preview/wireframe.html", Content: wireframerender.RenderPreview(semanticProject, previewTitle(raw))})
	}
	return files, nil
}

func indexFiles(raw *rawyaml.Project, resolver *placement.Resolver, renderFiles []File) []File {
	files := []File{{Path: "index.md", Content: masterIndexMarkdown(projectName(raw), resolver, renderFiles)}}
	for _, group := range resolver.Groups {
		files = append(files, File{Path: group.ID + "/index.md", Content: groupIndexMarkdown(group.ID, renderFiles)})
	}
	return files
}

func masterIndexMarkdown(name string, resolver *placement.Resolver, files []File) string {
	if name == "" {
		name = "brewprint"
	}
	counts := countByGroupKind(files)
	var b strings.Builder
	b.WriteString("# " + name + " render index\n\n")
	b.WriteString("| group | DAG | State | Sequence | Wireframe | ER | API |\n")
	b.WriteString("|---|---:|---:|---:|---:|---|---|\n")
	for _, group := range resolver.Groups {
		label := group.Label
		if label == "" {
			label = group.ID
		}
		groupCounts := counts[group.ID]
		b.WriteString("| [" + label + "](" + group.ID + "/index.md) | " + countText(groupCounts["DAG"]) + " | " + countText(groupCounts["State"]) + " | " + countText(groupCounts["Sequence"]) + " | " + countText(groupCounts["Wireframe"]) + " | - | - |\n")
	}
	if hasPath(files, "_cross/er.md") || hasPath(files, "_cross/api.md") {
		b.WriteString("| *(cross)* | - | - | - | - | " + linkOrDash(files, "_cross/er.md", "er") + " | " + linkOrDash(files, "_cross/api.md", "api") + " |\n")
	}
	if hasPath(files, "_preview/wireframe.html") {
		b.WriteString("| *(preview)* | - | - | - | [wireframe preview](_preview/wireframe.html) | - | - |\n")
	}
	return b.String()
}

func groupIndexMarkdown(groupID string, files []File) string {
	entries := filesForGroup(groupID, files)
	var b strings.Builder
	b.WriteString("# " + groupID + " render index\n\n")
	b.WriteString("| kind | title | path |\n")
	b.WriteString("|---|---|---|\n")
	for _, entry := range entries {
		fileName := strings.TrimPrefix(entry.Path, groupID+"/")
		kind := kindForFileName(fileName)
		if kind == "" {
			continue
		}
		b.WriteString("| " + kind + " | " + titleForFileName(fileName) + " | [" + fileName + "](" + fileName + ") |\n")
	}
	return b.String()
}

func sortedMainTasks(project *semantic.Project) []*semantic.Task {
	var tasks []*semantic.Task
	if project != nil {
		for _, task := range project.TasksByQID {
			if task.Main {
				tasks = append(tasks, task)
			}
		}
	}
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].QID < tasks[j].QID })
	return tasks
}

func sortedStateFileIDs(project *semantic.Project) []semantic.FileID {
	seen := map[semantic.FileID]struct{}{}
	if project != nil {
		for fileID, nodes := range project.NodesByFile {
			for _, node := range nodes {
				switch node.(type) {
				case *semantic.State, *semantic.Event:
					seen[fileID] = struct{}{}
				}
			}
		}
		for fileID, transitions := range project.TransitionsByFile {
			if len(transitions) > 0 {
				seen[fileID] = struct{}{}
			}
		}
	}
	fileIDs := make([]semantic.FileID, 0, len(seen))
	for fileID := range seen {
		fileIDs = append(fileIDs, fileID)
	}
	sort.Slice(fileIDs, func(i, j int) bool { return fileIDs[i] < fileIDs[j] })
	return fileIDs
}

func sortedScenarioIDs(project *semantic.Project) []string {
	if project == nil {
		return nil
	}
	ids := make([]string, 0, len(project.ScenariosByID))
	for id := range project.ScenariosByID {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func sortedERViewIDs(project *semantic.Project) []string {
	if project == nil {
		return nil
	}
	ids := make([]string, 0, len(project.ERViewsByID))
	for id := range project.ERViewsByID {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func sortedAPIViewIDs(project *semantic.Project) []string {
	if project == nil {
		return nil
	}
	ids := make([]string, 0, len(project.APIViewsByID))
	for id := range project.APIViewsByID {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func sortedWireframeStates(project *semantic.Project) []*semantic.State {
	var states []*semantic.State
	if project != nil {
		for _, state := range project.StatesByQID {
			if state.Wireframe != nil {
				states = append(states, state)
			}
		}
	}
	sort.Slice(states, func(i, j int) bool {
		if states[i].FileID != states[j].FileID {
			return states[i].FileID < states[j].FileID
		}
		return states[i].ID < states[j].ID
	})
	return states
}

func crossViewPath(kind string, ids []string, id string) string {
	if len(ids) == 1 {
		return "_cross/" + kind + ".md"
	}
	return "_cross/" + kind + "-" + safePathID(id) + ".md"
}

func fsmID(fileID semantic.FileID) string {
	path := strings.TrimSuffix(fileID.String(), ".yaml")
	path = strings.TrimSuffix(path, ".yml")
	if strings.HasSuffix(path, "/state") {
		path = strings.TrimSuffix(path, "/state")
	}
	path = strings.ReplaceAll(path, "\\", "/")
	path = strings.Trim(path, "/")
	return strings.ReplaceAll(path, "/", "-")
}

func projectName(raw *rawyaml.Project) string {
	if raw == nil || raw.Root == "" {
		return ""
	}
	root := filepath.Clean(raw.Root)
	if filepath.Base(root) == "yaml" {
		return filepath.Base(filepath.Dir(root))
	}
	return filepath.Base(root)
}

func previewTitle(raw *rawyaml.Project) string {
	name := projectName(raw)
	if name == "" {
		return "Wireframe Preview"
	}
	return name + " Wireframe Preview"
}

func countByGroupKind(files []File) map[string]map[string]int {
	counts := map[string]map[string]int{}
	for _, file := range files {
		parts := strings.Split(file.Path, "/")
		if len(parts) < 2 || strings.HasPrefix(parts[0], "_") {
			continue
		}
		kind := kindForFileName(parts[len(parts)-1])
		if kind == "" {
			continue
		}
		if counts[parts[0]] == nil {
			counts[parts[0]] = map[string]int{}
		}
		counts[parts[0]][kind]++
	}
	return counts
}

func filesForGroup(groupID string, files []File) []File {
	var out []File
	prefix := groupID + "/"
	for _, file := range files {
		if strings.HasPrefix(file.Path, prefix) && kindForFileName(strings.TrimPrefix(file.Path, prefix)) != "" {
			out = append(out, file)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		ki := kindForFileName(strings.TrimPrefix(out[i].Path, prefix))
		kj := kindForFileName(strings.TrimPrefix(out[j].Path, prefix))
		if kindRank(ki) != kindRank(kj) {
			return kindRank(ki) < kindRank(kj)
		}
		return out[i].Path < out[j].Path
	})
	return out
}

func kindForFileName(fileName string) string {
	switch {
	case strings.HasPrefix(fileName, "dag-") && strings.HasSuffix(fileName, ".md"):
		return "DAG"
	case strings.HasPrefix(fileName, "state-") && strings.HasSuffix(fileName, ".md"):
		return "State"
	case strings.HasPrefix(fileName, "seq-") && strings.HasSuffix(fileName, ".md"):
		return "Sequence"
	case strings.HasPrefix(fileName, "wireframe-") && strings.HasSuffix(fileName, ".html"):
		return "Wireframe"
	default:
		return ""
	}
}

func kindRank(kind string) int {
	switch kind {
	case "DAG":
		return 1
	case "State":
		return 2
	case "Sequence":
		return 3
	case "Wireframe":
		return 4
	default:
		return 99
	}
}

func titleForFileName(fileName string) string {
	name := strings.TrimSuffix(strings.TrimSuffix(fileName, ".md"), ".html")
	for _, prefix := range []string{"dag-", "state-", "seq-", "wireframe-"} {
		name = strings.TrimPrefix(name, prefix)
	}
	return name
}

func countText(n int) string {
	if n == 0 {
		return "-"
	}
	return fmt.Sprint(n)
}

func hasPath(files []File, path string) bool {
	for _, file := range files {
		if file.Path == path {
			return true
		}
	}
	return false
}

func linkOrDash(files []File, path string, label string) string {
	if hasPath(files, path) {
		return "[" + label + "](" + path + ")"
	}
	return "-"
}

func safePathID(id string) string {
	id = strings.ReplaceAll(id, "/", "-")
	id = strings.ReplaceAll(id, "\\", "-")
	return id
}

func countPlacementErrors(diagnostics []placement.Diagnostic) int {
	count := 0
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == placement.SeverityError {
			count++
		}
	}
	return count
}
