# UC-001: EC Checkout Flow

## 概要

UC-001 は、ECサイトのチェックアウト処理を題材に、brewprint の主要ノード種別・flow構文・view定義・render規則を横断的に検証するユースケースである。

このUCは説明用サンプルではなく、Go実装のテストfixtureとしても流用する想定。

## ファイル構成

```text
docs/uc/001-ec-checkout-flow/
  README.md           ← このファイル
  HANDOFF.md
  TASKS-UC-001.md
  yaml/
    actors.yaml
    auth/
      model/
        credential.yaml
        login_form.yaml
        login_log.yaml
        request_context.yaml
        token.yaml
      state.yaml
      store/
        request_context_store.yaml
        session_store.yaml
        user_db.yaml
      task/
        login.yaml
    cart/
      model/
        cart.yaml
        cart_item.yaml
        cart_item_list.yaml
      store/
        cart_session.yaml
      task/
        add_to_cart.yaml
        validate_cart.yaml
    catalog/
      model/
        item.yaml
        item_list.yaml
      store/
        inventory_db.yaml
        item_collection.yaml
      task/
        get_items.yaml
    inventory/
      state.yaml
    order/
      model/
        address.yaml
        order.yaml
        order_item.yaml
      state.yaml
      store/
        order_db.yaml
      task/
        checkout.yaml
        process_order.yaml
    payment/
      model/
        payment_event.yaml
      task/
      webhooks/
        task/
          process_payment.yaml
    views/
      api_table.yaml
      er.yaml
      scenarios/
        checkout_flow.yaml
        payment_webhook_flow.yaml
  docs/
    coverage.md         ← カバレッジ表（ノード種別・flow構文・task/model/state/view）
    render-dag.md       ← DAG render例 × 4
    render-state.md     ← State Diagram render例 × 3
    render-er.md        ← ER Diagram render例
    render-seq.md       ← Sequence Diagram render例 × 2
    render-api.md       ← API Table render例
```

## ドキュメント

| ファイル | 内容 |
|---|---|
| [coverage.md](docs/coverage.md) | ノード種別・flow構文・task/model/state/viewフィールドのカバレッジ表 |
| [render-dag.md](docs/render-dag.md) | DAG render期待値（シンプル / foreach / fork+join / branch） |
| [render-state.md](docs/render-state.md) | State Diagram render期待値（auth / order / inventory） |
| [render-er.md](docs/render-er.md) | ER Diagram render期待値（ec全体） |
| [render-seq.md](docs/render-seq.md) | Sequence Diagram render期待値（checkout_flow / payment_webhook_flow） |
| [render-api.md](docs/render-api.md) | API Table render期待値 |

## TODO / spec gap

未解決の仕様差分・追跡タスクは `TASKS-UC-001.md` を参照。
