# order

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/order/model/order.yaml |

### Fields

| field | type | note |
|---|---|---|
| id | str | 注文ID（PK）。order_item.order_id / payment_event.order_id のFK参照先 |
| user_id | str | 注文者。auth モジュールへのクロスモジュールFK（ADR-033） |
| shipping_address | address | 配送先住所。fk なし → DBではJSONカラム埋め込み（ADR-021） |
| total_price | float | 注文合計金額 |
| status | str | 注文状態（pending / processing / confirmed / failed） |
| created_at | datetime | 注文日時 |

