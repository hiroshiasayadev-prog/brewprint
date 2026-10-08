# Overview: Retrieval operations

- **id**: `spec:drmcp.design_records_mcp.operations.retrieval`
- **status**: draft
- **date**: 2026-07-14
- **parent**: `spec:drmcp.design_records_mcp.operations`

## What this is

Defines exact current-record retrieval and exact H2 section retrieval.

## Current contract

| operation | purpose |
|---|---|
| `get_records` | Retrieve exact current records with one request-wide projection. |
| `get_record_sections` | Retrieve exact H2 sections from exact current records. |

Each operation name is also the MCP tool name.
Both operations preserve one result for each supplied selector occurrence.

## Non-goals

- Record discovery, listing, search, validation, or legacy lookup.
- Reference repair, completion, normalization, inference, or fuzzy matching.
- Physical source-path exposure.
- Shared diagnostic object, code, context, or message rules.
- Parser, cache, storage, snapshot, concurrency, or filesystem architecture.

## Boundary

| concern | owner |
|---|---|
| Current-record admission and unique addressability | `spec:drmcp.design_records_mcp.current_record_model`. |
| Artifact metadata, H1, H2, and source formats | `spec:drmcp.design_records_mcp.artifacts`. |
| Shared batch-result invariants | `spec:drmcp.design_records_mcp.operations.retrieval.batch_retrieval_results`. |
| Exact record request and projection | `spec:drmcp.design_records_mcp.operations.retrieval.exact_record_retrieval`. |
| H2 section selector and section result | `spec:drmcp.design_records_mcp.operations.retrieval.h2_section_retrieval`. |
| Aggregate source-content limits | `spec:drmcp.design_records_mcp.operations.retrieval.retrieval_output_limits`. |
| Shared failure and advisory representation | `spec:drmcp.design_records_mcp.diagnostics`. |
| Detailed source defects and physical source paths | Validation Specifications. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Batch retrieval results | Contract | `spec:drmcp.design_records_mcp.operations.retrieval.batch_retrieval_results` | Defines selector-to-result correspondence, ordering, partial success, and request-wide failure boundaries. |
| Exact record retrieval | Contract | `spec:drmcp.design_records_mcp.operations.retrieval.exact_record_retrieval` | Defines `get_records`, its request-wide projection, and exact-record outcomes. |
| H2 section retrieval | Contract | `spec:drmcp.design_records_mcp.operations.retrieval.h2_section_retrieval` | Defines `get_record_sections`, exact H2 matching, section slices, and section-specific outcomes. |
| Retrieval output limits | Contract | `spec:drmcp.design_records_mcp.operations.retrieval.retrieval_output_limits` | Defines selector-count and aggregate source-content bounds. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.current_record_model` | Supplies uniquely addressable current records. |
| `spec:drmcp.design_records_mcp.artifacts` | Supplies artifact-specific source and projection formats. |
| `spec:drmcp.design_records_mcp.diagnostics` | Supplies shared operation diagnostic objects and canonical codes. |
