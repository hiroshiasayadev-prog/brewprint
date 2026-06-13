# PRODUCT-WORK-SPEC-008: ADR target specs traceability investigation and update plan

- **id**: PRODUCT-WORK-SPEC-008
- **status**: not_started
- **date**: 2026-06-13
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-007
- **impact_refs**:
- **tasks**:

## Summary

Investigate which existing ADRs semantically govern current spec files and define a follow-up update plan for ADR `target_specs` traceability.

This work captures the traceability follow-up from PRODUCT-WORK-SPEC-007 without reintroducing reverse traceability tables into spec documents.

## Scope

| area | in scope |
|---|---|
| ADR-to-spec mapping | Investigate which existing ADRs semantically govern current spec files. |
| target_specs decisions | Decide which ADRs should receive `target_specs`. |
| stale target handling | Define how to handle stale target specs when an ADR target becomes obsolete. |
| update planning | Identify required ADR metadata/update tasks. |

## Non-scope

| area | reason |
|---|---|
| `## Source records` in specs | Specs should not carry reverse traceability tables. |
| REQ/WORK/TASK as semantic sources | Spec changes should be governed by ADRs, not workflow records. |
| bulk ADR edits | This work plans ADR updates; actual edits require separately tasked work unless explicitly added later. |
| unrelated spec migration | Existing spec migration is owned elsewhere. |

## Done condition

| item | done when |
|---|---|
| candidate mapping listed | Candidate ADR-to-spec target mapping is documented. |
| update tasks identified | Required ADR metadata/update tasks are identified. |
| stale target policy recommended | Recommendation for stale target handling is documented. |

## Evidence

- Created from PRODUCT-TASK-SPEC-007-04 review finding that ADR-governed spec change traceability is needed.
- This work item preserves the decision not to reintroduce `## Source records` into specs and not to treat REQ/WORK/TASK as semantic sources for spec changes.

