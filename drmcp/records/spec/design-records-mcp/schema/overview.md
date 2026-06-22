# Overview: Schema

- **id**: `spec:drmcp.design_records_mcp.schema.overview`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Defines the data model, metadata grammar, field definitions, ID normalization, discovery rules, authoring guidance source, and authoring transaction schema for Design Records MCP.

## Current contract

The schema group covers:

- Which Markdown sources DRMCP reads and what metadata each record kind supplies.
- How record IDs are structured (public ID vs. bare ID grammar) and normalized.
- How records are discovered within a records tree.
- What the internal record model looks like and what bootstrap records seed MVP validation.
- How validation diagnostics are categorized and what fields they carry.
- How authoring guides are discovered and served.
- How authoring transactions (proposals, body cache, section selectors) are structured.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Record source | Reference | `spec:drmcp.design_records_mcp.schema.record_source` | Markdown source types and metadata sources per record kind. |
| Metadata grammar | Reference | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Bullet metadata grammar for ADR, investigation, and workflow artifacts. |
| Fields | Reference | `spec:drmcp.design_records_mcp.schema.fields` | Common and kind-specific field definitions, title extraction, and field validation rules. |
| ID normalization | Reference | `spec:drmcp.design_records_mcp.schema.id_normalization` | Public ID vs. bare ID grammar; namespace_prefix stripping and canonical ID derivation. |
| Discovery | Reference | `spec:drmcp.design_records_mcp.schema.discovery` | Discovery path rules and index inclusion conditions per record kind. |
| Record model | Reference | `spec:drmcp.design_records_mcp.schema.record_model` | Internal record model fields, headings/body inclusion rules, and bootstrap records. |
| Diagnostics | Reference | `spec:drmcp.design_records_mcp.schema.diagnostics` | Validation diagnostic categories, severity, required fields, and workflow/investigation-specific rules. |
| Authoring guidance source | Reference | `spec:drmcp.design_records_mcp.schema.authoring_guidance_source` | Guide discovery, ID derivation, title/abstract extraction, and tool response contracts. |
| Authoring transaction schema | Reference | `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | Proposal model, body cache, metadata block replacement target, and section selector model. |
