# Go M7-3 validate failure path test 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: `brewprint validate` の失敗系CLIテスト追加

---

## 1. 目的

M7-2で追加した `brewprint validate --yaml-root <path>` について、成功系だけでなく失敗系もCLIから固定する。

M7-3では、テスト内で一時YAML rootを作り、未解決model参照を含む最小node YAMLを置いて、validate commandがcode付きdiagnosticとnon-zero errorを返すことを確認する。

---

## 2. 実装範囲

- `cmd/brewprint/main_test.go` に `TestRunValidateInvalidProject` を追加
- `t.TempDir()` に最小YAML rootを作る
- `auth/task/login.yaml` に未解決 `params.model: missing_model` を置く
- stdoutが以下になることを確認する

```text
error unresolved_model auth/task/login.yaml: unresolved task params model: missing_model
```

- errorが以下を含むことを確認する

```text
validation failed: 1 error(s), 0 warning(s)
```

---

## 3. 境界

`validate` commandは引き続き以下だけを通す。

```text
cmd/brewprint -> source, resolve
```

query / mcp / renderer は呼ばない。

---

## 4. 受け入れ条件

- [x] `validate` 成功系テストがある
- [x] `validate` 失敗系テストがある
- [x] 失敗系stdoutにdiagnostic codeが含まれる
- [x] 失敗系errorにerror/warning件数が含まれる
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
