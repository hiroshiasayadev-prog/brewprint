# UC-001: API Table render例

## API Table

入力:

- `yaml/views/api_table.yaml`
- endpoint task群

### Render結果

```markdown
# ec_api

ECサイトAPI一覧

## Routes

- [auth](#auth)
- [catalog](#catalog)
- [cart](#cart)
- [order](#order)
- [payment.webhooks](#paymentwebhooks)

## auth

| task id | method | path | params | returns |
|---|---|---|---|---|
| login | POST | /api/login | login_form | token |

## catalog

| task id | method | path | params | returns |
|---|---|---|---|---|
| get_items | GET | /api/catalog_items | - | item_list |

## cart

| task id | method | path | params | returns |
|---|---|---|---|---|
| add_to_cart | POST | /api/cart_items | str<br/>str<br/>int | cart |

## order

| task id | method | path | params | returns |
|---|---|---|---|---|
| checkout | POST | /api/checkout | str<br/>address | order |

## payment.webhooks

| task id | method | path | params | returns |
|---|---|---|---|---|
| process_payment | POST | /api/stripe | payment_event | - |
```
