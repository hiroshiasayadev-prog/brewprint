# checkout

**API**: [POST /api/checkout](../_cross/api.md)

カート確定 → 注文レコード作成 → 在庫確保 + 決済要求送信 を並列実行し、
pending状態の注文（決済webhook待ち）を返す。

```mermaid
flowchart TD
  subgraph params
    cart_id([cart_id])
    shipping_address([shipping_address])
  end
  subgraph returns
    pending_order([pending_order])
  end

  _start([Start]) ==> build_order[build_order]
  cart_id --> build_order
  shipping_address --> build_order
  cart_session[(cart.store.cart_session)] -- "read" --> build_order
  user_db[(auth.store.user_db)] -- "read" --> build_order
  build_order -- "write" --> order_db[(order_db)]
  build_order --> draft_order([draft_order])

  build_order ==> parallel_processing{{parallel_processing}}
  draft_order --> reserve_inventory[reserve_inventory]
  draft_order --> notify_payment_gateway[notify_payment_gateway]

  parallel_processing == "parallel" ==> reserve_inventory
  parallel_processing == "parallel" ==> notify_payment_gateway

  reserve_inventory <-- "read/write" --> inventory_db[(catalog.store.inventory_db)]
  reserve_inventory --> reserved([reserved])
  notify_payment_gateway --> notified([notified])

  reserve_inventory ==> finalize_checkout{{finalize_checkout}}
  notify_payment_gateway ==> finalize_checkout
  reserved --> finalize_checkout
  notified --> finalize_checkout

  finalize_checkout --> pending_order
  finalize_checkout ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
  classDef forkNode     fill:#8A8A8A,stroke:#5A5A5A,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class build_order,reserve_inventory,notify_payment_gateway taskNode
  class draft_order,reserved,notified assetNode
  class cart_session,user_db,order_db,inventory_db storeNode
  class parallel_processing,finalize_checkout forkNode
  class _start,_end terminalNode
  class cart_id,shipping_address,pending_order boundaryNode
```

## Tasks

### checkout

#### Params

| name | model | note |
|---|---|---|
| cart_id | str | 対象カートID |
| shipping_address | address | 配送先（JSON埋め込み model、ADR-021） |

#### Returns

| name | model |
|---|---|
| pending_order | order |

### build_order

cart_session から cart を取得し、user_db でユーザー存在を確認した上で
draft order（status: pending）を order_db に新規作成して返す。

#### Params

| name | model | note |
|---|---|---|
| cart_id | str | — |
| shipping_address | address | — |

#### Returns

| name | model |
|---|---|
| draft_order | order |

#### Store access

| access | store |
|---|---|
| read | cart.store.cart_session |
| read | auth.store.user_db |
| write | order_db |

### reserve_inventory

draft_order の order_item から item_id / qty を読み出し、
catalog.store.inventory_db の stock を減算して在庫確保。
読み込みと書き込みは同一トランザクション境界（ADR-008/020）。

#### Params

| name | model | note |
|---|---|---|
| draft_order | order | — |

#### Returns

| name | model |
|---|---|
| reserved | order |

#### Store access

| access | store |
|---|---|
| read/write | catalog.store.inventory_db |

### notify_payment_gateway

Stripe API に決済要求を送信（外部HTTP呼び出し）。
決済の最終結果は payment.webhooks.task.process_payment が webhook で受信。
ここでは送信完了のみを保証して draft_order を passthrough で返す。

#### Params

| name | model | note |
|---|---|---|
| draft_order | order | — |

#### Returns

| name | model |
|---|---|
| notified | order |

### parallel_processing

在庫確保と決済要求送信を並列実行する分岐点。

#### Params

| name | model | note |
|---|---|---|
| draft_order | order | — |

### finalize_checkout

両ブランチ完了を待ち合わせ、結果統合済みの pending_order を返す。
在庫確保済み + 決済GW通知済み = pending（webhookで confirmed/failed に確定）。

#### Params

| name | model | note |
|---|---|---|
| reserved | order | — |
| notified | order | — |

#### Returns

| name | model |
|---|---|
| pending_order | order |
