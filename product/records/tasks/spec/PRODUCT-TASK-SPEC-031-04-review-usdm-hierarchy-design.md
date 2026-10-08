# PRODUCT-TASK-SPEC-031-04: Review USDM hierarchy design

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-031
- **task_type**: review
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-031-03
- **outputs**:

## Goal

Independently review the final combined hierarchical USDM design state.

## Work

- Verify the decision-to-ADR-to-Specification trace for D001-D008.
- Check row hierarchy, validation, direct and derived coverage, and warning semantics for contradictions.
- Check backward compatibility and explicit excluded scope.
- Return one integrated verdict without editing reviewed artifacts.

## Done condition

The independent review returns `PASS`, `NEEDS REVISION`, `NOT READY`, or `BLOCKED` with complete finding Evidence.

## Verification

The reviewer is independent of PRODUCT-TASK-SPEC-031-03 and reviews the exact final artifacts named in PRODUCT-WORK-SPEC-031.

## Evidence

Verdict: `NEEDS REVISION`.

Independent review session: `brewprint-usdm-hierarchy-review`.

Findings:

| id | severity | summary | required owner |
|---|---|---|---|
| REV-031-001 | Major | Hierarchical sibling allocation and non-reuse rules were authored without a terminal decision that explicitly selected them. | decision |
| REV-031-002 | Major | `coverage_tools` response shape leaves item fields, missing/error state, and direct/derived precedence under-specified. | decision |

No blocking findings were reported.

The reviewer confirmed D002-D008 semantics are otherwise consistent, including recursive derived coverage, direct-parent precedence, warning descendants beneath a directly covered ancestor, no implicit descendant direct coverage, and no semantic `optional` relation.

The review explicitly forbids proceeding to closure synchronization until the findings are decided, projected, and independently re-reviewed.
