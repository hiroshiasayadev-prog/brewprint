# USDM requirement: Lexical record search

- **id**: `usdm:drmcp.mcp_capabilities.search.lexical_record_search`
- **status**: draft
- **date**: 2026-07-12
- **kind**: requirement
- **parent**: `usdm:drmcp.mcp_capabilities.search`

## What this is

Product requirements for scoped lexical search across current Design Records, selectable search targets, match context, and bounded result output.

## Requirements: Search scope
> source: literal

| id | requirement | notes |
|---|---|---|
| R001 | When no search scope is supplied, DRMCP must search uniquely addressable current records across all configured app namespaces. |  |
| R002 | DRMCP must allow the caller to constrain search to one or more current-record scopes. | Available scopes are exposed by the discovery capability. |
| R003 | DRMCP must support a sequential-record search scope composed of an app namespace, an artifact kind, and a domain namespace. |  |
| R004 | DRMCP must support a tree-record search scope composed of an app namespace, an artifact kind, and a subtree root ref. |  |
| R005 | When multiple search scopes are supplied, DRMCP must search the union of their records and evaluate each overlapping record at most once. |  |
| R006 | When a valid search scope contains no searchable record, DRMCP must return a normal empty search result. |  |
| R007 | When a supplied search scope is unavailable or cannot be selected, DRMCP must make that outcome identifiable to the caller. | Exact diagnostic representation is defined by downstream Specifications. |

## Requirements: Lexical matching
> source: literal

| id | requirement | notes |
|---|---|---|
| R008 | DRMCP must support literal substring matching. |  |
| R009 | DRMCP must support pattern-based lexical matching. | The concrete pattern syntax and dialect are defined by downstream Specifications. |
| R010 | DRMCP must allow the caller to select whether lexical matching is case-sensitive. | The default behavior is defined by downstream Specifications. |
| R011 | DRMCP must allow H2 heading text to be selected as a search target. |  |
| R012 | DRMCP must allow H2 section content to be selected as a search target. | The H2 heading line is excluded from this target. H3 and deeper subsection content remains included. |
| R013 | DRMCP must allow record title text to be selected as a search target. |  |
| R014 | DRMCP must allow H1-adjacent metadata source text to be selected as a search target. | The searchable text includes metadata field names and field values. |
| R015 | DRMCP must allow the caller to combine multiple search targets in one search request. | The concrete request shape is defined by downstream Specifications. |
| R016 | DRMCP must determine lexical matches from source text without semantic rewriting or inferred synonyms. |  |
| R017 | DRMCP must not perform fuzzy matching, semantic similarity matching, canonical-ref completion, or reference repair as part of lexical search. | Exact retrieval, reference resolution, and semantic search are separate capabilities. |

## Requirements: Search result context
> source: literal

| id | requirement | notes |
|---|---|---|
| R018 | DRMCP must return the canonical ref of the record containing each reported match. |  |
| R019 | DRMCP must identify the selected search target in which each reported match occurred. | Search targets include H2 heading text, H2 section content, record title text, and H1-adjacent metadata source text. |
| R020 | DRMCP must identify the enclosing H2 heading for a match within H2 section content. | The heading can be used with H2 section retrieval. |
| R021 | DRMCP must return a bounded source snippet containing each reported match. |  |
| R022 | DRMCP must preserve matched source text in a snippet without summarization, normalization, or fabricated replacement text. |  |
| R023 | DRMCP must support reporting multiple matches from one record. | Per-record and aggregate result limits still apply. |
| R024 | DRMCP must return search results in a deterministic order. | The concrete ordering rule is defined by downstream Specifications. |
| R025 | DRMCP must not expose physical source paths in normal search results. |  |

## Requirements: Search output limits
> source: literal

| id | requirement | notes |
|---|---|---|
| R026 | DRMCP must impose an upper bound on the total number of search matches returned by one request. |  |
| R027 | DRMCP must allow the caller to specify the maximum number of search matches to return. | The numeric limit and request shape are defined by downstream Specifications. |
| R028 | DRMCP must apply a default search-result limit when the caller does not specify one. |  |
| R029 | DRMCP must enforce a server-side maximum for caller-specified search-result limits. |  |
| R030 | DRMCP must impose an upper bound on the number of matches returned from one record. | One record must not consume the complete result set. |
| R031 | DRMCP must allow the caller to control the maximum snippet size within a server-side maximum. | The measurement unit and numeric limits are defined by downstream Specifications. |
| R032 | DRMCP must make it identifiable when additional matches may exist beyond an applied result limit. |  |
| R033 | DRMCP must either return one complete search-result entry or omit that entry. | An aggregate output limit must not partially truncate a result entry. |
| R034 | DRMCP must not return complete record content or complete H2 section content as a search result. | Complete content is obtained through the retrieval capability. |
