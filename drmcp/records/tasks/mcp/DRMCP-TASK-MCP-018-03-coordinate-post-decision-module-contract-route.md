# DRMCP-TASK-MCP-018-03: Coordinate post-decision module-contract route

- **id**: DRMCP-TASK-MCP-018-03
- **status**: done
- **date**: 2026-07-06
- **work_item**: DRMCP-WORK-MCP-018
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-018-02
- **outputs**:
  - DRMCP-TASK-MCP-018-03
  - DRMCP-WORK-MCP-018
  - DRMCP-TASK-MCP-018-04
  - DRMCP-TASK-MCP-018-05
  - DRMCP-TASK-MCP-018-06
  - DRMCP-TASK-MCP-018-07
  - DRMCP-TASK-MCP-018-08

## Goal

Materialize the post-decision W018 graph for ADR routing, ADR authoring, canonical module-contract Specification authoring, integrated review, and closure synchronization.

## Work

- Read the completed W018 decision ledger.
- Add one ADR-routing Task.
- Add one ADR-authoring Task.
- Add one canonical module-contract Specification authoring Task.
- Add one integrated independent review Task.
- Add one closure synchronization Task.
- Update the parent Work Item task list, task flow, and task candidates.
- Preserve the D-012 handoff boundary that production implementation planning remains blocked.
- Do not author ADR body content or Specification body content in this coordination Task.

## Done condition

- The W018 graph contains the required downstream owners.
- The new Tasks have one primary outcome each.
- Dependencies enforce the order from routing to authoring to review to synchronization.
- Conditional correction and finding-closure review Tasks remain unmaterialized until review returns named findings.
- The parent Work Item records that detailed contract convergence and behavior Specification work remain downstream after W018 closure.

## Verification

- Confirmed T04 owns ADR routing.
- Confirmed T05 owns ADR authoring.
- Confirmed T06 owns canonical module-contract Specification authoring.
- Confirmed T07 owns integrated independent review.
- Confirmed T08 owns closure synchronization.
- Confirmed no correction or finding-closure review Task was created.
- Confirmed no implementation Task was created.

## Evidence

- Source decision Task: DRMCP-TASK-MCP-018-02.
- D-011 requires a coherent ADR route for D-001 through D-010.
- D-012 releases ADR authoring, canonical module-contract Specification authoring, and later component-scoped detailed contract convergence.
- D-012 blocks implementation planning until component-level detailed contracts and operation or feature behavior Specifications close.
- DRMCP is non-operational. Filesystem authoring was used as the required fallback.
