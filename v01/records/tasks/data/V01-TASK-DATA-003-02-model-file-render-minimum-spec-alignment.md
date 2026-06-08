# V01-TASK-DATA-003-02: Model-file render minimum spec alignment

- **id**: V01-TASK-DATA-003-02
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-DATA-003
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-DATA-003-01
- **outputs**:
  - V01-ADR-075 decision-basis confirmation without ADR text change
  - Model-file render minimum spec alignment
  - Implementation-entry notes for model-file render minimum

## Goal

Align the relevant specs for the model-file render minimum path under V01-WORK-DATA-003, using V01-ADR-075 as the decision basis without editing V01-ADR-075 in this task.

This task turns the V01-TASK-DATA-003-01 decision into concrete documentation: V01-ADR-075 is usable as the model-file render decision basis, while tagged union rendering is excluded from the current V01-WORK-DATA-003 execution scope and remains deferred to V01-ADR-073 or a separate tagged-union work item.

## Work

- Review V01-TASK-DATA-003-01 evidence.
- Use V01-ADR-075 as the model-file render decision basis without changing V01-ADR-075 text or status in this task.
- Identify whether a dedicated model-file render spec should be created, likely `docs/spec/views/model-file.md`.
- Draft or update spec wording for model-file render minimum:
  - one model YAML file produces one Markdown render
  - public main model section
  - private helper models section
  - struct fields table
  - enum values display
  - list element display
  - dict value display
  - output placement and index participation
- Update related spec references if needed:
  - `docs/spec/project-layout.md`
  - `docs/spec/nodes.md`
  - `docs/spec/file-types.md` if render output/type catalog needs mention
- Explicitly defer tagged union rendering.
- Produce implementation-entry notes for the next task.

## Included Scope

- V01-ADR-075 decision-basis confirmation without ADR text change.
- Model-file render minimum spec creation or update.
- Tagged union rendering is excluded from V01-WORK-DATA-003 current execution scope and deferred to a later tagged-union work item.
- Struct / enum / list / dict model display rules.
- File-private helper model table for model files.
- Output placement for model file renders.
- Spec-level implementation entry notes.

## Excluded Scope

- Tagged union discriminator / variants rendering.
- Tagged union model validation or YAML migration.
- DAG asset TypeRef hint.
- MCP helper model exposure / semantic identity.
- Model catalog implementation.
- UC-002 model response helper migration.
- Renderer implementation, fixture / golden update, and test execution.
- V01-WORK-DATA-002 task-file helper minimum reopening.
- V01-REQ-DATA-003 params / returns signature exposure policy.

## Done Condition

- V01-ADR-075 is recorded as the decision basis for this task without changing V01-ADR-075 text or status.
- The spec owner for model-file render minimum is created or identified.
- Model-file render minimum output format is defined for struct / enum / list / dict.
- Private helper model table behavior is defined for model files.
- Output placement / index participation is defined or explicitly deferred with rationale.
- Tagged union rendering is explicitly deferred.
- Implementation-entry notes for V01-TASK-DATA-003-03 are recorded.
- No renderer implementation, YAML migration, fixture update, or test execution is performed.

## Verification

- Confirm V01-TASK-DATA-003-01 decision is preserved.
- Confirm V01-WORK-DATA-003 current execution scope does not force V01-ADR-073 implementation.
- Confirm V01-ADR-070 human-visible helper constraint is satisfied for model-file helpers.
- Confirm V01-ADR-072 model catalog remains opt-in and not a substitute for model-file render.
- Confirm V01-WORK-DATA-002 remains closed and not reopened.

## Implementation-entry notes

Next implementation task should implement only the model-file render minimum defined in `docs/spec/views/model-file.md` and `docs/spec/project-layout.md`:

- generate one Markdown render for each `model/*.yaml` file
- place output as `renders/{group-id}/model-{model-id}.md`
- include model renders in master and group indexes as kind `Model`
- render public main model plus same-file private helper models
- support `struct`, `enum`, `list`, and `dict`
- do not implement tagged union discriminator / variants rendering in this path
- do not implement renderer changes for DAG TypeRef hint, MCP helper exposure, or UC-002 migration in this path

## Evidence

Final evidence:

- `docs/spec/views/model-file.md` exists and owns the model-file render minimum.
- `docs/spec/views/model-file.md` records V01-ADR-075 as decision basis without editing V01-ADR-075 text or status.
- `docs/spec/views/model-file.md` defines one model YAML file to one Markdown render.
- `docs/spec/views/model-file.md` defines public main model and private helper model sections.
- `docs/spec/views/model-file.md` defines minimum display rules for `struct`, `enum`, `list`, and `dict`.
- `docs/spec/views/model-file.md` explicitly excludes tagged union rendering, tagged union validation, model catalog implementation, DAG TypeRef hint, MCP helper model exposure, and UC-002 model response helper migration.
- `docs/spec/project-layout.md` now defines model-file render placement as `renders/{group-id}/model-{model-id}.md`, includes model renders in master index counts, and includes model renders in group index rows as kind `Model`.
- `docs/spec/nodes.md` now points model-file helper/render behavior to `docs/spec/views/model-file.md` while keeping model catalog and UC-002 migration out of the task-file helper minimum.
- No renderer implementation, YAML migration, fixture / golden update, or test execution was performed in this task.
