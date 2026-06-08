# UC-001: EC Checkout Flow

## 概要

UC-001 は、ECサイトのチェックアウト処理を題材に、brewprint の主要ノード種別・flow構文・view定義・render規則を横断的に検証するユースケースである。

このUCは説明用サンプルではなく、Go実装のテストfixtureとしても流用する想定。

## ファイル構成

```text
docs/uc/001-ec-checkout-flow/
  README.md             ← このファイル
  TASKS-UC-001.md       ← UC-001固有の追跡タスク
  render_index.yaml     ← render出力のgroup定義
  yaml/                 ← brewprint YAML群（single source of truth）
    actors.yaml
    auth/
      state.yaml
      model/
      store/
      task/
    cart/
      model/
      store/
      task/
    catalog/
      model/
      store/
      task/
    inventory/
      state.yaml
    order/
      state.yaml
      model/
      store/
      task/
    payment/
      model/
      webhooks/
        task/
    views/
      api_table.yaml
      er.yaml
      scenarios/
        checkout_flow.yaml
        payment_webhook_flow.yaml
  renders/              ← Go renderer golden fixture（canonical render output）
    index.md            ← master render index
    auth/
      index.md
      dag-login.md
      state-auth.md
      wireframe-auth-login_screen.html
      wireframe-auth-loading.html
    commerce/
      index.md
      dag-add_to_cart.md
      dag-checkout.md
      dag-process_order.md
      dag-process_payment.md
      dag-validate_cart.md
      seq-checkout_flow.md
      seq-payment_webhook_flow.md
      state-order.md
      wireframe-order-cart.html
      wireframe-order-checkout_screen.html
    catalog/
      index.md
      dag-get_items.md
      state-inventory.md
    _cross/
      er.md
      api.md
    _preview/
      wireframe.html
  docs/                 ← 人間が書く補助ドキュメント / 旧render参照
    coverage.md
    render-dag.md
    render-state.md
    render-er.md
    render-seq.md
    render-api.md
    render-wireframe.md
```

## render fixture 再生成

`renders/` は Go renderer の canonical fixture として管理する。
手編集ではなく、原則として以下のコマンドで再生成する。

```powershell
brewprint render --yaml-root docs/uc/001-ec-checkout-flow/yaml --out docs/uc/001-ec-checkout-flow/renders --clean
```

再生成後は以下を確認する。

```powershell
go test ./...
```

## ドキュメント

| ファイル | 内容 |
|---|---|
| [renders/index.md](renders/index.md) | canonical render output の master index |
| [coverage.md](docs/coverage.md) | ノード種別・flow構文・task/model/state/viewフィールドのカバレッジ表 |
| [render-dag.md](docs/render-dag.md) | 旧DAG render参照。canonical fixture は `renders/` を正とする |
| [render-state.md](docs/render-state.md) | 旧State Diagram render参照。canonical fixture は `renders/` を正とする |
| [render-er.md](docs/render-er.md) | 旧ER Diagram render参照。canonical fixture は `renders/` を正とする |
| [render-seq.md](docs/render-seq.md) | 旧Sequence Diagram render参照。canonical fixture は `renders/` を正とする |
| [render-api.md](docs/render-api.md) | 旧API Table render参照。canonical fixture は `renders/` を正とする |
| [render-wireframe.md](docs/render-wireframe.md) | 旧Wireframe HTML fragment render参照。canonical fixture は `renders/` を正とする |

## TODO / spec gap

未解決の仕様差分・追跡タスクは `TASKS-UC-001.md` を参照。
