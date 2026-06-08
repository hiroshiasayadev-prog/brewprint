# V01-ADR-055: MCP reference traversal は get_reference_tree 別toolとして定義する

- **status**: accepted
- **date**: 2026-04-30
- **depends on**: V01-ADR-047, V01-ADR-048, V01-ADR-049, V01-ADR-054

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

V01-ADR-049では、MCP / QueryService の参照語彙を `references` に統一し、`get_references` / `GetReferences` は v1 で direct references のみを返すと決定した。

V01-ADR-054では、MCP / QueryService を設計対話 coverage を基準に拡張する方針を決めた。
その中で、設計変更相談では direct references だけでは不足し、transitive impact traversal が必要になるため、将来拡張として以下の選択肢を保留していた。

- `get_reference_tree` のような別toolを追加する
- `get_references` に `depth` inputを追加する

M12では、まず `get_source` を実装し、次に direct references より深い影響範囲確認を扱う必要が出た。
このADRは、V01-ADR-054で保留していた reference traversal の設計判断を確定する。

## 決定

### 1. `get_reference_tree` を別toolとして追加する

MCP reference traversal は、`get_references(depth=...)` ではなく、別tool `get_reference_tree` として定義する。

`get_references` は引き続き direct references のみを返す。
`get_references` の外部契約に `depth` inputは追加しない。

```text
get_references       direct references API
get_reference_tree   bounded reference graph traversal API
```

これにより、V01-ADR-049の direct-only 方針を維持したまま、設計対話向けの traversal を追加する。

### 2. tool名は tree だが、出力は bounded reference graph とする

`get_reference_tree` という名前を採用するが、返却形式は純粋な木ではなく、`nodes[]` / `edges[]` からなる bounded reference graph とする。

理由:

- field FK は循環しうる
- asset / task / store / model は複数経路で同じobjectに到達しうる
- transition / scenario / state / event の関係も多重経路になりうる
- `children[]` 型の木にすると、同じobjectが重複し、LLMが影響範囲を読み誤りやすい

`tree` は「起点から深さ制限つきで辿る」操作感を表す名前であり、返却データ構造は graph とする。

### 3. traversal は BFS 固定とする

v1の `get_reference_tree` は BFS で traversal する。

depth の意味:

| depth | 意味 |
|---:|---|
| `0` | root のみ |
| `1` | direct references |
| `2..4` | N hop references |

`depth` は required とする。
暗黙defaultは持たない。

v1の最大depthは `4` とする。
`depth < 0` または `depth > 4` は tool error とし、error code は `invalid_depth` とする。

`depth=4` は、典型的な `task -> store -> model -> field -> primitive` 程度の確認を想定した上限である。
より深い FK chain や変更種別ごとの意味づけは、将来の上限見直し、または `analyze_impact` 側で扱う。

### 4. `direction` は required とする

`get_reference_tree` の `direction` は required とする。

`get_references` は省略時 `out` を維持するが、`get_reference_tree` は探索範囲が広がるtoolであるため、方向を暗黙にしない。

有効値:

- `out`
- `in`
- `both`

traversal rule:

| direction | traversal |
|---|---|
| `out` | 現在nodeがreferenceの `from` であるedgeを辿り、`to` へ進む |
| `in` | 現在nodeがreferenceの `to` であるedgeを辿り、`from` へ進む |
| `both` | `out` / `in` の両方を辿る |

### 5. 同一objectは再訪しない

BFS中、同一objectへ再訪しない。

同一objectへの再訪停止は正常動作であり、errorにも diagnostic にもしない。
`cycle_count` 等の summary も v1では返さない。

循環の有無を知りたい場合は、返却された `edges[]` から推論する。

### 6. `kinds` は traversal と返却の両方に適用する

`kinds` は optional な reference kind filter とする。

指定された場合、`kinds` に含まれるreference kindのみを:

1. traversal 経路として辿る
2. `edges[]` に含める

指定外のreference kindでしか到達できないobjectには到達しない。

将来、traversal対象と返却対象を分ける必要が出た場合は、`traverse_kinds` / `return_kinds` などを別途検討する。
ただし v1では単一の `kinds` のみを採用する。

### 7. root は常に `nodes[]` に含める

v1では `include_start` inputを持たない。

root object は常に `nodes[]` に含める。
root node の `depth` は `0`、`via` は空配列とする。

```json
{
  "object": { "object": "node", "kind": "task", "id": "order.task.checkout" },
  "depth": 0,
  "via": []
}
```

### 8. nodes は BFS最初到達経路の概要を持つ

`nodes[]` の各要素は、対象objectと到達depthを返す。

また、LLMが nodes だけを読んでも「なぜここに来たか」を把握できるように、`via` を返す。

`via` は、rootから当該nodeへ最初に到達したBFS経路の `Reference.kind` 列である。
同一nodeへ複数経路が存在する場合、`nodes[].via` は最短かつ最初に探索された経路のみを表す。
完全な経路復元には `edges[]` を使う。

```json
{
  "object": { "object": "node", "kind": "model", "id": "order.model.order" },
  "depth": 2,
  "via": ["writes", "store_of"]
}
```

### 9. truncation は明示する

`max_nodes` / `max_edges` を input として持つ。

v1のdefault:

| input | default |
|---|---:|
| `max_nodes` | `200` |
| `max_edges` | `500` |

これらの制限に達して traversal を打ち切った場合、responseに以下を返す。

```json
{
  "truncated": true,
  "truncated_reasons": ["max_nodes"]
}
```

`truncated_reasons` の値:

- `max_nodes`
- `max_edges`

指定depthに到達して止まることは正常動作であり、`truncated` とは扱わない。
`depth > 4` は丸めず `invalid_depth` error とする。

### 10. 起点selectorの対応範囲

`get_reference_tree` は、基本的に `get_references` の対応selectorを起点にする。

v1で起点として supported:

- `node: task`
- `node: model`
- `node: store`
- `node: state`
- `node: event`
- `node: actor`
- `view: sequence_diagram`
- `transition`
- `field`
- `file: node`
- `file: state_file`
- `asset`
- private sub node

v1で起点として unsupported:

- `primitive`
- `view: api_table`
- `view: er_diagram`
- `file: sequence_diagram`
- `file: api_table`
- `file: er_diagram`
- `file: render_index`

`primitive` は reference target only として到達可能だが、起点にはしない。
primitiveに到達した場合は、通常leafとして扱われる。

API Table / ER Diagram view は集約viewであり、v1では `inspect(view)` / `list_endpoints` で文脈を取得する。
reference traversal の起点化は将来検討とする。

### 11. flow wiring references は v1 traversal対象外とする

MCP v1の `get_references` は flow wiring references を返さない。
したがって `get_reference_tree` も v1では flow wiring references を traversal対象に含めない。

`flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over` は、引き続き `inspect(task).members.flow.entries` 内の flow inspect 用語彙として扱う。

将来、flow wiringを影響分析に含める場合は、`get_reference_tree` の reference kind 拡張、または `analyze_impact` 側の補完材料として扱う。

### 12. `analyze_impact` は別toolとして扱う

`get_reference_tree` は reference graph traversal を返す低レベルtoolである。
変更種別ごとの意味づけ、重要度、推奨対応は返さない。

以下は将来の `analyze_impact` で扱う。

- `rename`
- `remove`
- `change_type`
- `change_contract`
- impact severity
- recommended action
- render output mapping
- flow wiring を含む補完的影響分析

## 仕様案

### Input

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

| field | required | 内容 |
|---|---:|---|
| `selector` | yes | traversal root object |
| `direction` | yes | `out` / `in` / `both` |
| `depth` | yes | `0..4` |
| `kinds` | no | traversal / return対象のreference kind filter |
| `max_nodes` | no | node返却上限。default `200` |
| `max_edges` | no | edge返却上限。default `500` |

### Output

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

`edges[].depth` は、そのedgeを発見した traversal hop を表す。
rootから1 hop目で発見されたedgeは `depth: 1` とする。

## 理由

### 別toolにする理由

`get_references` は単体objectのdirect relationを確認するためのAPIである。
これに `depth` を追加すると、direct APIとtraversal APIの責務が混ざる。

設計対話では、direct references確認と影響範囲探索は利用場面が異なる。
別toolに分けることで、LLMは以下を明確に使い分けられる。

| 状況 | tool |
|---|---|
| 対象objectの直接参照だけ確認したい | `get_references` |
| 影響範囲をN hop辿りたい | `get_reference_tree` |
| 変更種別を含めた影響判断が欲しい | `analyze_impact` |

### graph返却にする理由

reference relation は一般に木ではない。
同じobjectへ複数経路で到達したり、循環したりする。

木として返すと、同一objectの重複、循環の表現、path explosion が問題になる。
そのため、bounded graphとして `nodes[]` / `edges[]` を返す。

### direction / depth を required にする理由

`get_reference_tree` は探索範囲が広がるtoolである。
`direction` や `depth` を暗黙defaultにすると、LLMが意図しない広い探索を行う可能性がある。

そのため、v1では両方を required とし、呼び出し側に探索意図を明示させる。

### cycleをdiagnosticにしない理由

同一objectの再訪停止は BFS traversal の正常動作である。
正常動作を diagnostic にすると、FK閉路や多重参照が多いprojectで大量のノイズが生じる。

MCP diagnostic は、tool実行は成立したが注意すべき情報に使う。
cycle検出はエラーでも注意でもなく、traversal ruleそのものなので、diagnosticには入れない。

## 影響

### V01-ADR-049への影響

V01-ADR-049は supersede しない。

V01-ADR-049の `get_references` / `GetReferences` は direct references APIとして維持する。
本ADRは、transitive traversalが必要になった場合の別tool追加方針を定める拡張ADRである。

### V01-ADR-054への影響

V01-ADR-054 §4で保留されていた選択肢のうち、`get_reference_tree` 別tool案を採用する。

V01-ADR-054の設計対話 coverage 方針と整合する。

### docs/spec/mcp.mdへの影響

実装時に `docs/spec/mcp.md` を更新する。

必要な更新:

- tool overview に `get_reference_tree` を追加する
- selector support matrix に `get_reference_tree` 列を追加する、または専用support表を追加する
- `get_references` の depth節に、direct-only維持と `get_reference_tree` 採用を注記する
- `get_reference_tree` の input / output / traversal semantics / error model を追加する
- error code に `invalid_depth` を追加する
- future extensions から `get_reference_tree` を外す
- LLM tool selection guidance に `get_reference_tree` を追加する

### QueryService / MCP実装への影響

実装時は、既存 `GetReferences` を内部的に再利用して BFS を構築する。

v1で新しい `Reference.kind` は追加しない。
flow wiring references も追加しない。

想定内部処理:

```text
root selector resolve
  ↓
queue root depth=0
  ↓
GetReferences(current, direction, kinds)
  ↓
edge dedupe / node dedupe
  ↓
next objects enqueue until depth / max_nodes / max_edges
```

### M12への影響

M12の `get_reference_tree または depth指定つきreference traversalを設計する` task は、本ADRにより設計判断が確定する。
実装タスクは、別途 `docs/spec/mcp.md` 更新後に分割する。

## Evidence

- commit: 8a6bc19
- impl commit: tbd
- 参考: V01-ADR-049, V01-ADR-054, docs/spec/mcp.md
