# PRODUCT-TASK-SPEC-007-02: Create split spec-format topic files

- **id**: PRODUCT-TASK-SPEC-007-02
- **status**: done
- **date**: 2026-06-11
- **work_item**: PRODUCT-WORK-SPEC-007
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **source_work_item**: PRODUCT-WORK-SPEC-001
- **estimate**: 0.4d
- **depends_on**:
  - PRODUCT-TASK-SPEC-007-01
- **outputs**:
  - product/records/spec/concepts/spec-format/overview.md
  - product/records/spec/concepts/spec-format/document-shape.md
  - product/records/spec/concepts/spec-format/topics-table.md
  - product/records/spec/concepts/spec-format/spec-id-as-ref.md
  - product/records/spec/concepts/spec-format/validation-policy.md
  - product/records/spec/concepts/spec-format/follow-up-boundary.md

## Goal

Split the substantive spec-format contract body from `index_old.md` into focused child topic specs.

## Work

Create child specs for overview, document shape, Topics table, spec ID-as-ref, validation policy, and follow-up boundary.

## Done condition

| item | done when |
|---|---|
| child specs created | The six focused child spec files exist under `product/records/spec/concepts/spec-format/`. |
| content moved by meaning | The major sections from `index_old.md` are represented in focused child specs instead of one all-in-one document. |
| scope held | No existing spec migration, DRMCP implementation patch, `v01/` change, or PRODUCT-WORK-SPEC-001 reopening is performed. |

## Verification

- Confirm each output file exists.
- Confirm each child spec has H1-adjacent `id`, `status`, `date`, and `parent` markers.
- Confirm `index_old.md` still exists for PRODUCT-TASK-SPEC-007-04 review and PRODUCT-TASK-SPEC-007-05 cleanup.

## Evidence

- Created `overview.md`, `document-shape.md`, `topics-table.md`, `spec-id-as-ref.md`, `validation-policy.md`, and `follow-up-boundary.md`.
- Split the all-in-one `index_old.md` content by meaning: overview/front matter policy, document shape, Topics table, ID-as-ref, validation policy, and follow-up boundary.
- Did not edit PRODUCT-WORK-SPEC-001, migrate existing specs, change DRMCP implementation code, or modify `v01/`.
