# V01-TASK-MCP-022-05: update authoring guides and tools.md examples for new status vocabulary

- **id**: V01-TASK-MCP-022-05
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-022
- **source_requirement**: V01-REQ-MCP-026
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-022-03
- **outputs**:
  - docs/guides/task-authoring.md updated
  - docs/guides/work-item-authoring.md updated
  - docs/spec/design-records-mcp/tools.md examples updated

## Goal

authoring guides と `tools.md` の response/request examples を新 status vocabulary に更新する。
codex review Finding 1 (tools.md 旧トークン) と Finding 2 (authoring guides stale) を解消する。

## Work

- `docs/guides/task-authoring.md`: status テーブルを `not_started` / `in_progress` / `blocked` / `done` に更新（`todo` / `doing` を削除）
- `docs/guides/work-item-authoring.md`: status テーブルを `not_started` / `in_progress` / `blocked` / `done` に更新（ステージ特化トークン7種を削除）
- `docs/spec/design-records-mcp/tools.md`:
  - L115: `"status": "design_spec_pending"` → `"status": "in_progress"`
  - L1057: `"status": "todo"` → `"status": "not_started"`

## Done condition

- `task-authoring.md` に旧トークン (`todo` / `doing`) が残っていない
- `work-item-authoring.md` にステージ特化トークンが残っていない
- `tools.md` の例示に旧トークンが残っていない

## Verification

旧トークンの grep で 0 件を確認。

## Evidence

- `docs/guides/task-authoring.md`: Format 行と Status テーブルを `not_started` / `in_progress` / `blocked` / `done` に更新（`todo`/`doing` 削除）
- `docs/guides/work-item-authoring.md`: Format 行と Status テーブルを `not_started` / `in_progress` / `blocked` / `done` に更新（ステージ特化トークン 7 種削除）
- `docs/spec/design-records-mcp/tools.md`: L115 `design_spec_pending` → `in_progress`、L1057 `todo` → `not_started`
