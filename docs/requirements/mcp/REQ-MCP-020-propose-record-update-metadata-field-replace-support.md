# REQ-MCP-020: propose_record_update metadata field replace support

- **id**: REQ-MCP-020
- **status**: captured
- **date**: 2026-06-05
- **source_refs**:
  - TASK-DATA-011-01
- **work_items**:
  - WORK-MCP-020

## Requirement

`propose_record_update` MUST support partial metadata field replacement for existing design records and workflow artifacts.

Callers MUST be able to update one or more metadata fields without reconstructing and resubmitting the complete metadata block.

This is required because common lifecycle operations, such as closing a task by changing only `status` to `done`, currently require a full `metadata_block_replace`. That operation fails unless required metadata fields such as `id` are also supplied, even though those fields already exist in the target record.

## Evidence

Runtime authoring attempt for `TASK-DATA-011-01` failed when trying to update only `status` via `metadata_block_replace`:

```json
{"proposal_created":false,"validation":{"ok":false,"diagnostics":[{"category":"missing_required_metadata","severity":"error","message":"missing required metadata field id","field":"id"}]},"diagnostics":[{"category":"missing_required_metadata","severity":"error","message":"missing required metadata field id","field":"id"}]}
```

The client then had to retry with a complete metadata block. That is unnecessarily error-prone for field-level lifecycle updates.

## Required Outcome

The authoring update contract supports a field-level metadata update operation, for example `metadata_field_replace` or `metadata_fields_replace`.

The operation MUST preserve unspecified existing metadata fields, apply the requested field changes, and then validate the resulting complete metadata block using the same metadata validation rules as full metadata replacement.

Acceptance criteria:

- A caller can update only `status` for an existing task without providing `id`.
- Unspecified existing metadata fields are preserved.
- The resulting complete metadata block is validated after the field-level patch is applied.
- Attempts to clear or invalidate required metadata fields still fail with existing metadata diagnostics.
- `metadata_block_replace` remains supported for intentional whole-block replacement.

## Explicitly Excluded Scope

- Removing support for `metadata_block_replace`.
- Changing workflow artifact status vocabularies.
- Changing required metadata rules.
- Automatically closing related work items or requirements.

## Boundary

This requirement belongs to the MCP authoring update contract. The observed target record was `TASK-DATA-011-01`, but the missing capability is not DATA-domain behavior. It is a Design Records MCP write/update ergonomics and correctness gap.
