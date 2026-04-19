---
scope: docs/spec/edges.md
status: wip
last_updated: 2026-04-20
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

> 出典: ADR-015, ADR-016

### 設計原則

`nodes:` 内の各ノード定義はsignature（型・名前）のみを持つ。
ノード間のwiring（どの出力がどの入力になるか）はすべて `flow:` セクションに集約する（ADR-015）。

参照: Airflow（`@task` + `>>`）、Prefect（`@task` + `@flow`）、Temporal（activity + workflow）と同じ構造。

### 制御フロースコープ

> 出典: ADR-023

**制御フロー構文（`branch` / `fork` / `foreach`）の内部で生成されたassetは、その構文のスコープ外から直接参照不可。**

スコープ外にデータを渡す必要がある場合は、`initializes` で事前宣言したstoreに `writes` で格納し、後続taskが `reads` で参照する。

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

> 出典: ADR-012, ADR-015

```yaml
flow:
  - fork: fan_out
    branches:
      - [static_analysis]
      - [dynamic_analysis]
      - [dep_check]
    join: aggregate
    params:
      request: fetch_data       # 各branchへの共通入力
```

#### forkエントリのフィールド

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `fork` | ✓ | fork node ID | ADR-012, ADR-015 |
| `branches` | ✓ | 各ブランチのstep列をリストで記述 | ADR-015 |
| `join` | ✓ | 対応するjoin node ID | ADR-015 |
| `params` | 任意 | 各ブランチへの共通入力のwiring | ADR-015 |

`fork` と `join` は必ずペアで使う。`fork` 単体・`join` 単体は不正（ADR-012）。

### 1-4. branchエントリ（排他分岐）

> 出典: ADR-012, ADR-023

```yaml
flow:
  - branch: route_by_role
    params:
      user: fetch_user
    cases:
      - label: admin
        step: admin_flow
      - label: user
        step: user_flow
```

#### branchエントリのフィールド

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `branch` | ✓ | branch node ID | ADR-012 |
| `params` | 任意 | 分岐判断に使う入力のwiring（stepエントリと同じルール） | ADR-023 |
| `cases` | ✓ | 各パスのエントリポイントのリスト | ADR-023 |

#### casesエントリのフィールド

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `label` | ✓ | 条件ラベル。人間とLLMへの意味記述。評価はbrewprintのスコープ外 | ADR-023 |
| `step` | ✓ | このケースのエントリポイントとなるtask ID（単一） | ADR-023 |

`step` は単一のnode IDのみ。エントリポイント以降の後続stepはwiring（`params`参照）から導出されるDAG構造によって決まるため、cases内での列挙は不要。`fork` の `branches` がstep列を持つのは並列ブランチの帰属を示す必要があるため。`branch` は1パスしか実行されないためエントリポイントのみで十分（ADR-023）。

### 1-5. foreachエントリ（ループ実行）

> 出典: ADR-013（superseded）, ADR-016

`foreach` はnode typeではなく、`flow:` セクションの制御構文（ADR-016）。

```yaml
flow:
  - foreach: process_item       # apply先taskのID
    mode: sequential            # sequential（デフォルト）or map
    over: fetch_items           # iterateするlistの参照元node ID
    params:
      item: $item               # 現在のイテレーション要素
      config: $params.config    # 他のparamも含めapply先taskのparams wiring（任意ではなくapply先に依存）
    returns: results            # applyの結果をcollectしたasset名
```

#### foreachエントリのフィールド

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `foreach` | ✓ | apply先taskのID（同ファイルのサブノードまたは外部main node） | ADR-016 |
| `over` | ✓ | iterateするlistの参照元node ID | ADR-016 |
| `mode` | 任意 | `sequential`（デフォルト）or `map`（並列実行） | ADR-016 |
| `params` | 任意 | apply先taskのparams wiring（stepエントリと同じルール）。apply先にparamsがある場合は必須。`$item` で現在のイテレーション要素を参照 | ADR-016 |
| `returns` | 任意 | applyの結果をcollectしたasset名 | ADR-016 |

#### modeの使い分け

| mode | 意味 | 実装パターンの例 |
|------|------|----------------|
| `sequential` | 要素を順番に1つずつ処理 | 通常のforループ |
| `map` | 全要素を並列処理 | ProcessPoolExecutor / asyncio.gather |

#### $item シジル

`$item` の型は `over` で参照したlistのelement型から暗黙に決まる（ADR-016）。apply先taskのparamのmodelとの型一致はGo実装の静的検証で担保する。

DAGレンダリングではforeachはapply先taskのboxに ↻ アイコンを装飾する形で表現。foreachが独立したboxとして描画されることはない（ADR-016）。

### 1-6. $シジル体系まとめ

> 出典: ADR-015, ADR-016

| シジル | 意味 | 利用場所 |
|--------|------|---------|
| `$params.field` | ファイル境界からの入力（main nodeのparams）の特定フィールドを参照。フィールド指定必須 | flow:内のwiring |
| `$item` | ループ境界からの入力（foreachの現在のイテレーション要素） | foreach.params内 |

シジルはnode IDと記法レベルで区別され、「外部からの注入」であることを明示する（ADR-015）。

---

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
