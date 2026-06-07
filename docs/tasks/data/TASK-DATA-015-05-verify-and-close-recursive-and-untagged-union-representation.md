# TASK-DATA-015-05: Verify and close recursive and untagged-union representation

- **id**: TASK-DATA-015-05
- **status**: not_started
- **date**: 2026-06-07
- **work_item**: WORK-DATA-015
- **source_requirement**: REQ-DATA-008
- **estimate**: 0.5d
- **depends_on**:
  - TASK-DATA-015-03
  - TASK-DATA-015-04
- **outputs**:
  - Verification summary for WORK-DATA-015
  - WORK-DATA-015 close evidence
  - REQ-DATA-008 close / accepted update input if appropriate

## Goal

Verify the selected recursive named reference / unsupported untagged union boundary and close WORK-DATA-015 when all selected follow-up work is complete.

## Work

- Review TASK-DATA-015-01 through TASK-DATA-015-04 results.
- Confirm recursive named model reference support is specified and either implemented or explicitly deferred with rationale.
- Confirm untagged union / general `oneOf` remains unsupported and no hidden broadening occurred.
- Confirm any UC-002 cleanup, fixture, or render follow-up is complete or explicitly deferred.
- Update WORK-DATA-015 close evidence when ready.
- Identify whether REQ-DATA-008 can move from `captured` to `accepted` or should remain open.

## Included Scope

- Verification and close synchronization.
- WORK-DATA-015 evidence update.
- Requirement status recommendation.

## Excluded Scope

- New implementation work.
- New UC-002 YAML migration.
- Golden regeneration.
- Untagged union / general `oneOf` support.

## Done condition

- WORK-DATA-015 is closed or explicitly left open with remaining blocker rationale.
- Verification evidence is recorded.
- Requirement follow-up status is identified.

## Verification

- Validate affected task / work item / requirement records after close updates.
- Record any repo-local test command outputs if implementation or YAML changed in prior tasks.

## Evidence

Not started.
