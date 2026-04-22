# UC-001 EC Checkout Flow ― Sonnet向け作業指示書

> このdocは UC-001 をSonnetが書ききるための引き継ぎ。完成後は削除してよい。

---

## 0. 前提

- doc-policy.md / ADR一覧把握 はセッション冒頭で必ず実施（プロジェクト指示通り）
- このUCは **「全view・全ノード種別を1ユースケースで網羅する検証ケース」** という性格を持つ。テスト的な機能網羅性が最優先

---

## 1. ゴール

`docs/uc/001-ec-checkout-flow/` 配下に以下を作成：

```
docs/uc/001-ec-checkout-flow/
  README.md                         ← UCの概要・目次・主要render結果
  yaml/                             ← brewprint YAML群（Goテストfixture流用想定）
    actors.yaml
    views/
      api_table.yaml
      scenarios/
        checkout_flow.yaml
        payment_webhook_flow.yaml
    auth/
      model/credential.yaml
      model/token.yaml
      model/login_form.yaml
      model/login_log.yaml
      model/request_context.yaml
      store/user_db.yaml
      store/session_store.yaml
      task/login.yaml
      state.yaml
    catalog/
      model/item.yaml
      model/item_list.yaml
      store/item_collection.yaml
      store/inventory_db.yaml
      task/get_items.yaml
    cart/
      model/cart.yaml
      model/cart_item.yaml
      store/cart_session.yaml
      task/add_to_cart.yaml
      task/validate_cart.yaml
    order/
      model/order.yaml
      model/order_item.yaml
      model/address.yaml
      store/order_db.yaml
      task/checkout.yaml
      task/process_order.yaml
      state.yaml
    payment/
      model/payment_event.yaml
      task/process_payment.yaml
    inventory/
      state.yaml
```

ファイル数 25〜27 程度。

---

## 2. 機能カバレッジ要件（チェックリスト）

このUCで以下を**すべて**使う。書き終わったら自分でチェックすること。

### ノード種別

- [ ] `task`（main: true 持ち、複数）
- [ ] `task`（endpoint: true 持ち）
- [ ] `task`（endpoint: false の内部task、サブノード扱い）
- [ ] `model` kind: struct（fields・pk・fk・unique含む）
- [ ] `model` kind: list
- [ ] `model` kind: dict
- [ ] `asset`（task.returnsから暗黙生成。明示記述は不要）
- [ ] `store` kind: db
- [ ] `store` kind: session
- [ ] `store` kind: collection（noteで検索条件記述）
- [ ] `store` kind: context（どこかで1個使う）
- [ ] `actor`（global, ファイル名任意 ― actors.yamlに集約）
- [ ] `event` source: ui
- [ ] `event` source: external（actor必須）
- [ ] `event` source: er（watches必須）
- [ ] `state`（initial: true / final: true 含む）
- [ ] `branch`（cases複数、order/process_order.yaml）
- [ ] `fork` + `join`（order/checkout.yaml）

### flow構文

- [ ] `step` エントリ
- [ ] `fork` + `join` エントリ
- [ ] `branch` エントリ
- [ ] `foreach` エントリ（cart/validate_cart.yaml で使用）
- [ ] `$params.field` シジル
- [ ] `$item` シジル

### task内フィールド

- [ ] `reads`
- [ ] `writes`
- [ ] `reads` + `writes` 両方
- [ ] `initializes`（task/login.yaml で login_log を初期化、等）

### model内フィールド

- [ ] `pk: true`
- [ ] `fk: <model>.<field>`
- [ ] `fk` + `unique: true`（1:1）
- [ ] N:M（中間model `order_item.yaml` で2本のfk）
- [ ] JSON埋め込み（fk なしのmodel ID参照、`address` 等）
- [ ] primitive全種（str, int, float, bool, datetime あたりは最低使う）

### state / event

- [ ] `transition.action` 同一ファイル内参照
- [ ] `transition.action` クロスファイル参照（フルパス: `auth.task.login` 等）
- [ ] `transition.guard`
- [ ] `state.wireframe`（複数stateで定義。container/leaf/interactive要素を使う）

### wireframe 要素

- [ ] container: col
- [ ] container: row
- [ ] container: card
- [ ] container: grid（cols, span）
- [ ] container: header / sidebar / footer のいずれか
- [ ] leaf interactive: button（fires持ち）
- [ ] leaf interactive: input
- [ ] leaf interactive: password
- [ ] leaf interactive: select / checkbox / radio のいずれか
- [ ] leaf non-interactive: text
- [ ] leaf non-interactive: badge
- [ ] leaf non-interactive: image / icon / divider のいずれか
- [ ] `disabled: true` のinteractive要素（loading state等）

### view定義

- [ ] `as: api_table`（views/api_table.yaml）
- [ ] `as: sequence_diagram` × 2本（views/scenarios/）

### クロスモジュール参照

- [ ] フルパス参照を最低1箇所（例: order の transition.action から `auth.task.login`、または cart task が catalog の store を reads する等）

### モジュールネスト

- 今回は1段モジュールでよい（ADR-027のネスト機能の検証は別UCで）

---

## 3. ドメイン設計（決定済み）

### モジュール責務

| モジュール | 責務 |
|---|---|
| `auth` | ログイン認証・セッション管理 |
| `catalog` | 商品一覧・在庫DB |
| `cart` | カート（セッション保持）・カートバリデーション |
| `order` | 注文処理・チェックアウト・注文FSM |
| `payment` | Stripe webhook受信・決済処理 |
| `inventory` | 在庫FSM（er source eventの起点） |

### actors（ec/actors.yaml）

```yaml
nodes:
  - id: end_user
    type: actor
    note: \"ECサイトを利用するエンドユーザー\"
  - id: stripe
    type: actor
    note: \"外部決済サービス\"
```

### task一覧（mainのもの）

| task | endpoint | 特徴 |
|---|---|---|
| `auth.task.login` | POST /login | initializes 使用 / reads+writes |
| `catalog.task.get_items` | GET /items | reads (collection) のみ |
| `cart.task.add_to_cart` | POST /cart/items | reads (collection) + writes (session) |
| `cart.task.validate_cart` | （内部呼び出しのみ。endpoint: false） | **foreach 使用**, sub task `validate_item` |
| `order.task.checkout` | POST /checkout | **fork+join 使用**, reads (multiple stores) + writes (multiple stores) |
| `order.task.process_order` | （内部）endpoint: false | **branch 使用**（在庫あり/なしで分岐） |
| `payment.task.process_payment` | POST /webhooks/stripe | external eventハンドラ。writes (order_db) |

### state.yaml

| ファイル | FSM対象 | 主要state | 主要event |
|---|---|---|---|
| `auth/state.yaml` | ログイン | idle / login_screen / loading / authenticated / error | source=ui の login_submitted |
| `order/state.yaml` | 注文 | cart / checkout_screen / processing / confirmed / failed | source=ui の checkout_submitted, source=external の payment_webhook_received（actor: stripe） |
| `inventory/state.yaml` | 在庫 | in_stock / low_stock / out_of_stock / restocking | source=er の inventory_changed（watches: catalog.store.inventory_db）|

`auth/state.yaml` と `order/state.yaml` の主要stateには `wireframe:` を持たせる。

### sequence scenarios

- `views/scenarios/checkout_flow.yaml`: state_file = `order/state.yaml`, steps = idle→checkout（source=uiパス検証）
- `views/scenarios/payment_webhook_flow.yaml`: state_file = `order/state.yaml`, steps = processing→payment_webhook_received（source=externalパス検証）

### api_table

`views/api_table.yaml`:
```yaml
as: api_table
id: ec_api
note: ECサイトAPI一覧
http_root_path: /api
modules:
  - module: auth
    include_submodules: false
  - module: catalog
    include_submodules: false
  - module: cart
    include_submodules: false
  - module: order
    include_submodules: false
  - module: payment
    include_submodules: false
```

---

## 4. 書く順序（依存順）

1. `actors.yaml`
2. 各moduleの `model/` 全部 ← これが最深の依存。先に全部書く
3. 各moduleの `store/` 全部
4. 各moduleの `task/`（mainファイル）
5. 各moduleの `state.yaml`
6. `views/api_table.yaml`
7. `views/scenarios/*.yaml`
8. `README.md`（最後）

**理由**: model→store→task→state→viewの順で参照方向が一方向。逆順だと未定義IDを書いて後で直す手間が発生する。

---

## 5. README.md の構成

```markdown
# UC-001: EC Checkout Flow

## 概要
（このUCで何を網羅検証するか）

## ファイル構成
（ツリー）

## カバレッジ
（§2のチェックリストの結果を「すべて✓」と書く）

## render例

### DAG例1: シンプル（auth.task.login）
（手書きMermaid）

### DAG例2: foreach（cart.task.validate_cart）
（手書きMermaid）

### DAG例3: fork+join（order.task.checkout）
（手書きMermaid）

### DAG例4: branch（order.task.process_order）
（手書きMermaid）

### State Diagram: auth
（手書きMermaid）

### State Diagram: order
（手書きMermaid）

### State Diagram: inventory
（手書きMermaid）

### ER Diagram: ec全体
（手書きMermaid）

### Sequence Diagram: checkout_flow
（手書きMermaid）

### Sequence Diagram: payment_webhook_flow
（手書きMermaid）

### API Table
（手書きMarkdown table）
```

合計10個のrender例。**各render結果は対応するview specのrenderルールに厳密に従うこと**（DAGはdag.md、ERはer.mdなど）。

---

## 6. ⚠️ 注意点（前セッションで議論済みのもの）

### main: は task のみ（前セッション最終確定）

`spec/nodes.md`に明記された通り、`main: true` は task ノードのみに付与。state.yaml内のstate/eventノードや、store/model/actorには `main:` フィールドを書かない。

### 1ファイル = 1FSM 厳守

`order/state.yaml` に在庫FSMを混ぜないこと。在庫FSMは `inventory/state.yaml` で別管理。state-diagram.md で「1ファイル=1FSM=1図」が明記されている。

### actorはglobal、モジュールパスなし

`event.actor` の値は `stripe`（フラットID）であって `payment.actor.stripe` ではない。ADR-031参照。

### sequence diagramの `state_file` パス記法

`views/scenarios/checkout_flow.yaml` の `state_file` は **yaml/ ルートからの相対パス** にする（例: `order/state.yaml`）。

### control flow scope（ADR-023）

branch / fork / foreach の内部assetは外部から直接参照不可。データを外に出す場合は store + reads/writes を使う。`order/task/checkout.yaml` の fork+join 設計時に注意。

### foreach の置き場

foreachはflow構文。`cart/task/validate_cart.yaml` 内で：
- main task: `validate_cart`
- sub task: `validate_item`（同ファイル内private）
- flow: で foreach: validate_item, over: $params.cart_items のように書く

### $params シジルは bareの暗黙解決禁止

`$params` 単独はNG。必ず `$params.field_name` の形（ADR-015）。

### task.returns は単一のみ

複数値を返したい場合はstruct modelでwrap（ADR-009）。

### context kind の store の使い所

`context` kindのstoreはどこかで1個使う（カバレッジ要件）。提案: `auth/store/` に `request_context_store` のようなものを置き、`auth.task.login` で reads させる。あるいは payment.task.process_payment で webhook context として使う。

### endpoint task の path はleaf path

`path: login` のように leaf のみ。フルパスは API Table view が組み立てる（ADR-028）。

### initializes は main task のみ

サブノードの task には `initializes` を書けない。

---

## 7. 不整合発見時の対応

書きながらspec/ADR間の不整合や曖昧な点を発見したら：

1. **書く手は止める**（推測で書き進めない）
2. README.md の「## 発見した不整合」セクションに列挙
3. ユーザーに判断を仰ぐ
4. 解消後、必要ならADR / specを更新してから書き続ける

特に注意して見るべき箇所：

- collection store の note 規約（complex query との境界）
- transition.action のクロスファイル参照のフルパス記法（`auth.task.login`形式）が ADR-027 のsentinel方式で正しくパースされるか
- inventory/state.yaml で er source event の `watches` が他モジュールのstoreを参照する際の記法（モジュールパス必要か？ → ADR-027準拠でフルパス `catalog.store.inventory_db` のはず）
- view定義ファイル（api_table.yaml / scenarios/*.yaml）に `nodes:` を書かないこと（ADR-030の「ノード定義 vs view定義」の区別）

---

## 8. commit戦略

ユーザーは bash で git が使えないので、Sonnetは適宜「ここでcommit推奨」とユーザーに提案する。提案タイミング目安：

1. actors.yaml + 全model 書き終わった時点
2. 全store + 全task 書き終わった時点
3. 全state.yaml 書き終わった時点
4. views 書き終わった時点
5. README 完成時点

各commitでADR Evidence更新は不要（このUCは新規ADRを起こさない想定）。万一新規ADR/spec更新が発生した場合のみ Evidence更新を依頼。

---

## 9. 完了条件

- [ ] §2 のカバレッジチェックリスト全項目 ✓
- [ ] §1 の全ファイル作成済み
- [ ] README.md の10個のrender例すべて記述済み
- [ ] 不整合があれば README.md に列挙済み
- [ ] TASKS.md の `[ ] UC-001` を `[x]` に更新

---

## 10. 完了後の片付け

このHANDOFF.md は完了後に削除してよい。残しておいても害はないが、UC本体（README.md + yaml/）が成果物。

---

## 付録: モデル一覧（参考、最終決定はSonnet判断）

参考までに想定model一覧。フィールド設計はSonnetが詳細化してよい。

| module | model | kind | 用途 |
|---|---|---|---|
| auth | credential | struct | username + password |
| auth | token | struct | access_token + expires_at |
| auth | login_form | struct | event payload |
| auth | login_log | struct | initializes example |
| auth | request_context | dict (value: str) | context store の `of:` に使用（context kind カバレッジ） |
| catalog | item | struct | id, name, price, stock |
| catalog | item_list | list (element: item) | get_items.returns に使用（list kind カバレッジ） |
| cart | cart | struct | id, user_id (fk: auth.model.credential.username), status |
| cart | cart_item | struct | id, cart_id (fk), item_id (fk: catalog.model.item.id), qty |
| order | order | struct | id, user_id (fk: auth.model.credential.username), shipping_address (type: address, fk なし=JSON埋め込み), status |
| order | order_item | struct | **N:M中間** order_id (fk) + item_id (fk: catalog.model.item.id) + qty |
| order | address | struct | JSON埋め込み専用（pk なし・fk なし）。order.shipping_address の型として使用 |
| payment | payment_event | struct | event payload |

> **注記（不整合解消済み）**
> - `ec.` モジュールプレフィックスは削除（ADR-027: フォルダ名がモジュール名）
> - address は JSON埋め込みに確定。1:1 FK の例は token.user_id → credential.username に振り替え
> - cart/order の user_id FK先は `auth.model.credential.username`（ADR-033: クロスモジュール FK フルパス必須）
> - login_log を auth/model/ に追加（§1 tree の漏れを修正）
> - list kind: `catalog/model/item_list.yaml`、dict kind: `auth/model/request_context.yaml`
> - クロスモジュール FK 記法は ADR-033 で確定

list / dict kindのmodelは上記テーブルの通り配置済み（list: catalog/model/item_list.yaml、dict: auth/model/request_context.yaml）。
