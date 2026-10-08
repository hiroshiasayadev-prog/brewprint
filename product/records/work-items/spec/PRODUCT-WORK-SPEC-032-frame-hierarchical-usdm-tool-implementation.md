# PRODUCT-WORK-SPEC-032: Frame hierarchical USDM tool implementation

- **status**: done
- **date**: 2026-09-30
- **source_refs**:
  - PRODUCT-REQ-SPEC-016
- **impact_refs**:
  - `spec:product.design_records.usdm.artifact_format`
  - `spec:product.design_records.usdm.coverage_format`
  - `spec:product.design_records.usdm.coverage_tools`
- **tasks**:
  - PRODUCT-TASK-SPEC-032-01
  - PRODUCT-TASK-SPEC-032-02

## Goal

Frame the downstream implementation boundary for the reviewed hierarchical USDM design.

## Boundary

This Work Item owns implementation framing and downstream Work Item creation.

It does not own production code changes, tool tests, or implementation review.

## Impact Scope

| target | impact |
|---|---|
| PRODUCT-REQ-SPEC-016 | Direct implementation source. |
| PRODUCT-WORK-SPEC-031 | Reviewed design closure establishing implementation-ready semantics. |
| PRODUCT-WORK-SPEC-033 | Downstream implementation Work Item. |

## Task flow

```text
PRODUCT-TASK-SPEC-032-01 decide implementation framing
  -> PRODUCT-TASK-SPEC-032-02 create implementation Work Item
```

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| PRODUCT-TASK-SPEC-032-01 | decision | Fix implementation Goal, Boundary, Completion Condition, unknown handling, and route. | none |
| PRODUCT-TASK-SPEC-032-02 | work_item_decomposition | Create the accepted implementation Work Item and bounded Task graph. | T01 |

## Completion Condition

- PRODUCT-REQ-SPEC-016 remains the direct motivating Requirement.
- PRODUCT-WORK-SPEC-031 is confirmed complete with a final `PASS` review.
- Tool implementation has one bounded Work Item separate from design closure.
- Implementation targets, verification ownership, and review boundary are explicit.

## Evidence

- PRODUCT-TASK-SPEC-032-01 records the `proceed` implementation framing decision.
- PRODUCT-TASK-SPEC-032-02 creates PRODUCT-WORK-SPEC-033.
- No new design judgment is delegated to implementation.
