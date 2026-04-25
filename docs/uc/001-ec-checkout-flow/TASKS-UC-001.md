# TASKS-UC-001: EC Checkout Flow

UC-001 固有の残作業・render例・spec gap を追跡するための作業台。
プロジェクト全体の入口は `docs/TASKS.md` に置き、細かい作業はこのファイルで管理する。

---

## ステータス

- [x] YAML群 作成完了
- [x] README.md 骨格作成
- [ ] docs/coverage.md を現行specに追随
- [ ] docs/render-*.md のTODOを解消
- [x] UC-001で発見したspec gapの扱いを確定
- [ ] 完了後に HANDOFF.md を削除

---

## 残作業

### カバレッジ（docs/coverage.md）

- [x] ノード種別
- [ ] flow構文（ADR-040 / foreach.over `$params.field` 解決済み状態へ更新）
- [x] task内フィールド
- [x] model内フィールド
- [x] state / event / wireframe
- [x] view定義

### render / view例（docs/render-*.md）

- [x] DAG: `auth.task.login`
- [x] DAG: `cart.task.validate_cart`
- [x] DAG: `order.task.checkout`
- [x] DAG: `order.task.process_order`
- [x] State Diagram: auth
- [x] State Diagram: order
- [x] State Diagram: inventory
- [x] ER Diagram: ec全体
- [ ] Sequence Diagram: checkout_flow（TODO placeholderあり）
- [ ] Sequence Diagram: payment_webhook_flow（TODO placeholderあり）
- [ ] API Table（TODO placeholderあり）

---

## UC-001で発見したspec gap

### 1. `foreach.over` に `$params.field` を指定可能か

- 対象: `docs/spec/edges.md`
- 発見元: `yaml/cart/task/validate_cart.yaml`
- 現状: `over: $params.cart_items`
- 論点: `foreach.over` が node ID のみか、`$params.field` も許可するか
- **解決**: `$params.field` も許可する。`edges.md` の `foreach.over` フィールド説明に追記済み。ADR不要（明示的な禁止がなかっただけで、設計変更ではない）。
  - `over: node_id` — 前段taskの `returns` assetを参照する既存の形式
  - `over: $params.field` — main taskの `params` から直接listを受け取る形式（自然なユースケース）
  - validator は `over` が `$params.field` の場合、`main.params.<field>.model` が `kind: list` であることを検証する

### 2. `fork.params` から branches 内 step へのwiring規則

- 対象: `docs/spec/edges.md`
- 発見元: `yaml/order/task/checkout.yaml`
- 旧状況: `fork.params.draft_order: build_order` を branches 内 task の同名paramへ暗黙伝播する前提で記述していた
- **解決**: ADR-040で暗黙伝播を廃止し、`fork.branches[].steps[].params` に明示する形式へ変更。
  - `docs/adr/040-control-flow-step-wiring.md` を起票済み
  - `docs/spec/edges.md` に反映済み
  - `yaml/order/task/checkout.yaml` を新構文へ更新済み

### 3. `branch.cases.step` 先taskへのparam wiring規則

- 対象: `docs/spec/edges.md`
- 発見元: `yaml/order/task/process_order.yaml`
- 旧状況: `branch.params.order: check_inventory` を `cases[].step` 先taskへ暗黙伝播する前提で記述していた
- **解決**: ADR-040で暗黙伝播を廃止し、`branch.cases[].params` に明示する形式へ変更。
  - `docs/adr/040-control-flow-step-wiring.md` を起票済み
  - `docs/spec/edges.md` に反映済み
  - `yaml/order/task/process_order.yaml` を新構文へ更新済み

---

## 完了条件

- [ ] docs/coverage.md が現行specに追随している
- [ ] docs/render-*.md の全render / view例がTODOなしで埋まっている
- [x] spec gap を ADR / spec / TASKS のどれで扱うか決まっている
- [ ] `HANDOFF.md` を削除済み
