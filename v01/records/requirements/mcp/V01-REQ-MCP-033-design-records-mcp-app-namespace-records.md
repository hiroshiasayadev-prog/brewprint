# V01-REQ-MCP-033: Design Records MCP がリポジトリルートから複数 app namespace の records ツリーを横断スキャンできること

- **id**: V01-REQ-MCP-033
- **status**: accepted
- **date**: 2026-06-09
- **source_refs**:
  - V01-REQ-PRODUCT-001
  - V01-REQ-PRODUCT-003
- **work_items**:
  - V01-WORK-MCP-030

## Requirement

Design Records MCP は、`--root <repo>` でリポジトリルートを指定した単一プロセスとして起動し、リポジトリ内の複数の app namespace records ツリー（例: `v01/records/`、`drmcp/records/`、`bpdsl/records/`）を同時にスキャンして統合 index を構築できなければならない。

各 records ツリーの namespace prefix はディレクトリ名から機械的に導出する（`strings.ToUpper(appNamespaceDir) + "-"`）。統合 index では各 record が自身の namespace prefix を持つ public ID で識別され、異なる app namespace の record を同一クエリで取得・参照解決できる。

## Required Outcome

- `design-records-mcp --root <repo>` の単一起動で、リポジトリ内の全 `*/records/` ツリーをスキャンして統合 index を構築する
- 各 record の public ID は `<NAMESPACE>-<KIND>-...` 形式（例: `V01-ADR-076`、`DRMCP-REQ-MCP-001`）で区別される
- `list_records`、`get_record`、`resolve_reference` 等の全 MCP tools が統合 index に対して動作する
- authoring tools（`propose_record_create`、`propose_record_update`）はターゲットの namespace を public ID の prefix から特定し、対応する records ツリーにファイルを生成・更新する
- cross-namespace の relation（例: `V01-WORK-DRMCP-001` が `DRMCP-TASK-MCP-001-01` を参照する）を index が保持し、`resolve_reference` / `validate_records` が解決できる

## Explicitly Excluded Scope

- 異なる app namespace 間での record ID 重複検出（namespace prefix が異なれば public ID は衝突しない）
- records ツリーが存在しない app namespace ディレクトリの自動作成
- `--records-root` による単一 root モードの廃止（後方互換として残す）
