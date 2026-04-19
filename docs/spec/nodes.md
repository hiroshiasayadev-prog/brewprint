---
scope: docs/spec/nodes.md
status: wip
last_updated: 2026-04-19
summary: >
  brewprintの全ノード種別のフィールド定義。
  各フィールドの必須/任意・型・意味・出典ADRを記載する。
  未定項目は ⚠️ 未定 として明示し、次sessionで詳細化する。
depends_on:
  - docs/adr/001-node-type-splitting.md
  - docs/adr/006-node-type-renaming.md
  - docs/adr/007-asset-store-boundary.md
  - docs/adr/008-dag-dataflow-and-field-validation.md
  - docs/adr/009-task-io-design.md
  - docs/adr/010-ca-enforcement-directory-structure-model-asset-split.md
  - docs/adr/011-file-main-node-and-sub-nodes.md
  - docs/adr/012-control-flow-nodes.md
  - docs/adr/014-initializes-field.md
  - docs/adr/016-foreach-as-flow-construct.md
  - docs/adr/017-diagram-layers-and-scope.md
  - docs/adr/018-event-node.md
  - docs/adr/019-state-node.md
  - docs/adr/020-cross-edge-management.md
---

# ノード定義仕様

## ファイル構造

1ファイル = 1メインノード（`main: true`）+ 0個以上のサブノード（ADR-011）。

```yaml
nodes:
  - id: <string>
    type: <node-type>
    main: true          # メインノードのみ。1ファイルに1つ
    ...

  - id: <sub-node>      # サブノード（ファイル内private）
    type: <node-type>
    ...

flow:                   # Processingレイヤーのwiring（ADR-015）
  - step: ...

transitions:            # Applicationレイヤーの状態遷移（ADR-019）
  - from: ...
```

サブノードは外部モジュールから参照不可。外部からアクセスできるのはメインノードのIDのみ（ADR-011）。

---

## 共通フィールド

全ノード種別に共通するフィールド。

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `id` | ✓ | string | モジュール内でユニークな識別子 |
| `type` | ✓ | enum | ノード種別（後述） |
| `main` | 任意 | bool | `true` でメインノード宣言。1ファイルに1つだけ許容（ADR-011） |
| `note` | 任意 | string | 人間向けdocstring兼LLMへのsemantic contract（ADR-008） |

---

## task

> 出典: ADR-009, ADR-010, ADR-011, ADR-014, ADR-017, ADR-020

Processingレイヤー。処理の単位。`returns` 宣言によってDAG上に `asset` ノードを暗黙的に生成する。

```yaml
- id: login
  type: task
  main: true
  endpoint: true
  method: POST
  path: /auth/login
  params:
    - name: credentials
      model: credential
  returns:
    name: auth_token
    model: token
  reads: [session_store]
  writes: [session_store]
  initializes:
    - name: login_log
      model: login_log
      note: "空のlogin_logで初期化"
  note: "認証情報を検証しトークンを発行する"
```

### フィールド定義

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| `params` | 任意 | list\<param\> | 入力。model IDへの参照リスト | ADR-009 |
| `returns` | 任意 | returns | 出力。model IDへの参照とasset名の宣言 | ADR-009 |
| `reads` | 任意 | list\<store-id\> | 参照するstore IDのリスト | ADR-020 |
| `writes` | 任意 | list\<store-id\> | 更新するstore IDのリスト | ADR-020 |
| `endpoint` | 任意 | bool | `true` で `list_endpoints` MCPツールに出力 | ADR-005, ADR-017 |
| `method` | endpoint時必須 | enum | HTTP method（GET/POST/PUT/DELETE/PATCH） | ADR-005 |
| `path` | endpoint時必須 | string | URLパス（例: `/auth/login`） | ADR-005 |
| `initializes` | 任意 | list\<init\> | このファイル内で使うstoreの初期宣言（main nodeのみ）| ADR-014 |

### param オブジェクト

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `name` | ✓ | string | このtask内でのパラメータ名 |
| `model` | ✓ | model-id | 型参照（model IDまたはprimitive） |

### returns オブジェクト

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `name` | ✓ | string | 生成されるassetの名前 |
| `model` | ✓ | model-id | 型参照（model ID） |

`returns` は単一のみ。複数返しが必要な場合はstruct modelでwrapして単一にする（ADR-009）。

### init オブジェクト（initializes内）

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `name` | ✓ | string | ファイル内でのstore参照名 |
| `model` | ✓ | model-id | storeの型（model ID） |
| `note` | 任意 | string | 初期値・初期化方法の自然言語記述 |

`initializes` で宣言されたstoreはファイル内にprivate。外部参照不可（ADR-014）。

---

## model

> 出典: ADR-007（superseded・内容はADR-010に継承）, ADR-008, ADR-010

Dataレイヤー。型定義に徹する。DAGには登場しない。`model/` サブディレクトリに1ファイル=1定義で置く。

```yaml
# struct
- id: user
  type: model
  kind: struct
  fields:
    - name: id
      type: str
      comment: "ユーザーID"
    - name: email
      type: str
      comment: "メールアドレス"
    - name: profile
      type: user_profile    # 別model IDを型として参照
      comment: "プロフィール情報"
```

### フィールド定義

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| `kind` | ✓ | enum | `scalar` / `struct` / `list` / `dict` | ADR-007→ADR-010 |
| `fields` | struct時必須 | list\<field\> | フィールド定義（structのみ） | ADR-008 |

### field オブジェクト（struct内）

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| `name` | ✓ | string | フィールド名。struct内でユニーク | ADR-008 |
| `type` | ✓ | string | 型。primitiveリテラル or model ID | ADR-008 |
| `comment` | 任意 | string | 人間向けdocstring兼LLM semantic contract | ADR-008 |

`type` の機械的validationの対象はprimitive or 定義済みmodel IDの存在チェックのみ（ADR-008）。

### kindごとの追加フィールド

> ⚠️ 未定: `scalar` / `list` / `dict` の具体的なフィールド構造は次sessionで確定する。
> 出典: ADR-007（superseded）にて `element: <model-id>` 記法の言及あり。ADR-010への継承後の正式構造は未定。

| kind | 追加フィールド（暫定） | 備考 |
|------|----------------------|------|
| `scalar` | `base: <primitive>` ? | primitiveの型エイリアス。フィールド名未定 |
| `list` | `element: <model-id>` | ADR-007由来。継続有効か未確認 |
| `dict` | `key: ?` / `value: <model-id>` ? | 未定 |

---

## asset

> 出典: ADR-010

Processingレイヤー。フロー上の存在。**独立ファイルは持たない**。`task` の `returns` 宣言から暗黙的に生まれる。

YAML上に直接定義する構文は存在しない。DAG図・MCPツールでは `task.returns.name` をasset IDとして扱う。

```
# task/login.yamlのreturnsから暗黙的に生まれるasset
asset ID: auth_token    # task.returns.name の値
model:   token          # task.returns.model の値
```

assetの「意味」はそれを生産するtaskによって与えられる。task定義の外に独立する情報がないため独立ファイルを持たない（ADR-010）。

---

## store

> 出典: ADR-007（kind定義・内容はADR-010に継承）, ADR-019（stateとの区別）

Processingレイヤー / Dataレイヤー。実行時にデータを保持する実体。FSMの `state` とは別概念（ADR-019）。

```yaml
- id: user_db
  type: store
  kind: db
  of: user
  note: "ユーザーテーブル"

- id: session_store
  type: store
  kind: session
  of: session
  note: "HTTPセッション"

- id: user_collection
  type: store
  kind: collection
  of: user
  note: |
    - find_by_email: emailが一致するもの
    - active_users: is_active = trueのもの
```

### フィールド定義

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| `kind` | ✓ | enum | `db` / `session` / `collection` / `context` | ADR-007→ADR-010 |
| `of` | 任意 | model-id | 保持するmodel ID | ADR-007 |

> ⚠️ 未定: `store.kind: db` がER図のためにフィールド（列）定義を独自に持つかどうか未確定。
> 現状の設計では `of: <model-id>` でmodelを参照する形のみ。
> ERに列レベルの詳細が必要な場合はmodel側に持つか、store側に持つかを次sessionで決定する。

### collectionのnote規約

| noteに書くもの | noteに書かないもの |
|--------------|-----------------|
| 単一collectionのフィールドへのフィルタ・検索 | 複数storeにまたがる結合 |
| 言葉で書ける程度の条件（等値・範囲・真偽フラグ） | 集計・変換・ソートを伴う複雑なクエリ |

collectionのnoteはLLMへのsemantic contract（ADR-007, ADR-008と同じ位置づけ）。

---

## actor

> 出典: ADR-004, ADR-006

Applicationレイヤー。人間・外部システム。Sequence Diagramのエントリーポイントとして使用。UML標準Actorに対応。

```yaml
- id: end_user
  type: actor
  note: "サービスを利用するエンドユーザー"

- id: payment_service
  type: actor
  note: "Stripeなど外部決済サービス"
```

### フィールド定義

> ⚠️ 未定: `actor` 固有フィールドの要否は次sessionで確定する。
> 現状はid/type/noteのみで十分と想定するが、`external: true` やシステム種別フラグが必要かは議論の余地あり。

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| （固有フィールドなし） | — | — | id/type/noteのみで表現 | ADR-004 |

---

## event

> 出典: ADR-017, ADR-018

Applicationレイヤー。制御フローの起点。DAGの `flow:` には登場しない。Sequence DiagramおよびState Diagramで使用。

```yaml
- id: login_submitted
  type: event
  source: ui
  payload:
    model: login_form
  note: "ログインフォームのsubmit"

- id: daily_batch_triggered
  type: event
  source: time
  note: "毎日0時に発火"

- id: payment_webhook_received
  type: event
  source: external
  payload:
    model: payment_event
  note: "Stripeからの決済完了通知"

- id: connection_status_changed
  type: event
  source: er
  watches: db_connection_store
  payload:
    model: connection_status
  note: "db_connection_storeのstatusが変化した時"
```

### フィールド定義

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| `source` | ✓ | enum | `ui` / `time` / `external` / `er` | ADR-018 |
| `payload` | 任意 | payload | イベントが運ぶデータのmodel参照 | ADR-018 |
| `watches` | `er`のみ必須 | store-id | 変化を監視するstore ID | ADR-018 |

### payload オブジェクト

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `model` | ✓ | model-id | ペイロードの型参照 |

### sourceの意味

| source | 意味 | payloadの要否 |
|--------|------|-------------|
| `ui` | ユーザー操作（クリック・フォーム送信等） | 任意（フォームデータ等） |
| `time` | cron・スケジュール実行 | 基本不要 |
| `external` | webhook・message queue・WebSocket | 任意（受信データ） |
| `er` | storeの値変化による発火。`watches`必須 | 任意 |

---

## state

> 出典: ADR-017, ADR-019

Applicationレイヤー。FSMの状態ノード。State Diagramで使用。`store`（データ保持）とは別概念（ADR-019）。

```yaml
- id: idle
  type: state
  initial: true
  note: "ユーザーが操作していない状態"

- id: loading
  type: state
  note: "認証リクエスト処理中"

- id: authenticated
  type: state
  final: true

- id: error
  type: state
  final: true
  note: "認証エラー状態"
```

### フィールド定義

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| `initial` | 任意 | bool | `true` でFSMの初期状態。1ファイルに1つ | ADR-019 |
| `final` | 任意 | bool | `true` で終端状態（複数可） | ADR-019 |

`state` と `store` の区別:

| ノード | レイヤー | 概念 |
|--------|---------|------|
| `store` | Processing / Data | 実行時にデータを保持するインスタンス。DAGでreads/writesされる |
| `state` | Application | FSMが取り得る状態の定義。遷移の起点・終点として使われる |

---

## branch

> 出典: ADR-012

Processingレイヤー。排他分岐。条件に応じて後続パスを**1本だけ**選ぶ。

```yaml
- id: route_by_role
  type: branch
  params:
    - name: user
      model: user
  note: "user.roleに応じてadmin_flow / user_flowのどちらかに進む"
```

### フィールド定義

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| `params` | 任意 | list\<param\> | 分岐判断に使う入力 | ADR-012 |

合流点は明示しない。後続taskがedge構造上で複数パスから受け取る形になる場合、それを暗黙の合流点として読む（ADR-012）。

---

## fork

> 出典: ADR-012

Processingレイヤー。並列分岐。後続パスを**すべて並列実行**する。必ず `join` とペアで使う。

```yaml
- id: fan_out
  type: fork
  params:
    - name: request
      model: analysis_request
  note: "静的解析・動的解析・依存チェックを並列実行する"
```

### フィールド定義

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| `params` | 任意 | list\<param\> | 各ブランチへの共通入力 | ADR-012 |

`fork` 単体での使用は不正。対応する `join` が必須（ADR-012）。wiringはflow:セクションで記述（ADR-015）。

---

## join

> 出典: ADR-012

Processingレイヤー。合流。対応する `fork` の全ブランチが揃うまで待ち、結果を統合する。必ず `fork` とペアで使う。

```yaml
- id: aggregate
  type: join
  params:
    - name: static_result
      model: static_result
    - name: dynamic_result
      model: dynamic_result
    - name: dep_result
      model: dep_result
  returns:
    name: full_report
    model: full_report
  note: "3ブランチの結果を結合してfull_reportを生成する"
```

### フィールド定義

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| `params` | 任意 | list\<param\> | 各ブランチからの入力。全ブランチ分を列挙する | ADR-012 |
| `returns` | 任意 | returns | 統合結果のasset宣言 | ADR-009, ADR-012 |

`join` 単体での使用は不正。対応する `fork` が必須（ADR-012）。

---

## ノード種別一覧

| 種別 | レイヤー | ファイル配置 | DAGに登場 |
|------|---------|------------|---------|
| `task` | Processing | `task/*.yaml` | ✅ |
| `model` | Data | `model/*.yaml` | ❌ |
| `asset` | Processing | なし（task.returnsから導出） | ✅ |
| `store` | Processing / Data | `store/*.yaml` | ✅（reads/writes参照） |
| `actor` | Application | （未定）| ❌ |
| `event` | Application | state.yaml等に同居 | ❌ |
| `state` | Application | `state.yaml` | ❌ |
| `branch` | Processing | `task/*.yaml`のサブノード等 | ✅ |
| `fork` | Processing | `task/*.yaml`のサブノード等 | ✅ |
| `join` | Processing | `task/*.yaml`のサブノード等 | ✅ |

> `foreach` はADR-016にてnode typeから廃止。`flow:` セクションの制御構文として記述する（→ `spec/edges.md`）。
