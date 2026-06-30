# DRMCP-ADR-MCP-006: Read-runtime Go package boundaries

- **status**: accepted
- **date**: 2026-06-30
- **depends_on**:
  - DRMCP-ADR-MCP-002
  - DRMCP-ADR-MCP-003
  - DRMCP-ADR-MCP-004
  - DRMCP-ADR-MCP-005
- **supersedes**: []
- **migrated_to_spec**: 2026-06-30

## Context

The logical architecture requires concrete Go package boundaries before implementation Task authoring can begin.

The retired implementation concentrated most read logic in `internal/designrecords` and transport logic in `internal/designrecordsmcp`.
That structure is not authority for the rebuild.

Package placement must preserve the accepted inward dependency direction without splitting every type into a separate package.

## Decision

The following target structure is the complete W011 package inventory for `list_records`, `get_records`, `resolve_reference`, and `validate_records`.
It is not a complete package inventory for the Design Records MCP server.

W011 does not decide package placement for authoring-guidance or authoring-transaction operations.

The four-operation read-runtime slice uses this target structure:

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

Package responsibilities:

| package area | responsibility |
|---|---|
| `cmd/design-records-mcp` | Startup, construction, wiring, run, and shutdown only. |
| `internal/app/snapshot` | Shared request-snapshot orchestration and application-owned source-access ports. |
| `internal/app/listrecords` | `list_records` input, output, use case, and focused tests. |
| `internal/app/getrecords` | `get_records` input, output, use case, and focused tests. |
| `internal/app/resolvereference` | `resolve_reference` input, output, use case, and focused tests. |
| `internal/app/validaterecords` | `validate_records` input, output, use case, and focused tests. |
| `internal/core/records` | Source models, parsed models, and parsers. |
| `internal/core/indexing` | Distinct current index and legacy lookup types plus index construction. |
| `internal/core/resolving` | Transport-neutral resolver logic. |
| `internal/core/validation` | Per-source and graph validators plus finding values. |
| `internal/adapters/mcp` | MCP schemas, protocol wrappers, decode, invocation, and encode for the four W011 read-runtime operations. Other tool mappings remain outside this package-placement decision. |
| `internal/adapters/filesystem` | Source enumeration and reading implementation. |
| `internal/adapters/config` | Concrete configuration loading and validation adapter. |

Parser and source models remain in `core/records` until implementation evidence justifies a separate parsing package.
Current and legacy indexes remain distinct types inside `core/indexing`.

No compatibility wrapper preserves the retired `internal/designrecords` API.
Package cycles are prohibited.
Imports follow the logical dependency direction in `DRMCP-ADR-MCP-003`.

## Rationale

The package tree mirrors the accepted component and operation boundaries for the four W011 read-runtime operations.
One application package for each of those operations keeps inputs, outputs, sequencing, and tests focused.

A shared `app/snapshot` package avoids duplicating request-snapshot orchestration while keeping that orchestration outside core.

Core packages use semantic responsibilities rather than one package per type.
The structure remains cohesive without recreating one monolithic package.

Separate adapter packages keep MCP, filesystem, and configuration dependencies from spreading across the runtime.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Preserve `internal/designrecords` as the primary package | The retired monolith does not match the accepted application, core, and adapter boundaries. |
| Add a compatibility wrapper for the retired API | The wrapper would preserve unwanted coupling and create a second internal surface. |
| Put all use cases in one `app` package | Operation contracts and tests would become coupled through one broad package. |
| Create one package per model or parser type | The package graph would fragment without an independent responsibility boundary. |
| Put snapshot orchestration in core | Core would own application sequencing and source-access ports. |
| Combine MCP, filesystem, and config adapters | Unrelated infrastructure dependencies would share one package boundary. |

## Consequences

- Implementation Tasks for the four W011 read-runtime operations must respect the package ownership table.
- Authoring-guidance and authoring-transaction package placement remains undecided by this ADR.
- Cross-operation sharing occurs through explicit application or core packages, not direct imports between operation packages.
- Old package API compatibility is intentionally absent.
- Tests can target operation packages, pure core packages, and adapters separately.
- A later package split requires evidence that one accepted responsibility has become independently changeable.
- Package placement changes that preserve the logical architecture may amend or supersede this ADR without reopening public tool semantics.

Affected Specification targets:

- `spec:drmcp.design_records_mcp.overview`;
- `spec:drmcp.design_records_mcp.responsibility_boundary`;
- future implementation-boundary specifications created or selected during Phase C.

## Evidence

- Source Requirement: `DRMCP-REQ-MCP-001`.
- Source Work Item: `DRMCP-WORK-MCP-011`.
- Accepted decision: D-009 in `DRMCP-TASK-MCP-011-01`.
- Logical architecture authorities: `DRMCP-ADR-MCP-002` through `DRMCP-ADR-MCP-005`.
