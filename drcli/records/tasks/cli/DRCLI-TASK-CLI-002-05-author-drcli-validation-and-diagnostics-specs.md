# DRCLI-TASK-CLI-002-05: Author DRCLI validation and diagnostics specs

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: authoring
- **estimate**: 1d-2d
- **depends_on**:
  - DRCLI-TASK-CLI-002-01
- **outputs**:
  - spec:drcli.design_records_cli.diagnostics
  - spec:drcli.design_records_cli.operations.validation

## Goal

Create DRCLI validation and diagnostic contracts from the current DRMCP baseline.

## Work

Write only:

- `drcli/records/spec/design-records-cli/diagnostics/**`
- `drcli/records/spec/design-records-cli/operations/validation/**`

Use current DRMCP validation and diagnostics as the primary baseline.

Preserve:

- broad-scope and exact-record validation separation;
- complete-state outcome despite bounded projections;
- contract violations, advisories, unadmitted sources, identity conflicts, selector failures, configuration failures, and execution failures;
- deterministic findings and limits;
- no repair, normalization, completion, or inference.

Separate semantic validation diagnostics from CLI usage diagnostics.
Concrete CLI help/usage rendering belongs to T08.

## Done condition

- Validation and diagnostic Specification areas are complete.
- Machine-readable diagnostic meaning is transport-independent.
- Validation semantics remain PRODUCT-owned where applicable.
- No CLI parser behavior is silently mixed into validation findings.

## Verification

- Compare against current DRMCP validation specs and diagnostic catalog.
- Confirm output limits cannot change the semantic outcome.
- Confirm assigned paths are the only modified files.

## Evidence

- Authored `spec:drcli.design_records_cli.diagnostics` under the assigned `diagnostics/**` writer boundary.
- Authored `spec:drcli.design_records_cli.operations.validation` under the assigned `operations/validation/**` writer boundary.
- Preserved DRMCP broad-scope versus exact-record validation separation, semantic diagnostic codes, deterministic ordering, and bounded complete-item projections.
- Preserved complete-state `ok` and summary semantics when result or finding projections are limited.
- Kept validation finding semantics with each canonical owning Specification through `contract_ref`; DRCLI only projects those findings.
- Separated semantic operation diagnostics from CLI usage/parser diagnostics; concrete help, usage, streams, exit status, and rendering remain owned by T08.
- Verification script completed with exit code 0: seven assigned specs had path-derived IDs, required Contract sections, no stale `command_runtime` refs, no MCP-specific input-schema/tool wording, all 26 DRMCP baseline diagnostic codes, and explicit complete-outcome limit invariants.
- Authoring writes were limited to the two assigned Specification path families plus this Task's required lifecycle/Evidence update. DRMCP and the parent Work Item were not modified.
