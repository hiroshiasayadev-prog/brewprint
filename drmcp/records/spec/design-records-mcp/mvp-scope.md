# Reference: MVP scope

- **id**: `spec:drmcp.design_records_mcp.mvp_scope`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Defines the MVP P0/P1 tool set and the list of items explicitly outside MVP for Design Records MCP.

## Current contract

### P0 tools (MVP required)

The following tools are required in MVP. All P0 tools are read-only; their purpose is candidate narrowing and metadata consistency validation before reading record body.

- `list_records`
- `get_record`
- `get_records`
- `validate_records`
- `resolve_reference`

> Source: V01-ADR-077 §P0: MVP必須tool, V01-ADR-090 §決定

### P1 tools (optional in MVP)

The following auxiliary tools may be included in MVP but are not required.

- `suggest_next_record`

> Source: V01-ADR-077 §P1: MVPに含めてもよい補助tool

### MVP exclusions

The following are explicitly out of scope for MVP.

| item | category |
|---|---|
| `trace_record` | write / traversal |
| `list_gaps` | gap analysis |
| `create_record` | write |
| `update_record` | write |
| `set_evidence` | write |
| Other write tools | write |
| Inferring dependencies from natural language body | NLP analysis |
| Strict semantic verification against spec body | NLP analysis |
| Git history analysis | git |
| Code static analysis | static analysis |
| Web UI | UI |
| Multi-project management | scope expansion |
| Public OSS CLI contract | contract |
| Full section-level traceability | traceability |
| `topics` / `affects` / `refines` / `conflicts_with` metadata | metadata |
| Legacy M-series task records as indexed record kind | indexing |
| UC docs / impl notes as indexed record kind | indexing |
| `internal-design:` / `coverage:` / `COV-*` resolve and semantic realization relation validation | ref resolve |
| Coverage mapping query | ref resolve |
| Orphan requirement / orphan work item / orphan task diagnostics | diagnostics |
| Deriving work item progress from task status | projection |
| Workflow-specific traversal / tree / graph query tools | traversal |
| Task dependency cycle detection / execution order projection | analysis |
| Routing `TASK-*` as canonical references from investigation metadata | ref resolve |

MVP operates solely on explicit information obtainable from ADR / investigation / requirement / work item / task bullet metadata, spec YAML front matter, H1 titles, and paths. Natural language body inference and operational gap diagnostics will be evaluated for inclusion after MVP validation against real data.

> Source: V01-ADR-076 §MVPスコープ外, V01-ADR-077 §MVP外, V01-ADR-092 §7
