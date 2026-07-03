# TRV-TASK-SPEC-001-15: Execute architecture-derived contract Specification Work Item

- **id**: TRV-TASK-SPEC-001-15
- **status**: blocked
- **date**: 2026-07-03
- **work_item**: TRV-WORK-SPEC-001
- **work_item_ref**: TRV-WORK-SPEC-005
- **task_type**: work_item_execution
- **estimate**: 2d
- **depends_on**:
  - TRV-TASK-SPEC-001-14
- **outputs**:
  - TRV-TASK-SPEC-001-15

## Goal

Represent TRV-WORK-SPEC-005 as one architecture-derived contract-Specification execution unit in the parent graph.

## Work

- Observe TRV-WORK-SPEC-005 without duplicating its Task graph or deliverables.
- Remain incomplete while the child Work Item is `not_started` or `in_progress`.
- Use `blocked` only when the child blocks the accepted parent route.
- Record the child terminal status after TRV-WORK-SPEC-005 becomes `done`.

This Task must not create, split, or redefine the child Work Item; change the graph; perform child-owned investigation, decision, authoring, review, or synchronization; or start implementation.

## Done condition

- `work_item_ref` identifies TRV-WORK-SPEC-005.
- TRV-WORK-SPEC-005 is `done`.
- Evidence records the observed child status.
- No child-owned detail is duplicated.

## Verification

- Confirm the parent owns this Task.
- Confirm W005 exists and differs from the parent Work Item.
- Confirm T14 is `done` before this Task starts.
- Confirm W005 is `done` before this Task becomes `done`.
- Confirm no graph or child deliverable changed.

## Evidence

- T14 created this execution relation after W005 existed.
- TRV-WORK-SPEC-005 is `blocked` by TRV-ADR-SPEC-006.
- T14 released the placement-first W005 route before semantic-validator delivery was suspended.
- This Task remains blocked until a later decision restores W005 execution.
