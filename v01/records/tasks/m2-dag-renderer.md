# Milestone 2: DAG renderer の対象を広げる

- **status**: closed
- **scope**: DAG renderer
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

---

## Tasks

- [x] **foreach をDAG rendererに追加する**
  - `cart.task.validate_cart`
  - `foreach.over: $params.field`
  - `$item`
  - apply先taskの `↻` 表示
  - `docs/uc/001-ec-checkout-flow/renders/commerce/dag-validate_cart.md` とgolden一致

- [x] **fork / join をDAG rendererに追加する**
  - `order.task.checkout`
  - V01-ADR-040 の `branches[].steps[].params`
  - join.params の branch終端step `returns.name` 解決
  - `docs/uc/001-ec-checkout-flow/renders/commerce/dag-checkout.md` とgolden一致

- [x] **branch / cases をDAG rendererに追加する**
  - `order.task.process_order`
  - `branch.params` と `cases[].params` の分離
  - floating node → `_end`
  - 制御フロースコープ（V01-ADR-023）
  - `docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_order.md` とgolden一致

- [x] **M2完了確認**
  - 既存M1の `auth.task.login` golden test維持
  - `go fmt ./...` / `go test ./...` 通過（2026-04-27）
