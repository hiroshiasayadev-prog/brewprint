# TASK-MCP-003-04: Workflow artifact MCP support を実装する

- **id**: TASK-MCP-003-04
- **status**: done
- **date**: 2026-05-26
- **work_item**: WORK-MCP-003
- **source_requirement**: REQ-MCP-003
- **estimate**: 1.5d
- **depends_on**:
  - TASK-MCP-003-03
- **outputs**:
  - 採用された workflow artifact MCP surface の implementation
  - 対象 parser / index / handler / validation の automated tests

## Goal

確定した public contract に従い、workflow artifact MCP support の採用範囲を実装し、仕様で定義していない capability を暗黙に追加しない。

## Work

- TASK-MCP-003-03 で確定した record/query / resolver / validation surface に限って implementation を更新する。
- `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref の処理と relation metadata の解析を、採用 contract に応じて実装する。
- Physical path を canonical workflow relation として受理しない境界を保持する。
- 採用した正常系・不正 relation・unsupported / unresolved 境界の unit / transport tests を追加する。

## Done condition

- 採用された public contract の implementation と automated tests が揃っている。
- MVP 外と判断された orphan diagnostics / progress projection 等を暗黙実装していない。
- TASK-MCP-003-05 が runtime verification を開始できる。

## Verification

- 対象 package tests と repository の必要な test suite を実行する。
- Spec と implementation / tests の response shape と diagnostic boundary を照合する。

## Evidence

- 2026-05-27: Pass 1: Record surface integration に着手した。今回の実装対象は workflow artifact の `requirement` / `work_item` / `task` を既存 Design Records MCP の record-oriented surface に追加する範囲に限定する。
- Pass 1 では workflow document discovery、metadata parsing、index registration、`list_records` / `get_record` / `get_records` の exact retrieval、stdio MCP tool schema の kind enum 同期、および対象 automated tests を扱う。
- Pass 1 では `resolve_reference` の workflow ID support、investigation metadata の workflow reference validation、workflow relation existence / bidirectional integrity validation、`task.source_requirement` / `task.depends_on` validation、orphan diagnostics、progress projection、workflow traversal、cycle detection は実装しない。
- 2026-05-27: Pass 1 implementation を完了。`internal/designrecords` に workflow record kind / status vocabulary / detail object / parser / discovery / list-get-batch retrieval response support / parse-status diagnostics を追加し、`internal/designrecordsmcp` の public tool schema enum と description を同期した。
- 2026-05-27: Pass 1 tests として parser / index / list / get / batch get / status diagnostic / stdio tool schema / tools/call 経由の workflow record response を追加・更新した。既存 decision / spec / investigation response shape と JSON error behavior の regression tests は維持した。
- 2026-05-27: Test result: `go test ./internal/designrecords ./internal/designrecordsmcp` passed。追加確認として `go test ./...` passed。
- Pass 2 / Pass 3 に残す項目: `resolve_reference` の `REQ-*` / `WORK-*` / `TASK-*` support、investigation metadata での `REQ-*` / `WORK-*` support と `TASK-*` unsupported handling、workflow relation existence validation、declared bidirectional relation integrity validation、`task.source_requirement` consistency validation、`task.depends_on` target validation。
- MVP 外として引き続き実装しない項目: orphan diagnostics、progress projection、workflow traversal / graph query、cycle detection / execution-order projection、physical path relation support、`req:` / `work:` / `task:` semantic prefix。
- 2026-05-27: Pass 2: resolver / investigation metadata boundary integration に着手した。今回の実装対象は `resolve_reference` の direct workflow ID-as-ref support と、investigation metadata の `source_refs` / `follow_up_results` / `follow_up_candidates` における `REQ-*` / `WORK-*` support および `TASK-*` unsupported boundary に限定する。
- 2026-05-27: Pass 2 implementation を完了。`internal/designrecords` で direct resolver の supported record ID-as-ref に `REQ-*` / `WORK-*` / `TASK-*` を追加し、canonical grammar に合う未存在 workflow ID を `unresolved`、malformed workflow-like input と physical path を `unsupported` とする boundary を固定した。
- 2026-05-27: Investigation metadata validation は direct resolver の supported set をそのまま流用せず、`source_refs` / `follow_up_results` / `follow_up_candidates` で `REQ-*` / `WORK-*` の canonical resolve を扱い、`TASK-*` は存在有無にかかわらず unsupported とした。`source_refs` / `follow_up_results` は unsupported error、`follow_up_candidates` は unsupported info とする。
- 2026-05-27: Pass 2 tests として resolver の workflow direct input / missing canonical workflow ID / malformed workflow-like unsupported / repository bootstrap direct resolve、investigation metadata の `REQ-*` / `WORK-*` valid / unresolved と `TASK-*` unsupported boundary、stdio `tools/call` 経由の workflow `resolve_reference` response を追加・更新した。
- 2026-05-27: Test result: `go test ./internal/designrecords ./internal/designrecordsmcp` passed。追加確認として `go test ./...` passed。
- Pass 3 に残す項目: `requirement.work_items` target validation、`work_item.source_requirement` target validation、`work_item.tasks` target validation、`task.work_item` target validation、`task.source_requirement` target validation、`task.depends_on` target validation、`REQ.work_items <-> WORK.source_requirement` bidirectional integrity、`WORK.tasks <-> TASK.work_item` bidirectional integrity、`TASK.source_requirement == parent WORK.source_requirement` consistency validation。
- MVP 外として引き続き実装しない項目: orphan diagnostics、progress projection、workflow traversal / tree / graph query、dependency cycle detection / execution-order projection、physical path support、`req:` / `work:` / `task:` semantic prefix。
- 2026-05-27: Pass 3: workflow declared relation validation に着手した。今回の実装対象は workflow metadata に明示された relation の target kind / target existence / reciprocal integrity / task source requirement consistency validation に限定し、orphan diagnostics、progress projection、workflow traversal、dependency cycle detection、same-work-item dependency rule、self-dependency rule は実装しない。
- 2026-05-27: Pass 3 implementation を完了。`validate_records` で `requirement.work_items`、`work_item.source_requirement`、`work_item.tasks`、`task.work_item`、`task.source_requirement`、`task.depends_on` の declared relation を検証し、`unresolved_workflow_relation`、`invalid_workflow_relation_target`、`workflow_relation_mismatch`、`workflow_source_requirement_mismatch` を error diagnostic として返すようにした。Diagnostic payload は `field` / `value` / `ref_status` / `target_id` を持つ。
- 2026-05-27: Duplicate handling は source field ごとの diagnostic とした。`record_id` / `field` / `value` が修正対象 metadata を直接示すため、filter なし validation では同一 broken reciprocal relation が両側 source から報告されうる。
- 2026-05-27: Pass 3 tests として happy path、各 workflow relation field の unresolved target、wrong-kind / actual-kind mismatch、requirement-work item と work item-task の bidirectional mismatch、task source requirement mismatch、`task.depends_on` boundary、kind filter behavior、repository bootstrap workflow relation validation を追加した。
- 2026-05-27: Test result: `go test ./internal/designrecords ./internal/designrecordsmcp` passed。追加確認として `go test ./...` passed。
- 2026-05-27: Pass 1〜3 の implementation と automated tests が完了し、MVP 外 scope を暗黙実装していないことを確認した。TASK-MCP-003-05 が runtime verification を開始できるため、本 task を `done` に更新した。
