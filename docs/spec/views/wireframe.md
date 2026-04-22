---
scope: docs/spec/views/wireframe.md
status: confirmed
last_updated: 2026-04-23
summary: >
  wireframe DSLのrenderルール定義。
  stateノードに付属するwireframeセクションをHTML+最小CSSで出力する際の
  要素種別・フィールド制約・renderルールを定義する。
depends_on:
  - docs/adr/019-state-node.md
  - docs/adr/018-event-node.md
  - docs/adr/029-wireframe-dsl.md
---

# wireframe renderルール

## スコープ

このspecはbrewprint YAMLの`wireframe`セクションのrenderルールを定義する。

- **対象ノード**: `type: state` のノードのみ
- **出力形式**: HTML + 最小CSS（インラインstyle）
- **対象外**: style（色・フォント・余白の詳細指定）、データバインディング

## 要素の分類

wireframeの要素は **container** と **leaf** に分類される。

### containerノード

子要素（`children`）を持てる要素。`children`は必須。

| type | 説明 |
|------|------|
| `col` | 縦並び（stack相当） |
| `row` | 横並び |
| `grid` | グリッドレイアウト。`cols`で列数指定（必須） |
| `card` | 枠付きの塊 |
| `sidebar` | サイドバー領域 |
| `header` | ヘッダー領域 |
| `footer` | フッター領域 |

### leafノード

`children`を持てない要素。

**interactive（`fires`を持てる）**

| type | 説明 |
|------|------|
| `button` | ボタン。`label`必須 |
| `input` | テキスト入力。`placeholder`任意 |
| `password` | パスワード入力。`placeholder`任意 |
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
| `id` | string | コンポーネントID。テスト・レンダラ内部参照・将来のdiff対象として使用 |
| `label` | string | 表示テキスト。type依存で意味が変わる。image/icon/divider/containerノードでは不要 |

### containerノード専用

| フィールド | 対象 | 型 | 説明 |
|-----------|------|----|------|
| `children` | 全container | array | 子ノードのリスト（必須） |
| `cols` | grid | integer | グリッドの列数（必須） |

### interactive要素専用

| フィールド | 対象 | 型 | 説明 |
|-----------|------|----|------|
| `fires` | button/input/password/select/checkbox/radio | string | 操作時に発火するevent ID |
| `disabled` | button/input/password/select/checkbox/radio | boolean | 非活性フラグ |
| `placeholder` | input/password | string | 入力欄のプレースホルダーテキスト |

### gridレイアウト用

| フィールド | 型 | 説明 |
|-----------|-----|------|
| `span` | integer | grid内で何列分を占有するか。gridの子ノードに記述する。バリデーションはしないが、grid外での使用は無意味 |

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

## renderルール

### col

```
display: flex; flex-direction: column; gap: 8px;
```

### row

```
display: flex; flex-direction: row; gap: 8px;
```

### grid

```
display: grid; grid-template-columns: repeat({cols}, 1fr); gap: 8px;
```

子ノードに`span`がある場合：

```
grid-column: span {span};
```

### card

```
border: 1px solid #ccc; border-radius: 4px; padding: 12px;
```

### sidebar / header / footer

配置の意味を示す骨格として描画する。詳細なpositioningはstyleの範囲のため定義しない。

### button

```html
<button disabled?>{label}</button>
```

`disabled: true`の場合は`disabled`属性を付与。

### input

```html
<div>
  <label>{label}</label>
  <input type="text" placeholder="{placeholder}" disabled? />
</div>
```

### password

```html
<div>
  <label>{label}</label>
  <input type="password" placeholder="{placeholder}" disabled? />
</div>
```

### select

```html
<div>
  <label>{label}</label>
  <select disabled?></select>
</div>
```

### checkbox

```html
<label><input type="checkbox" disabled? /> {label}</label>
```

### radio

```html
<label><input type="radio" disabled? /> {label}</label>
```

### text

```html
<span>{label}</span>
```

### badge

```html
<span style="border: 1px solid #ccc; border-radius: 12px; padding: 2px 8px;">{label}</span>
```

### image

```html
<div style="border: 1px dashed #ccc; background: #f5f5f5;">[image]</div>
```

### icon

```html
<span style="border: 1px dashed #ccc;">[icon]</span>
```

### divider

```html
<hr />
```

## render例

### ログイン画面

YAML:
```yaml
- id: login_screen
  type: state
  wireframe:
    type: col
    children:
      - type: input
        id: username_input
        label: ユーザー名
        placeholder: username
        fires: username_changed
      - type: password
        id: password_input
        label: パスワード
        fires: password_changed
      - type: button
        id: submit_button
        label: ログイン
        fires: submit_clicked
```

HTML出力:
```html
<div style="display: flex; flex-direction: column; gap: 8px;">
  <div>
    <label>ユーザー名</label>
    <input type="text" placeholder="username" />
  </div>
  <div>
    <label>パスワード</label>
    <input type="password" />
  </div>
  <button>ログイン</button>
</div>
```

### gridレイアウトを含む画面

YAML:
```yaml
- id: profile_screen
  type: state
  wireframe:
    type: col
    children:
      - type: grid
        cols: 2
        children:
          - type: input
            label: 姓
            span: 1
          - type: input
            label: 名
            span: 1
          - type: input
            label: メールアドレス
            span: 2
      - type: button
        label: 保存
        fires: save_clicked
```

HTML出力:
```html
<div style="display: flex; flex-direction: column; gap: 8px;">
  <div style="display: grid; grid-template-columns: repeat(2, 1fr); gap: 8px;">
    <div style="grid-column: span 1;">
      <label>姓</label>
      <input type="text" />
    </div>
    <div style="grid-column: span 1;">
      <label>名</label>
      <input type="text" />
    </div>
    <div style="grid-column: span 2;">
      <label>メールアドレス</label>
      <input type="text" />
    </div>
  </div>
  <button>保存</button>
</div>
```
