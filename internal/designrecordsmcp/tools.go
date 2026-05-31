package designrecordsmcp

type Tool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
}

func (s *Server) Tools() []Tool {
	return Tools()
}

func Tools() []Tool {
	return []Tool{
		{
			Name:        "list_records",
			Description: "List design records and workflow artifact records with optional metadata filters.",
			InputSchema: objectSchema(map[string]any{
				"kind":     enumStringSchema("decision", "spec", "investigation", "requirement", "work_item", "task"),
				"status":   map[string]any{"type": "string"},
				"id":       map[string]any{"type": "string"},
				"id_range": idRangeSchema(),
				"order_by": enumStringSchema("id"),
				"order":    enumStringSchema("asc", "desc"),
				"limit":    map[string]any{"type": "integer", "minimum": 1},
			}, nil),
		},
		{
			Name:        "validate_records",
			Description: "Validate indexed design record and workflow artifact metadata.",
			InputSchema: objectSchema(map[string]any{
				"kind":     enumStringSchema("decision", "spec", "investigation", "requirement", "work_item", "task"),
				"id_range": idRangeSchema(),
			}, nil),
		},
		{
			Name:        "get_record",
			Description: "Get one design record by ID, optionally including raw Markdown body.",
			InputSchema: objectSchema(map[string]any{
				"id":           map[string]any{"type": "string"},
				"include_body": map[string]any{"type": "boolean"},
			}, []string{"id"}),
		},
		{
			Name:        "get_records",
			Description: "Get multiple explicitly requested design records by exact ID, with item-level partial results.",
			InputSchema: objectSchema(map[string]any{
				"ids": map[string]any{
					"type":     "array",
					"minItems": 1,
					"items":    map[string]any{"type": "string"},
				},
				"include_body": map[string]any{"type": "boolean"},
			}, []string{"ids"}),
		},
		{
			Name:        "list_authoring_guides",
			Description: "List project authoring guides as guide ID, title, and abstract without exposing source paths.",
			InputSchema: objectSchema(map[string]any{}, nil),
		},
		{
			Name:        "get_authoring_guidance",
			Description: "Get project authoring guidance Markdown by exact guide ID.",
			InputSchema: objectSchema(map[string]any{
				"id": map[string]any{"type": "string"},
			}, []string{"id"}),
		},
		{
			Name:        "resolve_reference",
			Description: "Resolve one canonical semantic/artifact reference to a document, section, or record target.",
			InputSchema: objectSchema(map[string]any{
				"ref": map[string]any{"type": "string"},
			}, []string{"ref"}),
		},
		{
			Name:        "suggest_next_record",
			Description: "Suggest the next ADR ID and path for a new decision record.",
			InputSchema: objectSchema(map[string]any{
				"kind":  enumStringSchema("decision"),
				"title": map[string]any{"type": "string"},
			}, []string{"kind", "title"}),
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

func idRangeSchema() map[string]any {
	return objectSchema(map[string]any{
		"from": map[string]any{"type": "string"},
		"to":   map[string]any{"type": "string"},
	}, nil)
}

func enumStringSchema(values ...string) map[string]any {
	enums := make([]any, 0, len(values))
	for _, value := range values {
		enums = append(enums, value)
	}
	return map[string]any{"type": "string", "enum": enums}
}
