# V01-TASK-DRMCP-001-03: namespace prefix サポートを実装する

- **id**: V01-TASK-DRMCP-001-03
- **status**: done
- **date**: 2026-06-09
- **work_item**: V01-WORK-DRMCP-001
- **source_requirement**: (TBD)
- **estimate**: 2d
- **depends_on**:
  - V01-TASK-DRMCP-001-02
- **outputs**:
  - drmcp/src/internal/designrecords/config.go
  - drmcp/src/internal/designrecords/types.go
  - drmcp/src/internal/designrecords/index.go
  - drmcp/src/internal/designrecords/parser.go
  - drmcp/src/internal/designrecords/authoring.go
  - drmcp/src/internal/designrecords/tools.go
  - drmcp/src/internal/designrecords/resolver.go

## Goal

config.go / types.go / index.go / parser.go / authoring.go / tools.go / resolver.go に namespace prefix 導出・strip・付与の処理を実装する。

## Work

- config.go: `NewConfig(root, recordsRoot string)` の 2 引数化、`NamespacePrefix()` 導出メソッド追加
- types.go: `Index` に `NamespacePrefix` / `RecordsRoot` フィールド追加
- index.go: `BuildIndex` で `RecordsRoot` / `NamespacePrefix` を `Index` にセット
- parser.go: ADR H1 を `# ADR-NNN:` 形式に更新、namespace prefix ストリップ後に bare ID 検証
- authoring.go: `parseRecordByPath` の `v01/records/` ハードコードを除去し `recordsRoot` パラメータ化
- tools.go: `SuggestNextRecord` を ns-aware に、`suggestedDecisionRecordPath` を RecordsRoot ベースに修正
- resolver.go: `resolveReference` で V01- prefix 付き public ID を受け付けるよう修正

## Done condition

- `go test ./drmcp/src/internal/designrecords/...` 全件 PASS

## Verification

- `go test ./drmcp/src/internal/designrecords/...` を実行する

## Evidence

- 2026-06-09: 上記 7 ファイルに namespace prefix 対応を実装。`go test ./drmcp/src/internal/designrecords/...` 全件 PASS 確認。
