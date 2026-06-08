# Go M7-6 validate warning path test 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: `brewprint validate` のwarning-only挙動をCLIテストで固定

---

## 1. 目的

M7-2〜M7-5で `brewprint validate` のtext/json出力とdiagnostics順序を整えた。

M7-6では、errorがなくwarningだけが出るprojectについて、validate commandが成功扱いを維持しつつwarningを出力することを固定する。

---

## 2. 実装範囲

- `cmd/brewprint/main_test.go` にwarning-only fixture helperを追加
- `flow: - {}` による既存warning diagnosticを利用
- text出力でwarning diagnosticが表示されることを確認
- json出力で `error_count=0` / `warning_count=1` になることを確認
- warning-only projectでは `run(validate)` がnil errorを返すことを確認

---

## 3. warning fixture

```yaml
nodes: []
flow:
  - {}
```

期待diagnostic:

```text
warning unsupported_flow_entry auth/task/login.yaml: unsupported empty flow entry
```

---

## 4. 境界

`validate` commandは引き続き以下だけを通す。

```text
cmd/brewprint -> source, resolve
```

query / mcp / renderer は呼ばない。

---

## 5. 受け入れ条件

- [x] warning-only projectのtext validate testを追加
- [x] warning-only projectのjson validate testを追加
- [x] warning-only projectではnon-zero errorにしない
- [x] json出力の `error_count=0` / `warning_count=1` を固定
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
