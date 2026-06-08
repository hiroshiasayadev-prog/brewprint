# process_order

在庫状況に応じて confirm（在庫充足）/ cancel（在庫不足）に分岐する。
returns なし（各分岐先 task が order_db を更新して終端、DAG的には END へfloating）。

```mermaid
flowchart TD
  subgraph params
    order_id([order_id: str])
  end

  _start([Start]) ==> check_inventory[check_inventory]
  order_id --> check_inventory
  order_db[(order_db)] -- "read" --> check_inventory
  inventory_db[(catalog.store.inventory_db)] -- "read" --> check_inventory
  check_inventory --> order_asset([order: order])

  check_inventory ==> route_by_inventory{route_by_inventory}
  order_asset --> route_by_inventory
  order_asset --> confirm_order[confirm_order]
  order_asset --> cancel_order[cancel_order]

  route_by_inventory == "in_stock" ==> confirm_order
  route_by_inventory == "out_of_stock" ==> cancel_order

  confirm_order -- "write" --> order_db
  cancel_order -- "write" --> order_db
  confirm_order ==> _end([End])
  cancel_order ==> _end

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
  classDef branchNode   fill:#9B6BBD,stroke:#6B3D8F,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class check_inventory,confirm_order,cancel_order taskNode
  class order_asset assetNode
  class order_db,inventory_db storeNode
  class route_by_inventory branchNode
  class _start,_end terminalNode
  class order_id boundaryNode
```

## Tasks

### process_order

#### Params

| name | model | note |
|---|---|---|
| order_id | str | 対象注文ID |

### check_inventory

order_id から order を引き当て、含まれる order_item の在庫を inventory_db で照合。
条件評価そのものは branch ノード側の note + cases.label が表現する。

#### Params

| name | model | note |
|---|---|---|
| order_id | str | — |

#### Returns

| name | model | source |
|---|---|---|
| order | order | — |

#### Store access

| access | store |
|---|---|
| read | order_db |
| read | catalog.store.inventory_db |

### route_by_inventory

order に紐づく全 order_item について inventory_db.stock >= order_item.qty なら in_stock、
1件でも不足があれば out_of_stock。

#### Params

| name | model | note |
|---|---|---|
| order | order | — |

### confirm_order

在庫充足パス。order_db の order.status を confirmed に更新して終端。
下流 wiring 不要（floating node → END、V01-ADR-023）。

#### Params

| name | model | note |
|---|---|---|
| order | order | — |

#### Store access

| access | store |
|---|---|
| write | order_db |

### cancel_order

在庫不足パス。order_db の order.status を cancelled に更新して終端。
下流 wiring 不要（floating node → END、V01-ADR-023）。

#### Params

| name | model | note |
|---|---|---|
| order | order | — |

#### Store access

| access | store |
|---|---|
| write | order_db |

