# Go M9 QueryService state/event/scenario handoff

- **status**: in progress
- **last_updated**: 2026-04-29
- **repo**: `C:\Users\imved\projects\brewprint`
- **verified**: `gofmt ./...` / `go test ./...` 通過済み（2026-04-29、M9-3差分後）

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
  - transition endpoint fields を追加
  - `TransitionID` / `TransitionObjectKey` を追加

### resolve

- `internal/resolve/transitions.go`
  - transition index 構築を追加
- `internal/resolve/scenarios.go`
  - scenario step 解決を transition index lookup に変更
- `internal/resolve/references.go`
  - transition/event direct references を追加

### query

- `internal/query/types.go`
  - `TransitionRef`
  - transition endpoint fields
- `internal/query/signature.go`
  - state/event signature 対応
- `internal/query/references.go`
  - transition endpoint fields を response に反映
- `internal/query/inspect.go`
  - task action transition
  - state inspect
  - event inspect
  - sequence hints
- `internal/query/service_test.go`
  - M9-1〜M9-3 の UC-001 test を追加

### docs

- `docs/TASKS.md`
  - M9-1〜M9-3完了
  - 次候補 `inspect(scenario)` 追加
- `docs/impl/go-m9-summary.md`
  - この引継ぎファイル

---

## 5. 次にやること

### 最優先候補: inspect(scenario) を実装する

`docs/spec/mcp.md` には scenario inspect が定義済み。
次はこれを QueryService に通す。

候補スコープ:

1. selectorで sequence scenario view object を解決する
   - 既存 `nodeByID` は node 専用なので拡張または別 resolver が必要
   - selector.object / selector.kind は raw structに存在するが、Go `Selector` は現状 `ID` だけなので注意
2. scenario signature を返す
   - `id`
   - `title`
   - `state_file`
3. `members.steps` を返す
   - `index`
   - `from_state`
   - `via`
   - `guard`
   - `transition` as `TransitionRef`
   - `action` as task QualifiedID or empty/null相当
4. references を検討する
   - `scenario_state_file`
   - `scenario_step_transition`

実装候補ファイル:

- `internal/query/types.go`
  - `Selector` に `Object` / `Kind` を足すか検討
  - `ScenarioStepRef` などを足すか検討
- `internal/query/inspect.go`
  - scenario inspect entrypoint
- `internal/query/signature.go`
  - scenario signature helper
- `internal/query/references.go`
  - view/scenario reference を `GetReferences` に入れるか、scenario inspect 内だけにするか判断
- `internal/query/service_test.go`
  - `checkout_flow`
  - `payment_webhook_flow`

注意:

- 現状 `GetReferences` は `nodeByID` 前提。
- scenario は `semantic.Project.ScenariosByID` に入っているが node ではない。
- まず `Inspect` 専用で view object selector を扱い、`GetReferences(view)` は後続に分けるのが安全。
- `docs/spec/mcp.md` では selector.object / selector.kind が定義されているが、Go `Selector` はまだ `ID` のみ。ここを拡張すると MCP schema / wrapper test に影響する可能性がある。

---

## 6. 次セッション開始時の推奨コマンド

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

## 7. commitメモ

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
   - state/event inspect

M9だけをcommitするなら候補:

```powershell
git add .
git commit -m "feat(query): add transition references and state event inspect"
```

---

## 8. 直近検証ログ

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
