# Contract: Exact record validation

- **id**: `spec:drcli.design_records_cli.operations.validation.exact_record_validation`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.operations.validation`
- **contract_class**: `interface`

## What this is

Defines the logical `validate_records` operation for detailed validation of one or more exact current canonical refs.

## Non-goals

- Broad validation of repository, app, artifact-kind, domain, or subtree scopes.
- Selecting a source or tree node by physical path.
- Defining artifact-specific validation rules.
- Defining PRODUCT-owned conformance semantics as independent DRCLI authority.
- Defining the closed finding-code catalog or affected-subject variants owned by applicable validation contracts.
- Returning record source content, metadata, title, or H2 projections.
- Repairing, completing, normalizing, or inferring selectors or source content.
- Defining CLI parsing, help, usage, rendering, streams, or exit status.

## Request

The operation consumes a type-valid logical request with these fields:

| field | requirement | value | default |
|---|---|---|---|
| `record_refs` | required | Ordered list of `1..100` exact current canonical-ref strings. | - |
| `finding_limit` | optional | Maximum returned complete findings across the request. | `200`. |

Duplicate occurrences count toward the `100` selector bound.
DRCLI does not trim, repair, normalize, complete, infer, or fuzzy-match a supplied ref.
CLI syntax and conversion into this logical request belong to `spec:drcli.design_records_cli.cli`.
CLI usage failures are not validation findings or semantic operation diagnostics defined by this contract.

### Selector deduplication

DRCLI compares selector occurrences by exact string equality.
The first occurrence of each exact string is effective.
Later identical occurrences do not produce separate result items.
DRCLI aggregates later indexes into one top-level `duplicate_requested_ref` warning for that exact string.

DRCLI does not trim or normalize a string before duplicate comparison.
The duplicate rule also applies when the effective selector is malformed, unresolved, or conflicted.
Each selected current record is validated at most once per exact supplied string.

### Request processing order

The operation evaluates request-wide and selector outcomes in this order:

1. Selector count.
2. `finding_limit` bounds.
3. Selected-repository current-state construction.
4. Exact-string duplicate aggregation.
5. Effective selector processing in first-occurrence order.
6. Validation of every successfully selected current record.
7. Finding ordering and output-limit application.
8. Reliable operation completion.

The first established request-wide semantic error is returned.
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
The logical response does not prescribe CLI serialization or rendering.

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
| `code` | Stable machine-readable finding code defined by the canonical validation contract that owns the rule. |
| `message` | Non-empty human-readable English explanation. Exact wording is not stable. |
| `contract_ref` | Canonical `spec:` ref of the normative contract that establishes the finding. |
| `subject` | Discriminated object defined by the canonical validation contract that owns the rule. |

The parent result's `record_ref` identifies the affected current record.
A finding does not repeat `record_ref`.
This operation Specification does not define finding-code values, subject variants, source locators, rule composition, or artifact-specific trigger conditions.

The canonical contract referenced by `contract_ref` owns the finding's rule semantics.
DRCLI projects the finding without copying PRODUCT-owned semantics into independent DRCLI authority.
The applicable validation contract supplies semantic deduplication and canonical per-record finding order.
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
| Selected-repository current state cannot be constructed reliably | `configuration_failure` | None. |
| The operation cannot produce a reliable complete validation result | `execution_failure` | None. |

A request-wide semantic error returns exactly one top-level `error`.
The response contains no partial `results`, applied-limit fields, truncation fields, or top-level warnings.
Shared diagnostic objects, canonical codes, and context fields belong to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| Exact current-ref resolution and unique addressability | `spec:drcli.design_records_cli.current_record_model.identity_conflict_and_addressability`. |
| Continued exact-validation eligibility for nonconforming admitted records | `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission`. |
| Exact request, duplicate behavior, result correspondence, and selector outcomes | This Specification. |
| Exact-finding limit bounds and request-wide allocation | `spec:drcli.design_records_cli.operations.validation.validation_output_limits`. |
| Finding codes, affected-subject variants, semantic deduplication, canonical finding order, and conformance semantics | The canonical Specification that owns each validated rule. |
| Shared semantic operation diagnostics | `spec:drcli.design_records_cli.diagnostics`. |
| CLI usage and rendering behavior | `spec:drcli.design_records_cli.cli`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.current_record_model.identity_conflict_and_addressability` | Defines malformed, unresolved, conflicted, and uniquely addressable exact refs. |
| `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission` | Keeps nonconforming admitted records selectable for exact validation. |
| `spec:drcli.design_records_cli.operations.validation.scope_validation` | Discovers abnormal record refs through broad validation. |
| `spec:drcli.design_records_cli.operations.validation.validation_output_limits` | Defines complete-finding limit application. |
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_envelope` | Defines top-level warnings, selector-level errors, and request-wide failure exclusivity. |
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog` | Defines canonical exact-validation diagnostic codes and contexts. |
| `spec:drcli.design_records_cli.cli` | Owns CLI usage diagnostics and concrete output behavior. |
