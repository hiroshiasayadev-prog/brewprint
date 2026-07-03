# TRV-TASK-SPEC-002-01: Coordinate TRV architecture design graph

- **id**: TRV-TASK-SPEC-002-01
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-002
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-001-07
- **outputs**:
  - TRV-WORK-SPEC-002
  - TRV-TASK-SPEC-002-01
  - TRV-TASK-SPEC-002-02
  - TRV-TASK-SPEC-002-03

## Goal

Materialize the architecture ADR-routing and post-routing design graph for TRV-WORK-SPEC-002.

## Work

- Consume the fixed architecture decisions from TRV-TASK-SPEC-001-02.
- Confirm no unresolved architecture judgment blocks ADR routing.
- Create T02 as one architecture ADR-routing and boundary-partitioning decision Task.
- Create T03 as one post-routing coordination Task.
- Require T03 to materialize exact architecture ADR authoring, Specification authoring, integrated review, and closure synchronization from terminal T02 routing.
- Keep reconciliation, correction, and finding-closure review Tasks conditional on concrete conflicts or named findings.
- Keep contract, detailed-Specification, and implementation work outside W002.

This Task must not:

- make an architecture or ADR-routing decision;
- author an ADR or Specification;
- perform review, correction, synchronization, or implementation;
- create contract or detailed-Specification Tasks;
- stage or commit changes.

## Done condition

- T02 owns complete architecture ADR routing.
- T03 owns exact post-routing graph materialization.
- Dependencies form T01 to T02 to T03.
- Conditional conflict and finding branches remain unmaterialized.
- No child-owned canonical artifact or implementation work is authored.

## Verification

- Confirm T02 uses `decision` and T03 uses `coordination`.
- Confirm both Tasks belong to TRV-WORK-SPEC-002.
- Confirm dependencies are acyclic.
- Confirm no contract, detailed-design, review finding, or implementation Task was speculatively created.

## Evidence

- T07 created this child-local graph owner after W002 existed.
- TRV-TASK-SPEC-001-02 supplies terminal architecture inputs.
- TRV-TASK-SPEC-002-02 was created for architecture ADR routing.
- TRV-TASK-SPEC-002-03 was created for post-routing authoring and review coordination.
- W002 now lists T01 through T03 in dependency order.
- No ADR, Specification, review, correction, synchronization, implementation, stage, or commit work occurred.
- Result: `PASS`.
