# PRODUCT-TASK-SPEC-028-07: Synchronize retrospective evidence closure

- **id**: PRODUCT-TASK-SPEC-028-07
- **status**: not_started
- **date**: 2026-07-08
- **work_item**: PRODUCT-WORK-SPEC-028
- **task_type**: synchronization
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-028-06
- **outputs**:
  - PRODUCT-WORK-SPEC-028

## Goal

Synchronize lifecycle, Evidence, relations, and closure state after PRODUCT-WORK-SPEC-028 review passes.

## Work

- Read the T06 review verdict and findings.
- If review is PASS, synchronize Work Item Evidence and closure state.
- If review needs revision, stop and route through finding-specific coordination.
- Do not repair findings, alter the review verdict, or change the Task graph.

## Done condition

- PASS review state is reflected in the Work Item lifecycle and Evidence, or synchronization is blocked by named findings.
- Synchronization changes only mechanically derived lifecycle, Evidence, relation, and closure state.

## Verification

TBD

## Evidence

- PRODUCT-TASK-SPEC-028-01 D-005 requires post-review synchronization after integrated review.
