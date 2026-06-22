# Overview: Project artifact model

- **id**: `spec:product.concepts.project_artifact_model`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `root`

## What this is

Defines the artifact system of the brewprint project: what artifact layers exist, what each owns, and how the design, decision, and tooling mechanisms relate. The map for deciding where to place information, which layer a change affects, and where Design Records MCP fits.

## Current contract

The artifact system consists of three groups: design/implementation artifacts, decision/discovery/execution artifacts, and trace/tooling mechanisms. The responsibility matrix and flows are defined in child specs. MVP trace/tooling scope is limited to canonical reference resolution foundation — see `spec:product.concepts.traceability`.

## Boundary

What this concept owns:

- The kinds and responsibility boundaries of artifact layers
- Semantic relationships between artifacts and source-of-truth boundaries
- The overall picture of design flow / change flow / traceability flow
- The positioning of the traceability spec, Design Records MCP spec, and authoring guide IDs / directory entrypoints

What this concept does not own:

- Complete schema for front matter / bullet metadata
- Details of semantic ref grammar, coverage mapping schema, and resolver/validation rules
- Details of MCP request/response schemas and diagnostic categories
- Details of authoring templates, file naming, and lifecycle procedures for each artifact
- Migration state of legacy M-series records, implementation follow-ups

Detail contracts and authoring guidance are owned by leaf specs / authoring guide IDs / tasks, in accordance with the ownership boundary shown in this document.

## Artifact classes

The brewprint artifact system consists of three main groups.

### Design and implementation artifacts

The group of artifacts representing the target system / design model and the path to its realization.

| artifact | role |
|---|---|
| `<namespace>/records/spec/` | Canonical authority for current design specifications |
| `docs/internal-design/` | Internal wiring / route from spec to implementation |
| `yaml/` | Primary implementation source for the target design model in brewprint DSL |
| `renders/` | Human / AI readable views derived from brewprint DSL YAML |
| target implementation | Implementation artifact of the target system built from YAML and internal design |

### Decision, discovery, and execution artifacts

The group of artifacts for discovering design changes, making decisions, and dispatching them to execution.

| artifact | role |
|---|---|
| `<namespace>/records/adr/` | Records of design decisions and their rationale |
| `<namespace>/records/investigations/` | Records of research results, evidence, uncertainties, options, and follow-up artifact candidates |
| `<namespace>/records/requirements/` | Stable identity for requirements, gaps, and requests |
| `<namespace>/records/work-items/` | Resolution flow from requirements, goal state, cross-cutting progress, layer-by-layer impact tracking, task graph |
| `<namespace>/records/tasks/` | Concrete work closeable in the short term, completion conditions, and verification evidence |
| `docs/impl/` | Handover and review notes after implementation completion |

### Trace and tooling mechanisms

The mechanism / tool boundary that records, explores, and validates relationships without replacing the meaning of other artifacts.

| mechanism | role |
|---|---|
| trace metadata | Metadata for declaring and referencing `spec:` semantic refs, investigation canonical references, and workflow ID-as-ref relations |
| semantic ref / ID-as-ref | Reference identity independent of path |
| Design Records MCP | Tool boundary responsible for exploring design records / workflow artifacts and resolving / validating canonical refs / declared workflow relations |

The MVP is limited to the canonical reference resolution foundation. The active semantic ref is `spec:`; it handles record artifact ID-as-refs, resolve/validation of investigation `source_refs` / recorded `follow_up_results`, and integrity validation of declared ID-as-ref relations among workflow artifacts. Canonical references from investigation metadata to workflow artifacts are limited to `REQ-*` / `WORK-*`; `TASK-*` is limited to workflow artifact inter-relations and direct resolver inputs.

The `docs/internal-design/` layer continues to exist, but the `internal-design:` endpoint, semantic realization relations, and external relation / assurance artifacts are not treated as operational trace mechanisms in the MVP. These are future candidates to be decided with placement and responsibility if concrete requirements emerge for navigation/impact analysis, YAML cross-layer trace, gap/evidence/sign-off/lifecycle management, etc.

`trace metadata`, `semantic ref`, and Design Records MCP are not new source-of-truth layers. They are mechanism / tool boundaries for mechanically traversing the responsibilities of existing artifacts.

## Source of truth and documentation ownership

| concern | current owner | note |
|---|---|---|
| Existence, responsibilities, and relationships of artifact layers | This concept spec | Map of the project artifact system |
| Current design specifications | `<namespace>/records/spec/**` | Follows spec-first policy |
| Semantic ref / canonical reference / resolve / validation model | `product/records/spec/concepts/traceability/**` | Leaf specs of the canonical reference foundation |
| Design Records MCP record / tool contracts | `drmcp/records/spec/design-records-mcp/**` | |
| ADR authoring | guide ID `adr-authoring` | Design Records MCP authoring guidance |
| Artifact responsibility / boundary decisions | guide ID `artifact-boundary` | Design Records MCP authoring guidance |
| Investigation format / lifecycle / authoring guidance | guide ID `investigation-authoring` | Design Records MCP authoring guidance |
| Requirements / work-items / tasks authoring guidance | guide IDs `requirement-authoring` / `work-item-authoring` / `task-authoring` | Design Records MCP authoring guidance |
| Internal-design authoring guidance | No phase-1 guide | Read related tasks / specs as needed |
| External relation / assurance artifacts | No MVP owner / directory | Will be decided with placement and responsibility when completeness / evidence / sign-off triggers are confirmed |
| Rationale for decisions | `<namespace>/records/adr/**` | Historical decision snapshots |
| Concrete work / individual status / evidence for tasks | `<namespace>/records/tasks/**` | Short-term work units; not source of truth for specs |
| Legacy M-series records and their migration | migration work item | Existing `docs/tasks/m*.md` are treated as legacy milestone-shaped work records until migration; they do not become a new artifact layer |
| Implementation follow-up | `docs/impl/**` | Not source of truth for specs |

After authoring guidance canonicalization, investigation format / lifecycle authoring guidance is handled by guide ID `investigation-authoring`. This concept does not duplicate authoring format schemas. If a future decision consolidates all artifact formats under spec, this ownership boundary itself becomes a refinement target.

## Navigation

| topic | reference |
|---|---|
| Responsibilities and relationships of all artifacts | This document |
| Details of semantic ref / canonical reference / validation | `spec:product.concepts.traceability` |
| Rationale for this boundary | V01-ADR-081, V01-ADR-083, V01-ADR-085–V01-ADR-088, V01-ADR-091 |
| Design Records MCP record and tool contracts | `drmcp/records/spec/` — *(planned; drmcp spec not yet canonicalized)* |
| Internal-design layer artifacts and authoring guidance | no canonical path yet — *(planned; will be decided when `internal-design:` activation is triggered)* |
| Investigation, workflow artifact, and authoring guides | guide IDs `investigation-authoring` / `requirement-authoring` / `work-item-authoring` / `task-authoring` — *(planned; accessible via Design Records MCP once guides are canonicalized)* |

## MVP scope and future extensions

At the M18 / M19 boundary, the artifact layers continue with the responsibilities above. However, the operational scope of trace / tooling is limited to the canonical reference resolution foundation.

Items whose decisions are deferred as future extensions:

- Introducing the `internal-design:` endpoint and the spec / internal-design realization relation (when canonical navigation / impact analysis becomes an operational requirement)
- Creating external relation / assurance artifacts (when completeness / evidence / sign-off / central matrix / relation lifecycle becomes an operational requirement)
- Activating the `yaml:` semantic ref and YAML relation ownership
- Operationally introducing `maps_to` / `covers` relations
- Mapping groups
- Fixture / golden traceability
- Orphan diagnostics for workflow artifacts, task-status-derived progress projection, workflow-dedicated traversal queries, task dependency cycle / execution order projection
- MCP writer tools

These are not a denial of future necessity — they defer decisions on machine-readable trace / tool contracts and external artifact adoption/placement to subsequent decisions. The MVP does not provide directories or authoring entrypoints for external artifacts.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Artifact responsibility matrix | Reference | `spec:product.concepts.project_artifact_model.artifact_responsibility_matrix` | What each artifact layer owns and does not own. |
| Design artifact flow | Reference | `spec:product.concepts.project_artifact_model.design_flow` | Source-of-truth flow and derivation relationships among design and implementation artifacts. |
| Change and investigation flow | Reference | `spec:product.concepts.project_artifact_model.change_and_investigation_flow` | Change discovery, decision, and execution flow among design record and workflow artifacts. |
| Traceability boundary | Reference | `spec:product.concepts.project_artifact_model.traceability_boundary` | Traceability and tool boundary layer and its MVP scope. |

## Sources

- V01-ADR-081: requirement artifacts and spec traceability
- V01-ADR-083: project artifact boundary and YAML as primary implementation source
- V01-ADR-085: investigation artifact boundary
- V01-ADR-086: investigation artifact format and lifecycle
- V01-ADR-087: Design Records MCP investigation support and semantic ref resolve
- V01-ADR-088: Reduce semantic trace MVP to a canonical reference resolution foundation
- V01-ADR-091: Work item / task responsibility separation and legacy milestone migration
- V01-ADR-092: Design Records MCP workflow artifact record and relation boundary
- V01-INV-DOCS-002: external coverage artifact necessity for semantic trace MVP
- V01-INV-DOCS-003: internal-design endpoint necessity for semantic trace MVP
