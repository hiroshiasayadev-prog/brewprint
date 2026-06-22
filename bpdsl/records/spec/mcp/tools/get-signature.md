# Contract: `get_signature`

- **id**: `spec:bpdsl.mcp.tools.get_signature`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.mcp.overview`
- **contract_class**: `interface`

## What this is

`get_signature` returns the external shape of a single target object.

Returns: object identity, kind, source, signature, doc, diagnostics.

Does not return: deep surrounding context, transitive references, full inspect information, render output.

> Source: V01-ADR-018, V01-ADR-021, V01-ADR-026, V01-ADR-028, V01-ADR-031, V01-ADR-035, V01-ADR-062

## Request

```json
{
  "selector": {
    "id": "auth.task.login"
  }
}
```

| field | required | content |
|---|---:|---|
| `selector` | ✓ | Object selector. |

The selector's object/kind support range is governed by the selector support matrix in [`spec:bpdsl.mcp.schema`](../schema.md). If `get_signature` receives a selector marked `no` in the matrix, it returns `unsupported_object` tool error in principle. Retrieving the route list for `view: api_table` is `list_endpoints`'s responsibility, not `get_signature`'s.

## Response

Output envelope:

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
  "doc": "Validate credentials and issue a token",
  "diagnostics": []
}
```

### task signature

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
  "doc": "Validate credentials and issue a token",
  "diagnostics": []
}
```

For a task that is not an endpoint, the `signature.endpoint` field is omitted entirely — `endpoint.enabled: false` / `endpoint: null` are not used.

`signature.endpoint.leaf_path` is the task's own leaf path, not the full path composed by the API Table. The full path is returned via `list_endpoints`'s `endpoints[].path`.

`get_signature` is a lightweight tool returning a task's outward contract, so it does not return `returns.source`. `returns.source` is internal return wiring within the task; check `inspect(task).members.return_source` if needed.

### model signature

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
        "doc": "User ID"
      },
      {
        "name": "role_id",
        "type": "str",
        "fk": "auth.model.role.id",
        "unique": false,
        "doc": "Role ID"
      }
    ]
  },
  "doc": null,
  "diagnostics": []
}
```

### store signature

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
  "doc": "User table",
  "diagnostics": []
}
```

`signature.store_kind` derives from YAML `store.kind` and may be one of `db` / `session` / `collection` / `context`. When a `of` is present, it is returned as the model QualifiedID.

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
  "doc": "Cart session state",
  "diagnostics": []
}
```

The query spec for `store_kind=collection` is included in `doc` as natural language. MCP v1 defines no additional dedicated fields for `store_kind=context`.

### event signature

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
  "doc": "Payment-completed notification from Stripe",
  "diagnostics": []
}
```

### state signature

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
  "doc": "Checkout screen",
  "diagnostics": []
}
```

### transition signature

A transition is queried as a synthetic object, not a node. The selector specifies `object: "transition"` and the TransitionID.

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

| field | required | content |
|---|---:|---|
| `state_file` | ✓ | State FileID where the transition is defined. |
| `from` | ✓ | Source state local ID. |
| `on` | ✓ | Event local ID. |
| `to` | ✓ | Target state local ID. |
| `guard` | optional | Guard string. |
| `action` | optional | Resolved action task QualifiedID. |

### field signature

A model field is queried as a synthetic object. The selector specifies `object: "field"`, the parent model QualifiedID, and the field local ID.

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
  "doc": "Order ID (PK). FK target of order_item.order_id / payment_event.order_id",
  "diagnostics": []
}
```

| field | required | content |
|---|---:|---|
| `name` | ✓ | Field local ID. |
| `type` | ✓ | Field type as written in YAML. |
| `pk` | optional | `true` if primary key. |
| `fk` | optional | FK designation as written in YAML; returns the original bare form even if bare. |
| `unique` | optional | `true` if unique. |

## Errors

| code | condition |
|---|---|
| `unsupported_object` | Selector resolves to an object/kind marked `no` in the selector support matrix. |
| `not_found` | Selector does not resolve to any object. |
| `ambiguous` | Selector resolves to multiple candidates. |
| `kind_mismatch` | Resolved kind does not match `selector.kind`. |
| `invalid_selector` | Selector shape is malformed. |

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.mcp.overview` | Parent overview; tool catalog and selection guidance. |
| `spec:bpdsl.mcp.schema` | Selector support matrix, ObjectRef, TransitionRef, AssetRef shapes. |
| `spec:bpdsl.mcp.errors` | Error code catalog. |
| `spec:bpdsl.mcp.tools.list_endpoints` | Owns full-path computation for endpoint tasks. |
