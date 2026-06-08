# V01-TASK-MCP-020-03: Implement metadata field replacement and contract tests

- **id**: V01-TASK-MCP-020-03
- **status**: done
- **date**: 2026-06-05
- **work_item**: V01-WORK-MCP-020
- **source_requirement**: V01-REQ-MCP-020
- **estimate**: 1.5d
- **depends_on**:
  - V01-TASK-MCP-020-02
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
- Contract tests cover the acceptance criteria from `V01-REQ-MCP-020`.
- Existing metadata block replacement tests continue to pass.

## Verification

- Run targeted Go tests for designrecords authoring behavior.
- Run MCP tool call tests that cover request schema and handler behavior.

## Evidence
Verdict: PASS.

Implementation completed on 2026-06-05.

Changed files:
- `internal/designrecords/authoring.go`: Added `UpdateTypeMetadataFieldsReplace = "metadata_fields_replace"` constant; updated `bodyForbidden` check in `ProposeRecordUpdate` to include `metadata_fields_replace`; added `UpdateTypeMetadataFieldsReplace` case in `prepareUpdate` switch; added helper functions `patchMetadataFields`, `currentMetadataAsMap`, `currentSpecMetadataAsMap`, `rawMetadataScalarValue`.
- `internal/designrecords/authoring_test.go`: Added `TestAuthoringMetadataFieldsReplace`, `TestAuthoringMetadataFieldsReplaceBodyForbidden`, `TestAuthoringMetadataFieldsReplaceRequiredFieldValidation`.
- `internal/designrecordsmcp/tools.go`: Added `metadata_fields_replace` to the `update.type` enum in `updateSchema()`.

Implementation summary:
- `metadata_fields_replace` reads existing record metadata into a map via `currentMetadataAsMap`, patches only the caller-supplied fields via `patchMetadataFields`, then passes the merged map to the existing `replaceMetadataBlock` path, reusing all existing rendering and validation behavior.
- `body` and `body_cache_id` are forbidden for `metadata_fields_replace`; supplying either returns `invalid_body_source`.
- `metadata_block_replace` behavior is unchanged.
- `named_section_replace` behavior is unchanged.

Tests added:
- `TestAuthoringMetadataFieldsReplace`: task status-only update (only `status: done` supplied); verifies `id`, `date`, `work_item`, `source_requirement`, `estimate`, and `outputs` list item `initial` are all preserved in the diff.
- `TestAuthoringMetadataFieldsReplaceBodyForbidden`: `body` with `metadata_fields_replace` returns `invalid_body_source` and `ProposalCreated: false`.
- `TestAuthoringMetadataFieldsReplaceRequiredFieldValidation`: clearing `status` to empty string produces `Validation.OK: false` with `empty_required_metadata` diagnostic (post-proposal validation).

Commands run and results:
```
go test ./internal/designrecords -run TestAuthoringMetadata -v
# All 5 metadata tests PASS (3 new + 2 existing)

go test ./internal/designrecords ./internal/designrecordsmcp
# ok  internal/designrecords     3.245s
# ok  internal/designrecordsmcp  0.639s
# FAIL count: 0
```

Known unrelated failures: none in the targeted packages. Heading canonicalization tests (V01-REQ-MCP-021 / V01-WORK-MCP-019) are in pre-existing dirty state outside this scope.

Remaining concerns: none for this task scope. V01-TASK-MCP-020-04 runtime smoke will be the next step.
