---
scope: docs/spec/views/state-diagram.md
status: confirmed
last_updated: 2026-04-20
summary: >
  State DiagramのrenderルールとMermaid出力仕様を定義する。
  stateノード・eventノード・transitions:セクションを入力とし、
  stateDiagram-v2形式のMermaidを生成するルールを記述する。
depends_on:
  - docs/adr/017-diagram-layers-and-scope.md
  - docs/adr/018-event-node.md
  - docs/adr/019-state-node.md
---

# State Diagram renderルール

## 対象ノード

State Diagramに登場するノードは以下の2種のみ。

| ノード | 役割 |
|--------|------|
| `state` | FSMの状態として描画 |
| `event` | 遷移トリガーのラベルとして使用 |

`store` / `task` / `asset` 等はState Diagramに登場しない（ADR-017）。

---

## renderスコープ

State DiagramはYAML**ファイル単位**で描画する。  
1ファイル = 1FSM = 1枚の図。複数FSMを合成しない。

---

## stateの描画

### initial / final

| フィールド | Mermaid上の表現 |
|-----------|----------------|
| `initial: true` | `[*] --> <state-id>` として描画 |
| `final: true` | `<state-id> --> [*]` として描画 |
| どちらもなし | 通常のstate |

```mermaid
stateDiagram-v2
  [*] --> idle
  authenticated --> [*]
  error --> [*]
```

### stateのラベル

state IDをそのまま表示する。`note` は図上には出力しない（LLM向けのセマンティクス情報として保持）。

---

## transitionの描画

`transitions:` セクションの各エントリを1本のエッジとして描画する。

### エッジラベルの構成

エッジラベルは以下の形式で構成する：

```
<event-id> [<guard>] / <action>
```

各要素の出力ルール：

| 要素 | 出力条件 |
|------|---------|
| `<event-id>` | 常に出力（`on` フィールドの値） |
| `[<guard>]` | `guard` フィールドがある場合のみ出力 |
| `/ <action>` | `action` がクロスファイル参照（ドット区切り）の場合のみ出力。同一ファイル内参照は省略 |

### クロスファイル参照の判定

`action` フィールドの値にドット（`.`）が含まれる場合をクロスファイル参照とみなす。

| action の値 | 判定 | ラベル出力 |
|---|---|---|
| `login_task` | 同一ファイル内参照 | 省略 |
| `auth.task.login` | クロスファイル参照 | `/ auth.task.login` |

### ラベルパターン例

```
login_submitted                                      ← guard/actionなし（同一ファイルactionまたはaction省略）
login_submitted [retryCount < 3]                     ← guardのみ
login_submitted / auth.task.login                    ← クロスファイルactionのみ
login_submitted [retryCount < 3] / auth.task.login   ← guard + クロスファイルaction
```

---

## Mermaid出力イメージ

### YAMLの入力例

```yaml
nodes:
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

  - id: session_expired
    type: state

  - id: login_submitted
    type: event
    source: ui
    payload:
      model: login_form

  - id: login_succeeded
    type: event
    source: external
    actor: auth_server

  - id: login_failed
    type: event
    source: external
    actor: auth_server

  - id: session_timeout
    type: event
    source: external
    actor: scheduler

transitions:
  - from: idle
    on: login_submitted
    to: loading
    action: login_task              # 同一ファイル内 → ラベル省略

  - from: session_expired
    on: login_submitted
    to: loading
    action: auth.task.reauth        # クロスファイル → ラベル表示

  - from: loading
    on: login_succeeded
    to: authenticated

  - from: loading
    on: login_failed
    to: error
    guard: "retryCount < 3"

  - from: authenticated
    on: session_timeout
    to: session_expired
```

### 期待するMermaid出力

```mermaid
stateDiagram-v2
  [*] --> idle
  authenticated --> [*]
  error --> [*]

  idle --> loading : login_submitted
  session_expired --> loading : login_submitted / auth.task.reauth
  loading --> authenticated : login_succeeded
  loading --> error : login_failed [retryCount < 3]
  authenticated --> session_expired : session_timeout
```

---

## 出力フォーマット

````markdown
# {ファイルID}

{ファイルのnote（あれば）}

```mermaid
stateDiagram-v2
  ...
```

## States

| state | note |
|-------|------|
| idle | ユーザーが操作していない状態 |
| loading | — |

## Events

| event | source | note |
|-------|--------|------|
| login_submitted | ui | ログインフォームのsubmit |
| login_succeeded | external (auth_server) | — |
| login_failed | external (auth_server) | — |
| session_timeout | external (scheduler) | — |
````

- H1 = ファイルID
- 説明文 = ファイルレベルの `note`。ない場合は省略
- **States表** = `type: state` の全ノード。`note` がない場合は `—`
- **Events表** = `type: event` の全ノード。`source` と `note` を列挙。`note` がない場合は `—`
- Mermaid記法: `stateDiagram-v2`

---

## 図の生成元

State DiagramはYAMLを読み込んだbrewprintのMCPツール（`render_state_diagram`）が生成する。  
手書きのMermaid記述は存在しない。
