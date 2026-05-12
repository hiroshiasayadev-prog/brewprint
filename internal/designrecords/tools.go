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
	idRange  *numericIDRange
	order    string
	limit    int
	hasLimit bool
}

// ListRecords returns normalized ADR/spec record metadata with MVP filters,
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
		if req.Kind != RecordKindDecision && req.Kind != RecordKindSpec {
			return scope, newToolError(ErrorCodeInvalidRequest, fmt.Sprintf("unsupported kind %q", req.Kind))
		}
		scope.kind = req.Kind
		scope.hasKind = true
	}
	if req.IDRange != nil {
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
	if record.Kind != RecordKindDecision {
		return false
	}
	num, ok := decisionRecordNumber(record.ID)
	return ok && s.idRange.contains(num)
}

func sortRecordsByID(records []Record, order string) {
	sort.SliceStable(records, func(i, j int) bool {
		return compareRecordsByID(records[i], records[j], order) < 0
	})
}

func compareRecordsByID(a, b Record, order string) int {
	if rankA, rankB := recordKindSortRank(a.Kind), recordKindSortRank(b.Kind); rankA != rankB {
		return rankA - rankB
	}
	cmp := compareRecordID(a, b)
	if order == "desc" {
		return -cmp
	}
	return cmp
}

func recordKindSortRank(kind RecordKind) int {
	switch kind {
	case RecordKindDecision:
		return 0
	case RecordKindSpec:
		return 1
	default:
		return 2
	}
}

func compareRecordID(a, b Record) int {
	if a.Kind == RecordKindDecision && b.Kind == RecordKindDecision {
		if numA, okA := decisionRecordNumber(a.ID); okA {
			if numB, okB := decisionRecordNumber(b.ID); okB {
				switch {
				case numA < numB:
					return -1
				case numA > numB:
					return 1
				}
			}
		}
	}
	switch {
	case a.ID < b.ID:
		return -1
	case a.ID > b.ID:
		return 1
	default:
		return 0
	}
}

func listedRecord(record Record) ListedRecord {
	return ListedRecord{
		ID:             record.ID,
		Kind:           record.Kind,
		Title:          record.Title,
		Status:         record.Status,
		Path:           record.Path,
		DependsOn:      append([]string{}, record.DependsOn...),
		Supersedes:     append([]string{}, record.Supersedes...),
		MigratedToSpec: record.MigratedToSpec,
	}
}

func getRecordResponseRecord(record Record, includeBody bool) GetRecordRecord {
	out := GetRecordRecord{
		ID:             record.ID,
		Kind:           record.Kind,
		Title:          record.Title,
		Status:         record.Status,
		Path:           record.Path,
		DependsOn:      append([]string{}, record.DependsOn...),
		Supersedes:     append([]string{}, record.Supersedes...),
		MigratedToSpec: record.MigratedToSpec,
		Headings:       append([]Heading{}, record.Headings...),
	}
	if includeBody {
		body := record.RawBody
		out.Body = &body
	}
	return out
}

func suggestedDecisionRecordPath(num int, title string) string {
	prefix := fmt.Sprintf("docs/adr/%03d", num)
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
