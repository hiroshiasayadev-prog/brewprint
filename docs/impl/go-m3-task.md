# Go M3 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: Go実装 Milestone 3 第1段の QueryService vertical slice 実装チェックリスト（完了済み）

---

## 1. 目的

このファイルは、Milestone 3「QueryService vertical slice」を実装するための作業チェックリストである。

M3の目的は、rendererの次に、LLM向け問い合わせAPIをMCP transport抜きの通常Go APIとして通すこと。

### M3第1段完了メモ

M3第1段は実装完了済みとして扱う。

- `internal/query` を作成し、QueryServiceを通常のGo APIとして実装済み。
- `GetSignature` は task / model / store / branch / fork / join で到達済み。
- `GetReferences` は `referencesBySource` / `referencesByTarget` からdirect referencesのみ返す形で到達済み。
- `Inspect` は task / model / store で到達済み。
- UC-001に対するQueryService unit testを追加済み。
- `go fmt ./...` / `go test ./...` は 2026-04-28 に通過済み。
- state / event / transition / scenario と `actionsByTask` / `transitionsByStateEventGuard` / `scenariosByID` は第2段候補として残す。

最初の成功条件は、UC-001を以下の経路で読み込み、QueryServiceをunit testから直接叩ける状態にすること。

```text
docs/uc/001-ec-checkout-flow/yaml/
  ↓ source load / classify / decode
Raw YAML structs
  ↓ resolve / semantic build
semantic.Project
  ↓ internal/query.QueryService
GetSignature / GetReferences / Inspect
  ↓ unit test
```

M3ではMCP server wrapperを作らない。
QueryServiceは将来MCP toolから呼ばれる前提のGo APIだが、M3の検証はGo unit testで直接行う。

---

## 2. 前提ドキュメント

実装前に最低限読むこと。

- `docs/doc-policy.md`
- `docs/impl/go-skeleton.md`
- `docs/impl/go-m1-task.md`
- `docs/impl/go-m2-task.md`
- `docs/TASKS.md` の Milestone 3
- `docs/adr/047-go-semantic-model-query-layer-boundary.md`
- `docs/adr/048-resolved-project-index-strategy.md`
- `docs/adr/049-mcp-query-reference-vocabulary.md`
- `docs/spec/mcp.md` の以下の範囲
  - Common schema
  - Reference schema
  - `get_signature`
  - `get_references`
  - `inspect`
- `docs/spec/nodes.md` の task / model / store / branch / fork / join 範囲
- `docs/uc/001-ec-checkout-flow/yaml/auth/task/login.yaml`
- `docs/uc/001-ec-checkout-flow/yaml/cart/task/validate_cart.yaml`
- `docs/uc/001-ec-checkout-flow/yaml/order/task/checkout.yaml`
- `docs/uc/001-ec-checkout-flow/yaml/order/task/process_order.yaml`

---

## 3. 実装原則

- M3も縦切りを優先する。
- QueryServiceは `semantic.Project` を入力境界とする。
- QueryServiceは `rawyaml` をimportしない。
- QueryService内でRaw YAML走査をしない。
- QueryService内で名前解決を再実装しない。
- QueryService内でその場限りのreverse lookupを再構築しない。
- `GetReferences` はMCP v1方針に従いdirect referencesのみ返す。
- `Inspect` は task / model / store を最初の対象に限定する。
- state / event / transition / scenario はM3第2段または後続milestoneで扱う。
- MCP server wrapperはM3では作らない。

M3で守る境界:

```text
source      -> rawyaml
resolve     -> rawyaml, semantic
render/dag  -> semantic
query       -> semantic
```

禁止:

```text
query -> rawyaml
query内でRaw YAML走査
query内で名前解決の再実装
MCP server wrapperを先に作る
```

---

## 4. M3第1段の対象範囲

M3第1段は、DAG rendererで既に扱っているsemantic objectをQueryServiceから読めるようにする。

### 4.1 `GetSignature` 対象kind

- task
- model
- store
- branch
- fork
- join

### 4.2 `GetReferences` 対象reference kind

M3第1段では、以下のdirect referencesのみを対象にする。

- `param_model`
- `return_model`
- `produces_asset`
- `reads`
- `writes`
- `store_of`
- `field_type`
- `field_fk`

M3第1段では返さないreference kind:

- `transition_event`
- `transition_from`
- `transition_to`
- `transition_action`
- `event_payload`
- `event_actor`
- `event_watches`
- `scenario_state_file`
- `scenario_step_transition`

理由: state / event / transition / scenario decodeとsemantic buildはM5 Sequence / State renderer範囲を食い始めるため。

### 4.3 `Inspect` 対象kind

- task
- model
- store

`Inspect` は `GetSignature` と `GetReferences` を組み合わせて返す。
ただし、taskの `members.flow.entries` はMCP v1のdraft schemaに合わせ、最低限の概略順序だけを返す。

---

## 5. 作る / 更新する想定ファイル

M3で新規作成するpackage:

```text
internal/query/
  service.go       # QueryService本体 / constructor
  types.go         # request / response / common DTO
  signature.go     # GetSignature
  references.go    # GetReferences
  inspect.go       # Inspect
  service_test.go  # UC-001 vertical slice test
```

M3で更新する候補:

```text
internal/semantic/
  project.go       # ReferencesBySource / ReferencesByTarget の追加
  reference.go     # semantic.Reference / ReferenceKind / Direction の追加候補
  id.go            # 必要ならObjectID系helperを追加

internal/resolve/
  builder.go       # reference index初期化 / build呼び出し
  references.go    # direct reference index buildの追加候補
```

必要ならquery test helperを追加してよい。

```text
internal/query/query_test.go
internal/testutil/project/project.go
```

M3では以下を作らない。

```text
internal/mcp/
internal/cli/
rawyaml/view.go
rawyaml/render_index.go
state / event / scenario 専用semantic実装
```

---

## 6. タスク一覧

### 6.1 `internal/semantic`: Reference modelを追加する

実装対象:

- [ ] `ReferenceKind` を定義する
- [ ] `ReferenceDirection` を定義する
- [ ] `ObjectRef` またはQueryService向けに変換しやすいsemantic refを定義する
- [ ] `Reference` を定義する
- [ ] `Project` に `ReferencesBySource` を追加する
- [ ] `Project` に `ReferencesByTarget` を追加する
- [ ] `NewProject()` でreference indexを初期化する

M3第1段で扱うkind:

```text
param_model
return_model
produces_asset
reads
writes
store_of
field_type
field_fk
```

注意:

- `ReferencesBySource` / `ReferencesByTarget` のkeyは解決済みIDにする。
- Raw YAMLに書かれた未解決文字列をkeyにしない。
- primitive型をreference targetにする場合は、node QIDではなくprimitive objectとして表現する。
- model field / field fk のsource / target表現はQueryService DTOへ変換しやすい形にする。

完了条件:

- [ ] `semantic` が `rawyaml` に依存しない
- [ ] Project初期化直後にreference mapがnilでない
- [ ] M3第1段のdirect referencesを表現できる

---

### 6.2 `internal/resolve`: direct reference indexを構築する

実装対象:

- [ ] task paramsから `param_model` を作る
- [ ] task returnsから `return_model` を作る
- [ ] task returns assetから `produces_asset` を作る
- [ ] task readsから `reads` を作る
- [ ] task writesから `writes` を作る
- [ ] store.ofから `store_of` を作る
- [ ] model fieldsから `field_type` を作る
- [ ] model fieldsのfkから `field_fk` を作る
- [ ] branch paramsから `param_model` を作る
- [ ] join paramsから `param_model` を作る
- [ ] join returnsから `return_model` / `produces_asset` を作る
- [ ] `ReferencesBySource` / `ReferencesByTarget` の両方へ登録する
- [ ] referenceの返却順が安定するよう、QueryService側またはbuild側でsortする

M3第1段でやらないこと:

- [ ] transition / event / scenario reference build
- [ ] transitive reference graph build
- [ ] flow wiringを `GetReferences` に含める

完了条件:

- [ ] `auth.task.login` のparams / returns / reads / writes / produces_assetがindexされる
- [ ] `auth.store.user_db` へのincoming `reads` をindexから取得できる
- [ ] `auth.model.login_form` へのincoming `param_model` をindexから取得できる
- [ ] `order.join.finalize_checkout` のparams / returns / produces_assetをindexできる
- [ ] `go test ./...` が通る

---

### 6.3 `internal/query`: package skeletonを作る

実装対象:

- [ ] `internal/query` directoryを作る
- [ ] `Service` structを作る
- [ ] `NewService(project *semantic.Project) *Service` を作る
- [ ] nil projectに対する扱いを決める
- [ ] request / response DTOを定義する
- [ ] selectorはM3第1段では `id` 指定を必須にする

M3第1段の最小API候補:

```go
type Service struct {
    project *semantic.Project
}

func NewService(project *semantic.Project) *Service

func (s *Service) GetSignature(req GetSignatureRequest) (GetSignatureResponse, error)
func (s *Service) GetReferences(req GetReferencesRequest) (GetReferencesResponse, error)
func (s *Service) Inspect(req InspectRequest) (InspectResponse, error)
```

注意:

- MCP transport向けのJSON schemaをそのままGo型に固定しすぎない。
- ただしfield名と語彙は `docs/spec/mcp.md` と揃える。
- errorはまず通常のGo errorでよい。MCP-level error変換はMCP wrapper側で行う。

完了条件:

- [ ] `internal/query` が `internal/semantic` だけに依存する
- [ ] `internal/query` が `internal/rawyaml` をimportしない
- [ ] QueryServiceをunit testから直接作れる

---

### 6.4 `GetSignature` を通す

実装対象:

- [ ] selector.idからnodeを取得する
- [ ] task signatureを返す
- [ ] model signatureを返す
- [ ] store signatureを返す
- [ ] branch signatureを返す
- [ ] fork signatureを返す
- [ ] join signatureを返す
- [ ] object identityを返す
- [ ] docに `note` を返す
- [ ] sourceは保持情報がある範囲で返す
- [ ] unsupported kind / unknown idをerrorにする

M3第1段のtask signature:

- `main`
- `params`
- `returns`
- `reads`
- `writes`
- `endpoint`（endpoint taskのみ）

M3第1段のmodel signature:

- `model_kind`
- `fields`

M3第1段のstore signature:

- `store_kind`
- `of`

M3第1段のbranch / fork / join signature:

- `params`
- joinのみ `returns`

完了条件:

- [ ] `GetSignature(auth.task.login)` がtask外形を返す
- [ ] `GetSignature(auth.model.login_form)` がmodel fieldsを返す
- [ ] `GetSignature(auth.store.user_db)` がstore kind / ofを返す
- [ ] `GetSignature(order.branch.route_by_inventory)` 相当がbranch paramsを返す
- [ ] `GetSignature(order.fork.parallel_processing)` 相当がfork paramsを返す
- [ ] `GetSignature(order.join.finalize_checkout)` 相当がjoin params / returnsを返す

※ branch / fork / join の実際のQID形式はM2実装済みのsentinel方式に合わせること。

---

### 6.5 `GetReferences` をdirect referencesだけで通す

実装対象:

- [ ] selector.idからobjectを取得する
- [ ] direction省略時は `out` として扱う
- [ ] `direction=out` で `ReferencesBySource` を読む
- [ ] `direction=in` で `ReferencesByTarget` を読む
- [ ] `direction=both` で両方を読む
- [ ] `kinds` filterを実装する
- [ ] `depth` は常に `1` を返す
- [ ] 返却順を安定化する
- [ ] diagnosticsは空配列を返す

M3第1段で返すkind:

```text
param_model
return_model
produces_asset
reads
writes
store_of
field_type
field_fk
```

禁止:

- [ ] transitive closureを辿らない
- [ ] flow wiringをreferencesに混ぜない
- [ ] query package内でproject全体を走査してreferenceを作らない

完了条件:

- [ ] `GetReferences(auth.task.login, out)` がparam / return / asset / reads / writesを返す
- [ ] `GetReferences(auth.store.user_db, in)` がloginからのreadsを返す
- [ ] `GetReferences(auth.model.login_form, in)` がloginからのparam_modelを返す
- [ ] kind filterで `reads` のみ取得できる
- [ ] `direction=both` がin/outを返す

---

### 6.6 `Inspect` を task / model / store で通す

実装対象:

- [ ] task inspectを実装する
- [ ] model inspectを実装する
- [ ] store inspectを実装する
- [ ] common output shapeに合わせる
- [ ] signatureは `GetSignature` 相当の値を使う
- [ ] referencesは `GetReferences(direction=both)` 相当の主要referenceを使う
- [ ] task membersにreturns assetを含める
- [ ] task membersに同一file内sub task概要を含める
- [ ] task membersにflow概要を含める
- [ ] model membersはM3第1段ではfields中心でよい
- [ ] store membersにof model概要を含める
- [ ] unsupported kindはerrorにする

M3第1段でやらないこと:

- [ ] state inspect
- [ ] event inspect
- [ ] scenario inspect
- [ ] transition action reverse lookup
- [ ] scenario step解決

完了条件:

- [ ] `Inspect(auth.task.login)` がsignature / asset / referencesを返す
- [ ] `Inspect(order.task.checkout)` がflow概要とsub task概要を返す
- [ ] `Inspect(auth.model.login_form)` がfieldsとincoming referencesを返す
- [ ] `Inspect(auth.store.user_db)` がof model概要とincoming reads/writesを返す

---

### 6.7 unit testでUC-001を直接叩く

実装対象:

- [ ] UC-001 yaml rootをloadする
- [ ] semantic.Projectをbuildする
- [ ] QueryServiceを作る
- [ ] `GetSignature` testを追加する
- [ ] `GetReferences` testを追加する
- [ ] `Inspect` testを追加する
- [ ] 境界違反防止として `internal/query` のimportを目視または簡易testで確認する

最小test case:

```text
GetSignature:
  - auth.task.login
  - auth.model.login_form
  - auth.store.user_db
  - order.task.checkout
  - order.join.finalize_checkout

GetReferences:
  - auth.task.login direction=out
  - auth.store.user_db direction=in
  - auth.model.login_form direction=in
  - order.task.checkout direction=both

Inspect:
  - auth.task.login
  - order.task.checkout
  - auth.model.login_form
  - auth.store.user_db
```

完了条件:

- [ ] `go test ./...` が通る
- [ ] M1/M2のDAG golden testが壊れていない
- [ ] QueryService testがMCP wrapperなしで通る

---

## 7. M3第1段で実装しないもの

- MCP server wrapper
- CLI command設計
- `brewprint query` コマンド
- `list_endpoints`
- `get_source`
- `get_doc`
- transitive references
- reference tree
- State Diagram renderer
- Sequence Diagram renderer
- ER renderer
- API Table renderer
- Wireframe renderer
- state / event / transition / scenario decode
- `transition_event`
- `transition_from`
- `transition_to`
- `transition_action`
- `event_payload`
- `event_actor`
- `event_watches`
- `scenario_state_file`
- `scenario_step_transition`
- `actionsByTask`
- `transitionsByStateEventGuard`
- `scenariosByID`
- import boundary lint導入

---

## 8. 受け入れ条件

M3第1段完了条件:

- [x] `internal/query` が作成されている
- [x] QueryServiceが `semantic.Project` を入力として使う
- [x] QueryServiceが `rawyaml` をimportしていない
- [x] QueryService内で名前解決を再実装していない
- [x] `semantic.Project` に `ReferencesBySource` / `ReferencesByTarget` がある
- [x] resolve pipelineでM3第1段のdirect referencesが構築される
- [x] `GetSignature` が task / model / store / branch / fork / join を返せる
- [x] `GetReferences` がdirect referencesのみを返す
- [x] `GetReferences` が direction / kinds filter を扱える
- [x] `Inspect` が task / model / store を返せる
- [x] UC-001に対するQueryService unit testが通る
- [x] 既存M1/M2のDAG golden testが通る
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る

---

## 9. 推奨実装順

1. `internal/query` package skeletonを作る
2. `GetSignature` のDTOとService methodを作る
3. `GetSignature` を task / model / store で通す
4. `GetSignature` を branch / fork / join に広げる
5. `semantic.Reference` と `ReferencesBySource` / `ReferencesByTarget` を追加する
6. resolve pipelineで task params / returns / produces_asset / reads / writes をindexする
7. resolve pipelineで store_of / field_type / field_fk をindexする
8. resolve pipelineで branch / join params、join returns / assetをindexする
9. `GetReferences` を `direction=out` で通す
10. `GetReferences` を `direction=in` / `both` / `kinds` filterに広げる
11. `Inspect(task)` を signature / references / assetから通す
12. `Inspect(task)` にsub task / flow概要を足す
13. `Inspect(model)` を通す
14. `Inspect(store)` を通す
15. UC-001 QueryService unit testを整理する
16. `go fmt ./...`
17. `go test ./...`
18. import境界を確認する

---

## 10. 実装後に確認すること

- [ ] `go fmt ./...`
- [ ] `go test ./...`
- [ ] `git diff` で意図しないdocs / fixture変更がないこと
- [ ] `internal/query` が `internal/rawyaml` をimportしていないこと
- [ ] `internal/query` が名前解決を再実装していないこと
- [ ] `internal/query` がMCP protocolに依存していないこと
- [ ] `GetReferences` がdirect referencesのみを返していること
- [ ] state / event / scenario 実装をM3第1段に混ぜていないこと
- [ ] M3完了後、`docs/TASKS.md` のMilestone 3進捗更新を検討する
- [ ] M3完了後、このファイルのstatusを `completed` に更新する

---

## 11. 第2段で広げる候補

M3第1段完了後、必要になった場合のみ以下を検討する。

- event decode / semantic build
- state decode / semantic build
- transition semantic build
- scenario decode / semantic build
- `transitionsByStateEventGuard`
- `actionsByTask`
- `scenariosByID`
- `transition_event`
- `transition_from`
- `transition_to`
- `transition_action`
- `event_payload`
- `event_actor`
- `event_watches`
- `scenario_state_file`
- `scenario_step_transition`
- state / event / scenario の `GetSignature`
- state / event / scenario の `Inspect`

ただし、これらはM5のState / Sequence renderer範囲と強く重なるため、必要性が出るまで先取りしない。

---

## 12. 実装者向けメモ

M3は「LLM向け問い合わせAPIをGo APIとして通す」milestoneであり、MCP serverを作るmilestoneではない。

QueryServiceは `semantic.Project` の読み取り専用問い合わせ層にする。
QueryServiceでRaw YAMLを読むと、rendererとqueryで意味解釈が分岐するため禁止する。

`GetReferences` は便利に見えるため、transitive dependency traversalを入れたくなりやすいが、MCP v1ではdirect referencesだけで十分。
特にflow wiringはDAG内部の局所構造として扱い、`get_references` には混ぜない。

最初は `GetSignature(auth.task.login)` を通す。
その後、reference indexをsemantic.Projectに足して、`GetReferences` と `Inspect` を同じ材料から組み立てる。
