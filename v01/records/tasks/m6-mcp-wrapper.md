# Milestone 6: MCP server wrapper を実装する

- **status**: closed
- **scope**: MCP wrapper
- **source**: migrated from docs/TASKS.md
- **last_updated**: 2026-04-30

---

## Tasks

- [x] **MCP server wrapper を実装する**
  - QueryServiceをMCP toolとして公開する
  - JSON tool dispatch adapter 第1段完了
    - `get_signature`
    - `get_references`
    - `inspect`
    - `list_endpoints`
  - unknown / invalid args / not found のJSON error envelope
  - `internal/mcp -> query` 境界確認済み
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M6-2時点）
  - MCP stdio JSON-RPC 第1段完了
    - `initialize`
    - `tools/list`
    - `tools/call`
    - `brewprint mcp --yaml-root <path>`
  - MCP実用化仕上げ完了
    - QueryService response JSON field名をMCP spec寄りに整理
    - `tools/list` に `inputSchema` を追加
    - `tools/call` responseを `content: [{type:"text", text:"..."}]` / `isError` 形式に整理
    - `unsupported_direction` / `unsupported_detail` のtool error分類を追加
    - CLI usage / `--yaml-root` errorを整理
    - 起動例docを追加
  - `go fmt ./...` / `go test ./...` 通過済み（2026-04-28、M6-3差分後）
  - renderer / QueryService / MCP wrapper の責務が混ざっていないことを確認する
