# PRODUCT-TASK-SPEC-025-02: Author Investigation-Task lightweight Evidence exception

- **id**: PRODUCT-TASK-SPEC-025-02
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-025
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-025-01
- **outputs**:
  - product/records/spec/design-records/authoring-standards/task-authoring.md
  - PRODUCT-TASK-SPEC-025-02

## Goal

Add the Investigation-Task lightweight Evidence exception to `spec:product.design_records.authoring_standards.task_authoring`, using the wording accepted in PRODUCT-TASK-SPEC-025-01 D-006.

## Work

- Read PRODUCT-TASK-SPEC-025-01 D-006.
- Add a new `#### Investigation-Task lightweight Evidence exception` subsection immediately after the Task type contract table and before `#### Single responsibility`.
- Preserve the existing default `investigation` task_type contract (formal Investigation record required) outside the exception.

This Task does not populate `skills/task-boundary-vocabulary/`.

## Done condition

- The exception subsection exists in `task-authoring.md` using the accepted wording.
- The default investigation contract remains unchanged outside the exception.

## Verification

- Confirmed the exception text matches the user-accepted wording.
- Confirmed the exception does not loosen the `investigation` task_type's prohibited overlaps (decision adoption, canonical authoring, implementation, independent review, correction, synchronization).
- Confirmed no other section of `task-authoring.md` was altered.

## Evidence

- Added the "Investigation-Task lightweight Evidence exception" subsection to `product/records/spec/design-records/authoring-standards/task-authoring.md`, placed between the Task type contract table and `#### Single responsibility`.
- Wording matches the final user-accepted revision: the recorded result must receive at least one downstream conclusion, including a conclusion that explicitly determines no further action is required.
- No other section of `task-authoring.md` was changed; the diff was scoped to the single inserted subsection.
- DRMCP is non-operational. Filesystem authoring was the required fallback.
