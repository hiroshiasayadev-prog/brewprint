# DRMCP-WORK-MCP-022: Domain structural provisional contracts

- **id**: DRMCP-WORK-MCP-022
- **status**: not_started
- **date**: 2026-07-06
- **source_refs**:
  - DRMCP-WORK-MCP-020
  - DRMCP-WORK-MCP-018
  - DRMCP-WORK-MCP-021
  - DRMCP-TASK-MCP-021-03
  - DRMCP-ADR-MCP-013
- **impact_refs**:
  - `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`
- **tasks**:
  - DRMCP-TASK-MCP-022-01

## Goal

Decide and author a provisional detailed contract baseline for Domain structural components and structural transfer types.

The baseline covers Record Parser, Typed Source, Typed Record, Parse Finding, Current Records Logical Tree, Record Relation Graph, Logical-tree Lookup Result, and Relation-graph Query Result.

## Boundary

This Work Item owns:

- collecting accepted decisions and related specs for parser, logical-tree, and graph structure;
- consuming W021 T03 decisions about Application-needed Domain outputs;
- deciding provisional parser inputs and outputs;
- deciding provisional typed source and typed record boundaries;
- deciding provisional logical-tree lookup and relation-graph query transfer types;
- deciding structural parse and addressability states needed by downstream behavior contracts;
- authoring focused detailed contract Specifications under the Record Domain / Logical Tree scope.

This Work Item does not own:

- Reference Resolution behavior beyond structural lookup surfaces;
- Local Record Validation and Relation Graph Validation finding contracts;
- Application response projection;
- Infrastructure source access implementation;
- operation-level request or response contracts;
- implementation package layout, Go signatures, structs, interfaces, functions, algorithms, fixtures, or tests;
- changing W018 ownership decisions.

## Impact Scope

Likely target area:

- `spec:drmcp.implementation.contracts.record_domain_logical_tree`

Exact detailed Specification targets must be decided during this Work Item.

## Task flow

```text
DRMCP-TASK-MCP-022-01 decide Domain structural contract baseline from W021
  -> provisional Specification authoring
  -> local consistency check
  -> done when W023 can consume the baseline
```

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `DRMCP-TASK-MCP-022-01` | `decision` | Decide the provisional Domain structural contract shape from W021 T03 decisions. | DRMCP-TASK-MCP-021-03 |
| TBD | authoring | Author the accepted provisional detailed contract Specifications. | T01 decision complete |

DRMCP-TASK-MCP-022-01 is materialized.

## Completion Condition

- Record Parser inputs and outputs are provisionally specified.
- Typed Source, Typed Record, and Parse Finding transfer type boundaries are provisionally specified.
- Current Records Logical Tree structural responsibilities and lookup output surfaces are provisionally specified.
- Record Relation Graph structural responsibilities and query output surfaces are provisionally specified.
- Raw Markdown crosses into Domain only through parser outputs.
- Domain remains free of I/O.
- W023 can use the provisional structural baseline without guessing structural transfer types.
- Implementation planning remains blocked.

## Evidence

- W018 released component-scoped detailed contract convergence.
- DRMCP-ADR-MCP-013 assigns parser, logical tree, and relation graph responsibilities to Domain.
- W020 selected this Work Item as the second child in the provisional detailed contract wave.
- DRMCP-TASK-MCP-021-03 closes W021 requirement decisions and routes Domain-facing gaps to W022.
- DRMCP-TASK-MCP-022-01 materializes the first W022 decision Task.
