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
		DiagnosticInvalidWorkflowID,
		DiagnosticInvalidStatusForKind,
		DiagnosticSpecStatusMismatch,
		DiagnosticMissingDependsOnTarget,
		DiagnosticMissingSupersedesTarget,
		DiagnosticInvalidMigratedToSpec,
		DiagnosticMissingRecordPath,
		DiagnosticInvalidSemanticRefDeclaration,
		DiagnosticMissingSectionTarget,
		DiagnosticAmbiguousSectionTarget,
		DiagnosticDuplicateSemanticRef,
		DiagnosticUnresolvedSourceRef,
		DiagnosticUnresolvedFollowUpResult,
		DiagnosticUnresolvedFollowUpCandidate,
		DiagnosticNoncanonicalSourceRef,
		DiagnosticNoncanonicalFollowUpResult,
		DiagnosticNoncanonicalFollowUpCandidate,
		DiagnosticUnsupportedReference,
		DiagnosticUnresolvedReference,
		DiagnosticAmbiguousReference,
		DiagnosticUnresolvedWorkflowRelation,
		DiagnosticInvalidWorkflowRelationTarget,
		DiagnosticWorkflowRelationMismatch,
		DiagnosticWorkflowSourceReqMismatch,
		DiagnosticMissingRequiredMetadata,
		DiagnosticEmptyRequiredMetadata,
		DiagnosticInvalidMetadataValue,
		DiagnosticRecordNotFound,
		DiagnosticDuplicateRequestedIDIgnored,
	}
	want := []string{
		"duplicate_id",
		"filename_id_mismatch",
		"invalid_h1_title",
		"invalid_workflow_id",
		"invalid_status_for_kind",
		"spec_status_mismatch",
		"missing_depends_on_target",
		"missing_supersedes_target",
		"invalid_migrated_to_spec",
		"missing_record_path",
		"invalid_semantic_ref_declaration",
		"missing_section_target",
		"ambiguous_section_target",
		"duplicate_semantic_ref",
		"unresolved_source_ref",
		"unresolved_follow_up_result",
		"unresolved_follow_up_candidate",
		"noncanonical_source_ref",
		"noncanonical_follow_up_result",
		"noncanonical_follow_up_candidate",
		"unsupported_reference",
		"unresolved_reference",
		"ambiguous_reference",
		"unresolved_workflow_relation",
		"invalid_workflow_relation_target",
		"workflow_relation_mismatch",
		"workflow_source_requirement_mismatch",
		"missing_required_metadata",
		"empty_required_metadata",
		"invalid_metadata_value",
		"record_not_found",
		"duplicate_requested_id_ignored",
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
		ErrorCodeGuideNotFound,
		ErrorCodeInvalidRequest,
		ErrorCodeUnsupportedKind,
		ErrorCodeInvalidIDRange,
		ErrorCodeIDRangeRequiresDecisionKind,
	}
	want := []string{
		"record_not_found",
		"guide_not_found",
		"invalid_request",
		"unsupported_kind",
		"invalid_id_range",
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
		string(DiagnosticDuplicateID):                   true,
		string(DiagnosticFilenameIDMismatch):            true,
		string(DiagnosticInvalidH1Title):                true,
		string(DiagnosticInvalidWorkflowID):             true,
		string(DiagnosticInvalidStatusForKind):          true,
		string(DiagnosticSpecStatusMismatch):            true,
		string(DiagnosticMissingDependsOnTarget):        true,
		string(DiagnosticMissingSupersedesTarget):       true,
		string(DiagnosticInvalidMigratedToSpec):         true,
		string(DiagnosticMissingRecordPath):             true,
		string(DiagnosticInvalidSemanticRefDeclaration): true,
		string(DiagnosticMissingSectionTarget):          true,
		string(DiagnosticAmbiguousSectionTarget):        true,
		string(DiagnosticDuplicateSemanticRef):          true,
		string(DiagnosticUnresolvedSourceRef):           true,
		string(DiagnosticUnresolvedFollowUpResult):      true,
		string(DiagnosticUnresolvedFollowUpCandidate):   true,
		string(DiagnosticNoncanonicalSourceRef):         true,
		string(DiagnosticNoncanonicalFollowUpResult):    true,
		string(DiagnosticNoncanonicalFollowUpCandidate): true,
		string(DiagnosticUnsupportedReference):          true,
		string(DiagnosticUnresolvedReference):           true,
		string(DiagnosticAmbiguousReference):            true,
		string(DiagnosticUnresolvedWorkflowRelation):    true,
		string(DiagnosticInvalidWorkflowRelationTarget): true,
		string(DiagnosticWorkflowRelationMismatch):      true,
		string(DiagnosticWorkflowSourceReqMismatch):     true,
		string(DiagnosticMissingRequiredMetadata):       true,
		string(DiagnosticEmptyRequiredMetadata):         true,
		string(DiagnosticInvalidMetadataValue):          true,
		string(DiagnosticDuplicateRequestedIDIgnored):   true,
	}
	for _, code := range []ErrorCode{
		ErrorCodeRecordNotFound,
		ErrorCodeGuideNotFound,
		ErrorCodeInvalidRequest,
		ErrorCodeUnsupportedKind,
		ErrorCodeIDRangeRequiresDecisionKind,
	} {
		if diagnostics[string(code)] {
			t.Fatalf("tool error code %q is also a diagnostic category", code)
		}
	}
}
