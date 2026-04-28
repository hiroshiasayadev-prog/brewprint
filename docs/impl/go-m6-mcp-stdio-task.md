# Go M6 MCP stdio 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: Go実装 Milestone 6 第2段の MCP stdio transport vertical slice 実装チェックリスト

---

## 1. 目的

M6第2段では、M6第1段で実装した `internal/mcp` のtool dispatch adapterを、最小JSON-RPC stdio入口へ接続する。

対象:

```text
brewprint mcp --yaml-root <path>
```

起動後はstdinから1行1JSON-RPC requestを読み、stdoutへ1行1JSON-RPC responseを返す。

---

## 2. 実装範囲

- `initialize`
- `tools/list`
- `tools/call`
- JSON-RPC error envelope
- `cmd/brewprint mcp --yaml-root <path>`
- handler単位unit test

MCP SDK依存はこの段階では入れない。まずはtransport境界を薄く固定し、必要になった時点でSDK/厳密protocol追随を検討する。

---

## 3. 境界

```text
cmd/brewprint -> source, resolve, query, mcp
mcp           -> query
```

禁止:

```text
internal/mcp -> rawyaml
internal/mcp内でYAML load / resolveを実行
internal/mcp内でrendererを呼ぶ
```

---

## 4. 受け入れ条件

- [x] `initialize` にJSON-RPC responseを返せる
- [x] `tools/list` に4 toolを返せる
- [x] `tools/call` で `get_signature` を呼べる
- [x] `tools/call` で `list_endpoints` を呼べる
- [x] unknown JSON-RPC methodをprotocol errorで返せる
- [x] invalid JSONをparse errorで返せる
- [x] `brewprint mcp --yaml-root <path>` を起動できる
- [x] `internal/mcp` が `rawyaml` をimportしていない
- [x] handler unit testが通る
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る
