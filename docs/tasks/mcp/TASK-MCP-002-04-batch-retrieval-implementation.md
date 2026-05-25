# TASK-MCP-002-04: Batch retrieval capability を実装する

- **id**: TASK-MCP-002-04
- **status**: done
- **date**: 2026-05-26
- **work_item**: WORK-MCP-002
- **source_requirement**: REQ-MCP-002
- **estimate**: 1.5d
- **depends_on**:
  - TASK-MCP-002-03
- **outputs**:
  - Design Records MCP batch retrieval implementation
  - implementation-level tests

## Goal

確定した public contract に従って、複数 design record を取得する read-only capability と必要な tests を実装する。

## Work

- 確定した tool contract に対応する handler / service route を実装する。
- record ordering、本文取得、件数上限、partial result / diagnostic behavior を contract 通りに実装する。
- 既存 single-record retrieval と metadata parser / index の共有範囲を整理し、不要な重複を避ける。
- 主要な正常系・境界値・一部 unresolved input の unit / contract tests を追加する。

## Done condition

- Spec に定義された capability が実装されている。
- 実装レベル tests が追加され、想定 contract を検証できる。
- TASK-MCP-002-05 が runtime 確認へ進める。

## Verification

- Existing `list_records` / `get_record` / `validate_records` / `resolve_reference` の tests に回帰がないことを確認する。
- Public response に未定義 field や accidental compatibility behavior を持ち込んでいないことを確認する。

## Evidence

- 2026-05-26: `internal/designrecords/types.go` に `GetRecordsRequest` / `GetRecordsItem` / `GetRecordsResponse`、`RetrievalStatus`、request/retrieval diagnostic payload field (`requested_id`, `first_index`, `duplicate_indexes`) と diagnostic category を追加した。
- 2026-05-26: `internal/designrecords/tools.go` に `GetRecords` handler を実装した。明示 ID の exact lookup、first-occurrence ordering、item-level `not_found` partial result、duplicate requested ID の top-level `info` diagnostic、既存 `getRecordResponseRecord` による body / representation 再利用を扱う。
- 2026-05-26: `internal/designrecordsmcp/tools.go` / `tools_call.go` に `get_records` の tool schema と dispatch route を追加した。`ids` は required non-empty string array として公開する。
- 2026-05-26: `internal/designrecords/get_records_test.go` を追加し、mixed `decision` / `spec` / `investigation` retrieval、partial result、duplicate payload、raw body、all-missing exact lookup、invalid request を検証対象とした。
- 2026-05-26: `internal/designrecordsmcp/jsonrpc_test.go` / `tools_call_test.go` を更新し、tools/list での公開、tools/call の partial result / duplicate info、欠落・空・非 string `ids` の tool error を検証対象とした。
- ADR-090 / spec に従い、public 件数上限または response size 数値上限は実装 contract に追加しない。body の truncate も導入しない。
- 2026-05-26: ユーザー実行により、変更対象ファイルへ `gofmt` を適用した。
- 2026-05-26: ユーザー実行による対象 package test (`./internal/designrecords`, `./internal/designrecordsmcp`) は成功した。
- 2026-05-26: ユーザー実行による repository 全体の Go test suite (`./...`) は成功し、`cmd/design-records-mcp` および既存 package を含む regression がないことを確認した。
- 以上により、implementation と implementation-level test の完了条件を満たす。runtime を用いた contract 確認は TASK-MCP-002-05 に引き継ぐ。
