# PRODUCT-WORK-SPEC-026: Frame Brewprint design-governance term inventory

- **id**: PRODUCT-WORK-SPEC-026
- **status**: done
- **date**: 2026-07-04
- **source_refs**:
  - PRODUCT-REQ-SPEC-013
- **impact_refs**: []
- **tasks**:
  - PRODUCT-TASK-SPEC-026-01
  - PRODUCT-TASK-SPEC-026-02
  - PRODUCT-TASK-SPEC-026-03

## Goal

Decide how PRODUCT-REQ-SPEC-013 will be handled and fix one actionable downstream Work Item contract for the Brewprint design-governance term inventory.

## Boundary

This Work Item owns:

- Desired Outcome and Required Outcome alignment;
- source disposition;
- downstream Goal, Boundary, and Completion Condition;
- unknown handling;
- exact initial downstream route;
- materialization of only the Tasks uniquely required by accepted framing decisions.

This Work Item does not own:

- corpus scanning;
- JSONL observation authoring;
- helper-tool implementation;
- corpus partitioning or extraction Task execution;
- term classification, definition, consolidation, or retirement;
- Task-type use-case vocabulary.

## Impact Scope

| target | impact |
|---|---|
| PRODUCT-REQ-SPEC-013 | Decide the accepted handling and downstream resolution route. |
| downstream inventory Work Item | Fix its identity, completion meaning, unknown handling, and initial route before creation. |

## Task flow

```text
PRODUCT-TASK-SPEC-026-01 framing decision
  -> PRODUCT-TASK-SPEC-026-02 amend Requirement boundary
     -> PRODUCT-TASK-SPEC-026-03 create downstream inventory Work Item
```

T02 and T03 are materialized from the completed framing decision because the Requirement clarification and downstream Work Item contract are fixed.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| PRODUCT-TASK-SPEC-026-01 | decision | Align outcomes, select source disposition, and fix the downstream Work Item contract and route. | none |
| PRODUCT-TASK-SPEC-026-02 | authoring | Amend PRODUCT-REQ-SPEC-013 to exclude aggregation and helper-tool ownership. | T01 |
| PRODUCT-TASK-SPEC-026-03 | work_item_decomposition | Create the downstream inventory Work Item with one initial decision Task. | T02 |

## Completion Condition

- Desired Outcome and Required Outcome alignment is explicit.
- One source disposition and reason are explicit.
- If the disposition is `proceed`, the downstream Goal, Boundary, Completion Condition, direct source, unknown handling, and initial route are fixed.
- Every Task required by the accepted framing route is materialized and completed.
- No corpus extraction, aggregation, classification, definition, or helper-tool implementation is performed by this framing Work Item.

## Evidence

- The user accepted PRODUCT-REQ-SPEC-013 as the Requirement for machine-readable corpus discovery.
- The user directed proceeding to Work Item planning after Requirement creation.
- The user previously placed exact corpus scope, Task partitioning, and output placement in the downstream Work Item decision loop rather than the Requirement.
- The user excluded helper-tool implementation, validation tooling, aggregation, and use-case extraction from the downstream inventory Work Item.
- T01 fixed the complete framing contract.
- T02 amended PRODUCT-REQ-SPEC-013 to preserve the investigation-only boundary.
- T03 created PRODUCT-WORK-SPEC-027 and its single initial decision Task, PRODUCT-TASK-SPEC-027-01.
- Every materialized framing Task is `done`; the framing Completion Condition is satisfied.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
