# Overview: Retrieval operations

- **id**: `spec:drcli.design_records_cli.operations.retrieval`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations`

## What this is

Defines exact current-record retrieval and exact H2 section retrieval.

## Current contract

| operation | purpose |
|---|---|
| `get_records` | Retrieve exact current records with one request-wide projection. |
| `get_record_sections` | Retrieve exact H2 sections from exact current records. |

Operation names in this area are semantic identifiers. CLI command spelling and option names are defined elsewhere.
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
| Current-record admission and unique addressability | `spec:drcli.design_records_cli.current_record_model`. |
| Artifact metadata, H1, H2, and source formats | `spec:drcli.design_records_cli.artifacts`. |
| Shared batch-result invariants | `spec:drcli.design_records_cli.operations.retrieval.batch_retrieval_results`. |
| Exact record request and projection | `spec:drcli.design_records_cli.operations.retrieval.exact_record_retrieval`. |
| H2 section selector and section result | `spec:drcli.design_records_cli.operations.retrieval.h2_section_retrieval`. |
| Aggregate source-content limits | `spec:drcli.design_records_cli.operations.retrieval.retrieval_output_limits`. |
| Shared failure and advisory representation | `spec:drcli.design_records_cli.diagnostics`. |
| Detailed source defects and physical source paths | Validation Specifications. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Batch retrieval results | Contract | `spec:drcli.design_records_cli.operations.retrieval.batch_retrieval_results` | Defines selector-to-result correspondence, ordering, partial success, and request-wide failure boundaries. |
| Exact record retrieval | Contract | `spec:drcli.design_records_cli.operations.retrieval.exact_record_retrieval` | Defines `get_records`, its request-wide projection, and exact-record outcomes. |
| H2 section retrieval | Contract | `spec:drcli.design_records_cli.operations.retrieval.h2_section_retrieval` | Defines `get_record_sections`, exact H2 matching, section slices, and section-specific outcomes. |
| Retrieval output limits | Contract | `spec:drcli.design_records_cli.operations.retrieval.retrieval_output_limits` | Defines selector-count and aggregate source-content bounds. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.current_record_model` | Supplies uniquely addressable current records. |
| `spec:drcli.design_records_cli.artifacts` | Supplies artifact-specific source and projection formats. |
| `spec:drcli.design_records_cli.diagnostics` | Supplies shared operation diagnostic objects and canonical codes. |