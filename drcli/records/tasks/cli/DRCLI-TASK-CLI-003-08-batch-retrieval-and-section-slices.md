# DRCLI-TASK-CLI-003-08: Implement Exact records and H2 sections

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 2d
- **depends_on**:
  - DRCLI-TASK-CLI-003-05
  - DRCLI-TASK-CLI-003-06
- **outputs**:
  - drcli/internal/retrieval/retrieval.go
  - drcli/internal/retrieval/sections.go
  - drcli/internal/retrieval/retrieval_test.go

## Goal

Implement exact records and h2 sections in its sole-writer Go package.

## Work

### Read

- `DRCLI-TASK-CLI-003-08` (this Task).
- `drcli/records/spec/design-records-cli/operations/retrieval/batch-retrieval-results.md`
- `drcli/records/spec/design-records-cli/operations/retrieval/exact-record-retrieval.md`
- `drcli/records/spec/design-records-cli/operations/retrieval/h2-section-retrieval.md`
- `drcli/records/spec/design-records-cli/operations/retrieval/index.md`
- `drcli/records/spec/design-records-cli/operations/retrieval/retrieval-output-limits.md`
- Read predecessor-exported API source only at exact imported paths.

### Change

- `drcli/internal/retrieval/retrieval.go`
- `drcli/internal/retrieval/sections.go`
- `drcli/internal/retrieval/retrieval_test.go`

### Implement

- Export exactly the logical API: `Get(*current.Snapshot,RecordRequest)(Result,error); Sections(*current.Snapshot,SectionRequest)(Result,error); RecordRequest.Projection{Title,Metadata,H2Headings,SourceContent bool}`.
- Behavioral contract: Preserve 1..100 supplied selectors including duplicates, selector order, partial errors and projection omissions. Use exact H2 first-match with duplicate_h2 warning; return original source slice. Apply 1..16777216 UTF-8 byte aggregate content budget in selector order, skip oversize whole content with null and warning without consuming budget, continue later requests. Source limit is ignored with warning when source projection false.
- Model: Codex/Sonnet; dense semantic and failure-order boundary, not mechanical code transcription.

### Test

- Cases: Duplicate/unresolved/conflicted refs; additional # in section titles; first duplicate H2 and warnings; unavailable metadata/H2 projection; all-false projection; CRLF/UTF-8 preservation; skip oversized then accept smaller content; count and limit errors.
- Focused command: `go -C drcli test ./internal/retrieval -count=1` (exit 0).
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
| `drcli/internal/retrieval/retrieval.go` | Implement declared exact records and h2 sections symbols and rules. | Preserve 1..100 supplied selectors including duplicates, selector order, partial errors and projection omissions. | `go -C drcli test ./internal/retrieval -count=1` |
| `drcli/internal/retrieval/sections.go` | Implement declared exact records and h2 sections symbols and rules. | Preserve 1..100 supplied selectors including duplicates, selector order, partial errors and projection omissions. | `go -C drcli test ./internal/retrieval -count=1` |
| `drcli/internal/retrieval/retrieval_test.go` | Implement focused regression table. | Cover Duplicate/unresolved/conflicted refs. | `go -C drcli test ./internal/retrieval -count=1` |

## Done condition

- The fixed behavior and each Implementation contract row pass; `go -C drcli test ./internal/retrieval -count=1` exits 0.

## Verification

- Focused command `go -C drcli test ./internal/retrieval -count=1`, changed-path inspection; T18 owns aggregate tests.

## Evidence

TBD
