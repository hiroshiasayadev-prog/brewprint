# V01-TASK-DATA-002-02: Task-file helper minimum spec alignment

- **id**: V01-TASK-DATA-002-02
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-DATA-002
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-DATA-002-01
- **outputs**:
  - Spec alignment for Option A task-file helper minimum
  - Implementation-entry notes for V01-TASK-DATA-002-03

## Goal

Align the relevant specs for the selected V01-WORK-DATA-002 Option A boundary: task-file helper model minimum.

This task turns the V01-TASK-DATA-002-01 decision into concrete spec wording before implementation begins.

## Work

- Review V01-TASK-DATA-002-01 evidence.
- Identify the current spec sections that own:
  - model node placement and file-local helper model rules
  - TypeRef resolution and external reference prohibition
  - DAG Markdown render format
  - diagnostics for duplicate helper IDs and invalid helper references
- Update or draft updates for the affected spec files.
- Keep the scope limited to task-file helper model minimum.
- Preserve the V01-WORK-DATA-003 boundary for model-file helper render, V01-ADR-072 catalog follow-up, V01-ADR-075 resolution, and UC-002 model response helper-shape migration.

## Included Scope

- Task files may define file-private helper models for task-local schema shapes.
- Task-file helper models are same-file private and cannot be referenced from other YAML files.
- Same-file TypeRef resolution may resolve helper model IDs.
- Cross-file / cross-module TypeRef references to task-file helper models are invalid.
- DAG Markdown render may include `## Private models` for task files containing helper models.
- Helper models are not rendered as Mermaid DAG nodes.
- Diagnostics should cover duplicate helper IDs and invalid external helper references at the spec level.

## Excluded Scope

- Model-file helper model migration.
- Model-file render exposure.
- Model catalog render changes.
- V01-ADR-075 acceptance / revision / split.
- V01-ADR-073 tagged union model.
- V01-ADR-074 DAG asset TypeRef hint.
- V01-ADR-078 MCP helper model exposure / semantic identity.
- UC-002 model response migration candidates N-014 / N-015 / N-023.
- Implementation, renderer changes, YAML migration, fixture updates, and test execution.

## Candidate Spec Files

- `docs/spec/nodes.md`
- `docs/spec/type-ref.md` if it is the current TypeRef-owning spec section
- `docs/spec/views/dag.md` if present, otherwise the current DAG Markdown render-owning section
- `docs/spec/diagnostics.md` if present, otherwise the current diagnostics-owning section

If no dedicated spec file exists for a responsibility, record the actual owning spec section before editing.

## Done Condition

- The affected spec files / sections are identified.
- Spec wording for task-file helper model placement, visibility, and TypeRef resolution is updated or drafted.
- DAG Markdown render behavior for `## Private models` is updated or drafted.
- Diagnostics expectations are updated or drafted.
- The spec alignment explicitly excludes V01-WORK-DATA-003 scope.
- Implementation-entry notes for V01-TASK-DATA-002-03 are recorded.
- No implementation, renderer, YAML, fixture, or test execution is performed.

## Verification

- V01-TASK-DATA-002-01 Option A decision is preserved.
- Model-file helper render remains owned by V01-WORK-DATA-003.
- V01-ADR-072 catalog follow-up remains outside V01-WORK-DATA-002.
- V01-ADR-073 / V01-ADR-074 / V01-ADR-078 are not pulled into scope.
- V01-ADR-075 remains outside V01-WORK-DATA-002.
- No UC-002 model response migration is performed.

## Evidence

Initial evidence:

- V01-TASK-DATA-002-01 selected Option A.
- V01-TASK-DATA-002-01 deferred model-file helper render to V01-WORK-DATA-003.
- V01-ADR-070 defines file-private helper model semantics.
- V01-ADR-071 defines task-file helper model render exposure.
- V01-ADR-072 catalog is opt-in and excluded from Option A.
- V01-ADR-075 remains proposed and is owned by V01-WORK-DATA-003.

Completion evidence:

- Updated `docs/spec/nodes.md` to define task-file helper model placement, file-private visibility, local identity, same-file reference scope, external QualifiedID prohibition, DAG non-node behavior, and same-module public/helper model name collision.
- Updated `docs/spec/naming.md` to reference the canonical node semantics and define public model / file-private helper model collision rules.
- Updated `docs/spec/type-ref.md` to reference the canonical node semantics, define same-file helper model bare TypeRef resolution before same-module public model fallback, and state that QualifiedID TypeRefs only target public models.
- Updated `docs/spec/views/dag.md` to reference the canonical node semantics, define the task-file render `## Private models` section, and keep helper models out of the Mermaid DAG body.
- Updated `docs/spec/views/er.md` to state that task-file helper models do not appear in ER renders.
- Updated `docs/spec/diagnostics.md` to add `duplicate_model_id` for helper/public model collision and to leave any dedicated external-helper-reference diagnostic decision to V01-TASK-DATA-002-03; the existing unresolved TypeRef diagnostics remain acceptable for the minimum.
- Deferred V01-ADR-072 model/schema catalog, V01-ADR-075 model file render, model-file helper render exposure, and UC-002 helper-shape migration to V01-WORK-DATA-003.
- No implementation, renderer code, UC-002 YAML, fixture, or golden output was changed.

## Implementation-entry Notes for V01-TASK-DATA-002-03

- Parser / resolver should accept `type: model` helper nodes in task files as file-private local schema definitions.
- Same-file TypeRef resolution should search task-file helper models before same-module public models, while validation rejects public/helper same-name collisions within the module.
- QualifiedID TypeRefs should not resolve to task-file helper models.
- DAG renderer should render helper models only in Markdown `## Private models` and should not add them to Mermaid DAG nodes.
- Validation should implement `duplicate_model_id`; a dedicated invalid external helper reference diagnostic is optional and should be decided during implementation if existing `unresolved_model` / `unresolved_field_type` messages are too ambiguous.
