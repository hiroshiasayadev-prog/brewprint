# PRODUCT-WORK-SPEC-028: Reduce and review design-governance term identities

- **id**: PRODUCT-WORK-SPEC-028
- **status**: in_progress
- **date**: 2026-07-08
- **source_refs**:
  - PRODUCT-REQ-SPEC-014
  - PRODUCT-INV-SPEC-011
- **impact_refs**:
  - PRODUCT-REQ-SPEC-012
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/leaf-analysis/
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/trigger-reduction/
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/cross-trigger-review/
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/semantic-analysis/
- **tasks**:
  - PRODUCT-TASK-SPEC-028-01
  - PRODUCT-TASK-SPEC-028-02
  - PRODUCT-TASK-SPEC-028-03
  - PRODUCT-TASK-SPEC-028-04
  - PRODUCT-TASK-SPEC-028-05
  - PRODUCT-TASK-SPEC-028-06
  - PRODUCT-TASK-SPEC-028-07

## Goal

Recover the completed and in-progress semantic-analysis work over PRODUCT-INV-SPEC-011 into a product-owned Work Item boundary.

Establish reviewable evidence for trigger-level reduction, cross-trigger identity candidates, reviewed identity groups, unresolved candidates, and follow-up routing.

## Boundary

This Work Item owns:

- the retrospective evidence policy for analysis already performed under `tools/term-inventory-analysis/`;
- product-side evidence capture for trigger-level aggregation and reduction;
- product-side evidence capture for cross-trigger candidate generation and routing;
- product-side evidence capture for Tier A identity review completion and spot audit;
- separation of reviewed semantic-analysis evidence from canonical vocabulary approval;
- next-route decisions for vocabulary, conflict, qualified-term, deprecation, and PRODUCT-REQ-SPEC-012 restart work.

This Work Item does not own:

- raw corpus extraction;
- changes to the 32 batch-owned PRODUCT-INV-SPEC-011 observation files;
- canonical vocabulary approval;
- term definition authoring;
- term retirement or deprecation decisions;
- source-record rewrites;
- Specification, skill, authoring-guide, or validator projection;
- production implementation.

## Impact Scope

| target | impact |
|---|---|
| `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/` | Add product-owned semantic-analysis evidence derived from the raw corpus. |
| `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/semantic-analysis/index.md` | Index product-side capture status and accepted retrospective source stages. |
| `tools/term-inventory-analysis/` | Treated as ignored working output unless specific summaries are captured under product records. |
| PRODUCT-REQ-SPEC-012 | May receive restart evidence after foundational term boundaries are classified. |

## Task flow

```text
PRODUCT-TASK-SPEC-028-01 decide retrospective analysis scope and route
  -> PRODUCT-TASK-SPEC-028-02 capture leaf-analysis evidence
  -> PRODUCT-TASK-SPEC-028-05 coordinate follow-up routes
  -> PRODUCT-TASK-SPEC-028-06 review retrospective evidence route
  -> PRODUCT-TASK-SPEC-028-07 synchronize retrospective evidence closure

PRODUCT-TASK-SPEC-028-01 decide retrospective analysis scope and route
  -> PRODUCT-TASK-SPEC-028-03 capture trigger-reduction evidence
  -> PRODUCT-TASK-SPEC-028-05 coordinate follow-up routes

PRODUCT-TASK-SPEC-028-01 decide retrospective analysis scope and route
  -> PRODUCT-TASK-SPEC-028-04 synchronize cross-trigger review evidence
  -> PRODUCT-TASK-SPEC-028-05 coordinate follow-up routes
```

T02, T03, and T04 may run after T01.
T05 waits for all product-side evidence capture and synchronization Tasks.
T06 reviews the captured evidence and routing.
T07 runs only after the review route is satisfied.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| PRODUCT-TASK-SPEC-028-01 | decision | Decide the retrospective evidence policy, accepted analysis scope, exact product-side outputs, and next Task graph. | none |
| PRODUCT-TASK-SPEC-028-02 | investigation | Capture completed leaf semantic analysis into product-side evidence. | T01 |
| PRODUCT-TASK-SPEC-028-03 | investigation | Capture completed trigger-level reduction into product-side evidence. | T01 |
| PRODUCT-TASK-SPEC-028-04 | synchronization | Connect existing Tier A cross-trigger review evidence to this Work Item's evidence policy. | T01 |
| PRODUCT-TASK-SPEC-028-05 | coordination | Coordinate follow-up routes for canonical vocabulary, conflicting meanings, qualified terms, deprecation, and PRODUCT-REQ-SPEC-012 restart. | T02, T03, T04 |
| PRODUCT-TASK-SPEC-028-06 | review | Independently review captured evidence and follow-up routing. | T02, T03, T04, T05 |
| PRODUCT-TASK-SPEC-028-07 | synchronization | Synchronize lifecycle, Evidence, relations, and closure after review passes. | T06 |

## Completion Condition

- The Work Item records how completed tools-side analysis is admitted or rejected as product-side evidence.
- Product-owned evidence exists for every analysis stage the Work Item accepts as completed.
- Reviewed identity groups remain separate from canonical vocabulary decisions.
- Unresolved candidates and relationship hints remain explicit.
- Follow-up routes are decided for canonical vocabulary, conflicting meanings, qualified terms, deprecated wording, and PRODUCT-REQ-SPEC-012 restart criteria.
- No source artifact or normative Specification is changed without a separate accepted Work Item.

## Evidence

- PRODUCT-REQ-SPEC-014 requires semantic analysis over PRODUCT-INV-SPEC-011.
- PRODUCT-INV-SPEC-011 records 5,699 unclassified and unnormalized observations.
- PRODUCT-WORK-SPEC-027 completed raw inventory and excluded semantic aggregation.
- Commit-safe Tier A cross-trigger review evidence already exists under `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/cross-trigger-review/`.
- Raw per-job analysis output remains under ignored `tools/term-inventory-analysis/` output and is not product history by itself.
- PRODUCT-TASK-SPEC-028-01 decided the retrospective evidence policy, accepted completed analysis stages, required product-side evidence artifacts, review route, follow-up Task graph, and PRODUCT-REQ-SPEC-012 partial-restart boundary.
- PRODUCT-TASK-SPEC-028-02 through PRODUCT-TASK-SPEC-028-07 were materialized from PRODUCT-TASK-SPEC-028-01 D-005.
