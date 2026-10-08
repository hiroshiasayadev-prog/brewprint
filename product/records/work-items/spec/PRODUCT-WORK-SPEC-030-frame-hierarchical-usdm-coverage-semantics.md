# PRODUCT-WORK-SPEC-030: Frame hierarchical USDM coverage semantics

- **status**: done
- **date**: 2026-09-30
- **source_refs**:
  - PRODUCT-REQ-SPEC-016
- **impact_refs**:
  - `spec:product.design_records.usdm.artifact_format`
  - `spec:product.design_records.usdm.coverage_format`
  - `spec:product.design_records.usdm.coverage_tools`
  - `spec:product.design_records.authoring_standards.usdm_authoring`
- **tasks**:
  - PRODUCT-TASK-SPEC-030-01
  - PRODUCT-TASK-SPEC-030-02

## Goal

Decide how PRODUCT-REQ-SPEC-016 proceeds and establish one downstream design-convergence boundary.

## Boundary

This Work Item owns Requirement framing and downstream Work Item creation.

It does not own the ADR, normative USDM Specification changes, tool implementation, migration, or independent design review.

## Impact Scope

| target | impact |
|---|---|
| PRODUCT-REQ-SPEC-016 | Direct framing source. |
| USDM Specification area | Downstream design target. |
| PRODUCT-WORK-SPEC-031 | Downstream design-convergence Work Item created from the accepted framing route. |

## Task flow

```text
PRODUCT-TASK-SPEC-030-01 decide framing
  -> PRODUCT-TASK-SPEC-030-02 create downstream design Work Item
```

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| PRODUCT-TASK-SPEC-030-01 | decision | Fix source disposition, outcome alignment, downstream boundary, and route. | none |
| PRODUCT-TASK-SPEC-030-02 | work_item_decomposition | Create the accepted downstream design-convergence Work Item. | T01 |

## Completion Condition

- PRODUCT-REQ-SPEC-016 has one explicit source disposition.
- Desired Outcome and Required Outcome are aligned.
- The downstream Goal, Boundary, Completion Condition, direct source, unknown handling, and initial route are fixed.
- The downstream Work Item exists with a distinct completion boundary.

## Evidence

- PRODUCT-TASK-SPEC-030-01 records the accepted `proceed` framing decision.
- No formal Investigation is required because the current USDM contracts resolve the existing format and coverage behavior.
- PRODUCT-TASK-SPEC-030-02 created PRODUCT-WORK-SPEC-031 for design convergence.
- Framing is complete without waiting for downstream design execution.
