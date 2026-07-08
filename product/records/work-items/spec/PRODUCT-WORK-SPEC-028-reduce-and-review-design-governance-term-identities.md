# PRODUCT-WORK-SPEC-028: Reduce and review design-governance term identities

- **id**: PRODUCT-WORK-SPEC-028
- **status**: not_started
- **date**: 2026-07-08
- **source_refs**:
  - PRODUCT-REQ-SPEC-014
  - PRODUCT-INV-SPEC-011
- **impact_refs**:
  - PRODUCT-REQ-SPEC-012
  - product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/cross-trigger-review/
- **tasks**:
  - PRODUCT-TASK-SPEC-028-01

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
| `tools/term-inventory-analysis/` | Treated as ignored working output unless specific summaries are captured under product records. |
| PRODUCT-REQ-SPEC-012 | May receive restart evidence after foundational term boundaries are classified. |

## Task flow

```text
PRODUCT-TASK-SPEC-028-01 decide retrospective analysis scope and route
  -> later tasks materialized only after T01 fixes exact owners and outputs
```

T01 is the only initial Task.
No aggregation, review, correction, synchronization, or canonical authoring Task is materialized before T01 decides its exact responsibility boundary.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| PRODUCT-TASK-SPEC-028-01 | decision | Decide the retrospective evidence policy, accepted analysis scope, exact product-side outputs, and next Task graph. | none |
| candidate | investigation | Record trigger-level reduction evidence if T01 confirms that the existing tool output is admissible retrospective evidence. | T01 |
| candidate | investigation | Record cross-trigger identity candidate generation and routing evidence. | T01 |
| candidate | review | Review or audit product-side identity evidence when T01 requires an independent gate. | T01 and relevant evidence Task |
| candidate | coordination | Create follow-up Work Items for canonical vocabulary, conflicting meanings, qualified terms, deprecation, or PRODUCT-REQ-SPEC-012 restart. | T01 and accepted evidence Tasks |

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
