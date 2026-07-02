# Overview: Task Responsibility Validator

- **id**: `spec:trv`
- **status**: draft
- **date**: 2026-07-02
- **parent**: root

## What this is

Entry point for TRV-owned design records.
It records the application identity and its boundary with PRODUCT-owned validator semantics.

## Current contract

| area | current contract |
|---|---|
| App namespace | `TRV`. |
| Formal name | Task Responsibility Validator. |
| Records root | `trv/records/`. |
| PRODUCT-owned semantics | `spec:product.responsibility_boundary_validator`. |
| TRV-owned scope | App-local Requirement, ADR, Specification, and later implementation work. |
| Current design state | App-local design remains pending under `TRV-WORK-SPEC-001`. |

## Topic map

| topic | route |
|---|---|
| Cross-app semantic validator contract | `spec:product.responsibility_boundary_validator`. |
| TRV app-local design | `TRV-WORK-SPEC-001` and future child specs under `spec:trv.*`. |
| Future DRMCP integration | Separate Requirement or Work Item. |

## Non-goals

- Concrete interface, transport, runtime, model, provider, packaging, or deployment design.
- Current DRMCP integration.
- Implementation planning or implementation execution.
- Duplication of PRODUCT-owned semantic validation rules.

## Boundary

| content | owner |
|---|---|
| Cross-app semantic validator behavior and invocation policy | PRODUCT. |
| App-local design and later implementation | TRV. |
| Current DRMCP structural Design Record operations | DRMCP. |
| Future DRMCP integration | Separate Requirement or Work Item. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.responsibility_boundary_validator` | PRODUCT-owned semantic contract consumed by TRV. |
| `spec:product.brewprint.namespaces.app_namespaces` | Active TRV app namespace assignment. |
| `spec:product.brewprint.namespaces.domain_catalog` | Active TRV / SPEC domain assignment. |
| `spec:product.design_records.repository_layout` | App-independent placement contract for `trv/records/`. |
