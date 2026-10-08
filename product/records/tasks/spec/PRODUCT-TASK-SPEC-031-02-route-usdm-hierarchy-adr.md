# PRODUCT-TASK-SPEC-031-02: Route USDM hierarchy ADR

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-031
- **task_type**: decision
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-031-01
- **outputs**:

## Goal

Route the terminal hierarchy and coverage decisions to the correct ADR boundary.

## Work

- Check accepted ADRs for existing authority over hierarchical USDM coverage semantics.
- Classify the decision set against ADR-required criteria.
- Select one coherent ADR boundary and exact authoring targets.

## Done condition

Every decision from PRODUCT-TASK-SPEC-031-01 has one ADR route and the authoring boundary is fixed.

## Verification

No accepted ADR already owns the same durable choice, and one ADR can coherently own the shared hierarchy and coverage rationale.

## Evidence

| decisions | route | disposition | target | reason |
|---|---|---|---|---|
| D001-D008 | required | create | PRODUCT-ADR-SPEC-020 | The decisions change persistent ID grammar, validation, coverage success, warning behavior, and several Specifications. |

One ADR boundary is used because the hierarchy exists to support the selected effective-coverage semantics.
Separating the decisions would allow incompatible partial supersession.

Authoring targets:

- PRODUCT-ADR-SPEC-020;
- `spec:product.design_records.usdm.artifact_format`;
- `spec:product.design_records.usdm.coverage_format`;
- `spec:product.design_records.usdm.coverage_tools`;
- `spec:product.design_records.usdm`;
- `spec:product.design_records.authoring_standards.usdm_authoring`.

No accepted ADR was found that already owns this decision.
