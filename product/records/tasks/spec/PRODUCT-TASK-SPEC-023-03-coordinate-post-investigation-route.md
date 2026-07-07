# PRODUCT-TASK-SPEC-023-03: Coordinate post-Investigation route

- **id**: PRODUCT-TASK-SPEC-023-03
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-023
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-023-02
- **outputs**:
  - PRODUCT-WORK-SPEC-023
  - PRODUCT-TASK-SPEC-023-03
  - PRODUCT-TASK-SPEC-023-04
  - PRODUCT-TASK-SPEC-023-05

## Goal

Materialize one exact post-Investigation Task route for the decided `cancelled` lifecycle contract.

## Work

- Read the completed T01 decision ledger and PRODUCT-INV-SPEC-010.
- Confirm the originating Requirement and Work Item remain coherent.
- Determine whether Investigation findings require a reconciliation decision owner.
- Create only the required downstream Tasks.
- Set exact dependencies, blockers, outputs, writers, and review order.
- Serialize shared writers identified by the Investigation.
- Route ADR classification before canonical authoring.
- Place one integrated independent review after the final writer.
- Keep correction and finding-closure review abstract until named findings exist.
- Update the parent Work Item Task graph and this Task outputs with every materialized Task.

This Task must not:

- decide lifecycle semantics;
- perform Work Item decomposition;
- author ADR, Specification, skill, checklist, validation, or implementation content;
- repair findings;
- issue a review verdict;
- synchronize lifecycle closure;
- implement, stage, or commit changes.

## Done condition

- Every required post-Investigation responsibility has one Task owner.
- Unnecessary conditional Tasks remain unmaterialized.
- Every Task has one type, outcome, completion judgment, writer, and release condition.
- Dependencies, blockers, shared-writer order, and review order are deterministic.
- ADR routing precedes canonical authoring.
- One integrated independent review follows the final writer.
- Closure synchronization remains blocked until an accepted review route exists.
- The Work Item and Task ownership relation is coherent.

## Verification

- Confirm T01 and T02 are complete.
- Confirm every graph change traces to T01, PRODUCT-INV-SPEC-010, or a mechanically required owner repair.
- Confirm no speculative correction or finding-closure review Task exists.
- Confirm no completed Task is substantively rewritten.
- Confirm no circular dependency or concurrent shared writer exists.
- Confirm no deliverable authoring, review, synchronization, implementation, stage, or commit occurs.

## Evidence

- T01 is `done` and PRODUCT-INV-SPEC-010 is `concluded`.
- The Investigation found unresolved body-readiness, propagation-owner, and validator-invocation judgments.
- Created PRODUCT-TASK-SPEC-023-04 as the sole post-Investigation decision owner.
- Created PRODUCT-TASK-SPEC-023-05 as the later graph owner after T04.
- ADR routing, canonical authoring, integrated review, and closure Tasks remain unmaterialized until T04 completes.
- No speculative correction or finding-closure review Task exists.
- Updated PRODUCT-WORK-SPEC-023 to include T04 and T05.
- No lifecycle semantics were decided and no canonical artifact was authored.
- DRMCP is non-operational. Filesystem authoring was the required fallback.
- Result: `PASS`.
