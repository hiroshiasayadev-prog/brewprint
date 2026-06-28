# DRMCP-TASK-MCP-009-07: Implement current validation and diagnostics

- **id**: DRMCP-TASK-MCP-009-07
- **status**: not_started
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 2d
- **depends_on**:
  - DRMCP-TASK-MCP-009-04
- **outputs**:
  - drmcp/src/internal/designrecords/validation.go
  - drmcp/src/internal/designrecords/validation_test.go

## Goal

Implement current repository validation, cross-namespace current relation checks, and portable machine-readable diagnostics.

Keep retained specialized spec-format validators and all legacy validation behavior outside W009.

## Work

- Implement only `{}`, `{app_namespace}`, and `{ref}` validation selectors.
- Select current sources, addressable invalid sources, and current conflict groups deterministically.
- Emit accepted source, identity, field, section, relation, and current-conflict diagnostic categories.
- Validate exact current cross-namespace relations against the active current index.
- Emit repository-root-relative portable source locations only on diagnostics that permit them.
- Order and deduplicate diagnostics by the accepted stable key.
- Remove old kind/range selectors, semantic alias checks, and obsolete category projections.
- Exclude W-SPEC-001 per-file detector logic and W-SPEC-002 Topics graph logic.
- Exclude current-to-legacy target lookup and legacy validation subjects.
- Record own-branch targeted acceptance separately from the later P3 integration acceptance.

### Execution slices

| slice | owner model | parallel group | dependency | exact file boundary or inventory method | allowed changes | prohibited changes | commands | expected evidence | escalation condition |
|---|---|---|---|---|---|---|---|---|---|
| S07A current validation | Sonnet | P3 | T04 accepted | `validation.go`, `validation_test.go` only. | Replace selectors, current diagnostics, relations, ordering, and tests. | Shared types, config, parser, index, tools, resolver, authoring, fixtures, specialized spec validators, legacy behavior. | `gofmt -w`; targeted validation tests; full package test. | Selector matrix, category matrix, location examples, deterministic ordering output. | Any missing T03 diagnostic type or T04 index view; stop and escalate upstream. |
| S07B mechanical verification | Haiku | P3 | S07A complete | Same two files; read-only. | No changes. | Diagnostic design or test repair. | `gofmt -d`; targeted tests repeated for ordering; full package test; scoped Git check. | Raw outputs and exact changed paths. | Any nondeterminism, path leakage, or excluded validator behavior; escalate to Sonnet. |

## Done condition

- Validation accepts exactly the three corrected selector shapes.
- Current sequential, current spec, invalid addressable source, and conflict-group subjects are selected deterministically.
- Current relation targets resolve across configured current app roots by exact canonical identity.
- Accepted current diagnostic categories and severities are machine-readable.
- Diagnostic locations are portable and repository-root-relative.
- Normal read path hiding is not weakened by diagnostic source locations.
- Diagnostic ordering and deduplication are deterministic.
- Kind/range selectors, semantic alias checks, and obsolete categories are absent.
- W-SPEC-001, W-SPEC-002, and all legacy validation behavior remain unimplemented.
- W008 cases C09, C14, C15, R07, R08, and R10 are covered.
- Only `validation.go` and `validation_test.go` change.
- Own-branch targeted acceptance is recorded separately from P3 integration acceptance.
- Full package PASS is required before T08 only after T05, T06, and T07 are integrated.

## Verification

Run from repository root:

```powershell
gofmt -d drmcp/src/internal/designrecords/validation.go drmcp/src/internal/designrecords/validation_test.go
go test ./drmcp/src/internal/designrecords -run 'TestValidateRecords' -count=1
go test ./drmcp/src/internal/designrecords -run 'TestValidateRecords.*(Order|Determin|Dedup)' -count=10
go test ./drmcp/src/internal/designrecords -count=1
```

## Evidence

Record:

- selector acceptance and rejection matrix;
- diagnostic category and severity matrix;
- portable source-location examples;
- exact relation validation cases;
- repeated deterministic-order output;
- targeted and full package outputs;
- explicit exclusion of retained spec validators and legacy subjects.
- own-branch targeted acceptance and later P3 integration acceptance as separate evidence.

### Provisional implementation mapping

```yaml
implementation_mapping:
  status: provisional
  contract_refs:
    - DRMCP-REQ-MCP-001
    - DRMCP-ADR-MCP-001
    - DRMCP-WORK-MCP-006
    - DRMCP-WORK-MCP-007
    - DRMCP-WORK-MCP-008
  fixture_cases:
    - C09
    - C14
    - C15
    - R07
    - R08
    - R10
  implementation:
    - path: drmcp/src/internal/designrecords/validation.go
      symbols: []
  verification:
    - path: drmcp/src/internal/designrecords/validation_test.go
      tests: []
  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```

Populate `symbols` and `tests` with real names before Task closure.
Remove any path that does not contain the final contract-significant implementation or verification.
