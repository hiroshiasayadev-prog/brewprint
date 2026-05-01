# Go M13 analyze_impact implementation summary

- **status**: closed / hybrid v1 close
- **last_updated**: 2026-05-01
- **repo**: `C:\Users\imved\projects\brewprint`
- **verified**: `gofmt -w (Get-ChildItem -Recurse -Filter *.go -File | ForEach-Object FullName)` completed successfully; `go test ./...` passed after final M13 gap-fix changes and passed again on the final cached verification run.

---

## 1. Scope

This document summarizes the current Go implementation of MCP `analyze_impact` as observed in the implementation files.

Main implementation files reviewed:

- `internal/query/types.go`
- `internal/query/analyze_impact.go`
- `internal/query/analyze_impact_validation.go`
- `internal/query/analyze_impact_task.go`
- `internal/query/analyze_impact_field.go`
- `internal/query/analyze_impact_render.go`
- `internal/query/analyze_impact_add.go`
- `internal/query/reference_tree.go`
- `internal/query/references.go`
- `internal/query/service.go`
- `internal/query/service_test.go`
- `internal/mcp/server.go`
- `internal/mcp/server_test.go`
- `internal/mcp/jsonrpc_test.go`

Spec / task files reviewed:

- `docs/spec/mcp/tools/analyze-impact.md`
- `docs/spec/mcp/tools/get-reference-tree.md`
- `docs/spec/mcp/schema.md`
- `docs/spec/mcp/errors.md`
- `docs/spec/mcp/overview.md`
- `docs/spec/mcp/versioning.md`
- `docs/tasks/m13-mcp-analyze-impact-implementation.md`

---

## 2. Implemented analyze_impact change kinds

The current implementation validates and accepts the following `change.kind` values:

- `rename`
- `remove`
- `change_type`
- `change_contract`
- `change_transition_target`
- `add`

Validation behavior currently implemented:

- `rename` requires `new_id`.
- `remove` accepts no extra payload.
- `change_type` requires `new_type`.
- `change_contract` accepts optional `note`.
- `change_transition_target` requires at least one of `new_to` or `new_action`.
- `add` requires `added_id`.
- Unexpected payload keys are rejected as `invalid_change_payload`.

MCP wrapper maps invalid change payload errors to `invalid_change_payload` tool errors.

---

## 3. Collector coverage implemented

### 3.1 Orchestration

`Service.AnalyzeImpact` currently:

- validates `change` payloads;
- applies default `max_impacts = 200`;
- returns unsupported selectors as normal responses, not tool errors;
- resolves supported selectors with `referenceTarget`;
- dispatches to target/change-specific collectors;
- appends render output impacts for supported render-related changes;
- sorts impacts deterministically;
- truncates by `max_impacts` with `truncated_reasons=["max_impacts"]`;
- assigns deterministic IDs such as `impact-001`;
- summarizes impacts by severity, fixability, and kind;
- always returns `coverage`, `assumptions`, `truncated`, `truncated_reasons`, and `diagnostics`.

### 3.2 Task collector

For `node: task`, the current collector handles:

- `rename`
- `remove`
- `change_contract`

Collected impact kinds:

- `transition_action`
  - incoming `transition_action` references where the task is used as a transition action;
- `flow_step_task`
  - flow `step` references;
  - flow `foreach` task references;
  - task references inside `fork` branches;
  - task references inside `branch` cases;
- `sequence_step_action`
  - sequence scenario steps whose resolved transition action task is the target task.

Typical severity / fixability:

- `rename`: `breaking` + `suggested`, with `replace_reference` suggested fixes when source line/column is available. A shared mechanical judgement gate now exists, but task rename currently keeps `ReferenceStable=false`, so it remains conservative unless the gate can prove all requirements.
- `remove`: `breaking` + `manual_review`.
- `change_contract`: `warning` + `manual_review`.

### 3.3 Field collector

For `field` selectors, the current collector handles:

- `rename`
- `remove`
- `change_type`

Collector inputs:

- direct incoming references from `project.ReferencesByTarget`;
- `GetReferenceTree(direction="in", depth=2)` traversal results;
- v1-minimal flow param wiring scan.

Collected impact kinds:

- `field_consumer`
- `flow_param_field`

Current type analysis for `change_type`:

- Compares type identity at a shallow level.
- Model IDs are treated as model identities.
- Primitive names are treated as primitive identities.
- Model-to-model changes are downgraded to `warning`; other kind changes are `breaking`.

Current fixability:

- `rename`: `breaking` + `unknown` for file-only field consumer source fallback; `flow_param_field` uses `manual_review`.
- `remove`: `breaking` + `manual_review`.
- `change_type`: `breaking` or `warning` + `manual_review`.

### 3.4 Transition collector

For `transition` selectors, the current collector handles:

- `change_transition_target`
- `rename`
- `remove`

Collected impact kinds:

- `transition_target_resolution`
  - `change_transition_target.new_to` / `new_action` が既存 state / task に解決できない場合に返す;
- `transition_scenario_step`
  - sequence scenario steps that exact-match the transition;
- `transition_action_task`
  - the action task attached to the transition, when present.

Typical severity / fixability:

- unresolved `change_transition_target`: `breaking` + `manual_review`.
- resolved `change_transition_target`: contextual impacts are `warning` + `manual_review`.
- `rename`: `breaking` + `manual_review`.
- `remove`: `breaking` + `manual_review`.

### 3.5 Render output collector

The render collector handles these change kinds:

- `rename`
- `remove`
- `change_type`
- `change_contract`
- `change_transition_target`

It maps changed objects to render output file paths at file granularity.

Currently handled target categories:

- task → DAG group output and API cross-view output for endpoint tasks;
- model → dependent task DAG outputs and store / ER outputs;
- field → model-derived render outputs and ER outputs;
- transition → state diagram output and sequence diagram outputs that include the transition.

Collected impact kind:

- `render_output`

Recommended action asks the user to rerun `brewprint render` and inspect the affected render output file.

Typical severity / fixability:

- Most render impacts are `info`.
- `remove` render impacts are `warning`.
- If source line/column is unavailable, fixability is downgraded to `unknown`.

### 3.6 Add collector

For `add`, the current collector performs name collision checks for:

- nodes by qualified ID;
- model fields by full field ID;
- transitions by transition synthetic ID;
- assets by asset synthetic ID.

Collected impact kind:

- `name_collision`

If no collision exists and the `added_id` kind cannot be inferred, the collector returns an `unsupported_selector` diagnostic.

Current `add` impacts are `breaking` + `manual_review` for collisions.

---

## 4. Unsupported selector handling

Unsupported selectors are not returned as tool errors.

Current behavior:

- response is returned normally;
- `impacts` is an empty array;
- `summary` is the zero summary;
- `coverage.analyzed` is empty;
- `coverage.not_analyzed` starts with `unsupported_selector`;
- `diagnostics` contains one `unsupported_selector` warning.

Selectors currently treated as unsupported by `analyze_impact` include:

- `primitive`;
- `asset`;
- `file` / `state_file`;
- `view` selectors including API table, ER diagram, and sequence diagram;
- `render_index` kind.

This matches the public behavior expected by the spec: unsupported selector is a normal response, not a tool error.

---

## 5. Source location fallback

The current implementation uses source locations as follows:

- Task collector:
  - transition action source tries to locate the `action` line;
  - flow source tries to locate `step` / `foreach` task lines;
  - sequence source tries to locate the scenario step block.
- Transition collector:
  - scenario step source tries to locate the sequence step block;
  - action task source uses the transition block.
- Field collector:
  - current impacts generally fall back to source file only.
- Render collector:
  - uses the changed object source block when available;
  - may fall back to file-only source.
- Add collision collector:
  - returns source file when the collided object has a file.

When a source file is known but line/column is unavailable, the current implementation may:

- downgrade fixability to `unknown`;
- omit suggested fixes;
- add a `source_location_unavailable` warning diagnostic.

Known source-location limitation:

- Some impact paths only know `source.file` and do not know line/column.
- Field impacts currently use file-only fallback, so field rename fixability is `unknown`, not `mechanical`.

---

## 6. coverage.not_analyzed

The current default `coverage.not_analyzed` includes:

- `type_structural_compatibility`
- `semantic_contract_compatibility`
- `render_presentation_details`
- `wireframe_element_binding`

Unsupported selectors prepend:

- `unsupported_selector`

Current default assumptions include:

- `rename後のID衝突は検証対象外`
- `note内の自然言語参照は解析対象外`
- `semantic contract compatibility は解析対象外`

For `add`, `coverage.analyzed` is currently:

- `name_collision`

For `add`, `coverage.not_analyzed` additionally includes:

- `type_resolution`
- `writer_coverage`

Implementation note: the add collector intentionally implements name collision checks only in M13 v1. Full type resolution / writer coverage checks are documented limitations rather than claimed coverage.

---

## 7. Known limitations / closeout blockers

M13 is closed as a hybrid v1 close.

Closeout state:

1. The final gap-fix implementation added:
   - shared mechanical judgement gate;
   - v1-minimal `flow_param_field_resolution` collector;
   - `change_transition_target.new_to` / `new_action` resolution impact;
   - `add` coverage correction: only `name_collision` is analyzed, while `type_resolution` / `writer_coverage` are not analyzed;
   - field rename remains conservative when source line/column is missing.
2. Remaining known limitations are intentional M13 v1 limitations and are documented in `docs/spec/mcp/tools/analyze-impact.md`.
3. `gofmt` and `go test ./...` passed after the final gap-fix changes.

No public response shape / enum / diagnostic code change is proposed by this summary.
No spec public contract change was made.

---

## 8. Tests

Required closeout command:

```powershell
go test ./...
```

Result:

```powershell
PS C:\Users\imved\projects\brewprint> gofmt -w (Get-ChildItem -Recurse -Filter *.go -File | ForEach-Object FullName)
PS C:\Users\imved\projects\brewprint> go test ./...
ok      github.com/hiroshiasayadev-prog/brewprint/cmd/brewprint (cached)
ok      github.com/hiroshiasayadev-prog/brewprint/internal/mcp  (cached)
ok      github.com/hiroshiasayadev-prog/brewprint/internal/query        (cached)
?       github.com/hiroshiasayadev-prog/brewprint/internal/rawyaml      [no test files]
ok      github.com/hiroshiasayadev-prog/brewprint/internal/render/api   (cached)
ok      github.com/hiroshiasayadev-prog/brewprint/internal/render/dag   (cached)
ok      github.com/hiroshiasayadev-prog/brewprint/internal/render/er    (cached)
ok      github.com/hiroshiasayadev-prog/brewprint/internal/render/placement     (cached)
ok      github.com/hiroshiasayadev-prog/brewprint/internal/render/project       (cached)
ok      github.com/hiroshiasayadev-prog/brewprint/internal/render/sequence      (cached)
ok      github.com/hiroshiasayadev-prog/brewprint/internal/render/state (cached)
ok      github.com/hiroshiasayadev-prog/brewprint/internal/render/wireframe     (cached)
ok      github.com/hiroshiasayadev-prog/brewprint/internal/resolve      (cached)
?       github.com/hiroshiasayadev-prog/brewprint/internal/semantic     [no test files]
?       github.com/hiroshiasayadev-prog/brewprint/internal/source       [no test files]
?       github.com/hiroshiasayadev-prog/brewprint/internal/testutil/golden      [no test files]
```

Interpretation:

- `gofmt` passed when invoked with explicit Go file paths.
- `go test ./...` passed after the final M13 gap-fix changes and passed again on the final cached verification run.
- M13 is closed as hybrid v1 close.

---

## 9. Closeout note

M13 was closed as a hybrid v1 close after the final gap-fix changes and successful verification.

Recommended commit scope:

- analyze_impact collector gap fixes
- M13 v1 spec constraints
- M13 closeout docs
