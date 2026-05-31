# validate_cart

カート内アイテムを 1 件ずつ validate_item で検証し、
検証通過したアイテムのリストを返す。
mode: sequential（1件ずつ順に処理。後続でロールバック単位を揃えるため）

```mermaid
flowchart TD
  subgraph params
    cart_items([cart_items])
  end

  _start([Start]) ==> validate_item["↻ validate_item"]
  cart_items --"foreach"--> validate_item

  item_collection[(catalog.store.item_collection)] -- "read" --> validate_item
  inventory_db[(catalog.store.inventory_db)] -- "read" --> validate_item

  validate_item --> validated_items
  validate_item ==> _end([End])
  validated_items -- "returns as validated_items" --> _end

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class validate_item taskNode
  class validated_items assetNode
  class item_collection,inventory_db storeNode
  class _start,_end terminalNode
  class cart_items boundaryNode
```

## Tasks

### validate_cart

#### Params

| name | model | note |
|---|---|---|
| cart_items | cart_item_list | バリデート対象のカートアイテム一覧（list kind） |

#### Returns

| name | model | source |
|---|---|---|
| validated_items | cart_item_list | validated_items |

### validate_item

1件の cart_item について商品存在確認と在庫確認を行う。
- catalog.store.item_collection.find_by_id で cart_item.item_id の商品存在を確認
- catalog.store.inventory_db で item.stock >= cart_item.qty を確認
失敗時は実装で例外を上げる（YAML上は成功パスのみ表現）。

#### Params

| name | model | note |
|---|---|---|
| cart_item | cart_item | 1件のカートアイテム |

#### Returns

| name | model | source |
|---|---|---|
| validated | cart_item | — |

#### Store access

| access | store |
|---|---|
| read | catalog.store.item_collection |
| read | catalog.store.inventory_db |

