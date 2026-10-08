package designrecords

import (
	"context"
	"encoding/json"
	"testing"
)

func TestGetRecordsCurrentExactBatchSuccessWarningsAndNoPath(t *testing.T) {
	idx := currentReadTestIndex()
	resp, err := GetCurrentRecords(context.Background(), idx, CurrentGetRecordsRequest{
		Refs: []string{
			"DRMCP-REQ-MCP-003",
			"DRMCP-TASK-MCP-003-01",
			"DRMCP-REQ-MCP-003",
			"drmcp-req-mcp-003",
			" spec:product.fixture ",
			"spec:product.fixture",
			"DRMCP-REQ-MCP-999",
		},
		IncludeBody: true,
	})
	if err != nil {
		t.Fatalf("GetCurrentRecords: %v", err)
	}
	if got := len(resp.Records); got != 2 {
		t.Fatalf("records len = %d, want 2: %#v", got, resp.Records)
	}
	if resp.Records[0].Ref != "DRMCP-REQ-MCP-003" || resp.Records[1].Ref != "DRMCP-TASK-MCP-003-01" {
		t.Fatalf("record order = %#v", resp.Records)
	}
	if resp.Records[0].Body == nil || *resp.Records[0].Body == "" {
		t.Fatalf("body missing from first record: %#v", resp.Records[0])
	}
	if len(resp.Records[1].Headings) != 1 {
		t.Fatalf("headings = %#v", resp.Records[1].Headings)
	}

	wantWarnings := map[string]string{
		"duplicate_ref":   "DRMCP-REQ-MCP-003",
		"malformed_ref":   "drmcp-req-mcp-003",
		"unsupported_ref": "spec:product.fixture",
		"unresolved_ref":  "DRMCP-REQ-MCP-999",
	}
	for category, ref := range wantWarnings {
		if !hasOperationWarning(resp.Warnings, category, ref) {
			t.Fatalf("missing warning category=%s ref=%s in %#v", category, ref, resp.Warnings)
		}
	}

	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(encoded, &raw); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if _, ok := raw["items"]; ok {
		t.Fatalf("current get_records must not expose items: %s", encoded)
	}
	first := raw["records"].([]any)[0].(map[string]any)
	for _, forbidden := range []string{"id", "path", "decision", "requirement", "retrieval_status", "diagnostics"} {
		if _, ok := first[forbidden]; ok {
			t.Fatalf("get_records leaked %q: %s", forbidden, encoded)
		}
	}
}

func TestGetRecordsCurrentRequestErrors(t *testing.T) {
	idx := currentReadTestIndex()
	tooMany := make([]string, 21)
	for i := range tooMany {
		tooMany[i] = "DRMCP-REQ-MCP-003"
	}
	for _, test := range []struct {
		name string
		idx  *Index
		req  CurrentGetRecordsRequest
	}{
		{name: "missing refs", idx: idx, req: CurrentGetRecordsRequest{}},
		{name: "empty refs", idx: idx, req: CurrentGetRecordsRequest{Refs: []string{}}},
		{name: "too many refs", idx: idx, req: CurrentGetRecordsRequest{Refs: tooMany}},
		{name: "nil index", idx: nil, req: CurrentGetRecordsRequest{Refs: []string{"DRMCP-REQ-MCP-003"}}},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := GetCurrentRecords(context.Background(), test.idx, test.req)
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

func TestNormalReadGetRecordsPathIsNotSerialized(t *testing.T) {
	idx := currentReadTestIndex()
	resp, err := GetCurrentRecords(context.Background(), idx, CurrentGetRecordsRequest{Refs: []string{"DRMCP-REQ-MCP-003"}})
	if err != nil {
		t.Fatalf("GetCurrentRecords: %v", err)
	}
	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if jsonContainsKey(encoded, "path") {
		t.Fatalf("path key leaked into normal read JSON: %s", encoded)
	}
}

func hasOperationWarning(warnings []OperationWarning, category, ref string) bool {
	for _, warning := range warnings {
		if warning.Category == category && warning.Ref == ref {
			return true
		}
	}
	return false
}

func jsonContainsKey(data []byte, key string) bool {
	var value any
	if err := json.Unmarshal(data, &value); err != nil {
		return false
	}
	return jsonValueContainsKey(value, key)
}

func jsonValueContainsKey(value any, key string) bool {
	switch typed := value.(type) {
	case map[string]any:
		if _, ok := typed[key]; ok {
			return true
		}
		for _, child := range typed {
			if jsonValueContainsKey(child, key) {
				return true
			}
		}
	case []any:
		for _, child := range typed {
			if jsonValueContainsKey(child, key) {
				return true
			}
		}
	}
	return false
}
