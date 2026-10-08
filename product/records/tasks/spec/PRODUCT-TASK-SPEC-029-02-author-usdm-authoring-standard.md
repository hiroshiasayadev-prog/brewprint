# PRODUCT-TASK-SPEC-029-02: Author USDM authoring standard

- **id**: PRODUCT-TASK-SPEC-029-02
- **status**: done
- **date**: 2026-07-08
- **work_item**: PRODUCT-WORK-SPEC-029
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-029-01
- **outputs**:
  - `spec:product.design_records.authoring_standards.usdm_authoring`
  - `spec:product.design_records.authoring_standards`

## Goal

Author the MVP USDM authoring standard needed by PRODUCT-WORK-SPEC-029.

## Work

- Add `spec:product.design_records.authoring_standards.usdm_authoring`.
- Add USDM topic splitting, requirement-row writing, row-ID stability, source-spec, and coverage authoring rules.
- Add the USDM authoring topic row to the Authoring standards index.
- Do not implement USDM tools.
- Do not author USDM requirement records.
- Do not migrate existing records.

## Boundary

This Task owns only author-facing USDM writing rules.

This Task does not own USDM format contracts, standalone tool implementation, sample smoke fixtures, DRMCP integrated read support, artifact-model integration, repository-layout integration, or existing-record migration.

## Done condition

- `usdm-authoring.md` exists under `product/records/spec/design-records/authoring-standards/`.
- The Authoring standards index links to the USDM authoring standard.
- The standard explains how to split topics, write requirement rows, assign row IDs, preserve row stability, and declare coverage.
- The standard keeps USDM records separate from design decisions, implementation contracts, and task instructions.

## Verification

- Created `product/records/spec/design-records/authoring-standards/usdm-authoring.md`.
- Updated `product/records/spec/design-records/authoring-standards/index.md` with the USDM authoring topic row.
- Did not implement `tools/usdm/`.
- Did not create USDM requirement records.
- Did not migrate existing records.

## Evidence

- PRODUCT-REQ-SPEC-015 requires a USDM authoring standard.
- PRODUCT-WORK-SPEC-029 assigns USDM authoring-standard work to PRODUCT-TASK-SPEC-029-02.
- PRODUCT-TASK-SPEC-029-01 completed the USDM format and coverage contract prerequisites.
