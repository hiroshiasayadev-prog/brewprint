package mcp

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestHandleJSONRPC(t *testing.T) {
	server := newUC001Server(t)

	t.Run("initialize", func(t *testing.T) {
		res := handleLine(t, server, `{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`)
		if res.Error != nil {
			t.Fatalf("initialize error: %#v", res.Error)
		}
		result := resultMapAny(t, res.Result)
		if result["protocolVersion"] == "" {
			t.Fatalf("initialize result missing protocolVersion: %#v", result)
		}
	})

	t.Run("tools_list", func(t *testing.T) {
		res := handleLine(t, server, `{"jsonrpc":"2.0","id":"tools","method":"tools/list"}`)
		if res.Error != nil {
			t.Fatalf("tools/list error: %#v", res.Error)
		}
		result := resultMapAny(t, res.Result)
		tools := result["tools"].([]any)
		if len(tools) != 6 {
			t.Fatalf("tools len = %d, want 6: %#v", len(tools), tools)
		}
		firstTool := tools[0].(map[string]any)
		if _, ok := firstTool["inputSchema"]; !ok {
			t.Fatalf("tool inputSchema missing: %#v", firstTool)
		}
	})

	t.Run("tools_call_get_signature", func(t *testing.T) {
		res := handleLine(t, server, `{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"get_signature","arguments":{"selector":{"id":"auth.task.login"}}}}`)
		if res.Error != nil {
			t.Fatalf("tools/call error: %#v", res.Error)
		}
		content := toolTextMap(t, res)
		object := content["object"].(map[string]any)
		if object["id"] != "auth.task.login" {
			t.Fatalf("tool result object = %#v", object)
		}
	})

	t.Run("tools_call_get_source", func(t *testing.T) {
		res := handleLine(t, server, `{"jsonrpc":"2.0","id":"source","method":"tools/call","params":{"name":"get_source","arguments":{"selector":{"id":"auth.task.login"}}}}`)
		if res.Error != nil {
			t.Fatalf("tools/call get_source error: %#v", res.Error)
		}
		content := toolTextMap(t, res)
		source := content["source"].(map[string]any)
		if source["file"] != "auth/task/login.yaml" {
			t.Fatalf("get_source source = %#v", source)
		}
		snippet := content["snippet"].(map[string]any)
		if snippet["language"] != "yaml" || !strings.Contains(snippet["text"].(string), "id: login") {
			t.Fatalf("get_source snippet = %#v", snippet)
		}
	})

	t.Run("tools_call_list_endpoints", func(t *testing.T) {
		res := handleLine(t, server, `{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"list_endpoints","arguments":{"api_table_id":"ec_api"}}}`)
		if res.Error != nil {
			t.Fatalf("tools/call list_endpoints error: %#v", res.Error)
		}
		content := toolTextMap(t, res)
		if _, ok := content["tables"]; !ok {
			t.Fatalf("list_endpoints tables missing: %#v", content)
		}
	})
}

func TestHandleJSONRPCErrors(t *testing.T) {
	server := newUC001Server(t)

	t.Run("parse_error", func(t *testing.T) {
		data := server.HandleJSONRPCLine([]byte(`{`))
		var res JSONRPCResponse
		if err := json.Unmarshal(data, &res); err != nil {
			t.Fatalf("unmarshal response: %v", err)
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
	})

	t.Run("invalid_params", func(t *testing.T) {
		res := handleLine(t, server, `{"jsonrpc":"2.0","id":5,"method":"tools/call","params":{"arguments":{}}}`)
		if res.Error == nil || res.Error.Code != -32602 {
			t.Fatalf("invalid params response = %#v", res)
		}
	})

	t.Run("unknown_tool_is_tool_error", func(t *testing.T) {
		res := handleLine(t, server, `{"jsonrpc":"2.0","id":6,"method":"tools/call","params":{"name":"missing_tool","arguments":{}}}`)
		if res.Error != nil {
			t.Fatalf("unexpected protocol error: %#v", res.Error)
		}
		toolError := toolTextError(t, res)
		if toolError.Code != "unknown_tool" {
			t.Fatalf("unknown tool error = %#v", toolError)
		}
	})
}

func handleLine(t *testing.T, server *Server, line string) JSONRPCResponse {
	t.Helper()
	data := server.HandleJSONRPCLine([]byte(line))
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

func toolTextMap(t *testing.T, res JSONRPCResponse) map[string]any {
	t.Helper()
	text := toolText(t, res)
	var out map[string]any
	if err := json.Unmarshal([]byte(text), &out); err != nil {
		t.Fatalf("unmarshal tool text: %v\n%s", err, text)
	}
	return out
}

func toolTextError(t *testing.T, res JSONRPCResponse) ToolError {
	t.Helper()
	content := toolTextMap(t, res)
	errorMap, ok := content["error"].(map[string]any)
	if !ok {
		t.Fatalf("tool error missing: %#v", content)
	}
	data, err := json.Marshal(errorMap)
	if err != nil {
		t.Fatalf("marshal tool error: %v", err)
	}
	var toolError ToolError
	if err := json.Unmarshal(data, &toolError); err != nil {
		t.Fatalf("unmarshal tool error: %v", err)
	}
	return toolError
}

func toolText(t *testing.T, res JSONRPCResponse) string {
	t.Helper()
	result := resultMapAny(t, res.Result)
	content := result["content"].([]any)
	if len(content) != 1 {
		t.Fatalf("content len = %d, want 1: %#v", len(content), content)
	}
	item := content[0].(map[string]any)
	return item["text"].(string)
}
