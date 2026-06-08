# Go M7-10 transition diagnostic code specific化 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: transition系diagnosticの具体code化

## 1. 目的

M7-9でduplicate系 / flow系diagnosticを具体code化した。

M7-10では、transition系diagnosticに範囲を絞って、汎用 `semantic_validation` ではなく具体codeを返すようにする。

## 2. 実装範囲

追加したdiagnostic code:

- `unresolved_transition_state`
- `unresolved_transition_event`
- `duplicate_transition`
- `missing_transition_guard`

対象:

- transition `from` state未解決
- transition `to` state未解決
- transition `on` event未解決
- 同一 `from/on/guard` transition重複
- 同一 `from/on` の分岐transitionでguard欠落

## 3. 変更ファイル

- `internal/resolve/validation.go`
  - transition系diagnostic code定数を追加
- `internal/resolve/transitions.go`
  - transition系diagnosticを `addDiagnosticCode` へ置換
- `internal/resolve/validation_test.go`
  - transition diagnostic codeのunit testを追加

## 4. 境界

変更は `internal/resolve` に限定する。

```text
resolve -> rawyaml, semantic
```

query / mcp / renderer には触れない。

## 5. 受け入れ条件

- [x] transition state未解決が `unresolved_transition_state` を返す
- [x] transition event未解決が `unresolved_transition_event` を返す
- [x] transition重複が `duplicate_transition` を返す
- [x] branched transition guard欠落が `missing_transition_guard` を返す
- [x] unit testで各codeを固定する
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る

## 6. 検証メモ

2026-04-28にユーザー環境で以下を確認済み。

```sh
cd C:\Users\imved\projects\brewprint
go fmt ./...
go test ./...
```

全package通過。
