# PRODUCT-TASK-SPEC-009-03: Strict validation of migrated spec files

- **id**: PRODUCT-TASK-SPEC-009-03
- **status**: done
- **date**: 2026-06-22
- **work_item**: PRODUCT-WORK-SPEC-009
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-009-02
- **outputs**:
  - Validation output (inline in Evidence)

## Goal

Confirm all 9 migrated files pass `--strict` validation before handing off to the Opus reviewer. Fix any errors found before closing this task.

## Work

| area | required work |
|---|---|
| validate | Run `python product/src/tools/validate_spec.py product/records/spec/concepts/project-artifact-model/index.md product/records/spec/concepts/traceability/ product/records/spec/concepts/namespace-model/index.md --strict --no-color`. |
| fix errors | Fix any errors reported. Re-run until exit 0. Record final output in Evidence. |

## Done condition

| item | done when |
|---|---|
| exit 0 | `validate_spec.py --strict` exits 0 on all 9 files. |
| output recorded | Full validator output recorded in Evidence. |

## Verification

- Target: 0 errors, 0 warnings in `--strict` mode across all 9 files.
- Do not close this task if warnings remain — treat all `--strict` diagnostics as errors for migration completeness.

## Evidence

### Validation run (2026-06-22)

Command:
```
python product/src/tools/validate_spec.py product/records/spec/concepts/project-artifact-model/ product/records/spec/concepts/traceability/ product/records/spec/concepts/namespace-model/ --strict --no-color
```

Result: `[strict]  All 18 file(s) OK.` — exit 0.

Note: command targets full directories (not index.md only) to cover all 18 output files including the 11 new split files.
