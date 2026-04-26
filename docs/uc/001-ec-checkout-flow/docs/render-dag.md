# UC-001: DAG render例

> 注記: このファイルは履歴用に残す旧手書きrender参照であり、canonicalではない。
> 正式なgolden fixtureは `../renders/` 配下を正とする。

## DAG例1: シンプル（auth.task.login）

入力:

- `yaml/auth/task/login.yaml`

```mermaid
flowchart TD
  subgraph params
    form([form])
  end
  subgraph returns
    auth_token([auth_token])
  end

  _start([Start]) ==> login[login]
  form --> login

  user_db[(user_db)] -- "read" --> login
  request_context_store[(request_context_store)] -- "read" --> login
  login -- "write" --> session_store[(session_store)]
  login -- "write" --> login_log_db[(login_log_db)]

  login --> auth_token
  login ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class login taskNode
  class user_db,request_context_store,session_store,login_log_db storeNode
  class _start,_end terminalNode
  class form,auth_token boundaryNode
```

## DAG例2: foreach（cart.task.validate_cart）

入力:

- `yaml/cart/task/validate_cart.yaml`

```mermaid
flowchart TD
  subgraph params
    cart_items([cart_items])
  end
  subgraph returns
    validated_items([validated_items])
  end

  _start([Start]) ==> validate_item["↻ validate_item"]
  cart_items --"foreach"--> validate_item

  item_collection[(catalog.store.item_collection)] -- "read" --> validate_item
  inventory_db[(catalog.store.inventory_db)] -- "read" --> validate_item

  validate_item --> validated_items
  validate_item ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class validate_item taskNode
  class item_collection,inventory_db storeNode
  class _start,_end terminalNode
  class cart_items,validated_items boundaryNode
```

## DAG例3: fork+join（order.task.checkout）

入力:

- `yaml/order/task/checkout.yaml`

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

## DAG例4: branch（order.task.process_order）

入力:

- `yaml/order/task/process_order.yaml`

```mermaid
flowchart TD
  subgraph params
    order_id([order_id])
  end

  _start([Start]) ==> check_inventory[check_inventory]
  order_id --> check_inventory
  order_db[(order_db)] -- "read" --> check_inventory
  inventory_db[(catalog.store.inventory_db)] -- "read" --> check_inventory
  check_inventory --> order_asset([order])

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
