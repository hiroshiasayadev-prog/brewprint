# TASK-SELFHOST-001-03: Coverage / gap sync and close evidence

- **id**: TASK-SELFHOST-001-03
- **status**: done
- **date**: 2026-05-31
- **work_item**: WORK-SELFHOST-001
- **source_requirement**: REQ-SELFHOST-001
- **estimate**: 0.5d
- **depends_on**:
  - TASK-SELFHOST-001-02
- **outputs**:
  - UC-002 coverage render / review status synchronization
  - UC-002 TASKS Phase A render gap log synchronization
  - WORK-SELFHOST-001 close evidence

## Goal

UC-002 Phase A validate / render / generated render review の完了 evidence を `coverage.md` と `TASKS-UC-002.md` に反映し、`WORK-SELFHOST-001` を close できる状態にする。

## Work

- `coverage.md` §1 の未確認記述を更新し、Phase A validate / render / review が完了済みであることを記録する。
- `coverage.md` §5 の render coverage status を生成・review済みに更新する。
- `coverage.md` §8 に `TASK-SELFHOST-001-01` / `TASK-SELFHOST-001-02` の実行 evidence を反映する。
- `coverage.md` §9 から完了済み確認項目を除き、残項目を `WORK-SELFHOST-001` 範囲外の未完項目に整理する。
- `TASKS-UC-002.md` の Phase A render checkbox を done 化する。
- `WORK-SELFHOST-001` metadata に本taskを追加し、status を `done` にする。
- `WORK-SELFHOST-001` に close outcome を記録する。

## Done condition

- UC-002 coverage が Phase A validate / render / generated render review 完了状態を示している。
- `TASKS-UC-002.md` の Phase A render item が完了済みになっている。
- `WORK-SELFHOST-001` が `done` になり、close outcome が記録されている。
- M14 self-hosting 本体は `paused` のまま、M14a / M15 は `closed` のまま維持されている。
- UC-002 YAML、renders、legacy M14 / M14a / M15 records、`docs/TASKS.md` は変更していない。

## Verification

- `git status --short`
- Design Records MCP `validate_records(kind="work_item")`
- Design Records MCP `validate_records(kind="task")`

`go test ./...` は実行しない。今回の変更は docs synchronization のみであり、TASK-SELFHOST-001-01 で pass 済みのため。

## Evidence

- 2026-05-31: Task opened as the third and final task under `WORK-SELFHOST-001`, then completed in the same docs synchronization pass.
- 2026-05-31: `coverage.md` §1 now records that Phase A YAML placement, `go test ./...`, validate, canonical render generation, and generated render review are complete.
- 2026-05-31: `coverage.md` §5 now marks project index, MCP group index, 8 DAG renders, and the expected empty ER render as generated / reviewed.
- 2026-05-31: `coverage.md` §8 now records evidence from `TASK-SELFHOST-001-01` and `TASK-SELFHOST-001-02`, including `rendered 11 file(s)` and byte-level canonical / temp render match.
- 2026-05-31: `coverage.md` §9 now removes completed render/test/review checks and leaves only out-of-scope remaining UC-002 work.
- 2026-05-31: `TASKS-UC-002.md` Phase A render checkbox was marked done with the TASK-SELFHOST evidence.
- 2026-05-31: `WORK-SELFHOST-001` was moved from the fixture review stage to `done`. The earlier `fixture_pending` state is retained only as historical pre-close context and is no longer used as the final status.
- 2026-05-31: M14 self-hosting remained `paused` in `docs/TASKS.md`; M14a and M15 remained `closed`. No legacy M14 / M14a / M15 record was edited.
- 2026-05-31: UC-002 YAML and canonical renders were not changed by this task.
- 2026-05-31: `git status --short` after docs synchronization showed this task's modified / new files:
  - `M docs/uc/002-brewprint-self-hosting/TASKS-UC-002.md`
  - `M docs/uc/002-brewprint-self-hosting/docs/coverage.md`
  - `M docs/work-items/self-hosting/WORK-SELFHOST-001-uc002-phase-a-canonical-render-and-coverage-verification.md`
  - `?? docs/tasks/self-hosting/TASK-SELFHOST-001-03-coverage-gap-sync-and-close-evidence.md`
- 2026-05-31: The same `git status --short` also showed unrelated tracked modifications outside this task scope:
  - `M docs/requirements/data/REQ-DATA-002-helper-model-and-model-render-follow-up.md`
  - `M docs/work-items/data/WORK-DATA-002-helper-model-and-model-render-follow-up.md`
- 2026-05-31: The same `git status --short` also showed one pre-existing related untracked artifact not created by this task:
  - `?? docs/tasks/self-hosting/TASK-SELFHOST-001-02-generated-render-review-and-canonical-fixture-update.md`
- 2026-05-31: The same `git status --short` also showed unrelated untracked artifacts outside this task scope:
  - `?? docs/requirements/mcp/REQ-MCP-005-project-authoring-guidance-retrieval-support.md`
  - `?? docs/requirements/mcp/REQ-MCP-006-workflow-artifact-metadata-validation-strictness.md`
  - `?? docs/tasks/data/TASK-DATA-002-01-helper-model-render-boundary-review.md`
  - `?? docs/tasks/data/TASK-DATA-002-02-task-file-helper-minimum-path-decision-and-spec-alignment-scope.md`
  - `?? docs/tasks/mcp/TASK-MCP-005-01-guidance-catalog-and-boundary-decision.md`
  - `?? docs/work-items/data/WORK-DATA-003-model-file-helper-render-boundary.md`
  - `?? docs/work-items/mcp/WORK-MCP-005-project-authoring-guidance-retrieval-support.md`
  - `?? docs/work-items/mcp/WORK-MCP-006-workflow-artifact-metadata-validation-strictness.md`
- 2026-05-31: `git diff --check` reported no whitespace errors. It emitted only Git line-ending warnings for edited Markdown files.
- 2026-05-31: Design Records MCP `validate_records(kind="work_item")` returned `ok: true` with no diagnostics.
- 2026-05-31: Design Records MCP `validate_records(kind="task")` returned `ok: false` because of unrelated diagnostics:
  - `TASK-MCP-005-01`: invalid task status `not_started`.
  - `TASK-MCP-005-01`: `workflow_relation_mismatch` because `WORK-MCP-005.tasks` does not contain `TASK-MCP-005-01`.
  - No diagnostic was reported for `TASK-SELFHOST-001-03`.
