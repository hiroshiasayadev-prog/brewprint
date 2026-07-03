# TRV-TASK-SPEC-001-09: Execute TRV application contract Work Item

- **id**: TRV-TASK-SPEC-001-09
- **status**: blocked
- **date**: 2026-07-03
- **work_item**: TRV-WORK-SPEC-001
- **work_item_ref**: TRV-WORK-SPEC-003
- **task_type**: work_item_execution
- **estimate**: 2d
- **depends_on**:
  - TRV-TASK-SPEC-001-08
- **outputs**:
  - TRV-TASK-SPEC-001-09

## Goal

Represent TRV-WORK-SPEC-003 as one application-contract execution unit in the parent graph.

## Work

- Observe TRV-WORK-SPEC-003 without duplicating its Task graph or deliverables.
- Remain incomplete while the child Work Item is `not_started` or `in_progress`.
- Use `blocked` only when the child blocks the accepted parent route.
- Record the child terminal status after TRV-WORK-SPEC-003 becomes `done`.

This Task must not create, split, or redefine the child Work Item; change any graph or release order; perform child-owned work; or create implementation work.

## Done condition

- `work_item_ref` identifies TRV-WORK-SPEC-003.
- TRV-WORK-SPEC-003 is `done`.
- Evidence records the observed child status.
- No child-owned detail is duplicated.

## Verification

- Confirm the parent owns this Task.
- Confirm the child exists and differs from the parent Work Item.
- Confirm T08 is `done` before this Task starts.
- Confirm the child is `done` before this Task becomes `done`.
- Confirm no graph or child deliverable changed.

## Evidence

- T07 created this execution relation after TRV-WORK-SPEC-003 existed.
- W003 was operationally retired after its graph omitted the required architecture-derived Specification-placement decision.
- W003 remains `blocked` and will not satisfy this Task's Done condition.
- T13 blocked this historical execution relation.
- T14 created T15 as the active parent execution relation for replacement TRV-WORK-SPEC-005.
- Parent T10 now depends on T15 rather than this Task.
