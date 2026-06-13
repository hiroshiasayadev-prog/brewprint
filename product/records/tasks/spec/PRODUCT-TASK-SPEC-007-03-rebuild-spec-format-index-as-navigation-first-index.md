# PRODUCT-TASK-SPEC-007-03: Rebuild spec-format index as navigation-first index

- **id**: PRODUCT-TASK-SPEC-007-03
- **status**: done
- **date**: 2026-06-11
- **work_item**: PRODUCT-WORK-SPEC-007
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **source_work_item**: PRODUCT-WORK-SPEC-007
- **estimate**: 0.2d
- **depends_on**:
  - PRODUCT-TASK-SPEC-007-02
- **outputs**:
  - product/records/spec/concepts/spec-format/index.md

## Goal

Recreate `product/records/spec/concepts/spec-format/index.md` as a navigation-first Index for the split spec-format topic specs.

## Work

Create a small `# Index: Spec format` document with H1-adjacent metadata, a concise `## What this is`, an authoritative `## Topics` table for the six child specs, and minimal related specs.

## Done condition

| item | done when |
|---|---|
| index exists | `product/records/spec/concepts/spec-format/index.md` exists. |
| navigation-first | The index points to child specs and does not contain detailed contract body. |
| topics authoritative | The `## Topics` table declares the six child specs under `spec:product.concepts.spec_format`. |
| temporary source preserved | `index_old.md` remains present for later review and cleanup. |

## Verification

- Confirm `index.md` has H1 `# Index: Spec format`.
- Confirm `index.md` has id `spec:product.concepts.spec_format`.
- Confirm `index.md` contains `## Topics` rows for the six child specs.
- Confirm detailed contract sections live in child specs, not in `index.md`.

## Evidence

- Recreated `product/records/spec/concepts/spec-format/index.md` as a navigation-first Index.
- The `## Topics` table points to the six split child specs.
- The index omits detailed document-shape, Topics table, ID-as-ref, validation, and follow-up contract bodies.
