# TASK-MCP-026-04: smoke test — add multi-op operations array cases to tmp.py

- **id**: TASK-MCP-026-04
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-MCP-026
- **source_requirement**: REQ-MCP-025
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-026-03
- **outputs**:
  - Updated tmp.py

## Goal

`tmp.py` に multi-op update の smoke ケースを追加し、実際の MCP server 経由で動作を確認する。

## Work

`tmp.py` に以下の smoke ケースを追加する:
- `operations: [metadata_fields_replace, named_section_replace]` で Evidence + status:done を 1 proposal で適用
- `accept_proposed_write` して `written: true` を確認
- 既存の全 smoke ケースが通過することを確認

## Done condition

`python tmp.py` を実行して全ケース green。

## Verification

smoke スクリプトの出力で全ケース通過を確認する。

## Evidence
2026-06-07: スモークテスト完了。

`tmp.py` に [6]-[9] の4ケースを追加し `python tmp.py` で全9件 PASS を確認:

| Case | 内容 | 結果 |
|---|---|---|
| [6] | operations array (metadata_fields_replace + named_section_replace) → proposal_created:true, diff に両変更を含む | PASS |
| [7] | 2 named_section_replace → `multiple_section_replace_not_supported` | PASS |
| [8] | update + operations 両方 → `invalid_request` | PASS |
| [9] | 同一フィールドへの 2 metadata op → `conflicting_operations` | PASS |

合わせて `internal/designrecordsmcp/tools.go` の `propose_record_update` スキーマを更新:
- `operations` プロパティを追加（array of UpdateOperation）
- `required` から `update` を除外（`update` / `operations` どちらでも可）
- description を更新
- `operationsSchema()` ヘルパーを新設

`go test ./internal/designrecordsmcp/...` 全件 PASS 確認済み。
