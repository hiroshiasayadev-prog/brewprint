# Overview: Search operations

- **id**: `spec:drcli.design_records_cli.operations.search`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations`

## What this is

Defines scoped lexical search over uniquely addressable current Design Records.

## Current contract

| operation | purpose |
|---|---|
| `search_records` | Search selected current-record scopes across selected source-text targets. |

The operation name is a semantic identifier. CLI command spelling and option names are defined elsewhere.

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
| Current scope availability and included addressable records | `spec:drcli.design_records_cli.current_record_model`. |
| Artifact H1, metadata, and H2 source boundaries | `spec:drcli.design_records_cli.artifacts`. |
| Search request, scope collection, result envelope, and operation outcomes | `spec:drcli.design_records_cli.operations.search.lexical_record_search`. |
| Matching modes, target selection, source slices, and occurrence rules | `spec:drcli.design_records_cli.operations.search.lexical_match_model`. |
| Match projection, snippets, ordering, and output limits | `spec:drcli.design_records_cli.operations.search.search_results_and_limits`. |
| Complete record and H2 section content | `spec:drcli.design_records_cli.operations.retrieval`. |
| Shared diagnostic representation and canonical codes | `spec:drcli.design_records_cli.diagnostics`. |
| Detailed source defects and physical source paths | Validation Specifications. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Lexical record search | Contract | `spec:drcli.design_records_cli.operations.search.lexical_record_search` | Defines `search_records`, scope selection, request validation, response fields, and operation outcomes. |
| Lexical match model | Contract | `spec:drcli.design_records_cli.operations.search.lexical_match_model` | Defines literal and RE2 pattern matching across selectable source-text targets. |
| Search results and limits | Contract | `spec:drcli.design_records_cli.operations.search.search_results_and_limits` | Defines match entries, snippets, deterministic ordering, output limits, and additional-match indication. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.current_record_model` | Supplies current scopes and uniquely addressable current records. |
| `spec:drcli.design_records_cli.artifacts` | Supplies source boundaries consumed by search targets. |
| `spec:drcli.design_records_cli.operations.retrieval` | Owns complete record and H2 section retrieval. |
| `spec:drcli.design_records_cli.diagnostics` | Supplies shared operation diagnostic objects and canonical codes. |