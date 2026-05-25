package designrecords

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var adrIDPattern = regexp.MustCompile(`^ADR-(\d{3})$`)

type validationScope struct {
	kind    RecordKind
	hasKind bool
	idRange *numericIDRange
}

type numericIDRange struct {
	from    int
	hasFrom bool
	to      int
	hasTo   bool
}

func newValidationScope(req ValidateRecordsRequest) (validationScope, error) {
	scope := validationScope{}
	if req.Kind != "" {
		if req.Kind != RecordKindDecision && req.Kind != RecordKindSpec && req.Kind != RecordKindInvestigation {
			return scope, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unsupported kind %q", req.Kind))
		}
		scope.kind = req.Kind
		scope.hasKind = true
	}
	if req.IDRange == nil {
		return scope, nil
	}
	if req.Kind != "" && req.Kind != RecordKindDecision {
		return scope, newToolError(ErrorCodeIDRangeRequiresDecisionKind, "id_range requires kind decision")
	}
	parsed, err := parseDecisionIDRange(*req.IDRange)
	if err != nil {
		return scope, err
	}
	scope.kind = RecordKindDecision
	scope.hasKind = true
	scope.idRange = parsed
	return scope, nil
}

func parseDecisionIDRange(idRange IDRange) (*numericIDRange, error) {
	parsed := &numericIDRange{}
	if strings.TrimSpace(idRange.From) != "" {
		num, err := parseDecisionIDEndpoint(idRange.From)
		if err != nil {
			return nil, err
		}
		parsed.from = num
		parsed.hasFrom = true
	}
	if strings.TrimSpace(idRange.To) != "" {
		num, err := parseDecisionIDEndpoint(idRange.To)
		if err != nil {
			return nil, err
		}
		parsed.to = num
		parsed.hasTo = true
	}
	return parsed, nil
}

func parseDecisionIDEndpoint(value string) (int, error) {
	normalized := normalizeRecordID(value)
	match := adrIDPattern.FindStringSubmatch(normalized)
	if match == nil {
		return 0, newToolError(ErrorCodeIDRangeRequiresDecisionKind, "id_range endpoints must use ADR-NNN decision IDs")
	}
	num, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("invalid id_range endpoint %q", value))
	}
	return num, nil
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
	}
	return diagnostics
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
		case isPhysicalPathReference(value):
			diagnostics = append(diagnostics, investigationReferenceDiagnostic(record, noncanonicalCategory, severity, field, value, "noncanonical"))
		case isExplicitUnsupportedReference(value):
			diagnostics = append(diagnostics, investigationReferenceDiagnostic(record, DiagnosticUnsupportedReference, severity, field, value, "unsupported"))
		case activeSpecRefPattern.MatchString(value):
			targets := semanticByRef[value]
			if len(targets) == 0 {
				diagnostics = append(diagnostics, investigationReferenceDiagnostic(record, unresolvedCategory, severity, field, value, "unresolved"))
			}
		case recordIDRefPattern.MatchString(value):
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
	if record.Kind != RecordKindDecision {
		return false
	}
	num, ok := decisionRecordNumber(record.ID)
	return ok && s.idRange.contains(num)
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
	num, ok := decisionRecordNumber(issue.RecordID)
	return ok && s.idRange.contains(num)
}

func (s validationScope) selectCandidate(candidate RecordCandidate) bool {
	if s.hasKind && candidate.Kind != s.kind {
		return false
	}
	if s.idRange == nil {
		return true
	}
	if candidate.Kind != RecordKindDecision {
		return false
	}
	num, ok := decisionRecordNumber(candidate.ID)
	return ok && s.idRange.contains(num)
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
	num, ok := decisionFilenameNumber(issue.Path)
	return ok && s.idRange.contains(num)
}

func (r numericIDRange) contains(num int) bool {
	if r.hasFrom && num < r.from {
		return false
	}
	if r.hasTo && num > r.to {
		return false
	}
	return true
}

func decisionRecordNumber(id string) (int, bool) {
	match := adrIDPattern.FindStringSubmatch(normalizeRecordID(id))
	if match == nil {
		return 0, false
	}
	num, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, false
	}
	return num, true
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
