package query

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

const defaultAnalyzeImpactMaxImpacts = 200

func (s *Service) AnalyzeImpact(req AnalyzeImpactRequest) (AnalyzeImpactResponse, error) {
	if err := validateAnalyzeImpactChange(req.Change); err != nil {
		return AnalyzeImpactResponse{}, err
	}

	maxImpacts := req.MaxImpacts
	if maxImpacts <= 0 {
		maxImpacts = defaultAnalyzeImpactMaxImpacts
	}

	coverage := analyzeImpactCoverage(req.Change, req.Selector)
	assumptions := defaultAnalyzeImpactAssumptions()

	if unsupportedAnalyzeImpactSelector(req.Selector) {
		target := unsupportedAnalyzeImpactTarget(req.Selector)
		diagnostics := []semantic.Diagnostic{unsupportedAnalyzeImpactDiagnostic(req.Selector)}
		return AnalyzeImpactResponse{
			Target:           target,
			Change:           req.Change,
			Summary:          summarizeImpacts(nil),
			Impacts:          []ImpactEntry{},
			Coverage:         coverage,
			Assumptions:      assumptions,
			Truncated:        false,
			TruncatedReasons: []string{},
			Diagnostics:      diagnostics,
		}, nil
	}

	targetKey, target, err := s.referenceTarget(req.Selector)
	if err != nil {
		return AnalyzeImpactResponse{}, err
	}

	impacts, diagnostics := s.collectAnalyzeImpacts(req, targetKey, target)
	impacts, truncated, reasons := truncateImpacts(impacts, maxImpacts)
	assignImpactIDs(impacts)

	return AnalyzeImpactResponse{
		Target:           target,
		Change:           req.Change,
		Summary:          summarizeImpacts(impacts),
		Impacts:          impacts,
		Coverage:         coverage,
		Assumptions:      assumptions,
		Truncated:        truncated,
		TruncatedReasons: reasons,
		Diagnostics:      diagnostics,
	}, nil
}

func (s *Service) collectAnalyzeImpacts(req AnalyzeImpactRequest, targetKey semantic.ObjectKey, target ObjectRef) ([]ImpactEntry, []semantic.Diagnostic) {
	if req.Change.Kind == AnalyzeImpactChangeAdd {
		return s.collectAddAnalyzeImpacts(req, target)
	}

	var impacts []ImpactEntry
	var diagnostics []semantic.Diagnostic

	if taskImpacts, taskDiagnostics := s.collectTaskAnalyzeImpacts(req, targetKey, target); taskImpacts != nil || taskDiagnostics != nil {
		impacts = append(impacts, taskImpacts...)
		diagnostics = append(diagnostics, taskDiagnostics...)
	} else if target.Object == "field" || req.Selector.Object == "field" || req.Selector.Kind == "field" {
		fieldImpacts, fieldDiagnostics := s.collectFieldAnalyzeImpacts(req, target)
		impacts = append(impacts, fieldImpacts...)
		diagnostics = append(diagnostics, fieldDiagnostics...)
	} else if target.Object == "transition" && (req.Change.Kind == AnalyzeImpactChangeTransitionTarget || req.Change.Kind == AnalyzeImpactChangeRename || req.Change.Kind == AnalyzeImpactChangeRemove) {
		transitionImpacts, transitionDiagnostics := s.collectTransitionAnalyzeImpacts(req, target)
		impacts = append(impacts, transitionImpacts...)
		diagnostics = append(diagnostics, transitionDiagnostics...)
	}

	renderImpacts, renderDiagnostics := s.collectRenderAnalyzeImpacts(req, targetKey, target)
	impacts = append(impacts, renderImpacts...)
	diagnostics = append(diagnostics, renderDiagnostics...)

	sort.SliceStable(impacts, func(i, j int) bool {
		if impacts[i].Kind != impacts[j].Kind {
			return impacts[i].Kind < impacts[j].Kind
		}
		if impacts[i].Object.ID != impacts[j].Object.ID {
			return impacts[i].Object.ID < impacts[j].Object.ID
		}
		return sourceSortKey(impacts[i].Source) < sourceSortKey(impacts[j].Source)
	})
	return impacts, diagnostics
}

func (s *Service) collectTransitionAnalyzeImpacts(req AnalyzeImpactRequest, target ObjectRef) ([]ImpactEntry, []semantic.Diagnostic) {
	transition, err := s.transitionBySelector(req.Selector)
	if err != nil {
		return []ImpactEntry{}, []semantic.Diagnostic{}
	}

	var impacts []ImpactEntry
	var diagnostics []semantic.Diagnostic
	addImpact := func(impact ImpactEntry) {
		if !s.impactInScope(impact, target, req.ScopeModules) {
			return
		}
		if impact.Source == nil || impact.Source.Line == 0 || impact.Source.Column == 0 {
			diagnostics = appendSourceLocationUnavailableDiagnostic(diagnostics, impact.Source, impact.Object)
		}
		impacts = append(impacts, impact)
	}

	if impact, ok := s.collectTransitionTargetResolutionImpact(transition, req.Change); ok {
		addImpact(impact)
	}
	for _, impact := range s.collectTransitionScenarioStepImpacts(transition, req.Change) {
		addImpact(impact)
	}
	if impact, ok := s.collectTransitionActionTaskImpact(transition, req.Change); ok {
		addImpact(impact)
	}
	return impacts, diagnostics
}

func (s *Service) collectTransitionTargetResolutionImpact(transition semantic.Transition, change AnalyzeImpactChange) (ImpactEntry, bool) {
	if change.Kind != AnalyzeImpactChangeTransitionTarget {
		return ImpactEntry{}, false
	}
	transitionID := semantic.TransitionID(transition)
	var problems []string
	if change.NewTo != "" {
		if _, ok := s.project.StatesByQID[semantic.QualifiedID(change.NewTo)]; !ok {
			problems = append(problems, fmt.Sprintf("new_to %q cannot be resolved as state", change.NewTo))
		}
	}
	if change.NewAction != "" {
		if _, ok := s.project.TasksByQID[semantic.QualifiedID(change.NewAction)]; !ok {
			problems = append(problems, fmt.Sprintf("new_action %q cannot be resolved as task", change.NewAction))
		}
	}
	if len(problems) == 0 {
		return ImpactEntry{}, false
	}
	return ImpactEntry{
		Kind:       "transition_target_resolution",
		Severity:   "breaking",
		Fixability: "manual_review",
		Object:     transitionObjectRef(transition),
		Reason: fmt.Sprintf(
			"transition '%s' の変更先解決に失敗した: %s",
			transitionID,
			strings.Join(problems, "; "),
		),
		Via:               []string{"transition_target_resolution"},
		Source:            sourceLocationFromBlock(transition.FileID, s.findTransitionSource(transition)),
		RecommendedAction: "new_to / new_action が既存 state / task を指すように修正する",
	}, true
}

func (s *Service) collectTransitionScenarioStepImpacts(transition semantic.Transition, change AnalyzeImpactChange) []ImpactEntry {
	transitionID := semantic.TransitionID(transition)
	severity := transitionImpactSeverity(change)
	reason := fmt.Sprintf("sequence scenario step が transition '%s' を exact match で参照しているため、変更後にシナリオ上の意味が変わる可能性がある", transitionID)
	recommended := "scenario step の from_state / via / guard と遷移先の意味を確認する"
	if change.Kind == AnalyzeImpactChangeRemove || change.Kind == AnalyzeImpactChangeRename {
		reason = fmt.Sprintf("sequence scenario step が transition '%s' を exact match で参照しているため、変更後に参照解決できなくなる可能性が高い", transitionID)
		recommended = "scenario step の transition 参照を確認し、必要なら from_state / via / guard を更新する"
	}

	var scenarioIDs []string
	for id := range s.project.ScenariosByID {
		scenarioIDs = append(scenarioIDs, id)
	}
	sort.Strings(scenarioIDs)

	impacts := []ImpactEntry{}
	for _, scenarioID := range scenarioIDs {
		scenario := s.project.ScenariosByID[scenarioID]
		if scenario == nil {
			continue
		}
		for i, step := range scenario.Steps {
			if semantic.TransitionID(step.Transition) != transitionID {
				continue
			}
			source := s.sourceLocationForScenarioStep(scenario, i)
			impacts = append(impacts, ImpactEntry{
				Kind:              "transition_scenario_step",
				Severity:          severity,
				Fixability:        "manual_review",
				Object:            scenarioObjectRef(scenario),
				Reason:            reason,
				Via:               []string{"scenario_step_transition"},
				Source:            source,
				RecommendedAction: recommended,
			})
		}
	}
	return impacts
}

func (s *Service) collectTransitionActionTaskImpact(transition semantic.Transition, change AnalyzeImpactChange) (ImpactEntry, bool) {
	if transition.ActionTask == "" {
		return ImpactEntry{}, false
	}
	task := s.project.TasksByQID[transition.ActionTask]
	if task == nil {
		return ImpactEntry{}, false
	}

	transitionID := semantic.TransitionID(transition)
	severity := transitionImpactSeverity(change)
	reason := fmt.Sprintf("transition '%s' の遷移先や action が変わると、action task '%s' が実行される文脈の意味が変わる可能性がある", transitionID, transition.ActionTask)
	recommended := "action task の副作用と遷移先 state の整合を人間が確認する"
	if change.Kind == AnalyzeImpactChangeRemove {
		reason = fmt.Sprintf("transition '%s' が削除されると、action task '%s' への到達経路が失われる", transitionID, transition.ActionTask)
		recommended = "action task が別の transition から必要か確認する"
	}
	if change.Kind == AnalyzeImpactChangeRename {
		reason = fmt.Sprintf("transition '%s' が rename されると、action task '%s' との紐づけの意味を確認する必要がある", transitionID, transition.ActionTask)
		recommended = "transition rename 後も action task の文脈が意図通りか確認する"
	}

	return ImpactEntry{
		Kind:              "transition_action_task",
		Severity:          severity,
		Fixability:        "manual_review",
		Object:            objectRef(task),
		Reason:            reason,
		Via:               []string{"transition_action"},
		Source:            sourceLocationFromBlock(transition.FileID, s.findTransitionSource(transition)),
		RecommendedAction: recommended,
	}, true
}

func transitionImpactSeverity(change AnalyzeImpactChange) string {
	if change.Kind == AnalyzeImpactChangeRemove || change.Kind == AnalyzeImpactChangeRename {
		return "breaking"
	}
	return "warning"
}

func (s *Service) sourceLocationForScenarioStep(scenario *semantic.SequenceScenario, stepIndex int) *SourceLocation {
	if scenario == nil {
		return nil
	}
	content, ok := s.fileContent(scenario.FileID)
	if !ok {
		return &SourceLocation{File: scenario.FileID.String()}
	}
	lines := splitSourceLines(content)
	start, end := topLevelSectionRange(lines, "steps")
	if start < 0 {
		return &SourceLocation{File: scenario.FileID.String()}
	}
	current := 0
	for i := start; i < end; i++ {
		trimmed := strings.TrimSpace(lines[i])
		if !strings.HasPrefix(trimmed, "- ") && trimmed != "-" {
			continue
		}
		itemEnd := nextSequenceItemOrSectionEnd(lines, i+1, end, indentOf(lines[i]))
		if current == stepIndex {
			return sourceLocationFromBlock(scenario.FileID, makeBlock(lines, i, itemEnd))
		}
		current++
		i = itemEnd - 1
	}
	return &SourceLocation{File: scenario.FileID.String()}
}

func sourceLocationFromBlock(fileID semantic.FileID, block sourceBlock) *SourceLocation {
	loc := &SourceLocation{File: fileID.String()}
	if block.startLine > 0 {
		loc.Line = block.startLine
		loc.Column = block.column
		loc.EndLine = block.endLine
		loc.EndColumn = 1
	}
	return loc
}

func appendSourceLocationUnavailableDiagnostic(diagnostics []semantic.Diagnostic, source *SourceLocation, object ObjectRef) []semantic.Diagnostic {
	fileID := semantic.FileID(object.File)
	if source != nil && source.File != "" {
		fileID = semantic.FileID(source.File)
	}
	messageTarget := object.ID
	if messageTarget == "" {
		messageTarget = object.File
	}
	return append(diagnostics, semantic.Diagnostic{
		Severity: semantic.SeverityWarning,
		Code:     "source_location_unavailable",
		FileID:   fileID,
		Message:  fmt.Sprintf("source line/column is unavailable for impact object: %s", messageTarget),
	})
}

func sourceSortKey(source *SourceLocation) string {
	if source == nil {
		return ""
	}
	return fmt.Sprintf("%s:%08d:%08d", source.File, source.Line, source.Column)
}

func (s *Service) impactInScope(impact ImpactEntry, target ObjectRef, scopeModules []string) bool {
	if len(scopeModules) == 0 {
		return true
	}
	if objectMatchesScope(impact.Object, scopeModules) {
		return true
	}
	return objectMatchesScope(target, scopeModules)
}

func objectMatchesScope(object ObjectRef, scopeModules []string) bool {
	for _, module := range scopeModules {
		if module == "" {
			continue
		}
		if object.Module == module {
			return true
		}
		if object.File != "" && (object.File == module+".yaml" || strings.HasPrefix(object.File, module+"/")) {
			return true
		}
		if hasQualifiedIDModule(object.ID, module) || hasQualifiedIDModule(object.QualifiedID, module) {
			return true
		}
	}
	return false
}

func hasQualifiedIDModule(id string, module string) bool {
	if id == "" || module == "" {
		return false
	}
	return id == module || strings.HasPrefix(id, module+".")
}

func analyzeImpactFieldSelector(selector Selector) bool {
	return selector.Object == "field" || selector.Kind == "field"
}

func unsupportedAnalyzeImpactSelector(selector Selector) bool {
	if selector.Object == "primitive" || selector.Kind == "primitive" {
		return true
	}
	if selector.Object == "asset" || selector.Kind == "asset" {
		return true
	}
	if selector.Object == "file" || selector.Kind == "file" || selector.Kind == "state_file" {
		return true
	}
	if selector.Object == "view" {
		return true
	}
	if selector.Kind == "api_table" || selector.Kind == "er_diagram" || selector.Kind == "sequence_diagram" || selector.Kind == "render_index" {
		return true
	}
	return false
}

func unsupportedAnalyzeImpactTarget(selector Selector) ObjectRef {
	object := selector.Object
	if object == "" {
		object = "unsupported"
	}
	kind := selector.Kind
	if kind == "" {
		kind = object
	}
	id := selector.ID
	if id == "" {
		id = selector.File
	}
	return ObjectRef{
		Object:  object,
		Kind:    kind,
		ID:      id,
		File:    selector.File,
		LocalID: selector.LocalID,
	}
}

func unsupportedAnalyzeImpactDiagnostic(selector Selector) semantic.Diagnostic {
	id := selector.ID
	if id == "" {
		id = selector.File
	}
	return semantic.Diagnostic{
		Severity: semantic.SeverityWarning,
		Code:     "unsupported_selector",
		Message:  fmt.Sprintf("unsupported selector for analyze_impact: object=%s kind=%s id=%s", selector.Object, selector.Kind, id),
	}
}

func analyzeImpactCoverage(change AnalyzeImpactChange, selector Selector) ImpactCoverage {
	coverage := ImpactCoverage{
		Analyzed: []string{
			"direct_references",
			"reference_tree",
		},
		NotAnalyzed: []string{
			"type_structural_compatibility",
			"semantic_contract_compatibility",
			"render_presentation_details",
			"wireframe_element_binding",
		},
		Note: "v1では reference 経路の到達可能性と型 signature identity のみを判定する。",
	}

	switch change.Kind {
	case AnalyzeImpactChangeRename:
		coverage.Analyzed = append(coverage.Analyzed, "model_field_resolution", "transition_action_resolution", "render_output_files")
		if analyzeImpactFieldSelector(selector) {
			coverage.Analyzed = append(coverage.Analyzed, "flow_param_field_resolution")
		}
	case AnalyzeImpactChangeRemove:
		coverage.Analyzed = append(coverage.Analyzed, "transition_action_resolution", "flow_step_task_resolution", "sequence_step_task_resolution", "render_output_files")
		if analyzeImpactFieldSelector(selector) {
			coverage.Analyzed = append(coverage.Analyzed, "model_field_resolution")
		}
	case AnalyzeImpactChangeType:
		coverage.Analyzed = append(coverage.Analyzed, "model_field_resolution", "flow_param_field_resolution", "type_signature_identity", "render_output_files")
	case AnalyzeImpactChangeContract:
		coverage.Analyzed = append(coverage.Analyzed, "flow_step_task_resolution", "flow_param_field_resolution", "sequence_step_task_resolution", "render_output_files")
	case AnalyzeImpactChangeTransitionTarget:
		coverage.Analyzed = append(coverage.Analyzed, "transition_action_resolution", "sequence_step_task_resolution", "render_output_files")
	case AnalyzeImpactChangeAdd:
		coverage.Analyzed = []string{"name_collision"}
		coverage.NotAnalyzed = append(coverage.NotAnalyzed, "type_resolution", "writer_coverage")
	}

	if unsupportedAnalyzeImpactSelector(selector) {
		coverage.Analyzed = []string{}
		coverage.NotAnalyzed = append([]string{"unsupported_selector"}, coverage.NotAnalyzed...)
		coverage.Note = "selector is not supported by analyze_impact v1."
	}

	return coverage
}

func defaultAnalyzeImpactAssumptions() []string {
	return []string{
		"rename後のID衝突は検証対象外",
		"note内の自然言語参照は解析対象外",
		"semantic contract compatibility は解析対象外",
	}
}

func summarizeImpacts(impacts []ImpactEntry) ImpactSummary {
	summary := ImpactSummary{
		BySeverity:   map[string]int{"breaking": 0, "warning": 0, "info": 0},
		ByFixability: map[string]int{"mechanical": 0, "suggested": 0, "manual_review": 0, "unknown": 0},
		ByKind:       map[string]int{},
	}
	for _, impact := range impacts {
		if impact.Severity != "" {
			summary.BySeverity[impact.Severity]++
		}
		if impact.Fixability != "" {
			summary.ByFixability[impact.Fixability]++
		}
		if impact.Kind != "" {
			summary.ByKind[impact.Kind]++
		}
	}
	return summary
}

func truncateImpacts(impacts []ImpactEntry, maxImpacts int) ([]ImpactEntry, bool, []string) {
	if maxImpacts <= 0 {
		maxImpacts = defaultAnalyzeImpactMaxImpacts
	}
	if len(impacts) <= maxImpacts {
		return impacts, false, []string{}
	}
	return impacts[:maxImpacts], true, []string{"max_impacts"}
}

func assignImpactIDs(impacts []ImpactEntry) {
	for i := range impacts {
		if impacts[i].ID != "" {
			continue
		}
		impacts[i].ID = fmt.Sprintf("impact-%03d", i+1)
	}
}
