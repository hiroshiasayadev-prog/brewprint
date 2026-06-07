# TASK-DATA-015-06: Review spec and task split for recursive reference boundary

- **id**: TASK-DATA-015-06
- **status**: not_started
- **date**: 2026-06-07
- **work_item**: WORK-DATA-015
- **source_requirement**: REQ-DATA-008
- **estimate**: 0.5d
- **depends_on**:
  - TASK-DATA-015-02
- **outputs**:
  - Independent review result for TypeRef spec update and task split
  - Required revision list or PASS evidence

## Goal

Review the TypeRef spec update and the WORK-DATA-015 follow-up task split before implementation or UC-002 cleanup proceeds.

## Work

- Review the `docs/spec/type-ref.md` changes from TASK-DATA-015-02.
- Check consistency with ADR-073 and TASK-DATA-015-01.
- Check that untagged union / general `oneOf` is not silently introduced.
- Check that recursive named model reference support is bounded and implementable.
- Check that TASK-DATA-015-03 through TASK-DATA-015-05 form a coherent follow-up path.

## Included Scope

- Documentation / planning review.
- Required revision identification.

## Excluded Scope

- Implementation changes.
- UC-002 YAML migration.
- Golden regeneration.

## Done condition

- Review result is recorded as PASS or Needs revision.
- Any required revisions are assigned to a concrete task or applied before downstream work proceeds.

## Verification

- Validate affected records after review updates.

## Evidence

Not started.
