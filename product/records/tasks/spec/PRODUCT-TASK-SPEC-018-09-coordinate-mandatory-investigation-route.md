# PRODUCT-TASK-SPEC-018-09: Coordinate mandatory Investigation route

- **id**: PRODUCT-TASK-SPEC-018-09
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-06
- **outputs**:
  - PRODUCT-WORK-SPEC-018
  - PRODUCT-TASK-SPEC-018-07
  - PRODUCT-TASK-SPEC-018-09
  - PRODUCT-TASK-SPEC-018-10
  - PRODUCT-TASK-SPEC-018-11
  - PRODUCT-TASK-SPEC-018-12

## Goal

Repair the W018 Task graph after the T07 prerequisite failure by assigning every mandatory Investigation and reconvergence responsibility to an exact successor Task.

Preserve T01 through T06 as completed historical Evidence.

## Work

- Treat T07 prerequisite `P-BLK-01` as a mechanically necessary missing-owner repair.
- Do not treat `P-BLK-01` as an integrated-review finding.
- Materialize `PRODUCT-TASK-SPEC-018-10` as the formal `investigation` owner.
- Require T10 to produce one formal Investigation record for one bounded research question.
- Materialize `PRODUCT-TASK-SPEC-018-11` as the post-Investigation reconciliation `decision` owner.
- Materialize `PRODUCT-TASK-SPEC-018-12` as the ADR-routing revalidation `decision` owner.
- Sequence T10 after this Task, T11 after T10, and T12 after T11.
- Amend the incomplete T07 release route so integrated review resumes only after T12 and every authoring Task required by T11 or T12.
- Preserve `PRODUCT-TASK-SPEC-018-08` as the reserved closure-synchronization owner.
- Update the W018 Task graph, candidates, dependencies, and conditional routes without changing its resolution identity.
- Keep conditional authoring Tasks unmaterialized until T11 or T12 identifies exact required targets.

This Task does not perform Investigation, reconciliation judgment, ADR routing, authoring, independent review, synchronization, production implementation, stage, or commit.

## Done condition

- T10, T11, and T12 exist with one Task type, one primary outcome, exact dependencies, and bounded outputs.
- T10 owns reserved Investigation output `PRODUCT-INV-SPEC-006` without owning decision or authoring outcomes.
- T11 owns post-Investigation reconciliation judgment.
- T12 owns revalidation of the historical T02 ADR route.
- W018 and T07 express the same reconvergence and review-release order.
- T01 through T06 remain substantively unchanged.
- T08 remains unmaterialized and reserved for closure synchronization.
- No speculative authoring, correction, or finding-closure review Task exists.

## Verification

- Result: W018 lists T07, T09, T10, T11, and T12.
- Result: T09 through T12 reference `PRODUCT-WORK-SPEC-018` as their parent.
- Result: the dependency chain is T09 to T10 to T11 to T12, with no cycle.
- Result: T07 depends on T05, T06, and T12.
- Result: T07 remains blocked until T12 and every conditionally required writer complete.
- Result: T01 through T06 were not modified.
- Result: T08 and conditional authoring Tasks were not materialized.
- Result: only the declared six-file writable boundary changed.
- Result: stage and commit were not performed.

## Evidence

- Trigger: T07 recorded verdict `NOT READY` with prerequisite `P-BLK-01`.
- Materialized Task path: `product/records/tasks/spec/PRODUCT-TASK-SPEC-018-10-investigate-design-convergence-impact-and-conflicts.md`.
- Materialized Task path: `product/records/tasks/spec/PRODUCT-TASK-SPEC-018-11-decide-post-investigation-reconciliation.md`.
- Materialized Task path: `product/records/tasks/spec/PRODUCT-TASK-SPEC-018-12-revalidate-adr-routing-after-investigation.md`.
- Reserved Investigation ID: `PRODUCT-INV-SPEC-006`.
- Planned Investigation path: `product/records/investigations/spec/PRODUCT-INV-SPEC-006-design-convergence-impact-and-conflict-inventory.md`.
- Exact dependency chain: T10 depends on T09; T11 depends on T10; T12 depends on T11.
- W018 relation synchronization: T10, T11, and T12 were added to the Work Item `tasks` field and flow.
- T07 amendment: T12 was added as a dependency, and T09 through T12 plus `PRODUCT-INV-SPEC-006` were added to the final review boundary.
- Completed-record preservation: T01 through T06 remained unchanged.
- Reserved route preservation: T08 was not materialized.
- Investigation was not executed, and `PRODUCT-INV-SPEC-006` was not created.
- Reconciliation decision work was not executed.
- ADR-routing revalidation was not executed.
- ADR, Specification, Requirement, Work Item, and skill authoring was not executed.
- Integrated review was not resumed.
- Stage and commit were not performed.
- DRMCP is non-operational. Filesystem authoring was used under the current policy.
