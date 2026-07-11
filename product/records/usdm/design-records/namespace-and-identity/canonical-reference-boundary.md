# USDM requirement: Canonical reference boundary

- **id**: `usdm:product.design_records.namespace_and_identity.canonical_reference_boundary`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.namespace_and_identity`

## What this is

Defines the current canonical reference boundary for Design Records.

This record includes current record kinds, canonical reference forms, physical path prohibition, and inactive semantic endpoint boundaries.
This record does not include parser/resolver implementation, diagnostic names, reverse relation derivation, provenance-cycle validation, or field-specific allowed target sets.

## Requirements: spec:product.design_records.traceability

| id | requirement | notes |
|---|---|---|
| R001 | Implementations MUST support the current record kinds and canonical reference forms defined by Product traceability specs. |  |
| R002 | Implementations MUST treat physical paths as repository locations, not canonical relation or reference values. |  |
| R003 | Implementations MUST keep PRODUCT-owned reference semantics separate from DRMCP-owned request, response, diagnostic, parser, persistence, indexing, UI, tool API, and writer behavior. |  |
| R004 | Implementations MUST NOT activate `yaml:`, `internal-design:`, `coverage:`, `fixture:`, or workflow semantic prefixes. |  |
| R005 | Implementations MUST NOT define `maps_to`, `covers`, or `validates` as active trace relations. |  |
| R006 | Implementations MUST NOT define canonical section refs before a visible-table contract exists. |  |

## Requirements: spec:product.design_records.traceability.artifact_refs

| id | requirement | notes |
|---|---|---|
| R011 | Implementations MUST support only the current record kinds and canonical reference forms defined in `artifact_refs` as current Product reference forms. | Field-specific allowed target sets are defined by metadata or authoring contracts. |
| R012 | Implementations MUST NOT infer relations from physical paths, file names, parent directories, or ID string structure. |  |
| R014 | Implementations MUST use complete Task public IDs for Work Item `tasks`. |  |
| R015 | Implementations MUST use complete Work Item public IDs for Task `work_item`. |  |
| R016 | Implementations MUST use complete Task public IDs for Task `depends_on`. |  |
| R017 | Implementations MUST NOT define Task `source_refs`. |  |
| R019 | Implementations MUST NOT reserve or adopt additional semantic prefixes or realization relations. |  |
