package designrecords

import (
	"context"
	"fmt"
)

// ListRecords is a Phase 0 stub. It returns the current index records without
// applying filters, ordering, or limits.
//
// TODO(Phase 2): implement kind/status/id/id_range filtering plus id ordering
// and limit handling.
func ListRecords(ctx context.Context, idx *Index, req ListRecordsRequest) (ListRecordsResponse, error) {
	if err := ctx.Err(); err != nil {
		return ListRecordsResponse{}, err
	}
	if idx == nil {
		return ListRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	return ListRecordsResponse{Records: append([]Record(nil), idx.Records...)}, nil
}

// GetRecord is a Phase 0 lookup stub over the already-populated index.
//
// TODO(Phase 4): load raw Markdown body on demand when include_body is true.
func GetRecord(ctx context.Context, idx *Index, req GetRecordRequest) (GetRecordResponse, error) {
	if err := ctx.Err(); err != nil {
		return GetRecordResponse{}, err
	}
	if idx == nil {
		return GetRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	if req.ID == "" {
		return GetRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "id is required")
	}
	for _, record := range idx.Records {
		if record.ID == req.ID {
			return GetRecordResponse{Record: record}, nil
		}
	}
	return GetRecordResponse{}, newToolError(ErrorCodeRecordNotFound, fmt.Sprintf("record %s was not found", req.ID))
}

// ValidateRecords is a Phase 0 stub. It reports existing index diagnostics and
// treats an empty diagnostic set as ok.
//
// TODO(Phase 3): implement diagnostic generation and request filtering.
func ValidateRecords(ctx context.Context, idx *Index, req ValidateRecordsRequest) (ValidateRecordsResponse, error) {
	if err := ctx.Err(); err != nil {
		return ValidateRecordsResponse{}, err
	}
	if idx == nil {
		return ValidateRecordsResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	diagnostics := append([]Diagnostic(nil), idx.Diagnostics...)
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

// SuggestNextRecord is a P1 optional Phase 0 stub.
//
// TODO(Phase 6): suggest the next decision ADR id and path without writing files.
func SuggestNextRecord(ctx context.Context, idx *Index, req SuggestNextRecordRequest) (SuggestNextRecordResponse, error) {
	if err := ctx.Err(); err != nil {
		return SuggestNextRecordResponse{}, err
	}
	if idx == nil {
		return SuggestNextRecordResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	return SuggestNextRecordResponse{}, newToolError(ErrorCodeUnsupportedKind, "suggest_next_record is not implemented in Phase 0")
}
