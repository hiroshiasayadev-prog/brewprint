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

// ── Current read model type tests ─────────────────────────────────────────────

func TestCurrentGetRecordsResponseShape(t *testing.T) {
	resp := CurrentGetRecordsResponse{
		Records: []CurrentGetRecordsRecord{
			{Ref: "DRMCP-REQ-MCP-901", Kind: RecordKindRequirement, Status: RecordStatusAccepted},
		},
	}
	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal CurrentGetRecordsResponse: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if _, ok := decoded["records"]; !ok {
		t.Fatal("CurrentGetRecordsResponse JSON must have 'records' key, not 'items'")
	}
	if _, ok := decoded["items"]; ok {
		t.Fatal("CurrentGetRecordsResponse JSON must not have 'items' key")
	}
	if _, ok := decoded["warnings"]; ok {
		t.Fatal("warnings must be absent when empty")
	}
}

func TestCurrentGetRecordsResponseWithWarnings(t *testing.T) {
	resp := CurrentGetRecordsResponse{
		Records: []CurrentGetRecordsRecord{},
		Warnings: []OperationWarning{
			{Category: "not_found", Message: "ref not found", Ref: "DRMCP-REQ-MCP-999"},
		},
	}
	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if _, ok := decoded["warnings"]; !ok {
		t.Fatal("warnings must be present when non-empty")
	}
}

func TestCurrentValidateRecordsResponseShape(t *testing.T) {
	resp := CurrentValidateRecordsResponse{
		OK:          true,
		Scope:       "all",
		Summary:     ValidationSubjectSummary{Total: 5, Invalid: 0},
		Diagnostics: []Diagnostic{},
	}
	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal CurrentValidateRecordsResponse: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if ok, _ := decoded["ok"].(bool); !ok {
		t.Fatal("ok must be a boolean true")
	}
	if _, ok := decoded["summary"]; !ok {
		t.Fatal("summary must be present")
	}
	if _, ok := decoded["scope"]; !ok {
		t.Fatal("scope must be present")
	}
}

func TestCurrentValidateRecordsResponseOKFalseOnError(t *testing.T) {
	resp := CurrentValidateRecordsResponse{
		OK:      false,
		Scope:   "all",
		Summary: ValidationSubjectSummary{Total: 2, Invalid: 1},
		Diagnostics: []Diagnostic{
			{Category: DiagnosticMissingRequiredMetadata, Severity: DiagnosticSeverityError, Message: "missing parent"},
		},
	}
	if resp.OK {
		t.Fatal("OK must be false when diagnostics contain an error severity entry")
	}
}

func TestDiagnosticLocationShape(t *testing.T) {
	loc := DiagnosticLocation{
		SourceScope:  "current",
		RecordsRoot:  "current/product/records",
		Path:         "spec/fixture-baseline/overview.md",
		AppNamespace: "product",
	}
	encoded, err := json.Marshal(loc)
	if err != nil {
		t.Fatalf("Marshal DiagnosticLocation: %v", err)
	}
	want := `{"source_scope":"current","records_root":"current/product/records","path":"spec/fixture-baseline/overview.md","app_namespace":"product"}`
	if string(encoded) != want {
		t.Fatalf("DiagnosticLocation JSON = %s, want %s", encoded, want)
	}
}

func TestDiagnosticLocationAppNamespaceOmittedWhenEmpty(t *testing.T) {
	loc := DiagnosticLocation{
		SourceScope: "legacy",
		RecordsRoot: "legacy/v01/records",
		Path:        "adr/V01-ADR-901-fixture.md",
	}
	encoded, err := json.Marshal(loc)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if _, ok := decoded["app_namespace"]; ok {
		t.Fatal("app_namespace must be omitted when empty")
	}
}

func TestCurrentConflictShape(t *testing.T) {
	conflict := CurrentConflict{
		Ref: "DRMCP-REQ-MCP-990",
		Sources: []string{
			"arrangements/duplicate-current/root-a/records/requirements/mcp/DRMCP-REQ-MCP-990-duplicate-current-a.md",
			"arrangements/duplicate-current/root-b/records/requirements/mcp/DRMCP-REQ-MCP-990-duplicate-current-b.md",
		},
	}
	encoded, err := json.Marshal(conflict)
	if err != nil {
		t.Fatalf("Marshal CurrentConflict: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if decoded["ref"] != "DRMCP-REQ-MCP-990" {
		t.Fatalf("ref = %v, want DRMCP-REQ-MCP-990", decoded["ref"])
	}
	sources, _ := decoded["sources"].([]any)
	if len(sources) != 2 {
		t.Fatalf("sources count = %d, want 2", len(sources))
	}
}

// ── F-MAJ-01: CurrentListedRecord nullable JSON contract ─────────────────────

func TestCurrentListedRecordJSONAllFields(t *testing.T) {
	title := "Current read fixture baseline"
	status := RecordStatusAccepted
	date := "2026-06-28"
	rec := CurrentListedRecord{
		Ref:    "DRMCP-REQ-MCP-901",
		Kind:   RecordKindRequirement,
		Title:  &title,
		Status: &status,
		Date:   &date,
	}
	encoded, err := json.Marshal(rec)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if decoded["ref"] != "DRMCP-REQ-MCP-901" {
		t.Fatalf("ref = %v", decoded["ref"])
	}
	if decoded["title"] != "Current read fixture baseline" {
		t.Fatalf("title = %v", decoded["title"])
	}
	if decoded["status"] != "accepted" {
		t.Fatalf("status = %v", decoded["status"])
	}
	if decoded["date"] != "2026-06-28" {
		t.Fatalf("date = %v", decoded["date"])
	}
}

func TestCurrentListedRecordJSONNullFields(t *testing.T) {
	rec := CurrentListedRecord{
		Ref:    "DRMCP-REQ-MCP-901",
		Title:  nil,
		Status: nil,
		Date:   nil,
	}
	encoded, err := json.Marshal(rec)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	for _, key := range []string{"title", "status", "date"} {
		val, present := decoded[key]
		if !present {
			t.Fatalf("key %q must be present (got omitted)", key)
		}
		if val != nil {
			t.Fatalf("key %q must be null, got %v", key, val)
		}
	}
}

func TestCurrentListedRecordJSONKeyPresence(t *testing.T) {
	title := "Some title"
	rec := CurrentListedRecord{
		Ref:    "DRMCP-REQ-MCP-902",
		Title:  &title,
		Status: nil,
		Date:   nil,
	}
	encoded, err := json.Marshal(rec)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	for _, key := range []string{"ref", "title", "status", "date"} {
		if _, present := decoded[key]; !present {
			t.Fatalf("key %q must be present even when nil", key)
		}
	}
}

// ── F-MAJ-02: Diagnostic.Location ────────────────────────────────────────────

func TestDiagnosticWithLocationJSON(t *testing.T) {
	d := Diagnostic{
		Category: DiagnosticMissingRequiredMetadata,
		Severity: DiagnosticSeverityError,
		Message:  "missing parent",
		Location: &DiagnosticLocation{
			SourceScope:  "current",
			RecordsRoot:  "current/product/records",
			Path:         "spec/invalid-source/missing-parent.md",
			AppNamespace: "product",
		},
	}
	encoded, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal Diagnostic with Location: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	loc, ok := decoded["location"].(map[string]any)
	if !ok {
		t.Fatalf("location must be a JSON object; got %T (%v)", decoded["location"], decoded["location"])
	}
	if loc["source_scope"] != "current" {
		t.Fatalf("location.source_scope = %v", loc["source_scope"])
	}
	if loc["records_root"] != "current/product/records" {
		t.Fatalf("location.records_root = %v", loc["records_root"])
	}
	if loc["path"] != "spec/invalid-source/missing-parent.md" {
		t.Fatalf("location.path = %v", loc["path"])
	}
	if loc["app_namespace"] != "product" {
		t.Fatalf("location.app_namespace = %v", loc["app_namespace"])
	}
}

func TestDiagnosticWithoutLocationJSON(t *testing.T) {
	d := Diagnostic{
		Category: DiagnosticDuplicateID,
		Severity: DiagnosticSeverityError,
		Message:  "duplicate",
	}
	encoded, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if _, present := decoded["location"]; present {
		t.Fatal("location must be absent when nil")
	}
}

func TestDiagnosticExistingFieldsUnchangedWithLocation(t *testing.T) {
	recordID := "DRMCP-REQ-MCP-901"
	d := Diagnostic{
		Category: DiagnosticMissingRequiredMetadata,
		Severity: DiagnosticSeverityError,
		RecordID: recordID,
		Message:  "missing field",
		Location: &DiagnosticLocation{
			SourceScope: "current",
			RecordsRoot: "current/drmcp/records",
			Path:        "requirements/mcp/DRMCP-REQ-MCP-901.md",
		},
	}
	encoded, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if decoded["category"] != "missing_required_metadata" {
		t.Fatalf("category = %v", decoded["category"])
	}
	if decoded["severity"] != "error" {
		t.Fatalf("severity = %v", decoded["severity"])
	}
	if decoded["record_id"] != recordID {
		t.Fatalf("record_id = %v", decoded["record_id"])
	}
	if decoded["message"] != "missing field" {
		t.Fatalf("message = %v", decoded["message"])
	}
	if _, present := decoded["location"]; !present {
		t.Fatal("location must be present")
	}
}

func TestDiagnosticLegacyPathFieldPreserved(t *testing.T) {
	d := Diagnostic{
		Category: DiagnosticFilenameIDMismatch,
		Severity: DiagnosticSeverityError,
		Path:     "current/product/records/adr/spec/PRODUCT-ADR-SPEC-901.md",
		Message:  "mismatch",
	}
	encoded, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if decoded["path"] != "current/product/records/adr/spec/PRODUCT-ADR-SPEC-901.md" {
		t.Fatalf("path = %v", decoded["path"])
	}
}

// ── F-MIN-01: CurrentListRecordsRequest / Response ───────────────────────────

func TestCurrentListRecordsRequestJSONShape(t *testing.T) {
	limit := 20
	req := CurrentListRecordsRequest{
		AppNamespace: "drmcp",
		Kind:         RecordKindRequirement,
		Domain:       "mcp",
		Status:       RecordStatusAccepted,
		Order:        "desc",
		Limit:        &limit,
	}
	encoded, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Marshal CurrentListRecordsRequest: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if decoded["app_namespace"] != "drmcp" {
		t.Fatalf("app_namespace = %v", decoded["app_namespace"])
	}
	if decoded["kind"] != "requirement" {
		t.Fatalf("kind = %v", decoded["kind"])
	}
	if decoded["domain"] != "mcp" {
		t.Fatalf("domain = %v", decoded["domain"])
	}
}

func TestCurrentListRecordsRequestOptionalOmission(t *testing.T) {
	req := CurrentListRecordsRequest{
		AppNamespace: "drmcp",
		Kind:         RecordKindTask,
		Domain:       "mcp",
	}
	encoded, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	for _, key := range []string{"status", "order", "limit"} {
		if _, present := decoded[key]; present {
			t.Fatalf("optional field %q must be absent when zero", key)
		}
	}
}

func TestCurrentListRecordsResponseJSONShape(t *testing.T) {
	title := "Current read fixture baseline"
	status := RecordStatusAccepted
	date := "2026-06-28"
	resp := CurrentListRecordsResponse{
		Records: []CurrentListedRecord{
			{
				Ref:    "DRMCP-REQ-MCP-901",
				Kind:   RecordKindRequirement,
				Title:  &title,
				Status: &status,
				Date:   &date,
			},
		},
		HasMore:  false,
		Warnings: []OperationWarning{},
	}
	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal CurrentListRecordsResponse: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if _, ok := decoded["records"]; !ok {
		t.Fatal("records must be present")
	}
	if _, ok := decoded["has_more"]; !ok {
		t.Fatal("has_more must be present")
	}
	if _, ok := decoded["warnings"]; !ok {
		t.Fatal("warnings must be present even when empty")
	}
	if hasMore, _ := decoded["has_more"].(bool); hasMore {
		t.Fatal("has_more must be false")
	}
}

func TestCurrentListRecordsResponseWarningsAlwaysPresent(t *testing.T) {
	resp := CurrentListRecordsResponse{
		Records:  []CurrentListedRecord{},
		HasMore:  false,
		Warnings: []OperationWarning{},
	}
	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if _, present := decoded["warnings"]; !present {
		t.Fatal("warnings must be present even when empty (zero-match case)")
	}
	warnings, _ := decoded["warnings"].([]any)
	if len(warnings) != 0 {
		t.Fatalf("warnings must be empty array, got %v", decoded["warnings"])
	}
}

func TestCurrentListRecordsResponseNullableMetadataCombination(t *testing.T) {
	resp := CurrentListRecordsResponse{
		Records: []CurrentListedRecord{
			{Ref: "DRMCP-REQ-MCP-901", Title: nil, Status: nil, Date: nil},
		},
		HasMore:  true,
		Warnings: []OperationWarning{{Category: "missing_title", Message: "title is null"}},
	}
	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	records, _ := decoded["records"].([]any)
	if len(records) != 1 {
		t.Fatalf("records count = %d", len(records))
	}
	entry, _ := records[0].(map[string]any)
	for _, key := range []string{"title", "status", "date"} {
		val, present := entry[key]
		if !present {
			t.Fatalf("records[0].%s must be present (not omitted)", key)
		}
		if val != nil {
			t.Fatalf("records[0].%s must be null, got %v", key, val)
		}
	}
	if hasMore, _ := decoded["has_more"].(bool); !hasMore {
		t.Fatal("has_more must be true")
	}
	warnings, _ := decoded["warnings"].([]any)
	if len(warnings) != 1 {
		t.Fatalf("warnings count = %d", len(warnings))
	}
}
