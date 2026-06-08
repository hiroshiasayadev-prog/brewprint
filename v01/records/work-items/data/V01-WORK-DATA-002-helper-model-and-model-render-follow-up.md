# V01-WORK-DATA-002: Implement task-file helper model minimum

- **id**: V01-WORK-DATA-002
- **status**: done
- **date**: 2026-05-31
- **source_requirement**: V01-REQ-DATA-002
- **impact_refs**:
  - V01-INV-DATA-001
  - V01-INV-DATA-002
  - V01-REQ-DATA-001
  - V01-WORK-DATA-001
  - V01-WORK-DATA-003
  - V01-ADR-070
  - V01-ADR-071
- **tasks**:
  - V01-TASK-DATA-002-01
  - V01-TASK-DATA-002-02
  - V01-TASK-DATA-002-03

## Goal

Implement the first executable helper-model capability as the Option A task-file helper model minimum without reopening M15 / `v1.1.0-spec`.

This work item owns only the V01-ADR-070 / V01-ADR-071 path needed for task files: file-private helper models, same-file TypeRef resolution, validation for the task-file minimum, and DAG Markdown render exposure through `## Private models`.

## Boundary

### Included

- Apply V01-ADR-070 file-private helper model semantics to task files only.
- Apply V01-ADR-071 task-file helper model render exposure in DAG Markdown.
- Align the relevant specs for task-file helper model placement, visibility, same-file TypeRef resolution, diagnostics, and render format.
- Implement and validate task-file helper model parsing / resolution / diagnostics.
- Implement and verify DAG Markdown `## Private models` exposure for task files containing helper models.
- Keep helper models out of Mermaid DAG flow nodes.

### Deferred to V01-WORK-DATA-003 or Later

- V01-ADR-072 model / schema catalog follow-up.
- V01-ADR-075 model file render acceptance, revision, or split.
- Model-file helper model migration and model-file helper render exposure.
- UC-002 model response helper-shape migration, including `get_source_response.snippet`, `get_reference_tree_response.nodes`, and `get_reference_tree_response.edges`.
- MCP helper model exposure schema.

### Excluded

- V01-ADR-073 tagged union model.
- V01-ADR-074 DAG asset TypeRef hint.
- V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP / state identity work.
- M15 / `v1.1.0-spec` reopening.
- V01-REQ-DATA-001 / V01-WORK-DATA-001 edits.
- Treating all UC-002 notes retreat debt as one required migration.

## Impact Scope

| layer | current state | handling in this work item |
|---|---|---|
| source requirement | `V01-REQ-DATA-002` captured | This work item owns only the task-file helper model minimum |
| investigation evidence | `V01-INV-DATA-001` / `V01-INV-DATA-002` concluded | Use as boundary evidence, not as task status |
| previous data work | `V01-WORK-DATA-001` done | Preserve the M15 F1 close boundary |
| decision | V01-ADR-070 / V01-ADR-071 accepted | Use as the implementation basis for task-file helper models |
| deferred decisions | V01-ADR-072 accepted; V01-ADR-075 proposed | Send catalog, model-file render, and UC-002 model response migration to `V01-WORK-DATA-003` or later |
| spec | task-file helper model rules require alignment | `V01-TASK-DATA-002-02` owns spec alignment |
| implementation / validation | implemented and verified for task-file helper minimum | `V01-TASK-DATA-002-03` completed implementation, renderer exposure, and verification |
| UC-002 YAML | helper-shape notes remain | No UC-002 model response YAML migration in this work item |

## Task Flow

```mermaid
flowchart TD
  T1["V01-TASK-DATA-002-01: Option A boundary review and handoff"]
  T2["V01-TASK-DATA-002-02: Task-file helper minimum spec alignment"]
  T3["V01-TASK-DATA-002-03: Task-file helper minimum implementation and validation"]
  T1 --> T2 --> T3
```

## Tasks

- `V01-TASK-DATA-002-01`: Complete. Select Option A and record deferred model-file / catalog / UC-002 migration scope.
- `V01-TASK-DATA-002-02`: Align specs for task-file helper model placement, visibility, TypeRef resolution, diagnostics, and DAG Markdown render format.
- `V01-TASK-DATA-002-03`: Implement and verify task-file helper parsing / validation / TypeRef resolution and DAG Markdown `## Private models` render exposure.

## Follow-up Work

`V01-WORK-DATA-003` owns the model-file side of the helper-model chain. It receives V01-ADR-072 catalog follow-up, V01-ADR-075 model file render resolution, model-file helper render exposure, and UC-002 model response helper-shape migration candidates.

## Completion Condition

This work item can be marked `done` when the Option A task-file helper model minimum is reflected in the relevant specs, implementation, validation behavior, DAG Markdown render output, and verification evidence.

Completion does not require V01-ADR-072 catalog work, V01-ADR-075 model-file render work, model-file helper migration, UC-002 model response YAML migration, V01-ADR-073 tagged union work, V01-ADR-074 DAG TypeRef hint work, V01-ADR-078 MCP identity work, or any M15 / `v1.1.0-spec` reopening.

## Close Outcome

V01-WORK-DATA-002 is closed as `done` for the Option A task-file helper model minimum.

Closed scope:

- V01-TASK-DATA-002-01 selected and recorded the Option A boundary.
- V01-TASK-DATA-002-02 aligned the relevant specs for task-file helper placement, visibility, same-file TypeRef resolution, diagnostics, DAG Markdown `## Private models`, and ER exclusion.
- V01-TASK-DATA-002-03 implemented and verified task-file helper model parsing / indexing, same-file helper TypeRef resolution, external QualifiedID rejection by unresolved public lookup, `duplicate_model_id` diagnostics, DAG Markdown private model table rendering, and ER exclusion.
- V01-TASK-DATA-002-03 post-review corrections tightened helper model local identity collisions, rejected task-file `model main: true`, and kept private helper synthetic IDs out of public query reference targets.

Verification evidence:

- Close evidence commit: e1a70dc
- `gofmt -w ...`: completed for modified Go files.
- `go test ./internal/resolve ./internal/render/dag ./internal/render/er ./internal/query ./internal/mcp`: pass.
- `go test ./...`: pass.
- `validate_records(kind="task")`: pass.
- `validate_records(kind="work_item")`: pass.
- `git diff --check`: pass.

Deferred scope remains owned by V01-WORK-DATA-003 or later:

- V01-ADR-072 model / schema catalog follow-up.
- V01-ADR-075 model-file render resolution.
- Model-file helper render exposure.
- UC-002 model response helper-shape migration.
- V01-ADR-073 tagged union, V01-ADR-074 DAG TypeRef hint, V01-ADR-078 MCP helper exposure / semantic identity, and M15 / `v1.1.0-spec` reopening.
