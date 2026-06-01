# TASK-MCP-006-01: Current workflow metadata validation gap を整理する

- **id**: TASK-MCP-006-01
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-MCP-006
- **source_requirement**: REQ-MCP-006
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Current workflow metadata validation gap table
  - Strictness decision input for TASK-MCP-006-02

## Goal

WORK-MCP-006 の判断前提として、Design Records MCP の workflow artifact metadata validation について、現行 public contract / authoring guidance / implementation / tests / runtime behavior の差分を evidence として整理する。

## Scope

- 対象 artifact kind は `requirement` / `work_item` / `task` とする。
- `fixture_pending` は valid work item status として扱い、不正値扱いしない。
- REQ-MCP-003 / WORK-MCP-003 の完了済み scope は再オープンしない。
- SELFHOST artifact は観測対象にしてよいが、直接編集対象にしない。
- Orphan diagnostics、progress projection、workflow traversal、task dependency cycle / execution order projection、physical path relation support、`req:` / `work:` / `task:` semantic prefix support は対象外とする。

## Work

1. `docs/spec/design-records-mcp/schema.md` と `docs/spec/design-records-mcp/tools.md` から workflow artifact metadata / diagnostics / `validate_records` contract を確認する。
2. `docs/guides/requirement-authoring.md`、`docs/guides/work-item-authoring.md`、`docs/guides/task-authoring.md` の required metadata wording を確認する。
3. `internal/designrecords` と `internal/designrecordsmcp` の parser / validator / MCP response tests を確認する。
4. `requirement` / `work_item` / `task` ごとに、required field presence、empty scalar、missing list、empty list、invalid date、invalid status、invalid relation target の現行挙動を整理する。
5. 必要なら `validate_records` の runtime evidence を取得する。
6. TASK-MCP-006-02 で判断すべき strictness choice と unresolved points を列挙する。

## Expected output

- 現行 contract / authoring guidance / implementation / tests / runtime behavior の gap table。
- required metadata / empty value / date validation / diagnostic category に関する判断候補。
- TASK-MCP-006-02 に渡す decision input。

## Completion condition

TASK-MCP-006-02 の public contract decision に進めるだけの evidence が揃い、未判断点が明示されている。

## Verification

- 2026-06-01: Codex review により、TASK-MCP-006-02 に進めるだけの evidence review が揃っていることを確認した。
- Targeted packages passed: `go test ./internal/designrecords ./internal/designrecordsmcp`。
- Runtime observation では `validate_records(kind="requirement")`、`validate_records(kind="work_item")`、`validate_records(kind="task")` が `ok:true` を返した。

## Evidence

2026-06-01 review result: OK to proceed.

Reviewed files:

- `AGENTS.md`
- `docs/prompt_chappy.md`
- `docs/doc-policy.md`
- `docs/requirements/mcp/REQ-MCP-006-workflow-artifact-metadata-validation-strictness.md`
- `docs/work-items/mcp/WORK-MCP-006-workflow-artifact-metadata-validation-strictness.md`
- `docs/tasks/mcp/TASK-MCP-006-01-current-workflow-metadata-validation-gap-review.md`
- `docs/spec/design-records-mcp/schema.md`
- `docs/spec/design-records-mcp/tools.md`
- `docs/guides/requirement-authoring.md`
- `docs/guides/work-item-authoring.md`
- `docs/guides/task-authoring.md`
- `docs/requirements/README.md`
- `docs/work-items/README.md`
- `docs/tasks/README.md`
- `docs/adr/092-design-records-mcp-workflow-artifact-record-and-relation-boundary.md`
- `internal/designrecords`
- `internal/designrecordsmcp`

Current contract summary:

| kind | required metadata |
|---|---|
| `requirement` | `id`, `status`, `date`, `source_refs`, `work_items` |
| `work_item` | `id`, `status`, `date`, `source_requirement`, `impact_refs`, `tasks` |
| `task` | `id`, `status`, `date`, `work_item`, `source_requirement`, `estimate`, `depends_on`, `outputs` |

Current implementation behavior:

- Parser reads workflow id, status, relation/list/detail fields, but workflow metadata structs have no `Date` and no presence flags.
- `date` is ignored in requirement / work item / task parser branches.
- Missing list fields and empty list fields normalize to empty slices.
- Empty list items are dropped.
- Empty relation scalar values are skipped by relation validation.
- Missing or empty status becomes `invalid_status_for_kind`.
- Missing or empty id generally becomes `filename_id_mismatch`, not missing metadata diagnostic.
- Invalid non-empty relation target form becomes `invalid_workflow_relation_target`.
- Missing non-empty supported relation target becomes `unresolved_workflow_relation`.
- Reciprocal relation mismatch is validated for requirement work items, work item source requirement, work item tasks, task work item, and task source requirement consistency.

Current test coverage:

Covered:

- Workflow parser happy path for requirement / work item / task.
- Invalid workflow H1 / ID / metadata-H1-filename mismatch.
- Invalid status for requirement / work item / task.
- Workflow relation unresolved target.
- Invalid relation target kind / form.
- Bidirectional relation mismatch.
- Task source requirement mismatch.
- Task `depends_on` empty list accepted.
- MCP `validate_records` returns diagnostics as normal tool response.

Not covered:

- Missing required metadata presence for all workflow kinds.
- Invalid workflow date.
- Missing date.
- Empty scalar strictness except status/id side effects.
- Missing list vs empty list distinction.
- Empty required list strictness.
- Explicit `fixture_pending` regression as valid work item status.
- MCP/runtime tests for new metadata strictness categories.

Gap table:

| concern | contract/guidance | current behavior | gap |
|---|---|---|---|
| required field presence | required fields listed | no general presence validation | major |
| missing date | required | ignored / no diagnostic | major |
| invalid date | `YYYY-MM-DD` shown | ignored / no diagnostic | major |
| empty status | required enum | `invalid_status_for_kind` | covered indirectly |
| invalid status | enum | `invalid_status_for_kind` | covered |
| `fixture_pending` | valid work_item status | allowed in code | OK; add regression |
| empty `source_requirement` / `work_item` | required relation scalar | skipped / no diagnostic | major |
| missing `source_requirement` / `work_item` | required | default empty / no diagnostic | major |
| missing `estimate` | required task scalar | default empty / no diagnostic | major |
| empty `estimate` | required | no diagnostic | major |
| missing list field | required lists | default empty slice / no diagnostic | major |
| empty list field | required lists; task `depends_on` explicitly may be empty | no diagnostic for all lists | needs decision |
| invalid relation target | ID-as-ref kind-specific | diagnostic exists | covered |
| unresolved relation target | supported ID missing | diagnostic exists | covered |
| diagnostic category | none for metadata strictness | no category exists | decision needed |

Findings:

1. Required metadata is contracted but not validated. Required workflow fields are documented, but implementation validates only ID consistency, status enum, and relation integrity. `validate_records` can return OK for artifacts missing fields the public schema calls required.
2. Workflow `date` is required but operationally invisible. Parser discards it and validator cannot detect missing or invalid dates.
3. Missing vs empty lists cannot currently be distinguished. Missing field, empty field, and empty list all collapse to empty slices without parser presence tracking.

Non-issues:

- `fixture_pending` is valid work item status and should not be treated as a bug.
- `task.depends_on: []` is explicitly valid for workflow relation diagnostics.
- Orphan diagnostics, traversal, progress projection, dependency cycle/order projection, physical path relations, and `req:` / `work:` / `task:` prefixes remain out of scope.
- REQ-MCP-003 / WORK-MCP-003 completion scope does not need reopening.

TASK-MCP-006-02 decision points:

1. Add new diagnostic categories or reuse existing ones for missing / empty metadata.
2. Decide whether missing scalar and empty scalar are the same diagnostic.
3. Decide whether required list presence and empty list are distinct.
4. Decide which required lists may be empty, especially `source_refs`, `impact_refs`, `outputs`, `work_items`, and `tasks`.
5. Decide whether workflow date gets strict `YYYY-MM-DD` validation.
6. Decide whether `diagnostics:null` vs `[]` is in scope.
7. Add explicit `fixture_pending` valid regression test.

File candidates for follow-up:

- Specs: `docs/spec/design-records-mcp/schema.md`, `docs/spec/design-records-mcp/tools.md`
- Guidance: `docs/guides/requirement-authoring.md`, `docs/guides/work-item-authoring.md`, `docs/guides/task-authoring.md`, possibly matching README files
- Implementation: `internal/designrecords/parser.go`, `internal/designrecords/types.go`, `internal/designrecords/validation.go`
- Tests: `internal/designrecords/parser_index_test.go`, `internal/designrecords/validation_test.go`, `internal/designrecords/types_test.go`, `internal/designrecordsmcp/tools_call_test.go`
