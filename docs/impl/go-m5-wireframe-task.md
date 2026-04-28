# Go M5 Wireframe 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: Go実装 Milestone 5 第5段の Wireframe renderer vertical slice 実装チェックリスト

---

## 1. 目的

M5第5段では、Wireframe rendererをUC-001のstate.wireframeで縦切り実装する。

対象:

```text
auth/state.yaml  login_screen      -> renders/auth/wireframe-auth-login_screen.html
auth/state.yaml  loading           -> renders/auth/wireframe-auth-loading.html
order/state.yaml cart              -> renders/commerce/wireframe-order-cart.html
order/state.yaml checkout_screen   -> renders/commerce/wireframe-order-checkout_screen.html
```

まずは state YAML内の `wireframe` decode / semantic state build / HTML fragment render / golden test 4本までを通す。
`render/wireframe` は `semantic.Project` のみを読み、Raw YAML structs には依存しない。

---

## 2. 実装範囲

- state node の `wireframe` tree decode
- semantic.WireframeElement / Layout
- container / leaf HTML fragment render
- layout object の inline style変換
- HTML escape
- UC-001 golden test 4本
- preview harnessは次段で必要に応じて追加

---

## 3. 境界

```text
source           -> rawyaml
resolve          -> rawyaml, semantic
render/wireframe -> semantic
```

禁止:

```text
render/wireframe -> rawyaml
render/wireframe内でYAMLを直接読む
ER / API Table rendererを混ぜる
```

---

## 4. 受け入れ条件

- [x] state.wireframeをdecodeできる
- [x] `semantic.State` にwireframe treeを保持できる
- [x] auth.login_screen をrenderできる
- [x] auth.loading をrenderできる
- [x] order.cart をrenderできる
- [x] order.checkout_screen をrenderできる
- [x] layout object をinline styleへ変換できる
- [x] `render/wireframe` が `rawyaml` をimportしていない
- [x] golden test 4本が通る
- [x] preview harness golden test 1本が通る
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る
