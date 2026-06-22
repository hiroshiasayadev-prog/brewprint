# PRODUCT-TASK-SPEC-005-06: BPDSL — mcp/ spec authoring (format migration)

- **id**: PRODUCT-TASK-SPEC-005-06
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-05
- **outputs**:
  - `bpdsl/records/spec/mcp/overview.md` (migrated)
  - `bpdsl/records/spec/mcp/errors.md` (migrated)
  - `bpdsl/records/spec/mcp/schema.md` (migrated)
  - `bpdsl/records/spec/mcp/versioning.md` (migrated)
  - `bpdsl/records/spec/mcp/tools/get-source.md` (migrated)
  - `bpdsl/records/spec/mcp/tools/get-signature.md` (migrated)
  - `bpdsl/records/spec/mcp/tools/get-references.md` (migrated)
  - `bpdsl/records/spec/mcp/tools/get-reference-tree.md` (migrated)
  - `bpdsl/records/spec/mcp/tools/analyze-impact.md` (migrated)
  - `bpdsl/records/spec/mcp/tools/inspect.md` (migrated)
  - `bpdsl/records/spec/mcp/tools/list-objects.md` (migrated)
  - `bpdsl/records/spec/mcp/tools/list-endpoints.md` (migrated)

## Goal

Migrate all 12 files under `bpdsl/records/spec/mcp/` to the accepted spec format, in place. Use `bpdsl/old_mcp/` (staged in 005-05) as the unedited content source for comparison. Tool specs (`mcp/tools/*.md`) require restructuring to `Contract: interface` shape (`## Request`, `## Response`, `## Errors`).

## Work

| file | kind | required work |
|---|---|---|
| `mcp/overview.md` | Overview | Add H1 format, metadata, `## What this is`, `## Current contract` (design principles), `## Topics`. **Source H1 is mojibake** (`# MCP�d�l Overview`) — re-read the file with explicit UTF-8 decoding per the project's encoding rule before transcribing the title; do not guess the original text from corrupted bytes. |
| `mcp/errors.md` | Reference | Add H1 format, metadata, required sections. H1 is clean ASCII in source. |
| `mcp/schema.md` | Reference | Add H1 format, metadata, required sections. H1 is clean ASCII in source. |
| `mcp/versioning.md` | Reference | Add H1 format, metadata, required sections. H1 is clean ASCII in source. |
| `mcp/tools/*.md` (8 files) | Contract (interface) | Restructure numbered sections → `## What this is`, `## Request`, `## Response`, `## Errors`. Add `contract_class: interface`. H1s are clean ASCII (backtick tool names) in source. |

Run `validate_spec.py bpdsl/records/spec/mcp/ --strict` after completing all files.

## Done condition

| item | done when |
|---|---|
| all 12 files migrated | H1 format, metadata, required sections present in all files. |
| tool specs restructured | All 8 tool specs have `## Request`, `## Response`, `## Errors` H2 sections. |
| strict validation | `validate_spec.py bpdsl/records/spec/mcp/ --strict` exits 0. |
| YAML front matter removed | No front matter in any output file. |

## Verification

- Confirm each tool spec has `contract_class: interface`.
- Confirm `mcp/overview.md` `## Topics` covers all 11 child specs (3 non-tool + 8 tools).
- Confirm no mojibake / encoding artifacts remain in any migrated file, especially `overview.md`.

## Evidence

- All 12 files migrated in place under `bpdsl/records/spec/mcp/`: `overview.md`, `errors.md`, `schema.md`, `versioning.md` (Reference/Overview), plus 8 `tools/*.md` (Contract: interface).
- **Mojibake finding from the task description was wrong** — re-reading every source file with the Read tool (proper UTF-8 decode) showed clean, uncorrupted Japanese throughout, including `mcp/overview.md`'s H1 (`# MCP仕様 Overview`). The garbled bytes seen earlier (`# MCP�d�l Overview`) came from `grep`/Bash printing through the Windows console codepage, not actual file corruption. No re-read-and-guess was needed; all titles translated from the correctly-decoded source.
- `python product/src/tools/validate_spec.py bpdsl/records/spec/mcp/ --strict` → `[strict] All 12 file(s) OK.`
- `grep -l "contract_class.*interface" bpdsl/records/spec/mcp/tools/*.md` → 8/8 tool specs have `contract_class: interface`.
- `mcp/overview.md` `## Topics` has 11 rows (3 non-tool Reference specs + 8 tool Contract specs), matching the done condition.
- YAML front matter removed from all 12 files; H1-adjacent metadata (`id`/`status`/`date`/`parent`, plus `contract_class` on tool specs) added.
- All numbered source sections (`## 1. Purpose`, `## 2. Input`, `## 3. Output` etc.) restructured into `## What this is` / `## Request` / `## Response` / `## Errors`, with supplementary sections (selector support, design principles, etc.) kept as additional H2s.
- Content fidelity: all ADR source citations, selector-support-matrix rows, error-code tables, and worked JSON examples carried over; no section silently dropped.
