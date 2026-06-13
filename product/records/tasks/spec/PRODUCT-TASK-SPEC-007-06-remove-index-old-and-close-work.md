# PRODUCT-TASK-SPEC-007-06: Remove index_old and close work

- **id**: PRODUCT-TASK-SPEC-007-06
- **status**: done
- **date**: 2026-06-13
- **work_item**: PRODUCT-WORK-SPEC-007
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **source_work_item**: PRODUCT-WORK-SPEC-007
- **estimate**: 0.2d
- **depends_on**:
  - PRODUCT-TASK-SPEC-007-05
- **outputs**:
  - product/records/spec/concepts/spec-format/index.md
  - product/records/work-items/spec/PRODUCT-WORK-SPEC-007-split-spec-format-index-into-navigation-first-topic-specs.md

## Goal

Delete `product/records/spec/concepts/spec-format/index_old.md` and close PRODUCT-WORK-SPEC-007 after the review corrections are applied and reviewed.

## Work

Remove the temporary all-in-one review source only after PRODUCT-TASK-SPEC-007-05 has corrected the split spec-format contract. Then update PRODUCT-WORK-SPEC-007 for closure.

This task replaces the earlier intended cleanup task.

## Done condition

| item | done when |
|---|---|
| corrections complete | PRODUCT-TASK-SPEC-007-05 is done and reviewed. |
| temporary source removed | `index_old.md` is deleted after it is no longer needed for comparison. |
| duplicate identity risk gone | Only the navigation-first `index.md` carries `spec:product.concepts.spec_format`. |
| work item closed | PRODUCT-WORK-SPEC-007 is marked done with final evidence. |

## Verification

- Confirm PRODUCT-TASK-SPEC-007-05 is done before deleting `index_old.md`.
- Confirm `product/records/spec/concepts/spec-format/index_old.md` is absent after cleanup.
- Confirm `product/records/spec/concepts/spec-format/index.md` remains the canonical navigation-first Index.
- Confirm PRODUCT-WORK-SPEC-007 is marked done only after cleanup.

## Evidence

- User confirmed `product/records/spec/concepts/spec-format/index_old.md` was deleted after the split review and contract corrections completed.
- User confirmed PRODUCT-WORK-SPEC-002 file was renamed to `product/records/work-items/spec/PRODUCT-WORK-SPEC-002-path-derived-canonical-spec-refs-and-ref-first-topic-index.md` so its file path matches the reframed H1/title.
- Cleanup completed after PRODUCT-TASK-SPEC-007-05 corrected the spec-format contract.
- No PRODUCT-WORK-SPEC-001 reopening, `v01/` edits, DRMCP implementation changes, or unrelated spec migration were performed in this task.

- Reserved as the cleanup and closure task after review corrections.

