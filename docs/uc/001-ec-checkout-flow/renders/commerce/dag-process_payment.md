# process_payment

**API**: [POST /api/stripe](../_cross/api.md)

決済完了/失敗 webhook を受信し、対応する注文ステータスを更新する。
実装メモ:
- event.event_id は冪等性キー。同一 event_id の再送は idempotent に扱う
- event.order_id を主キーに order.store.order_db の order を引き当て
- event.status に応じて order.status を confirmed / failed に更新
- returns なし（webhookは200/空ボディ応答が一般的）

```mermaid
flowchart TD
  subgraph params
    event([event])
  end

  _start([Start]) ==> process_payment[process_payment]
  event --> process_payment

  process_payment -- "write" --> order.store.order_db[(order.store.order_db)]
  process_payment ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class process_payment taskNode
  class order.store.order_db storeNode
  class _start,_end terminalNode
  class event boundaryNode
```

## Tasks

### process_payment

#### Params

| name | model | note |
|---|---|---|
| event | payment.model.payment_event | Webhook payload。同 model は event.payload.model としても再利用される |

#### Store access

| access | store |
|---|---|
| write | order.store.order_db |
