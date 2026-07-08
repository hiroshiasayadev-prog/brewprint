# DRMCP-WORK-MCP-013: Read-runtime responsibility contract design

- **id**: DRMCP-WORK-MCP-013
- **status**: cancelled
- **date**: 2026-06-30
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-011
  - DRMCP-WORK-MCP-012
  - DRMCP-WORK-MCP-014
  - DRMCP-ADR-MCP-002
  - DRMCP-ADR-MCP-003
  - DRMCP-ADR-MCP-004
  - DRMCP-ADR-MCP-005
  - DRMCP-ADR-MCP-006
  - spec:drmcp.implementation
- **tasks**:
  - DRMCP-TASK-MCP-013-01

## Goal

Establish reviewed responsibility-level contracts for every component in the accepted read-runtime architecture.

Complete responsibility contracts before function-level internal specification or implementation planning begins.

## Boundary

This Work Item owns:

- investigation of the correct responsibility-contract Work Item partition;
- creation of one child Work Item per accepted partition unit;
- hub Tasks that track each child Work Item;
- a required child workflow for decision inventory, interactive decision loop, Specification synchronization, independent review, correction, re-review, and closure;
- an overall cross-contract consistency review after all child Work Items close;
- final responsibility-contract design closure.

This Work Item does not own:

- deciding responsibility contracts inside the partition-investigation Task;
- function-level types, functions, methods, signatures, or algorithms;
- execution-graph authoring;
- production implementation;
- tests or fixtures;
- W010 configured legacy fallback design;
- W-SPEC-001 or W-SPEC-002 implementation design.

Each child Work Item owns one bounded responsibility-contract design topic.
The child Work Item must use the repository-persistent design-decision workflow.

## Impact Scope

| ref | impact |
|---|---|
| `DRMCP-WORK-MCP-011` | Supplies the accepted package and responsibility architecture. |
| `spec:drmcp.implementation` | Supplies the current architecture and receives reviewed responsibility contracts. |
| `DRMCP-WORK-MCP-014` | May begin only after this Work Item completes reviewed contract closure. |
| `DRMCP-WORK-MCP-012` | Remains blocked until this Work Item and W014 complete. |
| `DRMCP-WORK-MCP-001` | Tracks this Work Item through its hub Task. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Work-partition investigation | W011 architecture closure | T01 identifies coherent responsibility-contract decision units and authors the persistent child graph. |
| B. Child contract workflows | Phase A | One hub Task tracks each child Work Item through decision inventory, interactive decisions, Specification synchronization, review, and closure. |
| C. Overall contract review | All child Work Items closed | A read-only reviewer checks cross-responsibility inputs, outputs, ownership, state, errors, and invariants. |
| D. Conditional correction and re-review | Overall review findings | Named owners correct findings and an independent reviewer closes them. |
| E. Contract design closure | Overall review accepted | Evidence and lifecycle state are synchronized before W014 starts. |

## Task Candidates

| Task | scope | dependency |
|---|---|---|
| `DRMCP-TASK-MCP-013-01` | Investigate the responsibility-contract Work Item partition and author the downstream hub graph. | W011 design closure. |
| Future T02 and later | Track partitioned child Work Items and own overall review and closure gates. Exact IDs are allocated by T01. | T01 output. |

T01 must not preselect partition units without repository evidence.
T01 must not decide a responsibility contract.

## Completion Condition

- T01 records an evidence-based and non-overlapping Work Item partition.
- Every partition unit has one child Work Item and one W013 hub Task.
- Every child Work Item includes decision-item investigation, an interactive decision loop, Specification synchronization, independent review, conditional correction and re-review, and closure synchronization.
- Every responsibility contract is represented in canonical Specification content.
- All child Work Items complete reviewed closure.
- An overall cross-contract review passes with no blocking or major finding.
- Required overall findings are independently closed.
- No function-level internal specification or production implementation begins inside this Work Item.
- W014 receives a reviewed contract baseline.

## Evidence

- `DRMCP-WORK-MCP-011`: accepted read-runtime architecture.
- `spec:drmcp.implementation`: current implementation-architecture Specification.
- Child partition and contract decisions: pending T01 and downstream child Work Items.
- Production implementation remains blocked.

### Cancellation disposition

- Cancellation date: 2026-07-03.
- Intentional-stop reason: Contract design must follow a newly framed DRMCP application architecture rather than the completed read-runtime-only W011 baseline.
- Unfinished Goal disposition: The contract need continues under `DRMCP-REQ-MCP-006` after `DRMCP-REQ-MCP-005` reaches an accepted architecture baseline.
- Owned Task `DRMCP-TASK-MCP-013-01` changed from `not_started` to `cancelled`.
- External tracking Task `DRMCP-TASK-MCP-001-18` changed from `not_started` to `cancelled`.
- The former downstream W014 route is cancelled and replaced by separate framing from `DRMCP-REQ-MCP-007`.
- No child Work Item, contract decision, Specification update, stage, or commit occurred under this Work Item.
