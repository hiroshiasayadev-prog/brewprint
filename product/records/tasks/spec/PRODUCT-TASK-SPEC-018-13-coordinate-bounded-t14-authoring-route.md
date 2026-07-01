# PRODUCT-TASK-SPEC-018-13: Coordinate bounded T14 authoring route

- **id**: PRODUCT-TASK-SPEC-018-13
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-12
- **outputs**:
  - PRODUCT-TASK-SPEC-018-13
  - PRODUCT-TASK-SPEC-018-14
  - PRODUCT-WORK-SPEC-018
  - PRODUCT-TASK-SPEC-018-07

## Goal

Persist one bounded authoring route for the `work_item_decomposition` Task-type addition and release T07 behind the T14 dependency.

## Work

- Preserve T12 as completed routing Evidence.
- Materialize T14 as one bounded `authoring` Task.
- Add T14 to W018 and the T07 dependency and review route.
- Change T07 from `blocked` to `not_started`.
- Keep T07 dependency-gated on T14 until authoring completes.
- Keep post-review routing conditional on the T07 verdict.

This Task must not:

- author ADR, Specification, or skill content;
- create another Work Item;
- create additional pre-review Tasks;
- perform review, synchronization, implementation, stage, or commit work.

## Done condition

- T14 exists with one authoring responsibility.
- W018 records the T13 to T14 to T07 route.
- T07 depends on T13 and T14.
- T07 is `not_started` and dependency-gated on T14.
- No unrelated Task or Work Item is created.

## Verification

- Confirm T13 and T14 IDs, types, dependencies, and outputs agree.
- Confirm W018 lists T13 and T14.
- Confirm T07 includes T13 and T14 and is no longer blocked.
- Confirm no additional pre-review Task or Work Item exists.
- Confirm stage and commit were not performed.

## Evidence

- T12 identified the bounded Task-type, ADR, Specification, and workflow-support changes.
- User direction removed all reciprocal-metadata migration and follow-up Work Item scope.
- W018 and T07 were aligned to the direct T14 to T07 route.
- T07 was released as `not_started` and remained gated by its T14 dependency.
- No review, synchronization, implementation, stage, or commit work was performed.
