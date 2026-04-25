# UC-001: State Diagram render例

## State Diagram: auth

入力:

- `yaml/auth/state.yaml`

```mermaid
stateDiagram-v2
  [*] --> idle
  authenticated --> [*]
  error --> [*]

  idle --> login_screen : enter_login
  error --> login_screen : enter_login
  login_screen --> loading : login_submitted [form.valid == true] / auth.task.login
  loading --> authenticated : login_succeeded
  loading --> error : login_failed
```

## State Diagram: order

入力:

- `yaml/order/state.yaml`

```mermaid
stateDiagram-v2
  state _choice_processing_payment_webhook_received <<choice>>

  [*] --> cart
  confirmed --> [*]
  failed --> [*]

  cart --> checkout_screen : view_checkout
  failed --> cart : view_cart
  checkout_screen --> processing : checkout_submitted / order.task.checkout
  processing --> _choice_processing_payment_webhook_received : payment_webhook_received
  _choice_processing_payment_webhook_received --> confirmed : [payload.status == 'succeeded'] / payment.webhooks.task.process_payment
  _choice_processing_payment_webhook_received --> failed : [payload.status == 'failed']
```

## State Diagram: inventory

入力:

- `yaml/inventory/state.yaml`

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
