# Contract: Exact record validation

- **id**: `spec:drmcp.design_records_mcp.operations.validation.exact_record_validation`
- **status**: draft
- **date**: 2026-07-14
- **parent**: `spec:drmcp.design_records_mcp.operations.validation`
- **contract_class**: `interface`
- **usdm_covers**:
  - usdm:drmcp.mcp_capabilities.validation.validation_scope_and_subjects#R005,R011
  - usdm:drmcp.mcp_capabilities.validation.validation_result_projection#R001-R002,R011,R013-R016,R023-R024

## What this is

Defines the `validate_records` operation for detailed validation of one or more exact current canonical refs.

## Non-goals

- Broad validation of repository, app, artifact-kind, domain, or subtree scopes.
- Selecting a source or tree node by physical path.
- Defining artifact-specific validation rules.
- Defining the closed finding-code catalog or affected-subject variants.
- Returning record source content, metadata, title, or H2 projections.
- Repairing, completing, normalizing, or inferring selectors or source content.

## Request

The request accepts these fields:

| field | requirement | value | default |
|---|---|---|---|
| `record_refs` | required | Ordered list of `1..100` exact current canonical-ref strings. | - |
| `finding_limit` | optional | Maximum returned complete findings across the request. | `200`. |

Duplicate occurrences count toward the `100` selector bound.
DRMCP does not trim, repair, normalize, complete, infer, or fuzzy-match a supplied ref.

Input-schema violations are rejected before operation result construction.
Schema violations include missing required fields, unknown fields, wrong value types, a non-array `record_refs`, non-string selector items, and a non-integer `finding_limit`.

### Selector deduplication

DRMCP compares selector occurrences by exact string equality.
The first occurrence of each exact string is effective.
Later identical occurrences do not produce separate result items.
DRMCP aggregates later indexes into one top-level `duplicate_requested_ref` warning for that exact string.

DRMCP does not trim or normalize a string before duplicate comparison.
The duplicate rule also applies when the effective selector is malformed, unresolved, or conflicted.
Each selected current record is validated at most once per exact supplied string.

### Request processing order

The operation evaluates request-wide and selector outcomes in this order:

1. Selector count.
2. `finding_limit` bounds.
3. Configured current-state construction.
4. Exact-string duplicate aggregation.
5. Effective selector processing in first-occurrence order.
6. Validation of every successfully selected current record.
7. Finding ordering and output-limit application.
8. Reliable operation completion.

The first established request-wide error is returned.
One selector-level error does not stop later effective selectors.

## Response

A normal response uses this logical shape:

```text
results:
  - <successful-record-result-or-selector-error>
applied_finding_limit
total_findings
findings_truncated
warnings?
```

The response does not contain a top-level `ok`.
`warnings` is present only when at least one duplicate-selector warning applies.
Every normal response contains `results`, including a response in which every effective selector fails.

### Result correspondence and ordering

Each effective selector produces exactly one result item.
Result order follows the first occurrence order of effective selectors.
Success results and selector-level errors are not grouped or reordered by state.
Every result preserves the exact effective selector string as `record_ref`.

A later duplicate occurrence produces no result item.
The duplicate warning identifies the omitted occurrences.

### Successful record result

A successful result uses this logical shape:

```text
record_ref
ok
summary:
  contract_violations
  advisories
findings:
  - <validation-finding>
```

| field | value |
|---|---|
| `record_ref` | Exact supplied effective selector string. |
| `ok` | Whether the complete record validation contains no contract-violation finding. |
| `summary.contract_violations` | Complete contract-violation finding count for the selected record. |
| `summary.advisories` | Complete advisory finding count for the selected record. |
| `findings` | Returned complete findings after request-wide limit application. |

The result does not repeat a `canonical_ref` field.
A successful exact canonical selector already identifies the selected current record.

`ok` is `false` when `summary.contract_violations` is greater than zero.
Advisory-only validation produces `ok: true`.
Both summary counts and `ok` describe the complete record validation, including findings omitted by `finding_limit`.
`findings` is always present for a successful result and may be empty.

### Validation finding envelope

Each returned finding uses this common logical envelope:

```text
classification
code
message
contract_ref
subject
```

| field | value |
|---|---|
| `classification` | `contract_violation` or `advisory`. |
| `code` | Stable machine-readable finding code defined by the applicable Validation finding contract. |
| `message` | Non-empty human-readable English explanation. Exact wording is not stable. |
| `contract_ref` | Canonical `spec:` ref of the normative contract that establishes the finding. |
| `subject` | Discriminated object defined by the applicable Validation finding contract. |

The parent result's `record_ref` identifies the affected current record.
A finding does not repeat `record_ref`.
The operation Specification does not define finding-code values, subject variants, source locators, rule composition, or artifact-specific trigger conditions.

The applicable Validation finding contract supplies semantic deduplication and canonical per-record finding order.
The operation preserves that order.

### Selector-level errors

An unsuccessful effective selector result contains `record_ref` and exactly one `error`.

| condition | classification |
|---|---|
| Supplied string is not a valid canonical record-ref form | `malformed_selector` |
| Valid ref has no current identity claim | `unresolved_record` |
| Valid ref has multiple current identity claims | `conflicted_record` |

A selector-level error omits `ok`, `summary`, and `findings`.
A selector-level error does not remove successful results for other effective selectors.
A request in which every effective selector fails remains a normal complete response with `total_findings: 0` and `findings_truncated: false`.

## Warnings

`duplicate_requested_ref` is the only top-level warning defined by this operation.

The warning context uses this logical shape:

```text
ref
first_index
duplicate_indexes:
  - <zero-based-index>
```

| field | value |
|---|---|
| `ref` | Exact supplied duplicate selector string. |
| `first_index` | Zero-based index of the effective first occurrence. |
| `duplicate_indexes` | Every later zero-based occurrence index in ascending order. |

The operation emits one warning per exact duplicate string.
Warnings use ascending `first_index` order.
A duplicate warning may coexist with successful results and selector-level errors.

## Errors

| condition | classification | operation-specific context |
|---|---|---|
| Selector count is outside `1..100` | `invalid_selector_count` | `minimum`, `maximum`, `actual`. |
| `finding_limit` is outside `1..2000` | `invalid_limit` | `minimum`, `maximum`, `default`, `actual`. |
| Configured current state cannot be constructed reliably | `configuration_failure` | None. |
| The operation cannot produce a reliable complete validation result | `execution_failure` | None. |

A request-wide error returns exactly one top-level `error`.
The response contains no partial `results`, applied-limit fields, truncation fields, or top-level warnings.
Shared diagnostic objects, canonical codes, context fields, and messages belong to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| Exact current-ref resolution and unique addressability | `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability`. |
| Continued exact-validation eligibility for nonconforming admitted records | `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission`. |
| Exact request, duplicate behavior, result correspondence, and selector outcomes | This Specification. |
| Exact-finding limit bounds and request-wide allocation | `spec:drmcp.design_records_mcp.operations.validation.validation_output_limits`. |
| Finding codes, affected-subject variants, semantic deduplication, and canonical finding order | Other Validation Specifications. |
| Artifact-specific rules and rule composition | Applicable artifact and Validation Specifications. |
| Shared operation diagnostics | `spec:drmcp.design_records_mcp.diagnostics`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability` | Defines malformed, unresolved, conflicted, and uniquely addressable exact refs. |
| `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission` | Keeps nonconforming admitted records selectable for exact validation. |
| `spec:drmcp.design_records_mcp.operations.validation.scope_validation` | Discovers abnormal record refs through broad validation. |
| `spec:drmcp.design_records_mcp.operations.validation.validation_output_limits` | Defines complete-finding limit application. |
| `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_envelope` | Defines top-level warnings, selector-level errors, and request-wide failure exclusivity. |
| `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_catalog` | Defines canonical exact-validation diagnostic codes and contexts. |
