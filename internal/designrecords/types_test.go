package designrecords

import "testing"

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
