# Go M7-11 view / scenario diagnostic code specific化 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: API View / ER View / Sequence Scenario系diagnosticの具体code化

## 1. 目的

M7-10までにduplicate / flow / transition系diagnosticを具体code化した。
M7-11ではview定義とsequence scenario定義のdiagnosticを具体code化する。

## 2. 追加したdiagnostic code

- `duplicate_view`
- `invalid_view_definition`
- `duplicate_view_module`
- `unresolved_sequence_step`
- `non_continuous_sequence`

## 3. 変更範囲

- `internal/resolve/validation.go`
  - view / scenario系diagnostic code定数を追加
- `internal/resolve/api_views.go`
  - API View系diagnosticを具体code化
- `internal/resolve/er_views.go`
  - ER View系diagnosticを具体code化
- `internal/resolve/scenarios.go`
  - Sequence Scenario系diagnosticを具体code化
- `internal/resolve/validation_test.go`
  - view / scenario diagnostic codeのunit testを追加

## 4. 境界

変更は `internal/resolve` に限定する。

```text
resolve -> rawyaml, semantic
```

query / mcp / renderer には触れない。

## 5. 受け入れ条件

- [x] API View / ER Viewの不正定義が `invalid_view_definition` を返す
- [x] API View / ER Viewの重複viewが `duplicate_view` を返す
- [x] API View / ER Viewの重複moduleが `duplicate_view_module` を返す
- [x] sequence step未解決が `unresolved_sequence_step` を返す
- [x] sequence step非連続が `non_continuous_sequence` を返す
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
