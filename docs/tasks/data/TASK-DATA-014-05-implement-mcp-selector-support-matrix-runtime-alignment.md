# TASK-DATA-014-05: Implement MCP selector support matrix runtime alignment

- **id**: TASK-DATA-014-05
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-DATA-014
- **source_requirement**: REQ-DATA-007
- **estimate**: 1d-2d
- **depends_on**:
  - TASK-DATA-014-04
- **outputs**:
  - MCP selector support matrix runtime enforcement
  - Aligned unsupported selector tool errors and diagnostics
  - Aligned file selector reference aggregation
  - Tests and smoke evidence for selector matrix behavior

## Goal

Implement the follow-up required by `TASK-DATA-014-04` so the MCP runtime / Go implementation aligns with the accepted selector support matrix and object-dependent vocabulary contract from `WORK-DATA-014`.

## Work

- Add explicit selector support classification for MCP tools where current behavior conflates unsupported-but-resolvable selectors with `not_found`.
- Align unsupported selector behavior for `get_signature`, `get_references`, `get_reference_tree`, and `inspect` with `unsupported_object` for matrix `no` cells.
- Preserve the `analyze_impact` exception: unsupported selectors should return a normal response with empty impact data and `unsupported_selector` diagnostic, not a tool error.
- Align `field` selector kind alias handling so `kind: field` and `kind: model_field` are accepted consistently where the shared schema allows both.
- Align `list_objects` object/kind filter validation so unknown object or kind returns `invalid_args` instead of silently returning empty results.
- Decide and implement the accepted behavior for `list_objects(kind: model_field)` or explicitly document a narrowed behavior if implementation should only list `kind: field`.
- Implement `get_references(file: node)` limited behavior so references whose source or target is a node defined in that file are aggregated.
- Implement `get_references(file: state_file)` behavior so references owned by the state file include transitions and resolvable `transition_from` / `transition_event` / `transition_to` / `transition_action` references.
- Align `get_reference_tree(file: node)` limited behavior with `get_references(file: node)`.
- Add or update unit/runtime tests covering the above behavior.

## Done condition

- `TASK-DATA-014-04` major findings are either fixed or explicitly split into smaller follow-up tasks with rationale.
- Supported selectors still pass for representative node/view/transition/field/asset/private-sub-node cases.
- Unsupported selectors in `get_signature`, `get_references`, `get_reference_tree`, and `inspect` return the accepted `unsupported_object` behavior when selector resolution succeeds but the tool does not support the object/kind.
- `analyze_impact` unsupported selector behavior remains normal response plus `unsupported_selector` diagnostic.
- `field` / `model_field` selector kind alias behavior is consistent.
- `list_objects` rejects unknown object/kind with `invalid_args`.
- File selector reference aggregation is covered by tests.
- Evidence records tests, runtime smoke results, and any intentional deviations.

## Verification

- Run targeted tests for selector support behavior.
- Run `go test ./internal/query ./internal/mcp`.
- Run `go test ./internal/designrecords ./internal/designrecordsmcp` if workflow artifact changes are made.
- Run a runtime smoke or MCP boundary test demonstrating key selector matrix behavior.
- Run Design Records MCP validation for this task and `WORK-DATA-014` after evidence is updated.

## Evidence
Verdict: PASS.

Files changed:

- `internal/query/selector_support.go`
- `internal/query/service.go`
- `internal/query/signature.go`
- `internal/query/references.go`
- `internal/query/reference_tree.go`
- `internal/query/inspect.go`
- `internal/query/list_objects.go`
- `internal/query/analyze_impact.go`
- `internal/query/analyze_impact_render.go`
- `internal/resolve/references.go`
- `internal/mcp/server.go`
- `internal/query/service_test.go`
- `internal/mcp/server_test.go`

Behavior fixed:

- Unsupported-but-resolvable selectors now return `unsupported_object` for the covered MCP boundary cases, including `get_signature(file: state_file)` and `inspect(primitive)`, while unresolved file selectors remain `not_found`.
- `analyze_impact(view: api_table / sequence matrix no)` behavior remains a normal response with empty impacts and `unsupported_selector` diagnostic; the field alias condition was widened without changing the unsupported-selector exception path.
- `field` selectors accept `kind: field` and `kind: model_field` consistently for full ID and `id + local_id`; returned ObjectRef remains `object: field`, `kind: field`.
- `list_objects` validates unknown `object` / `kind` as `invalid_args`; `kind: model_field` is treated as the narrow alias of `kind: field`.
- `get_references(file: node)` now aggregates direct references for nodes defined in the file and excludes raw flow-wiring `consumes_asset` references.
- `get_references(file: state_file)` now includes transition-owned `transition_from`, `transition_event`, `transition_to`, and `transition_action` references in addition to existing state-file references.
- `get_reference_tree(file: node)` uses the same limited node-file aggregation at the root and then continues existing bounded traversal.

Review repair summary:

- Updated `analyzeImpactFieldSelector` to treat `kind: model_field` as a field selector via `isFieldKind`, aligning coverage metadata with field impact collection.
- Added an `analyze_impact` model-field alias test using `kind: model_field` without `object: field`, including coverage for `flow_param_field_resolution`.
- Added MCP boundary tests for unsupported-but-resolvable `get_references(file: api_table)` and `get_reference_tree(view: api_table)`, both expecting `unsupported_object`.

Tests run:

- `git status --short`: repo still has pre-existing unrelated docs/spec/work item changes, untracked DATA/product files, and `tmp.py`; review repair changes are limited to `internal/query/analyze_impact.go`, `internal/query/service_test.go`, `internal/mcp/server_test.go`, plus this task evidence update.
- `go test ./internal/query ./internal/mcp`: PASS.
- `go test ./internal/designrecords ./internal/designrecordsmcp`: PASS.
- `go test ./internal/mcp -run "TestServerCallTool" -v`: PASS.

Runtime smoke / MCP boundary result:

- Covered MCP boundary behavior includes `get_signature(file: state_file) -> unsupported_object`, `inspect(primitive) -> unsupported_object`, `get_references(file: api_table) -> unsupported_object`, `get_reference_tree(view: api_table) -> unsupported_object`, unresolved file selector -> `not_found`, `list_objects(kind: unknown_kind) -> invalid_args`, `list_objects(kind: model_field)` returning field objects, `get_references(file: node)` limited aggregation, `get_references(file: state_file)` transition-owned references, and `get_reference_tree(file: node)` limited aggregation.

Remaining known gaps:

- No remaining review gaps from the `TASK-DATA-014-05` review findings.
- No intentional deviations from the accepted `WORK-DATA-014` selector matrix contract were left in this task.
- `WORK-DATA-014` status was not changed by this task evidence update; close/status synchronization can be done separately after validation/review if desired.
