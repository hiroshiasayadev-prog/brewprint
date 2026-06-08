---
scope: docs/spec/mcp/versioning.md
status: draft
last_updated: 2026-04-30
summary: >
  MCP v1で未定義とする範囲と将来拡張候補を定義する。
  設計対話coverageを広げるための候補toolやselectorを整理する。
depends_on:
  - docs/adr/054-mcp-query-coverage-for-design-conversation.md
  - docs/adr/055-mcp-reference-tree-traversal.md
  - docs/adr/056-mcp-analyze-impact-tool-design.md
---

# Versioning / future extensions

MCP v1では以下を未定義とする。

- unbounded transitive references
- reference graphの永続cache
- renderer outputを返すMCP tool
- code generation用tool

設計対話coverageを広げるための将来候補:

| 候補tool / selector | 目的 | 優先度 |
|---|---|---:|
| `list_objects` | project内objectの検索・一覧。LLMがquery対象を発見する入口 | high |
| `inspect(file)` | YAML file単位で定義内容・main node・sub node・view種別・diagnosticsを把握する | high |
| `inspect(view: api_table)` | API Table viewが集約するmodules / endpoints / computed routesを把握する | high |
| `inspect(view: er_diagram)` | ER Diagram viewが集約するmodules / stores / models / FKを把握する | high |
| flow wiring references | DAG上のflow step / param wiringをreferenceとして扱う | medium |
| `search_notes` | note/docに対するsemantic search | low |

これらは必要になった時点で `docs/spec/mcp/overview.md` または該当する `docs/spec/mcp/tools/*.md` を更新し、実装タスクは `docs/TASKS.md` で管理する。
既存ADRの方針を変更するほどの設計転換がある場合のみ、新ADRを起票する。

## v1に昇格済みの記録

- `analyze_impact` — V01-ADR-056で v1 採用。 仕様は [`tools/analyze-impact.md`](./tools/analyze-impact.md) 参照
