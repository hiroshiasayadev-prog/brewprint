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
| M12: MCP impact traversal / source assist | closed | [tasks/m12-mcp-impact-traversal.md](tasks/m12-mcp-impact-traversal.md) | ADR-054 / ADR-055 / ADR-056 |
| M13: MCP analyze_impact implementation | closed | [tasks/m13-mcp-analyze-impact-implementation.md](tasks/m13-mcp-analyze-impact-implementation.md) | ADR-056; hybrid v1 close; `docs/impl/go-m13-summary.md` |
| M14: brewprint self-hosting | paused | [tasks/m14-self-hosting.md](tasks/m14-self-hosting.md) | UC-002 / v1.0.0-spec実用検証 + editor/viewer要件抽出。M15完了まで一時停止 |
| M14a: subnode file-private scope fix | open | [tasks/m14a-subnode-scope-fix.md](tasks/m14a-subnode-scope-fix.md) | ADR-058 / ADR-059。M14 Phase Aで発覚したv1.0系実装バグ修正 |
| M15: data layer expressiveness (v1.1) | open | [tasks/m15-data-layer-expressiveness.md](tasks/m15-data-layer-expressiveness.md) | ADR-060 (TypeRef + flow wiring type compatibility)。Phase A/B/C構成。完了時点で v1.1.0-spec タグ発行 |
| M16: Design Records MCP MVP | closed | [tasks/m16-design-records-mcp-mvp.md](tasks/m16-design-records-mcp-mvp.md) | ADR-076 / ADR-077。ADR/spec record index + read-only MCP MVP |
| M17: Design Records MCP stdio transport | closed | [tasks/m17-design-records-mcp-stdio-transport.md](tasks/m17-design-records-mcp-stdio-transport.md) | M16 handlers を stdio MCP server として公開 |
| M18: semantic traceability foundation | closed | [tasks/m18-semantic-traceability-foundation.md](tasks/m18-semantic-traceability-foundation.md) | ADR-081〜088。ADR-088 により MVP を canonical reference resolution foundation に縮小し、spec / policy / handoff 同期と最終独立レビュー完了。MCP実装追従は M19 へ分離 |
| M19: Design Records MCP semantic trace support | open | [tasks/m19-design-records-semantic-trace-support.md](tasks/m19-design-records-semantic-trace-support.md) | REQ-MCP-001 / WORK-MCP-001 / ADR-087 / ADR-088。`spec:` + record ID + investigation canonical ref resolve / validation 実装 |

---

## 検討中

（現在検討中の項目はなし。v1後検討事項は ADR-057 §6 / doc-policy.md §11 を参照）
