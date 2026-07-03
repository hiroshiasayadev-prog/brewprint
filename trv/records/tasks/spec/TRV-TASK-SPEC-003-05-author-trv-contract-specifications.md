# TRV-TASK-SPEC-003-05: Author TRV contract Specifications

- **id**: TRV-TASK-SPEC-003-05
- **status**: blocked
- **date**: 2026-07-03
- **work_item**: TRV-WORK-SPEC-003
- **task_type**: authoring
- **estimate**: 1d
- **depends_on**:
  - TRV-TASK-SPEC-003-04
- **outputs**:
  - TRV-TASK-SPEC-003-05
  - spec:trv
  - spec:trv.mcp_interface
  - spec:trv.task_input
  - spec:trv.caller_integration
  - spec:trv.compatibility

## Goal

Project accepted TRV contract decisions into normative Specifications.

## Work

- Create MCP interface, Task input, caller integration, and compatibility Specifications.
- Register them under `spec:trv`.
- Preserve PRODUCT-ADR-SPEC-017 and closed W002 architecture.
- Leave exact implementation schemas and mechanics to W004.

This Task must not change accepted decisions, write detailed design, review its own work, or start implementation.

## Done condition

- Four contract Specifications exist and are registered.
- ADR and Specification content agree.
- PRODUCT semantics and W002 architecture remain unchanged.

## Verification

- Confirm IDs, paths, parent refs, required sections, and Topics rows.
- Confirm every contract decision has one normative projection.
- Confirm only declared outputs changed.

## Evidence

- T04 completed the routed ADR authoring boundary.
- T03 materialized this writer before the architecture-derived Specification placement gap was discovered.
- The user retired W003 and moved contract Specification design to TRV-WORK-SPEC-005.
- This Task must not execute because its four-file output set was selected before the required topic-tree and Markdown-placement decision.
- Replacement owner: TRV-WORK-SPEC-005.
- Result: `BLOCKED` by the retired W003 route.
