# PRODUCT-WORK-SPEC-009: Format-only migration of traceability and artifact model specs

- **id**: PRODUCT-WORK-SPEC-009
- **status**: not_started
- **date**: 2026-06-15
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **impact_refs**:
- **tasks**:

## Goal

Apply the accepted spec format (PRODUCT-WORK-SPEC-001 contract) to `project-artifact-model/index.md` and `traceability/**` without executing ownership relocation. This is the format-migration phase split from PRODUCT-WORK-SPEC-005, enabling format compliance before PRODUCT-WORK-SPEC-004 ownership boundary decisions.

## Boundary

Format migration only. No ownership relocation, no hybrid-section splits, and no changes outside the 8 PRODUCT-INV-SPEC-002 target files. Stale `docs/...` path cleanup identified by PRODUCT-INV-SPEC-002 is included as migration cleanup, not ownership change.

## Impact Scope

8 spec files under `product/records/spec/concepts/`:

- `project-artifact-model/index.md`
- `traceability/index.md`
- `traceability/artifact-refs.md`
- `traceability/coverage-mapping.md`
- `traceability/metadata-schema.md`
- `traceability/out-of-scope.md`
- `traceability/resolve-and-validation.md`
- `traceability/semantic-ref.md`

## Task flow

TBD

## Task Candidates

TBD

## Completion Condition

- All 8 target files pass PRODUCT-WORK-SPEC-006 format validation.
- Stale `docs/...` path references removed or updated in target files.
- No ownership relocation or hybrid-section split performed.
- PRODUCT-WORK-SPEC-004 can proceed against format-compliant files.

## Evidence

TBD
