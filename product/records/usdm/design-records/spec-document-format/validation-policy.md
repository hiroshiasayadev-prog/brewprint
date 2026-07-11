# USDM requirement: Spec validation policy

- **id**: `usdm:product.design_records.spec_document_format.validation_policy`
- **status**: draft
- **date**: 2026-07-09
- **kind**: `requirement`
- **parent**: `usdm:product.design_records.spec_document_format`

## What this is

This record defines implementation requirements for reporting Specification format validation diagnostics.

## Requirements: Validation policy
> source: spec:product.design_records.spec_format.validation_policy

| id | requirement | notes |
|---|---|---|
| R001 | Spec-format validation must report detectable spec-format violations as diagnostics. | Phase-specific severity is not fixed by this row. |
| R002 | Spec-format validation must treat a missing real ATX H1 as an error. | |
| R003 | Spec-format validation must treat multiple real ATX H1 headings as an error. | |
| R004 | Spec-format validation must treat an H1 kind outside the accepted Specification kind set as an error. | |
| R005 | Spec-format validation must treat an H1-adjacent `parent` value that violates the allowed parent grammar as an error. | |
| R006 | Spec-format validation must treat an unresolved child `ref` as an error. | |
| R007 | Spec-format validation must treat duplicate parent declarations for the same child Specification as an error. | |
