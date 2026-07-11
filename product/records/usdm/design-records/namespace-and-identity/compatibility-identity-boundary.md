# USDM requirement: Compatibility identity boundary

- **id**: `usdm:product.design_records.namespace_and_identity.compatibility_identity_boundary`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.namespace_and_identity`

## What this is

Defines the boundary between current canonical identity and Brewprint compatibility inputs.

This record includes compatibility-input ownership, old spec ref non-retention, and inactive alias/redirect defaults.
This record does not include current record kind definitions, current canonical reference forms, Brewprint profile facts, or authoring-time namespace selection.

## Requirements: spec:product.design_records.namespace_model.artifact_id_grammar

| id | requirement | notes |
|---|---|---|
| R003 | Implementations MUST route existing issued ADR ID compatibility through Brewprint legacy ID compatibility specs. | Brewprint compatibility mappings are compatibility inputs, not current canonical reference forms. |

## Requirements: spec:product.design_records.spec_format.spec_id_as_ref

| id | requirement | notes |
|---|---|---|
| R007 | Implementations MUST NOT preserve old spec refs by default after spec move or rename. | Explicit compatibility design is required before an old ref remains accepted. |
| R009 | Implementations MUST treat derived topic refs as prohibited under the current parent grammar. |  |
| R010 | Implementations MUST handle legacy `semantic_refs` and `sections` front matter through validation-policy severity rules, not through a parsing shim. | This applies to legacy material only. |

## Requirements: spec:product.design_records.traceability.semantic_ref

| id | requirement | notes |
|---|---|---|
| R014 | Implementations MUST NOT define aliases, redirects, stale-ref compatibility, split/merge compatibility, or superseded-ref tables in the current semantic-ref contract. | These concerns require explicit follow-up design. |

## Requirements: spec:product.design_records.traceability.artifact_refs

| id | requirement | notes |
|---|---|---|
| R015 | Implementations MUST treat complete legacy public IDs as compatibility inputs, not current canonical reference forms. |  |
| R016 | Implementations MUST preserve legacy issued-ID compatibility through Brewprint compatibility specs. |  |
