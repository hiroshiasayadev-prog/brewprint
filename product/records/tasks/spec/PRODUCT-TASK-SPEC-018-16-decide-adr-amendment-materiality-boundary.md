# PRODUCT-TASK-SPEC-018-16: Decide ADR amendment materiality boundary

- **id**: PRODUCT-TASK-SPEC-018-16
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-15
- **outputs**:
  - PRODUCT-TASK-SPEC-018-16

## Goal

Decide the F-BLK-01 amend-versus-supersede boundary for the `work_item_decomposition` responsibility split.

## Work

- Decide whether the responsibility split is a material ownership or architecture change.
- Decide the final ADR disposition for ADR-004, ADR-005, and ADR-010.
- Fix the durable amendment boundary and exact authoring targets.
- Preserve T07 and completed predecessor Tasks as historical Evidence.

This Task must not:

- edit T12 or T14;
- author ADR, Specification, skill, or instruction content;
- repair F-MAJ-01;
- perform finding-closure review;
- stage or commit changes.

## Done condition

- F-BLK-01 has one explicit materiality decision.
- ADR-004, ADR-005, and ADR-010 have one final disposition.
- The amendment and supersession boundaries are explicit.
- Exact authoring targets are fixed.

## Verification

- Confirm the selected disposition follows one coherent materiality rule.
- Confirm the decision does not conceal a selected-alternative reversal.
- Confirm every authoring target follows from the accepted decision.
- Confirm no canonical artifact was authored by this Task.

## Evidence

### Decision result

- Finding: F-BLK-01.
- Status: `decided`.
- User answer: `accepted`.
- Materiality: `non_material_responsibility_refinement`.
- ADR-004 disposition: `amend`.
- ADR-005 disposition: `amend`.
- ADR-010 disposition: `amend`.
- New ADR: none.

### Accepted boundary

The change extracts one overloaded responsibility from `coordination` into a named Task type.
The selected typed-responsibility architecture remains unchanged.
The closed taxonomy model remains unchanged.
The single-responsibility rule remains unchanged.
The typed convergence workflow remains unchanged.

An accepted ADR may be amended when:

- the selected alternative remains unchanged;
- the core architecture and rationale remain valid;
- a responsibility is clarified or extracted inside that architecture;
- the amendment does not conceal a reversal.

Supersession remains required when:

- the selected alternative changes;
- the core ownership architecture changes materially;
- an accepted constraint is removed or reversed;
- the prior rationale no longer justifies the current state.

### Authoring route

T17 must update:

- `PRODUCT-ADR-SPEC-006` with the accepted materiality distinction;
- `spec:product.design_records.authoring_standards.adr_authoring` with the normative amendment boundary;
- `skills/design-convergence-workflow/adr-routing.md` with the matching routing rule.

ADR-004, ADR-005, and ADR-010 retain their current amended content.
T12 and T14 remain historical Evidence and are not rewritten.
