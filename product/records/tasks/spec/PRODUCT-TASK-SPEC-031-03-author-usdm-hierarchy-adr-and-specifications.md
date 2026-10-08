# PRODUCT-TASK-SPEC-031-03: Author USDM hierarchy ADR and Specifications

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-031
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-031-02
- **outputs**:
  - PRODUCT-ADR-SPEC-020
  - `spec:product.design_records.usdm.artifact_format`
  - `spec:product.design_records.usdm.coverage_format`
  - `spec:product.design_records.usdm.coverage_tools`
  - `spec:product.design_records.usdm`
  - `spec:product.design_records.authoring_standards.usdm_authoring`

## Goal

Project the accepted hierarchy and effective-coverage decisions into one ADR and the current USDM normative Specifications.

## Work

- Create PRODUCT-ADR-SPEC-020 from the routed decision set.
- Update hierarchical row ID, parent, and validation rules.
- Update direct and derived coverage semantics and warning behavior.
- Update coverage-tool contracts and authoring guidance.
- Update the USDM overview without adding implementation work.

## Done condition

The ADR and every affected Specification express D001-D008 consistently with no stale flat-only coverage rule in the authoring boundary.

## Verification

Inspect the scoped diff, validate changed USDM records where applicable, and confirm no tool implementation or unrelated record changed.

## Evidence

- Created PRODUCT-ADR-SPEC-020 with accepted hierarchy, effective-coverage, warning, compatibility, and deferred-relation decisions.
- Updated `spec:product.design_records.usdm.artifact_format` with `RNNN(-NN)*`, parent derivation, sibling numbering, and missing-parent validation.
- Updated `spec:product.design_records.usdm.coverage_format` with direct coverage, recursive derived coverage, blocking uncovered requirements, and non-blocking refinement warnings.
- Updated `spec:product.design_records.usdm.coverage_tools` with effective coverage states and warning-aware response semantics.
- Updated `spec:product.design_records.usdm` and `spec:product.design_records.authoring_standards.usdm_authoring` for the current contract and authoring rules.
- Existing `RNNN` IDs and top-level `RNNN-RNNN` compact ranges remain valid.
- `validate_usdm` returned `ok: true` for 46 records and 353 requirements with no diagnostics after authoring.
- Scoped Git diff inspection returned `result: pass`; line-ending conversion notices are advisory only.
- No standalone or MCP tool implementation was changed.
- Design Records MCP authoring transaction tools are unavailable in this session, so filesystem authoring used the documented fallback.
- Independent review remains outside this authoring Task.
