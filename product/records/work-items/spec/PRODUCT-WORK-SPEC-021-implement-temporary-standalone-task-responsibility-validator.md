# PRODUCT-WORK-SPEC-021: Establish Task responsibility validator conceptual design

- **id**: PRODUCT-WORK-SPEC-021
- **status**: done
- **date**: 2026-07-02
- **source_refs**:
  - PRODUCT-TASK-SPEC-019-10
- **impact_refs**:
  - spec:product.responsibility_boundary_validator
  - spec:product.brewprint.namespaces.app_namespaces
  - spec:product.brewprint.namespaces.domain_catalog
  - spec:product.brewprint.layout
  - spec:trv
  - TRV-WORK-SPEC-001
- **tasks**:
  - PRODUCT-TASK-SPEC-021-01
  - PRODUCT-TASK-SPEC-021-02
  - PRODUCT-TASK-SPEC-021-03
  - PRODUCT-TASK-SPEC-021-04
  - PRODUCT-TASK-SPEC-021-05
  - PRODUCT-TASK-SPEC-021-06
  - PRODUCT-TASK-SPEC-021-07
  - PRODUCT-TASK-SPEC-021-08
  - PRODUCT-TASK-SPEC-021-09
  - PRODUCT-TASK-SPEC-021-10
  - PRODUCT-TASK-SPEC-021-11
  - PRODUCT-TASK-SPEC-021-12
  - PRODUCT-TASK-SPEC-021-13
  - PRODUCT-TASK-SPEC-021-14
  - PRODUCT-TASK-SPEC-021-16

## Goal

Establish one reviewed PRODUCT-owned conceptual-design Specification for semantic Task responsibility validation.

Bootstrap the `TRV` app namespace and create the independent app-local design successor without tracking its completion.

## Boundary

This Work Item owns:

- the corrected PRODUCT conceptual-design completion boundary;
- ADR routing and conditional ADR authoring for D-001 through D-005;
- activation of app namespace `TRV` and domain assignment `TRV` / `SPEC`;
- the minimal `spec:trv` namespace overview and current Brewprint layout projection;
- PRODUCT-owned semantic behavior and workflow invocation policy;
- one integrated independent review of the final PRODUCT state;
- creation of `TRV-WORK-SPEC-001` as an independent successor;
- PRODUCT lifecycle, Evidence, relation, and closure synchronization.

This Work Item does not own:

- TRV-local Requirement, ADR, Specification, interface, runtime, packaging, or implementation design;
- the internal Task graph or completion of `TRV-WORK-SPEC-001`;
- a `work_item_execution` relation for `TRV-WORK-SPEC-001`;
- implementation planning, implementation Tasks, executor prompts, tests, or release evidence;
- changes to PRODUCT-REQ-SPEC-007;
- checklist wording, storage contract, placement, or checklist authoring;
- current or future DRMCP integration;
- complete executor-readiness validation;
- multi-Task or Work Item graph validation;
- automatic Task correction or rewriting.

## Impact Scope

| target | impact |
|---|---|
| `spec:product.responsibility_boundary_validator` | Receives the reviewed PRODUCT semantic contract and corrected ownership boundary. |
| `spec:product.brewprint.namespaces.app_namespaces` | Registers `TRV` as an active app namespace. |
| `spec:product.brewprint.namespaces.domain_catalog` | Registers the active `TRV` / `SPEC` domain assignment. |
| `spec:product.brewprint.layout` | Records the established `trv/records/` namespace layout. |
| `spec:trv` | Provides the minimal app namespace overview without app-local design decisions. |
| `TRV-WORK-SPEC-001` | Receives the independent app-local design boundary after PRODUCT review. |

## Task flow

```text
PRODUCT-TASK-SPEC-021-01 initial graph coordination
  -> PRODUCT-TASK-SPEC-021-02 historical implementation-boundary decision
  -> PRODUCT-TASK-SPEC-021-03 historical implementation impact Investigation
     -> PRODUCT-TASK-SPEC-021-04 blocked obsolete executor-ready route
     -> PRODUCT-TASK-SPEC-021-05 PRODUCT-to-app graph repair
        -> PRODUCT-TASK-SPEC-021-06 PRODUCT-to-app bootstrap decision
           -> PRODUCT-TASK-SPEC-021-07 post-decision graph coordination
              -> PRODUCT-TASK-SPEC-021-08 ADR routing
              -> PRODUCT-TASK-SPEC-021-09 post-routing graph coordination
                 -> PRODUCT-TASK-SPEC-021-16 amend PRODUCT-ADR-SPEC-016
                 -> PRODUCT-TASK-SPEC-021-10 TRV namespace profile authoring
                 -> PRODUCT-TASK-SPEC-021-11 PRODUCT conceptual Specification authoring
                 -> PRODUCT-TASK-SPEC-021-12 integrated independent review
                    -> PASS, independently closed findings, or accepted exact mechanical correction
                       -> PRODUCT-TASK-SPEC-021-13 successor Work Item decomposition
                          -> PRODUCT-TASK-SPEC-021-14 PRODUCT closure synchronization
```

T09 does not create a no-op ADR authoring Task when T08 selects only `covered` or `not_required` routes.
T09 serializes every required ADR writer before T10.
T12 normally routes `NEEDS REVISION` through later finding-specific coordination.
For F-MIN-01, the user explicitly accepted the exact corrected mechanical projection without another review.
T13 creates `TRV-WORK-SPEC-001` but does not wait for its completion.
No `work_item_execution` or implementation route belongs to W021.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `PRODUCT-TASK-SPEC-021-01` | `coordination` | Materialize the initial decision, Investigation, and post-Investigation coordination route. | none |
| `PRODUCT-TASK-SPEC-021-02` | `decision` | Preserve the historical implementation-boundary ledger. | T01 |
| `PRODUCT-TASK-SPEC-021-03` | `investigation` | Produce PRODUCT-INV-SPEC-009 for the historical implementation impact question. | T02 |
| `PRODUCT-TASK-SPEC-021-04` | `coordination` | Preserve the blocked obsolete executor-ready route. | T03 |
| `PRODUCT-TASK-SPEC-021-05` | `coordination` | Repair the PRODUCT-to-app ownership graph and create successor owners. | T03 |
| `PRODUCT-TASK-SPEC-021-06` | `decision` | Decide W021 disposition, TRV identity, authority partition, successor boundary, and closure route. | T05 |
| `PRODUCT-TASK-SPEC-021-07` | `coordination` | Materialize the accepted post-decision graph. | T06 |
| `PRODUCT-TASK-SPEC-021-08` | `decision` | Route D-001 through D-005 to exact ADR boundaries and canonical targets. | T07 |
| `PRODUCT-TASK-SPEC-021-09` | `coordination` | Materialize required ADR authoring owners and release the canonical authoring chain. | T08 |
| `PRODUCT-TASK-SPEC-021-16` | `authoring` | Amend `PRODUCT-ADR-SPEC-016` with the non-material TRV ownership clarification. | T09 |
| `PRODUCT-TASK-SPEC-021-10` | `authoring` | Activate the TRV app and SPEC domain namespace profile and establish `spec:trv`. | T09, T16 |
| `PRODUCT-TASK-SPEC-021-11` | `authoring` | Project the fixed PRODUCT semantic contract and corrected ownership boundary. | T10 |
| `PRODUCT-TASK-SPEC-021-12` | `review` | Independently review the final combined W021 PRODUCT state. | T11 |
| `PRODUCT-TASK-SPEC-021-13` | `work_item_decomposition` | Create `TRV-WORK-SPEC-001` after an accepted W021 review route. | T12 |
| `PRODUCT-TASK-SPEC-021-14` | `synchronization` | Synchronize W021 lifecycle, Evidence, relations, and closure after T13. | T12, T13 |

Required ADR authoring Tasks are materialized only after T08 identifies exact boundaries.
Correction and finding-closure review Tasks remain conditional on named findings.
No implementation Task, executor prompt, or child execution-tracking Task belongs to W021.

## Completion Condition

- D-001 through D-005 remain terminal and unchanged.
- Every decision has a complete ADR routing outcome.
- Every required ADR is authored through the routed disposition.
- `TRV` is an active app namespace with active `SPEC` domain assignment.
- `spec:trv` and the current Brewprint layout reflect the established namespace.
- `spec:product.responsibility_boundary_validator` states the reviewed PRODUCT semantic contract and corrected TRV ownership boundary.
- Integrated independent review returns `PASS`, every required finding is independently closed, or an exact corrected mechanical-only finding is explicitly accepted by the user.
- `TRV-WORK-SPEC-001` exists with the accepted app-local design boundary.
- W021 lifecycle, Evidence, relations, and closure state are synchronized.
- W021 does not wait for or track `TRV-WORK-SPEC-001` completion.
- No implementation planning, implementation, executor prompt, current DRMCP integration, or checklist authoring is performed.

## Evidence

- PRODUCT-TASK-SPEC-019-09 separated implementation from checklist authoring.
- PRODUCT-TASK-SPEC-019-10 created the original implementation-oriented W021 boundary.
- W019 design closure is accepted through PRODUCT-TASK-SPEC-019-19.
- W020 checklist review returned `PASS` through PRODUCT-TASK-SPEC-020-05.
- PRODUCT-TASK-SPEC-020-06 synchronized W020 closure and released the checklist set.
- T01 materialized the initial decision, Investigation, and post-Investigation route.
- T02 and PRODUCT-INV-SPEC-009 remain historical Evidence.
- PRODUCT-INV-SPEC-009 exposed the missing app namespace and app-local design boundary.
- The user clarified that PRODUCT owns conceptual design only.
- T04 remains blocked because direct executor-ready implementation crossed the PRODUCT boundary.
- T05 repaired the graph and created T06 and T07.
- T06 continued W021 with a reviewed PRODUCT conceptual-design completion boundary.
- T06 selected `TRV`, `Task Responsibility Validator`, `trv/`, and `TRV-WORK-SPEC-001`.
- T06 excluded child completion tracking and every `work_item_execution` relation.
- T07 materialized T08 through T14 and released the PRODUCT design route.
- Corrected T08 classified the W021 workflow correction as `not_required` for ADR routing.
- Corrected T08 selected only one non-material ADR amendment.
- T09 materialized T16 as the sole ADR authoring owner before T10.
- T16 owns `PRODUCT-ADR-SPEC-016` amendment after T09.
- No namespace content, ADR, Specification, child Work Item, implementation Task, production implementation, or executor prompt was produced by T07 through T09.
- T10 activated `TRV`, `TRV` / `SPEC`, `trv/records/`, and `spec:trv`.
- T11 projected the PRODUCT-owned semantic contract and TRV ownership boundary.
- T12 independently reviewed the combined state and returned `NEEDS REVISION` only for F-MIN-01.
- F-MIN-01 identified a metadata-to-prose impact relation mismatch with one mechanically determined repair.
- The correction added `spec:trv` and `TRV-WORK-SPEC-001` to `impact_refs` and removed the read-only PRODUCT-WORK-SPEC-020 checklist row from `## Impact Scope`.
- The correction changed no decision, semantic contract, ownership boundary, completion condition, Task graph, lifecycle, or release behavior.
- The user explicitly accepted the corrected F-MIN-01 projection without another review and authorized W021 closure.
- T12 remains unchanged as historical review Evidence.
- T13 created `TRV-WORK-SPEC-001` as an independent app-local design successor.
- `TRV-WORK-SPEC-001` owns its own Requirement, decisions, ADRs, Specifications, review, and design closure.
- W021 does not wait for or track `TRV-WORK-SPEC-001` completion.
- T14 verified every W021 Completion Condition and synchronized lifecycle, Evidence, and relations.
- PRODUCT-WORK-SPEC-021 is closed as `done`.
- T04 remains validly `blocked` as the obsolete historical executor-ready route.
- No `work_item_execution`, implementation planning, implementation, current DRMCP integration, stage, or commit work was performed.
