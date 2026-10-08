# DRCLI-TASK-CLI-003-09: Implement Scoped lexical search

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 3d
- **depends_on**:
  - DRCLI-TASK-CLI-003-05
  - DRCLI-TASK-CLI-003-06
- **outputs**:
  - drcli/internal/search/search.go
  - drcli/internal/search/snippet.go
  - drcli/internal/search/search_test.go

## Goal

Implement scoped lexical search in its sole-writer Go package.

## Work

### Read

- `DRCLI-TASK-CLI-003-09` (this Task).
- `drcli/records/spec/design-records-cli/operations/search/index.md`
- `drcli/records/spec/design-records-cli/operations/search/lexical-match-model.md`
- `drcli/records/spec/design-records-cli/operations/search/lexical-record-search.md`
- `drcli/records/spec/design-records-cli/operations/search/search-results-and-limits.md`
- Read predecessor-exported API source only at exact imported paths.

### Change

- `drcli/internal/search/search.go`
- `drcli/internal/search/snippet.go`
- `drcli/internal/search/search_test.go`

### Implement

- Export exactly the logical API: `Search(*current.Snapshot,Request)(Result,error); Request{Query,Mode,CaseSensitive,Scopes,Targets,MatchLimit,SnippetLimit}; Result.Matches with span/truncated flags`.
- Behavioral contract: Apply exact selected scope union to uniquely addressable records; literal Unicode simple case fold or RE2 leftmost-first nonzero patterns. Search title, metadata_source, h2_heading and h2_section_content slices separately; preserve original source text. Return canonical-ref and target/instance/span ordered non-overlapping matches; snippet aligned to full line or centered match (scalar counts), omit matches exceeding snippet cap; per-record cap10 and global 1..200; warn each unprojectable record×target.
- Model: Codex/Sonnet; dense semantic and failure-order boundary, not mechanical code transcription.

### Test

- Cases: Pattern invalid and zero-length, literal case fold, multiline target bounds, snippets with odd centered capacity and CRLF, duplicate scope elimination, warning order, per-record cap, has_additional_matches for oversized matches.
- Focused command: `go -C drcli test ./internal/search -count=1` (exit 0).
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
| `drcli/internal/search/search.go` | Implement declared scoped lexical search symbols and rules. | Apply exact selected scope union to uniquely addressable records; literal Unicode simple case fold or RE2 leftmost-first nonzero patterns. | `go -C drcli test ./internal/search -count=1` |
| `drcli/internal/search/snippet.go` | Implement declared scoped lexical search symbols and rules. | Apply exact selected scope union to uniquely addressable records; literal Unicode simple case fold or RE2 leftmost-first nonzero patterns. | `go -C drcli test ./internal/search -count=1` |
| `drcli/internal/search/search_test.go` | Implement focused regression table. | Cover Pattern invalid and zero-length, literal case fold, multiline target bounds, snippets with odd centered capacity and CRLF, duplicate scope elimination, warning order, per-record cap, has_additional_matches for oversized matches.. | `go -C drcli test ./internal/search -count=1` |

## Done condition

- The fixed behavior and each Implementation contract row pass; `go -C drcli test ./internal/search -count=1` exits 0.

## Verification

- Focused command `go -C drcli test ./internal/search -count=1`, changed-path inspection; T18 owns aggregate tests.

## Evidence

TBD
