# V01-TASK-DATA-004-03: Private helper signature validation implementation

- **id**: V01-TASK-DATA-004-03
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-DATA-004
- **source_requirement**: V01-REQ-DATA-003
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-DATA-004-02
- **outputs**:
  - Validation implementation for task-file private helper signature exposure policy
  - Regression tests for params rejection and returns allowance
  - Verification evidence for V01-WORK-DATA-004 close

## Goal

Implement and verify the task-file private helper signature exposure policy defined by V01-REQ-DATA-003 and aligned by V01-TASK-DATA-004-02.

## Work

- Review V01-TASK-DATA-004-02 spec / diagnostic evidence.
- Implement validation for task-file private helper model references from `params[].model`.
- Preserve valid and silent behavior for task-file private helper model references from `returns.model`.
- Add or update focused tests for:
  - params private helper reference rejection
  - returns private helper reference allowance
  - diagnostic target / message behavior
  - suppression / non-cascade against unresolved model or duplicate helper diagnostics
- Confirm existing task-file helper model render behavior remains unchanged.
- Record verification evidence.

## Included Scope

- Validation implementation for task-file private helper signature exposure.
- Focused resolver / validation tests.
- Regression check that returns references remain valid.
- Regression check that render shape is not redesigned.
- V01-WORK-DATA-004 close evidence if all completion conditions are met.

## Excluded Scope

- Model-file helper render / V01-WORK-DATA-003.
- V01-ADR-075 model-file render.
- V01-ADR-073 tagged union.
- V01-ADR-074 DAG asset TypeRef hint.
- MCP helper exposure / semantic identity.
- UC-002 model response helper-shape migration.
- Rendering implementation changes unless required to keep existing behavior passing.
- V01-WORK-DATA-002 reopening.
- M15 / v1.1.0-spec reopening.

## Done Condition

- `params[].model` references to task-file private helper models produce the specified validation diagnostic.
- `returns.model` references to task-file private helper models remain valid and produce no diagnostic in the minimum scope.
- Tests cover positive and negative cases plus diagnostic suppression / non-cascade behavior.
- Existing task-file helper render behavior is not redesigned.
- Verification evidence is recorded.

## Verification

- Run focused validation tests for the changed behavior.
- Run related task-file helper model regression tests.
- Run broader tests as needed before closing V01-WORK-DATA-004.
- Confirm Design Records validation still accepts the updated task / work item artifacts.

## Evidence

### Implementation

- Added `invalid_private_model_reference` validation for task / branch / fork / join `params[].model` TypeRefs that resolve to same-file task-file private helper models.
- Kept same-file private helper references from `returns.model` valid and silent in the minimum scope.
- Preserved `unresolved_model` for genuinely unresolved model refs.
- Suppressed `invalid_private_model_reference` when helper identity is already invalid via `duplicate_model_id`.
- Kept DAG Markdown render implementation unchanged; only tests were adjusted to account for the new validation diagnostic.

### Regression tests

- `TestTaskFilePrivateHelperParamModelIsRejected`
- `TestTaskFilePrivateHelperParamModelIsRejectedForControlNodes`
- `TestTaskFilePrivateHelperReturnModelIsAccepted`
- `TestTaskFilePrivateHelperUnresolvedParamStillUsesUnresolvedModel`
- Existing duplicate helper identity tests now assert no cascading `invalid_private_model_reference`.

### Verification

- `go test ./internal/resolve -run "TestTaskFileHelper|TestTaskFilePrivateHelper|TestQualifiedTypeRefCannotReferenceTaskFileHelperModel|TestParamListMissingModelUsesUnresolvedModel"`: pass
- `go test ./internal/resolve`: pass
- `go test ./internal/render/dag ./internal/render/model ./internal/render/er`: pass
- `go test ./internal/query`: pass
- `go test ./internal/resolve ./internal/render/dag ./internal/query`: pass
- `go test ./...`: pass
