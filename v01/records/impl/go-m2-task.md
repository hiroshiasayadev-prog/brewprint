# Go M2 実装タスク

- **status**: completed
- **last_updated**: 2026-04-27
- **scope**: Go実装 Milestone 2 の実装チェックリスト（完了済み）

---

## 1. 目的

このファイルは、Milestone 2「DAG renderer の対象を広げる」を実装するための作業チェックリストである。

### M2完了メモ

M2は実装完了済みとして扱う。

- UC-001 の `cart.task.validate_cart` / `order.task.checkout` / `order.task.process_order` は Raw YAML → ResolvedProject → DAG renderer → golden test まで到達済み。
- 既存M1の `auth.task.login` golden testも維持済み。
- `go fmt ./...` / `go test ./...` は 2026-04-27 に通過済み。
- このファイル内の詳細タスク一覧は、M2実装前チェックリストとして残す。
- 完了判定は「8. 受け入れ条件」の完了済みサマリを正とする。

M2の最終ゴールは、M1で通した `auth.task.login` に加えて、UC-001の以下3本のDAGを Raw YAML → ResolvedProject → DAG renderer → golden test まで通すこと。

```text
docs/uc/001-ec-checkout-flow/yaml/cart/task/validate_cart.yaml
  ↓ source load / classify / decode
Raw YAML structs
  ↓ resolve / semantic build
ResolvedProject
  ↓ DAG renderer
docs/uc/001-ec-checkout-flow/renders/commerce/dag-validate_cart.md と比較
```

```text
docs/uc/001-ec-checkout-flow/yaml/order/task/checkout.yaml
  ↓ source load / classify / decode
Raw YAML structs
  ↓ resolve / semantic build
ResolvedProject
  ↓ DAG renderer
docs/uc/001-ec-checkout-flow/renders/commerce/dag-checkout.md と比較
```

```text
docs/uc/001-ec-checkout-flow/yaml/order/task/process_order.yaml
  ↓ source load / classify / decode
Raw YAML structs
  ↓ resolve / semantic build
ResolvedProject
  ↓ DAG renderer
docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_order.md と比較
```

---

## 2. 前提ドキュメント

実装前に最低限読むこと。

- `docs/doc-policy.md`
- `docs/impl/go-skeleton.md`
- `docs/impl/go-m1-task.md`
- `docs/TASKS.md` の Milestone 2
- `docs/adr/012-control-flow-nodes.md`
- `docs/adr/016-foreach-as-flow-construct.md`
- `docs/adr/023-control-flow-scope-and-branch-entry.md`
- `docs/adr/040-control-flow-step-wiring.md`
- `docs/adr/044-store-access-edge-labels.md`
- `docs/spec/edges.md` の `flow:` / `$` シジル / 制御フロースコープ範囲
- `docs/spec/nodes.md` の task / branch / fork / join 範囲
- `docs/spec/views/dag.md`
- `docs/uc/001-ec-checkout-flow/yaml/cart/task/validate_cart.yaml`
- `docs/uc/001-ec-checkout-flow/yaml/order/task/checkout.yaml`
- `docs/uc/001-ec-checkout-flow/yaml/order/task/process_order.yaml`
- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-validate_cart.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-checkout.md`
- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_order.md`

---

## 3. 実装原則

- M2もM1と同じく縦切りを優先する。
- M2の対象はDAG renderer拡張に限定する。
- rendererは引き続き `rawyaml` をimportしない。
- renderer内でRaw YAML走査・名前解決・reverse lookup index再構築をしない。
- `flow:` decodeは `rawyaml`、解決済みflowは `semantic` / `resolve` 側に置く。
- `foreach` はnode typeではなく `flow:` 制御構文として扱う。
- `branch` / `fork` / `join` はnodeとしてdecode / resolve / detail render対象にする。
- store access edgeはM1同様にV01-ADR-044の `read` / `write` / `read/write` ラベルを維持する。
- golden fixtureに合わせることを優先し、M3以降のQueryService向け汎用reference indexは作り込まない。

M2で守る境界:

```text
source      -> rawyaml
resolve     -> rawyaml, semantic
render/dag  -> semantic
```

禁止:

```text
render/dag -> rawyaml
semantic   -> rawyaml
rawyaml    -> yaml.Node
```

---

## 4. M2対象ファイルと期待パターン

| input YAML | main node | M2で検証する主機能 | golden |
|---|---|---|---|
| `cart/task/validate_cart.yaml` | `cart.task.validate_cart` | `foreach` / `$params.field` / `$item` / クロスモジュールstore read | `renders/commerce/dag-validate_cart.md` |
| `order/task/checkout.yaml` | `order.task.checkout` | `fork` / `join` / `branches[].steps[].params` / join params解決 | `renders/commerce/dag-checkout.md` |
| `order/task/process_order.yaml` | `order.task.process_order` | `branch` / `cases[].params` / floating node → `_end` | `renders/commerce/dag-process_order.md` |

---

## 5. 作る / 更新する想定ファイル

M1で作成済みのpackage構成を維持し、必要箇所だけ拡張する。

```text
internal/rawyaml/
  node.go       # branch / fork / join node field、flow root fieldの追加候補
  task.go       # 既存task field維持

internal/semantic/
  node.go       # branch / fork / join kindの追加候補
  task.go       # 既存task field維持
  project.go    # file単位flow index追加候補

internal/resolve/
  builder.go    # branch / fork / join node build、flow build
  symbols.go    # branch / fork / join symbol登録
  names.go      # flow内参照解決で必要なら拡張

internal/render/dag/
  viewmodel.go  # foreach / fork / join / branch用view model拡張
  renderer.go   # Mermaid / Tasks detail render拡張
  renderer_test.go # M2 golden test追加
```

必要なら新規ファイルを追加してよい。

```text
internal/rawyaml/flow.go
internal/semantic/flow.go
internal/render/dag/flow.go
```

M2では以下を作らない。

```text
internal/query/
internal/cli/
internal/mcp/
rawyaml/view.go
rawyaml/render_index.go
```

---

## 6. タスク一覧

### 6.1 `internal/rawyaml`: flow / control node decode

実装対象:

- [x] node file root structに `flow:` を追加する
- [x] `flow:` のentryをdecodeできるようにする
- [x] 通常step entryをdecodeする
- [x] foreach entryをdecodeする
- [x] fork entryをdecodeする
- [x] fork `branches[].steps[]` をdecodeする
- [x] branch entryをdecodeする
- [x] branch `cases[].params` をdecodeする
- [x] branch / fork / join nodeをdecode対象に追加する
- [x] `params` wiring mapを保持できるようにする
- [x] `$params.field` / `$item` / node ID文字列はRawの時点では未解決文字列として保持する

M2で扱うflow entry field:

```yaml
flow:
  - step: build_order
    params:
      cart_id: $params.cart_id

  - foreach: validate_item
    over: $params.cart_items
    mode: sequential
    params:
      cart_item: $item
    returns: validated_items

  - fork: parallel_processing
    branches:
      - steps:
          - step: reserve_inventory
            params:
              draft_order: build_order
    join: finalize_checkout

  - branch: route_by_inventory
    params:
      order: check_inventory
    cases:
      - label: in_stock
        step: confirm_order
        params:
          order: check_inventory
```

禁止:

- [x] `rawyaml` でflow参照解決をしない
- [x] `rawyaml` で `$params` / `$item` の型解決をしない
- [x] `rawyaml` に `*yaml.Node` を露出しない
- [x] foreachをnode typeとして復活させない

完了条件:

- [x] `validate_cart.yaml` の `foreach` をdecodeできる
- [x] `checkout.yaml` の `fork.branches[].steps[].params` をdecodeできる
- [x] `process_order.yaml` の `branch.cases[].params` をdecodeできる
- [x] branch / fork / join nodeのparams / returns / noteをdecodeできる

---

### 6.2 `internal/semantic`: control node / resolved flow model

実装対象:

- [x] branch node型を追加する
- [x] fork node型を追加する
- [x] join node型を追加する
- [x] node kindに branch / fork / join を追加する
- [x] resolved flow entryの最小表現を追加する
- [x] resolved step entryを表現できる
- [x] resolved foreach entryを表現できる
- [x] resolved fork entryを表現できる
- [x] resolved branch entryを表現できる
- [x] flowをFileID単位で参照できるindexを追加する

M2の `Project` 追加候補:

```go
type Project struct {
    // M1 fields...

    BranchesByQID map[QualifiedID]*Branch
    ForksByQID    map[QualifiedID]*Fork
    JoinsByQID    map[QualifiedID]*Join

    FlowByFile map[FileID][]FlowEntry
}
```

M2ではまだ作らないfield:

- `ReferencesBySource`
- `ReferencesByTarget`
- `TransitionsByStateEventGuard`
- `ActionsByTask`
- `ScenariosByID`
- `FilesByTopLevelModule`

完了条件:

- [x] DAG rendererが `semantic.Project` だけを読んでM2対象DAGを生成できる
- [x] `semantic` が `rawyaml` に依存しない
- [x] branch / fork / join が `## Tasks` 詳細セクションに出せる情報を持つ

---

### 6.3 `internal/resolve`: branch / fork / join symbol build

実装対象:

- [x] branch node QIDを生成する
- [x] fork node QIDを生成する
- [x] join node QIDを生成する
- [x] branch / fork / join を `NodesByQID` / `NodesByFile` に登録する
- [x] branch / fork / join のparamsをmodel参照へ解決する
- [x] join のreturnsをmodel参照へ解決する
- [x] duplicate QID検出をbranch / fork / joinにも適用する
- [x] `MainNodeByFile` は引き続きmain taskを指す

完了条件:

- [x] `order.branch.route_by_inventory` 相当のbranchがProject上に存在する
- [x] `order.fork.parallel_processing` 相当のforkがProject上に存在する
- [x] `order.join.finalize_checkout` 相当のjoinがProject上に存在する
- [x] branch / fork / join がTasks詳細セクションに出力可能な状態で解決される

※ 実際のQID形式は既存のsentinel方式に合わせること。上記は意図説明であり、既存実装の命名規則を優先する。

---

### 6.4 `internal/resolve`: flow wiring resolve

実装対象:

- [x] file内の `flow:` entryをResolvedProjectへ変換する
- [x] `step.params` の `$params.field` をmain task boundary paramへ解決する
- [x] `step.params` の node ID参照を同一ファイル内task returns assetへ解決する
- [x] foreach `over: $params.field` をboundary paramへ解決する
- [x] foreach `params` の `$item` をforeach item sourceとして表現する
- [x] foreach `returns` をmain task returns boundaryまたはcollect assetへ接続できる形で解決する
- [x] fork entryの `fork` IDをfork nodeへ解決する
- [x] fork `branches[].steps[].step` をtaskへ解決する
- [x] fork `branches[].steps[].params` を通常stepと同じ規則で解決する
- [x] fork `join` をjoin nodeへ解決する
- [x] join paramsをbranch終端stepの `returns.name` と同名一致で解決する
- [x] branch entryの `branch` IDをbranch nodeへ解決する
- [x] branch `params` をbranch node自身の判定入力として解決する
- [x] branch `cases[].step` をcase entry taskへ解決する
- [x] branch `cases[].params` をcase entry taskへのwiringとして解決する
- [x] branch case taskが後続から参照されない場合、floating nodeとして扱える情報を持たせる

参照解決ルール:

```text
flow params value resolution:
1. $params.<field>  -> main task params boundary asset
2. $item            -> foreach current item
3. <node id>        -> same-file node returns asset
4. <qualified id>   -> cross-file main node returns asset（M2で必要になった場合のみ）
```

M2で最低限必要な参照:

- `$params.cart_items`
- `$params.cart_id`
- `$params.shipping_address`
- `$params.order_id`
- `$item`
- `build_order`
- `check_inventory`

完了条件:

- [x] `validate_cart` の foreach overが `cart_items` boundaryに接続される
- [x] `validate_cart` の `$item` が `validate_item.params.cart_item` に接続される扱いになる
- [x] `checkout` の `build_order` returns `draft_order` が並列branch taskへ渡る
- [x] `checkout` の `reserve_inventory.returns.name = reserved` が `finalize_checkout.params.reserved` に接続される
- [x] `checkout` の `notify_payment_gateway.returns.name = notified` が `finalize_checkout.params.notified` に接続される
- [x] `process_order` の `branch.params.order` と `cases[].params.order` が別用途として解決される
- [x] `process_order` の `confirm_order` / `cancel_order` がfloating nodeとして扱われる

---

### 6.5 `internal/render/dag`: foreach render

実装対象:

- [x] foreach apply先taskを独立ノードではなくtaskノードとして描画する
- [x] apply先task labelに `↻` を付与する
- [x] `over` sourceからapply先taskへ `--"foreach"-->` のデータ線を描画する
- [x] foreach apply先taskへの制御線を描画する
- [x] foreach `returns` をmain task returns boundaryへ接続する
- [x] foreach apply先taskのreads / writes store accessを描画する
- [x] foreach apply先taskをTasks詳細セクションに含める

`validate_cart` の期待ポイント:

```mermaid
_start([Start]) ==> validate_item["↻ validate_item"]
cart_items --"foreach"--> validate_item
item_collection[(catalog.store.item_collection)] -- "read" --> validate_item
inventory_db[(catalog.store.inventory_db)] -- "read" --> validate_item
validate_item --> validated_items
validate_item ==> _end([End])
```

完了条件:

- [x] `docs/uc/001-ec-checkout-flow/renders/commerce/dag-validate_cart.md` とgolden一致する

---

### 6.6 `internal/render/dag`: fork / join render

実装対象:

- [x] fork nodeを `{{fork_id}}` で描画する
- [x] join nodeを `{{join_id}}` で描画する
- [x] fork前の通常stepからfork nodeへ制御線を描画する
- [x] fork nodeから各branch先頭stepへ `== "parallel" ==>` を描画する
- [x] branch taskへのparams data線を描画する
- [x] branch taskのreturns assetを描画する
- [x] branch終端taskからjoin nodeへ制御線を描画する
- [x] branch終端taskのreturns assetからjoin nodeへdata線を描画する
- [x] join returnsからmain task returns boundaryへdata線を描画する
- [x] fork / join をTasks詳細セクションに含める
- [x] fork / join に `forkNode` classを付与する

`checkout` の期待ポイント:

```mermaid
build_order ==> parallel_processing{{parallel_processing}}
draft_order --> reserve_inventory[reserve_inventory]
draft_order --> notify_payment_gateway[notify_payment_gateway]
parallel_processing == "parallel" ==> reserve_inventory
parallel_processing == "parallel" ==> notify_payment_gateway
reserve_inventory ==> finalize_checkout{{finalize_checkout}}
notify_payment_gateway ==> finalize_checkout
reserved --> finalize_checkout
notified --> finalize_checkout
finalize_checkout --> pending_order
finalize_checkout ==> _end([End])
```

完了条件:

- [x] `docs/uc/001-ec-checkout-flow/renders/commerce/dag-checkout.md` とgolden一致する

---

### 6.7 `internal/render/dag`: branch / cases render

実装対象:

- [x] branch nodeを `{branch_id}` で描画する
- [x] branch判定入力のdata線をbranch nodeへ描画する
- [x] `cases[].params` のdata線をcase entry taskへ描画する
- [x] branch nodeからcase entry taskへ `== "{label}" ==>` を描画する
- [x] case entry taskのstore accessを描画する
- [x] 後続参照のないcase entry taskをfloating nodeとして `_end` へ接続する
- [x] branch nodeをTasks詳細セクションに含める
- [x] branch nodeに `branchNode` classを付与する

`process_order` の期待ポイント:

```mermaid
check_inventory ==> route_by_inventory{route_by_inventory}
order_asset --> route_by_inventory
order_asset --> confirm_order[confirm_order]
order_asset --> cancel_order[cancel_order]
route_by_inventory == "in_stock" ==> confirm_order
route_by_inventory == "out_of_stock" ==> cancel_order
confirm_order -- "write" --> order_db
cancel_order -- "write" --> order_db
confirm_order ==> _end([End])
cancel_order ==> _end
```

注意:

- goldenでは `check_inventory.returns.name = order` のasset Mermaid IDが `order_asset` になっている。
- Mermaid IDがnode IDや予約語と衝突する場合は、既存rendererのID escape / sanitize方針に合わせる。
- `branch.params` と `cases[].params` は同じsourceを参照しても意味が違う。前者はbranch判定入力、後者はcase task実行入力。

完了条件:

- [x] `docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_order.md` とgolden一致する

---

### 6.8 Golden test追加

実装対象:

- [x] `internal/render/dag/renderer_test.go` にM2ケースを追加する
- [x] UC-001 yaml rootをloadする
- [x] ResolvedProjectをbuildする
- [x] `cart/task/validate_cart.yaml` のmain taskをrender対象にする
- [x] `order/task/checkout.yaml` のmain taskをrender対象にする
- [x] `order/task/process_order.yaml` のmain taskをrender対象にする
- [x] actual Markdownを生成する
- [x] 各golden fixtureと比較する

対象golden:

```text
docs/uc/001-ec-checkout-flow/renders/commerce/dag-validate_cart.md
docs/uc/001-ec-checkout-flow/renders/commerce/dag-checkout.md
docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_order.md
```

完了条件:

- [x] 既存の `dag-login.md` golden testが引き続き通る
- [x] `dag-validate_cart.md` golden testが通る
- [x] `dag-checkout.md` golden testが通る
- [x] `dag-process_order.md` golden testが通る
- [x] `go test ./...` が通る

---

## 7. M2で実装しないもの

- QueryService
- MCP server wrapper
- CLI command設計
- `brewprint render` コマンド
- view file decode
- render_index.yaml decode / validation
- output placement / renders index生成
- API Table renderer
- ER renderer
- State Diagram renderer
- Sequence Diagram renderer
- Wireframe renderer
- scenario decode
- state / event / transition semantic build
- branch内部assetの外部参照許可
- fork branchesの旧短縮形 `branches: - [step]` サポート
- foreachをnode typeとして扱う実装
- import boundary lint導入

---

## 8. 受け入れ条件

M2完了条件:

- [x] `go test ./...` が通る
- [x] 既存M1の `auth.task.login` golden testが通る
- [x] `cart.task.validate_cart` をRaw YAMLからResolvedProjectへbuildできる
- [x] `cart.task.validate_cart` のforeachをDAG renderできる
- [x] foreach apply先taskに `↻` を表示できる
- [x] foreach `over: $params.cart_items` をboundary paramからのforeach data線として描画できる
- [x] foreach `params.cart_item: $item` を解釈できる
- [x] `docs/uc/001-ec-checkout-flow/renders/commerce/dag-validate_cart.md` とgolden一致する
- [x] `order.task.checkout` をRaw YAMLからResolvedProjectへbuildできる
- [x] fork / join nodeをdecode / resolve / renderできる
- [x] `branches[].steps[].params` を解決し、branch taskへのdata線を描画できる
- [x] join.params をbranch終端stepの `returns.name` で解決できる
- [x] `docs/uc/001-ec-checkout-flow/renders/commerce/dag-checkout.md` とgolden一致する
- [x] `order.task.process_order` をRaw YAMLからResolvedProjectへbuildできる
- [x] branch nodeをdecode / resolve / renderできる
- [x] `branch.params` と `cases[].params` を別用途として扱える
- [x] branch case taskのfloating node → `_end` を描画できる
- [x] `docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_order.md` とgolden一致する
- [x] DAG rendererが `rawyaml` をimportしていない
- [x] `semantic` が `rawyaml` をimportしていない
- [x] `rawyaml` に `*yaml.Node` が露出していない

---

## 9. 推奨実装順

1. `rawyaml` に `flow:` decode構造を追加する
2. `rawyaml` / `semantic` に branch / fork / join nodeを追加する
3. `resolve` で branch / fork / join のsymbol登録とnode buildを追加する
4. 通常stepのflow解決を追加する
5. foreach flow解決を追加する
6. DAG rendererでforeachを描画し、`dag-validate_cart.md` goldenを通す
7. fork / join flow解決を追加する
8. DAG rendererでfork / joinを描画し、`dag-checkout.md` goldenを通す
9. branch flow解決を追加する
10. floating node判定を追加する
11. DAG rendererでbranch / casesを描画し、`dag-process_order.md` goldenを通す
12. `## Tasks` 詳細セクションに branch / fork / join を含める
13. 既存 `dag-login.md` goldenが壊れていないことを確認する
14. `go fmt ./...`
15. `go test ./...`

---

## 10. 実装後に確認すること

- [x] `go fmt ./...`
- [x] `go test ./...`
- [x] `git diff` で意図しないdocs / fixture変更がないこと
- [x] `internal/render/dag` が `internal/rawyaml` をimportしていないこと
- [x] `internal/semantic` が `internal/rawyaml` をimportしていないこと
- [x] `internal/rawyaml` に `*yaml.Node` が露出していないこと
- [x] M1の `dag-login.md` golden testが維持されていること
- [x] M2完了後、`docs/TASKS.md` のMilestone 2進捗更新を検討する
- [x] M2完了後、このファイルのstatusを `completed` に更新する

---

## 11. 実装者向けメモ

M2は「DAG renderer対象拡張」のmilestoneであり、QueryServiceやMCPのために汎用reference modelを先回りして作り込む段階ではない。

ただし、M2で初めて `flow:` を本格的に読むため、以下の理解が重要。

```text
nodes: 実行単位・制御ノードの定義
flow:  実行順序とwiringの定義
```

`foreach` は `nodes:` には存在しない。`flow:` 内の制御構文であり、apply先taskを `↻` 付きで描画する。

`branch.params` はbranch nodeの判定入力。`cases[].params` はcase entry taskへの実行入力。ここを混ぜるとDAGのdata線が崩れる。

`fork` の `join.params` は、各branch終端stepの `returns.name` から解決する。`fork.params` の暗黙伝播は採用しない。

branch / fork / foreach の内部assetは制御フロースコープ外から直接参照しない。外に出す必要がある場合はstore経由にする。ただしM2対象fixtureでは、branchはfloating終端、forkはjoinで集約、foreachはcollect returnsとして扱う。
