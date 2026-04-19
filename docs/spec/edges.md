---
scope: docs/spec/edges.md
status: wip
last_updated: 2026-04-19
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

### 1-1. stepエントリ（通常のtask）

```yaml
flow:
  - step: fetch_data
    params:
      config: $params           # main nodeのparamsから注入

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
| `source_node` | source_nodeの `returns` 全体を参照。型・name一致で解決 |
| `source_node.field` | 曖昧性解消が必要な場合のみ使用 |
| `$params` | ファイル境界からの入力（main nodeのparams）を参照。name一致で解決 |
| `$params.field` | main nodeのparams内の特定フィールドを指定 |
| `$item` | foreachのループ境界からの入力（現在のイテレーション要素） |

**wiringの単位は常にtaskのreturns全体**（ADR-015）。
`.field` 記法は曖昧性解消のためにのみ使う。returnsの一部フィールドを取り出すためには使わない。

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

### 1-3. foreachエントリ（ループ実行）

> 出典: ADR-013（superseded）, ADR-016

`foreach` はnode typeではなく、`flow:` セクションの制御構文（ADR-016）。

```yaml
flow:
  - foreach: process_item       # apply先taskのID
    mode: sequential            # sequential（デフォルト）or map
    over: fetch_items           # iterateするlistの参照元node ID
    params:
      item: $item               # 現在のイテレーション要素
      config: $params           # 補助入力（任意）
    returns: results            # applyの結果をcollectしたasset名
```

#### foreachエントリのフィールド

| フィールド | 必須 | 内容 | 出典 |
|-----------|------|------|------|
| `foreach` | ✓ | apply先taskのID（同ファイルのサブノードまたは外部main node） | ADR-016 |
| `over` | ✓ | iterateするlistの参照元node ID | ADR-016 |
| `mode` | 任意 | `sequential`（デフォルト）or `map`（並列実行） | ADR-016 |
| `params` | 任意 | 各イテレーションに渡すparams。`$item` で現在要素を参照 | ADR-016 |
| `returns` | 任意 | applyの結果をcollectしたasset名 | ADR-016 |

#### modeの使い分け

| mode | 意味 | 実装パターンの例 |
|------|------|----------------|
| `sequential` | 要素を順番に1つずつ処理 | 通常のforループ |
| `map` | 全要素を並列処理 | ProcessPoolExecutor / asyncio.gather |

#### $item シジル

`$item` の型は `over` で参照したlistのelement型から暗黙に決まる（ADR-016）。apply先taskのparamのmodelとの型一致はGo実装の静的検証で担保する。

DAGレンダリングではforeachはapply先taskのboxに ↻ アイコンを装飾する形で表現。foreachが独立したboxとして描画されることはない（ADR-016）。

### 1-4. $シジル体系まとめ

> 出典: ADR-015, ADR-016

| シジル | 意味 | 利用場所 |
|--------|------|---------|
| `$params` | ファイル境界からの入力（main nodeのparams） | flow:内のwiring |
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
