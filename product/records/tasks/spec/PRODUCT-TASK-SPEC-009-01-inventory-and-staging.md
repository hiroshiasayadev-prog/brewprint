# PRODUCT-TASK-SPEC-009-01: Inventory — current diagnostics and file staging

- **id**: PRODUCT-TASK-SPEC-009-01
- **status**: done
- **date**: 2026-06-22
- **work_item**: PRODUCT-WORK-SPEC-009
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - `product/records/old/` (staging directory, 9 pre-migration files)
  - Inventory diagnostics (inline in Evidence)

## Goal

Record the current format-validation state of all 9 target files, then stage them to `product/records/old/` so the Opus reviewer (PRODUCT-TASK-SPEC-009-04) can compare before and after.

## Work

| area | required work |
|---|---|
| inventory | Run `python product/src/tools/validate_spec.py product/records/spec/concepts/project-artifact-model/index.md product/records/spec/concepts/traceability/ product/records/spec/concepts/namespace-model/index.md` (default mode, no `--strict`). Record full output in Evidence. |
| staging | Copy all 9 target files to `product/records/old/`, preserving the relative directory structure. Do not edit the staged copies. |

Staging destination structure:

```
product/records/old/
  project-artifact-model/
    index.md
  traceability/
    index.md
    artifact-refs.md
    coverage-mapping.md
    metadata-schema.md
    out-of-scope.md
    resolve-and-validation.md
    semantic-ref.md
  namespace-model/
    index.md
```

## Done condition

| item | done when |
|---|---|
| inventory run | `validate_spec.py` default-mode output is recorded in Evidence for all 9 files. |
| staging complete | `product/records/old/` contains exactly 9 files in the structure above, unmodified. |
| no edits to staged copies | Staged copies match the pre-migration committed state of the source files. |

## Verification

- Do not run `--strict` — the pre-migration files are expected to have format warnings.
- Confirm staged files are byte-for-byte copies (no re-encoding, no line-ending changes).

## Evidence

### Inventory — validate_spec.py default mode (2026-06-22)

All 9 files: warnings only, 0 errors. Uniform pattern across all files:

| diagnostic | count |
|---|---|
| `YAML_FRONT_MATTER` | 1 per file |
| `FM_HIDDEN_REFS` (depends_on) | 1 per file |
| `FM_HIDDEN_REFS` (semantic_refs) | 1 per file |
| `FM_HIDDEN_REFS` (sections) | most files |
| `H1_FORMAT` | 1 per file |

Summary: `[inventory] 9 file(s) with issues: 44 warning(s)`

### Staging (2026-06-22)

`product/records/old/` created with subdirectory structure mirroring source:

```
product/records/old/
  namespace-model/index.md
  project-artifact-model/index.md
  traceability/artifact-refs.md
  traceability/coverage-mapping.md
  traceability/index.md
  traceability/metadata-schema.md
  traceability/out-of-scope.md
  traceability/resolve-and-validation.md
  traceability/semantic-ref.md
```

9 files staged, unmodified.
