# Go M7-13 diagnostic fallback code 定数化タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: generic fallback diagnostic codeの定数化

## 1. 目的

M7-1〜M7-12でdiagnostic codeを具体化し、specにも一覧化した。

M7-13では、fallback用の `semantic_validation` 文字列を定数化し、diagnostic code管理を `internal/resolve/validation.go` に寄せる。

## 2. 実装範囲

- `diagnosticSemanticValidation = "semantic_validation"` を追加
- `symbolTable.addDiagnostic` の直書き文字列を定数参照へ変更

## 3. 境界

挙動変更なし。

```text
resolve -> rawyaml, semantic
```

## 4. 受け入れ条件

- [x] fallback diagnostic codeが定数化される
- [x] `addDiagnostic` が定数を使う
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る

## 5. 検証メモ

2026-04-28にユーザー環境で以下を確認済み。

```sh
cd C:\Users\imved\projects\brewprint
go fmt ./...
go test ./...
```

全package通過。
