# PRODUCT-TASK-SPEC-005-05: BPDSL — mcp/ directory setup and file staging

- **id**: PRODUCT-TASK-SPEC-005-05
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-04
- **outputs**:
  - `bpdsl/old_mcp/` (staging directory — byte-identical snapshot of pre-migration mcp/ files)

## Goal

Snapshot the current `bpdsl/records/spec/mcp/` tree into `bpdsl/old_mcp/` before any migration edits, so the Opus review task (005-07) can diff pre- and post-migration content. The `mcp/` directory layout itself is not restructured — files are migrated in place in 005-06. (Suffix `_mcp` distinguishes this staging directory from a future `bpdsl/old_drmcp/` staged in a separate DRMCP migration session.)

## Work

| area | required work |
|---|---|
| staging | Copy all 12 files under `bpdsl/records/spec/mcp/` (including the `tools/` subdirectory) to `bpdsl/old_mcp/`, preserving relative paths. Do not edit the originals yet. |
| verification | Confirm `bpdsl/old_mcp/` is byte-identical to the current `bpdsl/records/spec/mcp/` content. |

## Done condition

| item | done when |
|---|---|
| staging | `bpdsl/old_mcp/` exists with 12 files matching `bpdsl/records/spec/mcp/` structure (`overview.md`, `errors.md`, `schema.md`, `versioning.md`, `tools/*.md` × 8). |
| originals untouched | `bpdsl/records/spec/mcp/` content is unchanged from before this task. |

## Verification

- Diff `bpdsl/old_mcp/` against `bpdsl/records/spec/mcp/` — must be identical at this point (copy, not yet migrated).
- No files under `bpdsl/old_mcp/` are edited after creation (read-only staging area).

## Evidence

- Copied all 12 files from `bpdsl/records/spec/mcp/` to `bpdsl/old_mcp/`, preserving `tools/` subdirectory structure.
- `diff -rq bpdsl/records/spec/mcp bpdsl/old_mcp` → no output, exit 0 (byte-identical).
- `bpdsl/records/spec/mcp/` originals untouched (copy operation only, no edits).
