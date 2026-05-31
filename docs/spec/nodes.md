---
scope: docs/spec/nodes.md
status: confirmed
last_updated: 2026-05-31
summary: >
  brewprintの全ノード種別のフィールド定義。
  各フィールドの必須/任意・型・意味・出典ADRを記載する。
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
  - docs/adr/021-model-field-structure.md
  - docs/adr/023-control-flow-scope-and-branch-entry.md
  - docs/adr/026-fk-cardinality-and-nm-relation.md
  - docs/adr/028-api-table-route-composition.md
  - docs/adr/031-actor-global-definition.md
  - docs/adr/034-internal-event-source.md
  - docs/adr/040-control-flow-step-wiring.md
  - docs/adr/042-wireframe-main-and-layout.md
  - docs/adr/058-subnode-file-private-scope-enforcement.md
  - docs/adr/060-flow-wiring-type-compatibility.md
  - docs/adr/062-task-return-source.md
  - docs/adr/063-task-return-source-initialized-store.md
  - docs/adr/065-asset-immutability-and-edge-role-contrast.md
  - docs/adr/067-enum-model.md
  - docs/adr/070-model-visibility-file-private-helper-model.md
  - docs/adr/071-file-private-helper-model-render-exposure.md
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

サブノードは file-private であり、外部モジュールまたは別ファイルから参照不可。外部からアクセスできるのはメインノードのIDのみ（ADR-011, ADR-058）。

メインノードの `id` は module / node type と組み合わせて public QualifiedID を構成し、project 全体の QualifiedID 一意性制約の対象となる。サブノードの `id` は同一 file 内の local ID であり、同一 file 内でのみ一意であればよい。別 file に存在する同名サブノードとは衝突しない。

flow / reads / writes 等に書かれた bare ID は、まず同一 file 内のサブノードまたは file-private source を優先して解決し、該当がない場合のみ同一 module のメインノードへフォールバックする。サブノードは public QualifiedID を持たず、外部 YAML から QualifiedID で参照できる対象ではない。

Task-file helper minimum では、task file 内に `type: model` の file-private helper model を定義できる。詳細な semantics は [model](#model) §Task-file private helper model semantics を参照する。

Model file 内 helper model、model file render、model catalog、tagged union、DAG TypeRef hint、MCP helper exposure schema はこの minimum の対象外とする。

---

## 共通フィールド

全ノード種別に共通するフィールド。

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `id` | ✓ | string | メインノードでは module / node type 内の public ID。サブノードでは同一 file 内で一意な local ID |
| `type` | ✓ | enum | ノード種別（後述） |
| `note` | 任意 | string | 人間向けdocstring兼LLMへのsemantic contract（ADR-008） |

---

## task

> 出典: ADR-009, ADR-010, ADR-011, ADR-014, ADR-017, ADR-020, ADR-070, ADR-071

Processingレイヤー。処理の単位。`returns` 宣言によってDAG上に `asset` ノードを暗黙的に生成する。

```yaml
- id: login
  type: task
  main: true
  endpoint: true
  method: POST
  path: login
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
| `main` | 任意 | bool | `true` でメインノード宣言。1ファイルに1つだけ許容。`task` にのみ適用（ADR-011） |
| `params` | 任意 | list\<param\> | 入力。TypeRefへの参照リスト | ADR-009, ADR-060 |
| `returns` | 任意 | returns | 出力。TypeRefへの参照、asset名の宣言、任意のreturn source wiring | ADR-009, ADR-060, ADR-062 |
| `reads` | 任意 | list\<store-id\> | 参照するstore IDのリスト | ADR-020 |
| `writes` | 任意 | list\<store-id\> | 更新するstore IDのリスト | ADR-020 |
| `endpoint` | 任意 | bool | `true` でAPI Tableの集計対象となり、`list_endpoints` MCPツールに出力 | ADR-005, ADR-017, ADR-028 |
| `method` | endpoint時必須 | enum | HTTP method（GET/POST/PUT/DELETE/PATCH） | ADR-005 |
| `path` | 任意 | string | endpointのleaf path（例: `login`）。省略時は `task.id` をleaf nameとして使う。`/` を含む複数セグメント不可（single segment限定）。full pathはAPI Table viewの `http_root_path` とmodule階層から構成する | ADR-005, ADR-028 |
| `initializes` | 任意 | list\<init\> | このファイル内で使うstoreの初期宣言（main nodeのみ）| ADR-014 |
| task-file helper model | 任意 | `type: model` node | 同一 task file 内の private schema 定義。詳細は [model](#model) §Task-file private helper model semantics | ADR-070, ADR-071 |

### param オブジェクト

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `name` | ✓ | string | このtask内でのパラメータ名 |
| `model` | ✓ | TypeRef | 型参照。primitive / named model / inline `list<T>` / inline `dict<T>` を指定できる |

### returns オブジェクト

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `name` | ✓ | string | 生成されるassetの名前 |
| `model` | ✓ | TypeRef | 型参照。primitive / named model / inline `list<T>` / inline `dict<T>` を指定できる |
| `source` | 任意 | string（wiring source 構文） | task が外部へ返す値の source。node id / collected asset source / initialized source / `$params.<name>` を指定できる。`$item` は指定不可 |

`returns` は単一のみ。複数返しが必要な場合はstruct modelでwrapして単一にする（ADR-009）。primitive return は `str` / `int` / `bool` / `any` 等の primitive TypeRef で表現する。

`returns.source` は task return wiring であり、`returns.name` / `returns.model` が表す外向き signature を満たす値を内部 flow または入力からどこで得るかを明示する。leaf task / note-only task / external boundary task では省略できる。`returns.source` を省略した場合、`returns.name` と flow source 名が一致していても暗黙接続は行わない。source 解決と型互換性の詳細は [edges.md](./edges.md) §1-8 を参照する。

> 由来: ADR-062 §1〜§7, ADR-063 §1

### init オブジェクト（initializes内）

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `name` | ✓ | string | ファイル内でのstore参照名 |
| `model` | ✓ | model-id | storeの型（model ID） |
| `note` | 任意 | string | 初期値・初期化方法の自然言語記述 |

`initializes` で宣言されたstoreはファイル内にprivate。外部参照不可（ADR-014）。

`initializes[].model` は v1.1 でも model-id 参照のまま扱う。TypeRef 適用対象は `params[].model` / `returns.model` / `fields[].type` / `model.element` / `model.value` に限定する。詳細は [type-ref.md](./type-ref.md) §1 を参照する。

`initializes[].name` は同一 file の bare wiring source 名前空間に参加する。`returns.source` および flow 内部 wiring（step.params / branch.params / fork.branches[].steps[].params / foreach.params / branch.cases[].params）の bare token から initialized source として参照できる。参照時の TypeRef は `initializes[].model` を named model TypeRef として扱い、ADR-060 の TypeRef compatibility ルールで target TypeRef と互換性検証する。詳細は [edges.md](./edges.md) §1-7 / §1-8 を参照する。

外部 file からの QualifiedID 参照や module 跨ぎ参照は ADR-014 通り引き続き不可である。

> 由来: ADR-014 §決定, ADR-063 §2〜§4

---

## model

> 出典: ADR-007（superseded・内容はADR-010に継承）, ADR-008, ADR-010, ADR-021, ADR-060, ADR-067

Dataレイヤー。型定義に徹する。DAGには登場しない。

Public model は `model/` サブディレクトリに置く module-level schema 定義である。Task-file helper minimum では、task file 内に `main: true` を持たない `type: model` node を file-private helper model として置ける。task-file helper model は、同一 YAML file 内の `params[].model` / `returns.model` / helper model の `fields[].type` / `element` / `value` から bare TypeRef で参照するための task-local schema であり、public model file ではない。

### Task-file private helper model semantics

Task-file helper model の visibility / identity / reference scope は以下の通り。

| property | rule |
|---|---|
| visibility | file-private。YAML field として `visibility:` は導入しない |
| public identity | public QualifiedID を持たない |
| local identity | defining task file 内の local model id |
| reference scope | 同一 YAML file 内の TypeRef からのみ bare id で参照できる |
| external reference | 外部 file / module から参照不可。QualifiedID で参照する構文もない |
| render role | DAG Mermaid 本体には描画しない。task-file render の `## Private models` に表示されうる |

同一 task file 内の helper model id は、他の file-private local node id と同じく同一 file 内で一意でなければならない。さらに、task-file helper model は同一 module 内の public model と同じ id を持ってはならない。これにより bare TypeRef の読み手が public model と file-private helper model を文脈依存で取り違えることを防ぐ。

Model file 内 helper model、model file render、model catalog、および UC-002 model response helper-shape migration は WORK-DATA-003 の範囲であり、この task-file helper minimum では定義しない。

```yaml
# struct
- id: user
  type: model
  kind: struct
  fields:
    - name: id
      type: str
      pk: true
      note: "ユーザーID（PK）"
    - name: email
      type: str
      note: "メールアドレス"
    - name: role_id
      type: str
      fk: role.id
      note: "ロールID（FK → role.id）"
    - name: profile
      type: user_profile    # named model TypeRef（fkなし → JSON埋め込み扱い）
      note: "プロフィール情報"
    - name: created_at
      type: datetime
      note: "レコード作成日時（DB管理）"

# list
- id: item_list
  type: model
  kind: list
  element: item             # TypeRef。model ID / primitive / inline list<T> / inline dict<T>

# dict
- id: config_map
  type: model
  kind: dict
  value: config             # TypeRef。keyは常にstr

# enum
- id: mcp_diagnostic_severity
  type: model
  kind: enum
  values:
    - error
    - warning
    - info
    - hint
```


### primitive予約語

以下の語はprimitive予約語。model IDとして定義不可（ADR-021）。TypeRef構文全体は [type-ref.md](./type-ref.md) で定義する。

| primitive | 意味 |
|-----------|------|
| `str` | 文字列 |
| `int` | 整数 |
| `float` | 浮動小数点数 |
| `bool` | 真偽値 |
| `bytes` | バイト列 |
| `datetime` | 日時 |
| `any` | 型不定（使用は最小限に） |

### フィールド定義

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| `kind` | ✓ | enum | `struct` / `list` / `dict` / `enum` | ADR-007→ADR-010, ADR-021, ADR-067 |
| `fields` | struct時必須 | list\<field\> | フィールド定義（structのみ） | ADR-008 |
| `element` | list時必須 | TypeRef | 要素の型 | ADR-021, ADR-060 |
| `value` | dict時必須 | TypeRef | 値の型。keyは常に `str` | ADR-021, ADR-060 |
| `values` | enum時必須 | list\<string\> | enum の許容値集合。non-empty、空文字不可、同一 enum model 内で重複不可 | ADR-067 |

`kind: scalar` は廃止。primitive literalを直接使う（ADR-021）。  
`kind: dict` のkeyは常に `str`。`key` フィールドは存在しない（ADR-021）。

`kind: enum` は string-valued finite vocabulary を表す model kind である。enum model は `values` を必須とし、`fields` / `element` / `value` を持たない。使用側は既存 TypeRef の named model 参照として enum model ID を指定する。inline enum TypeRef は存在しない。

`values` の順序は表示・schema 生成時の順序として保持してよいが、型互換性の意味には使わない。値ごとの `note` / `label` / `deprecated` 等の metadata は v1.1 では導入しない。

> 由来: ADR-067 §1〜§5

### field オブジェクト（struct内）

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| `name` | ✓ | string | フィールド名。struct内でユニーク | ADR-008 |
| `type` | ✓ | TypeRef | 型。primitive / named model / inline `list<T>` / inline `dict<T>` を指定できる | ADR-008, ADR-060 |
| `pk` | 任意 | bool | `true` でPKカラム。1 struct内に1つ | ADR-021 |
| `fk` | 任意 | `<model-id>.<field-name>` | FK参照先。省略時はJSON埋め込み扱い | ADR-021 |
| `unique` | 任意 | bool | `true` で1:1リレーション。`fk:` と併用。省略時はmany-to-one | ADR-026 |
| `note` | 任意 | string | 人間向けdocstring兼LLM semantic contract | ADR-008, ADR-021 |

`type` の機械的validationは TypeRef の構文チェック、および primitive予約語 or 定義済みmodel IDの存在チェックを行う。inline `list<T>` / `dict<T>` は要素・値の TypeRef を再帰的に検証する。TypeRef 構文自体が不正な場合は `invalid_type_ref`、構文はvalidだが named model が未解決の場合は `unresolved_field_type` を出す。

Task-file helper model の `fields[].type` / `element` / `value` も TypeRef として解決する。同一 file 内の helper model を bare TypeRef で参照できるが、外部 file から helper model を参照することはできない。詳細な解決順序は [type-ref.md](./type-ref.md) §4 を参照する。

**`fk` フィールドの意味**

| 記法 | DB上の扱い |
|------|-----------|
| `type: str, fk: role.id` | FKカラム。`role` modelの `id` フィールドを参照 |
| `type: user_profile`（fkなし） | JSON埋め込みカラム |
| `type: tag_list`（list kindのmodel） | variant/JSONカラム（実装依存） |

---

## asset

> 出典: ADR-010, ADR-065

Processingレイヤー。フロー上の存在。**独立ファイルは持たない**。`task` の `returns` 宣言から暗黙的に生まれる。

YAML上に直接定義する構文は存在しない。DAG図・MCPツールでは `task.returns.name` をasset IDとして扱う。

### immutability と書き込み文法の不在

asset は task の単一実行が生んだ output snapshot として扱う。brewprint の語彙レベルで以下が成り立つ。

- cross-edge `reads:` / `writes:` は `store-id` のみを受ける（[edges.md](./edges.md) §3）。asset はこの対象外
- task が値を生むのは `returns` 宣言を通してのみ。生成後の書き換えを表す YAML 記述は存在しない
- 「累積する箱」「後から append される対象」として asset を扱うことはできない

mutable な runtime instance が必要な場合は store を使う（§store / 同セクションの2形態表参照）。

実装言語上の参照渡しや mutation 挙動は brewprint のスコープ外（ADR-065 §決定 §4）。

> 由来: ADR-065 §決定 §1〜§5

```
# task/login.yamlのreturnsから暗黙的に生まれるasset
asset ID: auth_token    # task.returns.name の値
model:   token          # task.returns.model の値
```

assetの「意味」はそれを生産するtaskによって与えられる。task定義の外に独立する情報がないため独立ファイルを持たない（ADR-010）。

---

## store

> 出典: ADR-007（kind定義・内容はADR-010に継承）, ADR-019（stateとの区別）, ADR-065（2形態の整理）

Processingレイヤー / Dataレイヤー。実行時にデータを保持する実体。FSMの `state` とは別概念（ADR-019）。

brewprint における runtime data instance の総称が store であり、宣言経路で2形態に分かれる。

| 形態 | 宣言経路 | スコープ | 出典 |
|------|---------|---------|------|
| store node | `store/*.yaml` で `type: store` として宣言 | module-level（QualifiedID で外部参照可） | ADR-007 → ADR-010 |
| initialized store | task ファイルの `initializes[]` で宣言 | file-private（同一 file 内のみ参照可） | ADR-014, ADR-063 |

両者とも cross-edge `reads:` / `writes:` の対象であり、mutation を表現する語彙を持つ。`asset` との役割対比は §asset / ADR-065 §決定 §5 を参照。

以下の `kind` / `of` はいずれの形態にも適用される。`initialized store` は `kind` を持たず `model` 参照のみで宣言される（[init オブジェクト](#init-オブジェクトinitializes内) 参照）。

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

`store.kind: db` は独自フィールド（列定義）を持たない。`of: <model-id>` でmodelを参照するのみ。ER図は `store.of` → model → fields の参照を辿って列定義を描画する（ADR-021）。

`store.of` は v1.1 でも model-id 参照のまま扱う。TypeRef 適用対象外である理由と対象フィールドは [type-ref.md](./type-ref.md) §1 を参照する。

### collectionのnote規約

| noteに書くもの | noteに書かないもの |
|--------------|-----------------|
| 単一collectionのフィールドへのフィルタ・検索 | 複数storeにまたがる結合 |
| 言葉で書ける程度の条件（等値・範囲・真偽フラグ） | 集計・変換・ソートを伴う複雑なクエリ |

collectionのnoteはLLMへのsemantic contract（ADR-007, ADR-008と同じ位置づけ）。

---

## actor

> 出典: ADR-004, ADR-006, ADR-031

Applicationレイヤー。人間・外部システム。Sequence Diagramのparticipant列として登場。UML標準Actorに対応。

actorはプロジェクトglobalな存在であり、モジュールに属さない。任意のファイル名・任意の配置で定義できる。プロジェクト内でIDがユニークであることをパーサーが保証する（ADR-031）。

```yaml
# actors.yaml（ファイル名は任意）
nodes:
  - id: stripe
    type: actor
    note: "外部決済サービス"

  - id: scheduler
    type: actor
    note: "cronスケジューラー"

  - id: end_user
    type: actor
    note: "サービスを利用するエンドユーザー"
```

### フィールド定義

固有フィールドなし（ADR-021）。`id` / `type` / `note` のみで表現する。  
モジュールパスは不要。参照は常にID直参照（ADR-031）。

| フィールド | 必須 | 型 | 内容 | 出典 |
|-----------|------|-----|------|------|
| （固有フィールドなし） | — | — | id/type/noteのみで表現 | ADR-004, ADR-021 |

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

- id: payment_webhook_received
  type: event
  source: external
  actor: stripe
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
| `source` | ✓ | enum | `ui` / `external` / `er` / `internal` | ADR-018, ADR-034 |
| `actor` | `external`のみ必須 | actor-id | 発火元のactor ID。brewprintのいずれかのファイルに `type: actor` ノードとして宣言されていること | ADR-018 |
| `payload` | 任意 | payload | イベントが運ぶデータのmodel参照 | ADR-018 |
| `watches` | `er`のみ必須 | store-id | 変化を監視するstore ID | ADR-018 |

### payload オブジェクト

| フィールド | 必須 | 型 | 内容 |
|-----------|------|-----|------|
| `model` | ✓ | model-id | ペイロードの型参照。v1.1 でも TypeRef ではなく model-id 参照として扱う |

### sourceの意味

| source | 意味 | `actor`の要否 | payloadの要否 |
|--------|------|-------------|-------------|
| `ui` | ユーザー操作（クリック・フォーム送信等）。UI participant列を暗黙生成 | 不要 | 任意（フォームデータ等） |
| `external` | 外部システム・スケジューラーからの入力。`actor:` で発火元を明示 | **必須** | 任意（受信データ） |
| `er` | storeの値変化による発火。`watches`必須 | 不要 | 任意 |
| `internal` | アプリ内部での発火（task完了・FSM runtime内部処理・内部タイマー等）。何を監視するかは `note:` で人間向けに記述（ADR-034） | 不要 | 任意 |

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
| `note` | 任意 | string | 状態の意味・この状態にある条件 | ADR-019 |
| `wireframe` | 任意 | object | この状態のUI骨格。単一rootのcontainerノード。`main` containerと`layout` objectを含む詳細はwireframe specで定義 | ADR-029, ADR-042 |

`wireframe`フィールドの詳細は `docs/spec/views/wireframe.md` を参照。`main` containerと`layout` objectの設計判断は `docs/adr/042-wireframe-main-and-layout.md` を参照。

`state` と `store` の区別:

| ノード | レイヤー | 概念 |
|--------|---------|------|
| `store` | Processing / Data | 実行時にデータを保持するインスタンス。DAGでreads/writesされる |
| `state` | Application | FSMが取り得る状態の定義。遷移の起点・終点として使われる |

---

## branch

> 出典: ADR-012, ADR-023

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

**スコープルール**: branch内部で生成されたassetはbranch外から直接参照不可。スコープ外にデータを渡す場合は `initializes` + `writes`/`reads` を使う。各パスが独立して終端する場合、後続から参照されないtask（floatingノード）はDAGでENDに直行する形でrenderされる（ADR-023）。

flow:内での記法は `spec/edges.md` の `1-4. branchエントリ` を参照。

---

## fork

> 出典: ADR-012

Processingレイヤー。並列分岐。後続パスを**すべて並列実行**する。必ず `join` とペアで使う。

```yaml
- id: fan_out
  type: fork
  note: "静的解析・動的解析・依存チェックを並列実行する"
```

`fork` 単体での使用は不正。対応する `join` が必須（ADR-012）。branch内taskへの入力は `flow:` セクションの `branches[].steps[].params` に明示する（ADR-040）。

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

| 種別 | レイヤー | ファイル配置 | Seq | State | DAG | ER | API |
|------|---------|------------|:---:|:---:|:---:|:---:|:---:|
| `task` | Processing | `task/*.yaml` | △※1 | ❌ | ✅ | ❌ | △※1 |
| `model` | Data | `model/*.yaml` | ❌ | ❌ | ❌ | ✅ | ❌ |
| `asset` | Processing | なし（task.returnsから導出） | ❌ | ❌ | ✅ | ❌ | ❌ |
| `store` | Processing / Data | `store/*.yaml` | ❌ | ❌ | ✅ | ✅ | ❌ |
| `actor` | Application | プロジェクトglobal。任意のファイル名で定義（ADR-031） | ✅ | ❌ | ❌ | ❌ | ❌ |
| `event` | Application | state.yaml等に同居 | ✅ | ✅ | ❌ | ❌ | ❌ |
| `state` | Application | `state.yaml` | ❌ | ✅ | ❌ | ❌ | ❌ |
| `branch` | Processing | `task/*.yaml`のサブノード等 | ❌ | ❌ | ✅ | ❌ | ❌ |
| `fork` | Processing | `task/*.yaml`のサブノード等 | ❌ | ❌ | ✅ | ❌ | ❌ |
| `join` | Processing | `task/*.yaml`のサブノード等 | ❌ | ❌ | ✅ | ❌ | ❌ |

凡例: Seq = Sequence Diagram / State = State Diagram / ER = ER Diagram / API = API Table  
※1 `endpoint: true` の場合のみ。Sequence Diagramではtaskはlifelineではなく矢印ラベルとして参照される（ADR-017）。

> `foreach` はADR-016にてnode typeから廃止。`flow:` セクションの制御構文として記述する（→ `spec/edges.md`）。
