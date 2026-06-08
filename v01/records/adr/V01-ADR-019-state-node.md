# V01-ADR-019: stateノード設計（FSM）

- **status**: accepted
- **date**: 2026-04-19

## 背景

V01-ADR-017にてState DiagramはApplicationレイヤーの図として位置づけられた。
`store`（データ保持、V01-ADR-007）とは別概念として、FSMの状態を表現する`state`ノード種別を設計する必要があった。

合わせて以下を設計する：
- transition記法（V01-ADR-015の`flow:`セクションとの関係）
- eventノード（V01-ADR-018）との接合

## 決定

### 1. `state`はnode typeとして`nodes:`に書く

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
```

#### stateノードのフィールド定義

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `initial` | 任意 | `true`でFSMの初期状態。ファイルごとに1つ |
| `final` | 任意 | `true`で終端状態（複数可） |
| `note` | 任意 | 状態の意味・この状態にある条件 |

### 2. 遷移は`transitions:`セクションに分離

DAGの`flow:`と同じ「ノード定義とwiring分離」の思想を踏襲するが、セクション名は`transitions:`とする。`flow:`はProcessingレイヤー（DAG）の概念であり、Applicationレイヤーの状態遷移と混在させない。

```yaml
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

#### transitionエントリのフィールド定義

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `from` | ✓ | 遷移元state ID |
| `on` | ✓ | トリガーとなるevent ID（V01-ADR-018） |
| `to` | ✓ | 遷移先state ID |
| `action` | 任意 | 遷移に伴うtask ID（ProcessingレイヤーのCA参照） |
| `guard` | 任意 | 遷移条件テキスト（評価はbrewprintのスコープ外） |
| `note` | 任意 | 補足説明 |

#### 同一 `(from, on)` の複数transition（V01-ADR-035で追加）

同一 `(from, on)` ペアに対して、**互いに異なる `guard:` を持つ複数のtransitionを許容する**（Mealy machine / UML 2.x StateMachines準拠）。transitionの識別キーは `(from, on, guard)` の3タプル。

制約：

- 同一 `(from, on)` に複数transitionが存在する場合、**全エントリに `guard:` が必須**
- 同一 `(from, on, guard)` の完全一致は不可（パーサーエラー）
- guardのない単独transitionと、guardのある複数transitionを混在させることはできない

この分岐はstate diagramではchoice pseudostate、sequence diagramではstepの `guard:` で解決する（V01-ADR-035）。

### 3. `action`はtransitionに属し、eventには属さない

`action`を`transition`に置く理由は、actionが **「現在のstate × event」の組み合わせで決まるもの** だから。

同じeventでも遷移元stateが異なればactionが変わる：

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

eventはあくまで「何が起きたか」という純粋なトリガーであり、副作用（action）の指定は持たない。

### 4. eventはstate diagramファイルにstateと同居させる

eventノードはstate diagram YAMLの`nodes:`に`state`と同列で定義する（V01-ADR-018のnode typeとしての扱いは変わらない）。

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
```

Sequence Diagramから同じeventを参照する場合はcross-file参照になる。cross-file参照の解決ルールはV01-ADR-003に従う。

### 5. `store`との区別

| ノード | 種別 | 概念 |
|--------|------|------|
| `store` | データ保持（V01-ADR-007） | 実行時に状態を保持するインスタンス。DAGのflow内でreads/writesされる |
| `state` | FSMの状態（本ADR） | FSMが取り得る状態の定義。Applicationレイヤーに属する |

名称が似ているが概念は別。`store`はProcessingレイヤー、`state`はApplicationレイヤーに属する。

## 理由

### `transitions:`セクションを分離する

stateノードのYAML定義にtransitionを混在させると（XState形式のように）、「あるstateがどのstateと繋がるか」という情報が各nodeに分散する。`transitions:`セクションに集約することで、FSM全体のフローが1箇所で把握できる。

`flow:`と命名しないのは、`flow:`はDAGのデータフロー（Processingレイヤー）を指す語として既に確立しているため。レイヤーが異なれば構文も別名にすることで、YAMLを読んだ際に図の種別が即座に判別できる。

### `action`をtransitionに置く

actionは`(state, event)`の関数の値として決まる。eventに置くと「同じeventがstateによって異なるtaskを起動する」が表現できなくなる。FSMの基本セマンティクス（Mealy machine）に従い、transitionに属させる。

### `guard`をテキストのみにする

guardの評価はbrewprintのスコープ外（実装言語依存）。構造化するほどの複雑な条件は、extract taskまたはtaskのロジックとして実装すべきであり、brewprintのYAML上で表現する必要はない。

### eventをstate diagramファイルに同居させる

State Diagramが扱うeventはそのFSM専用のものが多く、別ファイルに分離するとcross-file参照が増えて見通しが悪くなる。Sequence DiagramなどからもアクセスされるeventはV01-ADR-003のcross-file参照で解決できるため、定義場所の制約は不要。

## 影響

- `spec/overview.md` の「ノード種別」テーブルに `state` を追加する
- `spec/overview.md` の「書ける図の一覧」のState Diagramの説明を本ADRに基づき更新する
- `store`と`state`の混同を避けるため、`spec/nodes.md`（または対応するspec）に両者の対比を記載する

## Evidence
- commit: 30611b3
- impl commit: tbd
- 参考: Mealy machine（FSM基本セマンティクス）、XState（transition記法の参考）
