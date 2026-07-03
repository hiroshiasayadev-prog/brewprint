# TRV-TASK-SPEC-001-07: Coordinate decomposed TRV design execution route

- **id**: TRV-TASK-SPEC-001-07
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-001
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-001-06
- **outputs**:
  - TRV-WORK-SPEC-001
  - TRV-TASK-SPEC-001-07
  - TRV-TASK-SPEC-001-08
  - TRV-TASK-SPEC-001-09
  - TRV-TASK-SPEC-001-10
  - TRV-TASK-SPEC-001-11
  - TRV-TASK-SPEC-001-12
  - TRV-TASK-SPEC-002-01
  - TRV-TASK-SPEC-003-01
  - TRV-TASK-SPEC-004-01

## Goal

Materialize the parent execution route and one child-local graph-bootstrap owner for each decomposed TRV design Work Item.

## Work

- Confirm TRV-WORK-SPEC-002 through TRV-WORK-SPEC-004 exist before creating execution relations.
- Create one child-local `coordination` Task in each child Work Item:
  - `TRV-TASK-SPEC-002-01` for the W002 architecture design graph;
  - `TRV-TASK-SPEC-003-01` for the W003 application-contract design graph;
  - `TRV-TASK-SPEC-004-01` for the W004 implementation-ready detailed-Specification graph.
- Create one parent `work_item_execution` Task for each child:
  - `TRV-TASK-SPEC-001-08` references TRV-WORK-SPEC-002;
  - `TRV-TASK-SPEC-001-09` references TRV-WORK-SPEC-003 and depends on T08;
  - `TRV-TASK-SPEC-001-10` references TRV-WORK-SPEC-004 and depends on T09.
- Make the W002 bootstrap depend on T07.
- Make the W003 bootstrap depend on parent T08.
- Make the W004 bootstrap depend on parent T09.
- Create `TRV-TASK-SPEC-001-11` as the final integrated parent review after T05 and T10.
- Create `TRV-TASK-SPEC-001-12` as verdict-gated parent closure synchronization after T11 or a later accepted finding-closure route.
- Update parent and child `tasks` lists and parent Task flow.
- Keep each child bootstrap responsible only for its child-internal graph.
- Leave finding correction and finding-closure review Tasks unmaterialized until named findings exist.

This Task must not:

- create or split a Work Item;
- author a Requirement, ADR, Specification, or implementation content;
- define child-owned canonical artifacts beyond accepted boundary inputs;
- duplicate child Task graphs in the parent;
- create an implementation Work Item, implementation Task, or executor prompt;
- issue a review verdict, repair findings, synchronize lifecycle, stage, or commit work.

## Done condition

- Each child Work Item has exactly one initial child-local coordination owner.
- The parent has exactly one `work_item_execution` Task per child Work Item.
- Parent execution order is W002, then W003, then W004.
- Child bootstrap release order matches the parent execution route.
- Parent integrated review follows Requirement authoring and all child execution Tasks.
- Parent synchronization follows an accepted review route.
- Parent and child Task ownership relations are bidirectionally consistent.
- No child internals, speculative finding Tasks, or implementation work is materialized.

## Verification

- Confirm every created Task has one `task_type`, primary outcome, and completion judgment.
- Confirm each `work_item_execution` Task has exactly one valid `work_item_ref`.
- Confirm child Tasks use the matching child Work Item ID and work sequence.
- Confirm dependencies are acyclic and enforce W002 before W003 before W004.
- Confirm the parent contains no duplicated child-internal procedures.
- Confirm correction and finding-closure review branches remain abstract.
- Confirm only the declared graph artifacts changed.

## Evidence

- T06 created TRV-WORK-SPEC-002 through TRV-WORK-SPEC-004 with empty Task lists.
- T02 fixed sequential architecture, contract, and detailed-Specification boundaries.
- T02 D-015 prohibits implementation decomposition before reviewed detailed-Specification closure.
- T04 materialized this post-decomposition graph owner.
- TRV-TASK-SPEC-002-01, TRV-TASK-SPEC-003-01, and TRV-TASK-SPEC-004-01 were created as child-local bootstrap owners.
- Parent T08 through T10 were created as one `work_item_execution` Task per child.
- T08, T09, and T10 serialize W002, W003, and W004.
- Parent T11 was created for final integrated review.
- Parent T12 was created for verdict-gated closure synchronization.
- Parent and child `tasks` relations were updated.
- No child canonical artifact, speculative finding Task, implementation Work Item, implementation Task, executor prompt, stage, or commit was created.
- Result: `PASS`.
