# REQ-DATA-004: Tagged union and discriminator payload support

- **id**: REQ-DATA-004
- **status**: captured
- **date**: 2026-06-01
- **source_refs**:
  - ADR-073
  - INV-DATA-002
  - TASK-DATA-005-01
  - TASK-DATA-005-02
- **work_items**:
  - WORK-DATA-010

## Requirement

The project needs explicit tagged union and discriminator payload support so that UC-002 request / response shapes with kind-specific payloads no longer remain hidden behind broad `any + note` descriptions.

This requirement captures ADR-073 as its own DATA expressiveness follow-up rather than burying tagged union work under helper model / model render scope.

## Evidence

M15, WORK-DATA-001, REQ-DATA-002, WORK-DATA-002, WORK-DATA-003, and WORK-DATA-004 all kept ADR-073 outside their implementation scope.

TASK-DATA-003-04 classified UC-002 tagged union / kind-specific response payloads as unchanged candidates. TASK-DATA-005-01 inventoried ADR-073 as still deferred. TASK-DATA-005-02 classified ADR-073 as DATA expressiveness that should receive its own requirement and work item.

## Required Outcome

- Decide and implement the accepted tagged union / discriminator payload contract.
- Keep tagged union work separate from helper model render, DAG TypeRef hint, MCP identity, and broad UC-002 cleanup.
- Update the relevant specs, implementation, fixtures, and verification evidence through successor work.

## Explicitly Excluded Scope

- DAG asset TypeRef hint from ADR-074.
- MCP semantic identity / state machine identity from ADR-078 / ADR-079 / ADR-080.
- UC-002 duplicate task QID / unresolved flow task issue.
- Remaining UC-002 notes retreat cleanup beyond tagged-union-specific candidates.
- Reopening M15 / WORK-DATA-001 / WORK-DATA-002 / WORK-DATA-003 / WORK-DATA-004.

## Boundary

This requirement owns the need for tagged union / discriminator payload support. It does not itself decide the final implementation sequence or validation details; those are owned by `WORK-DATA-010` and its tasks.
