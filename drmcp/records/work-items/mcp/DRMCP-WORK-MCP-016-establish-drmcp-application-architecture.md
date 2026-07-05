# DRMCP-WORK-MCP-016: Establish DRMCP application architecture

- **id**: DRMCP-WORK-MCP-016
- **status**: done
- **date**: 2026-07-03
- **source_refs**:
  - DRMCP-REQ-MCP-005
  - DRMCP-TASK-MCP-015-02
- **impact_refs**:
  - DRMCP-REQ-MCP-003
  - DRMCP-ADR-MCP-002
  - DRMCP-ADR-MCP-007
  - DRMCP-ADR-MCP-008
  - DRMCP-ADR-MCP-009
  - DRMCP-ADR-MCP-010
  - DRMCP-ADR-MCP-011
  - DRMCP-ADR-MCP-012
  - spec:drmcp.application_architecture
  - spec:drmcp.application_architecture.application_boundary_and_components
  - spec:drmcp.application_architecture.dependency_and_responsibility
  - spec:drmcp.application_architecture.runtime_and_state
  - spec:drmcp.application_architecture.failure_and_evolution
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.schema.overview
  - spec:drmcp.design_records_mcp.schema.discovery
  - spec:drmcp.design_records_mcp.schema.id_normalization
  - spec:drmcp.design_records_mcp.schema.record_source
  - spec:drmcp.design_records_mcp.schema.authoring_guidance_source
  - spec:drmcp.design_records_mcp.tools.list_authoring_guides
  - spec:drmcp.design_records_mcp.tools.get_authoring_guidance
- **tasks**:
  - DRMCP-TASK-MCP-016-01
  - DRMCP-TASK-MCP-016-02
  - DRMCP-TASK-MCP-016-03
  - DRMCP-TASK-MCP-016-04
  - DRMCP-TASK-MCP-016-05
  - DRMCP-TASK-MCP-016-06
  - DRMCP-TASK-MCP-016-07
  - DRMCP-TASK-MCP-016-08
  - DRMCP-TASK-MCP-016-09
  - DRMCP-TASK-MCP-016-10
  - DRMCP-TASK-MCP-016-11
  - DRMCP-TASK-MCP-016-12
  - DRMCP-TASK-MCP-016-13
  - DRMCP-TASK-MCP-016-14
  - DRMCP-TASK-MCP-016-15
  - DRMCP-TASK-MCP-016-16
  - DRMCP-TASK-MCP-016-17
  - DRMCP-TASK-MCP-016-18
  - DRMCP-TASK-MCP-016-19

## Goal

Establish canonical DRMCP application architecture that provides the application-level baseline for module contracts and implementation-ready detailed specifications.

## Boundary

This Work Item owns DRMCP whole-application architecture decisions and their reviewed canonical projection.

The architecture boundary includes:

- whole-application component boundaries;
- component responsibility ownership;
- dependency direction and forbidden dependencies;
- major runtime collaboration and stage ownership;
- startup, composition, runtime-state, and resource-lifecycle ownership;
- downstream change boundaries that distinguish local refinement from architecture change.

T03 records the initial application-level design. T12 records the revised portable-standards and Guidance decisions. T13 through T16 own replacement routing, authority authoring, canonical projection, and the first finding-closure review. T17 through T19 own the F-MAJ-02 correction route and independent closure review. T10 owns final closure synchronization.

## Impact Scope

The canonical architecture Specification topic is `spec:drmcp.application_architecture`.

The topic contains one Overview and four focused views:

- `spec:drmcp.application_architecture.application_boundary_and_components`;
- `spec:drmcp.application_architecture.dependency_and_responsibility`;
- `spec:drmcp.application_architecture.runtime_and_state`;
- `spec:drmcp.application_architecture.failure_and_evolution`.

The portable standards and Guidance realignment affects:

- `DRMCP-REQ-MCP-003`;
- `spec:drmcp.design_records_mcp.namespace_scanning`;
- `spec:drmcp.design_records_mcp.schema.overview`;
- `spec:drmcp.design_records_mcp.schema.discovery`;
- `spec:drmcp.design_records_mcp.schema.id_normalization`;
- `spec:drmcp.design_records_mcp.schema.record_source`;
- `spec:drmcp.design_records_mcp.schema.authoring_guidance_source`;
- `spec:drmcp.design_records_mcp.tools.list_authoring_guides`;
- `spec:drmcp.design_records_mcp.tools.get_authoring_guidance`.

T05 created the first three whole-application ADRs. T09 later found that their Guidance model conflicted with accepted portable-package authority.

T13 fixed three replacement ADR boundaries:

- create `DRMCP-ADR-MCP-010` for the five-component model, superseding ADR-007;
- create `DRMCP-ADR-MCP-011` for Guidance query-alias ownership, superseding ADR-008;
- create `DRMCP-ADR-MCP-012` for unified Current Records state and lifecycle, superseding ADR-009.

ADR-001 remains current authority. D-012 and D-013 remain outside binding authoring internals.

## Task flow

```text
DRMCP-TASK-MCP-016-01 initial graph coordination
  -> DRMCP-TASK-MCP-016-02 use-case and architecture-input inventory
     -> DRMCP-TASK-MCP-016-03 architecture graph and boundary decisions
        -> DRMCP-TASK-MCP-016-04 post-decision graph coordination
           -> DRMCP-TASK-MCP-016-05 ADR routing
              -> DRMCP-TASK-MCP-016-06 required ADR authoring
                 -> DRMCP-TASK-MCP-016-07 canonical architecture Specification authoring
              -> DRMCP-TASK-MCP-016-08 Guidance source-path Specification correction
           -> DRMCP-TASK-MCP-016-09 integrated independent review
              -> NEEDS REVISION: DRMCP-TASK-MCP-016-11 finding-specific coordination
                 -> DRMCP-TASK-MCP-016-12 revised architecture decision
                    -> DRMCP-TASK-MCP-016-13 replacement ADR routing
                       -> DRMCP-TASK-MCP-016-14 Requirement and ADR authoring
                          -> DRMCP-TASK-MCP-016-15 revised canonical Specification authoring
                             -> DRMCP-TASK-MCP-016-16 independent finding-closure review
                                -> NEEDS REVISION with F-MAJ-02: DRMCP-TASK-MCP-016-17 coordination
                                   -> DRMCP-TASK-MCP-016-18 namespace-scanning correction
                                      -> DRMCP-TASK-MCP-016-19 independent F-MAJ-02 closure review
                                         -> CLOSED: DRMCP-TASK-MCP-016-10 closure synchronization
                                         -> OPEN: new finding-specific coordination
```

T02 uses the accepted lightweight Investigation Evidence exception. It creates no separate Investigation record.

T06 and T08 may proceed after their declared dependencies. T09 waits for both T07 and T08.

T09 recorded named findings. T11 materialized T12 through T16. T16 closed F-BLK-01 and F-MAJ-01 but recorded F-MAJ-02. T17 materialized T18 and T19 and blocked T10 until F-MAJ-02 closure.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `DRMCP-TASK-MCP-016-01` | `coordination` | Materialize the minimum initial design-convergence graph. | none |
| `DRMCP-TASK-MCP-016-02` | `investigation` | Inventory current DRMCP use cases, lifecycle flows, external boundaries, and architecture inputs directly in Task Evidence. | T01 |
| `DRMCP-TASK-MCP-016-03` | `decision` | Decide the application component graph, ownership boundaries, dependency model, collaboration, lifecycle ownership, and downstream return rules. | T02 |
| `DRMCP-TASK-MCP-016-04` | `coordination` | Materialize the post-decision routing, authoring, review, and closure graph. | T03 |
| `DRMCP-TASK-MCP-016-05` | `decision` | Route every T03 decision to three exact ADR boundaries and non-ADR projections. | T04 |
| `DRMCP-TASK-MCP-016-06` | `authoring` | Create `DRMCP-ADR-MCP-007` through `009` and supersede `DRMCP-ADR-MCP-002`. | T05 |
| `DRMCP-TASK-MCP-016-07` | `authoring` | Author one Overview and four canonical architecture views. | T06 |
| `DRMCP-TASK-MCP-016-08` | `authoring` | Correct the stale Guidance source path in three existing Specifications. | T05 |
| `DRMCP-TASK-MCP-016-09` | `review` | Independently review the final combined ADR and Specification state. | T07, T08 |
| `DRMCP-TASK-MCP-016-10` | `synchronization` | Propagate independently closed finding Evidence and close W016 when every condition passes. | T19 |
| `DRMCP-TASK-MCP-016-11` | `coordination` | Materialize the exact F-BLK-01 and F-MAJ-01 reconvergence route. | T09 |
| `DRMCP-TASK-MCP-016-12` | `decision` | Decide the normal Current Records model for portable standards and Guidance. | T11 |
| `DRMCP-TASK-MCP-016-13` | `decision` | Route revised decisions to replacement ADRs and non-ADR projections. | T12 |
| `DRMCP-TASK-MCP-016-14` | `authoring` | Amend DRMCP-REQ-MCP-003 and author replacement ADR authority. | T13 |
| `DRMCP-TASK-MCP-016-15` | `authoring` | Author the revised application architecture, current-source, and Guidance Specifications. | T14 |
| `DRMCP-TASK-MCP-016-16` | `review` | Independently decide F-BLK-01 and F-MAJ-01 closure and record direct regressions. | T15 |
| `DRMCP-TASK-MCP-016-17` | `coordination` | Materialize the exact F-MAJ-02 correction and closure-review route. | T16 |
| `DRMCP-TASK-MCP-016-18` | `correction` | Correct the stale Guidance runtime snapshot boundary in namespace scanning. | T17 |
| `DRMCP-TASK-MCP-016-19` | `review` | Independently decide F-MAJ-02 closure. | T18 |

T19 releases T10 only when F-MAJ-02 and every earlier closure-blocking finding are independently `CLOSED`.

## Completion Condition

### Done conditions

- DRMCP whole-application component boundaries are defined.
- Each component's owned and excluded responsibilities are defined.
- Component dependency direction and forbidden dependencies are defined.
- Major runtime collaboration and stage ownership are defined.
- Startup, composition, runtime-state, and resource-lifecycle ownership are defined at architecture level.
- Changes allowed in downstream specifications are distinguished from changes that must return to architecture decision work.
- Current architecture is authored as one Overview and four canonical Specifications.
- Every durable architecture decision receives an ADR-routing result and any required ADR is authored.
- The portable standards package is loaded as normal Current Records under the `design_records` app namespace.
- Guidance operations project the `spec:design_records.authoring_standards.*` subtree through shared record-query orchestration.
- One integrated independent review passes, or every closure-blocking finding is independently closed.
- Closure synchronization confirms every Completion Condition before W016 becomes `done`.

### Non-goals

- Exact module contracts.
- Package or file layout.
- Exact interfaces, types, functions, methods, constructors, or algorithms.
- Implementation Task authoring.
- Production implementation, tests, or fixtures.
- Repair, reopening, or reuse of the cancelled legacy route.

## Evidence

- DRMCP-REQ-MCP-005 requires an application-architecture baseline before module contracts, detailed specifications, or implementation.
- DRMCP-TASK-MCP-015-01 accepted `proceed` and fixed the Goal, Boundary, completion contract, direct source, unknown handling, and design-convergence route.
- DRMCP-TASK-MCP-016-01 materialized T02 and T03 as the minimum initial graph.
- The user explicitly selected current-spec use-case inventory before architecture graph and boundary decisions.
- T02 records its bounded research directly in Task Evidence under the accepted lightweight Investigation exception.
- T03 completed D-001 through D-017 and selected six components, inward dependencies, request-scoped Current and Legacy state, PRODUCT authority placement, minimal Guidance, deferred authoring internals, three failure classes, architecture-return triggers, four canonical views, and three provisional ADR themes.
- T04 materialized T05 through T10 and fixed the downstream writer, review, conditional finding, and closure order.
- T05 completed final ADR routing. The result is three new ADRs: `DRMCP-ADR-MCP-007`, `DRMCP-ADR-MCP-008`, and `DRMCP-ADR-MCP-009`.
- `DRMCP-ADR-MCP-009` supersedes `DRMCP-ADR-MCP-002` because D-009 changes Legacy loading from unconditional to operation-specific.
- T05 rejected authoring sequencing as an ADR theme. D-002, D-012, D-013, and D-016 require no standalone ADR.
- T06 owns the exact three new ADRs and the superseded status update for `DRMCP-ADR-MCP-002`. T07 owns the canonical architecture topic tree. T08 owns the Guidance source-path correction.
- T09 is the historical integrated review and records `NEEDS REVISION` with F-BLK-01 and F-MAJ-01.
- T11 materialized T12 through T16 from those named findings.
- T12 selected ordinary Current Records treatment for the portable package and fixed-scope Guidance aliases.
- T13 routed ADR-010 through ADR-012 as replacements for ADR-007 through ADR-009.
- T14 and T15 are serialized authority and canonical Specification writers.
- T16 independently closed F-BLK-01 and F-MAJ-01 but recorded F-MAJ-02 as a new Major closure blocker.
- T17 materialized the bounded F-MAJ-02 correction and closure-review route.
- T18 corrects only `spec:drmcp.design_records_mcp.namespace_scanning`.
- T19 independently closed F-MAJ-02 and recorded no new closure blocker.
- T10 synchronized the accepted reviewed result and closed W016.
- D-012 and D-013 remain deferred. Their non-binding responsibility example is not canonical current architecture.
- DRMCP-WORK-MCP-011 remains historical read-runtime Evidence and is not the canonical whole-application architecture.
- DRMCP-WORK-MCP-012 through DRMCP-WORK-MCP-014 and their cancelled Tasks are not reopened or reused.
- DRMCP is non-operational. Filesystem authoring is the required fallback.

### Closure result

- Accepted route: finding-closure route after historical `NEEDS REVISION` reviews.
- T09 remains historical `NEEDS REVISION` Evidence for F-BLK-01 and F-MAJ-01.
- T16 independently closed F-BLK-01 and F-MAJ-01 and recorded F-MAJ-02.
- T16 remains historical `NEEDS REVISION` Evidence.
- T18 repaired the exact F-MAJ-02 projection defect.
- T19 returned `PASS`, independently closed F-MAJ-02, found no new closure blocker, and released T10.
- Every closure-blocking finding is independently `CLOSED`.
- Current accepted ADR authority is DRMCP-ADR-MCP-001 and DRMCP-ADR-MCP-010 through DRMCP-ADR-MCP-012.
- DRMCP-ADR-MCP-007 through DRMCP-ADR-MCP-009 remain historical `superseded` authority.
- The reviewed canonical state includes the five application-architecture Specifications and the portable-standards, Current Records, and Guidance Specifications listed in `impact_refs`.
- D-012 and D-013 remain validly deferred to future authoring architecture work with explicit application-architecture return triggers.
- Every Completion Condition is satisfied.
- T01 through T19 are `done`.
- The `tasks` relation contains T01 through T19, and every owned Task points back to DRMCP-WORK-MCP-016.
- T10 performed only lifecycle, Evidence, completion-result, and relation synchronization.
- No Requirement, ADR, Specification, completed Task outcome, review verdict, finding disposition, or Task graph was changed during closure.
- No production implementation, stage, or commit was performed.
- The standalone semantic responsibility validator was not executed because no operational invocation tool is available. No validator PASS was synthesized.
