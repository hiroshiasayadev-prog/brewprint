# Go M7-4 validate JSON format 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: `brewprint validate` のJSON出力追加

---

## 1. 目的

M7-2 / M7-3で追加した `brewprint validate --yaml-root <path>` を、CIや外部ツールから扱いやすくする。

M7-4では、既定のtext出力を維持したまま、machine-readableなJSON出力を追加する。

```sh
brewprint validate --yaml-root <path> --format json
```

---

## 2. 実装範囲

- `validate` commandに `--format text|json` を追加
- 既定値は `text`
- `--format text` は既存挙動を維持
  - diagnosticsなし: `ok`
  - diagnosticsあり: `severity code file: message`
- `--format json` はJSON objectをstdoutへ出力
- error diagnosticsがある場合も、JSONはstdoutへ出力してからnon-zero errorを返す
- unsupported formatはload前にerrorにする

---

## 3. JSON出力schema

```json
{
  "diagnostics": [],
  "error_count": 0,
  "warning_count": 0
}
```

`diagnostics` の各要素は `semantic.Diagnostic` のJSON形式。

```json
{
  "severity": "error",
  "code": "unresolved_model",
  "file": "auth/task/login.yaml",
  "message": "unresolved task params model: missing_model"
}
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

- [x] `brewprint validate --yaml-root <path> --format json` を追加
- [x] 既定のtext出力を維持
- [x] valid projectのJSON出力をCLI testで固定
- [x] invalid projectのJSON出力をCLI testで固定
- [x] unsupported formatをerrorにする
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
