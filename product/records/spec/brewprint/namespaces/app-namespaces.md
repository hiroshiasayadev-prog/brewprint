# Reference: Brewprint app namespaces

- **id**: `spec:product.brewprint.namespaces.app_namespaces`
- **status**: draft
- **date**: 2026-07-02
- **parent**: `spec:product.brewprint.namespaces`

## What this is

Records current Brewprint app namespace assignments as profile facts.

## Current assignments

| app namespace | formal name | profile status | nature |
|---|---|---|---|
| `DRMCP` | Design Records MCP | Active app namespace. | Design record management MCP server. |
| `BPDSL` | Brewprint YAML DSL | Active app namespace. | DSL for describing Brewprint designs. |
| `PRODUCT` | Product-wide / cross-app | Active product namespace. | Cross-app policy, governance, migration, and PRODUCT specification work. |
| `TRV` | Task Responsibility Validator | Active app namespace. | Semantic Task responsibility validation application. |
| `DRUI` | Design Records UI | Future candidate. | Not confirmed as an active app namespace. |

## Profile notes

| app namespace | profile note |
|---|---|
| `DRMCP` | Current app-aware DRMCP records exist under the `MCP` and `SPEC` domains. DRMCP architecture and tool behavior are not defined here. |
| `BPDSL` | BPDSL is an active app namespace in the Brewprint profile. Canonical BPDSL type, resolver, render, and self-hosting behavior are not defined here. |
| `PRODUCT` | PRODUCT is a product-level namespace rather than a runtime application. Current app-aware PRODUCT records exist under the `NAMESPACE` and `SPEC` domains. |
| `TRV` | TRV owns app-local design and later implementation. PRODUCT retains the cross-app semantic validator contract. |
| `DRUI` | DRUI remains a candidate only. It must not be listed as an active namespace until accepted by a later decision. |

## Boundary

| content | treatment |
|---|---|
| DRMCP tool API, authoring transaction, storage, UI, parser, or validation behavior | DRMCP app-local specifications. |
| BPDSL type system, identity resolution, render pipeline, YAML parser, or self-hosting behavior | BPDSL app-local specifications. |
| TRV interface, runtime, model, provider, packaging, deployment, or implementation behavior | TRV app-local specifications. |
| Current app namespace assignment labels | Profile facts in this spec. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.brewprint.namespaces` | Parent namespace profile overview. |
| `spec:product.design_records.namespace_model.app_namespaces` | Generic app namespace semantics. |
