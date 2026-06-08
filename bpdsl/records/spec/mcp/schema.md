---
scope: docs/spec/mcp/schema.md
status: draft
last_updated: 2026-06-07
summary: >
  MCP toolで共通利用するschemaを定義する。
  selector、ObjectRef、Reference、Diagnosticなどの共通表現を規定する。
depends_on:
  - docs/adr/027-module-nesting-and-name-resolution.md
  - docs/adr/031-actor-global-definition.md
  - docs/adr/035-fsm-guard-branch-and-transition-identification.md
  - docs/adr/043-project-root-layout-and-render-output.md
  - docs/adr/048-resolved-project-index-strategy.md
  - docs/adr/054-mcp-query-coverage-for-design-conversation.md
  - docs/adr/056-mcp-analyze-impact-tool-design.md
  - docs/adr/062-task-return-source.md
---

# MCP schema

## 1. Common schema

### 1.1 Object selector

各toolは、対象objectを指定するために `selector` を受け取る。

```json
{
  "selector": {
    "id": "auth.task.login"
  }
}
```

#### selector fields

| フィールド | 必須 | 型 | 内容 |
|---|---:|---|---|
| `id` | 任意 | string | 対象objectのID。通常はQualifiedID。actorはglobal actor ID。scenario等のview objectはview固有ID。transition / asset / private sub nodeは後述のsynthetic IDを使う。`file` + `local_id` で指定できるobjectでは省略できる |
| `object` | 任意 | enum | `node` / `view` / `transition` / `asset` / `field` / `file` / `primitive`。省略時は `node` として解決を試みる |
| `kind` | 任意 | string | 期待するkind。値集合は `object` に依存する。指定時、解決結果のkindが一致しなければ `kind_mismatch` tool error とする |
| `file` | 任意 | FileID | private sub node等、file-local objectを指定する場合に使う |
| `local_id` | 任意 | string | `file` 内のlocal object ID。sub task等のprivate object参照に使う |

`object` を省略した selector は `object: node` として解決を試みる。
`kind` だけを指定しても object class は推論しない。
node 以外を問い合わせる場合は、原則として `object` を明示する。

selector の形が壊れている場合は `invalid_selector` tool error とする。
selector が解決できない場合は `not_found`、複数候補に解決される場合は `ambiguous`、解決結果と `kind` が一致しない場合は `kind_mismatch` を返す。
解決は成立したが tool がその object / kind を扱わない場合は、§1.2 の selector support matrix に従う。

通常の外部参照可能nodeは `id` のみで指定する。
private sub nodeを直接問い合わせる必要がある場合は、synthetic IDまたは `file` + `local_id` の形式を使う。

```json
{
  "selector": {
    "object": "node",
    "id": "order/task/checkout.yaml#build_order"
  }
}
```

```json
{
  "selector": {
    "object": "node",
    "file": "order/task/checkout.yaml",
    "local_id": "build_order"
  }
}
```

assetを直接問い合わせる場合は、producerとasset nameからなるsynthetic ID、または `id` + `local_id` の形式を使う。

```json
{
  "selector": {
    "object": "asset",
    "id": "order.task.build_order#draft_order"
  }
}
```

private sub nodeの直接問い合わせは `get_signature` / `get_references` / `inspect` で対応する。
`inspect(main task)` の `members.sub_tasks` も同じObjectRef表現を使う。

#### object-dependent kind vocabulary

`kind` の値集合は `object` ごとに異なる。
この表は selector validation と ObjectRef response の共通語彙である。

| object | allowed kind values | notes |
|---|---|---|
| `node` | `task` / `model` / `store` / `state` / `event` / `actor` | brewprint node kind。private sub node も `object: node`, `kind: task` として表す |
| `view` | `sequence_diagram` / `api_table` / `er_diagram` | view object。API endpoints の computed route 一覧は `list_endpoints` の責務 |
| `transition` | `transition` | state file 内の synthetic transition object |
| `asset` | `asset` | task / join returns から暗黙生成される asset object |
| `field` | `field` / `model_field` | model field object。MCP v1 response では既存互換のため `field` を返してよい。新規specでは `model_field` を説明用語として使ってよい |
| `file` | `node` / `state_file` / `sequence_diagram` / `api_table` / `er_diagram` / `render_index` | `yaml/` 配下の file kind。node kind ではなく file kind を表す |
| `primitive` | `primitive` | primitive type reference target。直接query対象ではない |

この vocabulary は DATA DSL の dependent enum 機能ではない。
MCP schema / tool contract 上の runtime selector contract と response contract として扱う。

### 1.2 Selector support matrix

MCP v1のselector対応範囲は以下とする。

| object / kind | [`get_signature`](tools/get-signature.md) | [`get_references`](tools/get-references.md) | [`get_reference_tree`](tools/get-reference-tree.md) | [`analyze_impact`](tools/analyze-impact.md) | [`inspect`](tools/inspect.md) | status |
|---|---:|---:|---:|---:|---:|---|
| `node: task` | yes | yes | yes | yes | yes | supported |
| `node: model` | yes | yes | yes | yes | yes | supported |
| `node: store` | yes | yes | yes | yes | yes | supported |
| `node: state` | yes | yes | yes | yes | yes | supported |
| `node: event` | yes | yes | yes | yes | yes | supported |
| `node: actor` | yes | yes | yes | yes | limited | supported / limited inspect |
| `view: sequence_diagram` | yes | yes | yes | no | yes | supported |
| `transition` | yes | yes | yes | yes | yes | supported |
| `field` / `model_field` | yes | yes | yes | yes | yes | supported |
| `file: node` | no | limited | limited | no | yes | supported / limited references |
| `file: state_file` | no | yes | yes | no | yes | supported |
| `file: sequence_diagram` | no | no | no | no | yes | supported |
| `file: api_table` | no | no | no | no | yes | supported |
| `file: er_diagram` | no | no | no | no | yes | supported |
| `file: render_index` | no | no | no | no | yes | supported |
| `asset` | yes | yes | yes | no | yes | supported |
| `view: api_table` | no | no | no | no | yes | supported; `list_endpoints` はcomputed route一覧専用 |
| `view: er_diagram` | no | no | no | no | yes | supported |
| private sub node | yes | yes | yes | no | yes | supported |
| `primitive` | no | no | no | no | no | reference target only |

cell値の意味:

| value | 意味 |
|---|---|
| `yes` | tool が当該 selector を通常 query 対象として扱う |
| `no` | tool が当該 selector を通常 query 対象として扱わない |
| `limited` | tool は当該 selector を扱うが、返却範囲または情報量に制約がある |

statusの意味:

| status | 意味 |
|---|---|
| `supported` | MCP v1でquery対象として扱う |
| `supported / limited inspect` | signature / references は扱うが、専用inspectの情報量は限定的 |
| `supported / limited references` | file単位など、referenceの意味が限定される |
| `partial` | 一部toolのみ対応する |
| `future` | 設計対話coverage上は候補だが、現時点では未実装 |
| `v1 optional` | spec上は許容するが、実装必須ではない |
| `reference target only` | reference targetとして返すが、直接query対象ではない |

`no` の扱いは tool ごとに異なる。
`get_signature` / `get_references` / `get_reference_tree` / `inspect` で `no` の selector を受け取った場合は、原則として `unsupported_object` tool error とする。
`analyze_impact` の `no` は、tool error ではなく空 `impacts` と `unsupported_selector` diagnostic を返す対象を表す。

`get_reference_tree` における `file: node` の `limited` は、`get_references(file: node)` の対応範囲に従い、file内nodeへのreferenceのみ展開することを意味する。
`primitive` はreference targetとして到達可能だが、traversal rootとしては扱わない。

> 由来: V01-ADR-054 §決定

### 1.3 QualifiedID

モジュールスコープを持つnodeのQualifiedIDは、V01-ADR-027に従う。

```text
<module-path>.<node-kind>.<id>
```

例:

```text
auth.task.login
order.store.order_db
catalog.model.item
payment.stripe.task.receive_webhook
```

`<module-path>` は1段以上のdot区切りmodule path。
`<node-kind>` はbrewprintのnode kind sentinel。

#### actorの例外

actorはV01-ADR-031によりproject globalであり、モジュールに属さない。

actor参照は常にID直参照とする。

```text
stripe
scheduler
end_user
```

MCP responseでは、actorにも `qualified_id` フィールドを返すが、その値はglobal actor IDと同じでよい。

```json
{
  "object": "node",
  "kind": "actor",
  "id": "stripe",
  "qualified_id": "stripe"
}
```

### 1.4 FileID

FileIDは、V01-ADR-043で定義されるbrewprint projectの `yaml/` ディレクトリからの相対パスをslash正規化した文字列とする。

```text
auth/task/login.yaml
order/state.yaml
views/scenarios/checkout_flow.yaml
```

Windows上の `\\` はMCP responseでは `/` に正規化する。

### 1.5 Synthetic ID

QualifiedIDを持たないfile-local objectには、MCP response上の安定参照としてsynthetic IDを使う。

#### private sub node ID

private sub nodeのsynthetic IDは以下とする。

```text
<file-id>#<local-id>
```

例:

```text
order/task/checkout.yaml#build_order
order/task/checkout.yaml#reserve_inventory
```

sub nodeは外部モジュールから参照不可だが、MCP response内では `inspect(main task)` のmembersやreference targetとして識別する必要があるため、synthetic IDを返す。

#### asset ID

implicit assetのsynthetic IDは以下とする。

```text
<producer-qualified-id>#<asset-name>
```

例:

```text
order.task.build_order#draft_order
auth.task.login#auth_token
```

assetはtask / join の `returns` から暗黙生成される。直接queryする場合はこのsynthetic ID、またはproducerを `id`、asset nameを `local_id` として指定する。

#### transition ID

transition IDは、V01-ADR-035 / V01-ADR-048の `(stateFileID, fromStateID, eventID, guard?)` に対応する文字列として以下を使う。

```text
<state-file-id>#<from-state>:<event>
<state-file-id>#<from-state>:<event>[<guard>]
```

例:

```text
auth/state.yaml#idle:login_submitted
order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']
```

guard文字列はYAML decode後の文字列をそのまま使う。
trim、空白正規化、Unicode正規化、式AST比較は行わない。
これはV01-ADR-035 / V01-ADR-048のguard exact match方針と一致する。

### 1.6 SourceLocation

```json
{
  "file": "auth/task/login.yaml",
  "line": 12,
  "column": 5,
  "end_line": 42,
  "end_column": 1
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `file` | ✓ | FileID |
| `line` | 任意 | 1-origin line number |
| `column` | 任意 | 1-origin column number |
| `end_line` | 任意 | 範囲終端line |
| `end_column` | 任意 | 範囲終端column |

line / column が取得できない実装では `file` のみ返してよい。

### 1.7 ObjectRef

MCP response内でobjectを指す共通形式。

```json
{
  "object": "node",
  "kind": "task",
  "id": "auth.task.login",
  "qualified_id": "auth.task.login",
  "label": "login",
  "source": {
    "file": "auth/task/login.yaml",
    "line": 3
  }
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `object` | ✓ | `node` / `view` / `transition` / `asset` / `field` / `file` / `primitive` |
| `kind` | ✓ | object種別。nodeなら `task` / `model` 等 |
| `id` | ✓ | object ID。nodeならQualifiedID、actor global ID、またはfile-local synthetic ID |
| `qualified_id` | 任意 | 解決済みQualifiedID。QualifiedIDを持つobjectのみ。`id` と同じ場合も返してよい |
| `file` | 任意 | file-local objectの場合の所属FileID |
| `local_id` | 任意 | file-local objectの場合のlocal ID |
| `label` | 任意 | 人間向け短縮表示名 |
| `source` | 任意 | SourceLocation |
| `parent` | 任意 | field等の親ObjectRef |

private sub nodeは以下のように表す。

```json
{
  "object": "node",
  "kind": "task",
  "id": "order/task/checkout.yaml#build_order",
  "file": "order/task/checkout.yaml",
  "local_id": "build_order",
  "label": "build_order"
}
```

primitiveは以下のように表す。

```json
{
  "object": "primitive",
  "kind": "primitive",
  "id": "str",
  "label": "str"
}
```

model fieldは以下のように表す。

```json
{
  "object": "field",
  "kind": "model_field",
  "id": "auth.model.user.email",
  "label": "email",
  "parent": {
    "object": "node",
    "kind": "model",
    "id": "auth.model.user"
  }
}
```

### 1.8 TransitionRef

TransitionRefは、transitionを表すObjectRef拡張である。

```json
{
  "object": "transition",
  "kind": "transition",
  "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']",
  "state_file": "order/state.yaml",
  "from": "processing",
  "on": "payment_webhook_received",
  "to": "confirmed",
  "guard": "payload.status == 'succeeded'",
  "action": "order.task.confirm_order",
  "source": {
    "file": "order/state.yaml",
    "line": 42
  }
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `object` | ✓ | `transition` 固定 |
| `kind` | ✓ | `transition` 固定 |
| `id` | ✓ | TransitionID |
| `state_file` | ✓ | transition定義元のstate FileID |
| `from` | ✓ | 遷移元state ID |
| `on` | ✓ | event ID |
| `to` | ✓ | 遷移先state ID |
| `guard` | 任意 | guard文字列。exact match対象 |
| `action` | 任意 | action task QualifiedID |
| `source` | 任意 | SourceLocation |

state / event / taskへのObjectRefが必要な場合は、`references` に `transition_from` / `transition_event` / `transition_to` / `transition_action` として返す。
TransitionRef本体の `from` / `on` / `to` は、state diagram / scenario stepを読みやすくするための短縮情報である。

### 1.9 AssetRef

`asset` はYAML上に独立ファイルを持たず、taskの `returns` から暗黙生成される。
implicit assetはQualifiedIDを持たないため、producerとasset nameからなるsynthetic IDで直接queryできる。
`AssetRef` はproducer contextつきで返す。

```json
{
  "object": "asset",
  "id": "auth.task.login#auth_token",
  "name": "auth_token",
  "producer": "auth.task.login",
  "model": "auth.model.token",
  "scope_file": "auth/task/login.yaml"
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `object` | ✓ | `asset` 固定 |
| `id` | ✓ | `<producer>#<name>` のasset synthetic ID |
| `name` | ✓ | `task.returns.name` |
| `producer` | ✓ | assetを生成するtaskのQualifiedIDまたはfile-local synthetic ID |
| `model` | ✓ | assetのmodel QualifiedIDまたはprimitive |
| `scope_file` | 任意 | assetが生じるFileID |

`task.returns.name` はDAG上のasset labelとして扱うが、sub taskや別fileで同名returnsがあり得るため、MCP responseではproducer contextと一緒に返す。
直接query用のIDは `<producer>#<name>` のsynthetic IDを使う。

### 1.10 Diagnostic

```json
{
  "severity": "warning",
  "code": "uncovered_module",
  "message": "module catalog is not covered by render_index.yaml; implicit group will be used",
  "source": {
    "file": "render_index.yaml"
  }
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `severity` | ✓ | `error` / `warning` / `info` / `hint` |
| `code` | ✓ | machine-readable code |
| `message` | ✓ | human-readable message |
| `source` | 任意 | SourceLocation |
| `related` | 任意 | 関連SourceLocationまたはObjectRefの配列 |

---

## 2. Reference schema

### 2.1 Reference

`Reference` は、brewprint object間の直接参照を表す。

```json
{
  "kind": "reads",
  "direction": "out",
  "from": {
    "object": "node",
    "kind": "task",
    "id": "auth.task.login"
  },
  "to": {
    "object": "node",
    "kind": "store",
    "id": "auth.store.user_db"
  },
  "source": {
    "file": "auth/task/login.yaml",
    "line": 10
  }
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `kind` | ✓ | reference種別 |
| `direction` | ✓ | query対象から見た方向。`out` / `in` |
| `from` | ✓ | 参照元ObjectRef |
| `to` | ✓ | 参照先ObjectRef |
| `source` | 任意 | このreferenceが定義されているSourceLocation |
| `doc` | 任意 | referenceに関する自然言語補足。例: branch case label / transition note等 |

### 2.2 Reference kind

MCP v1で返すreference kindは以下とする。

| kind | from | to | 意味 |
|---|---|---|---|
| `param_model` | task / branch / join | model / primitive | paramがmodelまたはprimitiveを型参照する |
| `return_model` | task / join | model / primitive | returnsがmodelまたはprimitiveを型参照する |
| `produces_asset` | task / join | asset | returnsによりimplicit assetを生成する |
| `consumes_asset` | asset | task / join | flow wiringでimplicit assetがconsumer nodeへ渡される |
| `reads` | task | store | taskがstoreを読む |
| `writes` | task | store | taskがstoreへ書く |
| `store_of` | store | model | storeがmodelを保持する |
| `field_type` | model field | model / primitive | model fieldが型を参照する |
| `field_fk` | model field | model field | model fieldがFK参照する |
| `transition_event` | transition | event | transitionがeventをtriggerとして参照する |
| `transition_from` | transition | state | transitionのfrom state |
| `transition_to` | transition | state | transitionのto state |
| `transition_action` | transition | task | transitionがaction taskを呼ぶ |
| `event_payload` | event | model | event payloadがmodelを参照する |
| `event_actor` | event | actor | external eventがactorを参照する |
| `event_watches` | event | store | er eventがstoreを監視する |
| `scenario_state_file` | scenario | state file | sequence scenarioがstate_fileを参照する |
| `scenario_step_transition` | scenario step | transition | sequence scenario stepがtransitionを参照する |

flow wiring（`flow_step` / `flow_param` 相当の情報）はMCP v1の `get_references` では返さない。
flow wiringはDAG file内部の局所構造であり、MCP v1では `inspect(task).members.flow.entries` に閉じる。

`returns.source` による task return wiring も、MCP v1の `get_references` では返さない。`returns.source` は `inspect(task).members.return_source` に raw / resolved 情報として返し、global reverse index（`referencesBySource` / `referencesByTarget`）の必須対象にはしない。

M11ではこの方針を維持し、`flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over` / `task_return_source` は `Reference.kind` ではなく、flow / task inspect用の語彙として扱う。
これらは将来の `get_reference_tree` / `analyze_impact` の traversal 材料になりうるが、direct references v1 の返却対象には含めない。

### 2.3 Direction

`direction` は、query対象から見た向きを表す。

| direction | 意味 |
|---|---|
| `out` | query対象が参照している相手 |
| `in` | query対象を参照している相手 |

例: `inspect(auth.store.user_db)` で `auth.task.login` が `user_db` を読む場合、referenceは以下になる。

```json
{
  "kind": "reads",
  "direction": "in",
  "from": { "id": "auth.task.login", "kind": "task" },
  "to": { "id": "auth.store.user_db", "kind": "store" }
}
```

---
