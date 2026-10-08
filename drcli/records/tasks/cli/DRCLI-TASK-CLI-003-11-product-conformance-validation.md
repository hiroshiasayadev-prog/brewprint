# DRCLI-TASK-CLI-003-11: Implement Broad and exact PRODUCT-owned validation

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 3d
- **depends_on**:
  - DRCLI-TASK-CLI-003-04
  - DRCLI-TASK-CLI-003-05
  - DRCLI-TASK-CLI-003-06
- **outputs**:
  - drcli/internal/validation/rules.go
  - drcli/internal/validation/scope.go
  - drcli/internal/validation/records.go
  - drcli/internal/validation/validation_test.go
  - drcli/internal/validation/rules_test.go

## Goal

Implement broad and exact product-owned validation in its sole-writer Go package.

## Work

### Read

- `DRCLI-TASK-CLI-003-11` (this Task).
- `drcli/records/spec/design-records-cli/operations/validation/exact-record-validation.md`
- `drcli/records/spec/design-records-cli/operations/validation/index.md`
- `drcli/records/spec/design-records-cli/operations/validation/scope-validation.md`
- `drcli/records/spec/design-records-cli/operations/validation/validation-output-limits.md`
- Read predecessor-exported API source only at exact imported paths.

### Change

- `drcli/internal/validation/rules.go`
- `drcli/internal/validation/scope.go`
- `drcli/internal/validation/records.go`
- `drcli/internal/validation/validation_test.go`
- `drcli/internal/validation/rules_test.go`

### Implement

- Export exactly the logical API: `ValidateRecord(*current.Snapshot,model.Record) []Finding; Scope(*current.Snapshot,ScopeRequest)(ScopeResult,error); Records(*current.Snapshot,RecordsRequest)(RecordsResult,error); Finding{Classification,Code,Message,ContractRef,Subject}`.
- Behavioral contract: Consume PRODUCT canonical artifact/relations/semantic refs and DRCLI artifact projections; do not invent DRCLI finding codes. Scope handles repository/app/kind/domain/subtree, includes unadmitted sources/conflicts/nonrecord missing-index advisories; complete summary before limit with abnormal variant order. Exact refs preserve first occurrence, aggregate duplicates as warnings and maintain original selector order, canonical finding order and full summaries; output bounds 1..2000, no partial finding.
- Model: Codex/Sonnet; dense semantic and failure-order boundary, not mechanical code transcription.

### Test

- Cases: Rule codes and subjects from six artifact kinds, invalid/empty scopes, unadmitted and conflicted records, orphan tree index advisory, complete summary independent of limits, deduplicated refs warnings, duplicate warnings in first-occurrence order, exact findings and limit semantics.
- Focused command: `go -C drcli test ./internal/validation -count=1` (exit 0).
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
| `drcli/internal/validation/rules.go` | Implement declared broad and exact product-owned validation symbols and rules. | Consume PRODUCT canonical artifact/relations/semantic refs and DRCLI artifact projections; do not invent DRCLI finding codes. | `go -C drcli test ./internal/validation -count=1` |
| `drcli/internal/validation/scope.go` | Implement declared broad and exact product-owned validation symbols and rules. | Consume PRODUCT canonical artifact/relations/semantic refs and DRCLI artifact projections; do not invent DRCLI finding codes. | `go -C drcli test ./internal/validation -count=1` |
| `drcli/internal/validation/records.go` | Implement declared broad and exact product-owned validation symbols and rules. | Consume PRODUCT canonical artifact/relations/semantic refs and DRCLI artifact projections; do not invent DRCLI finding codes. | `go -C drcli test ./internal/validation -count=1` |
| `drcli/internal/validation/validation_test.go` | Implement focused regression table. | Cover Rule codes and subjects from six artifact kinds, invalid/empty scopes, unadmitted and conflicted records, orphan tree index advisory, complete summary independent of limits, deduplicated refs warnings, duplicate warnings in first-occurrence order, exact findings and limit semantics.. | `go -C drcli test ./internal/validation -count=1` |
| `drcli/internal/validation/rules_test.go` | Implement focused regression table. | Cover Rule codes and subjects from six artifact kinds, invalid/empty scopes, unadmitted and conflicted records, orphan tree index advisory, complete summary independent of limits, deduplicated refs warnings, duplicate warnings in first-occurrence order, exact findings and limit semantics.. | `go -C drcli test ./internal/validation -count=1` |

## Done condition

- The fixed behavior and each Implementation contract row pass; `go -C drcli test ./internal/validation -count=1` exits 0.

## Verification

- Focused command `go -C drcli test ./internal/validation -count=1`, changed-path inspection; T18 owns aggregate tests.

## Evidence

TBD
