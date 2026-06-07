# TASK-MCP-028-02: types + implementation — DiffMode constants, Diff struct, authoring response shaping, MCP schema

- **id**: TASK-MCP-028-02
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-MCP-028
- **source_requirement**: REQ-MCP-031
- **estimate**: 1d-2d
- **depends_on**:
  - TASK-MCP-028-01
- **outputs**:
  - internal/designrecords/types.go
  - internal/designrecords/authoring.go
  - internal/designrecordsmcp/tools.go

## Goal

spec レビュー承認後、以下の実装を行う。

1. `internal/designrecords/types.go`:
   - `DiffMode` 型定数 (`DiffModeSummary`, `DiffModePatch`, `DiffModeNone`) を追加
   - `Diff` 構造体に `Omitted bool` フィールドを追加
   - `ProposeRecordCreateRequest` / `ProposeRecordUpdateRequest` に `DiffMode string` フィールドを追加

2. `internal/designrecords/authoring.go`:
   - `diff_mode` バリデーション（空 → `summary` デフォルト、不正値 → `invalid_request`）
   - `buildDiff` の呼び出し後、mode に応じて response の `Diff` をシェーピング
   - `summary`: `Diff{Format: "unified", Files: [...]}` (Text 省略)
   - `patch`: 現行と同等
   - `none`: `Diff{Omitted: true}`

3. `internal/designrecordsmcp/tools.go`:
   - `propose_record_create` / `propose_record_update` スキーマに `diff_mode` optional フィールドを追加

## Done condition

- `diff_mode` なしリクエストが `summary` モード応答（`diff.text` なし）を返す。
- `diff_mode: "patch"` が従来と同等の完全 diff を返す。
- `diff_mode: "none"` が `diff: {"omitted": true}` を返す。
- 不正な `diff_mode` 値で `invalid_request` が返る。
- `get_proposed_write` は引き続き完全 diff を返す（変更なし）。

## Work

- `internal/designrecords/authoring.go` に `DiffMode` 型定数・`Diff` 構造体更新（`omitempty` + `Omitted bool`）・リクエスト構造体に `DiffMode string` フィールド追加
- `validateAndResolveDiffMode` / `shapeDiff` ヘルパーを `buildDiff` 直前に追加
- `persistProposal` シグネチャに `diffMode DiffMode` を追加し、`shapeDiff` で応答 diff をシェーピング
- `ProposeRecordCreate` / `ProposeRecordUpdate` の先頭で diff_mode をバリデーション・resolve して各経路へパス
- `proposeMultiOpUpdate` シグネチャに `diffMode DiffMode` を追加して `persistProposal` へ渡す
- `internal/designrecordsmcp/tools.go` の両ツールスキーマに `diff_mode` optional フィールドを追加
- `internal/designrecords/authoring_test.go` の `diff.text` アサーションを持つ既存テストに `DiffMode: "patch"` を追加してデフォルト変更の影響を隔離

## Verification

`go test ./internal/designrecords/... ./internal/designrecordsmcp/...` が通ること。

## Evidence

2026-06-07: 実装完了。`go test ./internal/designrecords/... ./internal/designrecordsmcp/...` 全件 PASS。
