# Contract: `get_record`

- **id**: `spec:drmcp.design_records_mcp.tools.get_record`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`get_record` is retired from the public Design Records MCP tool surface.

`get_records` is the sole exact-retrieval operation. A caller retrieves one record by sending one value in `refs`.

## Request

`get_record` has no current request contract.

The former single-record `id` request is not accepted through a compatibility alias.

Use this replacement request:

```json
{
  "refs": ["DRMCP-ADR-MCP-001"],
  "include_body": false
}
```

## Response

`get_record` has no current response contract.

The server tool catalog must omit `get_record`. The server must not route a `get_record` invocation to `get_records` implicitly.

The replacement response is defined by `spec:drmcp.design_records_mcp.tools.get_records`.

## Errors

`get_record` defines no tool-execution errors because the operation is not public or invokable.

Unknown-tool handling belongs to the MCP host boundary, not this retired operation contract.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.tools.get_records` | Sole public exact-retrieval operation. |
| `DRMCP-TASK-MCP-004-03` | Retirement and replacement contract owner. |
