# Go M4 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: Go実装 Milestone 4 第1段の render_index / output placement vertical slice 実装チェックリスト（完了済み）

---

## 1. 目的

このファイルは、Milestone 4「render_index / output placement」を実装するための作業チェックリストである。

M4の目的は、UC-001の `render_index.yaml` をGoで読み、moduleからgroupへの配置解決を行い、既存DAG rendererの出力先を `renders/{group}/dag-{task}.md` として決定できるようにすること。

### M4第1段完了メモ

M4第1段は実装完了済みとして扱う。

- `render_index.yaml` decodeを `rawyaml` / `source` に追加済み。
- `internal/render/placement` を作成し、group validation / module-to-group解決 / DAG output path決定を実装済み。
- UC-001の既存DAG fixture配置と一致するplacement unit testを追加済み。
- master index / group index markdown skeletonを追加済み。
- `go fmt ./...` / `go test ./...` は 2026-04-28 に通過済み。
- State / Sequence / Wireframe本体rendererと実ファイル書き出しは後続milestoneへ残す。

M4第1段では、State / Sequence / Wireframe本体のrendererは作らない。既存fixtureの配置規則に合わせ、DAG output placementとindex生成の土台だけを通す。

---

## 2. 前提ドキュメント

実装前に最低限読むこと。

- `docs/doc-policy.md`
- `docs/impl/go-skeleton.md`
- `docs/impl/go-m1-task.md`
- `docs/impl/go-m2-task.md`
- `docs/impl/go-m3-task.md`
- `docs/TASKS.md` の Milestone 4
- `docs/adr/043-project-root-layout-and-render-output.md`
- `docs/adr/045-render-index-schema.md`
- `docs/adr/046-render-output-placement-for-state-sequence-wireframe-preview.md`
- `docs/uc/001-ec-checkout-flow/render_index.yaml`
- `docs/uc/001-ec-checkout-flow/renders/index.md`
- `docs/uc/001-ec-checkout-flow/renders/auth/index.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/index.md`
- `docs/uc/001-ec-checkout-flow/renders/catalog/index.md`

---

## 3. 実装原則

- M4も縦切りを優先する。
- `render_index.yaml` decodeは `rawyaml` / `source` に置く。
- group validation / placement resolutionはRaw YAMLではなく専用packageまたはsemantic寄りのread modelで扱う。
- DAG renderer本体は引き続き `semantic.Project` だけを読む。
- output placementはrendererのMermaid生成ロジックに混ぜない。
- M4ではCLIを広げすぎない。まずunit testから直接叩けるGo APIで通す。

M4で守る境界:

```text
source          -> rawyaml
resolve         -> rawyaml, semantic
render/dag      -> semantic
render/placement -> rawyaml, semantic（またはrawyaml由来config + semantic.Project）
query           -> semantic
```

禁止:

```text
render/dag -> rawyaml
query -> rawyaml
render_index validationをDAG renderer内に埋め込む
State / Sequence / Wireframe rendererを先に作る
```

---

## 4. M4第1段の対象範囲

### 4.1 render_index validation

- `groups` 必須
- `groups` 空配列はerror
- group id必須
- group idは `[a-z0-9_]` のみ
- group idの `_` prefixは禁止
- `modules` 必須
- `modules` 空配列はerror
- module重複はerror
- moduleはtop-level moduleのみ許可
- moduleの `_` prefixは禁止
- uncovered moduleはwarningとして暗黙groupにする
- 暗黙groupは明示groupの後にアルファベット順で追加する

### 4.2 output placement

M4第1段で決定する出力先:

- main task DAG: `renders/{group-id}/dag-{task-id}.md`
- master index: `renders/index.md`
- group index: `renders/{group-id}/index.md`
- cross-cutting placeholder: `renders/_cross/`
- preview placeholder: `renders/_preview/`

M4第1段では、State / Sequence / Wireframeの実render生成は対象外。
ただし、V01-ADR-046のファイル名規則を壊さないよう、placement APIの拡張余地を残す。

---

## 5. 作る / 更新する想定ファイル

M4で新規作成する候補:

```text
internal/rawyaml/render_index.go
internal/render/placement/
  config.go
  resolver.go
  index.go
  resolver_test.go
```

M4で更新する候補:

```text
internal/rawyaml/node.go       # File.RenderIndex 追加
internal/source/loader.go      # render_index.yaml decode
internal/semantic/project.go   # filesByTopLevelModule相当が必要なら追加
internal/render/dag/...        # 既存renderer本体は原則触らない。必要なら薄い呼び出し側のみ
```

M4では以下を作らない。

```text
internal/cli/
internal/mcp/
State Diagram renderer
Sequence Diagram renderer
Wireframe renderer
Preview HTML renderer
```

---

## 6. タスク一覧

### 6.1 `rawyaml`: render_index.yaml decode

- [ ] `RenderIndex` structを追加する
- [ ] `RenderGroup` structを追加する
- [ ] `rawyaml.File` に `RenderIndex *RenderIndex` を追加する
- [ ] `source.loadFile` で `FileKindRenderIndex` の場合にdecodeする
- [ ] YAML decode失敗時はfile ID付きerrorにする

完了条件:

- [ ] UC-001の `render_index.yaml` をloadできる
- [ ] node file / view file / render_index file の既存分類が壊れていない

---

### 6.2 `render/placement`: group resolution

- [ ] raw projectからrender index configを取得する
- [ ] semantic.Projectからtop-level module一覧を取得する
- [ ] group id validationを実装する
- [ ] modules validationを実装する
- [ ] module重複errorを実装する
- [ ] uncovered module warningを実装する
- [ ] uncovered moduleから暗黙groupを作る
- [ ] group順序を安定化する
- [ ] top-level module -> group id のmapを作る

完了条件:

- [ ] `auth -> auth`
- [ ] `cart -> commerce`
- [ ] `order -> commerce`
- [ ] `payment -> commerce`
- [ ] `payment.webhooks -> commerce`
- [ ] `catalog -> catalog`
- [ ] `inventory -> catalog`

---

### 6.3 DAG output placement

- [ ] main taskのFileIDからtop-level moduleを求める
- [ ] group idを解決する
- [ ] DAG output pathを返す
- [ ] `dag-{task-id}.md` を返す
- [ ] nested moduleのtaskでも親top-level moduleでgroup解決する

完了条件:

- [ ] `auth.task.login -> auth/dag-login.md`
- [ ] `cart.task.validate_cart -> commerce/dag-validate_cart.md`
- [ ] `order.task.checkout -> commerce/dag-checkout.md`
- [ ] `order.task.process_order -> commerce/dag-process_order.md`

---

### 6.4 index generation skeleton

- [ ] master index markdownを生成する
- [ ] group index markdownを生成する
- [ ] M4第1段ではDAG entriesのみでもよい
- [ ] `_cross` / `_preview` はplaceholder rowまたは拡張余地として扱う

完了条件:

- [ ] UC-001のgroup一覧を決定できる
- [ ] auth / commerce / catalog のgroup indexを生成できる
- [ ] 既存fixtureの配置と矛盾しない

---

### 6.5 unit test

- [ ] UC-001 yaml rootをloadする
- [ ] semantic.Projectをbuildする
- [ ] placement resolverを作る
- [ ] group resolutionを検証する
- [ ] DAG output pathを検証する
- [ ] validation error caseを最小限追加する
- [ ] `go test ./...` を通す

---

## 7. M4第1段で実装しないもの

- CLI command設計
- 実ファイル書き出し
- renders/ directory clean
- State Diagram renderer
- Sequence Diagram renderer
- Wireframe renderer
- Preview HTML renderer
- ER renderer
- API Table renderer
- State / Sequence / Wireframe のgolden更新

---

## 8. 受け入れ条件

M4第1段完了条件:

- [x] `render_index.yaml` をdecodeできる
- [x] V01-ADR-045のvalidation errorを検出できる
- [x] uncovered moduleをwarning + 暗黙groupとして扱える
- [x] top-level moduleからgroupを解決できる
- [x] nested moduleが親top-level moduleのgroupに入る
- [x] DAG output pathを決定できる
- [x] UC-001の既存DAG fixture配置と一致する
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る

---

## 9. 推奨実装順

1. `rawyaml.RenderIndex` を追加する
2. `source.loadFile` で `render_index.yaml` をdecodeする
3. `internal/render/placement` package skeletonを作る
4. top-level module抽出を実装する
5. group validationを実装する
6. module -> group解決を実装する
7. DAG output pathを実装する
8. UC-001 unit testを追加する
9. validation error testを追加する
10. index markdown skeletonを追加する
11. `go fmt ./...`
12. `go test ./...`

---

## 10. 実装者向けメモ

M4は「出力の場所」を決めるmilestoneであり、各view rendererを増やすmilestoneではない。

既存のDAG rendererはMarkdownを返す責務に留め、どこへ書くかはplacement側へ分離する。
これにより、後続でCLIやState / Sequence rendererを追加しても、同じplacement resolverを使える。
