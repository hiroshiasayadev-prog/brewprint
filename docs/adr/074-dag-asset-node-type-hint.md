# 074: DAG asset node label の TypeRef hint 表示

- **status**: proposed
- **date**: 2026-05-11

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-060 で TypeRef と flow wiring type compatibility を導入し、task params / returns / model fields などに primitive / named model / inline `list<T>` / `dict<T>` を指定できるようにした。

一方、現行 DAG render の Mermaid 本体では、asset node の label は asset name を中心に表示される。
そのため、人間が DAG を見たとき、どの asset が `any` なのか、list なのか、dict なのか、どの named model なのかを Mermaid 図だけでは把握しづらい。

型の詳細は DAG Markdown の `## Tasks` detail section にある params / returns table で確認できる。
しかし、DAG 本体で dataflow を追っている最中に最低限の型ヒントが見えないと、特に UC-002 のような MCP contract / QueryService flow では `query_result: any` や `response: get_reference_tree_response` のような重要な情報を見落としやすい。

ただし、Mermaid DAG 本体に full TypeRef をすべて表示すると、node label が長くなり、処理フローの可読性を損なう。
例えば `dict<list<diagnostic>>` や深い nested TypeRef をそのまま表示すると、DAG が型展開の図に寄りすぎる。

本ADRは、DAG asset node label に full TypeRef ではなく top-level の TypeRef hint を表示する方針を定める。

## 決定

### 1. DAG asset node label に TypeRef hint を表示する

DAG render の Mermaid 本体において、asset node の label に TypeRef hint を表示する。

形式は以下とする。

```text
{asset_name}: {type_hint}
```

例:

```text
response: get_reference_tree_response
query_result: any
validated_items: list
config_by_file: dict
```

この表示は、人間が DAG 本体だけを見ても dataflow 上の値の大まかな型を把握できるようにするためのものである。

### 2. TypeRef hint は full TypeRef ではなく top-level 表示とする

DAG 本体に表示する `type_hint` は、full TypeRef ではなく top-level の型ヒントとする。

| TypeRef | DAG label の type_hint |
|---|---|
| primitive | primitive 名をそのまま表示する。例: `str`, `int`, `bool`, `any` |
| named model | model local id を表示する |
| inline `list<T>` | `list` と表示する |
| inline `dict<T>` | `dict` と表示する |

例:

```text
items: list                    # full TypeRef: list<cart_item>
diagnostics_by_file: dict      # full TypeRef: dict<list<diagnostic>>
response: get_reference_tree_response
query_result: any
```

inline container TypeRef の要素型 / value 型は DAG 本体には表示しない。
full TypeRef は Markdown detail section の params / returns table に残す。

### 3. named list / dict model は named model として表示する

TypeRef が named model の場合、その model の kind が `list` / `dict` であっても、DAG label では model local id を表示する。

例:

```yaml
nodes:
  - id: diagnostic_list
    type: model
    main: true
    kind: list
    element: diagnostic
```

この model を参照する asset は以下のように表示する。

```text
diagnostics: diagnostic_list
```

`diagnostics: list` とは表示しない。

理由は、named model は設計上の意味を持つ型であり、その名前を潰さない方が人間にとって有用だからである。

### 4. named model は QID ではなく local id を表示する

DAG node label では、named model TypeRef は QualifiedID ではなく local id を表示する。

例:

```text
response: get_reference_tree_response
```

以下のような長い QID 表示は行わない。

```text
response: mcp.model.get_reference_tree_response
```

DAG 本体は処理フローを読むための view であり、長い QID を label に含めると可読性が下がる。
必要な full identity は Markdown detail section、MCP inspect、または model render / catalog render で確認する。

ただし、同一 DAG render scope 内で別 module の同名 model が複数出現する場合、local id だけでは区別できない。
この場合、衝突した named model の TypeRef hint は shortened QID に fallback する。
shortened QID は、同一 render scope 内で一意に区別できる最短の module-qualified 表記とする。

shortened QID の算出は以下の規則とする。

1. まず local id を候補にする
2. 同一 DAG render scope 内で local id が衝突する場合、module path の末尾 segment を local id の前に付ける
3. それでも衝突する場合、module path を末尾から前方へ 1 segment ずつ伸ばす
4. 一意になった時点の suffix-qualified id を採用する
5. module path 全体を使っても一意にならない場合は full QID を使う

例:

```text
# auth/user_response と payment/user_response が衝突する場合
response: auth.user_response
response: payment.user_response

# foo/auth/user_response と bar/auth/user_response が衝突する場合
response: foo.auth.user_response
response: bar.auth.user_response
```

この fallback は衝突した named model TypeRef にだけ適用する。
衝突していない named model TypeRef は local id のまま表示する。

Markdown detail section には、衝突有無にかかわらず full TypeRef / full identity を残す。

### 5. 対象は asset node と params boundary asset とする

TypeRef hint 表示の対象は、DAG 上で asset として表示される node である。

対象:

- main task params の boundary asset
- task `returns` から生成される asset
- join `returns` から生成される asset
- `foreach.returns` から生成される collected asset

非対象:

- task node
- branch / fork / join control node
- store node
- initialized store node
- external task node
- `_start` / `_end`

store node に保持 model を表示するかどうかは、本ADRでは扱わない。

### 6. `subgraph returns` は対象外とする

ADR-064 により、DAG render の `subgraph returns` は廃止されている。
return は `returns.source` から `_end` への data line で表す。

そのため、本ADRでは returns boundary asset への TypeRef hint 表示は扱わない。

TypeRef hint は、実際に DAG 上に存在する asset node にのみ表示する。

### 7. full TypeRef は Markdown detail section に残す

DAG Mermaid 本体では top-level hint のみを表示するが、full TypeRef は Markdown detail section に残す。

例:

```text
# Mermaid asset node label
diagnostics_by_file: dict
```

```markdown
#### Returns

| name | model | source |
|---|---|---|
| diagnostics_by_file | dict<list<diagnostic>> | build_diagnostics |
```

これにより、DAG 本体の可読性と詳細情報の保持を両立する。

### 8. TypeRef が解決不能な場合は type hint を省略する

TypeRef が invalid / unresolved の場合、DAG renderer は type hint を省略し、asset name のみを表示してよい。

例:

```text
response
```

TypeRef の不正や未解決は diagnostic の責務であり、DAG render は追加の推測表示を行わない。

DAG 本体で `: ?` などの placeholder を出さない理由は、DAG view が正常系 dataflow の可読性を優先するためである。
型エラーの詳細は diagnostics に一元化し、Markdown detail section では解決できた full TypeRef のみを表示する。
TypeRef が invalid / unresolved の場合、該当箇所は diagnostics で確認する。

### 9. Mermaid escaping を避けるため full container TypeRef は表示しない

inline `list<T>` / `dict<T>` の `<` / `>` を Mermaid label に直接表示すると、escaping や readability の問題が起きうる。

本ADRでは top-level hint のみを表示するため、inline container は `list` / `dict` と表示し、Mermaid label の複雑化を避ける。

## 理由

### なぜ asset node に型ヒントを出すか

DAG は dataflow を読むための view である。
asset node に型ヒントがないと、どの dataflow が `any` なのか、どの named response model なのか、list / dict のような container なのかを Mermaid 図だけでは把握しづらい。

型ヒントを出すことで、人間は DAG 本体を見ながら dataflow の大まかな shape を追いやすくなる。

特に `any` は設計上の逃げ口であり、DAG 上で見えること自体にレビュー価値がある。

### なぜ full TypeRef ではなく top-level hint か

full TypeRef を表示すると、`dict<list<diagnostic>>` のような長い label が発生し、DAG 本体の可読性が下がる。
DAG 本体は処理フローを読むための view であり、型詳細の完全展開は Markdown detail section や model render / catalog render の責務である。

そのため、DAG 本体には top-level hint のみを表示する。

### なぜ named model は local id か

named model は設計上の意味を持つ型である。
`diagnostic_list` のような named list model を単に `list` と表示すると、設計概念が失われる。

一方で、QualifiedID を表示すると label が長くなりすぎる。
DAG 本体では local id を表示し、full identity は detail section / MCP / model render で確認する。

### なぜ params boundary asset も対象にするか

main task params は DAG の入力境界である。
入力境界の型が見えると、DAG の dataflow を読み始める時点で何が流入しているかを把握しやすい。

そのため、`subgraph params` 内の boundary asset も TypeRef hint 表示の対象とする。

## 却下した代替案

### 代替案A: asset node label に型を表示しない

- 利点: 現行表示を維持でき、DAG label が短い
- 欠点: DAG 本体だけでは dataflow の型が分かりづらい。`any` / list / dict / response model の違いを detail section まで見に行く必要がある

→ 却下。top-level hint を表示する。

### 代替案B: full TypeRef を表示する

```text
diagnostics_by_file: dict<list<diagnostic>>
```

- 利点: DAG 本体だけで完全な型が分かる
- 欠点: label が長くなり、DAG が読みにくくなる。Mermaid escaping の問題も増える

→ 却下。DAG 本体は top-level hint に留め、full TypeRef は detail section に出す。

### 代替案C: 型を asset name の前に表示する

```text
[get_reference_tree_response] response
```

- 利点: 型が先に目に入る
- 欠点: asset name の視認性が下がる。DAG では「何の値か」を示す asset name が主役である

→ 却下。`{asset_name}: {type_hint}` 形式にする。

### 代替案D: named list / dict model を list / dict として表示する

```text
diagnostics: list
```

- 利点: top-level container kind が分かる
- 欠点: `diagnostic_list` のような named model の設計意味が失われる

→ 却下。named model は local id を表示する。

### 代替案E: store node にも型を表示する

- 利点: store が保持する model も DAG 上で見える
- 欠点: store は asset ではなく runtime data instance であり、`store.of` / `initializes[].model` の表示は別論点である

→ 却下。本ADRでは asset node label に限定する。

## 影響

### spec への影響

本ADR受理後、以下の spec 更新が必要になる。

- `docs/spec/views/dag.md`
  - asset node label の形式を `{asset_name}: {type_hint}` に更新する
  - params boundary asset も TypeRef hint 表示対象に含める
  - top-level TypeRef hint の算出規則を定義する
  - named model local id が同一 render scope 内で衝突する場合の shortened QID fallback を定義する
  - shortened QID は module path の末尾 segment から順に前方へ伸ばし、一意になった suffix-qualified id を使うことを定義する
  - TypeRef 解決不能時の fallback を定義する
  - full TypeRef は Markdown detail section に残すことを明記する
  - invalid / unresolved TypeRef の説明は diagnostics に一元化することを明記する

### render 実装への影響

DAG renderer は asset node label 生成時に TypeRef hint を算出する必要がある。

対象は params boundary asset、task / join returns asset、foreach collected asset である。

TypeRef hint の算出は top-level のみとし、inline container の内部 TypeRef は展開しない。

named model local id が同一 DAG render scope 内で衝突する場合、renderer は該当 hint を shortened QID に fallback する。
shortened QID は module path の末尾 segment から順に前方へ伸ばし、一意になった suffix-qualified id を採用する。

### UC / fixture への影響

既存 DAG render golden は asset node label が変わるため更新が必要になる。

特に UC-002 では、`query_result: any` や `response: get_reference_tree_response` のように、dataflow 上の型が Mermaid 本体で見えるようになる。

具体的な golden 更新は task file / UC task file で追跡する。

### ADR-071 / ADR-075 との関係

ADR-071 / ADR-075 は file-private helper model の Markdown render exposure を扱う。
本ADRは DAG Mermaid 本体に asset の top-level 型ヒントを表示する。

helper model の詳細展開は本ADRの対象ではない。
詳細は `## Private models`、model file render、model catalog view で確認する。

## Evidence

- commit: 5ae7769
- impl commit: tbd
- close boundary: M15 / `v1.1.0-spec` では follow-up scope として deferred。実装は含めない。
- 参考: ADR-060 TypeRef、ADR-064 DAG returns.source render、ADR-071 task file helper model render exposure、ADR-075 model file render
