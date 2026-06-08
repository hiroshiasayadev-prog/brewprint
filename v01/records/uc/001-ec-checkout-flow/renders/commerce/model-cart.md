# cart

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/cart/model/cart.yaml |

### Fields

| field | type | note |
|---|---|---|
| id | str | カートID（PK） |
| user_id | str | 所有ユーザー。auth モジュールへのクロスモジュールFK（V01-ADR-033） |
| status | str | カート状態（active / locked / checked_out） |
| updated_at | datetime | 最終更新日時 |

