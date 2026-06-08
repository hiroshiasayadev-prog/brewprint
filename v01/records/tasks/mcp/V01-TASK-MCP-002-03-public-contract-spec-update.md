# V01-TASK-MCP-002-03: Batch retrieval の public contract / spec を更新する

- **id**: V01-TASK-MCP-002-03
- **status**: done
- **date**: 2026-05-26
- **work_item**: V01-WORK-MCP-002
- **source_requirement**: V01-REQ-MCP-002
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-MCP-002-02
- **outputs**:
  - Design Records MCP の更新済み public contract spec

## Goal

Batch retrieval capability の採用判断に従い、Design Records MCP の read-only public contract を、実装前に spec として確定する。

## Work

- 対象 tool の名称、request fields、response fields、ordering、件数上限を spec に定義する。
- 本文取得の可否と response size 制約を定義する。
- 一部 ID が存在しない場合の result / diagnostic contract を定義する。
- 既存 `list_records` / `get_record` の責務との重複・互換性を確認する。
- 必要な spec metadata / semantic ref / 由来注記を更新する。

## Done condition

- Public contract が implementation 前に spec 上で明文化されている。
- V01-TASK-MCP-002-04 が曖昧な contract 判断なしに実装へ進める。
- ADR を起票した場合、その判断と spec が矛盾していない。

## Verification

- `docs/spec/design-records-mcp/` の関連文書間で tool set と schema が整合していることを確認する。
- Existing tool の behavior を意図せず変更していないことを確認する。

## Evidence

- 2026-05-26: V01-ADR-090 のレビュー指摘を受け、実装前に concrete contract を spec へ明文化する方針を採用した。
- `docs/spec/design-records-mcp/tools.md` に `get_records` の request / response contract を追加した。response collection は `items` とし、found item は既存 `get_record.record` representation を再利用し、missing item は `retrieval_status: "not_found"` と item-level `record_not_found` diagnostic を返す。
- `ids` は required non-empty string array とし、欠落・空配列・非 array・非 string element は `invalid_request` とした。全 ID が missing の場合も normal response として扱う。
- duplicate requested ID は first occurrence の item のみ返し、ID ごとに top-level `duplicate_requested_id_ignored` / `info` diagnostic を返す。payload は `requested_id` / `first_index` / `duplicate_indexes` とし、index は zero-based とした。
- `ids[]` は exact record ID lookup key としてのみ評価し、resolver を呼ばない。string として受理されたが index に一致しない値は item-level `not_found` とした。
- `docs/spec/design-records-mcp/schema.md` に retrieval wrapper と request-level diagnostic vocabulary を反映し、`docs/spec/design-records-mcp/overview.md` に `get_records` の tool set / 対象責務を反映した。
- Spec の更新内容は V01-ADR-090 の責務境界、partial result、duplicate handling、raw body 非 truncate、public numeric size limit 非定義と整合する。
- Independent review で挙がった response shape、all-missing / empty `ids`、duplicate diagnostic payload、exact record ID と reference resolve の境界の四点は、実装前の public contract として明文化済みである。V01-TASK-MCP-002-04 は追加の contract 判断なしに実装へ進める。
