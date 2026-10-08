# Overview: USDM requirement artifacts

- **id**: `spec:product.design_records.usdm`
- **status**: draft
- **date**: 2026-09-30
- **parent**: `spec:product.design_records`

## What this is

This spec area defines MVP USDM requirement artifacts for Brewprint Design Records.

USDM records normalize Specification requirements or direct upstream requirements into topic-scoped rows. Rows may decompose hierarchically through stable row IDs, but one decomposition tree remains inside one USDM requirement record. Implementation Specifications may directly cover rows owned by another app namespace through full-ID `usdm_covers` references, while coverage checks may also derive parent coverage from complete child coverage.

The MVP treats USDM as an independent auxiliary artifact. Full artifact-model integration, repository-layout integration, migration, and DRMCP integrated read support are deferred.

## Current contract

USDM records live under each app namespace at `<app>/records/usdm/` during the MVP.

USDM record IDs use `usdm:<app_namespace>.<path.to.topic>`.

USDM requirement row IDs use `RNNN(-NN)*` within one USDM requirement record. The full USDM requirement ID is `usdm:<app_namespace>.<path.to.topic>#<row-local requirement ID>`.

Implementation Specifications may declare H1-adjacent `usdm_covers` metadata for direct row coverage. Coverage checks also derive non-leaf coverage recursively from complete direct-child coverage and keep uncovered descendants beneath a directly covered ancestor as non-blocking warnings.

Standalone repository tools validate USDM records and coverage state under `tools/usdm/`. These tools may later move behind DRMCP or another MCP surface.

## Non-goals

- Do not replace Product Specifications with USDM records.
- Do not make USDM a fully integrated canonical artifact kind in the MVP.
- Do not require section-level coverage in the MVP.
- Do not migrate existing records into USDM in the MVP.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| USDM artifact format | Contract | `spec:product.design_records.usdm.artifact_format` | MVP USDM placement, record kinds, ID grammar, metadata, section shape, requirement row IDs, and validation rules. |
| USDM coverage format | Contract | `spec:product.design_records.usdm.coverage_format` | `usdm_covers` metadata, direct and derived coverage, blocking uncovered requirements, refinement warnings, and dangling cover detection. |
| USDM coverage tools | Contract | `spec:product.design_records.usdm.coverage_tools` | Standalone MVP tool contracts for `validate_usdm`, `check_usdm_coverage`, `usdm_covered_by`, and scoped coverage reporting. |
| USDM requirement similarity collection | Contract | `spec:product.design_records.usdm.similarity_collection` | USDM-facing candidate collection operation for semantically similar requirement details. |

## Boundary

USDM owns normalized implementation requirements derived from Specification authorities or recorded directly as literal upstream requirements when no corresponding Specification exists.

USDM does not own design decisions, component architecture, implementation contracts, or public operation responses.

Coverage metadata records direct Specification-to-row claims. Effective coverage may also be derived from complete child coverage. Neither direct nor derived coverage proves implementation correctness.

## Related specs

| ref | relation |
|---|---|
| PRODUCT-REQ-SPEC-015 | Source requirement for MVP USDM artifacts and coverage checks. |
| PRODUCT-WORK-SPEC-029 | Source Work Item for MVP USDM Specification and tooling. |
| PRODUCT-REQ-SPEC-016 | Source requirement for hierarchical decomposition and derived coverage. |
| PRODUCT-ADR-SPEC-020 | Decision for hierarchical row identity and effective coverage. |
| PRODUCT-WORK-SPEC-031 | Design Work Item for the hierarchical USDM extension. |
| `spec:product.design_records.spec_format.document_shape` | Defines spec file shape rules used by this spec area. |
