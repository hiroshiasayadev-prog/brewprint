# TASK-MCP-018-01: Review current create contract and retry gap

- **id**: TASK-MCP-018-01
- **status**: done
- **date**: 2026-06-03
- **work_item**: WORK-MCP-018
- **source_requirement**: REQ-MCP-019
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Current contract / legacy mode / retry gap review result
  - Decision input for spec and implementation task

## Goal

`propose_record_create` の現行 contract、legacy full-record create mode、`fields + body_cache_id` retry gap を確認し、REQ-MCP-019 の実装前提を固める。

## Work

- `SPEC-design-records-mcp-tools` の `propose_record_create` input contract を確認する。
- `REQ-MCP-014` / `WORK-MCP-014` の close evidence と現行実装の前提を確認する。
- `REQ-MCP-015` / `WORK-MCP-012` の failed propose body_cache behavior を確認する。
- `body-only` / `body_cache_id-only` legacy create の残し方について、廃止・warning・compatibility boundary の選択肢を整理する。
- `fields` required化を schema required にするか validation required にするか、実装影響を分類する。

## Done condition

- valid / invalid / deprecated input combinations が一覧化されている。
- `fields + body_cache_id` を valid retry form にするための実装影響が説明されている。
- spec / guidance / implementation task に渡す判断材料が揃っている。

## Verification

- 関連 requirement / work item / spec / current implementation の少なくとも contract-relevant sections を確認する。
- 不明点や仕様判断待ちがある場合は、TASK-MCP-018-02 に渡せる形で明示する。

## Evidence
- Current `propose_record_create` contract and retry gap were reviewed before implementation.
- Review identified that `fields + body_cache_id` must become a valid retry form for `fields + body` create.
- User decision finalized the strict contract:
  - `fields` is required.
  - `fields`, `fields + body`, and `fields + body_cache_id` are valid create forms.
  - `body`-only / `body_cache_id`-only create are invalid.
  - invalid request with submitted string `body` should preserve the body in a new `body_cache`.
- The review output was carried forward into `TASK-MCP-018-02` impact assessment and subsequent spec / implementation tasks.
