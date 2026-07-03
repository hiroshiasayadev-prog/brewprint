# TRV-TASK-SPEC-004-01: Coordinate TRV detailed-Specification graph

- **id**: TRV-TASK-SPEC-004-01
- **status**: not_started
- **date**: 2026-07-03
- **work_item**: TRV-WORK-SPEC-004
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-001-15
- **outputs**:
  - TRV-WORK-SPEC-004
  - TRV-TASK-SPEC-004-01
  - TRV-TASK-SPEC-004-02
  - TRV-TASK-SPEC-004-03

## Goal

Materialize the remaining detailed-design decision and post-decision graph for TRV-WORK-SPEC-004.

## Work

- Consume the accepted W002 architecture and reviewed W005 contract Specification closure.
- Create T02 as one interactive decision Task for unresolved implementation-ready details.
- Create T03 as one post-decision coordination Task.
- Require T02 to decide exact packages, files, symbols, interfaces, schemas, parsing, validation, retry, configuration, launcher, fixture, test, build, and smoke contracts.
- Require T03 to materialize ADR routing, ADR authoring, detailed Specification authoring, integrated review, and closure synchronization after T02 completes.
- Keep reconciliation, correction, and finding-closure review Tasks conditional on concrete conflicts or named findings.
- Keep production implementation and implementation decomposition outside W004.

This Task must not:

- decide detailed-design outcomes;
- change reviewed architecture or contract without an explicit reconvergence route;
- author an ADR or Specification;
- perform review, correction, synchronization, or implementation;
- create an implementation Work Item, implementation Task, or executor prompt;
- stage or commit changes.

## Done condition

- T02 owns the bounded implementation-ready detailed decision ledger.
- T03 owns exact post-decision graph materialization.
- Dependencies form T01 to T02 to T03.
- T01 remains blocked until parent W005 execution T15 is `done`.
- Conditional conflict and finding branches remain unmaterialized.
- No canonical artifact or implementation work is authored.

## Verification

- Confirm T02 uses `decision` and T03 uses `coordination`.
- Confirm both Tasks belong to TRV-WORK-SPEC-004.
- Confirm T15 gates this child route.
- Confirm dependencies are acyclic.
- Confirm no review finding, implementation Work Item, implementation Task, or executor prompt was speculatively created.

## Evidence

- T07 created this child-local graph owner after W004 existed.
- TRV-TASK-SPEC-001-02 fixed the detailed-design topic inventory but not implementation-ready exactness.
- W003 and parent T09 were retired.
- Parent T15 now gates this route until W005 contract Specification closure.
- Graph coordination has not started.
