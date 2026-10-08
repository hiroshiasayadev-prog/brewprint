# Contract: Exact record retrieval

- **id**: `spec:drmcp.design_records_mcp.operations.retrieval.exact_record_retrieval`
- **status**: draft
- **date**: 2026-07-14
- **parent**: `spec:drmcp.design_records_mcp.operations.retrieval`
- **contract_class**: `interface`
- **usdm_covers**:
  - usdm:drmcp.mcp_capabilities.retrieval.record_content_retrieval#R001,R006-R012,R030

## What this is

Defines the `get_records` operation for exact current-record retrieval with one request-wide projection.

## Non-goals

- H2 section selectors.
- Record listing, search, reference repair, or legacy lookup.
- Detailed source-defect reporting.
- Physical source-path exposure.
- Source-content budget rules shared with section retrieval.

## Request

The request accepts these fields:

| field | requirement | value | default |
|---|---|---|---|
| `record_refs` | required | Ordered list of `1..100` exact current canonical-ref strings. | - |
| `projection` | required | Exact projection object defined below. | - |
| `source_content_limit_bytes` | optional | Integer source-content limit. Ignored with a warning when `projection.source_content` is `false`. | `1048576` only when source content is requested. |

`record_refs` accepts record refs only.
A section-selector string is not a valid record ref.
DRMCP does not trim, repair, complete, normalize, infer, or fuzzy-match a supplied ref.

Input-schema violations are rejected before operation result construction.
Schema violations include missing required fields, unknown fields, wrong value types, and malformed projection shape.

### Projection

`projection` contains exactly these required boolean fields:

| field | selected output |
|---|---|
| `title` | H1 title text. |
| `metadata` | Parsed H1-adjacent metadata. |
| `h2_headings` | Real H2 title list. |
| `source_content` | Complete source document. |

Unknown projection fields are prohibited.
All four projection values may be `false`.
An all-false projection requests canonical-ref-only results.

### Request validation order

The operation evaluates request-wide outcomes in this order:

1. Selector count.
2. Source-content limit applicability.
3. Source-content limit bounds when source content is requested.
4. Configured current-state construction.
5. Selector processing in supplied order.
6. Operation execution.

A non-applicable source-content limit records a top-level warning and does not stop this order.
The first established request-wide error is returned.

When `projection.source_content` is `false`, a supplied `source_content_limit_bytes` value is ignored without range validation.
The operation preserves normal selector processing and returns top-level `source_content_limit_ignored`.

## Response

A normal complete response contains `results` as defined by the batch retrieval contract.
The response contains `applied_source_content_limit_bytes` only when `projection.source_content` is `true`.
The response contains top-level `warnings` only when a request-wide warning applies.

### Successful result

Each successful result contains:

| field | presence | value |
|---|---|---|
| `record_ref` | always | Exact supplied selector string. |
| `canonical_ref` | always | Resolved current canonical ref. |
| `title` | when requested | H1 title text or `null`. |
| `metadata` | when requested | Parsed metadata object or `null`. |
| `h2_headings` | when requested | Ordered H2 title list, empty list, or `null`. |
| `source_content` | when requested | Complete source string, or `null` only when omitted by the applied limit. |
| `warnings` | when applicable | Ordered operation warnings. |

An unrequested projection field is absent.
A requested parsed projection field is present even when its value is unavailable.
Unavailable parsed projection data does not make the selector unsuccessful.
Detailed source defects belong to validation.

### Title projection

`title` is the source text after the H1's first `: ` delimiter.
The projection excludes the H1 marker and identity prefix.
DRMCP does not trim, normalize, or reformat the title text.
An unavailable title produces `title: null` without a retrieval warning.

### Metadata projection

`metadata` is an unordered JSON object.
Each object key is one successfully parsed source field name.

| metadata source form | projected value |
|---|---|
| `scalar` | String. |
| `inline_list` | Ordered string list. |
| `indented_list` | Ordered string list. |

An optional source field that is absent produces no object key.
A metadata block that cannot be projected reliably produces `metadata: null`.
DRMCP does not expose malformed duplicates, source positions, or detailed conformance findings.

### H2 heading projection

`h2_headings` contains real H2 title text in source order.
Each value excludes the leading `## ` marker.
The projection preserves exact title text and duplicate occurrences.
The projection excludes H1 and H3-or-deeper headings.
DRMCP does not trim, normalize, or deduplicate heading titles.

An empty list means the source was projected successfully with no real H2 headings.
A `null` value means DRMCP could not construct the heading projection reliably.
Neither state produces a retrieval warning.

### Source-content projection

`source_content` contains the complete source document.
DRMCP does not summarize, normalize, reformat, or truncate the content.

When the content does not fit the applied limit:

- `source_content` is `null`.
- `warnings` contains `source_content_omitted_by_limit`.
- The warning contains `content_size_bytes`.

A source read failure that prevents reliable completion is an execution failure.
DRMCP does not represent that failure as a successful `source_content: null` projection.

## Selector-level errors

| condition | classification | result fields established before `error` |
|---|---|---|
| Supplied string is not a valid canonical record-ref form | `malformed_selector` | `record_ref` only. |
| Valid ref has no current identity claim | `unresolved_record` | `record_ref` only. |
| Valid ref has multiple current identity claims | `conflicted_record` | `record_ref` only. |

A malformed section-selector form supplied to this operation is `malformed_selector`.
The result omits `canonical_ref` when exact resolution did not succeed.

## Warnings

| placement | condition | classification | required context |
|---|---|---|---|
| Top-level | `source_content_limit_bytes` is supplied while source content is unrequested | `source_content_limit_ignored` | `source_content_limit_bytes`. |
| Selector-level | Requested complete source cannot fit the remaining aggregate limit | `source_content_omitted_by_limit` | `content_size_bytes`. |

`source_content_limit_ignored` preserves all selector processing and may coexist with selector-level errors.
The response omits `applied_source_content_limit_bytes` when the supplied limit is ignored.
`get_records` currently has at most one warning at each placement.
Shared warning representation belongs to Diagnostics Specifications.

## Request-level errors

| condition | classification | operation-specific context |
|---|---|---|
| Selector count is outside `1..100` | `invalid_selector_count` | `minimum`, `maximum`, `actual`. |
| Source-content limit is outside the allowed range while source content is requested | `invalid_source_content_limit` | `minimum_bytes`, `maximum_bytes`, `actual_bytes`. |
| Configured current state cannot be constructed reliably | `configuration_failure` | None. |
| The operation cannot produce a reliable complete result | `execution_failure` | None. |

A request-level error returns no partial `results` collection or top-level `warnings`.

## Boundary

| concern | owner |
|---|---|
| Exact current-ref resolution and addressability | `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability`. |
| Continued retrieval eligibility for nonconforming admitted records | `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission`. |
| Artifact-specific metadata and source formats | Applicable artifact Specifications. |
| Selector-to-result correspondence | `spec:drmcp.design_records_mcp.operations.retrieval.batch_retrieval_results`. |
| Request fields and record projections | This Specification. |
| Source-content limit measurement and allocation | `spec:drmcp.design_records_mcp.operations.retrieval.retrieval_output_limits`. |
| Detailed source defects | Validation Specifications. |
| Shared diagnostics | `spec:drmcp.design_records_mcp.diagnostics`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.operations.retrieval.batch_retrieval_results` | Supplies common ordered batch behavior. |
| `spec:drmcp.design_records_mcp.operations.retrieval.retrieval_output_limits` | Supplies aggregate source-content limits. |
| `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability` | Defines exact current-ref resolution. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.h1_adjacent_metadata` | Defines shared metadata source forms. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.source` | Defines shared H1 and H2 source shape. |
| `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_envelope` | Defines top-level and selector-level diagnostic placement. |
| `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_catalog` | Defines canonical codes and context fields used by this operation. |
