# Contract: `get_reference_tree`

- **id**: `spec:bpdsl.mcp.tools.get_reference_tree`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.mcp.overview`
- **contract_class**: `interface`

## What this is

`get_reference_tree` performs a depth-limited BFS traversal of direct references starting from a target object as root.

Although the tool is named "tree," the return shape is not a pure tree — it is a bounded reference graph composed of `nodes[]` / `edges[]`.

Returns: root object, traversal direction/depth, reached objects, traversed references, truncation info, diagnostics.

Does not return: per-change-kind impact severity, recommended action, renderer output mapping, flow wiring references, the raw YAML AST.

Change-kind-aware impact judgment is handled by [`analyze_impact`](analyze-impact.md).

> Source: V01-ADR-049, V01-ADR-054, V01-ADR-055

## Request

```json
{
  "selector": {
    "object": "transition",
    "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
  },
  "direction": "out",
  "depth": 1,
  "kinds": ["transition_from", "transition_event", "transition_to", "transition_action"],
  "max_nodes": 200,
  "max_edges": 500
}
```

| field | required | content |
|---|---:|---|
| `selector` | ✓ | Traversal root object. |
| `direction` | ✓ | `out` / `in` / `both`. |
| `depth` | ✓ | `0..4`. |
| `kinds` | optional | Reference-kind filter for traversal / return. |
| `max_nodes` | optional | Node return cap. Defaults to `200`. |
| `max_edges` | optional | Edge return cap. Defaults to `500`. |

`depth < 0` or `depth > 4` is an `invalid_depth` tool error. This range is an MCP tool-contract runtime constraint, not a DATA DSL generic numeric-range construct. `direction` is required, since it must not implicitly default the search scope; a value outside `out` / `in` / `both` is an `unsupported_direction` tool error.

When `kinds` is specified, only the given reference kinds are traversed as path edges and included in `edges[]`. Objects reachable only via a non-listed kind are not reached.

The selector's object/kind support range is governed by the selector support matrix in [`spec:bpdsl.mcp.schema`](../schema.md). If `get_reference_tree` receives a selector marked `no` in the matrix, it returns `unsupported_object` tool error in principle. A `limited` selector may have its traversal-root / edge-expansion behavior restricted per that matrix and §Selector support below.

## Response

```json
{
  "root": {
    "object": "transition",
    "kind": "transition",
    "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']",
    "file": "order/state.yaml",
    "local_id": "processing:payment_webhook_received"
  },
  "direction": "out",
  "depth": 1,
  "nodes": [
    {
      "object": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
      },
      "depth": 0,
      "via": []
    },
    {
      "object": {
        "object": "node",
        "kind": "state",
        "id": "order.state.processing"
      },
      "depth": 1,
      "via": ["transition_from"]
    }
  ],
  "edges": [
    {
      "kind": "transition_from",
      "direction": "out",
      "from": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
      },
      "to": {
        "object": "node",
        "kind": "state",
        "id": "order.state.processing"
      },
      "depth": 1
    }
  ],
  "truncated": false,
  "truncated_reasons": [],
  "diagnostics": []
}
```

| field | required | content |
|---|---:|---|
| `root` | ✓ | Traversal root ObjectRef. |
| `direction` | ✓ | Direction actually used. |
| `depth` | ✓ | Max traversal depth actually used. |
| `nodes` | ✓ | List of reached objects. |
| `edges` | ✓ | List of References traversed. |
| `truncated` | ✓ | Whether truncated by `max_nodes` / `max_edges`. |
| `truncated_reasons` | ✓ | `max_nodes` / `max_edges`. |
| `diagnostics` | ✓ | Diagnostic list. |

### Node entry

Each entry of `nodes[]` has the shape:

```json
{
  "object": {
    "object": "node",
    "kind": "model",
    "id": "order.model.order"
  },
  "depth": 2,
  "via": ["writes", "store_of"]
}
```

| field | required | content |
|---|---:|---|
| `object` | ✓ | Reached ObjectRef. |
| `depth` | ✓ | Shortest hop count from root. |
| `via` | ✓ | Sequence of `Reference.kind` along the first/shortest BFS path discovered from root. |

When multiple paths reach the same node, `nodes[].via` represents only the shortest, first-discovered path. Use `edges[]` for full path reconstruction.

### Edge entry

`edges[]` uses the same base shape as `Reference`, plus `depth`.

`edges[].depth` is the traversal hop at which that edge was discovered. An edge discovered at hop 1 from root has `depth: 1`.

## Traversal semantics

- Traversal is fixed BFS.
- `depth=0` returns only the root.
- `depth=N` returns nodes and edges reached within `0..N` hops.
- Root is always included in `nodes[]`.
- The same object is not revisited.
- Revisit suppression is normal behavior and is not recorded in `diagnostics[]`.
- Cycle presence, if needed, is inferred from the returned `edges[]`.
- `direction=out` follows edges where the current node is the reference's `from`, advancing to `to`.
- `direction=in` follows edges where the current node is the reference's `to`, advancing to `from`.
- `direction=both` follows both `out` and `in`.

## Selector support

`get_reference_tree`'s root selector is, in principle, the same selector set [`get_references`](get-references.md) supports as a starting point.

Supported as a starting point:

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

Unsupported as a starting point:

- `primitive`
- `view: api_table`
- `view: er_diagram`
- `file: sequence_diagram`
- `file: api_table`
- `file: er_diagram`
- `file: render_index`

This list matches the `get_reference_tree` column of the shared selector support matrix. `primitive` is reachable as a reference target but cannot be used as a traversal root. With `file: node` as root, traversal follows `get_references(file: node)`'s limited range, expanding only references to in-file nodes.

Rooting at an actor with `direction=in` tends to reach many `event_actor` references — specify `kinds` / `max_nodes` / `max_edges` as needed.

## Flow wiring references

`get_reference_tree` v1 does not add new reference kinds.

Flow wiring (`flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over`) continues to be treated as flow-inspect vocabulary within `inspect(task).members.flow.entries`, and is not part of `get_reference_tree`'s traversal target.

If flow wiring is included in impact analysis in the future, it will be handled via a `get_reference_tree` reference-kind extension, or as supplementary material on the `analyze_impact` side.

## Errors

| code | condition |
|---|---|
| `invalid_depth` | `depth` outside `0..4`. |
| `unsupported_direction` | `direction` value outside `out` / `in` / `both`. |
| `unsupported_object` | Selector resolves to an object/kind marked `no` in the selector support matrix. |
| `not_found` | Selector does not resolve to any object. |
| `ambiguous` | Selector resolves to multiple candidates. |
| `kind_mismatch` | Resolved kind does not match `selector.kind`. |

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.mcp.overview` | Parent overview; tool catalog and selection guidance. |
| `spec:bpdsl.mcp.schema` | Selector support matrix, Reference shape. |
| `spec:bpdsl.mcp.errors` | Error code catalog. |
| `spec:bpdsl.mcp.tools.get_references` | Direct-references-only counterpart this tool extends with bounded traversal. |
| `spec:bpdsl.mcp.tools.analyze_impact` | Adds change-kind-aware interpretation on top of raw traversal. |
