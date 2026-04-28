# Go M7-2 diagnostics CLI 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: validation diagnosticsをCLIから確認する入口の追加

---

## 1. 目的

M7-1で追加したsemantic diagnosticsを、MCP起動やrenderer実行より前に確認できるようにする。

M7-2では、最小のCLI入口として以下を追加する。

```sh
brewprint validate --yaml-root <path>
```

---

## 2. 実装範囲

- `brewprint validate --yaml-root <path>` を追加
- `source.Loader` + `resolve.Build` のみを実行する
- diagnosticsが空なら `ok` をstdoutへ出す
- diagnosticsがあれば `severity code file: message` 形式でstdoutへ出す
- error diagnosticsが1件以上あればnon-zero errorを返す
- MCP起動時のsemantic diagnostic errorにもcodeを含める
- CLI unit testを追加

---

## 3. 境界

```text
cmd/brewprint -> source, resolve
```

`validate` commandでは、query / mcp / renderer を呼ばない。

---

## 4. 出力例

成功時:

```text
ok
```

失敗時:

```text
error unresolved_model order/state.yaml: unresolved event payload model: payment_event
```

終了時error:

```text
validation failed: 1 error(s), 0 warning(s)
```

---

## 5. 受け入れ条件

- [x] `brewprint validate --yaml-root <path>` が追加される
- [x] diagnosticsなしなら `ok` を出す
- [x] diagnostic formatterが `severity code file: message` を返す
- [x] missing `--yaml-root` がerrorになる
- [x] MCP起動時のsemantic diagnostic errorにcodeを含める
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
