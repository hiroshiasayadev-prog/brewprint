# V01-REQ-DATA-004: Tagged union and discriminator payload support

- **id**: V01-REQ-DATA-004
- **status**: accepted
- **date**: 2026-06-01
- **source_refs**:
  - V01-ADR-073
  - V01-INV-DATA-002
  - V01-TASK-DATA-005-01
  - V01-TASK-DATA-005-02
- **work_items**:
  - V01-WORK-DATA-010

## Requirement

The project needs explicit tagged union and discriminator payload support so that UC-002 request / response shapes with kind-specific payloads no longer remain hidden behind broad `any + note` descriptions.

This requirement captures V01-ADR-073 as its own DATA expressiveness follow-up rather than burying tagged union work under helper model / model render scope.

## Evidence

M15, V01-WORK-DATA-001, V01-REQ-DATA-002, V01-WORK-DATA-002, V01-WORK-DATA-003, and V01-WORK-DATA-004 all kept V01-ADR-073 outside their implementation scope.

V01-TASK-DATA-003-04 classified UC-002 tagged union / kind-specific response payloads as unchanged candidates. V01-TASK-DATA-005-01 inventoried V01-ADR-073 as still deferred. V01-TASK-DATA-005-02 classified V01-ADR-073 as DATA expressiveness that should receive its own requirement and work item.

V01-WORK-DATA-010 closed this requirement's successor work on 2026-06-03: V01-ADR-073 acceptance / split review, spec and diagnostics alignment, raw YAML / semantic model / resolver / validation implementation, model-file render support, representative UC-002 migration, fixture / golden regeneration, and final verification were completed. Deferred items (`change_contract`, `change_transition_target`, and model catalog renderer implementation) remain outside this requirement close because they require separate expressiveness or renderer work beyond the accepted same-object tagged union minimum.

## Required Outcome

- Decide and implement the accepted tagged union / discriminator payload contract.
- Keep tagged union work separate from helper model render, DAG TypeRef hint, MCP identity, and broad UC-002 cleanup.
- Update the relevant specs, implementation, fixtures, and verification evidence through successor work.

## Explicitly Excluded Scope

- DAG asset TypeRef hint from V01-ADR-074.
- MCP semantic identity / state machine identity from V01-ADR-078 / V01-ADR-079 / V01-ADR-080.
- UC-002 duplicate task QID / unresolved flow task issue.
- Remaining UC-002 notes retreat cleanup beyond tagged-union-specific candidates.
- Reopening M15 / V01-WORK-DATA-001 / V01-WORK-DATA-002 / V01-WORK-DATA-003 / V01-WORK-DATA-004.

## Boundary

This requirement owns the need for tagged union / discriminator payload support. It does not itself decide the final implementation sequence or validation details; those are owned by `V01-WORK-DATA-010` and its tasks.
