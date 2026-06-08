package mcp

import (
	"encoding/json"
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/v01/src/internal/query"
)

type Server struct {
	service *query.Service
}

type Tool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
}

type Envelope struct {
	Result any        `json:"result,omitempty"`
	Error  *ToolError `json:"error,omitempty"`
}

type ToolError struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Tool    string         `json:"tool,omitempty"`
	Args    map[string]any `json:"args,omitempty"`
}

func NewServer(service *query.Service) *Server {
	return &Server{service: service}
}

func (s *Server) Tools() []Tool {
	return []Tool{
		{
			Name:        "list_objects",
			Description: "List semantic objects in the project.",
			InputSchema: objectSchema(map[string]any{
				"object": enumStringSchema("node", "view", "transition", "field"),
				"kind":   map[string]any{"type": "string"},
				"module": map[string]any{"type": "string"},
				"file":   map[string]any{"type": "string"},
			}, nil),
		},
		{
			Name:        "get_signature",
			Description: "Return the signature for a semantic object.",
			InputSchema: objectSchema(map[string]any{
				"selector": selectorSchema(),
			}, []string{"selector"}),
		},
		{
			Name:        "get_source",
			Description: "Return the YAML source snippet for a semantic object.",
			InputSchema: objectSchema(map[string]any{
				"selector": selectorSchema(),
				"fallback": enumStringSchema("file", "error"),
			}, []string{"selector"}),
		},
		{
			Name:        "get_references",
			Description: "Return direct references for a semantic object.",
			InputSchema: objectSchema(map[string]any{
				"selector":  selectorSchema(),
				"direction": enumStringSchema("out", "in", "both"),
				"kinds": map[string]any{
					"type":  "array",
					"items": map[string]any{"type": "string"},
				},
			}, []string{"selector"}),
		},
		{
			Name:        "get_reference_tree",
			Description: "Return a bounded BFS reference graph for a semantic object.",
			InputSchema: objectSchema(map[string]any{
				"selector":  selectorSchema(),
				"direction": enumStringSchema("out", "in", "both"),
				"depth":     map[string]any{"type": "integer", "minimum": 0, "maximum": 4},
				"kinds": map[string]any{
					"type":  "array",
					"items": map[string]any{"type": "string"},
				},
				"max_nodes": map[string]any{"type": "integer", "minimum": 1},
				"max_edges": map[string]any{"type": "integer", "minimum": 1},
			}, []string{"selector", "direction", "depth"}),
		},
		{
			Name:        "analyze_impact",
			Description: "Return semantic impact analysis for a proposed change.",
			InputSchema: objectSchema(map[string]any{
				"selector": selectorSchema(),
				"change": objectSchema(map[string]any{
					"kind":       enumStringSchema("rename", "remove", "change_type", "change_contract", "change_transition_target", "add"),
					"new_id":     map[string]any{"type": "string"},
					"new_type":   map[string]any{"type": "string"},
					"note":       map[string]any{"type": "string"},
					"new_to":     map[string]any{"type": "string"},
					"new_action": map[string]any{"type": "string"},
					"added_id":   map[string]any{"type": "string"},
				}, []string{"kind"}),
				"scope_modules": map[string]any{
					"type":  "array",
					"items": map[string]any{"type": "string"},
				},
				"max_impacts": map[string]any{"type": "integer", "minimum": 1},
			}, []string{"selector", "change"}),
		},
		{
			Name:        "inspect",
			Description: "Return implementation context for a semantic object.",
			InputSchema: objectSchema(map[string]any{
				"selector": selectorSchema(),
				"detail":   enumStringSchema("brief", "normal", "full"),
			}, []string{"selector"}),
		},
		{
			Name:        "list_endpoints",
			Description: "Return endpoints from API Table views.",
			InputSchema: objectSchema(map[string]any{
				"api_table_id": map[string]any{"type": "string"},
			}, nil),
		},
	}
}

func objectSchema(properties map[string]any, required []string) map[string]any {
	schema := map[string]any{
		"type":                 "object",
		"properties":           properties,
		"additionalProperties": false,
	}
	if len(required) > 0 {
		schema["required"] = required
	}
	return schema
}

func selectorSchema() map[string]any {
	return objectSchema(map[string]any{
		"id":       map[string]any{"type": "string"},
		"object":   enumStringSchema("node", "view", "transition", "asset", "field", "file", "primitive"),
		"kind":     map[string]any{"type": "string"},
		"file":     map[string]any{"type": "string"},
		"local_id": map[string]any{"type": "string"},
	}, nil)
}

func enumStringSchema(values ...string) map[string]any {
	enums := make([]any, 0, len(values))
	for _, value := range values {
		enums = append(enums, value)
	}
	return map[string]any{"type": "string", "enum": enums}
}

func (s *Server) CallToolJSON(name string, args []byte) []byte {
	envelope := s.CallTool(name, args)
	data, err := json.Marshal(envelope)
	if err != nil {
		fallback, _ := json.Marshal(Envelope{Error: &ToolError{Code: "internal_error", Message: err.Error(), Tool: name}})
		return fallback
	}
	return data
}

func (s *Server) CallTool(name string, args []byte) Envelope {
	switch name {
	case "list_objects":
		var req query.ListObjectsRequest
		if err := decodeArgs(args, &req); err != nil {
			return errorEnvelope("invalid_args", name, err.Error(), args)
		}
		res, err := s.service.ListObjects(req)
		if err != nil {
			return errorEnvelope(errorCode(err), name, err.Error(), args)
		}
		return Envelope{Result: res}
	case "get_signature":
		var req query.GetSignatureRequest
		if err := decodeArgs(args, &req); err != nil {
			return errorEnvelope("invalid_args", name, err.Error(), args)
		}
		res, err := s.service.GetSignature(req)
		if err != nil {
			return errorEnvelope(errorCode(err), name, err.Error(), args)
		}
		return Envelope{Result: res}
	case "get_source":
		var req query.GetSourceRequest
		if err := decodeArgs(args, &req); err != nil {
			return errorEnvelope("invalid_args", name, err.Error(), args)
		}
		res, err := s.service.GetSource(req)
		if err != nil {
			return errorEnvelope(errorCode(err), name, err.Error(), args)
		}
		return Envelope{Result: res}
	case "get_references":
		var req query.GetReferencesRequest
		if err := decodeArgs(args, &req); err != nil {
			return errorEnvelope("invalid_args", name, err.Error(), args)
		}
		res, err := s.service.GetReferences(req)
		if err != nil {
			return errorEnvelope(errorCode(err), name, err.Error(), args)
		}
		return Envelope{Result: res}
	case "get_reference_tree":
		var req query.GetReferenceTreeRequest
		if err := decodeArgs(args, &req); err != nil {
			return errorEnvelope("invalid_args", name, err.Error(), args)
		}
		res, err := s.service.GetReferenceTree(req)
		if err != nil {
			return errorEnvelope(errorCode(err), name, err.Error(), args)
		}
		return Envelope{Result: res}
	case "analyze_impact":
		var req query.AnalyzeImpactRequest
		if err := decodeArgs(args, &req); err != nil {
			return errorEnvelope("invalid_args", name, err.Error(), args)
		}
		res, err := s.service.AnalyzeImpact(req)
		if err != nil {
			return errorEnvelope(errorCode(err), name, err.Error(), args)
		}
		return Envelope{Result: res}
	case "inspect":
		var req query.InspectRequest
		if err := decodeArgs(args, &req); err != nil {
			return errorEnvelope("invalid_args", name, err.Error(), args)
		}
		res, err := s.service.Inspect(req)
		if err != nil {
			return errorEnvelope(errorCode(err), name, err.Error(), args)
		}
		return Envelope{Result: res}
	case "list_endpoints":
		var req query.ListEndpointsRequest
		if err := decodeArgs(args, &req); err != nil {
			return errorEnvelope("invalid_args", name, err.Error(), args)
		}
		res, err := s.service.ListEndpoints(req)
		if err != nil {
			return errorEnvelope(errorCode(err), name, err.Error(), args)
		}
		return Envelope{Result: res}
	default:
		return errorEnvelope("unknown_tool", name, "unknown tool: "+name, args)
	}
}

func decodeArgs(args []byte, out any) error {
	if len(args) == 0 {
		args = []byte("{}")
	}
	return json.Unmarshal(args, out)
}

func errorEnvelope(code, tool, message string, args []byte) Envelope {
	return Envelope{Error: &ToolError{Code: code, Message: message, Tool: tool, Args: argsMap(args)}}
}

func argsMap(args []byte) map[string]any {
	if len(args) == 0 {
		return nil
	}
	var out map[string]any
	if err := json.Unmarshal(args, &out); err != nil {
		return nil
	}
	return out
}

func errorCode(err error) string {
	message := err.Error()
	if strings.Contains(message, "source_range_unavailable") {
		return "source_range_unavailable"
	}
	if strings.Contains(message, "invalid fallback") {
		return "invalid_args"
	}
	if strings.Contains(message, "invalid args") {
		return "invalid_args"
	}
	if strings.Contains(message, "not found") {
		return "not_found"
	}
	if strings.Contains(message, "unsupported detail") {
		return "unsupported_detail"
	}
	if strings.Contains(message, "unsupported direction") {
		return "unsupported_direction"
	}
	if strings.Contains(message, "invalid depth") {
		return "invalid_depth"
	}
	if strings.Contains(message, "invalid change payload") {
		return "invalid_change_payload"
	}
	if strings.Contains(message, "kind mismatch") {
		return "kind_mismatch"
	}
	if strings.Contains(message, "unsupported") {
		return "unsupported_object"
	}
	if strings.Contains(message, "selector") {
		return "invalid_selector"
	}
	return "internal_error"
}
