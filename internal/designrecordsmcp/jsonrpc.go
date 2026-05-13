package designrecordsmcp

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

func (s *Server) HandleJSONRPCLine(line []byte) ([]byte, bool) {
	var req JSONRPCRequest
	if err := json.Unmarshal(line, &req); err != nil {
		return marshalJSONRPCResponse(JSONRPCResponse{
			JSONRPC: jsonRPCVersion,
			Error:   &JSONRPCError{Code: -32700, Message: "parse error", Data: err.Error()},
		}), true
	}
	res, ok := s.HandleJSONRPC(req)
	if !ok {
		return nil, false
	}
	return marshalJSONRPCResponse(res), true
}

func (s *Server) HandleJSONRPC(req JSONRPCRequest) (JSONRPCResponse, bool) {
	if req.Method == "notifications/initialized" {
		return JSONRPCResponse{}, false
	}

	id := req.ID
	if isNotification(req) {
		return JSONRPCResponse{}, false
	}
	if req.JSONRPC != "" && req.JSONRPC != jsonRPCVersion {
		return jsonRPCProtocolError(id, -32600, "invalid request", "jsonrpc must be 2.0"), true
	}
	if req.Method == "" {
		return jsonRPCProtocolError(id, -32600, "invalid request", "method is required"), true
	}

	switch req.Method {
	case "initialize":
		return JSONRPCResponse{JSONRPC: jsonRPCVersion, ID: id, Result: InitializeResult{
			ProtocolVersion: "2024-11-05",
			ServerInfo: map[string]any{
				"name":    "brewprint-design-records-mcp",
				"version": "0.1.0",
			},
			Capabilities: map[string]any{"tools": map[string]any{}},
		}}, true
	case "tools/list":
		return JSONRPCResponse{JSONRPC: jsonRPCVersion, ID: id, Result: ToolsListResult{Tools: s.Tools()}}, true
	case "tools/call":
		return jsonRPCProtocolError(id, -32601, "method not found", "tools/call is not implemented in M17 Phase 1"), true
	default:
		return jsonRPCProtocolError(id, -32601, "method not found", fmt.Sprintf("unknown method: %s", req.Method)), true
	}
}

func isNotification(req JSONRPCRequest) bool {
	return len(req.ID) == 0
}

func jsonRPCProtocolError(id json.RawMessage, code int, message string, data any) JSONRPCResponse {
	return JSONRPCResponse{JSONRPC: jsonRPCVersion, ID: id, Error: &JSONRPCError{Code: code, Message: message, Data: data}}
}

func marshalJSONRPCResponse(res JSONRPCResponse) []byte {
	data, err := json.Marshal(res)
	if err != nil {
		fallback, _ := json.Marshal(JSONRPCResponse{
			JSONRPC: jsonRPCVersion,
			Error:   &JSONRPCError{Code: -32603, Message: "internal error", Data: err.Error()},
		})
		return fallback
	}
	return data
}
