# Contract: Artifact Specification set index template

- **id**: `spec:drmcp.design_records_mcp.artifacts.base.templates`
- **status**: draft
- **date**: 2026-07-10
- **parent**: `spec:drmcp.design_records_mcp.artifacts.base`
- **contract_class**: `format`

## What this is

Provides the template for an artifact-specific `index.md` Specification.

## Template

```markdown
# Reference: <ARTIFACT_NAME> artifact

- **id**: `<SPEC_REF>`
- **status**: draft
- **date**: `<YYYY-MM-DD>`
- **parent**: `spec:drmcp.design_records_mcp.artifacts`

## What this is

Defines the DRMCP Specification set for `<ARTIFACT_NAME>` records.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Identity and structure | Contract | `<IDENTITY_AND_STRUCTURE_SPEC_REF>` | Defines the record structure and identity declaration. |
| H1-adjacent metadata | Contract | `<H1_ADJACENT_METADATA_SPEC_REF>` | Defines the artifact metadata fields and value formats. |
| Source | Contract | `<SOURCE_SPEC_REF>` | Defines the source-document shape and allowed headings. |
```
