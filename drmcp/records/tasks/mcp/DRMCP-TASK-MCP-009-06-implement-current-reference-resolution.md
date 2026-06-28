# DRMCP-TASK-MCP-009-06: Implement current reference resolution

- **id**: DRMCP-TASK-MCP-009-06
- **status**: not_started
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-009-04
- **outputs**:
  - drmcp/src/internal/designrecords/resolver.go
  - drmcp/src/internal/designrecords/resolve_reference_test.go

## Goal

Implement exact current canonical reference classification and resolution against the active current index.

Return path-free resolved, unresolved, or unsupported results without aliasing, normalization, repair, or legacy fallback.

## Work

- Classify exact app-aware current sequential IDs and exact path-derived current `spec:` refs.
- Resolve only unique usable current targets.
- Return unresolved for accepted current grammar with no target or a current conflict group.
- Return unsupported for bare, path-shaped, partial, malformed, or otherwise unsupported inputs.
- Project `target_type`, canonical `ref`, kind, title, and status without physical paths.
- Remove semantic alias, section alias, normalized-ID, and mixed V01 resolution behavior.
- Leave the legacy fallback seam unimplemented for W010.
- Record own-branch targeted acceptance separately from the later P3 integration acceptance.

### Execution slices

| slice | owner model | parallel group | dependency | exact file boundary or inventory method | allowed changes | prohibited changes | commands | expected evidence | escalation condition |
|---|---|---|---|---|---|---|---|---|---|
| S06A current resolver | Sonnet | P3 | T04 accepted | `resolver.go`, `resolve_reference_test.go` only. | Replace current classification, lookup, result projection, and tests. | Shared types, config, parser, index, tools, validation, authoring, fixtures, legacy fallback. | `gofmt -w`; targeted resolver tests; full package test. | Classification matrix, unique/conflict results, no-path JSON proof, command output. | Any missing T03 type or T04 index view; stop and escalate rather than editing upstream files. |
| S06B mechanical verification | Haiku | P3 | S06A complete | Same two files; read-only. | No changes. | Resolver repair or interpretation. | `gofmt -d`; targeted and full package tests; scoped Git check. | Raw outputs and exact changed paths. | Any failure, legacy behavior, or path field; escalate to Sonnet. |

## Done condition

- Exact current sequential and spec refs resolve only through the current active index.
- Unique usable current target returns `resolved`.
- Missing accepted current target returns `unresolved`.
- Current identity conflict returns `unresolved` and never selects a winner.
- Bare IDs, physical paths, partial prefixes, and repair candidates return `unsupported`.
- Resolved targets contain no physical path.
- No semantic alias or section alias behavior remains.
- No legacy source is loaded and no legacy fallback executes.
- W008 cases C06, C11, C13, R02-R05, and R14 are covered.
- Only `resolver.go` and `resolve_reference_test.go` change.
- Own-branch targeted acceptance is recorded separately from P3 integration acceptance.
- Full package PASS is required before T08 only after T05, T06, and T07 are integrated.

## Verification

Run from repository root:

```powershell
gofmt -d drmcp/src/internal/designrecords/resolver.go drmcp/src/internal/designrecords/resolve_reference_test.go
go test ./drmcp/src/internal/designrecords -run 'TestResolveReference' -count=1
go test ./drmcp/src/internal/designrecords -count=1
```

## Evidence

Record:

- exact classification matrix;
- resolved, unresolved, conflict, and unsupported examples;
- no-repair and no-path assertions;
- proof that legacy fallback is absent;
- targeted and full package outputs;
- exact changed-file boundary.
- own-branch targeted acceptance and later P3 integration acceptance as separate evidence.

### Provisional implementation mapping

```yaml
implementation_mapping:
  status: provisional
  contract_refs:
    - DRMCP-REQ-MCP-001
    - DRMCP-ADR-MCP-001
    - DRMCP-WORK-MCP-005
    - DRMCP-WORK-MCP-006
    - DRMCP-WORK-MCP-008
  fixture_cases:
    - C06
    - C11
    - C13
    - R02
    - R03
    - R04
    - R05
    - R14
  implementation:
    - path: drmcp/src/internal/designrecords/resolver.go
      symbols: []
  verification:
    - path: drmcp/src/internal/designrecords/resolve_reference_test.go
      tests: []
  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```

Populate `symbols` and `tests` with real names before Task closure.
Remove any path that does not contain the final contract-significant implementation or verification.
