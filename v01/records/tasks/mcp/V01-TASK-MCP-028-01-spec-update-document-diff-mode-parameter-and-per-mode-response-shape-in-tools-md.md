# V01-TASK-MCP-028-01: spec update — document diff_mode parameter and per-mode response shape in tools.md

- **id**: V01-TASK-MCP-028-01
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-028
- **source_requirement**: V01-REQ-MCP-031
- **estimate**: 1d
- **depends_on**:
- **outputs**:
  - docs/spec/design-records-mcp/tools.md

## Goal

`docs/spec/design-records-mcp/tools.md` に以下の変更を加える。

1. `propose_record_create` Request テーブルに `diff_mode` パラメータ行を追加
2. `propose_record_update` Request テーブルに `diff_mode` パラメータ行を追加
3. Common authoring response fields の `diff` テーブルを更新:
   - `diff.text` を `patch` モード限定の条件付きフィールドとする
   - `diff.omitted` フィールド（`none` モード時）を追記
4. `diff_mode` の値・デフォルト・各モード応答形状を説明するセクションを追加

## Done condition

- `diff_mode` の 3 値（`summary` / `patch` / `none`）とデフォルト（`summary`）が明記されている。
- 各モードの `diff` 応答形状が表またはサンプルで明示されている。
- `get_proposed_write` は `diff_mode` 非対応（常に patch 相当）であることが記載されている。
- 不正な `diff_mode` 値で `invalid_request` が返る旨が記載されている。

## Work

`docs/spec/design-records-mcp/tools.md` を 4 箇所編集:

1. `### Common authoring response fields` の `diff` テーブルを `diff_mode` 対応に更新。`text` を `patch` 限定・`omitted` を `none` 限定として明記。
2. `### diff_mode` サブセクション新設（Authoring transaction model 内）。3 値の意味・デフォルト・invalid 動作・`get_proposed_write` 非対応を記載。
3. `propose_record_create` Request テーブルに `diff_mode` 行を追加。
4. `propose_record_update` Request テーブルに `diff_mode` 行を追加。

## Verification

spec レビュー（外部 LLM レビュー）を経てから次タスクへ進む。

## Evidence

2026-06-07: `docs/spec/design-records-mcp/tools.md` を初回 4 箇所編集後、外部 LLM レビュー（Needs revision）を受領。
以下の修正を追加実施:
- L942: `previewable diff` → `diff response object`
- L977: `previewable diff object` → `diff response object (shape depends on diff_mode)`
- L993: forward reference を Markdown anchor link に変更
- L1106 付近: `summary` の concise operation summaries = `diff.files[]` を明記
- L1615: multi-op response の `diff.text` 記述を `patch` モード限定に変更
- L1624/1627: update response の `diff.text MUST` 記述を `patch` モード限定に変更
spec レビュー再実施を推奨。
2回目レビュー（Needs revision）を受領。diff.files の無条件記述 4 箇所を追加修正:
- L1302: reciprocal update の diff.files → `When diff.files is present, it includes...`
- L1617: multi-op の diff.files → `When diff_mode is summary or patch, diff.files contains...`
- L1622: update Response の diff.files[].change → `When diff.files is present, diff.files[].change is modify`
- L1733: named section replacement の同上
spec レビュー 3 回目を推奨。
2026-06-07: 3 回目レビューで OK・承認取得。V01-TASK-MCP-028-02 に進む。
