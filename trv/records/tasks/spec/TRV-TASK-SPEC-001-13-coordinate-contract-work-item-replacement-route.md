# TRV-TASK-SPEC-001-13: Coordinate contract Work Item replacement route

- **id**: TRV-TASK-SPEC-001-13
- **status**: done
- **date**: 2026-07-03
- **work_item**: TRV-WORK-SPEC-001
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-003-04
- **outputs**:
  - TRV-WORK-SPEC-001
  - TRV-WORK-SPEC-003
  - TRV-TASK-SPEC-001-09
  - TRV-TASK-SPEC-001-13
  - TRV-TASK-SPEC-003-05
  - TRV-TASK-SPEC-003-06
  - TRV-TASK-SPEC-003-07
  - TRV-TASK-SPEC-003-08

## Goal

Retire the incomplete W003 execution route and materialize one Work Item decomposition owner for its replacement.

## Work

- Preserve completed W003 T01 through T04 as historical workflow Evidence.
- Mark W003 as blocked and operationally retired because its graph does not decide architecture-derived contract Specification placement.
- Block W003 T05 through T07 so the incomplete authoring, review, and closure route cannot execute.
- Block parent T09 because W003 will not reach `done` under its original completion contract.
- Create W003 T08 as the `work_item_decomposition` owner for the replacement Work Item.
- Keep TRV-ADR-SPEC-003 through TRV-ADR-SPEC-005 unchanged as historical accepted inputs.
- Leave replacement Work Item creation to T08.

This Task must not create the replacement Work Item, author ADR or Specification content, issue a review verdict, synchronize closure, or change W002 architecture.

## Done condition

- W003 and its unfinished downstream Tasks cannot be executed as the active contract route.
- T09 no longer presents W003 as a completable parent execution gate.
- T08 exists with one Work Item decomposition responsibility.
- Completed W003 Tasks remain historical Evidence.
- No replacement deliverable is authored by this Task.

## Verification

- Confirm W003, T05, T06, T07, and T09 record the exact blocked replacement reason.
- Confirm T08 uses `work_item_decomposition` and belongs to W003.
- Confirm W003 lists T08.
- Confirm no completed decision, routing, or ADR-authoring outcome was rewritten.
- Confirm no child Work Item or canonical content was created.

## Evidence

- Post-T04 inspection showed that W003 and T05 were limited to external MCP, Task-input, caller, and compatibility topics.
- The graph omitted the prior decision on how architecture-derived contract Specifications should be partitioned and placed.
- The user directed operational closure of W003 and creation of a separate replacement Work Item.
- W003 was changed to `blocked` rather than `done` because its original Completion Conditions remain unsatisfied.
- T05 through T07 and parent T09 were blocked with the replacement reason.
- T08 was created as the sole Work Item decomposition owner.
- No Work Item, ADR, Specification, review, synchronization, implementation, stage, or commit work occurred.
- Result: `PASS`.
