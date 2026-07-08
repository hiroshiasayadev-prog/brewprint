# PRODUCT-TASK-SPEC-028-01: Decide retrospective term-analysis scope

- **id**: PRODUCT-TASK-SPEC-028-01
- **status**: not_started
- **date**: 2026-07-08
- **work_item**: PRODUCT-WORK-SPEC-028
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - PRODUCT-WORK-SPEC-028

## Goal

Decide how PRODUCT-WORK-SPEC-028 will admit, capture, and route the semantic-analysis work already performed over PRODUCT-INV-SPEC-011.

## Work

Resolve these decision items:

| item | status | required judgment |
|---|---|---|
| D-001 | open | Whether existing `tools/term-inventory-analysis/` outputs may be used as retrospective evidence. |
| D-002 | open | Which analysis stages are accepted as already completed. |
| D-003 | open | Which product-side evidence artifacts must be created or updated. |
| D-004 | open | Whether the Tier A cross-trigger review evidence needs additional independent review. |
| D-005 | open | Which next Work Items or Tasks are required before canonical vocabulary decisions. |
| D-006 | open | Whether PRODUCT-REQ-SPEC-012 restart criteria can be partially or fully satisfied by this Work Item. |

Do not decide canonical vocabulary inside this Task.
Do not rewrite source records inside this Task.

## Done condition

- Every decision item is `decided`, `deferred`, or validly `blocked`.
- The accepted analysis scope is explicit.
- The retrospective evidence policy is explicit.
- The next Task graph or stop condition is explicit.

## Verification

Verify that PRODUCT-WORK-SPEC-028 still lists this Task.
Verify that no canonical vocabulary, deprecation, source rewrite, or Specification projection was performed by this Task.

## Evidence

- PRODUCT-REQ-SPEC-014 requires product-owned semantic-analysis evidence.
- PRODUCT-INV-SPEC-011 is the raw evidence source.
- The existing cross-trigger review summary is currently stored as commit-safe evidence, but its parent Work Item boundary was missing before PRODUCT-WORK-SPEC-028.
