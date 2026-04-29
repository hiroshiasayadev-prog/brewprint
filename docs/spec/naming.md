---
scope: docs/spec/naming.md
status: confirmed
last_updated: 2026-04-30
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

> 由来: ADR-002 §決定, ADR-027 §モジュールネストを許容する

## 2. QualifiedID

ノードのフルパス参照は以下の形式を取る。

```
<モジュールパス>.<ノード種別>.<ID>
```

- `<モジュールパス>`: 1段以上の任意の深さのドット区切りパス（例: `auth`, `commerce.cart`）
- `<ノード種別>`: 後述の予約語（sentinel）
- `<ID>`: モジュール内のノードID

例:

```yaml
auth.task.login
commerce.cart.store.cart_db
analysis.state.session
```

> 由来: ADR-027 §命名規則を拡張する

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

> 由来: ADR-027 §ノード種別予約語をsentinelとして正式化

## 4. 同モジュール内ID直書き

同モジュール内のノード参照では QualifiedID のフルパスは要求しない。ID 直書きで解決される。
モジュールを跨ぐ参照は QualifiedID（フルパス）必須。

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

> 由来: ADR-003 §決定, ADR-027 §フルパス必須（同モジュール内ID直書きは継承）

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

> 由来: ADR-031 §決定, ADR-025 §決定

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

> 由来: ADR-033 §決定, ADR-027 §命名規則を拡張する

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

> 由来: ADR-020 §決定, ADR-027 §命名規則を拡張する

## 8. resolve後のID表現

名前解決後の semantic model / reference index / MCP response では、入力がbare記法であっても、原則として解決済みの正規化IDを返す。

| YAML上の入力 | 文脈 | resolve後 |
|---|---|---|
| `order_db` | `order` module内の `reads` / `writes` | `order.store.order_db` |
| `order.id` | `order` module内の `fk:` | `order.model.order.id` |
| `cart.store.cart_session` | module跨ぎ store参照 | `cart.store.cart_session` |
| `commerce.order.model.order.id` | module跨ぎ FK参照 | `commerce.order.model.order.id` |
| `stripe` | actor参照 | `stripe` |

この方針により、YAML authoringでは同一module内参照を簡潔に書ける一方、実装・MCP・diagnosticsでは一意な参照IDを扱える。
