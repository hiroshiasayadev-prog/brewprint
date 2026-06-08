# V01-TASK-MCP-022-03: update spec schema.md status vocabulary for task and work_item

- **id**: V01-TASK-MCP-022-03
- **status**: done
- **date**: 2026-06-06
- **work_item**: V01-WORK-MCP-022
- **source_requirement**: V01-REQ-MCP-026
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-022-02
- **outputs**:
  - docs/spec/design-records-mcp/schema.md updated

## Goal

`docs/spec/design-records-mcp/schema.md` の status テーブルおよび関連記述を V01-ADR-094 の決定内容に合わせて更新する。

## Work

- § status テーブルの `task` / `work_item` 行を `not_started` / `in_progress` / `blocked` / `done` に更新
- schema.md 内で旧 status tokens を参照している箇所をすべて確認・更新

## Done condition

- schema.md の `task` / `work_item` status が `not_started` / `in_progress` / `blocked` / `done` になっている
- 旧トークン（`todo` / `doing` / `decision_pending` / `*_pending`）への参照が残っていない

## Verification

- `validate_records` が schema 変更後も pass する（Go 実装変更前なので validation エラーは expected だが、spec 記述として正しいことを確認）

## Evidence

`docs/spec/design-records-mcp/schema.md` を以下の通り更新した。

変更箇所:
- § status テーブル: `work_item` 行を `not_started` / `in_progress` / `blocked` / `done` に変更（旧: 10種のステージ特化トークン）
- § status テーブル: `task` 行を `not_started` / `in_progress` / `blocked` / `done` に変更（旧: `todo` / `doing` / `blocked` / `done`）
- § status の `> 由来` コメントに V01-ADR-094 を追記

旧トークン（`todo` / `doing` / `decision_pending` / `*_pending`）への参照は schema.md 内に他になし（grep 確認済み）。
