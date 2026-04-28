# Go M7 validation / diagnostics 強化 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: Go実装 Milestone 7 第1段のsemantic validation / diagnostics強化

---

## 1. 目的

M7では、ResolvedProject build時のvalidation / diagnosticsを強化する。

M7-1ではまず、既存のsemantic build pipelineを崩さず、診断情報の外形と参照検証を強化する。

---

## 2. 実装範囲

- `semantic.Diagnostic` に `code` を追加
- resolver内のdiagnostic追加helperをcode対応にする
- model / store / task / control / event の主要参照をsemantic validationする
- invalid fixtureをrawyaml.Projectで直接組むunit testを追加

---

## 3. 境界

```text
source  -> rawyaml
resolve -> rawyaml, semantic
query   -> semantic
mcp     -> query
```

M7-1でも境界は維持する。

- validationは `internal/resolve` に置く
- `internal/query` / `internal/mcp` でYAMLやRaw構造を読まない
- renderer側でvalidationを再実装しない

---

## 4. M7-1で追加したdiagnostic code

- `invalid_model_id`
- `unresolved_model`
- `unresolved_field_type`
- `unresolved_fk`
- `unresolved_store`
- `invalid_endpoint`
- `invalid_store_kind`
- `invalid_model_kind`
- `duplicate_model_field`
- `duplicate_primary_key`

既存のcode未分類diagnosticは `semantic_validation` とする。

---

## 5. M7-1 validation対象

### model

- primitive予約語をmodel IDとして使っていないか
- `kind` が `struct` / `list` / `dict` のいずれかか
- struct field名が重複していないか
- PKが複数ないか
- field `type` がprimitive予約語または定義済みmodelか
- field `fk` が存在するmodel fieldを指すか

### store

- `kind` が `db` / `session` / `collection` / `context` のいずれかか
- `of` が定義済みmodelを指すか

### task / branch / fork / join

- params modelがprimitive予約語または定義済みmodelか
- returns modelが定義済みmodelか
- initialized store modelが定義済みmodelか
- reads / writes storeが解決できるか
- endpoint methodが有効なHTTP methodか
- endpoint pathがsingle segmentか

### event

- payload modelが定義済みmodelか
- watches storeが解決できるか

---

## 6. 受け入れ条件

- [x] `semantic.Diagnostic` が `code` を持つ
- [x] 新規validationがcode付きdiagnosticを返す
- [x] unresolved reads/writesはwarningではなくvalidation層のcode付きerrorに一本化
- [x] invalid rawyaml.Project fixtureのunit testを追加
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る

## 7. 検証メモ

2026-04-28にユーザー環境で以下を確認済み。

```sh
cd C:\Users\imved\projects\brewprint
go fmt ./...
go test ./...
```

全package通過。

初回ローカル実行では、UC-001 fixtureの `payment/webhooks/task/process_payment.yaml` が `payment.model.payment_event` を `payment_event` と短縮参照していたため、`unresolved_model` で失敗した。
ADR-027ではクロスモジュール参照はフルパス必須のため、fixture側を `payment.model.payment_event` に修正した。

2回目ローカル実行では、UC-001 fixtureの `order/state.yaml` が `payment.model.payment_event` を `payment_event` と短縮参照していたため、`unresolved_model` で失敗した。
同じくADR-027に従い、event payload modelを `payment.model.payment_event` に修正した。
