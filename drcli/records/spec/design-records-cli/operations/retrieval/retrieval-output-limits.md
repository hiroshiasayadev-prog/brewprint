# Contract: Retrieval output limits

- **id**: `spec:drcli.design_records_cli.operations.retrieval.retrieval_output_limits`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations.retrieval`
- **contract_class**: `interface`
## What this is

Defines aggregate source-content limits shared by exact record and H2 section retrieval.

## Non-goals

- Limiting metadata, title, H2-heading, selector, diagnostic, or envelope output.
- Partial source-content truncation.
- Selector-count limits.
- Transport-wide response-byte limits.
- Shared diagnostic serialization.

## Request

The request field is `source_content_limit_bytes`.
The field measures source content as UTF-8 encoded bytes.

| constraint | value |
|---|---|
| minimum | `1` byte |
| default | `1048576` bytes |
| maximum | `16777216` bytes |

A caller-specified value is not clamped.
A value outside the range produces request-level `invalid_source_content_limit` when the field is applicable.
Selector processing does not begin after this error.
The response contains no `results` or top-level `warnings` after this error.

The error context contains:

| field | presence | value |
|---|---|---|
| `minimum_bytes` | always | `1`. |
| `maximum_bytes` | always | `16777216`. |
| `actual_bytes` | always | Exact supplied integer. |

A wrong input type is a logical input-shape violation.
Logical input-shape validation rejects the request before this result contract applies.

## Applicability

| operation | field availability | applied-limit response field |
|---|---|---|
| `get_records` with `projection.source_content: true` | Optional. Default applies when omitted. | `applied_source_content_limit_bytes` is present. |
| `get_records` with `projection.source_content: false` | Accepted but ignored without range validation. | Response field is absent. |
| `get_record_sections` | Optional. Default applies when omitted. | `applied_source_content_limit_bytes` is present. |

Supplying the field when record source content is unrequested produces top-level `source_content_limit_ignored`.
The warning preserves selector processing and every otherwise valid result.
The warning contains the exact supplied `source_content_limit_bytes` integer.
The ignored value is not range-validated.

## Response

The aggregate budget counts only returned source-content strings:

- `get_records.source_content`;
- `get_record_sections.section_content`.

The UTF-8 byte count includes every returned character.
The count includes source line endings, spaces, and other preserved source text.

The budget excludes:

- canonical refs and supplied selectors;
- titles;
- metadata;
- H2-heading lists;
- result envelopes;
- warnings and errors;
- `applied_source_content_limit_bytes`.

The aggregate limit applies independently to each request.

## Allocation order

DRCLI evaluates source content in supplied selector order.
For each selector, DRCLI compares the complete content size with the remaining budget.

| condition | action |
|---|---|
| Complete content fits | Return the complete content and consume its UTF-8 byte size. |
| Complete content does not fit | Omit that content, consume no budget, and continue with the next selector. |

A skipped large content does not prevent a later smaller content from using the remaining budget.
DRCLI does not reorder selectors to maximize returned content.
DRCLI does not stop selector processing after an omission.

## All-or-nothing content

Each requested content value is returned completely or omitted completely.
DRCLI does not truncate, summarize, normalize, or reformat content to fit the budget.

The operation-specific result keeps the content field present:

| operation | omitted field value | warning classification |
|---|---|---|
| `get_records` | `source_content: null` | `source_content_omitted_by_limit` |
| `get_record_sections` | `section_content: null` | `section_content_omitted_by_limit` |

Each omission warning contains `content_size_bytes`.
The value is the complete content's UTF-8 byte size.

An omission caused by the applied limit remains a successful selector result.
The omission does not suppress requested metadata, title, or H2-heading projections.

## Aggregate response indicators

The response does not include:

- `has_more`;
- an aggregate omission flag;
- an omission count;
- consumed-byte count;
- remaining-byte count.

Each result warning identifies its own omission.
All supplied selectors are processed within the normal complete-response boundary.

## Errors

| condition | classification | behavior |
|---|---|---|
| Applicable source-content limit is outside `1..16777216` | `invalid_source_content_limit` | Reject before selector processing. |
| Source-content limit has the wrong logical value type | Logical input-shape violation | Reject before this result contract applies. |

Request-wide configuration and execution failures are owned by the consuming retrieval operation.
This contract does not add a separate failure outcome for budget allocation.

## Boundary

| concern | owner |
|---|---|
| Selector-count limit | `spec:drcli.design_records_cli.operations.retrieval.batch_retrieval_results`. |
| Source-content field applicability, measurement, numeric bounds, and allocation | This Specification. |
| Complete record content | `spec:drcli.design_records_cli.operations.retrieval.exact_record_retrieval`. |
| Complete H2 section content | `spec:drcli.design_records_cli.operations.retrieval.h2_section_retrieval`. |
| Shared warning and error representation | `spec:drcli.design_records_cli.diagnostics`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.operations.retrieval.batch_retrieval_results` | Supplies selector order and partial-success behavior. |
| `spec:drcli.design_records_cli.operations.retrieval.exact_record_retrieval` | Consumes the budget for complete record source. |
| `spec:drcli.design_records_cli.operations.retrieval.h2_section_retrieval` | Consumes the budget for complete H2 section source. |
| `spec:drcli.design_records_cli.diagnostics` | Defines top-level and selector-level warning placement. |
| `spec:drcli.design_records_cli.diagnostics` | Defines source-content limit and omission diagnostic contexts. |