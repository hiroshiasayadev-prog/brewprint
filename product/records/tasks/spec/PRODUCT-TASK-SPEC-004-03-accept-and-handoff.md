# PRODUCT-TASK-SPEC-004-03: Accept ownership boundary decision and produce migration handoff

- **id**: PRODUCT-TASK-SPEC-004-03
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-004
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-004-02
- **outputs**:
  - Accepted ownership boundary decision (closes PRODUCT-WORK-SPEC-004)
  - PRODUCT-WORK-SPEC-009 scope update (if the review recommended adding `namespace-model/index.md`)
  - PRODUCT-WORK-SPEC-005 dependency/evidence note recording the accepted relocation plan as a future follow-up

## Goal

Apply must-fix corrections from PRODUCT-TASK-SPEC-004-02, accept the final ownership boundary decision and relocation plan, and hand off concrete prerequisites to PRODUCT-WORK-SPEC-005 / PRODUCT-WORK-SPEC-009. Close PRODUCT-WORK-SPEC-004.

## Work

| area | required work |
|---|---|
| corrections | Apply all must-fix findings from PRODUCT-TASK-SPEC-004-02. |
| WORK-SPEC-009 scope update | If the accepted decision adds `namespace-model/index.md` as a relocation target, add it to PRODUCT-WORK-SPEC-009's `Impact Scope` (format-only migration, no relocation — consistent with that work item's existing boundary). |
| WORK-SPEC-005 note | Record the accepted relocation plan as a future follow-up against the sections PRODUCT-TASK-SPEC-005-15 flagged as deferred relocation candidates, so the loop closes instead of dead-ending. |
| WORK-SPEC-004 closure | Mark PRODUCT-WORK-SPEC-004 `status: done`, with the accepted decision and relocation plan in its Evidence section. |

## Done condition

| item | done when |
|---|---|
| corrections applied | All must-fix findings from 004-02 are resolved. |
| decision accepted | PRODUCT-WORK-SPEC-004 Evidence records the final accepted decision and relocation plan. |
| WORK-SPEC-009 updated | Scope reflects whether `namespace-model/index.md` was added, with reasoning. |
| WORK-SPEC-005 updated | Evidence or Dependencies notes the accepted relocation plan and where it will be executed. |
| WORK-SPEC-004 closed | `status: done`. |

## Verification

- Confirm PRODUCT-TASK-SPEC-005-15's deferred-relocation-candidate findings now have a concrete destination in the accepted plan.
- Confirm no actual file relocation happened in this task — execution belongs to PRODUCT-WORK-SPEC-005 (or PRODUCT-WORK-SPEC-009 for the format-only step), per this work item's own non-scope.

## Evidence

- No must-fix corrections from PRODUCT-TASK-SPEC-004-02.
- PRODUCT-WORK-SPEC-009 `Impact Scope` updated to 9 files: `namespace-model/index.md` added. `Boundary` and `Completion Condition` updated to reflect 9 files and to add drift-guard requirement for `resolve-and-validation.md` hybrid sections.
- PRODUCT-WORK-SPEC-005 Evidence updated with accepted relocation plan summary and Phase 2 relocation batch note (future tasks, TBD after PRODUCT-TASK-SPEC-005-16 completes).
- PRODUCT-WORK-SPEC-004 `status` set to `done`.
