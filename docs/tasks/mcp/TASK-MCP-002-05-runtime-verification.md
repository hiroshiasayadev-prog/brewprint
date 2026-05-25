# TASK-MCP-002-05: Batch retrieval capability を検証し runtime で確認する

- **id**: TASK-MCP-002-05
- **status**: done
- **date**: 2026-05-26
- **work_item**: WORK-MCP-002
- **source_requirement**: REQ-MCP-002
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-002-04
- **outputs**:
  - verification result
  - runtime confirmation evidence

## Goal

確定した public contract と実装が一致し、実際の LLM 読み取り経路で batch retrieval capability が利用可能であることを確認する。

## Work

- 関連 package の unit / contract tests を実行する。
- 必要な範囲で全体回帰 tests を実行する。
- MCP runtime 経由で複数 record 取得の正常系と主要な境界挙動を確認する。
- 反復 `get_record` 利用と比べて、REQ-MCP-002 の負荷が実際に軽減されるかを確認する。

## Done condition

- Contract と implementation の一致が検証されている。
- Runtime で期待した query path が成立している。
- 不具合または不足があれば、close 前に修正対象として記録されている。

## Verification

- Tests と runtime confirmation の実行結果を evidence として残す。
- Verification で発見した追加要求を本 task 内で黙って拡張実装せず、必要なら requirement として捕捉する。

## Evidence

- 2026-05-26: ユーザー実行により `go run ./cmd/design-records-mcp -root .` を stdio JSON-RPC で起動し、`tools/call` の `name=get_records` を用いて runtime contract を確認した。ChatGPT 側に公開された MCP tool 一覧では新 tool が露出しなかったため、runtime への直接 JSON-RPC 呼出しを verification path とした。
- 正常系: `ADR-077` / `SPEC-design-records-mcp-tools` / `INV-DOCS-001` の batch retrieval は `items` 3件を返し、全 item が `retrieval_status: "found"`、top-level `diagnostics` は空であった。
- Partial result: `INV-DOCS-999` を含む request は tool error とならず normal response を返し、対象 item は `retrieval_status: "not_found"`、item-level diagnostic は `category: "record_not_found"`, `severity: "error"`, `requested_id: "INV-DOCS-999"` であった。
- Duplicate: `ADR-077` を重複指定した request は item を一件のみ返し、top-level diagnostic に `duplicate_requested_id_ignored`、`first_index: 0`、`duplicate_indexes: [1]` を返した。
- Body: `include_body: true` の found item で `record.body` が返り、`ADR-077` の raw Markdown 本文 prefix が確認できた。
- 追加回帰確認として、ユーザー実行による `go test ./internal/designrecordsmcp ./internal/designrecords` も成功した。
- 上記により、public contract と implementation の一致、および MCP runtime 経由で期待した batch retrieval query path が成立することを確認した。追加の capability 不具合は観測されていない。
- 観測事項: ChatGPT 側に露出する MCP tool registry / schema が再 build 後も新 tool を表示しない場面があった。runtime 自体では `get_records` が動作しており、本 capability の実装不具合とは切り分ける。接続再登録または tool schema refresh の運用上の確認事項として TASK-MCP-002-06 に引き継ぐ。
