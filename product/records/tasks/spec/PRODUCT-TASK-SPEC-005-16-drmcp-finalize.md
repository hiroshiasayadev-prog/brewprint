# PRODUCT-TASK-SPEC-005-16: DRMCP — Apply review corrections and finalize DRMCP migration

- **id**: PRODUCT-TASK-SPEC-005-16
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-15
- **outputs**:
  - `drmcp/records/spec/design-records-mcp/**` (corrections applied)
  - `drmcp/old/` removed

## Goal

Apply must-fix corrections from the Opus review, remove the `drmcp/old/` staging directory, and confirm the full DRMCP spec migration is clean.

## Work

| area | required work |
|---|---|
| corrections | Apply all must-fix findings from PRODUCT-TASK-SPEC-005-15. |
| old/ removal | Delete `drmcp/old/` after confirming `design-records-mcp/` is complete. |
| validator | Run `validate_spec.py drmcp/records/spec/design-records-mcp/ --strict --no-color` and confirm 0 errors. |
| WORK-005 evidence | Add DRMCP spec migration evidence entry to PRODUCT-WORK-SPEC-005, mirroring the BPDSL DSL batch entry format. Carry forward the PRODUCT-TASK-SPEC-005-15 deferred relocation candidates list (PRODUCT-INV-SPEC-004) into this entry so the open ownership question stays visible at the work-item level, not just buried in a closed task. |

## Done condition

| item | done when |
|---|---|
| corrections applied | All must-fix findings from 005-15 are resolved. |
| drmcp/old/ removed | Staging directory is deleted. |
| validator clean | `--strict` exits 0 on `drmcp/records/spec/design-records-mcp/`. |
| WORK-005 updated | Evidence section of WORK-005 reflects completion of the DRMCP spec migration batch, including the carried-forward deferred relocation candidates. |

## Verification

- Confirm `drmcp/old/` no longer exists.
- Confirm no broken `[]()` links to `drmcp/old/` remain in any spec file.
- Confirm `validate_spec.py drmcp/records/spec/design-records-mcp/ --strict` exits 0.

## Evidence

- 0 must-fix corrections (review PASS with no must-fix items).
- `drmcp/old/` removed.
- `validate_spec.py drmcp/records/spec/design-records-mcp/ --strict --no-color` → `[strict] All 30 file(s) OK.`
- WORK-005 Evidence updated with DRMCP batch entry and deferred relocation candidates table (4 files, Phase 2 pending).
