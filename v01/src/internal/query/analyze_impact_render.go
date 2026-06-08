package query

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

const impactKindRenderOutput = "render_output"

type renderOutputMapping struct {
	Path   string
	Source *SourceLocation
}

func (s *Service) collectRenderAnalyzeImpacts(req AnalyzeImpactRequest, targetKey semantic.ObjectKey, target ObjectRef) ([]ImpactEntry, []semantic.Diagnostic) {
	if !analyzeImpactRenderChangeSupported(req.Change.Kind) {
		return []ImpactEntry{}, []semantic.Diagnostic{}
	}

	mappings := s.renderOutputMappingsForTarget(req, targetKey, target)
	mappings = dedupeRenderOutputMappings(mappings)

	impacts := make([]ImpactEntry, 0, len(mappings))
	diagnostics := []semantic.Diagnostic{}
	diagnosticFiles := map[string]struct{}{}
	for _, mapping := range mappings {
		if strings.TrimSpace(mapping.Path) == "" {
			continue
		}
		impact := renderOutputImpact(req.Change, target, mapping)
		if !s.impactInScope(impact, target, req.ScopeModules) {
			continue
		}
		if impact.Source != nil && impact.Source.File != "" && (impact.Source.Line == 0 || impact.Source.Column == 0) {
			impact.Fixability = "unknown"
			if _, seen := diagnosticFiles[impact.Source.File]; !seen {
				diagnostics = append(diagnostics, semantic.Diagnostic{
					Severity: semantic.SeverityWarning,
					Code:     "source_location_unavailable",
					FileID:   semantic.FileID(impact.Source.File),
					Message:  "source line/column is unavailable for render output impact; source file only is known",
				})
				diagnosticFiles[impact.Source.File] = struct{}{}
			}
		}
		impacts = append(impacts, impact)
	}

	sort.SliceStable(diagnostics, func(i, j int) bool {
		if diagnostics[i].Code != diagnostics[j].Code {
			return diagnostics[i].Code < diagnostics[j].Code
		}
		return diagnostics[i].FileID < diagnostics[j].FileID
	})
	return impacts, diagnostics
}

func analyzeImpactRenderChangeSupported(kind string) bool {
	switch kind {
	case AnalyzeImpactChangeRename, AnalyzeImpactChangeRemove, AnalyzeImpactChangeType, AnalyzeImpactChangeContract, AnalyzeImpactChangeTransitionTarget:
		return true
	default:
		return false
	}
}

func (s *Service) renderOutputMappingsForTarget(req AnalyzeImpactRequest, targetKey semantic.ObjectKey, target ObjectRef) []renderOutputMapping {
	switch {
	case target.Object == "field" || req.Selector.Object == "field" || isFieldKind(req.Selector.Kind):
		model, field, err := s.modelFieldBySelector(req.Selector)
		if err != nil {
			return nil
		}
		return s.renderMappingsForField(model, field)
	case target.Object == "transition":
		transition, err := s.transitionBySelector(req.Selector)
		if err != nil {
			return nil
		}
		return s.renderMappingsForTransition(transition)
	case target.Object == "node" && target.Kind == string(semantic.NodeKindTask):
		if task := s.taskForRenderTarget(targetKey, target); task != nil {
			return s.renderMappingsForTask(task)
		}
	case target.Object == "node" && target.Kind == string(semantic.NodeKindModel):
		if model := s.project.ModelsByQID[semantic.QualifiedID(target.ID)]; model != nil {
			return s.renderMappingsForModel(model)
		}
	}
	return nil
}

func (s *Service) taskForRenderTarget(targetKey semantic.ObjectKey, target ObjectRef) *semantic.Task {
	if task := s.project.TasksByQID[semantic.QualifiedID(target.ID)]; task != nil {
		return task
	}
	rawKey := string(targetKey)
	if rawKey != "" && !strings.Contains(rawKey, ":") {
		return s.project.TasksByQID[semantic.QualifiedID(rawKey)]
	}
	if target.File != "" && target.LocalID != "" {
		for _, node := range s.project.NodesByFile[semantic.FileID(target.File)] {
			task, ok := node.(*semantic.Task)
			if ok && task.ID == target.LocalID {
				return task
			}
		}
	}
	return nil
}

func (s *Service) renderMappingsForTask(task *semantic.Task) []renderOutputMapping {
	if task == nil {
		return nil
	}
	source := sourceLocationFromBlock(task.FileID, s.findNodeSource(task))
	mappings := s.renderMappingsForTaskFile(task.FileID, source)
	if task.Endpoint {
		mappings = append(mappings, s.renderMappingsForAPIViews(moduleForFileID(task.FileID), source)...)
	}
	return mappings
}

func (s *Service) renderMappingsForTaskFile(fileID semantic.FileID, source *SourceLocation) []renderOutputMapping {
	mainTask := s.mainTaskForFile(fileID)
	if mainTask == nil {
		return nil
	}
	groupID := s.renderGroupForFile(fileID)
	if groupID == "" {
		return nil
	}
	return s.renderMappingsWithGroupIndexes(groupID, groupID+"/dag-"+mainTask.ID+".md", source)
}

func (s *Service) renderMappingsForModel(model *semantic.Model) []renderOutputMapping {
	if model == nil {
		return nil
	}
	source := sourceLocationFromBlock(model.FileID, s.findNodeSource(model))
	mappings := []renderOutputMapping{}
	for _, ref := range s.project.ReferencesByTarget[semantic.NodeObjectKey(model.QID)] {
		sourceID := semantic.QualifiedID(ref.From.ID)
		if task := s.project.TasksByQID[sourceID]; task != nil {
			for _, mapping := range s.renderMappingsForTask(task) {
				mapping.Source = source
				mappings = append(mappings, mapping)
			}
		}
	}
	for _, store := range s.project.StoresByQID {
		if store != nil && store.Of == model.QID {
			mappings = append(mappings, s.renderMappingsForStore(store, source)...)
		}
	}
	return mappings
}

func (s *Service) renderMappingsForField(model *semantic.Model, field semantic.ModelField) []renderOutputMapping {
	if model == nil {
		return nil
	}
	source := sourceLocationFromBlock(model.FileID, s.findFieldSource(model, field))
	mappings := []renderOutputMapping{}
	for _, mapping := range s.renderMappingsForModel(model) {
		mapping.Source = source
		mappings = append(mappings, mapping)
	}
	for _, store := range s.project.StoresByQID {
		if store != nil && store.Of == model.QID {
			mappings = append(mappings, s.renderMappingsForERViews(moduleForFileID(store.FileID), source)...)
		}
	}
	return mappings
}

func (s *Service) renderMappingsForStore(store *semantic.Store, source *SourceLocation) []renderOutputMapping {
	if store == nil {
		return nil
	}
	mappings := []renderOutputMapping{}
	mappings = append(mappings, s.renderMappingsForERViews(moduleForFileID(store.FileID), source)...)
	for _, taskID := range s.project.TasksReadingStore[store.QID] {
		if task := s.project.TasksByQID[taskID]; task != nil {
			mappings = append(mappings, s.renderMappingsForTaskFile(task.FileID, source)...)
		}
	}
	for _, taskID := range s.project.TasksWritingStore[store.QID] {
		if task := s.project.TasksByQID[taskID]; task != nil {
			mappings = append(mappings, s.renderMappingsForTaskFile(task.FileID, source)...)
		}
	}
	return mappings
}

func (s *Service) renderMappingsForTransition(transition semantic.Transition) []renderOutputMapping {
	source := sourceLocationFromBlock(transition.FileID, s.findTransitionSource(transition))
	mappings := s.renderMappingsForStateFile(transition.FileID, source)
	transitionID := semantic.TransitionID(transition)
	for _, scenario := range s.project.ScenariosByID {
		if scenario == nil {
			continue
		}
		for _, step := range scenario.Steps {
			if semantic.TransitionID(step.Transition) == transitionID {
				mappings = append(mappings, s.renderMappingsForScenario(scenario, source)...)
				break
			}
		}
	}
	return mappings
}

func (s *Service) renderMappingsForStateFile(fileID semantic.FileID, source *SourceLocation) []renderOutputMapping {
	groupID := s.renderGroupForFile(fileID)
	if groupID == "" {
		return nil
	}
	return s.renderMappingsWithGroupIndexes(groupID, groupID+"/state-"+renderFSMID(fileID)+".md", source)
}

func (s *Service) renderMappingsForScenario(scenario *semantic.SequenceScenario, source *SourceLocation) []renderOutputMapping {
	if scenario == nil {
		return nil
	}
	groupID := s.renderGroupForFile(scenario.StateFile)
	if groupID == "" {
		return nil
	}
	return s.renderMappingsWithGroupIndexes(groupID, groupID+"/seq-"+scenario.ID+".md", source)
}

func (s *Service) renderMappingsForERViews(module string, source *SourceLocation) []renderOutputMapping {
	if module == "" {
		return nil
	}
	ids := sortedERViewIDsForQuery(s.project)
	mappings := []renderOutputMapping{}
	for _, id := range ids {
		view := s.project.ERViewsByID[id]
		if view == nil || !erViewIncludesModule(view, module) {
			continue
		}
		mappings = append(mappings, renderOutputMapping{Path: renderCrossViewPath("er", ids, id), Source: source})
		mappings = append(mappings, renderOutputMapping{Path: "index.md", Source: source})
	}
	return mappings
}

func (s *Service) renderMappingsForAPIViews(module string, source *SourceLocation) []renderOutputMapping {
	if module == "" {
		return nil
	}
	ids := sortedAPIViewIDsForQuery(s.project)
	mappings := []renderOutputMapping{}
	for _, id := range ids {
		view := s.project.APIViewsByID[id]
		if view == nil || !apiViewIncludesModule(view, module) {
			continue
		}
		mappings = append(mappings, renderOutputMapping{Path: renderCrossViewPath("api", ids, id), Source: source})
		mappings = append(mappings, renderOutputMapping{Path: "index.md", Source: source})
	}
	return mappings
}

func (s *Service) renderMappingsWithGroupIndexes(groupID string, path string, source *SourceLocation) []renderOutputMapping {
	if groupID == "" || path == "" {
		return nil
	}
	return []renderOutputMapping{
		{Path: path, Source: source},
		{Path: groupID + "/index.md", Source: source},
		{Path: "index.md", Source: source},
	}
}

func renderOutputImpact(change AnalyzeImpactChange, target ObjectRef, mapping renderOutputMapping) ImpactEntry {
	severity := "info"
	fixability := "manual_review"
	if change.Kind == AnalyzeImpactChangeRemove {
		severity = "warning"
	}
	if mapping.Source == nil || mapping.Source.File == "" || mapping.Source.Line == 0 || mapping.Source.Column == 0 {
		fixability = "unknown"
	}

	targetLabel := target.ID
	if target.Label != "" {
		targetLabel = target.Label + " (" + target.ID + ")"
	}
	changeLabel := change.Kind
	return ImpactEntry{
		Kind:       impactKindRenderOutput,
		Severity:   severity,
		Fixability: fixability,
		Object: ObjectRef{
			Object: "file",
			Kind:   "render_output",
			ID:     mapping.Path,
			Label:  "render output: " + mapping.Path,
			File:   mapping.Path,
		},
		Reason:            fmt.Sprintf("%s の %s により render output '%s' が更新される可能性がある", targetLabel, changeLabel, mapping.Path),
		Via:               []string{"render_output_files"},
		Source:            mapping.Source,
		RecommendedAction: fmt.Sprintf("brewprint render を再実行し、render output '%s' を確認する", mapping.Path),
	}
}

func dedupeRenderOutputMappings(mappings []renderOutputMapping) []renderOutputMapping {
	seen := map[string]struct{}{}
	out := make([]renderOutputMapping, 0, len(mappings))
	for _, mapping := range mappings {
		if mapping.Path == "" {
			continue
		}
		if _, ok := seen[mapping.Path]; ok {
			continue
		}
		seen[mapping.Path] = struct{}{}
		out = append(out, mapping)
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	return out
}

func (s *Service) mainTaskForFile(fileID semantic.FileID) *semantic.Task {
	if qid := s.project.MainNodeByFile[fileID]; qid != "" {
		if task := s.project.TasksByQID[qid]; task != nil && task.Main {
			return task
		}
	}
	for _, node := range s.project.NodesByFile[fileID] {
		if task, ok := node.(*semantic.Task); ok && task.Main {
			return task
		}
	}
	return nil
}

func (s *Service) renderGroupForFile(fileID semantic.FileID) string {
	module := renderTopLevelModuleForFile(fileID)
	if module == "" || strings.HasPrefix(module, "_") {
		return ""
	}
	for _, group := range s.project.RenderGroups {
		for _, groupModule := range group.Modules {
			if groupModule == module {
				return group.ID
			}
		}
	}
	return module
}

func renderTopLevelModuleForFile(fileID semantic.FileID) string {
	parts := strings.Split(fileID.String(), "/")
	if len(parts) == 0 {
		return ""
	}
	if parts[0] == "actors.yaml" || parts[0] == "views" || parts[0] == "render_index.yaml" {
		return ""
	}
	return parts[0]
}

func renderFSMID(fileID semantic.FileID) string {
	path := strings.TrimSuffix(fileID.String(), ".yaml")
	path = strings.TrimSuffix(path, ".yml")
	if strings.HasSuffix(path, "/state") {
		path = strings.TrimSuffix(path, "/state")
	}
	path = strings.ReplaceAll(path, "\\", "/")
	path = strings.Trim(path, "/")
	return strings.ReplaceAll(path, "/", "-")
}

func renderCrossViewPath(kind string, ids []string, id string) string {
	if len(ids) == 1 {
		return "_cross/" + kind + ".md"
	}
	return "_cross/" + kind + "-" + renderSafePathID(id) + ".md"
}

func renderSafePathID(id string) string {
	id = strings.ReplaceAll(id, "/", "-")
	id = strings.ReplaceAll(id, "\\", "-")
	return id
}

func sortedERViewIDsForQuery(project *semantic.Project) []string {
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

func sortedAPIViewIDsForQuery(project *semantic.Project) []string {
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

func erViewIncludesModule(view *semantic.ERView, module string) bool {
	if view == nil || module == "" {
		return false
	}
	for _, viewModule := range view.Modules {
		if viewModule.Module == module {
			return true
		}
	}
	return false
}

func apiViewIncludesModule(view *semantic.APIView, module string) bool {
	if view == nil || module == "" {
		return false
	}
	for _, viewModule := range view.Modules {
		if viewModule.Module == module {
			return true
		}
		if viewModule.IncludeSubmodules && strings.HasPrefix(module, viewModule.Module+".") {
			return true
		}
	}
	return false
}
