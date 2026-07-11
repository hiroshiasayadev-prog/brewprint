# USDM requirement: Spec metadata shape

- **id**: `usdm:product.design_records.spec_document_format.metadata_shape`
- **status**: draft
- **date**: 2026-07-09
- **kind**: `requirement`
- **parent**: `usdm:product.design_records.spec_document_format`

## What this is

This record defines implementation requirements for visible H1-adjacent metadata in Specification records.

## Requirements: spec:product.design_records.spec_format.document_shape

| id | requirement | notes |
|---|---|---|
| R001 | A Specification file must place the visible H1-adjacent metadata block immediately after H1 and before `## What this is`. | |
| R002 | Every Specification file must contain H1-adjacent `id`, `status`, `date`, and `parent` metadata markers. | |
| R003 | A `Contract` Specification must contain H1-adjacent `contract_class` metadata. | |
| R004 | A non-`Contract` Specification must not contain H1-adjacent `contract_class` metadata. | |
| R005 | A `contract_class` value must be `interface` or `format`. | |
| R006 | A new or migrated Specification must not contain YAML front matter. | Existing unmigrated Specifications are handled by validation policy. |
