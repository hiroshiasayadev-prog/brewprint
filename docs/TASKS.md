# TASKS

brewprintの積みタスク一覧。会話をまたいでコンテキストを引き継ぐために使う。

---

## ステータス凡例

- `[ ]` 未着手
- `[~]` 進行中
- `[x]` 完了

---

## 仕様設計（ADR + spec）

- [x] **ADR-001〜039 完了**
  - ノード種別・エッジ設計・名前解決・各view renderルール・wireframe DSL・シナリオスキーマ等を確定
  - 最新: ADR-039（ER図横断view YAML: `as: er_diagram`）

- [x] **spec/nodes.md** — 全ノード種別フィールド定義 (status: confirmed)

- [~] **spec/edges.md** — flow: / transitions: / reads・writes / $シジル体系 (status: wip)
  - 内容は実質完成しているが Front Matter が wip のまま。confirmed に更新すること

- [x] **spec/views/dag.md** (status: confirmed)
- [x] **spec/views/er.md** (status: confirmed)
- [x] **spec/views/state-diagram.md** (status: confirmed)
- [x] **spec/views/sequence-diagram.md** (status: confirmed)
- [x] **spec/views/api-table.md** (status: confirmed)
- [x] **spec/views/wireframe.md** (status: confirmed)

---

## ユースケース

- [~] **UC-001: EC Checkout Flow**（`docs/uc/001-ec-checkout-flow/`）
  - [x] YAML群（actors / model / store / task / state / views）全ファイル作成完了
  - [x] `views/scenarios/checkout_flow.yaml` — 2 step構成（action なし遷移 + fork+join taskのADR-037/038検証）
  - [x] `views/scenarios/payment_webhook_flow.yaml`
  - [x] `views/api_table.yaml` / `views/er.yaml`
  - [ ] **README.md 未作成**（10個のrender例を含む。HANDOFF.md §5参照）
  - 完了条件: README.md 作成 → HANDOFF.md 削除

---

## 実装（Go）

- [ ] **GoのAST struct定義**
  - ノード種別ごとにstruct（ADR-001〜現在のスキーマに基づく）
  - spec/nodes.md + spec/edges.md が基準ドキュメント

- [ ] **YAMLパーサー実装**
  - AST構築 + 名前解決（ADR-003, ADR-027）
  - エラーはASTビルド時に返す

- [ ] **MCPツール実装**
  - `get_signature` / `get_deps` / `inspect`
  - リバースルックアップ（cross-edge逆引き）はここで実装（ADR-020）

- [ ] **Mermaidレンダラー実装**
  - DAG / ER / State Diagram / Sequence Diagram / API Table / Wireframe
  - 各viewのrenderルールはspec/views/配下のspecに準拠
