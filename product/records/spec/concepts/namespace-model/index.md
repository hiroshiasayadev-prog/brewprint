# Overview: Namespace model

- **id**: `spec:product.concepts.namespace_model`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `root`

## What this is

Defines the app namespace and domain namespace model for brewprint, providing the foundation for v2 artifact ID grammar, namespace catalog, and future namespace-first layout migration.

## Current contract

Three app namespaces are active: `DRMCP`, `BPDSL`, `PRODUCT`. Each app namespace owns one or more domain namespaces. See `## Topics` child specs for domain catalog, subdomain model, v2 ID grammar, and existing artifact ownership.

## Boundary

What this spec owns:

- Definitions of app namespaces and domain namespaces
- Architecture overview and domain namespace assignments for each app namespace
- Namespace attribution policy for existing artifacts
- Subdomain grouping model within domain namespaces (key-value label format, dynamic catalog, write-time advisory)

What this spec does not own:

- Machine-readable namespace registry formal schema, file format, or conversion implementation to physical placement
- Machine-readable namespace registry formal schema, file format, or physical placement
- Implementation specifications for namespace-aware MCP APIs
- Repository directory concerns, requiredness, and placement contracts (see `spec:product.concepts.repository_layout`)
- Current Brewprint repository inventory (see `spec:product.brewprint.layout`)

## Current placement and future layout

This spec is a human-readable concept definition located under `records/spec/concepts/`, distinct from the future machine-readable namespace registry.

Intended future operational placement:

- Near future: place formal schema and namespace definition files directly under the repo root, allowing MCP to mechanically read and resolve app/domain
- Future: place namespace declaration files directly under each app folder (app-local namespace declaration)

Defining this formal schema and placement is outside the scope of this spec and is treated as a separate requirement/specification.

The normative repository layout model, including `records/`, `dsl/`, `src/`, and design-record placement, is defined in `spec:product.concepts.repository_layout`.

The currently observed Brewprint repository tree is recorded in `spec:product.brewprint.layout`.

## App namespace and domain namespace

An **app namespace** identifies an application, subsystem, or cross-application product scope.

A **domain namespace** identifies a concern area within an app namespace.

| axis | example | role |
|---|---|---|
| app namespace | `DRMCP` / `BPDSL` / `PRODUCT` | Identifies owning application or product scope |
| domain namespace | `MCP` / `DATA` / `RESOLVE` | Identifies concern area within the app |

Current artifact IDs (`REQ-MCP-*` / `WORK-DATA-*`) are domain-first IDs from the era before app namespaces existed. Their attribution policy is described in `spec:product.concepts.namespace_model.existing_artifacts`.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| App namespaces | Reference | `spec:product.concepts.namespace_model.app_namespaces` | The three app namespaces (DRMCP, BPDSL, PRODUCT), their architecture overview, and domain namespace assignments. |
| Domain catalog | Reference | `spec:product.concepts.namespace_model.domain_catalog` | Canonical domain namespace catalog and existing artifact prefixes outside the catalog. |
| Subdomain model | Reference | `spec:product.concepts.namespace_model.subdomain_model` | Subdomain grouping model, representation, write-time advisory, and example. |
| v2 grammar | Reference | `spec:product.concepts.namespace_model.v2_grammar` | v2 artifact ID grammar, sequence format, and mapping rule from existing IDs. |
| Existing artifacts | Reference | `spec:product.concepts.namespace_model.existing_artifacts` | Namespace ownership of existing artifacts under V01-ADR-096. |
| v1 namespace resolution algorithm | Reference | `spec:product.concepts.namespace_model.v1_namespace_algorithm` | namespace_prefix derivation formula, kind-level prefix table, and multi-root scan behavior. |
| v1 record ID grammar | Reference | `spec:product.concepts.namespace_model.v1_id_grammar` | Public ID and bare ID grammar for v1 record kinds. |

## Sources

- V01-ADR-095: Coupling boundary between YAML DSL and Design Records MCP
- V01-ADR-096: PRODUCT namespace ownership of existing artifacts and non-execution of per-app migration
- V01-REQ-PRODUCT-001: App and domain namespace model for namespace-first design records
