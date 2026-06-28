# DRMCP-TASK-MCP-009-02: Implement configured current-root loading

- **id**: DRMCP-TASK-MCP-009-02
- **status**: not_started
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-009-01
- **outputs**:
  - drmcp/src/internal/designrecords/config.go
  - drmcp/src/internal/designrecords/config_test.go

## Goal

Replace automatic and V01-default root discovery with explicit configured current roots and app association.

Fail the complete current-index startup when any mandatory current root is invalid.

## Work

- Implement explicit repository-root and non-empty current-root configuration.
- Require one `app_namespace` and one repository-relative `records_root` per current root.
- Validate root containment, exact app-root shape, directory existence, readability, duplicate declarations, and duplicate app namespaces.
- Preserve valid empty records trees.
- Remove `*/records` auto-discovery and `v01/records` fallback.
- Keep legacy-root configuration and behavior outside this Task.

### Execution slices

| slice | owner model | parallel group | dependency | exact file boundary or inventory method | allowed changes | prohibited changes | commands | expected evidence | escalation condition |
|---|---|---|---|---|---|---|---|---|---|
| S02A implementation | Sonnet | P1 | T01 accepted | `config.go`, `config_test.go` only. | Replace current configuration structures and tests. | `types.go`, parsers, index, tools, legacy roots, authoring files, fixtures. | `gofmt -w` on the two files; targeted Go test below; package compile. | Config test names, exact accepted and rejected states, command output, `full_package_gate: deferred_to_T04`. | Any required shared type outside `config.go`; return to T01/T03 boundary review. |
| S02B verification | Haiku | P1 | S02A complete | Same two files; read-only verification. | No file changes. | Contract interpretation or test repair. | `gofmt -d`; targeted Go test; package compile; scoped Git check. | Raw pass/fail output, changed-file list, and deferred full-package gate note. | Any failure or unexpected changed path; stop and escalate to Sonnet. |

## Done condition

- No automatic current-root discovery remains.
- No default `v01/records` current root remains.
- Missing or empty current-root configuration is rejected.
- Every configured current root has an explicit app namespace.
- Invalid, escaped, non-directory, unreadable, duplicate, or app-mismatched current roots fail the complete configuration.
- A valid empty current root succeeds.
- W008 cases C08, C10, R08, and R10 are covered without implementing legacy behavior.
- Only `config.go` and `config_test.go` change.
- Targeted tests pass before individual closure.
- The package compiles before individual closure.
- Full-package PASS is deferred to T04.

## Verification

Run from repository root:

```powershell
gofmt -d drmcp/src/internal/designrecords/config.go drmcp/src/internal/designrecords/config_test.go
go test ./drmcp/src/internal/designrecords -run 'Test(NewConfig|NormalizeConfig|CurrentRoot)' -count=1
go test ./drmcp/src/internal/designrecords -run '^$' -count=1
```

Do not require full-package PASS for T02 individual acceptance.
Record `full_package_gate: deferred_to_T04`.

## Evidence

Record:

- exact changed files;
- removed auto-discovery and fallback behavior;
- accepted current-root cases;
- rejected root cases;
- targeted and package-compile command outputs;
- `full_package_gate: deferred_to_T04`;
- scoped whitespace result;
- any residual configuration serialization limitation.

No evidence is accepted from legacy-root loading or fallback behavior.

### Provisional implementation mapping

```yaml
implementation_mapping:
  status: provisional

  contract_refs:
    - DRMCP-REQ-MCP-001
    - DRMCP-ADR-MCP-001
    - DRMCP-WORK-MCP-003
    - DRMCP-WORK-MCP-008

  fixture_cases:
    - C08
    - C10
    - R08
    - R10

  implementation:
    - path: drmcp/src/internal/designrecords/config.go
      symbols: []

  verification:
    - path: drmcp/src/internal/designrecords/config_test.go
      tests: []

  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```

Populate `symbols` and `tests` with real names before Task closure.
Remove any path that does not contain the final contract-significant implementation or verification.
