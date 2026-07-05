# DRMCP-ADR-MCP-007: Whole-application component model

- **status**: superseded
- **date**: 2026-07-04
- **depends_on**:
  - DRMCP-ADR-MCP-001
- **supersedes**: []
- **migrated_to_spec**: null

## Context

DRMCP exposes eleven current public operations across read, validation, guidance, and authoring concerns.
The application requires stable component boundaries before module contracts or implementation details are designed.

The whole-application boundary includes authoring operations but preserves only their MCP-to-use-case extension seam.
The model must not invent proposal, transaction, package, or interface design.

A component model must also avoid one application-level component for every parser, resolver, validator, index, or source.

## Decision

DRMCP uses one whole-application boundary around all eleven current public operations and their shared application capabilities.

The baseline classifies operations into these families:

| family | operations | architecture treatment |
|---|---|---|
| Read and query | `list_records`, `get_records`, `resolve_reference` | Active application use cases. |
| Validation | `validate_records` | Active application use case. |
| Guidance | `list_authoring_guides`, `get_authoring_guidance` | Active application use cases. |
| Authoring extension | Five proposal and write operations | Preserve the MCP-to-use-case seam only. Concrete internals remain outside this decision. |

The MCP client or host is the inbound actor.
Configured repositories, PRODUCT standards, the guidance package, and runtime configuration are external providers.

The boundary excludes:

- retired operations and tombstones;
- existing brewprint MCP internals;
- model-provider behavior;
- unspecified host-level policy.

DRMCP uses six application-level components:

| component | owned responsibility | excluded responsibility |
|---|---|---|
| Composition / Lifecycle | Construction, dependency wiring, startup, and shutdown. | Operation semantics, record semantics, and source access. |
| MCP Inbound Adapter | MCP protocol mapping and application use-case invocation. | Application policy, domain semantics, and concrete source access. |
| Application Use Cases | Operation sequencing, request-scoped orchestration, and operation-specific result policy. | Record semantics and concrete I/O. |
| Record Domain / Logical Tree | Record models, immutable logical structures, identity, resolution, and validation behavior. | MCP transport, application sequencing, configuration, and I/O. |
| Guidance Domain | Path-independent guide identity and content-projection semantics. | Concrete storage, record snapshots, and retained runtime state. |
| Infrastructure I/O Adapters | Concrete source enumeration and reads behind inward-owned source capabilities. | Operation policy and domain-result classification. |

The six components are responsibility boundaries, not required packages or runtime objects.

Do not create separate application-level components for parsers, resolvers, validators, snapshot builders, indexes, or individual filesystem sources.
A later architecture decision may split a component only for an independent lifecycle, owner, substitution boundary, or cross-component contract.

This decision does not define authoring operation internals, proposal state, write transactions, packages, interfaces, types, functions, or methods.

## Rationale

The six-component model covers the active operation families without coupling protocol handling to application policy.

Separate Record and Guidance domains preserve their different semantics and state needs.
Application Use Cases retain orchestration without absorbing domain rules or concrete source access.

Broad component boundaries avoid layer proliferation.
The boundaries still preserve explicit ownership and future substitution points.

The MCP-to-use-case seam keeps deferred authoring operations inside the whole-application scope without inventing unsupported internal architecture.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Create separate application-level components for parsers, resolvers, validators, snapshot builders, indexes, or every source | Those units lack an accepted independent lifecycle, owner, substitution boundary, or cross-component contract. |

## Consequences

- Canonical architecture Specifications must describe the six components and the whole-application boundary.
- Active operations receive concrete placement within the accepted component graph.
- Authoring operations retain only an inbound-adapter-to-use-case extension seam.
- Downstream contract work may refine internals without creating new top-level components.
- A top-level component split or ownership move requires a return to application-architecture decision work.
- This ADR does not establish package, interface, type, function, method, or transaction contracts.

Affected design areas:

- `spec:drmcp.application_architecture`;
- `spec:drmcp.application_architecture.application_boundary_and_components`;
- `spec:drmcp.application_architecture.dependency_and_responsibility`.

## Evidence

- Source Requirement: `DRMCP-REQ-MCP-005`.
- Source Work Item: `DRMCP-WORK-MCP-016`.
- Routed decisions: D-001, D-003, D-005, and D-006 in `DRMCP-TASK-MCP-016-03`.
- ADR routing authority: B-01 in `DRMCP-TASK-MCP-016-05`.
- Contract baseline authority: `DRMCP-ADR-MCP-001`.
