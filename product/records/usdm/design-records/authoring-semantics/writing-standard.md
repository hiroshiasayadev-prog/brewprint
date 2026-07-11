# USDM requirement: Writing standard read and validation scope

- **id**: `usdm:product.design_records.authoring_semantics.writing_standard`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: `usdm:product.design_records.authoring_semantics`

## What this is

This record defines DRMCP read and static-validation requirements derived from the Product Design Records writing standard.

## Requirements: spec:product.design_records.authoring_standards.writing_standard

| id | requirement | notes |
|---|---|---|
| R001 | DRMCP static validation must enforce only writing-standard rules that can be determined from Markdown structure, H1-adjacent metadata, headings, tables, or fenced code blocks. | Semantic and prose-quality validation are out of scope. |
| R002 | DRMCP read operations must return design-record body structure without semantic interpretation, including H1, H1-adjacent metadata, heading hierarchy, tables, fenced code blocks, and plain body blocks. | Read-side structural requirement. |
