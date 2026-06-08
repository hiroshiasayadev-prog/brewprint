---
scope: docs/spec/mcp/tools/list-objects.md
status: draft
last_updated: 2026-06-07
summary: >
  list_objects toolの仕様を定義する。
  project内のsemantic object一覧を返す探索用tool。
  詳細情報は他toolで取得する前提とする。
depends_on:
  - docs/adr/054-mcp-query-coverage-for-design-conversation.md
---

# `list_objects`

## 1. Purpose

`list_objects` は、project内のsemantic object一覧を返す。

返す対象:

- `node`
- `view`
- `transition`
- `field`

`asset` / `file` / `primitive` は `list_objects` の一覧対象外とする。
これらは reference target、selector target、または file inspect target として扱う。

`list_objects` は探索用toolであり、各objectの詳細なsignature / references / inspect情報は返さない。詳細が必要な場合は、返された `object` / `kind` / `id` をselectorとして `get_signature` / `get_references` / `inspect` を呼ぶ。

## 2. Input
```json
{
  "object": "node",
  "kind": "task",
  "module": "order",
  "file": "order/task/checkout.yaml"
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `object` | 任意 | `node` / `view` / `transition` / `field` |
| `kind` | 任意 | object-dependent kind filter。値集合は `docs/spec/mcp/schema.md` の object-dependent kind vocabulary に従う |
| `module` | 任意 | module path。例: `order`, `payment.webhooks` |
| `file` | 任意 | FileID |

`object` を省略した場合は、`node` / `view` / `transition` / `field` の全 listable object を対象にする。
`kind` は指定された `object` に依存する filter である。
`object` 省略時に `kind` を指定した場合は、listable object 全体の中でその `kind` に一致するものだけを返す。
例えば `kind: task` は `node: task`、`kind: api_table` は `view: api_table` に一致する。
未知の `object` または `kind` は `invalid_args` tool error とする。

## 3. Output
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

---
