# PRODUCT-TASK-SPEC-019-22: Coordinate early checklist release graph

- **id**: PRODUCT-TASK-SPEC-019-22
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-21
- **outputs**:
  - PRODUCT-TASK-SPEC-019-22
  - PRODUCT-TASK-SPEC-019-19
  - PRODUCT-WORK-SPEC-019
  - PRODUCT-WORK-SPEC-020
  - PRODUCT-TASK-SPEC-020-01

## Goal

Apply the decided early W020 release boundary to the cross-Work-Item graph.

## Work

- Update W019 coarse routing for W020 early work.
- Update incomplete T19 so W019 closure does not own W020 start.
- Preserve T19 as the W019 closure synchronization owner.
- Preserve W021 release after W019 closure and accepted W020 review.
- Set W020 to `in_progress`.
- Create PRODUCT-TASK-SPEC-020-01 as the W020 internal graph owner.

This Task must not:

- rewrite T09 or T21;
- decide checklist content or storage;
- create checklist artifacts;
- create W020 child Tasks other than T01;
- modify ADRs or Specifications;
- perform implementation, review, correction, synchronization, stage, or commit work.

## Done condition

- W019 coarse routing reflects T21.
- T19 no longer treats W020 start as closure-owned.
- W021 retains its accepted release gate.
- W020 is `in_progress` and lists T01.
- T01 owns W020 internal Task materialization.

## Verification

- Confirm W019 lists T20 through T22.
- Confirm T19 retains one synchronization outcome.
- Confirm W020 lists only materialized child Tasks.
- Confirm no checklist content changed.
- Confirm W021 is unchanged.

## Evidence

- T21 decided the revised current release order.
- T09 remains the completed original release-order checkpoint.
- W020 internal graph ownership is delegated to PRODUCT-TASK-SPEC-020-01.
- Result: `PASS`.
