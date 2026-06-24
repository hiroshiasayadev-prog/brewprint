# PRODUCT-WORK-SPEC-010: Cognitive load writing standard for design records

- **id**: PRODUCT-WORK-SPEC-010
- **status**: completed
- **date**: 2026-06-15
- **source_requirement**: PRODUCT-REQ-SPEC-002
- **impact_refs**:
- **tasks**:
  - PRODUCT-TASK-SPEC-010-01
  - PRODUCT-TASK-SPEC-010-02
  - PRODUCT-TASK-SPEC-010-03
  - PRODUCT-TASK-SPEC-010-04

## Goal

Formalize PRODUCT-INV-SPEC-003 findings as a writing standard for brewprint design records. Produces a Reference-kind SPEC covering prose style rules and AI output constraints. Guide kind is deferred per `spec:product.design_records.spec_format.document_shape`; Reference kind used instead. Must complete before PRODUCT-WORK-SPEC-009 so that format-migrated specs comply with the writing standard from the start.

## Boundary

Writing standard only. No bulk rewrite of existing records.

## Impact Scope

- PRODUCT-INV-SPEC-003 (source investigation, concluded)
- PRODUCT-WORK-SPEC-009 (blocked on this; format migration must comply with the writing standard)

## Task flow

1. Rule selection from PRODUCT-INV-SPEC-003 candidates — user confirmed rule set and normative levels (2026-06-15).
2. Spec authoring — Reference-kind SPEC created at `spec:product.design_records.authoring_standards.writing_standard` (2026-06-15).
3. Spec review (Gate 4) — two findings applied: subordinating conjunction rule downgraded to SHOULD; verdict rule narrowed to embedded review verdict fields (2026-06-15).
4. Spec accepted, work item closed (2026-06-15).

## Completion Condition

- Writing standard published as a Reference-kind SPEC in the product namespace. (Guide kind deferred; Reference kind used. See `spec:product.design_records.spec_format.document_shape`.)
- AI output constraints for design record prose documented.

## Evidence

- `spec:product.design_records.authoring_standards.writing_standard` — writing standard spec, accepted 2026-06-15.
- `spec:product.design_records.authoring_standards` — parent Index created to support future authoring standard child specs.
