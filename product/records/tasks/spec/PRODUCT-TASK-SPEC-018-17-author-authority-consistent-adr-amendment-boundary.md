# PRODUCT-TASK-SPEC-018-17: Author authority-consistent ADR amendment boundary

- **id**: PRODUCT-TASK-SPEC-018-17
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-16
- **outputs**:
  - PRODUCT-TASK-SPEC-018-17
  - PRODUCT-ADR-SPEC-006
  - spec:product.design_records.authoring_standards.adr_authoring
  - skills/design-convergence-workflow/adr-routing.md

## Goal

Project the accepted F-BLK-01 materiality decision into the durable ADR, canonical ADR authoring rules, and active routing authority.

## Work

- Amend ADR-006 with the accepted non-material responsibility-refinement boundary.
- Update `adr_authoring` with the same amendment and supersession rule.
- Update `adr-routing.md` with the same materiality test.
- Preserve the current amended content of ADR-004, ADR-005, and ADR-010.
- Preserve T12 and T14 as historical Evidence.

This Task must not:

- change the selected `work_item_decomposition` design;
- rewrite T12 or T14;
- create a new ADR;
- repair F-MAJ-01;
- perform finding-closure review;
- stage or commit changes.

## Done condition

- ADR-006 defines responsibility extraction as non-material when the core architecture and rationale remain valid.
- `adr_authoring` contains the matching normative rule.
- `adr-routing.md` applies the same materiality test.
- Supersession remains required for a material core ownership or architecture change.
- ADR-004, ADR-005, and ADR-010 remain valid in-place amendments.

## Verification

- Compare the amendment criteria across all three outputs.
- Confirm the selected alternative and core typed-responsibility architecture remain unchanged.
- Confirm no new ADR or supersession chain was created.
- Confirm T12, T14, and T07 were not modified.
- Confirm stage and commit were not performed.

## Evidence

- T16 classified the responsibility split as `non_material_responsibility_refinement`.
- ADR-006 now records the durable materiality distinction.
- `adr_authoring` now permits bounded responsibility extraction inside an unchanged architecture.
- `adr-routing.md` now requires an explicit materiality judgment for responsibility-wording amendments.
- No new ADR was created.
- F-BLK-01 is repaired pending independent finding-closure review.
