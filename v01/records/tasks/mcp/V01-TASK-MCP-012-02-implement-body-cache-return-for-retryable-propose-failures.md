# V01-TASK-MCP-012-02: implement body cache return for retryable propose failures

- **id**: V01-TASK-MCP-012-02
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-MCP-012
- **source_requirement**: V01-REQ-MCP-015
- **estimate**: 1d-2d
- **depends_on**:
  - V01-TASK-MCP-012-01
- **outputs**:
  - Updated Design Records MCP authoring implementation
  - Regression tests for retryable failed-propose body_cache responses

## Goal

Implement and lock down the body_cache return classification for retryable propose failures. Ensure that:

- `body_cache` is NOT returned for request-level invalid cases where body was never received as a content source.
- `body_cache` IS returned for cases where body was received as a content source but proposal preparation failed after.

## Work

1. Reviewed `internal/designrecords/authoring.go`: `ProposeRecordCreate`, `ProposeRecordUpdate`, `resolveBodySource`, `prepareCreate`, `prepareUpdate`, `replaceNamedSection`, `failedProposalResponse`.
2. Traced all required classification paths and confirmed the existing implementation already handles all cases correctly:
   - Early rejection (body + body_cache_id, fields + body_cache_id) returns `failedProposalResponse(nil, ...)` — no body_cache.
   - After `resolveBodySource` receives a body string, bodyCache is set. Error paths in `ProposeRecordCreate` and `ProposeRecordUpdate` re-cache the body and pass it to `failedProposalResponse`.
3. Identified missing test assertions:
   - `TestAuthoringNamedSectionSelectors` ambiguous case did not assert `BodyCache != nil`.
   - `TestAuthoringCreateInputContractNormalization` `fields + full-record body` case did not assert `BodyCache != nil`.
   - No tests for legacy full-body create failures (domain mismatch, target already exists).
   - No test for non-string `body` at JSON decode boundary.
4. Added `|| multi.BodyCache == nil` assertion to existing ambiguous test in `TestAuthoringNamedSectionSelectors`.
5. Added `TestBodyCacheReturnClassification` in `internal/designrecords/authoring_test.go` covering all required classification cases.
6. Added non-string `body` decode error test cases to `TestToolsCallToolErrors` in `internal/designrecordsmcp/tools_call_test.go`.

## Done condition

- Regression tests cover the final classification.
- `go test ./internal/designrecords ./internal/designrecordsmcp` passes.
- `go test ./...` passes.
- No V01-WORK-MCP-014 behavior regressed.

## Verification

All classification test sub-cases pass:

- `no_cache_body_plus_body_cache_id_create` — rejected before body received, no body_cache ✓
- `no_cache_fields_plus_body_cache_id_create` — rejected before body received, no body_cache ✓
- `cache_returned_update_no_match_plus_body` — body received, section no match, body_cache returned ✓
- `cache_returned_update_ambiguous_plus_body` — body received, section ambiguous, body_cache returned ✓
- `cache_returned_create_fields_plus_full_record_body` — body received, rejected as stale full-record body, body_cache returned ✓
- `cache_returned_create_fields_plus_section_body_render_failure` — body received and validated, render failed after, body_cache returned ✓
- `cache_returned_legacy_full_body_create_domain_mismatch` — body received, domain mismatch, body_cache returned ✓
- `cache_returned_legacy_full_body_create_target_already_exists` — body received, target already exists, body_cache returned ✓
- `propose_record_create non-string body is decode error no body_cache` — decode error, no body_cache ✓
- `propose_record_update non-string body is decode error no body_cache` — decode error, no body_cache ✓

## Evidence

Implementation inspection confirmed no implementation change was required. All required classification behavior was already implemented in `authoring.go`. The task work was adding the regression tests and fixing the ambiguous assertion that was missing.

Commands run on 2026-06-03:

```
go test ./internal/designrecords ./internal/designrecordsmcp
```

Result: `ok` for both packages in 1.858s / 0.547s.

```
go test ./...
```

Result: all 17 test packages pass with no failures.

Files changed:

- `internal/designrecords/authoring_test.go`: Added `TestBodyCacheReturnClassification` (8 sub-cases); fixed `TestAuthoringNamedSectionSelectors` ambiguous assertion to include `|| multi.BodyCache == nil`.
- `internal/designrecordsmcp/tools_call_test.go`: Added 2 decode-error test cases to `TestToolsCallToolErrors` for non-string `body` on `propose_record_create` and `propose_record_update`.

No implementation changes to `authoring.go` or `tools_call.go` were required. No DATA/UC/ADR files were touched. No git add or commit was run.
