# V01-TASK-MCP-005-02: MCP tool / schema contract design

- **id**: V01-TASK-MCP-005-02
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-MCP-005
- **source_requirement**: V01-REQ-MCP-005
- **depends_on**:
  - V01-TASK-MCP-005-01
- **outputs**:
  - docs/spec/design-records-mcp/tools.md
  - docs/spec/design-records-mcp/schema.md

## Goal

Define the MCP request / response contract for authoring guidance discovery and retrieval before implementation.

## Work

- Define `list_authoring_guides` request shape.
- Define `list_authoring_guides` response shape.
- Define `get_authoring_guidance` request shape.
- Define `get_authoring_guidance` response shape.
- Confirm guide ID is the public reference and source path remains internal.
- Confirm guide content response returns Markdown content without converting the guide into a Design Records record.
- Identify the exact spec sections that must be updated in `docs/spec/design-records-mcp/tools.md` and `docs/spec/design-records-mcp/schema.md`.

## Proposed Contract Direction

### `list_authoring_guides`

Request:

```text
{}
```

Response:

```text
guides:
  - id: <guide filename stem>
    title: <first H1 text>
    abstract: <contents of ## Abstract>
```

Response MUST NOT expose source file path.

### `get_authoring_guidance`

Request:

```text
id: <guide filename stem>
```

Response:

```text
id: <guide filename stem>
title: <first H1 text>
content: <Markdown content>
```

The tool resolves `id` to guide source internally. The source file path is not part of the public response contract.

## Boundary

### Included

- Tool contract design.
- Schema field naming.
- Path non-exposure rule.
- Missing guide ID behavior as a contract item.
- Spec update target identification.

### Excluded

- Go implementation.
- MCP registration wiring.
- Tests.
- Moving or deleting existing guides.
- Creating guide canonical record IDs.

## Done condition

- Tool request / response contracts are decided.
- Path non-exposure is explicit.
- Missing guide ID behavior is decided.
- Spec update targets are identified.
- Next task can update `tools.md` and `schema.md` without reopening the boundary.

## Verification

- Review contract against `V01-TASK-MCP-005-01` decisions.
- Confirm response fields only use guide ID, title, abstract, and Markdown content.
- Confirm no physical path is exposed as public contract.

## Evidence

- Decided `list_authoring_guides` request / response contract.
- Decided `get_authoring_guidance` request / response contract.
- Confirmed guide ID is the public reference and source path is not exposed.
- Confirmed `get_authoring_guidance` returns raw Markdown content.
- Confirmed missing guide ID returns `guide_not_found` tool error.
- Updated `docs/spec/design-records-mcp/tools.md` with the new tool contracts.
- Updated `docs/spec/design-records-mcp/schema.md` with the authoring guidance source model.
