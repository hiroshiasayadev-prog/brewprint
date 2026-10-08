# Contract: DRCLI write completion

- **id**: `spec:drcli.design_records_cli.operations.authoring.write`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.operations.authoring`
- **contract_class**: interface

## What this is

Defines the shared pre-write guards and write-completion semantics used by direct create, direct update, and deferred-proposal acceptance.
It does not define command spelling, persistence technology, or PRODUCT-owned record semantics.

## Request

The write boundary receives a prepared authoring candidate with these semantic inputs.

| input | requirement |
|---|---|
| Selected repository | The repository against which the candidate was prepared and must be written. |
| Operation kind | Create or update. |
| Target identity | Requested and resolved canonical target identity. |
| Candidate affected set | Every record whose persisted content would change. |
| Base-state evidence | Sufficient state to detect target, allocation, identity, or affected-record staleness before modification. |
| Candidate validation result | Validation of the complete affected set in candidate state. |

The concrete representation of base-state evidence is an implementation detail.
It may not weaken the required stale-write detection behavior.

Before any repository modification begins, DRCLI must:

1. confirm that the selected repository is the repository for which the candidate was prepared;
2. re-resolve the target identity and applicable generated placement under current PRODUCT authority;
3. detect stale target state, changed target identity, and create-time identity or placement collisions;
4. re-establish every PRODUCT-required coupled change in the affected set against current repository state;
5. re-run or equivalently re-establish required candidate-state validation for the complete affected set;
6. confirm that no blocking diagnostic applies to the affected candidate;
7. only then begin repository modification.

Unrelated repository diagnostics outside the affected set do not become blocking authoring diagnostics.

A failed pre-write guard returns a non-writing outcome and must not modify any repository file.

Where PRODUCT authority requires a coupled multi-record lifecycle or relation change to be atomic, this write boundary must preserve that requirement.
This contract does not generalize those cases into an arbitrary multi-record transaction facility.

## Response

A successful write outcome communicates, independent of output formatting:

- that repository modification occurred;
- the resolved target identity;
- the complete set of records or files actually modified;
- the validation state relevant to write eligibility;
- any diagnostics that remain applicable after completion.

A pre-write rejection communicates:

- that no repository modification occurred;
- the target identity when it could be resolved safely;
- the failed guard or validation condition;
- an empty modified-set result.

If an execution failure is discovered after repository content was actually modified, the result must not falsely report that no write occurred.
This requirement preserves truthful write-state reporting and does not define rollback mechanics.

## Errors

| condition | machine-readable classification | required effect |
|---|---|---|
| Prepared repository does not match the selected repository | `authoring_repository_mismatch` (`phase=prepared_write` for direct write or `phase=accept` for proposal acceptance), exit 1 | Reject before modification; preserve the proposal binding. |
| Target state changed after candidate preparation | `authoring_stale_state` (`guard=target_state`), exit 1 | Reject before modification. |
| Create identity or generated placement is no longer available | `authoring_identity_collision` (`collision_kind=canonical_identity` or `generated_placement`) if occupied; otherwise `authoring_stale_state` (`guard=allocation`), exit 1 | Reject before modification. |
| Existing update target no longer resolves to the same record | `authoring_stale_state` (`guard=target_identity`), exit 1 | Reject before modification. |
| PRODUCT-required coupled affected set changed incompatibly | `authoring_stale_state` (`guard=affected_set`), exit 1 | Reject before modification. |
| Candidate-state validation now has a blocking error | `authoring_candidate_validation_failed`, exit 1 | Reject before modification; retain PRODUCT-owned finding identities. |
| Required target or affected record cannot be resolved safely | `authoring_stale_state` (`guard=target_resolution`), exit 1 when a changed/missing resolved identity is established; `authoring_operational_failure` (`reason=repository_state`), exit 3 when trustworthy resolution cannot be produced | Reject before modification. |
| Modification has already occurred before an execution failure is detected | `authoring_operational_failure` (`reason=write_completion`, `write_state=modified`), exit 3 | Report actual observed write state; never convert it into a false non-writing result. |

If write execution, repository I/O, or permission checks prevent a trustworthy result, the primary diagnostic is `authoring_operational_failure` (`reason=repository_state` before write, or `reason=write_completion` after write has begun), exit 3. Its required `write_state` and any confirmed modified identities are derived from actual evidence. A successfully rejected pre-write guard has `write_state=not_modified` and `modified_refs=[]` in its error context. The selected authoring failure code takes priority over the generic non-authoring `execution_failure` code.

There is no force-write path that bypasses blocking candidate validation or stale-target guards.

## Boundary

Create and update own request-specific candidate preparation.
Deferred proposal owns retained candidate lifecycle.
This contract owns the final eligibility check immediately before modification and the truthfulness of the write outcome.

The storage engine, locking mechanism, hash scheme, timestamp scheme, rollback implementation, command mapping, and rendered output are outside this contract.

## Related specs

| ref | relation |
|---|---|
| spec:drcli.design_records_cli.operations.authoring | Parent authoring overview. |
| spec:drcli.design_records_cli.operations.authoring.create | Create candidate preparation. |
| spec:drcli.design_records_cli.operations.authoring.update | Update candidate preparation. |
| spec:drcli.design_records_cli.operations.authoring.deferred_proposal | Deferred acceptance path. |
| DRMCP-REQ-MCP-002 | Baseline for candidate validation and accept-time guards. |
