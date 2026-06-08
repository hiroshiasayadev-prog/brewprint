# V01-TASK-DATA-002-01: Option A boundary review and handoff

- **id**: V01-TASK-DATA-002-01
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-DATA-002
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - Option A task-file helper model minimum selected
  - Deferred model-file / catalog / UC-002 migration scope recorded
  - Concrete V01-TASK-DATA-002-02 spec-alignment input

## Goal

Decide the first executable capability boundary for the helper model / model render follow-up without reopening M15 / `v1.1.0-spec`.

This task reviews the V01-ADR-070 / V01-ADR-071 / V01-ADR-072 / V01-ADR-075 chain and records which part should be implemented first, which parts must remain follow-up, and which UC-002 contract shapes must stay outside V01-WORK-DATA-002.

## Work

- Review the helper model / render ADR chain:
  - V01-ADR-070: file-private helper model
  - V01-ADR-071: task-file helper model Markdown render exposure
  - V01-ADR-072: model / schema catalog view
  - V01-ADR-075: model file render
- Confirm ADR status and dependency risks before implementation.
- Decide whether the first implementation boundary should be:
  - Option A: task-file helper minimum
  - Option B: model-file helper included
- Treat V01-ADR-075 as not implementation-ready unless its proposed status and V01-ADR-073 dependency are explicitly resolved.
- Keep the following out of scope:
  - V01-ADR-073 tagged union model
  - V01-ADR-074 DAG asset TypeRef hint
  - V01-ADR-078 MCP semantic identity / helper model exposure schema
  - M15 / `v1.1.0-spec` reopening
  - V01-REQ-DATA-001 / V01-WORK-DATA-001 edits
  - implementation changes

## Boundary Decision

Option A is selected for V01-WORK-DATA-002.

V01-WORK-DATA-002 proceeds with the task-file helper model minimum:

- file-private helper model parsing / validation / TypeRef resolution for same-file references in task files
- main node / private helper visibility rules needed for task-local helper models
- task-file helper model render exposure in DAG Markdown `## Private models` as required by V01-ADR-071

Option B is deferred. Model-file helper semantics, model-file render exposure, V01-ADR-072 catalog follow-up, V01-ADR-075 resolution, and UC-002 model response helper-shape migration are owned by V01-WORK-DATA-003 or later.

## Deferred UC-002 Migration Shapes

The following model response candidates from V01-INV-DATA-002 are valid only after model-file helper render exposure is resolved outside V01-WORK-DATA-002:

| candidate | V01-INV-DATA-002 ID | file / field | deferred owner | dependency notes |
|---|---|---|---|---|
| `get_source_response.snippet` | N-023 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_response.yaml` / `get_source_response.snippet` | V01-WORK-DATA-003 or later | Requires V01-ADR-075 or equivalent model-file render exposure. |
| `get_reference_tree_response.nodes` | N-014 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml` / `get_reference_tree_response.nodes` | V01-WORK-DATA-003 or later | Requires V01-ADR-075 or equivalent model-file render exposure. |
| `get_reference_tree_response.edges` | N-015 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml` / `get_reference_tree_response.edges` | V01-WORK-DATA-003 or later | Requires V01-ADR-075 or equivalent model-file render exposure; existing `reference` model reuse may need separate care. |

Avoid as first migration even after the model-file boundary is resolved:

- `analyze_impact_response.impacts` / N-005: large mixed shape with enum and possible tagged-union pressure
- `list_endpoints_response.tables` / N-029: deeply nested response shape
- `get_signature_response.signature` / N-021: tagged union candidate
- `inspect_response.signature` / N-026: tagged union candidate
- `inspect_response.members` / N-027: tagged union / object-kind-specific payload candidate
- identity-related ObjectRef shapes: V01-ADR-078 / identity semantics boundary

## Done Condition

- V01-ADR-070 / V01-ADR-071 / V01-ADR-072 / V01-ADR-075 handling order is recorded.
- V01-ADR-075 proposed status and V01-ADR-073 dependency risk are classified.
- Option A task-file helper minimum is selected.
- UC-002 model response migration candidates are explicitly deferred.
- Scope exclusions for V01-ADR-073, V01-ADR-074, V01-ADR-078, and M15 reopening remain explicit.
- V01-TASK-DATA-002-02 input is a concrete spec-alignment scope, not a vague follow-up note.

## Verification

- V01-REQ-DATA-002 and V01-WORK-DATA-002 boundaries are narrowed to Option A task-file helper minimum.
- V01-REQ-DATA-001 and V01-WORK-DATA-001 remain untouched.
- M15 / `v1.1.0-spec` is not reopened.
- V01-ADR-073 / V01-ADR-074 / V01-ADR-078 are treated only as deferred or excluded scope.
- No implementation, renderer, fixture, or YAML migration is performed in this task.

## Evidence

- V01-ADR-070 accepted file-private helper model semantics.
- V01-ADR-071 accepted task-file helper model render exposure.
- V01-ADR-072 is a catalog view and is not required for the Option A task-file minimum.
- V01-ADR-075 remains proposed and depends on V01-ADR-073, so model-file helper render is not implementation-ready for V01-WORK-DATA-002.
- V01-INV-DATA-002 model response helper-shape candidates require model-file helper render exposure and are deferred.
- V01-WORK-DATA-002 now owns only the Option A task-file helper model minimum.
- V01-WORK-DATA-003 receives V01-ADR-072 catalog follow-up, V01-ADR-075 model-file render resolution, model-file helper render exposure, and UC-002 model response helper-shape migration.

## V01-TASK-DATA-002-02 Input

Align specs for the task-file helper model minimum:

- task files may contain file-private helper model definitions
- task-file helper models are private to the defining YAML file
- same-file TypeRef resolution may resolve helper model IDs
- cross-file / cross-module references to task-file helper models are invalid
- DAG Markdown may render `## Private models` for task files containing helper models
- helper models are schema definitions and must not be rendered as Mermaid DAG flow nodes
- diagnostics cover duplicate helper IDs and invalid helper references at the spec level
