# Go M9 QueryService state/event/scenario handoff

- **status**: complete
- **last_updated**: 2026-04-29
- **repo**: `C:\Users\imved\projects\brewprint`
- **verified**: `gofmt ./...` / `go test ./...` 通過済み（2026-04-29、M9-4差分後）

---

## 1. 次セッションで最初に読むもの

次セッションでは以下の順で読む。

1. `docs/prompt_chappy.md`
2. `docs/doc-policy.md`
3. `docs/impl/go-m8-summary.md`
4. `docs/impl/go-m9-summary.md`（このファイル）
5. `docs/TASKS.md` の Milestone 3 付近
6. `docs/spec/mcp.md` の以下
   - `Reference kind`
   - `state inspect`
   - `event inspect`
   - `scenario inspect`

必要に応じて追加で読む実装ファイル:

- `internal/semantic/project.go`
- `internal/semantic/state.go`
- `internal/semantic/reference.go`
- `internal/resolve/transitions.go`
- `internal/resolve/scenarios.go`
- `internal/resolve/references.go`
- `internal/query/types.go`
- `internal/query/signature.go`
- `internal/query/references.go`
- `internal/query/inspect.go`
- `internal/query/service_test.go`

---

## 2. M9で完了したこと

M9では QueryService / ResolvedProject 周辺の state / event / scenario 系 reverse lookup と MCP response を拡張した。

### M9-1: state / event / scenario 系 reverse lookup index 第2段

追加した `semantic.Project` index:

- `TransitionsByStateEventGuard map[TransitionKey]TransitionRef`
- `TransitionsByStateEvent map[TransitionEventKey][]TransitionRef`
- `ActionsByTask map[QualifiedID][]TransitionRef`

追加した semantic 型:

- `TransitionKey`
- `TransitionEventKey`
- `TransitionRef`

実装内容:

- `resolve.buildTransitions` で transition index を構築するようにした。
- `ActionsByTask` は解決済み action task のみを index する。
- `resolve.buildScenarios` の scenario step 解決を、既存 slice 走査から transition index lookup に変更した。
- `query.Inspect(task)` に `members.action_transitions` を追加した。

テスト:

- UC-001で guarded transition index を固定
- 同一 `(state,event)` 候補2件を固定
- `actionsByTask` を固定
- `inspect(task)` の `action_transitions` を固定

検証:

```powershell
go test ./...
```

M9-1差分後に通過済み。

---

### M9-2: transition / event direct references を GetReferences へ追加

追加した reference kind:

- `transition_from`
- `transition_event`
- `transition_to`
- `transition_action`
- `event_payload`
- `event_actor`
- `event_watches`

追加した semantic helper:

- `semantic.TransitionID(transition)`
- `semantic.TransitionObjectKey(transition)`

追加した endpoint fields:

- `state_file`
- `from`
- `on`
- `to`
- `guard`
- `action`

実装内容:

- `resolve.buildReferences` に event / transition 系 direct reference を追加した。
- transition を reference source として扱うため、transition synthetic object key を導入した。
- `query.ReferenceEndpoint` に transition 詳細フィールドを通した。
- `query.transitionID` は `semantic.TransitionID` を使うようにした。

UC-001で固定した reference:

- task ← `transition_action`
  - `order.task.checkout`
  - `payment.webhooks.task.process_payment`
- event → `event_payload`
- event → `event_actor`
- event ← `transition_event`
- model ← `event_payload`
- store ← `event_watches`
- state ← `transition_from` / `transition_to`

検証:

```powershell
gofmt ./...
go test ./...
```

M9-2差分後に通過済み。

---

### M9-3: inspect(state/event) と GetSignature(state/event)

`GetSignature` 対応:

- `state`
  - `initial`
  - `final`
  - `wireframe.present`
- `event`
  - `source`
  - `actor`
  - `payload.model`
  - `watches`

`Inspect` 対応:

- `state`
  - `members.incoming_transitions`
  - `members.outgoing_transitions`
- `event`
  - `members.triggering_transitions`
  - `members.sequence_hints`

`sequence_hints` は ADR-036 / sequence render rule 由来の advisory 情報。
現状の mapping:

- `external` → participant `Actor`, actor from event, message label source `METHOD path`
- `ui` → participant `User`, message label source `event id`
- `internal` → participant `Task`, message label source `event id`
- `er` → participant `DB`, message label source `watched store`

UC-001で固定した test:

- `GetSignature(order.state.checkout_screen)`
- `GetSignature(order.event.payment_webhook_received)`
- `Inspect(order.state.processing)`
- `Inspect(order.event.payment_webhook_received)`

検証:

```powershell
gofmt ./...
go test ./...
```

M9-3差分後に通過済み。

注意:

- `go test` 実行時に `open ...\.git: The system cannot find the path specified.` が表示されるが、全packageは `ok`。
- 現時点では Go test failure ではなく環境側メッセージとして扱っている。

---

### M9-4: inspect(scenario) と scenario references

`GetSignature` 対応:

- `sequence_diagram` view object
  - `id`
  - `title`
  - `state_file`

`Inspect` 対応:

- `sequence_diagram` view object
  - `members.steps`
  - step `index`
  - `from_state`
  - `via`
  - `guard`
  - `guard_exact_match`
  - resolved `transition`
  - resolved `action`

`GetReferences` 対応:

- `scenario_state_file`
- `scenario_step_transition`

実装内容:

- `query.Selector` に `object` / `kind` / `file` / `local_id` を追加した。
- `selector.object=view` / `selector.kind=sequence_diagram` で scenario view object を解決できるようにした。
- `semantic.ReferenceKindScenarioStateFile` / `semantic.ReferenceKindScenarioStepTransition` を追加した。
- `resolve.buildReferences` で scenario reference を index に追加した。
- MCP selector schema を `docs/spec/mcp.md` の Object selector に寄せて拡張した。

UC-001で固定した test:

- `GetSignature(checkout_flow)`
- `GetReferences(checkout_flow)`
- `Inspect(checkout_flow)`
- `Inspect(payment_webhook_flow)`
- MCP `inspect` で `checkout_flow` を view selector 指定

検証:

```powershell
gofmt ./...
go test ./...
```

M9-4差分後に通過済み。

注意:

- `go test` 実行時に `open ...\.git: The system cannot find the path specified.` が表示されるが、全packageは `ok`。
- 現時点では Go test failure ではなく環境側メッセージとして扱っている。

---

## 3. 現在の実装境界

維持する境界:

```text
source  -> rawyaml
resolve -> rawyaml, semantic
query   -> semantic
mcp     -> query
renderer -> semantic
```

禁止:

```text
internal/mcp -> rawyaml
internal/mcp内でYAML load / resolve
internal/mcp内でrenderer呼び出し
renderer内でRaw YAML structsを直接読むこと
renderer内でname resolution / semantic validationを再実装すること
```

M9で追加した index / reference / inspect はすべて `resolve -> semantic -> query` の流れに収めている。
MCP wrapper は引き続き QueryService を呼ぶだけにする。

---

## 4. 変更ファイル一覧

### semantic

- `internal/semantic/project.go`
  - transition/action reverse lookup index を追加
- `internal/semantic/state.go`
  - `TransitionKey` / `TransitionEventKey` / `TransitionRef` を追加
- `internal/semantic/reference.go`
  - transition/event reference kind を追加
  - scenario reference kind を追加
  - transition endpoint fields を追加
  - `TransitionID` / `TransitionObjectKey` を追加
  - `ScenarioObjectKey` / `ScenarioStepObjectKey` / `StateFileObjectKey` を追加

### resolve

- `internal/resolve/transitions.go`
  - transition index 構築を追加
- `internal/resolve/scenarios.go`
  - scenario step 解決を transition index lookup に変更
- `internal/resolve/references.go`
  - transition/event direct references を追加
  - scenario state file / scenario step transition references を追加

### query

- `internal/query/types.go`
  - `Selector` に view selector fields を追加
  - `TransitionRef`
  - `ScenarioStepRef`
  - transition endpoint fields
- `internal/query/service.go`
  - scenario selector 解決を追加
- `internal/query/signature.go`
  - state/event signature 対応
  - scenario signature 対応
- `internal/query/references.go`
  - transition endpoint fields を response に反映
  - scenario object の reference lookup 対応
- `internal/query/inspect.go`
  - task action transition
  - state inspect
  - event inspect
  - scenario inspect
  - sequence hints
- `internal/query/service_test.go`
  - M9-1〜M9-4 の UC-001 test を追加

### mcp

- `internal/mcp/server.go`
  - selector inputSchema を `object` / `kind` / `file` / `local_id` 対応へ拡張
- `internal/mcp/server_test.go`
  - `inspect_scenario` test を追加

### docs

- `docs/TASKS.md`
  - M9-1〜M9-4完了
- `docs/impl/go-m9-summary.md`
  - この引継ぎファイル

---

## 5. Post-M9で完了したこと

### Post-M9-1: GetReferences(transition)

`selector.object=transition` / `selector.kind=transition` で transition synthetic ID を直接解決できるようにした。

対応した問い合わせ例:

```json
{
  "selector": {
    "object": "transition",
    "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
  },
  "direction": "both"
}
```

返せる references:

- `transition_from`
- `transition_event`
- `transition_to`
- `transition_action`
- incoming `scenario_step_transition`

実装内容:

- `query.transitionBySelector` を追加した。
- `query.isTransitionSelector` を追加した。
- `query.transitionObjectRef` を追加した。
- `GetReferences` の target resolver で transition object を扱えるようにした。
- QueryService testで guarded transition の out/in references を固定した。
- MCP wrapper testで `get_references_transition` を固定した。

検証:

```powershell
gofmt ./...
go test ./...
```

Post-M9-1差分後に通過済み。

注意:

- `go test` 実行時に `open ...\.git: The system cannot find the path specified.` が表示されるが、全packageは `ok`。
- 現時点では Go test failure ではなく環境側メッセージとして扱っている。

---

### Post-M9-2: inspect(transition)

`selector.object=transition` / `selector.kind=transition` で transition synthetic ID を直接 inspect できるようにした。

対応した問い合わせ例:

```json
{
  "selector": {
    "object": "transition",
    "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']"
  }
}
```

返せる内容:

- transition signature
  - `state_file`
  - `from`
  - `on`
  - `to`
  - `guard`
  - `action`
- `members.from_state`
- `members.event`
- `members.to_state`
- `members.action_task`
- direct `references`
  - `transition_from`
  - `transition_event`
  - `transition_to`
  - `transition_action`
  - incoming `scenario_step_transition`

実装内容:

- `GetSignature(transition)` を追加した。
- `signatureForTransition` を追加した。
- `Inspect` の transition selector 分岐を追加した。
- `inspectTransition` を追加した。
- `transitionMembers` を追加した。
- QueryService testで guarded transition の signature / inspect members / incoming scenario reference を固定した。
- MCP wrapper testで `inspect_transition` を固定した。

検証:

```powershell
gofmt ./...
go test ./...
```

Post-M9-2差分後に通過済み。

注意:

- `go test` 実行時に `open ...\.git: The system cannot find the path specified.` が表示されるが、全packageは `ok`。
- 現時点では Go test failure ではなく環境側メッセージとして扱っている。

---

### Post-M9-3: docs/spec/mcp.md transition spec追随

実装済みの transition object 対応に合わせて、MCP specを更新した。

追加した仕様:

- `get_signature` の `transition signature`
  - selector例
  - response例
  - `state_file` / `from` / `on` / `to` / `guard` / `action` のfield定義
- `inspect` の `transition inspect`
  - selector例
  - response例
  - `members.from_state`
  - `members.event`
  - `members.to_state`
  - `members.action_task`
  - transition direct references
  - incoming `scenario_step_transition`

変更ファイル:

- `docs/spec/mcp.md`

検証:

- docs only変更のため追加テストなし。

---

### Post-M9-4: GetReferences(field/file)

`get_references` の object selector 対応を field / file に広げた。

対応した selector:

```json
{
  "selector": {
    "object": "field",
    "id": "order.model.order",
    "local_id": "id"
  },
  "direction": "both"
}
```

```json
{
  "selector": {
    "object": "file",
    "kind": "state_file",
    "id": "order/state.yaml"
  },
  "direction": "in"
}
```

返せる references:

- field selector
  - `field_type`
  - `field_fk`
- file / state_file selector
  - incoming `scenario_state_file`

実装内容:

- `query.modelFieldBySelector` を追加した。
- `query.isFieldSelector` を追加した。
- `query.fileBySelector` を追加した。
- `query.isFileSelector` を追加した。
- `query.fieldObjectRef` を追加した。
- `query.fileObjectRef` を追加した。
- `GetReferences` の target resolver で field / file object を扱えるようにした。
- QueryService testで `order.model.order.id` の field references を固定した。
- QueryService testで `order/state.yaml` の incoming `scenario_state_file` references を固定した。
- MCP wrapper testで `get_references_field` / `get_references_file` を固定した。

検証:

```powershell
gofmt ./...
go test ./...
```

Post-M9-4差分後に通過済み。

注意:

- `order.model.order.id` に対する incoming `field_fk` は、現状 `payment.model.payment_event.order_id` のみ確認される。
- `order.model.order_item.order_id` の `fk: order.id` は、現状のreference indexでは `order.model.order.id` へ正規化されていない。
- bare FK正規化は次候補として扱う。
- `go test` 実行時に `open ...\.git: The system cannot find the path specified.` が表示されるが、全packageは `ok`。
- 現時点では Go test failure ではなく環境側メッセージとして扱っている。

---

## 6. 次にやること

### M9は完了

Milestone 3 の QueryService 拡張は M9-4 で完了。
Post-M9-1で transition object の direct references も問い合わせ可能になった。
Post-M9-2で transition object の signature / inspect も問い合わせ可能になった。
Post-M9-3で `docs/spec/mcp.md` も transition signature / inspect に追随した。
Post-M9-4で field / file object の direct references も問い合わせ可能になった。

次候補:

1. Post-M9-4差分をcommitする
2. bare FK正規化を追加する
   - 例: `fk: order.id` を `order.model.order.id` へ解決する
3. MCP / QueryService の object selector 対応をさらに広げる
   - asset object references
   - `inspect(field)`
   - `inspect(file)`
   - ER / API view object inspect
4. validation / diagnostics の追加強化
5. render CLI / docs の仕上げ

---

## 7. 次セッション開始時の推奨コマンド

```powershell
cd C:\Users\imved\projects\brewprint
git status
gofmt ./...
go test ./...
```

`go test` 時に以下が出る可能性がある。

```text
open ...\.git: The system cannot find the path specified.
```

ただし全packageが `ok` なら、現時点では作業継続してよい。

---

## 8. commitメモ

この環境から `git status` は確認していない。
次セッション冒頭で必ず確認する。

```powershell
git status
```

自然な commit 分割候補:

1. M8後半まとめ
   - render CLI `--clean`
   - project renderer index fixture stabilization
2. M9 query 拡張まとめ
   - reverse lookup index
   - transition/event references
   - state/event/scenario inspect
   - scenario references
3. Post-M9 query object selector拡張
   - `GetReferences(transition)`
   - `GetSignature(transition)`
   - `Inspect(transition)`
4. Post-M9 spec追随
   - transition signature spec
   - transition inspect spec
5. Post-M9 field/file references
   - `GetReferences(field)`
   - `GetReferences(file)`

Post-M9-4だけをcommitするなら候補:

```powershell
git add .
git commit -m "feat(query): support field and file references by selector"
```

---

## 9. 直近検証ログ

ユーザー実行:

```powershell
gofmt ./...
go test ./...
```

結果:

- `cmd/brewprint`: OK
- `internal/mcp`: OK
- `internal/query`: OK
- `internal/render/*`: OK
- `internal/resolve`: OK
- no test files packagesも問題なし

M9-3差分後に通過済み。
