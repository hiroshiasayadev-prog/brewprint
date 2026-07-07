# DRMCP-TASK-MCP-022-01: Decide Domain structural contract baseline from W021

- **id**: DRMCP-TASK-MCP-022-01
- **status**: not_started
- **date**: 2026-07-07
- **work_item**: DRMCP-WORK-MCP-022
- **task_type**: decision
- **estimate**: 1.0d
- **depends_on**:
  - DRMCP-TASK-MCP-021-03
- **outputs**:
  - DRMCP-TASK-MCP-022-01

## Goal

Decide the provisional Domain structural contract baseline needed by Application shared orchestration.

## Inputs

| source | accepted material |
|---|---|
| DRMCP-TASK-MCP-021-03 D-004a | Domain provides semantic views. Application owns use-case selection, aggregation, public projection, and operation-error selection. |
| DRMCP-TASK-MCP-021-03 D-004b | Snapshot construction may scan Markdown source. Snapshot state need not retain every verbatim body. |
| DRMCP-TASK-MCP-021-03 D-004c | Domain provides record summary view material. Application owns public response projection. |
| DRMCP-TASK-MCP-021-03 D-004d | Domain provides validation materials for Application use-case judgment. Application owns validation scope and aggregation. |
| DRMCP-TASK-MCP-021-03 D-005 | Domain gaps include semantic views, record summary view, validation materials, resolver inputs, validation inputs, relation inputs, and exact Domain output boundaries. |

## Decisions to make

- Define Domain semantic views consumed by Application shared orchestration.
- Define record summary view requirements.
- Define validation material boundaries.
- Define parser output boundaries for typed source, typed record, and parse finding.
- Define resolver input and relation-validation input boundaries at structural level.
- Define relation input and relation-graph query output boundaries at structural level.

## Boundary

This Task does not decide Application public response shapes, Application operation sequencing, Go contracts, tests, or production plans.

## Done condition

- Domain structural decisions are decided, deferred, or blocked.
- Domain output boundaries needed by W021 shared orchestration are known.
- W023 can consume the structural baseline without guessing structural transfer boundaries.
- Implementation planning remains blocked.

## Evidence

- DRMCP-TASK-MCP-021-03 closed W021 requirement decisions and routed Domain-facing gaps to W022.
- DRMCP-WORK-MCP-022 owns Domain structural provisional contracts.
