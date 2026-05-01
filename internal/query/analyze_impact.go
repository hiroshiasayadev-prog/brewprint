package query

import (
	"fmt"

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

	_, target, err := s.referenceTarget(req.Selector)
	if err != nil {
		return AnalyzeImpactResponse{}, err
	}

	impacts, diagnostics := s.collectAnalyzeImpacts(req, target)
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

func (s *Service) collectAnalyzeImpacts(req AnalyzeImpactRequest, target ObjectRef) ([]ImpactEntry, []semantic.Diagnostic) {
	return []ImpactEntry{}, []semantic.Diagnostic{}
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
	case AnalyzeImpactChangeRemove:
		coverage.Analyzed = append(coverage.Analyzed, "transition_action_resolution", "flow_step_task_resolution", "sequence_step_task_resolution", "render_output_files")
	case AnalyzeImpactChangeType:
		coverage.Analyzed = append(coverage.Analyzed, "model_field_resolution", "flow_param_field_resolution", "type_signature_identity")
	case AnalyzeImpactChangeContract:
		coverage.Analyzed = append(coverage.Analyzed, "flow_step_task_resolution", "flow_param_field_resolution", "sequence_step_task_resolution")
	case AnalyzeImpactChangeTransitionTarget:
		coverage.Analyzed = append(coverage.Analyzed, "transition_action_resolution")
	case AnalyzeImpactChangeAdd:
		coverage.Analyzed = []string{"name_collision", "type_resolution", "writer_coverage"}
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
