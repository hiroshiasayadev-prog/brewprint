# PRODUCT-TASK-SPEC-031-08: Author USDM hierarchy follow-up contracts

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-031
- **task_type**: authoring
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-031-07
- **outputs**:
  - PRODUCT-ADR-SPEC-020
  - `spec:product.design_records.usdm`
  - `spec:product.design_records.usdm.artifact_format`
  - `spec:product.design_records.usdm.coverage_format`
  - `spec:product.design_records.usdm.coverage_tools`
  - `spec:product.design_records.authoring_standards.usdm_authoring`

## Goal

Project D009-D012 and close REV-031-001 and REV-031-002 without changing the accepted core coverage algorithm.

## Work

- Amend PRODUCT-ADR-SPEC-020 with sibling allocation, record-local hierarchy, and cross-app coverage boundaries.
- Make the coverage-tool response schema exact and additive.
- Align artifact-format, coverage-format, overview, and authoring guidance.
- Preserve D001-D008 semantics and existing flat-ID compatibility.

## Done condition

The reviewed findings are fully projected and no implementation-time schema or hierarchy judgment remains.

## Verification

Inspect the scoped diff and rerun USDM validation before independent review.

## Evidence

- PRODUCT-ADR-SPEC-020 now records sibling-set allocation, record-local row hierarchy, cross-app coverage responsibility, and no cross-record refinement relation.
- PRODUCT-REQ-SPEC-016 explicitly keeps row decomposition within one requirement record/app namespace while allowing cross-app full-ID `usdm_covers`.
- `artifact_format` defines child `01` start, sibling-local monotonic allocation, removed-ID non-reuse, and record-local parent-child identity.
- `coverage_format` explicitly separates cross-app direct coverage from row hierarchy.
- `coverage_tools` defines additive direct/derived/warning output fields and nullable `coverage_state` with direct precedence.
- `usdm_authoring` instructs authors to use cross-app `usdm_covers` rather than cross-record parentage.
- `validate_usdm` returned `ok: true` for 46 records and 353 requirements with no diagnostics.
- Scoped Git whitespace inspection returned `result: pass`.
- No tool implementation was modified by this Task.
