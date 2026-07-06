# DRMCP-TASK-MCP-018-01: Coordinate initial module-contract convergence graph

- **id**: DRMCP-TASK-MCP-018-01
- **status**: done
- **date**: 2026-07-05
- **work_item**: DRMCP-WORK-MCP-018
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - DRMCP-WORK-MCP-018
  - DRMCP-TASK-MCP-018-01
  - DRMCP-TASK-MCP-018-02

## Goal

Materialize the minimum initial Task graph for architecture-derived module-contract design.

## Work

- Create one `decision` Task for bounded repository inventory and module-contract decisions.
- Keep the inventory inside the decision Task unless one durable bounded research question requires a formal Investigation.
- Keep ADR routing, canonical authoring, integrated review, finding repair, and closure synchronization abstract until the decisions determine exact owners and targets.
- Update DRMCP-WORK-MCP-018 with the initial Task list, dependency flow, and release boundary.
- Do not preselect the contract partition or map architecture components directly to contract artifacts.

## Done condition

- One decision Task owns the bounded inventory and contract-design judgment.
- DRMCP-TASK-MCP-018-02 depends on this coordination Task.
- No speculative Investigation, ADR-routing, authoring, review, correction, or synchronization Task exists.
- DRMCP-WORK-MCP-018 represents the persisted initial graph and abstract downstream route.
- No module contract or canonical design content is authored by this Task.

## Verification

- Confirm bidirectional Work Item–Task ownership for T01 and T02.
- Confirm unique Task IDs and dependency order.
- Confirm T02 owns design judgment but no canonical authoring.
- Confirm later conditional phases remain unmaterialized.
- Confirm the contract partition remains undecided.

## Evidence

- DRMCP-WORK-MCP-018 requires design convergence from the accepted architecture baseline.
- The initial repository inventory found enough accepted authority to start one decision ledger without a separate Investigation record.
- No durable bounded research question requiring independent ownership has been identified.
- DRMCP-TASK-MCP-018-02 was created as the only downstream Task.
- DRMCP-WORK-MCP-018 was updated to `in_progress` and records the initial graph.
- No contract partition, contract content, ADR, Specification, implementation Task, review outcome, or production implementation was authored.
- DRMCP is non-operational. Design Records MCP cannot author these records, so filesystem authoring is the required fallback.
