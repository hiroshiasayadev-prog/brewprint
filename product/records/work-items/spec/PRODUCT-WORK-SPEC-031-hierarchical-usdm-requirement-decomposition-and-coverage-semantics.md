# PRODUCT-WORK-SPEC-031: Hierarchical USDM requirement decomposition and coverage semantics

- **status**: done
- **date**: 2026-09-30
- **source_refs**:
  - PRODUCT-REQ-SPEC-016
  - PRODUCT-TASK-SPEC-030-02
- **impact_refs**:
  - `spec:product.design_records.usdm`
  - `spec:product.design_records.usdm.artifact_format`
  - `spec:product.design_records.usdm.coverage_format`
  - `spec:product.design_records.usdm.coverage_tools`
  - `spec:product.design_records.authoring_standards.usdm_authoring`
- **tasks**:
  - PRODUCT-TASK-SPEC-031-01
  - PRODUCT-TASK-SPEC-031-02
  - PRODUCT-TASK-SPEC-031-03
  - PRODUCT-TASK-SPEC-031-04
  - PRODUCT-TASK-SPEC-031-05
  - PRODUCT-TASK-SPEC-031-06
  - PRODUCT-TASK-SPEC-031-07
  - PRODUCT-TASK-SPEC-031-08
  - PRODUCT-TASK-SPEC-031-09

## Goal

Establish hierarchical USDM requirement decomposition and effective coverage semantics as accepted durable design and current normative Specification state.

## Boundary

This Work Item owns:

- the durable hierarchy and effective-coverage decision;
- one accepted ADR for that decision;
- normative updates to the USDM artifact, coverage, tool-contract, overview, and authoring Specifications;
- one integrated independent review;
- mechanical design-closure synchronization.

This Work Item does not own:

- standalone or MCP tool implementation;
- migration of existing USDM records;
- explicit optional, OR, conditional, weighted, or cardinality relation semantics;
- section-level coverage.

## Impact Scope

| target | impact |
|---|---|
| PRODUCT-ADR-SPEC-020 | Record the hierarchy and effective coverage decision. |
| `spec:product.design_records.usdm.artifact_format` | Extend row IDs and define parent-child structure. |
| `spec:product.design_records.usdm.coverage_format` | Define direct, derived, blocking, and warning coverage behavior. |
| `spec:product.design_records.usdm.coverage_tools` | Align tool contract outputs with effective coverage. |
| `spec:product.design_records.usdm` | Summarize the current hierarchical coverage contract. |
| `spec:product.design_records.authoring_standards.usdm_authoring` | Add hierarchy authoring and review rules. |

## Task flow

```text
PRODUCT-TASK-SPEC-031-01 decide hierarchy and coverage semantics
  -> PRODUCT-TASK-SPEC-031-02 route ADR
  -> PRODUCT-TASK-SPEC-031-03 author ADR and Specifications
  -> PRODUCT-TASK-SPEC-031-04 integrated independent review (NEEDS REVISION)
  -> PRODUCT-TASK-SPEC-031-06 coordinate findings
  -> PRODUCT-TASK-SPEC-031-07 decide follow-up contracts
  -> PRODUCT-TASK-SPEC-031-08 author follow-up contracts
  -> PRODUCT-TASK-SPEC-031-09 independent re-review
  -> PRODUCT-TASK-SPEC-031-05 closure synchronization
```

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| PRODUCT-TASK-SPEC-031-01 | decision | Fix hierarchy and effective coverage semantics. | none |
| PRODUCT-TASK-SPEC-031-02 | decision | Route the durable decision to one ADR boundary. | T01 |
| PRODUCT-TASK-SPEC-031-03 | authoring | Author the routed ADR and synchronized USDM Specifications. | T02 |
| PRODUCT-TASK-SPEC-031-04 | review | Independently review the final combined design state. | T03 |
| PRODUCT-TASK-SPEC-031-05 | synchronization | Propagate the final accepted re-review state and close this Work Item. | T09 |
| PRODUCT-TASK-SPEC-031-06 | coordination | Route REV-031-001/002 and the accepted namespace-boundary follow-up. | T04 |
| PRODUCT-TASK-SPEC-031-07 | decision | Decide sibling allocation, additive response shape, and record-local hierarchy boundary. | T06 |
| PRODUCT-TASK-SPEC-031-08 | authoring | Project D009-D012 and close the review findings. | T07 |
| PRODUCT-TASK-SPEC-031-09 | review | Independently re-review the revised combined design. | T08 |

## Completion Condition

- Hierarchical row IDs and parent derivation are normative.
- Effective coverage is direct coverage OR recursive all-child coverage.
- Direct parent coverage remains sufficient while uncovered descendants are visible as non-blocking warnings.
- No explicit optional or alternative relation semantic is introduced.
- PRODUCT-ADR-SPEC-020 is accepted and reflected into all affected USDM Specifications.
- The final integrated independent review returns `PASS`, or every required finding is independently closed.
- Tool implementation and migration remain explicit downstream scope rather than hidden design work.

## Evidence

- PRODUCT-REQ-SPEC-016 is the direct Requirement source.
- PRODUCT-TASK-SPEC-030-02 created this Work Item from the accepted framing route.
- PRODUCT-TASK-SPEC-031-01 fixed the hierarchy and effective coverage decisions.
- PRODUCT-TASK-SPEC-031-02 routed the durable choice to PRODUCT-ADR-SPEC-020.
- PRODUCT-TASK-SPEC-031-04 returned `NEEDS REVISION` with REV-031-001 and REV-031-002.
- PRODUCT-TASK-SPEC-031-06 through 031-09 own append-only reconvergence.
- PRODUCT-TASK-SPEC-031-08 projected D009-D012 and closed the two review findings in canonical artifacts.
- PRODUCT-TASK-SPEC-031-09 returned `PASS` and independently closed REV-031-001 and REV-031-002.
- PRODUCT-TASK-SPEC-031-05 synchronized the accepted result and verified every Completion Condition.
- Tool implementation and migration remain downstream work and are not part of this completed design boundary.
