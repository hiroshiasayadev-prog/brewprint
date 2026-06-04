# TASK-MCP-019-05: Run runtime smoke and close synchronization

- **id**: TASK-MCP-019-05
- **status**: todo
- **date**: 2026-06-05
- **work_item**: WORK-MCP-019
- **source_requirement**: REQ-MCP-021
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-MCP-019-04
- **outputs**:
  - Runtime smoke evidence for required heading case canonicalization
  - Close synchronization updates for the task, work item, and requirement
  - Final validation evidence before commit

## Goal

Complete runtime smoke and close synchronization for REQ-MCP-021 / WORK-MCP-019 after implementation.

## Work

- Run targeted runtime smoke for `propose_record_update` required-heading case-only fallback through the MCP boundary if feasible.
- Confirm `validate_records` public response exposes `section_heading_case_mismatch` with `actual_heading` when applicable.
- Confirm no broad fuzzy matching or optional/cross-kind heading canonicalization behavior was introduced.
- Run final validation for related workflow artifacts.
- Record smoke and validation evidence.
- Close this task, the parent work item, and the source requirement when evidence is complete.

## Done condition

- Runtime smoke demonstrates case-only required-heading repair through proposal flow or records why package-level tests are sufficient.
- Public validation diagnostic behavior is confirmed.
- `go test ./internal/designrecords ./internal/designrecordsmcp` is green.
- This task, the parent work item, and the source requirement are status-synchronized.
- Evidence is recorded before final commit.

## Verification

Recommended commands:

```powershell
cd C:\Users\imved\projects\brewprint
go test ./internal/designrecords -run TestRequiredSectionHeadingCaseMismatch -v
go test ./internal/designrecords -run TestProposeRecordUpdateRequiredHeadingCaseFallback -v
go test ./internal/designrecordsmcp -run TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields -v
go test ./internal/designrecords ./internal/designrecordsmcp
```

If MCP runtime smoke is executed manually, record the exact command and response summary in Evidence.

## Evidence

未実施。
