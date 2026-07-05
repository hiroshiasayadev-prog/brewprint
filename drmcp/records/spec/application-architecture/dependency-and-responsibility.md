# Concept: Dependency and responsibility

- **id**: `spec:drmcp.application_architecture.dependency_and_responsibility`
- **status**: draft
- **date**: 2026-07-04
- **parent**: `spec:drmcp.application_architecture`

## What this is

Inward dependency and policy-placement contract for DRMCP. This view owns allowed dependency direction, source-port ownership, standards placement, Guidance placement, and forbidden edges.

## Concept model

Dependencies point inward toward application and domain policy.

| dependent | allowed dependency | purpose |
|---|---|---|
| MCP Inbound Adapter | Application Use Cases | Map protocol requests to public application behavior. |
| Application Use Cases | Record Domain / Logical Tree | Invoke current record identity, query, retrieval, resolution, section, and validation behavior. |
| Application Use Cases | Application-owned source ports | Request only the external capabilities required by a use case. |
| Infrastructure I/O Adapters | Application-owned source-port contracts | Implement concrete source access without owning operation policy. |
| Composition / Lifecycle | Every concrete component | Construct and wire the application only. |

The detailed component responsibilities and exclusions are owned by `spec:drmcp.application_architecture.application_boundary_and_components`.

## Rules

### Dependency direction

- MCP protocol types do not enter Application Use Cases or Domain.
- Application Use Cases do not import concrete Infrastructure I/O Adapters.
- Record Domain / Logical Tree does not depend on Application, MCP, filesystem, or runtime configuration.
- Infrastructure I/O Adapters may depend inward on source-port contracts.
- Composition / Lifecycle is the only owner that wires concrete components.
- Runtime collaboration between operation-specific use cases and shared orchestration is defined by `spec:drmcp.application_architecture.runtime_and_state`.

### PRODUCT standards

PRODUCT is the external semantic authority for Design Records standards.
The portable package is a copied and ref-rewritten distribution of that authority.

| responsibility | owner |
|---|---|
| Define authoritative semantic rules and package-generation inputs | PRODUCT standards. |
| Generate the fixed `spec:design_records.*` package tree | PRODUCT package-production workflow. |
| Select and register the package spec-tree source as `design_records` | Composition / Lifecycle. |
| Enumerate and read package files through normal Current Records access | Infrastructure I/O Adapters. |
| Own parsed current Spec identity, logical structure, resolution, and validation semantics | Record Domain / Logical Tree. |
| Select package records and project operation-specific behavior | Application Use Cases. |

DRMCP does not rewrite package refs, redefine PRODUCT semantics, or maintain a second standards catalog.
MCP Inbound Adapter has no PRODUCT-standard interpretation responsibility.

### Guidance

Guidance is a fixed-scope application projection over normal Current Records.

- `list_authoring_guides` and `get_authoring_guidance` are operation-specific Application Use Cases.
- They reuse shared record-query orchestration and Record Domain behavior.
- They do not call the public `list_records` or `get_records` use cases.
- List scope fixes app namespace `design_records`, kind `spec`, and the `spec:design_records.authoring_standards.*` child subtree.
- List scope excludes `spec:design_records.authoring_standards` itself.
- Detail lookup accepts one exact canonical ref inside that child subtree.
- Record Domain owns canonical identity, parsed H1, headings, sections, and complete body representation.
- Application owns fixed scope, ASCII canonical-ref ordering, response projection, and operation errors.
- `id` is the canonical package Spec ref.
- `title` is first H1 text.
- `abstract` is the `## What this is` section body.
- `content` is complete Markdown verbatim.
- Guidance does not introduce a separate source port, domain, index, snapshot, cache, or lifecycle.

Concrete source access uses the normal Current Records source contracts.

### Forbidden dependencies and orchestration

| pattern | status |
|---|---|
| Application Use Case imports a concrete filesystem adapter | Prohibited. |
| Domain performs I/O | Prohibited. |
| Infrastructure adapter calls a public MCP tool | Prohibited. |
| Infrastructure adapters orchestrate laterally | Prohibited. |
| MCP Inbound Adapter reclassifies semantic outcomes or execution failures | Prohibited. |
| Application Use Cases interpret PRODUCT standards independently | Prohibited. |
| One public use case calls another public use case | Prohibited. See the collaboration contract in `spec:drmcp.application_architecture.runtime_and_state`. |
| A component bypasses an application-owned source port for concrete I/O | Prohibited. |

## Boundary

This view defines application-level dependency ownership. It does not select package imports, interface signatures, constructor shapes, adapter APIs, or rule encodings.

The accepted responsibility table remains in `spec:drmcp.application_architecture.application_boundary_and_components`. This view does not redefine those component boundaries.

## Non-goals

- Concrete port definitions.
- Package dependency graphs.
- Standards catalog formats.
- Concrete Guidance request serialization and error payloads.
- Parser, resolver, or validator algorithms.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.application_architecture` | Parent Overview and authoritative four-view map. |
| `spec:drmcp.application_architecture.application_boundary_and_components` | Owns the five components and their broad responsibilities. |
| `spec:drmcp.application_architecture.runtime_and_state` | Owns shared orchestration and request-state collaboration. |
| `spec:drmcp.application_architecture.failure_and_evolution` | Owns failure projection and dependency-change return triggers. |
