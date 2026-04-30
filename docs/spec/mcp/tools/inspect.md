---
scope: docs/spec/mcp/tools/inspect.md
status: draft
last_updated: 2026-04-30
summary: >
  inspect toolの仕様を定義する。
  対象objectの実装判断に必要な周辺文脈をkind別に返す。
  task / store / model / state / event / view / fileなどのinspect形を規定する。
depends_on:
  - docs/adr/021-model-field-structure.md
  - docs/adr/026-fk-cardinality-and-nm-relation.md
  - docs/adr/028-api-table-route-composition.md
  - docs/adr/031-actor-global-definition.md
  - docs/adr/032-sequence-diagram-scenario-schema.md
  - docs/adr/035-fsm-guard-branch-and-transition-identification.md
  - docs/adr/036-sequence-diagram-arrow-rules-per-source.md
  - docs/adr/038-sequence-diagram-sub-task-traversal.md
  - docs/adr/039-er-diagram-composed-view.md
---

# `inspect`

## 1. Purpose

`inspect` は、対象objectの実装判断に必要な周辺文脈をkind別にまとめて返す。

`get_signature` が薄い外形確認であるのに対し、`inspect` はLLMが実装・修正・レビュー時に読む濃い文脈取得toolである。

## 2. Input

```json
{
  "selector": {
    "id": "order.task.checkout"
  },
  "detail": "normal"
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `selector` | ✓ | Object selector |
| `detail` | 任意 | `brief` / `normal` / `full`。省略時は `normal` |

`detail` の意味:

| detail | 内容 |
|---|---|
| `brief` | signature + 主要referencesのみ |
| `normal` | 実装判断に必要な標準文脈 |
| `full` | source / members / references / diagnosticsを可能な範囲で最大限返す |

MCP v1では、`detail` による厳密な返却差分は実装任意とする。
ただし未知の値はerrorとする。

## 3. Common output shape

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

| フィールド | 必須 | 内容 |
|---|---:|---|
| `object` | ✓ | ObjectRef |
| `signature` | ✓ | `get_signature` 相当の外形 |
| `doc` | 任意 | note由来の説明 |
| `source` | 任意 | SourceLocation |
| `members` | 任意 | objectが内包する要素 |
| `references` | 任意 | 主要reference |
| `diagnostics` | ✓ | Diagnostic list |

## 4. task inspect

`task` の `inspect` は以下を返す。

- signature
- endpoint情報
- reads / writes
- 同一ファイル内sub task
- flow内での位置
- このtaskをactionとして呼ぶtransition
- このtaskが生成するasset
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
  "doc": "チェックアウトを開始し、注文をpendingで作成する",
  "source": { "file": "order/task/checkout.yaml" },
  "diagnostics": []
}
```

### flow.entries schema status

M11で `members.flow.entries` の最小schemaを確定する。

MCP v1で保証するのは以下。

- `members.flow.file` はflow定義元FileID
- `members.flow.entries[]` はflow内に登場するentryの概略順序を保持する
- 各entryは少なくとも `kind` を持つ
- `step` / `branch` / `fork` / `foreach` のflow構文は、QueryService側で正規化したflow entryとして返す
- wiring情報は `entries[].params[]` / `entries[].over` / `entries[].cases[]` など、flow inspect用schemaに閉じる
- flow inspect用の語彙として `flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over` を使ってよい
- 上記語彙は `Reference.kind` ではなく、`get_references` の返却対象にはしない

`entries[].params[]` は、task paramへのwiringを表す。

```json
{
  "name": "request",
  "source": { "kind": "main_param", "path": "$params.request" }
}
```

`source.kind` は以下を使う。

| source.kind | 意味 |
|---|---|
| `node_return` | 同一flow内の前段nodeの `returns` 全体 |
| `main_param` | `$params.<field>` によるmain task param参照 |
| `foreach_item` | `$item` によるforeach current item参照 |
| `implicit_join` | fork join.params の同名解決 |

`node_return` はreturns内部のfieldを直接参照しない。flow wiringの単位は `docs/spec/edges.md` と同じくtaskのreturns全体とする。

`branch` / `fork` / `foreach` は制御フロー構文であり、flow inspectではentryとして返してよい。ただし、それ自体をMCP selector化することはM11の範囲外とする。

### sub task reads/writes

ADR-038により、Sequence Diagram生成ではmain taskと同一ファイル内のsub taskのreads/writesを集約する。
`inspect(task)` でも、`detail=normal` 以上ではsub taskのreads/writesを辿れるようにする。

推奨形:

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

## 5. store inspect

`store` の `inspect` は以下を返す。

- store signature
- `of` modelのsignature概要
- このstoreを読むtask
- このstoreを書くtask
- kind=dbの場合、ER上のmodel field / FK概要

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
  "doc": "注文テーブル",
  "diagnostics": []
}
```

## 6. model inspect

`model` の `inspect` は以下を返す。

- model signature
- fields
- pk / fk / unique
- このmodelを `store.of` で使うstore
- このmodelをparam / returns / payload / field typeで参照するobject

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

## 7. state inspect

`state` の `inspect` は以下を返す。

- state signature
- incoming transitions
- outgoing transitions
- action task付きtransition
- wireframe有無

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

state inspectでは、incoming / outgoing transitions が中心情報であるため `members` に置く。
一方、`get_references(state)` は `transition_from` / `transition_to` を `references` として返す。

## 8. event inspect

`event` の `inspect` は以下を返す。

- event signature
- source / actor / payload / watches
- このeventをtriggerとして使うtransition
- source種別に基づくSequence Diagram上の補助hint

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

`members.sequence_hints` はADR-036のSequence Diagram render ruleから導ける補助情報である。
これはLLMがeventのsequence上の意味を理解するためのadvisory情報であり、ResolvedProjectの中核semantic relationではない。
Rendererのnormativeな出力規則は `docs/spec/views/sequence-diagram.md` に従う。

## 9. scenario inspect

Sequence Diagram scenarioはview objectとしてinspectできる。

```json
{
  "selector": {
    "object": "view",
    "kind": "sequence_diagram",
    "id": "checkout_flow"
  }
}
```

返す内容:

- scenario ID / title
- state_file
- resolved steps
- 各stepが解決したtransition
- 各stepのaction task
- guard exact match結果

```json
{
  "object": {
    "object": "view",
    "kind": "sequence_diagram",
    "id": "checkout_flow"
  },
  "signature": {
    "state_file": "order/state.yaml",
    "title": "チェックアウトフロー"
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

## 10. transition inspect

Transitionはsynthetic objectとしてinspectできる。

```json
{
  "selector": {
    "object": "transition",
    "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
  }
}
```

返す内容:

- transition signature
- 解決済みfrom state
- 解決済みevent
- 解決済みto state
- 解決済みaction task
- transitionが持つdirect references
- scenario step等からtransitionへのincoming references

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
    "from_state": {
      "object": "node",
      "kind": "state",
      "id": "order.state.processing"
    },
    "event": {
      "object": "node",
      "kind": "event",
      "id": "order.event.payment_webhook_received"
    },
    "to_state": {
      "object": "node",
      "kind": "state",
      "id": "order.state.confirmed"
    },
    "action_task": {
      "object": "node",
      "kind": "task",
      "id": "payment.webhooks.task.process_payment"
    }
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
      "kind": "transition_event",
      "direction": "out",
      "from": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
      },
      "to": { "object": "node", "kind": "event", "id": "order.event.payment_webhook_received" }
    },
    {
      "kind": "transition_to",
      "direction": "out",
      "from": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
      },
      "to": { "object": "node", "kind": "state", "id": "order.state.confirmed" }
    },
    {
      "kind": "transition_action",
      "direction": "out",
      "from": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
      },
      "to": { "object": "node", "kind": "task", "id": "payment.webhooks.task.process_payment" }
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

## 11. field inspect

Model fieldはsynthetic objectとしてinspectできる。

```json
{
  "selector": {
    "object": "field",
    "id": "order.model.order",
    "local_id": "id"
  }
}
```

返す内容:

- field signature
- parent model
- field type
- FK指定
- fieldが持つdirect references
- 他fieldからのincoming FK references

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
        "object": "model_field",
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
        "object": "model_field",
        "kind": "field",
        "id": "order.model.order_item.order_id",
        "qualified_id": "order.model.order_item",
        "name": "order_id",
        "file": "order/model/order_item.yaml"
      },
      "to": {
        "object": "model_field",
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
        "object": "model_field",
        "kind": "field",
        "id": "payment.model.payment_event.order_id",
        "qualified_id": "payment.model.payment_event",
        "name": "order_id",
        "file": "payment/model/payment_event.yaml"
      },
      "to": {
        "object": "model_field",
        "kind": "field",
        "id": "order.model.order.id",
        "qualified_id": "order.model.order",
        "name": "id"
      }
    }
  ],
  "doc": "注文ID（PK）。order_item.order_id / payment_event.order_id のFK参照先",
  "source": { "file": "order/model/order.yaml" },
  "diagnostics": []
}
```

## 12. API Table inspect

API Table viewはview objectとしてinspectできる。

```json
{
  "selector": {
    "object": "view",
    "kind": "api_table",
    "id": "ec_api"
  }
}
```

返す内容:

- API Table ID / `http_root_path`
- 対象modules / `include_submodules`
- moduleごとのendpoint件数
- `list_endpoints` と同じroute合成規則で計算したsections / endpoints

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

`inspect(view: api_table)` は、view定義が何を集約しているかを説明するための文脈取得である。
実装やroute確認でcomputed endpoint一覧だけが必要な場合は、`list_endpoints` を使う。

収集対象endpointが0件のmodule-entryは、API Table render / `list_endpoints` と同様に `sections` には出さない。
ただし `members.modules[]` には `endpoint_count: 0` として残してよい。

## 13. ER Diagram inspect

ER Diagram viewはview objectとしてinspectできる。

```json
{
  "selector": {
    "object": "view",
    "kind": "er_diagram",
    "id": "ec_er"
  }
}
```

返す内容:

- ER Diagram ID
- 対象modules
- included stores
- included models
- view内でrelationとして描画されるFK relations
- view対象外のmodelへ向くFKのsummary

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

view YAMLによる横断ERでは、`modules[]` に明示されたmodule直下の `store.kind: db` のみを対象にする。
サブモジュールは自動では含めない。
view内に含まれないmodelへのFKは `fk_relations` には含めず、`excluded_refs_summary` に入れる。

---

## 14. file inspect

`inspect(file)` は、FileID単位の実装判断用コンテキストを返す。
Raw YAML ASTは返さず、ResolvedProject上に構築済みのsemantic情報をfile単位に要約する。

input例:

```json
{
  "selector": {
    "object": "file",
    "kind": "state_file",
    "id": "order/state.yaml"
  }
}
```

node fileでは以下を返す。

- `members.nodes`: file内のnode一覧
- `members.main_node`: main nodeがある場合のObjectRef
- `members.flow`: flow entry summaryがある場合

state fileでは以下を返す。

- `members.states`
- `members.events`
- `members.transitions`
- `members.wireframes`: stateごとのwireframe有無

view fileでは、view種別に応じて以下を返す。

- `sequence_diagram`: `view`, `state_file`, `steps`
- `api_table`: `view`, `http_root_path`, `modules`
- `er_diagram`: `view`, `modules`

render index fileでは `members.groups` を返す。

---
