package designrecords

import (
	"context"
	"fmt"
	"sort"
	"strings"
)

type listRecordsScope struct {
	kind     RecordKind
	hasKind  bool
	status   RecordStatus
	id       string
	idRange  *recordIDRange
	order    string
	limit    int
	hasLimit bool
}

// ListRecords returns normalized design record metadata with MVP filters,
// deterministic ID ordering, and optional limit handling.
func ListRecords(ctx context.Context, idx *Index, req ListRecordsRequest) (ListRecordsResponse, error) {
	if err := ctx.Err(); err != nil {
		return ListRecordsResponse{}, err
	}
	if idx == nil {
		return ListRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	scope, err := newListRecordsScope(req)
	if err != nil {
		return ListRecordsResponse{}, err
	}

	records := make([]Record, 0, len(idx.Records))
	for _, record := range idx.Records {
		if scope.selectRecord(record) {
			records = append(records, record)
		}
	}
	sortRecordsByID(records, scope.order)
	if scope.hasLimit && len(records) > scope.limit {
		records = records[:scope.limit]
	}

	out := make([]ListedRecord, 0, len(records))
	for _, record := range records {
		out = append(out, listedRecord(record))
	}
	return ListRecordsResponse{Records: out}, nil
}

// GetRecord returns one exact ID match from the already-populated index.
// Duplicate IDs are reported by validate_records; this lookup deterministically
// returns the first record in index order and does not introduce another error.
func GetRecord(ctx context.Context, idx *Index, req GetRecordRequest) (GetRecordResponse, error) {
	if err := ctx.Err(); err != nil {
		return GetRecordResponse{}, err
	}
	if idx == nil {
		return GetRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	if strings.TrimSpace(req.ID) == "" {
		return GetRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "id is required")
	}
	for _, record := range idx.Records {
		if record.ID == req.ID {
			return GetRecordResponse{Record: getRecordResponseRecord(record, req.IncludeBody)}, nil
		}
	}
	return GetRecordResponse{}, newToolError(ErrorCodeRecordNotFound, fmt.Sprintf("record %s was not found", req.ID))
}

// GetRecords returns first-occurrence ordered results for explicitly requested
// record IDs. Missing IDs are item-level results; duplicate requested IDs are
// ignored after their first occurrence and reported as informational diagnostics.
func GetRecords(ctx context.Context, idx *Index, req GetRecordsRequest) (GetRecordsResponse, error) {
	if err := ctx.Err(); err != nil {
		return GetRecordsResponse{}, err
	}
	if idx == nil {
		return GetRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	if len(req.IDs) == 0 {
		return GetRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, "ids must be a non-empty array")
	}

	items := make([]GetRecordsItem, 0, len(req.IDs))
	firstIndexes := make(map[string]int, len(req.IDs))
	duplicateIndexes := make(map[string][]int)
	duplicateOrder := make([]string, 0)

	for index, id := range req.IDs {
		if _, seen := firstIndexes[id]; seen {
			if len(duplicateIndexes[id]) == 0 {
				duplicateOrder = append(duplicateOrder, id)
			}
			duplicateIndexes[id] = append(duplicateIndexes[id], index)
			continue
		}
		firstIndexes[id] = index

		item := GetRecordsItem{
			ID:          id,
			Diagnostics: []Diagnostic{},
		}
		for _, record := range idx.Records {
			if record.ID == id {
				recordResponse := getRecordResponseRecord(record, req.IncludeBody)
				item.RetrievalStatus = RetrievalStatusFound
				item.Record = &recordResponse
				break
			}
		}
		if item.Record == nil {
			item.RetrievalStatus = RetrievalStatusNotFound
			item.Diagnostics = []Diagnostic{{
				Category:    DiagnosticRecordNotFound,
				Severity:    DiagnosticSeverityError,
				RequestedID: id,
				Message:     fmt.Sprintf("record %s was not found", id),
			}}
		}
		items = append(items, item)
	}

	diagnostics := make([]Diagnostic, 0, len(duplicateOrder))
	for _, id := range duplicateOrder {
		firstIndex := firstIndexes[id]
		diagnostics = append(diagnostics, Diagnostic{
			Category:         DiagnosticDuplicateRequestedIDIgnored,
			Severity:         DiagnosticSeverityInfo,
			RequestedID:      id,
			FirstIndex:       &firstIndex,
			DuplicateIndexes: append([]int{}, duplicateIndexes[id]...),
			Message:          "duplicate requested record ID was ignored after its first occurrence",
		})
	}
	return GetRecordsResponse{Items: items, Diagnostics: diagnostics}, nil
}

// ValidateRecords checks the Phase 1 index materials and emits MVP validation
// diagnostics for the selected record scope.
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

// SuggestNextRecord suggests the next decision ADR ID and path from the
// already-populated index. It is read-only and does not create files.
func SuggestNextRecord(ctx context.Context, idx *Index, req SuggestNextRecordRequest) (SuggestNextRecordResponse, error) {
	if err := ctx.Err(); err != nil {
		return SuggestNextRecordResponse{}, err
	}
	if idx == nil {
		return SuggestNextRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	kind := RecordKind(strings.TrimSpace(string(req.Kind)))
	if kind == "" {
		return SuggestNextRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "kind is required")
	}
	if kind != RecordKindDecision {
		return SuggestNextRecordResponse{}, newToolError(ErrorCodeUnsupportedKind, fmt.Sprintf("suggest_next_record does not support kind %q", req.Kind))
	}
	title := strings.TrimSpace(req.Title)
	if title == "" {
		return SuggestNextRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "title is required")
	}

	maxNum := 0
	existingMaxID := ""
	for _, record := range idx.Records {
		if record.Kind != RecordKindDecision {
			continue
		}
		num, ok := decisionRecordNumber(record.ID)
		if !ok {
			continue
		}
		if num > maxNum {
			maxNum = num
			existingMaxID = fmt.Sprintf("ADR-%03d", num)
		}
	}

	nextNumber := maxNum + 1
	nextID := fmt.Sprintf("ADR-%03d", nextNumber)
	suggestedPath := suggestedDecisionRecordPath(nextNumber, title)
	return SuggestNextRecordResponse{
		Kind:          RecordKindDecision,
		Title:         title,
		NextID:        nextID,
		NextNumber:    nextNumber,
		SuggestedPath: suggestedPath,
		ExistingMaxID: existingMaxID,
	}, nil
}

func newListRecordsScope(req ListRecordsRequest) (listRecordsScope, error) {
	scope := listRecordsScope{order: "asc"}
	if req.Kind != "" {
		if !isListableRecordKind(req.Kind) {
			return scope, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unsupported kind %q", req.Kind))
		}
		scope.kind = req.Kind
		scope.hasKind = true
	}
	if req.IDRange != nil {
		parsed, err := parseRecordIDRange(*req.IDRange, req.Kind)
		if err != nil {
			return scope, err
		}
		scope.kind = parsed.kind
		scope.hasKind = true
		scope.idRange = parsed
	}
	if req.OrderBy != "" && req.OrderBy != "id" {
		return scope, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unsupported order_by %q", req.OrderBy))
	}
	if req.Order != "" {
		if req.Order != "asc" && req.Order != "desc" {
			return scope, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unsupported order %q", req.Order))
		}
		scope.order = req.Order
	}
	if req.Limit != nil && *req.Limit <= 0 {
		return scope, newToolError(ErrorCodeInvalidRequest, "limit must be greater than zero")
	}
	if req.Limit != nil {
		scope.limit = *req.Limit
		scope.hasLimit = true
	}
	scope.status = req.Status
	scope.id = req.ID
	return scope, nil
}

func (s listRecordsScope) selectRecord(record Record) bool {
	if s.hasKind && record.Kind != s.kind {
		return false
	}
	if s.status != "" && record.Status != s.status {
		return false
	}
	if s.id != "" && record.ID != s.id {
		return false
	}
	if s.idRange == nil {
		return true
	}
	return s.idRange.containsRecord(record)
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

// ResolveReference resolves an MVP canonical reference candidate without
// selecting an arbitrary target when the index is ambiguous.
func ResolveReference(ctx context.Context, idx *Index, req ResolveReferenceRequest) (ResolveReferenceResponse, error) {
	if err := ctx.Err(); err != nil {
		return ResolveReferenceResponse{}, err
	}
	if idx == nil {
		return ResolveReferenceResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	if req.Ref == "" {
		return ResolveReferenceResponse{}, newToolError(ErrorCodeInvalidRequest, "ref is required")
	}
	return resolveReference(idx, req.Ref), nil
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

func suggestedDecisionRecordPath(num int, title string) string {
	prefix := fmt.Sprintf("v01/records/adr/V01-ADR-%03d", num)
	slug := slugifyRecordTitle(title)
	if slug == "" {
		return prefix + ".md"
	}
	return prefix + "-" + slug + ".md"
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
