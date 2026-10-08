# DRCLI-TASK-CLI-003-15: Implement DRCLI command dispatch, help and transport

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 3d
- **depends_on**:
  - DRCLI-TASK-CLI-003-04
  - DRCLI-TASK-CLI-003-05
  - DRCLI-TASK-CLI-003-06
  - DRCLI-TASK-CLI-003-07
  - DRCLI-TASK-CLI-003-08
  - DRCLI-TASK-CLI-003-09
  - DRCLI-TASK-CLI-003-10
  - DRCLI-TASK-CLI-003-11
  - DRCLI-TASK-CLI-003-12
  - DRCLI-TASK-CLI-003-14
- **outputs**:
  - drcli/cmd/drcli/main.go
  - drcli/internal/cli/cli.go
  - drcli/internal/cli/parse.go
  - drcli/internal/cli/help.go
  - drcli/internal/cli/render.go
  - drcli/internal/cli/cli_test.go

## Goal

Implement DRCLI command dispatch, help and transport with its own non-overlapping writer and focused acceptance boundary.

## Work

### Read

- `DRCLI-TASK-CLI-003-15` (this Task).
- `drcli/records/spec/design-records-cli/cli/authoring-body-input.md`
- `drcli/records/spec/design-records-cli/cli/command-surface.md`
- `drcli/records/spec/design-records-cli/cli/help-and-usage-diagnostics.md`
- `drcli/records/spec/design-records-cli/cli/index.md`
- `drcli/records/spec/design-records-cli/cli/output-and-exit-status.md`
- `drcli/records/spec/design-records-cli/cli/repository-selection-and-traversal.md`
- `drcli/records/spec/design-records-cli/cli/runtime-and-distribution.md`
- Predecessor APIs only at declared imported paths; no parent Work Item or broad repository search.

### Change

- `drcli/cmd/drcli/main.go`
- `drcli/internal/cli/cli.go`
- `drcli/internal/cli/parse.go`
- `drcli/internal/cli/help.go`
- `drcli/internal/cli/render.go`
- `drcli/internal/cli/cli_test.go`

### Implement

- Export fixed logical API: `cli.Run(args []string, stdin io.Reader, stdout,stderr io.Writer, cwd,executablePath string) int; cli.Parse(args []string)(Invocation,UsageProblem); cli.Render(Outcome,format string)([]byte,int,bool); cmd/drcli.main delegates to Run.`.
- Required behaviors: Expose exactly seven root groups and every canonical command including five proposal commands; command and group --help with valid examples and exact route, no repository/standards/state access for help. Parse concrete flags --repo, --no-ignore, --format human/json, --body, --body-stdin, --body-cache-id; no stdin reads without explicit flag and mutually exclusive body forms. Unknown tokens/options, missing operands, invalid selector combinations return actionable correct alternatives and help command (status2); preserve operation type-safe request semantics. One JSON document stdout for trustworthy success or semantic negative outcome, one stderr document for untrustworthy failure; UTF-8 streams, status0/1/2/3/4, no TTY inference. On argv transport ambiguity, reject with actionable usage rather than inventing semantic aliases.
- Model routing: Codex/Sonnet: cross-package integration and exact usage semantics.

### Test

- Assertions: Root/group/command help without resources, exact all 19 command dispatches, unknown group/flag and help routes, invalid/absent/duplicate body sources, stdin opt-in, output/status 0..4 and one JSON document/stream, cwd default vs --repo, option parity.
- Focused command: `go -C drcli test ./internal/cli -count=1`; exit 0.
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
| `drcli/cmd/drcli/main.go` | Implement named API and declared behavior in sole-owned file. | Expose exactly seven root groups and every canonical command including five proposal commands; command and group --help with valid examples and exact route, no repository/standards/state access for help. | `go -C drcli test ./internal/cli -count=1` |
| `drcli/internal/cli/cli.go` | Implement named API and declared behavior in sole-owned file. | Expose exactly seven root groups and every canonical command including five proposal commands; command and group --help with valid examples and exact route, no repository/standards/state access for help. | `go -C drcli test ./internal/cli -count=1` |
| `drcli/internal/cli/parse.go` | Implement named API and declared behavior in sole-owned file. | Expose exactly seven root groups and every canonical command including five proposal commands; command and group --help with valid examples and exact route, no repository/standards/state access for help. | `go -C drcli test ./internal/cli -count=1` |
| `drcli/internal/cli/help.go` | Implement named API and declared behavior in sole-owned file. | Expose exactly seven root groups and every canonical command including five proposal commands; command and group --help with valid examples and exact route, no repository/standards/state access for help. | `go -C drcli test ./internal/cli -count=1` |
| `drcli/internal/cli/render.go` | Implement named API and declared behavior in sole-owned file. | Expose exactly seven root groups and every canonical command including five proposal commands; command and group --help with valid examples and exact route, no repository/standards/state access for help. | `go -C drcli test ./internal/cli -count=1` |
| `drcli/internal/cli/cli_test.go` | Add deterministic regressions for target implementation. | Root/group/command help without resources, exact all 19 command dispatches, unknown group/flag and help routes, invalid/absent/duplicate body sources, stdin opt-in, output/status 0..4 and one JSON document/stream, cwd default vs --repo, option parity.. | `go -C drcli test ./internal/cli -count=1` |

## Done condition

- Each implementation contract row is satisfied and `go -C drcli test ./internal/cli -count=1` returns 0.

## Verification

- Check `go -C drcli test ./internal/cli -count=1` and only the declared changed file list; T18 owns aggregate acceptance.

## Evidence

TBD
