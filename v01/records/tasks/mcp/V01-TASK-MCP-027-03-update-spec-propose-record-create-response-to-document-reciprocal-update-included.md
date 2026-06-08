# V01-TASK-MCP-027-03: update SPEC propose_record_create response to document reciprocal_update_included

- **id**: V01-TASK-MCP-027-03
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-027
- **source_requirement**: V01-REQ-MCP-030
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-027-02
- **outputs**:
  - docs/spec/design-records-mcp/tools.md

## Goal

SPEC の `propose_record_create` response セクションに `reciprocal_update_included` info diagnostic の記述を追加する。

## Work

- `reciprocal_update_mode` の表に `include_required` 時の info diagnostic の説明を追加
- Response セクションに `reciprocal_update_included` diagnostic の example を追加

## Done condition

SPEC が V01-REQ-MCP-030 の Required Outcome を網羅している

## Verification

目視確認

## Evidence
2026-06-07: SPEC `propose_record_create` の `reciprocal_update_mode` 表に `include_required` 時の info diagnostic 説明と「unsafe ケースは blocking/report-only」注記を追加。Response セクションに `reciprocal_update_included` ダイアグノスティックの動作説明と JSON example を追加。
