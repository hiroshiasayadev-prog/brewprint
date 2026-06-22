# PRODUCT-TASK-SPEC-005-17: DRMCP Phase 2 — relocate namespace algorithm and ID grammar to namespace-model

- **id**: PRODUCT-TASK-SPEC-005-17
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-16
- **outputs**:
  - `product/records/spec/concepts/namespace-model/` (2 new sections or sibling files, depending on how WORK-009 split the directory)
  - `drmcp/records/spec/design-records-mcp/namespace-scanning.md` (core content replaced with pointer)
  - `drmcp/records/spec/design-records-mcp/schema/id-normalization.md` (content replaced with pointer)

## Goal

Relocate PRODUCT-owned namespace semantics currently hosted in two DRMCP spec files into `namespace-model/index.md` as new sections, then replace those DRMCP files with one-paragraph pointers. Per PRODUCT-TASK-SPEC-004-01 Phase 2 relocation plan, steps 1–2.

**Prerequisite**: `product/records/spec/concepts/namespace-model/index.md` must be format-migrated under PRODUCT-WORK-SPEC-009 before this task starts. Do not begin until WORK-009 has migrated that file.

## Work

| step | source | destination | notes |
|---|---|---|---|
| 1 | `drmcp/records/spec/design-records-mcp/namespace-scanning.md` §core content | `namespace-model/` — new section in `index.md` or a new sibling file (e.g. `v1-algorithm.md`), depending on how WORK-009 split the directory | Namespace_prefix derivation formula, kind-別 prefix table, multi-root scan default. Remove deferred relocation note. Replace DRMCP file body with pointer to the receiving spec's `id`. |
| 2 | `drmcp/records/spec/design-records-mcp/schema/id-normalization.md` | `namespace-model/` — new section or sibling file (e.g. `v1-id-grammar.md`) | Public ID / bare ID grammar tables. Remove deferred relocation note. Replace DRMCP file body with pointer. |
| — | Validator | Run `validate_spec.py` on both namespaces | `product/records/spec/concepts/namespace-model/` and `drmcp/records/spec/design-records-mcp/` must both pass `--strict`. |

## Done condition

| item | done when |
|---|---|
| namespace-model content added | `namespace-model/` has the relocated content in new sections or sibling files; `namespace-model/` overview Topics covers them. |
| DRMCP pointer files | `namespace-scanning.md` and `schema/id-normalization.md` retain their H1 / metadata / `## What this is` with a short pointer replacing the body; deferred relocation notes removed. |
| validator clean | Both `product/records/spec/concepts/namespace-model/` and `drmcp/records/spec/design-records-mcp/` pass `--strict`. |

## Verification

- Confirm no content was dropped — content in new PRODUCT sections matches the migrated source.
- Confirm DRMCP files cross-ref `spec:product.concepts.namespace_model` (correct spec ref, not a path).
- Confirm deferred relocation notes are removed from both DRMCP files.

## Evidence

- Created `product/records/spec/concepts/namespace-model/v1-namespace-algorithm.md` (`spec:product.concepts.namespace_model.v1_namespace_algorithm`) — namespace_prefix derivation formula, kind-level prefix table, multi-root scan behavior.
- Created `product/records/spec/concepts/namespace-model/v1-id-grammar.md` (`spec:product.concepts.namespace_model.v1_id_grammar`) — public ID grammar, bare ID grammar tables.
- Added 2 Topics rows to `namespace-model/index.md`.
- `drmcp/records/spec/design-records-mcp/namespace-scanning.md` — body replaced with redirect table pointing to `spec:product.concepts.namespace_model.v1_namespace_algorithm`; deferred relocation note removed.
- `drmcp/records/spec/design-records-mcp/schema/id-normalization.md` — body replaced with redirect table pointing to `spec:product.concepts.namespace_model.v1_id_grammar`; deferred relocation note removed.
- `validate_spec.py product/records/spec/concepts/namespace-model/ drmcp/records/spec/design-records-mcp/ --strict` → `[strict] All 38 file(s) OK.`
