# TASK-MCP-019-04: Implement required heading case canonicalization support

- **id**: TASK-MCP-019-04
- **status**: todo
- **date**: 2026-06-05
- **work_item**: WORK-MCP-019
- **source_requirement**: REQ-MCP-021
- **estimate**: 1d-2d
- **depends_on**:
  - TASK-MCP-019-03
- **outputs**:
  - Diagnostic model support for section_heading_case_mismatch and actual_heading
  - Validation repair diagnostic implementation
  - Authoring selector fallback and canonical heading rewrite implementation
  - Passing regression tests for REQ-MCP-021

## Goal

Implement the REQ-MCP-021 behavior defined by the updated tools/schema spec and locked by `TASK-MCP-019-03` regression tests.

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

未実施。
