# Go M7 validation / diagnostics 総括

- **status**: completed
- **last_updated**: 2026-04-28
- **repo**: `C:\Users\imved\projects\brewprint`
- **verified**: `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M7-13差分後）

---

## 1. 完了したこと

M7では validation / diagnostics を強化した。

- `semantic.Diagnostic` に `code` を追加
- `resolve.Build` にsemantic validationを追加
- diagnosticsを安定ソート
- `brewprint validate --yaml-root <path>` を追加
- `brewprint validate --yaml-root <path> --format json` を追加
- warning-only projectは成功扱いに固定
-主要diagnosticをmachine-readableなcode付きに整理
- `docs/spec/diagnostics.md` を追加

---

## 2. 重要ファイル

### CLI

- `cmd/brewprint/main.go`
- `cmd/brewprint/main_test.go`

### semantic / resolve

- `internal/semantic/diagnostic.go`
- `internal/resolve/validation.go`
- `internal/resolve/diagnostics.go`
- `internal/resolve/validation_test.go`

### diagnostics specific化対象

- `internal/resolve/symbols.go`
- `internal/resolve/builder.go`
- `internal/resolve/flow.go`
- `internal/resolve/transitions.go`
- `internal/resolve/api_views.go`
- `internal/resolve/er_views.go`
- `internal/resolve/scenarios.go`

### docs

- `docs/spec/diagnostics.md`
- `docs/TASKS.md`

---

## 3. CLI仕様

```sh
brewprint validate --yaml-root <path>
```

Diagnosticsなし:

```text
ok
```

Diagnosticsあり:

```text
error unresolved_model order/state.yaml: unresolved event payload model: payment_event
```

JSON:

```sh
brewprint validate --yaml-root <path> --format json
```

```json
{
  "diagnostics": [],
  "error_count": 0,
  "warning_count": 0
}
```

Error diagnosticがある場合もJSONはstdoutへ出力し、その後non-zero errorを返す。

---

## 4. Diagnostic ordering

`resolve.Build` 返却前に以下で安定ソートする。

1. severity rank: `error` -> `warning` -> unknown
2. file id
3. code
4. message

---

## 5. Diagnostic code

詳細は `docs/spec/diagnostics.md` を参照。

現在の主なcode:

```text
semantic_validation
invalid_model_id
unresolved_model
unresolved_field_type
unresolved_fk
unresolved_store
invalid_endpoint
invalid_store_kind
invalid_model_kind
duplicate_model_field
duplicate_primary_key
missing_required_field
duplicate_node
duplicate_main_node
duplicate_actor
duplicate_initialized_store
unsupported_flow_entry
unresolved_flow_task
unresolved_flow_node
invalid_flow_branch
unmatched_join_param
unresolved_transition_state
unresolved_transition_event
duplicate_transition
missing_transition_guard
duplicate_view
invalid_view_definition
duplicate_view_module
unresolved_sequence_step
non_continuous_sequence
```

---

## 6. M7で修正したUC-001 fixture

M7 validationにより、ADR-027違反のクロスモジュール短縮参照が見つかった。

修正済み:

- `docs/uc/001-ec-checkout-flow/yaml/payment/webhooks/task/process_payment.yaml`
  - `payment_event` -> `payment.model.payment_event`
- `docs/uc/001-ec-checkout-flow/yaml/order/state.yaml`
  - `payment_event` -> `payment.model.payment_event`
- `docs/uc/001-ec-checkout-flow/renders/_cross/api.md`
  - API Table goldenを更新

---

## 7. 境界

維持済み:

```text
source  -> rawyaml
resolve -> rawyaml, semantic
query   -> semantic
mcp     -> query
```

禁止:

```text
internal/mcp -> rawyaml
internal/mcp内でYAML load / resolve
internal/mcp内でrenderer呼び出し
```

Validationは `internal/resolve` に置く。Query / MCP / renderer側でvalidationを再実装しない。

---

## 8. 次session向けメモ

次会話では以下だけ読めば復帰しやすい。

1. `docs/prompt_chappy.md`
2. `docs/doc-policy.md`
3. `docs/impl/go-m7-summary.md`
4. `docs/TASKS.md` の末尾

次に進むならM8候補:

- `brewprint render --yaml-root <path> --out <path>` CLI
- render_index駆動の一括render
- validate -> render pipeline

直近検証:

```powershell
cd C:\Users\imved\projects\brewprint
go fmt ./...
go test ./...
```

全package通過済み。
