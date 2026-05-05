---
scope: docs/spec/edges.md
status: confirmed
last_updated: 2026-05-05
summary: >
  brewprintのエッジ記法の定義。
  ファイル内データフロー（flow:セクション）・状態遷移（transitions:セクション）・
  クロスエッジ（task nodeのreads/writes）・$シジル体系を定義する。
depends_on:
  - docs/adr/003-name-resolution-rules.md
  - docs/adr/009-task-io-design.md
  - docs/adr/011-file-main-node-and-sub-nodes.md
  - docs/adr/012-control-flow-nodes.md
  - docs/adr/015-file-internal-edge-structure.md
  - docs/adr/016-foreach-as-flow-construct.md
  - docs/adr/018-event-node.md
  - docs/adr/019-state-node.md
  - docs/adr/020-cross-edge-management.md
  - docs/adr/023-control-flow-scope-and-branch-entry.md
  - docs/adr/040-control-flow-step-wiring.md
  - docs/adr/060-flow-wiring-type-compatibility.md
  - docs/adr/061-foreach-returns-collected-asset.md
  - docs/adr/062-task-return-source.md
---

# エッジ定義仕様

brewprintのエッジは記述場所によって3種に分かれる。

| 種別 | 記述場所 | 用途 |
|------|---------|------|
| データフローエッジ | `flow:` セクション | Processingレイヤー内のwiring |
| 状態遷移エッジ | `transitions:` セクション | Applicationレイヤーの状態遷移 |
| クロスエッジ | `task` nodeの `reads`/`writes` | task ↔ store間の参照・更新 |

---

## 1. データフローエッジ（flow:セクション）

> 出典: ADR-015, ADR-016, ADR-060, ADR-061

### 設計原則

`nodes:` 内の各ノード定義はsignature（型・名前）のみを持つ。
ノード間のwiring（どの出力がどの入力になるか）はすべて `flow:` セクションに集約する（ADR-015）。

参照: Airflow（`@task` + `>>`）、Prefect（`@task` + `@flow`）、Temporal（activity + workflow）と同じ構造。

### 制御フロースコープ

> 出典: ADR-023

**制御フロー構文（`branch` / `fork` / `foreach`）の内部で生成されたassetは、その構文のスコープ外から直接参照不可。**

例外として、`foreach.returns` で明示的に宣言された collected asset source は外部へ公開され、同一 flow file 内の後続 flow entry から bare source として参照できる。これは制御フロースコープに対する明示的な escape hatch であり、個別 iteration asset を外部参照可能にするものではない。

スコープ外にデータを渡す必要があり、`foreach.returns` の collect 結果では表せない場合は、`initializes` で事前宣言したstoreに `writes` で格納し、後続taskが `reads` で参照する。

```yaml
# NG: 分岐内のassetを外部wiringで直接参照
- step: finalize
  params:
    result: admin_flow     # admin_flowはbranchスコープ内のため参照不可

# OK: storeを介して分岐外にデータを渡す
- id: admin_flow
  type: task
  writes: [role_result_store]

- id: finalize
  type: task
  reads: [role_result_store]
```

収束不要なケース（各パスが独立して終端する）はstoreも不要。その場合、分岐後のtaskはfloatingノードとなりDAGでENDに直行する形でrenderされる（ADR-023）。

### 1-1. stepエントリ（通常のtask）

```yaml
flow:
  - step: fetch_data
    params:
      config: $params.config     # main nodeのparams内の"config"フィールドを参照

  - step: transform
    params:
      raw: fetch_data           # fetch_dataのreturns全体を参照
```

#### stepエントリのフィールド

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `step` | ✓ | 実行するnode ID | ADR-015 |
| `params` | 任意 | 入力のwiring。key=param名, value=参照元 | ADR-015 |

#### wiringの記法

| 記法 | 意味 |
|------|------|
| `source_node` | source_nodeの `returns` 全体を参照 |
| `collected_asset` | 先行する `foreach.returns` で宣言された collected asset source を参照 |
| `$params.field` | ファイル境界からの入力（main nodeのparams）の特定フィールドを参照 |
| `$item` | foreachのループ境界からの入力（現在のイテレーション要素） |

**wiringの単位は常にtaskのreturns全体**（ADR-015）。`source_node.field` 記法は存在しない。node IDはファイル内でユニークなため曖昧性が生じないこと、およびreturns内部への部分アクセスはNG（extract taskで対応）のため。

`$params.field` のフィールド指定は必須。外部taskのparam名とmain nodeのparam名が一致する保証はなく、`$params` bareによる暗黙のname一致解決は使わない（ADR-015）。

```yaml
# NG: returnsの一部フィールドを直接wiring
- step: static_analysis
  params:
    raw: fetch_data.rawa       # returnsの内部フィールドへの直接参照

# OK: extract taskを明示的に挟む
- step: extract_rawa
  params:
    raw: fetch_data
- step: static_analysis
  params:
    raw: extract_rawa
```

### 1-2. forkエントリ（並列実行）

> 出典: ADR-012, ADR-015, ADR-040

```yaml
flow:
  - fork: fan_out
    branches:
      - steps:
          - step: static_analysis
            params:
              raw: fetch_data
      - steps:
          - step: dynamic_analysis
            params:
              raw: fetch_data
      - steps:
          - step: dep_check
            params:
              raw: fetch_data
    join: aggregate
```

#### forkエントリのフィールド

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `fork` | ✓ | fork node ID | ADR-012, ADR-015 |
| `branches` | ✓ | branch object のリスト | ADR-040 |
| `join` | ✓ | 対応するjoin node ID | ADR-015 |

`fork` と `join` は必ずペアで使う。`fork` 単体・`join` 単体は不正（ADR-012）。

#### branch object

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `steps` | ✓ | このbranchで実行するstep objectのリスト | ADR-040 |

`steps[]` の各要素は通常の `flow` の `step` エントリと同じ形式を使う。

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `step` | ✓ | 実行するnode ID | ADR-015, ADR-040 |
| `params` | 任意 | 入力のwiring。key=param名, value=参照元 | ADR-015, ADR-040 |

旧形式の `branches: - [step_a, step_b]` は使用しない。paramsが不要な場合も `steps:` 配下に `step:` を明示する。

```yaml
branches:
  - steps:
      - step: static_analysis
  - steps:
      - step: dynamic_analysis
```

`fork.params` によるbranch内stepへの暗黙伝播は採用しない。branch内taskへの入力は、各 `steps[].params` に明示的に書く（ADR-040）。

#### join.params の解決

`join:` で指定されたjoin nodeの `params` は、各branch終端stepの `returns.name` と同名一致で解決する。

```yaml
nodes:
  - id: static_analysis
    type: task
    returns:
      name: static_result
      model: static_result

  - id: aggregate
    type: join
    params:
      - name: static_result
        model: static_result
```

上記では `static_analysis.returns.name == static_result` が `aggregate.params.static_result` に渡る。
一致するbranch終端stepのreturnsが存在しない場合はparser errorとする。

### 1-4. branchエントリ（排他分岐）

> 出典: ADR-012, ADR-023, ADR-040

```yaml
flow:
  - branch: route_by_role
    params:
      user: fetch_user
    cases:
      - label: admin
        step: admin_flow
        params:
          user: fetch_user
      - label: user
        step: user_flow
        params:
          user: fetch_user
```

#### branchエントリのフィールド

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `branch` | ✓ | branch node ID | ADR-012 |
| `params` | 任意 | branch node自身の分岐判断に使う入力のwiring（stepエントリと同じルール） | ADR-023, ADR-040 |
| `cases` | ✓ | 各パスのエントリポイントのリスト | ADR-023 |

`branch.params` はbranch node自身の判定入力としてのみ扱う。case entry taskへのwiringは `cases[].params` に明示的に書く（ADR-040）。

#### casesエントリのフィールド

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `label` | ✓ | 条件ラベル。人間とLLMへの意味記述。評価はbrewprintのスコープ外 | ADR-023 |
| `step` | ✓ | このケースのエントリポイントとなるtask ID（単一） | ADR-023 |
| `params` | 任意 | case entry taskへの入力wiring。key=param名, value=参照元 | ADR-040 |

`step` は単一のnode IDのみ。エントリポイント以降の後続stepはwiring（`params`参照）から導出されるDAG構造によって決まるため、cases内でのstep列挙は不要。`branch` は1パスしか実行されないためエントリポイントのみで十分（ADR-023）。

case内で生成されたassetは、ADR-023の制御フロースコープによりbranch外から直接参照できない。`cases[].params` はcase entry taskへの入力を明示するためのものであり、branch内部assetの外部参照を許可するものではない。

### 1-5. foreachエントリ（ループ実行）

> 出典: ADR-013（superseded）, ADR-016, ADR-060, ADR-061

`foreach` はnode typeではなく、`flow:` セクションの制御構文（ADR-016）。

```yaml
flow:
  - foreach: process_item       # apply先taskのID
    mode: sequential            # sequential（デフォルト）or map
    over: fetch_items           # iterateするlistの参照元node ID
    params:
      item: $item               # 現在のイテレーション要素
      config: $params.config    # 他のparamも含めapply先taskのparams wiring（任意ではなくapply先に依存）
    returns: results            # apply先taskのreturnsをiterationごとにcollectしたasset source名（任意）
```

#### foreachエントリのフィールド

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `foreach` | ✓ | apply先taskのID（同ファイルのサブノードまたは外部main node） | ADR-016 |
| `over` | ✓ | iterateするlistの参照元。**node ID**（前段task / join の `returns` asset）または **`$params.field`**（main taskのparams）を指定可。TypeRef 解決と `$item` 型導出の詳細は §1-7 を参照 | ADR-016, ADR-060 |
| `mode` | 任意 | `sequential`（デフォルト）or `map`（並列実行） | ADR-016 |
| `params` | 任意 | apply先taskのparams wiring（stepエントリと同じルール）。apply先にparamsがある場合は必須。`$item` で現在のイテレーション要素を参照 | ADR-016 |
| `returns` | 任意 | apply先taskの `returns` を iteration ごとに collect した collected asset source 名。後続 flow から参照する場合に指定する。side-effect only の foreach では省略可 | ADR-016, ADR-061 |

#### modeの使い分け

| mode | 意味 | 実装パターンの例 |
|------|------|----------------|
| `sequential` | 要素を順番に1つずつ処理 | 通常のforループ |
| `map` | 全要素を並列処理 | ProcessPoolExecutor / asyncio.gather |

#### $item シジル

`$item` は `foreach.params` 内だけで有効な wiring source である。foreach外で `$item` を使った場合は `invalid_wiring_source` を出す。

`$item` の型は `over` で参照した TypeRef から導出する。詳細な導出ルールと診断抑制は §1-7 を参照する。

#### foreach.returns collected asset source

`foreach.returns` は、apply 先 task の `returns` を iteration ごとに collect した collected asset source 名である。

```yaml
flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items

  - step: summarize_cart
    params:
      items: validated_items
```

上記の `validated_items` は、後続 flow から bare source として参照できる。
この名前は apply 先 task の `returns.name` ではなく、foreach invocation 単位で宣言される file-local source 名である。

`foreach.returns` は optional である。collect 結果を後続 flow から参照する場合は指定し、side-effect only の foreach では省略できる。省略時は collected asset source を semantic model に生成しない。renderer / inspect / MCP は、省略時の internal pseudo source を露出してはならない。

apply 先 task に `returns` がないにもかかわらず `foreach.returns` を指定した場合は `invalid_foreach_returns` とする。また、当該 foreach 自身の `params` 内から自分自身の `returns` 名を参照してはならない。この場合も `invalid_foreach_returns` とする。

apply 先 task の `returns.model` が TypeRef `T` の場合、`foreach.returns` の TypeRef は `list<T>` とする。apply 先 task の `returns.model` が `any` の場合は `list<any>` とする。apply 先 task の `returns.model` が解決不能な場合は collected asset source の TypeRef も解決不能として扱い、後続 wiring の `incompatible_wiring_type` は抑制する。

```yaml
nodes:
  - id: validate_item
    type: task
    returns:
      name: validated
      model: cart_item

flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items
```

上記では `validated_items` の TypeRef は `list<cart_item>` である。

`foreach.id` は導入しない。同じ apply 先 task を複数回 foreach する場合は、異なる `foreach.returns` 名で collect 結果を区別する。

本仕様は task return source を定義しない。`task.returns.source` および main task `returns.name` と `foreach.returns` の名前一致による暗黙接続は採用しない。

DAGレンダリングではforeachはapply先taskのboxに ↻ アイコンを装飾する形で表現。foreachが独立したboxとして描画されることはない（ADR-016）。

> 由来: ADR-061 §1〜§8

### 1-6. $シジル体系まとめ

> 出典: ADR-015, ADR-016

| シジル | 意味 | 利用場所 |
|--------|------|---------|
| `$params.field` | ファイル境界からの入力（main nodeのparams）の特定フィールドを参照。フィールド指定必須 | flow:内のwiring |
| `$item` | ループ境界からの入力（foreachの現在のイテレーション要素） | foreach.params内 |

シジルはnode IDと記法レベルで区別され、「外部からの注入」であることを明示する（ADR-015）。

---

### 1-7. flow wiring 型互換性

> 出典: ADR-060, ADR-061

flow wiring では、source TypeRef から target param TypeRef への代入互換性を検証する。

検証対象は以下。

| wiring箇所 | source | target |
|---|---|---|
| `step.params` | wiring source | step task の params |
| `branch.params` | wiring source | branch node 自身の params |
| `branch.cases[].params` | wiring source | case entry task の params |
| `fork.branches[].steps[].params` | wiring source | branch内 step task の params |
| `foreach.params` | wiring source | foreach apply task の params |
| `join.params` | fork branch terminal step の returns.name 一致による暗黙source | join node の params |

#### TypeRef互換ルール

source TypeRef `S` から target TypeRef `T` への代入は、named list/dict model を container TypeRef に正規化したうえで、以下の場合のみ有効である。

1. `S` または `T` が `any`
2. primitive 同士で同一
3. list/dict 以外の named model 同士で QID が同一
4. list 同士で element TypeRef が互換
5. dict 同士で value TypeRef が互換

それ以外は `incompatible_wiring_type` を出す。

```txt
str -> str                         OK
str -> int                         NG
user -> user                       OK
user -> order                      NG
any -> user                        OK
user -> any                        OK
list<user> -> list<user>           OK
list<user> -> list<order>          NG
list<any> -> list<user>            OK
user_list -> list<user>            OK
config_map -> dict<config>         OK
str -> user                        NG
```

named list/dict model の正規化および TypeRef 構文は [type-ref.md](./type-ref.md) を参照する。

#### wiring source の型解決

| source記法 | source TypeRef |
|---|---|
| node ID / QualifiedID | task または join の `returns.model` |
| `$params.<name>` | 同一ファイルの main task の `params[].name == <name>` の `model` |
| `$item` | foreach の `over` から導出した element TypeRef。`over` が `list<T>` または named list model の場合は `T`、`any` の場合は `any` |
| `foreach.returns` で宣言された collected asset source 名 | apply 先 task の `returns.model` `T` から導出した `list<T>` |

flow wiring source は以下の4種として解決する。

1. node ID / QualifiedID
2. `$params.<name>`
3. `$item`
4. `foreach.returns` で宣言された collected asset source 名

node ID が task / join 以外を指す場合、または task / join でも `returns` を持たない場合は、参照先は存在するが wiring source として使えないため `invalid_wiring_source` を出す。

`$params.<name>` は main task params の名前参照であり、struct field access ではない。

`$item` は `foreach.params` 内だけで有効である。foreach外で使った場合は、参照先の概念は存在するがその文脈では使えないため `invalid_wiring_source` を出す。

`foreach.returns` で宣言された collected asset source は、同一 flow file 内の後続 step / branch / fork / foreach から bare source として参照できる。当該 foreach 自身の `params` 内から自分の `returns` 名を参照した場合は `invalid_foreach_returns` を出す。

`foreach.returns` は同一 flow file 内の bare wiring source 名前空間に参加する。同一 flow file 内で、`foreach.returns` は node id または他の `foreach.returns` と重複してはならない。重複した場合は `duplicate_flow_source` を出す。ただし task の `returns.name` は通常 flow の wiring source ではないため、task `returns.name` と `foreach.returns` が同名でも衝突扱いしない。

wiring source が node ID / `$params.<name>` / `$item` / collected asset source のいずれとしても解決できない場合は `unresolved_wiring_source` を出す。

`foreach.over` の解決結果が list として扱えない場合は `invalid_foreach_over_type` を出し、その foreach 内の `$item` wiring に対する `incompatible_wiring_type` は抑制する。

> 由来: ADR-060 §5, ADR-061 §3〜§6, §9

#### 型解決失敗時の扱い

型互換性チェックは、source TypeRef と target TypeRef の両方が正常に解決できた場合のみ行う。

以下の場合は `incompatible_wiring_type` を発行しない。

- source TypeRef が解決不能
- target param TypeRef が解決不能
- foreach の `over` が list として扱えず、`$item` の型が導出できない
- collected asset source の TypeRef が解決不能

この方針により、未解決参照や TypeRef 構文エラーに対して二重に incompatible diagnostic を出さない。

### 1-8. task return wiring

> 出典: ADR-062

`task.returns.source` は、task が外部へ返す値の source を明示する task return wiring である。§1-7 の flow wiring が task 内部のノード間 wiring を扱うのに対し、return wiring は task の外向き signature と内部 source の接続を扱う。

`returns.name` / `returns.model` は task の外向き signature を表す。`returns.source` は、その signature を満たす値を内部 flow または入力からどこで得るかを指定する。

`returns.source` は optional である。leaf task / note-only task / external boundary task では指定しなくてよい。`flow:` を持つ main task / composite task が内部 flow の結果を返す場合、または入力をそのまま返す pass-through task を明示する場合に指定する。

```yaml
nodes:
  - id: validate_cart
    type: task
    main: true
    params:
      - name: cart_items
        model: cart_item_list
    returns:
      name: validated_items
      model: cart_item_list
      source: validated_items

flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items
```

`returns.source` に指定できる source は以下である。

| source | 意味 |
|---|---|
| node ID / QualifiedID | task / join など returns を持つ node の出力全体 |
| collected asset source | 同一 flow file 内に出現する `foreach.returns` 由来 collected asset |
| `$params.<name>` | main task params の `<name>` をそのまま返す |

`returns.source` から参照可能な collected asset source は、同一 flow file 内に出現するすべての `foreach.returns` 由来 collected asset source である。`returns.source` は task 全体の return を表すため、flow entry 順における前方参照という概念は適用しない。flow 内部 wiring の visibility とは扱いが異なる。

`$item` は `returns.source` では使えない。`$item` は foreach iteration 内部の source であり、task 全体の return source ではない。foreach 全体の結果を返す場合は、`foreach.returns` で collect した source を `returns.source` に指定する。

`returns.source` が node ID / `$params.<name>` / collected asset source のいずれとしても解決できない場合は `unresolved_return_source` を出す。source は解決できたが task return source として使えない場合は `invalid_return_source` を出す。例: returns を持たない node、`$item`。

`returns.source` を指定した場合、source TypeRef と `returns.model` の TypeRef は §1-7 と同じ TypeRef compatibility ルールで検証する。互換しない場合は `incompatible_return_type` を出す。

source TypeRef または `returns.model` の TypeRef が解決不能な場合は、重複して `incompatible_return_type` を発行しない。未解決 source / invalid source / unresolved model / invalid TypeRef など、一次診断を優先する。

`returns.name` と flow source 名が一致していても、それだけでは task return source として扱わない。返す値を明示する場合は `returns.source` を指定する。なお、`join.params` の `returns.name` 一致による暗黙接続は既存仕様として維持する。

`returns.source` は単一 return にのみ対応する。複数 return / named tuple / multi output task は本仕様では扱わない。必要な場合は struct model で wrap して単一 return として表現する。

> 由来: ADR-062 §1〜§8

## 2. 状態遷移エッジ（transitions:セクション）

> 出典: ADR-019

State Diagramファイルに記述。`flow:` はProcessingレイヤーのDAGを指す語として確立しているため、Applicationレイヤーの状態遷移は `transitions:` と別名にする（ADR-019）。

```yaml
nodes:
  - id: idle
    type: state
    initial: true

  - id: loading
    type: state

  - id: login_submitted
    type: event
    source: ui
    payload:
      model: login_form

transitions:
  - from: idle
    on: login_submitted
    to: loading
    action: auth.task.login
    note: "ログインAPIを呼び出す"

  - from: loading
    on: login_succeeded
    to: authenticated

  - from: loading
    on: login_failed
    to: error
    guard: "retryCount < 3"
    note: "リトライ上限未達の場合のみ"
```

### transitionエントリのフィールド

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `from` | ✓ | 遷移元state ID | ADR-019 |
| `on` | ✓ | トリガーとなるevent ID | ADR-019 |
| `to` | ✓ | 遷移先state ID | ADR-019 |
| `action` | 任意 | 遷移に伴い実行するtask ID（フルパスまたはローカルID） | ADR-019 |
| `guard` | 任意 | 遷移条件テキスト（評価はbrewprintのスコープ外） | ADR-019 |
| `note` | 任意 | 補足説明 | ADR-019 |

### actionの設計根拠

`action` は `transition` に置く（eventには置かない）。理由: actionは `(state, event)` の組み合わせで決まるもの。同じeventでも遷移元stateが異なればactionが変わりうる（ADR-019, Mealy machine参照）。

```yaml
transitions:
  - from: idle
    on: login_submitted
    to: loading
    action: auth.task.login        # 通常ログイン

  - from: session_expired
    on: login_submitted
    to: loading
    action: auth.task.reauth       # 再認証（別task）
```

### guardの扱い

`guard` の評価はbrewprintのスコープ外（実装言語依存）。複雑な条件は extract task またはtaskのロジックとして実装すべきであり、YAML上では自然言語テキストのみ記述する（ADR-019）。

### eventの定義場所

State Diagramが扱うeventはそのFSM専用のものが多いため、stateと同じファイルの `nodes:` に同居させる。他のファイル（Sequence Diagram等）から同じeventを参照する場合はクロスファイル参照（ADR-003）で解決する（ADR-019）。

---

## 3. クロスエッジ（task nodeのreads/writes）

> 出典: ADR-020

レイヤーをまたぐエッジは `write` / `read` の2種のみ。`task` nodeのフィールドとして記述する。flow:ステップではなく、node定義に属する（ADR-020）。

```yaml
- id: process_payment
  type: task
  reads: [balance_store]
  writes: [transaction_store, balance_store]
  note: "transactionとbalanceは同一トランザクションで更新"
```

| フィールド | 型 | 内容 | 出典 |
|-----------|-----|------|------|
| `reads` | list\<store-id\> | このtaskが参照するstore IDのリスト | ADR-020 |
| `writes` | list\<store-id\> | このtaskが更新するstore IDのリスト | ADR-020 |

### 設計根拠

- storeの参照・更新はtaskの実装に紐づく性質。どのflowから呼ばれようとも変わらない（ADR-020）
- 複数storeへの reads/writes を許容する。1task = 1store制約は設けない
- 複数storeにまたがるtransaction境界は `note` に自然言語で記述（ADR-008と同じ方針）

### 廃止されたクロスエッジ種別

| kind | 廃止理由 |
|------|---------|
| `trigger` | `transition.action`（ADR-019）で表現済み |
| `reflect` | event + transitionの連鎖で表現済み |
| `hydrate` | event + transitionの連鎖で表現済み |

（ADR-020）

---

## 4. クロスファイル参照

> 出典: ADR-003

同モジュール内はID直書きで解決。モジュールを跨ぐ場合はフルパスを要求する。

```yaml
# 同モジュール内（auth/task/login.yaml内での参照）
flow:
  - step: validate            # auth.task.validate に解決される

# モジュール跨ぎ
transitions:
  - from: idle
    on: login_submitted
    to: loading
    action: auth.task.login   # フルパス参照
```

フルパスのフォーマット: `<module>.<node-type>.<id>`
