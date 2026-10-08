# PRODUCT-TASK-SPEC-029-01: Author MVP USDM artifact Specification

- **id**: PRODUCT-TASK-SPEC-029-01
- **status**: done
- **date**: 2026-07-08
- **work_item**: PRODUCT-WORK-SPEC-029
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - `spec:product.design_records.usdm`
  - `spec:product.design_records.usdm.artifact_format`
  - `spec:product.design_records.usdm.coverage_format`
  - `spec:product.design_records.usdm.coverage_tools`

## Startup / Required reading

Before authoring, preserve these records and rules:

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/writing-standard.md`
3. `product/records/spec/design-records/authoring-standards/spec-authoring.md`
4. `product/records/spec/design-records/spec-format/document-shape.md`
5. `product/records/requirements/spec/PRODUCT-REQ-SPEC-015-mvp-usdm-requirement-artifacts-and-coverage-checks.md`
6. `product/records/work-items/spec/PRODUCT-WORK-SPEC-029-mvp-usdm-artifacts-and-coverage-tooling.md`

## Goal

Author the MVP USDM Specification records needed by PRODUCT-WORK-SPEC-029.

## Work

- Add `spec:product.design_records.usdm` as the USDM overview entry point.
- Add USDM artifact format rules for MVP USDM records.
- Add USDM coverage format rules for implementation Specification metadata.
- Add standalone USDM coverage tool contracts.
- Add the USDM topic row to the parent Design Records spec index.
- Keep `usdm-authoring.md` for PRODUCT-TASK-SPEC-029-02.
- Do not implement tools.
- Do not migrate existing records.

## Boundary

This Task owns only Product Specification authoring for the USDM MVP.

This Task does not own USDM authoring-standard prose, standalone tool implementation, sample smoke fixtures, DRMCP integrated read support, artifact-model integration, repository-layout integration, or existing-record migration.

## Done condition

- USDM overview, artifact-format, coverage-format, and coverage-tool specs exist under `product/records/spec/design-records/usdm/`.
- The parent Design Records index links to the USDM overview.
- The specs state that USDM is an independent auxiliary artifact during the MVP.
- The specs state that standalone tools may later move behind an MCP surface.
- No implementation task is made ready beyond the standalone tool contract surface.

## Verification

- Created `product/records/spec/design-records/usdm/index.md`.
- Created `product/records/spec/design-records/usdm/artifact-format.md`.
- Created `product/records/spec/design-records/usdm/coverage-format.md`.
- Created `product/records/spec/design-records/usdm/coverage-tools.md`.
- Updated `product/records/spec/design-records/index.md` with the USDM topic row.
- Did not create `usdm-authoring.md`.
- Did not implement `tools/usdm/`.
- Did not migrate existing records.

## Evidence

- PRODUCT-REQ-SPEC-015 requires MVP USDM Specification records and standalone coverage-tool behavior.
- PRODUCT-WORK-SPEC-029 assigns USDM Specification authoring to PRODUCT-TASK-SPEC-029-01.
