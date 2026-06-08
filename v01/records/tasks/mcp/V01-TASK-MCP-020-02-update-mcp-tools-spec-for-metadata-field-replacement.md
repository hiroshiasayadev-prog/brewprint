# V01-TASK-MCP-020-02: Update MCP tools spec for metadata field replacement

- **id**: V01-TASK-MCP-020-02
- **status**: done
- **date**: 2026-06-05
- **work_item**: V01-WORK-MCP-020
- **source_requirement**: V01-REQ-MCP-020
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-020-01
- **outputs**:
  - SPEC-design-records-mcp-tools update for metadata field replacement contract

## Goal

Update the MCP tools spec to define the field-level metadata replacement operation for `propose_record_update`.

## Work

- Define the operation name, preferably `metadata_fields_replace` unless investigation finds a stronger existing naming convention.
- Document request shape and field semantics.
- Specify that existing metadata is read first, requested fields are patched, unspecified fields are preserved, and the complete resulting metadata block is validated.
- Specify failure behavior for clearing or invalidating required metadata fields.
- State that `metadata_block_replace` remains supported for whole-block replacement.
- Update spec front matter metadata if required by the spec authoring guide.

## Done condition

- The current spec clearly describes the new operation and its validation behavior.
- The spec distinguishes field-level replacement from whole-block replacement.
- The spec does not include implementation checklist content.

## Verification

- Review the updated spec against `V01-REQ-MCP-020` acceptance criteria.
- Confirm terminology is consistent with the existing `propose_record_update` contract.

## Evidence
Verdict: PASS.

Spec updated: `docs/spec/design-records-mcp/tools.md`

Changes applied on 2026-06-05:

- **Purpose** (`propose_record_update`): Added `metadata_fields_replace` (field-level metadata patch) to the operation list.
- **Request examples**: Added `metadata_fields_replace` JSON example (status-only task update).
- **`update.type` values table**: Added `metadata_fields_replace` row.
- **New subsection `#### Metadata field replacement`**: Inserted between `Metadata block replacement` and `Named section replacement`. Specifies read-then-patch semantics, preservation of unspecified fields, post-patch complete-block validation reuse, body-forbidden rule, and continued `metadata_block_replace` availability.
- **Body source and body cache table**: Added `propose_record_update` `metadata_fields_replace` row with `body` / `body_cache_id` must be omitted.

`last_updated` was already `2026-06-05`; no front matter change needed.

V01-REQ-MCP-020 acceptance criteria coverage:
- Caller can update only `status` without providing `id`: covered by read-then-patch semantics.
- Unspecified fields preserved: explicitly stated.
- Post-patch complete-block validation: explicitly stated.
- Invalid/cleared required fields still fail: explicitly stated via existing diagnostics reuse.
- `metadata_block_replace` remains supported: explicitly stated.

Review result:

- Codex review verdict: PASS.
- Blocking findings: none.
- Non-blocking finding: workflow artifact files were untracked at review time and must be included in the intended spec-only changeset.
- Reviewed files: `docs/spec/design-records-mcp/tools.md`, `docs/tasks/mcp/TASK-MCP-020-02-update-mcp-tools-spec-for-metadata-field-replacement.md`, `docs/requirements/mcp/REQ-MCP-020-propose-record-update-metadata-field-replace-support.md`, `docs/work-items/mcp/WORK-MCP-020-add-metadata-field-replacement-support-to-propose-record-update.md`, and `docs/adr/093-design-records-mcp-authoring-transaction-model.md`.
- Validation confirmed OK for `V01-REQ-MCP-020`, `V01-WORK-MCP-020`, `V01-TASK-MCP-020-02`, `SPEC-design-records-mcp-tools`, and `V01-ADR-093`.
- `go test ./internal/designrecords ./internal/designrecordsmcp` failed only in heading canonicalization tests under `V01-REQ-MCP-021` / `V01-WORK-MCP-019`; this is outside `V01-TASK-MCP-020-02` scope.

ADR impact judgment:

- No `V01-ADR-093` update is needed. `V01-ADR-093` is migrated to the spec and delegates current tool names / request-response schemas to `SPEC-design-records-mcp-tools`.
