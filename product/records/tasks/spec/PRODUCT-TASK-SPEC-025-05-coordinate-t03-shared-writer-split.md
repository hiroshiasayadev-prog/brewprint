# PRODUCT-TASK-SPEC-025-05: Coordinate T03 shared-writer split into parallel investigation Tasks

- **id**: PRODUCT-TASK-SPEC-025-05
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-025
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-025-01
- **outputs**:
  - PRODUCT-TASK-SPEC-025-03
  - PRODUCT-TASK-SPEC-025-04
  - PRODUCT-TASK-SPEC-025-05
  - PRODUCT-TASK-SPEC-025-06
  - PRODUCT-TASK-SPEC-025-07
  - PRODUCT-TASK-SPEC-025-08
  - PRODUCT-TASK-SPEC-025-09
  - PRODUCT-TASK-SPEC-025-10
  - PRODUCT-WORK-SPEC-025

## Goal

Repair the T03 shared-writer conflict by replacing one sequential-session investigation Task with five independent, parallel-safe investigation Tasks, and re-route T04's dependency accordingly.

## Work

- Cancel `PRODUCT-TASK-SPEC-025-03`; record the shared-writer conflict as the cancellation reason.
- Create `PRODUCT-TASK-SPEC-025-06` through `PRODUCT-TASK-SPEC-025-10`, one per corpus range (W001–W007, W009–W012, W013–W017, W018–W019, W020–W025), each `investigation`, each depending only on `PRODUCT-TASK-SPEC-025-02`, each owning its own Evidence as its own write target.
- Update `PRODUCT-TASK-SPEC-025-04`'s `depends_on` to the five new investigation Tasks, and expand its Work to include cross-log coverage reconciliation and deduplication (the responsibility the cancelled Task's "Session 6" previously owned).
- Update `PRODUCT-WORK-SPEC-025`'s `tasks` list, Task flow, and Task Candidates table to match the new graph.

This Task does not extract vocabulary entries and does not author `skills/task-boundary-vocabulary/` content.

## Done condition

- `PRODUCT-TASK-SPEC-025-03` is `cancelled` with the write-conflict reason recorded.
- Five new investigation Tasks exist, each with a non-overlapping corpus range and no shared Evidence write target.
- `PRODUCT-TASK-SPEC-025-04` depends on all five new Tasks and owns cross-log reconciliation.
- `PRODUCT-WORK-SPEC-025` reflects the corrected Task list and flow.

## Verification

- Confirmed the five new investigation Tasks' corpus ranges are disjoint and jointly cover the same scope the cancelled Task covered.
- Confirmed no two Tasks write to the same Evidence target.
- Confirmed `PRODUCT-TASK-SPEC-025-04`'s dependency list and Work were updated.
- Confirmed the parent Work Item Task list and flow match the new graph.

## Evidence

- The user requested a true Task split instead of sequential sessions, specifically to avoid a write conflict from parallel execution against one shared Evidence target.
- `PRODUCT-TASK-SPEC-025-03` was `not_started` at cancellation, so cancellation from `not_started` is valid under the Task status lifecycle.
- `PRODUCT-TASK-SPEC-025-04` was the only directly dependent Task and was `not_started`; this coordination step re-points it to the five new Tasks in the same atomic operation, so it is not left `blocked`.
- Created `PRODUCT-TASK-SPEC-025-06` through `PRODUCT-TASK-SPEC-025-10` (investigation, one per corpus range, each independently write-safe).
- Updated `PRODUCT-TASK-SPEC-025-04` dependency and Work to absorb cross-log reconciliation.
- Updated `PRODUCT-WORK-SPEC-025` Task list, flow, and Task Candidates table.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
