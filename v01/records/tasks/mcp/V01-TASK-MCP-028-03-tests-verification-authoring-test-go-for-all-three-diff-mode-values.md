# V01-TASK-MCP-028-03: tests + verification — authoring_test.go for all three diff_mode values

- **id**: V01-TASK-MCP-028-03
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-028
- **source_requirement**: V01-REQ-MCP-031
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-MCP-028-02
- **outputs**:
  - internal/designrecords/authoring_test.go

## Goal

`internal/designrecords/authoring_test.go` に `diff_mode` 対応のテストを追加する。

カバレッジ:
- `diff_mode` 省略 → `summary` モード（`diff.text` 空・`diff.files` あり）
- `diff_mode: "patch"` → `diff.text` が存在する
- `diff_mode: "none"` → `diff.omitted: true`・files/text なし
- 不正値 → `invalid_request` エラー
- create / update 両方をカバーする

## Done condition

`go test ./internal/designrecords/... ./internal/designrecordsmcp/...` がすべて PASS。

## Work

`internal/designrecords/authoring_test.go` に `TestDiffModeRequestParameter` を追加。
サブテスト 8 件（create × 4 / update × 4）:

- `create_default_is_summary` — DiffMode 省略 → `diff.text` 空・`diff.files` あり・`diff.omitted` false
- `create_patch_includes_diff_text` — `DiffMode: "patch"` → `diff.text` 非空・`diff.files` あり
- `create_none_omits_diff` — `DiffMode: "none"` → `diff.omitted: true`・text/files 空
- `create_invalid_diff_mode` — 不正値 → `invalid_request` diagnostic・proposal 未作成
- `update_default_is_summary` — 同上 (update)
- `update_patch_includes_diff_text` — 同上 (update)
- `update_none_omits_diff` — 同上 (update)
- `update_invalid_diff_mode` — 同上 (update)

## Verification

テスト PASS を確認後、V01-WORK-MCP-028 / V01-REQ-MCP-031 のステータス同期とクローズ処理を行う。

## Evidence

2026-06-07: `TestDiffModeRequestParameter` 全 8 サブテスト PASS。`go test ./internal/designrecords/... ./internal/designrecordsmcp/...` 全件 PASS 確認。
