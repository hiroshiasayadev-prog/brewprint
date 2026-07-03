# TRV-WORK-SPEC-005: Define architecture-derived TRV contract Specifications

- **id**: TRV-WORK-SPEC-005
- **status**: blocked
- **date**: 2026-07-03
- **source_refs**:
  - TRV-TASK-SPEC-003-08
  - TRV-WORK-SPEC-003
  - spec:trv.application_architecture
- **impact_refs**:
  - spec:trv
- **tasks**:
  - TRV-TASK-SPEC-005-01
  - TRV-TASK-SPEC-005-02
  - TRV-TASK-SPEC-005-03

## Goal

Define and independently review the TRV contract Specification set derived from the reviewed application architecture.

Decide the Specification topic tree and Markdown placement before normative contract authoring begins.

## Boundary

This Work Item owns:

- investigation of architecture-derived application-contract boundaries;
- the `spec:trv` topic tree and Markdown placement for TRV contract Specifications;
- mapping W002 components, application ports, application models, validation-flow boundaries, and outcomes into coherent Specification ownership;
- the boundary between application contracts, external MCP contracts, caller contracts, compatibility contracts, and W004 implementation-ready detail;
- ADR routing or amendment when the accepted Specification partition exposes a durable choice;
- normative contract Specification authoring after placement is decided;
- one integrated independent review of the complete contract Specification state;
- lifecycle, Evidence, relation, and closure synchronization.

This Work Item does not own:

- changes to the reviewed W002 architecture without an explicit return-to-decision route;
- exact Go packages, files, declarations, method signatures, structs, schemas, constructors, path algorithms, retry mechanics, commands, or tests;
- production implementation or implementation decomposition;
- current DRMCP integration.

## Impact Scope

| target | impact |
|---|---|
| `spec:trv` | Receives the decided contract topic tree and reviewed child Specifications. |
| W002 architecture Specifications | Supply the component, port, model, flow, outcome, and ownership boundaries from which contracts are derived. |
| TRV contract ADRs | May be amended, reused, or superseded only after explicit ADR routing. |
| W004 | Remains the owner of exact implementation-ready representations and mechanics. |

## Task flow

```text
TRV-TASK-SPEC-005-01 architecture-to-contract placement Investigation
  -> TRV-TASK-SPEC-005-02 topic-tree and Markdown-placement decision
     -> TRV-TASK-SPEC-005-03 post-decision graph coordination
        -> later ADR routing, authoring, integrated review, and closure Tasks
```

The first design judgment is the contract Specification partition and placement.
Canonical authoring does not begin before T02 is complete and T03 materializes exact writers.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `TRV-TASK-SPEC-005-01` | `investigation` | Inventory architecture-derived contract boundaries and candidate Specification topic trees and Markdown placement. | Parent T14 |
| `TRV-TASK-SPEC-005-02` | `decision` | Decide the contract Specification topic tree, file placement, and ownership partition. | T01 |
| `TRV-TASK-SPEC-005-03` | `coordination` | Materialize exact ADR routing, authoring, review, and closure ownership after the placement decision. | T02 |

## Completion Condition

- The contract Specification topic tree and Markdown placement are explicitly decided.
- Every W002 application port, application model boundary, validation-flow handoff, and application outcome has one normative contract owner or an explicit exclusion.
- External MCP, Task-input, caller, compatibility, and application-contract concerns have non-overlapping Specification ownership.
- Required ADR routing and authoring are complete.
- Current contract Specifications state normative behavior rather than summarize ADR rationale.
- W004-owned implementation representations and mechanics remain excluded.
- One integrated independent review returns `PASS`, or every required finding is independently closed.
- Lifecycle, Evidence, relations, and Work Item closure are synchronized.

## Evidence

- W002 completed the reviewed architecture and defined component, port, model, dependency, validation-flow, and outcome ownership.
- W003 materialized only an external-interface-oriented contract route and did not decide the architecture-derived contract Specification tree.
- The user directed retirement of W003 and creation of a separate Work Item whose first responsibility is Specification Markdown placement.
- TRV-TASK-SPEC-003-08 created this replacement Work Item.
- TRV-TASK-SPEC-001-14 materialized the placement-first child graph and parent execution relation.
- No contract Specification placement or normative contract content has been decided by this Work Item yet.
- TRV-ADR-SPEC-006 blocks this Work Item pending controlled responsibility vocabulary and renewed semantic-feasibility evidence.
- Resume requires an explicit decision that restores TRV contract design.
