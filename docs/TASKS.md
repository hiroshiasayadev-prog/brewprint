# TASKS

brewprintの積みタスク一覧。会話をまたいでコンテキストを引き継ぐために使う。

---

## ステータス凡例

- `[ ]` 未着手
- `[~]` 進行中
- `[x]` 完了

---

## 仕様設計（ADR + spec）

- [x] **ADR-001〜046 完了**
  - ノード種別・エッジ設計・名前解決・各view renderルール・wireframe DSL・シナリオスキーマ等を確定
  - ADR-040: 制御フロー内step wiring明示化（`fork.branches[].steps[].params` / `branch.cases[].params`）
  - ADR-041: Sequence Diagram message rendering（step index prefix / DB片方向 / autonumber不使用）
  - ADR-042: wireframeのmain containerとlayout object
  - ADR-043: プロジェクトルートレイアウトとrender出力構造
  - ADR-044: DAG store access edgeにread/writeラベルを付与
  - ADR-045: render_index.yaml スキーマ
  - ADR-046: State / Sequence / Wireframe / Preview のrender出力配置

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

実装は横切り（AST → parser → renderer）ではなく、**Raw YAML → ResolvedProject → QueryService / Renderer** の縦切りで進める。
renderer と MCP wrapper が Raw YAML structs を直接読まない方針を前提にする。

### Milestone 0: 実装境界を固定する

- [x] **ADR-047: Go semantic model / query layer boundary を起票**
  - Raw YAML structs / ResolvedProject / QueryService / Renderer の責務境界を決める
  - renderer と MCP wrapper が Raw YAML structs を直接読まない方針を明文化する
  - validation / name resolution / derived model build の責務範囲を決める
  - renderer は ResolvedProject または view-specific view model を読む方針を明文化する
  - 具体的なview model型はDAG vertical slice実装中に固める

- [x] **ADR-048: ResolvedProject index strategy を起票**
  - reverse lookup index を ResolvedProject build 時に構築する方針を決める
  - MCP tool / renderer は都度Raw構造を走査せず、ResolvedProjectのindexを読む
  - 初期index候補を決める
    - referencesBySource
    - referencesByTarget
    - tasksReadingStore
    - tasksWritingStore
    - transitionsByStateEventGuard
    - actionsByTask
    - scenariosByID

- [x] **ADR-049: MCP / QueryService の reference 語彙統一 を起票**
  - 外部MCP tool名を `get_references` に統一する
  - 内部QueryService method名を `GetReferences` に統一する
  - `get_deps` / `GetDeps` は採用しない
  - MCP responseの中心語彙を `references` とする

- [x] **docs/spec/mcp.md を薄く作る**
  - `get_signature` / `get_references` / `inspect` / `list_endpoints` の input / output を決める
  - QualifiedID / FileID / TransitionID / private sub node synthetic ID の指定形式を決める
  - ObjectRef / TransitionRef / AssetRef / Diagnostic / Reference の共通schemaを決める
  - not found 時の error 形式を決める
  - references は v1 では direct のみとする
  - inspect の粒度と reverse lookup の返却範囲を決める
  - まだ MCP server transport は実装しない

### Milestone 1: UC-001のDAG 1本を縦に通す

- [x] **Go project skeleton を作る**
  - package構成
  - UC-001 fixture 読み込み
  - YAML decode の入口
  - golden test の入口

- [x] **YAML loader + file classifier を実装する**
  - node file / view file / `render_index.yaml` を分類する
  - view file は `as:` で判定する
  - node file は `nodes:` を入口にする

- [x] **Raw YAML structs を実装する**
  - M1範囲として task / model / store / actor / params / returns / reads / writes / initializes を実装
  - view file / render_index.yaml は分類のみで、decodeは後続milestoneに回す
  - validationは薄くてよい

- [x] **Symbol table / QualifiedID parser を実装する**
  - ADR-027 sentinel方式
  - actor global（ADR-031）
  - module nesting
  - 同モジュール内ID直書き
  - クロスモジュールフルパス

- [x] **ResolvedProject build の最小版を実装する**
  - main node by file
  - model / store / task / actor index
  - implicit asset
  - file-private initialized store
  - reads/writes index
  - nodesByQID / nodesByFile
  - DAG renderer が必要な範囲だけ先に実装する

- [x] **DAG renderer vertical slice を通す**
  - UC-001の `auth.task.login` を Raw → Resolved → Render まで通す
  - `docs/uc/001-ec-checkout-flow/renders/auth/dag-login.md` とgolden一致させる
  - params / returns boundary
  - implicit asset
  - store read/write edge label（ADR-044）
  - Tasks詳細セクション
  - `go test ./...` 通過（2026-04-27）

### Milestone 2: DAG renderer の対象を広げる

- [ ] **foreach をDAG rendererに追加する**
  - `cart.task.validate_cart`
  - `foreach.over: $params.field`
  - `$item`
  - apply先taskの `↻` 表示

- [ ] **fork / join をDAG rendererに追加する**
  - `order.task.checkout`
  - ADR-040 の `branches[].steps[].params`
  - join.params の branch終端step `returns.name` 解決

- [ ] **branch / cases をDAG rendererに追加する**
  - `order.task.process_order`
  - `branch.params` と `cases[].params` の分離
  - floating node → `_end`
  - 制御フロースコープ（ADR-023）

### Milestone 3: QueryService を通す

- [ ] **QueryService vertical slice を実装する**
  - `GetSignature`
  - `GetReferences` はdirect referencesのみでよい
  - `Inspect`
  - unit test / CLI相当から直接叩けるようにする

- [ ] **reverse lookup index を整備する**
  - ADR-048に準拠して実装する
  - referencesBySource
  - referencesByTarget
  - tasksReadingStore
  - tasksWritingStore
  - transitionsByStateEventGuard
  - actionsByTask
  - scenariosByID

### Milestone 4: render_index / output placement を実装する

- [ ] **render_index.yaml validation を実装する**
  - ADR-045に準拠
  - group id命名規則
  - module重複禁止
  - uncovered module の warning
  - nested module の親group所属

- [ ] **renders/ 出力配置を実装する**
  - ADR-043 / ADR-046に準拠
  - `renders/{group}/...`
  - `_cross/`
  - `_preview/`
  - master index / group index

### Milestone 5: 残りrendererを順に実装する

- [ ] **State Diagram renderer を実装する**
  - guard分岐の choice pseudostate（ADR-035）
  - FSMファイル単位の `state-{fsm-id}.md`（ADR-046）

- [ ] **Sequence Diagram renderer を実装する**
  - event source別矢印ルール（ADR-036）
  - actionなしtransition（ADR-037）
  - sub task reads/writes traversal（ADR-038）
  - step index prefix / DB片方向message（ADR-041）

- [ ] **ER renderer を実装する**
  - 横断view YAML（ADR-039）
  - cross module FK
  - `store.kind: db` から model fields を辿る

- [ ] **API Table renderer を実装する**
  - `as: api_table`
  - `http_root_path`
  - endpoint task の leaf path 合成（ADR-028）

- [ ] **Wireframe renderer を実装する**
  - HTML fragment
  - `main` container / `layout` object（ADR-042）
  - fixed CSS profile
  - preview harness

### Milestone 6: MCP server wrapper を実装する

- [ ] **MCP server wrapper を実装する**
  - QueryServiceをMCP toolとして公開する
  - MCP transportは最後に薄く被せる
  - renderer / QueryService / MCP wrapper の責務が混ざっていないことを確認する
