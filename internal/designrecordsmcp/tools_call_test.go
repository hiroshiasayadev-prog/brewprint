package designrecordsmcp

import (
	"context"
	"encoding/json"
	"errors"
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
				if got := len(resp.Records); got != 3 {
					t.Fatalf("records len = %d, want 3", got)
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
				if got := len(resp.Records); got != 3 {
					t.Fatalf("records len = %d, want 3", got)
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
	return &designrecords.Index{Records: []designrecords.Record{
		toolsCallRecord("ADR-001", "One"),
		toolsCallRecord("ADR-002", "Two"),
		{
			ID:           "SPEC-one",
			NormalizedID: "SPEC-ONE",
			Kind:         designrecords.RecordKindSpec,
			Title:        "Spec One",
			Status:       designrecords.RecordStatusDraft,
			Path:         "docs/spec/one.md",
		},
	}}
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
