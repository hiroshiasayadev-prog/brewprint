---
scope: docs/spec/mcp/tools/get-reference-tree.md
status: draft
last_updated: 2026-04-30
summary: >
  get_reference_tree toolの仕様を定義する。
  対象objectからreferenceをBFSでtraversalし、bounded graphとして返却する。
  入出力スキーマおよびtraversal semanticsを規定する。
depends_on:
  - docs/adr/049-mcp-query-reference-vocabulary.md
  - docs/adr/054-mcp-query-coverage-for-design-conversation.md
  - docs/adr/055-mcp-reference-tree-traversal.md
---

# `get_reference_tree`

## 1. Purpose

`get_reference_tree` は、対象objectをrootとして、direct referencesをdepth制限つきでBFS traversalする。

tool名は `tree` だが、返却形式は純粋な木ではなく、`nodes[]` / `edges[]` からなる bounded reference graph とする。

返すもの:

- root object
- traversal direction / depth
- 到達したobjects
- traversalで辿ったreferences
- truncation情報
- diagnostics

返さないもの:

- 変更種別ごとのimpact severity
- recommended action
- renderer output mapping
- flow wiring references
- Raw YAML AST

変更種別を含む影響判断は、将来の `analyze_impact`（[../versioning.md](../versioning.md)）で扱う。

## 2. Input

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

| フィールド | 必須 | 内容 |
|---|---:|---|
| `selector` | ✓ | traversal root object |
| `direction` | ✓ | `out` / `in` / `both` |
| `depth` | ✓ | `0..4` |
| `kinds` | 任意 | traversal / return対象のreference kind filter |
| `max_nodes` | 任意 | node返却上限。省略時 `200` |
| `max_edges` | 任意 | edge返却上限。省略時 `500` |

`depth < 0` または `depth > 4` は `invalid_depth` error とする。
`direction` は探索範囲を暗黙化しないため必須とする。

`kinds` を指定した場合、指定されたreference kindのみを traversal 経路として辿り、`edges[]` に含める。
指定外kindでしか到達できないobjectには到達しない。

## 3. Output

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
    },
    {
      "object": {
        "object": "node",
        "kind": "event",
        "id": "order.event.payment_webhook_received"
      },
      "depth": 1,
      "via": ["transition_event"]
    },
    {
      "object": {
        "object": "node",
        "kind": "state",
        "id": "order.state.confirmed"
      },
      "depth": 1,
      "via": ["transition_to"]
    },
    {
      "object": {
        "object": "node",
        "kind": "task",
        "id": "payment.webhooks.task.process_payment"
      },
      "depth": 1,
      "via": ["transition_action"]
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
    },
    {
      "kind": "transition_event",
      "direction": "out",
      "from": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
      },
      "to": {
        "object": "node",
        "kind": "event",
        "id": "order.event.payment_webhook_received"
      },
      "depth": 1
    },
    {
      "kind": "transition_to",
      "direction": "out",
      "from": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
      },
      "to": {
        "object": "node",
        "kind": "state",
        "id": "order.state.confirmed"
      },
      "depth": 1
    },
    {
      "kind": "transition_action",
      "direction": "out",
      "from": {
        "object": "transition",
        "kind": "transition",
        "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
      },
      "to": {
        "object": "node",
        "kind": "task",
        "id": "payment.webhooks.task.process_payment"
      },
      "depth": 1
    }
  ],
  "truncated": false,
  "truncated_reasons": [],
  "diagnostics": []
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `root` | ✓ | traversal root ObjectRef |
| `direction` | ✓ | 実際に使ったdirection |
| `depth` | ✓ | 実際に使ったmax traversal depth |
| `nodes` | ✓ | 到達object一覧 |
| `edges` | ✓ | traversalで辿ったReference一覧 |
| `truncated` | ✓ | `max_nodes` / `max_edges` により打ち切ったか |
| `truncated_reasons` | ✓ | `max_nodes` / `max_edges` |
| `diagnostics` | ✓ | Diagnostic list |

## 4. Node entry

`nodes[]` の各entryは以下の形を持つ。

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

| フィールド | 必須 | 内容 |
|---|---:|---|
| `object` | ✓ | 到達したObjectRef |
| `depth` | ✓ | rootからの最短到達hop数 |
| `via` | ✓ | rootから最初に到達したBFS経路のReference.kind列 |

同一nodeへ複数経路が存在する場合、`nodes[].via` は最短かつ最初に探索された経路のみを表す。
完全な経路復元には `edges[]` を使う。

## 5. Edge entry

`edges[]` は `Reference` と同じ基本形に `depth` を加えたものとする。

`edges[].depth` は、そのedgeを発見した traversal hop を表す。
rootから1 hop目で発見されたedgeは `depth: 1` とする。

## 6. Traversal semantics

- traversal は BFS 固定とする
- `depth=0` はrootのみを返す
- `depth=N` は `0..N` hop までの到達nodeとedgeを返す
- rootは常に `nodes[]` に含める
- 同一objectへの再訪は行わない
- 同一objectへの再訪停止は正常動作であり、`diagnostics[]` には記録しない
- 循環の有無を知りたい場合は、返却された `edges[]` から推論する
- `direction=out` は現在nodeがreferenceの `from` であるedgeを辿り、`to` へ進む
- `direction=in` は現在nodeがreferenceの `to` であるedgeを辿り、`from` へ進む
- `direction=both` は `out` / `in` の両方を辿る

## 7. Selector support

`get_reference_tree` のroot selectorは、基本的に [`get_references`](./get-references.md) の対応selectorを起点にする。

起点として supported:

- `node: task`
- `node: model`
- `node: store`
- `node: state`
- `node: event`
- `node: actor`
- `view: sequence_diagram`
- `transition`
- `field`
- `file: node` limited
- `file: state_file`
- `asset`
- private sub node

起点として unsupported:

- `primitive`
- `view: api_table`
- `view: er_diagram`
- `file: sequence_diagram`
- `file: api_table`
- `file: er_diagram`
- `file: render_index`

`primitive` は reference target として到達可能だが、traversal rootにはできない。
`file: node` をrootにした場合は、`get_references(file: node)` の limited 対応範囲に従い、file内nodeへのreferenceのみ展開する。

actorをrootにして `direction=in` を指定すると、多数の `event_actor` reference に到達しやすい。
必要に応じて `kinds` / `max_nodes` / `max_edges` を指定する。

## 8. Flow wiring references

`get_reference_tree` v1は、新しいreference kindを追加しない。

flow wiring（`flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over`）は、引き続き `inspect(task).members.flow.entries` 内のflow inspect用語彙として扱い、`get_reference_tree` のtraversal対象には含めない。

将来、flow wiringを影響分析に含める場合は、`get_reference_tree` のreference kind拡張、または `analyze_impact` 側の補完材料として扱う。

---
