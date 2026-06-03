# REQ-DATA-005: DAG asset TypeRef hint render support

- **id**: REQ-DATA-005
- **status**: accepted
- **date**: 2026-06-01
- **source_refs**:
  - ADR-074
  - WORK-DATA-001
  - TASK-DATA-005-01
  - TASK-DATA-005-02
- **work_items**:
  - WORK-DATA-007

## Requirement

The project needs DAG asset TypeRef hint render support so that DAG asset nodes can expose useful top-level type information without changing the underlying TypeRef model.

This requirement captures ADR-074 as a DATA render / view follow-up separate from tagged union implementation and UC-002 notes retreat cleanup.

## Evidence

WORK-DATA-001 explicitly closed M15 without introducing DAG asset node label TypeRef hints, while preserving enum as machine-readable named model information. M15 close notes, REQ-DATA-002, WORK-DATA-002, WORK-DATA-003, and WORK-DATA-004 all kept ADR-074 outside scope.

TASK-DATA-005-01 inventoried ADR-074 as still deferred. TASK-DATA-005-02 classified ADR-074 as DATA render / view support that should receive its own requirement and work item.

## Required Outcome

- Decide and implement the DAG asset TypeRef hint render support described by ADR-074 or its accepted successor wording.
- Keep render hint behavior separate from tagged union support, MCP identity, UC-002 diagnostic blockers, and broad notes retreat cleanup.
- Update the relevant view specs, renderer behavior, fixtures / golden files, and verification evidence through successor work.

## Explicitly Excluded Scope

- Tagged union / discriminator payload support from ADR-073.
- MCP semantic identity / state machine identity from ADR-078 / ADR-079 / ADR-080.
- UC-002 duplicate task QID / unresolved flow task issue.
- Remaining UC-002 notes retreat cleanup.
- Reopening M15 / WORK-DATA-001 / WORK-DATA-002 / WORK-DATA-003 / WORK-DATA-004.

## Boundary

This requirement owns the need for DAG asset TypeRef hint render support. Render output updates were owned and completed by `WORK-DATA-007` and its tasks.

## Close Evidence

Accepted on 2026-06-02.

- ADR-074 was accepted under REQ-DATA-005 / WORK-DATA-007.
- `docs/spec/views/dag.md` was aligned with the accepted asset TypeRef hint label rule.
- WORK-DATA-007 implemented DAG renderer TypeRef hint labels for params boundary assets, task returns, join returns, and foreach collected assets.
- UC-001 DAG golden files were updated for TypeRef hint labels.
- `go test ./...` passed.
- UC-001 / UC-002 validation passed with 0 errors and 0 warnings.
- Explicitly excluded scopes stayed excluded: ADR-073 tagged union support, MCP identity work, UC-002 duplicate task QID repair, remaining notes retreat cleanup, and reopening M15 / WORK-DATA-001〜004.
