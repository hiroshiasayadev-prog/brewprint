# V01-WORK-MCP-030: Design Records MCP に multi-root スキャンを実装する

- **id**: V01-WORK-MCP-030
- **status**: not_started
- **date**: 2026-06-09
- **source_requirement**: V01-REQ-MCP-033
- **impact_refs**:
  - V01-REQ-MCP-033
- **tasks**:

## Goal

`design-records-mcp --root <repo>` の単一起動で、リポジトリ内の全 `*/records/` ツリーを自動検出して統合 index を構築できるようにする。これにより `v01/records`（V01-）・`drmcp/records`（DRMCP-）等を単一 MCP プロセスで横断できる。

## Background

現状は `--records-root` で1つの records ツリーのみを scan する単一 root 構成。V01-REQ-MCP-033 で multi-root スキャンが要件として確定した。`drmcp/records/` 配下の workflow artifact を MCP tools でオーサリングするためにも必須。

## Boundary

このWORKが所有するもの:
- `drmcp/src/internal/designrecords/`: 複数 records root からのレコードを統合する multi-root index 構築
- `drmcp/src/cmd/design-records-mcp/main.go`: リポジトリルートからの `*/records/` 自動検出（または複数 `--records-root` 指定対応）
- authoring tools の namespace ルーティング（public ID の prefix からターゲット records ツリーを特定してファイル生成）
- spec 更新（`drmcp/records/spec/design-records-mcp/overview.md` の multi-root セクション改訂）

このWORKが所有しないもの:
- cross-namespace relation validation（将来課題）
- 既存 artifact の namespace 間移行
