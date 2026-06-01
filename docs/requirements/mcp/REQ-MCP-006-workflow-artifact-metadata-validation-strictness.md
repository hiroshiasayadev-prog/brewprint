# REQ-MCP-006: Workflow artifact metadata validation strictness が必要

- **id**: REQ-MCP-006
- **status**: accepted
- **date**: 2026-05-31
- **decision_date**: 2026-06-01
- **source_refs**:
  - ADR-092
  - SPEC-design-records-mcp-schema
  - SPEC-design-records-mcp-tools
  - REQ-MCP-003
- **work_items**:
  - WORK-MCP-006

## 要求

Design Records MCP の `validate_records` は、workflow artifact の relation integrity だけでなく、workflow artifact metadata の required field presence と field value strictness を、spec / authoring guidance と一致する形で検証できる必要がある。

特に、requirement / work item / task の metadata について、どの field が必須で、空値を許容するか、field value の形式不正をどの diagnostic として返すかを public contract と implementation / tests で揃える必要がある。

## 発見根拠

`WORK-SELFHOST-001` の `status` に `fixture_pending` が入り、Design Records MCP の `validate_records(kind="work_item")` は `ok: true` を返した。

調査の結果、`fixture_pending` は非標準値ではなく、以下の現行 source で work item status として許容されている。

- `docs/work-items/README.md`
- `docs/spec/design-records-mcp/schema.md`
- `internal/designrecords/types.go`
- `internal/designrecords/validation.go`

したがって、この観測は `fixture_pending` を不正値として通した status enum bug ではない。一方で、同じ確認により、workflow artifact metadata strictness の範囲が relation validation と status enum validation に寄っており、spec が required とする workflow metadata field の presence / value validation が十分に閉じていないことが分かった。

## Current behavior summary

- `validate_records` は kind 別 status enum を検査し、`invalid_status_for_kind` を返す実装を持つ。
- `work_item` status として `fixture_pending` は許可値である。
- `validate_records` は workflow relation field の target kind / existence / reciprocal integrity を検査する。
- Workflow artifact の `date` は parser で読まれるが、record response field には持たず、workflow artifact 側では日付形式 diagnostic も定義されていない。
- `docs/spec/design-records-mcp/schema.md` は workflow artifact metadata の required field を定義しているが、missing required metadata、空値、日付形式、list field の presence と empty list の扱いを diagnostic contract として十分に分解していない。

## Candidate capability

後続判断では、少なくとも以下を検討する。

- requirement / work item / task の required metadata presence validation
- `date` field の形式 validation を workflow artifact に適用するか
- required list field が存在しない場合と、存在して空 list の場合を区別するか
- scalar relation field の空値を missing metadata として扱うか、relation validation の非対象として扱うか
- 新規 diagnostic category を追加するか、既存 `invalid_status_for_kind` / `invalid_workflow_id` / `invalid_workflow_relation_target` 等へ寄せるか
- `fixture_pending` を含む許可済み work item status が valid であること、および許可外 status が invalid であることを regression test で明示するか

## Boundary

- 本 requirement は `WORK-SELFHOST-001` または `TASK-SELFHOST-001-02` を修正しない。
- `fixture_pending` を work item status から削除する要求ではない。
- REQ-MCP-003 / WORK-MCP-003 の完了済み workflow artifact MCP support を再オープンしない。
- Orphan diagnostics、progress projection、workflow traversal、task dependency cycle / execution order projection は本 requirement の主対象ではない。
- Workflow artifact 間 relation の physical path support、`req:` / `work:` / `task:` semantic prefix support は導入しない。

## Progress

- 2026-06-01: `accepted`。WORK-MCP-006 で扱い、current validation gap review から開始する。

本件は REQ-MCP-003 の target chain 完了を否定する blocker ではない。REQ-MCP-003 は workflow artifact record / resolver / declared relation validation の最小 public contractを完了済みである。

ただし、required workflow metadata と field value strictness は別の validation completeness issue として追跡する必要があるため、REQ-MCP-006 として分離し、WORK-MCP-006 で strictness 範囲の判断と実現を進める。

## Close evidence

- 2026-06-01: `WORK-MCP-006` で requirement を解消した。
- Public contract:
  - `docs/spec/design-records-mcp/schema.md` に workflow artifact required metadata strictness と metadata diagnostic category を反映した。
  - `docs/spec/design-records-mcp/tools.md` に `validate_records` diagnostic category / response field contract を反映した。
  - `docs/guides/requirement-authoring.md`、`docs/guides/work-item-authoring.md`、`docs/guides/task-authoring.md` に required metadata validation wording を反映した。
- Implementation / tests:
  - Parser / validation model に workflow artifact required metadata presence tracking、empty list item tracking、raw workflow `date` tracking を追加した。
  - Validator に `missing_required_metadata` / `empty_required_metadata` / `invalid_metadata_value` を追加した。
  - MCP diagnostic response で `field` と explicit empty `value` を返せることを test で確認した。
  - `fixture_pending` が valid work item status であることを regression test で確認した。
- Verification:
  - `go test -count=1 ./internal/designrecords ./internal/designrecordsmcp` passed.
  - `validate_records(kind="task", id_range=TASK-MCP-006-01..TASK-MCP-006-05)` returned `ok: true`, `diagnostics: null`.
  - `validate_records(kind="work_item")` returned `ok: true`, `diagnostics: null`.
  - Build verification passed after adding the local Windows build helper script.
- Scope boundaries preserved:
  - `diagnostics:null` versus `diagnostics:[]` remains out of scope.
  - README / legacy authoring guide cleanup is tracked separately by `REQ-MCP-009`.
  - Orphan diagnostics, traversal, progress projection, cycle/order projection, physical path workflow relations, and `req:` / `work:` / `task:` semantic prefixes remain out of scope.

