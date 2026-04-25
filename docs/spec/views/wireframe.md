---
scope: docs/spec/views/wireframe.md
status: confirmed
last_updated: 2026-04-26
summary: >
  wireframe DSLのrenderルール定義。
  stateノードに付属するwireframeセクションをHTML+最小CSSで出力する際の
  要素種別・フィールド制約・layout object・renderルールを定義する。
depends_on:
  - docs/adr/019-state-node.md
  - docs/adr/018-event-node.md
  - docs/adr/029-wireframe-dsl.md
  - docs/adr/042-wireframe-main-and-layout.md
---

# wireframe renderルール

## スコープ

このspecはbrewprint YAMLの`wireframe`セクションのrenderルールを定義する。

- **対象ノード**: `type: state` のノードのみ
- **入力**: `state.wireframe` 以下の単一root wireframe tree
- **出力形式**: HTML fragment + fixed CSS profile（`.wf-*` namespace）
- **対象**: 画面構造、semantic container、構造layout
- **対象外**: 視覚style（色・フォント・border/radius/shadow等）、任意CSS、class直指定、JS生成、データバインディング

`layout`で扱うのは、画面構造を成立させるための構造的layout情報に限定する。色・フォント・枠線・角丸・影などの視覚styleはwireframe DSLでは扱わない。

## 要素の分類

wireframeの要素は **container** と **leaf** に分類される。

### containerノード

子要素（`children`）を持つ要素。`children`は必須。

| type | 説明 | HTML element |
|------|------|--------------|
| `col` | 縦並び（stack相当） | `div.wf-col` |
| `row` | 横並び | `div.wf-row` |
| `grid` | グリッドレイアウト。`cols`で列数指定（必須） | `div.wf-grid` |
| `card` | 意味的なまとまり | `section.wf-card` |
| `sidebar` | サイドバー領域 | `aside.wf-sidebar` |
| `header` | ヘッダー領域 | `header.wf-header` |
| `footer` | フッター領域 | `footer.wf-footer` |
| `main` | 主要コンテンツ領域 | `main.wf-main` |

`main` は主要コンテンツ領域を表すcontainerであり、HTML renderでは `<main class="wf-main">` に対応する。

### leafノード

`children`を持てない要素。

**interactive（`fires`を持てる）**

| type | 説明 |
|------|------|
| `button` | ボタン。`label`必須 |
| `input` | テキスト入力。`label`必須、`placeholder`任意 |
| `password` | パスワード入力。`label`必須、`placeholder`任意 |
| `select` | ドロップダウン。`label`必須 |
| `checkbox` | チェックボックス。`label`必須 |
| `radio` | ラジオボタン。`label`必須 |

**non-interactive（`fires`を持たない）**

| type | 説明 |
|------|------|
| `text` | テキスト・ラベル。`label`必須 |
| `badge` | バッジ・タグ。`label`必須 |
| `image` | 画像プレースホルダー |
| `icon` | アイコンプレースホルダー |
| `divider` | 区切り線 |

## フィールド定義

### 全要素共通

| フィールド | 型 | 説明 |
|-----------|-----|------|
| `type` | string | 要素種別（必須） |
| `id` | string | コンポーネントID。HTMLでは`id`属性ではなく`data-wf-id`へ出力する |
| `label` | string | 表示テキスト。type依存で意味が変わる。image/icon/divider/containerノードでは不要 |
| `layout` | object | 構造layout指定。全要素で任意だが、指定可能fieldはcontainer/leafで異なる |

### containerノード専用

| フィールド | 対象 | 型 | 説明 |
|-----------|------|----|------|
| `children` | 全container | array | 子ノードのリスト（必須） |
| `cols` | grid | integer | グリッドの列数（必須） |

### interactive要素専用

| フィールド | 対象 | 型 | 説明 |
|-----------|------|----|------|
| `fires` | button/input/password/select/checkbox/radio | string | 操作時に発火するevent ID。HTMLでは`data-wf-fires`へ出力する |
| `disabled` | button/input/password/select/checkbox/radio | boolean | 非活性フラグ |
| `placeholder` | input/password | string | 入力欄のプレースホルダーテキスト |

### gridレイアウト用

| フィールド | 型 | 説明 |
|-----------|-----|------|
| `span` | integer | grid内で何列分を占有するか。gridの子ノードに記述する |

`span` は `grid` の直接子でのみ意味を持つ。`grid` 外の `span` はparser errorとしてよい。

## layout object

`layout` は全wireframe要素で任意に指定できる。ただしv1では、扱うfieldを構造layoutに限定する。

```yaml
layout:
  width: 220
  height: 56
  min_width: 120
  min_height: 80
  grow: true
  gap: 16
  padding: 16
  align: center
  justify: between
  scroll: y
```

### layout field一覧

| フィールド | 型 | 対象 | 内容 |
|---|---|---|---|
| `width` | size | 全要素 | 幅。数値はpx。予約語として `fill` / `fit` を許可 |
| `height` | size | 全要素 | 高さ。数値はpx。予約語として `fill` / `fit` を許可 |
| `min_width` | int | 全要素 | 最小幅px。`fill` / `fit` は不可 |
| `min_height` | int | 全要素 | 最小高さpx。`fill` / `fit` は不可 |
| `grow` | bool | 全要素 | `true`でrow/colの子として残り領域を占有 |
| `gap` | int | container | 子要素間のgap px |
| `padding` | int or object | container | container内側の構造的余白。leafではv1非対応 |
| `align` | enum | container | 交差軸方向の配置。`start` / `center` / `end` / `stretch` |
| `justify` | enum | container | 主軸方向の配置。`start` / `center` / `end` / `between` |
| `scroll` | enum | container | `none` / `x` / `y` / `both` |

### size値

`width` / `height` の値は以下のいずれか。

| 値 | 意味 | HTML/CSS変換 |
|---|---|---|
| number | px固定値 | `Npx` |
| `fill` | 親の利用可能サイズに対する100% | `width: 100%` / `height: 100%` |
| `fit` | 内容サイズ | widthでは `fit-content`、heightでは `auto` |

```yaml
layout:
  width: 220       # width: 220px
  height: 56       # height: 56px; min-height: 56px
```

```yaml
layout:
  width: fill      # width: 100%
  height: fit      # height: auto
```

数値はpxとして扱う。`"80%"` / `"12rem"` / `"calc(...)"` のような任意CSS文字列は許可しない。

`min_width` / `min_height` はpx整数のみ。`fill` / `fit` は許可しない。

### padding object

`padding` はcontainerでのみ指定できる。数値またはobjectで指定する。

```yaml
layout:
  padding: 16
```

```yaml
layout:
  padding:
    x: 16
    y: 12
```

```yaml
layout:
  padding:
    top: 8
    right: 16
    bottom: 8
    left: 16
```

`x` は left/right、`y` は top/bottom の省略形。
`top` / `right` / `bottom` / `left` が指定された場合は個別指定を優先する。

leaf要素のpaddingはv1では扱わない。button/input/text等の見た目調整になりやすく、構造layoutと視覚styleの境界が曖昧になるため。

### fill と grow

`width: fill` / `height: fill` は、親の利用可能サイズに対する `100%` 指定である。flexの残り領域占有は意味しない。

`grow: true` は、`row` / `col` の子として残り領域を占有する指定である。HTML/CSSでは `flex: 1 1 0%; min-width: 0; min-height: 0` を基本変換とする。

```yaml
- type: row
  children:
    - type: sidebar
      layout:
        width: 220
      children: [...]
    - type: main
      layout:
        grow: true
        scroll: y
      children: [...]
```

画面の残り領域を取らせたい場合は `fill` ではなく `grow: true` を使う。

## validation方針

parser / validator は以下を検証する。

- `wireframe:` 直下は単一のcontainerノード。複数rootは不可。
- containerノードは `children` 必須。
- leafノードは `children` を持てない。
- `main` はcontainerなので `children` 必須。
- `grid` は `cols` 必須。`cols` は1以上の整数。
- `layout` は全要素で任意。
- `width` / `height` / `min_width` / `min_height` / `grow` はcontainer/leafのどちらにも指定可能。
- `gap` / `padding` / `align` / `justify` / `scroll` はcontainerでのみ有効。leafで指定した場合はparser error。
- `grow: true` は `row` / `col` の直接子でのみ有効。それ以外の位置ではparser error。
- `layout` に未定義fieldがある場合はparser error。
- `style` / `class` / `css` 等、HTML/CSS実装詳細を直接指定するfieldはparser error。
- `fires` はinteractive leafでのみ指定可能。
- `id` / `label` / `placeholder` / `fires` はHTML出力時にescapeする。

## ルート構造

`wireframe:`直下は**単一のcontainerノード**。複数rootは不可。

```yaml
# OK
wireframe:
  type: col
  children: [...]

# NG（複数root）
wireframe:
  - type: col
    children: [...]
  - type: col
    children: [...]
```

## 状態ごとの表示差分

`loading`・`error`などの状態は**別stateノードとして定義**し、それぞれに`wireframe`を持たせる。
同一`wireframe`内での表示条件フィールド（`visible_when`など）は存在しない。

```yaml
# OK: stateノードを分ける
- id: login_screen
  type: state
  wireframe:
    type: col
    children:
      - type: button
        label: ログイン
        fires: submit_clicked

- id: login_loading
  type: state
  wireframe:
    type: col
    children:
      - type: text
        label: 送信中...
      - type: button
        label: ログイン
        disabled: true

# NG: wireframe内で状態を条件分岐
wireframe:
  type: col
  children:
    - type: text
      label: 送信中...
      state: login_loading   # このフィールドは存在しない
```

## HTML/CSS render profile

HTML rendererは、wireframe treeをHTML fragmentとして出力する。DOCTYPE、`html`、`head`、`body`は生成しない。

### 出力契約

- 全wireframe要素に `.wf-*` namespaceのclassを付与する。
- YAMLの `id` はHTMLの `id` 属性にはせず、`data-wf-id` に出力する。
- YAMLの `fires` は `data-wf-fires` に出力する。
- `label` / `placeholder` / `id` / `fires` はHTML escapeする。
- JSは生成しない。
- 任意CSSは受け付けない。
- `layout` は決定的なinline styleへ変換する。実装が固定utility classを使う場合も、同じ意味に変換される必要がある。

### HTML element対応

| type | 出力 |
|---|---|
| `col` | `<div class="wf-col">` |
| `row` | `<div class="wf-row">` |
| `grid` | `<div class="wf-grid">` |
| `card` | `<section class="wf-card">` |
| `sidebar` | `<aside class="wf-sidebar">` |
| `header` | `<header class="wf-header">` |
| `footer` | `<footer class="wf-footer">` |
| `main` | `<main class="wf-main">` |
| `button` | `<button class="wf-button">` |
| `input` | `<div class="wf-field"><label>...</label><input type="text" /></div>` |
| `password` | `<div class="wf-field"><label>...</label><input type="password" /></div>` |
| `select` | `<div class="wf-field"><label>...</label><select></select></div>` |
| `checkbox` | `<label class="wf-checkbox"><input type="checkbox" /> ...</label>` |
| `radio` | `<label class="wf-radio"><input type="radio" /> ...</label>` |
| `text` | `<span class="wf-text">...</span>` |
| `badge` | `<span class="wf-badge">...</span>` |
| `image` | `<div class="wf-image">[image]</div>` |
| `icon` | `<span class="wf-icon">[icon]</span>` |
| `divider` | `<hr class="wf-divider" />` |

### fixed CSS profile

rendererは以下の意味を持つ固定CSSを前提としてよい。

```css
.wf-col { display: flex; flex-direction: column; gap: 8px; }
.wf-row { display: flex; flex-direction: row; gap: 8px; }
.wf-grid { display: grid; gap: 8px; }
.wf-header, .wf-footer, .wf-sidebar, .wf-main, .wf-card { display: flex; flex-direction: column; gap: 8px; }
.wf-field { display: flex; flex-direction: column; gap: 4px; }
```

これはrender profile側の固定CSSであり、YAMLから任意に指定するstyleではない。

### layout変換

| YAML | HTML/CSS上の意味 |
|---|---|
| `layout.width: 220` | `width: 220px` |
| `layout.width: fill` | `width: 100%` |
| `layout.width: fit` | `width: fit-content` |
| `layout.height: 56` | `height: 56px; min-height: 56px` |
| `layout.height: fill` | `height: 100%` |
| `layout.height: fit` | `height: auto` |
| `layout.min_width: 120` | `min-width: 120px` |
| `layout.min_height: 80` | `min-height: 80px` |
| `layout.grow: true` | `flex: 1 1 0%; min-width: 0; min-height: 0` |
| `layout.gap: 16` | `gap: 16px` |
| `layout.padding: 16` | `padding: 16px` |
| `layout.padding.x: 16` + `layout.padding.y: 8` | `padding: 8px 16px 8px 16px` |
| `layout.padding.top/right/bottom/left` | `padding: {top}px {right}px {bottom}px {left}px`。未指定方向は0px、`x` / `y` があればその値で補完 |
| `layout.align: start` | `align-items: flex-start` |
| `layout.align: center` | `align-items: center` |
| `layout.align: end` | `align-items: flex-end` |
| `layout.align: stretch` | `align-items: stretch` |
| `layout.justify: start` | `justify-content: flex-start` |
| `layout.justify: center` | `justify-content: center` |
| `layout.justify: end` | `justify-content: flex-end` |
| `layout.justify: between` | `justify-content: space-between` |
| `layout.scroll: none` | `overflow: visible` |
| `layout.scroll: x` | `overflow-x: auto` |
| `layout.scroll: y` | `overflow-y: auto` |
| `layout.scroll: both` | `overflow: auto` |

`grow: true` と `min_width` / `min_height` を同時指定した場合、明示された `min_width` / `min_height` を優先する。

### grid render

`grid` の基礎displayは fixed CSS profile の `.wf-grid { display: grid; ... }` が担う。
`grid.cols` は `grid-template-columns` に変換し、inline styleへ出力する。
`display: grid` はinline styleへ重複出力しない。

```css
grid-template-columns: repeat({cols}, 1fr)
```

子ノードに `span` がある場合は、その子のstyleへ以下を出力する。

```css
grid-column: span {span}
```

## render例

### main + layoutを含む画面

YAML:
```yaml
- id: cart
  type: state
  wireframe:
    type: col
    children:
      - type: header
        layout:
          height: 56
        children:
          - type: text
            label: ショッピングカート
      - type: row
        layout:
          grow: true
        children:
          - type: sidebar
            layout:
              width: 220
            children:
              - type: text
                label: 注文サマリー
          - type: main
            layout:
              grow: true
              scroll: y
            children:
              - type: text
                label: カート内アイテム
```

HTML出力:
```html
<div class="wf-col">
  <header class="wf-header" style="height: 56px; min-height: 56px;">
    <span class="wf-text">ショッピングカート</span>
  </header>
  <div class="wf-row" style="flex: 1 1 0%; min-width: 0; min-height: 0;">
    <aside class="wf-sidebar" style="width: 220px;">
      <span class="wf-text">注文サマリー</span>
    </aside>
    <main class="wf-main" style="flex: 1 1 0%; min-width: 0; min-height: 0; overflow-y: auto;">
      <span class="wf-text">カート内アイテム</span>
    </main>
  </div>
</div>
```
