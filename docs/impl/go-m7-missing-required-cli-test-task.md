# Go M7-8 missing required field CLI test 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: `missing_required_field` diagnostic のCLI JSON出力テスト追加

---

## 1. 目的

M7-7で追加した `missing_required_field` diagnostic が、resolver unit testだけでなく `brewprint validate --format json` からも確認できることを固定する。

---

## 2. 実装範囲

- `cmd/brewprint/main_test.go` に `TestRunValidateMissingRequiredFieldsJSON` を追加
- `t.TempDir()` に最小YAML rootを作る
- `model.kind` と `model field.type` が欠落したfixtureを置く
- `brewprint validate --yaml-root <path> --format json` がnon-zero errorを返すことを確認
- JSON出力に `missing_required_field` diagnostic codeが含まれることを確認

---

## 3. fixture

```yaml
nodes:
  - id: broken
    type: model
    fields:
      - name: id
```

期待:

- `error_count > 0`
- `diagnostics` に `missing_required_field` が含まれる

---

## 4. 境界

`validate` commandは引き続き以下だけを通す。

```text
cmd/brewprint -> source, resolve
```

query / mcp / renderer は呼ばない。

---

## 5. 受け入れ条件

- [x] `missing_required_field` のCLI JSON出力テストを追加
- [x] invalid fixture helperを追加
- [x] JSON diagnostics内のcodeを確認
- [x] validation failedのnon-zero errorを確認
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
