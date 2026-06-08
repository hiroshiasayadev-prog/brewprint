# V01-TASK-DATA-003-01: V01-ADR-075 dependency and split review

- **id**: V01-TASK-DATA-003-01
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-DATA-003
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - V01-ADR-075 dependency review
  - Decision input for model-file helper render boundary

## Goal

Decide whether V01-ADR-075 can be used as the model-file helper render boundary for V01-WORK-DATA-003, or whether it must be revised or split before implementation.

This task exists because V01-ADR-075 is still proposed and depends on V01-ADR-073, while V01-WORK-DATA-003 needs a model-file helper render path that may not intrinsically require tagged union support.

## Work

- Review V01-ADR-075 model file render decision and dependency metadata.
- Review V01-ADR-073 tagged union scope only enough to determine whether V01-ADR-075's dependency is intrinsic.
- Review V01-ADR-070 / V01-ADR-071 / V01-ADR-072 responsibilities as context:
  - V01-ADR-070: helper models must not become invisible in human-facing render.
  - V01-ADR-071: task-file helper render exposure.
  - V01-ADR-072: opt-in model catalog, not automatic model-file render.
- Decide whether V01-ADR-075 should be:
  - accepted as-is,
  - revised before acceptance,
  - split into a smaller model-file-render-minimum decision plus later tagged-union rendering detail,
  - or deferred.
- Identify the minimum model-file render capability required before UC-002 model response helper migration can proceed.
- Keep implementation, YAML migration, fixture updates, and renderer changes out of this task.

## Decision Questions

1. Is tagged union rendering intrinsic to model-file render, or only one kind-specific rendering extension?
2. Can model-file render minimum be defined for struct / enum / list / dict helper models without V01-ADR-073?
3. Should V01-ADR-075 be accepted with a narrowed minimum scope, or should a new ADR supersede/split it?
4. What is the earliest safe boundary for migrating UC-002 model response helper candidates such as N-014 / N-015 / N-023?

## Included Scope

- V01-ADR-075 status / dependency review.
- V01-ADR-073 dependency classification.
- Model-file helper render minimum boundary.
- Follow-up recommendation for ADR / spec / implementation tasks under V01-WORK-DATA-003.

## Excluded Scope

- Implementing model-file render.
- Implementing tagged union model.
- Implementing DAG TypeRef hint.
- Implementing MCP helper model exposure / semantic identity.
- UC-002 YAML migration.
- Fixture / golden update.
- M15 / v1.1.0-spec reopening.
- V01-WORK-DATA-002 task-file helper minimum, which is already closed.
- V01-REQ-DATA-003 params / returns signature exposure policy.

## Done Condition

- V01-ADR-075 dependency on V01-ADR-073 is classified as intrinsic, non-intrinsic, or unresolved.
- One of the following paths is selected:
  - accept V01-ADR-075 as-is,
  - revise V01-ADR-075 before acceptance,
  - split / supersede V01-ADR-075 with a smaller model-file-render-minimum ADR,
  - defer model-file render.
- Minimum model-file render capability needed for helper model visibility is described.
- Impact on UC-002 model response helper migration candidates is recorded.
- Follow-up task input for V01-WORK-DATA-003 is concrete enough to create the next task.

## Verification

- Confirm V01-ADR-070 human-visible helper constraint is preserved.
- Confirm V01-ADR-072 catalog is not treated as a substitute for automatic model-file render.
- Confirm V01-ADR-073 is not implemented or implicitly accepted by this task.
- Confirm no spec, implementation, renderer, YAML, fixture, or test changes are performed.
- Confirm V01-WORK-DATA-002 closed boundary is not reopened.

## Evidence

### Decision

V01-ADR-075 can proceed as the model-file render decision for V01-WORK-DATA-003, but its initial implementation boundary must exclude tagged union rendering.

V01-ADR-075's dependency on V01-ADR-073 is classified as non-intrinsic for the model-file render minimum. Tagged union rendering is a kind-specific rendering extension, not a prerequisite for rendering struct / enum / list / dict model files and their private helper models.

Selected path: accept / use V01-ADR-075 for model-file render minimum, with tagged union display deferred to V01-ADR-073 or a separate tagged-union work item.

### Rationale

- V01-ADR-075's core decision is model file render: one model YAML file produces one human-readable Markdown render.
- The minimum render can cover public main model and private helper models for struct / enum / list / dict without tagged union support.
- V01-ADR-070 requires helper models not to be invisible to human-facing render; V01-ADR-075 provides the automatic file-local render path for model files.
- V01-ADR-072 model catalog is an opt-in curated view and does not replace automatic model-file render.
- V01-ADR-073 tagged union model has different implementation and validation complexity and should not block model-file render minimum.

### Minimum model-file render capability

The V01-WORK-DATA-003 minimum should cover:

- public main model section
- private helper models section
- struct fields table
- enum values display
- list element display
- dict value display
- output placement for model file render

The V01-WORK-DATA-003 minimum should defer:

- tagged union discriminator / variants rendering
- tagged union validation
- tagged union UC migration
- DAG TypeRef hint
- MCP helper model exposure / semantic identity

### UC-002 migration impact

UC-002 model response helper candidates can proceed after model-file render minimum is specified and implemented, provided they do not require tagged union semantics.

Candidate interpretation:

- N-014 `get_reference_tree_response.nodes`: model-file helper candidate; not intrinsically tagged union.
- N-015 `get_reference_tree_response.edges`: model-file helper candidate; not intrinsically tagged union, but existing `reference` model reuse needs care.
- N-023 `get_source_response.snippet`: small model-file helper candidate; not intrinsically tagged union.

### Follow-up task input

Create the next V01-WORK-DATA-003 task to align V01-ADR-075 / spec scope for model-file render minimum:

`V01-TASK-DATA-003-02: Model-file render minimum spec alignment`

That task should update or draft spec changes for model-file render covering struct / enum / list / dict and private helper model table, while explicitly deferring tagged union rendering to V01-ADR-073 or a separate tagged-union work item.

### Verification result

- V01-ADR-070 human-visible helper constraint is preserved.
- V01-ADR-072 catalog is not treated as a substitute for automatic model-file render.
- V01-ADR-073 is not implemented or implicitly accepted by this task.
- No spec, implementation, renderer, YAML, fixture, or test changes were performed.
- V01-WORK-DATA-002 closed boundary is not reopened.
