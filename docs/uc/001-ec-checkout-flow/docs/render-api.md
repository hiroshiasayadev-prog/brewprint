# UC-001: API Table render例

## API Table

入力:

- `yaml/views/api_table.yaml`
- endpoint task群

## Routes

- `/api/login`
- `/api/catalog_items`
- `/api/cart_items`
- `/api/checkout`
- `/api/stripe`

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
| stripe | POST | /api/stripe | payment_event | - |
