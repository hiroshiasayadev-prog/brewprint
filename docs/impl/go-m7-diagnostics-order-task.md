# Go M7-5 diagnostics order stabilization 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: diagnostics出力順の安定化

---

## 1. 目的

M7-1〜M7-4でdiagnosticsのcode化とCLI出力を追加した。

M7-5では、map iteration由来でdiagnosticsの順序が揺れないように、`resolve.Build` の返却直前でdiagnosticsを安定ソートする。

---

## 2. 実装範囲

- `internal/resolve/diagnostics.go` を追加
- `sortedDiagnostics` を追加
- `resolve.Build` の返却直前で `sortedDiagnostics(symbols.diags)` を適用
- sorting順をunit testで固定

---

## 3. ソート順

以下の順で昇順ソートする。

1. severity rank
   - `error`
   - `warning`
   - unknown severity
2. file id
3. diagnostic code
4. message

---

## 4. 受け入れ条件

- [x] diagnosticsがBuild返却前に安定ソートされる
- [x] errorがwarningより先に出る
- [x] file id / code / message の順で安定化される
- [x] unit testで順序を固定する
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
