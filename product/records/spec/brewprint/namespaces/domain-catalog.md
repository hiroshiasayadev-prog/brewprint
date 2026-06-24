# Reference: Brewprint domain catalog

- **id**: `spec:product.brewprint.namespaces.domain_catalog`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.brewprint.namespaces`

## What this is

Records Brewprint domain namespace assignments as profile and registry content.

## Current assignments

| app namespace | domain namespace | status | concern area | evidence |
|---|---|---|---|---|
| `DRMCP` | `MCP` | Active assignment. | MCP tool contract and namespace-aware authoring work. | Current IDs include `DRMCP-REQ-MCP-001`, `DRMCP-REQ-MCP-002`, and `DRMCP-INV-MCP-001`. |
| `DRMCP` | `SPEC` | Active assignment. | DRMCP-owned spec-format implementation work. | Current IDs include `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002`. |
| `PRODUCT` | `NAMESPACE` | Active assignment. | Namespace-model cleanup and migration work. | Current IDs include `PRODUCT-WORK-NAMESPACE-001` and `PRODUCT-TASK-NAMESPACE-001-01`. |
| `PRODUCT` | `SPEC` | Active assignment. | PRODUCT specification format and semantic restructuring work. | Current IDs include `PRODUCT-REQ-SPEC-001`, `PRODUCT-WORK-SPEC-012`, and `PRODUCT-TASK-SPEC-012-03`. |

## Legacy-effective assignments

| app namespace | domain namespace | status | compatibility source |
|---|---|---|---|
| `BPDSL` | `DATA` | Legacy-effective assignment. | V01 `DATA` records have effective attribution to BPDSL compatibility history. |
| `BPDSL` | `RESOLVE` | Legacy-effective assignment. | V01 `RESOLVE` records have effective attribution to BPDSL compatibility history. |

Legacy-effective assignment records historical attribution.
It is not evidence that app-aware BPDSL `DATA` or `RESOLVE` records currently exist.

## Future candidates

| app namespace | domain namespace | status | note |
|---|---|---|---|
| `PRODUCT` | `GOVERNANCE` | Future candidate. | No scoped current app-aware record ID was found for this domain during T03 verification. |
| `PRODUCT` | `MIGRATION` | Future candidate. | No scoped current app-aware record ID was found for this domain during T03 verification. |

## Cross-app and legacy prefixes

| prefix | status | treatment |
|---|---|---|
| `SELFHOST` | Cross-app legacy activity. | Compatibility and effective attribution material belong under `spec:product.brewprint.compatibility`. |
| `PRODUCT` in V01 IDs | Legacy prefix. | Compatibility and effective attribution material belong under `spec:product.brewprint.compatibility`. |

## Objective contradiction recorded

The old generic domain catalog listed `PRODUCT` / `GOVERNANCE` and `PRODUCT` / `MIGRATION` as canonical domains.
Scoped current record IDs show active `PRODUCT` / `SPEC`, `PRODUCT` / `NAMESPACE`, and `DRMCP` / `SPEC` usage.
This profile corrects the Brewprint registry facts without changing generic namespace semantics.

## Related specs

| ref | relation |
|---|---|
| `spec:product.brewprint.namespaces` | Parent namespace profile overview. |
| `spec:product.design_records.namespace_model` | Generic namespace model. |
| `spec:product.brewprint.compatibility.existing_artifacts` | Legacy effective attribution and cross-app activity history. |
