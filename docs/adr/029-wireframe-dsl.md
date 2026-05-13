# 029: wireframe DSLをYAML内に持つ

- **status**: accepted
- **date**: 2026-04-21
- **supersedes**:

## 背景

brewprintはシステム設計言語として、DAG・ER・state diagramをYAMLから導出できる。
しかしUIの「どこに何を配置するか」という構造情報を表現する手段がなかった。

UMLにはwireframeの標準記法が存在しない。既存ツール（Figma・draw.io・PlantUML Salt）はいずれも以下の問題を持つ：

- Figma / draw.io: テキストベースでない、LLMが生成・読解できない
- PlantUML Salt: レイアウト自由度が低く、brewprintのYAMLと統合できない

「LLMが生成できる + LLMが読める + brewprintのeventと紐付けられる」を同時に満たすwireframe記法が必要だった。

## 決定

brewprint YAMLの`wireframe`セクションとして、独自のwireframe DSLを持つ。

### 設計方針

- **styleは対象外**。配置・構造・コンポーネント種別のみを記述する
- **データバインディングは対象外**。wireframeはUIの骨格図であり、「何を表示するか」はstateのmodelとDAGで表現する
- **HTML標準レベルの要素セット**に絞る。UIフレームワーク依存を持たない
- **`fires`フィールド**でbrewprintのevent IDと紐付ける（interactive要素のみ）
- **wireframeはstateノード専用**。stateノードのみが`wireframe`フィールドを持てる
- **単一root**。`wireframe:`直下は1つのcontainerノード
- **状態ごとの表示差分はstateノードとして独立させる**。`loading`や`error`は別stateノードを定義し、それぞれに`wireframe`を持たせる

### 要素セット

**レイアウト**

| 種別 | 説明 |
|------|------|
| `row` | 横並び |
| `col` | 縦並び（stack相当） |
| `grid` | グリッド。`cols`で列数指定 |
| `sidebar` | サイドバー領域 |
| `header` | ヘッダー領域 |
| `footer` | フッター領域 |
| `card` | カード（枠付きの塊） |

**コンポーネント**

| 種別 | 説明 |
|------|------|
| `text` | テキスト・ラベル |
| `button` | ボタン |
| `input` | テキスト入力 |
| `password` | パスワード入力 |
| `select` | ドロップダウン |
| `checkbox` | チェックボックス |
| `radio` | ラジオボタン |
| `image` | 画像プレースホルダー |
| `icon` | アイコンプレースホルダー |
| `badge` | バッジ・タグ |
| `divider` | 区切り線 |

**共通フィールド（全要素）**

| フィールド | 説明 |
|-----------|------|
| `id` | このコンポーネントのID（テスト・レンダラ内部参照・将来のdiff対象） |
| `label` | 表示テキスト（type依存で意味が変わる。image/icon/divider/layoutノードでは不要） |

**containerノード専用フィールド**

| フィールド | 対象 | 説明 |
|-----------|------|------|
| `children` | row/col/grid/sidebar/header/footer/card | 子ノードのリスト（必須） |
| `cols` | grid | グリッドの列数（必須） |

**interactive要素専用フィールド**

| フィールド | 対象 | 説明 |
|-----------|------|------|
| `fires` | button/input/password/select/checkbox/radio | 操作時に発火するevent ID |
| `disabled` | button/input/password/select/checkbox/radio | 非活性フラグ |
| `placeholder` | input/password | 入力欄のプレースホルダー |

**gridレイアウト用フィールド**

| フィールド | 説明 |
|-----------|------|
| `span` | grid内で何列分を占有するか。gridの子ノードに書く。バリデーションはしないが、grid外での使用は無意味 |

### 記述例（ログイン画面）

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

### レンダリング

wireframeセクションはHTML+最小CSSで機械的にレンダリングする。
styleの付与はしない。配置構造のみを視覚化する。

## 理由

### HTML標準レベルに絞った理由

Mantineなどのフレームワークのコンポーネント名・propsをそのまま使う案も検討したが、以下の理由で却下した：

- フレームワークのバージョンアップでprops名が変わるとbrewprint側も追従が必要になる
- `button`・`input`・`select`などはHTML標準の概念であり、どのフレームワークに依存しても変わらない
- wireframeの目的は「どこに何を置くか」の構造把握であり、フレームワーク固有の表現力は不要

### stateノード専用にした理由

wireframeは「この状態の画面がこう見える」という表現であり、stateとの対応が自然。
taskノードやactorノードにwireframeを持たせても設計上の意味が薄い。

### 状態ごとの表示差分をstateノードで表現する理由

「ローディング中はspinnerを表示」「エラー時はエラーメッセージを表示」などの表示差分を
wireframe内の`state`フィールドで表現する案を検討したが却下した。

`loading`や`error`はbrewprintの意味論上まさにstateである。
同一wireframe内で別stateを参照する構造は意味論が複雑になる。
各stateが自身のwireframeを持つことでシンプルさと一貫性が保たれる。

### データバインディングを対象外にした理由

wireframeの本来の定義は「UIの構造・レイアウト・要素の配置を示す骨格図。見た目とデータは含まない」。
「何を表示するか」はwireframeの責務ではない。

将来的にバインディングが必要になった場合は、別ADRでフィールドを定義することで
既存のwireframe DSLを壊さずに拡張できる。
その際の設計候補：

- `context`をcontainerノードの一種として追加する（Reactの`<Context.Provider>`相当）
- `context`ノードは`store`フィールドでbrewprintのstoreノードを参照する
- `context`は任意の深さに配置できる
- `context`配下のコンポーネントが`bind`フィールドでstoreの値を参照する

```yaml
# 将来の拡張イメージ（現時点では未定義）
wireframe:
  type: context
  store: session_store
  children:
    - type: col
      children:
        - type: text
          bind: store.username  # 未定義フィールド
```

### `fires`をinteractive要素限定にした理由

image・divider・layoutノードなど、ユーザー操作が発生しない要素にevent IDを紐付けるユースケースがない。
全要素共通にすると「どの要素がeventを発火できるか」がLLMに不明瞭になる。

### 単一rootにした理由

ReactのコンポーネントツリーはContext → Layout → Childrenという単一rootのtree構造であり、
UIのrenderingモデルとして自然。複数rootを許すとレンダラがトップレベルの並置をどう扱うか曖昧になる。

### YAMLに埋め込む理由

wireframeを別ファイル（SVGなど）に分離する案も検討したが：

- YAML内に収めることでsingle source of truthが保たれる
- LLMがYAML全体をコンテキストとして読む際に、wireframeも一体で参照できる
- MCPツールでのコンテキスト供給が単純になる

## 影響

- `docs/spec/nodes.md` のstateノードに`wireframe`フィールドを追加する必要がある
- wireframeのレンダリング仕様を `docs/spec/views/wireframe.md` として新規作成する
- MCPツールの`inspect`がwireframeセクションを返せるよう拡張が必要になる（実装時に判断）

## Evidence

- commit: e7dd532
- impl commit: tbd
- 参考: PlantUML Salt参考、wireweave DSL参考
