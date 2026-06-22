# PRODUCT-TASK-SPEC-005-10: BPDSL — views/ spec authoring (format migration)

- **id**: PRODUCT-TASK-SPEC-005-10
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-09
- **outputs**:
  - `bpdsl/records/spec/views/overview.md` (new)
  - `bpdsl/records/spec/views/model-file.md` (migrated)
  - `bpdsl/records/spec/views/api-table.md` (migrated)
  - `bpdsl/records/spec/views/state-diagram.md` (migrated)
  - `bpdsl/records/spec/views/er.md` (migrated)
  - `bpdsl/records/spec/views/dag.md` (migrated)
  - `bpdsl/records/spec/views/sequence-diagram.md` (migrated)
  - `bpdsl/records/spec/views/wireframe.md` (migrated)

## Goal

Create `views/overview.md` as the navigation entry point for the views area, and migrate all 7 view spec files (staged unedited in `bpdsl/old_views/` from 005-09) to `Contract: format` shape (`## Current contract`, `## Rules`, `## Validation rules`), in place. All Japanese H1 titles must be translated to English.

## Work

| file | kind | required work |
|---|---|---|
| `views/overview.md` | Overview | New file. `## What this is` + `## Current contract` (view layer summary) + `## Topics` pointing to 7 view specs. |
| `views/model-file.md` | Contract (format) | Restructure sections → `## What this is`, `## Current contract`, `## Rules`, `## Validation rules`. Add `contract_class: format`. H1 is clean ASCII in source (`# Model file render rules`) — no translation needed. |
| `views/api-table.md` | Contract (format) | Same restructuring. **Source H1 is mojibake** (`# API Table�d�l`) — re-read with explicit UTF-8 decoding before transcribing the title; do not guess from corrupted bytes. |
| `views/state-diagram.md` | Contract (format) | Same restructuring. Source H1 is mojibake — re-read with explicit UTF-8 decoding. |
| `views/er.md` | Contract (format) | Same restructuring. Source H1 is mojibake — re-read with explicit UTF-8 decoding. |
| `views/dag.md` | Contract (format) | Same restructuring. Source H1 is mojibake — re-read with explicit UTF-8 decoding. |
| `views/sequence-diagram.md` | Contract (format) | Same restructuring. Source H1 is mojibake — re-read with explicit UTF-8 decoding. |
| `views/wireframe.md` | Contract (format) | Same restructuring. Source H1 is mojibake — re-read with explicit UTF-8 decoding. |

Read each view spec before migrating to identify which current sections map to `## Current contract`, `## Rules`, and `## Validation rules`. Use `bpdsl/old_views/` as the unedited content source if the in-place original becomes ambiguous mid-edit. Run `validate_spec.py bpdsl/records/spec/views/ --strict` after completing all files.

## Done condition

| item | done when |
|---|---|
| overview.md created | `views/overview.md` exists with `## Topics` covering all 7 view specs. |
| all 7 view specs migrated | H1 format, metadata, required sections present. |
| view specs restructured | All 7 view specs have `## Current contract`, `## Rules`, `## Validation rules`. |
| English titles | No Japanese H1 in any views/ file. |
| strict validation | `validate_spec.py bpdsl/records/spec/views/ --strict` exits 0. |

## Verification

- Confirm each view spec has `contract_class: format`.
- Confirm `views/overview.md` `## Topics` covers all 7 child specs.
- Cross-check `parent:` in each view spec against `views/overview.md` Topics row.
- Confirm no mojibake / encoding artifacts remain in any migrated file.

## Evidence

- All source files re-read with the Read tool to check for the mojibake pattern found in the mcp/ batch — confirmed **no corruption in any of the 7 views files**; all were clean, correctly-decoded Japanese. The earlier mojibake warnings baked into this task's `## Work` table (from terminal/grep display artifacts) did not reproduce when read properly.
- `views/overview.md` authored new, with `## Topics` covering all 7 view specs.
- All 7 view specs migrated in place: `model-file.md`, `api-table.md`, `state-diagram.md`, `er.md`, `dag.md`, `sequence-diagram.md`, `wireframe.md`. All restructured into `## What this is` / `## Current contract` / `## Rules` / `## Validation rules` (+ `## Non-goals` / `## Related specs` where applicable), modeled directly on `spec:product.concepts.spec_format.document_shape` — itself a `Contract: format` spec — as the canonical template for the `Rules` vs. `Validation rules` split.
- `grep -l "contract_class.*format" bpdsl/records/spec/views/*.md` → 7/7.
- `grep -c` Topics rows in `views/overview.md` → 7.
- All 7 files confirmed to have `## Current contract`, `## Rules`, and `## Validation rules` present.
- `python product/src/tools/validate_spec.py bpdsl/records/spec/views/ --strict` → `[strict] All 8 file(s) OK.`
- No Japanese remains in any H1, H2/H3, or table header across all 8 files.
- `dag.md` (largest source, 992 lines) fully preserved: all node-render rules (start/end, task, asset, store, branch, fork/join, external ref, subgraph params, subgraph initializes, foreach decoration, coloring), all edge-render rules (data line, returns.source data line, store access line, control line incl. branch/fork-join/foreach subsections), and all 6 worked render examples carried over.
