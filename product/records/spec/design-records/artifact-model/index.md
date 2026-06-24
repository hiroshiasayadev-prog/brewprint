# Overview: Project artifact model

- **id**: `spec:product.design_records.artifact_model`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records`

## What this is

Defines app-independent Design Records artifact responsibilities.
It defines workflow roles, source-of-truth boundaries, and the boundary between traceability contracts and implementation tools.

## Current contract

| area | owns |
|---|---|
| Artifact responsibilities | Responsibilities of ADRs, investigations, requirements, work items, tasks, and specs. |
| Artifact relationships | Relationships among Design Records artifacts from discovery through decision and execution. |
| Source-of-truth boundaries | Separation between current specs, decisions, research, requirements, execution state, and evidence. |
| Workflow roles | App-independent flow from investigation through decision, requirement resolution, work planning, task execution, and spec updates. |
| Traceability boundary | The semantic boundary between Design Records traceability contracts and implementation tools. |

This concept does not own BPDSL language, schema, resolver, render, generation, runtime, or MCP behavior.
It does not own app-specific source generation or target implementation architecture.
It does not own DRMCP parser, storage, UI, tool, or request/response behavior.

## Artifact classes

| artifact | role |
|---|---|
| `<namespace>/records/adr/` | Records accepted design decisions and rationale. |
| `<namespace>/records/investigations/` | Preserves research results, evidence, options, uncertainties, and follow-up candidates. |
| `<namespace>/records/requirements/` | Records stable requirements, gaps, or requests. |
| `<namespace>/records/work-items/` | Coordinates requirement resolution, cross-cutting progress, and task graph ownership. |
| `<namespace>/records/tasks/` | Records concrete closeable work, completion conditions, status, and verification evidence. |
| `<namespace>/records/spec/` | Owns current design contracts and currently valid normative specifications. |

Implementation-flow artifacts are not generic Design Records artifact classes in this concept.
Preserved BPDSL and implementation-flow material is staged under `spec:product.bpdsl`.

## Source of truth and documentation ownership

| concern | owner |
|---|---|
| Current design contracts | `<namespace>/records/spec/**`. |
| Accepted decisions | `<namespace>/records/adr/**`. |
| Research evidence and options | `<namespace>/records/investigations/**`. |
| Requirements and stable gaps | `<namespace>/records/requirements/**`. |
| Cross-cutting resolution flow and task graph | `<namespace>/records/work-items/**`. |
| Short-term task completion evidence | `<namespace>/records/tasks/**`. |
| App-independent record placement | `spec:product.design_records.repository_layout`. |
| Current Brewprint repository inventory | `spec:product.brewprint.layout`. |
| Canonical reference and traceability semantics | `spec:product.design_records.traceability`. |
| DRMCP operational behavior | DRMCP app-local specifications. |
| BPDSL internals and implementation flow | BPDSL app-local specifications after migration review. |

Tasks record execution evidence.
Tasks do not become canonical authority for accepted specs or decisions.

## Disposition of previous ownership statements

The previous source-of-truth table included authoring, migration, and implementation-follow-up statements.
Those statements are preserved here so T04 does not silently drop them.

| previous statement | disposition | preserved location |
|---|---|---|
| ADR authoring is represented by guide ID `adr-authoring`. | Pointer retained. Authoring rules belong to the authoring standards area or guide records. | This section and `spec:product.design_records.authoring_standards`. |
| Artifact responsibility and boundary decisions are represented by guide ID `artifact-boundary`. | Pointer retained. Responsibility semantics stay in this artifact-model area; authoring guidance belongs to the authoring standards area or guide records. | This section, this artifact-model area, and `spec:product.design_records.authoring_standards`. |
| Investigation authoring is represented by guide ID `investigation-authoring`. | Pointer retained. Authoring rules belong to the authoring standards area or guide records. | This section and `spec:product.design_records.authoring_standards`. |
| Requirement, work-item, and task authoring are represented by guide IDs `requirement-authoring`, `work-item-authoring`, and `task-authoring`. | Pointer retained. Authoring rules belong to the authoring standards area or guide records. | This section and `spec:product.design_records.authoring_standards`. |
| Internal-design authoring guidance has no phase-1 guide. | Current non-goal. Internal-design canonical trace and authoring endpoint placement are not accepted current artifact-model contracts. | This section. |
| External relation or assurance artifacts have no MVP owner or directory. | Existing evidence owner. V01-INV-DOCS-002 and V01-ADR-088 preserve the rationale and reintroduction triggers. | T05 Evidence. |
| Legacy M-series records and migration are owned by migration work. | Historical ownership pointer retained. T04 does not create a T05 action for this statement. | This section. |
| `docs/impl/**` stores implementation follow-up and is not source of truth for specs. | Non-normative implementation-follow-up pointer retained. T04 does not create a T05 action for this statement. | This section. |
| After authoring guidance canonicalization, investigation authoring guidance is handled by guide ID `investigation-authoring`. | Pointer retained as historical authoring-guidance context. | This section and `spec:product.design_records.authoring_standards`. |
| If a future decision consolidates all artifact formats under spec, this ownership boundary becomes a refinement target. | Delete. No current requirement or trigger adopts this broad consolidation mechanism. | This section. |

## Navigation

| topic | reference |
|---|---|
| Artifact responsibilities | `spec:product.design_records.artifact_model.artifact_responsibility_matrix`. |
| Change and investigation flow | `spec:product.design_records.artifact_model.change_and_investigation_flow`. |
| Traceability boundary | `spec:product.design_records.artifact_model.traceability_boundary`. |
| Temporary BPDSL staging | `spec:product.bpdsl`. |
| Repository layout | `spec:product.design_records.repository_layout`. |
| Current Brewprint repository inventory | `spec:product.brewprint.layout`. |
| Authoring guidance entrypoints | `spec:product.design_records.authoring_standards` and historical guide IDs preserved in `## Disposition of previous ownership statements`. |
| Artifact model boundary | Concise current boundary in `## Artifact model boundary` in this file. |

## Artifact model boundary

The artifact model owns current Design Records artifact roles: ADR, investigation, requirement, work item, task, and spec.
Semantic realization endpoints and external assurance artifacts are not current artifact classes.
BPDSL-specific preserved material remains in temporary non-canonical staging under `spec:product.bpdsl` until BPDSL app-local migration.
DRMCP operational behavior belongs to DRMCP app-local specifications.

Historical per-item disposition evidence is recorded in T05.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Artifact responsibility matrix | Reference | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | Design Records artifact roles and source-of-truth boundaries. |
| Change and investigation flow | Reference | `spec:product.design_records.artifact_model.change_and_investigation_flow` | App-independent discovery, decision, requirement, work, task, and spec flow. |
| Traceability boundary | Reference | `spec:product.design_records.artifact_model.traceability_boundary` | Traceability as a semantic contract boundary rather than an implementation owner. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.traceability` | Defines canonical reference and validation semantics. |
| `spec:product.design_records.repository_layout` | Defines app-independent record placement. |
| `spec:product.bpdsl` | Temporary preservation area for BPDSL and implementation-flow material. |
| PRODUCT-ADR-SPEC-001 | Accepts the PRODUCT semantic ownership boundary. |

## Sources

- V01-ADR-081: requirement artifacts and spec traceability.
- V01-ADR-083: project artifact boundary and YAML as primary implementation source.
- V01-ADR-085: investigation artifact boundary.
- V01-ADR-086: investigation artifact format and lifecycle.
- V01-ADR-087: Design Records MCP investigation support and semantic ref resolve.
- V01-ADR-088: Reduce semantic trace MVP to a canonical reference resolution foundation.
- V01-ADR-091: Work item/task responsibility separation and legacy milestone migration.
- V01-ADR-092: Design Records MCP workflow artifact record and relation boundary.
- V01-INV-DOCS-002: external coverage artifact necessity for semantic trace MVP.
- V01-INV-DOCS-003: internal-design endpoint necessity for semantic trace MVP.
