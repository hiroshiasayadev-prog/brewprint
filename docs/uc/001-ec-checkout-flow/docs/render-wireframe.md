# UC-001 wireframe render期待値

## スコープ

このdocは `docs/uc/001-ec-checkout-flow` の既存state wireframeを対象に、HTML fragment renderの期待値を定義する。

対象state:

- `auth.login_screen`
- `auth.loading`
- `order.cart`
- `order.checkout_screen`

入力YAML:

- `docs/uc/001-ec-checkout-flow/yaml/auth/state.yaml`
- `docs/uc/001-ec-checkout-flow/yaml/order/state.yaml`

準拠仕様:

- `docs/spec/views/wireframe.md`
- `docs/adr/042-wireframe-main-and-layout.md`

## render共通ルール

- HTML fragmentのみを出力する。`DOCTYPE` / `html` / `head` / `body` は出力しない。
- YAMLの `id` はHTMLの `id` 属性ではなく `data-wf-id` に出力する。
- YAMLの `fires` は `data-wf-fires` に出力する。
- `label` / `placeholder` / `id` / `fires` はHTML escapeする。
- `layout` は決定的なinline styleに変換する。
- JSは生成しない。
- 任意CSSは受け付けない。

---

## auth.login_screen

### 入力YAML

```yaml
# docs/uc/001-ec-checkout-flow/yaml/auth/state.yaml
- id: login_screen
  type: state
  wireframe:
    type: col
    layout:
      width: fill
      min_height: 480
    children:
      - type: main
        id: login_main
        layout:
          grow: true
          padding: 24
          gap: 16
          align: stretch
          justify: center
        children:
          - type: card
            id: login_card
            layout:
              width: 360
              gap: 12
              padding: 16
            children:
              - type: text
                id: login_title
                label: "ログイン"
              - type: input
                id: username_input
                label: "ユーザー名"
                placeholder: "username"
              - type: password
                id: password_input
                label: "パスワード"
                placeholder: "パスワードを入力"
              - type: button
                id: submit_button
                label: "ログイン"
                fires: login_submitted
```

### HTML出力期待値

```html
<div class="wf-col" style="width: 100%; min-height: 480px;">
  <main class="wf-main" data-wf-id="login_main" style="flex: 1 1 0%; min-width: 0; min-height: 0; gap: 16px; padding: 24px; align-items: stretch; justify-content: center;">
    <section class="wf-card" data-wf-id="login_card" style="width: 360px; gap: 12px; padding: 16px;">
      <span class="wf-text" data-wf-id="login_title">ログイン</span>
      <div class="wf-field" data-wf-id="username_input">
        <label>ユーザー名</label>
        <input type="text" placeholder="username" />
      </div>
      <div class="wf-field" data-wf-id="password_input">
        <label>パスワード</label>
        <input type="password" placeholder="パスワードを入力" />
      </div>
      <button class="wf-button" data-wf-id="submit_button" data-wf-fires="login_submitted">ログイン</button>
    </section>
  </main>
</div>
```

---

## auth.loading

### 入力YAML

```yaml
# docs/uc/001-ec-checkout-flow/yaml/auth/state.yaml
- id: loading
  type: state
  wireframe:
    type: col
    layout:
      width: fill
      min_height: 320
    children:
      - type: main
        id: loading_main
        layout:
          grow: true
          padding: 24
          gap: 12
          align: center
          justify: center
        children:
          - type: icon
            id: loading_spinner
          - type: text
            id: loading_text
            label: "認証中..."
          - type: button
            id: cancel_button
            label: "キャンセル"
            disabled: true
```

### HTML出力期待値

```html
<div class="wf-col" style="width: 100%; min-height: 320px;">
  <main class="wf-main" data-wf-id="loading_main" style="flex: 1 1 0%; min-width: 0; min-height: 0; gap: 12px; padding: 24px; align-items: center; justify-content: center;">
    <span class="wf-icon" data-wf-id="loading_spinner">[icon]</span>
    <span class="wf-text" data-wf-id="loading_text">認証中...</span>
    <button class="wf-button" data-wf-id="cancel_button" disabled>キャンセル</button>
  </main>
</div>
```

---

## order.cart

### 入力YAML

```yaml
# docs/uc/001-ec-checkout-flow/yaml/order/state.yaml
- id: cart
  type: state
  wireframe:
    type: col
    layout:
      width: fill
      height: fill
      min_height: 640
    children:
      - type: header
        id: cart_header
        layout:
          height: 56
          padding:
            x: 16
            y: 8
          justify: center
        children:
          - type: text
            id: cart_header_title
            label: "ショッピングカート"
      - type: row
        id: cart_shell
        layout:
          grow: true
          gap: 16
        children:
          - type: sidebar
            id: cart_sidebar
            layout:
              width: 220
              padding: 16
              gap: 8
            children:
              - type: text
                id: cart_summary_title
                label: "注文サマリー"
              - type: text
                id: cart_total_label
                label: "合計金額"
          - type: main
            id: cart_main
            layout:
              grow: true
              padding: 16
              gap: 12
              scroll: y
            children:
              - type: text
                id: cart_items_label
                label: "カート内アイテム"
              - type: divider
                id: items_divider
              - type: row
                id: item_row
                layout:
                  gap: 12
                  align: center
                children:
                  - type: image
                    id: item_thumbnail
                    layout:
                      width: 96
                      height: 96
                  - type: col
                    id: item_info_col
                    layout:
                      grow: true
                      gap: 8
                    children:
                      - type: text
                        id: item_name
                        label: "商品名"
                      - type: badge
                        id: shipping_badge
                        label: "送料無料"
              - type: button
                id: proceed_button
                label: "購入手続きへ"
                fires: view_checkout
      - type: footer
        id: cart_footer
        layout:
          height: 48
          padding:
            x: 16
            y: 8
        children:
          - type: text
            id: cart_footer_note
            label: "価格は税込表示です"
```

### HTML出力期待値

```html
<div class="wf-col" style="width: 100%; height: 100%; min-height: 640px;">
  <header class="wf-header" data-wf-id="cart_header" style="height: 56px; min-height: 56px; padding: 8px 16px 8px 16px; justify-content: center;">
    <span class="wf-text" data-wf-id="cart_header_title">ショッピングカート</span>
  </header>
  <div class="wf-row" data-wf-id="cart_shell" style="flex: 1 1 0%; min-width: 0; min-height: 0; gap: 16px;">
    <aside class="wf-sidebar" data-wf-id="cart_sidebar" style="width: 220px; gap: 8px; padding: 16px;">
      <span class="wf-text" data-wf-id="cart_summary_title">注文サマリー</span>
      <span class="wf-text" data-wf-id="cart_total_label">合計金額</span>
    </aside>
    <main class="wf-main" data-wf-id="cart_main" style="flex: 1 1 0%; min-width: 0; min-height: 0; gap: 12px; padding: 16px; overflow-y: auto;">
      <span class="wf-text" data-wf-id="cart_items_label">カート内アイテム</span>
      <hr class="wf-divider" data-wf-id="items_divider" />
      <div class="wf-row" data-wf-id="item_row" style="gap: 12px; align-items: center;">
        <div class="wf-image" data-wf-id="item_thumbnail" style="width: 96px; height: 96px; min-height: 96px;">[image]</div>
        <div class="wf-col" data-wf-id="item_info_col" style="flex: 1 1 0%; min-width: 0; min-height: 0; gap: 8px;">
          <span class="wf-text" data-wf-id="item_name">商品名</span>
          <span class="wf-badge" data-wf-id="shipping_badge">送料無料</span>
        </div>
      </div>
      <button class="wf-button" data-wf-id="proceed_button" data-wf-fires="view_checkout">購入手続きへ</button>
    </main>
  </div>
  <footer class="wf-footer" data-wf-id="cart_footer" style="height: 48px; min-height: 48px; padding: 8px 16px 8px 16px;">
    <span class="wf-text" data-wf-id="cart_footer_note">価格は税込表示です</span>
  </footer>
</div>
```

---

## order.checkout_screen

### 入力YAML

```yaml
# docs/uc/001-ec-checkout-flow/yaml/order/state.yaml
- id: checkout_screen
  type: state
  wireframe:
    type: col
    layout:
      width: fill
      min_height: 640
    children:
      - type: header
        id: checkout_header
        layout:
          height: 56
          padding:
            x: 16
            y: 8
        children:
          - type: text
            id: checkout_header_title
            label: "チェックアウト"
      - type: main
        id: checkout_main
        layout:
          grow: true
          padding: 16
          gap: 16
          scroll: y
        children:
          - type: card
            id: shipping_card
            layout:
              gap: 12
              padding: 16
            children:
              - type: text
                id: shipping_title
                label: "配送先情報"
              - type: grid
                id: address_grid
                cols: 2
                layout:
                  gap: 12
                children:
                  - type: input
                    id: postal_code_input
                    label: "郵便番号"
                    placeholder: "123-4567"
                    span: 1
                  - type: input
                    id: prefecture_input
                    label: "都道府県"
                    span: 1
                  - type: input
                    id: address_full_input
                    label: "番地・建物名"
                    placeholder: "東京都千代田区..."
                    span: 2
          - type: card
            id: payment_card
            layout:
              gap: 12
              padding: 16
            children:
              - type: text
                id: payment_title
                label: "支払い方法"
              - type: select
                id: payment_method_select
                label: "支払い方法を選択"
              - type: checkbox
                id: tos_agree_checkbox
                label: "利用規約に同意する"
          - type: button
            id: confirm_button
            label: "注文を確定する"
            fires: checkout_submitted
      - type: footer
        id: checkout_footer
        layout:
          height: 48
          padding:
            x: 16
            y: 8
        children:
          - type: button
            id: back_to_cart_button
            label: "カートに戻る"
            fires: view_cart
```

### HTML出力期待値

```html
<div class="wf-col" style="width: 100%; min-height: 640px;">
  <header class="wf-header" data-wf-id="checkout_header" style="height: 56px; min-height: 56px; padding: 8px 16px 8px 16px;">
    <span class="wf-text" data-wf-id="checkout_header_title">チェックアウト</span>
  </header>
  <main class="wf-main" data-wf-id="checkout_main" style="flex: 1 1 0%; min-width: 0; min-height: 0; gap: 16px; padding: 16px; overflow-y: auto;">
    <section class="wf-card" data-wf-id="shipping_card" style="gap: 12px; padding: 16px;">
      <span class="wf-text" data-wf-id="shipping_title">配送先情報</span>
      <div class="wf-grid" data-wf-id="address_grid" style="grid-template-columns: repeat(2, 1fr); gap: 12px;">
        <div class="wf-field" data-wf-id="postal_code_input" style="grid-column: span 1;">
          <label>郵便番号</label>
          <input type="text" placeholder="123-4567" />
        </div>
        <div class="wf-field" data-wf-id="prefecture_input" style="grid-column: span 1;">
          <label>都道府県</label>
          <input type="text" />
        </div>
        <div class="wf-field" data-wf-id="address_full_input" style="grid-column: span 2;">
          <label>番地・建物名</label>
          <input type="text" placeholder="東京都千代田区..." />
        </div>
      </div>
    </section>
    <section class="wf-card" data-wf-id="payment_card" style="gap: 12px; padding: 16px;">
      <span class="wf-text" data-wf-id="payment_title">支払い方法</span>
      <div class="wf-field" data-wf-id="payment_method_select">
        <label>支払い方法を選択</label>
        <select></select>
      </div>
      <label class="wf-checkbox" data-wf-id="tos_agree_checkbox"><input type="checkbox" /> 利用規約に同意する</label>
    </section>
    <button class="wf-button" data-wf-id="confirm_button" data-wf-fires="checkout_submitted">注文を確定する</button>
  </main>
  <footer class="wf-footer" data-wf-id="checkout_footer" style="height: 48px; min-height: 48px; padding: 8px 16px 8px 16px;">
    <button class="wf-button" data-wf-id="back_to_cart_button" data-wf-fires="view_cart">カートに戻る</button>
  </footer>
</div>
```
