# Contract: H2 section retrieval

- **id**: `spec:drcli.design_records_cli.operations.retrieval.h2_section_retrieval`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations.retrieval`
- **contract_class**: `interface`
## What this is

Defines the `get_record_sections` operation for exact retrieval of H2 source sections.

## Non-goals

- Defining a canonical section-ref form outside retrieval.
- Record-level projection selection.
- H3-or-deeper headings as independent selectors.
- Heading repair, normalization, completion, or fuzzy matching.
- Detailed source-defect reporting.
- Physical source-path exposure.

## Request

The request accepts these fields:

| field | requirement | value | default |
|---|---|---|---|
| `section_selectors` | required | Ordered list of `1..100` H2 section-selector strings. | - |
| `source_content_limit_bytes` | optional | Integer source-content limit. | `1048576`. |

A section selector has this retrieval-specific logical form:

```text
<exact-record-ref>#<exact-H2-title>
```

The first `#` separates the record ref from the H2 title.
Every character after the first `#` belongs to the H2 title.
The title may therefore contain additional `#` characters.

Both selector parts must be non-empty.
The record-ref part must satisfy the canonical record-ref grammar.
Transport quoting is not part of selector content.

DRCLI does not trim, strip quotes, repair, normalize, complete, infer, or fuzzy-match either selector part.
A plain record ref without `#` is not a valid section selector.
This selector syntax does not establish a canonical section ref outside retrieval.

Logical input-shape violations are rejected before operation result construction.
Logical input-shape violations include missing required fields, unknown fields, and wrong value types.

### Request validation order

The operation evaluates request-wide outcomes in this order:

1. Selector count.
2. Source-content limit bounds.
3. Selected-repository current-state construction.
4. Selector processing in supplied order.
5. Operation execution.

The first established request-wide error is returned.

## Response

A normal complete response contains `results` as defined by the batch retrieval contract.
The response always contains `applied_source_content_limit_bytes`.

### Successful result

Each successful result contains:

| field | presence | value |
|---|---|---|
| `section_selector` | always | Exact supplied selector string. |
| `record_ref` | always | Parsed record-ref part. |
| `canonical_ref` | always | Resolved current canonical ref. |
| `h2_title` | always | Parsed requested H2 title. |
| `section_content` | always | Complete section string, or `null` only when omitted by the applied limit. |
| `warnings` | when applicable | Ordered operation warnings. |

### Exact H2 matching

DRCLI compares `h2_title` with real source H2 title text.
The comparison is exact.
DRCLI does not trim, normalize case, repair, complete, or fuzzy-match title text.

When exactly one H2 matches, DRCLI returns that section.
When multiple H2 headings match, DRCLI returns the first match in source order.
The duplicate state produces `duplicate_h2`.

### Section content

`section_content` begins with the selected H2 heading line.
The content ends immediately before the next real H2 heading or at end of source.
The content includes every contained H3-or-deeper subsection.

DRCLI preserves the selected source slice without summarization, normalization, reformatting, or truncation.

When the content does not fit the applied limit:

- `section_content` is `null`.
- `warnings` contains `section_content_omitted_by_limit`.
- The warning contains `content_size_bytes`.

## Errors

| condition | classification | result fields established before `error` |
|---|---|---|
| Selector lacks `#`, has an empty part, or has a malformed record-ref part | `malformed_selector` | `section_selector` only. |
| Parsed record ref has no current identity claim | `unresolved_record` | `section_selector`, `record_ref`, `h2_title`. |
| Parsed record ref has multiple current identity claims | `conflicted_record` | `section_selector`, `record_ref`, `h2_title`. |
| Resolved record has no exact H2 title match | `section_not_found` | `section_selector`, `record_ref`, `canonical_ref`, `h2_title`. |
| H2 structure cannot be projected reliably | `section_data_unavailable` | `section_selector`, `record_ref`, `canonical_ref`, `h2_title`. |

### Section-not-found context

`section_not_found.error` contains `available_h2_headings`.
The list uses the same rules as the record H2-heading projection:

- source order;
- exact title text without the leading `## ` marker;
- duplicate occurrences preserved;
- H1 and H3-or-deeper headings excluded;
- no trimming, normalization, or deduplication.

An H2 projection failure produces `section_data_unavailable` instead.
`section_data_unavailable` does not contain `available_h2_headings`.
Detailed source defects belong to validation.

## Warnings

| order | condition | classification | required context |
|---|---|---|---|
| 1 | Multiple exact H2 title matches exist | `duplicate_h2` | None. |
| 2 | Selected section cannot fit the remaining aggregate limit | `section_content_omitted_by_limit` | `content_size_bytes`. |

Only applicable warnings appear.
A successful result without warnings omits the `warnings` field.
The fixed warning order does not depend on parser or execution order.

## Request-level errors

| condition | classification | operation-specific context |
|---|---|---|
| Selector count is outside `1..100` | `invalid_selector_count` | `minimum`, `maximum`, `actual`. |
| Source-content limit is outside the allowed range | `invalid_source_content_limit` | `minimum_bytes`, `maximum_bytes`, `actual_bytes`. |
| Selected-repository current state cannot be constructed reliably | `configuration_failure` | None. |
| The operation cannot produce a reliable complete result | `execution_failure` | None. |

A request-level error returns no partial `results` collection.

## Boundary

| concern | owner |
|---|---|
| Exact current-ref resolution and addressability | `spec:drcli.design_records_cli.current_record_model`. |
| Artifact H2 source shape | Applicable artifact source Specifications. |
| Selector-to-result correspondence | `spec:drcli.design_records_cli.operations.retrieval.batch_retrieval_results`. |
| Section selector, matching, slice, outcomes, and warnings | This Specification. |
| Source-content limit measurement and allocation | `spec:drcli.design_records_cli.operations.retrieval.retrieval_output_limits`. |
| Detailed source defects | Validation Specifications. |
| Shared diagnostics | `spec:drcli.design_records_cli.diagnostics`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.operations.retrieval.batch_retrieval_results` | Supplies common ordered batch behavior. |
| `spec:drcli.design_records_cli.operations.retrieval.retrieval_output_limits` | Supplies aggregate section-content limits. |
| `spec:drcli.design_records_cli.current_record_model` | Defines exact current-ref resolution. |
| `spec:drcli.design_records_cli.artifacts` | Defines shared H2 source shape. |
| `spec:drcli.design_records_cli.diagnostics` | Defines selector-level error and warning placement. |
| `spec:drcli.design_records_cli.diagnostics` | Defines canonical codes and context fields used by this operation. |