# Contract: `suggest_next_record`

- **id**: `spec:drmcp.design_records_mcp.tools.suggest_next_record`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`suggest_next_record` is retired from the public Design Records MCP tool surface.

The operation is not public, invokable, or available through a compatibility alias.

The former ADR-number and V01 path-suggestion behavior is not part of the current contract.

## Request

`suggest_next_record` has no current request contract.

The former `kind` and `title` request is not accepted.

## Response

`suggest_next_record` has no current response contract.

The server tool catalog must omit `suggest_next_record`. The server must not route a `suggest_next_record` invocation to another operation implicitly.

No current operation returns the former `next_id`, `next_number`, `suggested_path`, or `existing_max_id` projection.

## Errors

`suggest_next_record` defines no tool-execution errors because the operation is not public or invokable.

Unknown-tool handling belongs to the MCP host boundary, not this retired operation contract.

## Related records

| ref | relation |
|---|---|
| `DRMCP-ADR-MCP-001` | Accepted removal from the current public tool surface. |
| `DRMCP-WORK-MCP-004` | Query and exact-retrieval tool-surface realignment owner. |
| `DRMCP-TASK-MCP-004-05` | Final retirement correction and validation owner. |
