# payment_event

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/payment/model/payment_event.yaml |

### Fields

| field | type | note |
|---|---|---|
| event_id | str | Stripe側のevent ID（PK / 冪等性キー） |
| order_id | str | 対象注文。order モジュールへのクロスモジュールFK（ADR-033） |
| amount | float | 決済金額 |
| status | str | 決済結果（succeeded / failed） |
| received_at | datetime | webhook受信日時 |

