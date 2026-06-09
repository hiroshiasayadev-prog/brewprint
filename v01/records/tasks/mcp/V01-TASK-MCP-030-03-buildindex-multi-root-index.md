# V01-TASK-MCP-030-03: BuildIndex multi-root 統合 index 構築

- **id**: V01-TASK-MCP-030-03
- **status**: done
- **date**: 2026-06-09
- **work_item**: V01-WORK-MCP-030
- **source_requirement**: V01-REQ-MCP-033
- **estimate**: 1.5d
- **depends_on**:
  - V01-TASK-MCP-030-02
- **outputs**:
  - drmcp/src/internal/designrecords/index.go 更新
  - drmcp/src/internal/designrecords/types.go 更新
  - drmcp/src/internal/designrecords/index_test.go 更新

## Goal

`BuildIndex` を複数 records root から record をスキャンして統合 index を構築できるように拡張する。

## Work

- `Index` struct から単数フィールド `NamespacePrefix` / `RecordsRoot` を削除または multi-root 対応に変更
- `BuildIndex` を `Config.RecordsRoots` をループして各 root の discover 関数を呼び出し、結果を `Index.Records` にマージする形に変更
- 各 record は自身の namespace prefix を含む public ID で識別されるため、ID 衝突は発生しない（prefix が異なる）
- `index_test.go` を multi-root index 構築のシナリオでカバー

## Done condition

- `BuildIndex` が複数 records root をスキャンし、全 record を統合した単一 `Index` を返す
- 各 record の public ID が各 namespace prefix を持つ
- `index_test.go` がすべて pass する

## Verification

`go test ./drmcp/src/internal/designrecords/... -run TestIndex` が pass する。

## Evidence
`Index` に `RecordsEntries []RecordsEntry` 追加。`BuildIndex` を `cfg.RecordsRoots` のループにリファクタリング。`idx.RecordsEntries` / `idx.RecordsRoot` / `idx.NamespacePrefix` 初期化。全テスト pass。実行確認: multi-root 時 405 records (bpdsl/drmcp/product/v01)。
