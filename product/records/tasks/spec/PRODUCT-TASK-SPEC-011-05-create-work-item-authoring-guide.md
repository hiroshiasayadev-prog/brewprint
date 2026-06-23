# PRODUCT-TASK-SPEC-011-05: Create work-item authoring guide

- **id**: PRODUCT-TASK-SPEC-011-05
- **status**: done
- **date**: 2026-06-23
- **work_item**: PRODUCT-WORK-SPEC-011
- **source_requirement**: PRODUCT-REQ-SPEC-002
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-011-02
  - PRODUCT-TASK-SPEC-011-03
  - PRODUCT-TASK-SPEC-011-04
- **outputs**:
  - `product/records/spec/concepts/authoring-standards/work-item-authoring.md`
  - `product/records/spec/concepts/authoring-standards/index.md`

## Goal

Create the canonical work-item authoring guide under PRODUCT authoring standards.

## Work

- Use the common per-artifact guide structure.
- Define namespace-aware work-item ID grammar and file layout.
- Define canonical English body headings.
- Define work-item metadata meaning and requiredness for create, partial update, and persisted state.
- Define work-item status lifecycle.
- Define requirement linkage, impact tracking, task graph, and completion evidence rules.
- Distinguish work-item goal state from task-level execution details.
- Reference `artifact_boundary` for cross-artifact selection.
- Define author-facing exact-ID and `new` forms.
- Exclude concrete DRMCP request, response, and diagnostic schemas.
- Update the authoring-standards Index.

## Done condition

| item | done when |
|---|---|
| Guide published | `work-item-authoring.md` exists as a Reference spec. |
| Common shape applied | The guide follows PRODUCT-WORK-SPEC-011 section structure. |
| English headings canonical | All prescribed work-item headings use English. |
| Metadata states separated | Create, partial update, and persisted requirements are explicit. |
| Workflow boundary clear | Requirement linkage, task graph, progress, and completion responsibilities are precise. |
| Index updated | The guide appears with its canonical ref and current summary. |

## Verification

- Confirm the guide uses abstract v2 IDs as primary forms.
- Confirm work-item headings and table headers use English.
- Confirm reciprocal requirement and task relations are expressed semantically, not as DRMCP implementation details.
- Confirm the work item owns cross-cutting progress but not subordinate task status as canonical state.
- Confirm no current DRMCP operating status appears.

## Evidence

- `product/records/spec/concepts/authoring-standards/work-item-authoring.md` created as a Reference spec.
- Common guide shape from PRODUCT-WORK-SPEC-011 applied: all sections present in order.
- All prescribed headings use English.
- Metadata schema distinguishes create input, partial update, and persisted state.
- `id` field: persisted work item carries explicit `id` in bullet metadata; create input does not supply `id` as a metadata field — MCP generates it from the top-level exact ID or `new` resolution.
- `source_requirement` marked as correction-only in partial update with three correction rules: preserve reciprocal Requirement linkage; align child Task `source_requirement` values.
- Persisted workflow relation invariants stated: source Requirement lists this Work Item in `work_items`; every listed Task references this Work Item in `work_item`; every child Task shares the same `source_requirement`. Responsibility for reciprocal updates deferred to DRMCP contracts.
- Body section table separates heading presence (always required) from substantive content (required only for Goal, Boundary, Evidence when `done`). Remaining sections may remain `TBD`.
- `TBD` placeholder prohibition added: `TBD` must not remain in `## Goal`, `## Boundary`, or `## Evidence` when `status` is `done`.
- Tooling-specific rationale ("ensures named-section updates can target sections without filesystem edits") removed; replaced with normative "preserve a stable document shape."
- `impact_refs` vs `## Impact Scope` distinction recorded: metadata carries machine-readable refs; body section carries human-readable impact descriptions.
- Task flow defined as a view, not the canonical task status source; no task status duplication rule stated.
- `done` gate: `## Goal`, `## Boundary`, `## Evidence` must be non-empty — sourced from v01 guide validation contract.
- Guide cites `spec:product.concepts.authoring_standards.artifact_boundary`; does not duplicate the ownership matrix.
- No DRMCP operating status recorded.
- `product/records/spec/concepts/authoring-standards/index.md` updated with new entry.
