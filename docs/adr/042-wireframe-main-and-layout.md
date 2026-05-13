# 042: wireframeのmain containerとlayout object

- **status**: accepted
- **date**: 2026-04-26
- **supersedes**:

## 背景

ADR-029でwireframe DSLをYAML内に持つことを決定し、`docs/spec/views/wireframe.md` で `row` / `col` / `grid` / `header` / `sidebar` / `footer` / `card` などの要素種別を定義した。

しかし、UC-001のwireframeを画面として実際にrenderする検討を進めると、以下の不足が明確になった。

1. **main領域が表現できない**
   - `header` / `sidebar` / `footer` はあるが、主要コンテンツ領域を示す `main` がない。
   - `row` + `sidebar` + `col` で代用は可能だが、HTML出力上も意味上も「これは主要コンテンツ」という情報が失われる。

2. **layout指定がないと画面を組めない**
   - 一般的なWebアプリ画面では、headerの高さ、sidebarの幅、mainのスクロール、row/col間のgap、padding、align/justifyなどが最低限必要になる。
   - 現行DSLは「何を置くか」は表現できるが、「画面領域としてどう組むか」を十分に表現できない。

3. **render profileを定義する前提が不足している**
   - `.wf-page` / `.wf-header` / `.wf-sidebar` / `.wf-main` / `.wf-row` / `.wf-col` のような安定したHTML/CSS構造を生成するには、YAML側に対応する意味情報が必要。
   - Go rendererが暗黙に `col` を `main` 扱いするなどの補完をすると、YAMLがsingle source of truthでなくなる。

ADR-029では「styleは対象外」としていたが、ここで必要なのは色・フォント・角丸・影などの装飾ではなく、画面構造を成立させるための**構造的layout情報**である。

## 決定

### 1. `main` container typeを追加する

wireframe DSLのcontainer要素に `main` を追加する。

```yaml
wireframe:
  type: col
  children:
    - type: header
      children: [...]
    - type: row
      children:
        - type: sidebar
          children: [...]
        - type: main
          children: [...]
```

`main` はHTMLの `<main>` に対応する主要コンテンツ領域を表す。

| type | 分類 | children | HTML出力 |
|---|---|---|---|
| `main` | container | 必須 | `<main class="wf-main">...</main>` |

### 2. `layout` objectをwireframe要素の共通optional fieldとして追加する

全wireframe要素は任意で `layout:` を持てる。

```yaml
- type: sidebar
  id: nav_sidebar
  layout:
    width: 220
    padding:
      x: 12
      y: 8
  children: [...]

- type: main
  id: content_main
  layout:
    grow: true
    scroll: y
    padding: 16
  children: [...]
```

`layout` は視覚装飾ではなく、画面構造を成立させるための制約だけを扱う。

### 3. `layout` v1のフィールドセット

初期バージョンでは、以下のフィールドだけを定義する。

| フィールド | 型 | 対象 | 内容 |
|---|---|---|---|
| `width` | size | 全要素 | 幅。数値はpx。予約語として `fill` / `fit` を許可 |
| `height` | size | 全要素 | 高さ。数値はpx。予約語として `fill` / `fit` を許可 |
| `min_width` | int | 全要素 | 最小幅px |
| `min_height` | int | 全要素 | 最小高さpx |
| `grow` | bool | flex item | `true` で残り領域を占有する |
| `gap` | int | container | 子要素間のgap px |
| `padding` | int または padding object | container | padding。数値なら全方向px。leafではv1非対応 |
| `align` | enum | container | 交差軸方向の配置。`start` / `center` / `end` / `stretch` |
| `justify` | enum | container | 主軸方向の配置。`start` / `center` / `end` / `between` |
| `scroll` | enum | container | `none` / `x` / `y` / `both` |

#### size値

`width` / `height` の値は以下のいずれか。

```yaml
layout:
  width: 220       # px
  height: 56       # px
```

```yaml
layout:
  width: fill      # 親の利用可能な幅を埋める。flex growは意味しない
  height: fit      # 内容に合わせる
```

| 値 | 意味 | HTML/CSS変換の基本方針 |
|---|---|---|
| number | px固定値 | `Npx` |
| `fill` | 親の利用可能サイズを埋める | `width: 100%` / `height: 100%`。flexの残り領域占有はしない |
| `fit` | 内容サイズに合わせる | widthでは `fit-content`、heightでは `auto` を基本とする |

数値はpxとして扱う。任意のCSS文字列は許可しない。

`min_width` / `min_height` はpx整数のみとする。`fill` / `fit` は許可しない。最小サイズとして「親を埋める」「内容に合わせる」を指定しても意味が不安定であり、実用上も必要性が薄いため。

#### grow と fill の使い分け

`width: fill` / `height: fill` はサイズ指定であり、親の利用可能サイズに対して `100%` を指定するだけである。

`grow: true` はflex itemとして親の残り領域を占有する指定である。`row` / `col` の子要素でのみ有効とし、HTML/CSSでは `flex: 1 1 0%; min-width: 0; min-height: 0` を基本変換とする。

```yaml
# OK: sidebarは固定幅、mainは残り領域を占有
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

`fill` と `grow` は同義ではない。画面の残り領域を取らせたい場合は `grow: true` を使う。

#### padding object

`padding` は数値またはobjectで指定できる。

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
`top/right/bottom/left` が指定された場合は個別指定を優先する。

### 4. layoutに含めないもの

以下は `layout` に含めない。

| 項目 | 理由 |
|---|---|
| 色 | wireframeはstyleを扱わない |
| font family / font size | 見た目のstyleであり、構造ではない |
| border / radius / shadow | 装飾であり、構造ではない |
| arbitrary CSS | rendererの決定性と静的検証性を壊す |
| class名の直接指定 | HTML/CSS実装詳細への依存を生む |
| responsive breakpoint | 初期DSLとして過剰。必要になったら別ADRで追加 |

### 5. HTML render上の扱い

HTML rendererは、`layout` を決定的なinline styleまたは固定utility classへ変換してよい。
ただし、出力契約は `docs/spec/views/wireframe.md` に定義する。

推奨される基本対応は以下。

| YAML | HTML/CSS上の意味 |
|---|---|
| `type: main` | `<main class="wf-main">` |
| `layout.width: 220` | `width: 220px` |
| `layout.width: fill` | `width: 100%`。flexの残り領域占有はしない |
| `layout.height: 56` | `height: 56px; min-height: 56px` |
| `layout.grow: true` | `flex: 1 1 0%; min-width: 0; min-height: 0` |
| `layout.gap: 16` | `gap: 16px` |
| `layout.scroll: y` | `overflow-y: auto` |
| `layout.align: center` | `align-items: center` |
| `layout.justify: between` | `justify-content: space-between` |

詳細な変換ルールはspec側に置く。

### 6. validation方針

- `main` はcontainerなので `children` 必須。
- `layout` は全要素で任意。
- `gap` / `padding` / `align` / `justify` / `scroll` はcontainerでのみ有効。
- `width` / `height` / `min_width` / `min_height` / `grow` はleafにも指定可能。
- `grow: true` は `row` / `col` の子要素でのみ有効。それ以外の位置ではparser error。
- `layout` に未定義フィールドがある場合はparser error。
- 任意のCSS文字列はparser error。

## 理由

### mainを追加する理由

`header` / `sidebar` / `footer` が存在するなら、主要コンテンツ領域である `main` も同じ抽象度で存在すべき。
`col` で代用すると、HTMLのsemantic elementとしての `<main>` を生成できず、LLMや人間が画面構造を読むときにも「ここが主要領域」という情報が失われる。

### layoutを追加する理由

現行DSLは要素ツリーの構造は表現できるが、画面として成立する最低限のレイアウト制約を表現できない。
例えば以下は構造情報であり、装飾ではない。

- sidebarは固定幅
- headerは固定高さ
- mainは残り領域を占有
- mainだけ縦スクロールする
- row/colの子要素間にgapがある
- header内の要素を左右に分ける

これらをrendererが暗黙補完すると、YAMLから画面の意図を復元できなくなる。
brewprintの「YAMLがsingle source of truth」という方針に従い、layout意図はYAMLに明示する。

### style対象外方針との整合

ADR-029の「styleは対象外」は、色・フォント・余白の美的調整などの視覚デザインを対象外にする意図だった。
しかし、padding/gap/width/height/align/scrollは、画面骨格を成立させるための構造制約である。

したがって本ADRでは、以下の境界を定義する。

| 種類 | brewprintで扱うか |
|---|---|
| 構造layout | 扱う |
| 視覚style | 扱わない |

### arbitrary CSSを禁止する理由

`style: "..."` のような任意CSSを許すと、brewprintの静的検証性が失われる。
また、LLMがCSS実装詳細を直接生成することになり、wireframe DSLの抽象化が崩れる。

layoutは閉じたfield setとして持ち、rendererがHTML/CSSへ変換する。

### padding/gapをcontainerに許可する理由

padding/gapは見た目の調整にも使えるが、wireframeでは「領域と要素の分離」を示す骨格情報として必要。
これらを完全に禁止すると、header/sidebar/mainのような画面構成を表現できず、実用的なwireframeにならない。

一方で、leaf要素の内側paddingはbutton/input/text等の見た目に直結し、視覚styleとの境界が曖昧になる。そのためv1では `padding` をcontainer限定とする。leaf要素の余白が必要な場合は、親containerの `gap` / `padding` で表現する。

## 影響

- `docs/spec/views/wireframe.md` を更新する必要がある。
  - containerノード一覧に `main` を追加
  - 共通フィールドに `layout` を追加
  - `layout` objectのschemaを追加
  - renderルールに `main` とlayout変換を追加
  - HTML/CSS render profileの出力契約を追加
- `docs/spec/nodes.md` の `depends_on` に本ADRを追加し、state.wireframe 説明から `main` / `layout` を参照できるように更新する。
- UC-001 のwireframe YAMLは、`main` / `layout` を使う形に更新できる。
- Go rendererは `layout` objectを解釈してHTML/CSSへ変換する。

## Evidence
- commit: 16fc169
- impl commit: tbd
- 参考: HTML semantic elements（main/header/footer/aside）、CSS Flexbox / Grid の基本layout概念
