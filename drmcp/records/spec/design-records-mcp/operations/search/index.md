# Overview: Search operations

- **id**: `spec:drmcp.design_records_mcp.operations.search`
- **status**: draft
- **date**: 2026-07-14
- **parent**: `spec:drmcp.design_records_mcp.operations`

## What this is

Defines scoped lexical search over uniquely addressable current Design Records.

## Current contract

| operation | purpose |
|---|---|
| `search_records` | Search selected current-record scopes across selected source-text targets. |

The operation name is also the MCP tool name.

## Non-goals

- Semantic search, fuzzy matching, synonym expansion, stemming, or morphological analysis.
- Canonical-ref search, completion, repair, relation traversal, or graph search.
- Legacy search.
- Complete record or H2 section retrieval.
- Validation findings, physical source paths, or repair guidance.
- Search-index, parser, cache, storage, watcher, or concurrency architecture.

## Boundary

| concern | owner |
|---|---|
| Current scope availability and included addressable records | `spec:drmcp.design_records_mcp.current_record_model`. |
| Artifact H1, metadata, and H2 source boundaries | `spec:drmcp.design_records_mcp.artifacts`. |
| Search request, scope collection, result envelope, and operation outcomes | `spec:drmcp.design_records_mcp.operations.search.lexical_record_search`. |
| Matching modes, target selection, source slices, and occurrence rules | `spec:drmcp.design_records_mcp.operations.search.lexical_match_model`. |
| Match projection, snippets, ordering, and output limits | `spec:drmcp.design_records_mcp.operations.search.search_results_and_limits`. |
| Complete record and H2 section content | `spec:drmcp.design_records_mcp.operations.retrieval`. |
| Shared diagnostic representation and canonical codes | `spec:drmcp.design_records_mcp.diagnostics`. |
| Detailed source defects and physical source paths | Validation Specifications. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Lexical record search | Contract | `spec:drmcp.design_records_mcp.operations.search.lexical_record_search` | Defines `search_records`, scope selection, request validation, response fields, and operation outcomes. |
| Lexical match model | Contract | `spec:drmcp.design_records_mcp.operations.search.lexical_match_model` | Defines literal and RE2 pattern matching across selectable source-text targets. |
| Search results and limits | Contract | `spec:drmcp.design_records_mcp.operations.search.search_results_and_limits` | Defines match entries, snippets, deterministic ordering, output limits, and additional-match indication. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.current_record_model` | Supplies current scopes and uniquely addressable current records. |
| `spec:drmcp.design_records_mcp.artifacts` | Supplies source boundaries consumed by search targets. |
| `spec:drmcp.design_records_mcp.operations.retrieval` | Owns complete record and H2 section retrieval. |
| `spec:drmcp.design_records_mcp.diagnostics` | Supplies shared operation diagnostic objects and canonical codes. |
