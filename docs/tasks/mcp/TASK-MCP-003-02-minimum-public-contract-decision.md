# TASK-MCP-003-02: Workflow artifact MCP support の最小 public contract と ADR 要否を判断する

- **id**: TASK-MCP-003-02
- **status**: done
- **date**: 2026-05-26
- **work_item**: WORK-MCP-003
- **source_requirement**: REQ-MCP-003
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-003-01
- **outputs**:
  - workflow artifact MCP support の採用・延期・不採用判断
  - MVP public surface と非対象 scope の判断
  - ADR 起票要否の判断、または必要な ADR

## Goal

TASK-MCP-003-01 の evidence を基に、`REQ-*` / `WORK-*` / `TASK-*` を扱う最小 public contract を確定し、設計判断を ADR として残す必要があるか判断する。

## Work

- Existing record-oriented surface の extension と workflow 専用 surface の選択肢を比較する。
- `get_record(s)`、`resolve_reference`、relation validation の採否と一貫した MVP boundary を判断する。
- Orphan diagnostics と task status 由来 progress projection を MVP に含めるか、後続へ送るか判断する。
- Physical path 非対応、ID-as-ref、manual status copy 非導入の境界を維持する。
- 採用時は spec 更新範囲と implementation 前提を TASK-MCP-003-03 に渡す。

## Done condition

- 採用・延期・不採用のいずれかが、対象 surface ごとに根拠付きで確定している。
- ADR が必要なら起票・判断結果が存在し、不要なら不要理由が記録されている。
- 採用時は TASK-MCP-003-03 が開始可能である。

## Verification

- ADR-091 の workflow artifact boundary と矛盾しないことを確認する。
- 現行 MCP tool responsibility を不用意に重複させていないことを確認する。
- MVP 外へ送る capability が未確定のまま実装 task に流入しないことを確認する。

## Evidence

### 判断 draft

- TASK-MCP-003-01 の evidence と補正結果を入力として、workflow artifact MCP support の最小 public contract を `ADR-092: Design Records MCP workflow artifact record and relation boundary` として `proposed` で起票した。
- ADR draft は、`requirement` / `work_item` / `task` を既存 record-oriented surface の record kind として追加し、`list_records` / `get_record` / `get_records` の対象に含める方針を採用候補として記録した。
- ADR draft は、`resolve_reference` の supported record ID-as-ref に `REQ-*` / `WORK-*` / `TASK-*` を追加し、現行 authoring convention の workflow ID form を MCP-supported grammar として採用する方針を記録した。Relation は path や ID 文字列構造から推測せず、明示 metadata の ID-as-ref として扱う。
- ADR draft は、workflow relation の参照先存在確認に加え、`REQ.work_items` と `WORK.source_requirement`、`WORK.tasks` と `TASK.work_item` を双方向に照合し、`TASK.source_requirement` と parent work item の source requirement の一致を検査する MVP boundary を記録した。
- ADR draft は、ADR-087 の investigation canonical reference validation semantics を維持したまま、`source_refs` / `follow_up_results` / `follow_up_candidates` が `REQ-*` / `WORK-*` を扱えるようにする boundary を採用候補として記録した。`TASK-*` は workflow relation と direct resolver input として support する一方、investigation metadata relation には追加しない。
- ADR draft は、orphan diagnostics、task status 由来 progress projection、workflow 専用 traversal tool、dependency cycle / execution projection、investigation `trigger` / optional `related_*` の validation 拡張を MVP 外へ送る方針を記録した。

### ADR 要否

- 新 ADR が必要である。
- 理由は、ADR-087 / ADR-090 の既存 public record / retrieval / resolver / validation boundary に新しい public record kind と canonical ID-as-ref、relation integrity validation を追加する判断であり、単なる spec clarification ではないためである。

### Decision result

- `ADR-092` は Codex review で major 二件・minor 一件の修正要請を受け、双方向 relation integrity、investigation metadata から `TASK-*` を除外する boundary、workflow ID grammar の public contract 化を draft に反映した。
- 再 review で `OK to proceed to acceptance` が確認され、`ADR-092` を `accepted` とした。
- `REQ-MCP-003` は採用結果を反映して `accepted` に進め、`WORK-MCP-003` は spec 反映工程へ進める。

### Verification evidence

- Record-oriented surface、direct resolver input、declared relation integrity validation、investigation metadata boundary、MVP 外 scope が `ADR-092` に確定している。
- Orphan diagnostics、progress projection、workflow 専用 traversal、investigation metadata での `TASK-*` support を implementation task に流入させない境界を確認した。
- `TASK-MCP-003-03` は accepted decision に基づく spec 更新へ開始可能である。
