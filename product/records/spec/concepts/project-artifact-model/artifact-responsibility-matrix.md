# Reference: Artifact responsibility matrix

- **id**: `spec:product.concepts.project_artifact_model.artifact_responsibility_matrix`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `spec:product.concepts.project_artifact_model`

## What this is

Defines what each artifact layer in the brewprint project owns and does not own, establishing the responsibility boundary for the artifact system.

## Artifact responsibility matrix

| artifact | owns | does not own |
|---|---|---|
| ADR | Design decisions, rationale for adoption, rejected alternatives | Current specifications, exploration logs, progress management |
| investigation | Research results, evidence, uncertainties, options, follow-up artifact candidates | Decisions, current specs, cross-cutting progress, completion state |
| requirement | The need / gap / request, stable requirement identity | Research process, implementation procedures, design decisions |
| work item | Entire resolution flow for the requirement, goal state, cross-cutting progress, layer-by-layer impact tracking, task graph | Spec body, decision history, research reports, canonical status of subordinate tasks |
| task | Concrete work closeable in the short term, completion conditions, individual status, verification evidence | Canonical requirements, design decisions, current specs, work item goal state / task graph |
| spec | Current specifications, scope, currently valid contracts | Decision history, exploration logs, progress state |
| internal design | Wiring / route from spec semantics to implementation | Canonical authority for spec semantics, primary source for target model |
| brewprint DSL definition | Target implementation source when the app's DSL pipeline is operational | Decision history, responsibility boundaries of design records |
| source implementation | Generated realization of DSL definitions, or handwritten bootstrap implementation while DSL support is insufficient | Canonical authority for design contracts |
| render | Views derived from DSL definitions | Editable source of truth |
| target implementation | Executable or deployable realization built from source implementation | Canonical authority for design contracts |
| impl note | Handover and review notes for completed implementation | Current specs, future tasks |

External relation / assurance artifacts are not active artifacts in the MVP. If the need arises for managing completeness / evidence / sign-off / approval in ways that endpoint metadata alone cannot express, the decision to introduce a standalone artifact — including its placement, name, and responsibilities — will be made at that time.
