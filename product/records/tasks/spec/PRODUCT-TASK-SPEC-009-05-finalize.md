# PRODUCT-TASK-SPEC-009-05: Apply review corrections and finalize migration

- **id**: PRODUCT-TASK-SPEC-009-05
- **status**: done
- **date**: 2026-06-22
- **work_item**: PRODUCT-WORK-SPEC-009
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-009-04
- **outputs**:
  - All 9 migrated spec files (corrections applied)
  - `product/records/old/` removed
  - PRODUCT-WORK-SPEC-009 Evidence updated

## Goal

Apply must-fix corrections from the Opus review, remove the staging directory, confirm the migration is clean, and close the work item.

## Work

| area | required work |
|---|---|
| corrections | Apply all must-fix findings from PRODUCT-TASK-SPEC-009-04. |
| old/ removal | Delete `product/records/old/` after confirming all 9 migrated files are correct. |
| final validation | Run `python product/src/tools/validate_spec.py product/records/spec/concepts/project-artifact-model/index.md product/records/spec/concepts/traceability/ product/records/spec/concepts/namespace-model/index.md --strict --no-color`. Confirm exit 0. |
| WORK-009 evidence | Add Evidence section to PRODUCT-WORK-SPEC-009 recording: files migrated, validator result, review verdict and disposition of findings. |
| WORK-009 status | Update PRODUCT-WORK-SPEC-009 status to `done`. |

## Done condition

| item | done when |
|---|---|
| corrections applied | All must-fix findings from PRODUCT-TASK-SPEC-009-04 resolved. |
| old/ removed | `product/records/old/` no longer exists. |
| final validator clean | `--strict` exits 0 on all 9 files. |
| WORK-009 updated | Evidence and status updated in PRODUCT-WORK-SPEC-009. |

## Verification

- Confirm `product/records/old/` no longer exists.
- Confirm `validate_spec.py --strict` exits 0 on all 9 files.
- Confirm no broken links into `product/records/old/` remain in any spec file.

## Evidence

### Finalization (2026-06-22)

- All must-fix findings from T04 (M1 + D1/D2/D3) applied
- `product/records/old/` confirmed no longer referenced in any spec file; directory removed
- Final validation: `[strict]  All 18 file(s) OK.`
- PRODUCT-WORK-SPEC-009 Evidence section added; status updated to `done`
