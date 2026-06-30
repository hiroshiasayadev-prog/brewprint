# DRMCP-WORK-MCP-012: Rebuild current read runtime implementation

- **id**: DRMCP-WORK-MCP-012
- **status**: blocked
- **date**: 2026-06-30
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-WORK-MCP-011
  - DRMCP-WORK-MCP-013
  - DRMCP-WORK-MCP-014
  - DRMCP-ADR-MCP-002
  - DRMCP-ADR-MCP-003
  - DRMCP-ADR-MCP-004
  - DRMCP-ADR-MCP-005
  - DRMCP-ADR-MCP-006
  - spec:drmcp.implementation
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-004
  - DRMCP-WORK-MCP-005
  - DRMCP-WORK-MCP-006
  - DRMCP-WORK-MCP-008
- **tasks**:
  - DRMCP-TASK-MCP-012-01

## Goal

Build a clean replacement current-read runtime for the four accepted W011 operations.

Implement `list_records`, `get_records`, `resolve_reference`, and `validate_records` against the accepted architecture.

## Boundary

This Work Item owns:

- composition-root wiring for the four operations;
- application use cases;
- request-scoped snapshot orchestration;
- core record, index, resolving, and validation infrastructure;
- MCP, filesystem, and configuration adapters;
- general current-read automated tests;
- current-only integration against W008 fixtures;
- aggregate implementation verification;
- independent implementation review and closure.

This Work Item consumes:

- the completed W011 architecture;
- reviewed responsibility contracts from `DRMCP-WORK-MCP-013`;
- reviewed function-level internal specifications from `DRMCP-WORK-MCP-014`;
- `DRMCP-ADR-MCP-002` through `DRMCP-ADR-MCP-006`;
- `spec:drmcp.implementation`;
- the accepted W003, W004, W005, W006, and W008 outputs.

This Work Item does not own:

- authoring-guidance application architecture;
- authoring transaction behavior;
- portable package loading;
- configured legacy archive fallback implementation;
- W-SPEC-001 per-file detector implementation;
- W-SPEC-002 Topics graph implementation;
- PRODUCT semantic rules;
- fixture authoring;
- reuse of W009 production structure, Task graph, writer allocation, or extension seams.

Old code may be inspected only as non-authoritative evidence.
Old code must not be preserved merely to reduce implementation effort.

## Impact Scope

| ref | impact |
|---|---|
| `DRMCP-WORK-MCP-011` | Supplies the accepted replacement read-runtime architecture. |
| `DRMCP-WORK-MCP-013` | Supplies reviewed responsibility-level contracts. |
| `DRMCP-WORK-MCP-014` | Supplies reviewed function-level internal specifications. |
| `DRMCP-ADR-MCP-002` through `DRMCP-ADR-MCP-006` | Supply the accepted snapshot, layering, model, validation, and package decisions. |
| `spec:drmcp.implementation` | Supplies the implementation-architecture authority for the four operations. |
| `DRMCP-WORK-MCP-003` | Supplies current discovery, parsing, and active-index contracts. |
| `DRMCP-WORK-MCP-004` | Supplies compact listing and exact batch-retrieval contracts. |
| `DRMCP-WORK-MCP-005` | Supplies current-first resolution contracts. |
| `DRMCP-WORK-MCP-006` | Supplies validation, diagnostics, and path-exposure contracts. |
| `DRMCP-WORK-MCP-008` | Supplies current-format fixtures consumed by implementation tests. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Responsibility-contract gate | W013 reviewed closure | Responsibility-level contracts are complete and represented in canonical Specification content. |
| B. Function-level internal-specification gate | Phase A | W014 completes reviewed type, function, signature, processing, state, error-flow, and test-seam specifications. |
| C. Execution graph authoring | Phase B and W001 T19 | T01 maps reviewed detailed specifications to exact files, symbols, writers, dependencies, model routes, and verification owners. |
| D. Independent graph review | Phase C | A read-only reviewer judges the persistent execution graph. |
| E. Graph release synchronization | Phase D `PASS` | Explicit executor leaves become eligible. |
| F. Replacement implementation | Phase E | Released executor Tasks implement the four current-read operations. |
| G. Aggregate verification and review | Phase F | One aggregate owner verifies integration, then independent review and closure complete. |

Production implementation must not begin before Phase E.

## Task Candidates

| Task | scope | dependency |
|---|---|---|
| `DRMCP-TASK-MCP-012-01` | Author and freeze the persistent execution graph from reviewed detailed specifications. Stop before production implementation. | `DRMCP-TASK-MCP-001-19` and W014 reviewed closure. |

T01 remains blocked until W013 and W014 complete reviewed closure.
T01 must create the remaining executor, aggregate-verification, independent-review, and release-synchronization Tasks.
No production executor Task exists yet.

## Completion Condition

This Work Item is complete when all of the following are true:

- W013 responsibility contracts complete reviewed closure;
- W014 function-level internal specifications complete reviewed closure;
- the persistent execution graph is independently reviewed and released;
- the four accepted read-runtime operations are implemented through the W011 architecture and reviewed detailed specifications;
- composition-root, application, snapshot, core, and adapter boundaries match the accepted architecture;
- general current-read automated tests pass;
- current-only integration passes against W008 fixtures;
- aggregate implementation verification passes;
- independent implementation review reports no blocking or major findings;
- W009 structure, writer allocation, extension seams, and code reuse are not treated as authority;
- configured legacy fallback remains owned by W010;
- per-file and Topics graph validation remain owned by W-SPEC-001 and W-SPEC-002;
- final implementation and review evidence is recorded before status changes to `done`.

## Evidence

- `DRMCP-WORK-MCP-011`: accepted architecture input.
- `DRMCP-TASK-MCP-001-13`: historical authoring authority for this replacement Work Item and its initial graph-authoring Task.
- `DRMCP-WORK-MCP-013`: responsibility-contract design predecessor.
- `DRMCP-WORK-MCP-014`: function-level internal-specification predecessor.

Blocker:

```text
Awaiting reviewed closure of DRMCP-WORK-MCP-013 and
DRMCP-WORK-MCP-014 before execution-graph authoring.
```

- Production implementation: not released.
- Execution graph: blocked before T01 authoring.
