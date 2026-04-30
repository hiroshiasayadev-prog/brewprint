---
scope: docs/spec/mcp/tools/list-objects.md
status: draft
last_updated: 2026-04-30
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
| `kind` | 任意 | `task` / `model` / `api_table` / `transition` / `field` 等 |
| `module` | 任意 | module path。例: `order`, `payment.webhooks` |
| `file` | 任意 | FileID |

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
