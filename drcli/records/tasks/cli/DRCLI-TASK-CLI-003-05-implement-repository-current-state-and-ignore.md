# DRCLI-TASK-CLI-003-05: Implement selected-repository current record state

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 3d
- **depends_on**:
  - DRCLI-TASK-CLI-003-04
- **outputs**:
  - drcli/internal/current/source.go
  - drcli/internal/current/ignore.go
  - drcli/internal/current/admission.go
  - drcli/internal/current/scopes.go
  - drcli/internal/current/current_test.go
  - drcli/internal/current/ignore_test.go

## Goal

Implement selected-repository current record state in a bounded, separately testable Go package.

## Work

### Read

- This Task: `DRCLI-TASK-CLI-003-05`.
  - `drcli/records/spec/design-records-cli/cli/repository-selection-and-traversal.md`
  - `drcli/records/spec/design-records-cli/current-record-model/artifact-candidate-and-admission.md`
  - `drcli/records/spec/design-records-cli/current-record-model/current-record-scopes.md`
  - `drcli/records/spec/design-records-cli/current-record-model/current-source-corpus.md`
  - `drcli/records/spec/design-records-cli/current-record-model/identity-conflict-and-addressability.md`
  - `drcli/records/spec/design-records-cli/current-record-model/index.md`
- Consume only the completed predecessor exported APIs named below; do not reread siblings or change their code.

### Change

- `drcli/internal/current/source.go`
- `drcli/internal/current/ignore.go`
- `drcli/internal/current/admission.go`
- `drcli/internal/current/scopes.go`
- `drcli/internal/current/current_test.go`
- `drcli/internal/current/ignore_test.go`

### Implement

- Fixed private Go API: `current.Open(repo string,noIgnore bool) (*Snapshot,error); (*Snapshot).Find(ref string) (model.Record,LookupState); (*Snapshot).Scope(selector ScopeSelector) (Scope,error); (*Snapshot).Tree(ref string) (*model.TreeNode,bool); (*Snapshot).Records() []model.Record; Snapshot exposes SourceIssues and Conflicts with repository-relative member paths.`
- Fixed behavior: Use exact cwd or --repo from caller, never installation path or ancestors. Walk only selected repo with Git-compatible .gitignore and nested ignore pruning before finding directories named records; --no-ignore only skips pruning. Collect namespace evidence only from in-corpus syntactically extractable H1 or spec id. Reject zero/mixed namespace roots and duplicate app roots as request-wide configuration failure. Admit sequential domain-direct Markdown and tree index/leaf records; aggregate claims repo-wide with no collision winner; nonidentity faults retain addressability. Preserve empty app/artifact/domain/tree nodes and missing-index nodes.
- Model routing: Codex/Sonnet: source and admission invariants.
- Preserve exact canonical refs, error/warning order, null-versus-omitted fields, and accepted PRODUCT authority; no inferred aliases or behavior.

### Test

- Assertions: Test ignore negation, nested patterns, root ignored/unignored, no roots empty success, zero/mixed/duplicate roots failure, malformed namespace evidence, H1/domain mismatch, duplicate claims, nonidentity admission, empty tree nodes, provenance and corpus exclusions.
- Focused command: `go -C drcli test ./internal/current -count=1`; require exit 0.
- Aggregate owner: DRCLI-TASK-CLI-003-18. Independent code verdict: DRCLI-TASK-CLI-003-19.

### Stop

- Return BLOCKED with exact missing input, field, symbol, or owning Task for absent predecessor output, a needed out-of-Change edit, a contract conflict, or a failed environment command. Do not infer missing semantics.

### Prohibited operations

- No modifications to PRODUCT, DRMCP, accepted DRCLI Specifications or closed DRCLI-WORK-CLI-002; no other executor's files, no repository-wide traversal, no independent review, no lifecycle synchronization, and no stage/commit.

### Output

- Record changed files, focused command and exit code, blocking diagnostics if any, and handoff to T18.

## Implementation contract

| target | required change | acceptance criterion | verification |
|---|---|---|---|
| `drcli/internal/current/source.go` | Implement selected-repository current record state API and its named ownership boundary. | Use exact cwd or --repo from caller, never installation path or ancestors. | `go -C drcli test ./internal/current -count=1` |
| `drcli/internal/current/ignore.go` | Implement selected-repository current record state API and its named ownership boundary. | Use exact cwd or --repo from caller, never installation path or ancestors. | `go -C drcli test ./internal/current -count=1` |
| `drcli/internal/current/admission.go` | Implement selected-repository current record state API and its named ownership boundary. | Use exact cwd or --repo from caller, never installation path or ancestors. | `go -C drcli test ./internal/current -count=1` |
| `drcli/internal/current/scopes.go` | Implement selected-repository current record state API and its named ownership boundary. | Use exact cwd or --repo from caller, never installation path or ancestors. | `go -C drcli test ./internal/current -count=1` |
| `drcli/internal/current/current_test.go` | Add focused table-driven regression cases. | Test ignore negation, nested patterns, root ignored/unignored, no roots empty success, zero/mixed/duplicate roots failure, malformed namespace evidence, H1/domain mismatch, duplicate claims, nonidentity admission, empty tree nodes, provenance and corpus exclusions.. | `go -C drcli test ./internal/current -count=1` |
| `drcli/internal/current/ignore_test.go` | Add focused table-driven regression cases. | Test ignore negation, nested patterns, root ignored/unignored, no roots empty success, zero/mixed/duplicate roots failure, malformed namespace evidence, H1/domain mismatch, duplicate claims, nonidentity admission, empty tree nodes, provenance and corpus exclusions.. | `go -C drcli test ./internal/current -count=1` |

## Done condition

- All declared API, behavior, and test-contract rows pass without modifying other owners. `go -C drcli test ./internal/current -count=1` exits 0.

## Verification

- Run `go -C drcli test ./internal/current -count=1` and inspect only declared changed files and test assertions. T18 owns full integration.

## Evidence

TBD
