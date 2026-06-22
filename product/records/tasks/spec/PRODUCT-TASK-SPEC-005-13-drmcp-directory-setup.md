# PRODUCT-TASK-SPEC-005-13: DRMCP — directory setup and file staging

- **id**: PRODUCT-TASK-SPEC-005-13
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - `drmcp/old/` (staging directory for pre-migration flat files)
  - `drmcp/records/spec/design-records-mcp/schema/` (empty directory, ready for authoring)
  - `drmcp/records/spec/design-records-mcp/tools/` (empty directory, ready for authoring)

## Goal

Stage the 3 pre-migration DRMCP spec files for reference during migration, and create the `schema/` and `tools/` sub-directories under `design-records-mcp/` that the authoring task will populate. Unlike the BPDSL batch, DRMCP has only one domain (`design-records-mcp/`) — no peer-domain restructuring of a root `overview.md` is needed; `design-records-mcp/overview.md` itself is re-authored in place as the root entry point (`parent: -`) in PRODUCT-TASK-SPEC-005-14.

## Work

| area | required work |
|---|---|
| staging | Move all 3 existing files (`overview.md`, `schema.md`, `tools.md`) from `drmcp/records/spec/design-records-mcp/` to `drmcp/old/`. Do not edit them there. |
| directory structure | Create empty `drmcp/records/spec/design-records-mcp/schema/` and `drmcp/records/spec/design-records-mcp/tools/` directories. |

## Done condition

| item | done when |
|---|---|
| staging | `drmcp/old/` exists and contains the 3 pre-migration files (`overview.md`, `schema.md`, `tools.md`). |
| schema/ exists | `drmcp/records/spec/design-records-mcp/schema/` directory exists. |
| tools/ exists | `drmcp/records/spec/design-records-mcp/tools/` directory exists. |

## Verification

- No files under `drmcp/old/` are edited (read-only staging area).
- `drmcp/records/spec/design-records-mcp/` contains no files at this point — only the two new empty sub-directories.
- No validator run in this task — no target-format files exist yet to validate.

## Evidence

| item | detail |
|---|---|
| `drmcp/old/` created | 3 pre-migration files staged: `overview.md` (220 lines), `schema.md` (834 lines), `tools.md` (1944 lines). |
| `drmcp/records/spec/design-records-mcp/` | Empty after staging — contains only the two new empty subdirectories. |
| `drmcp/records/spec/design-records-mcp/schema/` created | Empty directory, ready for authoring. |
| `drmcp/records/spec/design-records-mcp/tools/` created | Empty directory, ready for authoring. |
| no validator run | No target-format files exist yet to validate. |
