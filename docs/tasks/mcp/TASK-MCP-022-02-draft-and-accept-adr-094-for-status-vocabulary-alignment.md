# TASK-MCP-022-02: draft and accept ADR-094 for status vocabulary alignment

- **id**: TASK-MCP-022-02
- **status**: done
- **date**: 2026-06-06
- **work_item**: WORK-MCP-022
- **source_requirement**: REQ-MCP-026
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-022-01
- **outputs**:
  - ADR-094 accepted

## Goal

`task` / `work_item` status vocabulary の統一に関する設計判断を ADR-094 として記録し、accepted にする。
ADR-092 が定めた work_item status vocabulary を refine する。

## Work

- ADR-094 を `propose_record_create` で起票（status: `proposed`）
- 内容: 背景（LLM recall 乖離の evidence）、決定（4トークン統一）、理由、却下した代替案、影響範囲
- ユーザーレビュー後 accepted に更新

## Done condition

- ADR-094 が `accepted` になっている

## Verification

- `get_record("ADR-094")` で status = `accepted` を確認

## Evidence

ADR-094 を `proposed` で起票し、`metadata_fields_replace` で `accepted` に更新した。

- `get_record("ADR-094")` で status = `accepted` を確認済み
- depends_on: ADR-092, ADR-089
- migration mapping テーブル、却下代替案（全 kind 統一案・ステージトークン維持案・task のみ案）を記録
