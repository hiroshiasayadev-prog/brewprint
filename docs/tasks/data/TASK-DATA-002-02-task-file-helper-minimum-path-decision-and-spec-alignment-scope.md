# TASK-DATA-002-02: Task-file helper minimum spec alignment

- **id**: TASK-DATA-002-02
- **status**: todo
- **date**: 2026-05-31
- **work_item**: WORK-DATA-002
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-002-01
- **outputs**:
  - Spec alignment for Option A task-file helper minimum
  - Implementation-entry notes for TASK-DATA-002-03

## Goal

Align the relevant specs for the selected WORK-DATA-002 Option A boundary: task-file helper model minimum.

This task turns the TASK-DATA-002-01 decision into concrete spec wording before implementation begins.

## Work

- Review TASK-DATA-002-01 evidence.
- Identify the current spec sections that own:
  - model node placement and file-local helper model rules
  - TypeRef resolution and external reference prohibition
  - DAG Markdown render format
  - diagnostics for duplicate helper IDs and invalid helper references
- Update or draft updates for the affected spec files.
- Keep the scope limited to task-file helper model minimum.
- Preserve the WORK-DATA-003 boundary for model-file helper render, ADR-072 catalog follow-up, ADR-075 resolution, and UC-002 model response helper-shape migration.

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
- ADR-075 acceptance / revision / split.
- ADR-073 tagged union model.
- ADR-074 DAG asset TypeRef hint.
- ADR-078 MCP helper model exposure / semantic identity.
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
- The spec alignment explicitly excludes WORK-DATA-003 scope.
- Implementation-entry notes for TASK-DATA-002-03 are recorded.
- No implementation, renderer, YAML, fixture, or test execution is performed.

## Verification

- TASK-DATA-002-01 Option A decision is preserved.
- Model-file helper render remains owned by WORK-DATA-003.
- ADR-072 catalog follow-up remains outside WORK-DATA-002.
- ADR-073 / ADR-074 / ADR-078 are not pulled into scope.
- ADR-075 remains outside WORK-DATA-002.
- No UC-002 model response migration is performed.

## Evidence

Initial evidence:

- TASK-DATA-002-01 selected Option A.
- TASK-DATA-002-01 deferred model-file helper render to WORK-DATA-003.
- ADR-070 defines file-private helper model semantics.
- ADR-071 defines task-file helper model render exposure.
- ADR-072 catalog is opt-in and excluded from Option A.
- ADR-075 remains proposed and is owned by WORK-DATA-003.
