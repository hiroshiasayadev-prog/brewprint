# TRV-WORK-SPEC-003: Define TRV application contract

- **id**: TRV-WORK-SPEC-003
- **status**: blocked
- **date**: 2026-07-03
- **source_refs**:
  - TRV-TASK-SPEC-001-06
  - TRV-TASK-SPEC-001-02
  - TRV-REQ-SPEC-001
- **impact_refs**:
  - spec:trv
- **tasks**:
  - TRV-TASK-SPEC-003-01
  - TRV-TASK-SPEC-003-02
  - TRV-TASK-SPEC-003-03
  - TRV-TASK-SPEC-003-04
  - TRV-TASK-SPEC-003-05
  - TRV-TASK-SPEC-003-06
  - TRV-TASK-SPEC-003-07
  - TRV-TASK-SPEC-003-08

## Goal

Define and independently review the TRV external and application contract after the application architecture closes.

## Boundary

This Work Item owns:

- standalone stdio MCP interface boundary;
- Task-path input and repository-root safety contract;
- MCP server and tool identity;
- tagged semantic, structural-failure, and execution-failure result envelopes;
- caller-owned human acceptance and rejection boundary;
- compatibility boundary for future DRMCP integration;
- contract ADR routing and required ADR authoring;
- contract Specification and namespace projection;
- one integrated independent contract review;
- contract lifecycle and closure synchronization.

This Work Item is operationally retired.
TRV-WORK-SPEC-005 owns the replacement architecture-derived contract Specification route.
The original Completion Conditions remain unsatisfied and are not weakened to manufacture a `done` state.

This Work Item does not own:

- application architecture changes;
- exact package, file, symbol, Go type, JSON schema implementation, retry implementation, fixture, or command layout;
- production implementation or implementation decomposition;
- current DRMCP integration.

## Impact Scope

| target | impact |
|---|---|
| `spec:trv` | Receives reviewed application-contract registration and child Specification routing. |
| Future TRV contract ADRs | Record durable interface and compatibility choices when routing requires them. |
| Future TRV contract Specifications | Define current inputs, outputs, failure classes, caller boundary, and compatibility behavior. |

## Task flow

```text
historical route:
parent T08 architecture execution complete
  -> TRV-TASK-SPEC-003-01 child graph coordination
     -> TRV-TASK-SPEC-003-02 contract ADR routing
        -> TRV-TASK-SPEC-003-03 post-routing coordination
           -> TRV-TASK-SPEC-003-04 ADR authoring

retired branch:
  -> TRV-TASK-SPEC-003-05 blocked
  -> TRV-TASK-SPEC-003-06 blocked
  -> TRV-TASK-SPEC-003-07 blocked

replacement route:
TRV-TASK-SPEC-001-13 replacement coordination
  -> TRV-TASK-SPEC-003-08 Work Item decomposition
     -> TRV-WORK-SPEC-005
```

T01 through T04 remain historical Evidence.
T08 owns only replacement Work Item decomposition.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `TRV-TASK-SPEC-003-01` | `coordination` | Materialize contract ADR routing and post-routing graph ownership. | Parent T08 |
| `TRV-TASK-SPEC-003-02` | `decision` | Route the fixed application-contract decisions into coherent ADR boundaries. | T01 |
| `TRV-TASK-SPEC-003-03` | `coordination` | Materialize exact contract authoring, review, and closure ownership. | T02 |
| `TRV-TASK-SPEC-003-04` | `authoring` | Author TRV-ADR-SPEC-003 through TRV-ADR-SPEC-005. | T03 |
| `TRV-TASK-SPEC-003-05` | `authoring` | Retired four-file Specification writer. | T04 |
| `TRV-TASK-SPEC-003-06` | `review` | Retired review route without a complete Specification state. | T04 and T05 |
| `TRV-TASK-SPEC-003-07` | `synchronization` | Retired closure route; blocked because no accepted integrated contract review exists. | T06 |
| `TRV-TASK-SPEC-003-08` | `work_item_decomposition` | Create W005 as the replacement architecture-derived contract Specification boundary. | Parent T13 |

## Completion Condition

- Contract decisions from TRV-TASK-SPEC-001-02 have complete ADR routing outcomes.
- Required contract ADRs are accepted.
- Current TRV contract Specifications define external interface, Task input, MCP identity, tagged results, caller interaction, and compatibility.
- Contract content preserves reviewed architecture and leaves implementation-ready detail to W004.
- One integrated independent review returns `PASS`, or every required finding is independently closed.
- Lifecycle, Evidence, relations, and Work Item closure are synchronized.

## Evidence

- TRV-TASK-SPEC-001-02 fixed the external interface, path input, MCP identity, result classes, caller boundary, and compatibility policy.
- TRV-TASK-SPEC-001-06 created this independent contract boundary.
- TRV-TASK-SPEC-003-01 was created as the child-local contract graph owner.
- Parent T08 completed after W002 architecture closure.
- T01 created T02 contract ADR routing and T03 post-routing coordination.
- T02 selected three new contract ADR boundaries and reused PRODUCT-ADR-SPEC-017 for caller-owned human judgment.
- T03 created T04 ADR authoring, T05 Specification authoring, T06 integrated review, and T07 closure synchronization.
- T04 authored the routed external-interface ADR set.
- Post-T04 inspection found that W003 did not first decide the architecture-derived contract Specification topic tree and Markdown placement.
- The user directed retirement of W003 and creation of a separate replacement Work Item.
- W003 is `blocked`, not `done`, because its original Completion Conditions remain unsatisfied.
- T05 through T07 are blocked and must not execute.
- Parent T09 is blocked because W003 will not become `done` under the retired route.
- T13 created T08 as the replacement Work Item decomposition owner.
- T08 created TRV-WORK-SPEC-005 for placement-first architecture-derived contract Specification design.
