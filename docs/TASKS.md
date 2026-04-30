# TASKS

brewprintの積みタスク一覧。
詳細は milestone 別ファイルを参照する。

---

## 読み方

- セッション開始時はこのファイルだけ読む
- 作業対象 milestone が決まってから `docs/tasks/mXX-*.md` を読む
- closed milestone は原則読まない
- 実装後の詳細な引継ぎは `docs/impl/go-mX-summary.md` を参照する
- task file は未来向きのチェックリスト、impl summary は過去向きの実装引継ぎ

---

## Status

- `open`: 未完了taskあり
- `closed`: milestone内task完了
- `paused`: 一時停止中

---

## Milestones

| milestone | status | detail | notes |
|---|---|---|---|
| M0: implementation boundary | closed | [tasks/m0-boundary.md](tasks/m0-boundary.md) | ADR-047〜049 |
| M1: DAG vertical slice | closed | [tasks/m1-dag-vertical-slice.md](tasks/m1-dag-vertical-slice.md) | UC-001 auth.login |
| M2: DAG renderer expansion | closed | [tasks/m2-dag-renderer.md](tasks/m2-dag-renderer.md) | foreach / fork / branch |
| M3: QueryService | closed | [tasks/m3-query-service.md](tasks/m3-query-service.md) | `docs/impl/go-m9-summary.md` も参照 |
| M4: render_index / placement | closed | [tasks/m4-render-index-placement.md](tasks/m4-render-index-placement.md) | ADR-045 / 046 |
| M5: renderers | closed | [tasks/m5-renderers.md](tasks/m5-renderers.md) | State / Sequence / ER / API / Wireframe |
| M6: MCP wrapper | closed | [tasks/m6-mcp-wrapper.md](tasks/m6-mcp-wrapper.md) | stdio JSON-RPC |
| M7: diagnostics | closed | [tasks/m7-diagnostics.md](tasks/m7-diagnostics.md) | validate CLI / diagnostic codes |
| M8: render CLI / pipeline | closed | [tasks/m8-render-cli.md](tasks/m8-render-cli.md) | M8 tasks completed |
| M9: QueryService state/event/scenario expansion | closed | [tasks/m9-query-service-expansion.md](tasks/m9-query-service-expansion.md) | `docs/impl/go-m9-summary.md` |
| M10: MCP project exploration / view inspect | closed | [tasks/m10-mcp-project-exploration.md](tasks/m10-mcp-project-exploration.md) | ADR-054 |
| M11: MCP diagram element query | closed | [tasks/m11-mcp-diagram-element-query.md](tasks/m11-mcp-diagram-element-query.md) | ADR-054 |
| M12: MCP impact traversal / source assist | open | [tasks/m12-mcp-impact-traversal.md](tasks/m12-mcp-impact-traversal.md) | ADR-054 |
