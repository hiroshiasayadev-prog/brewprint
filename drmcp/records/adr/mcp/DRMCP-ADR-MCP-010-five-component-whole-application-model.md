# DRMCP-ADR-MCP-010: Five-component whole-application model

- **status**: accepted
- **date**: 2026-07-04
- **depends_on**:
  - DRMCP-ADR-MCP-001
- **supersedes**:
  - DRMCP-ADR-MCP-007
- **migrated_to_spec**: null

## Context

ADR-007 introduced a separate Guidance Domain because Guidance was treated as a distinct guide source.
T09 found that model conflicted with the accepted portable standards package authority.

The portable package contains ordinary current Spec records under `spec:design_records.*`.
The application can reuse the existing Record Domain, logical tree, active index, retrieval, resolution, and validation model.

A separate Guidance Domain would duplicate normal Spec semantics without an independent lifecycle or substitution boundary.

## Decision

DRMCP uses five application-level components:

| component | owned responsibility | excluded responsibility |
|---|---|---|
| Composition / Lifecycle | Configuration validation, dependency construction, wiring, startup, and shutdown. | Operation policy, record semantics, and source access. |
| MCP Inbound Adapter | MCP protocol mapping and Application Use Case invocation. | Application policy, domain semantics, and concrete source access. |
| Application Use Cases | Operation sequencing, request-scoped orchestration, source selection, fixed Guidance scope, response projection, and result policy. | Record semantics and concrete I/O. |
| Record Domain / Logical Tree | Current record models, immutable logical structures, identity, retrieval primitives, resolution, and validation behavior. | MCP transport, application sequencing, configuration, and I/O. |
| Infrastructure I/O Adapters | Concrete Current Records and Legacy Archive source access behind inward-owned source contracts. | Operation policy and domain-result classification. |

The portable standards package is a normal Current Records source associated with the reserved `design_records` app namespace.
It does not create another top-level component.

Guidance operations remain active Application Use Cases.
They reuse normal record-query orchestration and Record Domain semantics.
Guidance adds no independent domain, source capability, parser, logical tree, index, snapshot, cache, or lifecycle.

The five components are responsibility boundaries.
They do not require one package, module, type, or runtime object per component.

Do not create top-level components for parsers, resolvers, validators, snapshot builders, indexes, individual sources, or Guidance projection helpers.
A later split requires an independent lifecycle, owner, substitution boundary, or cross-component contract.

Deferred Authoring keeps the MCP-to-use-case extension seam only.
This ADR does not decide proposal state, write transactions, packages, interfaces, types, functions, methods, or concrete configuration serialization.

## Rationale

Portable package Specs follow the normal current Spec format.
Using the Record Domain avoids a second identity, parsing, logical-tree, and validation model.

Guidance-specific behavior is operation policy and response projection.
Those responsibilities fit Application Use Cases and do not justify a top-level domain component.

The five-component model preserves inward dependencies and keeps the architecture smaller without weakening portability.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Retain Guidance Domain over the portable package | The package contains ordinary Specs and needs no independent semantic model. |
| Keep a dedicated package or Guidance index | A second index would duplicate Current Records identity and lookup behavior. |
| Merge operation policy into Record Domain | Fixed Guidance scope and response projection are Application concerns. |
| Create one component for each source or parser | Those units lack independent application-level ownership or lifecycle. |

## Consequences

- ADR-007 becomes historical and superseded.
- Canonical architecture Specifications use five components.
- Portable standards participate in the normal Current Records model.
- Guidance operations remain separate public use cases but share internal record-query orchestration.
- Downstream module design may share parser, query, and projection helpers inside the accepted boundaries.
- Adding a separate Guidance or package component requires new architecture decision work.

Affected design areas:

- `spec:drmcp.application_architecture`;
- `spec:drmcp.application_architecture.application_boundary_and_components`;
- `spec:drmcp.application_architecture.dependency_and_responsibility`;
- `spec:drmcp.application_architecture.runtime_and_state`.

## Evidence

- Source Requirements: `DRMCP-REQ-MCP-003` and `DRMCP-REQ-MCP-005`.
- Source Work Item: `DRMCP-WORK-MCP-016`.
- Review findings: F-BLK-01 and F-MAJ-01 in `DRMCP-TASK-MCP-016-09`.
- Revised decisions: D-018 and D-020 in `DRMCP-TASK-MCP-016-12`.
- ADR routing authority: B-04 in `DRMCP-TASK-MCP-016-13`.
- Portable package authority: `DRMCP-ADR-MCP-001` and `PRODUCT-REQ-SPEC-003`.
