# PRODUCT-TASK-SPEC-031-05: Synchronize USDM hierarchy design closure

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-031
- **task_type**: synchronization
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-031-09
- **outputs**:
  - PRODUCT-WORK-SPEC-031

## Goal

Synchronize an accepted integrated-review result into the design Work Item lifecycle and closure Evidence.

## Work

- Read the final accepted re-review result.
- Confirm every Completion Condition is mechanically satisfied.
- Update only this Task and PRODUCT-WORK-SPEC-031 lifecycle and closure Evidence.
- Stop if review did not pass or any design judgment remains.

## Done condition

The accepted review result and Work Item lifecycle express the same closed design state.

## Verification

No ADR, Specification, completed decision Task, authoring Task, or review verdict is changed during synchronization.

## Evidence

- Accepted review route: PRODUCT-TASK-SPEC-031-09 `PASS`.
- Prior review PRODUCT-TASK-SPEC-031-04 returned `NEEDS REVISION`; REV-031-001 and REV-031-002 are independently `CLOSED` by PRODUCT-TASK-SPEC-031-09.
- Accepted ADR: PRODUCT-ADR-SPEC-020, status `accepted`, reflected in Specification state dated 2026-09-30.
- Accepted Specification refs:
  - `spec:product.design_records.usdm`
  - `spec:product.design_records.usdm.artifact_format`
  - `spec:product.design_records.usdm.coverage_format`
  - `spec:product.design_records.usdm.coverage_tools`
  - `spec:product.design_records.authoring_standards.usdm_authoring`
- PRODUCT-REQ-SPEC-016 reflects record-local hierarchy and cross-app full-ID `usdm_covers`.
- Tool implementation and migration remain downstream scope and are not claimed by design closure.
- Every PRODUCT-WORK-SPEC-031 Completion Condition is satisfied.
- PRODUCT-WORK-SPEC-031 is synchronized to `done`.
- No canonical design content, Task graph decision, or review verdict was changed by this synchronization.
