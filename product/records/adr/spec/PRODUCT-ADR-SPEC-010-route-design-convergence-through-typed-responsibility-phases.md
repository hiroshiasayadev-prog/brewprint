# PRODUCT-ADR-SPEC-010: Route design convergence through typed responsibility phases

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**:
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-009
- **supersedes**: []
- **migrated_to_spec**: null

## Context

A complete design workflow contains investigation, judgment, graph change, canonical authoring, review, correction, and synchronization.

A fixed linear sequence would create unnecessary Tasks when no conflict or graph change exists.
A loosely defined sequence would allow one Task to absorb several independent responsibilities.

The workflow needs an explicit responsibility architecture and a stable method for routing discovered mismatches.

## Decision

Use the following responsibility units:

1. decision inventory and interactive decision loop;
2. decision-impact and conflict investigation;
3. conflict resolution and originating-artifact reconciliation decision;
4. conditional execution-graph amendment;
5. conditional Work Item decomposition;
6. ADR routing and ADR-boundary partitioning;
7. conditional ADR authoring;
8. Specification and originating-artifact authoring;
9. integrated independent review;
10. named-finding correction;
11. independent finding-closure review;
12. lifecycle, Evidence, and relation synchronization.

Decision-impact investigation and integrated independent review are mandatory.
Other units are conditional when their owned outcome is unnecessary.

Classify each decision-to-originating-artifact relationship as follows:

| class | meaning | route |
|---|---|---|
| `consistent_refinement` | The decision concretizes the source without changing its meaning or boundary. | Continue without originating-artifact amendment. |
| `stale_representation` | The intended disposition is decided, but source wording or references are stale. | Use `authoring`. |
| `semantic_conflict` | The decision and source disagree materially, and disposition remains unresolved. | Return to `decision`. |
| `workflow_graph_drift` | The current graph no longer supports the accepted route. | Use `coordination`. |

A reconciliation `decision` Task selects the resolution policy.
It does not change the graph or author canonical artifacts.

A `coordination` Task changes Tasks, dependencies, owners, blockers, writer order, review order, and release conditions.
It does not select the design resolution, create child Work Items, or author child deliverables.

A `work_item_decomposition` Task creates or splits child Work Items after the reconciliation decision fixes the Work Item identity boundary.
It does not change the Task graph or author child-owned deliverables.

## Rationale

Typed responsibility phases preserve one outcome and one completion judgment per Task.

Conditional materialization avoids empty or speculative Tasks.
The mismatch classes provide a stable routing rule without imposing one fixed sequence on every topic.

Separating reconciliation, graph amendment, Work Item decomposition, and authoring prevents hidden design decisions during coordination or writing.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Use one mandatory linear Task sequence for every topic. | Topics without conflicts or graph changes would receive unnecessary Tasks. |
| Combine reconciliation and graph amendment. | Selecting a resolution and changing the execution graph have different completion judgments. |
| Let authoring resolve source conflicts. | Authoring must consume decided inputs and stop on material ambiguity. |
| Insert Tasks whenever an inconsistency is detected. | A detected inconsistency may already have a decided disposition and an existing owner. |
| Route only by artifact kind. | The required route depends on semantic state and graph validity, not only the target artifact. |

## Consequences

- The successor skill must define each phase and its stop conditions.
- Investigation records must identify conflicts, graph-change candidates, and shared-writer candidates.
- Work Items must show conditional routing rather than a universal fixed graph.
- Accepted Work Item creation or split routes through `work_item_decomposition`.
- Authoring must stop when decided inputs permit materially different interpretations.
- Synchronization must stop when judgment or graph change is required.

## Evidence

- `PRODUCT-TASK-SPEC-018-01`: D-006 and D-011.
- `PRODUCT-TASK-SPEC-018-02`: B-002 routing boundary.
- `PRODUCT-ADR-SPEC-004`: accepted Task taxonomy.
- `PRODUCT-ADR-SPEC-005`: accepted single-responsibility boundaries.
