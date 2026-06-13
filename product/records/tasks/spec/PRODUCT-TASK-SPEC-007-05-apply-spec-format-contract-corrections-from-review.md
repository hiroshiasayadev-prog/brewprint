# PRODUCT-TASK-SPEC-007-05: Apply spec-format contract corrections from review

- **id**: PRODUCT-TASK-SPEC-007-05
- **status**: done
- **date**: 2026-06-13
- **work_item**: PRODUCT-WORK-SPEC-007
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **source_work_item**: PRODUCT-WORK-SPEC-007
- **estimate**: 0.6d
- **depends_on**:
  - PRODUCT-TASK-SPEC-007-04
- **outputs**:
  - product/records/spec/concepts/spec-format/document-shape.md
  - product/records/spec/concepts/spec-format/topics-table.md
  - product/records/spec/concepts/spec-format/validation-policy.md
  - product/records/work-items/spec/PRODUCT-WORK-SPEC-002-stable-spec-id-as-ref-and-derived-topic-compatibility.md

## Goal

Apply the spec-format contract corrections found during split review.

## Scope

| area | in scope |
|---|---|
| spec provenance sections | Remove `## Source records` sections from spec-format child specs. |
| required section matrix | Update the required section matrix to remove the Source records recommendation. |
| contract subtype marker | Add a `contract_class` marker for `Contract` specs. |
| interface contracts | Define `contract_class: interface` for API, tool, and external-boundary contracts. |
| format contracts | Define `contract_class: format` for document, table, metadata, and validation-shape contracts. |
| Contract section expectations | Require `## Request` and `## Response` only for `contract_class: interface` contracts. |
| format contract sections | Rewrite format contracts to use `## Current contract`, `## Rules`, and `## Validation rules` rather than `## Request` / `## Response`. |
| Topics table target column | Change the Topics table contract from a `file` column to a `ref` column. |
| spec ref mapping | Define path-derived canonical spec refs and one-to-one ref-to-path / path-to-ref mapping. |
| rename and move behavior | Change rename/move behavior so the spec ID changes with location. |
| ID mismatch validation | Change validation policy so path-derived ID mismatch is an error for new/migrated specs and a warning only during inventory, migration, or transient work. |
| PRODUCT-WORK-SPEC-002 wording | Reframe PRODUCT-WORK-SPEC-002 away from stable-ID-survives-rename/move and toward path-derived canonical spec refs and a ref-first topic index. |

## Non-scope

| area | reason |
|---|---|
| `index_old.md` deletion | Reserved for PRODUCT-TASK-SPEC-007-06 after corrections are applied and reviewed. |
| PRODUCT-WORK-SPEC-007 closure | Reserved for PRODUCT-TASK-SPEC-007-06. |
| `v01/` edits | Historical snapshot; not part of this correction task. |
| DRMCP implementation patches | Later DRMCP implementation work owns durable tooling changes. |
| unrelated existing spec migration | PRODUCT-WORK-SPEC-005 owns broader migration. |

## Work

Update the split spec-format child specs and directly related work-item wording to reflect the review decisions recorded by PRODUCT-TASK-SPEC-007-04.

## Done condition

| item | done when |
|---|---|
| child specs corrected | Spec-format child specs reflect the corrected Source records, contract subtype, Topics ref, path-derived ID, and validation-severity design. |
| PRODUCT-WORK-SPEC-002 reframed | PRODUCT-WORK-SPEC-002 no longer implies stable IDs survive rename/move by default and is consistent with path-derived canonical spec refs. |
| cleanup still pending | `index_old.md` remains available until PRODUCT-TASK-SPEC-007-06. |

## Verification

- Confirm spec-format child specs no longer recommend or carry `## Source records` as reverse traceability.
- Confirm format contracts use `contract_class: format` and do not require `## Request` / `## Response`.
- Confirm Topics table uses `ref` as canonical child target data rather than `file`.
- Confirm ID mismatch severity distinguishes new/migrated specs from inventory/migration/transient work.
- Confirm PRODUCT-WORK-SPEC-002 wording aligns with path-derived canonical spec refs.

## Evidence

- Created from PRODUCT-TASK-SPEC-007-04 review findings.
- This task intentionally keeps cleanup and work-item closure separate so major corrections can be reviewed before deleting `index_old.md`.
- Removed `## Source records` from split spec-format specs and replaced reverse traceability with ADR-governed traceability policy owned by PRODUCT-WORK-SPEC-008.
- Updated the parent Index and Topics table contract to use required columns `title`, `kind`, `ref`, and `summary`; `file` and row-level `parent` are not canonical required columns.
- Reframed spec ID behavior around path-derived canonical refs: path/ref mapping is one-to-one for new or migrated specs, and move/rename changes the canonical spec ID unless a later compatibility exception applies.
- Added `contract_class: format` to format Contract specs and updated the document-shape matrix so only `contract_class: interface` requires `## Request` and `## Response`.
- Rewrote `document-shape.md`, `topics-table.md`, and `validation-policy.md` away from artificial `## Request` / `## Response` sections.
- Updated validation severity so H1-adjacent `id` mismatch with path-derived canonical ref is an error for new/migrated specs and a warning during inventory, migration, or transient working states.
- Reframed PRODUCT-WORK-SPEC-002 as `Path-derived canonical spec refs and ref-first topic index` while keeping status `not_started`.
