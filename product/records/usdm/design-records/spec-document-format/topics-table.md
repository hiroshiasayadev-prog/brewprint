# USDM requirement: Spec Topics table

- **id**: `usdm:product.design_records.spec_document_format.topics_table`
- **status**: draft
- **date**: 2026-07-10
- **kind**: `requirement`
- **parent**: `usdm:product.design_records.spec_document_format`

## What this is

This record defines implementation requirements for `## Topics` tables in Index Specifications and Overview Specifications that declare child topics.

## Requirements: Topics table
> source: spec:product.design_records.spec_format.topics_table

| id | requirement | notes |
|---|---|---|
| R001 | An `Index` Specification must contain a `## Topics` table. | |
| R002 | Validation must reject an `Overview` `## Topics` child declaration when the child file's directory placement, H1-adjacent `id`, and row `ref` do not identify the same Specification through the path-derived canonical mapping. | |
| R003 | A `## Topics` table must contain `title`, `kind`, `ref`, and `summary` columns. | |
| R004 | A `## Topics` row must use a canonical `spec:` ref as the child target. | |
| R005 | Tooling must not treat a `file` column as the canonical child target in a `## Topics` table. | |
| R006 | A child Specification must be declared as a child by a `## Topics` row in an `Index` Specification or an `Overview` Specification. | |
| R007 | A child Specification's H1-adjacent `parent` marker must match the `id` of the parent Specification that declares the child in a `## Topics` row. | |
| R008 | Validation must treat the same child `ref` declared by multiple parent Specifications' `## Topics` rows as invalid. | |
| R010 | A `## Topics` row `kind` value must be one of the accepted Specification kinds. | |
| R011 | A `## Topics` row `ref` must resolve to exactly one Specification file through the path-derived canonical mapping. | |
