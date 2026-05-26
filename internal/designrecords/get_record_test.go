package designrecords

import (
	"context"
	"encoding/json"
	"testing"
)

func TestGetRecordBasicDecisionAndSpecResponseShape(t *testing.T) {
	root := t.TempDir()
	adrRaw := "# 076: Design Records MCP\n\n" +
		"- **status**: accepted\n" +
		"- **depends_on**: ADR-050, ADR-068\n" +
		"- **supersedes**: ADR-001\n" +
		"- **migrated_to_spec**:\n\n" +
		"## Background\n\n" +
		"```md\n" +
		"## Not a heading\n" +
		"```\n\n" +
		"## Decision\n"
	writeTestFile(t, root, "docs/adr/076-design-records-mcp.md", adrRaw)
	specRaw := "---\n" +
		"status: draft\n" +
		"design_record:\n" +
		"  id: SPEC-design-records-mcp-tools\n" +
		"  kind: spec\n" +
		"  status: draft\n" +
		"  depends_on:\n" +
		"    - ADR-076\n" +
		"---\n\n" +
		"# Design Records MCP tools\n\n" +
		"## get_record\n"
	writeTestFile(t, root, "docs/spec/design-records-mcp/tools.md", specRaw)
	investigationRaw := "# INV-DOCS-001: Test investigation\n\n" +
		"- **status**: concluded\n" +
		"- **date**: 2026-05-19\n" +
		"- **trigger**: ADR-076\n" +
		"- **scope**: test\n" +
		"- **non_scope**: none\n" +
		"- **source_refs**:\n" +
		"  - ADR-076\n" +
		"- **follow_up_candidates**:\n" +
		"  - SPEC-design-records-mcp-tools\n" +
		"## Result\n"
	writeTestFile(t, root, "docs/investigations/docs/INV-DOCS-001-test.md", investigationRaw)
	writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-003-test.md", "# REQ-MCP-003: Test requirement\n- **id**: REQ-MCP-003\n- **status**: accepted\n- **date**: 2026-05-25\n- **source_refs**:\n  - ADR-076\n- **work_items**:\n  - WORK-MCP-003\n")
	writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-003-test.md", "# WORK-MCP-003: Test work item\n- **id**: WORK-MCP-003\n- **status**: implementation_pending\n- **date**: 2026-05-26\n- **source_requirement**: REQ-MCP-003\n- **impact_refs**:\n  - ADR-076\n- **tasks**:\n  - TASK-MCP-003-01\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-003-01-test.md", "# TASK-MCP-003-01: Test task\n- **id**: TASK-MCP-003-01\n- **status**: todo\n- **date**: 2026-05-26\n- **work_item**: WORK-MCP-003\n- **source_requirement**: REQ-MCP-003\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n  - implementation\n")
	idx := buildTestIndex(t, root)

	adrResp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "ADR-076"})
	if err != nil {
		t.Fatalf("GetRecord ADR-076: %v", err)
	}
	adr := adrResp.Record
	if adr.ID != "ADR-076" || adr.Kind != RecordKindDecision || adr.Title != "Design Records MCP" || adr.Status != RecordStatusAccepted || adr.Path != "docs/adr/076-design-records-mcp.md" {
		t.Fatalf("ADR metadata = %#v", adr)
	}
	if adr.Decision == nil {
		t.Fatalf("Decision detail missing: %#v", adr)
	}
	assertStrings(t, adr.Decision.DependsOn, []string{"ADR-050", "ADR-068"})
	assertStrings(t, adr.Decision.Supersedes, []string{"ADR-001"})
	if adr.Decision.MigratedToSpec != nil {
		t.Fatalf("MigratedToSpec = %#v, want nil", adr.Decision.MigratedToSpec)
	}
	assertHeadings(t, adr.Headings, []Heading{
		{Level: 1, Text: "076: Design Records MCP"},
		{Level: 2, Text: "Background"},
		{Level: 2, Text: "Decision"},
	})
	if adr.Body != nil {
		t.Fatalf("body = %#v, want nil when include_body is omitted", *adr.Body)
	}
	assertGetRecordJSONShape(t, adrResp, false)

	specResp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "SPEC-design-records-mcp-tools"})
	if err != nil {
		t.Fatalf("GetRecord spec: %v", err)
	}
	spec := specResp.Record
	if spec.ID != "SPEC-design-records-mcp-tools" || spec.Kind != RecordKindSpec || spec.Title != "Design Records MCP tools" || spec.Status != RecordStatusDraft || spec.Path != "docs/spec/design-records-mcp/tools.md" {
		t.Fatalf("spec metadata = %#v", spec)
	}
	if spec.Spec == nil {
		t.Fatalf("Spec detail missing: %#v", spec)
	}
	assertStrings(t, spec.Spec.DependsOn, []string{"ADR-076"})
	assertHeadings(t, spec.Headings, []Heading{
		{Level: 1, Text: "Design Records MCP tools"},
		{Level: 2, Text: "get_record"},
	})
	if spec.Body != nil {
		t.Fatalf("spec body = %#v, want nil when include_body is omitted", *spec.Body)
	}

	invResp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "INV-DOCS-001"})
	if err != nil {
		t.Fatalf("GetRecord investigation: %v", err)
	}
	inv := invResp.Record
	if inv.Kind != RecordKindInvestigation || inv.Investigation == nil || inv.Investigation.Trigger != "ADR-076" {
		t.Fatalf("investigation response = %#v", inv)
	}
	assertStrings(t, inv.Investigation.SourceRefs, []string{"ADR-076"})
	assertStrings(t, inv.Investigation.FollowUpCandidates, []string{"SPEC-design-records-mcp-tools"})
	assertGetRecordJSONShape(t, invResp, false)

	reqResp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "REQ-MCP-003"})
	if err != nil {
		t.Fatalf("GetRecord requirement: %v", err)
	}
	if reqResp.Record.Kind != RecordKindRequirement || reqResp.Record.Requirement == nil {
		t.Fatalf("requirement response = %#v", reqResp.Record)
	}
	assertStrings(t, reqResp.Record.Requirement.SourceRefs, []string{"ADR-076"})
	assertStrings(t, reqResp.Record.Requirement.WorkItems, []string{"WORK-MCP-003"})
	assertGetRecordJSONShape(t, reqResp, false)

	workResp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "WORK-MCP-003"})
	if err != nil {
		t.Fatalf("GetRecord work item: %v", err)
	}
	if workResp.Record.Kind != RecordKindWorkItem || workResp.Record.WorkItem == nil || workResp.Record.WorkItem.SourceRequirement != "REQ-MCP-003" {
		t.Fatalf("work item response = %#v", workResp.Record)
	}
	assertStrings(t, workResp.Record.WorkItem.Tasks, []string{"TASK-MCP-003-01"})
	assertGetRecordJSONShape(t, workResp, false)

	taskResp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "TASK-MCP-003-01"})
	if err != nil {
		t.Fatalf("GetRecord task: %v", err)
	}
	if taskResp.Record.Kind != RecordKindTask || taskResp.Record.Task == nil || taskResp.Record.Task.WorkItem != "WORK-MCP-003" {
		t.Fatalf("task response = %#v", taskResp.Record)
	}
	assertStrings(t, taskResp.Record.Task.DependsOn, []string{})
	assertStrings(t, taskResp.Record.Task.Outputs, []string{"implementation"})
	assertGetRecordJSONShape(t, taskResp, false)
}

func TestGetRecordIncludeBodyReturnsRawMarkdown(t *testing.T) {
	root := t.TempDir()
	specRaw := "---\r\n" +
		"status: draft\r\n" +
		"design_record:\r\n" +
		"  id: SPEC-raw-body\r\n" +
		"  kind: spec\r\n" +
		"  status: draft\r\n" +
		"  depends_on: []\r\n" +
		"---\r\n" +
		"# Raw body spec\r\n" +
		"\r\n" +
		"## Section\r\n" +
		"Original newline style is CRLF.\r\n"
	writeTestFile(t, root, "docs/spec/raw-body.md", specRaw)
	idx := buildTestIndex(t, root)

	withoutBody, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "SPEC-raw-body", IncludeBody: false})
	if err != nil {
		t.Fatalf("GetRecord without body: %v", err)
	}
	if withoutBody.Record.Body != nil {
		t.Fatalf("body = %#v, want nil when include_body=false", *withoutBody.Record.Body)
	}
	assertHeadings(t, withoutBody.Record.Headings, []Heading{
		{Level: 1, Text: "Raw body spec"},
		{Level: 2, Text: "Section"},
	})
	assertGetRecordJSONShape(t, withoutBody, false)

	withBody, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "SPEC-raw-body", IncludeBody: true})
	if err != nil {
		t.Fatalf("GetRecord with body: %v", err)
	}
	if withBody.Record.Body == nil {
		t.Fatal("body = nil, want raw Markdown")
	}
	if *withBody.Record.Body != specRaw {
		t.Fatalf("body was changed:\ngot  %#v\nwant %#v", *withBody.Record.Body, specRaw)
	}
	assertHeadings(t, withBody.Record.Headings, []Heading{
		{Level: 1, Text: "Raw body spec"},
		{Level: 2, Text: "Section"},
	})
	assertGetRecordJSONShape(t, withBody, true)
}

func TestGetRecordErrors(t *testing.T) {
	idx := &Index{Records: []Record{
		{ID: "ADR-001", Kind: RecordKindDecision},
		{ID: "ADR-076", Kind: RecordKindDecision},
	}}

	assertGetRecordErrorCode(t, idx, GetRecordRequest{}, ErrorCodeInvalidRequest)
	assertGetRecordErrorCode(t, idx, GetRecordRequest{ID: ""}, ErrorCodeInvalidRequest)
	assertGetRecordErrorCode(t, idx, GetRecordRequest{ID: " \t"}, ErrorCodeInvalidRequest)
	assertGetRecordErrorCode(t, idx, GetRecordRequest{ID: "ADR-999"}, ErrorCodeRecordNotFound)
	assertGetRecordErrorCode(t, idx, GetRecordRequest{ID: "adr-076"}, ErrorCodeRecordNotFound)
	assertGetRecordErrorCode(t, idx, GetRecordRequest{ID: "ADR-076 "}, ErrorCodeRecordNotFound)
	assertGetRecordErrorCode(t, nil, GetRecordRequest{ID: "ADR-001"}, ErrorCodeInvalidRequest)
}

func TestGetRecordDuplicateIDReturnsFirstIndexRecord(t *testing.T) {
	idx := &Index{Records: []Record{
		{ID: "ADR-001", Kind: RecordKindDecision, Title: "First", Path: "docs/adr/001-first.md", RawBody: "# 001: First\n"},
		{ID: "ADR-001", Kind: RecordKindDecision, Title: "Second", Path: "docs/adr/001-second.md", RawBody: "# 001: Second\n"},
	}}

	resp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "ADR-001", IncludeBody: true})
	if err != nil {
		t.Fatalf("GetRecord duplicate: %v", err)
	}
	if resp.Record.Title != "First" || resp.Record.Path != "docs/adr/001-first.md" {
		t.Fatalf("duplicate lookup returned %#v, want first index record", resp.Record)
	}
	if resp.Record.Body == nil || *resp.Record.Body != "# 001: First\n" {
		t.Fatalf("duplicate body = %#v, want first raw body", resp.Record.Body)
	}
}

func TestGetRecordRepositoryBootstrapRecords(t *testing.T) {
	root := findRepoRoot(t)
	cfg, err := NewConfig(root)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}

	adrResp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "ADR-076", IncludeBody: true})
	if err != nil {
		t.Fatalf("GetRecord ADR-076: %v", err)
	}
	if adrResp.Record.Kind != RecordKindDecision || adrResp.Record.Title != "Design Records MCP" || adrResp.Record.Path != "docs/adr/076-design-records-mcp.md" {
		t.Fatalf("ADR-076 = %#v", adrResp.Record)
	}
	indexedADR := findRecord(idx.Records, "ADR-076")
	if indexedADR == nil {
		t.Fatal("ADR-076 missing from index")
	}
	if adrResp.Record.Body == nil || *adrResp.Record.Body == "" || *adrResp.Record.Body != indexedADR.RawBody {
		t.Fatalf("ADR-076 raw body not returned from index")
	}
	if !hasHeading(adrResp.Record.Headings, Heading{Level: 1, Text: "076: Design Records MCP"}) || !hasHeadingLevel(adrResp.Record.Headings, 2) {
		t.Fatalf("ADR-076 headings = %#v", adrResp.Record.Headings)
	}
	assertGetRecordJSONShape(t, adrResp, true)

	withoutBody, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "ADR-076"})
	if err != nil {
		t.Fatalf("GetRecord ADR-076 without body: %v", err)
	}
	if withoutBody.Record.Body != nil {
		t.Fatalf("ADR-076 body present when include_body=false")
	}

	specResp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "SPEC-design-records-mcp-tools"})
	if err != nil {
		t.Fatalf("GetRecord spec: %v", err)
	}
	if specResp.Record.Kind != RecordKindSpec || specResp.Record.Title != "Design Records MCP tools" {
		t.Fatalf("spec record = %#v", specResp.Record)
	}

	for _, tt := range []struct {
		id   string
		kind RecordKind
	}{
		{"REQ-MCP-003", RecordKindRequirement},
		{"WORK-MCP-003", RecordKindWorkItem},
		{"TASK-MCP-003-01", RecordKindTask},
	} {
		resp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: tt.id})
		if err != nil {
			t.Fatalf("GetRecord %s: %v", tt.id, err)
		}
		if resp.Record.Kind != tt.kind {
			t.Fatalf("%s kind = %q, want %q", tt.id, resp.Record.Kind, tt.kind)
		}
	}
}

func TestGetRecordEmptyHeadingsJSONShape(t *testing.T) {
	idx := &Index{Records: []Record{{
		ID:       "ADR-001",
		Kind:     RecordKindDecision,
		Title:    "No headings",
		Status:   RecordStatusAccepted,
		Path:     "docs/adr/001-no-headings.md",
		Headings: nil,
	}}}

	resp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "ADR-001"})
	if err != nil {
		t.Fatalf("GetRecord: %v", err)
	}
	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal response: %v", err)
	}
	var raw map[string]map[string]any
	if err := json.Unmarshal(encoded, &raw); err != nil {
		t.Fatalf("Unmarshal response: %v", err)
	}
	headings, ok := raw["record"]["headings"].([]any)
	if !ok {
		t.Fatalf("headings = %#v, want JSON array: %s", raw["record"]["headings"], encoded)
	}
	if len(headings) != 0 {
		t.Fatalf("headings = %#v, want empty array", headings)
	}
	assertGetRecordJSONShape(t, resp, false)
}

func assertGetRecordErrorCode(t *testing.T, idx *Index, req GetRecordRequest, code ErrorCode) {
	t.Helper()
	resp, err := GetRecord(context.Background(), idx, req)
	if err == nil {
		t.Fatalf("GetRecord response = %#v, want ToolError %s", resp, code)
	}
	toolErr, ok := err.(*ToolError)
	if !ok {
		t.Fatalf("error = %T %v, want *ToolError", err, err)
	}
	if toolErr.Code != code {
		t.Fatalf("error code = %q, want %q", toolErr.Code, code)
	}
}

func assertGetRecordJSONShape(t *testing.T, resp GetRecordResponse, wantBody bool) {
	t.Helper()
	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal response: %v", err)
	}
	var raw map[string]map[string]any
	if err := json.Unmarshal(encoded, &raw); err != nil {
		t.Fatalf("Unmarshal response: %v", err)
	}
	record := raw["record"]
	_, hasBody := record["body"]
	if hasBody != wantBody {
		t.Fatalf("body presence = %v, want %v: %s", hasBody, wantBody, encoded)
	}
	allowed := map[string]bool{
		"id":       true,
		"kind":     true,
		"title":    true,
		"status":   true,
		"path":     true,
		"headings": true,
	}
	switch resp.Record.Kind {
	case RecordKindDecision:
		allowed["decision"] = true
	case RecordKindSpec:
		allowed["spec"] = true
	case RecordKindInvestigation:
		allowed["investigation"] = true
	case RecordKindRequirement:
		allowed["requirement"] = true
	case RecordKindWorkItem:
		allowed["work_item"] = true
	case RecordKindTask:
		allowed["task"] = true
	}
	if wantBody {
		allowed["body"] = true
	}
	for key := range allowed {
		if _, ok := record[key]; !ok {
			t.Fatalf("response missing %q: %s", key, encoded)
		}
	}
	for key := range record {
		if !allowed[key] {
			t.Fatalf("response unexpectedly includes %q: %s", key, encoded)
		}
	}
}

func assertHeadings(t *testing.T, got, want []Heading) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("headings = %#v, want %#v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("headings = %#v, want %#v", got, want)
		}
	}
}

func hasHeading(headings []Heading, want Heading) bool {
	for _, heading := range headings {
		if heading == want {
			return true
		}
	}
	return false
}

func hasHeadingLevel(headings []Heading, level int) bool {
	for _, heading := range headings {
		if heading.Level == level {
			return true
		}
	}
	return false
}
