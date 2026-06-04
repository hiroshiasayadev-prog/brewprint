# TASK-MCP-020-01: Review current metadata update contract and implementation boundary

- **id**: TASK-MCP-020-01
- **status**: done
- **date**: 2026-06-05
- **work_item**: WORK-MCP-020
- **source_requirement**: REQ-MCP-020
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Current behavior inventory for propose_record_update metadata replacement
  - Implementation and test touchpoint list

## Goal

Review the current `propose_record_update` metadata update contract and implementation boundary before changing the tool surface.

## Work

- Read the current `SPEC-design-records-mcp-tools` `propose_record_update` section.
- Inspect current authoring update implementation and MCP request handling.
- Identify where `metadata_block_replace` validates required metadata fields.
- Identify the minimum implementation and test touchpoints for field-level metadata replacement.
- Check whether `ADR-093` needs a compatibility note or whether spec-only update is sufficient.

## Done condition

- Current behavior and failure point are summarized.
- Implementation touchpoints are listed.
- Test touchpoints are listed.
- ADR impact is classified as update-needed or no-change-needed.

## Verification

- Confirm findings against current spec and implementation files.
- Confirm the findings explain the `TASK-DATA-011-01` status-only update failure captured by `REQ-MCP-020`.

## Evidence
Verdict: PASS.

Current behavior:

- `propose_record_update` currently supports `metadata_block_replace` and `named_section_replace` only.
- No `metadata_fields_replace` or equivalent field-level metadata operation exists yet.
- `metadata_block_replace` calls the metadata rendering path as a whole-block replacement.
- For task metadata, `renderTaskMetadata()` / `renderWorkflowMetadata()` require the complete task metadata set, including `id`, `status`, `date`, `work_item`, `source_requirement`, `estimate`, `depends_on`, and `outputs`.
- A status-only metadata update such as `{ "status": "done" }` therefore fails because required fields such as `id` are absent, matching the failure captured by `REQ-MCP-020`.

Implementation touchpoints identified:

- `internal/designrecords/authoring.go`
  - add `UpdateTypeMetadataFieldsReplace = "metadata_fields_replace"`
  - include the new operation in the `ProposeRecordUpdate()` body-forbidden path
  - add a new `prepareUpdate()` switch branch
  - add helper logic that builds current metadata from the target record, applies requested field changes, and reuses existing whole-block metadata replacement / validation behavior
- `internal/designrecordsmcp/tools.go`
  - add `metadata_fields_replace` to the `update.type` enum in the tool schema

Test touchpoints identified:

- Existing `metadata_block_replace` tests must continue to pass.
- Add regression coverage for status-only task update with preserved unspecified fields.
- Add regression coverage that invalid required metadata changes still fail through existing metadata diagnostics.
- Add body / body_cache forbidden coverage for `metadata_fields_replace`.
- Prefer empty / invalid required fields such as empty `status` or `date` for required metadata failure coverage. Changing `id` to a different non-empty value may exercise identity validation rather than the core required-metadata acceptance criteria.

Spec / ADR impact:

- `SPEC-design-records-mcp-tools` needs an update to define `metadata_fields_replace`, request semantics, body-source rules, and validation behavior.
- `metadata_fields_replace` should preserve unspecified existing metadata fields, patch only supplied fields, and validate the resulting complete metadata block using existing metadata validation rules.
- `metadata_block_replace` remains supported for intentional whole-block replacement.
- `ADR-093` update is not required for correctness because current behavior is owned by the spec. A short compatibility note may be added only if reviewers want to reduce confusion around ADR-093's historical set-only wording.

Known unrelated test state:

- Reported failures in heading canonicalization tests belong to `REQ-MCP-021` / `WORK-MCP-019` and are outside `REQ-MCP-020` / `WORK-MCP-020` scope.
