# TRV-WORK-SPEC-001: Define Task Responsibility Validator app-local design

- **id**: TRV-WORK-SPEC-001
- **status**: blocked
- **date**: 2026-07-03
- **source_refs**:
  - PRODUCT-TASK-SPEC-021-13
  - TRV-REQ-SPEC-001
  - spec:product.responsibility_boundary_validator
- **impact_refs**:
  - spec:trv
  - TRV-WORK-SPEC-002
  - TRV-WORK-SPEC-003
  - TRV-WORK-SPEC-004
  - TRV-WORK-SPEC-005
- **tasks**:
  - TRV-TASK-SPEC-001-01
  - TRV-TASK-SPEC-001-02
  - TRV-TASK-SPEC-001-03
  - TRV-TASK-SPEC-001-04
  - TRV-TASK-SPEC-001-05
  - TRV-TASK-SPEC-001-06
  - TRV-TASK-SPEC-001-07
  - TRV-TASK-SPEC-001-08
  - TRV-TASK-SPEC-001-09
  - TRV-TASK-SPEC-001-10
  - TRV-TASK-SPEC-001-11
  - TRV-TASK-SPEC-001-12
  - TRV-TASK-SPEC-001-13
  - TRV-TASK-SPEC-001-14
  - TRV-TASK-SPEC-001-15

## Goal

Define the reviewed TRV app-local design needed to implement the Task Responsibility Validator under the PRODUCT semantic contract.

Produce a closed design boundary that can hand off to a separate TRV implementation Work Item.

## Boundary

This Work Item owns:

- establishment of the TRV app-local Requirement boundary;
- app-local design decision work and downstream graph coordination;
- ADR routing and required ADR authoring;
- TRV-owned Specification authoring;
- one integrated independent review of the final TRV design state;
- correction and finding-closure routing when required;
- lifecycle, Evidence, relation, and design-closure synchronization;
- definition of the implementation-readiness gate and future decomposition handoff after detailed-Specification closure.

This Work Item does not own:

- changes to PRODUCT-owned semantic validator behavior;
- checklist wording or PRODUCT checklist ownership;
- production implementation or implementation Tasks;
- executor prompts, build output, runtime smoke, or release evidence;
- current DRMCP integration;
- future DRMCP integration design unless a separate Requirement accepts it;
- completion tracking by PRODUCT-WORK-SPEC-021.

No implementation Work Item is created by this Work Item.
Its shape and identity are decided only after reviewed implementation-ready detailed Specifications expose the final writer map.

## Impact Scope

| target | impact |
|---|---|
| `spec:trv` | Receives routing to reviewed TRV app-local Specifications and the closed design state. |
| `TRV-WORK-SPEC-002` | Owns reviewed TRV application architecture. |
| `TRV-WORK-SPEC-003` | Historical external-interface-oriented contract route. Operationally retired after the Specification-placement gap was found. |
| `TRV-WORK-SPEC-005` | Owns reviewed architecture-derived contract Specification placement and normative contract state after W002. |
| `TRV-WORK-SPEC-004` | Owns reviewed implementation-ready detailed Specifications after W005. |

Future TRV-local Requirement, ADR, and Specification refs are added when their identities are decided.

## Task flow

```text
TRV-TASK-SPEC-001-01 initial graph coordination
  -> TRV-TASK-SPEC-001-02 app-local design decision loop
     -> TRV-TASK-SPEC-001-03 remove unnecessary impact-Investigation route
        -> TRV-TASK-SPEC-001-04 downstream design graph coordination
           -> TRV-TASK-SPEC-001-05 accepted Requirement authoring
              -> TRV-TASK-SPEC-001-06 initial design Work Item decomposition
                 -> TRV-TASK-SPEC-001-07 initial child execution coordination
                    -> TRV-TASK-SPEC-001-08 execute W002 architecture

historical retired branch:
  -> TRV-TASK-SPEC-001-09 execute W003 application contract [blocked]

replacement route:
TRV-TASK-SPEC-001-13 retire W003 and create decomposition owner
  -> TRV-TASK-SPEC-003-08 create W005
     -> TRV-TASK-SPEC-001-14 coordinate W005 execution
        -> TRV-TASK-SPEC-001-15 execute W005 architecture-derived contract Specifications
           -> TRV-TASK-SPEC-001-10 execute W004 implementation-ready detailed Specifications
              -> TRV-TASK-SPEC-001-11 final integrated parent review
                 -> TRV-TASK-SPEC-001-12 parent closure synchronization
```

T04 completed the bounded repository-alignment check and materialized the initial child route.
T13 and T14 preserve the historical W003 branch while replacing the active contract gate with W005.
W005 begins with the Specification topic-tree and Markdown-placement boundary.

Production implementation, implementation Tasks, executor prompts, and current DRMCP integration remain outside this graph.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `TRV-TASK-SPEC-001-01` | `coordination` | Materialize the initial decision and downstream coordination route. | none |
| `TRV-TASK-SPEC-001-02` | `decision` | Decide one bounded TRV app-local design ledger through an interactive loop. | T01 |
| `TRV-TASK-SPEC-001-03` | `coordination` | Amend the graph to remove the unnecessary standalone impact-Investigation route. | T02 |
| `TRV-TASK-SPEC-001-04` | `coordination` | Confirm repository alignment and materialize T05 through T07. | T03 |
| `TRV-TASK-SPEC-001-05` | `authoring` | Author the accepted technology-neutral TRV Requirement and establish the parent Requirement relation. | T04 |
| `TRV-TASK-SPEC-001-06` | `work_item_decomposition` | Create W002 architecture, W003 contract, and W004 implementation-ready detailed-Specification boundaries. | T05 |
| `TRV-TASK-SPEC-001-07` | `coordination` | Create child bootstrap owners, parent child-execution Tasks, final parent review, and closure routing after decomposition. | T06 |
| `TRV-TASK-SPEC-001-08` | `work_item_execution` | Represent W002 as one parent architecture execution unit. | T07 |
| `TRV-TASK-SPEC-001-09` | `work_item_execution` | Historical W003 execution relation; blocked after W003 retirement. | T08 |
| `TRV-TASK-SPEC-001-10` | `work_item_execution` | Represent W004 as one parent detailed-design execution unit. | T15 |
| `TRV-TASK-SPEC-001-11` | `review` | Independently review the final combined TRV design state. | T05 and T10 |
| `TRV-TASK-SPEC-001-12` | `synchronization` | Synchronize accepted parent design closure after the review route. | T11 or later accepted finding closure |
| `TRV-TASK-SPEC-001-13` | `coordination` | Retire W003 and create the W005 decomposition owner. | W003 T04 |
| `TRV-TASK-SPEC-001-14` | `coordination` | Create the W005 initial graph and active parent execution relation. | T13 and W003 T08 |
| `TRV-TASK-SPEC-001-15` | `work_item_execution` | Represent W005 as the active architecture-derived contract-Specification execution unit. | T14 |

Conditional successor candidates after T11:

- child-local reconciliation decision only when a concrete unresolved conflict appears;
- child-local ADR routing and required ADR authoring;
- architecture, contract, detailed Specification, and namespace-overview authoring;
- one integrated independent review per child Work Item;
- one final integrated parent review after all child executions;
- finding-specific correction and independent closure review only after named findings;
- lifecycle, Evidence, relation, and design-closure synchronization;
- later implementation decomposition only after reviewed implementation-ready detailed-Specification closure.

## Completion Condition

- `TRV-REQ-SPEC-001` defines the accepted app-local design need and exclusions.
- Every required app-local decision is terminal.
- The W005 architecture-to-contract placement Investigation and decision are complete before contract authoring begins.
- Every decision has an ADR routing outcome.
- Every required ADR is authored at the accepted lifecycle state.
- TRV Specifications define the current app-local interface, runtime, transport, configuration, packaging, testing, and operational boundaries required for implementation planning.
- PRODUCT-owned semantics remain referenced rather than duplicated or changed.
- Current DRMCP integration remains excluded.
- One integrated independent review returns `PASS`, or every required finding is independently closed under the active review rules.
- Lifecycle, Evidence, relations, and design closure are synchronized.
- The future implementation-readiness gate and decomposition owner are identified without prematurely creating an implementation Work Item.
- No production implementation is performed by this Work Item.

## Evidence

- PRODUCT-TASK-SPEC-021-06 D-004 selected this Work Item identity and design-only completion boundary.
- PRODUCT-TASK-SPEC-021-13 created this independent successor after PRODUCT conceptual design and namespace bootstrap.
- `spec:product.responsibility_boundary_validator` is the controlling cross-app semantic contract.
- `spec:trv` is the active TRV namespace overview.
- PRODUCT-WORK-SPEC-021 does not wait for or track this Work Item's completion.
- TRV-TASK-SPEC-001-01 materialized T01 through T04.
- TRV-TASK-SPEC-001-02 owns the app-local design decision ledger.
- TRV-TASK-SPEC-001-03 corrected the graph by removing the unnecessary standalone impact-Investigation route; no TRV Investigation record was created.
- TRV-TASK-SPEC-001-04 completed bounded repository alignment and materialized T05 through T07.
- TRV-TASK-SPEC-001-05 authored accepted `TRV-REQ-SPEC-001` and established the parent Requirement relation.
- TRV-TASK-SPEC-001-06 created W002 through W004 with empty child Task lists and the coarse W002 to W003 to W004 route.
- TRV-TASK-SPEC-001-07 created child bootstrap owners, T08 through T10 child execution Tasks, T11 parent review, and T12 closure synchronization.
- T09 remains historical and blocked after W003 retirement.
- T13 retired W003 and created the W005 decomposition owner.
- W003 T08 created W005.
- T14 created the W005 placement-first graph and parent T15 execution relation.
- T15 and T10 serialize W005 contract-Specification design before W004 detailed design.
- TRV-ADR-SPEC-006 suspends semantic-validator delivery pending controlled responsibility vocabulary and renewed feasibility evidence.
- T15, W005, and W004 are blocked by that prerequisite.
- T11 owns the final integrated parent review only after a later decision restores the design route.
- T12 owns verdict-gated parent closure synchronization.
- DRMCP is non-operational, so the initial graph was authored through the filesystem fallback.
- TRV-INV-SPEC-001 through TRV-INV-SPEC-005 did not establish reliable lightweight-model semantic validation.
- TRV-ADR-SPEC-006 suspends further contract, detailed-design, and implementation work.
- Restart requires reviewed controlled vocabulary, a new cross-model evaluation, and an explicit restoration decision.
