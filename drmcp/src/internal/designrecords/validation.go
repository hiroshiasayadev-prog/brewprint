package designrecords

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type validationScope struct {
	appNamespace string
	ref          string
}

const (
	validationRequestAppNamespacePrefix = "\x00app_namespace:"
	validationRequestRefPrefix          = "\x00ref:"
)

func (r *ValidateRecordsRequest) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if raw == nil {
		return fmt.Errorf("validate_records request must be a JSON object")
	}
	for key := range raw {
		switch key {
		case "app_namespace", "ref":
		case "kind", "domain", "id_range":
			return fmt.Errorf("%s selector is not supported by current validate_records", key)
		default:
			return fmt.Errorf("%s selector is not supported by current validate_records", key)
		}
	}

	appNamespaceRaw, hasAppNamespace := raw["app_namespace"]
	refRaw, hasRef := raw["ref"]
	if hasAppNamespace && hasRef {
		return fmt.Errorf("app_namespace and ref selectors are mutually exclusive")
	}

	switch {
	case hasAppNamespace:
		var appNamespace *string
		if err := json.Unmarshal(appNamespaceRaw, &appNamespace); err != nil || appNamespace == nil || *appNamespace == "" {
			return fmt.Errorf("app_namespace must be a non-empty string")
		}
		r.Kind = RecordKind(validationRequestAppNamespacePrefix + *appNamespace)
	case hasRef:
		var ref *string
		if err := json.Unmarshal(refRaw, &ref); err != nil || ref == nil || *ref == "" {
			return fmt.Errorf("ref must be a non-empty string")
		}
		r.Kind = RecordKind(validationRequestRefPrefix + *ref)
	default:
		*r = ValidateRecordsRequest{}
	}
	return nil
}

func newValidationScope(req ValidateRecordsRequest) (validationScope, error) {
	if req.IDRange != nil {
		return validationScope{}, newToolError(ErrorCodeInvalidRequest, "id_range selector is not supported by current validate_records")
	}
	value := string(req.Kind)
	switch {
	case value == "":
		return validationScope{}, nil
	case strings.HasPrefix(value, validationRequestAppNamespacePrefix):
		return validationScope{appNamespace: strings.TrimPrefix(value, validationRequestAppNamespacePrefix)}, nil
	case strings.HasPrefix(value, validationRequestRefPrefix):
		return validationScope{ref: strings.TrimPrefix(value, validationRequestRefPrefix)}, nil
	default:
		return validationScope{}, newToolError(ErrorCodeInvalidRequest, "kind selector is not supported by current validate_records")
	}
}

func generateValidationDiagnostics(idx *Index, scope validationScope) []Diagnostic {
	var diagnostics []Diagnostic

	recordsByPath := make(map[string]Record, len(idx.Records))
	recordsByID := make(map[string][]Record, len(idx.Records))
	for _, record := range idx.Records {
		recordsByPath[record.Path] = record
		if record.NormalizedID != "" {
			recordsByID[record.ID] = append(recordsByID[record.ID], record)
		}
	}

	candidatesByPath := make(map[string]RecordCandidate, len(idx.Candidates))
	for _, candidate := range idx.Candidates {
		if candidate.Path != "" {
			candidatesByPath[candidate.Path] = candidate
		}
	}

	diagnostics = append(diagnostics, currentConflictDiagnostics(idx, scope)...)
	diagnostics = append(diagnostics, parseIssueDiagnostics(idx, idx.ParseIssues, recordsByPath, candidatesByPath, scope)...)
	diagnostics = append(diagnostics, recordDiagnostics(idx, idx.Records, recordsByID, scope)...)
	diagnostics = append(diagnostics, pathIssueDiagnostics(idx, idx.PathIssues, candidatesByPath, scope)...)
	return sortAndDedupeDiagnostics(idx, diagnostics)
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
			if !scope.selectRecord(nil, record) {
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

func currentConflictDiagnostics(idx *Index, scope validationScope) []Diagnostic {
	var diagnostics []Diagnostic
	for _, group := range idx.ConflictGroups {
		if !scope.selectConflictGroup(idx, group) {
			continue
		}
		for _, source := range group.Sources {
			if scope.appNamespace != "" && pathAppNamespace(idx, source) != scope.appNamespace {
				continue
			}
			diagnostics = append(diagnostics, Diagnostic{
				Category: DiagnosticCategory("current_conflict"),
				Severity: DiagnosticSeverityError,
				RecordID: group.Ref,
				Path:     source,
				Message:  fmt.Sprintf("current ref %s has conflicting current sources", group.Ref),
				Location: currentDiagnosticLocation(idx, source),
			})
		}
	}
	return diagnostics
}

func parseIssueDiagnostics(idx *Index, issues []ParseIssue, recordsByPath map[string]Record, candidatesByPath map[string]RecordCandidate, scope validationScope) []Diagnostic {
	diagnostics := make([]Diagnostic, 0, len(issues))
	for _, issue := range issues {
		if !scope.selectIssue(idx, issue, recordsByPath, candidatesByPath) {
			continue
		}
		diagnostics = append(diagnostics, Diagnostic{
			Category: issue.Category,
			Severity: DiagnosticSeverityError,
			RecordID: issue.RecordID,
			Path:     issue.Path,
			Message:  issue.Message,
			Location: currentDiagnosticLocation(idx, issue.Path),
		})
	}
	return diagnostics
}

func recordDiagnostics(idx *Index, records []Record, recordsByID map[string][]Record, scope validationScope) []Diagnostic {
	var diagnostics []Diagnostic
	for _, record := range records {
		if !scope.selectRecord(idx, record) {
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
			diagnostics = append(diagnostics, investigationReferenceDiagnostics(idx, record, recordsByID)...)
		}
		diagnostics = append(diagnostics, workflowRelationDiagnostics(idx, record, recordsByID)...)
	}
	for i := range diagnostics {
		if diagnostics[i].Path != "" && diagnostics[i].Location == nil {
			diagnostics[i].Location = currentDiagnosticLocation(idx, diagnostics[i].Path)
		}
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

func investigationReferenceDiagnostics(idx *Index, record Record, recordsByID map[string][]Record) []Diagnostic {
	var diagnostics []Diagnostic
	diagnostics = append(diagnostics, diagnosticsForInvestigationRefs(idx, record, "source_refs", record.Investigation.SourceRefs, DiagnosticUnresolvedSourceRef, DiagnosticSeverityError, recordsByID)...)
	diagnostics = append(diagnostics, diagnosticsForInvestigationRefs(idx, record, "follow_up_results", record.Investigation.FollowUpResults, DiagnosticUnresolvedFollowUpResult, DiagnosticSeverityError, recordsByID)...)
	diagnostics = append(diagnostics, diagnosticsForInvestigationRefs(idx, record, "follow_up_candidates", record.Investigation.FollowUpCandidates, DiagnosticUnresolvedFollowUpCandidate, DiagnosticSeverityInfo, recordsByID)...)
	return diagnostics
}

func diagnosticsForInvestigationRefs(idx *Index, record Record, field string, values []string, unresolvedCategory DiagnosticCategory, severity DiagnosticSeverity, recordsByID map[string][]Record) []Diagnostic {
	var diagnostics []Diagnostic
	for _, value := range values {
		switch {
		case strings.HasPrefix(value, "yaml:"):
			continue
		case isPhysicalPathReference(value):
			diagnostics = append(diagnostics, investigationReferenceDiagnostic(record, DiagnosticUnsupportedReference, severity, field, value, "unsupported"))
		case isExplicitUnsupportedReference(value):
			diagnostics = append(diagnostics, investigationReferenceDiagnostic(record, DiagnosticUnsupportedReference, severity, field, value, "unsupported"))
		case activeSpecRefPattern.MatchString(value):
			targets := recordsByID[value]
			if len(targets) == 0 {
				diagnostics = append(diagnostics, investigationReferenceDiagnostic(record, unresolvedCategory, severity, field, value, "unresolved"))
			}
		case isCurrentRecordIDReference(idx, value) && isSupportedInvestigationRecordIDReference(stripKnownNamespacePrefix(idx, value)):
			targets := recordsByID[value]
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

func workflowRelationDiagnostics(idx *Index, record Record, recordsByID map[string][]Record) []Diagnostic {
	switch record.Kind {
	case RecordKindRequirement:
		if record.Requirement == nil {
			return nil
		}
		return requirementWorkflowRelationDiagnostics(idx, record, recordsByID)
	case RecordKindWorkItem:
		if record.WorkItem == nil {
			return nil
		}
		return workItemWorkflowRelationDiagnostics(idx, record, recordsByID)
	case RecordKindTask:
		if record.Task == nil {
			return nil
		}
		return taskWorkflowRelationDiagnostics(idx, record, recordsByID)
	default:
		return nil
	}
}

func requirementWorkflowRelationDiagnostics(idx *Index, record Record, recordsByID map[string][]Record) []Diagnostic {
	var diagnostics []Diagnostic
	for _, value := range record.Requirement.WorkItems {
		target, ok, targetDiagnostics := validateWorkflowRelationTarget(idx, record, "work_items", value, RecordKindWorkItem, recordsByID)
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

func workItemWorkflowRelationDiagnostics(idx *Index, record Record, recordsByID map[string][]Record) []Diagnostic {
	var diagnostics []Diagnostic
	if record.WorkItem.SourceRequirement != "" {
		target, ok, targetDiagnostics := validateWorkflowRelationTarget(idx, record, "source_requirement", record.WorkItem.SourceRequirement, RecordKindRequirement, recordsByID)
		diagnostics = append(diagnostics, targetDiagnostics...)
		if ok && (target.Requirement == nil || !containsString(target.Requirement.WorkItems, record.ID)) {
			diagnostics = append(diagnostics, workflowMismatchDiagnostic(record, "source_requirement", record.WorkItem.SourceRequirement, target.ID, fmt.Sprintf("%s.source_requirement is %s but %s.work_items does not contain %s", record.ID, record.WorkItem.SourceRequirement, target.ID, record.ID)))
		}
	}
	for _, value := range record.WorkItem.Tasks {
		target, ok, targetDiagnostics := validateWorkflowRelationTarget(idx, record, "tasks", value, RecordKindTask, recordsByID)
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

func taskWorkflowRelationDiagnostics(idx *Index, record Record, recordsByID map[string][]Record) []Diagnostic {
	var diagnostics []Diagnostic
	var parent Record
	parentOK := false
	if record.Task.WorkItem != "" {
		target, ok, targetDiagnostics := validateWorkflowRelationTarget(idx, record, "work_item", record.Task.WorkItem, RecordKindWorkItem, recordsByID)
		diagnostics = append(diagnostics, targetDiagnostics...)
		parent = target
		parentOK = ok
		if ok && (target.WorkItem == nil || !containsString(target.WorkItem.Tasks, record.ID)) {
			diagnostics = append(diagnostics, workflowMismatchDiagnostic(record, "work_item", record.Task.WorkItem, target.ID, fmt.Sprintf("%s.work_item is %s but %s.tasks does not contain %s", record.ID, record.Task.WorkItem, target.ID, record.ID)))
		}
	}
	sourceRequirementOK := false
	if record.Task.SourceRequirement != "" {
		_, ok, targetDiagnostics := validateWorkflowRelationTarget(idx, record, "source_requirement", record.Task.SourceRequirement, RecordKindRequirement, recordsByID)
		diagnostics = append(diagnostics, targetDiagnostics...)
		sourceRequirementOK = ok
	}
	for _, value := range record.Task.DependsOn {
		_, _, targetDiagnostics := validateWorkflowRelationTarget(idx, record, "depends_on", value, RecordKindTask, recordsByID)
		diagnostics = append(diagnostics, targetDiagnostics...)
	}
	if parentOK && sourceRequirementOK && parent.WorkItem != nil && parent.WorkItem.SourceRequirement != "" && record.Task.SourceRequirement != parent.WorkItem.SourceRequirement {
		diagnostics = append(diagnostics, workflowSourceRequirementMismatchDiagnostic(record, parent.WorkItem.SourceRequirement))
	}
	return diagnostics
}

func validateWorkflowRelationTarget(idx *Index, record Record, field, value string, expectedKind RecordKind, recordsByID map[string][]Record) (Record, bool, []Diagnostic) {
	if strings.TrimSpace(value) == "" {
		return Record{}, false, nil
	}
	targets := recordsByID[value]
	for _, target := range targets {
		if target.Kind == expectedKind {
			return target, true, nil
		}
	}
	if len(targets) > 0 {
		// Found at least one record with this ID but none matched expectedKind.
		return Record{}, false, []Diagnostic{workflowTargetDiagnostic(record, DiagnosticInvalidWorkflowRelationTarget, field, value, "invalid_target", value)}
	}
	if !validCurrentWorkflowIDForKind(idx, value, expectedKind) {
		return Record{}, false, []Diagnostic{workflowTargetDiagnostic(record, DiagnosticInvalidWorkflowRelationTarget, field, value, "invalid_target", value)}
	}
	return Record{}, false, []Diagnostic{workflowTargetDiagnostic(record, DiagnosticUnresolvedWorkflowRelation, field, value, "unresolved", value)}
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

func pathIssueDiagnostics(idx *Index, issues []PathIssue, candidatesByPath map[string]RecordCandidate, scope validationScope) []Diagnostic {
	diagnostics := make([]Diagnostic, 0, len(issues))
	for _, issue := range issues {
		if !scope.selectPathIssue(idx, issue, candidatesByPath) {
			continue
		}
		diagnostics = append(diagnostics, Diagnostic{
			Category: DiagnosticMissingRecordPath,
			Severity: DiagnosticSeverityError,
			Path:     issue.Path,
			Message:  fmt.Sprintf("%s failed for record path %s: %v", issue.Operation, issue.Path, issue.Err),
			Location: currentDiagnosticLocation(idx, issue.Path),
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
		return status == RecordStatusNotStarted || status == RecordStatusInProgress || status == RecordStatusBlocked || status == RecordStatusDone
	case RecordKindTask:
		return status == RecordStatusNotStarted || status == RecordStatusInProgress || status == RecordStatusBlocked || status == RecordStatusDone
	default:
		return false
	}
}

func recordIDExists(recordsByID map[string][]Record, id string) bool {
	records := recordsByID[id]
	return len(records) > 0
}

func (s validationScope) selectRecord(idx *Index, record Record) bool {
	if s.ref != "" {
		return record.ID == s.ref
	}
	if s.appNamespace != "" {
		return pathAppNamespace(idx, record.Path) == s.appNamespace
	}
	return true
}

func (s validationScope) selectIssue(idx *Index, issue ParseIssue, recordsByPath map[string]Record, candidatesByPath map[string]RecordCandidate) bool {
	if record, ok := recordsByPath[issue.Path]; ok {
		return s.selectRecord(idx, record)
	}
	if candidate, ok := candidatesByPath[issue.Path]; ok {
		return s.selectCandidate(idx, candidate)
	}
	if s.ref != "" {
		return issue.RecordID == s.ref
	}
	if s.appNamespace != "" {
		return pathAppNamespace(idx, issue.Path) == s.appNamespace
	}
	return true
}

func (s validationScope) selectCandidate(idx *Index, candidate RecordCandidate) bool {
	if s.ref != "" {
		return candidate.ID == s.ref
	}
	if s.appNamespace != "" {
		return pathAppNamespace(idx, candidate.Path) == s.appNamespace
	}
	return true
}

func (s validationScope) selectPathIssue(idx *Index, issue PathIssue, candidatesByPath map[string]RecordCandidate) bool {
	if candidate, ok := candidatesByPath[issue.Path]; ok {
		return s.selectCandidate(idx, candidate)
	}
	if s.ref != "" {
		return false
	}
	if s.appNamespace != "" {
		return pathAppNamespace(idx, issue.Path) == s.appNamespace
	}
	return true
}

func (s validationScope) selectConflictGroup(idx *Index, group CurrentConflict) bool {
	if s.ref != "" {
		return group.Ref == s.ref
	}
	if s.appNamespace != "" {
		for _, source := range group.Sources {
			if pathAppNamespace(idx, source) == s.appNamespace {
				return true
			}
		}
		return false
	}
	return true
}

func decisionFilenameNumber(path, ns string) (int, bool) {
	value := filenameNumber(path, ns)
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
	case strings.HasPrefix(path, "v01/records/adr/"):
		return RecordKindDecision, true
	case strings.HasPrefix(path, "v01/records/spec/"):
		return RecordKindSpec, true
	case strings.HasPrefix(path, "v01/records/investigations/"):
		return RecordKindInvestigation, true
	case strings.HasPrefix(path, "v01/records/requirements/"):
		return RecordKindRequirement, true
	case strings.HasPrefix(path, "v01/records/work-items/"):
		return RecordKindWorkItem, true
	case strings.HasPrefix(path, "v01/records/tasks/"):
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
			if actual, ok := findCaseOnlyMismatch(record.Headings, sectionName); ok {
				candidates := make([]CandidateHeading, len(record.Headings))
				for i, h := range record.Headings {
					candidates[i] = CandidateHeading{Heading: h.Text, Level: h.Level, Ordinal: i + 1}
				}
				diagnostics = append(diagnostics, Diagnostic{
					Category:          DiagnosticSectionHeadingCaseMismatch,
					Severity:          DiagnosticSeverityInfo,
					RecordID:          record.ID,
					Path:              record.Path,
					Message:           fmt.Sprintf("canonical required section %q found with non-canonical case as %q; repair heading to canonical text", sectionName, actual),
					Section:           sectionName,
					Status:            string(p.gatedStatus),
					ActualHeading:     actual,
					CandidateHeadings: candidates,
				})
			}
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

func findCaseOnlyMismatch(headings []Heading, canonical string) (string, bool) {
	for _, h := range headings {
		if strings.EqualFold(h.Text, canonical) && h.Text != canonical {
			return h.Text, true
		}
	}
	return "", false
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

func validateCurrentValidationScope(idx *Index, req CurrentValidateRecordsRequest) (validationScope, string, error) {
	if req.AppNamespace != "" && req.Ref != "" {
		return validationScope{}, "", newToolError(ErrorCodeInvalidRequest, "app_namespace and ref selectors are mutually exclusive")
	}
	if req.AppNamespace != "" {
		if !knownAppNamespace(idx, req.AppNamespace) {
			return validationScope{}, "", newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unknown app_namespace %q", req.AppNamespace))
		}
		return validationScope{appNamespace: req.AppNamespace}, "app_namespace", nil
	}
	if req.Ref != "" {
		if !validCurrentRefSelector(idx, req.Ref) {
			return validationScope{}, "", newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unsupported current ref selector %q", req.Ref))
		}
		return validationScope{ref: req.Ref}, "ref", nil
	}
	return validationScope{}, "all", nil
}

func knownAppNamespace(idx *Index, appNamespace string) bool {
	for _, entry := range idx.RecordsEntries {
		if entry.AppNamespace == appNamespace {
			return true
		}
	}
	return false
}

func validCurrentRefSelector(idx *Index, ref string) bool {
	if ref == "" || isPhysicalPathReference(ref) || strings.HasPrefix(ref, "V01-") {
		return false
	}
	if validCurrentSpecRefSelector(ref) {
		return true
	}
	return isCurrentRecordIDReference(idx, ref)
}

func validCurrentSpecRefSelector(ref string) bool {
	if !strings.HasPrefix(ref, "spec:") || strings.ContainsAny(ref, `/\`) || strings.HasSuffix(strings.ToLower(ref), ".md") {
		return false
	}
	rest := strings.TrimPrefix(ref, "spec:")
	if rest == "" || strings.HasPrefix(rest, ".") || strings.HasSuffix(rest, ".") || strings.Contains(rest, "..") {
		return false
	}
	for _, r := range rest {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_' || r == '.' {
			continue
		}
		return false
	}
	return true
}

func isCurrentRecordIDReference(idx *Index, ref string) bool {
	bare := stripKnownNamespacePrefix(idx, ref)
	return bare != ref && currentNamespaceIDPattern.MatchString(ref)
}

func validCurrentWorkflowIDForKind(idx *Index, ref string, kind RecordKind) bool {
	bare := stripKnownNamespacePrefix(idx, ref)
	return bare != ref && validWorkflowIDForKind(bare, kind)
}

func stripKnownNamespacePrefix(idx *Index, ref string) string {
	if idx != nil {
		for _, entry := range idx.RecordsEntries {
			if entry.NamespacePrefix != "" && strings.HasPrefix(ref, entry.NamespacePrefix) {
				return strings.TrimPrefix(ref, entry.NamespacePrefix)
			}
		}
		if idx.NamespacePrefix != "" && strings.HasPrefix(ref, idx.NamespacePrefix) {
			return strings.TrimPrefix(ref, idx.NamespacePrefix)
		}
	}
	return ref
}

func pathAppNamespace(idx *Index, path string) string {
	if idx == nil {
		return ""
	}
	entry, ok := recordsEntryForPath(idx, path)
	if !ok {
		return ""
	}
	return entry.AppNamespace
}

func currentDiagnosticLocation(idx *Index, path string) *DiagnosticLocation {
	if idx == nil || path == "" {
		return nil
	}
	entry, ok := recordsEntryForPath(idx, path)
	if !ok {
		return nil
	}
	recordsRoot := strings.TrimSuffix(entry.RecordsRoot, "/")
	if !isPortableRepositoryRelativePath(recordsRoot) || !isPortableRepositoryRelativePath(path) || path == recordsRoot {
		return nil
	}
	prefix := recordsRoot + "/"
	if !strings.HasPrefix(path, prefix) {
		return nil
	}
	rel := strings.TrimPrefix(path, prefix)
	if !isPortableRepositoryRelativePath(rel) {
		return nil
	}
	return &DiagnosticLocation{
		SourceScope:  "current",
		RecordsRoot:  recordsRoot,
		Path:         rel,
		AppNamespace: entry.AppNamespace,
	}
}

func isPortableRepositoryRelativePath(path string) bool {
	if path == "" || path == "." || path == ".." {
		return false
	}
	if strings.Contains(path, `\`) || strings.HasPrefix(path, "/") || strings.Contains(path, ":") {
		return false
	}
	if len(path) >= 2 && ((path[0] >= 'A' && path[0] <= 'Z') || (path[0] >= 'a' && path[0] <= 'z')) && path[1] == ':' {
		return false
	}
	for _, segment := range strings.Split(path, "/") {
		if segment == "" || segment == "." || segment == ".." {
			return false
		}
	}
	return true
}

func recordsEntryForPath(idx *Index, path string) (RecordsEntry, bool) {
	var best RecordsEntry
	bestLen := -1
	for _, entry := range idx.RecordsEntries {
		root := strings.TrimSuffix(entry.RecordsRoot, "/")
		if path == root || strings.HasPrefix(path, root+"/") {
			if len(root) > bestLen {
				best = entry
				bestLen = len(root)
			}
		}
	}
	if bestLen >= 0 {
		return best, true
	}
	return RecordsEntry{}, false
}

func sortAndDedupeDiagnostics(idx *Index, diagnostics []Diagnostic) []Diagnostic {
	for i := range diagnostics {
		if diagnostics[i].Path != "" && diagnostics[i].Location == nil {
			diagnostics[i].Location = currentDiagnosticLocation(idx, diagnostics[i].Path)
		}
	}
	sort.SliceStable(diagnostics, func(i, j int) bool {
		return diagnosticStableKey(diagnostics[i]) < diagnosticStableKey(diagnostics[j])
	})
	out := diagnostics[:0]
	seen := map[string]bool{}
	for _, diagnostic := range diagnostics {
		key := diagnosticStableKey(diagnostic)
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, diagnostic)
	}
	return out
}

func diagnosticStableKey(d Diagnostic) string {
	encoded, _ := json.Marshal(d)
	var structured map[string]json.RawMessage
	_ = json.Unmarshal(encoded, &structured)
	delete(structured, "message")
	delete(structured, "Message")
	key, _ := json.Marshal(structured)
	return string(key)
}

func currentValidationSubjectSummary(idx *Index, scope validationScope, diagnostics []Diagnostic) ValidationSubjectSummary {
	subjects := map[string]bool{}
	for _, record := range idx.Records {
		if scope.selectRecord(idx, record) {
			subjects[record.Path] = true
		}
	}
	for _, candidate := range idx.Candidates {
		if scope.selectCandidate(idx, candidate) {
			subjects[candidate.Path] = true
		}
	}
	for _, group := range idx.ConflictGroups {
		if !scope.selectConflictGroup(idx, group) {
			continue
		}
		for _, source := range group.Sources {
			if scope.appNamespace == "" || pathAppNamespace(idx, source) == scope.appNamespace {
				subjects[source] = true
			}
		}
	}
	invalid := map[string]bool{}
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity != DiagnosticSeverityError || diagnostic.Path == "" {
			continue
		}
		invalid[diagnostic.Path] = true
	}
	return ValidationSubjectSummary{Total: len(subjects), Invalid: len(invalid)}
}

func ValidateCurrentRecords(ctx context.Context, idx *Index, req CurrentValidateRecordsRequest) (CurrentValidateRecordsResponse, error) {
	if err := ctx.Err(); err != nil {
		return CurrentValidateRecordsResponse{}, err
	}
	if idx == nil {
		return CurrentValidateRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	scope, scopeName, err := validateCurrentValidationScope(idx, req)
	if err != nil {
		return CurrentValidateRecordsResponse{}, err
	}
	diagnostics := generateValidationDiagnostics(idx, scope)
	ok := true
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == DiagnosticSeverityError {
			ok = false
			break
		}
	}
	return CurrentValidateRecordsResponse{
		OK:          ok,
		Scope:       scopeName,
		Summary:     currentValidationSubjectSummary(idx, scope, diagnostics),
		Diagnostics: diagnostics,
	}, nil
}

func ValidateRecords(ctx context.Context, idx *Index, req ValidateRecordsRequest) (ValidateRecordsResponse, error) {
	if err := ctx.Err(); err != nil {
		return ValidateRecordsResponse{}, err
	}
	if idx == nil {
		return ValidateRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	scope, err := newValidationScope(req)
	if err != nil {
		return ValidateRecordsResponse{}, err
	}
	if scope.appNamespace != "" && !knownAppNamespace(idx, scope.appNamespace) {
		return ValidateRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unknown app_namespace %q", scope.appNamespace))
	}
	if scope.ref != "" && !validCurrentRefSelector(idx, scope.ref) {
		return ValidateRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unsupported current ref selector %q", scope.ref))
	}
	diagnostics := generateValidationDiagnostics(idx, scope)
	ok := true
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == DiagnosticSeverityError {
			ok = false
			break
		}
	}
	return ValidateRecordsResponse{
		OK:          ok,
		Diagnostics: diagnostics,
	}, nil
}
