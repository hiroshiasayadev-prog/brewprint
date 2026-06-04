# TASK-MCP-020-03: Implement metadata field replacement and contract tests

- **id**: TASK-MCP-020-03
- **status**: todo
- **date**: 2026-06-05
- **work_item**: WORK-MCP-020
- **source_requirement**: REQ-MCP-020
- **estimate**: 1.5d
- **depends_on**:
  - TASK-MCP-020-02
- **outputs**:
  - metadata field replacement implementation
  - contract tests for preserved metadata and required metadata validation

## Goal

Implement the field-level metadata replacement operation for `propose_record_update` and cover it with contract tests.

## Work

- Add the new update operation to request parsing and validation.
- Reuse existing record metadata loading and metadata validation where possible.
- Apply requested metadata field changes on top of the existing complete metadata block.
- Preserve unspecified metadata fields.
- Ensure attempts to clear or invalidate required metadata fields fail with existing metadata diagnostics.
- Keep `metadata_block_replace` behavior unchanged.
- Add regression tests for status-only task update, preserved list metadata, invalid required metadata, and existing whole-block replacement compatibility.

## Done condition

- The new operation is implemented in the authoring update path.
- Contract tests cover the acceptance criteria from `REQ-MCP-020`.
- Existing metadata block replacement tests continue to pass.

## Verification

- Run targeted Go tests for designrecords authoring behavior.
- Run MCP tool call tests that cover request schema and handler behavior.

## Evidence

Not started.
