# Reference: Domain catalog

- **id**: `spec:product.concepts.namespace_model.domain_catalog`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `spec:product.concepts.namespace_model`

## What this is

Catalogs the canonical domain namespaces assigned to each app namespace and the existing artifact prefixes that fall outside the canonical catalog.

## Canonical domain namespaces

| app namespace | domain namespace | concern area |
|---|---|---|
| `DRMCP` | `MCP` | All tools, authoring, schema, and validation for Design Records MCP |
| `BPDSL` | `DATA` | Data models, type system, rendering |
| `BPDSL` | `RESOLVE` | Identity resolution, sub-node identity enforcement |
| `PRODUCT` | `NAMESPACE` | Namespace model, catalog, artifact ID grammar |
| `PRODUCT` | `GOVERNANCE` | Cross-app governance (future) |
| `PRODUCT` | `MIGRATION` | Migration policies (future) |

## Existing prefixes outside the canonical catalog

Prefixes used in existing artifacts that do not belong to a canonical domain namespace.

| prefix | existing artifacts | nature |
|---|---|---|
| `SELFHOST` | REQ-SELFHOST-\* / WORK-SELFHOST-\* | Cross-app verification activity. Not a domain of a specific app, but a dogfooding / verification effort applicable to any app. Planned to also apply to DRMCP etc. in the future |
