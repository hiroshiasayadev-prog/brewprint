# V01-TASK-MCP-026-02: implementation — extend ProposeRecordUpdateRequest and prepareUpdate for multi-op

- **id**: V01-TASK-MCP-026-02
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-026
- **source_requirement**: V01-REQ-MCP-025
- **estimate**: 1d-2d
- **depends_on**:
  - V01-TASK-MCP-026-01
- **outputs**:
  - Updated internal/designrecords/types.go
  - Updated internal/designrecords/authoring.go

## Goal

`ProposeRecordUpdateRequest` に `operations` 配列を追加し、`ProposeRecordUpdate` が `operations` 配列を処理できるよう `authoring.go` を拡張する。既存の `update: {...}` 形式の後方互換を維持する。

## Work

**types.go:**
- `UpdateOperation` struct を新設（`Type`, `Metadata`, `SectionSelector`, `Body *string`, `BodyCacheID string`）
- `ProposeRecordUpdateRequest` に `Operations []UpdateOperation` を追加

**authoring.go:**
- `ProposeRecordUpdate`: `update` / `operations` 排他チェック（両方指定はエラー）
- `ProposeRecordUpdate`: `operations` 配列の各 op から body source を resolve する処理
- `prepareMultiOpUpdate` を新設:
  - conflict detection: metadata 系で同一フィールドに複数 op → `conflicting_operations` diagnostic
  - conflict detection: named_section_replace が 2 op 以上 → `multiple_section_replace_not_supported` diagnostic
  - operation 適用順: metadata 系を先に適用 → named_section_replace を適用
  - 全 op 適用後の最終状態に対して validation を実行

## Done condition

- `operations: [metadata_fields_replace, named_section_replace]` の proposal が正常に作成される
- 既存 `update: {...}` 形式が壊れない（回帰なし）
- `go test ./internal/designrecords/...` が通る

## Verification

authoring_test.go の既存テストが全通過することを確認する。

## Evidence
2026-06-07: 実装完了。

変更ファイル:
- `internal/designrecords/types.go`: `ErrorCodeConflictingOperations`, `ErrorCodeMultipleSectionReplaceNotSupported` を追加。
- `internal/designrecords/authoring.go`:
  - `ProposeRecordUpdateRequest` に `Operations []UpdateOperation` フィールドを追加。
  - `UpdateOperation` 構造体を新設（Type / Metadata / SectionSelector / Body / BodyCacheID）。
  - `ProposeRecordUpdate` に operations / update 排他ガード、top-level body 禁止ガードを追加。
  - `proposeMultiOpUpdate` 関数: per-op body source 解決 → `prepareMultiOpUpdate` 呼び出し。
  - `prepareMultiOpUpdate` 関数: conflict detection → metadata ops pass → section op pass → single-file ProposedFile を返す。
  - `detectOperationConflicts` 関数: `multiple_section_replace_not_supported` / `conflicting_operations` を検出。

`go build ./internal/designrecords/...` および `go test ./internal/designrecords/... ./internal/designrecordsmcp/...` の両方がパスしたことを確認済み。

deferred work（複数セクション replace 対応）は V01-TASK-MCP-026-02 本文の Evidence 末尾に記載済み。
