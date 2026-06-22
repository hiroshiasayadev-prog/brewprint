# Concept: DSL design philosophy

- **id**: `spec:bpdsl.dsl.design_philosophy`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.overview`

## What this is

Foundational design constraints behind brewprint: why it prioritizes machine clarity over human ergonomics, and the invariants that follow from treating YAML as the single source of truth.

## Concept model

### YAML as single source of truth

YAML is the canonical representation of system design. All other artifacts — diagrams, MCP responses, and implementation context — are derived from it.

Invariants:
- An implementation without a prior YAML definition is architecturally invalid.
- No diagram exists independently of a YAML node.
- Design and documentation divergence is structurally prevented.

Analogues: Terraform (infrastructure as code), Prisma (schema as single source of truth).

### AI-first design

Brewprint treats AI as the primary implementer. The governing principle:

> Prioritize "AI does not get confused" over "human can write it easily."

Design decisions derived from this principle:

| decision | rule |
|---|---|
| Static verifiability | ID references must be unique and unambiguous; verifiability takes precedence over expressiveness. |
| Inline definition prohibited | Named definitions required even for disposable structs (e.g., anonymous payload shapes). |
| Ambiguous notation eliminated | Any notation admitting multiple interpretations is not adopted. |

Human authoring ergonomics are addressed at the UI/tool layer, not in the YAML schema itself.

### Happy-path boundary

Brewprint captures happy-path structure only. Exceptions, concurrency, rollback, and bidirectional sync are implementation-level concerns expressed via `note` fields or impl task entries, not as DSL constructs.

### Scope boundary

Brewprint's responsibility ends at design-language and MCP context supply. Code generation is explicitly out of scope.

```text
brewprint (design language + MCP)
    ↓ supplies design context
Claude Code (implements per CLAUDE.md)
    ↓ halts on ambiguity
impl_tasks.md (pending implementation items)
```

## Boundary

| in scope | out of scope |
|---|---|
| YAML design contract | Code generation |
| MCP context supply | Concrete style or visual layout |
| Static structural validation | Concurrency, rollback, bidirectional sync |
| note-based semantic annotation | Runtime validation of note content |

## Non-goals

Human-ergonomic YAML syntax is not a goal.

Non-functional attributes (`retry`, `idempotent`, `async`) as first-class fields are deferred to post-dogfood evaluation.

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.dsl.overview` | Parent overview for the DSL spec area. |
| `spec:bpdsl.dsl.nodes.overview` | Node definitions built on these constraints. |
| `spec:bpdsl.dsl.edges.overview` | Edge syntax built on these constraints. |
