# V01-TASK-DATA-015-06: Review spec and task split for recursive reference boundary

- **id**: V01-TASK-DATA-015-06
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-DATA-015
- **source_requirement**: V01-REQ-DATA-008
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-DATA-015-02
- **outputs**:
  - Independent review result for TypeRef spec update and task split
  - Required revision list or PASS evidence

## Goal

Review the TypeRef spec update and the V01-WORK-DATA-015 follow-up task split before implementation or UC-002 cleanup proceeds.

## Work

- Review the `docs/spec/type-ref.md` changes from V01-TASK-DATA-015-02.
- Check consistency with V01-ADR-073 and V01-TASK-DATA-015-01.
- Check that untagged union / general `oneOf` is not silently introduced.
- Check that recursive named model reference support is bounded and implementable.
- Check that V01-TASK-DATA-015-03 through V01-TASK-DATA-015-05 form a coherent follow-up path.

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

Completed on 2026-06-07.

### Review result

Verdict: PASS after required cleanup.

Codex review first returned `Needs revision` because accidental placeholder / discarded artifacts remained, unrelated dirty files were present, and V01-WORK-DATA-015 task graph / V01-TASK-DATA-015-04 dependency wiring needed cleanup.

Required V01-WORK-DATA-015 fixes were applied:

- `V01-WORK-DATA-015.tasks` now lists `V01-TASK-DATA-015-01` through `V01-TASK-DATA-015-06`.
- `V01-WORK-DATA-015.status` was changed from `blocked` to `in_progress`.
- `V01-WORK-DATA-015` Task Flow was updated to the actual task split.
- `V01-TASK-DATA-015-04.depends_on` now includes `V01-TASK-DATA-015-06`, so the review gates UC-002 cleanup.
- Accidental placeholder / discarded artifacts were removed before commit.

### Non-issues confirmed by review

- `docs/spec/type-ref.md` allows recursive named model references.
- `docs/spec/type-ref.md` rejects inline recursive shapes.
- `docs/spec/type-ref.md` rejects inline `union<...>`, `oneOf<...>`, `anyOf<...>`, and scalar union syntax.
- V01-ADR-073 was not silently broadened.
- Untagged-like machine-readable surfaces are routed to tagged union envelope models.
- `V01-TASK-DATA-015-02` evidence is spec-only and does not claim implementation, YAML migration, fixture, or golden changes.

### Validation

- `validate_records` for `V01-TASK-DATA-015-01` through `V01-TASK-DATA-015-06`: PASS.
- `validate_records` for `V01-WORK-DATA-015`: PASS.
- Repository-wide spec validation was reported as failing only on unrelated pre-existing namespace-model `missing_section_target` diagnostics, not on the TypeRef update.
- `go test` was not run because the reviewed scope changed only docs/spec/task/work-item files.

### Follow-up

`V01-TASK-DATA-015-03` should be delegated to Codex because it requires repo-local runtime behavior investigation, including focused validate / render checks for recursive named model references.
