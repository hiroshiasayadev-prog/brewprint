# PRODUCT-TASK-SPEC-019-04: Coordinate mandatory impact Investigation route

- **id**: PRODUCT-TASK-SPEC-019-04
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-03
- **outputs**:
  - PRODUCT-WORK-SPEC-019
  - PRODUCT-TASK-SPEC-019-04
  - PRODUCT-TASK-SPEC-019-05

## Goal

Materialize the mandatory impact Investigation route after the corrected temporary standalone validator decision.

## Work

- Add T04 and T05 to the W019 Task graph.
- Materialize T05 as the sole owner of one bounded formal Investigation.
- Persist T05 responsibility, dependency, outputs, excluded scope, and release condition.
- Reserve `PRODUCT-INV-SPEC-007` as the bounded Investigation output.
- Sequence T05 after T04.
- Synchronize the W019 Task flow and Task Candidates with the materialized route.
- Preserve T01 through T03 as completed historical Evidence.
- Keep every post-Investigation branch conditional until T05 concludes.

`PRODUCT-INV-SPEC-006` is unavailable because W018 already owns that Investigation ID.
The next available sequence, `PRODUCT-INV-SPEC-007`, is used without changing the Investigation responsibility.

This Task does not perform Investigation, reconciliation judgment, ADR routing, authoring, review, implementation, synchronization, stage, or commit.

## Done condition

- W019 lists T04 and T05.
- T04 owns only mandatory Investigation-route coordination.
- T05 exists with status `not_started` and task type `investigation`.
- T05 depends on T04.
- T05 outputs only T05 and `PRODUCT-INV-SPEC-007`.
- T05 reserves one formal Investigation record for one bounded research question.
- T01 through T03 remain substantively unchanged.
- No conditional successor Task is materialized before T05 concludes.

## Verification

- Result: W019 lists T04 and T05 in its `tasks` field.
- Result: T04 and T05 reference `PRODUCT-WORK-SPEC-019` as their parent.
- Result: T04 has task type `coordination` and one graph-materialization outcome.
- Result: T05 has task type `investigation` and one formal Investigation outcome.
- Result: T05 depends only on T04.
- Result: T05 outputs only T05 and `PRODUCT-INV-SPEC-007`.
- Result: T01 through T03 were not modified.
- Result: no correction, review, ADR, Specification, implementation, or closure Task was materialized.
- Result: no DRMCP artifact was changed.

## Evidence

- Materialized Task path: `product/records/tasks/spec/PRODUCT-TASK-SPEC-019-04-coordinate-mandatory-impact-investigation-route.md`.
- Materialized Task path: `product/records/tasks/spec/PRODUCT-TASK-SPEC-019-05-investigate-temporary-standalone-validator-impact-and-conflicts.md`.
- Reserved Investigation ID: `PRODUCT-INV-SPEC-007`.
- Planned Investigation path: `product/records/investigations/spec/PRODUCT-INV-SPEC-007-temporary-standalone-task-validator-impact-and-conflicts.md`.
- ID collision Evidence: `PRODUCT-INV-SPEC-006` exists at `product/records/investigations/spec/PRODUCT-INV-SPEC-006-design-convergence-impact-and-conflict-inventory.md` and is owned by W018 T10.
- W019 relation synchronization adds T04 and T05 to the Work Item `tasks` field and flow.
- T01 through T03 remained unchanged.
- The Investigation was not executed, and `PRODUCT-INV-SPEC-007` was not created.
- No downstream conditional Task was created.
- No DRMCP artifact was read or modified for graph materialization.
- No stage or commit was performed.
- DRMCP is non-operational under the current agent authoring policy. Filesystem authoring was used.
