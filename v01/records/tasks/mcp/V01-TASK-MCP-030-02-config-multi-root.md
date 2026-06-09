# V01-TASK-MCP-030-02: Config multi-root 対応

- **id**: V01-TASK-MCP-030-02
- **status**: done
- **date**: 2026-06-09
- **work_item**: V01-WORK-MCP-030
- **source_requirement**: V01-REQ-MCP-033
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-MCP-030-01
- **outputs**:
  - drmcp/src/internal/designrecords/config.go 更新
  - drmcp/src/internal/designrecords/config_test.go 更新

## Goal

`Config` を複数 records root を保持できる構造に拡張し、auto-detect ロジックの基盤を整える。

## Work

- `Config` の `RecordsRoot string` を `RecordsRoots []RecordsEntry` に変更（各 entry は `RecordsRoot` + 導出 `NamespacePrefix` を持つ）
- `NewConfig` を更新：`--root` 指定時に `<root>/*/records/` を glob して自動検出、`--records-root` 指定時は単一 root として backward compat を維持
- `normalizeConfig` を更新
- `config_test.go` を multi-root 構造に合わせて更新

## Done condition

- `Config` が複数 `RecordsEntry` を保持できる
- `--root` のみ指定時に `*/records/` が自動検出される
- `--records-root` 指定時は従来通り単一 root として動作する
- `config_test.go` がすべて pass する

## Verification

`go test ./drmcp/src/internal/designrecords/... -run TestConfig` が pass する。

## Evidence
`RecordsEntry` struct 追加。`Config.RecordsRoot string` を `Config.RecordsRoots []RecordsEntry` に置換。`NamespacePrefix()` / `primaryRecordsRoot()` メソッド追加。`NewConfig` に `discoverRecordsEntries` による auto-detect 実装。`makeRecordsEntry` ヘルパー追加。`normalizeConfig` 更新。コンパイルエラー 0。
