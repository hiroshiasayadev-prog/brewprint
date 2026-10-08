# DRCLI-TASK-CLI-003-12: Implement portable authoring guidance

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 2d
- **depends_on**:
  - DRCLI-TASK-CLI-003-04
  - DRCLI-TASK-CLI-003-05
  - DRCLI-TASK-CLI-003-06
- **outputs**:
  - drcli/internal/guidance/guidance.go
  - drcli/internal/guidance/guidance_test.go

## Goal

Implement portable authoring guidance with its own non-overlapping writer and focused acceptance boundary.

## Work

### Read

- `DRCLI-TASK-CLI-003-12` (this Task).
- `drcli/records/spec/design-records-cli/operations/guidance/exact-read.md`
- `drcli/records/spec/design-records-cli/operations/guidance/index.md`
- `drcli/records/spec/design-records-cli/operations/guidance/list.md`
- Predecessor APIs only at declared imported paths; no parent Work Item or broad repository search.

### Change

- `drcli/internal/guidance/guidance.go`
- `drcli/internal/guidance/guidance_test.go`

### Implement

- Export fixed logical API: `guidance.List(packageRoot string)([]Entry,error); guidance.Get(packageRoot,id string)(Entry,error); Entry{ID,Title,Abstract,Content}`.
- Required behaviors: Use only bundled design_records Specs, never repository guidance or a separate legacy guide source. List only exact children of spec:design_records.authoring_standards excluding root, sorted by canonical ref ASCII. Require readable first H1, What this is abstract and whole source per candidate; fail entire catalog on conflict/unreadable. Get one exact child only, verbatim untruncated Markdown; no inferred basename/alias/fuzzy fallback. Resolve package root supplied by executable-relative CLI, not cwd.
- Model routing: Codex/Sonnet because package semantics and all-or-nothing availability are significant.

### Test

- Assertions: Canonical child filtering, ASCII order, non-index leaf, exact root/external ID rejection, invalid source must fail whole list, verbatim CRLF bytes, empty package scope, missing package error.
- Focused command: `go -C drcli test ./internal/guidance -count=1`; exit 0.
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
| `drcli/internal/guidance/guidance.go` | Implement named API and declared behavior in sole-owned file. | Use only bundled design_records Specs, never repository guidance or a separate legacy guide source. | `go -C drcli test ./internal/guidance -count=1` |
| `drcli/internal/guidance/guidance_test.go` | Add deterministic regressions for target implementation. | Canonical child filtering, ASCII order, non-index leaf, exact root/external ID rejection, invalid source must fail whole list, verbatim CRLF bytes, empty package scope, missing package error.. | `go -C drcli test ./internal/guidance -count=1` |

## Done condition

- Each implementation contract row is satisfied and `go -C drcli test ./internal/guidance -count=1` returns 0.

## Verification

- Check `go -C drcli test ./internal/guidance -count=1` and only the declared changed file list; T18 owns aggregate acceptance.

## Evidence

TBD
