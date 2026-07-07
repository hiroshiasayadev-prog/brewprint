# DRMCP-WORK-MCP-021: Application shared orchestration provisional contracts

- **id**: DRMCP-WORK-MCP-021
- **status**: done
- **date**: 2026-07-06
- **source_refs**:
  - DRMCP-WORK-MCP-020
  - DRMCP-WORK-MCP-018
  - DRMCP-ADR-MCP-013
- **impact_refs**:
  - `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary`
- **tasks**:
  - DRMCP-TASK-MCP-021-01
  - DRMCP-TASK-MCP-021-02
  - DRMCP-TASK-MCP-021-03

## Goal

Decide a provisional Application shared-orchestration baseline for Current Records Snapshot, Legacy exact lookup map, and the Domain outputs that Application needs.

The baseline derives minimum request-state requirements from public use-case contracts before Domain detailed-contract work.

## Boundary

This Work Item owns:

- collecting accepted decisions and related specs for Application shared orchestration;
- deciding minimum Current Records Snapshot requirements;
- deciding minimum Legacy exact lookup-map requirements;
- deciding Application-to-Domain output requirements at requirement level;
- classifying existing-spec coverage and downstream gaps;
- routing Domain-facing gaps to W022.

This Work Item does not own:

- public operation request and response contracts;
- Domain parser, logical tree, graph, resolver, or validator internals;
- Domain output exact shape after W022 starts;
- Infrastructure concrete access fields beyond source contract needs exposed to assembly;
- implementation package layout, Go signatures, structs, interfaces, functions, algorithms, fixtures, or tests;
- changing W018 ownership decisions.

## Impact Scope

Likely target area:

- `spec:drmcp.implementation.contracts.application_use_cases`

W021 does not author focused detailed Specifications. W022 consumes the accepted W021 baseline first.

## Task flow

```text
DRMCP-TASK-MCP-021-01 decision ledger
  -> DRMCP-TASK-MCP-021-02 investigate public-use-case request-state requirements
  -> DRMCP-TASK-MCP-021-03 confirm snapshot requirements through decision loop
  -> DRMCP-TASK-MCP-022-01 consumes Domain-facing gaps
  -> done
```

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `DRMCP-TASK-MCP-021-01` | `decision` | Inventory accepted authority and decide the provisional Application shared-orchestration contract shape. | none |
| `DRMCP-TASK-MCP-021-02` | `investigation` | Collect public-use-case record requirements, map them to implementation-side Domain components, and preserve candidate gaps for decision. Uses the no-separate-INV-record exception. | T01 |
| `DRMCP-TASK-MCP-021-03` | `decision` | Confirm snapshot and lookup-map requirements, Application-to-Domain construction-output requirements, existing-spec coverage, and downstream route through a decision loop. | T02 |
| `DRMCP-TASK-MCP-022-01` | `decision` | Decide the Domain structural contract baseline from W021 T03 decisions. | W021 T03 |

## Completion Condition

- Current Records Snapshot minimum requirements are confirmed.
- Legacy exact lookup-map minimum requirements are confirmed.
- Application-to-Domain construction-output requirements are confirmed at requirement level.
- Existing-spec coverage and concrete gaps are classified.
- Domain-facing gaps are routed to W022.
- W022 can use the provisional baseline without guessing Application-owned assembly responsibilities.
- Implementation planning remains blocked.

## Evidence

- W018 released component-scoped detailed contract convergence.
- DRMCP-ADR-MCP-013 assigns Application ownership of request-level assembly orchestration.
- W020 selected this Work Item as the first child in the provisional detailed contract wave.
- DRMCP-TASK-MCP-021-01 materializes the initial W021 decision ledger.
- DRMCP-TASK-MCP-021-02 materializes the investigation step for public-use-case request-state requirements before W021 authoring.
- DRMCP-TASK-MCP-021-03 confirms snapshot, lookup-map, Application-to-Domain, coverage, gap, and route decisions.
- DRMCP-TASK-MCP-021-03 D-006 routes Domain-facing gaps to W022.
- DRMCP-TASK-MCP-022-01 materializes the W022 decision Task that consumes the W021 T03 decisions.
- DRMCP is non-operational. Design Records MCP cannot author this record, so filesystem authoring is the required fallback.
