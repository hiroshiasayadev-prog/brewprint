# Go M6 MCP wrapper 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: Go実装 Milestone 6 の MCP server wrapper vertical slice 実装チェックリスト

---

## 1. 目的

M6では、QueryServiceをMCP toolとして公開する薄いwrapperを実装する。

まずは実transportに入る前に、tool名 + JSON input を QueryService 呼び出しへdispatchし、JSON output / JSON errorを返せるところまでを縦切りで通す。
この段階ではstdio server本体ではなく、unit test可能な薄いadapterを先に固定する。

対象tool:

```text
get_signature
get_references
inspect
list_endpoints
```

`MCP wrapper` は transport / protocol adapter であり、semantic queryの実体は `internal/query` に閉じる。

---

## 2. 実装範囲

- QueryService `ListEndpoints`
- `internal/mcp` package
- tool name dispatch
- JSON args decode
- JSON result encode
- not_found / invalid_args / unknown_tool のJSON error envelope
- UC-001 fixtureを使ったwrapper unit test

---

## 3. 境界

```text
source  -> rawyaml
resolve -> rawyaml, semantic
query   -> semantic
mcp     -> query
```

禁止:

```text
mcp -> rawyaml
mcp内でname resolution / semantic traversalを再実装
mcp内でrendererを呼ぶ
```

---

## 4. 受け入れ条件

- [x] `get_signature` をJSON tool callで呼べる
- [x] `get_references` をJSON tool callで呼べる
- [x] `inspect` をJSON tool callで呼べる
- [x] `list_endpoints` をJSON tool callで呼べる
- [x] unknown toolをJSON errorで返せる
- [x] invalid argsをJSON errorで返せる
- [x] not foundをJSON errorで返せる
- [x] `internal/mcp` が `rawyaml` をimportしていない
- [x] wrapper unit testが通る
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る
