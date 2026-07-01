# PRODUCT-TASK-SPEC-018-15: Coordinate T07 finding repair route

- **id**: PRODUCT-TASK-SPEC-018-15
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-07
- **outputs**:
  - PRODUCT-TASK-SPEC-018-15
  - PRODUCT-TASK-SPEC-018-16
  - PRODUCT-TASK-SPEC-018-17
  - PRODUCT-TASK-SPEC-018-18
  - PRODUCT-TASK-SPEC-018-19
  - PRODUCT-WORK-SPEC-018

## Goal

Materialize the exact repair and closure-review route for T07 findings F-BLK-01 and F-MAJ-01.

## Work

- Create T16 as the new decision owner for F-BLK-01.
- Create T17 as the bounded authoring owner after T16.
- Create T18 as the correction owner for F-MAJ-01.
- Create T19 as the independent finding-closure review owner.
- Set T17 to depend on T16.
- Set T19 to depend on T17 and T18.
- Keep T08 reserved until T19 closes every required finding.
- Update W018 with the exact post-review route.

This Task must not:

- decide F-BLK-01;
- author ADR, Specification, skill, or instruction content;
- repair either finding;
- perform finding-closure review;
- materialize T08;
- stage or commit changes.

## Done condition

- T16 through T19 exist with one responsibility each.
- Dependencies serialize decision, authoring, correction, and closure review correctly.
- W018 records the exact finding route.
- T08 remains unmaterialized.

## Verification

- Confirm every new Task references W018.
- Confirm W018 lists T15 through T19.
- Confirm T17 depends on T16.
- Confirm T19 depends on T17 and T18.
- Confirm no correction, review, or closure work was performed by T15.
- Confirm stage and commit were not performed.

## Evidence

- T07 returned `NEEDS REVISION` with F-BLK-01 and F-MAJ-01.
- F-BLK-01 requires a new decision before authoring.
- F-MAJ-01 requires a bounded correction with no new judgment.
- T16 through T19 were materialized as the exact finding route.
- T08 remains reserved.
