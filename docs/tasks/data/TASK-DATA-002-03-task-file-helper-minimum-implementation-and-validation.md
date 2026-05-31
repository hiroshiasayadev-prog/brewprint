# TASK-DATA-002-03: Task-file helper minimum implementation and validation

- **id**: TASK-DATA-002-03
- **status**: todo
- **date**: 2026-05-31
- **work_item**: WORK-DATA-002
- **source_requirement**: REQ-DATA-002
- **estimate**: 1d-2d
- **depends_on**:
  - TASK-DATA-002-02
- **outputs**:
  - Task-file helper model implementation
  - Task-file helper TypeRef validation
  - DAG Markdown `## Private models` render exposure
  - Verification evidence for the Option A minimum

## Goal

Implement and verify the selected WORK-DATA-002 Option A boundary: task-file helper model minimum.

This task starts only after TASK-DATA-002-02 has aligned the relevant specs.

## Work

- Implement task-file helper model parsing / loading for file-private schema definitions.
- Implement same-file TypeRef resolution for task-file helper models.
- Reject cross-file / cross-module references to task-file helper models.
- Validate duplicate local helper IDs within a task file.
- Validate helper/public name conflicts to the extent required by the TASK-DATA-002-02 spec alignment.
- Implement DAG Markdown `## Private models` render exposure for task files containing helper models.
- Keep helper models out of Mermaid DAG flow nodes.
- Add or update focused tests / fixtures / golden evidence for the Option A minimum.

## Included Scope

- Task-file helper model parsing and validation.
- Same-file helper TypeRef resolution.
- Invalid external helper reference diagnostics.
- DAG Markdown private model render exposure for task files.
- Verification evidence for parser / resolver / diagnostics / render behavior.

## Excluded Scope

- Model-file helper model migration.
- Model-file render exposure.
- Model catalog render changes.
- ADR-075 acceptance / revision / split.
- ADR-073 tagged union model.
- ADR-074 DAG asset TypeRef hint.
- ADR-078 MCP helper model exposure / semantic identity.
- UC-002 model response migration candidates N-014 / N-015 / N-023.
- M15 / `v1.1.0-spec` reopening.

## Done Condition

- Task-file helper models can be represented according to the aligned specs.
- Same-file TypeRef resolution works for task-file helper models.
- External references to task-file helper models are rejected.
- Duplicate helper IDs and required helper/public conflicts are diagnosed.
- DAG Markdown renders `## Private models` for task-file helper models.
- Helper models are not emitted as Mermaid DAG flow nodes.
- Focused tests / fixtures / golden evidence cover the implemented behavior.
- WORK-DATA-003 scope remains untouched.

## Verification

- Confirm TASK-DATA-002-02 spec alignment is the implementation source.
- Confirm model-file helper render remains owned by WORK-DATA-003.
- Confirm ADR-072 / ADR-073 / ADR-074 / ADR-075 / ADR-078 are not pulled into WORK-DATA-002 implementation.
- Confirm no UC-002 model response YAML migration is performed.
- Run the focused validation / render test suite identified during implementation.

## Evidence

Initial evidence:

- TASK-DATA-002-01 selected Option A.
- TASK-DATA-002-02 owns the required spec alignment.
- ADR-070 defines file-private helper model semantics.
- ADR-071 defines task-file helper model render exposure.
- WORK-DATA-003 owns model-file helper render and UC-002 model response migration.
