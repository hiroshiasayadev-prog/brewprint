# PRODUCT-TASK-SPEC-023-07: Coordinate post-ADR-routing authoring and review

- **id**: PRODUCT-TASK-SPEC-023-07
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-023
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-023-06
- **outputs**:
  - PRODUCT-WORK-SPEC-023
  - PRODUCT-TASK-SPEC-023-07
  - PRODUCT-TASK-SPEC-023-08
  - PRODUCT-TASK-SPEC-023-09
  - PRODUCT-TASK-SPEC-023-10
  - PRODUCT-TASK-SPEC-023-11

## Goal

Materialize the exact ADR, canonical authoring, workflow-support, integrated review, and accepted-route closure graph from T06.

## Work

- Read T06 routing outcomes and ADR boundaries.
- Create exact ADR authoring Tasks required by T06.
- Create exact Specification and workflow-support authoring Tasks required by T06.
- Serialize every shared writer and preserve ADR-before-Specification order where required.
- Create one integrated independent review after the final writer.
- Create direct closure synchronization only when the accepted review route and writable targets can be stated without speculation.
- Keep correction and finding-closure review Tasks unmaterialized until named findings exist.
- Update Work Item tasks, Task flow, Task Candidates, and this Task outputs.

This Task must not:

- alter ADR routing outcomes;
- author ADR, Specification, workflow-support, checklist, or implementation content;
- perform review, correction, synchronization, implementation, stage, or commit work.

## Done condition

- Every T06 `required` or direct-projection outcome has one exact writer.
- Writer order is deterministic and shared files have one owner.
- One integrated independent review follows the final writer.
- Closure materialization is limited to an exact accepted-input route.
- No speculative finding-repair Task exists.
- Work Item and Task relations are coherent.

## Verification

- Confirm T06 is complete.
- Confirm every created writer traces to an exact T06 route.
- Confirm ADR authoring precedes dependent Specification and workflow-support authoring.
- Confirm exactly one integrated review exists.
- Confirm no correction or closure-review placeholders exist.
- Confirm no canonical authoring, review, synchronization, implementation, stage, or commit occurs.

## Evidence

- T06 is `done` and supplies four exact ADR boundaries.
- Created T08 as the sole ADR writer for one new ADR and five amendments.
- Created T09 as the sole canonical and workflow-support writer after T08.
- Created T10 as the one integrated independent review after T09.
- Created T11 as the direct-PASS-only closure synchronization owner after T10.
- Shared writers are serialized T08 -> T09 -> T10 -> T11.
- Correction and finding-closure review remain unmaterialized until named findings exist.
- Updated PRODUCT-WORK-SPEC-023 with T08 through T11.
- No ADR, Specification, workflow-support, review, synchronization, implementation, stage, or commit work occurred.
- DRMCP is non-operational. Filesystem authoring was the required fallback.
- Result: `PASS`.
