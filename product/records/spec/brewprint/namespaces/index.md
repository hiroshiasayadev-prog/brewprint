# Overview: Brewprint namespace profile

- **id**: `spec:product.brewprint.namespaces`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.brewprint`

## What this is

Records current Brewprint app namespace and domain namespace assignments.
This profile instantiates the generic namespace model without redefining it.

## Current contract

| profile area | owner |
|---|---|
| Current app namespace assignments | `spec:product.brewprint.namespaces.app_namespaces`. |
| Current domain namespace assignments | `spec:product.brewprint.namespaces.domain_catalog`. |
| Future candidates | Recorded only when labeled as candidates. |
| Legacy prefixes and effective historical attribution | `spec:product.brewprint.compatibility`. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Brewprint app namespaces | Reference | `spec:product.brewprint.namespaces.app_namespaces` | Current app namespace assignments and future candidate labels. |
| Brewprint domain catalog | Reference | `spec:product.brewprint.namespaces.domain_catalog` | Current domain assignments, legacy-effective domains, and cross-app activity labels. |

## Rules

- Label registry facts as Brewprint profile content.
- Distinguish active assignment, future candidate, legacy prefix, and cross-app activity.
- Do not present future candidates as active namespaces or domains.
- Do not copy DRMCP or BPDSL operational contracts into this profile.

## Boundary

| content | owner |
|---|---|
| App-independent namespace semantics | `spec:product.design_records.namespace_model`. |
| Current Brewprint app and domain assignments | This profile area. |
| V01 compatibility, historical attribution, and issued-ID retention | `spec:product.brewprint.compatibility`. |
| DRMCP operational behavior | DRMCP app-local specifications. |
| BPDSL internal behavior | BPDSL app-local specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.brewprint` | Parent Brewprint overview. |
| `spec:product.design_records.namespace_model` | Generic namespace model instantiated by this profile. |
| `spec:product.brewprint.compatibility` | Brewprint compatibility and legacy attribution owner. |
