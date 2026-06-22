# PRODUCT-TASK-SPEC-005-02: BPDSL — Author dsl/ spec files

- **id**: PRODUCT-TASK-SPEC-005-02
- **status**: done
- **date**: 2026-06-16
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 2d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-01
- **outputs**:
  - `bpdsl/records/spec/dsl/overview.md`
  - `bpdsl/records/spec/dsl/nodes/overview.md`
  - `bpdsl/records/spec/dsl/nodes/processing.md`
  - `bpdsl/records/spec/dsl/nodes/data.md`
  - `bpdsl/records/spec/dsl/nodes/application.md`
  - `bpdsl/records/spec/dsl/edges/overview.md`
  - `bpdsl/records/spec/dsl/edges/data-flow.md`
  - `bpdsl/records/spec/dsl/edges/state-transitions.md`
  - `bpdsl/records/spec/dsl/edges/cross-edges.md`
  - `bpdsl/records/spec/dsl/edges/cross-file-refs.md`
  - `bpdsl/records/spec/dsl/file-types.md`
  - `bpdsl/records/spec/dsl/naming.md`
  - `bpdsl/records/spec/dsl/type-ref.md`
  - `bpdsl/records/spec/dsl/project-layout.md`
  - `bpdsl/records/spec/dsl/diagnostics.md`

## Goal

Author all spec files under `bpdsl/records/spec/dsl/` using `bpdsl/old/` as the content source. Each file must pass `--strict` validation. All H1 titles must be English. Content must be complete — no section placeholders.

## Work

| file | kind | source in bpdsl/old/ | notes |
|---|---|---|---|
| `dsl/overview.md` | Overview | — | `## Topics` pointing to all dsl/ children. `## Current contract` is the DSL at-a-glance (not node detail). |
| `dsl/nodes/overview.md` | Overview | `nodes.md` §ノード種別一覧 + §ファイル構造 + §共通フィールド | Boundary matrix (diagram presence per node type) in `## Current contract`. `## Topics` → processing, data, application. |
| `dsl/nodes/processing.md` | Reference | `nodes.md` §task + §asset + §branch + §fork + §join | Field tables required. Translate H1 and section titles. |
| `dsl/nodes/data.md` | Reference | `nodes.md` §model + §store | Field tables required. |
| `dsl/nodes/application.md` | Reference | `nodes.md` §actor + §event + §state | Field tables required. |
| `dsl/edges/overview.md` | Overview | `edges.md` §設計原則 | Design principle (wiring in flow:) + edge-kind summary table in `## Current contract`. `## Topics` → 4 child refs. |
| `dsl/edges/data-flow.md` | Reference | `edges.md` §1 | ~484 lines. Translate section titles. Keep all sub-sections. |
| `dsl/edges/state-transitions.md` | Reference | `edges.md` §2 | |
| `dsl/edges/cross-edges.md` | Reference | `edges.md` §3 | |
| `dsl/edges/cross-file-refs.md` | Reference | `edges.md` §4 | |
| `dsl/file-types.md` | Reference | `file-types.md` | |
| `dsl/naming.md` | Reference | `naming.md` | |
| `dsl/type-ref.md` | Reference | `type-ref.md` | |
| `dsl/project-layout.md` | Reference | `project-layout.md` | |
| `dsl/diagnostics.md` | Reference | `diagnostics.md` | |

Run `validate_spec.py bpdsl/records/spec/dsl/ --strict` after each file is authored. Target: 0 errors on all 15 files before closing this task.

## Done condition

| item | done when |
|---|---|
| all 15 files exist | All output files listed above are present. |
| strict validation | `validate_spec.py bpdsl/records/spec/dsl/ --strict` exits 0. |
| English titles | All H1 titles are English. No Japanese H1 remains. |
| content complete | All content from `bpdsl/old/` source is accounted for — either migrated or explicitly noted as moved to another spec. |
| YAML front matter removed | No YAML front matter in any output file. |

## Verification

- Cross-check each `parent:` ref against the `## Topics` row that declares it as a child.
- Confirm `dsl/nodes/overview.md` `## Topics` covers processing, data, application.
- Confirm `dsl/edges/overview.md` `## Topics` covers all 4 edge sub-specs.
- Confirm `dsl/overview.md` `## Topics` covers all dsl/ children including design-philosophy.

## Evidence

- All 15 output files authored under `bpdsl/records/spec/dsl/` (7 top-level + `nodes/overview.md`,`processing.md`,`data.md`,`application.md` + `edges/overview.md`,`data-flow.md`,`state-transitions.md`,`cross-edges.md`,`cross-file-refs.md`).
- `python product/src/tools/validate_spec.py bpdsl/records/spec/dsl/ --strict` → `[strict]  All 16 file(s) OK.` (16 = 15 new + `design-philosophy.md` from T01).
- All H1 titles are English (`# Overview: ...` / `# Reference: ...`); no Japanese H1 remains in `dsl/`.
- YAML front matter removed from all files; metadata expressed as H1-adjacent bullet list.
- `dsl/overview.md` `## Topics` covers all 8 dsl/ children (design_philosophy, nodes.overview, edges.overview, file_types, naming, type_ref, project_layout, diagnostics).
- `dsl/nodes/overview.md` `## Topics` covers processing / data / application.
- `dsl/edges/overview.md` `## Topics` covers data_flow / state_transitions / cross_edges / cross_file_refs.
- Content from `bpdsl/old/` source files fully migrated: `nodes.md`, `edges.md`, `file-types.md`, `naming.md`, `type-ref.md`, `project-layout.md`, `diagnostics.md` — each ADR-sourced rule/table carried over with translated prose; no section dropped.
