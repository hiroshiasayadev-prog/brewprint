# V01-TASK-SELFHOST-001-02: Generated render review and canonical fixture update

- **id**: V01-TASK-SELFHOST-001-02
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-SELFHOST-001
- **source_requirement**: V01-REQ-SELFHOST-001
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-SELFHOST-001-01
- **outputs**:
  - UC-002 Phase A generated render review result
  - canonical fixture acceptance decision
  - stale coverage / TASKS-UC-002 update list for V01-TASK-SELFHOST-001-03

## Goal

Commit `8617d9c` で生成済みの UC-002 Phase A canonical render を review し、fixture として採用してよいかを判定する。

## Work

- `git status --short` で unrelated untracked files を確認する。
- `git show --stat 8617d9c` で canonical render commit の対象を確認する。
- Review 対象 render files と `render_index.yaml` / Phase A scope の対応を確認する。
- 8 MCP tool DAG render が Phase A MCP tool task と対応しているか確認する。
- DAG の obvious breakage を確認する。
- `_cross/er.md` の空 `erDiagram` が Phase A の context store / contract schema model 中心の前提と矛盾しないか確認する。
- 一時 render を生成し、canonical render と比較する。
- `coverage.md` / `TASKS-UC-002.md` に反映すべき stale 記述を V01-TASK-SELFHOST-001-03 へ送る。

## Done condition

- Review 対象 render files が確認されている。
- Canonical render と一時 render の比較結果が evidence に記録されている。
- Blocking issue の有無が分類されている。
- Canonical fixture として commit 済みにしてよいかが判定されている。
- V01-TASK-SELFHOST-001-03 に送る stale docs 更新事項が列挙されている。

## Verification

- `git status --short`
- `git show --stat 8617d9c`
- `go run ./cmd/brewprint render --yaml-root docs\uc\002-brewprint-self-hosting\yaml --out $env:TEMP\brewprint-uc002-render-review --clean`
- Byte-level comparison between `docs\uc\002-brewprint-self-hosting\renders` and `$env:TEMP\brewprint-uc002-render-review`
- Design Records MCP `validate_records(kind="task")`
- Design Records MCP `validate_records(kind="work_item")`

## Evidence

- 2026-05-31: Initial `git status --short` showed unrelated untracked files outside this task scope:
  - `?? docs/requirements/mcp/REQ-MCP-005-project-authoring-guidance-retrieval-support.md`
  - `?? docs/tasks/data/TASK-DATA-002-01-helper-model-render-boundary-review.md`
  - `?? docs/tasks/mcp/TASK-MCP-005-01-guidance-catalog-and-boundary-decision.md`
  - `?? docs/work-items/mcp/WORK-MCP-005-project-authoring-guidance-retrieval-support.md`
- 2026-05-31: `git show --stat 8617d9c` confirmed that the commit added V01-TASK-SELFHOST-001-01 evidence, 11 UC-002 render files, and a small V01-WORK-SELFHOST-001 update.
- Reviewed render files:
  - `docs/uc/002-brewprint-self-hosting/renders/index.md`
  - `docs/uc/002-brewprint-self-hosting/renders/mcp/index.md`
  - `docs/uc/002-brewprint-self-hosting/renders/mcp/dag-analyze_impact.md`
  - `docs/uc/002-brewprint-self-hosting/renders/mcp/dag-get_reference_tree.md`
  - `docs/uc/002-brewprint-self-hosting/renders/mcp/dag-get_references.md`
  - `docs/uc/002-brewprint-self-hosting/renders/mcp/dag-get_signature.md`
  - `docs/uc/002-brewprint-self-hosting/renders/mcp/dag-get_source.md`
  - `docs/uc/002-brewprint-self-hosting/renders/mcp/dag-inspect.md`
  - `docs/uc/002-brewprint-self-hosting/renders/mcp/dag-list_endpoints.md`
  - `docs/uc/002-brewprint-self-hosting/renders/mcp/dag-list_objects.md`
  - `docs/uc/002-brewprint-self-hosting/renders/_cross/er.md`
- Render count / scope review:
  - `render_index.yaml` defines one group, `mcp`, for module `mcp`.
  - `renders/index.md` reports 8 DAG renders and one cross ER render.
  - `renders/mcp/index.md` lists the 8 Phase A MCP tool DAG renders.
  - The 11 files match the Phase A render output set: project index, MCP group index, 8 tool DAGs, and cross ER.
- DAG review:
  - Each DAG has the expected main task section for its MCP tool.
  - Each DAG uses the expected public contract flow: `request -> validate_request -> query_service -> build_response -> response`.
  - Each `query_service` reads `resolved_project_store`.
  - No missing main task, missing flow step, invalid edge, or obvious private sub task exposure issue was found.
  - `returns.source` / initialized source rendering is not exercised by these Phase A tool DAGs; no expression collapse was observed.
- ER review:
  - `_cross/er.md` contains an empty `erDiagram`.
  - This is consistent with the recorded Phase A premise that `resolved_project_store` is a context store and MCP request / response / common models are contract schema models, while the v1 ER renderer draws DB-store-derived entities.
  - Therefore the empty ER render is a non-issue for Phase A and should not be treated as a new spec gap.
- Temp render comparison:
  - `go run ./cmd/brewprint render --yaml-root docs\uc\002-brewprint-self-hosting\yaml --out $env:TEMP\brewprint-uc002-render-review --clean` completed with `rendered 11 file(s)`.
  - The temp render file list matched canonical exactly.
  - Byte-level comparison found no missing files, extra files, or content diffs.
- Findings:
  - Blocking: none.
  - Minor: `coverage.md` and `TASKS-UC-002.md` are stale because they still describe Phase A render / `go test ./...` / render review as unexecuted or unconfirmed.
  - Non-issue: empty ER render is expected for Phase A.
- Canonical fixture decision:
  - The generated render output is acceptable as the UC-002 Phase A canonical fixture baseline.
  - No UC-002 YAML or renderer implementation change is required by this review.
- V01-TASK-SELFHOST-001-03 stale docs update list:
  - Update `docs/uc/002-brewprint-self-hosting/docs/coverage.md` §1 to state that `go test ./...`, `brewprint render`, and generated render review have been executed.
  - Update `docs/uc/002-brewprint-self-hosting/docs/coverage.md` §5 render coverage statuses from `未生成 / 未確認` to current reviewed status for project index, MCP group index, DAG renders, and ER render.
  - Update `docs/uc/002-brewprint-self-hosting/docs/coverage.md` §8 execution results with the V01-TASK-SELFHOST-001-01 and V01-TASK-SELFHOST-001-02 evidence.
  - Update `docs/uc/002-brewprint-self-hosting/docs/coverage.md` §9 next confirmation points so it no longer asks to do already completed render/test/review work.
  - Update `docs/uc/002-brewprint-self-hosting/TASKS-UC-002.md` Phase A render checkbox to done and replace the stale "未実行" note.
  - Update `docs/uc/002-brewprint-self-hosting/TASKS-UC-002.md` completion condition status only for Phase A render confirmation; do not close M14 self-hosting or mark Phase B items done.
  - Consider adding a short note that V01-WORK-SELFHOST-001 remains `fixture_pending` until coverage / TASKS-UC-002 synchronization and close evidence are completed.
- V01-WORK-SELFHOST-001 status:
  - Keep `fixture_pending`.
  - Do not mark `done` until V01-TASK-SELFHOST-001-03 synchronizes coverage / TASKS-UC-002 and records close evidence.
- 2026-05-31: Design Records MCP `validate_records(kind="work_item")` returned `ok: true` with no diagnostics.
- 2026-05-31: Design Records MCP `validate_records(kind="task")` returned `ok: false` because of unrelated untracked workflow artifacts outside this task scope:
  - `V01-TASK-DATA-002-01`: `workflow_relation_mismatch` because `V01-WORK-DATA-002.tasks` does not contain `V01-TASK-DATA-002-01`.
  - `V01-TASK-MCP-005-01`: invalid status `not_started`.
  - `V01-TASK-MCP-005-01`: `workflow_relation_mismatch` because `V01-WORK-MCP-005.tasks` does not contain `V01-TASK-MCP-005-01`.
  - No diagnostic was reported for `V01-TASK-SELFHOST-001-02` or `V01-WORK-SELFHOST-001`.
