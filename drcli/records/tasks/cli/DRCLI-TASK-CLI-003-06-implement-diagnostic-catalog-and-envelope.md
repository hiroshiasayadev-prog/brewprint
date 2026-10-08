# DRCLI-TASK-CLI-003-06: Implement closed operation diagnostic catalog

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 2d
- **depends_on**:
  - DRCLI-TASK-CLI-003-04
- **outputs**:
  - drcli/internal/diagnostic/catalog.go
  - drcli/internal/diagnostic/envelope.go
  - drcli/internal/diagnostic/diagnostic_test.go

## Goal

Implement closed operation diagnostic catalog in a bounded, separately testable Go package.

## Work

### Read

- This Task: `DRCLI-TASK-CLI-003-06`.
  - `drcli/records/spec/design-records-cli/diagnostics/index.md`
  - `drcli/records/spec/design-records-cli/diagnostics/operation-diagnostic-catalog.md`
  - `drcli/records/spec/design-records-cli/diagnostics/operation-diagnostic-envelope.md`
- Consume only the completed predecessor exported APIs named below; do not reread siblings or change their code.

### Change

- `drcli/internal/diagnostic/catalog.go`
- `drcli/internal/diagnostic/envelope.go`
- `drcli/internal/diagnostic/diagnostic_test.go`

### Implement

- Fixed private Go API: `diagnostic.Code, diagnostic.Item{Code,Message,Context}; diagnostic.Error(code,context), diagnostic.Warning(code,context); Validate(Item,Placement) error; diagnostic.Failure{Kind,Item,WriteState,ModifiedRefs}; diagnostic.OutcomeKind.Success/Contract/Usage/Execution/Internal.`
- Fixed behavior: Freeze every cataloged error/warning code, placement, exact required/forbidden context and variant discriminator; enforce no context:null/{}, severity, operation or unregistered code. Preserve first request-wide error and selector-level warning placement. Authoring uses one top-level error, sorted affected_findings groups by their own record_ref and within-record PRODUCT ordering; write_state truth, status1 stdout versus status3 stderr, exclusive retained body_cache_id/cache_error.
- Model routing: Codex/Sonnet: context variants and negative-result placement.
- Preserve exact canonical refs, error/warning order, null-versus-omitted fields, and accepted PRODUCT authority; no inferred aliases or behavior.

### Test

- Assertions: Test all catalog codes and numeric variant contexts, unknown context rejection, nil omission, invalid mixed response, 12 authoring codes, nested-only cache error, three-group independently sorted affected_findings, modified/nonmodified/unknown write-state, warning ordering.
- Focused command: `go -C drcli test ./internal/diagnostic -count=1`; require exit 0.
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
| `drcli/internal/diagnostic/catalog.go` | Implement closed operation diagnostic catalog API and its named ownership boundary. | Freeze every cataloged error/warning code, placement, exact required/forbidden context and variant discriminator; enforce no context:null/{}, severity, operation or unregistered code. | `go -C drcli test ./internal/diagnostic -count=1` |
| `drcli/internal/diagnostic/envelope.go` | Implement closed operation diagnostic catalog API and its named ownership boundary. | Freeze every cataloged error/warning code, placement, exact required/forbidden context and variant discriminator; enforce no context:null/{}, severity, operation or unregistered code. | `go -C drcli test ./internal/diagnostic -count=1` |
| `drcli/internal/diagnostic/diagnostic_test.go` | Add focused table-driven regression cases. | Test all catalog codes and numeric variant contexts, unknown context rejection, nil omission, invalid mixed response, 12 authoring codes, nested-only cache error, three-group independently sorted affected_findings, modified/nonmodified/unknown write-state, warning ordering.. | `go -C drcli test ./internal/diagnostic -count=1` |

## Done condition

- All declared API, behavior, and test-contract rows pass without modifying other owners. `go -C drcli test ./internal/diagnostic -count=1` exits 0.

## Verification

- Run `go -C drcli test ./internal/diagnostic -count=1` and inspect only declared changed files and test assertions. T18 owns full integration.

## Evidence

TBD
