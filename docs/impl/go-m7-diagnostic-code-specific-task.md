# Go M7-9 diagnostic code specific化 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: 汎用 `semantic_validation` diagnostic の一部を具体codeへ置換

---

## 1. 目的

M7-1〜M7-8でdiagnostic codeとvalidate CLIを整えた。

M7-9では、まだ汎用 `semantic_validation` として出ていたdiagnosticのうち、意味が明確な範囲だけを具体codeへ置き換える。

---

## 2. 実装範囲

対象を以下に絞る。

### symbol / duplicate系

- `duplicate_node`
- `duplicate_main_node`
- `duplicate_actor`
- `duplicate_initialized_store`

### flow系

- `unsupported_flow_entry`
- `unresolved_flow_task`
- `unresolved_flow_node`
- `invalid_flow_branch`
- `unmatched_join_param`

---

## 3. 実装メモ

- `internal/resolve/symbols.go`
  - node / main node / actor重複を具体code化
- `internal/resolve/builder.go`
  - initialized store重複を具体code化
- `internal/resolve/flow.go`
  - flow entry / unresolved task / unresolved node / invalid branch / unmatched join paramを具体code化
- `cmd/brewprint/main_test.go`
  - warning-only fixtureの期待codeを `semantic_validation` から `unsupported_flow_entry` へ更新
- `internal/resolve/validation_test.go`
  - flow diagnostic codeのunit testを追加
- `docs/impl/go-m7-validate-warning-test-task.md`
  - 期待diagnostic表記を現行codeに更新

---

## 4. 境界

変更は `internal/resolve` と CLI test に限定する。

```text
resolve -> rawyaml, semantic
cmd/brewprint -> source, resolve
```

query / mcp / renderer には触れない。

---

## 5. 受け入れ条件

- [x] duplicate系diagnosticが具体codeを持つ
- [x] flow系diagnosticが具体codeを持つ
- [x] warning-only validate testが `unsupported_flow_entry` を期待する
- [x] flow diagnostic codeをunit testで固定する
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
