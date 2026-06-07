# TASK-MCP-026-01: spec draft — add operations array contract to propose_record_update

- **id**: TASK-MCP-026-01
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-MCP-026
- **source_requirement**: REQ-MCP-025
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Updated docs/spec/design-records-mcp/tools.md

## Goal

`propose_record_update` に `operations: [...]` 配列の contract を spec（tools.md）に追記する。実装前に spec を確定させる。

## Work

`docs/spec/design-records-mcp/tools.md` の `propose_record_update` セクションに以下を追記する。

- `operations` 配列の request shape（各 op に `body` / `body_cache_id` をインラインで持つ Y 形式）
- `update` と `operations` の排他ルール
- body source rules（operations 配列使用時は各 op にインライン）
- operation ordering: metadata 系（metadata_fields_replace / metadata_block_replace）→ named_section_replace の順
- conflict detection rules: 同一フィールドに複数の metadata 系 op → error、named_section_replace が 2 op 以上 → error
- named_section_replace は 1 op のみサポート（制約）の明示
- multi-op proposal の response shape（diff, validation affected set）
- no-op detection: 全 op が no-op の場合の扱い

## Done condition

tools.md の `propose_record_update` セクションに `operations` 配列 contract が追記され、ユーザーが内容を確認して TASK-MCP-026-02 への GO を出した。

## Verification

追記内容が REQ-MCP-025 の acceptance criteria と矛盾していないことを照合する。

## Evidence
2026-06-07: `docs/spec/design-records-mcp/tools.md` の `propose_record_update` セクションに operations 配列 contract を追記した。

変更内容:
- Purpose: single-op / multi-op 両対応と明記
- Request: operations 配列例 JSON、フィールドテーブル更新（update を conditional に、operations 追加）、排他ルール
- 新設 `#### Operations array` サブセクション: cardinality、各 op のフィールドテーブル、full semantics 継承（heading-safe normalization・fallback 含む）、body source rules (MUST/MUST NOT)、body cache retry、operation ordering、conflict detection、validation、no-op detection、response shape
- `### Diagnostic categories`: `conflicting_operations`・`multiple_section_replace_not_supported` を bullet list と説明に追加

Codex レビュー 4 件（cardinality 未定義・body cache retry 未記述・diagnostic catalog 未登録・normative language 弱い）を全対応。ユーザーから TASK-MCP-026-02 への GO を得た。
