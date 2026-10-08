# PRODUCT-WORK-SPEC-033: Implement hierarchical USDM tooling

- **status**: done
- **date**: 2026-09-30
- **source_refs**:
  - PRODUCT-REQ-SPEC-016
  - PRODUCT-TASK-SPEC-032-02
  - `spec:product.design_records.usdm.artifact_format`
  - `spec:product.design_records.usdm.coverage_format`
  - `spec:product.design_records.usdm.coverage_tools`
- **impact_refs**:
  - `spec:product.design_records.usdm.artifact_format`
  - `spec:product.design_records.usdm.coverage_format`
  - `spec:product.design_records.usdm.coverage_tools`
  - `spec:product.design_records.usdm.similarity_collection`
- **tasks**:
  - PRODUCT-TASK-SPEC-033-01
  - PRODUCT-TASK-SPEC-033-02
  - PRODUCT-TASK-SPEC-033-03
  - PRODUCT-TASK-SPEC-033-04

## Goal

Implement the reviewed hierarchical USDM row and effective-coverage contracts in standalone tools and the MCP wrapper.

## Boundary

This Work Item owns:

- hierarchical requirement-ID parsing and validation;
- immediate-parent existence validation;
- hierarchical compact coverage tokens while retaining top-level ranges;
- direct, derived, blocking-uncovered, and refinement-warning evaluation;
- exact additive coverage response projection;
- hierarchical requirement scope support in the similarity loader;
- MCP request exposure for the revised scope-coverage contract;
- scoped automated tests and current-corpus validation;
- independent implementation review.

This Work Item does not own:

- new USDM semantics or ADR changes;
- cross-record row hierarchy;
- explicit optional/OR/conditional relation semantics;
- migration of existing USDM records;
- unrelated DRMCP, PuppyDSL, or validation work.

## Impact Scope

| target | impact |
|---|---|
| `tools/usdm/usdm_tools.py` | Parse/validate hierarchy and compute effective coverage. |
| `tools/usdm/tests/` | Add regression tests for hierarchical IDs and coverage states. |
| `tools/usdm/similarity/usdm_loader.py` | Accept and load hierarchical full requirement IDs. |
| `tools/usdm/similarity/tests/` | Add hierarchical scope/load regression coverage. |
| `tools/usdm-mcp/server.py` | Expose `include_warnings` and updated public tool descriptions. |

## Task flow

```text
PRODUCT-TASK-SPEC-033-01 implement and verify
  -> PRODUCT-TASK-SPEC-033-02 independent implementation review (NEEDS REVISION)
  -> PRODUCT-TASK-SPEC-033-03 correct REV-033-001
  -> PRODUCT-TASK-SPEC-033-04 finding-closure re-review
```

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| PRODUCT-TASK-SPEC-033-01 | implementation | Implement reviewed hierarchical USDM tooling and run scoped verification. | none |
| PRODUCT-TASK-SPEC-033-02 | review | Independently review implementation/spec conformance and test evidence. | T01 |
| PRODUCT-TASK-SPEC-033-03 | implementation | Correct REV-033-001 diagnostic propagation and add regression coverage. | T02 |
| PRODUCT-TASK-SPEC-033-04 | review | Independently verify REV-033-001 closure and regression safety. | T03 |

## Completion Condition

- Hierarchical IDs `RNNN(-NN)*` are accepted by all affected standalone USDM parsing/scope paths.
- Orphan hierarchical rows are rejected.
- Top-level compact ranges continue to work; hierarchical range shorthand is rejected.
- Effective coverage and warning semantics match PRODUCT-ADR-SPEC-020.
- Cross-app `usdm_covers` continues to count as direct coverage of the named target row.
- Scope responses preserve existing direct `covered` maps and add derived/warning categories exactly as specified.
- `usdm_covered_by` reports exact state, direct-only `covered_by`, and nearest direct-covered ancestor for warnings.
- Similarity requirement loading and exact requirement scopes accept hierarchical IDs.
- MCP exposes the revised scope request.
- Scoped automated tests pass and the current repository USDM corpus validates.
- Final independent implementation review passes before Work Item closure.

## Evidence

- PRODUCT-TASK-SPEC-033-01 implemented the reviewed design and passed its scoped verification.
- PRODUCT-TASK-SPEC-033-02 returned `NEEDS REVISION` with REV-033-001.
- PRODUCT-TASK-SPEC-033-03 corrected REV-033-001 and added regression coverage; 36 USDM tests and 23 similarity tests passed, and current-corpus validation remained clean.
- PRODUCT-TASK-SPEC-033-04 returned `PASS` and independently closed REV-033-001.
- Every Completion Condition above is satisfied; no further implementation correction or design decision is required for this Work Item.
