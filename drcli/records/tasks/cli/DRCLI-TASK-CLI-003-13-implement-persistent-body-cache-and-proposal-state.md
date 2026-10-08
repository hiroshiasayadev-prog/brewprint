# DRCLI-TASK-CLI-003-13: Implement process-independent retained authoring state

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 3d
- **depends_on**:
  - DRCLI-TASK-CLI-003-04
- **outputs**:
  - drcli/internal/state/state.go
  - drcli/internal/state/state_test.go

## Goal

Implement process-independent retained authoring state with its own non-overlapping writer and focused acceptance boundary.

## Work

### Read

- `DRCLI-TASK-CLI-003-13` (this Task).
- `drcli/records/spec/design-records-cli/operations/authoring/retry-cache.md`
- `drcli/records/spec/design-records-cli/operations/authoring/deferred-proposal.md`
- `drcli/records/spec/design-records-cli/cli/runtime-and-distribution.md`
- Predecessor APIs only at declared imported paths; no parent Work Item or broad repository search.

### Change

- `drcli/internal/state/state.go`
- `drcli/internal/state/state_test.go`

### Implement

- Export fixed logical API: `state.Open(installationRoot string)(*Store,error); (*Store).PutBody([]byte)(CacheID,error); (*Store).GetBody(CacheID)([]byte,error); (*Store).PutProposal(Proposal)(ProposalID,error); (*Store).GetProposal(ProposalID)(Proposal,error); (*Store).SetProposalState(ProposalID,expected,next State)(Proposal,error)`.
- Required behaviors: Persist exact byte-preserving opaque cached body independent of repository, process, session and cwd; retry reuse is non-consuming. Persist proposal repository canonical binding, complete candidate/base-state payload, lifecycle pending/accepted/discarded and validity, atomically guard duplicate acceptance. Distinguish proven unknown/expired/invalid ID from storage I/O; concurrent isolated opaque IDs cannot collide across repos. Persist store outside selected repository and write only own runtime state, never Design Records. No hardcoded normative retention period.
- Model routing: Codex/Sonnet because durable state and atomicity decisions need strong coding judgment.

### Test

- Assertions: Restart two child processes, later cwd/repo independent cache retrieval, byte exact roundtrip UTF-8/CRLF, proposal repository binding, two-repo IDs, corrupt/unknown ID versus I/O failure, atomic accepted/discarded state guard, concurrent unique IDs.
- Focused command: `go -C drcli test ./internal/state -count=1`; exit 0.
- T18 is exclusive aggregate verification owner; T19 independently reviews all code.

### Stop

- BLOCKED on unavailable predecessor input, unknown accepted contract, required out-of-Change write, or nonrunnable test. Identify exact owner and condition; no independent design.

### Prohibited operations

- No PRODUCT/DRMCP/Specification edits, historical review modification, sibling code, arbitrary refactor, repo-wide traversal, stage/commit, or review/closure.

### Output

- Changes, focused test and exit, blockers, integration readiness.

## Implementation contract

| target | required change | acceptance criterion | verification |
|---|---|---|---|
| `drcli/internal/state/state.go` | Implement named API and declared behavior in sole-owned file. | Persist exact byte-preserving opaque cached body independent of repository, process, session and cwd; retry reuse is non-consuming. | `go -C drcli test ./internal/state -count=1` |
| `drcli/internal/state/state_test.go` | Add deterministic regressions for target implementation. | Restart two child processes, later cwd/repo independent cache retrieval, byte exact roundtrip UTF-8/CRLF, proposal repository binding, two-repo IDs, corrupt/unknown ID versus I/O failure, atomic accepted/discarded state guard, concurrent unique IDs.. | `go -C drcli test ./internal/state -count=1` |

## Done condition

- Each implementation contract row is satisfied and `go -C drcli test ./internal/state -count=1` returns 0.

## Verification

- Check `go -C drcli test ./internal/state -count=1` and only the declared changed file list; T18 owns aggregate acceptance.

## Evidence

TBD
