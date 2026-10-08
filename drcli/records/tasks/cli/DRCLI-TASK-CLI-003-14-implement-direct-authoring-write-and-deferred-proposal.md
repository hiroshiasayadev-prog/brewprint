# DRCLI-TASK-CLI-003-14: Implement create, update, guarded write and proposal actions

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 3d
- **depends_on**:
  - DRCLI-TASK-CLI-003-04
  - DRCLI-TASK-CLI-003-05
  - DRCLI-TASK-CLI-003-06
  - DRCLI-TASK-CLI-003-11
  - DRCLI-TASK-CLI-003-13
- **outputs**:
  - drcli/internal/authoring/prepare.go
  - drcli/internal/authoring/write.go
  - drcli/internal/authoring/proposal.go
  - drcli/internal/authoring/authoring_test.go
  - drcli/internal/authoring/proposal_test.go

## Goal

Implement create, update, guarded write and proposal actions with its own non-overlapping writer and focused acceptance boundary.

## Work

### Read

- `DRCLI-TASK-CLI-003-14` (this Task).
- `drcli/records/spec/design-records-cli/operations/authoring/create.md`
- `drcli/records/spec/design-records-cli/operations/authoring/deferred-proposal.md`
- `drcli/records/spec/design-records-cli/operations/authoring/index.md`
- `drcli/records/spec/design-records-cli/operations/authoring/retry-cache.md`
- `drcli/records/spec/design-records-cli/operations/authoring/update.md`
- `drcli/records/spec/design-records-cli/operations/authoring/write.md`
- `drcli/records/spec/design-records-cli/operations/validation/exact-record-validation.md`
- `drcli/records/spec/design-records-cli/diagnostics/operation-diagnostic-catalog.md`
- Predecessor APIs only at declared imported paths; no parent Work Item or broad repository search.

### Change

- `drcli/internal/authoring/prepare.go`
- `drcli/internal/authoring/write.go`
- `drcli/internal/authoring/proposal.go`
- `drcli/internal/authoring/authoring_test.go`
- `drcli/internal/authoring/proposal_test.go`

### Implement

- Export fixed logical API: `authoring.Create(ctx Context, req CreateRequest)(Outcome,error); authoring.Update(ctx Context,req UpdateRequest)(Outcome,error); authoring.ProposalCreate/ProposalUpdate/Get/Accept/Discard(ctx Context,req ProposalRequest)(Outcome,error); Context{Snapshot,Validator,Store,Repository}; Outcome{TargetRef,WriteState,ModifiedRefs,BodyCacheID,ProposalID,Diagnostics}`.
- Required behaviors: PRODUCT owns six record authoring fields, IDs, paths, headings, lifecycle and coupled relations. Direct create/update writes same invocation by default; valid no-op is nonwriting; deferred prepare explicit-only. Apply final combined partial sections and fields, reject conflicts and duplicate/missing H2 selection, enforce complete affected-set candidate validation with PRODUCT finding order, then ordered repository/identity/staleness/collision/affected-set/revalidation guards immediately before any write. Coupled PRODUCT-atomic multi-record change cannot partial-success. Capture actual modified/unknown state truthfully; never force overwrite. On retryable body-bearing failure retain exact received body ID or nested cache_error, not two diagnostics. Bind proposals to creation repo; reject retarget/accepted/discarded and stale state.
- Model routing: Codex/Sonnet: transaction semantics and cross-record atomicity are high judgment.

### Test

- Assertions: Create immediate write; update no-op; combined partial update and preserved omitted fields; coupled Work Item Task relation; candidate-only findings; stale/collision and pre-write no-modification; proposal create/read/discard/accept restart and reject reuse; cache preservation success/fail; unknown write-state if I/O uncertain; exact 12 catalog codes and status distinction.
- Focused command: `go -C drcli test ./internal/authoring -count=1`; exit 0.
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
| `drcli/internal/authoring/prepare.go` | Implement named API and declared behavior in sole-owned file. | PRODUCT owns six record authoring fields, IDs, paths, headings, lifecycle and coupled relations. | `go -C drcli test ./internal/authoring -count=1` |
| `drcli/internal/authoring/write.go` | Implement named API and declared behavior in sole-owned file. | PRODUCT owns six record authoring fields, IDs, paths, headings, lifecycle and coupled relations. | `go -C drcli test ./internal/authoring -count=1` |
| `drcli/internal/authoring/proposal.go` | Implement named API and declared behavior in sole-owned file. | PRODUCT owns six record authoring fields, IDs, paths, headings, lifecycle and coupled relations. | `go -C drcli test ./internal/authoring -count=1` |
| `drcli/internal/authoring/authoring_test.go` | Add deterministic regressions for target implementation. | Create immediate write. | `go -C drcli test ./internal/authoring -count=1` |
| `drcli/internal/authoring/proposal_test.go` | Add deterministic regressions for target implementation. | Create immediate write. | `go -C drcli test ./internal/authoring -count=1` |

## Done condition

- Each implementation contract row is satisfied and `go -C drcli test ./internal/authoring -count=1` returns 0.

## Verification

- Check `go -C drcli test ./internal/authoring -count=1` and only the declared changed file list; T18 owns aggregate acceptance.

## Evidence

TBD
