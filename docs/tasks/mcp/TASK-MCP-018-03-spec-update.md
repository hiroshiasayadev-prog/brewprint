# TASK-MCP-018-03: Update spec and guidance for strict fields-required create contract

- **id**: TASK-MCP-018-03
- **status**: done
- **date**: 2026-06-03
- **work_item**: WORK-MCP-018
- **source_requirement**: REQ-MCP-019
- **estimate**: 1d
- **depends_on**:
  - TASK-MCP-018-02
- **outputs**:
  - Updated SPEC-design-records-mcp-tools contract text
  - Updated authoring guidance for strict create form

## Goal

`propose_record_create` の strict fields-required create contract、section-only body source、`fields + body_cache_id` retry form、legacy-free create boundary を spec / authoring guidance に反映する。

## Work

- `TASK-MCP-018-01` / `TASK-MCP-018-02` の review / impact result を前提に、valid / invalid input combinations を spec に反映する。
- `fields` を schema-level required として明文化する。
- `fields + body_cache_id` を `fields + body` retry form として明文化する。
- `body` / `body_cache_id` は `fields` と組み合わせた section-only content source として扱うことを明記する。
- `body`-only / `body_cache_id`-only create は invalid と明記する。
- `body + body_cache_id` と `fields + body + body_cache_id` が invalid であることを維持して明記する。
- invalid request でも submitted `body` が string として受け取れている場合、新しい `body_cache` を返してよいことを明記する。
- requirement / work item / task / ADR authoring guidance の MCP create note と矛盾がないか確認する。

## Done condition

- `SPEC-design-records-mcp-tools` に strict `propose_record_create` contract と retry behavior が反映されている。
- Authoring guidance が new code に `fields`, `fields + body`, `fields + body_cache_id` だけを使わせる内容になっている。
- 旧 create mode が invalid boundary として明確化されている。

## Verification

- Design Records MCP validation が spec / workflow artifact metadata に対して問題を出さない。
- Spec / guidance / REQ-MCP-019 / WORK-MCP-018 の boundary が矛盾していないことを確認する。

## Evidence
- `docs/spec/design-records-mcp/tools.md` を strict `propose_record_create` contract に更新した。
- `docs/spec/design-records-mcp/schema.md` の body cache model / authoring diagnostic wording を strict contract に更新した。
- Spec review で stale wording が 2 件見つかったため修正した。
  - `tools.md` Error handling の `invalid_request` 例示を required `fields` 省略に変更した。
  - `schema.md` `invalid_body_source` から旧 fields/body conflict wording を削除した。
- Follow-up review confirmed both blocking findings are resolved and `TASK-MCP-018-04` can proceed.
- Design Records MCP validation for `spec` passed with no diagnostics after the fix.
