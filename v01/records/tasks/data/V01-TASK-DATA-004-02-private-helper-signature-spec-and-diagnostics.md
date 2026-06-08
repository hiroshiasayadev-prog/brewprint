# V01-TASK-DATA-004-02: Private helper signature spec and diagnostics alignment

- **id**: V01-TASK-DATA-004-02
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-DATA-004
- **source_requirement**: V01-REQ-DATA-003
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-DATA-004-01
- **outputs**:
  - Spec alignment for task-file private helper signature exposure policy
  - Diagnostic specification for disallowed params references
  - Implementation-entry notes for V01-TASK-DATA-004-03

## Goal

Reflect the V01-TASK-DATA-004-01 diagnostic boundary in the relevant specs before implementation.

This task updates only the documentation needed for task-file signature exposure policy. It must not redesign render output or pull in model-file render work.

## Work

- Review V01-TASK-DATA-004-01 evidence.
- Identify the current spec sections that own:
  - task-file private helper model semantics
  - same-file TypeRef resolution
  - task params / returns signature rules
  - diagnostics
  - DAG Markdown render `## Private models` wording if clarification is needed
- Update or draft spec wording for:
  - `params[].model` private helper references as validation errors
  - `returns.model` private helper references as valid and silent
  - diagnostic identifier / severity / target for disallowed params references
  - suppression / non-cascade behavior
  - render exposure as structural visibility rather than validation success, only if needed
- Produce implementation-entry notes for V01-TASK-DATA-004-03.

## Included Scope

- Spec wording for task-file private helper model signature exposure.
- Diagnostic specification for disallowed params references.
- TypeRef resolution boundary clarification where needed.
- Minimal render-vs-validation wording if needed to avoid misreading `used by` as validity.
- Implementation-entry notes.

## Excluded Scope

- Renderer implementation or render table redesign.
- Model-file helper render / V01-WORK-DATA-003.
- V01-ADR-075 model-file render.
- V01-ADR-072 model / schema catalog.
- V01-ADR-073 tagged union.
- V01-ADR-074 DAG asset TypeRef hint.
- MCP helper exposure / semantic identity.
- UC-002 model response helper-shape migration.
- Validation implementation or test execution.
- V01-WORK-DATA-002 reopening.
- M15 / v1.1.0-spec reopening.

## Candidate Spec Files

- `docs/spec/nodes.md`
- `docs/spec/type-ref.md`
- `docs/spec/diagnostics.md`
- `docs/spec/views/dag.md` only if a wording clarification is needed; do not redesign the table.

If a responsibility is owned elsewhere, record the actual owning spec before editing.

## Done Condition

- The affected spec files / sections are identified.
- Params rejection and returns allowance are reflected in spec wording.
- Diagnostic behavior for disallowed params references is reflected in diagnostics spec.
- Suppression / non-cascade expectations are documented.
- Render-vs-validation boundary is clarified only if needed and without changing render shape.
- Implementation-entry notes for V01-TASK-DATA-004-03 are recorded.
- No implementation, fixture, render, or test execution is performed.

## Verification

- Confirm V01-TASK-DATA-004-01 decision is preserved.
- Confirm V01-WORK-DATA-003 model-file render remains excluded.
- Confirm V01-WORK-DATA-002 remains closed.
- Confirm the existing `## Private models.used by` render table is not redesigned.

## Evidence

### Spec owners

Affected spec sections:

- `docs/spec/diagnostics.md`: diagnostic code table and semantic reference / model validation notes.
- `docs/spec/nodes.md`: task signature exposure policy for task-file private helper models.
- `docs/spec/type-ref.md`: TypeRef resolution versus use-site validation boundary.

`docs/spec/views/dag.md` was not changed because `## Private models.used by` already exists as direct reference inventory and no render table redesign is needed for this task.

### Changes

- Added `invalid_private_model_reference` as an error diagnostic for task / branch / fork / join `params[].model` references to same-file private helper models.
- Documented that same-file helper resolution is not an unresolved reference when the helper exists; the use-site is rejected by signature exposure policy.
- Documented that `returns.model` may reference same-file private helper models and remains silent in the minimum scope.
- Documented suppression / non-cascade behavior: do not add `unresolved_model` for a resolved-but-forbidden params reference; suppress `invalid_private_model_reference` when helper identity is already invalid such as `duplicate_model_id`.
- Added task signature exposure policy wording to `docs/spec/nodes.md` without changing render behavior.
- Clarified in `docs/spec/type-ref.md` that TypeRef resolution success does not imply every use-site is valid.

### Implementation-entry notes for V01-TASK-DATA-004-03

V01-TASK-DATA-004-03 should implement validation after TypeRef resolution:

- Detect same-file private helper model references in task / branch / fork / join `params[].model`.
- Emit `invalid_private_model_reference` with severity `error` at the offending `params[].model` location.
- Keep `returns.model` private helper references valid and diagnostic-free.
- Preserve existing `unresolved_model` behavior for genuinely unresolved references.
- Avoid diagnostic cascades when helper identity is already invalid, such as `duplicate_model_id`.
- Keep DAG Markdown `## Private models` render behavior unchanged.

### Verification result

- V01-TASK-DATA-004-01 decision is preserved.
- V01-WORK-DATA-003 model-file render remains excluded.
- V01-WORK-DATA-002 remains closed.
- Existing `## Private models.used by` render table is not redesigned.
- No implementation, fixture, render, or test execution was performed.
