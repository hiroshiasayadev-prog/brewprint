# Contract: DRCLI deferred proposal

- **id**: `spec:drcli.design_records_cli.operations.authoring.deferred_proposal`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.operations.authoring`
- **contract_class**: interface

## What this is

Defines the optional process-independent retained proposal path for callers that explicitly want authoring preparation and repository modification to occur in separate invocations.
It is not the normal create or update path.

## Request

Deferred proposal behavior is entered only by explicit caller choice during create or update preparation.
Absence of that choice uses direct write.

A retained proposal represents one prepared create or update candidate and contains, semantically:

- an opaque proposal identity;
- the repository against which the proposal was prepared;
- the authoring operation kind;
- requested and resolved target identity;
- the complete candidate affected set or sufficient retained material to reconstruct it exactly;
- base-state evidence needed for stale-write detection;
- candidate-state validation result;
- proposal lifecycle state;
- validity information under the installation's retention policy.

The concrete persisted representation is not part of the contract.

Proposal state is process-independent and caller-session-independent.
One shared DRCLI installation may retain proposals for multiple repositories.
Each proposal is bound to the repository against which it was prepared; later inspection, discard, or acceptance must not silently redirect it to another repository.

The proposal lifecycle supports these semantic actions:

| action | contract |
|---|---|
| Prepare | Build and validate the candidate without modifying repository files. |
| Inspect | Retrieve retained proposal identity, state, target, affected-set information, and candidate validation sufficient to understand the pending change. |
| Discard | Mark an unaccepted retained proposal unusable for later acceptance without modifying repository files. |
| Accept | Attempt the retained write through the shared write-completion contract. |

Proposal preparation never modifies repository files.

The retained proposal may expose a preview of the candidate change, but concrete diff format, output rendering, and command spelling belong outside this contract.

A proposal identity is opaque.
Callers must not infer repository, path, timestamp, or storage location from its spelling.

A proposal remains usable while it is retained, not discarded, not already accepted, and valid under the installation's retention policy.
This contract does not choose a concrete retention duration.

## Response

Successful preparation returns a non-writing outcome with an opaque proposal identity and the proposal's validated target and affected-set semantics.

Inspection returns the current retained lifecycle state and pending candidate information without modification.

Successful discard returns a non-writing discarded outcome.
Discard does not undo a proposal that was already accepted.

Acceptance delegates to the shared write contract and therefore re-checks, immediately before modification:

- repository binding;
- target and affected-record staleness;
- create identity and placement availability when applicable;
- update target identity when applicable;
- PRODUCT-required coupled changes;
- candidate-state validation.

Successful acceptance records the proposal as accepted and reports the actual write result.

A stale proposal cannot be force-accepted.
The caller must prepare a new candidate against current repository state.

## Errors

| condition | machine-readable classification | required effect |
|---|---|---|
| Proposal identity is unknown | `authoring_proposal_unavailable` (`reason=unknown`), exit 3 | Reject without repository modification. |
| Proposal is expired or otherwise no longer valid | `authoring_proposal_unavailable` (`reason=expired` or `invalid`), exit 3 | Reject without repository modification. |
| Proposal was discarded | `authoring_proposal_ineligible` (`state=discarded`), exit 1 on attempted acceptance | Reject acceptance without modification. |
| Proposal was already accepted | `authoring_proposal_ineligible` (`state=accepted`), exit 1 on attempted acceptance | Do not repeat the write. |
| Proposal repository does not match the repository against which it was prepared | `authoring_repository_mismatch` (`phase=accept` and supplied `proposal_id`), exit 1 | Reject without modification; no retargeting. |
| Target or affected state became stale | `authoring_stale_state` (`guard=target_state`, `target_identity`, `affected_set`, or `target_resolution`), exit 1 | Reject without modification and require a new proposal for the changed state. |
| Create identity or placement is no longer available | `authoring_identity_collision` (`collision_kind=canonical_identity` or `generated_placement`), or `authoring_stale_state` (`guard=allocation`), exit 1 | Reject without modification. |
| Candidate validation now fails | `authoring_candidate_validation_failed`, exit 1 | Reject without modification and preserve canonical finding identities. |

Failure to reliably read retained proposal state uses `authoring_operational_failure` (`reason=retained_state_io`), exit 3, instead of falsely claiming an unknown/expired/invalid identity. Failure to reliably persist a newly prepared proposal or lifecycle state uses `authoring_operational_failure` (`reason=proposal_persistence`), exit 3. Failure of repository inspection or write completion during acceptance is classified by the shared write contract; the returned write-state must remain truthful. These operational failures do not silently mutate proposal binding or imply successful acceptance.

A failed acceptance does not become a successful write merely because proposal preparation previously passed.

## Boundary

This contract owns optional retained candidate lifecycle and process-independent proposal state.
Create and update own candidate preparation inputs.
The shared write contract owns final stale-state, validation, and modification eligibility.

Mandatory propose-then-accept, MCP session affinity, storage engine choice, fixed retention duration, command flags, and CLI output formatting are outside this contract.

## Related specs

| ref | relation |
|---|---|
| spec:drcli.design_records_cli.operations.authoring | Parent authoring overview. |
| spec:drcli.design_records_cli.operations.authoring.create | Create preparation. |
| spec:drcli.design_records_cli.operations.authoring.update | Update preparation. |
| spec:drcli.design_records_cli.operations.authoring.write | Acceptance write guards. |
| DRMCP-REQ-MCP-002 | Baseline for retained proposal and accept-time behavior. |
