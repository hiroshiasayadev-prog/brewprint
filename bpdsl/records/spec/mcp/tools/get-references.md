# Contract: `get_references`

- **id**: `spec:bpdsl.mcp.tools.get_references`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.mcp.overview`
- **contract_class**: `interface`

## What this is

`get_references` returns a target object's direct references.

MCP v1 returns direct references only.

> Source: V01-ADR-049

## Request

```json
{
  "selector": {
    "id": "auth.task.login"
  },
  "direction": "both",
  "kinds": ["reads", "writes", "transition_action"]
}
```

| field | required | content |
|---|---:|---|
| `selector` | ✓ | Object selector. |
| `direction` | optional | `out` / `in` / `both`. Defaults to `out` when omitted. |
| `kinds` | optional | Reference-kind filter. All kinds when omitted. |

When `direction` is omitted, the tool uses `out`, and the response's `direction` field reports the `out` actually used. A `direction` value outside `out` / `in` / `both` is an `unsupported_direction` tool error. This default is MCP tool-contract runtime behavior, not a DATA DSL default construct.

The selector's object/kind support range is governed by the selector support matrix in [`spec:bpdsl.mcp.schema`](../schema.md). If `get_references` receives a selector marked `no` in the matrix, it returns `unsupported_object` tool error in principle. A `limited` selector may have its returned references restricted per that matrix and this tool's tool-specific section below.

## Response

```json
{
  "object": {
    "object": "node",
    "kind": "task",
    "id": "auth.task.login"
  },
  "direction": "both",
  "depth": 1,
  "references": [
    {
      "kind": "param_model",
      "direction": "out",
      "from": { "object": "node", "kind": "task", "id": "auth.task.login" },
      "to": { "object": "node", "kind": "model", "id": "auth.model.credential" }
    },
    {
      "kind": "reads",
      "direction": "out",
      "from": { "object": "node", "kind": "task", "id": "auth.task.login" },
      "to": { "object": "node", "kind": "store", "id": "auth.store.user_db" }
    },
    {
      "kind": "transition_action",
      "direction": "in",
      "from": {
        "object": "transition",
        "kind": "transition",
        "id": "auth/state.yaml#idle:login_submitted",
        "state_file": "auth/state.yaml",
        "from": "idle",
        "on": "login_submitted",
        "to": "loading",
        "action": "auth.task.login"
      },
      "to": { "object": "node", "kind": "task", "id": "auth.task.login" }
    }
  ],
  "diagnostics": []
}
```

`depth` always returns `1`. MCP v1's `get_references` request has no `depth` field — transitive reference traversal is handled by the separate [`get_reference_tree`](get-reference-tree.md) tool per V01-ADR-055.

## Errors

| code | condition |
|---|---|
| `unsupported_object` | Selector resolves to an object/kind marked `no` in the selector support matrix. |
| `unsupported_direction` | `direction` value outside `out` / `in` / `both`. |
| `not_found` | Selector does not resolve to any object. |
| `ambiguous` | Selector resolves to multiple candidates. |
| `kind_mismatch` | Resolved kind does not match `selector.kind`. |

## Selector support

Supported selectors:

- `node: task`
- `node: model`
- `node: store`
- `node: state`
- `node: event`
- `node: actor`
- `view: sequence_diagram`
- `transition`
- `field` / `model_field`
- `file: node` (limited)
- `file: state_file`
- `asset`
- private sub node

Unsupported selectors:

- `primitive`
- `view: api_table`
- `view: er_diagram`
- `file: sequence_diagram`
- `file: api_table`
- `file: er_diagram`
- `file: render_index`

`file: node` is limited. It returns references whose source or target is a node defined in that file. It does not return raw flow wiring, render output links, or references for non-node YAML entries in the file.

`file: state_file` returns references owned by the state file, including transitions defined in that file and their `transition_from` / `transition_event` / `transition_to` / `transition_action` references when resolvable.

For unsupported selectors, `get_references` returns an `unsupported_object` tool error.

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.mcp.overview` | Parent overview; tool catalog and selection guidance. |
| `spec:bpdsl.mcp.schema` | Selector support matrix, Reference / Reference kind shapes. |
| `spec:bpdsl.mcp.errors` | Error code catalog. |
| `spec:bpdsl.mcp.tools.get_reference_tree` | Handles transitive traversal beyond direct references. |
