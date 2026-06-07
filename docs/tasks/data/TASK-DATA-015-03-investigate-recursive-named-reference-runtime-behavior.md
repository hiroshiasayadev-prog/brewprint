# TASK-DATA-015-03: Investigate recursive named reference runtime behavior

- **id**: TASK-DATA-015-03
- **status**: not_started
- **date**: 2026-06-07
- **work_item**: WORK-DATA-015
- **source_requirement**: REQ-DATA-008
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-015-02
- **outputs**:
  - Current resolver / validator / renderer behavior for recursive named model references
  - Implementation gap classification if current behavior rejects or expands recursion incorrectly
  - Follow-up implementation / fixture task input

## Goal

Investigate whether current implementation already handles the `TASK-DATA-015-01` / `TASK-DATA-015-02` recursive named model reference boundary.

## Work

- Create or use a minimal fixture with a struct model field referencing the same named model, for example `object_ref.parent: object_ref`.
- Check parser / resolver behavior for recursive named model TypeRef.
- Check validator behavior for recursive named model TypeRef.
- Check renderer / model-file behavior and confirm it does not infinitely expand recursive references.
- Classify the result as already-supported, implementation-gap, or spec-only without runtime impact.
- Identify follow-up implementation and fixture tasks if needed.

## Included Scope

- Runtime behavior investigation.
- Focused fixture or temporary test input if needed.
- Gap classification for resolver / validator / renderer behavior.

## Excluded Scope

- Permanent UC-002 YAML migration.
- Golden regeneration unless a later implementation task selects it.
- Untagged union / general oneOf support.
- Broadening ADR-073 or WORK-DATA-010.

## Done condition

- Current recursive named model reference behavior is classified.
- Required implementation / fixture follow-up is identified, or explicit no-change evidence is recorded.
- No untagged union support is introduced.

## Verification

- Run focused validation / render checks or equivalent repo-local tests.
- Record commands and results in Evidence.

## Evidence

Not started.
