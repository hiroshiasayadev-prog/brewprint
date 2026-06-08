# V01-TASK-DATA-003-03: Model-file render implementation and verification

- **id**: V01-TASK-DATA-003-03
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-DATA-003
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 1d-3d
- **depends_on**:
  - V01-TASK-DATA-003-02
- **outputs**:
  - Model-file render minimum implementation
  - Model render fixture / golden coverage
  - Verification evidence for model-file render placement and index participation

## Goal

Implement and verify the model-file render minimum defined by `docs/spec/views/model-file.md` and `docs/spec/project-layout.md`.

This task turns the completed V01-TASK-DATA-003-02 spec alignment into renderer behavior and fixture / golden evidence for model files that use `struct`, `enum`, `list`, and `dict` model kinds.

## Work

- Review the implementation-entry notes in V01-TASK-DATA-003-02.
- Locate the existing render pipeline for DAG / state / sequence / wireframe outputs and project / group index generation.
- Add model-file render generation for `model/*.yaml` files.
- Place model-file render outputs under the owning group as `model-{model-id}.md`.
- Include model-file renders in master index counts as the `Model` column.
- Include model-file renders in group index rows with kind `Model`.
- Render the public main model section for supported model kinds:
  - `struct`
  - `enum`
  - `list`
  - `dict`
- Render same-file private helper models in a `Private models` section when present.
- Add or update fixture / golden coverage for model-file render output and index participation.
- Run the relevant tests / renderer verification commands and record evidence.

## Included Scope

- Renderer implementation for the model-file render minimum.
- Placement under `renders/{group-id}/model-{model-id}.md`.
- Master index and group index participation for `Model` renders.
- Markdown table output for `struct`, `enum`, `list`, and `dict`.
- Same-file private helper model visibility in model-file render.
- Fixture / golden update required to prove the above behavior.
- Verification commands needed to prove no unintended renderer breakage.

## Excluded Scope

- Tagged union discriminator / variants rendering.
- Tagged union validation or YAML migration.
- DAG asset TypeRef hint.
- MCP helper model exposure / semantic identity.
- V01-REQ-DATA-003 private helper model signature exposure policy.
- UC-002 model response helper migration.
- V01-WORK-DATA-002 reopening.
- V01-ADR-075 text or status changes.
- V01-ADR-073 implementation.

## Done Condition

- Model YAML files produce model-file Markdown renders according to `docs/spec/views/model-file.md`.
- Model-file renders are placed according to `docs/spec/project-layout.md`.
- Master index and group index include model-file renders as kind `Model`.
- Public main model rendering works for `struct`, `enum`, `list`, and `dict`.
- Same-file private helper models are visible in model-file render output.
- Fixture / golden coverage exists for the implemented behavior.
- Relevant renderer / test commands pass, or any failures are classified as unrelated / follow-up with evidence.
- Excluded scope remains unimplemented in this task.

## Verification

Recommended verification commands for the implementation agent:

```powershell
# inspect working tree before changes
git status --short

# run targeted tests first if available
go test ./internal/...

# run broader project tests if targeted package boundary is unclear
go test ./...

# inspect final diff
git diff -- docs/spec docs/uc internal cmd

git diff --check
```

The exact command set may be adjusted by the implementation agent after locating the renderer package and existing fixture / golden test conventions.

## Evidence

Implemented and verified on 2026-05-31.

Implementation:

- Added model-file Markdown rendering for public model files under `internal/render/model`.
- Connected model renders to project render generation as `renders/{group-id}/model-{model-id}.md`.
- Added model-file helper classification for `model/*.yaml` files that declare a `main: true` public model; same-file non-main models are file-private helpers.
- Updated master index generation to include the `Model` column before `DAG`.
- Updated group index ordering to `Model` -> `DAG` -> `State` -> `Sequence` -> `Wireframe`.
- Kept tagged union rendering / validation, V01-ADR-073 implementation, DAG TypeRef hints, MCP helper exposure, V01-REQ-DATA-003 policy, UC-002 migration, and V01-WORK-DATA-004 fixes out of this task.

Fixture / golden:

- Updated UC-001 model YAML files to declare their public model with `main: true`.
- Added same-file private helper models to `docs/uc/001-ec-checkout-flow/yaml/auth/model/login_form.yaml` covering:
  - `struct`
  - `enum`
  - `list`
  - `dict`
- Regenerated UC-001 render golden files with:

```powershell
go run ./cmd/brewprint render --yaml-root docs/uc/001-ec-checkout-flow/yaml --out docs/uc/001-ec-checkout-flow/renders --clean
```

Result:

```text
rendered 37 file(s)
```

Golden evidence includes:

- `docs/uc/001-ec-checkout-flow/renders/auth/model-login_form.md`
- `docs/uc/001-ec-checkout-flow/renders/auth/model-request_context.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/model-cart_item_list.md`
- `docs/uc/001-ec-checkout-flow/renders/index.md`
- `docs/uc/001-ec-checkout-flow/renders/auth/index.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/index.md`
- `docs/uc/001-ec-checkout-flow/renders/catalog/index.md`

Verification:

```powershell
go test ./internal/render/model ./internal/resolve
go test ./internal/...
go test ./cmd/...
go test ./...
git diff --check
```

Results:

- `go test ./internal/render/model ./internal/resolve`: passed after golden generation.
- `go test ./internal/...`: passed.
- `go test ./cmd/...`: passed.
- `go test ./...`: passed.
- `git diff --check`: passed; output contained only existing CRLF conversion warnings from Git.
- `validate_records(kind=task)`: passed.
- `validate_records(kind=work_item)`: failed only on the known unrelated `V01-WORK-DATA-004` invalid status `todo` error.
