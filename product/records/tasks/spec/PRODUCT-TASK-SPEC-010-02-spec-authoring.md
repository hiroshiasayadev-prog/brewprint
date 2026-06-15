# PRODUCT-TASK-SPEC-010-02: Author writing standard spec

- **id**: PRODUCT-TASK-SPEC-010-02
- **status**: done
- **date**: 2026-06-15
- **work_item**: PRODUCT-WORK-SPEC-010
- **source_requirement**: PRODUCT-REQ-SPEC-002
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-010-01
- **outputs**:
  - `product/records/spec/concepts/authoring-standards/index.md`
  - `product/records/spec/concepts/authoring-standards/writing-standard.md`

## Goal

Create the Reference-kind writing standard spec and its parent Index under `spec:product.concepts.authoring_standards`.

## Work

- Confirm spec kind: Guide kind is deferred per `spec:product.concepts.spec_format.document_shape`; use Reference kind instead.
- Create parent Index at `product/records/spec/concepts/authoring-standards/index.md` (`spec:product.concepts.authoring_standards`).
- Create writing standard spec at `product/records/spec/concepts/authoring-standards/writing-standard.md` (`spec:product.concepts.authoring_standards.writing_standard`) with all confirmed rules and correct normative levels.

## Done condition

| item | done when |
|---|---|
| Parent Index exists | `authoring-standards/index.md` present with Topics row pointing to writing-standard. |
| Writing standard spec exists | `authoring-standards/writing-standard.md` present with correct id, parent, status, and all confirmed rules. |
| Spec kind correct | H1 uses `Reference`, not `Guide`. |

## Verification

- Confirm both files exist at correct paths.
- Confirm id derivation: `spec:product.concepts.authoring_standards` and `spec:product.concepts.authoring_standards.writing_standard`.
- Confirm parent declared correctly in index Topics table.

## Evidence

- `authoring-standards/index.md` created on 2026-06-15.
- `authoring-standards/writing-standard.md` created on 2026-06-15 with status `draft`.
- Guide kind replacement noted in PRODUCT-WORK-SPEC-010 Goal and Completion Condition.
