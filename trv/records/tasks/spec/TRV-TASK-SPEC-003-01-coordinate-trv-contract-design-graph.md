# TRV-TASK-SPEC-003-01: Coordinate TRV contract design graph

- **id**: TRV-TASK-SPEC-003-01
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-003
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-001-08
- **outputs**:
  - TRV-WORK-SPEC-003
  - TRV-TASK-SPEC-003-01
  - TRV-TASK-SPEC-003-02
  - TRV-TASK-SPEC-003-03

## Goal

Materialize the contract ADR-routing and post-routing design graph for TRV-WORK-SPEC-003.

## Work

- Consume the fixed contract decisions from TRV-TASK-SPEC-001-02 and the accepted W002 architecture closure.
- Confirm no unresolved contract judgment blocks ADR routing.
- Create T02 as one contract ADR-routing and boundary-partitioning decision Task.
- Create T03 as one post-routing coordination Task.
- Require T03 to materialize exact contract ADR authoring, Specification authoring, integrated review, and closure synchronization from terminal T02 routing.
- Keep reconciliation, correction, and finding-closure review Tasks conditional on concrete conflicts or named findings.
- Keep detailed-Specification and implementation work outside W003.

This Task must not:

- make a contract or ADR-routing decision;
- change reviewed architecture;
- author an ADR or Specification;
- perform review, correction, synchronization, or implementation;
- create detailed-Specification or implementation Tasks;
- stage or commit changes.

## Done condition

- T02 owns complete contract ADR routing.
- T03 owns exact post-routing graph materialization.
- Dependencies form T01 to T02 to T03.
- Parent architecture execution T08 is done.
- Conditional conflict and finding branches remain unmaterialized.
- No child-owned canonical artifact or implementation work is authored.

## Verification

- Confirm T02 uses `decision` and T03 uses `coordination`.
- Confirm both Tasks belong to TRV-WORK-SPEC-003.
- Confirm T08 is done.
- Confirm dependencies are acyclic.
- Confirm no detailed-design, review finding, or implementation Task was speculatively created.

## Evidence

- T07 created this child-local graph owner after W003 existed.
- TRV-TASK-SPEC-001-02 supplies terminal contract inputs.
- Parent T08 completed after W002 architecture closure.
- TRV-TASK-SPEC-003-02 was created for contract ADR routing.
- TRV-TASK-SPEC-003-03 was created for post-routing authoring and review coordination.
- W003 lists T01 through T03 in dependency order.
- No ADR, Specification, review, correction, synchronization, detailed design, implementation, stage, or commit work occurred.
- Result: `PASS`.
