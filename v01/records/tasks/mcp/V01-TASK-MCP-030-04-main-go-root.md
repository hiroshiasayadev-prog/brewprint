# V01-TASK-MCP-030-04: main.go `--root` 自動検出対応

- **id**: V01-TASK-MCP-030-04
- **status**: done
- **date**: 2026-06-09
- **work_item**: V01-WORK-MCP-030
- **source_requirement**: V01-REQ-MCP-033
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-030-02
  - V01-TASK-MCP-030-03
- **outputs**:
  - drmcp/src/cmd/design-records-mcp/main.go 更新
  - drmcp/src/cmd/design-records-mcp/main_test.go 更新

## Goal

`design-records-mcp --root <repo>` の単一起動で `*/records/` を自動検出し、multi-root index を構築して MCP server を起動できるようにする。

## Work

- `--records-root` 未指定時に `Config` の auto-detect ロジックを呼び出す形に `run()` を変更
- `--records-root` 指定時は単一 root モードとして動作（backward compat）
- `--summary` フラグの出力を multi-root 情報（各 root / namespace prefix / records 数）に対応
- `main_test.go` を auto-detect シナリオでカバー

## Done condition

- `--root <repo>` のみ指定でリポジトリ内の全 `*/records/` が検出され、統合 index が構築される
- `--records-root` 指定時は従来通り動作する
- `main_test.go` がすべて pass する

## Verification

`go test ./drmcp/src/cmd/design-records-mcp/...` が pass する。

## Evidence
`main.go` の summary 出力を multi-root 対応に更新。`cfg.RecordsRoot` 参照を削除し、`idx.RecordsEntries` ループで各 entry の `records: <root> (prefix: <ns>)` を出力。`--root <repo> --summary` 実行で bpdsl/drmcp/product/v01 の 4 ツリーを表示確認。
