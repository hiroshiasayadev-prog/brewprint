package query

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

const analyzeImpactSourceLocationUnavailable = "source_location_unavailable"

type fieldImpactCandidate struct {
	Object ObjectRef
	Via    []string
	File   string
}

func (s *Service) collectFieldAnalyzeImpacts(req AnalyzeImpactRequest, target ObjectRef) ([]ImpactEntry, []semantic.Diagnostic) {
	switch req.Change.Kind {
	case AnalyzeImpactChangeRename, AnalyzeImpactChangeRemove, AnalyzeImpactChangeType:
		// supported below
	default:
		return []ImpactEntry{}, []semantic.Diagnostic{}
	}

	targetKey, _, err := s.referenceTarget(req.Selector)
	if err != nil {
		return []ImpactEntry{}, []semantic.Diagnostic{}
	}
	model, field, err := s.modelFieldBySelector(req.Selector)
	if err != nil {
		return []ImpactEntry{}, []semantic.Diagnostic{}
	}

	oldTypeID, oldTypeKind := s.fieldTypeIdentity(model, field)
	newTypeID, newTypeKind := "", ""
	if req.Change.Kind == AnalyzeImpactChangeType {
		newTypeID, newTypeKind = s.normalizeFieldTypeIdentity(model, req.Change.NewType)
		if oldTypeID == newTypeID && oldTypeKind == newTypeKind {
			return []ImpactEntry{}, []semantic.Diagnostic{}
		}
	}

	candidates := s.fieldImpactCandidates(req, targetKey)
	candidates = filterFieldImpactCandidatesByScope(candidates, req.ScopeModules)

	impacts := make([]ImpactEntry, 0, len(candidates))
	diagnostics := []semantic.Diagnostic{}
	diagnosticFiles := map[string]struct{}{}
	for _, candidate := range candidates {
		if candidate.Object.ID == "" || candidate.Object.ID == target.ID {
			continue
		}
		impact := fieldImpactForChange(req.Change, target, candidate, oldTypeID, oldTypeKind, newTypeID, newTypeKind)
		if impact.Object.ID == "" {
			continue
		}
		impacts = append(impacts, impact)
		if impact.Source != nil && impact.Source.File != "" && (impact.Source.Line == 0 || impact.Source.Column == 0) {
			if _, seen := diagnosticFiles[impact.Source.File]; !seen {
				diagnostics = append(diagnostics, semantic.Diagnostic{
					Severity: semantic.SeverityWarning,
					Code:     analyzeImpactSourceLocationUnavailable,
					FileID:   semantic.FileID(impact.Source.File),
					Message:  "source location is unavailable for analyze_impact impact; only source file is known",
				})
				diagnosticFiles[impact.Source.File] = struct{}{}
			}
		}
	}

	sort.Slice(impacts, func(i, j int) bool {
		if impacts[i].Object.ID != impacts[j].Object.ID {
			return impacts[i].Object.ID < impacts[j].Object.ID
		}
		return strings.Join(impacts[i].Via, ",") < strings.Join(impacts[j].Via, ",")
	})
	sort.Slice(diagnostics, func(i, j int) bool { return diagnostics[i].FileID < diagnostics[j].FileID })
	return impacts, diagnostics
}

func (s *Service) fieldImpactCandidates(req AnalyzeImpactRequest, targetKey semantic.ObjectKey) []fieldImpactCandidate {
	candidates := []fieldImpactCandidate{}
	for _, ref := range s.project.ReferencesByTarget[targetKey] {
		candidates = append(candidates, fieldImpactCandidateFromEndpoint(ref.From, []string{string(ref.Kind)}))
	}

	tree, err := s.GetReferenceTree(GetReferenceTreeRequest{
		Selector:  req.Selector,
		Direction: string(semantic.ReferenceDirectionIn),
		Depth:     2,
	})
	if err == nil {
		for _, node := range tree.Nodes {
			if node.Depth == 0 || len(node.Via) == 0 {
				continue
			}
			candidates = append(candidates, fieldImpactCandidate{
				Object: node.Object,
				Via:    append([]string{}, node.Via...),
				File:   sourceFileForObjectRef(node.Object),
			})
		}
	}

	return dedupeFieldImpactCandidates(candidates)
}

func fieldImpactCandidateFromEndpoint(endpoint semantic.ReferenceEndpoint, via []string) fieldImpactCandidate {
	object := objectRefFromReferenceEndpoint(endpoint)
	return fieldImpactCandidate{
		Object: object,
		Via:    append([]string{}, via...),
		File:   sourceFileForObjectRef(object),
	}
}

func dedupeFieldImpactCandidates(candidates []fieldImpactCandidate) []fieldImpactCandidate {
	seen := map[string]struct{}{}
	out := make([]fieldImpactCandidate, 0, len(candidates))
	for _, candidate := range candidates {
		key := candidate.Object.Object + "\x00" + candidate.Object.Kind + "\x00" + candidate.Object.ID + "\x00" + strings.Join(candidate.Via, ",")
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, candidate)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Object.ID != out[j].Object.ID {
			return out[i].Object.ID < out[j].Object.ID
		}
		return strings.Join(out[i].Via, ",") < strings.Join(out[j].Via, ",")
	})
	return out
}

func filterFieldImpactCandidatesByScope(candidates []fieldImpactCandidate, scopeModules []string) []fieldImpactCandidate {
	if len(scopeModules) == 0 {
		return candidates
	}
	scope := map[string]struct{}{}
	for _, module := range scopeModules {
		if module != "" {
			scope[module] = struct{}{}
		}
	}
	if len(scope) == 0 {
		return candidates
	}
	out := make([]fieldImpactCandidate, 0, len(candidates))
	for _, candidate := range candidates {
		module := moduleForObjectRef(candidate.Object)
		if _, ok := scope[module]; ok {
			out = append(out, candidate)
		}
	}
	return out
}

func fieldImpactForChange(change AnalyzeImpactChange, target ObjectRef, candidate fieldImpactCandidate, oldTypeID, oldTypeKind, newTypeID, newTypeKind string) ImpactEntry {
	impact := ImpactEntry{
		Kind:              "field_consumer",
		Object:            candidate.Object,
		Via:               append([]string{}, candidate.Via...),
		Source:            sourceLocationForCandidate(candidate),
		SuggestedFixes:    []SuggestedFix{},
		RecommendedAction: "参照している field を見直す",
	}

	subject := target.ID
	if target.Label != "" {
		subject = target.Label
	}
	via := strings.Join(candidate.Via, " -> ")
	if via == "" {
		via = "reference"
	}

	switch change.Kind {
	case AnalyzeImpactChangeRename:
		impact.Severity = "breaking"
		impact.Fixability = "unknown"
		impact.Reason = fmt.Sprintf("%s が field %q を %s で参照しているため、rename 後に参照解決できなくなる可能性がある", candidate.Object.ID, subject, via)
		impact.RecommendedAction = fmt.Sprintf("field reference を %s へ更新する", change.NewID)
	case AnalyzeImpactChangeRemove:
		impact.Severity = "breaking"
		impact.Fixability = "manual_review"
		impact.Reason = fmt.Sprintf("%s が field %q を %s で参照しているため、remove 後に参照解決できなくなる", candidate.Object.ID, subject, via)
		impact.RecommendedAction = "field への参照を削除するか、代替 field / model 関係へ設計を更新する"
	case AnalyzeImpactChangeType:
		impact.Severity = fieldChangeTypeSeverity(oldTypeKind, newTypeKind)
		impact.Fixability = "manual_review"
		impact.Reason = fmt.Sprintf("field %q の type が %s から %s に変わるため、%s の参照元で型整合性確認が必要", subject, oldTypeID, newTypeID, candidate.Object.ID)
		impact.RecommendedAction = "参照元の field / FK / consumer が新しい type identity と整合するか確認する"
	}

	return impact
}

func sourceLocationForCandidate(candidate fieldImpactCandidate) *SourceLocation {
	file := candidate.File
	if file == "" {
		file = sourceFileForObjectRef(candidate.Object)
	}
	if file == "" {
		return nil
	}
	return &SourceLocation{File: file}
}

func sourceFileForObjectRef(ref ObjectRef) string {
	if ref.File != "" {
		return ref.File
	}
	if ref.Source != nil {
		return ref.Source["file"]
	}
	return ""
}

func (s *Service) fieldTypeIdentity(model *semantic.Model, field semantic.ModelField) (string, string) {
	if model == nil {
		return field.Type, fieldTypeIdentityKind(field.Type)
	}
	fieldKey := semantic.ModelFieldObjectKey(model.QID, field.Name)
	for _, ref := range s.project.ReferencesBySource[fieldKey] {
		if ref.Kind != semantic.ReferenceKindFieldType {
			continue
		}
		return ref.To.ID, fieldTypeIdentityKindFromEndpoint(ref.To)
	}
	return s.normalizeFieldTypeIdentity(model, field.Type)
}

func (s *Service) normalizeFieldTypeIdentity(model *semantic.Model, raw string) (string, string) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", ""
	}
	if _, ok := s.project.ModelsByQID[semantic.QualifiedID(raw)]; ok {
		return raw, "model"
	}
	if model != nil {
		module := moduleForQualifiedID(model.QID.String())
		if module != "" && !strings.Contains(raw, ".") {
			candidate := module + ".model." + raw
			if _, ok := s.project.ModelsByQID[semantic.QualifiedID(candidate)]; ok {
				return candidate, "model"
			}
		}
	}
	return raw, fieldTypeIdentityKind(raw)
}

func fieldTypeIdentityKindFromEndpoint(endpoint semantic.ReferenceEndpoint) string {
	if endpoint.Object == "primitive" || endpoint.Kind == "primitive" {
		return "primitive"
	}
	if endpoint.Kind == "model" || strings.Contains(endpoint.ID, ".model.") {
		return "model"
	}
	return fieldTypeIdentityKind(endpoint.ID)
}

func fieldTypeIdentityKind(id string) string {
	if strings.Contains(id, ".model.") {
		return "model"
	}
	if id == "" {
		return ""
	}
	return "primitive"
}

func fieldChangeTypeSeverity(oldKind, newKind string) string {
	if oldKind == "model" && newKind == "model" {
		return "warning"
	}
	return "breaking"
}
