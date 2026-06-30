# DRMCP-ADR-MCP-003: Layered read-runtime application and adapter boundaries

- **status**: accepted
- **date**: 2026-06-30
- **depends_on**:
  - DRMCP-ADR-MCP-001
  - DRMCP-ADR-MCP-002
- **supersedes**: []
- **migrated_to_spec**: 2026-06-30

## Context

The rebuild requires an internal architecture that preserves the accepted public read contracts without inheriting the retired W009 implementation structure.

MCP transport, operation orchestration, record semantics, and filesystem access have different change drivers.
A direct MCP-to-filesystem design would couple public protocol types to scanning and indexing.
A generic application router would duplicate MCP dispatch without adding a separate responsibility.

## Decision

The W011 read-runtime slice for `list_records`, `get_records`, `resolve_reference`, and `validate_records` uses these component boundaries:

```text
composition root
├─ inbound adapter: MCP
├─ application
│  ├─ operation use cases
│  └─ request-snapshot orchestration
├─ core
│  ├─ source and parsed models
│  ├─ parsers
│  ├─ current index
│  ├─ legacy lookup
│  ├─ resolver
│  └─ validation
└─ outbound adapters
   ├─ filesystem
   └─ configuration
```

Dependencies point inward:

```text
MCP adapter ───────┐
filesystem adapter ├──> application ───> core
config adapter ────┘
composition root wires every component
core depends on no outer component
```

Each of `list_records`, `get_records`, `resolve_reference`, and `validate_records` invokes one dedicated application use case.
The MCP adapter maps public request and response schemas for those operations to transport-neutral application contracts.
No generic App Router is introduced for those four operations.
Their MCP handlers do not call core services or outbound adapters directly.

This ADR does not decide application use cases, snapshot lifecycle, or package placement for authoring-guidance or authoring-transaction operations.
Other MCP tool architecture remains with the owning Requirement or Work Item.

Application use cases depend on application-owned narrow ports for:

- configuration access;
- source enumeration;
- source reading.

Filesystem traversal, file reading, and concrete configuration loading belong to outbound adapters.
Parsers and index builders remain pure core services where practical.

The architecture does not define one generic Repository interface that combines unrelated configuration, source, index, resolver, and validation operations.

The composition root owns construction and wiring only.
Outbound adapters do not invoke application use cases.
Core does not import application, MCP, filesystem, or configuration packages.

## Rationale

Dedicated use cases preserve one operation boundary for each of the four W011 read-runtime operations.
The MCP adapter remains a protocol adapter instead of a semantic owner.

Application-owned ports let orchestration define the capabilities it needs.
Narrow ports avoid coupling use cases to a concrete filesystem or a speculative all-purpose repository abstraction.

Pure core services keep parsing, indexing, resolution, and validation testable without MCP or filesystem setup.
The same boundaries permit a future database-backed source adapter without changing operation sequencing.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| MCP handlers call scanners, indexes, or validators directly | Transport would own application sequencing and infrastructure coupling. |
| Introduce one generic App Router | MCP SDK dispatch already selects the operation; another router adds indirection without a distinct policy. |
| Put filesystem traversal inside core services | Core would depend on infrastructure and lose pure-test boundaries. |
| Use one generic Repository interface | Unrelated capabilities would be coupled and future adapters would implement methods they do not own. |
| Let outbound adapters call use cases | Dependency direction would become cyclic or bidirectional. |
| Preserve the retired W009 internal package structure | The retired structure is not architecture authority for the rebuild. |

## Consequences

- Each of `list_records`, `get_records`, `resolve_reference`, and `validate_records` requires a dedicated application use case.
- Authoring-guidance and authoring-transaction architecture remains outside this ADR.
- MCP SDK and protocol types remain adapter-local.
- Application ports must preserve source identity, path-diagnostic inputs, and freshness semantics explicitly.
- Filesystem and future database adapters may differ internally while satisfying the same narrow capabilities.
- Cross-layer shortcuts are prohibited even when they reduce local code size.
- Concrete package placement follows a separate ADR so package names can change without reopening this logical architecture.

Affected Specification targets:

- `spec:drmcp.design_records_mcp.overview`;
- `spec:drmcp.design_records_mcp.responsibility_boundary`;
- `spec:drmcp.design_records_mcp.namespace_scanning`;
- `spec:drmcp.design_records_mcp.tools.overview`.

## Evidence

- Source Requirement: `DRMCP-REQ-MCP-001`.
- Source Work Item: `DRMCP-WORK-MCP-011`.
- Accepted decisions: D-001, D-004, and D-007 in `DRMCP-TASK-MCP-011-01`.
- Runtime lifecycle dependency: `DRMCP-ADR-MCP-002`.
