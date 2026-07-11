# USDM requirement: Spec document shape

- **id**: `usdm:product.design_records.spec_document_format.document_shape`
- **status**: draft
- **date**: 2026-07-10
- **kind**: `requirement`
- **parent**: `usdm:product.design_records.spec_document_format`

## What this is

This record defines implementation requirements for the visible document shape of Specification records.

## Requirements: Spec document shape
> source: spec:product.design_records.spec_format.document_shape

| id | requirement | notes |
|---|---|---|
| R001 | A new or migrated Specification file must contain exactly one real ATX H1 outside fenced code blocks. | YAML front matter prohibition is covered by metadata shape. |
| R002 | A Specification H1 must match `# <SpecKind>: <Title>`. | |
| R003 | A Specification H1 kind must be one of `Overview`, `Index`, `Concept`, `Reference`, or `Contract`. | |
| R004 | A new or migrated Specification must reject deferred H1 kinds until this Specification or a successor explicitly accepts them. | Deferred kinds are `Guide`, `Process`, `Architecture`, and `Glossary`. |
| R005 | Every Specification kind must contain a `## What this is` section. | |
| R006 | An `Overview` Specification must contain a `## Current contract` section. | |
| R007 | A `Concept` Specification must contain a `## Concept model` section. | |
| R008 | A `Reference` Specification must contain at least one body H2 section with a Markdown table. | |
| R009 | An `interface` class `Contract` Specification must contain `## Request`, `## Response`, and `## Errors` sections. | |
| R010 | A `format` class `Contract` Specification must contain `## Current contract`, `## Rules`, and `## Validation rules` sections. | |
| R011 | A Specification file must not contain sections prohibited for its H1 kind or contract class. | |
