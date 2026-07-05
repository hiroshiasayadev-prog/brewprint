# DRMCP-TASK-MCP-016-04: Coordinate post-decision application-architecture route

- **id**: DRMCP-TASK-MCP-016-04
- **status**: done
- **date**: 2026-07-04
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-016-03
- **outputs**:
  - DRMCP-WORK-MCP-016
  - DRMCP-TASK-MCP-016-04
  - DRMCP-TASK-MCP-016-05
  - DRMCP-TASK-MCP-016-06
  - DRMCP-TASK-MCP-016-07
  - DRMCP-TASK-MCP-016-08
  - DRMCP-TASK-MCP-016-09
  - DRMCP-TASK-MCP-016-10

## Goal

Materialize the post-decision Task graph that carries the accepted DRMCP application architecture through routing, canonical authoring, independent review, and closure synchronization.

## Work

- Use the completed T03 decision ledger as the canonical graph input.
- Create one ADR-routing Task covering every T03 decision and the three provisional ADR themes.
- Create one bounded ADR-authoring Task for the complete set that T05 classifies as required.
- Create one bounded canonical Specification-authoring Task for one Overview and four architecture views.
- Create one separate authoring Task for the stale Guidance source-path correction.
- Create one integrated independent review Task after every required writer.
- Create one closure-synchronization Task for the direct `PASS` route.
- Keep finding correction and finding-closure review abstract until T09 records named findings.
- Update DRMCP-WORK-MCP-016 with the exact Task list, dependency flow, target set, conditional finding route, and release order.

## Done condition

- Every T03 decision is routed to ADR routing, architecture Specification authoring, Guidance correction, deferred scope, review, or closure as applicable.
- The canonical four-view Specification authoring route is explicit.
- The three provisional ADR themes are explicit inputs to a separate ADR-routing Task.
- The stale `docs/guides/*.md` source is routed to a separate Specification correction Task.
- Deferred authoring internals remain excluded from current ADR and Specification authoring.
- One integrated independent review follows all required writers.
- A `NEEDS REVISION` verdict routes to a newly materialized finding-specific coordination route.
- One closure-synchronization Task exists for accepted reviewed state.
- Dependencies and release order are explicit.
- No implementation Task exists.

## Verification

- Confirm Work Item ownership is bidirectional for T01 through T10.
- Confirm T05 through T10 use unique sequential IDs.
- Confirm T06 depends on T05.
- Confirm T07 depends on T06.
- Confirm T08 depends on T05 and remains separate from architecture-view authoring.
- Confirm T09 depends on T07 and T08.
- Confirm T10 depends on T09 and stops on an unclosed `NEEDS REVISION` route.
- Confirm no speculative correction or finding-closure review Task exists.
- Confirm no ADR, Specification, implementation, review verdict, finding repair, or closure state was authored.

## Evidence

- Repository-root `prompt_chappy.md` was read before every other repository file in this session.
- `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational. Design Records MCP could not be used, so all Design Record reads and writes used the filesystem fallback.
- `spec:product.design_records.authoring_standards.task_authoring`, `spec:product.design_records.authoring_standards.work_item_authoring`, `spec:product.design_records.authoring_standards.spec_authoring`, `spec:product.design_records.authoring_standards.agent_authoring_policy`, and the design-convergence companions governed this graph change.
- T03 D-001 through D-017 are terminal and remain unchanged.
- T05 owns routing outcomes for all T03 decisions and the three provisional ADR themes.
- T06 owns all ADR files that T05 classifies as required because they share one decision source, writer boundary, and integrated review boundary.
- T07 owns one canonical `spec:drmcp.application_architecture` topic tree with one Overview and four focused view Specifications.
- T08 owns only the stale Guidance source-path correction in the three current Guidance Specifications.
- T09 owns the one final integrated verdict for the combined W016 design state.
- T10 owns only mechanically derived closure propagation after `PASS` or independently closed findings.
- A T09 `NEEDS REVISION` result does not execute T10. It requires a new coordination Task that materializes exact correction and independent finding-closure review Tasks from named findings.
- D-012 and D-013 remain deferred. Their non-binding responsibility example was not promoted into an ADR or canonical Specification contract.
- No ADR routing result, ADR body, Specification body, Guidance correction, review, finding repair, lifecycle closure, implementation, test, stage, or commit was performed.
- The standalone semantic responsibility validator was not invoked because current DRMCP does not integrate it and no operational TRV invocation tool is available in this session.
