# 商取引 render index

| type | render | source |
|---|---|---|
| DAG | [validate_cart](dag-validate_cart.md) | `yaml/cart/task/validate_cart.yaml` |
| DAG | [checkout](dag-checkout.md) | `yaml/order/task/checkout.yaml` |
| DAG | [process_order](dag-process_order.md) | `yaml/order/task/process_order.yaml` |
| State | [order](state-order.md) | `yaml/order/state.yaml` |
| Sequence | [checkout_flow](seq-checkout_flow.md) | `yaml/views/scenarios/checkout_flow.yaml` |
| Sequence | [payment_webhook_flow](seq-payment_webhook_flow.md) | `yaml/views/scenarios/payment_webhook_flow.yaml` |
| Wireframe | [cart](wireframe-order-cart.html) | `yaml/order/state.yaml#cart` |
| Wireframe | [checkout_screen](wireframe-order-checkout_screen.html) | `yaml/order/state.yaml#checkout_screen` |
