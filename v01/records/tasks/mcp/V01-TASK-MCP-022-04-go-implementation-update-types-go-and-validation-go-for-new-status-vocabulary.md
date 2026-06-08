# V01-TASK-MCP-022-04: Go implementation: update types.go and validation.go for new status vocabulary

- **id**: V01-TASK-MCP-022-04
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-022
- **source_requirement**: V01-REQ-MCP-026
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-022-03
- **outputs**:
  - internal/designrecords/types.go updated
  - internal/designrecords/validation.go updated
  - tests pass

## Goal

`internal/designrecords/types.go` と `internal/designrecords/validation.go` を V01-ADR-094 に従って更新し、新 status vocabulary (`not_started` / `in_progress` / `blocked` / `done`) を Go 実装に反映する。

## Work

- `internal/designrecords/types.go`:
  - `RecordStatusTodo` (`"todo"`) を削除
  - `RecordStatusDoing` (`"doing"`) を削除
  - `RecordStatusInProgress = "in_progress"` を追加
  - `work_item` ステージ定数をすべて削除: `RecordStatusDecisionPending` / `RecordStatusDesignSpecPending` / `RecordStatusInternalDesignPending` / `RecordStatusYAMLPending` / `RecordStatusImplementationPending` / `RecordStatusFixturePending` / `RecordStatusVerificationPending`
- `internal/designrecords/validation.go`:
  - `isAllowedStatusForKind` の `work_item` 分岐: `not_started` / `in_progress` / `blocked` / `done` のみ
  - `isAllowedStatusForKind` の `task` 分岐: `not_started` / `in_progress` / `blocked` / `done` のみ
- 関連テスト (`validation_test.go`、`parser_index_test.go` 等) を更新

## Done condition

- `go test ./internal/designrecords/...` が pass
- 旧トークン (`todo` / `doing` / `decision_pending` 等) が `invalid_status_for_kind` を返す
- `not_started` / `in_progress` / `blocked` / `done` が両 kind で受理される

## Verification

`go test ./internal/designrecords/...`

## Evidence

- `internal/designrecords/types.go`: `RecordStatusTodo`/`RecordStatusDoing` 削除、`RecordStatusInProgress = "in_progress"` 追加、work_item ステージ定数 7 種削除
- `internal/designrecords/validation.go`: `statusAllowedForKind` の `work_item` / `task` 分岐を `not_started` / `in_progress` / `blocked` / `done` に変更
- テストファイル 5 件更新: `authoring_test.go`、`validation_test.go`、`parser_index_test.go`、`get_records_test.go`、`list_records_test.go`
- `internal/designrecordsmcp/tools_call_test.go`: 旧定数参照 1 件・フィクスチャ文字列 3 件を更新
- `go test ./...` 全パッケージ pass 確認
