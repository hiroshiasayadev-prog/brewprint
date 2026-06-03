package designrecords

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type validationScope struct {
	kind    RecordKind
	hasKind bool
	idRange *recordIDRange
}

func newValidationScope(req ValidateRecordsRequest) (validationScope, error) {
	scope := validationScope{}
	if req.Kind != "" {
		if !isListableRecordKind(req.Kind) {
			return scope, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unsupported kind %q", req.Kind))
		}
		scope.kind = req.Kind
		scope.hasKind = true
	}
	if req.IDRange == nil {
		return scope, nil
	}
	parsed, err := parseRecordIDRange(*req.IDRange, req.Kind)
	if err != nil {
		return scope, err
	}
	scope.kind = parsed.kind
	scope.hasKind = true
	scope.idRange = parsed
	return scope, nil
}

func generateValidationDiagnostics(idx *Index, scope validationScope) []Diagnostic {
	var diagnostics []Diagnostic

	recordsByPath := make(map[string]Record, len(idx.Records))
	recordsByID := make(map[string][]Record, len(idx.Records))
	for _, record := range idx.Records {
		recordsByPath[record.Path] = record
		if record.NormalizedID != "" {
			recordsByID[record.NormalizedID] = append(recordsByID[record.NormalizedID], record)
		}
	}

	candidatesByPath := make(map[string]RecordCandidate, len(idx.Candidates))
	for _, candidate := range idx.Candidates {
		if candidate.Path != "" {
			candidatesByPath[candidate.Path] = candidate
		}
	}

	diagnostics = append(diagnostics, duplicateIDDiagnostics(recordsByID, scope)...)
	diagnostics = append(diagnostics, parseIssueDiagnostics(idx.ParseIssues, recordsByPath, candidatesByPath, scope)...)
	diagnostics = append(diagnostics, semanticRefDiagnostics(idx, scope)...)
	diagnostics = append(diagnostics, recordDiagnostics(idx.Records, recordsByID, semanticTargetsByRef(idx), scope)...)
	diagnostics = append(diagnostics, pathIssueDiagnostics(idx.PathIssues, candidatesByPath, scope)...)
	return diagnostics
}

func duplicateIDDiagnostics(recordsByID map[string][]Record, scope validationScope) []Diagnostic {
	var diagnostics []Diagnostic
	ids := make([]string, 0, len(recordsByID))
	for normalizedID := range recordsByID {
		ids = append(ids, normalizedID)
	}
	sort.Strings(ids)
	for _, normalizedID := range ids {
		records := append([]Record(nil), recordsByID[normalizedID]...)
		if len(records) < 2 {
			continue
		}
		sort.Slice(records, func(i, j int) bool {
			if records[i].Path == records[j].Path {
				return records[i].ID < records[j].ID
			}
			return records[i].Path < records[j].Path
		})
		for _, record := range records {
			if !scope.selectRecord(record) {
				continue
			}
			diagnostics = append(diagnostics, Diagnostic{
				Category: DiagnosticDuplicateID,
				Severity: DiagnosticSeverityError,
				RecordID: record.ID,
				Path:     record.Path,
				Message:  fmt.Sprintf("duplicate normalized record ID %s", normalizedID),
			})
		}
	}
	return diagnostics
}

func parseIssueDiagnostics(issues []ParseIssue, recordsByPath map[string]Record, candidatesByPath map[string]RecordCandidate, scope validationScope) []Diagnostic {
	diagnostics := make([]Diagnostic, 0, len(issues))
	for _, issue := range issues {
		if !scope.selectIssue(issue, recordsByPath, candidatesByPath) {
			continue
		}
		diagnostics = append(diagnostics, Diagnostic{
			Category: issue.Category,
			Severity: DiagnosticSeverityError,
			RecordID: issue.RecordID,
			Path:     issue.Path,
			Message:  issue.Message,
		})
	}
	return diagnostics
}

func recordDiagnostics(records []Record, recordsByID map[string][]Record, semanticByRef map[string][]SemanticRefDecl, scope validationScope) []Diagnostic {
	var diagnostics []Diagnostic
	for _, record := range records {
		if !scope.selectRecord(record) {
			continue
		}
		diagnostics = append(diagnostics, workflowMetadataDiagnostics(record)...)
		diagnostics = append(diagnostics, requiredNarrativeSectionDiagnostics(record)...)
		if !statusAllowedForKind(record.Kind, record.Status) {
			diagnostics = append(diagnostics, Diagnostic{
				Category: DiagnosticInvalidStatusForKind,
				Severity: DiagnosticSeverityError,
				RecordID: record.ID,
				Path:     record.Path,
				Message:  fmt.Sprintf("%s status %q is not valid for kind %s", record.ID, record.Status, record.Kind),
			})
		}
		for _, targetID := range recordDependsOn(record) {
			if !recordIDExists(recordsByID, targetID) {
				diagnostics = append(diagnostics, Diagnostic{
					Category: DiagnosticMissingDependsOnTarget,
					Severity: DiagnosticSeverityError,
					RecordID: record.ID,
					Path:     record.Path,
					Message:  fmt.Sprintf("depends_on references missing record %s", targetID),
					TargetID: targetID,
				})
			}
		}
		for _, targetID := range recordSupersedes(record) {
			if !recordIDExists(recordsByID, targetID) {
				diagnostics = append(diagnostics, Diagnostic{
					Category: DiagnosticMissingSupersedesTarget,
					Severity: DiagnosticSeverityError,
					RecordID: record.ID,
					Path:     record.Path,
					Message:  fmt.Sprintf("supersedes references missing record %s", targetID),
					TargetID: targetID,
				})
			}
		}
		if record.Kind == RecordKindInvestigation && record.Investigation != nil {
			diagnostics = append(diagnostics, investigationReferenceDiagnostics(record, recordsByID, semanticByRef)...)
		}
		diagnostics = append(diagnostics, workflowRelationDiagnostics(record, recordsByID)...)
	}
	return diagnostics
}

type workflowRequiredField struct {
	name   string
	scalar bool
}

func workflowMetadataDiagnostics(record Record) []Diagnostic {
	if record.WorkflowMeta == nil {
		return nil
	}
	var fields []workflowRequiredField
	switch record.Kind {
	case RecordKindRequirement:
		fields = []workflowRequiredField{
			{name: "id", scalar: true},
			{name: "status", scalar: true},
			{name: "date", scalar: true},
			{name: "source_refs"},
			{name: "work_items"},
		}
	case RecordKindWorkItem:
		fields = []workflowRequiredField{
			{name: "id", scalar: true},
			{name: "status", scalar: true},
			{name: "date", scalar: true},
			{name: "source_requirement", scalar: true},
			{name: "impact_refs"},
			{name: "tasks"},
		}
	case RecordKindTask:
		fields = []workflowRequiredField{
			{name: "id", scalar: true},
			{name: "status", scalar: true},
			{name: "date", scalar: true},
			{name: "work_item", scalar: true},
			{name: "source_requirement", scalar: true},
			{name: "estimate", scalar: true},
			{name: "depends_on"},
			{name: "outputs"},
		}
	default:
		return nil
	}

	var diagnostics []Diagnostic
	for _, required := range fields {
		field, ok := record.WorkflowMeta.Fields[required.name]
		if !ok || !field.Present {
			diagnostics = append(diagnostics, workflowMetadataDiagnostic(record, DiagnosticMissingRequiredMetadata, required.name, "", false, fmt.Sprintf("%s is missing required metadata field %s", record.ID, required.name)))
			continue
		}
		if required.scalar {
			if strings.TrimSpace(field.Value) == "" {
				diagnostics = append(diagnostics, workflowMetadataDiagnostic(record, DiagnosticEmptyRequiredMetadata, required.name, field.Value, true, fmt.Sprintf("%s.%s is empty", record.ID, required.name)))
				continue
			}
			if required.name == "date" && !validDateOnly(field.Value) {
				diagnostics = append(diagnostics, workflowMetadataDiagnostic(record, DiagnosticInvalidMetadataValue, required.name, field.Value, true, fmt.Sprintf("%s.date must use strict YYYY-MM-DD format", record.ID)))
			}
			continue
		}
		for range field.EmptyItems {
			diagnostics = append(diagnostics, workflowMetadataDiagnostic(record, DiagnosticEmptyRequiredMetadata, required.name, "", true, fmt.Sprintf("%s.%s contains an empty list item", record.ID, required.name)))
		}
	}
	return diagnostics
}

func workflowMetadataDiagnostic(record Record, category DiagnosticCategory, field, value string, valuePresent bool, message string) Diagnostic {
	return Diagnostic{
		Category:     category,
		Severity:     DiagnosticSeverityError,
		RecordID:     record.ID,
		Path:         record.Path,
		Message:      message,
		Field:        field,
		Value:        value,
		ValuePresent: valuePresent,
	}
}

func semanticRefDiagnostics(idx *Index, scope validationScope) []Diagnostic {
	var diagnostics []Diagnostic
	if scope.hasKind && scope.kind != RecordKindSpec {
		return diagnostics
	}
	byRef := map[string][]SemanticRefDecl{}
	for _, source := range idx.SemanticRefSources {
		for _, decl := range source.Decls {
			if !activeSpecRefPattern.MatchString(decl.Ref) {
				diagnostics = append(diagnostics, Diagnostic{
					Category: DiagnosticInvalidSemanticRefDeclaration,
					Severity: DiagnosticSeverityError,
					RecordID: source.RecordID,
					Path:     decl.Path,
					Message:  fmt.Sprintf("semantic reference declaration %q is invalid", decl.Ref),
				})
				continue
			}
			if decl.TargetType == SemanticTargetSection {
				matches := matchingHeadingCount(source.Headings, decl.Section)
				switch matches {
				case 0:
					diagnostics = append(diagnostics, Diagnostic{
						Category: DiagnosticMissingSectionTarget,
						Severity: DiagnosticSeverityError,
						RecordID: source.RecordID,
						Path:     decl.Path,
						Message:  fmt.Sprintf("section target %q for %s was not found", decl.Section, decl.Ref),
					})
				case 1:
				default:
					diagnostics = append(diagnostics, Diagnostic{
						Category: DiagnosticAmbiguousSectionTarget,
						Severity: DiagnosticSeverityError,
						RecordID: source.RecordID,
						Path:     decl.Path,
						Message:  fmt.Sprintf("section target %q for %s matches multiple headings", decl.Section, decl.Ref),
					})
				}
			}
			byRef[decl.Ref] = append(byRef[decl.Ref], decl)
		}
	}
	refs := make([]string, 0, len(byRef))
	for ref, decls := range byRef {
		if len(decls) > 1 {
			refs = append(refs, ref)
		}
	}
	sort.Strings(refs)
	for _, ref := range refs {
		decls := byRef[ref]
		sort.Slice(decls, func(i, j int) bool {
			if decls[i].Path == decls[j].Path {
				return decls[i].Section < decls[j].Section
			}
			return decls[i].Path < decls[j].Path
		})
		for _, decl := range decls {
			diagnostics = append(diagnostics, Diagnostic{
				Category: DiagnosticDuplicateSemanticRef,
				Severity: DiagnosticSeverityError,
				Path:     decl.Path,
				Message:  fmt.Sprintf("semantic reference %s has multiple targets", ref),
			})
		}
	}
	return diagnostics
}

func investigationReferenceDiagnostics(record Record, recordsByID map[string][]Record, semanticByRef map[string][]SemanticRefDecl) []Diagnostic {
	var diagnostics []Diagnostic
	diagnostics = append(diagnostics, diagnosticsForInvestigationRefs(record, "source_refs", record.Investigation.SourceRefs, DiagnosticUnresolvedSourceRef, DiagnosticNoncanonicalSourceRef, DiagnosticSeverityError, recordsByID, semanticByRef)...)
	diagnostics = append(diagnostics, diagnosticsForInvestigationRefs(record, "follow_up_results", record.Investigation.FollowUpResults, DiagnosticUnresolvedFollowUpResult, DiagnosticNoncanonicalFollowUpResult, DiagnosticSeverityError, recordsByID, semanticByRef)...)
	diagnostics = append(diagnostics, diagnosticsForInvestigationRefs(record, "follow_up_candidates", record.Investigation.FollowUpCandidates, DiagnosticUnresolvedFollowUpCandidate, DiagnosticNoncanonicalFollowUpCandidate, DiagnosticSeverityInfo, recordsByID, semanticByRef)...)
	return diagnostics
}

func diagnosticsForInvestigationRefs(record Record, field string, values []string, unresolvedCategory, noncanonicalCategory DiagnosticCategory, severity DiagnosticSeverity, recordsByID map[string][]Record, semanticByRef map[string][]SemanticRefDecl) []Diagnostic {
	var diagnostics []Diagnostic
	for _, value := range values {
		switch {
		case strings.HasPrefix(value, "yaml:"):
			continue
		case strings.HasPrefix(value, "TASK-"):
			diagnostics = append(diagnostics, investigationReferenceDiagnostic(record, DiagnosticUnsupportedReference, severity, field, value, "unsupported"))
		case isPhysicalPathReference(value):
			diagnostics = append(diagnostics, investigationReferenceDiagnostic(record, noncanonicalCategory, severity, field, value, "noncanonical"))
		case isExplicitUnsupportedReference(value):
			diagnostics = append(diagnostics, investigationReferenceDiagnostic(record, DiagnosticUnsupportedReference, severity, field, value, "unsupported"))
		case activeSpecRefPattern.MatchString(value):
			targets := semanticByRef[value]
			if len(targets) == 0 {
				diagnostics = append(diagnostics, investigationReferenceDiagnostic(record, unresolvedCategory, severity, field, value, "unresolved"))
			}
		case isSupportedInvestigationRecordIDReference(value):
			targets := recordsByID[normalizeRecordID(value)]
			if len(targets) == 0 {
				diag := investigationReferenceDiagnostic(record, unresolvedCategory, severity, field, value, "unresolved")
				diag.TargetID = value
				diagnostics = append(diagnostics, diag)
			}
		default:
			diagnostics = append(diagnostics, investigationReferenceDiagnostic(record, DiagnosticUnsupportedReference, severity, field, value, "unsupported"))
		}
	}
	return diagnostics
}

func investigationReferenceDiagnostic(record Record, category DiagnosticCategory, severity DiagnosticSeverity, field, value, refStatus string) Diagnostic {
	return Diagnostic{
		Category:  category,
		Severity:  severity,
		RecordID:  record.ID,
		Path:      record.Path,
		Message:   fmt.Sprintf("%s contains %s reference %q", field, refStatus, value),
		Field:     field,
		Value:     value,
		RefStatus: refStatus,
	}
}

func workflowRelationDiagnostics(record Record, recordsByID map[string][]Record) []Diagnostic {
	switch record.Kind {
	case RecordKindRequirement:
		if record.Requirement == nil {
			return nil
		}
		return requirementWorkflowRelationDiagnostics(record, recordsByID)
	case RecordKindWorkItem:
		if record.WorkItem == nil {
			return nil
		}
		return workItemWorkflowRelationDiagnostics(record, recordsByID)
	case RecordKindTask:
		if record.Task == nil {
			return nil
		}
		return taskWorkflowRelationDiagnostics(record, recordsByID)
	default:
		return nil
	}
}

func requirementWorkflowRelationDiagnostics(record Record, recordsByID map[string][]Record) []Diagnostic {
	var diagnostics []Diagnostic
	for _, value := range record.Requirement.WorkItems {
		target, ok, targetDiagnostics := validateWorkflowRelationTarget(record, "work_items", value, RecordKindWorkItem, recordsByID)
		diagnostics = append(diagnostics, targetDiagnostics...)
		if !ok {
			continue
		}
		if target.WorkItem == nil || target.WorkItem.SourceRequirement != record.ID {
			diagnostics = append(diagnostics, workflowMismatchDiagnostic(record, "work_items", value, target.ID, fmt.Sprintf("%s.work_items contains %s but %s.source_requirement is %q", record.ID, value, target.ID, workflowSourceRequirement(target))))
		}
	}
	return diagnostics
}

func workItemWorkflowRelationDiagnostics(record Record, recordsByID map[string][]Record) []Diagnostic {
	var diagnostics []Diagnostic
	if record.WorkItem.SourceRequirement != "" {
		target, ok, targetDiagnostics := validateWorkflowRelationTarget(record, "source_requirement", record.WorkItem.SourceRequirement, RecordKindRequirement, recordsByID)
		diagnostics = append(diagnostics, targetDiagnostics...)
		if ok && (target.Requirement == nil || !containsString(target.Requirement.WorkItems, record.ID)) {
			diagnostics = append(diagnostics, workflowMismatchDiagnostic(record, "source_requirement", record.WorkItem.SourceRequirement, target.ID, fmt.Sprintf("%s.source_requirement is %s but %s.work_items does not contain %s", record.ID, record.WorkItem.SourceRequirement, target.ID, record.ID)))
		}
	}
	for _, value := range record.WorkItem.Tasks {
		target, ok, targetDiagnostics := validateWorkflowRelationTarget(record, "tasks", value, RecordKindTask, recordsByID)
		diagnostics = append(diagnostics, targetDiagnostics...)
		if !ok {
			continue
		}
		if target.Task == nil || target.Task.WorkItem != record.ID {
			diagnostics = append(diagnostics, workflowMismatchDiagnostic(record, "tasks", value, target.ID, fmt.Sprintf("%s.tasks contains %s but %s.work_item is %q", record.ID, value, target.ID, workflowTaskWorkItem(target))))
		}
	}
	return diagnostics
}

func taskWorkflowRelationDiagnostics(record Record, recordsByID map[string][]Record) []Diagnostic {
	var diagnostics []Diagnostic
	var parent Record
	parentOK := false
	if record.Task.WorkItem != "" {
		target, ok, targetDiagnostics := validateWorkflowRelationTarget(record, "work_item", record.Task.WorkItem, RecordKindWorkItem, recordsByID)
		diagnostics = append(diagnostics, targetDiagnostics...)
		parent = target
		parentOK = ok
		if ok && (target.WorkItem == nil || !containsString(target.WorkItem.Tasks, record.ID)) {
			diagnostics = append(diagnostics, workflowMismatchDiagnostic(record, "work_item", record.Task.WorkItem, target.ID, fmt.Sprintf("%s.work_item is %s but %s.tasks does not contain %s", record.ID, record.Task.WorkItem, target.ID, record.ID)))
		}
	}
	sourceRequirementOK := false
	if record.Task.SourceRequirement != "" {
		_, ok, targetDiagnostics := validateWorkflowRelationTarget(record, "source_requirement", record.Task.SourceRequirement, RecordKindRequirement, recordsByID)
		diagnostics = append(diagnostics, targetDiagnostics...)
		sourceRequirementOK = ok
	}
	for _, value := range record.Task.DependsOn {
		_, _, targetDiagnostics := validateWorkflowRelationTarget(record, "depends_on", value, RecordKindTask, recordsByID)
		diagnostics = append(diagnostics, targetDiagnostics...)
	}
	if parentOK && sourceRequirementOK && parent.WorkItem != nil && parent.WorkItem.SourceRequirement != "" && record.Task.SourceRequirement != parent.WorkItem.SourceRequirement {
		diagnostics = append(diagnostics, workflowSourceRequirementMismatchDiagnostic(record, parent.WorkItem.SourceRequirement))
	}
	return diagnostics
}

func validateWorkflowRelationTarget(record Record, field, value string, expectedKind RecordKind, recordsByID map[string][]Record) (Record, bool, []Diagnostic) {
	if strings.TrimSpace(value) == "" {
		return Record{}, false, nil
	}
	if !validWorkflowIDForKind(value, expectedKind) {
		return Record{}, false, []Diagnostic{workflowTargetDiagnostic(record, DiagnosticInvalidWorkflowRelationTarget, field, value, "invalid_target", value)}
	}
	targets := recordsByID[normalizeRecordID(value)]
	if len(targets) == 0 {
		return Record{}, false, []Diagnostic{workflowTargetDiagnostic(record, DiagnosticUnresolvedWorkflowRelation, field, value, "unresolved", value)}
	}
	for _, target := range targets {
		if target.Kind == expectedKind {
			return target, true, nil
		}
	}
	return Record{}, false, []Diagnostic{workflowTargetDiagnostic(record, DiagnosticInvalidWorkflowRelationTarget, field, value, "invalid_target", value)}
}

func workflowTargetDiagnostic(record Record, category DiagnosticCategory, field, value, refStatus, targetID string) Diagnostic {
	message := fmt.Sprintf("%s.%s contains %s but target is absent", record.ID, field, value)
	if category == DiagnosticInvalidWorkflowRelationTarget {
		message = fmt.Sprintf("%s.%s contains invalid target %s", record.ID, field, value)
	}
	return Diagnostic{
		Category:  category,
		Severity:  DiagnosticSeverityError,
		RecordID:  record.ID,
		Path:      record.Path,
		Message:   message,
		TargetID:  targetID,
		Field:     field,
		Value:     value,
		RefStatus: refStatus,
	}
}

func workflowMismatchDiagnostic(record Record, field, value, targetID, message string) Diagnostic {
	return Diagnostic{
		Category:  DiagnosticWorkflowRelationMismatch,
		Severity:  DiagnosticSeverityError,
		RecordID:  record.ID,
		Path:      record.Path,
		Message:   message,
		TargetID:  targetID,
		Field:     field,
		Value:     value,
		RefStatus: "mismatch",
	}
}

func workflowSourceRequirementMismatchDiagnostic(record Record, expected string) Diagnostic {
	return Diagnostic{
		Category:  DiagnosticWorkflowSourceReqMismatch,
		Severity:  DiagnosticSeverityError,
		RecordID:  record.ID,
		Path:      record.Path,
		Message:   fmt.Sprintf("%s.source_requirement is %s but parent work item source_requirement is %s", record.ID, record.Task.SourceRequirement, expected),
		TargetID:  expected,
		Field:     "source_requirement",
		Value:     record.Task.SourceRequirement,
		RefStatus: "mismatch",
	}
}

func workflowSourceRequirement(record Record) string {
	if record.WorkItem == nil {
		return ""
	}
	return record.WorkItem.SourceRequirement
}

func workflowTaskWorkItem(record Record) string {
	if record.Task == nil {
		return ""
	}
	return record.Task.WorkItem
}

func containsString(values []string, needle string) bool {
	for _, value := range values {
		if value == needle {
			return true
		}
	}
	return false
}

func pathIssueDiagnostics(issues []PathIssue, candidatesByPath map[string]RecordCandidate, scope validationScope) []Diagnostic {
	diagnostics := make([]Diagnostic, 0, len(issues))
	for _, issue := range issues {
		if !scope.selectPathIssue(issue, candidatesByPath) {
			continue
		}
		diagnostics = append(diagnostics, Diagnostic{
			Category: DiagnosticMissingRecordPath,
			Severity: DiagnosticSeverityError,
			Path:     issue.Path,
			Message:  fmt.Sprintf("%s failed for record path %s: %v", issue.Operation, issue.Path, issue.Err),
		})
	}
	return diagnostics
}

func statusAllowedForKind(kind RecordKind, status RecordStatus) bool {
	switch kind {
	case RecordKindDecision:
		return status == RecordStatusProposed || status == RecordStatusAccepted || status == RecordStatusSuperseded
	case RecordKindSpec:
		return status == RecordStatusConfirmed || status == RecordStatusDraft || status == RecordStatusWIP
	case RecordKindInvestigation:
		return status == RecordStatusInvestigating || status == RecordStatusConcluded || status == RecordStatusSuperseded
	case RecordKindRequirement:
		return status == RecordStatusCaptured || status == RecordStatusDecisionNeeded || status == RecordStatusAccepted || status == RecordStatusDeferred || status == RecordStatusRejected
	case RecordKindWorkItem:
		return status == RecordStatusNotStarted || status == RecordStatusDecisionPending || status == RecordStatusDesignSpecPending || status == RecordStatusInternalDesignPending || status == RecordStatusYAMLPending || status == RecordStatusImplementationPending || status == RecordStatusFixturePending || status == RecordStatusVerificationPending || status == RecordStatusDone || status == RecordStatusBlocked
	case RecordKindTask:
		return status == RecordStatusTodo || status == RecordStatusDoing || status == RecordStatusBlocked || status == RecordStatusDone
	default:
		return false
	}
}

func recordIDExists(recordsByID map[string][]Record, id string) bool {
	records := recordsByID[normalizeRecordID(id)]
	return len(records) > 0
}

func (s validationScope) selectRecord(record Record) bool {
	if s.hasKind && record.Kind != s.kind {
		return false
	}
	if s.idRange == nil {
		return true
	}
	return s.idRange.containsRecord(record)
}

func (s validationScope) selectIssue(issue ParseIssue, recordsByPath map[string]Record, candidatesByPath map[string]RecordCandidate) bool {
	if record, ok := recordsByPath[issue.Path]; ok {
		return s.selectRecord(record)
	}
	if candidate, ok := candidatesByPath[issue.Path]; ok {
		return s.selectCandidate(candidate)
	}
	kind, ok := kindFromPath(issue.Path)
	if !ok {
		return !s.hasKind && s.idRange == nil
	}
	if s.hasKind && kind != s.kind {
		return false
	}
	if s.idRange == nil {
		return true
	}
	return s.idRange.containsID(kind, issue.RecordID)
}

func (s validationScope) selectCandidate(candidate RecordCandidate) bool {
	if s.hasKind && candidate.Kind != s.kind {
		return false
	}
	if s.idRange == nil {
		return true
	}
	return s.idRange.containsID(candidate.Kind, candidate.ID)
}

func (s validationScope) selectPathIssue(issue PathIssue, candidatesByPath map[string]RecordCandidate) bool {
	if candidate, ok := candidatesByPath[issue.Path]; ok {
		return s.selectCandidate(candidate)
	}
	kind, ok := kindFromPath(issue.Path)
	if !ok {
		return !s.hasKind && s.idRange == nil
	}
	if s.hasKind && kind != s.kind {
		return false
	}
	if s.idRange == nil {
		return true
	}
	if s.idRange.kind != RecordKindDecision {
		return false
	}
	num, ok := decisionFilenameNumber(issue.Path)
	return ok && s.idRange.contains(num)
}

func decisionFilenameNumber(path string) (int, bool) {
	value := filenameNumber(path)
	if value == "" {
		return 0, false
	}
	num, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}
	return num, true
}

func kindFromPath(path string) (RecordKind, bool) {
	switch {
	case strings.HasPrefix(path, "docs/adr/"):
		return RecordKindDecision, true
	case strings.HasPrefix(path, "docs/spec/"):
		return RecordKindSpec, true
	case strings.HasPrefix(path, "docs/investigations/"):
		return RecordKindInvestigation, true
	case strings.HasPrefix(path, "docs/requirements/"):
		return RecordKindRequirement, true
	case strings.HasPrefix(path, "docs/work-items/"):
		return RecordKindWorkItem, true
	case strings.HasPrefix(path, "docs/tasks/"):
		return RecordKindTask, true
	default:
		return "", false
	}
}

func recordDependsOn(record Record) []string {
	switch record.Kind {
	case RecordKindDecision:
		if record.Decision != nil {
			return record.Decision.DependsOn
		}
	case RecordKindSpec:
		if record.Spec != nil {
			return record.Spec.DependsOn
		}
	}
	return nil
}

func recordSupersedes(record Record) []string {
	switch record.Kind {
	case RecordKindDecision:
		if record.Decision != nil {
			return record.Decision.Supersedes
		}
	case RecordKindInvestigation:
		if record.Investigation != nil {
			return record.Investigation.Supersedes
		}
	}
	return nil
}

func collectSemanticRefs(records []Record) []SemanticRefDecl {
	var out []SemanticRefDecl
	for _, record := range records {
		out = append(out, record.SemanticRefs...)
	}
	return out
}

func matchingHeadingCount(headings []Heading, text string) int {
	count := 0
	for _, heading := range headings {
		if heading.Text == text {
			count++
		}
	}
	return count
}

type requiredSectionPolicy struct {
	sections    []string
	gatedStatus RecordStatus
}

func requiredSectionPolicyFor(record Record) *requiredSectionPolicy {
	switch record.Kind {
	case RecordKindWorkItem:
		if record.Status == RecordStatusDone {
			return &requiredSectionPolicy{
				sections:    []string{"Goal", "Boundary", "Evidence"},
				gatedStatus: RecordStatusDone,
			}
		}
	case RecordKindTask:
		if record.Status == RecordStatusDone {
			return &requiredSectionPolicy{
				sections:    []string{"Goal", "Work", "Done condition", "Verification", "Evidence"},
				gatedStatus: RecordStatusDone,
			}
		}
	case RecordKindRequirement:
		if record.Status == RecordStatusAccepted {
			return &requiredSectionPolicy{
				sections:    []string{"Requirement", "Required Outcome"},
				gatedStatus: RecordStatusAccepted,
			}
		}
	}
	return nil
}

func requiredNarrativeSectionDiagnostics(record Record) []Diagnostic {
	p := requiredSectionPolicyFor(record)
	if p == nil {
		return nil
	}
	var diagnostics []Diagnostic
	for _, sectionName := range p.sections {
		level, found := findHeadingLevel(record.Headings, sectionName)
		if !found {
			diagnostics = append(diagnostics, Diagnostic{
				Category: DiagnosticMissingRequiredSection,
				Severity: DiagnosticSeverityError,
				RecordID: record.ID,
				Path:     record.Path,
				Message:  fmt.Sprintf("required section %q must be present when %s status is %q", sectionName, record.Kind, p.gatedStatus),
				Section:  sectionName,
				Status:   string(p.gatedStatus),
			})
			continue
		}
		body := extractSectionBody(record.RawBody, sectionName, level)
		if strings.TrimSpace(body) == "" {
			diagnostics = append(diagnostics, Diagnostic{
				Category: DiagnosticEmptyRequiredSection,
				Severity: DiagnosticSeverityError,
				RecordID: record.ID,
				Path:     record.Path,
				Message:  fmt.Sprintf("required section %q must be non-empty when %s status is %q", sectionName, record.Kind, p.gatedStatus),
				Section:  sectionName,
				Status:   string(p.gatedStatus),
			})
		}
	}
	return diagnostics
}

func findHeadingLevel(headings []Heading, text string) (int, bool) {
	for _, h := range headings {
		if h.Text == text {
			return h.Level, true
		}
	}
	return 0, false
}

func extractSectionBody(raw, headingText string, headingLevel int) string {
	lines := contentLinesOutsideFrontMatterAndFences(raw)
	inSection := false
	var bodyLines []string
	for _, line := range lines {
		trimmed := trimLineEnd(line)
		match := atxHeadingPattern.FindStringSubmatch(trimmed)
		if match != nil {
			level := len(match[1])
			text := strings.TrimSpace(match[2])
			if !inSection {
				if level == headingLevel && text == headingText {
					inSection = true
				}
				continue
			}
			if level <= headingLevel {
				break
			}
		}
		if inSection {
			bodyLines = append(bodyLines, line)
		}
	}
	return strings.Join(bodyLines, "\n")
}
