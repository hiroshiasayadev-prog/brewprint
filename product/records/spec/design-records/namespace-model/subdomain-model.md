# Reference: Subdomain model

- **id**: `spec:product.design_records.namespace_model.subdomain_model`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records.namespace_model`

## What this is

Defines the subdomain grouping model for organizing artifacts within a domain namespace.

## Current contract

When the number of artifacts within a domain namespace grows large, **subdomains** provide grouping by concept area.
Subdomains do not change public artifact IDs.

## Definition and representation

Subdomains are not included in artifact IDs. They are expressed as a `subdomain` field in record metadata. For Design Records workflow artifacts (requirement / work_item / task), this is recorded in the bullet-list metadata immediately after H1.

```
- **subdomain**: AUTHORING
```

- **Index is flat per domain**: sequence numbers do not reset per subdomain
- **No predefined catalog**: valid values are derived dynamically from the `subdomain:` fields of existing records. The set of `subdomain` values present within a domain constitutes the catalog
- **Not enforced on all domains**: domains that do not need subdomains do not need to carry this field

## Write-time advisory

When a propose-type tool detects a new value in the `subdomain` field, it lists the existing subdomain values within the same domain for the author to review (no blocking, no similarity algorithm — the author makes the final call).

## Boundary

| content | owner |
|---|---|
| Subdomain concept and metadata representation | This spec. |
| Write-time advisory semantics | This spec. |
| Concrete subdomain catalog for an app domain | The owning app or Brewprint profile, depending on whether the catalog is current profile data or app-local operational detail. |
| DRMCP MCP-domain subdomain example and tool concern split | DRMCP app-local specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.namespace_model` | Parent namespace model overview. |
| `spec:product.brewprint.compatibility.existing_artifacts` | Legacy effective attribution context. |
