# V01-TASK-MCP-022-02: draft and accept V01-ADR-094 for status vocabulary alignment

- **id**: V01-TASK-MCP-022-02
- **status**: done
- **date**: 2026-06-06
- **work_item**: V01-WORK-MCP-022
- **source_requirement**: V01-REQ-MCP-026
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-022-01
- **outputs**:
  - V01-ADR-094 accepted

## Goal

`task` / `work_item` status vocabulary の統一に関する設計判断を V01-ADR-094 として記録し、accepted にする。
V01-ADR-092 が定めた work_item status vocabulary を refine する。

## Work

- V01-ADR-094 を `propose_record_create` で起票（status: `proposed`）
- 内容: 背景（LLM recall 乖離の evidence）、決定（4トークン統一）、理由、却下した代替案、影響範囲
- ユーザーレビュー後 accepted に更新

## Done condition

- V01-ADR-094 が `accepted` になっている

## Verification

- `get_record("V01-ADR-094")` で status = `accepted` を確認

## Evidence

V01-ADR-094 を `proposed` で起票し、`metadata_fields_replace` で `accepted` に更新した。

- `get_record("V01-ADR-094")` で status = `accepted` を確認済み
- depends_on: V01-ADR-092, V01-ADR-089
- migration mapping テーブル、却下代替案（全 kind 統一案・ステージトークン維持案・task のみ案）を記録
