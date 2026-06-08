# V01-TASK-DATA-002-03: Task-file helper minimum implementation and validation

- **id**: V01-TASK-DATA-002-03
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-DATA-002
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 1d-2d
- **depends_on**:
  - V01-TASK-DATA-002-02
- **outputs**:
  - Task-file helper model implementation
  - Task-file helper TypeRef validation
  - DAG Markdown `## Private models` render exposure
  - Verification evidence for the Option A minimum

## Goal

Implement and verify the selected V01-WORK-DATA-002 Option A boundary: task-file helper model minimum.

This task starts only after V01-TASK-DATA-002-02 has aligned the relevant specs.

## Work

- Implement task-file helper model parsing / loading for file-private schema definitions.
- Implement same-file TypeRef resolution for task-file helper models.
- Reject cross-file / cross-module references to task-file helper models.
- Validate duplicate local helper IDs within a task file.
- Validate helper/public name conflicts to the extent required by the V01-TASK-DATA-002-02 spec alignment.
- Implement DAG Markdown `## Private models` render exposure for task files containing helper models.
- Keep helper models out of Mermaid DAG flow nodes.
- Add or update focused tests / fixtures / golden evidence for the Option A minimum.

## Included Scope

- Task-file helper model parsing and validation.
- Same-file helper TypeRef resolution.
- Invalid external helper reference diagnostics.
- DAG Markdown private model render exposure for task files.
- Verification evidence for parser / resolver / diagnostics / render behavior.

## Excluded Scope

- Model-file helper model migration.
- Model-file render exposure.
- Model catalog render changes.
- V01-ADR-075 acceptance / revision / split.
- V01-ADR-073 tagged union model.
- V01-ADR-074 DAG asset TypeRef hint.
- V01-ADR-078 MCP helper model exposure / semantic identity.
- UC-002 model response migration candidates N-014 / N-015 / N-023.
- M15 / `v1.1.0-spec` reopening.

## Done Condition

- Task-file helper models can be represented according to the aligned specs.
- Same-file TypeRef resolution works for task-file helper models.
- External references to task-file helper models are rejected.
- Duplicate helper IDs and required helper/public conflicts are diagnosed.
- DAG Markdown renders `## Private models` for task-file helper models.
- Helper models are not emitted as Mermaid DAG flow nodes.
- Focused tests / fixtures / golden evidence cover the implemented behavior.
- V01-WORK-DATA-003 scope remains untouched.

## Verification

- Confirm V01-TASK-DATA-002-02 spec alignment is the implementation source.
- Confirm model-file helper render remains owned by V01-WORK-DATA-003.
- Confirm V01-ADR-072 / V01-ADR-073 / V01-ADR-074 / V01-ADR-075 / V01-ADR-078 are not pulled into V01-WORK-DATA-002 implementation.
- Confirm no UC-002 model response YAML migration is performed.
- Run the focused validation / render test suite identified during implementation.

## Evidence

Initial evidence:

- V01-TASK-DATA-002-01 selected Option A.
- V01-TASK-DATA-002-02 owns the required spec alignment.
- V01-ADR-070 defines file-private helper model semantics.
- V01-ADR-071 defines task-file helper model render exposure.
- V01-WORK-DATA-003 owns model-file helper render and UC-002 model response migration.

Completion evidence:

- Added task-file private helper model indexing through `PrivateModelsByFile`; task-file helper models use file-local internal identity and are not registered as public `ModelsByQID`.
- Added same-file scoped TypeRef resolution for task / control params, returns, and helper model `fields[].type` / `element` / `value`, with same-file helper model lookup before public model fallback.
- Added `duplicate_model_id` diagnostics for same-file helper model duplicate IDs and same-module public/helper model ID collisions.
- Kept QualifiedID TypeRefs public-only; external references to task-file helper models remain unresolved.
- Added DAG Markdown `## Private models` rendering for both simple and flow DAG render paths; helper models remain out of the Mermaid body.
- Kept ER render behavior public/store-driven; task-file helper models do not appear in ER output.
- Added focused tests:
  - `internal/resolve/helper_model_test.go`
  - `internal/render/dag/private_models_test.go`
  - `internal/render/er/private_models_test.go`
  - `internal/query/service_test.go`
- Implementation evidence commit: e1a70dc
- Updated existing tests that used task-file model fixtures as public model fixtures so public model assumptions now live in `model/` files instead of task files.
- Review correction:
  - Added `duplicate_model_id` validation for task-file helper model IDs that collide with same-file main / private sub node local IDs.
  - Rejected `main: true` on task-file `type: model` nodes with `semantic_validation` instead of silently treating them as private helpers.
  - Kept private helper models out of public query reference targets and displayed raw/local TypeRef names in public task / asset signatures.
  - Kept private helper models out of `list_objects` and file inspect node listings; MCP private helper exposure schema remains deferred.
- Verification:
  - `gofmt -w ...`: completed for modified Go files.
  - `go test ./internal/resolve ./internal/render/dag ./internal/render/er ./internal/query ./internal/mcp`: pass.
  - `go test ./...`: pass.
  - `validate_records(kind="task")`: pass.
  - `validate_records(kind="work_item")`: pass.
  - `git diff --check`: pass.
- Scope guard:
  - No UC-002 YAML migration was performed.
  - `docs/uc/003-task-file-helper-model/` remains excluded from this commit candidate.
  - V01-ADR-072 model/schema catalog, V01-ADR-075 model file render, V01-ADR-073 tagged union, V01-ADR-074 DAG TypeRef hint, V01-ADR-078 MCP identity / helper exposure, and M15 / `v1.1.0-spec` reopening remain outside this task.
