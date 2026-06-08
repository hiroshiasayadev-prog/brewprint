# V01-REQ-DATA-003: Private helper model signature exposure boundary

- **id**: V01-REQ-DATA-003
- **status**: accepted
- **date**: 2026-05-31
- **source_refs**:
  - V01-REQ-DATA-002
  - V01-WORK-DATA-002
  - V01-TASK-DATA-002-03
  - V01-ADR-070
  - V01-ADR-071
- **work_items**:
  - V01-WORK-DATA-004

## Requirement

The project needs a separate policy for file-private helper models in task signatures, because task `params` and `returns` have different contract directions.

During V01-TASK-DATA-002-03 discussion and the follow-up boundary review, the following direction emerged:

- file-private helper model in task `params`: validation error, because callers must construct the input contract and cannot reference file-local schemas from another YAML file.
- file-private helper model in task `returns`: allowed, because the task owns the emitted output shape and a task-local response schema is a useful minimum helper-model capability.

This requirement is resolved through `V01-WORK-DATA-004` without reopening the already-completed V01-WORK-DATA-002 Option A task-file helper minimum.

## Evidence

TASK-DATA-002 Option A allows task-file helper models and same-file TypeRef resolution. However, task signatures have two different contract directions:

- `params` describe values the caller must provide.
- `returns` describe values the task produces.

Treating both directions identically may create an unstable contract. A private helper model in `params` may force external callers to depend on a file-local schema, while a private helper model in `returns` can still be useful as a task-local response schema.

The current DAG Markdown spec already defines `## Private models.used by` as direct reference inventory. That render surface should remain separate from validation success so that invalid structural references are still visible to humans and LLMs. See `docs/spec/views/dag.md` for the current render surface.

This requirement does not require changing DAG render output unless wording clarification is needed.

## Required Outcome

- Preserve the accepted asymmetric policy for file-private helper model references in `params[].model` and `returns.model`.
- Define the concrete validation behavior for disallowed params references.
- Keep allowed returns references silent in the minimum scope.
- Keep render exposure separate from validation success.
- Preserve `## Private models.used by` as direct TypeRef reference inventory, not as proof that the reference is valid as a public contract.

## Explicitly Excluded Scope

- Model-file helper render boundary owned by V01-WORK-DATA-003.
- V01-ADR-073 tagged union model.
- V01-ADR-074 DAG asset TypeRef hint.
- V01-ADR-078 MCP helper exposure / semantic identity.
- UC-002 model response helper-shape migration.
- Remaining UC-002 notes retreat debt.
- M15 / v1.1.0-spec reopening.

## Boundary

This requirement captures the signature exposure policy for task-file helper models and is resolved by `V01-WORK-DATA-004`. It does not reopen `V01-WORK-DATA-002`, decide model-file helper render, or block the already-completed task-file helper minimum.

The accepted policy is asymmetric. `V01-WORK-DATA-004` finalizes the concrete spec wording, diagnostic behavior, implementation, and tests:

- `params[].model` private helper reference is a validation error because params are caller-provided input contracts.
- `returns.model` private helper reference is allowed because returns are task-produced output contracts and can use task-local response schemas.

Render exposure is best-effort structural visibility, not proof of validation success. `used by` may show direct references even when a validation rule classifies the reference as invalid.
