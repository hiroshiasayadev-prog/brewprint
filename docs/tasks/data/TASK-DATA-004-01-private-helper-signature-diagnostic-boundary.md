# TASK-DATA-004-01: Private helper signature diagnostic boundary

- **id**: TASK-DATA-004-01
- **status**: done
- **date**: 2026-05-31
- **work_item**: WORK-DATA-004
- **source_requirement**: REQ-DATA-003
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Diagnostic boundary for disallowed params private helper references
  - Implementation-entry decision for spec alignment

## Goal

Finalize the validation diagnostic boundary for task-file private helper model references in task signatures.

This task records how the accepted REQ-DATA-003 policy becomes a concrete validation rule before spec wording and implementation begin.

## Work

- Review REQ-DATA-003 accepted policy.
- Confirm that `params[].model` references to task-file private helper models are validation errors.
- Confirm that `returns.model` references to task-file private helper models are valid and silent in the minimum scope.
- Decide the diagnostic identifier, severity, target location, and message intent for disallowed params references.
- Decide suppression / non-cascade expectations against unresolved model diagnostics and existing helper model diagnostics.
- Keep render table behavior out of scope except for preserving the render-vs-validation boundary.

## Included Scope

- Task-file private helper model references from task `params[].model`.
- Task-file private helper model references from task `returns.model`.
- Diagnostic boundary and minimum validation behavior.
- Suppression / non-cascade policy for related diagnostics.
- Inputs for TASK-DATA-004-02 spec alignment.

## Excluded Scope

- Spec editing.
- Implementation or tests.
- Model-file helper render / WORK-DATA-003.
- ADR-075 model-file render.
- ADR-073 tagged union.
- ADR-074 DAG asset TypeRef hint.
- MCP helper exposure / semantic identity.
- UC-002 model response helper-shape migration.
- Redesigning `## Private models` render output.
- WORK-DATA-002 reopening.
- M15 / v1.1.0-spec reopening.

## Done Condition

- Diagnostic behavior for disallowed `params[].model` private helper references is selected.
- Allowed `returns.model` private helper references are confirmed to require no diagnostic in the minimum scope.
- Suppression / non-cascade expectations are described.
- TASK-DATA-004-02 has concrete input for spec / diagnostics wording.
- No spec, implementation, fixture, render, or test changes are performed.

## Verification

- Confirm REQ-DATA-003 accepted policy is preserved.
- Confirm WORK-DATA-003 is not pulled into scope.
- Confirm WORK-DATA-002 remains closed.
- Confirm render exposure is not treated as validation success.

## Evidence

### Decision

- Diagnostic id: `invalid_private_model_reference`.
- Severity: `error`.
- Target: the offending `params[].model` field.
- Applies to params on task / branch / fork / join signatures.
- A same-file private helper model referenced from `params[].model` is resolved first, then rejected as an invalid public input contract.
- `returns.model` may reference same-file private helper models and produces no diagnostic in the minimum scope.
- If the referenced helper model cannot be resolved, keep the existing `unresolved_model` behavior.
- If helper model identity is already invalid, such as `duplicate_model_id`, `invalid_private_model_reference` may be suppressed to avoid cascading diagnostics.
- Render behavior is unchanged. `## Private models.used by` may still show the direct reference because render visibility is not validation success.

### Rationale

- The diagnostic id follows the existing diagnostics style and the prior TASK-DATA-002-03 candidate name recorded in `docs/spec/diagnostics.md`.
- Params are caller-provided input contracts. Requiring a file-private helper schema from params would expose a file-local schema outside its visibility boundary.
- Returns are task-produced output contracts. A task-local response schema can be represented by a same-file private helper model without requiring callers to construct it.
- The issue is not unresolved resolution when the helper exists in the same file; it is an invalid use-site under the signature exposure policy.

### Follow-up task input

TASK-DATA-004-02 should update the relevant specs to record:

- `invalid_private_model_reference` as the diagnostic for params references to same-file private helper models.
- severity `error` and target `params[].model`.
- applicability to task / branch / fork / join params.
- `returns.model` private helper references as valid and silent.
- suppression / non-cascade behavior relative to `unresolved_model` and `duplicate_model_id`.
- render output remains structural visibility and is not proof of validation success.

### Verification result

- REQ-DATA-003 accepted policy is preserved.
- WORK-DATA-003 is not pulled into scope.
- WORK-DATA-002 remains closed.
- Render exposure is not treated as validation success.
- No spec, implementation, fixture, render, or test changes were performed.
