# Concept: Infrastructure I/O adapter boundary

- **id**: `spec:drmcp.implementation.contracts.infrastructure_io_adapters.contract_boundary`
- **status**: draft
- **date**: 2026-07-06
- **parent**: `spec:drmcp.implementation.contracts.infrastructure_io_adapters`

## What this is

Contract baseline for Infrastructure I/O Adapters as module-contract owners.

## Current contract

Infrastructure I/O Adapters own concrete source-family access behind inward-owned source contracts.
They do not own operation policy, Domain result classification, or MCP response encoding.

Contract-level components:

| component | responsibility |
|---|---|
| Current Records Source Access | Current source enumeration, reading, provenance, boundary enforcement, and access-failure reporting. |
| Legacy Archive Source Access | Legacy Archive enumeration, reading, provenance, boundary enforcement, and access-failure reporting. |

Enumeration and reading are internal operations or downstream port refinements.
They are not separate W018 contract-level components.

## Non-goals

- Application operation sequencing.
- Record identity, parser, resolver, or validation semantics.
- MCP protocol mapping.
- Concrete filesystem library choice.
- Database or persistence technology selection.

## Concept model

| surface | direction | meaning |
|---|---|---|
| Current source input or source-access result | Infrastructure to Application or parser-facing flow | Discovered source, raw Markdown, provenance, and access state. |
| Legacy source input or legacy-access result | Infrastructure to Legacy lookup-state build flow | Legacy source, issued-ID candidate, read or access state, and provenance. |

## Rules

| rule | contract |
|---|---|
| Source-family split | Current Records access and Legacy Archive access remain separate. |
| Inward-owned contracts | Infrastructure implements source contracts owned inward by Application and Domain boundaries. |
| Provenance preservation | Source results preserve identity and location material required by diagnostics. |
| Access failure reporting | Source access reports access failures as source-family states. |
| No semantic classification | Infrastructure does not classify Domain outcomes or operation responses. |
| No request-state observation | Infrastructure does not observe Current Records Snapshot or Legacy Lookup State internals. |
| No use-case invocation | Infrastructure does not invoke Application Use Cases. |

## Boundary

Forbidden bypasses:

- Infrastructure must not call public MCP tools.
- Infrastructure must not invoke use cases.
- Infrastructure must not select duplicate winners from source ordering.
- Infrastructure must not merge Current and Legacy access into one generic index.
- Infrastructure must not reuse stale request state after source failure.

Downstream detailed contract convergence must define exact source result fields, access-state values, port boundaries, and fixture expectations.
This module-contract baseline is not implementation-ready.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.implementation.contracts` | Module-contract root. |
| `spec:drmcp.application_architecture.dependency_and_responsibility` | Accepted dependency authority. |
| `DRMCP-ADR-MCP-013` | Source ADR. |
