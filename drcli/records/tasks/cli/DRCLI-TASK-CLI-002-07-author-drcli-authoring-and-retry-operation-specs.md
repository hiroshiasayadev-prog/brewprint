# DRCLI-TASK-CLI-002-07: Author DRCLI authoring and retry operation specs

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: authoring
- **estimate**: 2d
- **depends_on**:
  - DRCLI-TASK-CLI-002-01
- **outputs**:
  - spec:drcli.design_records_cli.operations.authoring
  - spec:drcli.design_records_cli.operations.authoring.create
  - spec:drcli.design_records_cli.operations.authoring.update
  - spec:drcli.design_records_cli.operations.authoring.write
  - spec:drcli.design_records_cli.operations.authoring.retry_cache
  - spec:drcli.design_records_cli.operations.authoring.deferred_proposal

## Goal

Create DRCLI create, update, write, retry-cache, and optional deferred-proposal operation contracts.

## Work

Write only:

- `drcli/records/spec/design-records-cli/operations/authoring/**`

Use DRMCP-REQ-MCP-002 and old DRMCP authoring-transaction contracts as design baselines, but adapt to DRCLI-REQ-CLI-001.

Required DRCLI differences:

- normal successful create/update is one CLI invocation and does not require mandatory propose-then-accept;
- optional deferred proposal behavior may remain for explicit callers;
- body sources support inline argument, stdin, or existing body_cache_id as defined by the operation;
- retryable failure after receiving body preserves that body and returns opaque body_cache_id;
- cache/proposal state survives process and caller-session termination;
- one shared DRCLI installation may retain state for multiple repositories;
- accept/write paths re-check stale targets and required validation before modification;
- authoring remains non-interactive by default;
- no temporary-file upload step is required.

Do not choose concrete command flags, database engines, or CLI output formatting.
T08 owns CLI mapping and runtime-visible common behavior.

## Done condition

- Create and partial-update contracts cover supported Design Records kinds according to PRODUCT authoring authority.
- Direct write and optional deferred-proposal paths are unambiguous.
- Retry cache semantics preserve the useful DRMCP behavior.
- Candidate-state validation and stale-write guards are retained.
- No MCP session dependency remains.

## Verification

- Trace every DRCLI-REQ-CLI-001 authoring/body/cache outcome.
- Compare retained proposal/cache semantics with DRMCP baselines.
- Confirm only `operations/authoring/**` changed.

## Evidence

- Authored the DRCLI authoring specification subtree under drcli/records/spec/design-records-cli/operations/authoring/.
- Covered create and partial update for ADR, Requirement, Work Item, Task, Specification, and Investigation by reference to PRODUCT-owned authoring authority rather than duplicating per-kind semantics.
- Recorded direct same-invocation write as the normal path and retained deferred proposals only for explicit callers.
- Recorded inline, standard-input, and body-cache body sources; retryable received-body preservation; process-independent cache/proposal state; and shared-installation multi-repository isolation.
- Centralized affected-candidate validation, stale-target checks, create collision checks, and required pre-write validation in the shared write contract.
- Added an explicit DRCLI-REQ-CLI-001 authoring/body/cache trace table and a DRMCP baseline adaptation table.
- Verified all five child contracts contain the required interface Contract sections and the authoring overview declares every child in its Topics table.
- Specification writes were confined to drcli/records/spec/design-records-cli/operations/authoring/**. The only non-specification write was this Task's permitted lifecycle, outputs, and Evidence update.
- No parent Work Item lifecycle, sibling lane output, root DRCLI spec index, DRMCP record, implementation file, or Git state was modified.
- No independent review was performed in this authoring Task.
