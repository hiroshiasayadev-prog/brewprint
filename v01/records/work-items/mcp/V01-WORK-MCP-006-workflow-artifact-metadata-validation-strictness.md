# V01-WORK-MCP-006: Workflow artifact metadata validation strictness を判断・実現する

- **id**: V01-WORK-MCP-006
- **status**: done
- **date**: 2026-05-31
- **decision_date**: 2026-06-01
- **source_requirement**: V01-REQ-MCP-006
- **impact_refs**:
  - V01-ADR-092
  - SPEC-design-records-mcp-schema
  - SPEC-design-records-mcp-tools
  - V01-REQ-MCP-003
- **tasks**:
  - V01-TASK-MCP-006-01
  - V01-TASK-MCP-006-02
  - V01-TASK-MCP-006-03
  - V01-TASK-MCP-006-04
  - V01-TASK-MCP-006-05

## Goal

V01-REQ-MCP-006 を解消するため、Design Records MCP の workflow artifact metadata validation strictness を public contract、implementation、tests、runtime validation evidence まで一貫して揃える。

## Boundary

- 本 work item は workflow artifact metadata strictness の判断と実現フローを所有する。
- `fixture_pending` は現行 work item status の許可値として扱い、不正値への変更対象にしない。
- SELFHOST 関連 artifact は本 work item の直接編集対象にしない。
- V01-REQ-MCP-003 / V01-WORK-MCP-003 の完了済み scope は再オープンしない。必要な追加 validation は新規 follow-up として扱う。
- Orphan diagnostics、progress projection、workflow traversal、task dependency cycle / execution order projection、physical path relation support、`req:` / `work:` / `task:` semantic prefix support は対象外とする。

## Impact scope

| layer | expected handling |
|---|---|
| requirement | V01-REQ-MCP-006 の採否と strictness 範囲を判断する |
| decision / spec | 必要なら Design Records MCP schema / tools spec に diagnostic contract を追加する |
| implementation | `internal/designrecords` の parser / validation / tests を contract に合わせる |
| MCP transport | `validate_records` runtime schema / response が spec と一致することを確認する |
| docs | authoring guidance と public spec の status / required field wording を同期する |

## Task flow

1. `V01-TASK-MCP-006-01`: Current validator behavior と metadata field gap の evidence を整理する。
2. `V01-TASK-MCP-006-02`: Required metadata / empty value / date validation / diagnostic category の public contract を判断する。
3. `V01-TASK-MCP-006-03`: 採用 contract を Design Records MCP spec と authoring guidance に反映する。
4. `V01-TASK-MCP-006-04`: Parser / validator / MCP tests を実装する。
5. `V01-TASK-MCP-006-05`: Runtime validation と close review を実施する。

## Completion condition

以下のいずれかを満たしたとき、本 work item を `done` にできる。

1. Workflow artifact metadata strictness の採用範囲を決め、必要な spec / implementation / tests / runtime verification / close evidence が完了している。
2. 追加 validation を延期または不採用と判断し、理由、残す risk、V01-REQ-MCP-006 への結果反映、close evidence が完了している。

## Current blockers

None.

## Progress summary

- 2026-05-31: `V01-WORK-SELFHOST-001.status=fixture_pending` に対する `validate_records(kind="work_item")` の `ok:true` 観測を起点に調査した。`fixture_pending` は work item status の許可値であり、status enum bug ではないと判定した。
- 2026-05-31: 一方で、workflow artifact の required metadata / date / field value strictness は、V01-REQ-MCP-003 で完了した declared relation validation とは別の validation completeness issue として残るため、V01-REQ-MCP-006 / V01-WORK-MCP-006 として分離起票した。
- 2026-06-01: `V01-TASK-MCP-006-01`〜`V01-TASK-MCP-006-05` に分解し、まず current validator behavior と metadata field gap の evidence review から開始する形にした。
- 2026-06-01: `V01-TASK-MCP-006-01` で current validation gap review を完了した。
- 2026-06-01: `V01-TASK-MCP-006-02` で required metadata presence、required scalar empty validation、required list presence、empty list allowed、empty list item error、workflow date validation、metadata-specific diagnostic categories を採用した。work item status を `design_spec_pending` に更新した。
- 2026-06-01: `V01-TASK-MCP-006-03` で Design Records MCP schema / tools spec と requirement / work item / task authoring guidance に採用 contract を反映した。README 系は V01-REQ-MCP-009 の cleanup scope として分離し、work item status を `implementation_pending` に更新した。
- 2026-06-01: `V01-TASK-MCP-006-04` で parser / validator / tests への反映と implementation review を完了した。`go test -count=1 ./internal/designrecords ./internal/designrecordsmcp`、task range validation、work item validation が pass した。work item status を `verification_pending` に更新した。
- 2026-06-01: `V01-TASK-MCP-006-05` で runtime verification、build verification、close review を完了した。V01-REQ-MCP-006 の close evidence を記録し、work item status を `done` に更新した。

