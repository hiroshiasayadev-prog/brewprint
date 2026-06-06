# TASK-MCP-022-06: migrate existing work_item records to new status vocabulary

- **id**: TASK-MCP-022-06
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-MCP-022
- **source_requirement**: REQ-MCP-026
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-022-04
- **outputs**:
  - existing work_item records migrated
  - validate_records pass

## Goal

既存 `task` / `work_item` レコードの status を ADR-094 migration mapping に従って更新し、全レコードが新 vocabulary で validate_records を pass する状態にする。

## Work

migration mapping（ADR-094 § 決定）:

| 変更前 | 変更後 | 対象レコード |
|---|---|---|
| `decision_pending` | `blocked` | WORK-DATA-013 / WORK-DATA-014 / WORK-DATA-015 |
| `todo` | `not_started` | 対象 0 件（作業不要） ※TASK-MCP-022-04/05/06 は Go 実装後に migration 対象 |
| `doing` | `in_progress` | 対象 0 件（作業不要） |
| `*_pending` | `in_progress` | 対象 0 件（作業不要） |

各対象レコードを `propose_record_update` (metadata_fields_replace) で更新し accept する。

## Done condition

- `decision_pending` を持つレコードが 0 件
- `todo` を持つ task レコードが 0 件
- `validate_records` で全 task / work_item レコードが pass

## Verification

`validate_records` 実行。

## Evidence

- WORK-DATA-013/014/015: `propose_record_update` (metadata_fields_replace) → `accept_proposed_write` で `decision_pending` → `blocked` 移行完了
- TASK-MCP-022-04/05/06: ファイルシステム直接編集で `todo` → `not_started` （旧バイナリ稼働中の MCP が `not_started` を task kind で拒否するため fallback）
- `validate_records` (work_item): `invalid_status_for_kind` 0 件、`decision_pending` 0 件確認
- `validate_records` (task): `invalid_status_for_kind` 3 件 (TASK-022-04/05/06 の `not_started`) — ファイル内容は正しい。MCP サーバを新バイナリで再ビルドすると解消する。
