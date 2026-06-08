package query

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

const impactKindNameCollision = "name_collision"

func (s *Service) collectAddAnalyzeImpacts(req AnalyzeImpactRequest, target ObjectRef) ([]ImpactEntry, []semantic.Diagnostic) {
	addedID := strings.TrimSpace(req.Change.AddedID)
	collisions := s.addNameCollisionObjects(addedID)

	impacts := make([]ImpactEntry, 0, len(collisions))
	for _, collision := range collisions {
		impact := ImpactEntry{
			Kind:              impactKindNameCollision,
			Severity:          "breaking",
			Fixability:        "manual_review",
			Object:            collision,
			Reason:            fmt.Sprintf("added_id %q は既存の %s:%s %q と衝突するため、このまま add すると名前解決が一意にならない", addedID, collision.Object, collision.Kind, collision.ID),
			Via:               []string{impactKindNameCollision},
			Source:            sourceLocationForAddCollision(collision),
			RecommendedAction: "added_id を変更するか、既存 object との統合・置き換えが意図通りか確認する",
		}
		if s.impactInScope(impact, target, req.ScopeModules) {
			impacts = append(impacts, impact)
		}
	}

	diagnostics := []semantic.Diagnostic{}
	if len(collisions) == 0 && !s.addedIDKindKnown(addedID) {
		diagnostics = append(diagnostics, addUnsupportedTargetDiagnostic(addedID, target))
	}

	sort.SliceStable(impacts, func(i, j int) bool {
		if impacts[i].Object.Object != impacts[j].Object.Object {
			return impacts[i].Object.Object < impacts[j].Object.Object
		}
		if impacts[i].Object.Kind != impacts[j].Object.Kind {
			return impacts[i].Object.Kind < impacts[j].Object.Kind
		}
		return impacts[i].Object.ID < impacts[j].Object.ID
	})
	return impacts, diagnostics
}

func (s *Service) addNameCollisionObjects(addedID string) []ObjectRef {
	if addedID == "" || s == nil || s.project == nil {
		return nil
	}

	var collisions []ObjectRef
	if node := s.project.NodesByQID[semantic.QualifiedID(addedID)]; node != nil {
		collisions = append(collisions, objectRef(node))
	}
	for _, model := range s.project.ModelsByQID {
		if model == nil {
			continue
		}
		for _, field := range model.Fields {
			if model.QID.String()+"."+field.Name == addedID {
				collisions = append(collisions, fieldObjectRef(model, field))
			}
		}
	}
	for _, transition := range s.allTransitions() {
		if semantic.TransitionID(transition) == addedID {
			collisions = append(collisions, transitionObjectRef(transition))
		}
	}
	for _, asset := range s.allAssets() {
		if asset == nil {
			continue
		}
		if semantic.AssetID(asset.ProducedBy, asset.Name) == addedID {
			collisions = append(collisions, assetObjectRef(asset))
		}
	}

	return dedupeAndSortObjectRefs(collisions)
}

func (s *Service) allTransitions() []semantic.Transition {
	if s == nil || s.project == nil {
		return nil
	}
	var transitions []semantic.Transition
	for _, byFile := range s.project.TransitionsByFile {
		transitions = append(transitions, byFile...)
	}
	sort.SliceStable(transitions, func(i, j int) bool {
		return semantic.TransitionID(transitions[i]) < semantic.TransitionID(transitions[j])
	})
	return transitions
}

func (s *Service) allAssets() []*semantic.Asset {
	if s == nil || s.project == nil {
		return nil
	}
	var assets []*semantic.Asset
	for _, task := range s.project.TasksByQID {
		if task != nil && task.Returns != nil && task.Returns.Asset != nil {
			assets = append(assets, task.Returns.Asset)
		}
	}
	for _, join := range s.project.JoinsByQID {
		if join != nil && join.Returns != nil && join.Returns.Asset != nil {
			assets = append(assets, join.Returns.Asset)
		}
	}
	sort.SliceStable(assets, func(i, j int) bool {
		return semantic.AssetID(assets[i].ProducedBy, assets[i].Name) < semantic.AssetID(assets[j].ProducedBy, assets[j].Name)
	})
	return assets
}

func dedupeAndSortObjectRefs(refs []ObjectRef) []ObjectRef {
	seen := map[string]struct{}{}
	out := make([]ObjectRef, 0, len(refs))
	for _, ref := range refs {
		key := ref.Object + "\x00" + ref.Kind + "\x00" + ref.ID
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, ref)
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Object != out[j].Object {
			return out[i].Object < out[j].Object
		}
		if out[i].Kind != out[j].Kind {
			return out[i].Kind < out[j].Kind
		}
		return out[i].ID < out[j].ID
	})
	return out
}

func sourceLocationForAddCollision(ref ObjectRef) *SourceLocation {
	file := ref.File
	if file == "" && ref.Source != nil {
		file = ref.Source["file"]
	}
	if file == "" {
		return nil
	}
	return &SourceLocation{File: file}
}

func (s *Service) addedIDKindKnown(addedID string) bool {
	if addedID == "" {
		return false
	}
	if len(s.addNameCollisionObjects(addedID)) > 0 {
		return true
	}
	if s.addedIDLooksLikeModelField(addedID) {
		return true
	}
	if strings.Contains(addedID, "#") {
		left, right, ok := splitSyntheticID(addedID)
		return ok && left != "" && right != ""
	}
	parts := strings.Split(addedID, ".")
	for i, part := range parts {
		if isAddNodeKindSegment(part) && i > 0 && i < len(parts)-1 {
			return true
		}
	}
	return false
}

func (s *Service) addedIDLooksLikeModelField(addedID string) bool {
	if s == nil || s.project == nil {
		return false
	}
	for _, model := range s.project.ModelsByQID {
		if model == nil {
			continue
		}
		prefix := model.QID.String() + "."
		if strings.HasPrefix(addedID, prefix) {
			fieldName := strings.TrimPrefix(addedID, prefix)
			return fieldName != "" && !strings.Contains(fieldName, ".")
		}
	}
	return false
}

func isAddNodeKindSegment(part string) bool {
	if part == "actor" {
		return true
	}
	return isKindSegment(part)
}

func addUnsupportedTargetDiagnostic(addedID string, target ObjectRef) semantic.Diagnostic {
	diagnostic := semantic.Diagnostic{
		Severity: semantic.SeverityWarning,
		Code:     "unsupported_selector",
		Message:  fmt.Sprintf("added_id kind cannot be inferred by analyze_impact add collector: %s", addedID),
	}
	if target.File != "" {
		diagnostic.FileID = semantic.FileID(target.File)
	}
	return diagnostic
}
