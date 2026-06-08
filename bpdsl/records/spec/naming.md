---
scope: docs/spec/naming.md
status: confirmed
last_updated: 2026-05-31
summary: >
  brewprint の名前空間とID解決ルールを定義する。
  モジュール=フォルダ階層、QualifiedID形式、同モジュール内ID直書き、sentinel方式パース、
  actor global namespace、FK field参照、reads/writesのstore参照解決を含む。
depends_on:
  - docs/adr/002-folder-as-namespace.md
  - docs/adr/003-name-resolution-rules.md
  - docs/adr/020-cross-edge-management.md
  - docs/adr/027-module-nesting-and-name-resolution.md
  - docs/adr/031-actor-global-definition.md
  - docs/adr/033-fk-name-resolution.md
  - docs/adr/058-subnode-file-private-scope-enforcement.md
  - docs/adr/070-model-visibility-file-private-helper-model.md
  - docs/adr/078-mcp-semantic-anchor-synthetic-id-policy.md
---

# 名前解決仕様

## 1. モジュールとフォルダ階層

`yaml/` 配下のディレクトリ階層がそのままモジュール階層となる。

```
yaml/
  auth/
    dag.yaml        ← モジュール: auth
  commerce/
    cart/
      dag.yaml      ← モジュール: commerce.cart
    order/
      dag.yaml      ← モジュール: commerce.order
```

ディレクトリの深さに制限はない（任意の段数のネストを許容）。モジュール名は `.` 区切りで表現する。

> 由来: V01-ADR-002 §決定, V01-ADR-027 §モジュールネストを許容する

## 2. QualifiedID

ノードのフルパス参照は以下の形式を取る。

```
<モジュールパス>.<ノード種別>.<ID>
```

- `<モジュールパス>`: 1段以上の任意の深さのドット区切りパス（例: `auth`, `commerce.cart`）
- `<ノード種別>`: 後述の予約語（sentinel）
- `<ID>`: public node ID。メインノードのみが public QualifiedID の対象となる

サブノードは file-private local ID を持つが、public QualifiedID は持たない。したがって、別 file の同名サブノード local ID とは衝突せず、外部 YAML から `<module>.<type>.<sub-node-id>` 形式で参照することもできない。

Task-file helper model の基本 semantics は [nodes.md](./nodes.md#task-file-private-helper-model-semantics) が定義する。本仕様では、`<module>.model.<id>` QualifiedID が public model のみを対象とすることを定義する。

MCP query layer が file-private / generated object を返す場合、V01-ADR-078 の方針により `<semantic-anchor-id>#<local-id>` 形式の synthetic ID を使える。ただしこれは public QualifiedID ではなく、YAML authoring の外部参照形式でもない。MCP schema / ObjectRef migration details are outside this spec section and are not changed by this resolver work.

例:

```yaml
auth.task.login
commerce.cart.store.cart_db
analysis.state.session
```

> 由来: V01-ADR-027 §命名規則を拡張する

## 3. ノード種別予約語（sentinel）

QualifiedID のパースは、ノード種別予約語の出現位置でモジュールパスと ID の境界を判定する。

| 予約語 | ノード種別 |
|--------|-----------|
| `task` | タスク |
| `model` | モデル |
| `store` | ストア |
| `actor` | アクター |
| `event` | イベント |
| `state` | ステート |
| `branch` | ブランチ |
| `fork` | フォーク |
| `join` | ジョイン |

パース例:

```
auth.oauth.task.login
           ↑
           sentinel を発見 → 左側がモジュールパス、右側が ID
→ module: auth.oauth
→ type:   task
→ id:     login
```

> 由来: V01-ADR-027 §ノード種別予約語をsentinelとして正式化

## 4. 同モジュール内ID直書き

同モジュール内のメインノード参照では QualifiedID のフルパスは要求しない。ID 直書きで解決される。
モジュールを跨ぐ参照は QualifiedID（フルパス）必須。

同一 file 内に書かれた `flow.step` / `reads` / `writes` 等の bare ID は、まず同一 file 内の file-private sub node / source を優先して解決する。該当がない場合のみ、同一 module のメインノードへフォールバックする。

TypeRef における bare model name の解決は [type-ref.md](./type-ref.md) §4 が所有する。本仕様では、task-file helper model と public model の名前衝突 rule だけを定義する。

```yaml
# 同モジュール内（auth/dag.yaml内での参照）
flow:
  - from: login           # auth.task.login に解決される
    to: token

# モジュール跨ぎ（フルパス必須）
edges:
  - from: auth.task.login
    to: analysis.state.session
```

> 由来: V01-ADR-003 §決定, V01-ADR-027 §フルパス必須（同モジュール内ID直書きは継承）

### 4.1 task-file helper model の名前衝突

Task-file helper model の visibility / identity / reference scope は [nodes.md](./nodes.md#task-file-private-helper-model-semantics) を正とする。本節では、TypeRef の可読性を守るための名前衝突 rule を定義する。

| case | result |
|---|---|
| 同一 file 内で helper model 同士が同じ id を持つ | invalid |
| 同一 file 内で main node / private sub node / helper model の local id が衝突する | invalid |
| 同一 module 内の public model と task-file helper model が同じ id を持つ | invalid |
| 同一 module 内の別 file にある task-file helper model 同士が同じ id を持つ | valid |
| 別 module の public model と task-file helper model が同じ id を持つ | valid |
| 別 module の task-file helper model 同士が同じ id を持つ | valid |

外部から再利用する必要が出た helper shape は public model へ昇格させる。

> 由来: V01-ADR-070 §7〜§8

## 5. actor の global namespace

`actor` はプロジェクトglobalなノード種別であり、モジュールに属さない。
そのため actor の QualifiedID は `<module-path>.actor.<id>` 形式を取らず、常に actor ID そのものを参照IDとして使う。

```yaml
nodes:
  - id: stripe
    type: actor
    note: "外部決済サービス"
```

上記の actor は、どの module からも以下のように参照する。

```yaml
actor: stripe
```

actor ID はプロジェクト全体で一意でなければならない。複数ファイルに同じ actor ID が定義された場合は validation error とする。

> 由来: V01-ADR-031 §決定, V01-ADR-025 §決定

## 6. FK field参照の解決

model field の `fk:` は、参照先 model field を指す。
FK参照には、同一module内でのみ使える bare FK と、moduleを跨ぐ場合に使う Qualified field reference の2形式がある。

### 6.1 同一module内 bare FK

同一module内の model field を参照する場合は、以下の短縮形式を許容する。

```text
<model-id>.<field-name>
```

例:

```yaml
# yaml/order/model/order_item.yaml
nodes:
  - id: order_item
    type: model
    kind: struct
    fields:
      - name: order_id
        type: str
        fk: order.id
```

上記の `fk: order.id` は、現在のmoduleが `order` の場合、resolve後に以下へ正規化される。

```text
order.model.order.id
```

つまり、bare FK は以下の規則で解決する。

```text
<model-id>.<field-name>
→ <current-module>.model.<model-id>.<field-name>
```

### 6.2 module跨ぎ FK

moduleを跨いで別moduleの model field を参照する場合は、Qualified field reference を必須とする。

```text
<module-path>.model.<model-id>.<field-name>
```

例:

```yaml
# yaml/payment/model/payment_event.yaml
nodes:
  - id: payment_event
    type: model
    kind: struct
    fields:
      - name: order_id
        type: str
        fk: order.model.order.id
```

nested module を参照する場合も同じ形式を使う。

```yaml
fk: commerce.order.model.order.id
```

moduleを跨ぐ参照で bare FK を使うことはできない。`fk: order.id` のようなbare FKは、常に現在のmodule内の `model.order.id` として解釈される。

### 6.3 resolve後の正規化

YAML上の入力が bare FK であっても、semantic model / reference index / MCP response では常に正規化済みの field ID を使う。

```text
入力: fk: order.id
解決: order.model.order.id
reference: field_fk from order.model.order_item.order_id to order.model.order.id
```

この正規化により、`get_references` などの逆参照取得では、bare FK と Qualified field reference を同一形式で扱える。

> 由来: V01-ADR-033 §決定, V01-ADR-027 §命名規則を拡張する

## 7. reads / writes の store参照解決

`task.reads` / `task.writes` は store node を参照する。
store参照も、同一module内は bare ID を許容し、moduleを跨ぐ場合は QualifiedID を必須とする。

### 7.1 同一module内 store参照

同一module内の store を参照する場合は、store ID をそのまま書ける。

```yaml
# yaml/order/task/checkout.yaml
nodes:
  - id: checkout
    type: task
    reads:
      - order_db
    writes:
      - order_db
```

現在のmoduleが `order` の場合、上記は resolve後に以下へ正規化される。

```text
order.store.order_db
```

### 7.2 module跨ぎ store参照

別moduleの store を参照する場合は QualifiedID を使う。

```yaml
# yaml/order/task/checkout.yaml
nodes:
  - id: checkout
    type: task
    reads:
      - cart.store.cart_session
      - auth.store.user_db
    writes:
      - order_db
```

resolve後の参照は以下になる。

```text
reads:
  - cart.store.cart_session
  - auth.store.user_db
writes:
  - order.store.order_db
```

moduleを跨ぐ参照で bare store ID を使うことはできない。bare store ID は、常に現在のmodule内の store として解釈される。

> 由来: V01-ADR-020 §決定, V01-ADR-027 §命名規則を拡張する

## 8. resolve後のID表現

名前解決後の semantic model / reference index / MCP response では、入力がbare記法であっても、原則として解決済みの正規化IDを返す。

| YAML上の入力 | 文脈 | resolve後 |
|---|---|---|
| `order_db` | `order` module内の `reads` / `writes` | `order.store.order_db` |
| `order.id` | `order` module内の `fk:` | `order.model.order.id` |
| `cart.store.cart_session` | module跨ぎ store参照 | `cart.store.cart_session` |
| `commerce.order.model.order.id` | module跨ぎ FK参照 | `commerce.order.model.order.id` |
| `stripe` | actor参照 | `stripe` |
| `validate_request` | `mcp/task/get_signature.yaml` 内の private sub task 参照 | file-local sub node identity。public QualifiedID へ正規化しない |
| `mcp.task.get_signature#validate_request` | MCP query layer 上の private sub task synthetic ID | semantic anchor based synthetic ID。public QualifiedID ではない |

この方針により、YAML authoringでは同一module内参照を簡潔に書ける一方、実装・MCP・diagnosticsでは一意な参照IDを扱える。

ただし、すべての解決済みIDが public QualifiedID になるわけではない。file-private sub node は public QualifiedID を持たず、同一 file 内の local identity として扱う。MCP response 上で安定参照が必要な場合のみ、V01-ADR-078 の `<semantic-anchor-id>#<local-id>` synthetic ID を使う。MCP schema / ObjectRef migration details are outside this spec section and are not changed by this resolver work.
