# PRODUCT-TASK-SPEC-005-18: DRMCP Phase 2 — relocate discovery path patterns to repository-layout

- **id**: PRODUCT-TASK-SPEC-005-18
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-16
- **outputs**:
  - `product/records/spec/concepts/repository-layout/` (1 new section or sibling file, depending on how WORK-009 split the directory)
  - `drmcp/records/spec/design-records-mcp/schema/discovery.md` (path-pattern content replaced; DRMCP kind-filter stays)

## Goal

Relocate discovery path-pattern conventions (PRODUCT-owned) from `schema/discovery.md` into `repository-layout/index.md` as a new section. The DRMCP kind-inclusion filter (e.g. requiring `design_record` front matter) is not relocated — it stays in the remaining `discovery.md` body. Per PRODUCT-TASK-SPEC-004-01 Phase 2 relocation plan, step 3.

**Prerequisite**: `product/records/spec/concepts/repository-layout/index.md` must be format-migrated under PRODUCT-WORK-SPEC-009 before this task starts. Do not begin until WORK-009 has migrated that file.

## Work

| step | source | destination | notes |
|---|---|---|---|
| 3 | `drmcp/records/spec/design-records-mcp/schema/discovery.md` path-pattern content | `repository-layout/` — new section in `index.md` or a new sibling file (e.g. `record-discovery-paths.md`), depending on how WORK-009 split the directory | Kind → `<records_root>/.../*.md` path-pattern table. DRMCP kind-inclusion filter (design_record front-matter requirement) stays in the remaining DRMCP file body. Remove deferred relocation note from the portion moved. |
| — | Validator | Run `validate_spec.py` on both namespaces | Both must pass `--strict`. |

## Done condition

| item | done when |
|---|---|
| repository-layout content added | `repository-layout/` has the path-pattern content in a new section or sibling file; overview Topics covers it. |
| DRMCP file trimmed | `schema/discovery.md` retains H1 / metadata / `## What this is` (updated), DRMCP kind-filter content, and a cross-ref to `spec:product.concepts.repository_layout`; deferred relocation note for the relocated portion removed. |
| validator clean | Both namespaces pass `--strict`. |

## Verification

- Confirm path-pattern convention moved completely — no duplication between PRODUCT and DRMCP.
- Confirm DRMCP kind-filter content (design_record front-matter requirement) is still present in `discovery.md`.
- Confirm cross-ref uses canonical `spec:product.concepts.repository_layout`, not a path.

## Evidence

- `product/records/spec/concepts/repository-layout/index.md` — format-migrated from old YAML front-matter format to new spec format (Overview, id `spec:product.concepts.repository_layout`). Was not in WORK-009 scope; migrated as part of this task.
- Created `product/records/spec/concepts/repository-layout/record-discovery-paths.md` (`spec:product.concepts.repository_layout.record_discovery_paths`) — kind-level path-pattern table relocated from DRMCP.
- `drmcp/records/spec/design-records-mcp/schema/discovery.md` — path-pattern table removed; DRMCP-specific kind-filter inclusion conditions retained; pointer to `spec:product.concepts.repository_layout.record_discovery_paths` added; deferred relocation note removed.
- `validate_spec.py product/records/spec/concepts/repository-layout/ drmcp/records/spec/design-records-mcp/ --strict` → `[strict] All 32 file(s) OK.`
