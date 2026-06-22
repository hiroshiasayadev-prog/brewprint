# Reference: Versioning / future extensions

- **id**: `spec:bpdsl.mcp.versioning`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.mcp.overview`

## What this is

What MCP v1 leaves undefined, and the candidate tools/selectors under consideration for extending design-conversation coverage in future versions.

> Source: V01-ADR-054, V01-ADR-055, V01-ADR-056

## Undefined in v1

MCP v1 leaves the following undefined:

- Unbounded transitive references.
- A persistent cache of the reference graph.
- An MCP tool that returns renderer output.
- A code-generation tool.

## Future candidates for design-conversation coverage

| candidate tool / selector | purpose | priority |
|---|---|---:|
| `list_objects` | Search / list objects in the project. Entry point for an LLM to discover query targets. | high |
| `inspect(file)` | Grasp a YAML file's definition contents, main node, sub nodes, view kind, and diagnostics at file granularity. | high |
| `inspect(view: api_table)` | Grasp the modules / endpoints / computed routes aggregated by an API Table view. | high |
| `inspect(view: er_diagram)` | Grasp the modules / stores / models / FKs aggregated by an ER Diagram view. | high |
| flow wiring references | Treat DAG flow-step / param wiring as a reference. | medium |
| `search_notes` | Semantic search over note/doc text. | low |

When any of these becomes needed, update `bpdsl/records/spec/mcp/overview.md` or the relevant `bpdsl/records/spec/mcp/tools/*.md` and track the implementation task separately. Only file a new ADR if the change represents a design shift significant enough to alter an existing ADR's policy.

## Already promoted to v1

- `analyze_impact` — adopted into v1 by V01-ADR-056. Spec: [`tools/analyze-impact.md`](tools/analyze-impact.md).

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.mcp.overview` | Parent overview; tool catalog. |
| `spec:bpdsl.mcp.tools.analyze_impact` | Tool spec promoted to v1, referenced here as a precedent. |
