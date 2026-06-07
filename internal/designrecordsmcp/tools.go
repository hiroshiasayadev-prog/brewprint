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
		{
			Name:        "propose_record_create",
			Description: "Create a retained no-write proposal for a new decision or workflow artifact record.",
			InputSchema: objectSchema(map[string]any{
				"kind":                   enumStringSchema("decision", "requirement", "work_item", "task"),
				"id":                     map[string]any{"type": "string"},
				"domain":                 map[string]any{"type": "string"},
				"parent_id":              map[string]any{"type": "string"},
				"title":                  map[string]any{"type": "string"},
				"fields":                 map[string]any{"type": "object", "additionalProperties": true},
				"body":                   map[string]any{"type": "string"},
				"body_cache_id":          map[string]any{"type": "string"},
				"reciprocal_update_mode": enumStringSchema("include_required", "report_required_follow_up"),
			}, []string{"kind", "id", "title", "fields"}),
		},
		{
			Name:        "propose_record_update",
			Description: "Create a retained no-write proposal for one or more update operations. Supply either update (single-op) or operations (multi-op array); they are mutually exclusive.",
			InputSchema: objectSchema(map[string]any{
				"kind":          enumStringSchema("decision", "spec", "requirement", "work_item", "task"),
				"id":            map[string]any{"type": "string"},
				"update":        updateSchema(),
				"operations":    operationsSchema(),
				"body":          map[string]any{"type": "string"},
				"body_cache_id": map[string]any{"type": "string"},
			}, []string{"kind", "id"}),
		},
		{
			Name:        "get_proposed_write",
			Description: "Get one retained authoring proposal by proposal ID.",
			InputSchema: objectSchema(map[string]any{
				"proposal_id": map[string]any{"type": "string"},
			}, []string{"proposal_id"}),
		},
		{
			Name:        "accept_proposed_write",
			Description: "Accept a retained authoring proposal and write repository files after guards pass.",
			InputSchema: objectSchema(map[string]any{
				"proposal_id": map[string]any{"type": "string"},
			}, []string{"proposal_id"}),
		},
		{
			Name:        "discard_proposed_write",
			Description: "Discard a retained authoring proposal without writing repository files.",
			InputSchema: objectSchema(map[string]any{
				"proposal_id": map[string]any{"type": "string"},
			}, []string{"proposal_id"}),
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

func updateSchema() map[string]any {
	return objectSchema(map[string]any{
		"type":             enumStringSchema("metadata_block_replace", "metadata_fields_replace", "named_section_replace"),
		"metadata":         map[string]any{"type": "object", "additionalProperties": true},
		"section_selector": sectionSelectorSchema(),
	}, []string{"type"})
}

func operationsSchema() map[string]any {
	item := objectSchema(map[string]any{
		"type":             enumStringSchema("metadata_block_replace", "metadata_fields_replace", "named_section_replace"),
		"metadata":         map[string]any{"type": "object", "additionalProperties": true},
		"section_selector": sectionSelectorSchema(),
		"body":             map[string]any{"type": "string"},
		"body_cache_id":    map[string]any{"type": "string"},
	}, []string{"type"})
	return map[string]any{
		"type":     "array",
		"items":    item,
		"minItems": 1,
	}
}

func sectionSelectorSchema() map[string]any {
	return objectSchema(map[string]any{
		"heading": map[string]any{"type": "string"},
		"match":   enumStringSchema("exact"),
		"level":   map[string]any{"type": "integer", "minimum": 1, "maximum": 6},
	}, []string{"heading"})
}
