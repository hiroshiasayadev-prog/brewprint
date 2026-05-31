# add_to_cart

**API**: [POST /api/cart_items](../_cross/api.md)

指定商品をカートに追加する。
実装メモ:
- catalog.store.item_collection で item_id の存在を確認（find_by_id）
- cart_session から既存cartを取得し、cart_item を append または qty を加算
- 同一トランザクションで cart_session を更新

```mermaid
flowchart TD
  subgraph params
    cart_id([cart_id])
    item_id([item_id])
    qty([qty])
  end

  _start([Start]) ==> add_to_cart[add_to_cart]
  cart_id --> add_to_cart
  item_id --> add_to_cart
  qty --> add_to_cart

  catalog.store.item_collection[(catalog.store.item_collection)] -- "read" --> add_to_cart
  add_to_cart -- "write" --> cart_session[(cart_session)]

  add_to_cart --> updated_cart([updated_cart])
  add_to_cart ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class add_to_cart taskNode
  class catalog.store.item_collection,cart_session storeNode
  class updated_cart assetNode
  class _start,_end terminalNode
  class cart_id,item_id,qty boundaryNode
```

## Tasks

### add_to_cart

#### Params

| name | model | note |
|---|---|---|
| cart_id | str | 対象カートID（primitive str） |
| item_id | str | 追加する商品ID（catalog.model.item.id へのID参照） |
| qty | int | 数量（primitive int） |

#### Returns

| name | model | source |
|---|---|---|
| updated_cart | cart | — |

#### Store access

| access | store |
|---|---|
| read | catalog.store.item_collection |
| write | cart_session |
