# PRODUCT-TASK-SPEC-005-12: BPDSL — Apply review corrections and finalize views/ migration

- **id**: PRODUCT-TASK-SPEC-005-12
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-11
- **outputs**:
  - `bpdsl/records/spec/views/**` (corrections applied)
  - `bpdsl/old_views/` removed

## Goal

Apply must-fix corrections from the Opus review, remove the `bpdsl/old_views/` staging directory, and confirm the views/ migration — and with it, the full BPDSL namespace spec-format migration (dsl/ + mcp/ + views/) — is clean.

## Work

| area | required work |
|---|---|
| corrections | Apply all must-fix findings from PRODUCT-TASK-SPEC-005-11. |
| old_views/ removal | Delete `bpdsl/old_views/` after confirming `views/` is complete. |
| validator | Run `validate_spec.py bpdsl/records/spec/views/ --strict --no-color` and confirm 0 errors on all views/ files. |
| full-namespace validator | Run `validate_spec.py bpdsl/records/spec/ --strict --no-color` and confirm 0 errors across dsl/ + mcp/ + views/ (DRMCP migration is tracked separately and is not part of this check). |
| WORK-005 evidence | Add views/ migration evidence entry to PRODUCT-WORK-SPEC-005, noting BPDSL namespace spec-format migration complete. |

## Done condition

| item | done when |
|---|---|
| corrections applied | All must-fix findings from 005-11 are resolved. |
| bpdsl/old_views/ removed | Staging directory is deleted. |
| validator clean (views/) | `--strict` exits 0 on all `bpdsl/records/spec/views/` files. |
| validator clean (full namespace) | `--strict` exits 0 on all `bpdsl/records/spec/` files. |
| WORK-005 updated | Evidence section of WORK-005 reflects completion of the views/ batch and the full BPDSL namespace migration. |

## Verification

- Confirm `bpdsl/old_views/` no longer exists.
- Confirm no broken `[]()` links to `bpdsl/old_views/` remain in any spec file.
- Confirm `validate_spec.py bpdsl/records/spec/views/ --strict` exits 0.
- Confirm `validate_spec.py bpdsl/records/spec/ --strict` exits 0 (full BPDSL namespace, excluding DRMCP which is tracked separately).

## Evidence

- T11 must-fix correction already applied in T11 pass (`er.md` `## Rules` generation-source sentence added); two false-alarm findings confirmed via direct file read and discarded.
- `bpdsl/old_views/` removed. Confirmed no `old_views` links remain in any `bpdsl/records/spec/` file (grep clean).
- `validate_spec.py bpdsl/records/spec/views/ --strict --no-color` → `[strict] All 8 file(s) OK.`
- `validate_spec.py bpdsl/records/spec/ --strict --no-color` → `[strict] All 37 file(s) OK.` (dsl/ 16 + mcp/ 12 + views/ 8 + overview.md 1 = 37)
- WORK-005 evidence section updated with views/ batch entry noting full BPDSL namespace migration complete.
