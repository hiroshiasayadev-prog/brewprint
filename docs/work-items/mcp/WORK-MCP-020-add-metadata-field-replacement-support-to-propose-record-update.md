# WORK-MCP-020: Add metadata field replacement support to propose_record_update

- **id**: WORK-MCP-020
- **status**: done
- **date**: 2026-06-05
- **source_requirement**: REQ-MCP-020
- **impact_refs**:
  - SPEC-design-records-mcp-tools
  - ADR-093
  - TASK-DATA-011-01
- **tasks**:
  - TASK-MCP-020-01
  - TASK-MCP-020-02
  - TASK-MCP-020-03
  - TASK-MCP-020-04

## Goal

Add field-level metadata replacement support to `propose_record_update` so callers can update one or more metadata fields without reconstructing the complete metadata block.

The primary lifecycle case is updating only `status` on an existing workflow artifact while preserving required fields such as `id`, `date`, `work_item`, `source_requirement`, `depends_on`, and `outputs` from the current record.

## Boundary

In scope:

- Define the MCP authoring update contract for field-level metadata replacement.
- Preserve unspecified metadata fields from the target record.
- Validate the resulting complete metadata block after applying the field-level patch.
- Add implementation and regression coverage for task status-only update behavior.
- Keep `metadata_block_replace` available for intentional whole-block replacement.

Out of scope:

- Removing or weakening `metadata_block_replace`.
- Changing workflow artifact status vocabularies.
- Changing required metadata rules.
- Automatically closing related work items or requirements.
- DATA-domain behavior changes.

## Impact Scope

- `SPEC-design-records-mcp-tools`: document the new `propose_record_update` operation and request semantics.
- `ADR-093`: check whether the authoring transaction model needs a short compatibility note or no change.
- `internal/designrecords`: implement metadata field patching and validation behavior.
- `internal/designrecordsmcp`: expose the new update operation through the MCP tool request schema and call handling.
- Tests: add regression coverage for status-only task update and invalid required metadata changes.

## Task flow

TASK-MCP-020-01 -> TASK-MCP-020-02 -> TASK-MCP-020-03 -> TASK-MCP-020-04

## Task Candidates

- `TASK-MCP-020-01`: Review current metadata update contract and implementation boundary.
- `TASK-MCP-020-02`: Update the MCP tools spec for metadata field replacement.
- `TASK-MCP-020-03`: Implement metadata field replacement and contract tests.
- `TASK-MCP-020-04`: Run runtime smoke and close synchronization.

## Completion Condition

- `propose_record_update` supports a field-level metadata replacement operation for existing records.
- Updating only `status` for an existing task succeeds without resubmitting `id` or other required metadata fields.
- Unspecified metadata fields are preserved.
- Invalid changes to required metadata still fail through existing metadata diagnostics.
- `metadata_block_replace` remains supported.
- Tests and runtime smoke evidence are recorded.
- `REQ-MCP-020`, `WORK-MCP-020`, and all related tasks are synchronized to their final statuses.

## Evidence

Verdict: PASS.

Completed scope:

- `SPEC-design-records-mcp-tools` defines `metadata_fields_replace` for `propose_record_update`.
- Implementation supports status-only metadata field replacement while preserving unspecified metadata fields.
- Existing `metadata_block_replace` behavior remains supported.
- Invalid required metadata changes still fail through existing metadata diagnostics.
- Contract tests pass.
- Runtime smoke through MCP JSON-RPC passes in a temp root.

Verification:

- `go test ./internal/designrecords -run TestAuthoringMetadata -v`: PASS
- `go test ./internal/designrecords ./internal/designrecordsmcp`: PASS
- Runtime smoke confirmed `metadata_fields_replace` can update only `TASK-MCP-020-04.status` to `doing` via proposal / accept flow while preserving `work_item`, `source_requirement`, `depends_on`, and `outputs`.

Close synchronization:

- `TASK-MCP-020-01`: done
- `TASK-MCP-020-02`: done
- `TASK-MCP-020-03`: done
- `TASK-MCP-020-04`: done
- `WORK-MCP-020`: done
- `REQ-MCP-020`: accepted
