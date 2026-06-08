# V01-WORK-MCP-026: atomic multi-operation record update proposal support

- **id**: V01-WORK-MCP-026
- **status**: done
- **date**: 2026-06-07
- **source_requirement**: V01-REQ-MCP-025
- **impact_refs**:
  - docs/spec/design-records-mcp/tools.md
  - internal/designrecords/types.go
  - internal/designrecords/authoring.go
  - internal/designrecords/authoring_test.go
  - tmp.py
- **tasks**:
  - V01-TASK-MCP-026-01
  - V01-TASK-MCP-026-02
  - V01-TASK-MCP-026-03
  - V01-TASK-MCP-026-04
  - V01-TASK-MCP-026-05

## Goal

`propose_record_update` を拡張し、`operations: [...]` 配列で複数の update operation を 1 retained proposal に原子的にまとめて適用できるようにする。Evidence + status:done を同一 proposal で完結させる高頻度ユースケースを解消する。

## Boundary

- 単一レコードに対する atomic multi-operation update のみ。マルチレコードトランザクションは対象外。
- API 形状: `operations: [...]` 配列を新設。既存の `update: {...}` 単体形式は後方互換維持。`update` と `operations` は排他。
- `named_section_replace` は 1 `operations` 配列につき 1 op のみサポート（複数は conflict error）。複数 named_section_replace のサポートは将来 REQ として別途起票。
- `body` / `body_cache_id` は各 op にインラインで持たせる（Y 形式）。
- Operation ordering: metadata 系（metadata_fields_replace / metadata_block_replace）→ named_section_replace の順で適用。
- Validation は全 operation 適用後の最終状態に対して実行。

## Impact Scope

- `docs/spec/design-records-mcp/tools.md` — `propose_record_update` の `operations` 配列 contract 追記
- `internal/designrecords/types.go` — `UpdateRequest` / `ProposeRecordUpdateRequest` 拡張
- `internal/designrecords/authoring.go` — `ProposeRecordUpdate` / `prepareUpdate` 拡張
- `internal/designrecords/authoring_test.go` — 回帰テスト追加
- `tmp.py` — smoke ケース追加

## Task flow

V01-TASK-MCP-026-01 → V01-TASK-MCP-026-02 → V01-TASK-MCP-026-03 → V01-TASK-MCP-026-04 → V01-TASK-MCP-026-05

V01-TASK-MCP-026-01 完了後にユーザーが spec を確認してから V01-TASK-MCP-026-02 に入る（ゲート A）。

## Task Candidates

tasks フィールドに記録。

## Completion Condition

- `propose_record_update` で `operations: [...]` が動作する
- Evidence + status:done を 1 proposal で適用できる
- 既存 `update: {...}` 形式の後方互換が保たれている
- 全テスト通過
- V01-REQ-MCP-025 が `accepted` になっている

## Evidence

2026-06-07: 全タスク完了。

- V01-TASK-MCP-026-01: spec 更新（operations array contract 追加）
- V01-TASK-MCP-026-02: 実装（ProposeRecordUpdateRequest + proposeMultiOpUpdate + prepareMultiOpUpdate + detectOperationConflicts）
- V01-TASK-MCP-026-03: テスト11件追加・全件 PASS
- V01-TASK-MCP-026-04: スモークテスト [6]-[9] 追加・全9件 PASS
- V01-TASK-MCP-026-05: クローズ同期

`go test ./internal/designrecords/... ./internal/designrecordsmcp/...` 全件 PASS を確認。
