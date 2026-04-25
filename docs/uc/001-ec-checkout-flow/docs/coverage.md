# UC-001: カバレッジ

## ノード種別

| 種別 | 使用箇所 | カバー内容 |
|---|---|---|
| `actor` | `yaml/actors.yaml` | global actor 定義。`end_user` / `stripe` を定義し、`event.source=external` の `actor: stripe` から参照する。 |
| `model` / `kind: struct` | `yaml/auth/model/credential.yaml`, `yaml/catalog/model/item.yaml`, `yaml/order/model/order.yaml`, `yaml/payment/model/payment_event.yaml` など | PK / FK / unique / primitive field / JSON埋め込み用model参照をカバーする。 |
| `model` / `kind: list` | `yaml/cart/model/cart_item_list.yaml`, `yaml/catalog/model/item_list.yaml` | `element:` によるlist型定義をカバーする。 |
| `model` / `kind: dict` | `yaml/auth/model/request_context.yaml` | `value:` のみを持つdict型定義をカバーする。 |
| `store` / `kind: db` | `yaml/auth/store/user_db.yaml`, `yaml/catalog/store/inventory_db.yaml`, `yaml/order/store/order_db.yaml` | ER図対象のDB storeをカバーする。 |
| `store` / `kind: session` | `yaml/auth/store/session_store.yaml`, `yaml/cart/store/cart_session.yaml` | セッション保持storeをカバーする。 |
| `store` / `kind: collection` | `yaml/catalog/store/item_collection.yaml` | noteによる検索semantic contractをカバーする。 |
| `store` / `kind: context` | `yaml/auth/store/request_context_store.yaml` | リクエストスコープのcontext storeをカバーする。 |
| `task` / main task | `yaml/auth/task/login.yaml`, `yaml/order/task/checkout.yaml` など | `main: true`、endpointあり/なし、params/returns、reads/writesをカバーする。 |
| `task` / sub task | `yaml/cart/task/validate_cart.yaml`, `yaml/order/task/checkout.yaml`, `yaml/order/task/process_order.yaml` | 1ファイル内private sub taskをカバーする。 |
| `branch` | `yaml/order/task/process_order.yaml` | 在庫状況に応じた排他分岐をカバーする。 |
| `fork` / `join` | `yaml/order/task/checkout.yaml` | 並列分岐と合流をカバーする。 |
| `state` | `yaml/auth/state.yaml`, `yaml/order/state.yaml`, `yaml/inventory/state.yaml` | initial / final / guard付き遷移 / action付き遷移 / actionなし遷移をカバーする。 |
| `event` | `yaml/auth/state.yaml`, `yaml/order/state.yaml`, `yaml/inventory/state.yaml` | `source: ui` / `internal` / `external` / `er` をカバーする。 |
| implicit `asset` | 各 `task.returns` | 独立ファイルを持たず、taskの `returns` からDAG上のassetとして導出される前提をカバーする。 |
| view YAML | `yaml/views/api_table.yaml`, `yaml/views/er.yaml`, `yaml/views/scenarios/*.yaml` | `as: api_table` / `er_diagram` / `sequence_diagram` をカバーする。 |

## flow構文

| 構文 | 使用箇所 | カバー内容 |
|---|---|---|
| `step` | `yaml/order/task/checkout.yaml`, `yaml/order/task/process_order.yaml` | 通常stepと `params:` wiringをカバーする。 |
| `$params.field` | `yaml/order/task/checkout.yaml`, `yaml/order/task/process_order.yaml`, `yaml/cart/task/validate_cart.yaml` | main task paramsからflowへ入力を注入する境界シジルをカバーする。 |
| `foreach` | `yaml/cart/task/validate_cart.yaml` | `foreach: validate_item`, `mode: sequential`, `params: { cart_item: $item }`, `returns:` をカバーする。 |
| `$item` | `yaml/cart/task/validate_cart.yaml` | foreachの現在要素をapply先taskへ渡す記法をカバーする。 |
| `foreach.over` | `yaml/cart/task/validate_cart.yaml` | `over: $params.cart_items` を使用。`node_id` と `$params.field` の両方を許可する現行 `edges.md` に追随済み。 |
| `fork` / `branches` / `join` | `yaml/order/task/checkout.yaml` | ADR-040 の `branches[].steps[].step` + `steps[].params` 形式と `join: finalize_checkout` をカバーする。 |
| `fork.branches[].steps[].params` | `yaml/order/task/checkout.yaml` | fork内taskへのparams wiringを、暗黙伝播ではなく各stepの `params:` で明示する形式をカバーする。 |
| `branch` / `cases` | `yaml/order/task/process_order.yaml` | `cases[].label` / `cases[].step` / `cases[].params` による排他分岐をカバーする。 |
| `branch.params` / `cases[].params` | `yaml/order/task/process_order.yaml` | `branch.params` はbranch node自身の判定入力、`cases[].params` はcase先taskへのwiringとして分離するADR-040形式をカバーする。 |

## task内フィールド

| フィールド | 使用箇所 | カバー内容 |
|---|---|---|
| `id` / `type` | 全task YAML | task nodeの基本識別子をカバーする。 |
| `main` | `yaml/*/task/*.yaml` | ファイル代表taskの明示をカバーする。 |
| `endpoint` | `yaml/auth/task/login.yaml`, `yaml/catalog/task/get_items.yaml`, `yaml/cart/task/add_to_cart.yaml`, `yaml/order/task/checkout.yaml`, `yaml/payment/webhooks/task/process_payment.yaml` | API Table対象taskと内部専用taskの差分をカバーする。 |
| `method` / `path` | endpoint task群 | leaf path方式をカバーする。`path: catalog_items`, `path: cart_items`, `path: checkout`, `path: stripe` など。 |
| `params` | `yaml/auth/task/login.yaml`, `yaml/cart/task/add_to_cart.yaml`, `yaml/order/task/checkout.yaml` など | model参照とprimitive参照（`str` / `int`）の両方をカバーする。 |
| `returns` | `yaml/auth/task/login.yaml`, `yaml/catalog/task/get_items.yaml`, `yaml/order/task/checkout.yaml` など | 単一return asset導出をカバーする。returnsなしtaskも `process_order` / `process_payment` でカバーする。 |
| `reads` | `yaml/auth/task/login.yaml`, `yaml/cart/task/validate_cart.yaml`, `yaml/order/task/checkout.yaml` など | 同モジュールstore参照とクロスモジュールstore参照をカバーする。 |
| `writes` | `yaml/auth/task/login.yaml`, `yaml/cart/task/add_to_cart.yaml`, `yaml/order/task/checkout.yaml`, `yaml/payment/webhooks/task/process_payment.yaml` | 同モジュールstore更新とクロスモジュールstore更新をカバーする。 |
| `initializes` | `yaml/auth/task/login.yaml` | ファイル内private storeの初期化をカバーする。 |
| `note` | task / branch / fork / join 各所 | 機械検証外のsemantic contractをカバーする。 |
| `flow` | `yaml/cart/task/validate_cart.yaml`, `yaml/order/task/checkout.yaml`, `yaml/order/task/process_order.yaml` | task定義とwiring定義の分離をカバーする。 |

## model内フィールド

| フィールド / パターン | 使用箇所 | カバー内容 |
|---|---|---|
| `kind: struct` + `fields` | `yaml/auth/model/credential.yaml`, `yaml/catalog/model/item.yaml`, `yaml/order/model/order.yaml` など | 構造体modelをカバーする。 |
| `fields[].name` / `type` / `note` | struct model全般 | field基本構造をカバーする。 |
| `pk: true` | `credential.username`, `item.id`, `order.id`, `payment_event.event_id` など | 単一PKをカバーする。 |
| `fk:` 同モジュール | `auth/model/token.yaml`, `cart/model/cart_item.yaml`, `order/model/order_item.yaml` | `credential.username`, `cart.id`, `order.id` などの同モジュールFKをカバーする。 |
| `fk:` クロスモジュール | `cart/model/cart.yaml`, `cart/model/cart_item.yaml`, `order/model/order.yaml`, `payment/model/payment_event.yaml` | `auth.model.credential.username`, `catalog.model.item.id`, `order.model.order.id` などをカバーする。 |
| `unique: true` | `yaml/auth/model/token.yaml` | FKの1:1表現をカバーする。 |
| N:M 中間model | `yaml/cart/model/cart_item.yaml`, `yaml/order/model/order_item.yaml` | surrogate key + 2本FKのパターンをカバーする。 |
| JSON埋め込み | `yaml/order/model/order.yaml` → `shipping_address: address` | `fk` なしのmodel参照をJSON埋め込みとして扱うパターンをカバーする。 |
| `kind: list` + `element` | `yaml/cart/model/cart_item_list.yaml`, `yaml/catalog/model/item_list.yaml` | list modelをカバーする。 |
| `kind: dict` + `value` | `yaml/auth/model/request_context.yaml` | dict modelをカバーする。 |
| primitive予約語 | `str`, `int`, `float`, `bool`, `datetime` | UC-001内で使用済み。`bytes` / `any` は未使用。 |

## state / event / wireframe

| 項目 | 使用箇所 | カバー内容 |
|---|---|---|
| `state.initial` | `auth.state: idle`, `order.state: cart`, `inventory.state: in_stock` | 初期状態をカバーする。 |
| `state.final` | `auth.state: authenticated/error`, `order.state: confirmed/failed` | 終端状態をカバーする。 |
| `transitions` | `yaml/auth/state.yaml`, `yaml/order/state.yaml`, `yaml/inventory/state.yaml` | FSM遷移定義をカバーする。 |
| `transition.action` | `auth.state -> auth.task.login`, `order.state -> order.task.checkout`, `order.state -> payment.webhooks.task.process_payment` | Application → Processing の参照をカバーする。 |
| actionなしtransition | `order.state: cart -> checkout_screen`, `auth.state: idle -> login_screen` | sequence diagram上のUI self-message対象をカバーする。 |
| `guard` | `auth/state.yaml`, `order/state.yaml`, `inventory/state.yaml` | 単独guardと同一 `(from, on)` のguard分岐をカバーする。 |
| `event.source: ui` | `auth/state.yaml`, `order/state.yaml` | UI起点eventをカバーする。 |
| `event.source: internal` | `auth/state.yaml` | task完了起点eventをカバーする。 |
| `event.source: external` + `actor` | `order/state.yaml` | Stripe webhook eventをカバーする。 |
| `event.source: er` + `watches` | `inventory/state.yaml` | DB変化監視eventをカバーする。 |
| `event.payload.model` | `auth/state.yaml`, `order/state.yaml` | event payloadのmodel参照をカバーする。 |
| wireframe layout | `auth/state.yaml`, `order/state.yaml` | `col` / `row` / `grid` / `header` / `footer` / `card` をカバーする。`sidebar` は未使用。 |
| wireframe component | `auth/state.yaml`, `order/state.yaml` | `text` / `button` / `input` / `password` / `select` / `checkbox` / `image` / `icon` / `badge` / `divider` をカバーする。`radio` は未使用。 |
| wireframe interactive field | `auth/state.yaml`, `order/state.yaml` | `fires` / `disabled` / `placeholder` をカバーする。 |
| wireframe grid field | `order/state.yaml` | `cols` / `span` をカバーする。 |

## view定義

| view | 使用箇所 | カバー内容 |
|---|---|---|
| API Table | `yaml/views/api_table.yaml` | `as: api_table`, `id`, `note`, `http_root_path`, `modules[].module`, `include_submodules` をカバーする。 |
| ER Diagram | `yaml/views/er.yaml` | `as: er_diagram`, 横断view、複数module集約、クロスモジュールFK描画対象をカバーする。 |
| Sequence Diagram | `yaml/views/scenarios/checkout_flow.yaml` | `as: sequence_diagram`, `state_file`, `steps`, actionなしtransition、UI起点API呼び出しをカバーする。 |
| Sequence Diagram + guard | `yaml/views/scenarios/payment_webhook_flow.yaml` | `guard:` によるtransition一意特定、external actor起点、returnsなしAPI応答をカバーする。 |
