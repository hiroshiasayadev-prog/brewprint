# TASK-DATA-015-04: Clean up UC-002 recursive and untagged-like surfaces

- **id**: TASK-DATA-015-04
- **status**: not_started
- **date**: 2026-06-07
- **work_item**: WORK-DATA-015
- **source_requirement**: REQ-DATA-008
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-015-03
  - TASK-DATA-015-06
- **outputs**:
  - UC-002 cleanup decision for N-044 recursive `object_ref.parent`
  - UC-002 cleanup decision for N-009 untagged-like related list
  - Fixture / render update input if migration is selected

## Goal

Apply the accepted recursive / untagged-union boundary to UC-002 surfaces after runtime behavior is understood.

## Work

- Review N-044 and decide whether `object_ref.parent` should be represented as a recursive named model reference in UC-002 YAML.
- Review N-009 and decide whether the untagged-like related list remains `any + note` / prose or becomes a tagged union envelope model.
- Keep untagged union / general `oneOf` unsupported.
- Identify any fixture / render regeneration required if YAML migration is selected.

## Included Scope

- UC-002 cleanup decision and possible YAML migration input.
- Tagged union envelope decision for untagged-like surfaces.
- Traceability to `TASK-DATA-015-01` and `TASK-DATA-015-02`.

## Excluded Scope

- Introducing untagged union / general `oneOf`.
- Broadening ADR-073 or WORK-DATA-010.
- Runtime implementation work.
- Golden regeneration unless selected as a later task.

## Done condition

- N-044 is either migrated, explicitly deferred, or marked no-action with rationale.
- N-009 is either modeled via tagged union envelope, explicitly left opaque, deferred, or marked no-action with rationale.
- Follow-up fixture / render tasks are identified if needed.

## Verification

- Validate affected records after update.
- If YAML is changed, run focused validate / render checks in the later implementation or cleanup task.

## Evidence

Not started.
