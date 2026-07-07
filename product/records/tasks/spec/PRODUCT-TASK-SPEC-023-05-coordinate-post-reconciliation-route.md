# PRODUCT-TASK-SPEC-023-05: Coordinate post-reconciliation route

- **id**: PRODUCT-TASK-SPEC-023-05
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-023
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-023-04
- **outputs**:
  - PRODUCT-WORK-SPEC-023
  - PRODUCT-TASK-SPEC-023-05
  - PRODUCT-TASK-SPEC-023-06
  - PRODUCT-TASK-SPEC-023-07

## Goal

Materialize the exact ADR-routing, authoring, review, and closure route after T04 resolves the Investigation findings.

## Work

- Read T01, PRODUCT-INV-SPEC-010, and completed T04.
- Create one bounded ADR-routing Task.
- Create only the authoring Tasks required by the routing result.
- Set exact ADR, Specification, workflow-support, and originating-artifact writer order.
- Place one integrated independent review after the final writer.
- Materialize direct closure synchronization only when its accepted-input boundary can be stated safely.
- Keep correction and finding-closure review abstract until named findings exist.
- Update the Work Item Task flow, candidate table, tasks metadata, and this Task outputs.

This Task must not:

- decide lifecycle semantics;
- classify ADR routes itself;
- author ADR, Specification, skill, checklist, or implementation content;
- perform review, correction, synchronization, implementation, stage, or commit work.

## Done condition

- Every post-T04 responsibility has one owner.
- ADR routing precedes canonical authoring.
- Every writer and dependency order is deterministic.
- One integrated independent review follows the final writer.
- Unused conditional branches remain unmaterialized.
- The Work Item and Task ownership relation is coherent.

## Verification

- Confirm T04 is complete.
- Confirm every created Task traces to T04 and PRODUCT-INV-SPEC-010.
- Confirm no speculative correction or finding-closure review Task exists.
- Confirm no completed Task is substantively rewritten.
- Confirm no shared writer is concurrent.
- Confirm no canonical authoring, review, synchronization, implementation, stage, or commit occurs.

## Evidence

- T04 is `done`; J-001 through J-005 are terminal.
- Created PRODUCT-TASK-SPEC-023-06 as the sole ADR-routing owner.
- Created PRODUCT-TASK-SPEC-023-07 as the sole post-routing graph owner.
- Concrete ADR, Specification, workflow-support, review, and closure Tasks remain unmaterialized until T06 completes.
- No speculative correction or finding-closure review Task exists.
- Updated PRODUCT-WORK-SPEC-023 with T06 and T07.
- No routing outcome was selected and no canonical artifact was authored.
- DRMCP is non-operational. Filesystem authoring was the required fallback.
- Result: `PASS`.
