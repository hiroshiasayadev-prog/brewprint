# Milestone 3: QueryService を通す

- **status**: closed
- **scope**: QueryService
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

---

## Tasks

- [x] **QueryService vertical slice 第1段を実装する**
  - `GetSignature`
  - `GetReferences` はdirect referencesのみ
  - `Inspect`
  - unit testから直接叩けるようにする
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28）

- [x] **direct reference index を整備する**
  - V01-ADR-048 / V01-ADR-049に準拠して実装する
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
