# V01-TASK-MCP-006-03: Workflow metadata strictness contract を spec / guidance に反映する

- **id**: V01-TASK-MCP-006-03
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-MCP-006
- **source_requirement**: V01-REQ-MCP-006
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-006-02
- **outputs**:
  - Design Records MCP schema / tools spec updates
  - Requirement / work item / task authoring guidance updates

## Goal

V01-TASK-MCP-006-02 で確定した workflow artifact metadata validation strictness contract を、Design Records MCP の public spec と authoring guidance に反映する。

## Scope

- 更新対象候補は `docs/spec/design-records-mcp/schema.md`、`docs/spec/design-records-mcp/tools.md`、`docs/guides/requirement-authoring.md`、`docs/guides/work-item-authoring.md`、`docs/guides/task-authoring.md` とする。
- README 系 wording cleanup は V01-REQ-MCP-009 scope とし、本 task では扱わない。
- V01-ADR-092 は起票時点の decision snapshot として扱い、現在仕様は spec 側に反映する。
- V01-REQ-MCP-003 / V01-WORK-MCP-003 の完了済み scope は再オープンしない。

## Work

1. V01-TASK-MCP-006-02 の decision output を読む。
2. Schema spec に required field / empty value / date / diagnostic category contract を反映する。
3. Tools spec の `validate_records` response / diagnostic wording を同期する。
4. Requirement / work item / task authoring guidance の required metadata wording を同期する。
5. Spec / guidance 間で、required / optional / empty allowed の表現が矛盾していないことを確認する。

## Expected output

- Public spec updates。
- Authoring guidance updates。
- Implementation task に渡す diagnostic / validation contract summary。

## Completion condition

V01-TASK-MCP-006-04 が実装可能な形で、public spec と authoring guidance が同期している。

## Verification

- `docs/spec/design-records-mcp/schema.md` に workflow artifact required metadata strictness を反映した。
- `docs/spec/design-records-mcp/tools.md` に `validate_records` diagnostic category / field contract を反映した。
- `docs/guides/requirement-authoring.md`、`docs/guides/work-item-authoring.md`、`docs/guides/task-authoring.md` に required metadata validation wording を反映した。
- README 系は V01-REQ-MCP-009 の cleanup scope として分離し、本 task では更新しない。

## Evidence

Updated public spec:

- `docs/spec/design-records-mcp/schema.md`
  - Required scalar field は presence + non-empty が必要。
  - Required list field は presence が必要。
  - Empty required list は artifact-specific non-empty rule がない限り valid。
  - Empty list item は validation error とし、metadata diagnostic category は `empty_required_metadata`。
  - Workflow `date` は strict `YYYY-MM-DD` format。
  - `missing_required_metadata` / `empty_required_metadata` / `invalid_metadata_value` を diagnostic category に追加。
  - Workflow metadata diagnostic は `field` を必須、入力 value が存在する場合は `value` も返す。
- `docs/spec/design-records-mcp/tools.md`
  - `validate_records` diagnostic category list に metadata strictness categories を追加。
  - Metadata diagnostics の severity と field / value response contract を追加。

Updated authoring guidance:

- `docs/guides/requirement-authoring.md`
  - `id` / `status` / `date` は required scalar metadata。
  - `source_refs` / `work_items` は required list metadata、empty list allowed、empty item は `empty_required_metadata`。
- `docs/guides/work-item-authoring.md`
  - `id` / `status` / `date` / `source_requirement` は required scalar metadata。
  - `impact_refs` / `tasks` は required list metadata、empty list allowed、empty item は `empty_required_metadata`。
- `docs/guides/task-authoring.md`
  - `id` / `status` / `date` / `work_item` / `source_requirement` / `estimate` は required scalar metadata。
  - `depends_on` / `outputs` は required list metadata、empty list allowed、empty item は `empty_required_metadata`。

Implementation input for V01-TASK-MCP-006-04:

- Parser needs presence tracking for required metadata fields.
- Parser or validation model needs raw workflow `date` value for `requirement` / `work_item` / `task`.
- Validator needs `missing_required_metadata`, `empty_required_metadata`, and `invalid_metadata_value` diagnostics.
- Existing relation diagnostics remain responsible for non-empty malformed / unresolved / mismatched relation targets.
- `fixture_pending` remains valid and should get explicit regression coverage.

Review follow-up:

- Codex review result: OK with minor fixes.
- Fixed ambiguous empty list item wording by mapping empty list item diagnostics to `empty_required_metadata` in schema and authoring guides.
- Replaced the README synchronization scope wording with explicit V01-REQ-MCP-009 separation.
- Updated `last_updated` front matter in `docs/spec/design-records-mcp/schema.md` and `docs/spec/design-records-mcp/tools.md` to `2026-06-01`.
