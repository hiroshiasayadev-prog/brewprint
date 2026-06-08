---
scope: docs/spec/mcp/tools/get-references.md
status: draft
last_updated: 2026-06-07
summary: >
  get_references toolの仕様を定義する。
  対象objectの直接referenceを返す。
  directionとreference kind filter、およびdepthの扱いを規定する。
depends_on:
  - docs/adr/049-mcp-query-reference-vocabulary.md
---

# `get_references`

## 1. Purpose

`get_references` は、対象objectの直接referenceを返す。

MCP v1ではdirect referencesのみを返す。

## 2. Input

```json
{
  "selector": {
    "id": "auth.task.login"
  },
  "direction": "both",
  "kinds": ["reads", "writes", "transition_action"]
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `selector` | ✓ | Object selector |
| `direction` | 任意 | `out` / `in` / `both`。省略時は `out` |
| `kinds` | 任意 | reference kind filter。省略時は全kind |

`direction` が省略された場合、toolは `out` を使用し、response の `direction` には実際に使用した `out` を返す。
`direction` が `out` / `in` / `both` 以外の場合は `unsupported_direction` tool error とする。
この default は MCP tool contract の実行時挙動であり、DATA DSL の default 構文としては扱わない。

`selector` の object / kind 対応範囲は `docs/spec/mcp/schema.md` の selector support matrix を正本とする。
`get_references` で matrix が `no` の selector を受け取った場合は、原則として `unsupported_object` tool error とする。
`limited` の selector は同matrixと本toolの tool-specific section に従って、返却対象referenceを限定してよい。

## 3. Output

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

## 4. depth

`depth` は常に `1` を返す。

MCP v1では、`get_references` inputに `depth` を持たない。
transitive reference traversal は、V01-ADR-055に従い、別tool `get_reference_tree` ([get-reference-tree.md](./get-reference-tree.md)) で扱う。

## 5. Selector support

`get_references` の selector support は `docs/spec/mcp/schema.md` の selector support matrix を正本とする。

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
- `file: node` limited
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

---
