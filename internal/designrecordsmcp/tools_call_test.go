package designrecordsmcp

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/designrecords"
)

func TestToolsCallSuccess(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		wantID     string
		assertText func(t *testing.T, text string)
	}{
		{
			name:   "list_records",
			line:   `{"jsonrpc":"2.0","id":10,"method":"tools/call","params":{"name":"list_records","arguments":{"kind":"decision","limit":1}}}`,
			wantID: "10",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.ListRecordsResponse
				unmarshalToolText(t, text, &resp)
				if got := len(resp.Records); got != 1 {
					t.Fatalf("records len = %d, want 1", got)
				}
				if resp.Records[0].ID != "ADR-001" {
					t.Fatalf("record ID = %q, want ADR-001", resp.Records[0].ID)
				}
			},
		},
		{
			name:   "validate_records",
			line:   `{"jsonrpc":"2.0","id":11,"method":"tools/call","params":{"name":"validate_records","arguments":{}}}`,
			wantID: "11",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.ValidateRecordsResponse
				unmarshalToolText(t, text, &resp)
				if !resp.OK || len(resp.Diagnostics) != 0 {
					t.Fatalf("validate response = %#v, want ok with no diagnostics", resp)
				}
			},
		},
		{
			name:   "list_records omitted arguments",
			line:   `{"jsonrpc":"2.0","id":14,"method":"tools/call","params":{"name":"list_records"}}`,
			wantID: "14",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.ListRecordsResponse
				unmarshalToolText(t, text, &resp)
				if got := len(resp.Records); got != 6 {
					t.Fatalf("records len = %d, want 6", got)
				}
			},
		},
		{
			name:   "list_records null arguments",
			line:   `{"jsonrpc":"2.0","id":15,"method":"tools/call","params":{"name":"list_records","arguments":null}}`,
			wantID: "15",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.ListRecordsResponse
				unmarshalToolText(t, text, &resp)
				if got := len(resp.Records); got != 6 {
					t.Fatalf("records len = %d, want 6", got)
				}
			},
		},
		{
			name:   "list_records workflow requirement",
			line:   `{"jsonrpc":"2.0","id":151,"method":"tools/call","params":{"name":"list_records","arguments":{"kind":"requirement"}}}`,
			wantID: "151",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.ListRecordsResponse
				unmarshalToolText(t, text, &resp)
				if len(resp.Records) != 1 || resp.Records[0].ID != "REQ-MCP-003" || resp.Records[0].Requirement == nil {
					t.Fatalf("requirement list response = %#v", resp)
				}
			},
		},
		{
			name:   "get_record workflow task",
			line:   `{"jsonrpc":"2.0","id":152,"method":"tools/call","params":{"name":"get_record","arguments":{"id":"TASK-MCP-003-01"}}}`,
			wantID: "152",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.GetRecordResponse
				unmarshalToolText(t, text, &resp)
				if resp.Record.ID != "TASK-MCP-003-01" || resp.Record.Task == nil || resp.Record.Task.WorkItem != "WORK-MCP-003" {
					t.Fatalf("task get_record response = %#v", resp)
				}
			},
		},
		{
			name:   "get_record",
			line:   `{"jsonrpc":"2.0","id":12,"method":"tools/call","params":{"name":"get_record","arguments":{"id":"ADR-001","include_body":true}}}`,
			wantID: "12",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.GetRecordResponse
				unmarshalToolText(t, text, &resp)
				if resp.Record.ID != "ADR-001" || resp.Record.Body == nil || *resp.Record.Body == "" {
					t.Fatalf("get_record response = %#v", resp)
				}
			},
		},
		{
			name:   "get_records partial result and duplicate info",
			line:   `{"jsonrpc":"2.0","id":121,"method":"tools/call","params":{"name":"get_records","arguments":{"ids":["ADR-001","SPEC-one","ADR-001","INV-DOCS-999"],"include_body":true}}}`,
			wantID: "121",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.GetRecordsResponse
				unmarshalToolText(t, text, &resp)
				if len(resp.Items) != 3 || resp.Items[0].ID != "ADR-001" || resp.Items[1].ID != "SPEC-one" || resp.Items[2].RetrievalStatus != designrecords.RetrievalStatusNotFound {
					t.Fatalf("get_records response = %#v", resp)
				}
				if resp.Items[0].Record == nil || resp.Items[0].Record.Body == nil || *resp.Items[0].Record.Body == "" {
					t.Fatalf("get_records found body = %#v", resp.Items[0])
				}
				if len(resp.Items[2].Diagnostics) != 1 || resp.Items[2].Diagnostics[0].RequestedID != "INV-DOCS-999" {
					t.Fatalf("get_records missing diagnostic = %#v", resp.Items[2].Diagnostics)
				}
				if len(resp.Diagnostics) != 1 || resp.Diagnostics[0].RequestedID != "ADR-001" || resp.Diagnostics[0].FirstIndex == nil || *resp.Diagnostics[0].FirstIndex != 0 {
					t.Fatalf("get_records duplicate diagnostic = %#v", resp.Diagnostics)
				}
			},
		},
		{
			name:   "resolve_reference",
			line:   `{"jsonrpc":"2.0","id":16,"method":"tools/call","params":{"name":"resolve_reference","arguments":{"ref":"spec:one.doc"}}}`,
			wantID: "16",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.ResolveReferenceResponse
				unmarshalToolText(t, text, &resp)
				if resp.Status != "resolved" || resp.Target == nil || resp.Target.Path != "docs/spec/one.md" {
					t.Fatalf("resolve_reference response = %#v", resp)
				}
			},
		},
		{
			name:   "resolve_reference workflow requirement",
			line:   `{"jsonrpc":"2.0","id":161,"method":"tools/call","params":{"name":"resolve_reference","arguments":{"ref":"REQ-MCP-003"}}}`,
			wantID: "161",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.ResolveReferenceResponse
				unmarshalToolText(t, text, &resp)
				if resp.Status != "resolved" || resp.RefKind != "record_id" || resp.Target == nil || resp.Target.RecordID != "REQ-MCP-003" || resp.Target.RecordKind != designrecords.RecordKindRequirement {
					t.Fatalf("resolve_reference requirement response = %#v", resp)
				}
			},
		},
		{
			name:   "resolve_reference workflow task",
			line:   `{"jsonrpc":"2.0","id":162,"method":"tools/call","params":{"name":"resolve_reference","arguments":{"ref":"TASK-MCP-003-01"}}}`,
			wantID: "162",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.ResolveReferenceResponse
				unmarshalToolText(t, text, &resp)
				if resp.Status != "resolved" || resp.RefKind != "record_id" || resp.Target == nil || resp.Target.RecordID != "TASK-MCP-003-01" || resp.Target.RecordKind != designrecords.RecordKindTask {
					t.Fatalf("resolve_reference task response = %#v", resp)
				}
			},
		},
		{
			name:   "suggest_next_record",
			line:   `{"jsonrpc":"2.0","id":13,"method":"tools/call","params":{"name":"suggest_next_record","arguments":{"kind":"decision","title":"Next Thing"}}}`,
			wantID: "13",
			assertText: func(t *testing.T, text string) {
				var resp designrecords.SuggestNextRecordResponse
				unmarshalToolText(t, text, &resp)
				if resp.NextID != "ADR-003" || resp.NextNumber != 3 || resp.ExistingMaxID != "ADR-002" {
					t.Fatalf("suggest_next_record response = %#v", resp)
				}
				if resp.SuggestedPath != "docs/adr/003-next-thing.md" {
					t.Fatalf("suggested path = %q", resp.SuggestedPath)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			server := NewServerWithIndexBuilder(designrecords.Config{Root: "."}, func(context.Context, designrecords.Config) (*designrecords.Index, error) {
				calls++
				return toolsCallTestIndex(), nil
			})

			res := handleLine(t, server, tt.line)
			if res.Error != nil {
				t.Fatalf("tools/call protocol error: %#v", res.Error)
			}
			if string(res.ID) != tt.wantID {
				t.Fatalf("response ID = %s, want %s", string(res.ID), tt.wantID)
			}
			result := assertToolCallResult(t, res, false)
			tt.assertText(t, result.Content[0].Text)
			if calls != 1 {
				t.Fatalf("BuildIndex calls = %d, want 1", calls)
			}
		})
	}
}

func TestToolsCallValidateDiagnosticsAreNormalResponse(t *testing.T) {
	server := NewServerWithIndexBuilder(designrecords.Config{Root: "."}, func(context.Context, designrecords.Config) (*designrecords.Index, error) {
		idx := toolsCallTestIndex()
		idx.Records = append(idx.Records, designrecords.Record{
			ID:           "ADR-003",
			NormalizedID: "ADR-003",
			Kind:         designrecords.RecordKindDecision,
			Title:        "Bad status",
			Status:       designrecords.RecordStatusDraft,
			Path:         "docs/adr/003-bad-status.md",
		})
		return idx, nil
	})

	res := handleLine(t, server, `{"jsonrpc":"2.0","id":"validate","method":"tools/call","params":{"name":"validate_records","arguments":{}}}`)
	result := assertToolCallResult(t, res, false)

	var text map[string]any
	unmarshalToolText(t, result.Content[0].Text, &text)
	if text["ok"] != false {
		t.Fatalf("ok = %#v, want false", text["ok"])
	}
	diagnostics, ok := text["diagnostics"].([]any)
	if !ok || len(diagnostics) == 0 {
		t.Fatalf("diagnostics missing from response text: %#v", text)
	}
}

func TestToolsListWorkflowKindEnums(t *testing.T) {
	tools := Tools()
	for _, toolName := range []string{"list_records", "validate_records"} {
		tool := findToolForTest(tools, toolName)
		if tool == nil {
			t.Fatalf("missing tool %s", toolName)
		}
		kind, ok := tool.InputSchema["properties"].(map[string]any)["kind"].(map[string]any)
		if !ok {
			t.Fatalf("%s kind schema missing: %#v", toolName, tool.InputSchema)
		}
		enumValues, ok := kind["enum"].([]any)
		if !ok {
			t.Fatalf("%s kind enum missing: %#v", toolName, kind)
		}
		for _, want := range []string{"decision", "spec", "investigation", "requirement", "work_item", "task"} {
			if !hasEnumValue(enumValues, want) {
				t.Fatalf("%s kind enum missing %q: %#v", toolName, want, enumValues)
			}
		}
	}
	if !strings.Contains(Tools()[0].Description, "workflow") {
		t.Fatalf("list_records description does not mention workflow artifacts: %q", Tools()[0].Description)
	}
}

func TestToolsListAuthoringGuidanceTools(t *testing.T) {
	tools := Tools()
	listTool := findToolForTest(tools, "list_authoring_guides")
	if listTool == nil {
		t.Fatal("missing list_authoring_guides")
	}
	if got := listTool.InputSchema["additionalProperties"]; got != false {
		t.Fatalf("list_authoring_guides additionalProperties = %#v, want false", got)
	}

	getTool := findToolForTest(tools, "get_authoring_guidance")
	if getTool == nil {
		t.Fatal("missing get_authoring_guidance")
	}
	required, ok := getTool.InputSchema["required"].([]string)
	if !ok || len(required) != 1 || required[0] != "id" {
		t.Fatalf("get_authoring_guidance required = %#v", getTool.InputSchema["required"])
	}
}

func TestToolsCallAuthoringGuidance(t *testing.T) {
	root := t.TempDir()
	writeToolsCallTestFile(t, root, "docs/guides/zeta.md", "# Zeta Guide\n\n## Abstract\n\nZeta summary.\n\n## Body\n\nZeta body.\n")
	content := "# Alpha Guide\n\n## Abstract\n\nAlpha summary.\n\n## Body\n\nAlpha body.\n"
	writeToolsCallTestFile(t, root, "docs/guides/alpha.md", content)

	calls := 0
	server := NewServerWithIndexBuilder(designrecords.Config{Root: root}, func(context.Context, designrecords.Config) (*designrecords.Index, error) {
		calls++
		return toolsCallTestIndex(), nil
	})

	list := handleLine(t, server, `{"jsonrpc":"2.0","id":501,"method":"tools/call","params":{"name":"list_authoring_guides","arguments":{}}}`)
	listResult := assertToolCallResult(t, list, false)
	var listResp designrecords.ListAuthoringGuidesResponse
	unmarshalToolText(t, listResult.Content[0].Text, &listResp)
	if len(listResp.Guides) != 2 || listResp.Guides[0].ID != "alpha" || listResp.Guides[1].ID != "zeta" {
		t.Fatalf("list_authoring_guides response = %#v", listResp)
	}
	assertToolTextJSONKeys(t, listResult.Content[0].Text, []string{"guides"})
	assertToolValueJSONKeys(t, listResp.Guides[0], []string{"abstract", "id", "title"})

	get := handleLine(t, server, `{"jsonrpc":"2.0","id":502,"method":"tools/call","params":{"name":"get_authoring_guidance","arguments":{"id":"alpha"}}}`)
	getResult := assertToolCallResult(t, get, false)
	var getResp designrecords.GetAuthoringGuidanceResponse
	unmarshalToolText(t, getResult.Content[0].Text, &getResp)
	if getResp.ID != "alpha" || getResp.Title != "Alpha Guide" || getResp.Content != content {
		t.Fatalf("get_authoring_guidance response = %#v", getResp)
	}
	assertToolValueJSONKeys(t, getResp, []string{"content", "id", "title"})

	if calls != 2 {
		t.Fatalf("BuildIndex calls = %d, want 2", calls)
	}
}

func TestToolsCallToolErrors(t *testing.T) {
	tests := []struct {
		name string
		line string
		code designrecords.ErrorCode
	}{
		{
			name: "get_record missing id",
			line: `{"jsonrpc":"2.0","id":20,"method":"tools/call","params":{"name":"get_record","arguments":{}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "get_record unknown id",
			line: `{"jsonrpc":"2.0","id":21,"method":"tools/call","params":{"name":"get_record","arguments":{"id":"ADR-999"}}}`,
			code: designrecords.ErrorCodeRecordNotFound,
		},
		{
			name: "get_records missing ids",
			line: `{"jsonrpc":"2.0","id":211,"method":"tools/call","params":{"name":"get_records","arguments":{}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "get_records empty ids",
			line: `{"jsonrpc":"2.0","id":212,"method":"tools/call","params":{"name":"get_records","arguments":{"ids":[]}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "get_records non-string id",
			line: `{"jsonrpc":"2.0","id":213,"method":"tools/call","params":{"name":"get_records","arguments":{"ids":["ADR-001",7]}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "get_authoring_guidance missing id",
			line: `{"jsonrpc":"2.0","id":214,"method":"tools/call","params":{"name":"get_authoring_guidance","arguments":{}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "get_authoring_guidance unknown id",
			line: `{"jsonrpc":"2.0","id":215,"method":"tools/call","params":{"name":"get_authoring_guidance","arguments":{"id":"unknown-guide"}}}`,
			code: designrecords.ErrorCodeGuideNotFound,
		},
		{
			name: "get_authoring_guidance non-string id",
			line: `{"jsonrpc":"2.0","id":216,"method":"tools/call","params":{"name":"get_authoring_guidance","arguments":{"id":7}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "list_authoring_guides unknown argument",
			line: `{"jsonrpc":"2.0","id":217,"method":"tools/call","params":{"name":"list_authoring_guides","arguments":{"extra":true}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "suggest_next_record spec kind",
			line: `{"jsonrpc":"2.0","id":22,"method":"tools/call","params":{"name":"suggest_next_record","arguments":{"kind":"spec","title":"Spec"}}}`,
			code: designrecords.ErrorCodeUnsupportedKind,
		},
		{
			name: "list_records spec kind with id range",
			line: `{"jsonrpc":"2.0","id":23,"method":"tools/call","params":{"name":"list_records","arguments":{"kind":"spec","id_range":{"from":"ADR-001"}}}}`,
			code: designrecords.ErrorCodeIDRangeRequiresDecisionKind,
		},
		{
			name: "unknown tool name",
			line: `{"jsonrpc":"2.0","id":24,"method":"tools/call","params":{"name":"missing_tool","arguments":{}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "missing tool name",
			line: `{"jsonrpc":"2.0","id":241,"method":"tools/call","params":{"arguments":{}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "empty tool name",
			line: `{"jsonrpc":"2.0","id":242,"method":"tools/call","params":{"name":"","arguments":{}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "malformed arguments",
			line: `{"jsonrpc":"2.0","id":25,"method":"tools/call","params":{"name":"list_records","arguments":{"limit":"bad"}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "resolve_reference unknown argument",
			line: `{"jsonrpc":"2.0","id":251,"method":"tools/call","params":{"name":"resolve_reference","arguments":{"ref":"ADR-001","extra":true}}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
		{
			name: "omitted arguments is empty object",
			line: `{"jsonrpc":"2.0","id":26,"method":"tools/call","params":{"name":"get_record"}}`,
			code: designrecords.ErrorCodeInvalidRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewServerWithIndexBuilder(designrecords.Config{Root: "."}, func(context.Context, designrecords.Config) (*designrecords.Index, error) {
				return toolsCallTestIndex(), nil
			})
			res := handleLine(t, server, tt.line)
			result := assertToolCallResult(t, res, true)
			assertToolErrorCode(t, result.Content[0].Text, tt.code)
		})
	}
}

func TestToolsCallBuildIndexErrorIsToolError(t *testing.T) {
	server := NewServerWithIndexBuilder(designrecords.Config{Root: "."}, func(context.Context, designrecords.Config) (*designrecords.Index, error) {
		return nil, errors.New("boom")
	})

	res := handleLine(t, server, `{"jsonrpc":"2.0","id":27,"method":"tools/call","params":{"name":"list_records","arguments":{}}}`)
	result := assertToolCallResult(t, res, true)

	var envelope struct {
		Error designrecords.ToolError `json:"error"`
	}
	unmarshalToolText(t, result.Content[0].Text, &envelope)
	if envelope.Error.Code != designrecords.ErrorCodeInvalidRequest {
		t.Fatalf("tool error code = %q, want %q; text=%s", envelope.Error.Code, designrecords.ErrorCodeInvalidRequest, result.Content[0].Text)
	}
	if envelope.Error.Message == "" {
		t.Fatalf("tool error message is empty: %s", result.Content[0].Text)
	}
}

func TestToolsCallToolErrorsStillBuildIndexOnce(t *testing.T) {
	tests := []string{
		`{"jsonrpc":"2.0","id":30,"method":"tools/call","params":{"name":"missing_tool","arguments":{}}}`,
		`{"jsonrpc":"2.0","id":31,"method":"tools/call","params":{"name":"list_records","arguments":{"limit":"bad"}}}`,
		`{"jsonrpc":"2.0","id":32,"method":"tools/call","params":{"arguments":{}}}`,
	}
	for _, line := range tests {
		t.Run(line, func(t *testing.T) {
			calls := 0
			server := NewServerWithIndexBuilder(designrecords.Config{Root: "."}, func(context.Context, designrecords.Config) (*designrecords.Index, error) {
				calls++
				return toolsCallTestIndex(), nil
			})
			res := handleLine(t, server, line)
			assertToolCallResult(t, res, true)
			if calls != 1 {
				t.Fatalf("BuildIndex calls = %d, want 1", calls)
			}
		})
	}
}

func TestToolsCallIndexRebuildPolicy(t *testing.T) {
	calls := 0
	server := NewServerWithIndexBuilder(designrecords.Config{Root: "."}, func(context.Context, designrecords.Config) (*designrecords.Index, error) {
		calls++
		if calls == 1 {
			return &designrecords.Index{Records: []designrecords.Record{toolsCallRecord("ADR-001", "First")}}, nil
		}
		return &designrecords.Index{Records: []designrecords.Record{toolsCallRecord("ADR-002", "Second")}}, nil
	})

	initialize := handleLine(t, server, `{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`)
	if initialize.Error != nil {
		t.Fatalf("initialize error: %#v", initialize.Error)
	}
	toolsList := handleLine(t, server, `{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}`)
	if toolsList.Error != nil {
		t.Fatalf("tools/list error: %#v", toolsList.Error)
	}
	if calls != 0 {
		t.Fatalf("initialize/tools_list BuildIndex calls = %d, want 0", calls)
	}

	first := handleLine(t, server, `{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"list_records","arguments":{}}}`)
	firstResult := assertToolCallResult(t, first, false)
	assertListRecordsTextIDs(t, firstResult.Content[0].Text, []string{"ADR-001"})
	if calls != 1 {
		t.Fatalf("after first tools/call BuildIndex calls = %d, want 1", calls)
	}

	second := handleLine(t, server, `{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"list_records","arguments":{}}}`)
	secondResult := assertToolCallResult(t, second, false)
	assertListRecordsTextIDs(t, secondResult.Content[0].Text, []string{"ADR-002"})
	if calls != 2 {
		t.Fatalf("after second tools/call BuildIndex calls = %d, want 2", calls)
	}
}

func toolsCallTestIndex() *designrecords.Index {
	idx := &designrecords.Index{Records: []designrecords.Record{
		toolsCallRecord("ADR-001", "One"),
		toolsCallRecord("ADR-002", "Two"),
		{
			ID:           "SPEC-one",
			NormalizedID: "SPEC-ONE",
			Kind:         designrecords.RecordKindSpec,
			Title:        "Spec One",
			Status:       designrecords.RecordStatusDraft,
			Path:         "docs/spec/one.md",
			Spec:         &designrecords.SpecDetail{},
			SemanticRefs: []designrecords.SemanticRefDecl{{Ref: "spec:one.doc", Path: "docs/spec/one.md", TargetType: designrecords.SemanticTargetDocument}},
		},
		{
			ID:           "REQ-MCP-003",
			NormalizedID: "REQ-MCP-003",
			Kind:         designrecords.RecordKindRequirement,
			Title:        "Workflow support",
			Status:       designrecords.RecordStatusAccepted,
			Path:         "docs/requirements/mcp/REQ-MCP-003-workflow-support.md",
			RawBody:      "# REQ-MCP-003: Workflow support\n",
			Requirement: &designrecords.RequirementDetail{
				SourceRefs: []string{"ADR-092"},
				WorkItems:  []string{"WORK-MCP-003"},
			},
		},
		{
			ID:           "WORK-MCP-003",
			NormalizedID: "WORK-MCP-003",
			Kind:         designrecords.RecordKindWorkItem,
			Title:        "Workflow implementation",
			Status:       designrecords.RecordStatusImplementationPending,
			Path:         "docs/work-items/mcp/WORK-MCP-003-workflow-implementation.md",
			RawBody:      "# WORK-MCP-003: Workflow implementation\n",
			WorkItem: &designrecords.WorkItemDetail{
				SourceRequirement: "REQ-MCP-003",
				ImpactRefs:        []string{"ADR-092"},
				Tasks:             []string{"TASK-MCP-003-01"},
			},
		},
		{
			ID:           "TASK-MCP-003-01",
			NormalizedID: "TASK-MCP-003-01",
			Kind:         designrecords.RecordKindTask,
			Title:        "Workflow evidence",
			Status:       designrecords.RecordStatusDone,
			Path:         "docs/tasks/mcp/TASK-MCP-003-01-workflow-evidence.md",
			RawBody:      "# TASK-MCP-003-01: Workflow evidence\n",
			Task: &designrecords.TaskDetail{
				WorkItem:          "WORK-MCP-003",
				SourceRequirement: "REQ-MCP-003",
				Estimate:          "0.5d",
				DependsOn:         []string{},
				Outputs:           []string{"evidence"},
			},
		},
	}}
	idx.SemanticRefs = []designrecords.SemanticRefDecl{{Ref: "spec:one.doc", Path: "docs/spec/one.md", TargetType: designrecords.SemanticTargetDocument}}
	return idx
}

func findToolForTest(tools []Tool, name string) *Tool {
	for i := range tools {
		if tools[i].Name == name {
			return &tools[i]
		}
	}
	return nil
}

func hasEnumValue(values []any, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func toolsCallRecord(id, title string) designrecords.Record {
	return designrecords.Record{
		ID:           id,
		NormalizedID: strings.ToUpper(id),
		Kind:         designrecords.RecordKindDecision,
		Title:        title,
		Status:       designrecords.RecordStatusAccepted,
		Path:         "docs/adr/" + strings.TrimPrefix(id, "ADR-") + "-" + strings.ToLower(title) + ".md",
		Headings:     []designrecords.Heading{{Level: 1, Text: strings.TrimPrefix(id, "ADR-") + ": " + title}},
		RawBody:      "# " + strings.TrimPrefix(id, "ADR-") + ": " + title + "\n- **status**: accepted\n",
	}
}

func assertToolCallResult(t *testing.T, res JSONRPCResponse, wantError bool) ToolsCallResult {
	t.Helper()
	if res.Error != nil {
		t.Fatalf("unexpected protocol error: %#v", res.Error)
	}
	data, err := json.Marshal(res.Result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	var result ToolsCallResult
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("unmarshal tools/call result: %v\n%s", err, data)
	}
	if result.IsError != wantError {
		t.Fatalf("isError = %v, want %v; result=%#v", result.IsError, wantError, result)
	}
	if len(result.Content) != 1 {
		t.Fatalf("content len = %d, want 1", len(result.Content))
	}
	if result.Content[0].Type != "text" {
		t.Fatalf("content[0].type = %q, want text", result.Content[0].Type)
	}
	if !json.Valid([]byte(result.Content[0].Text)) {
		t.Fatalf("content[0].text is not valid JSON: %s", result.Content[0].Text)
	}
	return result
}

func unmarshalToolText(t *testing.T, text string, out any) {
	t.Helper()
	if err := json.Unmarshal([]byte(text), out); err != nil {
		t.Fatalf("unmarshal tool text: %v\n%s", err, text)
	}
}

func assertToolErrorCode(t *testing.T, text string, want designrecords.ErrorCode) {
	t.Helper()
	var envelope struct {
		Error designrecords.ToolError `json:"error"`
	}
	unmarshalToolText(t, text, &envelope)
	if envelope.Error.Code != want {
		t.Fatalf("tool error code = %q, want %q; text=%s", envelope.Error.Code, want, text)
	}
	if envelope.Error.Message == "" {
		t.Fatalf("tool error message is empty: %s", text)
	}
}

func assertListRecordsTextIDs(t *testing.T, text string, want []string) {
	t.Helper()
	var resp designrecords.ListRecordsResponse
	unmarshalToolText(t, text, &resp)
	got := make([]string, 0, len(resp.Records))
	for _, record := range resp.Records {
		got = append(got, record.ID)
	}
	if len(got) != len(want) {
		t.Fatalf("record IDs = %#v, want %#v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("record IDs = %#v, want %#v", got, want)
		}
	}
}

func writeToolsCallTestFile(t *testing.T, root, rel, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
}

func assertToolTextJSONKeys(t *testing.T, text string, want []string) {
	t.Helper()
	var got map[string]any
	unmarshalToolText(t, text, &got)
	assertMapKeys(t, got, want)
}

func assertToolValueJSONKeys(t *testing.T, value any, want []string) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	assertMapKeys(t, got, want)
}

func assertMapKeys(t *testing.T, got map[string]any, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("JSON keys = %#v, want %#v", got, want)
	}
	for _, key := range want {
		if _, ok := got[key]; !ok {
			t.Fatalf("JSON missing key %q in %#v", key, got)
		}
	}
}
