# TASK-DATA-002-01: Option A boundary review and handoff

- **id**: TASK-DATA-002-01
- **status**: done
- **date**: 2026-05-31
- **work_item**: WORK-DATA-002
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - Option A task-file helper model minimum selected
  - Deferred model-file / catalog / UC-002 migration scope recorded
  - Concrete TASK-DATA-002-02 spec-alignment input

## Goal

Decide the first executable capability boundary for the helper model / model render follow-up without reopening M15 / `v1.1.0-spec`.

This task reviews the ADR-070 / ADR-071 / ADR-072 / ADR-075 chain and records which part should be implemented first, which parts must remain follow-up, and which UC-002 contract shapes must stay outside WORK-DATA-002.

## Work

- Review the helper model / render ADR chain:
  - ADR-070: file-private helper model
  - ADR-071: task-file helper model Markdown render exposure
  - ADR-072: model / schema catalog view
  - ADR-075: model file render
- Confirm ADR status and dependency risks before implementation.
- Decide whether the first implementation boundary should be:
  - Option A: task-file helper minimum
  - Option B: model-file helper included
- Treat ADR-075 as not implementation-ready unless its proposed status and ADR-073 dependency are explicitly resolved.
- Keep the following out of scope:
  - ADR-073 tagged union model
  - ADR-074 DAG asset TypeRef hint
  - ADR-078 MCP semantic identity / helper model exposure schema
  - M15 / `v1.1.0-spec` reopening
  - REQ-DATA-001 / WORK-DATA-001 edits
  - implementation changes

## Boundary Decision

Option A is selected for WORK-DATA-002.

WORK-DATA-002 proceeds with the task-file helper model minimum:

- file-private helper model parsing / validation / TypeRef resolution for same-file references in task files
- main node / private helper visibility rules needed for task-local helper models
- task-file helper model render exposure in DAG Markdown `## Private models` as required by ADR-071

Option B is deferred. Model-file helper semantics, model-file render exposure, ADR-072 catalog follow-up, ADR-075 resolution, and UC-002 model response helper-shape migration are owned by WORK-DATA-003 or later.

## Deferred UC-002 Migration Shapes

The following model response candidates from INV-DATA-002 are valid only after model-file helper render exposure is resolved outside WORK-DATA-002:

| candidate | INV-DATA-002 ID | file / field | deferred owner | dependency notes |
|---|---|---|---|---|
| `get_source_response.snippet` | N-023 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_response.yaml` / `get_source_response.snippet` | WORK-DATA-003 or later | Requires ADR-075 or equivalent model-file render exposure. |
| `get_reference_tree_response.nodes` | N-014 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml` / `get_reference_tree_response.nodes` | WORK-DATA-003 or later | Requires ADR-075 or equivalent model-file render exposure. |
| `get_reference_tree_response.edges` | N-015 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml` / `get_reference_tree_response.edges` | WORK-DATA-003 or later | Requires ADR-075 or equivalent model-file render exposure; existing `reference` model reuse may need separate care. |

Avoid as first migration even after the model-file boundary is resolved:

- `analyze_impact_response.impacts` / N-005: large mixed shape with enum and possible tagged-union pressure
- `list_endpoints_response.tables` / N-029: deeply nested response shape
- `get_signature_response.signature` / N-021: tagged union candidate
- `inspect_response.signature` / N-026: tagged union candidate
- `inspect_response.members` / N-027: tagged union / object-kind-specific payload candidate
- identity-related ObjectRef shapes: ADR-078 / identity semantics boundary

## Done Condition

- ADR-070 / ADR-071 / ADR-072 / ADR-075 handling order is recorded.
- ADR-075 proposed status and ADR-073 dependency risk are classified.
- Option A task-file helper minimum is selected.
- UC-002 model response migration candidates are explicitly deferred.
- Scope exclusions for ADR-073, ADR-074, ADR-078, and M15 reopening remain explicit.
- TASK-DATA-002-02 input is a concrete spec-alignment scope, not a vague follow-up note.

## Verification

- REQ-DATA-002 and WORK-DATA-002 boundaries are narrowed to Option A task-file helper minimum.
- REQ-DATA-001 and WORK-DATA-001 remain untouched.
- M15 / `v1.1.0-spec` is not reopened.
- ADR-073 / ADR-074 / ADR-078 are treated only as deferred or excluded scope.
- No implementation, renderer, fixture, or YAML migration is performed in this task.

## Evidence

- ADR-070 accepted file-private helper model semantics.
- ADR-071 accepted task-file helper model render exposure.
- ADR-072 is a catalog view and is not required for the Option A task-file minimum.
- ADR-075 remains proposed and depends on ADR-073, so model-file helper render is not implementation-ready for WORK-DATA-002.
- INV-DATA-002 model response helper-shape candidates require model-file helper render exposure and are deferred.
- WORK-DATA-002 now owns only the Option A task-file helper model minimum.
- WORK-DATA-003 receives ADR-072 catalog follow-up, ADR-075 model-file render resolution, model-file helper render exposure, and UC-002 model response helper-shape migration.

## TASK-DATA-002-02 Input

Align specs for the task-file helper model minimum:

- task files may contain file-private helper model definitions
- task-file helper models are private to the defining YAML file
- same-file TypeRef resolution may resolve helper model IDs
- cross-file / cross-module references to task-file helper models are invalid
- DAG Markdown may render `## Private models` for task files containing helper models
- helper models are schema definitions and must not be rendered as Mermaid DAG flow nodes
- diagnostics cover duplicate helper IDs and invalid helper references at the spec level
