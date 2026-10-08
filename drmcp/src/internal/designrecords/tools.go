package designrecords

import (
	"context"
	"fmt"
	"sort"
	"strings"
)

const (
	defaultCurrentListLimit = 20
	maxCurrentListLimit     = 100
	maxCurrentGetRefs       = 20
)

// ListCurrentRecords returns the accepted compact current-list projection for
// the MCP list_records tool.
func ListCurrentRecords(ctx context.Context, idx *Index, req CurrentListRecordsRequest) (CurrentListRecordsResponse, error) {
	if err := ctx.Err(); err != nil {
		return CurrentListRecordsResponse{}, err
	}
	if idx == nil {
		return CurrentListRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	scope, err := newCurrentListScope(req)
	if err != nil {
		return CurrentListRecordsResponse{}, err
	}

	records := make([]Record, 0, len(idx.Records))
	for _, record := range idx.Records {
		if scope.selectRecord(record) {
			records = append(records, record)
		}
	}
	sort.SliceStable(records, func(i, j int) bool {
		if scope.order == "asc" {
			return records[i].ID < records[j].ID
		}
		return records[i].ID > records[j].ID
	})

	hasMore := len(records) > scope.limit
	if hasMore {
		records = records[:scope.limit]
	}

	out := make([]CurrentListedRecord, 0, len(records))
	for _, record := range records {
		out = append(out, currentListedRecord(record))
	}
	return CurrentListRecordsResponse{Records: out, HasMore: hasMore, Warnings: []OperationWarning{}}, nil
}

// GetCurrentRecords returns successful exact current-record lookups only.
// Per-ref misses and request duplicates are reported as top-level warnings.
func GetCurrentRecords(ctx context.Context, idx *Index, req CurrentGetRecordsRequest) (CurrentGetRecordsResponse, error) {
	if err := ctx.Err(); err != nil {
		return CurrentGetRecordsResponse{}, err
	}
	if idx == nil {
		return CurrentGetRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	if len(req.Refs) == 0 {
		return CurrentGetRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, "refs must be a non-empty array")
	}
	if len(req.Refs) > maxCurrentGetRefs {
		return CurrentGetRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, "refs must contain at most 20 entries")
	}

	records := make([]CurrentGetRecordsRecord, 0, len(req.Refs))
	warnings := make([]OperationWarning, 0)
	seen := map[string]bool{}
	for _, ref := range req.Refs {
		if seen[ref] {
			warnings = append(warnings, operationWarning("duplicate_ref", ref, "duplicate requested ref was ignored after its first occurrence"))
			continue
		}
		seen[ref] = true

		if category, ok := invalidCurrentRefCategory(ref); ok {
			warnings = append(warnings, operationWarning(category, ref, fmt.Sprintf("requested ref %q is %s", ref, strings.TrimSuffix(category, "_ref"))))
			continue
		}

		matches := matchingCurrentRecords(idx, ref)
		switch len(matches) {
		case 0:
			warnings = append(warnings, operationWarning("unresolved_ref", ref, fmt.Sprintf("requested ref %q was not found", ref)))
		case 1:
			records = append(records, currentGetRecord(matches[0], req.IncludeBody))
		default:
			warnings = append(warnings, operationWarning("duplicate_ref", ref, fmt.Sprintf("requested ref %q matched multiple current records", ref)))
		}
	}
	return CurrentGetRecordsResponse{Records: records, Warnings: warnings}, nil
}

type currentListScope struct {
	appNamespace string
	prefix       string
	kind         RecordKind
	domain       string
	status       RecordStatus
	order        string
	limit        int
}

func newCurrentListScope(req CurrentListRecordsRequest) (currentListScope, error) {
	scope := currentListScope{order: "desc", limit: defaultCurrentListLimit}
	appNamespace := strings.TrimSpace(req.AppNamespace)
	if appNamespace == "" {
		return scope, newToolError(ErrorCodeInvalidRequest, "app_namespace is required")
	}
	if !isCurrentListKind(req.Kind) {
		return scope, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unsupported kind %q", req.Kind))
	}
	domain := strings.TrimSpace(req.Domain)
	if domain == "" {
		return scope, newToolError(ErrorCodeInvalidRequest, "domain is required")
	}
	if req.Order != "" && req.Order != "asc" && req.Order != "desc" {
		return scope, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unsupported order %q", req.Order))
	}
	if req.Limit != nil {
		if *req.Limit < 1 || *req.Limit > maxCurrentListLimit {
			return scope, newToolError(ErrorCodeInvalidRequest, "limit must be between 1 and 100")
		}
		scope.limit = *req.Limit
	}
	if req.Order != "" {
		scope.order = req.Order
	}
	scope.appNamespace = appNamespace
	scope.prefix = strings.ToUpper(appNamespace) + "-"
	scope.kind = req.Kind
	scope.domain = strings.ToUpper(domain)
	scope.status = req.Status
	return scope, nil
}

func (s currentListScope) selectRecord(record Record) bool {
	if record.Kind != s.kind || record.Kind == RecordKindSpec {
		return false
	}
	if !strings.HasPrefix(record.ID, s.prefix) {
		return false
	}
	if currentRecordDomain(record.ID, s.prefix, record.Kind) != s.domain {
		return false
	}
	if s.status != "" && record.Status != s.status {
		return false
	}
	return true
}

func sortRecordsByID(records []Record, order string) {
	sort.SliceStable(records, func(i, j int) bool {
		return compareRecordsByID(records[i], records[j], order) < 0
	})
}

func compareRecordsByID(a, b Record, order string) int {
	cmp := compareRecordID(a, b)
	if order == "desc" {
		return -cmp
	}
	return cmp
}

func compareRecordID(a, b Record) int {
	switch {
	case a.ID < b.ID:
		return -1
	case a.ID > b.ID:
		return 1
	case a.Path < b.Path:
		return -1
	case a.Path > b.Path:
		return 1
	default:
		return 0
	}
}

func listedRecord(record Record) ListedRecord {
	return ListedRecord{
		ID:            record.ID,
		Kind:          record.Kind,
		Title:         record.Title,
		Status:        record.Status,
		Path:          record.Path,
		Decision:      responseDecisionDetail(record),
		Spec:          responseSpecDetail(record),
		Investigation: responseInvestigationDetail(record),
		Requirement:   responseRequirementDetail(record),
		WorkItem:      responseWorkItemDetail(record),
		Task:          responseTaskDetail(record),
	}
}

func currentListedRecord(record Record) CurrentListedRecord {
	return CurrentListedRecord{
		Ref:    record.ID,
		Title:  nullableString(record.Title),
		Status: nullableStatus(record.Status),
		Date:   nullableString(record.Date),
	}
}

func currentGetRecord(record Record, includeBody bool) CurrentGetRecordsRecord {
	out := CurrentGetRecordsRecord{
		Ref:      record.ID,
		Kind:     record.Kind,
		Title:    record.Title,
		Status:   record.Status,
		Date:     record.Date,
		Headings: append([]Heading{}, record.Headings...),
	}
	if includeBody {
		body := record.RawBody
		out.Body = &body
	}
	return out
}

func getRecordResponseRecord(record Record, includeBody bool) GetRecordRecord {
	out := GetRecordRecord{
		ID:            record.ID,
		Kind:          record.Kind,
		Title:         record.Title,
		Status:        record.Status,
		Path:          record.Path,
		Decision:      responseDecisionDetail(record),
		Spec:          responseSpecDetail(record),
		Investigation: responseInvestigationDetail(record),
		Requirement:   responseRequirementDetail(record),
		WorkItem:      responseWorkItemDetail(record),
		Task:          responseTaskDetail(record),
		Headings:      append([]Heading{}, record.Headings...),
	}
	if includeBody {
		body := record.RawBody
		out.Body = &body
	}
	return out
}

func cloneDecisionDetail(in *DecisionDetail) *DecisionDetail {
	if in == nil {
		return nil
	}
	return &DecisionDetail{
		DependsOn:      append([]string{}, in.DependsOn...),
		Supersedes:     append([]string{}, in.Supersedes...),
		MigratedToSpec: in.MigratedToSpec,
	}
}

func responseDecisionDetail(record Record) *DecisionDetail {
	if record.Kind != RecordKindDecision {
		return nil
	}
	if record.Decision == nil {
		return &DecisionDetail{DependsOn: []string{}, Supersedes: []string{}}
	}
	return cloneDecisionDetail(record.Decision)
}

func cloneSpecDetail(in *SpecDetail) *SpecDetail {
	if in == nil {
		return nil
	}
	return &SpecDetail{DependsOn: append([]string{}, in.DependsOn...)}
}

func responseSpecDetail(record Record) *SpecDetail {
	if record.Kind != RecordKindSpec {
		return nil
	}
	if record.Spec == nil {
		return &SpecDetail{DependsOn: []string{}}
	}
	return cloneSpecDetail(record.Spec)
}

func cloneInvestigationDetail(in *InvestigationDetail) *InvestigationDetail {
	if in == nil {
		return nil
	}
	return &InvestigationDetail{
		Trigger:               in.Trigger,
		Scope:                 in.Scope,
		NonScope:              in.NonScope,
		SourceRefs:            append([]string{}, in.SourceRefs...),
		FollowUpCandidates:    append([]string{}, in.FollowUpCandidates...),
		Supersedes:            append([]string{}, in.Supersedes...),
		RelatedRequirements:   append([]string{}, in.RelatedRequirements...),
		RelatedWorkItems:      append([]string{}, in.RelatedWorkItems...),
		RelatedADRs:           append([]string{}, in.RelatedADRs...),
		RelatedSpecs:          append([]string{}, in.RelatedSpecs...),
		RelatedInternalDesign: append([]string{}, in.RelatedInternalDesign...),
		RelatedCoverage:       append([]string{}, in.RelatedCoverage...),
		FollowUpResults:       append([]string{}, in.FollowUpResults...),
	}
}

func responseInvestigationDetail(record Record) *InvestigationDetail {
	if record.Kind != RecordKindInvestigation {
		return nil
	}
	if record.Investigation == nil {
		return &InvestigationDetail{SourceRefs: []string{}, FollowUpCandidates: []string{}}
	}
	return cloneInvestigationDetail(record.Investigation)
}

func cloneRequirementDetail(in *RequirementDetail) *RequirementDetail {
	if in == nil {
		return nil
	}
	return &RequirementDetail{
		SourceRefs: append([]string{}, in.SourceRefs...),
		WorkItems:  append([]string{}, in.WorkItems...),
	}
}

func responseRequirementDetail(record Record) *RequirementDetail {
	if record.Kind != RecordKindRequirement {
		return nil
	}
	if record.Requirement == nil {
		return &RequirementDetail{SourceRefs: []string{}, WorkItems: []string{}}
	}
	return cloneRequirementDetail(record.Requirement)
}

func cloneWorkItemDetail(in *WorkItemDetail) *WorkItemDetail {
	if in == nil {
		return nil
	}
	return &WorkItemDetail{
		SourceRequirement: in.SourceRequirement,
		ImpactRefs:        append([]string{}, in.ImpactRefs...),
		Tasks:             append([]string{}, in.Tasks...),
	}
}

func responseWorkItemDetail(record Record) *WorkItemDetail {
	if record.Kind != RecordKindWorkItem {
		return nil
	}
	if record.WorkItem == nil {
		return &WorkItemDetail{ImpactRefs: []string{}, Tasks: []string{}}
	}
	return cloneWorkItemDetail(record.WorkItem)
}

func cloneTaskDetail(in *TaskDetail) *TaskDetail {
	if in == nil {
		return nil
	}
	return &TaskDetail{
		WorkItem:          in.WorkItem,
		SourceRequirement: in.SourceRequirement,
		Estimate:          in.Estimate,
		DependsOn:         append([]string{}, in.DependsOn...),
		Outputs:           append([]string{}, in.Outputs...),
	}
}

func responseTaskDetail(record Record) *TaskDetail {
	if record.Kind != RecordKindTask {
		return nil
	}
	if record.Task == nil {
		return &TaskDetail{DependsOn: []string{}, Outputs: []string{}}
	}
	return cloneTaskDetail(record.Task)
}

func isListableRecordKind(kind RecordKind) bool {
	switch kind {
	case RecordKindDecision, RecordKindSpec, RecordKindInvestigation, RecordKindRequirement, RecordKindWorkItem, RecordKindTask:
		return true
	default:
		return false
	}
}

func isCurrentListKind(kind RecordKind) bool {
	switch kind {
	case RecordKindDecision, RecordKindInvestigation, RecordKindRequirement, RecordKindWorkItem, RecordKindTask:
		return true
	default:
		return false
	}
}

func nullableString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func nullableStatus(value RecordStatus) *RecordStatus {
	if value == "" {
		return nil
	}
	return &value
}

func operationWarning(category, ref, message string) OperationWarning {
	return OperationWarning{Category: category, Ref: ref, Message: message}
}

func invalidCurrentRefCategory(ref string) (string, bool) {
	if ref == "" || strings.TrimSpace(ref) != ref {
		return "malformed_ref", true
	}
	if strings.HasPrefix(ref, "spec:") {
		return "unsupported_ref", true
	}
	parts := strings.Split(ref, "-")
	if len(parts) < 2 {
		return "malformed_ref", true
	}
	if parts[1] == "SPEC" || strings.HasPrefix(ref, "V01-") {
		return "unsupported_ref", true
	}
	if ref != strings.ToUpper(ref) {
		return "malformed_ref", true
	}
	if !isCurrentRecordRefParts(parts) {
		return "malformed_ref", true
	}
	return "", false
}

func isCurrentRecordRefParts(parts []string) bool {
	switch parts[1] {
	case "ADR", "INV", "REQ", "WORK":
		return len(parts) == 4 && isThreeDigitSeq(parts[3])
	case "TASK":
		return len(parts) == 5 && isThreeDigitSeq(parts[3]) && isTwoDigitSeq(parts[4])
	default:
		return false
	}
}

func matchingCurrentRecords(idx *Index, ref string) []Record {
	matches := []Record{}
	for _, record := range idx.Records {
		if record.ID == ref && record.Kind != RecordKindSpec {
			matches = append(matches, record)
		}
	}
	return matches
}

func currentRecordDomain(ref, prefix string, kind RecordKind) string {
	if !strings.HasPrefix(ref, prefix) {
		return ""
	}
	parts := strings.Split(strings.TrimPrefix(ref, prefix), "-")
	if len(parts) < 3 {
		return ""
	}
	switch kind {
	case RecordKindDecision:
		if parts[0] == "ADR" && len(parts) == 3 && isThreeDigitSeq(parts[2]) {
			return parts[1]
		}
	case RecordKindInvestigation:
		if parts[0] == "INV" && len(parts) == 3 && isThreeDigitSeq(parts[2]) {
			return parts[1]
		}
	case RecordKindRequirement:
		if parts[0] == "REQ" && len(parts) == 3 && isThreeDigitSeq(parts[2]) {
			return parts[1]
		}
	case RecordKindWorkItem:
		if parts[0] == "WORK" && len(parts) == 3 && isThreeDigitSeq(parts[2]) {
			return parts[1]
		}
	case RecordKindTask:
		if parts[0] == "TASK" && len(parts) == 4 && isThreeDigitSeq(parts[2]) && isTwoDigitSeq(parts[3]) {
			return parts[1]
		}
	}
	return ""
}

func isThreeDigitSeq(value string) bool {
	return len(value) == 3 && isAllDigits(value)
}

func isTwoDigitSeq(value string) bool {
	return len(value) == 2 && isAllDigits(value)
}

func isAllDigits(value string) bool {
	for _, r := range value {
		if r < '0' || r > '9' {
			return false
		}
	}
	return value != ""
}

func slugifyRecordTitle(title string) string {
	var b strings.Builder
	previousSeparator := true
	for _, r := range title {
		switch {
		case r >= 'A' && r <= 'Z':
			b.WriteRune(r + ('a' - 'A'))
			previousSeparator = false
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
			previousSeparator = false
		case r >= '0' && r <= '9':
			b.WriteRune(r)
			previousSeparator = false
		default:
			if !previousSeparator {
				b.WriteByte('-')
				previousSeparator = true
			}
		}
	}
	return strings.Trim(b.String(), "-")
}
