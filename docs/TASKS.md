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
  - [x] **golden fixture 再配置**（ADR-043 / ADR-046）
    - [x] `render_index.yaml` 追加
    - [x] `renders/{group}/` へ DAG / State / Sequence / Wireframe を配置
    - [x] `_cross/` へ ER / API を配置
    - [x] `_preview/wireframe.html` を配置
    - [x] README / coverage / legacy docs 表記を追随
    - [x] Go renderer 実装後に `renders/` を正式 golden fixture として再生成・検証
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

- [x] **foreach をDAG rendererに追加する**
  - `cart.task.validate_cart`
  - `foreach.over: $params.field`
  - `$item`
  - apply先taskの `↻` 表示
  - `docs/uc/001-ec-checkout-flow/renders/commerce/dag-validate_cart.md` とgolden一致

- [x] **fork / join をDAG rendererに追加する**
  - `order.task.checkout`
  - ADR-040 の `branches[].steps[].params`
  - join.params の branch終端step `returns.name` 解決
  - `docs/uc/001-ec-checkout-flow/renders/commerce/dag-checkout.md` とgolden一致

- [x] **branch / cases をDAG rendererに追加する**
  - `order.task.process_order`
  - `branch.params` と `cases[].params` の分離
  - floating node → `_end`
  - 制御フロースコープ（ADR-023）
  - `docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_order.md` とgolden一致

- [x] **M2完了確認**
  - 既存M1の `auth.task.login` golden test維持
  - `go fmt ./...` / `go test ./...` 通過（2026-04-27）

### Milestone 3: QueryService を通す

- [x] **QueryService vertical slice 第1段を実装する**
  - `GetSignature`
  - `GetReferences` はdirect referencesのみ
  - `Inspect`
  - unit testから直接叩けるようにする
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **direct reference index を整備する**
  - ADR-048 / ADR-049に準拠して実装する
  - referencesBySource
  - referencesByTarget
  - tasksReadingStore
  - tasksWritingStore

- [x] **state / event / scenario 系reverse lookup indexを第2段で実装する**
  - `scenariosByID` は既に実装済み
  - `transitionsByStateEventGuard` を追加
  - `transitionsByStateEvent` を候補列挙用補助indexとして追加
  - `actionsByTask` を追加
  - sequence scenario step解決を transition index lookup に変更
  - `inspect(task)` に `members.action_transitions` を追加
  - UC-001で transition index / actionsByTask / inspect action_transitions をtest固定
  - `go test ./...` 通過済み（2026-04-29、M9-1差分後）

- [x] **transition / event direct referencesをGetReferencesへ載せる**
  - MCP specの reference kind に合わせて `transition_from` / `transition_event` / `transition_to` / `transition_action` を追加する
  - `event_payload` / `event_actor` / `event_watches` を追加する
  - transition endpointの `state_file` / `from` / `on` / `to` / `guard` / `action` を返す
  - UC-001で task / event / model / store / state からのincoming/outgoing referenceをtest固定する
  - `gofmt ./...` / `go test ./...` 通過済み（2026-04-29、M9-2差分後）

- [x] **inspect(state/event) を実装する**
  - state inspectで incoming_transitions / outgoing_transitions を返す
  - event inspectで triggering_transitions / sequence_hints を返す
  - GetSignatureも state / event に対応
  - UC-001で state/event signature / inspect をtest固定する
  - `gofmt ./...` / `go test ./...` 通過済み（2026-04-29、M9-3差分後）

- [x] **inspect(scenario) を実装する**
  - selectorで sequence scenario view object を解決する
  - scenario signatureで id / title / state_file を返す
  - members.steps に resolved transition / action task / guard exact match結果を返す
  - scenario_state_file / scenario_step_transition reference を追加する
  - UC-001の checkout_flow / payment_webhook_flow でtest固定する
  - MCP selector schemaを `object` / `kind` / `file` / `local_id` 対応へ拡張する
  - `gofmt ./...` / `go test ./...` 通過済み（2026-04-29、M9-4差分後）

### Milestone 4: render_index / output placement を実装する

- [x] **render_index.yaml validation 第1段を実装する**
  - ADR-045に準拠
  - group id命名規則
  - module重複禁止
  - uncovered module の warning
  - nested module の親group所属
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **renders/ 出力配置 第1段を実装する**
  - ADR-043 / ADR-046に準拠
  - `renders/{group}/...`
  - `_cross/`
  - `_preview/`
  - master index / group index skeleton

### Milestone 5: 残りrendererを順に実装する

- [x] **State Diagram renderer 第1段を実装する**
  - guard分岐の choice pseudostate（ADR-035）
  - FSMファイル単位の `state-{fsm-id}.md`（ADR-046）
  - UC-001 golden test 3本
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **Sequence Diagram renderer を実装する**
  - event source別矢印ルール（ADR-036）
  - actionなしtransition（ADR-037）
  - sub task reads/writes traversal（ADR-038）
  - step index prefix / DB片方向message（ADR-041）
  - UC-001 golden test 2本
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **ER renderer を実装する**
  - 横断view YAML（ADR-039）
  - cross module FK
  - `store.kind: db` から model fields を辿る
  - UC-001 golden test 1本
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **API Table renderer を実装する**
  - `as: api_table`
  - `http_root_path`
  - endpoint task の leaf path 合成（ADR-028）
  - UC-001 golden test 1本
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **Wireframe renderer を実装する**
  - HTML fragment
  - `main` container / `layout` object（ADR-042）
  - fixed CSS profile
  - preview harness
  - UC-001 golden test fragment 4本 + preview 1本
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

### Milestone 6: MCP server wrapper を実装する

- [x] **MCP server wrapper を実装する**
  - QueryServiceをMCP toolとして公開する
  - JSON tool dispatch adapter 第1段完了
    - `get_signature`
    - `get_references`
    - `inspect`
    - `list_endpoints`
  - unknown / invalid args / not found のJSON error envelope
  - `internal/mcp -> query` 境界確認済み
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M6-2時点）
  - MCP stdio JSON-RPC 第1段完了
    - `initialize`
    - `tools/list`
    - `tools/call`
    - `brewprint mcp --yaml-root <path>`
  - MCP実用化仕上げ完了
    - QueryService response JSON field名をMCP spec寄りに整理
    - `tools/list` に `inputSchema` を追加
    - `tools/call` responseを `content: [{type:"text", text:"..."}]` / `isError` 形式に整理
    - `unsupported_direction` / `unsupported_detail` のtool error分類を追加
    - CLI usage / `--yaml-root` errorを整理
    - 起動例docを追加
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M6-3差分後）
  - renderer / QueryService / MCP wrapper の責務が混ざっていないことを確認する

### Milestone 7: validation / diagnostics を強化する

- [x] **semantic validation / diagnostics 第1段を実装する**
  - `semantic.Diagnostic` に `code` を追加
  - resolver diagnostic helperをcode対応
  - model / store / task / control / event の主要参照validationを追加
  - unresolved reads/writesはvalidation層のcode付きerrorに一本化
  - invalid rawyaml.Project fixtureのunit testを追加
  - UC-001 fixtureのクロスモジュール短縮参照をADR-027準拠に修正
  - API Table goldenを更新
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-1差分後）

- [x] **diagnostics CLI 第1段を実装する**
  - `brewprint validate --yaml-root <path>` を追加
  - diagnosticsなしなら `ok` をstdoutへ出力
  - diagnosticsを `severity code file: message` 形式で表示
  - error diagnosticsが1件以上あればnon-zero error
  - MCP起動時のsemantic diagnostic errorにもcodeを含める
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-2差分後）

- [x] **validate failure path testを追加する**
  - `t.TempDir()` でinvalid YAML rootを作成
  - 未解決 `params.model` を含む最小node YAMLを配置
  - stdoutのcode付きdiagnosticを固定
  - validation failed errorのerror/warning件数を固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-3差分後）

- [x] **validate JSON formatを追加する**
  - `brewprint validate --yaml-root <path> --format json` を追加
  - 既定のtext出力を維持
  - valid projectのJSON出力をCLI testで固定
  - invalid projectのJSON出力をCLI testで固定
  - unsupported formatをerrorにする
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-4差分後）

- [x] **diagnostics出力順を安定化する**
  - `resolve.Build` 返却前にdiagnosticsを安定ソート
  - severity / file id / code / message の順で固定
  - unit testで順序を固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-5差分後）

- [x] **validate warning-only path testを追加する**
  - warning-only YAML root helperを追加
  - text validateでwarning diagnosticを固定
  - json validateで `error_count=0` / `warning_count=1` を固定
  - warning-only projectではnon-zero errorにしないことを固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-6差分後）

- [x] **required fields validation 第1段を実装する**
  - `missing_required_field` diagnostic codeを追加
  - node id欠落を検出
  - model kind / model field name/type欠落を検出
  - store kind欠落を検出
  - param / return / initialized storeのname/model欠落を検出
  - unit testで `missing_required_field` を固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-7差分後）

- [x] **missing required fieldのCLI JSON出力を固定する**
  - `missing_required_field` を含むinvalid YAML root helperを追加
  - `brewprint validate --format json` のdiagnostics内codeを確認
  - validation failedのnon-zero errorを確認
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-8差分後）

- [x] **diagnostic code specific化 第1段を実装する**
  - duplicate系diagnosticを具体code化
  - flow系diagnosticを具体code化
  - warning-only validate testを `unsupported_flow_entry` 期待へ更新
  - flow diagnostic codeをunit testで固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-9差分後）

- [x] **transition diagnostic code specific化を実装する**
  - `unresolved_transition_state` / `unresolved_transition_event` を追加
  - `duplicate_transition` / `missing_transition_guard` を追加
  - transition系diagnosticを具体code化
  - transition diagnostic codeをunit testで固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-10差分後）

- [x] **view / scenario diagnostic code specific化を実装する**
  - API View / ER View / Sequence Scenario系diagnosticを具体code化
  - `duplicate_view` / `invalid_view_definition` / `duplicate_view_module` を追加
  - `unresolved_sequence_step` / `non_continuous_sequence` を追加
  - view / scenario diagnostic codeをunit testで固定
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-11差分後）

- [x] **diagnostics specを追加する**
  - `docs/spec/diagnostics.md` を追加
  - Diagnostic object / severity / text format / JSON formatを記述
  - diagnostic orderingを記述
  - 現在のdiagnostic code一覧をカテゴリ別に記述
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-12差分後）

- [x] **diagnostic fallback codeを定数化する**
  - `semantic_validation` fallback codeを定数化
  - `symbolTable.addDiagnostic` の直書き文字列を削除
  - 挙動変更なし
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-13差分後）

- [x] **M7総括docを追加する**
  - `docs/impl/go-m7-summary.md` を追加
  - M7で追加したCLI / diagnostics / validation / spec / 境界 / 次候補を整理
  - 次sessionの復帰用メモを記載

### Milestone 8: render CLI / 一括render pipeline を実装する

- [x] **render CLI 第1段を実装する**
  - `brewprint render --yaml-root <path> --out <path>` を追加
  - validate相当のsemantic diagnosticsを通過したprojectだけrenderする
  - `internal/render/project` に一括render orchestratorを追加
  - 既存rendererを束ねて、DAG / State / Sequence / ER / API / Wireframe / Previewを出力する
  - `render_index.yaml` に基づき `renders/{group}/...` / `_cross/` / `_preview/` へ配置する
  - master `index.md` / group `index.md` を生成する
  - CLI testではUC-001を `t.TempDir()` に出力し、主要出力パスと生成件数を固定する
  - renderer本文のgolden一致は各 `internal/render/*` package testに委譲し、CLI testは一括生成・配置確認に限定する
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-29、M8-1差分後）

- [x] **一括render manifest testを追加する**
  - `internal/render/project/renderer_test.go` を追加
  - UC-001の一括render出力23件をmanifestとして固定
  - 全render fileのcontentが空でないことを確認する
  - `Write` がネストした出力ディレクトリを作成して書き込めることを固定
  - CLI testではなく一括render本体側で出力パス契約を担保する
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-29、M8-2差分後）

- [x] **UC-001 render fixtureを一括render出力へ追随する 第1段**
  - M8一括render manifestで不足していたDAG fixture 3本を追加
    - `renders/commerce/dag-add_to_cart.md`
    - `renders/catalog/dag-get_items.md`
    - `renders/commerce/dag-process_payment.md`
  - master `renders/index.md` のDAG件数を現行一括render出力に合わせて更新
  - group indexの完全追随は後続差分で扱う
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-29、M8-3差分後）

- [x] **DAG renderer golden testを追加fixtureへ拡張する**
  - `internal/render/dag/renderer_test.go` のgolden対象に追加DAG 3本を追加
    - `cart/task/add_to_cart.yaml`
    - `catalog/task/get_items.yaml`
    - `payment/webhooks/task/process_payment.yaml`
  - `dag-process_payment.md` を実レンダー出力に合わせて更新
  - `go test ./internal/render/dag` / `go test ./...` 通過済み（2026-04-29、M8-4差分後）

- [x] **UC-001 group index fixtureを一括render出力へ追随する**
  - `renders/commerce/index.md` に追加DAG 2本を反映
    - `dag-add_to_cart.md`
    - `dag-process_payment.md`
  - `renders/catalog/index.md` に `dag-get_items.md` を反映
  - `renders/auth/index.md` / `renders/catalog/index.md` / `renders/commerce/index.md` を project renderer 形式へ統一
  - `internal/render/project/renderer_test.go` に group index fixture 一致テストを追加
  - `go test ./...` 通過済み（2026-04-29、M8-5差分後）

- [x] **render CLI `--clean` を追加する**
  - `brewprint render --yaml-root <path> --out <path> --clean` を追加
  - render / validation / placement 成功後、書き込み直前に `--out` を削除して作り直す
  - `--out` が空・`.`・filesystem root・`..` を含む場合は拒否する
  - `--out` が `--yaml-root` を含む場合は YAML 削除事故防止のため拒否する
  - stale file 削除のCLI testを追加
  - unsafe clean path のunit testを追加
  - `go test ./...` 通過済み（2026-04-29、M8-6差分後）

- [x] **UC-001 render fixture再生成コマンドをREADMEへ明記する**
  - `docs/uc/001-ec-checkout-flow/README.md` に `brewprint render --yaml-root ... --out ... --clean` を追記
  - READMEの `renders/` ツリーに M8で追加したDAG fixtureを反映
  - docsのみの変更のため追加テストなし

- [x] **project renderer index生成仕様を固定する**
  - master index titleを project directory名から人間向けに整形する
  - group index titleに `render_index.yaml` の `label` を使う
  - master `index.md` も project renderer fixture一致テスト対象に戻す
  - preview titleも project名整形に追随
  - `go test ./...` 通過済み（2026-04-29、M8-8差分後）

- [ ] **render output path collision validationを実装する**
  - ADR-053 / `docs/spec/project-layout.md` に準拠
  - v1では nested module path を出力ファイル名に含めず、local ID filenameを維持する
  - 同一group内で複数render outputが同じrelative pathに解決された場合はerrorにする
  - silent overwriteは禁止する
  - collision対象source objectをerror messageに含める
  - project renderer / placement周辺のunit testで固定する

### Milestone 10: MCP project exploration / view inspect を拡張する

- [ ] **`list_objects` を実装する**
  - ADR-054 / `docs/spec/mcp.md` の設計対話coverageに準拠
  - project内の主要semantic objectを一覧・絞り込みできるようにする
  - object kind / module / file / label / source を返す
  - node / view / transition / field を対象に含める
  - asset / private sub node はM11で直接selector対応するまで対象外または任意扱いにする
  - MCP tool inputSchema / server wrapper testを追加する
  - QueryService unit testでUC-001の一覧を固定する
  - 実装後に `docs/spec/mcp.md` の tool仕様を追記する

- [ ] **`inspect(file)` を実装する**
  - FileID selectorで YAML file単位の定義内容を返す
  - node fileでは main node / sub nodes / flow summary / diagnostics を返す
  - state fileでは states / events / transitions / wireframe presence を返す
  - view fileでは view kind / id / target files or modules を返す
  - `render_index.yaml` は group summary / uncovered module warning 等を返す候補として扱う
  - `get_references(file)` の既存 `state_file` partial対応と整合させる
  - MCP wrapper test / QueryService testで代表fileを固定する
  - 実装後に `docs/spec/mcp.md` の selector support matrixを更新する

- [ ] **`inspect(view: api_table)` を実装する**
  - API Table view objectを直接inspectできるようにする
  - `http_root_path` / modules / include_submodules / collected endpoints / computed routes を返す
  - `list_endpoints` との役割分担を明確にする
  - excluded endpoint候補や収集対象0件sectionの扱いは実装時に判断し、必要ならspecへ追記する
  - UC-001の API Table view でQueryService / MCP wrapper testを追加する

- [ ] **`inspect(view: er_diagram)` を実装する**
  - ER Diagram view objectを直接inspectできるようにする
  - 対象modules / included stores / included models / FK relations / excluded refs summary を返す
  - default module単位ERと view YAMLによる横断ERの扱いを整理する
  - UC-001の ER view でQueryService / MCP wrapper testを追加する
  - 実装後に `docs/spec/mcp.md` の view inspect仕様を追記する

### Milestone 11: MCP diagram element query を拡張する

- [ ] **implicit asset selectorを実装する**
  - DAG上のasset nodeを直接queryできるようにする
  - selector形式は producer + name を基本候補とし、stable synthetic IDが必要か実装時に決める
  - asset signatureで name / producer / model / scope_file を返す
  - asset referencesで producer / consumer task への関係を返す
  - `produces_asset` との整合を保つ
  - UC-001のDAG assetでQueryService / MCP wrapper testを追加する
  - 実装後に `docs/spec/mcp.md` の AssetRef / selector support matrixを更新する

- [ ] **private sub node selectorを実装する**
  - file-local task / branch / fork / join を直接queryできるようにする
  - selectorは `<file-id>#<local-id>` または `file` + `local_id` を使う
  - get_signature / get_references / inspect の対応範囲を決める
  - main task inspect内の `members.sub_tasks` と同じObjectRef表現に揃える
  - UC-001の checkout sub task / branch / fork / join 相当でtestを追加する

- [ ] **flow wiring referencesを実装する**
  - DAG上のflow step / param wiringをMCP referenceとして辿れるようにする
  - reference kind候補: `flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over`
  - 既存方針「MCP v1ではflow wiringをget_referencesに返さない」を変更するため、必要なら小ADRまたはspec更新で扱う
  - `inspect(task).members.flow.entries` のdraft schemaと整合させる
  - DAG rendererのview modelとQueryServiceの責務境界を崩さない

### Milestone 12: MCP impact traversal / source assist を拡張する

- [ ] **`get_source` を実装する**
  - semantic objectに対応するYAML snippetを返す
  - Raw YAML AST全体公開ではなく、ResolvedProject objectのsource補助情報として扱う
  - selector / source range / fallback挙動を `docs/spec/mcp.md` に定義する
  - source line/columnが未取得の場合の返却形式を決める
  - MCP wrapper test / QueryService testを追加する

- [ ] **`get_reference_tree` または depth指定つきreference traversalを設計する**
  - direct referencesだけでは不足する変更影響範囲を辿れるようにする
  - 別tool `get_reference_tree` にするか、`get_references` に `depth` inputを追加するかを比較する
  - cycle detection / max depth / kind filter / direction の仕様を決める
  - ADR-049のdirect reference方針をどう拡張するか整理し、必要なら新ADRを起票する
  - 設計確定後に実装タスクを分割する
