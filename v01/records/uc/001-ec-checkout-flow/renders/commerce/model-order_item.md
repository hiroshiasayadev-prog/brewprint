# order_item

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/order/model/order_item.yaml |

### Fields

| field | type | note |
|---|---|---|
| id | str | 中間テーブルのsurrogate key（PK / V01-ADR-026） |
| order_id | str | 所属注文。同モジュール内FK（bare ID参照） |
| item_id | str | 商品参照。catalog モジュールへのクロスモジュールFK（V01-ADR-033） |
| qty | int | 数量 |
| price_snapshot | float | 注文時点の単価（後日価格が変わってもこの値は不変） |

