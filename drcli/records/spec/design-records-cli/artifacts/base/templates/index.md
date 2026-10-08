# Overview: Artifact Specification templates

- **id**: `spec:drcli.design_records_cli.artifacts.base.templates`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.artifacts.base`

## What this is

Indexes the shared artifact Specification templates and provides the template for an artifact-specific `index.md` Specification.

## Current contract

This Specification defines the DRCLI-local projection and declarations for this topic.
Where this Specification cites PRODUCT authority, the cited PRODUCT Specification remains authoritative for PRODUCT-owned semantics.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| H1-adjacent metadata template | Contract | `spec:drcli.design_records_cli.artifacts.base.templates.h1_adjacent_metadata` | Template for artifact-specific H1-adjacent metadata Specifications. |
| Identity and structure templates | Index | `spec:drcli.design_records_cli.artifacts.base.templates.identity_and_structure` | Templates for sequential and tree identity-and-structure Specifications. |
| Source template | Contract | `spec:drcli.design_records_cli.artifacts.base.templates.source` | Template for artifact-specific source-document Specifications. |

## Rules

- DRCLI-local declarations must preserve cited PRODUCT-owned semantics.
- A DRCLI-local declaration must not create an independent replacement for cited PRODUCT authority.

## Validation rules

- A DRCLI interpretation that conflicts with cited PRODUCT authority is nonconforming.
- Validation of PRODUCT-owned meaning uses the cited PRODUCT authority.

## Template

```markdown
# Index: <ARTIFACT_NAME> artifact

- **id**: `<SPEC_REF>`
- **status**: draft
- **date**: `<YYYY-MM-DD>`
- **parent**: `spec:drcli.design_records_cli.artifacts`

## What this is

Defines the DRCLI Specification set for `<ARTIFACT_NAME>` records.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Identity and structure | Contract | `<IDENTITY_AND_STRUCTURE_SPEC_REF>` | Defines the record structure and identity declaration. |
| H1-adjacent metadata | Contract | `<H1_ADJACENT_METADATA_SPEC_REF>` | Defines the artifact metadata fields and value formats. |
| Source | Contract | `<SOURCE_SPEC_REF>` | Defines the source-document shape and allowed headings. |
```
