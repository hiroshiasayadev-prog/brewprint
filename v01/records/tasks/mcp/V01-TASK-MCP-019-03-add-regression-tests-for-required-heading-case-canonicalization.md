# V01-TASK-MCP-019-03: Add regression tests for required heading case canonicalization

- **id**: V01-TASK-MCP-019-03
- **status**: done
- **date**: 2026-06-05
- **work_item**: V01-WORK-MCP-019
- **source_requirement**: V01-REQ-MCP-021
- **estimate**: 1d-2d
- **depends_on**:
  - V01-TASK-MCP-019-02
- **outputs**:
  - Regression tests for validation repair diagnostics
  - Regression tests for authoring selector fallback and proposal diff canonicalization
  - MCP public contract tests for new diagnostic fields

## Goal

Add regression tests that lock the V01-REQ-MCP-021 spec contract before implementation.

The tests should cover the narrow case-only heading canonicalization behavior for validation-required workflow headings, while proving that optional headings, cross-kind headings, and non-case differences remain outside the fallback.

## Work

Add targeted tests for:

- Validation: gated task with case-only mismatch emits `missing_required_section` plus info `section_heading_case_mismatch` with `section`, `actual_heading`, and `status`.
- Validation: same case-only mismatch on non-gated status does not emit the repair diagnostic.
- Authoring: `propose_record_update` for a task selector matching a case-only existing required heading creates a proposal.
- Authoring: retained proposal diff rewrites the matched heading to canonical text.
- Authoring: fallback works even when the target record is not currently in gated status.
- Authoring: ambiguous case-insensitive headings fail with `section_selector_ambiguous`, no proposal, and candidate headings.
- Authoring: non-case mismatch fails with `section_selector_no_match`; no fuzzy matching.
- Authoring: optional or user-defined headings are not canonicalized.
- Authoring: headings required for a different target record kind are not canonicalized.
- Authoring: exact case-sensitive match remains the default and does not take fallback.
- Authoring: fallback honors `section_selector.level` when supplied.
- MCP boundary: public tool response exposes the new diagnostic fields, including `actual_heading`.

## Done condition

- Tests express the current spec/schema contract clearly enough for implementation in V01-TASK-MCP-019-04.
- Tests fail against current implementation where behavior is not yet implemented.
- Tests include no broad fuzzy matching expectation.
- Tests confirm optional and cross-kind headings are outside the fallback.
- Test names and assertions make the target-kind and gated-status split explicit.

## Verification

Run targeted tests as appropriate, then the package test set:

```powershell
cd C:\Users\imved\projects\brewprint
go test ./internal/designrecords ./internal/designrecordsmcp
```

If tests are intentionally failing before implementation, record the failing test names and failure reason in Evidence.

## Evidence
Regression tests were added by Codex.

Files changed:

- `internal/designrecords/authoring_test.go`
- `internal/designrecords/validation_test.go`
- `internal/designrecordsmcp/tools_call_test.go`

Tests added:

- `TestRequiredSectionHeadingCaseMismatchDiagnostics`
  - gated task expects `missing_required_section` plus info `section_heading_case_mismatch`.
  - non-gated task confirms no repair info diagnostic.
- `TestProposeRecordUpdateRequiredHeadingCaseFallback`
  - proposal creation and heading canonicalization diff.
  - fallback on non-gated task.
  - ambiguous case-insensitive failure.
  - non-case mismatch remains no-match.
  - optional heading not canonicalized.
  - cross-kind required heading not canonicalized.
  - exact match remains default.
  - level constraint honored.
- `TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields`
  - public validate_records response must expose `actual_heading` and `candidate_headings`.

Commands run:

```powershell
go test ./internal/designrecords -run TestRequiredSectionHeadingCaseMismatch -v
go test ./internal/designrecords -run TestProposeRecordUpdateRequiredHeadingCaseFallback -v
go test ./internal/designrecordsmcp -run TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields -v
go test ./internal/designrecords ./internal/designrecordsmcp
```

Expected current failures before implementation:

- `TestRequiredSectionHeadingCaseMismatchDiagnostics/gated_task_emits_strict_missing_error_plus_repair_info`
  - current code emits only `missing_required_section`; it does not emit `section_heading_case_mismatch` yet.
- `TestProposeRecordUpdateRequiredHeadingCaseFallback`
  - positive fallback cases currently return `section_selector_no_match`.
  - ambiguous case-only headings currently return `section_selector_no_match`, not `section_selector_ambiguous`.
  - guardrail cases for non-case mismatch, optional headings, cross-kind headings, exact match, and level no-match already pass.
- `TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields`
  - public response lacks `section_heading_case_mismatch` and therefore lacks `actual_heading`.

Implementation input for `V01-TASK-MCP-019-04`:

- Add `DiagnosticSectionHeadingCaseMismatch` category.
- Add `ActualHeading string` with JSON field `actual_heading,omitempty` to `Diagnostic` and custom JSON marshaling.
- Validation: when gated required-section validation finds canonical section missing, emit info `section_heading_case_mismatch` if exactly one same-text-different-case heading exists. Keep `missing_required_section`.
- Authoring: after exact selector matching finds zero matches, apply case-only fallback only for workflow artifact kinds and only for required headings of the target kind.
- Authoring fallback must respect `section_selector.level`, fail ambiguous case-insensitive matches with `section_selector_ambiguous`, and rewrite the matched heading line to canonical selector heading in the proposal diff.
- Keep optional/user-defined headings, cross-kind headings, and non-case differences on existing exact/no-match behavior.
