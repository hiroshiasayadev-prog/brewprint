package query

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/semantic"
)

const impactKindFlowParamField = "flow_param_field"

type flowParamFieldHit struct {
	FileID      semantic.FileID
	Owner       ObjectRef
	TargetParam string
	SourceRaw   string
	SourceKind  semantic.FlowSourceKind
	SourceNode  semantic.QualifiedID
	SourceAsset string
	Source      *SourceLocation
	Description string
}

func (s *Service) collectFlowParamFieldAnalyzeImpacts(req AnalyzeImpactRequest, target ObjectRef, model *semantic.Model, field semantic.ModelField, oldTypeID, newTypeID string) ([]ImpactEntry, []semantic.Diagnostic) {
	if model == nil {
		return nil, nil
	}
	if req.Change.Kind != AnalyzeImpactChangeRename && req.Change.Kind != AnalyzeImpactChangeRemove && req.Change.Kind != AnalyzeImpactChangeType {
		return nil, nil
	}

	hits := s.flowParamFieldHits(model, field)
	impacts := make([]ImpactEntry, 0, len(hits))
	diagnostics := []semantic.Diagnostic{}
	diagnosticFiles := map[string]struct{}{}
	for _, hit := range hits {
		impact := flowParamFieldImpactForChange(req.Change, target, hit, oldTypeID, newTypeID)
		if !s.impactInScope(impact, target, req.ScopeModules) {
			continue
		}
		if impact.Source != nil && impact.Source.File != "" && (impact.Source.Line == 0 || impact.Source.Column == 0) {
			impact.Fixability = "unknown"
			impact.SuggestedFixes = nil
			if _, seen := diagnosticFiles[impact.Source.File]; !seen {
				diagnostics = append(diagnostics, semantic.Diagnostic{
					Severity: semantic.SeverityWarning,
					Code:     analyzeImpactSourceLocationUnavailable,
					FileID:   semantic.FileID(impact.Source.File),
					Message:  "source location is unavailable for flow param impact; only source file is known",
				})
				diagnosticFiles[impact.Source.File] = struct{}{}
			}
		}
		impacts = append(impacts, impact)
	}

	sort.SliceStable(impacts, func(i, j int) bool {
		if impacts[i].Object.ID != impacts[j].Object.ID {
			return impacts[i].Object.ID < impacts[j].Object.ID
		}
		return sourceSortKey(impacts[i].Source) < sourceSortKey(impacts[j].Source)
	})
	sort.SliceStable(diagnostics, func(i, j int) bool { return diagnostics[i].FileID < diagnostics[j].FileID })
	return impacts, diagnostics
}

func (s *Service) flowParamFieldHits(model *semantic.Model, field semantic.ModelField) []flowParamFieldHit {
	fieldName := strings.TrimSpace(field.Name)
	modelID := model.QID
	var hits []flowParamFieldHit

	fileIDs := make([]string, 0, len(s.project.FlowByFile))
	for fileID := range s.project.FlowByFile {
		fileIDs = append(fileIDs, fileID.String())
	}
	sort.Strings(fileIDs)
	for _, rawFileID := range fileIDs {
		fileID := semantic.FileID(rawFileID)
		owner := s.flowImpactOwnerObject(fileID)
		for _, entry := range s.project.FlowByFile[fileID] {
			hits = append(hits, s.flowParamFieldHitsForEntry(fileID, owner, entry, fieldName, modelID, "flow")...)
		}
	}
	return dedupeFlowParamFieldHits(hits)
}

func (s *Service) flowParamFieldHitsForEntry(fileID semantic.FileID, owner ObjectRef, entry semantic.FlowEntry, fieldName string, modelID semantic.QualifiedID, path string) []flowParamFieldHit {
	var hits []flowParamFieldHit
	switch entry.Kind {
	case semantic.FlowKindStep:
		hits = append(hits, s.flowParamFieldHitsForWirings(fileID, owner, entry.Step.Params, fieldName, modelID, path+".step")...)
	case semantic.FlowKindForeach:
		hits = append(hits, s.flowParamFieldHitsForWirings(fileID, owner, entry.Foreach.Params, fieldName, modelID, path+".foreach")...)
	case semantic.FlowKindFork:
		for branchIndex, branch := range entry.Fork.Branches {
			for stepIndex, step := range branch.Steps {
				hits = append(hits, s.flowParamFieldHitsForWirings(fileID, owner, step.Params, fieldName, modelID, fmt.Sprintf("%s.fork.branch[%d].step[%d]", path, branchIndex+1, stepIndex+1))...)
			}
		}
		hits = append(hits, s.flowParamFieldHitsForWirings(fileID, owner, entry.Fork.JoinParams, fieldName, modelID, path+".fork.join")...)
	case semantic.FlowKindBranch:
		hits = append(hits, s.flowParamFieldHitsForWirings(fileID, owner, entry.Branch.Params, fieldName, modelID, path+".branch")...)
		for caseIndex, branchCase := range entry.Branch.Cases {
			hits = append(hits, s.flowParamFieldHitsForWirings(fileID, owner, branchCase.Step.Params, fieldName, modelID, fmt.Sprintf("%s.branch.case[%d]", path, caseIndex+1))...)
		}
	}
	return hits
}

func (s *Service) flowParamFieldHitsForWirings(fileID semantic.FileID, owner ObjectRef, wirings []semantic.ParamWiring, fieldName string, modelID semantic.QualifiedID, path string) []flowParamFieldHit {
	var hits []flowParamFieldHit
	for _, wiring := range wirings {
		if !flowParamFieldWiringMatches(s.project, wiring, fieldName, modelID) {
			continue
		}
		hits = append(hits, flowParamFieldHit{
			FileID:      fileID,
			Owner:       owner,
			TargetParam: wiring.TargetParam,
			SourceRaw:   wiring.Source.Raw,
			SourceKind:  wiring.Source.Kind,
			SourceNode:  wiring.Source.Node,
			SourceAsset: wiring.Source.AssetName,
			Source:      s.sourceForFlowParam(fileID, wiring.TargetParam, wiring.Source.Raw),
			Description: path,
		})
	}
	return hits
}

func flowParamFieldWiringMatches(project *semantic.Project, wiring semantic.ParamWiring, fieldName string, modelID semantic.QualifiedID) bool {
	if fieldName != "" && (wiring.TargetParam == fieldName || wiring.Source.ParamName == fieldName || wiring.Source.AssetName == fieldName) {
		return true
	}
	if modelID == "" || project == nil {
		return false
	}
	if wiring.Source.Node != "" {
		if task := project.TasksByQID[wiring.Source.Node]; task != nil && task.Returns != nil && task.Returns.Model == modelID {
			return true
		}
		if join := project.JoinsByQID[wiring.Source.Node]; join != nil && join.Returns != nil && join.Returns.Model == modelID {
			return true
		}
	}
	return false
}

func dedupeFlowParamFieldHits(hits []flowParamFieldHit) []flowParamFieldHit {
	seen := map[string]struct{}{}
	out := make([]flowParamFieldHit, 0, len(hits))
	for _, hit := range hits {
		key := strings.Join([]string{hit.FileID.String(), hit.Owner.ID, hit.TargetParam, hit.SourceRaw, hit.Description}, "\x00")
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, hit)
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Owner.ID != out[j].Owner.ID {
			return out[i].Owner.ID < out[j].Owner.ID
		}
		if out[i].TargetParam != out[j].TargetParam {
			return out[i].TargetParam < out[j].TargetParam
		}
		return sourceSortKey(out[i].Source) < sourceSortKey(out[j].Source)
	})
	return out
}

func flowParamFieldImpactForChange(change AnalyzeImpactChange, target ObjectRef, hit flowParamFieldHit, oldTypeID string, newTypeID string) ImpactEntry {
	subject := target.ID
	if target.Label != "" {
		subject = target.Label
	}
	impact := ImpactEntry{
		Kind:              impactKindFlowParamField,
		Object:            hit.Owner,
		Via:               []string{"flow_param_field_resolution"},
		Source:            hit.Source,
		Fixability:        "manual_review",
		RecommendedAction: "flow params の source/target と field 変更の整合を確認する",
	}
	where := hit.Description
	if where == "" {
		where = "flow params"
	}
	wire := hit.TargetParam
	if hit.SourceRaw != "" {
		wire += " <- " + hit.SourceRaw
	}
	switch change.Kind {
	case AnalyzeImpactChangeRename:
		impact.Severity = "breaking"
		impact.Reason = fmt.Sprintf("%s の %s wiring %q が field %q に関係するため、rename 後に flow param 解決が変わる可能性がある", hit.Owner.ID, where, wire, subject)
		impact.RecommendedAction = fmt.Sprintf("flow param wiring を %s へ更新する必要があるか確認する", change.NewID)
	case AnalyzeImpactChangeRemove:
		impact.Severity = "breaking"
		impact.Reason = fmt.Sprintf("%s の %s wiring %q が field %q に関係するため、remove 後に flow param 解決できなくなる可能性がある", hit.Owner.ID, where, wire, subject)
		impact.RecommendedAction = "flow param wiring を削除するか、代替 field / return value に更新する"
	case AnalyzeImpactChangeType:
		impact.Severity = "warning"
		if oldTypeID != "" && newTypeID != "" && oldTypeID != newTypeID {
			impact.Severity = "breaking"
		}
		impact.Reason = fmt.Sprintf("%s の %s wiring %q が field %q に関係するため、type identity が %s から %s に変わる影響を確認する必要がある", hit.Owner.ID, where, wire, subject, oldTypeID, newTypeID)
		impact.RecommendedAction = "flow param の source/target model identity が新しい field type と整合するか確認する"
	}
	return impact
}

func (s *Service) sourceForFlowParam(fileID semantic.FileID, targetParam string, sourceRaw string) *SourceLocation {
	content, ok := s.fileContent(fileID)
	if !ok {
		return sourceLocationFromBlock(fileID, sourceBlock{})
	}
	lines := splitSourceLines(content)
	start, end := topLevelSectionRange(lines, "flow")
	if start < 0 {
		return sourceLocationFromBlock(fileID, sourceBlock{})
	}
	for i := start; i < end; i++ {
		key, value, ok := yamlKeyValue(lines[i])
		if !ok {
			continue
		}
		if key == targetParam && (sourceRaw == "" || value == sourceRaw) {
			column := indentOf(lines[i]) + 1
			return &SourceLocation{File: fileID.String(), Line: i + 1, Column: column, EndLine: i + 1, EndColumn: column + len(strings.TrimSpace(lines[i]))}
		}
	}
	return sourceLocationFromBlock(fileID, sourceBlock{})
}

func (s *Service) flowImpactOwnerObject(fileID semantic.FileID) ObjectRef {
	if mainQID := s.project.MainNodeByFile[fileID]; mainQID != "" {
		if node := s.project.NodesByQID[mainQID]; node != nil {
			return objectRef(node)
		}
	}
	return fileObjectRef(fileID, s.fileKind(fileID))
}
