# PRODUCT-TASK-SPEC-005-08: BPDSL — Apply review corrections and finalize mcp/ migration

- **id**: PRODUCT-TASK-SPEC-005-08
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-07
- **outputs**:
  - `bpdsl/records/spec/mcp/**` (corrections applied)
  - `bpdsl/old_mcp/` removed

## Goal

Apply must-fix corrections from the Opus review, remove the `bpdsl/old_mcp/` staging directory, and confirm the mcp/ migration is clean.

## Work

| area | required work |
|---|---|
| corrections | Apply all must-fix findings from PRODUCT-TASK-SPEC-005-07. |
| old_mcp/ removal | Delete `bpdsl/old_mcp/` after confirming `mcp/` is complete. |
| validator | Run `validate_spec.py bpdsl/records/spec/mcp/ --strict --no-color` and confirm 0 errors on all mcp/ files. |
| WORK-005 evidence | Add mcp/ migration evidence entry to PRODUCT-WORK-SPEC-005. |

## Done condition

| item | done when |
|---|---|
| corrections applied | All must-fix findings from 005-07 are resolved. |
| bpdsl/old_mcp/ removed | Staging directory is deleted. |
| validator clean | `--strict` exits 0 on all `bpdsl/records/spec/mcp/` files. |
| WORK-005 updated | Evidence section of WORK-005 reflects completion of the mcp/ batch. |

## Verification

- Confirm `bpdsl/old_mcp/` no longer exists.
- Confirm no broken `[]()` links to `bpdsl/old_mcp/` remain in any spec file.
- Confirm `validate_spec.py bpdsl/records/spec/mcp/ --strict` exits 0.

## Evidence

- All must-fix-equivalent corrections from 005-07 were already applied during that task's review pass (see 005-07 Evidence disposition table) — nothing further to apply here.
- `bpdsl/old_mcp/` deleted. `ls bpdsl/` shows only `records/`.
- Grepped `*.md` repo-wide for `bpdsl/old_mcp` — only the 4 task records under `product/records/tasks/spec/` reference it (historical evidence text); no spec file links into it.
- `python product/src/tools/validate_spec.py bpdsl/records/spec/mcp/ --strict --no-color` → `[strict] All 12 file(s) OK.`
- WORK-005 evidence updated to record completion of the BPDSL mcp/ migration batch (T05–T08).
