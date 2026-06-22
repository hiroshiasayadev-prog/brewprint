# Contract: `list_objects`

- **id**: `spec:bpdsl.mcp.tools.list_objects`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.mcp.overview`
- **contract_class**: `interface`

## What this is

`list_objects` returns a list of semantic objects in the project.

Returned object kinds: `node`, `view`, `transition`, `field`. `asset` / `file` / `primitive` are not listed by `list_objects` — they are handled as reference targets, selector targets, or file-inspect targets.

`list_objects` is an exploration tool — it does not return detailed signature / references / inspect information for each object. When detail is needed, call `get_signature` / `get_references` / `inspect` using the returned `object` / `kind` / `id` as a selector.

> Source: V01-ADR-054

## Request

```json
{
  "object": "node",
  "kind": "task",
  "module": "order",
  "file": "order/task/checkout.yaml"
}
```

| field | required | content |
|---|---:|---|
| `object` | optional | `node` / `view` / `transition` / `field`. |
| `kind` | optional | Object-dependent kind filter. Value set follows the object-dependent kind vocabulary in [`spec:bpdsl.mcp.schema`](../schema.md). |
| `module` | optional | Module path, e.g. `order`, `payment.webhooks`. |
| `file` | optional | FileID. |

When `object` is omitted, all listable objects across `node` / `view` / `transition` / `field` are targeted. `kind` is a filter dependent on the given `object`. If `kind` is specified while `object` is omitted, only listable objects matching that `kind` across all object types are returned — e.g. `kind: task` matches `node: task`, `kind: api_table` matches `view: api_table`. An unknown `object` or `kind` is an `invalid_args` tool error.

## Response

```json
{
  "objects": [
    {
      "object": "node",
      "kind": "task",
      "id": "order.task.checkout",
      "qualified_id": "order.task.checkout",
      "label": "checkout",
      "module": "order",
      "file": "order/task/checkout.yaml",
      "source": { "file": "order/task/checkout.yaml" }
    }
  ],
  "diagnostics": []
}
```

## Errors

| code | condition |
|---|---|
| `invalid_args` | Unknown `object` or `kind` value. |

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.mcp.overview` | Parent overview; tool catalog and selection guidance. |
| `spec:bpdsl.mcp.schema` | Object-dependent kind vocabulary used by the `kind` filter. |
| `spec:bpdsl.mcp.errors` | Error code catalog. |
