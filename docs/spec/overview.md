---
scope: docs/spec/overview.md
status: draft
last_updated: 2026-04-19
summary: >
  brewprintの全体概要。コンセプト・ノード種別・クロスエッジ・伝搬方向・
  責務分離方針・未解決課題を定義する。
depends_on:
  - docs/adr/001-node-type-splitting.md
  - docs/adr/002-folder-as-namespace.md
  - docs/adr/003-name-resolution-rules.md
  - docs/adr/004-sequence-diagram-participants.md
  - docs/adr/007-asset-store-boundary.md
  - docs/adr/008-dag-dataflow-and-field-validation.md
  - docs/adr/009-task-io-design.md
  - docs/adr/010-ca-enforcement-directory-structure-model-asset-split.md
  - docs/adr/011-file-main-node-and-sub-nodes.md
  - docs/adr/017-diagram-layers-and-scope.md
  - docs/adr/018-event-node.md
open_issues:
  - stateノード（FSM）のスキーマ未定（ADR-019）
  - Edgeの別管理 vs Node内adjacency list未決
  - non-functional属性（retry/idempotent/async）のfirst-class化はdogfood後に判断
---

# brewprint 概要仕様

## コンセプト

brewprintは**人間とLLMの共通設計言語**。

| 対象 | 出力 |
|------|------|
| 人間 | Mermaid図（md形式） |
| LLM | signature / dep tree / inspect（MCP経由） |
| YAML | 裏側にある中間表現にすぎない |

同一YAMLから複数の図をrenderする。ER・class diagram・DAG・state diagramは**同じシステムの別の切り口**であり、図はviewであって実体はひとつ。

### YAMLはsingle source of truth

YAMLがシステム設計の唯一の真実であり、実装・図・MCPレスポンスはすべてそこから導出される。

- 「YAMLに存在しない実装」は原理的に生まれない
- 「実装だけ追加」はYAMLを先に書かないと始まらない
- 設計とドキュメントの乖離を構造的に防ぐ

参照した公知概念：Terraform（infrastructure as code）・Prisma（schema as single source of truth）

### brewprintのスコープ

brewprintは**設計言語とMCPによるコンテキスト供給**までを責務とする。コード生成はbrewprintのスコープ外。

```
brewprint（設計言語 + MCP）
    ↓ 設計コンテキストを供給
Claude Code（CLAUDE.mdに従い実装）
    ↓ 曖昧な点は実装を止める
impl_tasks.md（未解決実装項目の積み場）
```

Claude CodeはbrewprintのMCPツールで設計を参照しながら実装する。CLAUDE.mdには以下を記述する：
- YAMLのディレクトリ構造に従ってファイルを作ること
- 設計上曖昧な点があれば即座に実装を停止し `impl_tasks.md` に記録すること
- brewprintのMCPツールで設計コンテキストを参照すること

### AI実装前提の設計哲学

brewprintの**実装者はほぼAIを想定**している。この前提から導かれる設計方針：

> **「人間が書きやすい」より「AIが迷わない」を優先する**

具体的には：
- **静的検証性に極振り** — IDによる参照の一意性を最優先にする
- **インライン定義禁止** — 使い捨てstructであっても名前付き定義を強制する（ADR 009）
- **曖昧な記法を排除** — 複数の解釈が生まれうる記法は採用しない

人間向けの「書きやすさ」はUIやツールで補う。YAML自体は厳密さを優先する。

---

## ノード種別

> 命名の根拠は `docs/adr/006-node-type-renaming.md` を参照。
> model/asset 分離の根拠は `docs/adr/010-ca-enforcement-directory-structure-model-asset-split.md` を参照。

| 種別 | レイヤー | 内容 |
|------|---------|------|
| `task` | Processing | 処理。`returns` 宣言によって `asset` ノードを暗黙的に生成する |
| `model` | Data | 型定義。`scalar` / `struct` / `list` / `dict`。`model/` サブディレクトリに1ファイル=1定義 |
| `asset` | Processing | フロー上の存在。`task` の `returns` から暗黙的に生まれる。独立ファイルは持たない |
| `store` | Processing / Data | 実行時にデータを保持する実体。DB・session_state・context・collectionなど |
| `actor` | Application | 人間・外部システム。sequence diagramのエントリーポイントとして使用 |
| `event` | Application | 制御フローの起点。`source`（ui/time/external/er）でタグ付け。DAGには登場しない |
| `state` | Application | FSMの状態ノード。State Diagramで使用。`store`（データ保持）とは別概念（ADR-019参照） |
| `branch` | Processing | 排他分岐。条件に応じて後続パスを1本だけ選ぶ。合流点は暗黙（edge構造から読む） |
| `fork` | Processing | 並列分岐。後続パスをすべて並列実行する。必ず `join` とペアで使う |
| `join` | Processing | 合流。対応する `fork` の全ブランチが揃うまで待つ。必ず `fork` とペアで使う |

> `foreach` はADR-016にてnode typeから廃止。`flow:` セクションの制御構文として記述する。

### taskのendpointフラグ

`endpoint: true` を付与したtaskはバックエンドエンドポイントとして扱われ、`list_endpoints` MCPツールで出力される。class diagram viewには出力しない（ADR-017）。

```yaml
- id: login
  type: task
  endpoint: true
  method: POST
  path: /auth/login
  params:
    - name: request
      model: login_request    # modelのID
  returns:
    name: auth_token
    model: token              # このassetノードがDAG上に暗黙的に生える
```

| フィールド | 必須 | 内容 |
|---|---|---|
| `endpoint` | ✓ | `true` のとき `list_endpoints` MCPツールで出力 |
| `method` | ✓ | HTTP method（GET / POST / PUT / DELETE / PATCH） |
| `path` | ✓ | URLパス |
| `params` | 任意 | リクエストbodyのmodel ID |
| `returns` | 任意 | レスポンスbodyのmodel IDと生成されるasset名 |

### classについて

classは独立したノード種別として持たない。

> **「structのmethodsはnoteで明示できる程度のもの。それを超えるものはドメインロジックであり、taskとしてDAGに出す」**

これはクリーンアーキテクチャのentity vs use caseの境界線とほぼ一致する。

```yaml
- id: voltage_result
  type: model
  kind: struct
  fields:
    - name: value
      type: scalar.float
    - name: unit
      type: scalar.text
  note: "is_valid: valueが正かつunitが空でない"
```

---

## クロスエッジの種類

レイヤーをまたぐエッジは `write` / `read` の2種のみ。task nodeの `writes` / `reads` フィールドとして記述する（ADR-020）。

| kind | 向き | 意味 |
|------|------|------|
| `write` | task → store | taskがstoreを更新する |
| `read` | task → store | taskがstoreを参照する |

`trigger` / `reflect` / `hydrate` はADR-018/019のevent・transition・`transition.action` で表現済みのため廃止。

### クロスエッジの連鎖

クロスエッジは連鎖することがある。例：`write`の副作用としてER変化が発生し、別の`trigger`が発火する。

```
task → (write) → ER table → (trigger) → 別のtask
```

この連鎖はスコープ内だが、**ハッピーパスの範囲で記述できる連鎖のみ**を対象とする。
無限ループや競合が生じる連鎖はハッピーパス外としてスコープ外。

### task間の参照（refクロスエッジ）

別モジュールのtaskをrefで呼ぶ場合、DAG内部のエッジではなくクロスエッジ的に振る舞う。
フルパス参照（ADR 003）で記述し、依存関係として扱う。

---

## 伝搬の方向性

**正方向（イベント起点）：**

```
trigger発生
  → state transition (state diagram)
    → task発動 (DAG)
      → asset変形 (DAG)
        → storage書き込み (ER)
```

**負方向（データ反映）：**

```
ER変化
  → store反映
    → UI更新
```

### triggerの発生源（4種）

> eventノードの詳細設計は `docs/adr/018-event-node.md` を参照。

| 発生源 | 例 |
|--------|----|
| `ui` | ボタンクリック、フォーム送信 |
| `time` | cron、scheduled batch |
| `external` | webhook、message queue、WebSocket |
| `er` | storeの値が変化したことによって発火 |

---

## 責務分離の方針

### ハッピーパス前提

```
DAG/UML       → ハッピーパスの構造を示す
impl design   → 例外・並列・トランザクションを詰める
実装          → コードで担保
```

### 図で表せないことの扱い

- **図で表せること** → YAMLの構造として定義、validationが効く
- **図で表せないこと** → `note`フィールドに自然言語で補足、validationスコープ外

### noteの運用規約

`note`に書いた内容は機械検証できない。乖離を防ぐための運用規約：

- **noteに書いたことは実装時にテストで担保する**
- LLMがnoteを過信して実装しないよう、noteはあくまで補足情報として扱う
- 将来的には「noteからテストコードを生成する」拡張を検討（dogfood後に判断）

### non-functional属性のfirst-class化

dogfoodしながら必要なものだけ昇格させる運用。候補：`retry` / `idempotent` / `async`

---

## スコープ

**スコープ内：** UIコンポーネントのI/O・処理・状態など内部的な設計

**スコープ外：**
- 具体的なstyle・視覚的な配置
- 並列・競合・ロールバック・双方向同期（ハッピーパス外）

---

## 書ける図の一覧

> 図レイヤーとスコープの根拠は `docs/adr/017-diagram-layers-and-scope.md` を参照。

### Applicationレイヤー

| 図 | renderの元となる要素 | 備考 |
|---|---|---|
| Sequence Diagram | `actor` / `event` / `task`（endpoint=true）/ `store`（kind=db） | 誰が・何をいつやるか |
| State Diagram | `state`（FSM）/ `event` / `store` | 何がどんな状態を持ち、どう遷移するか（ADR-019） |
| API Table | `task`（endpoint=true） | `list_endpoints` MCPツールで出力。Mermaid描画なし |

### Processingレイヤー

| 図 | renderの元となる要素 | 備考 |
|---|---|---|
| DAG | `task` / `asset` / `branch` / `fork` / `join` ＋ flow:エッジ | 処理の中身はどうなっているか |

### Dataレイヤー

| 図 | renderの元となる要素 | 備考 |
|---|---|---|
| ER Diagram | `store`（kind=db）＋フィールド定義 | データはどんな構造か |

### レイヤー間の依存方向

```
Application → Processing → Data
```

- ApplicationレイヤーはProcessingレイヤーのノード（task等）を参照できる
- ProcessingレイヤーはApplicationレイヤーのノード（event / state等）を参照しない

### sequence diagramのparticipant対応

| participant | brewprintの実体 | 他の図へのリンク |
|---|---|---|
| Actor | `actor` ノード | なし |
| UI | `event`（source=ui） | なし |
| API | `task`（endpoint=true）のグループ | API Table（MCPツール） |
| DB | `store`（kind=db） | ER Diagram |

矢印のラベルにはtask IDを記載する（例：`auth.task.login`）。
リンクにはならないが、IDがあればMCP経由で詳細を参照できる。

---

## 未解決課題

### stateノード（FSM）の設計

State Diagram用のFSM状態ノード。`store`（データ保持）とは別概念。ADR-019にて確定済み。

### Edgeの管理方式

クロスエッジは `write` / `read` の2種に絞り、task nodeの `reads`/`writes` フィールドとして記述する。ADR-020にて確定済み。
