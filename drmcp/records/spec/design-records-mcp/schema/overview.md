# Overview: Schema

- **id**: `spec:drmcp.design_records_mcp.schema.overview`
- **status**: draft
- **date**: 2026-07-04
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Defines the data model, metadata grammar, field definitions, ID normalization, discovery rules, authoring guidance source, and authoring transaction schema for Design Records MCP.

## Current contract

The schema group covers:

- Which current Markdown sources DRMCP discovers under configured records-root and spec-tree sources.
- Which metadata and H1 values each record kind supplies.
- How complete current sequential IDs and path-derived spec refs establish canonical identity.
- How valid, invalid, identity-less, and duplicate sources participate in the active index and validation inputs.
- How source, parsed, index, conflict, finding, and public-projection states remain separate internally.
- How validation diagnostics are categorized and what fields they carry.
- How authoring-standard Specs are selected and projected through Guidance operations.
- How authoring transactions use proposals, body cache, and section selectors.

`spec:drmcp.implementation` owns the concrete internal layer and Go package architecture. Schema Specifications own the public and semantic structures consumed by that architecture.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Record source | Reference | `spec:drmcp.design_records_mcp.schema.record_source` | Markdown source types and metadata sources per record kind. |
| Metadata grammar | Reference | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | H1-adjacent visible metadata grammar for current sequential artifacts and specs. |
| Fields | Reference | `spec:drmcp.design_records_mcp.schema.fields` | Common and kind-specific parsed fields and source mapping. |
| ID normalization | Reference | `spec:drmcp.design_records_mcp.schema.id_normalization` | Complete app-aware sequential IDs and path-derived current spec refs. |
| Discovery | Reference | `spec:drmcp.design_records_mcp.schema.discovery` | Current candidate paths, invalid-source handling, and addressability boundaries. |
| Record model | Reference | `spec:drmcp.design_records_mcp.schema.record_model` | Source provenance, active-index outcomes, and validation-input retention. |
| Diagnostics | Reference | `spec:drmcp.design_records_mcp.schema.diagnostics` | Validation diagnostic categories, severity, required fields, and workflow/investigation-specific rules. |
| Authoring guidance source | Reference | `spec:drmcp.design_records_mcp.schema.authoring_guidance_source` | Fixed Current Records scope, canonical-ref identity, Guidance projection, and tool response contracts. |
| Authoring transaction schema | Reference | `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | Proposal model, body cache, metadata block replacement target, and section selector model. |
