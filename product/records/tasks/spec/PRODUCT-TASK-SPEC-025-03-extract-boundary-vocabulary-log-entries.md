# PRODUCT-TASK-SPEC-025-03: Extract boundary-vocabulary log entries from done Tasks

- **id**: PRODUCT-TASK-SPEC-025-03
- **status**: cancelled
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-025
- **task_type**: investigation
- **estimate**: 3d
- **depends_on**:
  - PRODUCT-TASK-SPEC-025-02
- **outputs**: []

## Goal

Extract boundary-vocabulary log entries from every `product/records/tasks/spec/` Task that is `status: done` at this Task's corpus cutoff, and record them directly in this Task's own `## Evidence`, under the Investigation-Task lightweight Evidence exception (PRODUCT-TASK-SPEC-025-02).

## Done condition

Superseded. See Evidence.

## Evidence

- This Task was cancelled before starting. Its single-Evidence, multi-session design created a shared-writer conflict when the user wanted the corpus range extraction to run in parallel.
- `PRODUCT-TASK-SPEC-025-05` (coordination) replaced this Task with five independent, parallel-safe investigation Tasks (`PRODUCT-TASK-SPEC-025-06` through `PRODUCT-TASK-SPEC-025-10`), one per corpus range, each with its own Evidence write target.
- `PRODUCT-TASK-SPEC-025-04` was re-pointed to depend on the five replacement Tasks instead of this one, in the same coordination step, so it is not left `blocked`.
- No work was performed under this Task before cancellation.
