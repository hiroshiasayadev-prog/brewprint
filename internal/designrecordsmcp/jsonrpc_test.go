package designrecordsmcp

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/hiroshiasayadev-prog/brewprint/internal/designrecords"
)

func TestHandleJSONRPC(t *testing.T) {
	server, calls := newCountingServer()

	t.Run("initialize", func(t *testing.T) {
		res := handleLine(t, server, `{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`)
		if res.Error != nil {
			t.Fatalf("initialize error: %#v", res.Error)
		}
		if string(res.ID) != "1" {
			t.Fatalf("response id = %s, want 1", string(res.ID))
		}
		result := resultMapAny(t, res.Result)
		serverInfo := result["serverInfo"].(map[string]any)
		if !strings.Contains(strings.ToLower(serverInfo["name"].(string)), "design-records") {
			t.Fatalf("serverInfo.name = %#v", serverInfo["name"])
		}
		if result["protocolVersion"] == "" {
			t.Fatalf("initialize result missing protocolVersion: %#v", result)
		}
		if *calls != 0 {
			t.Fatalf("initialize called BuildIndex %d time(s), want 0", *calls)
		}
	})

	t.Run("initialized_notification", func(t *testing.T) {
		response, ok := server.HandleJSONRPCLine([]byte(`{"jsonrpc":"2.0","method":"notifications/initialized","params":{}}`))
		if ok || len(response) != 0 {
			t.Fatalf("initialized notification response = ok:%v %s, want none", ok, string(response))
		}
	})

	t.Run("tools_list", func(t *testing.T) {
		res := handleLine(t, server, `{"jsonrpc":"2.0","id":"tools","method":"tools/list","params":{}}`)
		if res.Error != nil {
			t.Fatalf("tools/list error: %#v", res.Error)
		}
		result := resultMapAny(t, res.Result)
		tools := result["tools"].([]any)
		if len(tools) != 8 {
			t.Fatalf("tools len = %d, want 8: %#v", len(tools), tools)
		}
		for _, name := range []string{"list_records", "validate_records", "get_record", "get_records", "list_authoring_guides", "get_authoring_guidance", "resolve_reference", "suggest_next_record"} {
			if !hasToolName(tools, name) {
				t.Fatalf("%s missing from tools/list: %#v", name, tools)
			}
		}
		first := tools[0].(map[string]any)
		if _, ok := first["inputSchema"]; !ok {
			t.Fatalf("tool inputSchema missing: %#v", first)
		}
		if *calls != 0 {
			t.Fatalf("tools/list called BuildIndex %d time(s), want 0", *calls)
		}
	})
}

func TestHandleJSONRPCErrors(t *testing.T) {
	server, _ := newCountingServer()

	t.Run("parse_error", func(t *testing.T) {
		data, ok := server.HandleJSONRPCLine([]byte(`{`))
		if !ok {
			t.Fatalf("parse error returned no response")
		}
		var res JSONRPCResponse
		if err := json.Unmarshal(data, &res); err != nil {
			t.Fatalf("unmarshal parse error response: %v\n%s", err, string(data))
		}
		if res.Error == nil || res.Error.Code != -32700 {
			t.Fatalf("parse error response = %#v", res)
		}
	})

	t.Run("unknown_method", func(t *testing.T) {
		res := handleLine(t, server, `{"jsonrpc":"2.0","id":4,"method":"missing/method"}`)
		if res.Error == nil || res.Error.Code != -32601 {
			t.Fatalf("unknown method response = %#v", res)
		}
		if string(res.ID) != "4" {
			t.Fatalf("unknown method response id = %s, want 4", string(res.ID))
		}
	})

	t.Run("tools_call", func(t *testing.T) {
		res := handleLine(t, server, `{"jsonrpc":"2.0","id":"call","method":"tools/call","params":{"name":"list_records","arguments":{}}}`)
		if res.Error != nil {
			t.Fatalf("tools/call response = %#v", res)
		}
		if string(res.ID) != `"call"` {
			t.Fatalf("tools/call response id = %s, want \"call\"", string(res.ID))
		}
		result := assertToolCallResult(t, res, false)
		var text designrecords.ListRecordsResponse
		unmarshalToolText(t, result.Content[0].Text, &text)
	})
}

func TestServeJSONRPCLinesSkipsNotifications(t *testing.T) {
	server, _ := newCountingServer()
	input := strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`,
		`{"jsonrpc":"2.0","method":"notifications/initialized","params":{}}`,
		`{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}`,
	}, "\n") + "\n"

	var out bytes.Buffer
	if err := server.ServeJSONRPCLines(strings.NewReader(input), &out); err != nil {
		t.Fatalf("serve lines: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("response lines len = %d, want 2\n%s", len(lines), out.String())
	}
}

func TestServeJSONRPCLinesToolsCall(t *testing.T) {
	server, _ := newCountingServer()

	var out bytes.Buffer
	if err := server.ServeJSONRPCLines(strings.NewReader(`{"jsonrpc":"2.0","id":9,"method":"tools/call","params":{"name":"list_records","arguments":{}}}`+"\n"), &out); err != nil {
		t.Fatalf("serve lines: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 1 {
		t.Fatalf("response lines len = %d, want 1\n%s", len(lines), out.String())
	}
	var res JSONRPCResponse
	if err := json.Unmarshal([]byte(lines[0]), &res); err != nil {
		t.Fatalf("tools/call response is not JSON: %v\n%s", err, lines[0])
	}
	if res.Error != nil {
		t.Fatalf("tools/call response = %#v", res)
	}
	if string(res.ID) != "9" {
		t.Fatalf("tools/call response id = %s, want 9", string(res.ID))
	}
	assertToolCallResult(t, res, false)
}

func newCountingServer() (*Server, *int) {
	calls := 0
	server := NewServerWithIndexBuilder(designrecords.Config{Root: "."}, func(context.Context, designrecords.Config) (*designrecords.Index, error) {
		calls++
		return &designrecords.Index{}, nil
	})
	return server, &calls
}

func handleLine(t *testing.T, server *Server, line string) JSONRPCResponse {
	t.Helper()
	data, ok := server.HandleJSONRPCLine([]byte(line))
	if !ok {
		t.Fatalf("line returned no response: %s", line)
	}
	var res JSONRPCResponse
	if err := json.Unmarshal(data, &res); err != nil {
		t.Fatalf("unmarshal JSON-RPC response: %v\n%s", err, string(data))
	}
	return res
}

func resultMapAny(t *testing.T, result any) map[string]any {
	t.Helper()
	data, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	var out map[string]any
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("unmarshal result map: %v", err)
	}
	return out
}

func hasToolName(tools []any, name string) bool {
	for _, item := range tools {
		tool, ok := item.(map[string]any)
		if ok && tool["name"] == name {
			return true
		}
	}
	return false
}
