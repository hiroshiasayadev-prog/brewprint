# PRODUCT-TASK-SPEC-007-04: Review split content against index_old

- **id**: PRODUCT-TASK-SPEC-007-04
- **status**: done
- **date**: 2026-06-13
- **work_item**: PRODUCT-WORK-SPEC-007
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **source_work_item**: PRODUCT-WORK-SPEC-007
- **estimate**: 0.2d
- **depends_on**:
  - PRODUCT-TASK-SPEC-007-03
- **outputs**:

## Goal

Review the split spec-format content against `product/records/spec/concepts/spec-format/index_old.md` before cleanup.

## Work

Compare the navigation-first `index.md` and split child specs against `index_old.md` for split mechanics, content representation, and responsibility boundaries.

## Done condition

| item | done when |
|---|---|
| split mechanics reviewed | `index.md` and child specs are compared against `index_old.md`. |
| review result recorded | The review verdict and major findings are captured for follow-up correction work. |
| cleanup deferred | `index_old.md` remains preserved until corrections are applied and reviewed. |

## Verification

- Confirm `index.md` is navigation-first.
- Confirm major content from `index_old.md` is represented in child specs.
- Confirm no immediate split-mechanics blocker was found.
- Confirm major design findings are carried into PRODUCT-TASK-SPEC-007-05.

## Evidence

- `product/records/spec/concepts/spec-format/index.md` is navigation-first: it contains a concise `## What this is`, a `## Topics` table, and related specs instead of the detailed contract body.
- Major content from `index_old.md` is represented in child specs for overview/front matter policy, document shape, Topics table, spec ID-as-ref, validation policy, and follow-up boundary.
- No immediate split-mechanics blocker was found: the split structure is usable for review and follow-up corrections.
- Review found major design corrections needed before cleanup:
  - `## Source records` should be removed from specs because it bloats specs and duplicates traceability.
  - ADR-governed spec change traceability is needed; follow-up must investigate which ADRs should receive `target_specs`.
  - `## Topics` should use `ref`, not `file`.
  - Spec IDs are path-derived canonical refs; move/rename changes the ID.
  - `Contract` needs `contract_class: interface | format`; `Request` and `Response` should not be forced onto format contracts.
- `index_old.md` remains preserved for comparison and later cleanup.

