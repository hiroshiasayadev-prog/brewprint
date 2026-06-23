# PRODUCT-TASK-SPEC-011-04: Create requirement authoring guide

- **id**: PRODUCT-TASK-SPEC-011-04
- **status**: done
- **date**: 2026-06-23
- **work_item**: PRODUCT-WORK-SPEC-011
- **source_requirement**: PRODUCT-REQ-SPEC-002
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-011-02
  - PRODUCT-TASK-SPEC-011-03
- **outputs**:
  - `product/records/spec/concepts/authoring-standards/requirement-authoring.md`
  - `product/records/spec/concepts/authoring-standards/index.md`

## Goal

Create the canonical requirement authoring guide under PRODUCT authoring standards.

## Work

- Use the common per-artifact guide structure.
- Define namespace-aware requirement ID grammar and file layout.
- Define canonical English body headings.
- Define requirement metadata meaning and requiredness for create, partial update, and persisted state.
- Define requirement status lifecycle.
- Define requirement-specific writing boundaries.
- Reference `artifact_boundary` for cross-artifact selection.
- Define author-facing exact-ID and `new` forms.
- Exclude concrete DRMCP request, response, and diagnostic schemas.
- Update the authoring-standards Index.

## Done condition

| item | done when |
|---|---|
| Guide published | `requirement-authoring.md` exists as a Reference spec. |
| Common shape applied | The guide follows PRODUCT-WORK-SPEC-011 section structure. |
| English headings canonical | All prescribed requirement headings use English. |
| Metadata states separated | Create, partial update, and persisted requirements are explicit. |
| Boundary clear | Requirement ownership is distinguished from investigation, ADR, spec, work item, and task. |
| Index updated | The guide appears with its canonical ref and current summary. |

## Verification

- Confirm the guide uses abstract v2 IDs as primary forms.
- Confirm requirement headings and table headers use English.
- Confirm status and metadata rules match canonical PRODUCT semantics.
- Confirm the guide cites `artifact_boundary` and does not duplicate the full ownership matrix.
- Confirm no current DRMCP operating status appears.

## Evidence

- `product/records/spec/concepts/authoring-standards/requirement-authoring.md` created as a Reference spec.
- Common guide shape from PRODUCT-WORK-SPEC-011 applied: all sections present in order.
- All prescribed headings use English.
- Metadata schema distinguishes create input, partial update, and persisted state.
- `id` field: persisted requirement carries explicit `id` in bullet metadata (unlike ADR); create input does not supply `id` as a metadata field — MCP generates it from the top-level exact ID or `new` resolution.
- File path corrected to include domain subdirectory per `spec:product.concepts.repository_layout.record_discovery_paths` discovery pattern.
- `source_refs` meaning broadened to canonical artifact IDs or active semantic refs, not limited to investigations/requirements; raw observations directed to `## Evidence`.
- `rejected` terminal-state rule removed — lifecycle transition table is not established in the source specs.
- `date` update condition aligned between Metadata rules and Update section (both now include evidence base).
- `## Required Outcome` required only when `status: accepted`.
- Guide cites `spec:product.concepts.authoring_standards.artifact_boundary`; does not duplicate the ownership matrix.
- No DRMCP operating status recorded.
- `product/records/spec/concepts/authoring-standards/index.md` updated with new entry.
