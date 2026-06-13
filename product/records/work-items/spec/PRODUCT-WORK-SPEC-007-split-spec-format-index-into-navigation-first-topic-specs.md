# PRODUCT-WORK-SPEC-007: Split spec-format index into navigation-first topic specs

- **id**: PRODUCT-WORK-SPEC-007
- **status**: done
- **date**: 2026-06-11
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-001
- **task_refs**:
  - PRODUCT-TASK-SPEC-007-01
  - PRODUCT-TASK-SPEC-007-02
  - PRODUCT-TASK-SPEC-007-03
  - PRODUCT-TASK-SPEC-007-04
  - PRODUCT-TASK-SPEC-007-05
  - PRODUCT-TASK-SPEC-007-06

## Summary

Split `product/records/spec/concepts/spec-format/index.md` from an all-in-one spec format contract into a navigation-first Index plus focused child topic specs.

This work corrects the self-example produced by PRODUCT-WORK-SPEC-001 without reopening PRODUCT-WORK-SPEC-001. The goal is to make the spec-format area demonstrate the same anti-bloat structure that it requires from future specs.

## Scope

| area | in scope |
|---|---|
| temporary source preservation | Keep the previous all-in-one file as `index_old.md` until split review completes. |
| navigation-first index | Recreate `index.md` as `# Index: Spec format` with concise navigation content and `## Topics`. |
| child topic specs | Split substantive content into focused spec files under `product/records/spec/concepts/spec-format/`. |
| content parity review | Compare split specs against `index_old.md` before deleting the temporary file. |
| cleanup | Remove `index_old.md` before closing this work item. |

## Non-scope

| area | owner |
|---|---|
| PRODUCT-WORK-SPEC-001 reopening | PRODUCT-WORK-SPEC-001 remains done; this work is a follow-up correction. |
| path-derived spec ref compatibility exceptions | PRODUCT-WORK-SPEC-002. |
| ADR `target_specs` traceability investigation | PRODUCT-WORK-SPEC-008. |
| authoring guide update | PRODUCT-WORK-SPEC-003. |
| existing spec migration | PRODUCT-WORK-SPEC-005. |
| temporary validator tooling | PRODUCT-WORK-SPEC-006. |
| current DRMCP implementation changes | Deferred to later DRMCP redesign/reimplementation work. |
| v01 updates | `v01/` is historical and must not be modified. |

## Planned split

| file | kind | role |
|---|---|---|
| `index.md` | Index | Navigation-first entry and authoritative `## Topics` table. |
| `overview.md` | Overview | Purpose, scope, non-goals, and current capability summary. |
| `document-shape.md` | Contract | H1 format, H1-adjacent metadata, spec kinds, and required section matrix. |
| `topics-table.md` | Contract | `## Topics` table columns, parent row behavior, and parent declaration rules. |
| `spec-id-as-ref.md` | Concept | Path-derived canonical `spec:` refs, move/rename ID behavior, and underscore policy. |
| `validation-policy.md` | Contract | Parser-aware validation rules and migration warning/error policy. |
| `follow-up-boundary.md` | Concept | PRODUCT / DRMCP ownership boundary and follow-up ordering. |

## Done condition

| item | done when |
|---|---|
| source preserved | PRODUCT-TASK-SPEC-007-01 is done and `index_old.md` exists as temporary review source. |
| split specs created | Focused child spec files exist and carry the substantive contract content from `index_old.md`. |
| index rebuilt | `index.md` is navigation-first and does not contain detailed contract body. |
| parity reviewed | Split content is reviewed against `index_old.md` for unintentional loss, duplication, and responsibility mixing. |
| corrections applied | Major design corrections from split review are applied and reviewed before cleanup. |
| cleanup complete | `index_old.md` is deleted and duplicate `spec:product.concepts.spec_format` identity risk is gone. |
| no scope creep | No existing spec migration, DRMCP implementation patch, or `v01/` update is performed. |

## Source records

| ref | role |
|---|---|
| PRODUCT-REQ-SPEC-001 | Requirement for MCP-readable spec format and topic tree support. |
| PRODUCT-WORK-SPEC-001 | Completed work that produced the initial all-in-one format contract and follow-up split. |
| `spec:product.concepts.spec_format` | Target spec area to restructure into navigation-first topic specs. |

## Evidence

- PRODUCT-TASK-SPEC-007-01 created as done after the user renamed `index.md` to `index_old.md`.
- PRODUCT-TASK-SPEC-007-02 created as done after splitting `index_old.md` into six focused child topic specs.
- PRODUCT-TASK-SPEC-007-03 created as done after rebuilding `index.md` as a navigation-first Index with child `## Topics` rows.
- PRODUCT-TASK-SPEC-007-04 completed the split review and recorded major findings before cleanup.
- PRODUCT-TASK-SPEC-007-05 created to apply the design corrections found by the split review.
- PRODUCT-TASK-SPEC-007-05 completed the spec-format corrections: Source records sections removed from specs, Topics table changed to `ref`, path-derived canonical spec refs defined, `contract_class` added for format contracts, and PRODUCT-WORK-SPEC-002 reframed.
- PRODUCT-TASK-SPEC-007-06 created as the reserved cleanup and close task after corrections are applied and reviewed.
- PRODUCT-TASK-SPEC-007-06 completed cleanup: `index_old.md` was removed, PRODUCT-WORK-SPEC-002 file path was renamed to match its path-derived canonical ref topic title, and this work item was closed.
- Final scope boundary held: PRODUCT-WORK-SPEC-001 was not reopened, `v01/` was not modified, DRMCP implementation code was not changed, and unrelated existing specs were not migrated.
