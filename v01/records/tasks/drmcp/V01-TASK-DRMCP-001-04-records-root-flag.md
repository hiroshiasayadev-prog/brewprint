# V01-TASK-DRMCP-001-04: design-records-mcp に --records-root フラグを追加する

- **id**: V01-TASK-DRMCP-001-04
- **status**: done
- **date**: 2026-06-09
- **work_item**: V01-WORK-DRMCP-001
- **source_requirement**: (TBD)
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-DRMCP-001-03
- **outputs**:
  - drmcp/src/cmd/design-records-mcp/main.go

## Goal

drmcp/src/cmd/design-records-mcp/main.go に `--records-root` フラグを追加し、design-records-mcp が任意の records ディレクトリを scan できるようにする。

## Work

- `run()` 関数に `--records-root` フラグ追加（デフォルト: 空文字列 → `NewConfig` で `v01/records` に fallback）
- `NewConfig(*root, *recordsRoot)` を使う形に更新

## Done condition

- `go build ./drmcp/src/cmd/design-records-mcp` が成功する
- `--root <repo>` 単体でも `v01/records` を scan できる

## Verification

- `go build ./drmcp/src/cmd/design-records-mcp` を実行する

## Evidence

- 2026-06-09: `--records-root` フラグを追加し、デフォルト `v01/records` で動作することを確認。`go build ./drmcp/src/cmd/design-records-mcp` 成功。
