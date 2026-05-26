package designrecords

import (
	"context"
	"encoding/json"
	"testing"
)

func TestListRecordsBasicFiltersAndResponseShape(t *testing.T) {
	idx := buildListRecordsTestIndex(t)

	resp, err := ListRecords(context.Background(), idx, ListRecordsRequest{})
	if err != nil {
		t.Fatalf("ListRecords: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"ADR-066", "ADR-067", "ADR-076", "ADR-077", "ADR-078", "INV-DOCS-001", "REQ-MCP-003", "SPEC-design-records-mcp-overview", "SPEC-design-records-mcp-schema", "TASK-MCP-003-01", "WORK-MCP-003"}) {
		t.Fatalf("record IDs = %#v", got)
	}

	record := findListedRecord(resp.Records, "ADR-076")
	if record == nil {
		t.Fatal("missing ADR-076")
	}
	if record.Kind != RecordKindDecision || record.Title != "Design Records MCP" || record.Status != RecordStatusAccepted || record.Path != "docs/adr/076-design-records-mcp.md" {
		t.Fatalf("ADR-076 metadata = %#v", record)
	}
	if record.Decision == nil {
		t.Fatalf("Decision detail missing: %#v", record)
	}
	assertStrings(t, record.Decision.DependsOn, []string{"ADR-067"})
	assertStrings(t, record.Decision.Supersedes, []string{"ADR-066"})
	if record.Decision.MigratedToSpec != nil {
		t.Fatalf("MigratedToSpec = %#v, want nil", record.Decision.MigratedToSpec)
	}

	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal response: %v", err)
	}
	var raw map[string][]map[string]any
	if err := json.Unmarshal(encoded, &raw); err != nil {
		t.Fatalf("Unmarshal response: %v", err)
	}
	first := raw["records"][0]
	for _, unexpected := range []string{"headings", "body", "RawBody", "raw_body", "depends_on", "supersedes", "migrated_to_spec"} {
		if _, ok := first[unexpected]; ok {
			t.Fatalf("response unexpectedly includes %q: %s", unexpected, encoded)
		}
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindDecision})
	if err != nil {
		t.Fatalf("ListRecords decision: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"ADR-066", "ADR-067", "ADR-076", "ADR-077", "ADR-078"}) {
		t.Fatalf("decision IDs = %#v", got)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindSpec})
	if err != nil {
		t.Fatalf("ListRecords spec: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"SPEC-design-records-mcp-overview", "SPEC-design-records-mcp-schema"}) {
		t.Fatalf("spec IDs = %#v", got)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindInvestigation})
	if err != nil {
		t.Fatalf("ListRecords investigation: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"INV-DOCS-001"}) {
		t.Fatalf("investigation IDs = %#v", got)
	}
	investigation := resp.Records[0]
	if investigation.Investigation == nil || investigation.Investigation.Trigger != "ADR-076" {
		t.Fatalf("investigation detail = %#v", investigation.Investigation)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindRequirement})
	if err != nil {
		t.Fatalf("ListRecords requirement: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"REQ-MCP-003"}) {
		t.Fatalf("requirement IDs = %#v", got)
	}
	if resp.Records[0].Requirement == nil || resp.Records[0].Requirement.WorkItems[0] != "WORK-MCP-003" {
		t.Fatalf("requirement detail = %#v", resp.Records[0].Requirement)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindWorkItem})
	if err != nil {
		t.Fatalf("ListRecords work_item: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"WORK-MCP-003"}) {
		t.Fatalf("work item IDs = %#v", got)
	}
	if resp.Records[0].WorkItem == nil || resp.Records[0].WorkItem.SourceRequirement != "REQ-MCP-003" {
		t.Fatalf("work item detail = %#v", resp.Records[0].WorkItem)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindTask})
	if err != nil {
		t.Fatalf("ListRecords task: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"TASK-MCP-003-01"}) {
		t.Fatalf("task IDs = %#v", got)
	}
	if resp.Records[0].Task == nil || resp.Records[0].Task.WorkItem != "WORK-MCP-003" {
		t.Fatalf("task detail = %#v", resp.Records[0].Task)
	}

	assertListRecordsErrorCode(t, idx, ListRecordsRequest{Kind: RecordKind("milestone")}, ErrorCodeInvalidRequest)
}

func TestListRecordsStatusAndIDFilters(t *testing.T) {
	idx := buildListRecordsTestIndex(t)

	resp, err := ListRecords(context.Background(), idx, ListRecordsRequest{Status: RecordStatusProposed})
	if err != nil {
		t.Fatalf("ListRecords status: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"ADR-067"}) {
		t.Fatalf("status IDs = %#v", got)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{Status: RecordStatusWIP})
	if err != nil {
		t.Fatalf("ListRecords status no-match: %v", err)
	}
	if len(resp.Records) != 0 {
		t.Fatalf("status no-match records = %#v, want empty", resp.Records)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{ID: "ADR-076"})
	if err != nil {
		t.Fatalf("ListRecords id: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"ADR-076"}) {
		t.Fatalf("id IDs = %#v", got)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{ID: "ADR-999"})
	if err != nil {
		t.Fatalf("ListRecords non-existing id: %v", err)
	}
	if len(resp.Records) != 0 {
		t.Fatalf("non-existing id records = %#v, want empty", resp.Records)
	}
}

func TestListRecordsIDRangeFilter(t *testing.T) {
	idx := buildListRecordsTestIndex(t)

	tests := []struct {
		name string
		req  ListRecordsRequest
		want []string
	}{
		{
			name: "inclusive endpoints",
			req:  ListRecordsRequest{IDRange: &IDRange{From: "ADR-067", To: "ADR-077"}},
			want: []string{"ADR-067", "ADR-076", "ADR-077"},
		},
		{
			name: "one sided from",
			req:  ListRecordsRequest{IDRange: &IDRange{From: "ADR-076"}},
			want: []string{"ADR-076", "ADR-077", "ADR-078"},
		},
		{
			name: "one sided to",
			req:  ListRecordsRequest{IDRange: &IDRange{To: "ADR-067"}},
			want: []string{"ADR-066", "ADR-067"},
		},
		{
			name: "omitted kind behaves as decision",
			req:  ListRecordsRequest{IDRange: &IDRange{From: "ADR-067", To: "ADR-077"}},
			want: []string{"ADR-067", "ADR-076", "ADR-077"},
		},
		{
			name: "decision kind works",
			req:  ListRecordsRequest{Kind: RecordKindDecision, IDRange: &IDRange{From: "ADR-067", To: "ADR-077"}},
			want: []string{"ADR-067", "ADR-076", "ADR-077"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ListRecords(context.Background(), idx, tt.req)
			if err != nil {
				t.Fatalf("ListRecords: %v", err)
			}
			if got := listedRecordIDs(resp.Records); !sameStrings(got, tt.want) {
				t.Fatalf("record IDs = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestListRecordsRequestErrors(t *testing.T) {
	idx := buildListRecordsTestIndex(t)
	tests := []struct {
		name string
		req  ListRecordsRequest
		code ErrorCode
	}{
		{
			name: "invalid kind",
			req:  ListRecordsRequest{Kind: RecordKind("milestone")},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "kind spec with id range",
			req:  ListRecordsRequest{Kind: RecordKindSpec, IDRange: &IDRange{From: "ADR-067"}},
			code: ErrorCodeIDRangeRequiresDecisionKind,
		},
		{
			name: "kind requirement with id range",
			req:  ListRecordsRequest{Kind: RecordKindRequirement, IDRange: &IDRange{From: "ADR-067"}},
			code: ErrorCodeIDRangeRequiresDecisionKind,
		},
		{
			name: "kind work_item with id range",
			req:  ListRecordsRequest{Kind: RecordKindWorkItem, IDRange: &IDRange{From: "ADR-067"}},
			code: ErrorCodeIDRangeRequiresDecisionKind,
		},
		{
			name: "kind task with id range",
			req:  ListRecordsRequest{Kind: RecordKindTask, IDRange: &IDRange{From: "ADR-067"}},
			code: ErrorCodeIDRangeRequiresDecisionKind,
		},
		{
			name: "SPEC range endpoint",
			req:  ListRecordsRequest{IDRange: &IDRange{From: "SPEC-design-records-mcp-schema"}},
			code: ErrorCodeIDRangeRequiresDecisionKind,
		},
		{
			name: "malformed ADR range endpoint",
			req:  ListRecordsRequest{IDRange: &IDRange{From: "ADR-x"}},
			code: ErrorCodeIDRangeRequiresDecisionKind,
		},
		{
			name: "invalid order by",
			req:  ListRecordsRequest{OrderBy: "status"},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "invalid order",
			req:  ListRecordsRequest{Order: "newest"},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "zero limit",
			req:  ListRecordsRequest{Limit: intPtr(0)},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "negative limit",
			req:  ListRecordsRequest{Limit: intPtr(-1)},
			code: ErrorCodeInvalidRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertListRecordsErrorCode(t, idx, tt.req, tt.code)
		})
	}
}

func TestListRecordsSortOrderAndLimit(t *testing.T) {
	idx := buildListRecordsTestIndex(t)

	resp, err := ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindDecision, OrderBy: "id", Order: "asc"})
	if err != nil {
		t.Fatalf("ListRecords asc: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"ADR-066", "ADR-067", "ADR-076", "ADR-077", "ADR-078"}) {
		t.Fatalf("asc IDs = %#v", got)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindDecision, OrderBy: "id", Order: "desc"})
	if err != nil {
		t.Fatalf("ListRecords desc: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"ADR-078", "ADR-077", "ADR-076", "ADR-067", "ADR-066"}) {
		t.Fatalf("desc IDs = %#v", got)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{OrderBy: "id", Order: "desc"})
	if err != nil {
		t.Fatalf("ListRecords mixed desc: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"WORK-MCP-003", "TASK-MCP-003-01", "SPEC-design-records-mcp-schema", "SPEC-design-records-mcp-overview", "REQ-MCP-003", "INV-DOCS-001", "ADR-078", "ADR-077", "ADR-076", "ADR-067", "ADR-066"}) {
		t.Fatalf("mixed desc IDs = %#v", got)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindDecision})
	if err != nil {
		t.Fatalf("ListRecords default order: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"ADR-066", "ADR-067", "ADR-076", "ADR-077", "ADR-078"}) {
		t.Fatalf("default order IDs = %#v", got)
	}

	resp, err = ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindDecision, Order: "desc", Limit: intPtr(2)})
	if err != nil {
		t.Fatalf("ListRecords limit: %v", err)
	}
	if got := listedRecordIDs(resp.Records); !sameStrings(got, []string{"ADR-078", "ADR-077"}) {
		t.Fatalf("limit IDs = %#v", got)
	}
}

func TestListRecordsRepositoryBootstrapQueries(t *testing.T) {
	root := findRepoRoot(t)
	cfg, err := NewConfig(root)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}

	latest, err := ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindDecision, OrderBy: "id", Order: "desc", Limit: intPtr(1)})
	if err != nil {
		t.Fatalf("ListRecords latest ADR: %v", err)
	}
	if len(latest.Records) != 1 || latest.Records[0].Kind != RecordKindDecision {
		t.Fatalf("latest ADR response = %#v", latest.Records)
	}
	if want := latestDecisionRecordID(t, idx); latest.Records[0].ID != want {
		t.Fatalf("latest ADR ID = %q, want %q", latest.Records[0].ID, want)
	}

	rangeResp, err := ListRecords(context.Background(), idx, ListRecordsRequest{IDRange: &IDRange{From: "ADR-067", To: "ADR-077"}})
	if err != nil {
		t.Fatalf("ListRecords ADR range: %v", err)
	}
	for _, id := range []string{"ADR-067", "ADR-068", "ADR-069", "ADR-070", "ADR-071", "ADR-072", "ADR-073", "ADR-074", "ADR-075", "ADR-076", "ADR-077"} {
		if findListedRecord(rangeResp.Records, id) == nil {
			t.Fatalf("ADR range missing %s in %#v", id, listedRecordIDs(rangeResp.Records))
		}
	}

	specResp, err := ListRecords(context.Background(), idx, ListRecordsRequest{Kind: RecordKindSpec})
	if err != nil {
		t.Fatalf("ListRecords specs: %v", err)
	}
	for _, id := range []string{"SPEC-design-records-mcp-overview", "SPEC-design-records-mcp-schema", "SPEC-design-records-mcp-tools"} {
		if findListedRecord(specResp.Records, id) == nil {
			t.Fatalf("spec records missing %s in %#v", id, listedRecordIDs(specResp.Records))
		}
	}

	for _, tt := range []struct {
		kind RecordKind
		id   string
	}{
		{RecordKindRequirement, "REQ-MCP-003"},
		{RecordKindWorkItem, "WORK-MCP-003"},
		{RecordKindTask, "TASK-MCP-003-01"},
	} {
		resp, err := ListRecords(context.Background(), idx, ListRecordsRequest{Kind: tt.kind})
		if err != nil {
			t.Fatalf("ListRecords %s: %v", tt.kind, err)
		}
		if findListedRecord(resp.Records, tt.id) == nil {
			t.Fatalf("%s records missing %s in %#v", tt.kind, tt.id, listedRecordIDs(resp.Records))
		}
	}

	assertListRecordsErrorCode(t, idx, ListRecordsRequest{Kind: RecordKindSpec, IDRange: &IDRange{From: "ADR-067"}}, ErrorCodeIDRangeRequiresDecisionKind)
}

func buildListRecordsTestIndex(t *testing.T) *Index {
	t.Helper()
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/066-old.md", "# 066: Old\n- **status**: accepted\n- **depends_on**:\n- **supersedes**:\n")
	writeTestFile(t, root, "docs/adr/067-foundation.md", "# 067: Foundation\n- **status**: proposed\n- **depends_on**:\n- **supersedes**:\n")
	writeTestFile(t, root, "docs/adr/076-design-records-mcp.md", "# 076: Design Records MCP\n- **status**: accepted\n- **depends_on**: ADR-067\n- **supersedes**: ADR-066\n")
	writeTestFile(t, root, "docs/adr/077-boundary.md", "# 077: Boundary\n- **status**: accepted\n- **depends_on**: ADR-076\n- **supersedes**:\n")
	writeTestFile(t, root, "docs/adr/078-next.md", "# 078: Next\n- **status**: superseded\n- **depends_on**:\n- **supersedes**:\n- **migrated_to_spec**: 2026-05-12\n")
	writeTestFile(t, root, "docs/spec/design-records-mcp/overview.md", "---\nstatus: draft\ndesign_record:\n  id: SPEC-design-records-mcp-overview\n  kind: spec\n  status: draft\n  depends_on:\n    - ADR-076\n---\n# Design Records MCP overview\n")
	writeTestFile(t, root, "docs/spec/design-records-mcp/schema.md", "---\nstatus: confirmed\ndesign_record:\n  id: SPEC-design-records-mcp-schema\n  kind: spec\n  status: confirmed\n  depends_on:\n    - ADR-076\n---\n# Design Records MCP schema\n")
	writeTestFile(t, root, "docs/investigations/docs/INV-DOCS-001-test.md", "# INV-DOCS-001: Test investigation\n- **status**: concluded\n- **date**: 2026-05-19\n- **trigger**: ADR-076\n- **scope**: test\n- **non_scope**: none\n- **source_refs**:\n  - ADR-076\n- **follow_up_candidates**:\n  - SPEC-design-records-mcp-schema\n")
	writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-003-test.md", "# REQ-MCP-003: Test requirement\n- **id**: REQ-MCP-003\n- **status**: accepted\n- **date**: 2026-05-25\n- **source_refs**:\n  - ADR-076\n- **work_items**:\n  - WORK-MCP-003\n")
	writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-003-test.md", "# WORK-MCP-003: Test work item\n- **id**: WORK-MCP-003\n- **status**: implementation_pending\n- **date**: 2026-05-26\n- **source_requirement**: REQ-MCP-003\n- **impact_refs**:\n  - ADR-076\n- **tasks**:\n  - TASK-MCP-003-01\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-003-01-test.md", "# TASK-MCP-003-01: Test task\n- **id**: TASK-MCP-003-01\n- **status**: todo\n- **date**: 2026-05-26\n- **work_item**: WORK-MCP-003\n- **source_requirement**: REQ-MCP-003\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n  - test\n")
	return buildTestIndex(t, root)
}

func assertListRecordsErrorCode(t *testing.T, idx *Index, req ListRecordsRequest, code ErrorCode) {
	t.Helper()
	_, err := ListRecords(context.Background(), idx, req)
	if err == nil {
		t.Fatal("ListRecords error = nil, want ToolError")
	}
	toolErr, ok := err.(*ToolError)
	if !ok {
		t.Fatalf("error = %T %v, want *ToolError", err, err)
	}
	if toolErr.Code != code {
		t.Fatalf("error code = %q, want %q", toolErr.Code, code)
	}
}

func listedRecordIDs(records []ListedRecord) []string {
	ids := make([]string, 0, len(records))
	for _, record := range records {
		ids = append(ids, record.ID)
	}
	return ids
}

func findListedRecord(records []ListedRecord, id string) *ListedRecord {
	for i := range records {
		if records[i].ID == id {
			return &records[i]
		}
	}
	return nil
}

func latestDecisionRecordID(t *testing.T, idx *Index) string {
	t.Helper()
	maxNum := -1
	maxID := ""
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
			maxID = record.ID
		}
	}
	if maxID == "" {
		t.Fatal("no decision records in index")
	}
	return maxID
}

func intPtr(value int) *int {
	return &value
}
