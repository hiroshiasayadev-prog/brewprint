---
scope: docs/spec/overview.md
status: draft
last_updated: 2026-04-17
summary: >
  brewprintの全体概要。コンセプト・ノード種別・クロスエッジ・伝搬方向・
  責務分離方針・未解決課題を定義する。
depends_on:
  - docs/adr/001-node-type-splitting.md
  - docs/adr/002-folder-as-namespace.md
  - docs/adr/003-name-resolution-rules.md
  - docs/adr/004-sequence-diagram-participants.md
  - docs/adr/005-class-diagram-as-endpoint-view.md
open_issues:
  - eventノードのスキーマ未定
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

---

## ノード種別

| 種別 | 内容 |
|------|------|
| `procedure` | 処理。インライン記述またはrefで別ファイル参照 |
| `artifact` | 成果物。`scalar` / `struct` / `list` / `dict` |
| `state` | DB・グローバル変数・session_stateなど永続化されるもの |
| `actor` | 人間・外部システム。sequence diagramのエントリーポイントとして使用 |

### procedureのendpointフラグ

`endpoint: true` を付与したprocedureはバックエンドエンドポイントとして扱われ、class diagram viewに出力される。

```yaml
- id: login
  type: procedure
  endpoint: true
  method: POST
  path: /auth/login
  input: login_request    # artifactのID
  output: auth_token      # artifactのID
```

| フィールド | 必須 | 内容 |
|---|---|---|
| `endpoint` | ✓ | `true` のとき class diagram viewに出力 |
| `method` | ✓ | HTTP method（GET / POST / PUT / DELETE / PATCH） |
| `path` | ✓ | URLパス |
| `input` | 任意 | リクエストbodyのartifact ID |
| `output` | 任意 | レスポンスbodyのartifact ID |

### classについて

classは独立したノード種別として持たない。

> **「structのmethodsはnoteで明示できる程度のもの。それを超えるものはドメインロジックであり、procedureとしてDAGに出す」**

これはクリーンアーキテクチャのentity vs use caseの境界線とほぼ一致する。

```yaml
- id: voltage_result
  type: artifact
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

| エッジ | 意味 |
|--------|------|
| DAG procedure → ER table | `write` |
| ER table → DAG artifact | `read` |
| state transition → DAG procedure | `trigger` |
| DAG artifact → state | `reflect` |
| state → ER table | `hydrate` |

### クロスエッジの連鎖

クロスエッジは連鎖することがある。例：`write`の副作用としてER変化が発生し、別の`trigger`が発火する。

```
procedure → (write) → ER table → (trigger) → 別のprocedure
```

この連鎖はスコープ内だが、**ハッピーパスの範囲で記述できる連鎖のみ**を対象とする。
無限ループや競合が生じる連鎖はハッピーパス外としてスコープ外。

### procedure間の参照（refクロスエッジ）

別モジュールのprocedureをrefで呼ぶ場合、DAG内部のエッジではなくクロスエッジ的に振る舞う。
フルパス参照（ADR 003）で記述し、依存関係として扱う。

---

## 伝搬の方向性

**正方向（イベント起点）：**

```
trigger発生
  → state transition (state diagram)
    → procedure発動 (DAG)
      → artifact変形 (DAG)
        → storage書き込み (ER)
```

**負方向（データ反映）：**

```
ER変化
  → state反映
    → UI更新
```

### triggerの発生源（4種）

| 発生源 | 例 |
|--------|----|
| `ui` | ボタンクリック、フォーム送信 |
| `time` | cron、scheduled batch |
| `external` | webhook、message queue、WebSocket |
| `er` | テーブル変化による派生データ再計算 |

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

| 図 | renderの元となる要素 | 備考 |
|---|---|---|
| DAG | `procedure` ノード＋エッジ | 制御フロー・データフロー |
| ER | `state`（kind=db）＋フィールド定義 | テーブル構造 |
| state diagram | `state` ノード＋遷移エッジ | 状態遷移 |
| class diagram | `endpoint: true` なprocedureをモジュール単位でグルーピング | APIのI/Oシグネチャ |
| sequence diagram | Actor / UI / API / DB の4種のparticipantとそのやりとり | レイヤー間の粗い粒度の流れ |

### sequence diagramのparticipant対応

| participant | brewprintの実体 | 他の図へのリンク |
|---|---|---|
| Actor | `actor` ノード | なし |
| UI | `event`（source=ui） | なし |
| API | `procedure`（endpoint=true）のグループ | class diagram |
| DB | `state`（kind=db） | ER diagram |

矢印のラベルにはprocedure IDを記載する（例：`auth.procedure.login`）。
リンクにはならないが、IDがあればMCP経由で詳細を参照できる。

---

## 未解決課題

### eventノードの設計（最優先）

制御フローの起点として`event`ノードをDAGに導入する方向。
`source`属性（ui/time/external/er）でタグ付けする。具体的なスキーマは未定。

### Edgeの管理方式

クロスエッジにkind属性が必要なため、Nodeのadjacency listではなく別管理が有力。未決。
