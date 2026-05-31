# チェックアウトフロー

```mermaid
sequenceDiagram
  participant UI
  participant API as order.task.checkout
  participant DB

  UI->>UI: 1. view_checkout
  UI->>API: 2. POST /checkout
  API->>DB: 2. reads
  API->>DB: 2. writes
  API-->>UI: 2. pending_order
```

## DB操作

| step | task | sub_task | store | 操作 |
|---|---|---|---|---|
| 2 | order.task.checkout | build_order | order_db | writes |
| 2 | order.task.checkout | build_order | auth.store.user_db | reads |
| 2 | order.task.checkout | reserve_inventory | catalog.store.inventory_db | reads |
| 2 | order.task.checkout | reserve_inventory | catalog.store.inventory_db | writes |

