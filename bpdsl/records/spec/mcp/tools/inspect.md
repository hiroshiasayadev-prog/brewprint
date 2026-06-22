# Contract: `inspect`

- **id**: `spec:bpdsl.mcp.tools.inspect`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.mcp.overview`
- **contract_class**: `interface`

## What this is

`inspect` returns, per object kind, the surrounding context an implementer needs in order to make implementation decisions.

Where `get_signature` is a thin external-shape check, `inspect` is the dense-context tool an LLM reads when implementing, fixing, or reviewing.

> Source: V01-ADR-021, V01-ADR-026, V01-ADR-028, V01-ADR-031, V01-ADR-032, V01-ADR-035, V01-ADR-036, V01-ADR-038, V01-ADR-039, V01-ADR-062, V01-ADR-063

## Request

```json
{
  "selector": {
    "id": "order.task.checkout"
  },
  "detail": "normal"
}
```

| field | required | content |
|---|---:|---|
| `selector` | ✓ | Object selector. |
| `detail` | optional | `brief` / `normal` / `full`. Defaults to `normal` when omitted. |

`detail` meaning:

| detail | content |
|---|---|
| `brief` | signature + main references only. |
| `normal` | Standard context needed for implementation decisions. |
| `full` | Returns source / members / references / diagnostics as fully as feasible. |

MCP v1 leaves the exact return-shape difference per `detail` to implementation discretion. When `detail` is omitted, the tool uses `normal`. A `detail` value outside `brief` / `normal` / `full` is an `unsupported_detail` tool error. This default is MCP tool-contract runtime behavior, not a DATA DSL default construct.

The selector's object/kind support range is governed by the selector support matrix in [`spec:bpdsl.mcp.schema`](../schema.md). If `inspect` receives a selector marked `no` in the matrix, it returns `unsupported_object` tool error in principle. A `limited` selector may have the information volume of returned `members` / `references` / `diagnostics` restricted per that matrix and this tool's kind-specific sections below.

## Response

Common output shape:

```json
{
  "object": {},
  "signature": {},
  "doc": "...",
  "source": {},
  "members": {},
  "references": [],
  "diagnostics": []
}
```

| field | required | content |
|---|---:|---|
| `object` | ✓ | ObjectRef. |
| `signature` | ✓ | Equivalent to `get_signature`'s external shape. |
| `doc` | optional | Description derived from `note`. |
| `source` | optional | SourceLocation. |
| `members` | optional | Elements the object contains. |
| `references` | optional | Main references. |
| `diagnostics` | ✓ | Diagnostic list. |

### task inspect

`inspect` on a `task` returns:

- signature
- endpoint info
- reads / writes
- sub tasks within the same file
- position within the flow
- raw / resolved info for the task return source (`returns.source`)
- transitions that call this task as an action
- assets this task produces
- source
- doc

```json
{
  "object": {
    "object": "node",
    "kind": "task",
    "id": "order.task.checkout"
  },
  "signature": {
    "main": true,
    "params": [
      { "name": "request", "model": "order.model.checkout_request" }
    ],
    "returns": {
      "name": "pending_order",
      "model": "order.model.order"
    },
    "endpoint": {
      "method": "POST",
      "leaf_path": "checkout"
    }
  },
  "members": {
    "return_source": {
      "raw": "build_order",
      "source": {
        "kind": "node_return",
        "node": "build_order",
        "type_ref": "order.model.order"
      }
    },
    "assets": [
      {
        "object": "asset",
        "name": "pending_order",
        "producer": "order.task.checkout",
        "model": "order.model.order",
        "scope_file": "order/task/checkout.yaml"
      }
    ],
    "sub_tasks": [
      {
        "object": "node",
        "kind": "task",
        "id": "order/task/checkout.yaml#build_order",
        "file": "order/task/checkout.yaml",
        "local_id": "build_order",
        "label": "build_order",
        "signature": {
          "reads": ["cart.store.cart_session", "auth.store.user_db"],
          "writes": ["order.store.order_db"]
        },
        "source": { "file": "order/task/checkout.yaml" }
      },
      {
        "object": "node",
        "kind": "task",
        "id": "order/task/checkout.yaml#reserve_inventory",
        "file": "order/task/checkout.yaml",
        "local_id": "reserve_inventory",
        "label": "reserve_inventory",
        "signature": {
          "reads": ["inventory.store.inventory_db"],
          "writes": ["inventory.store.inventory_db"]
        },
        "source": { "file": "order/task/checkout.yaml" }
      }
    ],
    "flow": {
      "file": "order/task/checkout.yaml",
      "entries": [
        {
          "kind": "step",
          "step": "build_order",
          "params": [
            {
              "name": "request",
              "source": { "kind": "main_param", "path": "$params.request" }
            }
          ]
        },
        {
          "kind": "step",
          "step": "reserve_inventory",
          "params": [
            {
              "name": "order",
              "source": { "kind": "node_return", "node": "build_order" }
            }
          ]
        }
      ],
      "schema_status": "confirmed"
    }
  },
  "references": [
    {
      "kind": "produces_asset",
      "direction": "out",
      "from": { "object": "node", "kind": "task", "id": "order.task.checkout" },
      "to": { "object": "asset", "name": "pending_order", "producer": "order.task.checkout" }
    },
    {
      "kind": "transition_action",
      "direction": "in",
      "from": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#checkout_screen:checkout_submitted",
        "state_file": "order/state.yaml",
        "from": "checkout_screen",
        "on": "checkout_submitted",
        "to": "processing",
        "action": "order.task.checkout"
      },
      "to": { "object": "node", "kind": "task", "id": "order.task.checkout" }
    }
  ],
  "doc": "Start checkout and create the order as pending",
  "source": { "file": "order/task/checkout.yaml" },
  "diagnostics": []
}
```

#### flow.entries schema status

M11 fixes the minimum schema of `members.flow.entries`.

MCP v1 guarantees the following:

- `members.flow.file` is the FileID where the flow is defined.
- `members.flow.entries[]` preserves the approximate appearance order of entries in the flow.
- Each entry has at least `kind`.
- `step` / `branch` / `fork` / `foreach` flow constructs are returned as QueryService-normalized flow entries.
- Wiring info stays scoped within flow-inspect schema: `entries[].params[]` / `entries[].over` / `entries[].cases[]` etc.
- `flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over` may be used as flow-inspect vocabulary.
- The above vocabulary is not `Reference.kind` and is not part of `get_references`'s return scope.

`entries[].params[]` represents wiring into a task param.

```json
{
  "name": "request",
  "source": { "kind": "main_param", "path": "$params.request" }
}
```

`source.kind` uses one of:

| source.kind | meaning |
|---|---|
| `node_return` | The entire `returns` of a preceding node in the same flow. |
| `collected_asset` | A collected asset source declared via `foreach.returns`. |
| `initialized_source` | An initialized source declared via `initializes[].name`. |
| `main_param` | A main task param reference via `$params.<field>`. |
| `foreach_item` | The current foreach item via `$item`. |
| `implicit_join` | Same-name resolution of a fork join's params. |

`node_return` does not reference a field inside `returns` directly. The flow-wiring unit is the entire task `returns`, matching [`spec:bpdsl.dsl.edges.data_flow`](../../dsl/edges/data-flow.md).

For a task with `returns.source` specified, `inspect(task)` returns `members.return_source`. `get_signature` returns only the outward contract and does not include `returns.source`.

```json
{
  "members": {
    "return_source": {
      "raw": "validated_items",
      "source": {
        "kind": "collected_asset",
        "name": "validated_items",
        "type_ref": "list<cart.model.cart_item>"
      }
    }
  }
}
```

`members.return_source.raw` preserves the `returns.source` string as written in YAML. `members.return_source.source` represents the source resolved on `ResolvedProject`; `kind` is one of `node_return` / `collected_asset` / `initialized_source` / `main_param`. `$item` is invalid as a task return source — it is subject to the `invalid_return_source` diagnostic and never becomes a resolved source.

When `returns.source` is unspecified, or the task has no `returns`, `members.return_source` is omitted. When the source is unresolved or invalid, diagnostics take priority and the implementation may return only `raw`, omitting the resolved `source`.

`branch` / `fork` / `foreach` are control-flow constructs and may be returned as flow-inspect entries. Making them individually selectable as MCP selectors is out of M11 scope.

#### sub task reads/writes

Per V01-ADR-038, Sequence Diagram generation aggregates the reads/writes of sub tasks within the same file as the main task. `inspect(task)` should also let sub-task reads/writes be traced at `detail=normal` or above.

Recommended shape:

```json
{
  "members": {
    "sub_tasks": [
      {
        "id": "order/task/checkout.yaml#build_order",
        "file": "order/task/checkout.yaml",
        "local_id": "build_order",
        "signature": {
          "reads": ["cart.store.cart_session", "auth.store.user_db"],
          "writes": ["order.store.order_db"]
        }
      }
    ]
  }
}
```

### store inspect

`inspect` on a `store` returns:

- store signature
- summary signature of the `of` model
- tasks that read this store
- tasks that write this store
- for `kind=db`, an ER-level model field / FK summary

```json
{
  "object": {
    "object": "node",
    "kind": "store",
    "id": "order.store.order_db"
  },
  "signature": {
    "store_kind": "db",
    "of": "order.model.order"
  },
  "members": {
    "model": {
      "object": "node",
      "kind": "model",
      "id": "order.model.order",
      "fields": [
        { "name": "id", "type": "str", "pk": true },
        { "name": "user_id", "type": "str", "fk": "auth.model.credential.username" }
      ]
    }
  },
  "references": [
    {
      "kind": "reads",
      "direction": "in",
      "from": { "object": "node", "kind": "task", "id": "order.task.load_order" },
      "to": { "object": "node", "kind": "store", "id": "order.store.order_db" }
    },
    {
      "kind": "writes",
      "direction": "in",
      "from": { "object": "node", "kind": "task", "id": "order.task.checkout" },
      "to": { "object": "node", "kind": "store", "id": "order.store.order_db" }
    }
  ],
  "doc": "Order table",
  "diagnostics": []
}
```

### model inspect

`inspect` on a `model` returns:

- model signature
- fields
- pk / fk / unique
- stores that use this model via `store.of`
- objects that reference this model via param / returns / payload / field type

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
      { "name": "id", "type": "str", "pk": true },
      { "name": "email", "type": "str" }
    ]
  },
  "references": [
    {
      "kind": "store_of",
      "direction": "in",
      "from": { "object": "node", "kind": "store", "id": "auth.store.user_db" },
      "to": { "object": "node", "kind": "model", "id": "auth.model.user" }
    }
  ],
  "diagnostics": []
}
```

### state inspect

`inspect` on a `state` returns:

- state signature
- incoming transitions
- outgoing transitions
- transitions with an action task
- whether a wireframe exists

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
    "wireframe": { "present": true }
  },
  "members": {
    "incoming_transitions": [
      {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#cart:view_checkout",
        "state_file": "order/state.yaml",
        "from": "cart",
        "on": "view_checkout",
        "to": "checkout_screen"
      }
    ],
    "outgoing_transitions": [
      {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#checkout_screen:checkout_submitted",
        "state_file": "order/state.yaml",
        "from": "checkout_screen",
        "on": "checkout_submitted",
        "to": "processing",
        "action": "order.task.checkout"
      }
    ]
  },
  "diagnostics": []
}
```

State inspect places incoming/outgoing transitions in `members` since they are the central information. By contrast, `get_references(state)` returns `transition_from` / `transition_to` as `references`.

### event inspect

`inspect` on an `event` returns:

- event signature
- source / actor / payload / watches
- transitions that use this event as a trigger
- Sequence Diagram auxiliary hints based on source kind

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
    "payload": { "model": "payment.model.payment_event" }
  },
  "references": [
    {
      "kind": "transition_event",
      "direction": "in",
      "from": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']",
        "state_file": "order/state.yaml",
        "from": "processing",
        "on": "payment_webhook_received",
        "to": "confirmed",
        "guard": "payload.status == 'succeeded'"
      },
      "to": { "object": "node", "kind": "event", "id": "order.event.payment_webhook_received" }
    }
  ],
  "members": {
    "sequence_hints": {
      "advisory": true,
      "participant": "Actor",
      "actor": "stripe",
      "message_label_source": "METHOD path"
    }
  },
  "diagnostics": []
}
```

`members.sequence_hints` is supplementary info derivable from the V01-ADR-036 Sequence Diagram render rule. It is advisory information helping the LLM understand an event's sequence meaning, not a core ResolvedProject semantic relation. The renderer's normative output rule follows [`spec:bpdsl.views.sequence_diagram`](../../views/sequence-diagram.md).

### scenario inspect

A Sequence Diagram scenario can be inspected as a view object.

```json
{
  "selector": {
    "object": "view",
    "kind": "sequence_diagram",
    "id": "checkout_flow"
  }
}
```

Returns:

- scenario ID / title
- state_file
- resolved steps
- the transition each step resolves to
- each step's action task
- guard exact-match result

```json
{
  "object": {
    "object": "view",
    "kind": "sequence_diagram",
    "id": "checkout_flow"
  },
  "signature": {
    "state_file": "order/state.yaml",
    "title": "Checkout flow"
  },
  "members": {
    "steps": [
      {
        "index": 1,
        "from_state": "cart",
        "via": "view_checkout",
        "transition": {
          "object": "transition",
          "kind": "transition",
          "id": "order/state.yaml#cart:view_checkout",
          "state_file": "order/state.yaml",
          "from": "cart",
          "on": "view_checkout",
          "to": "checkout_screen"
        },
        "action": null
      },
      {
        "index": 2,
        "from_state": "checkout_screen",
        "via": "checkout_submitted",
        "transition": {
          "object": "transition",
          "kind": "transition",
          "id": "order/state.yaml#checkout_screen:checkout_submitted",
          "state_file": "order/state.yaml",
          "from": "checkout_screen",
          "on": "checkout_submitted",
          "to": "processing",
          "action": "order.task.checkout"
        },
        "action": "order.task.checkout"
      }
    ]
  },
  "references": [
    {
      "kind": "scenario_state_file",
      "direction": "out",
      "from": { "object": "view", "kind": "sequence_diagram", "id": "checkout_flow" },
      "to": { "object": "file", "kind": "state_file", "id": "order/state.yaml" }
    }
  ],
  "diagnostics": []
}
```

### transition inspect

A transition can be inspected as a synthetic object.

```json
{
  "selector": {
    "object": "transition",
    "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
  }
}
```

Returns:

- transition signature
- resolved from state
- resolved event
- resolved to state
- resolved action task
- direct references this transition holds
- incoming references from scenario steps etc. to this transition

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
  "members": {
    "from_state": { "object": "node", "kind": "state", "id": "order.state.processing" },
    "event": { "object": "node", "kind": "event", "id": "order.event.payment_webhook_received" },
    "to_state": { "object": "node", "kind": "state", "id": "order.state.confirmed" },
    "action_task": { "object": "node", "kind": "task", "id": "payment.webhooks.task.process_payment" }
  },
  "references": [
    {
      "kind": "transition_from",
      "direction": "out",
      "from": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
      },
      "to": { "object": "node", "kind": "state", "id": "order.state.processing" }
    },
    {
      "kind": "scenario_step_transition",
      "direction": "in",
      "from": {
        "object": "scenario_step",
        "kind": "sequence_step",
        "id": "scenario_step:payment_webhook_flow:1"
      },
      "to": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
      }
    }
  ],
  "diagnostics": []
}
```

### field inspect

A model field can be inspected as a synthetic object.

```json
{
  "selector": {
    "object": "field",
    "id": "order.model.order",
    "local_id": "id"
  }
}
```

Returns:

- field signature
- parent model
- field type
- FK designation
- direct references this field holds
- incoming FK references from other fields

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
  "members": {
    "model": {
      "object": "node",
      "kind": "model",
      "id": "order.model.order",
      "qualified_id": "order.model.order",
      "label": "order",
      "file": "order/model/order.yaml"
    },
    "type": "str"
  },
  "references": [
    {
      "kind": "field_type",
      "direction": "out",
      "from": {
        "object": "field",
        "kind": "field",
        "id": "order.model.order.id",
        "qualified_id": "order.model.order",
        "name": "id",
        "file": "order/model/order.yaml"
      },
      "to": { "object": "primitive", "kind": "primitive", "id": "str", "name": "str" }
    },
    {
      "kind": "field_fk",
      "direction": "in",
      "from": {
        "object": "field",
        "kind": "field",
        "id": "order.model.order_item.order_id",
        "qualified_id": "order.model.order_item",
        "name": "order_id",
        "file": "order/model/order_item.yaml"
      },
      "to": {
        "object": "field",
        "kind": "field",
        "id": "order.model.order.id",
        "qualified_id": "order.model.order",
        "name": "id"
      }
    },
    {
      "kind": "field_fk",
      "direction": "in",
      "from": {
        "object": "field",
        "kind": "field",
        "id": "payment.model.payment_event.order_id",
        "qualified_id": "payment.model.payment_event",
        "name": "order_id",
        "file": "payment/model/payment_event.yaml"
      },
      "to": {
        "object": "field",
        "kind": "field",
        "id": "order.model.order.id",
        "qualified_id": "order.model.order",
        "name": "id"
      }
    }
  ],
  "doc": "Order ID (PK). FK target of order_item.order_id / payment_event.order_id",
  "source": { "file": "order/model/order.yaml" },
  "diagnostics": []
}
```

### API Table inspect

An API Table view can be inspected as a view object.

```json
{
  "selector": {
    "object": "view",
    "kind": "api_table",
    "id": "ec_api"
  }
}
```

Returns:

- API Table ID / `http_root_path`
- target modules / `include_submodules`
- endpoint count per module
- `sections` / `endpoints` computed with the same route-composition rule as `list_endpoints`

```json
{
  "object": {
    "object": "view",
    "kind": "api_table",
    "id": "ec_api"
  },
  "signature": {
    "id": "ec_api",
    "http_root_path": "/api",
    "modules": [
      { "module": "auth", "include_submodules": false }
    ]
  },
  "members": {
    "modules": [
      { "module": "auth", "include_submodules": false, "endpoint_count": 1 }
    ],
    "sections": [
      {
        "module": "auth",
        "include_submodules": false,
        "endpoints": [
          {
            "method": "POST",
            "path": "/api/login",
            "leaf_path": "login",
            "task": "auth.task.login"
          }
        ]
      }
    ],
    "collected_endpoints": [
      {
        "module": "auth",
        "task": "auth.task.login",
        "method": "POST",
        "path": "/api/login",
        "leaf_path": "login"
      }
    ]
  },
  "diagnostics": []
}
```

`inspect(view: api_table)` is context retrieval explaining what the view definition aggregates. When only the computed endpoint list is needed for implementation or route checking, use `list_endpoints` instead.

A module-entry with zero collected endpoints is not shown in `sections`, matching API Table render / `list_endpoints` behavior. It may still appear in `members.modules[]` with `endpoint_count: 0`.

### ER Diagram inspect

An ER Diagram view can be inspected as a view object.

```json
{
  "selector": {
    "object": "view",
    "kind": "er_diagram",
    "id": "ec_er"
  }
}
```

Returns:

- ER Diagram ID
- target modules
- included stores
- included models
- FK relations drawn as relations within the view
- a summary of FKs pointing to models outside the view

```json
{
  "object": {
    "object": "view",
    "kind": "er_diagram",
    "id": "ec_er"
  },
  "signature": {
    "id": "ec_er",
    "modules": [
      { "module": "auth" },
      { "module": "order" }
    ]
  },
  "members": {
    "modules": [
      { "module": "auth", "store_count": 1, "model_count": 1 }
    ],
    "included_stores": [
      { "object": "node", "kind": "store", "id": "order.store.order_db" }
    ],
    "included_models": [
      { "object": "node", "kind": "model", "id": "order.model.order" }
    ],
    "fk_relations": [
      {
        "from_model": "order.model.order_item",
        "from_field": "order_id",
        "to_model": "order.model.order",
        "to_field": "id",
        "fk": "order.id",
        "cardinality": "many_to_one"
      }
    ],
    "excluded_refs_summary": {
      "count": 0
    }
  },
  "diagnostics": []
}
```

A cross-module ER view defined by a view YAML targets only `store.kind: db` directly under modules explicitly listed in `modules[]`. Submodules are not automatically included. FKs pointing to a model not included in the view are not added to `fk_relations` — they go into `excluded_refs_summary`.

### file inspect

`inspect(file)` returns FileID-granularity context for implementation decisions. It does not return the raw YAML AST — it summarizes already-built semantic information per file.

Input example:

```json
{
  "selector": {
    "object": "file",
    "kind": "state_file",
    "id": "order/state.yaml"
  }
}
```

A node file returns:

- `members.nodes`: list of nodes in the file.
- `members.main_node`: ObjectRef of the main node, if present.
- `members.flow`: flow entry summary, if present.

A state file returns:

- `members.states`
- `members.events`
- `members.transitions`
- `members.wireframes`: whether a wireframe exists, per state.

A view file returns, depending on view kind:

- `sequence_diagram`: `view`, `state_file`, `steps`.
- `api_table`: `view`, `http_root_path`, `modules`.
- `er_diagram`: `view`, `modules`.

A render index file returns `members.groups`.

## Errors

| code | condition |
|---|---|
| `unsupported_object` | Selector resolves to an object/kind marked `no` in the selector support matrix. |
| `unsupported_detail` | `detail` value outside `brief` / `normal` / `full`. |
| `not_found` | Selector does not resolve to any object. |
| `ambiguous` | Selector resolves to multiple candidates. |
| `kind_mismatch` | Resolved kind does not match `selector.kind`. |

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.mcp.overview` | Parent overview; tool catalog and selection guidance. |
| `spec:bpdsl.mcp.schema` | Selector support matrix, ObjectRef, AssetRef, TransitionRef shapes. |
| `spec:bpdsl.mcp.errors` | Error code catalog. |
| `spec:bpdsl.mcp.tools.get_signature` | Thin external-shape counterpart this tool extends with dense context. |
| `spec:bpdsl.mcp.tools.list_endpoints` | Computed endpoint list, used instead of `inspect(view: api_table)` for route-only needs. |
| `spec:bpdsl.dsl.edges.data_flow` | Flow wiring unit (`returns` as a whole) referenced by `members.flow.entries`. |
| `spec:bpdsl.views.sequence_diagram` | Normative Sequence Diagram render rule that `members.sequence_hints` is advisory to. |
