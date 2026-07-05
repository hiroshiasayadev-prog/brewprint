# DRMCP-TASK-MCP-016-01: Coordinate initial application-architecture convergence graph

- **id**: DRMCP-TASK-MCP-016-01
- **status**: done
- **date**: 2026-07-03
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - DRMCP-WORK-MCP-016
  - DRMCP-TASK-MCP-016-01
  - DRMCP-TASK-MCP-016-02
  - DRMCP-TASK-MCP-016-03

## Goal

Materialize the minimum initial Task graph needed to derive DRMCP whole-application architecture from current canonical specifications.

## Work

- Create one lightweight `investigation` Task for current use-case and architecture-input inventory.
- Create one dependent `decision` Task for application component graph, ownership boundaries, dependency direction, collaboration, lifecycle ownership, and downstream return rules.
- Keep ADR routing, canonical authoring, integrated review, finding repair, and closure synchronization abstract until the architecture decisions determine their exact owners and targets.
- Update DRMCP-WORK-MCP-016 with the initial Task list, dependency flow, and release boundary.

## Done condition

- The use-case inventory and architecture decision responsibilities have separate Tasks and completion judgments.
- T03 depends on T02.
- No speculative Investigation record, ADR, authoring, review, correction, or synchronization Task exists.
- DRMCP-WORK-MCP-016 represents the persisted initial graph and downstream abstract route.

## Verification

- Confirm bidirectional Work Item–Task ownership for T01 through T03.
- Confirm unique Task IDs and dependency order.
- Confirm T02 uses the accepted lightweight Investigation Evidence exception.
- Confirm T03 owns design judgment but no canonical authoring.
- Confirm later conditional phases remain unmaterialized.

## Evidence

- The user selected a current-spec use-case inventory before component and boundary decisions.
- The user explicitly judged a separate Investigation record disproportionate because no durable bounded research artifact is needed at this stage.
- `spec:product.design_records.authoring_standards.task_authoring` permits an `investigation` Task to record the bounded result directly in Task Evidence when a downstream Task consumes it.
- DRMCP-TASK-MCP-016-02 and DRMCP-TASK-MCP-016-03 were created with T02 preceding T03.
- DRMCP-WORK-MCP-016 was updated to `in_progress` and now records the initial graph.
- No architecture content, ADR, Specification, implementation, or review outcome was authored by this coordination Task.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
