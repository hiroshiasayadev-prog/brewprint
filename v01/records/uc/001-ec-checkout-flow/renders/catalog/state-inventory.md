# inventory

```mermaid
stateDiagram-v2
  state _choice_in_stock_inventory_changed <<choice>>
  state _choice_low_stock_inventory_changed <<choice>>
  state _choice_out_of_stock_inventory_changed <<choice>>
  state _choice_restocking_inventory_changed <<choice>>

  [*] --> in_stock

  in_stock --> _choice_in_stock_inventory_changed : inventory_changed
  _choice_in_stock_inventory_changed --> low_stock : [stock > 0 AND stock < threshold]
  _choice_in_stock_inventory_changed --> out_of_stock : [stock == 0]

  low_stock --> _choice_low_stock_inventory_changed : inventory_changed
  _choice_low_stock_inventory_changed --> in_stock : [stock >= threshold]
  _choice_low_stock_inventory_changed --> out_of_stock : [stock == 0]

  out_of_stock --> _choice_out_of_stock_inventory_changed : inventory_changed
  _choice_out_of_stock_inventory_changed --> restocking : [restock_in_progress == true]
  _choice_out_of_stock_inventory_changed --> in_stock : [stock >= threshold]

  restocking --> _choice_restocking_inventory_changed : inventory_changed
  _choice_restocking_inventory_changed --> in_stock : [stock >= threshold]
  _choice_restocking_inventory_changed --> low_stock : [stock > 0 AND stock < threshold]
```

## States

| state | note |
|---|---|
| in_stock | 在庫充足（閾値 threshold 以上） |
| low_stock | 在庫少量（0 < stock < threshold） |
| out_of_stock | 在庫切れ（stock == 0） |
| restocking | 補充作業中。バックヤードの restock 処理が inventory_db に書き込むことで発火・遷移する想定。restock_in_progress フラグ（inventory_db レコード内）で判定。 |

## Events

| event | source | actor | note |
|---|---|---|---|
| inventory_changed | er | — | catalog.store.inventory_db（商品在庫DB）の変化を監視するER起点event（V01-ADR-018）。クロスモジュールstore参照のためフルパス記法（V01-ADR-027）。payloadは省略（どのレコードが変化したかは watches 側の監視機構で解決される前提）。 |

