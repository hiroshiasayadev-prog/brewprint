# PRODUCT-TASK-SPEC-033-02: Review hierarchical USDM tool implementation

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-033
- **task_type**: review
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-033-01
- **outputs**:

## Goal

Independently review hierarchical USDM tooling against the accepted Specifications and implementation Evidence.

## Work

- Inspect only PRODUCT-TASK-SPEC-033-01 owned implementation paths and tests.
- Verify hierarchical parsing, parent validation, recursive coverage, warning behavior, cross-app coverage, compatibility, similarity loading, and MCP forwarding.
- Verify tests exercise the accepted boundary.
- Return a review verdict without modifying implementation.

## Done condition

The independent implementation review returns `PASS`, `NEEDS REVISION`, `NOT READY`, or `BLOCKED` with concrete findings.

## Verification

The reviewer is independent of the implementation Task and reads actual source/test diffs and execution Evidence.

## Evidence

Verdict: `NEEDS REVISION`.

Independent reviewer inspected the live working-tree implementation and reran the scoped unit tests and current-corpus validation in read-only mode.

Finding:

| id | severity | summary | user judgment |
|---|---|---|---|
| REV-033-001 | Major / closure-blocking | `check_usdm_scope_coverage` drops `scan_coverage(...).diagnostics` and `scan_usdm(...).diagnostics`, so malformed coverage or invalid hierarchy can return `ok: true` with empty diagnostics. | not required |

Confirmed passing areas include hierarchical parsing, local-parent validation, cross-app full-ID coverage, compact tokens/ranges, recursive effective coverage, direct precedence, warning/blocking classification, nearest direct ancestor, all `usdm_covered_by` states, scope category/filter/sort behavior, similarity loading, and MCP `include_warnings` forwarding.

Reviewer reproduced REV-033-001 with:
- valid direct coverage plus malformed `usdm_covers` -> `ok: true`, `diagnostics: []`;
- orphan hierarchical row -> `ok: true`, `diagnostics: []`.

Required correction:
- propagate relevant USDM and coverage scan error diagnostics into `check_usdm_scope_coverage`;
- error diagnostics must make `ok: false`;
- add regression tests for malformed coverage and orphan hierarchy;
- preserve existing scope relevance semantics.

Review Evidence also confirmed the implementer-recorded test results:
- 34 USDM tests: PASS;
- 23 similarity tests: PASS;
- current corpus: 46 records, 353 requirements, diagnostics 0.

PRODUCT-WORK-SPEC-033 is not ready for closure until correction and finding-closure re-review pass.
