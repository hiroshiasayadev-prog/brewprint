# Reference: Cross-edges

- **id**: `spec:bpdsl.dsl.edges.cross_edges`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.edges.overview`

## What this is

Field definitions for cross-layer edges: `reads:` and `writes:` fields on `task` nodes that express access to `store` nodes across layers. These are declared on the task node definition, not in the `flow:` section.

## Cross-edge fields

```yaml
- id: process_payment
  type: task
  reads: [balance_store]
  writes: [transaction_store, balance_store]
  note: "transaction and balance updated in the same transaction"
```

| field | type | description | source |
|---|---|---|---|
| `reads` | list\<store-id\> | Store IDs this task reads from. | V01-ADR-020 |
| `writes` | list\<store-id\> | Store IDs this task writes to. | V01-ADR-020 |

## Design rationale

| rule | reason |
|---|---|
| Declared on task node, not in flow | Store read/write is a property of the task's implementation, independent of which flow calls it. |
| Multiple stores allowed | No 1-task = 1-store constraint. |
| Transaction boundary in `note` | Cross-store transaction scope is described in natural language (same policy as V01-ADR-008). |

## Deprecated cross-edge kinds

| kind | reason for removal |
|---|---|
| `trigger` | Expressed via `transition.action` (V01-ADR-019). |
| `reflect` | Expressed via event + transition chain. |
| `hydrate` | Expressed via event + transition chain. |

(V01-ADR-020)

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.dsl.edges.overview` | Parent overview; edge kind summary. |
| `spec:bpdsl.dsl.nodes.data` | `store` node definitions. |
| `spec:bpdsl.dsl.nodes.processing` | `task` node definition including `reads`/`writes` fields. |
