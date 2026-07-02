# PRODUCT-TASK-SPEC-020-01: Coordinate checklist authoring graph

- **id**: PRODUCT-TASK-SPEC-020-01
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-020
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-22
- **outputs**:
  - PRODUCT-TASK-SPEC-020-01
  - PRODUCT-TASK-SPEC-020-02
  - PRODUCT-TASK-SPEC-020-03
  - PRODUCT-TASK-SPEC-020-04
  - PRODUCT-TASK-SPEC-020-05
  - PRODUCT-TASK-SPEC-020-06
  - PRODUCT-WORK-SPEC-020

## Goal

Materialize the W020 checklist contract, Investigation, authoring, review, and closure graph.

## Work

- Create T02 as the checklist artifact contract decision owner.
- Create T03 as the mandatory checklist impact Investigation owner.
- Reserve PRODUCT-INV-SPEC-008 for T03.
- Create T04 as the checklist artifact authoring owner.
- Create T05 as the integrated independent checklist review owner.
- Create T06 as the lifecycle, Evidence, and relation synchronization owner.
- Gate T05 on accepted W019 integrated review.
- Keep correction and finding-closure review Tasks conditional on named findings.

This Task must not:

- decide checklist format, placement, schema, or wording;
- author the Investigation or checklist artifacts;
- issue a review verdict;
- create speculative correction Tasks;
- implement the validator;
- modify W019 canonical artifacts;
- perform stage or commit work.

## Done condition

- T02 through T06 exist with one responsibility each.
- Dependencies follow decision, Investigation, authoring, review, and synchronization order.
- PRODUCT-INV-SPEC-008 is reserved only for T03.
- T05 records the accepted W019 review gate.
- Conditional finding work remains unmaterialized.

## Verification

- Confirm every Task uses one canonical `task_type`.
- Confirm every Task references PRODUCT-WORK-SPEC-020.
- Confirm Work Item `tasks` and Task `work_item` relations match.
- Confirm no checklist artifact or Investigation was authored.
- Confirm W021 remains outside this graph.

## Evidence

- W020 already defines one checklist-authoring completion boundary.
- PRODUCT-TASK-SPEC-019-21 permits W020 decision, Investigation, and authoring before W019 closure.
- PRODUCT-TASK-SPEC-019-21 gates W020 integrated review on accepted W019 review.
- Result: `PASS`.
