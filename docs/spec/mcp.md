---
scope: docs/spec/mcp.md
status: draft
last_updated: 2026-04-30
summary: >
  brewprintのMCP query tool外部仕様。
  Python inspectに近い語彙で、ResolvedProject上のsemantic objectに対する
  signature / source / references / inspect / endpoint query のinput/outputと、
  設計対話に必要なquery coverageを定義する。
depends_on:
  - docs/adr/017-diagram-layers-and-scope.md
  - docs/adr/018-event-node.md
  - docs/adr/021-model-field-structure.md
  - docs/adr/026-fk-cardinality-and-nm-relation.md
  - docs/adr/027-module-nesting-and-name-resolution.md
  - docs/adr/028-api-table-route-composition.md
  - docs/adr/030-yaml-file-type-declaration.md
  - docs/adr/031-actor-global-definition.md
  - docs/adr/032-sequence-diagram-scenario-schema.md
  - docs/adr/035-fsm-guard-branch-and-transition-identification.md
  - docs/adr/036-sequence-diagram-arrow-rules-per-source.md
  - docs/adr/038-sequence-diagram-sub-task-traversal.md
  - docs/adr/043-project-root-layout-and-render-output.md
  - docs/adr/047-go-semantic-model-query-layer-boundary.md
  - docs/adr/048-resolved-project-index-strategy.md
  - docs/adr/049-mcp-query-reference-vocabulary.md
  - docs/adr/054-mcp-query-coverage-for-design-conversation.md
  - docs/adr/055-mcp-reference-tree-traversal.md
---

# MCP仕様

## 1. Scope

このspecは、brewprintがMCP経由でLLMへ提供するquery toolの**外部I/O契約**を定義する。

対象:

- MCP tool名
- tool input schema
- tool output schema
- 共通ID表現
- 共通レスポンス語彙
- reference表現
- diagnostic / error表現
- LLMが各toolをどう使うべきかの意図

対象外:

- Go package名 / struct名 / interface名
- Raw YAML decode用struct
- `ResolvedProject` 内部のmap/index具体実装
- RendererのMermaid / HTML / Markdown出力仕様
- MCP transport実装の詳細
- transitive dependency graphの事前構築方式

本specは、ADR-047 / ADR-048で定義されたGo実装境界、およびADR-049で定義されたreference語彙統一を前提に、`QueryService` が外部へ返す情報の形を固定する。

---

## 2. Design principles

### 2.1 MCPはRaw YAML ASTを公開しない

brewprint MCPは、Raw YAML ASTをLLMへ公開するためのAPIではない。

MCP toolは常に、semantic buildを通過した `ResolvedProject` 上の情報を返す。

```text
YAML files
  ↓ load / classify / decode
Raw YAML structs
  ↓ validate / name resolution / derived model build / index build
ResolvedProject
  ↓
QueryService
  ↓
MCP response
```

MCP response内の参照は、原則として名前解決済みのIDを使う。
Raw YAMLに書かれた未解決文字列は、diagnosticやsource表示のために必要な場合のみ補助情報として返す。

### 2.2 Python inspect風の語彙を採用する

LLMが既に学習している一般的な introspection 語彙に寄せるため、MCP responseはPythonの `inspect` に近い操作感を持つ。

| 語彙 | 意味 |
|---|---|
| `signature` | 対象objectの外形。params / returns / fields / endpoint等 |
| `doc` | YAMLの `note` に由来する自然言語説明。semantic contractだが機械検証済み構造ではない |
| `source` | 定義元file / line / column等 |
| `members` | 対象objectが内包する要素。sub task / fields / transitions等 |
| `references` | 対象objectが参照する、または対象objectを参照する関係 |
| `diagnostics` | warning / error / hint 等の診断情報 |

ただし、brewprint MCPはPython AST互換APIではない。
公開対象はsyntax treeではなく、`ResolvedProject` 上のsemantic objectである。

### 2.3 dependenciesではなくreferencesを中心語彙にする

MCP responseでは、依存・参照・逆参照をまとめて `references` と呼ぶ。

理由:

- `dependency` は「ビルド依存」「実行依存」「型依存」などに意味が寄りやすい
- brewprintでは `reads` / `writes` / `transition.action` / `model field type` / `scenario step` など、依存というより参照として読むべき関係が多い
- Python / IDE文脈の `references` に近い語彙の方がLLMが解釈しやすい

ADR-049により、外部MCP tool名は `get_references`、内部QueryService method名は `GetReferences` に統一する。
`get_deps` / `GetDeps` は採用しない。

### 2.4 構造情報とdocを分離する

MCP responseは、機械的に確定した構造情報と、`note` 由来の自然言語説明を分けて返す。

```json
{
  "signature": {
    "params": [
      { "name": "credentials", "model": "auth.model.credential" }
    ]
  },
  "doc": "認証情報を検証しトークンを発行する"
}
```

`doc` はLLMへのsemantic contractとして重要だが、機械検証済みの事実として扱ってはならない。

### 2.5 v1のreferencesはdirectのみ

ADR-048 / ADR-049に従い、MCP v1では完全なtransitive dependency graphを事前構築しない。

そのため `get_references` は、初期仕様では**直接referenceのみ**を返す。

transitive closure / depth指定 / dependency graph cacheは、QueryService vertical sliceで実需が出た時点で別途拡張する。

### 2.6 設計対話 coverage を拡張判断基準にする

MCP / QueryService は、単なる実装補助APIではなく、DAG / State Diagram / Sequence Diagram / ER / API Table / Wireframe などの図・viewを見ながらLLMと設計対話するためのquery layerである。

そのため、renderされた図やviewに現れる主要semantic objectは、原則としてMCPからquery可能にする。

対象例:

- task / model / store / state / event / actor
- model field
- transition
- sequence scenario view
- API Table view
- ER Diagram view
- implicit asset
- file-local sub task / branch / fork / join
- flow entry / flow wiring
- source file

すべてをMCP v1で即時実装する必要はない。ただし、今後のMCP拡張では「そのobjectが図やview上で利用者に見えており、会話対象になりうるか」を優先判断基準とする。

MCP responseは引き続きRaw YAML ASTを公開せず、ResolvedProject上のsemantic object queryとして返す。source snippet取得が必要な場合も、semantic objectに対応するsource補助情報として扱う。

> 由来: ADR-054 §決定

---

## 3. Common schema

### 3.1 Object selector

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
| `kind` | 任意 | string | 期待するkind。指定時、解決結果のkindが一致しなければerror |
| `file` | 任意 | FileID | private sub node等、file-local objectを指定する場合に使う |
| `local_id` | 任意 | string | `file` 内のlocal object ID。sub task等のprivate object参照に使う |

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

### 3.2 Selector support matrix

MCP v1のselector対応範囲は以下とする。

| object / kind | `get_signature` | `get_references` | `get_reference_tree` | `inspect` | status |
|---|---:|---:|---:|---:|---|
| `node: task` | yes | yes | yes | yes | supported |
| `node: model` | yes | yes | yes | yes | supported |
| `node: store` | yes | yes | yes | yes | supported |
| `node: state` | yes | yes | yes | yes | supported |
| `node: event` | yes | yes | yes | yes | supported |
| `node: actor` | yes | yes | yes | limited | supported / limited inspect |
| `view: sequence_diagram` | yes | yes | yes | yes | supported |
| `transition` | yes | yes | yes | yes | supported |
| `field` | yes | yes | yes | yes | supported |
| `file: node` | no | limited | limited | yes | supported / limited references |
| `file: state_file` | no | yes | yes | yes | supported |
| `file: sequence_diagram` | no | no | no | yes | supported |
| `file: api_table` | no | no | no | yes | supported |
| `file: er_diagram` | no | no | no | yes | supported |
| `file: render_index` | no | no | no | yes | supported |
| `asset` | yes | yes | yes | yes | supported |
| `view: api_table` | no | no | no | yes | supported; `list_endpoints` はcomputed route一覧専用 |
| `view: er_diagram` | no | no | no | yes | supported |
| private sub node | yes | yes | yes | yes | supported |
| `primitive` | no | no | no | no | reference target only |

statusの意味:

| status | 意味 |
|---|---|
| `supported` | MCP v1でquery対象として扱う |
| `supported / limited inspect` | signature / references は扱うが、専用inspectの情報量は限定的 |
| `partial` | 一部toolのみ対応する |
| `future` | 設計対話coverage上は候補だが、現時点では未実装 |
| `v1 optional` | spec上は許容するが、実装必須ではない |
| `reference target only` | reference targetとして返すが、直接query対象ではない |

`get_reference_tree` における `file: node` の `limited` は、`get_references(file: node)` の対応範囲に従い、file内nodeへのreferenceのみ展開することを意味する。
`primitive` はreference targetとして到達可能だが、traversal rootとしては扱わない。

> 由来: ADR-054 §決定

### 3.3 QualifiedID

モジュールスコープを持つnodeのQualifiedIDは、ADR-027に従う。

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

actorはADR-031によりproject globalであり、モジュールに属さない。

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

### 3.4 FileID

FileIDは、ADR-043で定義されるbrewprint projectの `yaml/` ディレクトリからの相対パスをslash正規化した文字列とする。

```text
auth/task/login.yaml
order/state.yaml
views/scenarios/checkout_flow.yaml
```

Windows上の `\\` はMCP responseでは `/` に正規化する。

### 3.5 Synthetic ID

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

transition IDは、ADR-035 / ADR-048の `(stateFileID, fromStateID, eventID, guard?)` に対応する文字列として以下を使う。

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
これはADR-035 / ADR-048のguard exact match方針と一致する。

### 3.6 SourceLocation

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

### 3.7 ObjectRef

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

### 3.8 TransitionRef

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

### 3.9 AssetRef

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

### 3.10 Diagnostic

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

## 4. Reference schema

### 4.1 Reference

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

### 4.2 Reference kind

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

M11ではこの方針を維持し、`flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over` は `Reference.kind` ではなく、flow inspect用の語彙として扱う。
これらは将来の `get_reference_tree` / `analyze_impact` の traversal 材料になりうるが、direct references v1 の返却対象には含めない。

### 4.3 Direction

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

## 5. Tool overview

MCP v1のquery toolは以下の7つとする。

| tool | 目的 | 主な利用場面 |
|---|---|---|
| `list_objects` | project内のsemantic object一覧を取得する | 実装・設計対話の起点として対象objectを探す |
| `get_signature` | object単体の外形を取得する | 実装前にtask/model/store等の型・I/Oを確認する |
| `get_source` | semantic objectに対応するYAML snippetを取得する | 設計対話中に定義元YAMLを確認する |
| `get_references` | objectの直接referenceを取得する | 影響範囲・依存・逆参照を確認する |
| `get_reference_tree` | objectからdepth制限つきでreference graphを辿る | 変更影響範囲や周辺objectをN hopで確認する |
| `inspect` | object kind別に実装判断用の文脈を取得する | Claude Code等が実装・修正時に読む |
| `list_endpoints` | API Table viewに基づくendpoint一覧を取得する | API実装・ルーティング確認 |

---

### 5.1 `list_objects`

#### 5.1.1 Purpose

`list_objects` は、project内のsemantic object一覧を返す。

返す対象:

- `node`
- `view`
- `transition`
- `field`

`list_objects` は探索用toolであり、各objectの詳細なsignature / references / inspect情報は返さない。詳細が必要な場合は、返された `object` / `kind` / `id` をselectorとして `get_signature` / `get_references` / `inspect` を呼ぶ。

#### 5.1.2 Input
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

#### 5.1.3 Output
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

## 6. `get_signature`

### 6.1 Purpose

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

### 6.2 Input

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

### 6.3 Output envelope

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

### 6.4 task signature

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

### 6.5 model signature

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

### 6.6 store signature

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

### 6.7 event signature

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

### 6.8 state signature

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

### 6.9 transition signature

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

### 6.10 field signature

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

## 7. `get_source`

### 7.1 Purpose

`get_source` は、対象semantic objectに対応するYAML source snippetを返す。

返すもの:

- object identity
- source file / range
- YAML snippet
- fallbackした場合の理由を示すdiagnostic

返さないもの:

- Raw YAML AST
- semantic build前の未解決構造全体
- project外fileの内容
- renderer output

`get_source` はADR-054の方針に従い、Raw YAML AST公開APIではなく、ResolvedProject上のsemantic objectに紐づくsource補助情報として扱う。

### 7.2 Input

```json
{
  "selector": {
    "id": "auth.task.login"
  },
  "fallback": "file"
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `selector` | ✓ | Object selector |
| `fallback` | 任意 | `file` / `error`。省略時は `file` と同等 |

`fallback=file` または省略時は、object単位のrangeが特定できない場合に、同じFileIDのYAML全体を返し、`diagnostics[]` に `source_range_unavailable` warningを入れる。
`fallback=error` の場合は、object単位のrangeが特定できないとtool errorを返す。

### 7.3 Output

```json
{
  "object": {
    "object": "node",
    "kind": "task",
    "id": "auth.task.login",
    "qualified_id": "auth.task.login",
    "label": "login",
    "file": "auth/task/login.yaml"
  },
  "source": {
    "file": "auth/task/login.yaml",
    "line": 3,
    "column": 5,
    "end_line": 18,
    "end_column": 1
  },
  "snippet": {
    "language": "yaml",
    "text": "  - id: login\n    type: task\n    ..."
  },
  "diagnostics": []
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `object` | ✓ | 対象ObjectRef |
| `source` | ✓ | SourceLocation。line / columnが取得できない場合は `file` のみでもよい |
| `snippet` | ✓ | `language: yaml` と snippet text |
| `fallback` | 任意 | fallbackした場合は `file` |
| `diagnostics` | ✓ | Diagnostic list |

### 7.4 Selector support

`get_source` は、MCP v1でquery可能なsemantic objectを対象にする。

初期実装でsnippet rangeを返す対象:

- `node` / private sub node: `nodes[]` 内の該当item
- `field`: parent modelの `fields[]` 内の該当item
- `transition`: `transitions[]` 内の該当item
- `asset`: producer nodeの `returns` block。特定不能時はproducer nodeへfallbackしてよい
- `view`: view file全体
- `file`: file全体

source line/columnが取得できない実装、または対応objectの局所rangeを特定できない実装では、`fallback=file` により同一FileID全体を返してよい。
この場合は `diagnostics[]` に以下を入れる。

```json
{
  "severity": "warning",
  "code": "source_range_unavailable",
  "file": "auth/task/login.yaml",
  "message": "source range is unavailable; returned whole file"
}
```

> 由来: ADR-054 §決定 §5

---

## 8. `get_references`

### 8.1 Purpose

`get_references` は、対象objectの直接referenceを返す。

MCP v1ではdirect referencesのみを返す。

### 8.2 Input

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

### 8.3 Output

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

### 8.4 depth

`depth` は常に `1` を返す。

MCP v1では、`get_references` inputに `depth` を持たない。
transitive reference traversal は、ADR-055に従い、別tool `get_reference_tree` で扱う。

---

## 9. `get_reference_tree`

### 9.1 Purpose

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

変更種別を含む影響判断は、将来の `analyze_impact` で扱う。

### 9.2 Input

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

### 9.3 Output

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

### 9.4 Node entry

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

### 9.5 Edge entry

`edges[]` は `Reference` と同じ基本形に `depth` を加えたものとする。

`edges[].depth` は、そのedgeを発見した traversal hop を表す。
rootから1 hop目で発見されたedgeは `depth: 1` とする。

### 9.6 Traversal semantics

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

### 9.7 Selector support

`get_reference_tree` のroot selectorは、基本的に `get_references` の対応selectorを起点にする。

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

### 9.8 Flow wiring references

`get_reference_tree` v1は、新しいreference kindを追加しない。

flow wiring（`flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over`）は、引き続き `inspect(task).members.flow.entries` 内のflow inspect用語彙として扱い、`get_reference_tree` のtraversal対象には含めない。

将来、flow wiringを影響分析に含める場合は、`get_reference_tree` のreference kind拡張、または `analyze_impact` 側の補完材料として扱う。

---

## 10. `inspect`

### 10.1 Purpose

`inspect` は、対象objectの実装判断に必要な周辺文脈をkind別にまとめて返す。

`get_signature` が薄い外形確認であるのに対し、`inspect` はLLMが実装・修正・レビュー時に読む濃い文脈取得toolである。

### 10.2 Input

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

### 10.3 Common output shape

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

### 10.4 task inspect

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

#### flow.entries schema status

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

#### sub task reads/writes

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

### 10.5 store inspect

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

### 10.6 model inspect

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

### 10.7 state inspect

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

### 10.8 event inspect

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

### 10.9 scenario inspect

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

### 10.10 transition inspect

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

### 10.11 field inspect

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

### 10.12 API Table inspect

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

### 10.13 ER Diagram inspect

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

### 10.14 file inspect

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

## 11. `list_endpoints`

### 11.1 Purpose

`list_endpoints` は、API Table view YAMLに基づいてendpoint一覧を返す。

`task(endpoint=true)` を単純列挙するだけではなく、ADR-028のroute合成規則に従いfull pathを返す。

### 11.2 Input

```json
{
  "api_table_id": "ec_api"
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `api_table_id` | 任意 | API Table view ID。省略時はproject内の全API Tableを返す |

API Tableが複数存在し、`api_table_id` が省略された場合は、全API Tableを `tables[]` に分けて返す。

### 11.3 Output

```json
{
  "tables": [
    {
      "id": "ec_api",
      "http_root_path": "/api",
      "sections": [
        {
          "module": "auth",
          "include_submodules": true,
          "endpoints": [
            {
              "method": "POST",
              "path": "/api/login",
              "leaf_path": "login",
              "task": "auth.task.login",
              "params": "auth.model.login_request",
              "returns": "auth.model.token",
              "source": {
                "file": "auth/task/login.yaml"
              }
            }
          ]
        }
      ]
    }
  ],
  "diagnostics": []
}
```

この例ではsection起点moduleが `auth` のため、ADR-028のroute合成規則により、section起点moduleからの相対module pathは空になる。
そのためfull pathは `/api/auth/login` ではなく `/api/login` になる。
`/api/auth/login` を返したい場合は、API Table view側で `http_root_path: /api/auth` とするか、section起点moduleを上位moduleにする。

### 11.4 endpoint object

| フィールド | 必須 | 内容 |
|---|---:|---|
| `method` | ✓ | HTTP method |
| `path` | ✓ | API Table viewにより合成されたfull path |
| `leaf_path` | ✓ | task側のleaf path。省略時はtask.id由来 |
| `task` | ✓ | endpoint task QualifiedID |
| `params` | 任意 | request model QualifiedID |
| `returns` | 任意 | response model QualifiedID |
| `source` | 任意 | endpoint taskのSourceLocation |

---

## 12. Error model

### 12.1 MCP-level error vs diagnostic

MCP toolの実行自体が成立しない場合はtool errorを返す。

例:

- projectがsemantic validationを通過していない
- selectorの形式が壊れている
- 対象objectが存在しない
- guard未指定でtransitionが曖昧

一方、tool実行は成立したが注意すべき情報がある場合は `diagnostics` に入れる。

例:

- source lineが取得できない
- uncovered moduleが暗黙groupになった
- noteが存在しない
- optionalな周辺情報が未実装

### 12.2 Error code

MCP v1で定義するerror code:

| code | 意味 |
|---|---|
| `project_invalid` | semantic buildに失敗しておりqueryできない |
| `invalid_selector` | selectorの形式が不正 |
| `not_found` | 対象objectが存在しない |
| `kind_mismatch` | selector.kind と解決結果のkindが一致しない |
| `ambiguous` | 候補が複数あり一意に解決できない |
| `unsupported_object` | v1ではquery対象外のobject |
| `unsupported_detail` | `detail` の値が未対応 |
| `unsupported_direction` | `direction` の値が未対応 |
| `invalid_depth` | traversal depth が未対応範囲外 |
| `internal_error` | 実装内部エラー |

### 12.3 Error payload

```json
{
  "error": {
    "code": "not_found",
    "message": "object not found: auth.task.missing_login",
    "selector": {
      "id": "auth.task.missing_login"
    },
    "diagnostics": []
  }
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `code` | ✓ | error code |
| `message` | ✓ | human-readable message |
| `selector` | 任意 | 入力selector |
| `diagnostics` | 任意 | 関連diagnostic |

---

## 13. Tool selection guidance for LLM

LLMは以下の使い分けを基本とする。

| 状況 | 使うtool |
|---|---|
| 対象nodeのI/Oだけ確認したい | `get_signature` |
| 対象objectの定義元YAML snippetを確認したい | `get_source` |
| 何に依存しているか / 何から参照されているか確認したい | `get_references` |
| 変更影響範囲や周辺objectをN hopで確認したい | `get_reference_tree` |
| 実装・修正・レビューのために周辺文脈が必要 | `inspect` |
| API route一覧が必要 | `list_endpoints` |

原則:

- 実装前にはまず `inspect` を使う
- 小さな型確認だけなら `get_signature` を使う
- 直接参照確認では `get_references(direction="in")` または `both` を使う
- N hopの影響範囲確認では `get_reference_tree` を使い、`direction` と `depth` を明示する
- Raw YAMLを直接読む前に、まず `get_source` でsemantic objectに対応するsnippetを確認する

---

## 14. Versioning / future extensions

MCP v1では以下を未定義とする。

- unbounded transitive references
- reference graphの永続cache
- renderer outputを返すMCP tool
- code generation用tool

設計対話coverageを広げるための将来候補:

| 候補tool / selector | 目的 | 優先度 |
|---|---|---:|
| `list_objects` | project内objectの検索・一覧。LLMがquery対象を発見する入口 | high |
| `inspect(file)` | YAML file単位で定義内容・main node・sub node・view種別・diagnosticsを把握する | high |
| `inspect(view: api_table)` | API Table viewが集約するmodules / endpoints / computed routesを把握する | high |
| `inspect(view: er_diagram)` | ER Diagram viewが集約するmodules / stores / models / FKを把握する | high |
| flow wiring references | DAG上のflow step / param wiringをreferenceとして扱う | medium |
| `analyze_impact` | 変更種別つき影響分析。direct references / reference tree / flow wiring / render output mapping を材料に、影響範囲・重要度・確認事項を返す | medium |
| `search_notes` | note/docに対するsemantic search | low |

これらは必要になった時点で `docs/spec/mcp.md` を更新し、実装タスクは `docs/TASKS.md` で管理する。
既存ADRの方針を変更するほどの設計転換がある場合のみ、新ADRを起票する。
