# TASK-DATA-012-03: Apply UC-002 enum and vocabulary YAML cleanup

- **id**: TASK-DATA-012-03
- **status**: done
- **date**: 2026-06-05
- **work_item**: WORK-DATA-012
- **source_requirement**: REQ-DATA-002
- **estimate**: 1d-2d
- **depends_on**:
  - TASK-DATA-012-02
- **outputs**:
  - UC-002 YAML enum and usage-site vocabulary cleanup for selected WORK-DATA-012 candidates
  - Validation / render / test evidence for selected YAML cleanup
  - Residual note-only or no-action confirmation for excluded candidates

## Goal

Apply the selected UC-002 enum and usage-site vocabulary YAML cleanup decisions from `TASK-DATA-012-02`.

This task owns only the YAML cleanup and verification for the selected WORK-DATA-012 candidates. It must not reopen closed DATA work, broaden enum minimum scope, or absorb behavior/default, selector matrix, recursive/untagged-union, tagged-union, or MCP identity work.

## Work

- Create or select named enum models for the selected candidates:
  - `mcp_error_code` for N-034;
  - `reference_kind` for N-045;
  - `reference_direction` for N-046;
  - `list_objects_object_filter` for N-030;
  - `reference_query_direction` for N-019;
  - `analyze_impact_coverage_vocabulary` for N-006 residual.
- Update selected UC-002 YAML usage sites:
  - `mcp_error.code` to use `mcp_error_code`;
  - `reference.kind` to use `reference_kind`;
  - `reference.direction` to use `reference_direction`;
  - `get_reference_tree_edge.kind` to use `reference_kind`;
  - `get_reference_tree_edge.direction` to use `reference_direction`;
  - `list_objects_request.object` to use `list_objects_object_filter`;
  - `get_references_response.direction` to use `reference_query_direction`;
  - `analyze_impact_coverage.analyzed` and `analyze_impact_coverage.not_analyzed` to use a typed list of `analyze_impact_coverage_vocabulary`, if supported by current model syntax;
  - `get_reference_tree_node.via` to use a typed list of `reference_kind`, if supported by current model syntax.
- Preserve `string_list` as a generic helper for open/free-text usage sites.
- Keep N-023 residual note-only: `get_source_snippet.language` remains `str` with literal `yaml` note.
- Keep N-029 residual no-action: do not create an HTTP method enum without a closed value set.
- If typed list-of-enum syntax is not supported by current model syntax, stop those list migrations and record the blocker / follow-up path instead of inventing syntax.
- Run the relevant validation, render, golden, and Go test commands needed for UC-002 YAML cleanup.

## Done condition

- Selected YAML fields are updated according to the `TASK-DATA-012-02` decision table, or explicitly blocked with evidence where the current model syntax cannot represent the selected form.
- No behavior/default, selector/support-matrix, recursive/untagged-union, tagged-union, or MCP identity scope is changed by this task.
- N-023 and N-029 remain note-only / no-action unless new evidence proves the previous decision invalid.
- UC-002 validation and relevant render/golden/test checks are run and recorded.
- The task evidence lists exact files changed, commands run, results, and any deferred follow-up.

## Verification

- Confirm only DATA-012-selected YAML cleanup and this task's evidence/status are changed, aside from generated render/golden outputs if required by the project's normal workflow.
- Confirm `WORK-DATA-001`, `WORK-DATA-006`, `WORK-DATA-009`, and `WORK-DATA-010` are not reopened or edited for scope changes.
- Confirm `REQ-DATA-006` / `WORK-DATA-013`, `REQ-DATA-007` / `WORK-DATA-014`, `REQ-DATA-008` / `WORK-DATA-015`, `REQ-DATA-004` / `WORK-DATA-010`, and `REQ-MCP-004` / `WORK-MCP-004` remain only routed references, not modified owner scopes.
- Run project validation / rendering commands for UC-002 and relevant Go tests.
- Run Design Records MCP validation for this task and `WORK-DATA-012` after updates.

## Evidence
Completed on 2026-06-05.

### Sources reviewed

- `docs/prompt_chappy.md`
- `AGENTS.md`
- `docs/AGENTS.md`
- `docs/doc-policy.md`
- `docs/tasks/data/TASK-DATA-012-03-apply-uc-002-enum-and-vocabulary-yaml-cleanup.md`
- `docs/tasks/data/TASK-DATA-012-02-decide-uc-002-enum-literal-and-vocabulary-cleanup-path.md`
- `docs/tasks/data/TASK-DATA-012-01-review-uc-002-enum-literal-and-vocabulary-cleanup-candidates.md`
- `docs/work-items/data/WORK-DATA-012-uc-002-enum-literal-vocabulary-cleanup.md`
- `docs/adr/067-enum-model.md`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/common.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/mcp_error.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/reference.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/string_list.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_request.yaml`

### Enum models created

| enum model ID | placement | values (count) |
|---|---|---|
| `reference_kind` | `common.yaml` | 18 values |
| `reference_direction` | `common.yaml` | 2 values: `out`, `in` |
| `mcp_error_code` | `mcp_error.yaml` | 12 values |
| `list_objects_object_filter` | `list_objects_request.yaml` | 4 values: `node`, `view`, `transition`, `field` |
| `reference_query_direction` | `get_references_response.yaml` | 3 values: `out`, `in`, `both` |
| `analyze_impact_coverage_vocabulary` | `analyze_impact_response.yaml` | 14 values |

Placement rationale: `reference_kind` and `reference_direction` placed in `common.yaml` because they are cross-file shared (used in `reference.yaml`, `get_reference_tree_response.yaml`, and `get_reference_tree_node.via`). All other enums are local to their respective model files, consistent with existing UC-002 model organization.

### Usage sites updated

| field path | before | after | candidate |
|---|---|---|---|
| `mcp_error.code` | `str` | `mcp_error_code` | N-034 |
| `reference.kind` | `str` | `reference_kind` | N-045 |
| `reference.direction` | `str` | `reference_direction` | N-046 |
| `get_reference_tree_edge.kind` | `str` | `reference_kind` | N-015 residual |
| `get_reference_tree_edge.direction` | `str` | `reference_direction` | N-015 residual |
| `list_objects_request.object` | `str` | `list_objects_object_filter` | N-030 |
| `get_references_response.direction` | `str` | `reference_query_direction` | N-019 |
| `analyze_impact_coverage.analyzed` | `string_list` | `list<analyze_impact_coverage_vocabulary>` | N-006 residual |
| `analyze_impact_coverage.not_analyzed` | `string_list` | `list<analyze_impact_coverage_vocabulary>` | N-006 residual |
| `get_reference_tree_node.via` | `string_list` | `list<reference_kind>` | N-051 |

### Typed list-of-enum migration

`list<reference_kind>` and `list<analyze_impact_coverage_vocabulary>` were applied without inventing new syntax. Rationale: ADR-067 §3 states enum models are treated as named model TypeRef (same namespace as struct models). ADR-060 includes `list<T>` as an inline TypeRef where T is any valid TypeRef. The existing `get_reference_tree_response.yaml` already uses `list<get_reference_tree_node>` and `list<get_reference_tree_edge>` (struct TypeRef in list). Applying the same inline list syntax with enum TypeRef is consistent. Validation passed with no diagnostics, confirming the syntax is accepted.

### Note-only / no-action confirmations

- **N-023 residual**: `get_source_snippet.language` remains `type: str`. Not modified. Literal `yaml` constraint remains note-only per TASK-DATA-012-02 decision.
- **N-029 residual**: No enum created for HTTP method. No-action confirmed; no closed HTTP method value set is defined by current spec evidence.
- **N-051 note**: `get_reference_tree_request.kinds` remains `string_list`. The `kinds` field is a reference kind filter (open/multi-select query input) and is out of scope for this task. Only the confirmed closed-vocabulary usage sites (`via`, `analyzed`, `not_analyzed`) were migrated.
- `string_list` remains unchanged as a generic helper.

### Commands run

| command | result |
|---|---|
| `go run ./cmd/brewprint validate --yaml-root ./docs/uc/002-brewprint-self-hosting` | `ok` — no diagnostics |
| `go run ./cmd/brewprint render --yaml-root ./docs/uc/002-brewprint-self-hosting --out ./docs/uc/002-brewprint-self-hosting/render` | `rendered 47 file(s)` with pre-existing `render_index` warning (unrelated to enum changes) |
| `go test ./...` | All packages except `internal/designrecords` and `internal/designrecordsmcp` passed. Failures in `TestProposeRecordUpdateRequiredHeadingCaseFallback`, `TestRequiredSectionHeadingCaseMismatchDiagnostics`, `TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields` are pre-existing failures about REQ-MCP-021 heading case canonicalization implementation, unrelated to UC-002 YAML enum changes. |

### Design Records MCP validation

- `validate_records` for `TASK-DATA-012-03`: `ok: true`, `diagnostics: null`
- `validate_records` for `WORK-DATA-012`: `ok: true`, `diagnostics: null`

### Scoped git diff summary

Changed files within TASK-DATA-012-03 scope:
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/common.yaml` — added `reference_kind` (18 values) and `reference_direction` (2 values) enums
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/mcp_error.yaml` — added `mcp_error_code` (12 values) enum; updated `mcp_error.code` from `str` to `mcp_error_code`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/reference.yaml` — updated `reference.kind` from `str` to `reference_kind`; updated `reference.direction` from `str` to `reference_direction`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_request.yaml` — added `list_objects_object_filter` (4 values) enum; updated `list_objects_request.object` from `str` to `list_objects_object_filter`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_response.yaml` — added `reference_query_direction` (3 values) enum; updated `get_references_response.direction` from `str` to `reference_query_direction`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml` — updated `get_reference_tree_edge.kind` from `str` to `reference_kind`; updated `get_reference_tree_edge.direction` from `str` to `reference_direction`; updated `get_reference_tree_node.via` from `string_list` to `list<reference_kind>`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml` — added `analyze_impact_coverage_vocabulary` (14 values) enum; updated `analyze_impact_coverage.analyzed` from `string_list` to `list<analyze_impact_coverage_vocabulary>`; updated `analyze_impact_coverage.not_analyzed` from `string_list` to `list<analyze_impact_coverage_vocabulary>`

Other dirty files in git status (`REQ-MCP-020`, `REQ-MCP-021`, `REQ-MCP-022`, `schema.md`, `tools.md`, `authoring_test.go`, `validation_test.go`, `tools_call_test.go`) are pre-existing unrelated MCP implementation work and are not part of this task scope.

### Out-of-scope routing confirmation

- `WORK-DATA-001`, `WORK-DATA-006`, `WORK-DATA-009`, `WORK-DATA-010` — not reopened or modified.
- `REQ-DATA-006` / `WORK-DATA-013` (behavior/default) — not touched.
- `REQ-DATA-007` / `WORK-DATA-014` (selector/support matrix) — not touched.
- `REQ-DATA-008` / `WORK-DATA-015` (recursive/untagged-union) — not touched.
- `REQ-DATA-004` / `WORK-DATA-010` (tagged union) — not touched.
- `REQ-MCP-004` / `WORK-MCP-004` (MCP identity) — not touched.
