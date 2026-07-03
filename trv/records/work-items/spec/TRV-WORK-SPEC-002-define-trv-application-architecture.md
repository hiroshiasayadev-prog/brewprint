# TRV-WORK-SPEC-002: Define TRV application architecture

- **id**: TRV-WORK-SPEC-002
- **status**: done
- **date**: 2026-07-02
- **source_refs**:
  - TRV-TASK-SPEC-001-06
  - TRV-TASK-SPEC-001-02
  - TRV-REQ-SPEC-001
- **impact_refs**:
  - spec:trv
  - spec:trv.application_architecture
  - spec:trv.application_architecture.component_model
  - spec:trv.application_architecture.dependency_model
  - spec:trv.application_architecture.validation_flow
  - spec:trv.application_architecture.boundary
  - spec:trv.model_runtime
  - TRV-ADR-SPEC-001
  - TRV-ADR-SPEC-002
- **tasks**:
  - TRV-TASK-SPEC-002-01
  - TRV-TASK-SPEC-002-02
  - TRV-TASK-SPEC-002-03
  - TRV-TASK-SPEC-002-04
  - TRV-TASK-SPEC-002-05
  - TRV-TASK-SPEC-002-06
  - TRV-TASK-SPEC-002-07
  - TRV-TASK-SPEC-002-08
  - TRV-TASK-SPEC-002-09
  - TRV-TASK-SPEC-002-10
  - TRV-TASK-SPEC-002-11
  - TRV-TASK-SPEC-002-12
  - TRV-TASK-SPEC-002-13
  - TRV-TASK-SPEC-002-14
  - TRV-TASK-SPEC-002-15
  - TRV-TASK-SPEC-002-16
  - TRV-TASK-SPEC-002-17
  - TRV-TASK-SPEC-002-18

## Goal

Define and independently review the TRV application architecture required before application-contract authoring begins.

## Boundary

This Work Item owns:

- ports-and-adapters architecture projection;
- application-core responsibility and dependency direction;
- logical component ownership;
- remote Ollama runtime and provider-adapter ownership;
- startup and dependency-wiring responsibility;
- architecture ADR routing and required ADR authoring;
- architecture Overview, component, dependency, validation-flow, boundary, and runtime Specification projection;
- one integrated independent architecture review;
- architecture lifecycle and closure synchronization.

This Work Item does not own:

- external MCP tool names, input, output, or caller workflow contracts;
- exact package, file, symbol, interface, schema, fixture, or command design;
- production implementation or implementation decomposition;
- current DRMCP integration.

## Impact Scope

| target | impact |
|---|---|
| `spec:trv` | Registers the application-architecture and model-runtime topics. |
| `spec:trv.application_architecture` | Provides the architecture-area Overview, whole-system composition view, and child-Specification navigation. |
| `spec:trv.application_architecture.component_model` | Defines the five top-level components and application-core internal responsibility elements. |
| `spec:trv.application_architecture.dependency_model` | Defines static source dependencies, startup construction edges, and forbidden dependencies. |
| `spec:trv.application_architecture.validation_flow` | Defines the runtime validation sequence and stage ownership. |
| `spec:trv.application_architecture.boundary` | Defines ADR, W002, W003, W004, and PRODUCT handoff boundaries. |
| `spec:trv.model_runtime` | Defines the model port, Ollama adapter, decoded-result, failure, and external runtime boundary. |
| `TRV-ADR-SPEC-001` | Records the ports-and-adapters and component-ownership decision. |
| `TRV-ADR-SPEC-002` | Records the provider-neutral port and remote-runtime decision. |

## Task flow

```text
TRV-TASK-SPEC-002-01 child graph coordination
  -> TRV-TASK-SPEC-002-02 architecture ADR routing
     -> TRV-TASK-SPEC-002-03 post-routing coordination
        -> TRV-TASK-SPEC-002-04 ADR authoring
           -> TRV-TASK-SPEC-002-05 architecture Specification authoring
              -> TRV-TASK-SPEC-002-06 initial integrated independent review
                 -> TRV-TASK-SPEC-002-07 initial closure synchronization
                    -> TRV-TASK-SPEC-002-08 architecture reconvergence coordination
                       -> TRV-TASK-SPEC-002-09 architecture documentation and dependency decision loop
                          -> TRV-TASK-SPEC-002-10 revised architecture ADR routing
                             -> TRV-TASK-SPEC-002-11 revised authoring-route coordination
                                -> TRV-TASK-SPEC-002-12 amend architecture ADRs
                                   -> TRV-TASK-SPEC-002-13 author revised architecture Specifications
                                      -> TRV-TASK-SPEC-002-14 new integrated independent review
                                         -> TRV-TASK-SPEC-002-16 post-review finding-route coordination
                                            -> TRV-TASK-SPEC-002-17 correct revised architecture findings
                                               -> TRV-TASK-SPEC-002-18 independently review finding closure
                                                  -> TRV-TASK-SPEC-002-15 verdict-gated closure synchronization
```

T01 through T07 remain the historical initial closure route. T08 added an append-only reconvergence route after the closed projection was found to lack a proper architecture document set and to permit materially different module dependency graphs.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `TRV-TASK-SPEC-002-01` | `coordination` | Materialize architecture ADR routing and post-routing graph ownership. | Parent T07 |
| `TRV-TASK-SPEC-002-02` | `decision` | Route D-017, D-005, and D-011 into coherent architecture ADR boundaries. | T01 |
| `TRV-TASK-SPEC-002-03` | `coordination` | Materialize exact architecture authoring, integrated-review, and closure ownership. | T02 |
| `TRV-TASK-SPEC-002-04` | `authoring` | Author TRV-ADR-SPEC-001 and TRV-ADR-SPEC-002. | T03 |
| `TRV-TASK-SPEC-002-05` | `authoring` | Author architecture Specifications and register them under `spec:trv`. | T04 |
| `TRV-TASK-SPEC-002-06` | `review` | Independently review the initial combined W002 architecture state. | T04 and T05 |
| `TRV-TASK-SPEC-002-07` | `synchronization` | Synchronize the accepted initial W002 closure and ADR migration state. | T06 or accepted waiver route |
| `TRV-TASK-SPEC-002-08` | `coordination` | Materialize the missing module-dependency decision and reconvergence route. | T07 |
| `TRV-TASK-SPEC-002-09` | `decision` | Decide the architecture document structure, architectural views, module ownership, ports, flow, and dependency graph. | T08 |
| `TRV-TASK-SPEC-002-10` | `decision` | Route the revised architecture decisions into coherent ADR boundaries. | T09 |
| `TRV-TASK-SPEC-002-11` | `coordination` | Materialize revised writers, a new integrated review, and a new closure owner. | T10 |
| `TRV-TASK-SPEC-002-12` | `authoring` | Amend TRV-ADR-SPEC-001 and TRV-ADR-SPEC-002 without reversing their selected alternatives. | T11 |
| `TRV-TASK-SPEC-002-13` | `authoring` | Author the revised architecture Overview, four child Specifications, model-runtime update, and parent registration. | T12 |
| `TRV-TASK-SPEC-002-14` | `review` | Independently review the complete revised ADR and Specification state. | T12 and T13 |
| `TRV-TASK-SPEC-002-15` | `synchronization` | Synchronize ADR migration, Evidence, lifecycle, and W002 closure after accepted review. | T18 with all required findings closed |
| `TRV-TASK-SPEC-002-16` | `coordination` | Materialize the T14 finding correction, independent closure review, and T15 blocker route. | T14 |
| `TRV-TASK-SPEC-002-17` | `correction` | Repair F-MAJ-01, F-MAJ-02, and F-MIN-01 in the component, dependency, and validation-flow Specifications. | T16 |
| `TRV-TASK-SPEC-002-18` | `review` | Independently decide closure of the three T14 findings and release or retain the T15 blocker. | T17 |

## Completion Condition

- Architecture decisions from TRV-TASK-SPEC-001-02 have complete ADR routing outcomes.
- Required architecture ADRs are accepted.
- Current TRV architecture Specifications provide a readable Overview and separate component, dependency, validation-flow, boundary, and runtime views.
- The Overview contains a whole-system composition diagram and authoritative child-topic navigation without duplicating detailed child contracts.
- The architecture fixes one unambiguous module dependency graph and architecture-level port responsibility boundary.
- W003 can define contracts without choosing module ownership or dependency edges.
- Architecture content introduces no external contract or implementation-ready detail owned by later Work Items.
- One new integrated independent review covers the revised final architecture state and returns `PASS`, or every required revised-state finding is independently closed.
- Lifecycle, Evidence, relations, and Work Item closure are synchronized.

## Evidence

- TRV-TASK-SPEC-001-02 fixed the architecture style, component boundary, runtime ownership, and dependency policy.
- TRV-TASK-SPEC-001-06 created this independent architecture boundary.
- TRV-TASK-SPEC-002-01 created T02 architecture ADR routing and T03 post-routing coordination.
- T02 selected two new ADR boundaries: TRV-ADR-SPEC-001 and TRV-ADR-SPEC-002.
- T03 created T04 ADR authoring, T05 Specification authoring, T06 integrated review, and T07 closure synchronization.
- The initial architecture projection did not uniquely fix inbound port ownership, outbound port partitioning, or port data boundaries.
- T08 created T09 through T11 while preserving T06 and T07 as historical completed records.
- T09 completed D-001 through D-013 and fixed the architecture document set, component model, dependency view, validation flow, handoff boundary, port ownership, prompt ownership, provider boundary, application outcome boundary, and startup responsibility.
- T10 routed the durable refinements to non-material amendments of TRV-ADR-SPEC-001 and TRV-ADR-SPEC-002; document structure and derived views route directly to Specifications.
- The user explicitly waived a separate impact Investigation because TRV is a new application without existing implementation, migration, or operational integration impact.
- T11 created T12 through T15 for serialized ADR authoring, Specification authoring, new integrated review, and verdict-gated closure.
- T14 returned `NEEDS REVISION` with F-MAJ-01, F-MAJ-02, and F-MIN-01. All three findings are projection defects and require no new user judgment.
- T16 created T17 correction and T18 independent finding-closure review, then changed T15 to depend on T18 and remain blocked until all findings are closed.
- T04 and T05 completed the two architecture ADRs and their normative Specification projections.
- T06 returned `NEEDS REVISION` only for F-MAJ-01 and F-MIN-01 in `spec:trv`; architecture choices, ADR boundaries, and child Specifications passed.
- Both findings were corrected directly in `spec:trv` without new design judgment.
- The user explicitly waived a separate finding-closure review for these bounded projection defects.
- T07 accepted that explicit waiver route, synchronized both ADR migration fields, and closed the initial W002 state.
- A later architecture gap showed that the closed state permits materially different module and port dependency graphs.
- The initial architecture Specifications also lack a proper architecture Overview and separate component, dependency, validation-flow, and boundary views.
- W002 returned to `in_progress` through T08 append-only reconvergence.
- T17 corrected F-MAJ-01, F-MAJ-02, and F-MIN-01 without changing the accepted architecture decisions.
- T18 independently closed all three findings, reported no direct regression, and released T15.
- T15 synchronized both ADR migration fields to `2026-07-02`, confirmed every W002 Completion Condition as `PASS`, and closed this Work Item.
- T06 and T07 remain historical initial review and closure records; T14 through T18 record the accepted revised-state route.
- W003 may proceed with external contract design without choosing module ownership, port ownership, or dependency direction.
