# Reference: Error model

- **id**: `spec:bpdsl.mcp.errors`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.mcp.overview`

## What this is

Error model for MCP tools: the distinction between a tool error and a diagnostic, the error code catalog, and the error payload shape.

> Source: V01-ADR-054

## MCP-level error vs. diagnostic

When MCP tool execution itself cannot proceed, the tool returns a tool error.

Examples:
- The project has not passed semantic validation.
- The selector is malformed.
- The target object does not exist.
- A transition is ambiguous because no guard was specified.

When tool execution succeeded but there is information worth noting, it goes into `diagnostics` instead.

Examples:
- A source line could not be obtained.
- An uncovered module became an implicit group.
- A `note` is absent.
- An optional piece of surrounding info is not yet implemented.

## Error code

MCP v1 defines the following error codes.

| code | meaning |
|---|---|
| `project_invalid` | Semantic build failed; cannot query. |
| `invalid_args` | Tool input JSON or input schema is invalid. |
| `invalid_selector` | Selector is malformed. |
| `invalid_change_payload` | `analyze_impact.change`'s kind/payload combination is invalid. |
| `not_found` | Target object does not exist. |
| `kind_mismatch` | `selector.kind` does not match the resolved kind. |
| `ambiguous` | Multiple candidates resolve and cannot be disambiguated. |
| `unsupported_object` | Object is out of query scope in v1. |
| `unsupported_detail` | `detail` value is unsupported. |
| `unsupported_direction` | `direction` value is unsupported. |
| `invalid_depth` | Traversal depth is outside the supported range. |
| `source_range_unavailable` | Cannot determine the source range for the object. |
| `internal_error` | Internal implementation error. |

Selector support behavior:

- If the selector shape is malformed, return `invalid_selector`.
- If the selector cannot be resolved, return `not_found`.
- If the selector resolves to multiple candidates, return `ambiguous`.
- If `selector.kind` does not match the resolved kind, return `kind_mismatch`.
- If the selector resolves but the target tool's selector support matrix marks it `no`, return `unsupported_object` in principle.
- `analyze_impact` does not turn an unsupported selector into a tool error — it returns a normal response with empty `impacts`, `coverage`, and an `unsupported_selector` diagnostic.

When `change.kind` is missing a required payload field, or the kind/payload combination is invalid, return `invalid_change_payload`.

Request option behavior:

- When an enum-like request option falls outside its specified value set, prefer a tool-specific code if one exists; otherwise return `invalid_args`.
- An unknown `direction` value returns `unsupported_direction`.
- An unknown `detail` value returns `unsupported_detail`.
- `get_reference_tree.depth` outside `0..4` returns `invalid_depth`.
- `get_source.fallback` outside `file` / `error` returns `invalid_args`.
- `source_range_unavailable` represents the same underlying condition: the source range for an object cannot be determined.
- `get_source(fallback=file)`, or `fallback` omitted, returns `source_range_unavailable` as a warning diagnostic.
- `get_source(fallback=error)` returns `source_range_unavailable` as a tool error.
- The surface / severity of `source_range_unavailable` is determined by the request's `fallback` option.
- Request-option default / omitted behavior is treated as each tool spec's contract, not as a DATA DSL default construct.

## Error payload

```json
{
  "error": {
    "code": "not_found",
    "message": "object not found: auth.task.missing_login",
    "selector": {
      "id": "auth.task.missing_login"
    },
    "diagnostics": []
  }
}
```

| field | required | content |
|---|---:|---|
| `code` | ✓ | Error code. |
| `message` | ✓ | Human-readable message. |
| `selector` | optional | The input selector. |
| `diagnostics` | optional | Related diagnostics. |

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.mcp.overview` | Parent overview; tool catalog. |
| `spec:bpdsl.mcp.schema` | Selector / ObjectRef / Diagnostic common schema referenced by error payloads. |
