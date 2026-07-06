# Overview: DRMCP implementation

- **id**: `spec:drmcp.implementation`
- **status**: draft
- **date**: 2026-06-30
- **parent**: `-`

## What this is

Defines the current implementation architecture for the W011 read-runtime slice: `list_records`, `get_records`, `resolve_reference`, and `validate_records`.

Public operation behavior remains owned by `spec:drmcp.design_records_mcp` and its child contracts.
Authoring-guidance and authoring-transaction use cases, snapshot lifecycle, and package placement are outside this Specification.

## Current contract

### Authority boundary

| concern | owner |
|---|---|
| Public requests, responses, statuses, warnings, diagnostics, and operation errors | `spec:drmcp.design_records_mcp` child contracts. |
| PRODUCT artifact semantics, identity, layout, and authoring rules | PRODUCT Design Records Specifications. |
| Runtime lifecycle, internal layers, ports, state boundaries, validation orchestration, and package placement for the four W011 read-runtime operations | This Specification. |
| Decision rationale and rejected alternatives | `DRMCP-ADR-MCP-002` through `DRMCP-ADR-MCP-006`. |

The implementation consumes public contracts.
It does not redefine them.

### Runtime lifecycle

The composition root loads and validates configuration at server startup.
Invalid startup configuration prevents server start.

Each invocation of `list_records`, `get_records`, `resolve_reference`, or `validate_records` builds one fresh immutable snapshot.
The snapshot contains:

- one current active index;
- one optional legacy exact lookup map;
- retained current source and parsed state required by the operation.

The current index and legacy lookup map remain separate.
No generic merged index may hide source-family behavior.

Filesystem changes become visible on the next invocation of one of those four operations.
This lifecycle does not apply this architecture to authoring-guidance or authoring-transaction operations.
The runtime does not retain a mutable process-wide record index.
The runtime does not incrementally patch an index after writes.

A post-start scan, read, or index failure fails the current operation.
The operation must not return a partial normal response from an untrustworthy snapshot.

### Component structure

The following structure applies to the W011 read-runtime slice for the four named operations.
It does not define the complete internal architecture of the Design Records MCP server.

```text
composition root
├─ inbound adapter: MCP
├─ application
│  ├─ operation use cases
│  └─ request-snapshot orchestration
├─ core
│  ├─ records and parsers
│  ├─ current index and legacy lookup
│  ├─ resolver
│  └─ validation
└─ outbound adapters
   ├─ filesystem
   └─ configuration
```

Dependency direction:

```text
MCP adapter ───────┐
filesystem adapter ├──> application ───> core
config adapter ────┘
composition root wires every component
```

Rules:

- The composition root owns startup, construction, wiring, run, and shutdown.
- The composition root owns no parsing, indexing, resolution, validation, or operation policy.
- MCP is the inbound adapter for the four W011 read-runtime operations.
- Application use cases for those operations own operation sequencing and snapshot orchestration.
- Core owns transport-neutral parsing, indexing, resolution, and validation logic.
- Outbound adapters implement application-owned ports.
- Core imports no application, MCP, filesystem, or configuration package.
- Outbound adapters do not invoke application use cases.

### Operation use cases and MCP adapter

Each of `list_records`, `get_records`, `resolve_reference`, and `validate_records` invokes one dedicated application use case.
No generic App Router is introduced for those operations.

Authoring-guidance and authoring-transaction operation architecture remains with their owning contracts.
This Specification does not assign their use cases, snapshots, or package placement.

Each W011 use case owns one typed input and one typed output.
Application types contain no MCP SDK or protocol wrapper types.

For the four W011 operations, the MCP adapter owns:

- protocol request and response schemas;
- decoding;
- structural JSON validation;
- use-case invocation;
- response encoding.

Other tool mappings are outside this architecture scope.

The MCP adapter maps between MCP schemas and application contracts.
The adapter does not redesign operation fields, statuses, warnings, diagnostics, or error meaning.

### Ports and outbound adapters

The four W011 application use cases depend on narrow ports for:

- configuration access;
- source enumeration;
- source reading.

Filesystem traversal and file reading belong to the filesystem adapter.
Concrete configuration loading belongs to the configuration adapter.

Parsers and index builders remain pure core services where practical.
No all-purpose repository interface combines configuration, source access, indexing, resolution, and validation.

A future database-backed adapter may implement the same narrow capabilities.
The adapter must preserve source identity, diagnostic location, and freshness semantics.

### Internal state separation

The runtime uses distinct internal types for:

| state | responsibility |
|---|---|
| Source state | Raw bytes, source identity, and provenance. |
| Parsed state | Parsed fields and document structure, including invalid present values. |
| Current index entry | One uniquely addressable current identity. |
| Current identity conflict | Every source claiming one identity, with no winner. |
| Invalid but addressable state | One canonical identity whose other content is invalid. |
| Legacy lookup state | Exact issued legacy ID to unique, absent, unreadable, or conflicted source state. |
| Validation finding | Transport-neutral semantic or structural finding data. |
| Operation projection | The public meaning returned by one application use case. |

No nullable record type spans parsing, indexing, validation, application, and MCP transport.

Missing parsed values remain missing.
Invalid present values remain available as invalid data.
Conflicts do not select a filesystem-order winner.

Core returns domain values, lookup states, conflicts, and validation findings.
Core does not construct complete tool responses.

### Result and execution-error boundary

Expected semantic states remain application result data when the public contract defines a normal response.
Examples include:

- unresolved accepted references;
- duplicate identity conflicts;
- disabled legacy lookup;
- validation findings;
- partial exact-retrieval success.

A failure is an application execution error when no trustworthy result can be constructed.
Examples include:

- an unbuildable mandatory snapshot;
- an incomplete required conflict collection;
- a required diagnostic location that cannot be represented.

The MCP adapter preserves this classification.

### Validation orchestration

Standalone `validate_records` uses one fresh request-scoped snapshot.
Source enumeration and reading finish before validation passes start.

Validation order:

1. Parse and retain source state.
2. Build current and configured legacy lookup state.
3. Run per-source syntax, metadata, identity, and document-shape detectors.
4. Run relation and Topics graph validation against complete index state.
5. Aggregate, suppress semantic duplicates, order, and project findings.

Individual validators:

- perform no filesystem I/O;
- do not rescan roots;
- do not call public MCP tools;
- do not invoke another validator through transport;
- return findings as data;
- do not format MCP responses.

Legacy sources remain lookup targets only.
They do not become current repository-validation subjects.

Any future integration that validates after persisted filesystem writes must discard candidate and pre-write state.
That integration must rebuild validation input from persisted files.

Current authoring-transaction integration is deferred to `DRMCP-REQ-MCP-002`.
W011 does not define authoring-transaction behavior.
YAML or V01-SPEC authoring semantics are not integrated with the current-format validation contract by this Specification.

### Go package layout

This tree is the complete package inventory for the four W011 read-runtime operations.
It is not the complete package inventory for the Design Records MCP server.
Authoring-guidance and authoring-transaction package placement are outside this Specification.

```text
drmcp/src/
├─ cmd/design-records-mcp/
│  └─ main.go
└─ internal/
   ├─ app/
   │  ├─ snapshot/
   │  ├─ listrecords/
   │  ├─ getrecords/
   │  ├─ resolvereference/
   │  └─ validaterecords/
   ├─ core/
   │  ├─ records/
   │  ├─ indexing/
   │  ├─ resolving/
   │  └─ validation/
   └─ adapters/
      ├─ mcp/
      ├─ filesystem/
      └─ config/
```

| package | responsibility |
|---|---|
| `cmd/design-records-mcp` | Composition root only. |
| `internal/app/snapshot` | Shared snapshot orchestration and application-owned source ports. |
| `internal/app/listrecords` | `list_records` application contract and use case. |
| `internal/app/getrecords` | `get_records` application contract and use case. |
| `internal/app/resolvereference` | `resolve_reference` application contract and use case. |
| `internal/app/validaterecords` | `validate_records` application contract and use case. |
| `internal/core/records` | Source models, parsed models, and parsers. |
| `internal/core/indexing` | Distinct current index and legacy lookup types. |
| `internal/core/resolving` | Resolver logic. |
| `internal/core/validation` | Per-source and graph validators plus finding values. |
| `internal/adapters/mcp` | MCP schemas and protocol mapping for the four W011 read-runtime operations. Other tool internals are not decided here. |
| `internal/adapters/filesystem` | Filesystem source enumeration and reading. |
| `internal/adapters/config` | Concrete configuration loading. |

Package cycles are prohibited.

No compatibility wrapper preserves the retired `internal/designrecords` API.
Parsers remain in `core/records` unless implementation evidence justifies a separate package.
Current and legacy index types remain distinct inside `core/indexing`.

## Non-goals

- Changing public read-operation behavior accepted by W003 through W006.
- Reusing the retired W009 production structure.
- Defining or integrating authoring transaction runtime behavior. That work belongs to `DRMCP-REQ-MCP-002`.
- Selecting database technology or another persistence implementation.
- Defining framework-specific MCP SDK code.
- Adding caches, watchers, incremental index mutation, or background refresh.
- Defining fixture placement.
- Planning implementation Tasks.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.overview` | Public Design Records MCP contract entry point. |
| `spec:drmcp.design_records_mcp.responsibility_boundary` | External and internal ownership boundary. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Configured roots and current or legacy index semantics. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Current source and addressability contract. |
| `spec:drmcp.design_records_mcp.schema.diagnostics` | Public warning and diagnostic representation. |
| `spec:drmcp.design_records_mcp.tools.overview` | Public tool catalog and operation contracts. |
| `spec:drmcp.design_records_mcp.tools.validate_records` | Public validation operation contract. |
| `spec:drmcp.implementation.contracts` | Architecture-derived module-contract baseline for the current whole-application model. |
| `DRMCP-ADR-MCP-001` | Current and legacy separation authority. |
| `DRMCP-ADR-MCP-002` | Snapshot and lifecycle rationale. |
| `DRMCP-ADR-MCP-003` | Layer and adapter rationale. |
| `DRMCP-ADR-MCP-004` | Internal state and operation-contract rationale. |
| `DRMCP-ADR-MCP-005` | Validation orchestration rationale. |
| `DRMCP-ADR-MCP-006` | Go package-layout rationale. |
