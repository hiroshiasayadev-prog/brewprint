package designrecords

import (
	"encoding/json"
	"testing"
)

func TestDiagnosticCategoryConstants(t *testing.T) {
	got := []DiagnosticCategory{
		DiagnosticDuplicateID,
		DiagnosticFilenameIDMismatch,
		DiagnosticInvalidH1Title,
		DiagnosticInvalidStatusForKind,
		DiagnosticSpecStatusMismatch,
		DiagnosticMissingDependsOnTarget,
		DiagnosticMissingSupersedesTarget,
		DiagnosticInvalidMigratedToSpec,
		DiagnosticMissingRecordPath,
	}
	want := []string{
		"duplicate_id",
		"filename_id_mismatch",
		"invalid_h1_title",
		"invalid_status_for_kind",
		"spec_status_mismatch",
		"missing_depends_on_target",
		"missing_supersedes_target",
		"invalid_migrated_to_spec",
		"missing_record_path",
	}
	if len(got) != len(want) {
		t.Fatalf("category count = %d, want %d", len(got), len(want))
	}
	for i := range got {
		if string(got[i]) != want[i] {
			t.Fatalf("category[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestErrorCodeConstants(t *testing.T) {
	got := []ErrorCode{
		ErrorCodeRecordNotFound,
		ErrorCodeInvalidRequest,
		ErrorCodeUnsupportedKind,
		ErrorCodeIDRangeRequiresDecisionKind,
	}
	want := []string{
		"record_not_found",
		"invalid_request",
		"unsupported_kind",
		"id_range_requires_decision_kind",
	}
	if len(got) != len(want) {
		t.Fatalf("error code count = %d, want %d", len(got), len(want))
	}
	for i := range got {
		if string(got[i]) != want[i] {
			t.Fatalf("error code[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestToolErrorJSONShape(t *testing.T) {
	encoded, err := json.Marshal(struct {
		Error *ToolError `json:"error"`
	}{
		Error: newToolError(ErrorCodeRecordNotFound, "record ADR-999 was not found"),
	})
	if err != nil {
		t.Fatalf("Marshal ToolError: %v", err)
	}
	want := `{"error":{"code":"record_not_found","message":"record ADR-999 was not found"}}`
	if string(encoded) != want {
		t.Fatalf("ToolError JSON = %s, want %s", encoded, want)
	}
}

func TestToolErrorCodesAreNotDiagnosticCategories(t *testing.T) {
	diagnostics := map[string]bool{
		string(DiagnosticDuplicateID):             true,
		string(DiagnosticFilenameIDMismatch):      true,
		string(DiagnosticInvalidH1Title):          true,
		string(DiagnosticInvalidStatusForKind):    true,
		string(DiagnosticSpecStatusMismatch):      true,
		string(DiagnosticMissingDependsOnTarget):  true,
		string(DiagnosticMissingSupersedesTarget): true,
		string(DiagnosticInvalidMigratedToSpec):   true,
		string(DiagnosticMissingRecordPath):       true,
	}
	for _, code := range []ErrorCode{
		ErrorCodeRecordNotFound,
		ErrorCodeInvalidRequest,
		ErrorCodeUnsupportedKind,
		ErrorCodeIDRangeRequiresDecisionKind,
	} {
		if diagnostics[string(code)] {
			t.Fatalf("tool error code %q is also a diagnostic category", code)
		}
	}
}
