# DRMCP-WORK-MCP-018: Establish architecture-derived module contracts

- **id**: DRMCP-WORK-MCP-018
- **status**: done
- **date**: 2026-07-06
- **source_refs**:
  - DRMCP-REQ-MCP-006
  - DRMCP-TASK-MCP-017-02
  - DRMCP-WORK-MCP-016
- **impact_refs**: []
- **tasks**:
  - DRMCP-TASK-MCP-018-01
  - DRMCP-TASK-MCP-018-02
  - DRMCP-TASK-MCP-018-03
  - DRMCP-TASK-MCP-018-04
  - DRMCP-TASK-MCP-018-05
  - DRMCP-TASK-MCP-018-06
  - DRMCP-TASK-MCP-018-07
  - DRMCP-TASK-MCP-018-08
  - DRMCP-TASK-MCP-018-09
  - DRMCP-TASK-MCP-018-10
  - DRMCP-TASK-MCP-018-11

## Goal

Establish reviewed canonical contracts for modules and collaboration boundaries derived from the accepted DRMCP application architecture.

The contracts provide the stable responsibility and interaction baseline required before implementation-ready detailed specifications begin.

## Boundary

This Work Item owns:

- inventorying the accepted architecture and existing operation contracts;
- deriving a contract-level component model below the accepted application architecture;
- giving every active Application Use Case at least one explicit component;
- identifying smaller shared components for reusable orchestration, Domain behavior, source access, state, or failure ownership;
- deriving contract decision units from that component model;
- deciding a coherent, non-overlapping contract partition;
- defining owned responsibility for each module or collaboration boundary;
- defining accepted inputs and outputs;
- defining owned state and state visibility;
- defining dependency and collaboration rules;
- defining failure ownership and failure propagation boundaries;
- defining invariants, preconditions, and postconditions;
- preserving architecture alignment for Current Records, Legacy Archive, Guidance aliases, and source contracts;
- defining downstream detailed-design freedom and architecture-return conditions;
- authoring canonical contract Specifications;
- running integrated independent review and finding closure;
- synchronizing lifecycle, Evidence, and closure state.

This Work Item does not preselect the contract partition.
It does not map five architecture components to five contract Work Items.

Unknown handling:

- Start design convergence with a bounded repository inventory of accepted architecture and existing operation contracts.
- Create a formal Investigation only when the inventory identifies one durable, bounded research question that cannot be resolved inside the decision Task.

## Impact Scope

The W018 canonical module-contract target is `spec:drmcp.implementation.contracts` and its first subdomains under the five accepted architecture components.
The W018 baseline does not decide every component-scoped detailed contract target.
Downstream detailed contract convergence must identify finer-grained canonical targets before detailed authoring.

Existing operation-contract impacts remain outside this Work Item unless a direct contradiction is found during review or closure.
The accepted application architecture is an input authority, not a writable target.
A proposed architecture change returns to the architecture-design route.

## Task flow

```text
DRMCP-TASK-MCP-018-01 initial graph coordination
  -> DRMCP-TASK-MCP-018-02 bounded inventory and contract decisions
     -> conditional Investigation route when required
     -> DRMCP-TASK-MCP-018-03 post-decision graph coordination
        -> DRMCP-TASK-MCP-018-04 ADR routing
           -> DRMCP-TASK-MCP-018-05 ADR authoring
              -> DRMCP-TASK-MCP-018-06 canonical module-contract Specification authoring
                 -> DRMCP-TASK-MCP-018-07 integrated independent review
                    -> DRMCP-TASK-MCP-018-09 finding-route coordination after NEEDS REVISION
                       -> DRMCP-TASK-MCP-018-10 finding correction
                          -> DRMCP-TASK-MCP-018-11 independent finding-closure review
                             -> DRMCP-TASK-MCP-018-08 closure synchronization
                       -> downstream component-scoped detailed contract convergence
                       -> downstream operation or feature behavior Specification work
                       -> implementation planning only after relevant behavior Specifications close
```

T02 owns the completed decision ledger.
T03 through T08 own the post-decision route through ADR authoring, canonical module-contract Specification authoring, review, and closure synchronization.
T09 through T11 own the finding route created after T07 returned `NEEDS REVISION`.
Detailed contract convergence and behavior Specification work remain downstream of W018.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `DRMCP-TASK-MCP-018-01` | `coordination` | Materialize the minimum initial design-convergence graph. | none |
| `DRMCP-TASK-MCP-018-02` | `decision` | Inventory accepted authority and decide the architecture-derived contract partition and boundaries. | T01 |
| `DRMCP-TASK-MCP-018-03` | `coordination` | Materialize the post-decision route through ADR routing, authoring, review, and closure. | T02 |
| `DRMCP-TASK-MCP-018-04` | `coordination` | Route W018 decisions to ADR boundaries without writing ADR body content. | T03 |
| `DRMCP-TASK-MCP-018-05` | `authoring` | Author the required module-contract ADR. | T04 |
| `DRMCP-TASK-MCP-018-06` | `authoring` | Author the canonical module-contract Specification baseline. | T05 |
| `DRMCP-TASK-MCP-018-07` | `review` | Independently review the final W018 decision, ADR, and Specification state. | T06 |
| `DRMCP-TASK-MCP-018-08` | `synchronization` | Synchronize W018 closure after findings close. | T11 |
| `DRMCP-TASK-MCP-018-09` | `coordination` | Materialize the finding-specific correction and closure-review route. | T07 |
| `DRMCP-TASK-MCP-018-10` | `correction` | Correct F-MAJ-W018-07-01 and F-MIN-W018-07-01. | T09 |
| `DRMCP-TASK-MCP-018-11` | `review` | Independently review closure of F-MAJ-W018-07-01 and F-MIN-W018-07-01. | T10 |

No further correction or finding-closure review Task is materialized unless T11 records an open finding or direct regression.

## Completion Condition

### Done conditions

- The contract partition derives from accepted architecture.
- The partition is coherent, non-overlapping, and traceable.
- Every architecture-defined responsibility boundary has the required contract coverage.
- Each contract defines responsibility, input, output, owned state, failure boundary, and invariants.
- Cross-contract collaboration is consistent.
- Inward dependency direction remains intact.
- Public use case chaining is not introduced.
- Current Records and Legacy Archive remain separate.
- The Guidance alias model remains intact.
- The portable `design_records` Current Records model remains intact.
- Any contract change to architecture returns to the architecture route.
- Canonical contract Specifications exist.
- Durable decisions receive ADR routing when required.
- One integrated independent review passes, or every closure-blocking finding is independently closed.
- Closure synchronization completes before this Work Item becomes `done`.
- The contract baseline is sufficient to start component-scoped detailed contract convergence.
- Operation or feature behavior Specifications remain downstream of detailed component contracts.
- Implementation planning remains blocked until the relevant behavior Specifications close.

### Non-goals

- Re-deciding application architecture.
- Exact Go packages or directory layout.
- Exact interfaces, types, structs, functions, methods, or signatures.
- Concrete algorithms.
- Concrete configuration serialization.
- Production implementation or test implementation.
- Implementation Task authoring or implementation execution graph design.
- Reopening W013, reusing its Task, or repairing the cancelled route.

Conceptual request, response, and state shapes are allowed when needed to express a contract.
Implementation-language-specific exact signatures remain excluded.

## Evidence

- DRMCP-REQ-MCP-006 requires architecture-derived module contracts before implementation-ready detailed design.
- DRMCP-WORK-MCP-016 is the completed accepted application-architecture baseline.
- DRMCP-TASK-MCP-017-01 selected `proceed` and fixed the Goal, Boundary, completion contract, Non-goals, unknown handling, direct sources, and initial route.
- DRMCP-TASK-MCP-017-02 created this Work Item from that accepted framing contract.
- DRMCP-WORK-MCP-013 remains a cancelled historical route and is not reopened or reused.
- DRMCP-TASK-MCP-018-01 materialized T02 as the minimum initial graph and changed this Work Item to `in_progress`.
- T02 records the bounded repository inventory and D-001 through D-012.
- T02 is done and its decision loop state is `decision_complete`.
- D-001 fixes component-first derivation before contract partitioning.
- D-002 fixes six explicit active use-case components, one for each active public Application Use Case.
- D-011 routes D-001 through D-010 to one required ADR boundary.
- D-012 blocks implementation planning until component-level detailed contracts and operation or feature behavior Specifications close.
- T03 materialized the post-decision W018 route.
- T04 routed the module-contract decisions to DRMCP-ADR-MCP-013.
- T05 authored DRMCP-ADR-MCP-013.
- T06 authored the canonical module-contract Specification baseline under `spec:drmcp.implementation.contracts`.
- T07 integrated independent review returned `NEEDS REVISION` with F-MAJ-W018-07-01 and F-MIN-W018-07-01.
- T09 materialized the finding-specific route.
- T10 corrected F-MAJ-W018-07-01 and F-MIN-W018-07-01.
- T11 independent finding-closure review returned `PASS`; F-MAJ-W018-07-01 and F-MIN-W018-07-01 are independently `CLOSED`.
- T08 synchronized W018 closure state.
- W018 is `done`.
- No implementation Task or production implementation has been authored.
- DRMCP is non-operational. Design Records MCP cannot author this record, so filesystem authoring is the required fallback.

### Closure result

W018 completed reviewed design closure for the architecture-derived module-contract baseline.

Accepted closure artifacts:

- DRMCP-ADR-MCP-013;
- `spec:drmcp.implementation.contracts`;
- `spec:drmcp.implementation.contracts.composition_lifecycle.contract_boundary`;
- `spec:drmcp.implementation.contracts.mcp_inbound_adapter.contract_boundary`;
- `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary`;
- `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`;
- `spec:drmcp.implementation.contracts.infrastructure_io_adapters.contract_boundary`.

Review route:

- T07 returned `NEEDS REVISION`.
- T10 corrected the named findings.
- T11 returned `PASS` and closed every required finding.
- T08 synchronized closure.

W018 closure releases downstream component-scoped detailed contract convergence.
It does not release production implementation planning.
Operation or feature behavior Specifications remain downstream of detailed component contracts.
