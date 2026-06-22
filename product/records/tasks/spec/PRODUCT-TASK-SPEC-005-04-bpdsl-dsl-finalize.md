# PRODUCT-TASK-SPEC-005-04: BPDSL — Apply review corrections and finalize DSL migration

- **id**: PRODUCT-TASK-SPEC-005-04
- **status**: done
- **date**: 2026-06-16
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-03
- **outputs**:
  - `bpdsl/records/spec/dsl/**` (corrections applied)
  - `bpdsl/old/` removed

## Goal

Apply must-fix corrections from the Opus review, remove the `bpdsl/old/` staging directory, and confirm the full BPDSL DSL migration is clean.

## Work

| area | required work |
|---|---|
| corrections | Apply all must-fix findings from PRODUCT-TASK-SPEC-005-03. |
| old/ removal | Delete `bpdsl/old/` after confirming `dsl/` is complete. |
| validator | Run `validate_spec.py bpdsl/records/spec/dsl/ --strict --no-color` (plus `bpdsl/records/spec/overview.md`) and confirm 0 errors. `mcp/` and `views/` are not yet migrated (PRODUCT-TASK-SPEC-005-05 / -06) and are out of scope for this task's validator pass — running the validator across all of `bpdsl/records/spec/` would surface their pre-existing, expected failures. |
| WORK-005 evidence | Add BPDSL DSL migration evidence entry to PRODUCT-WORK-SPEC-005. |

## Done condition

| item | done when |
|---|---|
| corrections applied | All must-fix findings from 005-03 are resolved. |
| bpdsl/old/ removed | Staging directory is deleted. |
| validator clean | `--strict` exits 0 on `bpdsl/records/spec/dsl/` and `bpdsl/records/spec/overview.md`. |
| WORK-005 updated | Evidence section of WORK-005 reflects completion of BPDSL DSL batch. |

## Verification

- Confirm `bpdsl/old/` no longer exists.
- Confirm no broken `[]()` links to `bpdsl/old/` remain in any spec file.
- Confirm `validate_spec.py bpdsl/records/spec/dsl/ --strict` and `validate_spec.py bpdsl/records/spec/overview.md --strict` both exit 0.

## Evidence

- No must-fix findings remained from 005-03 — the two broken section anchors and three dropped-content judgment calls were already corrected during the 005-03 review pass itself (see 005-03 Evidence disposition table).
- `bpdsl/old/` deleted. Confirmed `ls bpdsl/` shows only `records/`.
- Grepped `*.md` repo-wide for `bpdsl/old` — only the 4 task records under `product/records/tasks/spec/` reference it (as historical evidence text); no spec file links into it.
- `python product/src/tools/validate_spec.py bpdsl/records/spec/dsl/ --strict --no-color` → `[strict] All 16 file(s) OK.`
- `python product/src/tools/validate_spec.py bpdsl/records/spec/overview.md --strict --no-color` → `[strict] All 1 file(s) OK.`
- Full-tree `validate_spec.py bpdsl/records/spec/ --strict` still fails (57 errors, all in `mcp/` and `views/` — pending PRODUCT-TASK-SPEC-005-05/-06). This is expected and out of scope for this task; the task's verification step was corrected accordingly (see `## Work` / `## Verification` above) rather than silently run against the wrong scope.
- WORK-005 evidence updated to record completion of the BPDSL DSL migration batch (T01–T04).
