# TASK-DATA-014-01: Decide selector support matrix and object-dependent vocabulary contract boundary

- **id**: TASK-DATA-014-01
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-DATA-014
- **source_requirement**: REQ-DATA-007
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Selector support matrix contract boundary decision
  - Object-dependent kind vocabulary ownership decision
  - Follow-up split recommendation for spec, YAML cleanup, implementation verification, and close

## Goal

Decide the contract boundary for `REQ-DATA-007` / `WORK-DATA-014` before any spec, YAML, fixture, or implementation change is introduced.

This task determines where selector support matrices and object-dependent `kind` vocabulary are represented and validated.

## Work

- Review the WORK-DATA-014 source bucket: N-020, N-031, N-037, N-040, and N-042.
- Compare the bucket with the existing MCP schema surfaces: `Object selector`, `Selector support matrix`, and `ObjectRef`.
- Decide whether the contract belongs in DATA DSL expressiveness, MCP schema/tool contracts, YAML notes, implementation-only behavior, or a new ADR.
- Identify the minimal follow-up task split after the boundary decision.
- Keep tagged union, recursive/untagged union, numeric/default behavior, DAG TypeRef hint, and MCP identity work out of this task.

## Done condition

- The selected contract owner is recorded in Evidence.
- Each source candidate N-020, N-031, N-037, N-040, and N-042 has an ownership decision.
- The next follow-up tasks are identified without performing spec, YAML, fixture, or implementation edits in this task.

## Verification

- Confirm the decision aligns with `REQ-DATA-007` and `WORK-DATA-014`.
- Confirm the decision does not reopen `WORK-DATA-009`, `WORK-DATA-010`, `WORK-DATA-013`, MCP identity, tagged union, recursive/untagged-union, or DAG TypeRef hint scope.
- Confirm no spec, YAML, fixture, or implementation files are changed by this task.

## Evidence

Verdict: PASS.

Decision result:

- WORK-DATA-014 will not introduce a DATA DSL extension for selector support matrices, object-dependent `kind` vocabulary, dependent enums, or selector-field combination validation at this stage.
- The main contract owner is the MCP schema and MCP tool contract surface:
  - `docs/spec/mcp/schema.md`
  - `docs/spec/mcp/tools/*.md`
  - related MCP error / diagnostic documentation when unsupported selector behavior must be made explicit.
- YAML remains an example / fixture surface. UC-002 YAML notes may point to the MCP schema/tool specs, but YAML should not become the source of truth for the selector matrix.
- No new ADR is required for this boundary decision at this time because the selected direction keeps the contract in existing MCP spec surfaces and avoids a new cross-cutting DATA model capability.

Candidate ownership:

| ID | Decision |
|---|---|
| N-020 | `get_signature_request.selector` support is owned by the MCP selector support matrix and `get_signature` tool contract. |
| N-031 | `list_objects_request.kind` object-dependent vocabulary is owned by the `list_objects` tool contract plus shared MCP object/kind vocabulary. |
| N-037 | `object_selector.kind` validation against resolved object type is owned by the shared MCP Object selector schema. |
| N-040 | valid `object_selector` field combinations are owned by the shared MCP Object selector schema and tool-specific selector support rules. |
| N-042 | `object_ref.kind` object-dependent response vocabulary is owned by the shared MCP ObjectRef schema. |

Follow-up recommendation:

- Next task should update `docs/spec/mcp/schema.md` and the relevant MCP tool specs so the selector support matrix, object-dependent kind vocabulary, and unsupported / limited behavior are explicit.
- A later task should clean up UC-002 YAML notes for N-020, N-031, N-037, N-040, and N-042 so they point to the accepted MCP schema/tool contracts instead of carrying broad prose notes.
- Implementation/runtime verification should be split after the spec wording is accepted. It should check whether current Go behavior matches the accepted selector matrix and unsupported-selector behavior.
- Fixture/golden regeneration should only be done if YAML cleanup or implementation changes affect rendered outputs.

Verification performed:

- Compared against `REQ-DATA-007` / `WORK-DATA-014` scope: the decision covers selector support matrices and object-dependent vocabulary without mixing in excluded tagged union, recursive/untagged union, numeric/default behavior, DAG TypeRef hint, or MCP identity work.
- Confirmed the source bucket is N-020, N-031, N-037, N-040, and N-042 from the UC-002 notes-retreat classification lineage.
- Confirmed no spec, YAML, fixture, or implementation files are intentionally changed by this task.
