# PRODUCT-TASK-SPEC-018-18: Correct decomposition companion activation

- **id**: PRODUCT-TASK-SPEC-018-18
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: correction
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-15
- **outputs**:
  - PRODUCT-TASK-SPEC-018-18
  - prompt_chappy.md

## Goal

Repair T07 finding F-MAJ-01 in the active design-convergence instruction.

## Work

- Add the exact `work-item-decomposition.md` companion pointer.
- State that `coordination` owns Task graph change.
- State that `work_item_decomposition` owns decided parent-to-child Work Item creation or split.
- Preserve every unrelated instruction and workflow rule.

This Task must not:

- change the accepted workflow design;
- edit successor skill files;
- repair F-BLK-01;
- close F-MAJ-01;
- perform lifecycle synchronization;
- stage or commit changes.

## Done condition

- The active instruction names `work-item-decomposition.md` in the phase companion list.
- The ownership summary distinguishes coordination from Work Item decomposition.
- No unrelated instruction content changes.

## Verification

- Confirm the companion path exactly matches the successor skill file.
- Confirm the ownership summary matches ADR-010 and `SKILL.md`.
- Confirm no other `prompt_chappy.md` section changed.
- Confirm stage and commit were not performed.

## Evidence

- Added `work-item-decomposition.md` to the mandatory companion list.
- Added the distinct `coordination` and `work_item_decomposition` ownership summary.
- F-MAJ-01 is repaired pending independent finding-closure review.
