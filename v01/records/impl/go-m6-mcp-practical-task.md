# Go M6-3 MCP実用化仕上げ 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: Go実装 Milestone 6 第3段の MCP互換寄せ・CLI usage整理・起動例整理

---

## 1. 目的

M6-1 / M6-2で実装したMCP tool dispatch adapterとstdio JSON-RPC transportを、実際のMCP clientから使いやすい形へ寄せる。

対象:

```text
brewprint mcp --yaml-root <path>
```

MCP server wrapperは引き続きtransport / protocol adapterであり、semantic queryの実体は `internal/query` に閉じる。

---

## 2. 実装範囲

- QueryService response JSON field名をMCP spec寄りに整理
- `tools/list` responseに `inputSchema` を追加
- `tools/call` responseをMCP互換寄りに整理
  - `content: [{ "type": "text", "text": "..." }]`
  - `isError: true` をtool error時に返す
  - 成功時の `text` はtool result本体のJSON文字列
  - 失敗時の `text` は `{ "error": ... }` のJSON文字列
- `unsupported_direction` / `unsupported_detail` のtool error分類追加
- CLI usage / `--yaml-root` error整理
- 起動例doc追加

---

## 3. 境界

```text
cmd/brewprint -> source, resolve, query, mcp
mcp           -> query
query         -> semantic
```

禁止:

```text
internal/mcp -> rawyaml
internal/mcp内でYAML load / resolveを実行
internal/mcp内でrendererを呼ぶ
```

今回の変更でも、YAML load / resolveは `cmd/brewprint` 側に残し、`internal/mcp` は `query.Service` を呼ぶだけにする。

---

## 4. 起動例

UC-001 fixtureを使う場合:

```sh
brewprint mcp --yaml-root docs/uc/001-ec-checkout-flow/yaml
```

手動疎通例:

```sh
printf '%s\n' \
  '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}' \
  '{"jsonrpc":"2.0","id":2,"method":"tools/list"}' \
  '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"get_signature","arguments":{"selector":{"id":"auth.task.login"}}}}' \
  | brewprint mcp --yaml-root docs/uc/001-ec-checkout-flow/yaml
```

`tools/call` 成功時の `result.content[0].text` は、tool result本体のJSON文字列になる。

例:

```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "{\"object\":{...},\"signature\":{...},\"diagnostics\":[]}"
      }
    ]
  }
}
```

Tool error時はJSON-RPC protocol errorではなく、`isError: true` のtool resultとして返す。

```json
{
  "jsonrpc": "2.0",
  "id": 4,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "{\"error\":{\"code\":\"unknown_tool\",\"message\":\"unknown tool: missing_tool\",\"tool\":\"missing_tool\"}}"
      }
    ],
    "isError": true
  }
}
```

JSON-RPCとして壊れている入力はprotocol errorで返す。

- invalid JSON: `-32700 parse error`
- unknown method: `-32601 method not found`
- invalid params: `-32602 invalid params`

---

## 5. 受け入れ条件

- [x] QueryService response JSON field名がsnake_case / lower camel寄りになる
- [x] `tools/list` が4 toolと `inputSchema` を返す
- [x] `tools/call` 成功時に `content[0].text` へtool result本体JSONを返す
- [x] `tools/call` tool error時に `isError: true` と `{ "error": ... }` text JSONを返す
- [x] unknown toolをtool errorで返す
- [x] invalid argsをtool errorで返す
- [x] not foundをtool errorで返す
- [x] unsupported directionを `unsupported_direction` で返す
- [x] unsupported inspect detailを `unsupported_detail` で返す
- [x] `brewprint mcp --yaml-root <path>` のusage / missing arg errorを整理する
- [x] `internal/mcp` が `rawyaml` をimportしていない
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
