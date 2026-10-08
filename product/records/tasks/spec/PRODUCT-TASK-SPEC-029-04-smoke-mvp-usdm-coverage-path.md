# PRODUCT-TASK-SPEC-029-04: Smoke MVP USDM coverage path

- **id**: PRODUCT-TASK-SPEC-029-04
- **status**: done
- **date**: 2026-07-09
- **work_item**: PRODUCT-WORK-SPEC-029
- **task_type**: verification
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-029-03
- **outputs**:
  - PRODUCT-TASK-SPEC-029-04
  - PRODUCT-WORK-SPEC-029

## Goal

Verify the standalone USDM tools against a minimal non-empty smoke repository.

## Work

- Generate a temporary repository using inline Python from PowerShell.
- Create one USDM `requirement` record in the temporary repository.
- Create one implementation Specification that declares `usdm_covers` for that requirement.
- Run the positive path for `validate_usdm`, `check_usdm_coverage`, and `usdm_covered_by`.
- Run a negative uncovered coverage path.
- Record the smoke result without creating persistent fixture records.

## Done condition

- The inline smoke creates a non-empty temporary USDM repository.
- `validate_usdm` succeeds on the temporary repository.
- `check_usdm_coverage` succeeds when the requirement is covered.
- `usdm_covered_by` returns a covering Specification for the covered requirement.
- `check_usdm_coverage` fails as expected after removing coverage.
- The observed smoke output is recorded.

## Verification

The inline Python smoke executed from PowerShell and printed:

```text
USDM smoke PASS
```

Verified paths:

| check | expected result | observed result |
|---|---|---|
| `validate_usdm` over temporary repo | exit 0 and `ok: true` | PASS |
| `check_usdm_coverage` with `usdm_covers` present | exit 0 and `ok: true` | PASS |
| `usdm_covered_by` for covered requirement | exit 0, `ok: true`, and non-empty `covered_by` | PASS |
| `check_usdm_coverage` after removing coverage | non-zero exit, `ok: false`, and requirement listed in `uncovered` | PASS |

## Evidence

- The smoke used a temporary repository instead of persistent app records.
- The temporary repository contained `demo/records/usdm/identity.md` and `demo/records/spec/identity/index.md`.
- The smoke covered `usdm:demo.identity#R001` from `spec:demo.identity`.
- The negative case removed the `usdm_covers` entry and confirmed uncovered detection.
- The observed terminal output was `USDM smoke PASS`.
- No fixture files were intentionally persisted by this Task.
