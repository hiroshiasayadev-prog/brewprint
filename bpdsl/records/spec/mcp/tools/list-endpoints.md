# Contract: `list_endpoints`

- **id**: `spec:bpdsl.mcp.tools.list_endpoints`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.mcp.overview`
- **contract_class**: `interface`

## What this is

`list_endpoints` returns the endpoint list based on an API Table view YAML.

Rather than simply enumerating `task(endpoint=true)`, it returns the full path following V01-ADR-028's route-composition rule.

> Source: V01-ADR-028

## Request

```json
{
  "api_table_id": "ec_api"
}
```

| field | required | content |
|---|---:|---|
| `api_table_id` | optional | API Table view ID. When omitted, returns all API Tables in the project. |

When `api_table_id` is specified, the tool targets only that API Table view and the response's `tables[]` contains exactly that one entry. When omitted, the tool targets every API Table view in the project and `tables[]` contains all of them. This response shape is MCP tool-contract runtime behavior, not a DATA DSL default/union construct.

## Response

`list_endpoints` always returns `tables[]`. When `api_table_id` is specified, `tables[]` contains exactly that one API Table. When omitted, `tables[]` contains every API Table in the project. A top-level single-table response is not part of the MCP v1 contract.

```json
{
  "tables": [
    {
      "id": "ec_api",
      "http_root_path": "/api",
      "sections": [
        {
          "module": "auth",
          "include_submodules": true,
          "endpoints": [
            {
              "method": "POST",
              "path": "/api/login",
              "leaf_path": "login",
              "task": "auth.task.login",
              "params": "auth.model.login_request",
              "returns": "auth.model.token",
              "source": {
                "file": "auth/task/login.yaml"
              }
            }
          ]
        }
      ]
    }
  ],
  "diagnostics": []
}
```

In this example, the section's anchor module is `auth`, so per V01-ADR-028's route-composition rule, the relative module path from the section's anchor module is empty. The full path is therefore `/api/login`, not `/api/auth/login`. To get `/api/auth/login`, either set `http_root_path: /api/auth` on the API Table view, or make the section's anchor module a parent module.

### endpoint object

| field | required | content |
|---|---:|---|
| `method` | ✓ | HTTP method. |
| `path` | ✓ | Full path composed by the API Table view. |
| `leaf_path` | ✓ | Task-side leaf path. Derived from `task.id` when omitted. |
| `task` | ✓ | Endpoint task QualifiedID. |
| `params` | optional | Request model QualifiedID. |
| `returns` | optional | Response model QualifiedID. |
| `source` | optional | SourceLocation of the endpoint task. |

## Errors

| code | condition |
|---|---|
| `not_found` | `api_table_id` does not resolve to an existing API Table view. |
| `invalid_args` | Malformed input. |

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.mcp.overview` | Parent overview; tool catalog and selection guidance. |
| `spec:bpdsl.mcp.tools.get_signature` | `signature.endpoint.leaf_path` is task-side only; this tool owns the composed full path. |
| `spec:bpdsl.mcp.tools.inspect` | `inspect(view: api_table)` uses the same route-composition rule for context retrieval. |
