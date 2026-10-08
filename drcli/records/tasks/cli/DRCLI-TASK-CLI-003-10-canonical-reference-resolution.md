# DRCLI-TASK-CLI-003-10: Implement Canonical and V01 reference resolution

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 2d
- **depends_on**:
  - DRCLI-TASK-CLI-003-05
  - DRCLI-TASK-CLI-003-06
- **outputs**:
  - drcli/internal/reference/resolve.go
  - drcli/internal/reference/resolve_test.go

## Goal

Implement canonical and v01 reference resolution in its sole-writer Go package.

## Work

### Read

- `DRCLI-TASK-CLI-003-10` (this Task).
- `drcli/records/spec/design-records-cli/operations/reference-resolution/index.md`
- Read predecessor-exported API source only at exact imported paths.

### Change

- `drcli/internal/reference/resolve.go`
- `drcli/internal/reference/resolve_test.go`

### Implement

- Export exactly the logical API: `Resolve(*current.Snapshot,ref string,compat CompatibilityLookup)(Resolution,error); Resolution{Ref,Status,Target}; Target{TargetType,Ref,Kind,Title,Status}`.
- Behavioral contract: Use PRODUCT canonical ref grammar and current-first exact current lookup, then only PRODUCT-approved V01 issued ID families with exact compatibility lookup. Return resolved/unresolved/unsupported, nullable title/status for current, no title/status for legacy, never fuzzy/alias/path/section fallback.
- Model: Codex/Sonnet; dense semantic and failure-order boundary, not mechanical code transcription.

### Test

- Cases: Resolved current, missing current, conflict unresolved, accepted V01 issued ID resolved/unresolved, unsupported prefix/physical path/whitespace, current precedence over compatibility.
- Focused command: `go -C drcli test ./internal/reference -count=1` (exit 0).
- Broader integration belongs solely to T18; independent code review to T19.

### Stop

- Return BLOCKED when required symbols, fixtures, PRODUCT semantics, or Go environment are unavailable or an out-of-Change file must change. Name owning Task, no exploratory cross-repo changes.

### Prohibited operations

- No other executor writes, unrelated refactor, source-spec/PRODUCT/DRMCP/historical record change, repo-wide traversal, stage/commit, self-review, or lifecycle synchronization.

### Output

- Changed paths, focused command/exit, blockers, review handoff.

## Implementation contract

| target | required change | acceptance criterion | verification |
|---|---|---|---|
| `drcli/internal/reference/resolve.go` | Implement declared canonical and v01 reference resolution symbols and rules. | Use PRODUCT canonical ref grammar and current-first exact current lookup, then only PRODUCT-approved V01 issued ID families with exact compatibility lookup. | `go -C drcli test ./internal/reference -count=1` |
| `drcli/internal/reference/resolve_test.go` | Implement focused regression table. | Cover Resolved current, missing current, conflict unresolved, accepted V01 issued ID resolved/unresolved, unsupported prefix/physical path/whitespace, current precedence over compatibility.. | `go -C drcli test ./internal/reference -count=1` |

## Done condition

- The fixed behavior and each Implementation contract row pass; `go -C drcli test ./internal/reference -count=1` exits 0.

## Verification

- Focused command `go -C drcli test ./internal/reference -count=1`, changed-path inspection; T18 owns aggregate tests.

## Evidence

TBD
