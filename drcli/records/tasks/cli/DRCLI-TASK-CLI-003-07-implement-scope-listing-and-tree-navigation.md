# DRCLI-TASK-CLI-003-07: Implement scope listing and tree navigation

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 2d
- **depends_on**:
  - DRCLI-TASK-CLI-003-05
  - DRCLI-TASK-CLI-003-06
- **outputs**:
  - drcli/internal/discovery/discovery.go
  - drcli/internal/discovery/tree.go
  - drcli/internal/discovery/discovery_test.go

## Goal

Implement scope listing and tree navigation in a bounded, separately testable Go package.

## Work

### Read

- This Task: `DRCLI-TASK-CLI-003-07`.
  - `drcli/records/spec/design-records-cli/operations/discovery-and-listing/index.md`
  - `drcli/records/spec/design-records-cli/operations/discovery-and-listing/record-scope-discovery.md`
  - `drcli/records/spec/design-records-cli/operations/discovery-and-listing/sequential-record-listing.md`
  - `drcli/records/spec/design-records-cli/operations/discovery-and-listing/tree-child-listing.md`
  - `drcli/records/spec/design-records-cli/operations/discovery-and-listing/tree-overview.md`
- Consume only the completed predecessor exported APIs named below; do not reread siblings or change their code.

### Change

- `drcli/internal/discovery/discovery.go`
- `drcli/internal/discovery/tree.go`
- `drcli/internal/discovery/discovery_test.go`

### Implement

- Fixed private Go API: `discovery.ScopeList(*current.Snapshot,ScopeRequest) (Result,error); discovery.Records(*current.Snapshot,ListRequest) (Result,error); discovery.Children(*current.Snapshot,TreeRequest) (Result,error); discovery.Inspect(*current.Snapshot,TreeInspectRequest) (Result,error).`
- Fixed behavior: Implement selector-validation precedence and only one scope result variant; app/artifact/domain ascending simple string order; sequential default descending tuple order by segment priority, tie canonical ref always ascending, offset then limit, null requested projections. Tree children include segment/has_children/has_record with optional null title/status, terminal-node distinction, 1..500 range; tree inspection depth0..5, limit1..500, breadth-first selection and nested structural rendering with truthful truncated.
- Model routing: Codex/Sonnet: sequence and tree projection.
- Preserve exact canonical refs, error/warning order, null-versus-omitted fields, and accepted PRODUCT authority; no inferred aliases or behavior.

### Test

- Assertions: Test invalid selector combination precedence, empty but available domains, numeric limits, sequence tuple ties, repeated H1 IDs conflict exclusion, terminal tree nodes with no record, BFS budget, depth boundary not truncation and canonical child identity.
- Focused command: `go -C drcli test ./internal/discovery -count=1`; require exit 0.
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
| `drcli/internal/discovery/discovery.go` | Implement scope listing and tree navigation API and its named ownership boundary. | Implement selector-validation precedence and only one scope result variant; app/artifact/domain ascending simple string order; sequential default descending tuple order by segment priority, tie canonical ref always ascending, offset then limit, null requested projections. | `go -C drcli test ./internal/discovery -count=1` |
| `drcli/internal/discovery/tree.go` | Implement scope listing and tree navigation API and its named ownership boundary. | Implement selector-validation precedence and only one scope result variant; app/artifact/domain ascending simple string order; sequential default descending tuple order by segment priority, tie canonical ref always ascending, offset then limit, null requested projections. | `go -C drcli test ./internal/discovery -count=1` |
| `drcli/internal/discovery/discovery_test.go` | Add focused table-driven regression cases. | Test invalid selector combination precedence, empty but available domains, numeric limits, sequence tuple ties, repeated H1 IDs conflict exclusion, terminal tree nodes with no record, BFS budget, depth boundary not truncation and canonical child identity.. | `go -C drcli test ./internal/discovery -count=1` |

## Done condition

- All declared API, behavior, and test-contract rows pass without modifying other owners. `go -C drcli test ./internal/discovery -count=1` exits 0.

## Verification

- Run `go -C drcli test ./internal/discovery -count=1` and inspect only declared changed files and test assertions. T18 owns full integration.

## Evidence

TBD
