# Overview: MCP query layer

- **id**: `spec:bpdsl.mcp.overview`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.overview`

## What this is

External I/O contract for the MCP query tools brewprint exposes to LLMs. Covers MCP tool names, tool input schema, tool output schema, common ID representation, common response vocabulary, reference representation, diagnostic/error representation, and the intent behind how LLMs should use each tool.

Out of scope: Go package / struct / interface names, raw YAML decode structs, `ResolvedProject` internal map/index implementation, renderer Mermaid/HTML/Markdown output spec, MCP transport implementation details, transitive dependency graph pre-build strategy.

This spec assumes the Go implementation boundary defined in V01-ADR-047 / V01-ADR-048 and the reference-vocabulary unification defined in V01-ADR-049, and fixes the shape of information `QueryService` returns externally.

## Current contract

### Design principles

| principle | summary |
|---|---|
| MCP does not expose the raw YAML AST | MCP tools always return information from `ResolvedProject` after semantic build (load/classify/decode → validate/name-resolution/derived-model-build/index-build → `ResolvedProject` → `QueryService` → MCP response). References in MCP responses use resolved IDs by default; unresolved raw strings from YAML are returned only as auxiliary info when needed for diagnostics or source display. |
| Adopt Python `inspect`-style vocabulary | MCP response shape mirrors Python's `inspect` introspection feel, so LLMs can reuse vocabulary they already know: `signature` (object's external shape — params/returns/fields/endpoint etc.), `doc` (natural-language description derived from YAML `note` — a semantic contract, not machine-verified structure), `source` (defining file/line/column etc.), `members` (elements the object contains — sub task/fields/transitions etc.), `references` (relations the object refers to or is referred by), `diagnostics` (warning/error/hint info). brewprint MCP is not a Python-AST-compatible API — it exposes semantic objects on `ResolvedProject`, not a syntax tree. |
| Center on "references", not "dependencies" | MCP responses group dependency/reference/reverse-reference relations under the single term `references`. Rationale: `dependency` skews toward "build dependency" / "runtime dependency" / "type dependency"; brewprint relations like `reads` / `writes` / `transition.action` / `model field type` / `scenario step` read more naturally as references than dependencies; vocabulary close to the IDE/Python `references` concept is easier for an LLM to interpret. Per V01-ADR-049, the external MCP tool name is unified to `get_references` and the internal QueryService method to `GetReferences`; `get_deps` / `GetDeps` are not adopted. |
| Separate structural info from doc | MCP responses separate mechanically-determined structural information from natural-language description derived from `note`. `doc` is an important semantic contract for the LLM but must not be treated as a machine-verified fact. |
| v1 references are direct-only | Per V01-ADR-048 / V01-ADR-049, MCP v1 does not pre-build a full transitive dependency graph. `get_references` returns **direct references only** in the initial spec. Transitive closure / depth specification / dependency-graph caching are separate extensions to be added once real need emerges in the QueryService vertical slice. |
| Design-conversation coverage drives extension | MCP / QueryService is not just an implementation-helper API — it is a query layer for design conversation with an LLM while looking at diagrams/views (DAG / State Diagram / Sequence Diagram / ER / API Table / Wireframe). Major semantic objects that appear in rendered diagrams/views should in principle be queryable from MCP (task / model / store / state / event / actor, model field, transition, sequence scenario view, API Table view, ER Diagram view, implicit asset, file-local sub task / branch / fork / join, flow entry / flow wiring, source file). Not everything needs immediate v1 implementation, but future MCP extension should prioritize "is this object visible to the user in a diagram/view and a plausible conversation subject." MCP responses continue to avoid exposing the raw YAML AST — even source-snippet retrieval is treated as source auxiliary info attached to a semantic object. |

> Source: V01-ADR-054 §Decision

### Tool overview

MCP v1 defines the following 8 query tools.

| tool | purpose | typical use |
|---|---|---|
| [`list_objects`](tools/list-objects.md) | Get a list of semantic objects in the project. | Starting point for implementation / design conversation — finding the target object. |
| [`get_signature`](tools/get-signature.md) | Get the external shape of a single object. | Check a task/model/store's types and I/O before implementing. |
| [`get_source`](tools/get-source.md) | Get the YAML snippet corresponding to a semantic object. | Check the defining YAML during design conversation. |
| [`get_references`](tools/get-references.md) | Get an object's direct references. | Check impact range / dependencies / reverse references. |
| [`get_reference_tree`](tools/get-reference-tree.md) | Walk the reference graph from an object with a depth limit. | Check change impact or surrounding objects N hops out. |
| [`analyze_impact`](tools/analyze-impact.md) | Get an interpreted impact analysis based on a change kind. | Judge "what breaks / how to fix it" during a design-change discussion. |
| [`inspect`](tools/inspect.md) | Get implementation-judgment context per object kind. | Read by Claude Code etc. when implementing/fixing. |
| [`list_endpoints`](tools/list-endpoints.md) | Get the endpoint list based on the API Table view. | API implementation / routing check. |

### Tool selection guidance for LLM

LLMs should use the following basic decision rules.

| situation | tool to use |
|---|---|
| Just need to check a target node's I/O | [`get_signature`](tools/get-signature.md) |
| Need to check the defining YAML snippet of a target object | [`get_source`](tools/get-source.md) |
| Need to check what it depends on / what references it | [`get_references`](tools/get-references.md) |
| Need to check change impact or surrounding objects N hops out | [`get_reference_tree`](tools/get-reference-tree.md) |
| Need to judge the impact and fix of a design change (rename / remove / type change etc.) | [`analyze_impact`](tools/analyze-impact.md) |
| Need surrounding context for implementation / fix / review | [`inspect`](tools/inspect.md) |
| Need a list of API routes | [`list_endpoints`](tools/list-endpoints.md) |

Principles:

- Use `inspect` first, before implementing.
- For a small type check only, use `get_signature`.
- For direct-reference checks, use `get_references(direction="in")` or `both`.
- For N-hop impact-range checks, use `get_reference_tree` and specify `direction` and `depth` explicitly.
- For design-change discussions, use `analyze_impact` and specify `change.kind` explicitly; drop down to `get_reference_tree` only when raw reference exploration is needed.
- Before reading raw YAML directly, first check the snippet corresponding to the semantic object via `get_source`.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Error model | Reference | `spec:bpdsl.mcp.errors` | MCP-level error vs. diagnostic distinction, error code catalog, error payload shape. |
| MCP schema | Reference | `spec:bpdsl.mcp.schema` | Common schema: selector, ObjectRef, Reference, Diagnostic, QualifiedID/FileID/synthetic ID representations. |
| Versioning / future extensions | Reference | `spec:bpdsl.mcp.versioning` | What v1 leaves undefined and candidate future tools/selectors. |
| `list_objects` | Contract | `spec:bpdsl.mcp.tools.list_objects` | Get a list of semantic objects in the project. |
| `get_signature` | Contract | `spec:bpdsl.mcp.tools.get_signature` | Get the external shape of a single object. |
| `get_source` | Contract | `spec:bpdsl.mcp.tools.get_source` | Get the YAML snippet corresponding to a semantic object. |
| `get_references` | Contract | `spec:bpdsl.mcp.tools.get_references` | Get an object's direct references. |
| `get_reference_tree` | Contract | `spec:bpdsl.mcp.tools.get_reference_tree` | Walk the reference graph from an object with a depth limit. |
| `analyze_impact` | Contract | `spec:bpdsl.mcp.tools.analyze_impact` | Get an interpreted impact analysis based on a change kind. |
| `inspect` | Contract | `spec:bpdsl.mcp.tools.inspect` | Get implementation-judgment context per object kind. |
| `list_endpoints` | Contract | `spec:bpdsl.mcp.tools.list_endpoints` | Get the endpoint list based on the API Table view. |
