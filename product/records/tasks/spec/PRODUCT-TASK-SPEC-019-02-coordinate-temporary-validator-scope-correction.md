# PRODUCT-TASK-SPEC-019-02: Coordinate temporary-validator scope correction

- **id**: PRODUCT-TASK-SPEC-019-02
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-01
- **outputs**:
  - PRODUCT-WORK-SPEC-019
  - PRODUCT-TASK-SPEC-019-02
  - PRODUCT-TASK-SPEC-019-03

## Goal

Repair the W019 Task graph after the validator delivery boundary was clarified as a temporary standalone tool rather than DRMCP integration.

## Work

- Preserve completed T01 as historical decision Evidence.
- Amend the incomplete W019 boundary and downstream route.
- Add one successor decision Task for the corrected standalone-tool scope.
- Route the mandatory impact Investigation after the successor decision.
- Keep current DRMCP Specifications, tools, diagnostics, and implementation outside W019.

This Task does not decide the corrected product boundary, perform Investigation, author canonical Specifications, integrate DRMCP, or implement the validator.

## Done condition

W019 identifies the temporary standalone tool boundary, preserves T01 unchanged, and contains one successor decision owner before impact Investigation.

## Verification

- Confirm T01 remains `done` and unchanged by this coordination.
- Confirm W019 lists T02 and T03 in `tasks`.
- Confirm T03 owns the corrected decision rather than coordination.
- Confirm the Investigation dependency moves from T01 to T03.
- Confirm no DRMCP artifact or implementation file changed.

## Evidence

- The user clarified that the validator is an interim standalone tool.
- The user deferred DRMCP integration to a later design and delivery boundary.
- `spec:product.design_records.authoring_standards.task_authoring` prohibits substantive reopening of completed decision Tasks.
- `skills/design-convergence-workflow/graph-coordination.md` requires a successor decision route when a completed decision boundary changes.
- W019 now lists T01, T02, and T03.
- W019 now excludes DRMCP integration and changes to current DRMCP artifacts.
- No DRMCP file was modified.
