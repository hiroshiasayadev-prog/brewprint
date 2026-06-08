# V01-WORK-MCP-027: emit reciprocal_update_included info diagnostic on default include_required create

- **id**: V01-WORK-MCP-027
- **status**: done
- **date**: 2026-06-07
- **source_requirement**: V01-REQ-MCP-030
- **impact_refs**:
  - docs/spec/design-records-mcp/tools.md
  - internal/designrecords/types.go
  - internal/designrecords/authoring.go
  - internal/designrecords/authoring_test.go
- **tasks**:
  - V01-TASK-MCP-027-01
  - V01-TASK-MCP-027-02
  - V01-TASK-MCP-027-03
  - V01-TASK-MCP-027-04

## Goal

`propose_record_create` のデフォルト動作 (`reciprocal_update_mode: include_required`) で、親レコードへの相互更新が提案 diff に含まれた際に `reciprocal_update_included` info diagnostic を emit する。

現状は `include_required` 時に reciprocal ファイルが diff に追加されるが、diagnostic が emit されない。V01-REQ-MCP-030 はこの info diagnostic を必須としている。

## Boundary

- `DiagnosticReciprocalUpdateIncluded` カテゴリの追加のみ。既存の `report_required_follow_up` 挙動は変更しない。
- SPEC の `propose_record_create` response セクション更新を含む。
- 実装対象: task create (parent work item `tasks` append) と work item create (source requirement `work_items` append) の両パス。

## Evidence
2026-06-07: 全タスク完了。

- V01-TASK-MCP-027-01: `DiagnosticReciprocalUpdateIncluded` を types.go に追加。`requiredReciprocalUpdates` の `include_required` パス（WorkItem・Task）に `reciprocalUpdateIncludedDiagnostic` を追加。
- V01-TASK-MCP-027-02: `TestProposeRecordCreateReciprocalUpdateIncludedDiagnostic` 3サブテスト追加・全件 PASS。全スイート回帰なし。
- V01-TASK-MCP-027-03: SPEC `propose_record_create` 表・レスポンスセクションを更新。
- V01-TASK-MCP-027-04: クローズ同期完了。V01-REQ-MCP-030 accepted、V01-WORK-MCP-027 done。

`go test ./internal/designrecords/... ./internal/designrecordsmcp/...` 全件 PASS 確認。
