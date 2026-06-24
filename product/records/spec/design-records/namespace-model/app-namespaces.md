# Reference: App namespaces

- **id**: `spec:product.design_records.namespace_model.app_namespaces`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records.namespace_model`

## What this is

Defines app namespace semantics for Design Records.
It does not list the current Brewprint registry.

## Current contract

An app namespace identifies the application, subsystem, or cross-application product scope that owns a record.

| rule | contract |
|---|---|
| Ownership axis | App namespace is the first ownership axis for app-aware design record IDs. |
| Domain relationship | Each app namespace may contain one or more domain namespaces. |
| Cross-app concerns | Cross-app or unclear concerns use the product-level namespace until ownership is confirmed. |
| Registry facts | Concrete app names and domain assignments are profile data, not generic rules. |

## Rules

- Do not encode application architecture in a generic app namespace definition.
- Do not infer current registry assignments from future candidate names.
- Do not treat a tool implementation as the owner of the generic namespace concept.
- Use a Brewprint profile spec to record current app namespace assignments.

## Boundary

| content | owner |
|---|---|
| App namespace concept | This spec. |
| Current Brewprint app namespace list | `spec:product.brewprint.namespaces.app_namespaces`. |
| Current Brewprint domain assignments | `spec:product.brewprint.namespaces.domain_catalog`. |
| DRMCP MCP server architecture and tool behavior | DRMCP app-local specifications. |
| BPDSL type, resolver, render, and self-hosting behavior | BPDSL app-local specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.namespace_model` | Parent namespace model overview. |
| `spec:product.brewprint.namespaces` | Current Brewprint namespace profile. |
