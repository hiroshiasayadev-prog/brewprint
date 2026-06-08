# V01-TASK-MCP-005-03: MCP tool / schema spec update

- **id**: V01-TASK-MCP-005-03
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-MCP-005
- **source_requirement**: V01-REQ-MCP-005
- **depends_on**:
  - V01-TASK-MCP-005-02
- **outputs**:
  - docs/spec/design-records-mcp/tools.md
  - docs/spec/design-records-mcp/schema.md

## Goal

Reflect the authoring guidance discovery / retrieval contract in Design Records MCP specs.

## Work

- Add `list_authoring_guides` to the Design Records MCP tool set.
- Add `get_authoring_guidance` to the Design Records MCP tool set.
- Document request / response shape for both tools.
- Document guide source model in schema spec.
- Document path non-exposure and guide ID based resolution.
- Document missing guide ID error handling.

## Done condition

- `docs/spec/design-records-mcp/tools.md` contains the new tool contracts.
- `docs/spec/design-records-mcp/schema.md` contains the authoring guidance source model.
- Both specs state that guide source path is not public response contract.
- Missing guide ID behavior is represented as `guide_not_found` tool error.

## Verification

- Checked `tools.md` for tool set, request / response, and error handling updates.
- Checked `schema.md` for authoring guidance source model and non-record boundary.

## Evidence

- Updated `docs/spec/design-records-mcp/tools.md`:
  - `last_updated: 2026-05-31`
  - added `list_authoring_guides`
  - added `get_authoring_guidance`
  - added `guide_not_found`
- Updated `docs/spec/design-records-mcp/schema.md`:
  - `last_updated: 2026-05-31`
  - added authoring guidance source model
  - clarified guide is not a Design Records record kind
