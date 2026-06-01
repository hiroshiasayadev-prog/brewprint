# TASKS

brewprint の旧 M-series index。

ADR-091 により、新規の実行計画と到達点は work item が所有し、具体作業は work item 配下の短期 task として管理する。下表の `M*` は移行前の歴史的ラベルであり、参照先の `docs/tasks/m*.md` は work item 相当の計画と具体 task が混在した legacy record である。

---

## 読み方

- セッション開始時はこのファイルだけ読む
- 新規 work は requirement -> work item -> `docs/tasks/<domain>/TASK-*.md` の短期 task として追跡する
- 旧 M-series 記録の確認が必要な場合のみ、対応する `docs/tasks/mXX-*.md` を legacy record として読む
- closed な旧 M-series 記録の詳細は原則読まない
- 実装後の詳細な引継ぎは `docs/impl/go-mX-summary.md` を参照する
- 新形式 task は未来向きの短期実行単位、legacy record は旧計画・完了履歴、impl summary は過去向きの実装引継ぎとして扱う

---

## Legacy M-series status

- `open`: 当該 legacy record が未完了である
- `closed`: 当該 legacy record が完了済みである
- `paused`: 当該 legacy record を一時停止中である

この status は旧 M-series 記録の状態であり、新形式の work item または task artifact の `status` contract ではない。

---

## Legacy M-series records

| legacy label | status | detail | notes |
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
| M14a: subnode file-private scope fix | closed | [tasks/m14a-subnode-scope-fix.md](tasks/m14a-subnode-scope-fix.md) | ADR-058 / ADR-059。B1 は REQ-RESOLVE-001 / WORK-RESOLVE-001 で解消。B2 は後続 v1.1 系 TypeRef / return handling に回収済み |
| M15: data layer expressiveness (v1.1) | closed | [tasks/m15-data-layer-expressiveness.md](tasks/m15-data-layer-expressiveness.md) | REQ-DATA-001 / WORK-DATA-001 により minimum-expressiveness release として close。v1.1.0-spec は commit 後 tag ready |
| M16: Design Records MCP MVP | closed | [tasks/m16-design-records-mcp-mvp.md](tasks/m16-design-records-mcp-mvp.md) | ADR-076 / ADR-077。ADR/spec record index + read-only MCP MVP |
| M17: Design Records MCP stdio transport | closed | [tasks/m17-design-records-mcp-stdio-transport.md](tasks/m17-design-records-mcp-stdio-transport.md) | M16 handlers を stdio MCP server として公開 |
| M18: semantic traceability foundation | closed | [tasks/m18-semantic-traceability-foundation.md](tasks/m18-semantic-traceability-foundation.md) | ADR-081〜088。ADR-088 により MVP を canonical reference resolution foundation に縮小し、spec / policy / handoff 同期と最終独立レビュー完了。MCP実装追従は M19 へ分離 |
| M19: Design Records MCP semantic trace support | closed | [tasks/m19-design-records-semantic-trace-support.md](tasks/m19-design-records-semantic-trace-support.md) | REQ-MCP-001 / WORK-MCP-001 / ADR-087 / ADR-088。investigation integration + canonical ref resolve / validation 実装・runtime validation確認完了 |

---

## Migration note

- 既存 `docs/tasks/m*.md` の archive 化は未実施であり、別 migration work で扱う。
- Open / paused な M14 は、再開時に `WORK-*` / short `TASK-*` へ明示的に移行する。
- 新規 task authoring guidance は guide ID `task-authoring`、work item authoring guidance は guide ID `work-item-authoring` を Design Records MCP 経由で参照する。
- 新形式に `milestone` field、milestone artifact、または milestone relation は導入しない。

## 検討中

- ADR-091 追従として、legacy M-series record の archive 移行と open record の分解方針を後続 work で扱う。
