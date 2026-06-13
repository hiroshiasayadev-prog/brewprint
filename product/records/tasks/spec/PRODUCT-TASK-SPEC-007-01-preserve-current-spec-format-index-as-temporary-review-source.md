# PRODUCT-TASK-SPEC-007-01: Preserve current spec-format index as temporary review source

- **id**: PRODUCT-TASK-SPEC-007-01
- **status**: done
- **date**: 2026-06-11
- **work_item**: PRODUCT-WORK-SPEC-007
- **source_work_item**: PRODUCT-WORK-SPEC-001
- **estimate**: 0.1d
- **depends_on**:
- **outputs**:
  - product/records/spec/concepts/spec-format/index_old.md

## Goal

Preserve the current all-in-one spec-format index as a temporary review source before splitting it into navigation-first index and child topic specs.

## Work

Rename `product/records/spec/concepts/spec-format/index.md` to `product/records/spec/concepts/spec-format/index_old.md` before rebuilding `index.md` as a navigation-first Index.

## Done condition

| item | done when |
|---|---|
| temporary source preserved | `index_old.md` exists and contains the previous all-in-one spec-format contract. |
| split source clear | Later split/review tasks can compare new child specs against `index_old.md`. |
| no permanent duplicate | Any temporary duplicate spec ID risk is explicitly carried forward to the cleanup task. |

## Verification

- Confirm `product/records/spec/concepts/spec-format/index_old.md` exists.
- Confirm `product/records/spec/concepts/spec-format/index.md` is absent before task 007-03 recreates it.

## Evidence

- User reported that `product/records/spec/concepts/spec-format/index.md` was renamed to `product/records/spec/concepts/spec-format/index_old.md`.
- This preserves the previous all-in-one spec-format contract as the comparison source for PRODUCT-TASK-SPEC-007-04.
- Temporary duplicate/missing index risk is accepted only during PRODUCT-WORK-SPEC-007 and must be resolved before close by deleting `index_old.md` after review.
