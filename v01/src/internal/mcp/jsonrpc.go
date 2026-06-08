package mcp

import (
	"encoding/json"
	"fmt"
)

const jsonRPCVersion = "2.0"

type JSONRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type JSONRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Result  any             `json:"result,omitempty"`
	Error   *JSONRPCError   `json:"error,omitempty"`
}

type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type InitializeResult struct {
	ProtocolVersion string         `json:"protocolVersion"`
	ServerInfo      map[string]any `json:"serverInfo"`
	Capabilities    map[string]any `json:"capabilities"`
}

type ToolsListResult struct {
	Tools []Tool `json:"tools"`
}

type ToolsCallParams struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments"`
}

type ToolsCallResult struct {
	Content []ToolContent `json:"content"`
	IsError bool          `json:"isError,omitempty"`
}

type ToolContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (s *Server) HandleJSONRPCLine(line []byte) []byte {
	var req JSONRPCRequest
	if err := json.Unmarshal(line, &req); err != nil {
		return marshalJSONRPCResponse(JSONRPCResponse{
			JSONRPC: jsonRPCVersion,
			Error:   &JSONRPCError{Code: -32700, Message: "parse error", Data: err.Error()},
		})
	}
	return marshalJSONRPCResponse(s.HandleJSONRPC(req))
}

func (s *Server) HandleJSONRPC(req JSONRPCRequest) JSONRPCResponse {
	id := req.ID
	if req.JSONRPC != "" && req.JSONRPC != jsonRPCVersion {
		return jsonRPCProtocolError(id, -32600, "invalid request", "jsonrpc must be 2.0")
	}
	if req.Method == "" {
		return jsonRPCProtocolError(id, -32600, "invalid request", "method is required")
	}

	switch req.Method {
	case "initialize":
		return JSONRPCResponse{JSONRPC: jsonRPCVersion, ID: id, Result: InitializeResult{
			ProtocolVersion: "2024-11-05",
			ServerInfo: map[string]any{
				"name":    "brewprint",
				"version": "0.1.0",
			},
			Capabilities: map[string]any{"tools": map[string]any{}},
		}}
	case "tools/list":
		return JSONRPCResponse{JSONRPC: jsonRPCVersion, ID: id, Result: ToolsListResult{Tools: s.Tools()}}
	case "tools/call":
		var params ToolsCallParams
		if err := json.Unmarshal(req.Params, &params); err != nil {
			return jsonRPCProtocolError(id, -32602, "invalid params", err.Error())
		}
		if params.Name == "" {
			return jsonRPCProtocolError(id, -32602, "invalid params", "tool name is required")
		}
		envelope := s.CallTool(params.Name, params.Arguments)
		return JSONRPCResponse{JSONRPC: jsonRPCVersion, ID: id, Result: toolsCallResult(envelope)}
	default:
		return jsonRPCProtocolError(id, -32601, "method not found", fmt.Sprintf("unknown method: %s", req.Method))
	}
}

func jsonRPCProtocolError(id json.RawMessage, code int, message string, data any) JSONRPCResponse {
	return JSONRPCResponse{JSONRPC: jsonRPCVersion, ID: id, Error: &JSONRPCError{Code: code, Message: message, Data: data}}
}

func marshalJSONRPCResponse(res JSONRPCResponse) []byte {
	data, err := json.Marshal(res)
	if err != nil {
		fallback, _ := json.Marshal(JSONRPCResponse{JSONRPC: jsonRPCVersion, Error: &JSONRPCError{Code: -32603, Message: "internal error", Data: err.Error()}})
		return fallback
	}
	return data
}

func toolsCallResult(envelope Envelope) ToolsCallResult {
	payload := envelope.Result
	isError := envelope.Error != nil
	if isError {
		payload = map[string]any{"error": envelope.Error}
	}
	return ToolsCallResult{
		Content: []ToolContent{{Type: "text", Text: mustMarshalString(payload)}},
		IsError: isError,
	}
}

func mustMarshalString(value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		fallback, _ := json.Marshal(map[string]any{"error": ToolError{Code: "internal_error", Message: err.Error()}})
		return string(fallback)
	}
	return string(data)
}
