# Overview: Validation operations

- **id**: `spec:drcli.design_records_cli.operations.validation`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations`

## What this is

Defines transport-independent broad-scope validation and detailed validation of exact current records for DRCLI.

## Current contract

| logical operation | purpose |
|---|---|
| `validate_scope` | Validate one repository, app, artifact-kind, sequential-domain, or tree-subtree scope and return aggregates with abnormal subject summaries. |
| `validate_records` | Validate one or more exact current canonical refs and return detailed findings for every successfully selected current record. |

Broad-scope validation and exact-record validation are independent operations.
The logical operation names do not define CLI command spelling or hierarchy.
Command mapping and presentation belong to `spec:drcli.design_records_cli.cli`.

## Non-goals

- Defining artifact-specific conformance rules.
- Defining validation rule components or PRODUCT-owned finding semantics.
- Defining metadata-field, H2-section, reference, or structure locator schemas owned by other Specifications.
- Repairing, completing, normalizing, or inferring nonconforming source content.
- Selecting unadmitted sources or tree nodes by physical source path.
- Defining CLI argument parsing, help, usage, output rendering, stream placement, or exit status.
- Defining parser, cache, storage, watcher, concurrency, or filesystem architecture.

## Boundary

| concern | owner |
|---|---|
| Current source, admission, addressability, scope, and tree-node semantics | `spec:drcli.design_records_cli.current_record_model`. |
| Base and artifact-specific normative conformance rules | The canonical Design Records Specifications that own those rules. |
| Broad-scope request, abnormal summaries, aggregates, and operation outcomes | `spec:drcli.design_records_cli.operations.validation.scope_validation`. |
| Exact-record request, selector outcomes, and detailed finding placement | `spec:drcli.design_records_cli.operations.validation.exact_record_validation`. |
| Broad-result and exact-finding output limits | `spec:drcli.design_records_cli.operations.validation.validation_output_limits`. |
| Shared semantic operation errors and warnings | `spec:drcli.design_records_cli.diagnostics`. |
| CLI usage diagnostics and concrete command/output behavior | `spec:drcli.design_records_cli.cli`. |

Validation operations select subjects and project outcomes.
The canonical Specification owning a validated rule defines its semantics.
DRCLI does not acquire independent authority over PRODUCT-owned rules by reporting their findings.

## Topic map

Broad validation discovers abnormal subjects without returning detailed record findings.
Exact validation returns detailed findings for explicitly selected current records.
Output limiting changes only the returned projection, never the complete semantic outcome.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Scope validation | Contract | `spec:drcli.design_records_cli.operations.validation.scope_validation` | Defines broad selectors, abnormal result summaries, aggregates, ordering, and request-wide outcomes. |
| Exact record validation | Contract | `spec:drcli.design_records_cli.operations.validation.exact_record_validation` | Defines exact selector handling, detailed finding placement, duplicate handling, and selector-level errors. |
| Validation output limits | Contract | `spec:drcli.design_records_cli.operations.validation.validation_output_limits` | Defines request-wide broad-result and exact-finding limits without partial result objects or semantic truncation. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.current_record_model` | Supplies current sources, scopes, tree nodes, and uniquely addressable current records. |
| `spec:drcli.design_records_cli.diagnostics` | Supplies shared semantic operation diagnostic objects and canonical codes. |
| `spec:drcli.design_records_cli.cli` | Maps logical operations to the external CLI and owns usage/rendering behavior. |
