# V01-TASK-MCP-017-03: Implement exact ID gap warning diagnostics and tests

- **id**: V01-TASK-MCP-017-03
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-017
- **source_requirement**: V01-REQ-MCP-018
- **estimate**: 1d-2d
- **depends_on**:
  - V01-TASK-MCP-017-02
- **outputs**:
  - Proposal-time exact ID gap warning implementation
  - Regression tests for WORK and TASK exact ID gap warning cases
  - Regression tests showing exact ID create remains allowed

## Goal

Implement non-blocking proposal feedback for exact ID create requests that may create workflow artifact sequence gaps.

## Work

- Add gap-warning detection in the propose_record_create path for workflow artifacts.
- Emit a warning diagnostic or proposal note without rejecting the proposal.
- Cover WORK domain sequence cases.
- Cover TASK parent work item sequence cases.
- Cover any REQ behavior decided by V01-TASK-MCP-017-01.
- Ensure new placeholder create behavior is unchanged.

## Done condition

Regression tests prove that risky exact ID create returns warning feedback while acceptance remains possible and new placeholder create remains unaffected.

## Verification

- go test ./internal/designrecords ./internal/designrecordsmcp
- Targeted tests for authoring create warnings

## Evidence
Claude Code implementation completed and passed review for V01-TASK-MCP-017-03.

Files changed:

- `internal/designrecords/types.go`
- `internal/designrecords/authoring.go`
- `internal/designrecords/authoring_test.go`
- `internal/designrecordsmcp/tools_call_test.go`

Implementation summary:

- Added `DiagnosticExactIDSequenceGap` with category `exact_id_sequence_gap`.
- Updated `ProposeRecordCreate` gating to block only error diagnostics, allowing info diagnostics to produce proposals.
- Added exact ID sequence-gap detection for workflow artifact create.
- REQ and WORK warning scope is same kind plus same domain.
- TASK warning scope is same domain plus parent work item sequence.
- ADR / decision create remains outside the warning scope.
- `*-new` placeholder create remains unchanged and does not emit the warning.
- Exact ID create that fills an existing gap does not emit the warning.

Diagnostic placement confirmed:

- `exact_id_sequence_gap` is returned in proposal-level `diagnostics`.
- `exact_id_sequence_gap` is not returned in `validation.diagnostics`.
- Warning severity is `info`.
- Proposal creation and acceptance remain non-blocking for this warning.

Tests added / updated:

- Added `TestExactIDSequenceGapWarning` with 11 subtests covering WORK gap/no-gap/fill-gap/first-record cases, REQ gap/no-gap, TASK gap/no-gap, WORK-new no-check, and ADR exact skip outside scope.
- Added `TestToolsCallProposeExactWorkIDGapDiagnostic` for MCP-layer response behavior, including proposal creation, top-level diagnostic placement, validation diagnostic separation, and accept success.

Commands run by Claude Code:

- `go test ./internal/designrecords ./internal/designrecordsmcp` — PASS
- `go test ./...` — PASS
- `go test -run TestExactIDSequenceGapWarning|TestToolsCallProposeExactWorkIDGapDiagnostic -v` — PASS

Notes:

- `prepareUpdate` still gates on any diagnostic; this is out of scope because `exact_id_sequence_gap` applies only to create.
- V01-TASK-MCP-017-03 done condition is satisfied. Proceed to V01-TASK-MCP-017-04 for runtime smoke and close synchronization.
