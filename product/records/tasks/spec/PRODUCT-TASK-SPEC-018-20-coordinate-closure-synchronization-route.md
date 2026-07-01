# PRODUCT-TASK-SPEC-018-20: Coordinate closure synchronization route

- **id**: PRODUCT-TASK-SPEC-018-20
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-19
- **outputs**:
  - PRODUCT-TASK-SPEC-018-20
  - PRODUCT-TASK-SPEC-018-08
  - PRODUCT-WORK-SPEC-018

## Goal

Materialize the accepted W018 closure-synchronization owner after independent finding closure.

## Work

- Create T08 as the only closure-synchronization Task.
- Set T08 to depend on T19 and T20.
- Add T08 and T20 to the W018 Task graph.
- Route the accepted T19 result to T08.
- Preserve T07 and T19 verdict Evidence unchanged.

This Task must not:

- perform closure synchronization;
- change a review verdict or finding disposition;
- author or correct canonical content;
- change another completed Task;
- start production implementation;
- stage or commit changes.

## Done condition

- T08 exists with one synchronization responsibility.
- T08 depends on the accepted finding-closure route.
- W018 lists T08 and T20.
- The next owner is unambiguous.
- No canonical or lifecycle closure was performed by T20.

## Verification

- Confirm T08 and T20 use the W018 Task sequence.
- Confirm both Tasks reference W018.
- Confirm T08 depends on T19 and T20.
- Confirm W018 records the T19 to T20 to T08 route.
- Confirm no reviewed artifact or verdict changed.
- Confirm stage and commit were not performed.

## Evidence

- T19 returned `PASS`.
- F-BLK-01 is `CLOSED`.
- F-MAJ-01 is `CLOSED`.
- T20 materialized T08 as the exact closure-synchronization owner.
- T08 is the only remaining W018 workflow owner.
- No canonical authoring, correction, review, synchronization, implementation, stage, or commit work was performed by T20.
