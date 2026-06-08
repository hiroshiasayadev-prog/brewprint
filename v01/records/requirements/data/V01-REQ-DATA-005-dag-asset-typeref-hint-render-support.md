# V01-REQ-DATA-005: DAG asset TypeRef hint render support

- **id**: V01-REQ-DATA-005
- **status**: accepted
- **date**: 2026-06-01
- **source_refs**:
  - V01-ADR-074
  - V01-WORK-DATA-001
  - V01-TASK-DATA-005-01
  - V01-TASK-DATA-005-02
- **work_items**:
  - V01-WORK-DATA-007

## Requirement

The project needs DAG asset TypeRef hint render support so that DAG asset nodes can expose useful top-level type information without changing the underlying TypeRef model.

This requirement captures V01-ADR-074 as a DATA render / view follow-up separate from tagged union implementation and UC-002 notes retreat cleanup.

## Evidence

V01-WORK-DATA-001 explicitly closed M15 without introducing DAG asset node label TypeRef hints, while preserving enum as machine-readable named model information. M15 close notes, V01-REQ-DATA-002, V01-WORK-DATA-002, V01-WORK-DATA-003, and V01-WORK-DATA-004 all kept V01-ADR-074 outside scope.

V01-TASK-DATA-005-01 inventoried V01-ADR-074 as still deferred. V01-TASK-DATA-005-02 classified V01-ADR-074 as DATA render / view support that should receive its own requirement and work item.

## Required Outcome

- Decide and implement the DAG asset TypeRef hint render support described by V01-ADR-074 or its accepted successor wording.
- Keep render hint behavior separate from tagged union support, MCP identity, UC-002 diagnostic blockers, and broad notes retreat cleanup.
- Update the relevant view specs, renderer behavior, fixtures / golden files, and verification evidence through successor work.

## Explicitly Excluded Scope

- Tagged union / discriminator payload support from V01-ADR-073.
- MCP semantic identity / state machine identity from V01-ADR-078 / V01-ADR-079 / V01-ADR-080.
- UC-002 duplicate task QID / unresolved flow task issue.
- Remaining UC-002 notes retreat cleanup.
- Reopening M15 / V01-WORK-DATA-001 / V01-WORK-DATA-002 / V01-WORK-DATA-003 / V01-WORK-DATA-004.

## Boundary

This requirement owns the need for DAG asset TypeRef hint render support. Render output updates were owned and completed by `V01-WORK-DATA-007` and its tasks.

## Close Evidence

Accepted on 2026-06-02.

- V01-ADR-074 was accepted under V01-REQ-DATA-005 / V01-WORK-DATA-007.
- `docs/spec/views/dag.md` was aligned with the accepted asset TypeRef hint label rule.
- V01-WORK-DATA-007 implemented DAG renderer TypeRef hint labels for params boundary assets, task returns, join returns, and foreach collected assets.
- UC-001 DAG golden files were updated for TypeRef hint labels.
- `go test ./...` passed.
- UC-001 / UC-002 validation passed with 0 errors and 0 warnings.
- Explicitly excluded scopes stayed excluded: V01-ADR-073 tagged union support, MCP identity work, UC-002 duplicate task QID repair, remaining notes retreat cleanup, and reopening M15 / V01-WORK-DATA-001〜004.
