# Contract: Validation output limits

- **id**: `spec:drcli.design_records_cli.operations.validation.validation_output_limits`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.operations.validation`
- **contract_class**: `interface`

## What this is

Defines request-wide logical output limits for broad abnormal results and exact-record findings without changing complete validation outcomes.

## Non-goals

- Transport-wide response-byte limits.
- Partial truncation of one abnormal result or finding.
- Selector-count limits.
- Limiting request errors, selector-level errors, operation warnings, or scope aggregates.
- Defining broad-result or finding ordering.
- Pagination, offsets, cursors, or continuation tokens.
- Defining CLI option spelling, parsing, output rendering, streams, or exit status.

## Request

The validation operations consume these optional logical limit fields:

| operation | field | measured output | minimum | default | maximum |
|---|---|---|---:|---:|---:|
| `validate_scope` | `result_limit` | Complete abnormal result items. | `1` | `200` | `2000` |
| `validate_records` | `finding_limit` | Complete findings across the request. | `1` | `200` | `2000` |

A caller-specified value is not clamped.
A type-valid integer outside the accepted range produces request-level `invalid_limit`.
CLI decoding of an argument into this type-valid logical field belongs to `spec:drcli.design_records_cli.cli`.
Limit validation precedes selected-repository current-state construction.

## Response

### Broad validation limit

A normal `validate_scope` response always contains:

```text
applied_result_limit
total_results
results_truncated
```

| field | value |
|---|---|
| `applied_result_limit` | Effective `result_limit`. |
| `total_results` | Complete abnormal result count before output limiting. |
| `results_truncated` | Whether one or more complete abnormal results were omitted by the limit. |

`total_results` counts every complete result variant:

```text
unadmitted_source
identity_conflict
record
tree_node
```

`scope_validation` supplies the deterministic complete result order.
DRCLI retains the first `applied_result_limit` items from that ordered collection.
An abnormal result is returned completely or omitted completely.

The complete scope summary and top-level `ok` remain unchanged by output limiting.
`results_truncated` is `true` exactly when `total_results` is greater than the returned `results` count.
Reaching the limit remains normal success.

### Exact validation limit

A normal `validate_records` response always contains:

```text
applied_finding_limit
total_findings
findings_truncated
```

| field | value |
|---|---|
| `applied_finding_limit` | Effective `finding_limit`. |
| `total_findings` | Complete finding count across every successfully selected current record. |
| `findings_truncated` | Whether one or more complete findings were omitted by the limit. |

DRCLI constructs the request-wide finding order as follows:

```text
effective selector first-occurrence order
  -> applicable validation contract's canonical finding order for that record
```

DRCLI retains the first `applied_finding_limit` findings from that complete order.
A finding is returned completely or omitted completely.
Selector-level errors and top-level operation warnings do not consume the finding limit.

Each successful record result retains its complete `summary` and `ok` values.
A successful result always contains `findings`, including when every finding for that record is omitted.
The response does not contain a per-record truncation field.
A consumer can compare the record summary count with the returned `findings` count.

`findings_truncated` is `true` exactly when `total_findings` is greater than the total returned finding count.
A response with no successfully selected current record uses `total_findings: 0` and `findings_truncated: false`.
Reaching the limit remains normal success.

## Errors

| condition | classification | context |
|---|---|---|
| Applicable validation output limit is outside `1..2000` | `invalid_limit` | `minimum`, `maximum`, `default`, `actual`. |
| Reliable ordered limit application or complete response construction fails | `execution_failure` | None. |

For `invalid_limit`, the context values are:

| field | value |
|---|---|
| `minimum` | `1` |
| `maximum` | `2000` |
| `default` | `200` |
| `actual` | Exact supplied integer in the logical operation request. |

A request-wide error returns no normal result fields, applied-limit fields, truncation fields, or top-level warnings.
Shared diagnostic representation and context contracts belong to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| Broad abnormal result variants and deterministic ordering | `spec:drcli.design_records_cli.operations.validation.scope_validation`. |
| Exact selector order and successful result placement | `spec:drcli.design_records_cli.operations.validation.exact_record_validation`. |
| Canonical per-record finding order and semantic deduplication | The canonical validation contract that owns each finding rule. |
| Numeric bounds, complete-item application, and truncation indicators | This Specification. |
| Shared semantic operation errors | `spec:drcli.design_records_cli.diagnostics`. |
| CLI option mapping and output presentation | `spec:drcli.design_records_cli.cli`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.operations.validation.scope_validation` | Supplies the complete ordered broad abnormal result collection. |
| `spec:drcli.design_records_cli.operations.validation.exact_record_validation` | Supplies effective selector order and successful record result envelopes. |
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_envelope` | Defines request-wide error exclusivity. |
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog` | Defines `invalid_limit` and `execution_failure`. |
| `spec:drcli.design_records_cli.cli` | Owns CLI option mapping and concrete output behavior. |
