---
scope: docs/spec/mcp/tools/get-references.md
status: draft
last_updated: 2026-04-30
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
transitive reference traversal は、ADR-055に従い、別tool `get_reference_tree` ([get-reference-tree.md](./get-reference-tree.md)) で扱う。

---
