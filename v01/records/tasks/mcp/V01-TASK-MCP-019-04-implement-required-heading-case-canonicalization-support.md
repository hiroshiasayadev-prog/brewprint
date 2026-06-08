# V01-TASK-MCP-019-04: Implement required heading case canonicalization support

- **id**: V01-TASK-MCP-019-04
- **status**: done
- **date**: 2026-06-05
- **work_item**: V01-WORK-MCP-019
- **source_requirement**: V01-REQ-MCP-021
- **estimate**: 1d-2d
- **depends_on**:
  - V01-TASK-MCP-019-03
- **outputs**:
  - Diagnostic model support for section_heading_case_mismatch and actual_heading
  - Validation repair diagnostic implementation
  - Authoring selector fallback and canonical heading rewrite implementation
  - Passing regression tests for V01-REQ-MCP-021

## Goal

Implement the V01-REQ-MCP-021 behavior defined by the updated tools/schema spec and locked by `V01-TASK-MCP-019-03` regression tests.

## Work

- Add `DiagnosticSectionHeadingCaseMismatch` category.
- Add `ActualHeading string` with JSON field `actual_heading,omitempty` to `Diagnostic` and custom JSON marshaling.
- Update required-section validation so gated workflow artifacts keep strict `missing_required_section` behavior and additionally emit info `section_heading_case_mismatch` when exactly one same-text-different-case heading exists for the missing canonical required heading.
- Update authoring named section replacement so exact matching remains the default.
- After exact selector matching finds zero matches, apply case-only fallback only for workflow artifact kinds and only for required headings of the target record kind.
- Ensure fallback respects `section_selector.level`.
- Ensure ambiguous case-insensitive matches fail with `section_selector_ambiguous` and candidate headings.
- Ensure successful fallback rewrites the matched heading line to canonical `section_selector.heading` text in the retained proposal diff.
- Preserve existing exact/no-match behavior for optional/user-defined headings, cross-kind headings, and non-case differences.

## Done condition

- `TestRequiredSectionHeadingCaseMismatchDiagnostics` passes.
- `TestProposeRecordUpdateRequiredHeadingCaseFallback` passes.
- `TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields` passes.
- `go test ./internal/designrecords ./internal/designrecordsmcp` passes.
- Implementation does not introduce fuzzy matching or optional heading canonicalization.
- Evidence records the implemented files and test results.

## Verification

Run targeted tests and full package tests:

```powershell
cd C:\Users\imved\projects\brewprint
go test ./internal/designrecords -run TestRequiredSectionHeadingCaseMismatch -v
go test ./internal/designrecords -run TestProposeRecordUpdateRequiredHeadingCaseFallback -v
go test ./internal/designrecordsmcp -run TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields -v
go test ./internal/designrecords ./internal/designrecordsmcp
```

## Evidence
Implementation completed by Claude Code.

Verdict: PASS.

Files changed:

- `internal/designrecords/types.go`
  - Added `DiagnosticSectionHeadingCaseMismatch`.
  - Added `ActualHeading string` to `Diagnostic`.
  - Added `actual_heading,omitempty` to custom JSON marshaling.
- `internal/designrecords/validation.go`
  - Added `findCaseOnlyMismatch` helper.
  - Updated required narrative section diagnostics to emit info `section_heading_case_mismatch` after `missing_required_section` when the canonical required section is missing and exactly one case-only mismatch exists.
- `internal/designrecords/authoring.go`
  - Added `requiredSectionsForKind` helper.
  - Updated `replaceNamedSection` to accept optional record kind context and implement required-heading case-only fallback.
  - Updated `prepareUpdate` call path to pass `req.Kind`.

Implementation summary:

- Validation keeps strict `missing_required_section` behavior.
- Validation emits `section_heading_case_mismatch` only for gated workflow artifacts when the required canonical heading is missing and a single case-only non-canonical heading exists.
- Authoring exact matching remains default and case-sensitive.
- Authoring fallback applies only after exact matching finds zero matches.
- Authoring fallback is limited to workflow artifact kinds and required headings of the target kind.
- Authoring fallback respects `section_selector.level`.
- Ambiguous case-insensitive matches return `section_selector_ambiguous`.
- Successful fallback rewrites the matched heading line to canonical `section_selector.heading` in the proposal diff.
- Optional/user-defined headings, cross-kind headings, and non-case differences remain on existing exact/no-match behavior.

Tests run:

```powershell
go test ./internal/designrecords -run TestRequiredSectionHeadingCaseMismatch -v
go test ./internal/designrecords -run TestProposeRecordUpdateRequiredHeadingCaseFallback -v
go test ./internal/designrecordsmcp -run TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields -v
go test ./internal/designrecords ./internal/designrecordsmcp
```

Results:

- `TestRequiredSectionHeadingCaseMismatchDiagnostics`: PASS, all 2 subtests.
- `TestProposeRecordUpdateRequiredHeadingCaseFallback`: PASS, all 8 subtests.
- `TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields`: PASS.
- `go test ./internal/designrecords ./internal/designrecordsmcp`: PASS.

Closure assessment:

- All `V01-TASK-MCP-019-04` done conditions are satisfied.
- No fuzzy matching was introduced.
- Optional heading canonicalization and cross-kind heading canonicalization remain excluded.
