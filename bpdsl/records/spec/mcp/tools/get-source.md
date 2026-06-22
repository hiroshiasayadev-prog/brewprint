# Contract: `get_source`

- **id**: `spec:bpdsl.mcp.tools.get_source`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.mcp.overview`
- **contract_class**: `interface`

## What this is

`get_source` returns the YAML source snippet corresponding to a target semantic object.

Returns: object identity, source file/range, YAML snippet, a diagnostic indicating the reason when a fallback occurred.

Does not return: the raw YAML AST, the full unresolved structure before semantic build, the content of files outside the project, renderer output.

Per V01-ADR-054, `get_source` is not a raw-YAML-AST-exposing API — it is treated as source auxiliary info attached to a semantic object on `ResolvedProject`.

## Request

```json
{
  "selector": {
    "id": "auth.task.login"
  },
  "fallback": "file"
}
```

| field | required | content |
|---|---:|---|
| `selector` | ✓ | Object selector. |
| `fallback` | optional | `file` / `error`. Defaults to `file` when omitted. |

`fallback=file`, or omitted, returns the whole YAML file with the same FileID when the object-level range cannot be determined, with `fallback: "file"` and a `source_range_unavailable` warning in `diagnostics[]`. `fallback=error` returns a `source_range_unavailable` tool error instead of the file-fallback response when the object-level range cannot be determined. A `fallback` value outside `file` / `error` is an `invalid_args` tool error. This default and fallback branching is MCP tool-contract runtime behavior, not a DATA DSL default/fallback construct.

## Response

```json
{
  "object": {
    "object": "node",
    "kind": "task",
    "id": "auth.task.login",
    "qualified_id": "auth.task.login",
    "label": "login",
    "file": "auth/task/login.yaml"
  },
  "source": {
    "file": "auth/task/login.yaml",
    "line": 3,
    "column": 5,
    "end_line": 18,
    "end_column": 1
  },
  "snippet": {
    "language": "yaml",
    "text": "  - id: login\n    type: task\n    ..."
  },
  "diagnostics": []
}
```

| field | required | content |
|---|---:|---|
| `object` | ✓ | Target ObjectRef. |
| `source` | ✓ | SourceLocation. May contain only `file` if line/column are unavailable. |
| `snippet` | ✓ | `language: yaml` plus snippet text. |
| `fallback` | optional | `file` if a fallback occurred. Absent when `fallback=error` (no fallback response is returned in that case). |
| `diagnostics` | ✓ | Diagnostic list. |

### Selector support

`get_source` targets semantic objects queryable in MCP v1.

Objects for which the initial implementation returns a snippet range:

- `node` / private sub-node: the matching item within `nodes[]`.
- `field`: the matching item within the parent model's `fields[]`.
- `transition`: the matching item within `transitions[]`.
- `asset`: the producer node's `returns` block. May fall back to the producer node if the asset itself can't be pinpointed.
- `view`: the whole view file.
- `file`: the whole file.

Implementations that cannot obtain line/column, or cannot pinpoint a local range for a supported object, may use `fallback=file` to return the whole file at the same FileID. In that case, include in `diagnostics[]`:

```json
{
  "severity": "warning",
  "code": "source_range_unavailable",
  "file": "auth/task/login.yaml",
  "message": "source range is unavailable; returned whole file"
}
```

> Source: V01-ADR-054 §Decision §5

## Errors

| code | condition |
|---|---|
| `invalid_args` | `fallback` value outside `file` / `error`. |
| `source_range_unavailable` | Object-level range cannot be determined and `fallback=error`. |
| `unsupported_object` | Selector resolves to an object/kind marked `no` in the selector support matrix. |
| `not_found` | Selector does not resolve to any object. |

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.mcp.overview` | Parent overview; tool catalog and selection guidance. |
| `spec:bpdsl.mcp.schema` | Selector support matrix, ObjectRef, SourceLocation shapes. |
| `spec:bpdsl.mcp.errors` | Error code catalog, `source_range_unavailable` severity/surface rules. |
