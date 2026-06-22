# Reference: Subdomain model

- **id**: `spec:product.concepts.namespace_model.subdomain_model`
- **status**: draft
- **date**: 2026-06-22
- **parent**: `spec:product.concepts.namespace_model`

## What this is

Defines the subdomain grouping model for organizing artifacts within a domain namespace, including representation, write-time advisory behavior, and the DRMCP example.

## Subdomain model

When the number of artifacts within a domain namespace grows large, **subdomains** provide grouping by concept area.

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

## DRMCP MCP domain example

Existing V01-REQ-MCP-001–032 remain PRODUCT-owned per V01-ADR-096, but as a logical projection of their effective attribution / mapping to the DRMCP MCP domain, the following subdomains are identified:

| subdomain | target concern area |
|---|---|
| `TOOLS` | Read/retrieval tools (list_records / get_record / get_records etc.) |
| `SCHEMA` | Data models, metadata schema, validation rules |
| `AUTHORING` | Authoring transactions (propose/accept/discard) and authoring guidance |
| `UI` | Record browser UI |
