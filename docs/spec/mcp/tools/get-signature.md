---
scope: docs/spec/mcp/tools/get-signature.md
status: draft
last_updated: 2026-04-30
summary: >
  get_signature toolの仕様を定義する。
  対象object単体の外形、signature、doc、diagnosticsを返す。
  kind別signatureの返却形を規定する。
depends_on:
  - docs/adr/018-event-node.md
  - docs/adr/021-model-field-structure.md
  - docs/adr/026-fk-cardinality-and-nm-relation.md
  - docs/adr/028-api-table-route-composition.md
  - docs/adr/031-actor-global-definition.md
  - docs/adr/035-fsm-guard-branch-and-transition-identification.md
---

# `get_signature`

## 1. Purpose

`get_signature` は、対象object単体の外形を返す。

返すもの:

- object identity
- kind
- source
- signature
- doc
- diagnostics

返さないもの:

- 深い周辺文脈
- transitive references
- full inspect情報
- render出力

## 2. Input

```json
{
  "selector": {
    "id": "auth.task.login"
  }
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `selector` | ✓ | Object selector |

## 3. Output envelope

```json
{
  "object": {
    "object": "node",
    "kind": "task",
    "id": "auth.task.login",
    "qualified_id": "auth.task.login",
    "label": "login",
    "source": {
      "file": "auth/task/login.yaml",
      "line": 3
    }
  },
  "signature": {},
  "doc": "認証情報を検証しトークンを発行する",
  "diagnostics": []
}
```

## 4. task signature

```json
{
  "object": {
    "object": "node",
    "kind": "task",
    "id": "auth.task.login",
    "qualified_id": "auth.task.login",
    "label": "login",
    "source": { "file": "auth/task/login.yaml" }
  },
  "signature": {
    "main": true,
    "params": [
      {
        "name": "credentials",
        "model": "auth.model.credential"
      }
    ],
    "returns": {
      "name": "auth_token",
      "model": "auth.model.token",
      "asset": {
        "object": "asset",
        "name": "auth_token",
        "producer": "auth.task.login",
        "model": "auth.model.token",
        "scope_file": "auth/task/login.yaml"
      }
    },
    "reads": ["auth.store.user_db"],
    "writes": ["auth.store.session_store"],
    "endpoint": {
      "method": "POST",
      "leaf_path": "login"
    }
  },
  "doc": "認証情報を検証しトークンを発行する",
  "diagnostics": []
}
```

`endpoint` でないtaskでは、`signature.endpoint` フィールド自体を省略する。
`endpoint.enabled: false` / `endpoint: null` は使わない。

`signature.endpoint.leaf_path` はtask側のleaf pathであり、API Tableで合成されたfull pathではない。
full pathは `list_endpoints` の `endpoints[].path` で返す。

## 5. model signature

```json
{
  "object": {
    "object": "node",
    "kind": "model",
    "id": "auth.model.user"
  },
  "signature": {
    "model_kind": "struct",
    "fields": [
      {
        "name": "id",
        "type": "str",
        "pk": true,
        "doc": "ユーザーID"
      },
      {
        "name": "role_id",
        "type": "str",
        "fk": "auth.model.role.id",
        "unique": false,
        "doc": "ロールID"
      }
    ]
  },
  "doc": null,
  "diagnostics": []
}
```

## 6. store signature

```json
{
  "object": {
    "object": "node",
    "kind": "store",
    "id": "auth.store.user_db"
  },
  "signature": {
    "store_kind": "db",
    "of": "auth.model.user"
  },
  "doc": "ユーザーテーブル",
  "diagnostics": []
}
```

`signature.store_kind` はYAMLの `store.kind` に由来し、`db` / `session` / `collection` / `context` の4種を返しうる。
いずれも `of` を持つ場合はmodel QualifiedIDを返す。

```json
{
  "object": {
    "object": "node",
    "kind": "store",
    "id": "cart.store.cart_session"
  },
  "signature": {
    "store_kind": "session",
    "of": "cart.model.cart"
  },
  "doc": "カートのセッション状態",
  "diagnostics": []
}
```

`store_kind=collection` のquery仕様は `doc` に自然言語で含める。
`store_kind=context` の追加固有フィールドはMCP v1では定義しない。

## 7. event signature

```json
{
  "object": {
    "object": "node",
    "kind": "event",
    "id": "order.event.payment_webhook_received"
  },
  "signature": {
    "source": "external",
    "actor": "stripe",
    "payload": {
      "model": "payment.model.payment_event"
    }
  },
  "doc": "Stripeからの決済完了通知",
  "diagnostics": []
}
```

## 8. state signature

```json
{
  "object": {
    "object": "node",
    "kind": "state",
    "id": "order.state.checkout_screen"
  },
  "signature": {
    "initial": false,
    "final": false,
    "wireframe": {
      "present": true
    }
  },
  "doc": "チェックアウト画面",
  "diagnostics": []
}
```

## 9. transition signature

Transitionはnodeではなくsynthetic objectとして問い合わせる。
selectorには `object: "transition"` とTransitionIDを指定する。

```json
{
  "selector": {
    "object": "transition",
    "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
  }
}
```

```json
{
  "object": {
    "object": "transition",
    "kind": "transition",
    "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']",
    "file": "order/state.yaml",
    "local_id": "processing:payment_webhook_received"
  },
  "signature": {
    "state_file": "order/state.yaml",
    "from": "processing",
    "on": "payment_webhook_received",
    "to": "confirmed",
    "guard": "payload.status == 'succeeded'",
    "action": "payment.webhooks.task.process_payment"
  },
  "diagnostics": []
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `state_file` | ✓ | transition定義元のstate FileID |
| `from` | ✓ | 遷移元state local ID |
| `on` | ✓ | event local ID |
| `to` | ✓ | 遷移先state local ID |
| `guard` | 任意 | guard文字列 |
| `action` | 任意 | 解決済みaction task QualifiedID |

## 10. field signature

Model fieldはsynthetic objectとして問い合わせる。
selectorには `object: "field"`、親model QualifiedID、field local IDを指定する。

```json
{
  "selector": {
    "object": "field",
    "id": "order.model.order",
    "local_id": "id"
  }
}
```

```json
{
  "object": {
    "object": "field",
    "kind": "field",
    "id": "order.model.order.id",
    "qualified_id": "order.model.order",
    "label": "id",
    "file": "order/model/order.yaml",
    "local_id": "id"
  },
  "signature": {
    "name": "id",
    "type": "str",
    "pk": true
  },
  "doc": "注文ID（PK）。order_item.order_id / payment_event.order_id のFK参照先",
  "diagnostics": []
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `name` | ✓ | field local ID |
| `type` | ✓ | YAML上のfield type |
| `pk` | 任意 | primary keyならtrue |
| `fk` | 任意 | YAML上のFK指定。bare FKの場合も元の記述を返す |
| `unique` | 任意 | uniqueならtrue |

---
