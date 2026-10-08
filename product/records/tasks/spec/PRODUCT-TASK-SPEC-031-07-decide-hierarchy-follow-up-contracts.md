# PRODUCT-TASK-SPEC-031-07: Decide USDM hierarchy follow-up contracts

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-031
- **task_type**: decision
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-031-06
- **outputs**:

## Goal

Resolve the two integrated-review findings and fix the namespace boundary for hierarchical USDM rows.

## Work

- Decide sibling allocation and removed-ID stability.
- Decide exact additive coverage-tool response shape.
- Decide whether row decomposition may cross USDM records or app namespaces.
- Preserve existing coverage and compact-range compatibility.

## Done condition

REV-031-001, REV-031-002, and the cross-app row-hierarchy question have terminal decisions.

## Verification

The decisions uniquely determine the ADR/spec changes without requiring implementation-time interpretation.

## Evidence

| decision | state | selected outcome | reason |
|---|---|---|---|
| D009 | decided | Each sibling set owns an independent numeric sequence. A newly authored child sibling set starts at `01`; allocation increases monotonically; gaps after reviewed removal are allowed; removed sibling IDs are not reused for a different requirement. | Stable descendant IDs need the same non-reuse property as existing top-level IDs without imposing one record-wide child sequence. |
| D010 | decided | Extend coverage-tool responses additively. Existing `covered` item maps remain direct-coverage maps. Add `derived_covered`, `refinement_warnings`, and retain `not_covered` as compact row-ID lists. `usdm_covered_by.coverage_state` is nullable and uses direct precedence over derived; missing/malformed/unavailable evaluation returns null. `covered_by` contains only direct covering Specification refs. | Preserves current callers while making all new effective-coverage states explicit and implementable. |
| D011 | decided | Requirement decomposition is local to exactly one USDM requirement record. A row cannot parent or child a row in another USDM record or app namespace. Cross-app implementation responsibility is represented by a Specification's full-ID `usdm_covers` reference to the target row. | Keeps recursive tree closure locally enumerable while preserving cross-app realization traceability. |
| D012 | decided | No new cross-record `refines` or `derived_from` relation is introduced now. | Cross-app `usdm_covers` satisfies the demonstrated need; a row-to-row relation can be added only when a concrete traceability need appears. |
