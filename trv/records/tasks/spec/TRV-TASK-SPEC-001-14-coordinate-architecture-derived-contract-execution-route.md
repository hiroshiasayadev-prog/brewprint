# TRV-TASK-SPEC-001-14: Coordinate architecture-derived contract execution route

- **id**: TRV-TASK-SPEC-001-14
- **status**: done
- **date**: 2026-07-03
- **work_item**: TRV-WORK-SPEC-001
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-001-13
  - TRV-TASK-SPEC-003-08
- **outputs**:
  - TRV-WORK-SPEC-001
  - TRV-WORK-SPEC-005
  - TRV-WORK-SPEC-004
  - TRV-TASK-SPEC-001-10
  - TRV-TASK-SPEC-001-14
  - TRV-TASK-SPEC-001-15
  - TRV-TASK-SPEC-005-01
  - TRV-TASK-SPEC-005-02
  - TRV-TASK-SPEC-005-03
  - TRV-TASK-SPEC-004-01

## Goal

Materialize the replacement W005 execution relation and its placement-first initial Task graph.

## Work

- Confirm TRV-WORK-SPEC-005 exists before creating an execution relation.
- Create parent T15 as one `work_item_execution` Task referencing W005.
- Make T15 depend on this coordination Task.
- Replace parent T10 and W004 T01 dependencies on retired T09 with T15.
- Update W004 provenance and coarse flow to consume reviewed W005 contract Specifications.
- Create W005 T01 as the architecture-to-contract placement Investigation owner.
- Create W005 T02 as the Specification topic-tree and Markdown-placement decision owner after T01.
- Create W005 T03 as the post-decision graph coordination owner after T02.
- Update W001 and W005 Task lists and coarse execution flow.
- Keep ADR routing, canonical authoring, review, and closure Tasks unmaterialized until T02 decides placement and T03 coordinates the route.

This Task must not decide Specification placement, create an Investigation record, author an ADR or Specification, perform review or closure, or start implementation.

## Done condition

- W005 has one Investigation, one decision, and one post-decision coordination Task in deterministic order.
- W001 has exactly one active execution Task for W005.
- W004 execution waits for W005 rather than retired W003.
- Parent and child Task ownership relations are bidirectionally consistent.
- No downstream canonical writer or speculative finding Task exists.

## Verification

- Confirm T15 uses `work_item_execution` and references exactly W005.
- Confirm W005 T01 through T03 use the correct Work Item sequence and task types.
- Confirm dependencies enforce T01, then T02, then T03.
- Confirm T10 depends on T15 and no longer depends on T09.
- Confirm no ADR, Specification, Investigation record, review, closure, or implementation content changed.

## Evidence

- T13 retired the incomplete W003 route and created the W005 decomposition owner.
- W003 T08 created W005 with an independent completion boundary.
- Created parent T15 for W005 execution.
- Created W005 T01 investigation, T02 decision, and T03 coordination Tasks.
- Updated W001 and W005 Task relations and the parent execution route.
- Rerouted W004 parent execution and child bootstrap from T09 to T15.
- Updated W004 provenance and coarse flow to consume W005.
- No canonical contract content, review, synchronization, implementation, stage, or commit work occurred.
- Result: `PASS`.
