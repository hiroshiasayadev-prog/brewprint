# PRODUCT-TASK-SPEC-005-09: BPDSL — views/ directory setup and file staging

- **id**: PRODUCT-TASK-SPEC-005-09
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-08
- **outputs**:
  - `bpdsl/old_views/` (staging directory — byte-identical snapshot of pre-migration views/ spec files)

## Goal

Snapshot the current 7 view spec files under `bpdsl/records/spec/views/` into `bpdsl/old_views/` before any migration edits. `views/overview.md` does not exist yet — it is authored new in 005-10, not staged here. The two CSS assets (`wireframe.css`, `wireframe.preview.css`) are not spec files and are out of scope for staging. (Suffix `_views` distinguishes this staging directory from a future `bpdsl/old_drmcp/` staged in a separate DRMCP migration session.)

## Work

| area | required work |
|---|---|
| staging | Copy the 7 existing `.md` files under `bpdsl/records/spec/views/` (`model-file.md`, `api-table.md`, `state-diagram.md`, `er.md`, `dag.md`, `sequence-diagram.md`, `wireframe.md`) to `bpdsl/old_views/`. Do not edit the originals yet. |
| verification | Confirm `bpdsl/old_views/` is byte-identical to the current 7 files. |

## Done condition

| item | done when |
|---|---|
| staging | `bpdsl/old_views/` exists with the 7 staged files. |
| originals untouched | `bpdsl/records/spec/views/` content is unchanged from before this task. |

## Verification

- Diff `bpdsl/old_views/` against the source files — must be identical at this point.
- No files under `bpdsl/old_views/` are edited after creation (read-only staging area).

## Evidence

- Confirmed `bpdsl/records/spec/views/` contents: 7 `.md` files (`model-file.md`, `api-table.md`, `state-diagram.md`, `er.md`, `dag.md`, `sequence-diagram.md`, `wireframe.md`) + 2 `.css` assets (`wireframe.css`, `wireframe.preview.css`) — CSS excluded from staging per task scope.
- Copied all 7 `.md` files to `bpdsl/old_views/`.
- `diff -q` on each of the 7 pairs → no output, exit 0 (byte-identical).
- `bpdsl/records/spec/views/` originals untouched (copy operation only, no edits).
