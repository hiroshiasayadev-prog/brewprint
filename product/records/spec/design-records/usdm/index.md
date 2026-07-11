# Overview: USDM requirement artifacts

- **id**: `spec:product.design_records.usdm`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:product.design_records`

## What this is

This spec area defines MVP USDM requirement artifacts for Brewprint Design Records.

USDM records normalize Specification requirements or direct upstream requirements into topic-scoped rows. Implementation Specifications use coverage metadata to declare which USDM requirement rows they cover.

The MVP treats USDM as an independent auxiliary artifact. Full artifact-model integration, repository-layout integration, migration, and DRMCP integrated read support are deferred.

## Current contract

USDM records live under each app namespace at `<app>/records/usdm/` during the MVP.

USDM record IDs use `usdm:<app_namespace>.<path.to.topic>`.

USDM requirement row IDs use `RNNN` within one USDM requirement record. The full USDM requirement ID is `usdm:<app_namespace>.<path.to.topic>#RNNN`.

Implementation Specifications may declare H1-adjacent `usdm_covers` metadata to list covered USDM requirement IDs or compact same-record row lists.

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
| USDM coverage format | Contract | `spec:product.design_records.usdm.coverage_format` | `usdm_covers` metadata, compact coverage row-list syntax, coverage relation semantics, uncovered requirement detection, and dangling cover detection. |
| USDM coverage tools | Contract | `spec:product.design_records.usdm.coverage_tools` | Standalone MVP tool contracts for `validate_usdm`, `check_usdm_coverage`, `usdm_covered_by`, and scoped coverage reporting. |
| USDM requirement similarity collection | Contract | `spec:product.design_records.usdm.similarity_collection` | USDM-facing candidate collection operation for semantically similar requirement details. |

## Boundary

USDM owns normalized implementation requirements derived from Specification authorities or recorded directly as literal upstream requirements when no corresponding Specification exists.

USDM does not own design decisions, component architecture, implementation contracts, or public operation responses.

Coverage metadata records that an implementation Specification claims coverage. Coverage metadata does not prove implementation correctness.

## Related specs

| ref | relation |
|---|---|
| PRODUCT-REQ-SPEC-015 | Source requirement for MVP USDM artifacts and coverage checks. |
| PRODUCT-WORK-SPEC-029 | Source Work Item for MVP USDM Specification and tooling. |
| `spec:product.design_records.spec_format.document_shape` | Defines spec file shape rules used by this spec area. |
