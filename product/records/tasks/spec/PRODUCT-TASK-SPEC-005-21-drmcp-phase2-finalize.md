# PRODUCT-TASK-SPEC-005-21: DRMCP Phase 2 — finalize and close

- **id**: PRODUCT-TASK-SPEC-005-21
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-20
- **outputs**:
  - `drmcp/records/spec/design-records-mcp/` (must-fix corrections from -20 applied)
  - `product/records/spec/concepts/` (must-fix corrections from -20 applied)
  - `product/records/work-items/spec/PRODUCT-WORK-SPEC-005.md` (Phase 2 evidence entry)

## Goal

Apply any must-fix corrections from the Phase 2 Opus review, confirm validator cleanliness, and update WORK-005 Evidence to mark Phase 2 complete. Also confirm whether PRODUCT-WORK-SPEC-005 itself can now close (pending WORK-009 completion).

## Work

| area | required work |
|---|---|
| corrections | Apply all must-fix findings from PRODUCT-TASK-SPEC-005-20. |
| validator | `validate_spec.py product/records/spec/concepts/ --strict` and `drmcp/records/spec/design-records-mcp/ --strict` both exit 0. |
| WORK-005 evidence | Add Phase 2 relocation evidence entry, mirroring the format of the Phase 1 DRMCP batch entry. |
| WORK-005 status | If PRODUCT-WORK-SPEC-009 is also complete, evaluate closing WORK-005. If WORK-009 is not yet done, leave WORK-005 `in_progress`. |

## Done condition

| item | done when |
|---|---|
| corrections applied | All must-fix findings from -20 are resolved. |
| validator clean | Both namespaces pass `--strict`. |
| WORK-005 updated | Evidence records Phase 2 relocation completion. |

## Verification

- Confirm no deferred relocation notes remain in `drmcp/records/spec/design-records-mcp/` (all 4 candidates resolved in Phase 2).
- Confirm `validate_spec.py product/records/spec/concepts/ --strict` exits 0.
- Confirm `validate_spec.py drmcp/records/spec/design-records-mcp/ --strict` exits 0.

## Evidence

- 1 correction applied from -20 review: ADR discovery pattern updated in `record-discovery-paths.md` and noted in `schema/discovery.md`.
- `validate_spec.py product/records/spec/concepts/ drmcp/records/spec/design-records-mcp/ --strict` → `[strict] All 62 file(s) OK.`
- `grep -rl "Deferred relocation note" drmcp/records/spec/design-records-mcp/` → 0 hits.
- WORK-005 Evidence updated with Phase 2 relocation entry.
- PRODUCT-WORK-SPEC-005 `status` set to `done` — all tasks (-01 .. -21) complete, WORK-009 complete.
