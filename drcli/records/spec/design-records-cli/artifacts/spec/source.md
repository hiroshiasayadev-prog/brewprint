# Contract: Specification source

- **id**: `spec:drcli.design_records_cli.artifacts.spec.source`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.artifacts.spec`
- **contract_class**: `format`
- **usdm_covers**:
  - usdm:product.design_records.spec_document_format.document_shape#R001,#R003-R011
  - usdm:product.design_records.spec_document_format.metadata_shape#R006

## What this is

Defines the source-document shape for Specification records.

## Current contract

This Specification defines the DRCLI-local projection and declarations for this topic.
Where this Specification cites PRODUCT authority, the cited PRODUCT Specification remains authoritative for PRODUCT-owned semantics.

## Rules

- DRCLI-local declarations must preserve cited PRODUCT-owned semantics.
- A DRCLI-local declaration must not create an independent replacement for cited PRODUCT authority.

## Validation rules

- A DRCLI interpretation that conflicts with cited PRODUCT authority is nonconforming.
- Validation of PRODUCT-owned meaning uses the cited PRODUCT authority.

## H1 prefix

- **type**: `enum`

### Values

| value |
|---|
| `Overview` |
| `Index` |
| `Concept` |
| `Reference` |
| `Contract` |

A Specification source contains exactly one real ATX H1 outside fenced code blocks.
The H1-adjacent metadata block uses the fields defined by `spec:drcli.design_records_cli.artifacts.spec.h1_adjacent_metadata`.
A new or migrated Specification source must not contain YAML front matter.

## H2 heading policy

- **unlisted headings**: `allowed`

## H2 headings

| heading | condition | format reference |
|---|---|---|
| `## What this is` | `always` | `spec:product.design_records.authoring_standards.spec_authoring` |
| `## Current contract` | Overview: always;<br/>Index: prohibited;<br/>Concept: optional;<br/>Reference: optional;<br/>Contract/interface: optional;<br/>Contract/format: always; | `-` |
| `## Non-goals` | Overview: recommended;<br/>Index: prohibited;<br/>Concept: recommended;<br/>Reference: optional;<br/>Contract/interface: recommended;<br/>Contract/format: recommended; | `-` |
| `## Topic map` | Overview: recommended;<br/>Index: prohibited;<br/>Concept: optional;<br/>Reference: optional;<br/>Contract/interface: optional;<br/>Contract/format: optional; | `-` |
| `## Topics` | Overview: optional;<br/>Index: always;<br/>Concept: prohibited;<br/>Reference: prohibited;<br/>Contract/interface: prohibited;<br/>Contract/format: prohibited; | `spec:product.design_records.spec_format.topics_table` |
| `## Concept model` | Overview: optional;<br/>Index: prohibited;<br/>Concept: always;<br/>Reference: prohibited;<br/>Contract/interface: optional;<br/>Contract/format: optional; | `-` |
| `## Rules` | Overview: optional;<br/>Index: prohibited;<br/>Concept: recommended;<br/>Reference: optional;<br/>Contract/interface: optional;<br/>Contract/format: always; | `-` |
| `## Boundary` | Overview: optional;<br/>Index: prohibited;<br/>Concept: recommended;<br/>Reference: optional;<br/>Contract/interface: recommended;<br/>Contract/format: recommended; | `-` |
| `## Request` | Overview: prohibited;<br/>Index: prohibited;<br/>Concept: prohibited;<br/>Reference: prohibited;<br/>Contract/interface: always;<br/>Contract/format: prohibited; | `-` |
| `## Response` | Overview: prohibited;<br/>Index: prohibited;<br/>Concept: prohibited;<br/>Reference: prohibited;<br/>Contract/interface: always;<br/>Contract/format: prohibited; | `-` |
| `## Errors` | Overview: optional;<br/>Index: prohibited;<br/>Concept: prohibited;<br/>Reference: optional;<br/>Contract/interface: always;<br/>Contract/format: optional; | `-` |
| `## Validation rules` | Overview: optional;<br/>Index: prohibited;<br/>Concept: optional;<br/>Reference: optional;<br/>Contract/interface: optional;<br/>Contract/format: always; | `-` |
| `## Related specs` | `recommended` | `-` |

A `Reference` Specification contains at least one body H2 section with a Markdown table.
A `Concept` or `Contract` Specification may contain a body H2 section with a Markdown table.
An `Overview` or `Index` Specification does not contain a body H2 section with a Markdown table.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.source` | Shared source-document rules. |
| `spec:drcli.design_records_cli.artifacts.spec.identity_and_structure` | Specification identity and tree structure. |
| `spec:drcli.design_records_cli.artifacts.spec.h1_adjacent_metadata` | Specification H1-adjacent metadata fields. |
| `spec:product.design_records.spec_format.document_shape` | Product authority for Specification kinds, source shape, and kind-specific section conditions. |
| `spec:product.design_records.spec_format.topics_table` | Product authority for the `## Topics` body format. |
| `spec:product.design_records.authoring_standards.spec_authoring` | Product authority for Specification authoring meaning. |
