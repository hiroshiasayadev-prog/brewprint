# PRODUCT-TASK-SPEC-021-05: Repair PRODUCT-to-app design route

- **id**: PRODUCT-TASK-SPEC-021-05
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-021-03
- **outputs**:
  - PRODUCT-WORK-SPEC-021
  - PRODUCT-TASK-SPEC-021-04
  - PRODUCT-TASK-SPEC-021-05
  - PRODUCT-TASK-SPEC-021-06
  - PRODUCT-TASK-SPEC-021-07

## Goal

Repair the W021 Task graph so PRODUCT work stops before app-local design and implementation.

## Work

- Preserve completed T02 and T03 as historical Evidence.
- Block the obsolete direct executor-ready route in T04.
- Block W021 from releasing implementation work.
- Create one successor decision owner for the PRODUCT-to-app ownership correction.
- Create one successor coordination owner for the accepted post-decision route.
- Preserve app namespace creation, child Work Item creation, and app-local design as downstream responsibilities.
- Keep implementation Tasks, executor prompts, and implementation review outside the repaired graph.

This Task must not:

- select an app namespace or top-level directory;
- change T02 decisions or PRODUCT-INV-SPEC-009 findings;
- create a child Work Item;
- author namespace profile or app-local Specification content;
- materialize implementation or executor Tasks;
- perform review, synchronization, stage, or commit work.

## Done condition

- W021 and T04 cannot release direct implementation work.
- T06 owns the unresolved PRODUCT-to-app boundary and app namespace decisions.
- T07 owns graph materialization after T06 reaches a terminal result.
- Completed T02 and T03 remain unchanged.
- No app-local deliverable or implementation Task is created.

## Verification

- Confirm W021 lists T05 through T07.
- Confirm T04 is blocked and names the replacement route.
- Confirm T06 is a `decision` Task with no selected outcome.
- Confirm T07 depends on T06 and owns only graph coordination.
- Confirm no completed Task was substantively rewritten.
- Confirm no namespace, child Work Item, source, test, or prompt was created.

## Evidence

- The user clarified that PRODUCT owns conceptual design only.
- App namespace bootstrap and app-local implementation Specifications must precede implementation planning.
- PRODUCT-INV-SPEC-009 found that no active standalone validator app namespace exists.
- The prior T04 contract incorrectly routed from PRODUCT Investigation directly to executor-ready implementation planning.
- W021 was set to `blocked`, and T04 was set to `blocked`.
- Created T06 for the PRODUCT-to-app ownership and bootstrap decision.
- Created T07 for post-decision graph materialization.
- T02 and T03 remain historical inputs and were not modified.
- No implementation, child Work Item, namespace authoring, review, synchronization, stage, or commit was performed.
