# WORK-DATA-007: Implement DAG asset TypeRef hint render support

- **id**: WORK-DATA-007
- **status**: done
- **date**: 2026-06-01
- **source_requirement**: REQ-DATA-005
- **impact_refs**:
  - REQ-DATA-005
  - ADR-074
  - WORK-DATA-001
  - TASK-DATA-005-02
- **tasks**:
  - TASK-DATA-007-01
  - TASK-DATA-007-02
  - TASK-DATA-007-03
  - TASK-DATA-007-04

## Goal

Implement DAG asset TypeRef hint render support as a focused DATA render / view successor to ADR-074.

## Boundary

### Included

- Review ADR-074 and decide whether it can be accepted as-is, revised, or split before implementation.
- Define the exact DAG asset label TypeRef hint render boundary.
- Align relevant view specs and render behavior.
- Add or update fixture / golden evidence for asset node and params boundary asset hint rendering.
- Verify that full container TypeRef details remain outside Mermaid label output when required by the accepted boundary.

### Excluded

- ADR-073 tagged union / discriminator payload support.
- ADR-078 / ADR-079 / ADR-080 MCP semantic identity / state machine identity.
- UC-002 duplicate task QID / unresolved flow task issue.
- Remaining UC-002 notes retreat cleanup.
- Reopening M15, WORK-DATA-001, WORK-DATA-002, WORK-DATA-003, or WORK-DATA-004.

## Impact Scope

| layer | current state | handling in this work item |
|---|---|---|
| source requirement | REQ-DATA-005 captured | Owns DAG TypeRef hint render support |
| decision | ADR-074 proposed | Review and accept / revise / split before implementation |
| completed M15 work | WORK-DATA-001 done | Preserve M15 close boundary; implement only as successor work |
| render views | DAG Markdown / Mermaid | Update only the accepted TypeRef hint surfaces |

## Task Flow

Initial task artifacts:

- TASK-DATA-007-01
- TASK-DATA-007-02
- TASK-DATA-007-03
- TASK-DATA-007-04

Expected later split:

```mermaid
flowchart TD
  T1["TASK-DATA-007-01 ADR-074 acceptance / split review"]
  T2["TASK-DATA-007-02 ADR-074 revision, collision inventory, and DAG view spec alignment"]
  T3["TASK-DATA-007-03 renderer and fixture update"]
  T4["TASK-DATA-007-04 verification and close"]
  T1 --> T2 --> T3 --> T4
```

## Completion Condition

This work item is `done`.

Close evidence:

- ADR-074 is accepted under REQ-DATA-005 / WORK-DATA-007.
- `docs/spec/views/dag.md` defines accepted DAG asset TypeRef hint label behavior.
- DAG renderer implementation supports params boundary assets, task returns, join returns, and foreach collected assets.
- UC-001 DAG golden files were updated for TypeRef hint labels.
- `go test ./...` passed.
- UC-001 / UC-002 validate passed with 0 errors and 0 warnings.
- Shortened QID fallback was not implemented.
- No new diagnostics, TypeRef compatibility changes, UC YAML changes, or unrelated work were introduced.
- Close synchronization recorded in TASK-DATA-007-04.
