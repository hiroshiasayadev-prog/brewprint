package designrecords

import (
	"context"
	"encoding/json"
	"testing"
)

func TestGetRecordsMixedKindsPartialResultsAndDuplicates(t *testing.T) {
	idx := &Index{Records: []Record{
		{
			ID:       "ADR-077",
			Kind:     RecordKindDecision,
			Title:    "Boundary",
			Status:   RecordStatusAccepted,
			Path:     "docs/adr/077-boundary.md",
			Decision: &DecisionDetail{DependsOn: []string{"ADR-076"}, Supersedes: []string{}},
			Headings: []Heading{{Level: 1, Text: "077: Boundary"}},
			RawBody:  "# 077: Boundary\n",
		},
		{
			ID:       "SPEC-design-records-mcp-tools",
			Kind:     RecordKindSpec,
			Title:    "Design Records MCP tools",
			Status:   RecordStatusDraft,
			Path:     "docs/spec/design-records-mcp/tools.md",
			Spec:     &SpecDetail{DependsOn: []string{"ADR-077"}},
			Headings: []Heading{{Level: 1, Text: "Design Records MCP tools"}},
			RawBody:  "# Design Records MCP tools\n",
		},
		{
			ID:     "INV-DOCS-001",
			Kind:   RecordKindInvestigation,
			Title:  "Investigation",
			Status: RecordStatusConcluded,
			Path:   "docs/investigations/docs/INV-DOCS-001-test.md",
			Investigation: &InvestigationDetail{
				Trigger:            "ADR-085",
				Scope:              "test",
				NonScope:           "none",
				SourceRefs:         []string{"ADR-085"},
				FollowUpCandidates: []string{},
			},
			Headings: []Heading{{Level: 1, Text: "INV-DOCS-001: Investigation"}},
			RawBody:  "# INV-DOCS-001: Investigation\n",
		},
	}}

	resp, err := GetRecords(context.Background(), idx, GetRecordsRequest{
		IDs:         []string{"ADR-077", "SPEC-design-records-mcp-tools", "INV-DOCS-001", "ADR-077", "INV-DOCS-999", "ADR-077"},
		IncludeBody: true,
	})
	if err != nil {
		t.Fatalf("GetRecords: %v", err)
	}
	if got := len(resp.Items); got != 4 {
		t.Fatalf("items len = %d, want 4", got)
	}
	wantIDs := []string{"ADR-077", "SPEC-design-records-mcp-tools", "INV-DOCS-001", "INV-DOCS-999"}
	for i, wantID := range wantIDs {
		if resp.Items[i].ID != wantID {
			t.Fatalf("items[%d].id = %q, want %q", i, resp.Items[i].ID, wantID)
		}
	}
	for i := 0; i < 3; i++ {
		if resp.Items[i].RetrievalStatus != RetrievalStatusFound || resp.Items[i].Record == nil {
			t.Fatalf("items[%d] = %#v, want found record", i, resp.Items[i])
		}
		if resp.Items[i].Record.Body == nil || *resp.Items[i].Record.Body == "" {
			t.Fatalf("items[%d] raw body missing", i)
		}
		if len(resp.Items[i].Diagnostics) != 0 {
			t.Fatalf("items[%d].diagnostics = %#v, want empty", i, resp.Items[i].Diagnostics)
		}
	}
	missing := resp.Items[3]
	if missing.RetrievalStatus != RetrievalStatusNotFound || missing.Record != nil {
		t.Fatalf("missing item = %#v", missing)
	}
	if len(missing.Diagnostics) != 1 || missing.Diagnostics[0].Category != DiagnosticRecordNotFound || missing.Diagnostics[0].Severity != DiagnosticSeverityError || missing.Diagnostics[0].RequestedID != "INV-DOCS-999" {
		t.Fatalf("missing diagnostics = %#v", missing.Diagnostics)
	}
	if len(resp.Diagnostics) != 1 {
		t.Fatalf("top-level diagnostics = %#v, want one duplicate diagnostic", resp.Diagnostics)
	}
	duplicate := resp.Diagnostics[0]
	if duplicate.Category != DiagnosticDuplicateRequestedIDIgnored || duplicate.Severity != DiagnosticSeverityInfo || duplicate.RequestedID != "ADR-077" {
		t.Fatalf("duplicate diagnostic = %#v", duplicate)
	}
	if duplicate.FirstIndex == nil || *duplicate.FirstIndex != 0 {
		t.Fatalf("first_index = %#v, want 0", duplicate.FirstIndex)
	}
	if len(duplicate.DuplicateIndexes) != 2 || duplicate.DuplicateIndexes[0] != 3 || duplicate.DuplicateIndexes[1] != 5 {
		t.Fatalf("duplicate_indexes = %#v, want [3 5]", duplicate.DuplicateIndexes)
	}

	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(encoded, &raw); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if _, ok := raw["items"]; !ok {
		t.Fatalf("response does not contain items: %s", encoded)
	}
	if _, ok := raw["records"]; ok {
		t.Fatalf("response unexpectedly contains records: %s", encoded)
	}
}

func TestGetRecordsAllMissingIsNormalExactLookupResponse(t *testing.T) {
	idx := &Index{Records: []Record{{ID: "ADR-077", Kind: RecordKindDecision}}}
	resp, err := GetRecords(context.Background(), idx, GetRecordsRequest{IDs: []string{"adr-077", " ADR-077 ", "spec:trace", "REQ-MCP-002"}})
	if err != nil {
		t.Fatalf("GetRecords: %v", err)
	}
	if len(resp.Items) != 4 || len(resp.Diagnostics) != 0 {
		t.Fatalf("response = %#v", resp)
	}
	for i, item := range resp.Items {
		if item.RetrievalStatus != RetrievalStatusNotFound || item.Record != nil {
			t.Fatalf("items[%d] = %#v, want not_found", i, item)
		}
		if len(item.Diagnostics) != 1 || item.Diagnostics[0].Category != DiagnosticRecordNotFound {
			t.Fatalf("items[%d].diagnostics = %#v", i, item.Diagnostics)
		}
	}
}

func TestGetRecordsRequestErrors(t *testing.T) {
	idx := &Index{Records: []Record{{ID: "ADR-077", Kind: RecordKindDecision}}}
	for _, test := range []struct {
		name string
		idx  *Index
		req  GetRecordsRequest
	}{
		{name: "missing ids", idx: idx, req: GetRecordsRequest{}},
		{name: "empty ids", idx: idx, req: GetRecordsRequest{IDs: []string{}}},
		{name: "nil index", idx: nil, req: GetRecordsRequest{IDs: []string{"ADR-077"}}},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := GetRecords(context.Background(), test.idx, test.req)
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
