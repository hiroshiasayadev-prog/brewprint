# TASKS

brewprintの積みタスク一覧。会話をまたいでコンテキストを引き継ぐために使う。

---

## ステータス凡例

- `[ ]` 未着手
- `[~]` 進行中
- `[x]` 完了

---

## 仕様設計（ADR + spec）

- [x] **ADR-001〜041 完了**
  - ノード種別・エッジ設計・名前解決・各view renderルール・wireframe DSL・シナリオスキーマ等を確定
  - ADR-040: 制御フロー内step wiring明示化（`fork.branches[].steps[].params` / `branch.cases[].params`）
  - ADR-041: Sequence Diagram message rendering（step index prefix / DB片方向 / autonumber不使用）

- [x] **spec/nodes.md** — 全ノード種別フィールド定義 (status: confirmed)

- [x] **spec/edges.md** — flow: / transitions: / reads・writes / $シジル体系 (status: confirmed)

- [x] **spec/views/dag.md** (status: confirmed)
- [x] **spec/views/er.md** (status: confirmed)
- [x] **spec/views/state-diagram.md** (status: confirmed)
- [x] **spec/views/sequence-diagram.md** (status: confirmed)
- [x] **spec/views/api-table.md** (status: confirmed)
- [x] **spec/views/wireframe.md** (status: confirmed)

---

## ユースケース

- [x] **UC-001: EC Checkout Flow**（`docs/uc/001-ec-checkout-flow/`）
  - [x] YAML群（actors / model / store / task / state / views）全ファイル作成完了
  - [x] `views/scenarios/checkout_flow.yaml` — 2 step構成（action なし遷移 + fork+join taskのADR-037/038検証）
  - [x] `views/scenarios/payment_webhook_flow.yaml`
  - [x] `views/api_table.yaml` / `views/er.yaml`
  - [x] `README.md` 骨格作成
  - [x] `TASKS-UC-001.md` 作成
  - [x] `docs/coverage.md` / `docs/render-*.md` 記入完了
  - [x] UC-001で発見したspec gapの解消方針決定
  - [~] **golden fixture 再配置**（ADR-043 / ADR-046）
    - [x] `render_index.yaml` 追加
    - [x] `renders/{group}/` へ DAG / State / Sequence / Wireframe を配置
    - [x] `_cross/` へ ER / API を配置
    - [x] `_preview/wireframe.html` を配置
    - [x] README / coverage / legacy docs 表記を追随
    - [ ] Go renderer 実装後に `renders/` を正式 golden fixture として再生成・検証
  - 詳細: `docs/uc/001-ec-checkout-flow/TASKS-UC-001.md`

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
