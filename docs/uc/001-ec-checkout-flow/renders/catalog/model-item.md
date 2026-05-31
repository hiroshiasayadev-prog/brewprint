# item

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/catalog/model/item.yaml |

### Fields

| field | type | note |
|---|---|---|
| id | str | 商品ID（PK）。cart_item.item_id / order_item.item_id のクロスモジュールFK参照先 |
| name | str | 商品名 |
| price | float | 単価（float primitive の例） |
| stock | int | 在庫数（int primitive の例） |
| is_available | bool | 販売可能フラグ（bool primitive の例） |
| created_at | datetime | 登録日時（datetime primitive の例） |

