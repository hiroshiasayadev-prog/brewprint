# Overview: Namespace model

- **id**: `spec:product.design_records.namespace_model`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records`

## What this is

Defines the app-independent Design Records namespace model.
The model provides vocabulary for app namespaces, domain namespaces, artifact ownership selection, artifact ID grammar, and subdomain grouping.

## Current contract

| concept | contract |
|---|---|
| App namespace | Identifies the application, subsystem, or cross-application scope that owns a record. |
| Domain namespace | Identifies a concern area within an app namespace. |
| Artifact ownership selection | New records use the owning app namespace when the owner is confirmed. Cross-app or unclear concerns use the product namespace. |
| Artifact ID grammar | Sequential workflow artifacts use the canonical app-aware ID grammar. |
| Subdomain grouping | A domain may use metadata subdomains for advisory grouping without changing public IDs. |

The current Brewprint app and domain assignments are profile facts.
They are recorded in `spec:product.brewprint.namespaces`.

Brewprint V01 compatibility and issued-ID retention are project compatibility facts.
They are recorded in `spec:product.brewprint.compatibility`.

## Boundary

| content | owner |
|---|---|
| App-independent app namespace and domain namespace semantics | This namespace-model area. |
| Artifact ID grammar relationship to namespaces | `spec:product.design_records.namespace_model.artifact_id_grammar`. |
| New-artifact ownership selection | `spec:product.design_records.namespace_model.existing_artifacts`. |
| Subdomain metadata representation and advisory behavior | `spec:product.design_records.namespace_model.subdomain_model`. |
| Current Brewprint app and domain assignments | `spec:product.brewprint.namespaces`. |
| Brewprint V01 compatibility and migration state | `spec:product.brewprint.compatibility`. |
| DRMCP parser, MCP, storage, UI, projection, or tool behavior | DRMCP app-local specifications. |
| Canonical BPDSL language, type, resolver, render, or runtime behavior | BPDSL app-local specifications. |
| Repository directory requiredness and Design Records placement contracts | `spec:product.design_records.repository_layout`. |
| Current Brewprint repository inventory | `spec:product.brewprint.layout`. |

## Current placement and future layout

This section records the current disposition of earlier registry-placement alternatives.
It does not define current Brewprint placement or current registry facts.

This spec is a human-readable concept definition located under `records/spec/design-records/namespace-model/`.
It is distinct from a future machine-readable namespace registry.

Earlier alternatives considered root-level formal registry files and app-local namespace declaration files.
No accepted requirement currently chooses either placement.
The alternatives are deleted from the current contract after evidence transfer to PRODUCT-INV-SPEC-005 and PRODUCT-ADR-SPEC-001.

The normative Design Records repository layout, including `records/` and design-record placement, is defined in `spec:product.design_records.repository_layout`. Historical `dsl/`, `src/`, and DSL-to-source implementation-flow material is preserved in `spec:product.bpdsl.repository_implementation_flow`.

The currently observed Brewprint repository tree is recorded in `spec:product.brewprint.layout`.

## App namespace and domain namespace

An **app namespace** identifies an application, subsystem, or cross-application product scope.

A **domain namespace** identifies a concern area within an app namespace.

| axis | example value | role |
|---|---|---|
| app namespace | `EXAMPLEAPP` | Identifies the owning application or product scope. |
| domain namespace | `AUTHORING` | Identifies the concern area within the app. |

Current Brewprint app names, domain names, and legacy prefixes are not generic namespace rules.
They are profile or compatibility facts under `spec:product.brewprint`.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| App namespaces | Reference | `spec:product.design_records.namespace_model.app_namespaces` | App namespace concept, domain relationship, and ownership selection semantics. |
| Subdomain model | Reference | `spec:product.design_records.namespace_model.subdomain_model` | Subdomain grouping model, representation, and write-time advisory semantics. |
| Artifact ID grammar | Reference | `spec:product.design_records.namespace_model.artifact_id_grammar` | Canonical artifact ID grammar, sequence format, allocation scopes, and canonical-ref rule. |
| New artifact ownership | Reference | `spec:product.design_records.namespace_model.existing_artifacts` | App-independent rule for selecting the namespace of new artifacts. |

## Sources

- V01-ADR-095: Coupling boundary between YAML DSL and Design Records MCP
- V01-ADR-096: PRODUCT namespace ownership of existing artifacts and non-execution of per-app migration
- V01-REQ-PRODUCT-001: App and domain namespace model for namespace-first design records

## Related specs

| ref | relation |
|---|---|
| `spec:product.brewprint.namespaces` | Current Brewprint app and domain namespace profile. |
| `spec:product.brewprint.compatibility` | Brewprint V01 compatibility and migration history. |
