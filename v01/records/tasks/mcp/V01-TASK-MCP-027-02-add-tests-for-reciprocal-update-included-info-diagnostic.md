# V01-TASK-MCP-027-02: add tests for reciprocal_update_included info diagnostic

- **id**: V01-TASK-MCP-027-02
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-027
- **source_requirement**: V01-REQ-MCP-030
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-027-01
- **outputs**:
  - internal/designrecords/authoring_test.go

## Goal

task create (default) と work_item create (default) で `reciprocal_update_included` info diagnostic が emit されることをテストで検証する。

## Work

- `TestProposeRecordCreateReciprocalUpdateIncludedDiagnostic_Task`: TASK-MCP-001-new を作成、info diagnostic のカテゴリ・severity・RecordID・Field・Value を検証
- `TestProposeRecordCreateReciprocalUpdateIncludedDiagnostic_WorkItem`: WORK-MCP-new を作成、info diagnostic を検証
- report_required_follow_up モードでは info diagnostic が emit されないことも確認

## Done condition

全テスト PASS

## Verification

`go test ./internal/designrecords/...`

## Evidence
2026-06-07: `TestProposeRecordCreateReciprocalUpdateIncludedDiagnostic` を3サブテストで追加。task_create_default_mode_emits_info_diagnostic / work_item_create_default_mode_emits_info_diagnostic / report_required_follow_up_mode_does_not_emit_included_diagnostic の全件 PASS 確認。全スイート回帰なし。
