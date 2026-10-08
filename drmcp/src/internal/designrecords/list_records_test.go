package designrecords

import (
	"context"
	"encoding/json"
	"testing"
)

func TestListRecordsCurrentCompactFiltersDefaultsAndNoPath(t *testing.T) {
	idx := currentReadTestIndex()

	resp, err := ListCurrentRecords(context.Background(), idx, CurrentListRecordsRequest{
		AppNamespace: "drmcp",
		Kind:         RecordKindRequirement,
		Domain:       "mcp",
	})
	if err != nil {
		t.Fatalf("ListCurrentRecords: %v", err)
	}
	if len(resp.Records) != 2 {
		t.Fatalf("records = %#v, want two requirements", resp.Records)
	}
	if got := []string{resp.Records[0].Ref, resp.Records[1].Ref}; !sameStrings(got, []string{"DRMCP-REQ-MCP-021", "DRMCP-REQ-MCP-003"}) {
		t.Fatalf("default order refs = %#v", got)
	}
	if resp.HasMore {
		t.Fatalf("has_more = true, want false")
	}
	if resp.Records[0].Title == nil || *resp.Records[0].Title != "Later requirement" {
		t.Fatalf("title = %#v", resp.Records[0].Title)
	}
	if resp.Records[0].Status == nil || *resp.Records[0].Status != RecordStatusCaptured {
		t.Fatalf("status = %#v", resp.Records[0].Status)
	}
	if resp.Records[0].Date == nil || *resp.Records[0].Date != "2026-06-21" {
		t.Fatalf("date = %#v", resp.Records[0].Date)
	}
	if len(resp.Warnings) != 0 {
		t.Fatalf("warnings = %#v, want empty", resp.Warnings)
	}

	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(encoded, &raw); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if _, ok := raw["warnings"]; !ok {
		t.Fatalf("warnings must always be present: %s", encoded)
	}
	first := raw["records"].([]any)[0].(map[string]any)
	for _, forbidden := range []string{"id", "kind", "path", "decision", "requirement", "headings", "body"} {
		if _, ok := first[forbidden]; ok {
			t.Fatalf("compact list leaked %q: %s", forbidden, encoded)
		}
	}
}

func TestListRecordsCurrentStatusOrderLimitAndHasMore(t *testing.T) {
	idx := currentReadTestIndex()

	resp, err := ListCurrentRecords(context.Background(), idx, CurrentListRecordsRequest{
		AppNamespace: "drmcp",
		Kind:         RecordKindRequirement,
		Domain:       "MCP",
		Status:       RecordStatusCaptured,
		Order:        "asc",
		Limit:        intPtr(1),
	})
	if err != nil {
		t.Fatalf("ListCurrentRecords: %v", err)
	}
	if len(resp.Records) != 1 || resp.Records[0].Ref != "DRMCP-REQ-MCP-003" {
		t.Fatalf("records = %#v", resp.Records)
	}
	if !resp.HasMore {
		t.Fatalf("has_more = false, want true")
	}
}

func TestListRecordsCurrentRejectsObsoleteAndInvalidInputs(t *testing.T) {
	idx := currentReadTestIndex()
	for _, test := range []struct {
		name string
		req  CurrentListRecordsRequest
	}{
		{name: "missing app namespace", req: CurrentListRecordsRequest{Kind: RecordKindRequirement, Domain: "MCP"}},
		{name: "missing domain", req: CurrentListRecordsRequest{AppNamespace: "drmcp", Kind: RecordKindRequirement}},
		{name: "spec normal listing", req: CurrentListRecordsRequest{AppNamespace: "product", Kind: RecordKindSpec, Domain: "MCP"}},
		{name: "zero limit", req: CurrentListRecordsRequest{AppNamespace: "drmcp", Kind: RecordKindRequirement, Domain: "MCP", Limit: intPtr(0)}},
		{name: "over max limit", req: CurrentListRecordsRequest{AppNamespace: "drmcp", Kind: RecordKindRequirement, Domain: "MCP", Limit: intPtr(101)}},
		{name: "bad order", req: CurrentListRecordsRequest{AppNamespace: "drmcp", Kind: RecordKindRequirement, Domain: "MCP", Order: "newest"}},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := ListCurrentRecords(context.Background(), idx, test.req)
			if err == nil {
				t.Fatal("error = nil, want invalid_request")
			}
			toolErr, ok := err.(*ToolError)
			if !ok || toolErr.Code != ErrorCodeInvalidRequest {
				t.Fatalf("error = %#v, want invalid_request", err)
			}
		})
	}
}

func intPtr(value int) *int {
	return &value
}

func currentReadTestIndex() *Index {
	return &Index{Records: []Record{
		{ID: "DRMCP-REQ-MCP-003", Kind: RecordKindRequirement, Title: "Workflow support", Status: RecordStatusCaptured, Date: "2026-06-03", Path: "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-003.md", RawBody: "# DRMCP-REQ-MCP-003: Workflow support\n", Requirement: &RequirementDetail{}},
		{ID: "DRMCP-REQ-MCP-021", Kind: RecordKindRequirement, Title: "Later requirement", Status: RecordStatusCaptured, Date: "2026-06-21", Path: "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-021.md", RawBody: "# DRMCP-REQ-MCP-021: Later requirement\n", Requirement: &RequirementDetail{}},
		{ID: "DRMCP-WORK-MCP-003", Kind: RecordKindWorkItem, Title: "Workflow work", Status: RecordStatusInProgress, Date: "2026-06-04", Path: "drmcp/records/work-items/mcp/DRMCP-WORK-MCP-003.md", RawBody: "# DRMCP-WORK-MCP-003: Workflow work\n", WorkItem: &WorkItemDetail{}},
		{ID: "DRMCP-TASK-MCP-003-01", Kind: RecordKindTask, Title: "Workflow task", Status: RecordStatusDone, Date: "2026-06-05", Path: "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-003-01.md", RawBody: "# DRMCP-TASK-MCP-003-01: Workflow task\n", Headings: []Heading{{Level: 1, Text: "DRMCP-TASK-MCP-003-01: Workflow task"}}, Task: &TaskDetail{}},
		{ID: "DRMCP-INV-MCP-001", Kind: RecordKindInvestigation, Title: "Investigation", Status: RecordStatusConcluded, Date: "2026-06-06", Path: "drmcp/records/investigations/mcp/DRMCP-INV-MCP-001.md", Investigation: &InvestigationDetail{}},
		{ID: "DRMCP-ADR-MCP-001", Kind: RecordKindDecision, Title: "Decision", Status: RecordStatusAccepted, Date: "2026-06-01", Path: "drmcp/records/adr/mcp/DRMCP-ADR-MCP-001.md", Decision: &DecisionDetail{}},
		{ID: "PRODUCT-REQ-MCP-001", Kind: RecordKindRequirement, Title: "Other app", Status: RecordStatusCaptured, Date: "2026-06-07", Path: "product/records/requirements/mcp/PRODUCT-REQ-MCP-001.md", Requirement: &RequirementDetail{}},
		{ID: "DRMCP-REQ-DATA-001", Kind: RecordKindRequirement, Title: "Other domain", Status: RecordStatusCaptured, Date: "2026-06-08", Path: "drmcp/records/requirements/data/DRMCP-REQ-DATA-001.md", Requirement: &RequirementDetail{}},
		{ID: "DRMCP-SPEC-MCP-001", Kind: RecordKindSpec, Title: "Spec hidden", Status: RecordStatusDraft, Date: "2026-06-09", Path: "drmcp/records/spec/mcp/spec.md", Spec: &SpecDetail{}},
		{ID: "V01-REQ-MCP-999", Kind: RecordKindRequirement, Title: "Legacy hidden", Status: RecordStatusCaptured, Date: "2026-06-10", Path: "v01/records/requirements/mcp/V01-REQ-MCP-999.md", Requirement: &RequirementDetail{}},
	}}
}
