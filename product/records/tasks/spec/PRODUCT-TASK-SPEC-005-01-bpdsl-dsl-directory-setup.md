# PRODUCT-TASK-SPEC-005-01: BPDSL — DSL directory setup and file staging

- **id**: PRODUCT-TASK-SPEC-005-01
- **status**: done
- **date**: 2026-06-16
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - `bpdsl/old/` (staging directory for pre-migration flat files)
  - `bpdsl/records/spec/overview.md` (Topics updated to 3 rows)
  - `bpdsl/records/spec/dsl/` (empty directory, ready for authoring)
  - `bpdsl/records/spec/dsl/design-philosophy.md` (moved from root, id updated)

## Goal

Establish the `dsl/` directory under `bpdsl/records/spec/` to mirror the `mcp/` and `views/` peer structure, stage the 8 flat root-level BPDSL spec files for reference during migration, and update the top-level `overview.md` Topics table to the 3-category shape (dsl / mcp / views).

## Work

| area | required work |
|---|---|
| staging | Move 8 flat root-level files (`nodes.md`, `edges.md`, `file-types.md`, `naming.md`, `type-ref.md`, `project-layout.md`, `diagnostics.md`, old `overview.md`) to `bpdsl/old/`. Do not edit them there. |
| design-philosophy.md | Move `bpdsl/records/spec/design-philosophy.md` → `bpdsl/records/spec/dsl/design-philosophy.md`. Update `id` to `spec:bpdsl.dsl.design_philosophy` and `parent` to `spec:bpdsl.dsl.overview`. |
| overview.md Topics | Reduce top-level `bpdsl/records/spec/overview.md` Topics table from 16 rows to 3: `spec:bpdsl.dsl.overview`, `spec:bpdsl.mcp.overview`, `spec:bpdsl.views.overview`. |
| validator | Run `validate_spec.py bpdsl/records/spec/overview.md bpdsl/records/spec/dsl/design-philosophy.md --strict` and confirm 0 errors. |

## Done condition

| item | done when |
|---|---|
| staging | `bpdsl/old/` exists and contains the 8 pre-migration flat files. |
| dsl/ exists | `bpdsl/records/spec/dsl/` directory exists. |
| design-philosophy.md | File is at `dsl/design-philosophy.md` with updated id and parent. |
| overview.md | Top-level Topics has exactly 3 rows. |
| validator | `--strict` passes on modified files. |

## Verification

- No files under `bpdsl/old/` are edited (read-only staging area).
- `bpdsl/records/spec/nodes.md` and other root-level flat files are gone from root (moved to `bpdsl/old/`).
- `bpdsl/records/spec/dsl/design-philosophy.md` id matches `spec:bpdsl.dsl.design_philosophy`.

## Evidence

| item | detail |
|---|---|
| bpdsl/old/ created | 7 flat files staged: nodes.md, edges.md, file-types.md, naming.md, type-ref.md, project-layout.md, diagnostics.md |
| bpdsl/records/spec/dsl/ created | Empty directory ready for authoring |
| design-philosophy.md moved | `bpdsl/records/spec/dsl/design-philosophy.md`. id updated to `spec:bpdsl.dsl.design_philosophy`, parent to `spec:bpdsl.dsl.overview`. Related specs refs updated. |
| overview.md updated | Topics reduced to 3 rows: spec:bpdsl.dsl.overview / spec:bpdsl.mcp.overview / spec:bpdsl.views.overview. Current contract trimmed to domain summary table. |
| validator | `--strict` passes on overview.md and dsl/design-philosophy.md. |
