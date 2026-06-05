# WORK-MCP-019: section heading canonicalization tolerant authoring update support

- **id**: WORK-MCP-019
- **status**: done
- **date**: 2026-06-05
- **source_requirement**: REQ-MCP-021
- **impact_refs**:
  - SPEC-design-records-mcp-tools
  - ADR-093
- **tasks**:
  - TASK-MCP-019-01
  - TASK-MCP-019-02
  - TASK-MCP-019-03
  - TASK-MCP-019-04
  - TASK-MCP-019-05

## Goal

REQ-MCP-021 の required section heading case mismatch support を実現する。

## Boundary

この work item は Design Records MCP authoring update と workflow artifact validation の境界を扱う。

対象外:

- fuzzy heading matching
- canonical section name 自体の変更
- proposal / accept を経由しない mutation
- required-section validation の緩和
- DATA-domain task semantics の変更

## Impact Scope

- `SPEC-design-records-mcp-tools`
- `ADR-093`
- `internal/designrecords` authoring / validation implementation
- `internal/designrecordsmcp` tool call boundary and tests

## Task flow

```text
TASK-MCP-019-01 current gap review
  -> TASK-MCP-019-02 spec contract update
  -> TASK-MCP-019-03 regression tests
  -> TASK-MCP-019-04 implementation
  -> TASK-MCP-019-05 runtime smoke and close synchronization
```

## Task Candidates

- TASK-MCP-019-01: Review current heading selector, required-section validation, and authoring update behavior.
- TASK-MCP-019-02: Update MCP tools spec for safe case-only heading canonicalization proposal behavior.
- TASK-MCP-019-03: Add regression tests for case-only repair, ambiguous failure, non-case mismatch behavior, validation repair diagnostics, and proposal diff canonicalization.
- TASK-MCP-019-04: Implement heading canonicalization support in the authoring update path.
- TASK-MCP-019-05: Run runtime smoke, record evidence, and synchronize close status.

## Completion Condition
- REQ-MCP-021 acceptance criteria are satisfied.
- Case-only required-section heading mismatches can be repaired through MCP proposal / accept flow.
- Ambiguous heading matches fail safely with candidate headings.
- Non-case heading differences remain governed by existing selector and validation rules.
- Tests and runtime smoke evidence are recorded.
- REQ-MCP-021, this work item, and child tasks are status-synchronized.

## Evidence

Implementation and smoke completed for REQ-MCP-021.

Completed tasks:

- `TASK-MCP-019-01`: current behavior / implementation boundary review.
- `TASK-MCP-019-02`: tools/schema spec update and review fixes.
- `TASK-MCP-019-03`: regression tests for required heading case canonicalization.
- `TASK-MCP-019-04`: implementation.
- `TASK-MCP-019-05`: smoke and close synchronization.

Implemented behavior:

- Validation emits `section_heading_case_mismatch` info diagnostics with `actual_heading` for gated workflow artifacts when a required canonical heading is missing and a single case-only mismatch exists.
- `missing_required_section` remains strict and is not suppressed.
- `propose_record_update` named section replacement applies case-only fallback only for validation-required headings of the target workflow artifact kind.
- Successful fallback rewrites the matched heading line to canonical `section_selector.heading` in the retained proposal diff.
- Ambiguous case-insensitive matches fail closed with `section_selector_ambiguous` and candidate headings.
- Optional/user-defined headings, cross-kind headings, and non-case differences remain on existing exact/no-match behavior.

Verification:

```powershell
go test ./internal/designrecords -run TestRequiredSectionHeadingCaseMismatch -v
go test ./internal/designrecords -run TestProposeRecordUpdateRequiredHeadingCaseFallback -v
go test ./internal/designrecordsmcp -run TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields -v
go test ./internal/designrecords ./internal/designrecordsmcp
```

All commands passed. Package-level MCP boundary JSON-RPC test covers the public `validate_records` response fields, including `section_heading_case_mismatch`, `actual_heading`, `section`, `status`, and `candidate_headings`.
