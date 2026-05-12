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
	idx := buildTestIndex(t, root)

	adrResp, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "ADR-076"})
	if err != nil {
		t.Fatalf("GetRecord ADR-076: %v", err)
	}
	adr := adrResp.Record
	if adr.ID != "ADR-076" || adr.Kind != RecordKindDecision || adr.Title != "Design Records MCP" || adr.Status != RecordStatusAccepted || adr.Path != "docs/adr/076-design-records-mcp.md" {
		t.Fatalf("ADR metadata = %#v", adr)
	}
	assertStrings(t, adr.DependsOn, []string{"ADR-050", "ADR-068"})
	assertStrings(t, adr.Supersedes, []string{"ADR-001"})
	if adr.MigratedToSpec != nil {
		t.Fatalf("MigratedToSpec = %#v, want nil", adr.MigratedToSpec)
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
	assertStrings(t, spec.DependsOn, []string{"ADR-076"})
	assertHeadings(t, spec.Headings, []Heading{
		{Level: 1, Text: "Design Records MCP tools"},
		{Level: 2, Text: "get_record"},
	})
	if spec.Body != nil {
		t.Fatalf("spec body = %#v, want nil when include_body is omitted", *spec.Body)
	}
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
		"id":               true,
		"kind":             true,
		"title":            true,
		"status":           true,
		"path":             true,
		"depends_on":       true,
		"supersedes":       true,
		"migrated_to_spec": true,
		"headings":         true,
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
