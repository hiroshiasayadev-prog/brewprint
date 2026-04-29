---
scope: docs/spec/mcp.md
status: draft
last_updated: 2026-04-27
summary: >
  brewprintのMCP query tool外部仕様。
  Python inspectに近い語彙で、ResolvedProject上のsemantic objectに対する
  signature / references / inspect / endpoint query のinput/outputを定義する。
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
| `id` | ✓ | string | 対象objectのID。通常はQualifiedID。actorはglobal actor ID。scenario等のview objectはview固有ID。transitionやprivate sub nodeは後述のsynthetic IDを使う |
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

MCP v1では、private sub nodeの直接問い合わせは実装任意とする。
ただし `inspect(main task)` の `members.sub_tasks` には、同一ファイル内のsub task情報を含める。

### 3.2 QualifiedID

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

### 3.3 FileID

FileIDは、ADR-043で定義されるbrewprint projectの `yaml/` ディレクトリからの相対パスをslash正規化した文字列とする。

```text
auth/task/login.yaml
order/state.yaml
views/scenarios/checkout_flow.yaml
```

Windows上の `\\` はMCP responseでは `/` に正規化する。

### 3.4 Synthetic ID

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

### 3.5 SourceLocation

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

### 3.6 ObjectRef

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

### 3.7 TransitionRef

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

### 3.8 AssetRef

`asset` はYAML上に独立ファイルを持たず、taskの `returns` から暗黙生成される。
MCP v1では、implicit assetをグローバルにselectableなQualifiedIDとしては要求しない。
代わりに、`AssetRef` としてproducer contextつきで返す。

```json
{
  "object": "asset",
  "name": "auth_token",
  "producer": "auth.task.login",
  "model": "auth.model.token",
  "scope_file": "auth/task/login.yaml"
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `object` | ✓ | `asset` 固定 |
| `name` | ✓ | `task.returns.name` |
| `producer` | ✓ | assetを生成するtaskのQualifiedIDまたはfile-local synthetic ID |
| `model` | ✓ | assetのmodel QualifiedIDまたはprimitive |
| `scope_file` | 任意 | assetが生じるFileID |

`task.returns.name` はDAG / MCP上のasset IDとして扱うが、sub taskや別fileで同名returnsがあり得るため、MCP responseではproducer contextと一緒に返す。

将来、implicit assetを直接query対象にする必要が出た場合は、stable synthetic ID形式を別途spec/ADRで定義する。

### 3.9 Diagnostic

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
flow wiringはDAG file内部の局所構造であり、MCP v1では `inspect(task).members.flow.entries` のdraft schema内に閉じる。

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

MCP v1のquery toolは以下の4つとする。

| tool | 目的 | 主な利用場面 |
|---|---|---|
| `get_signature` | object単体の外形を取得する | 実装前にtask/model/store等の型・I/Oを確認する |
| `get_references` | objectの直接referenceを取得する | 影響範囲・依存・逆参照を確認する |
| `inspect` | object kind別に実装判断用の文脈を取得する | Claude Code等が実装・修正時に読む |
| `list_endpoints` | API Table viewに基づくendpoint一覧を取得する | API実装・ルーティング確認 |

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

## 7. `get_references`

### 7.1 Purpose

`get_references` は、対象objectの直接referenceを返す。

MCP v1ではdirect referencesのみを返す。

### 7.2 Input

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

### 7.3 Output

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

### 7.4 depth

`depth` は常に `1` を返す。

MCP v1では、inputに `depth` を持たない。
将来、transitive referencesが必要になった場合は以下のどちらかで拡張する。

- `depth` inputを追加する
- `get_reference_tree` 等の別toolを追加する

---

## 8. `inspect`

### 8.1 Purpose

`inspect` は、対象objectの実装判断に必要な周辺文脈をkind別にまとめて返す。

`get_signature` が薄い外形確認であるのに対し、`inspect` はLLMが実装・修正・レビュー時に読む濃い文脈取得toolである。

### 8.2 Input

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

### 8.3 Common output shape

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

### 8.4 task inspect

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
        { "kind": "step", "step": "build_order" },
        { "kind": "step", "step": "reserve_inventory" }
      ],
      "schema_status": "draft"
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

`members.flow.entries` の詳細schemaはMCP v1ではdraft扱いとする。
DAG vertical slice実装中に、`step` / `branch` / `fork` / `foreach` のview-specific modelと合わせて確定する。

MCP v1で保証するのは以下のみ。

- `members.flow.file` はflow定義元FileID
- `members.flow.entries[]` はflow内に登場するentryの概略順序
- 各entryは少なくとも `kind` を持つ

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

### 8.5 store inspect

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

### 8.6 model inspect

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

### 8.7 state inspect

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

### 8.8 event inspect

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

### 8.9 scenario inspect

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

### 8.10 transition inspect

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

### 8.11 field inspect

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

---

## 9. `list_endpoints`

### 9.1 Purpose

`list_endpoints` は、API Table view YAMLに基づいてendpoint一覧を返す。

`task(endpoint=true)` を単純列挙するだけではなく、ADR-028のroute合成規則に従いfull pathを返す。

### 9.2 Input

```json
{
  "api_table_id": "ec_api"
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `api_table_id` | 任意 | API Table view ID。省略時はproject内の全API Tableを返す |

API Tableが複数存在し、`api_table_id` が省略された場合は、全API Tableを `tables[]` に分けて返す。

### 9.3 Output

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

### 9.4 endpoint object

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

## 10. Error model

### 10.1 MCP-level error vs diagnostic

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

### 10.2 Error code

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
| `internal_error` | 実装内部エラー |

### 10.3 Error payload

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

## 11. Tool selection guidance for LLM

LLMは以下の使い分けを基本とする。

| 状況 | 使うtool |
|---|---|
| 対象nodeのI/Oだけ確認したい | `get_signature` |
| 何に依存しているか / 何から参照されているか確認したい | `get_references` |
| 実装・修正・レビューのために周辺文脈が必要 | `inspect` |
| API route一覧が必要 | `list_endpoints` |

原則:

- 実装前にはまず `inspect` を使う
- 小さな型確認だけなら `get_signature` を使う
- 影響範囲確認では `get_references(direction="in")` または `both` を使う
- Raw YAMLを直接読むのは、MCP responseのsource位置を確認したあとでよい

---

## 12. Versioning / future extensions

MCP v1では以下を未定義とする。

- `get_source`
- transitive references
- reference graphの永続cache
- implicit assetの直接selector
- renderer outputを返すMCP tool
- code generation用tool

将来候補:

| 候補tool | 目的 |
|---|---|
| `get_source` | `inspect.getsource` 相当。対象objectのYAML snippetを返す |
| `get_reference_tree` | depth指定つきreference traversal |
| `list_objects` | project内objectの検索・一覧 |
| `search_notes` | note/docに対するsemantic search |

これらは必要になった時点でspec更新またはADR起票を行う。
