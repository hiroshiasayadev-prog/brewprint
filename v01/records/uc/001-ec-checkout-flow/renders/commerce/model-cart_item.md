# cart_item

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/cart/model/cart_item.yaml |

### Fields

| field | type | note |
|---|---|---|
| id | str | 中間テーブルのsurrogate key（PK / V01-ADR-026） |
| cart_id | str | 所属カート。同モジュール内FK（bare ID参照） |
| item_id | str | 商品参照。catalog モジュールへのクロスモジュールFK（V01-ADR-033） |
| qty | int | 数量 |

