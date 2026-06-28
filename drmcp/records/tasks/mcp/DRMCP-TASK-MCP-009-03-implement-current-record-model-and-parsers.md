# DRMCP-TASK-MCP-009-03: Implement current record model and parsers

- **id**: DRMCP-TASK-MCP-009-03
- **status**: not_started
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 2d
- **depends_on**:
  - DRMCP-TASK-MCP-009-01
- **outputs**:
  - drmcp/src/internal/designrecords/types.go
  - drmcp/src/internal/designrecords/types_test.go
  - drmcp/src/internal/designrecords/parser.go
  - drmcp/src/internal/designrecords/parser_index_test.go

## Goal

Establish the shared current read model and parse exact current sequential records and H1-adjacent current specs.

Freeze downstream request, response, diagnostic, source, and conflict structures before parallel tool work begins.

## Work

- Define app-aware canonical current identities without case, prefix, whitespace, or fuzzy normalization.
- Parse current sequential records from accepted filename, H1, and visible metadata authority.
- Parse current specs from real H1-adjacent visible metadata.
- Derive current spec identity from app namespace and spec-relative path.
- Reject YAML front matter as current spec metadata authority.
- Retain uniquely path-addressable invalid current sources for validation.
- Define path-free normal projections and portable diagnostic locations.
- Define all shared request and response types needed by T05 through T07.
- Remove retired read-operation and ID-range types when no authoring dependency remains.
- Preserve authoring compilation and behavior without editing authoring files.

### Execution slices

| slice | owner model | parallel group | dependency | exact file boundary or inventory method | allowed changes | prohibited changes | commands | expected evidence | escalation condition |
|---|---|---|---|---|---|---|---|---|---|
| S03A shared model and parser | Sonnet | P1 | T01 accepted | `types.go`, `types_test.go`, `parser.go`, `parser_index_test.go` only. | Redesign current read types and parser tests; preserve authoring-compatible fields where required. | `authoring.go`, authoring tests, config, index, tools, resolver, validation, fixtures, legacy parsing. | `gofmt -w`; targeted parser/type tests; package compile. | Frozen shared API summary, fixture-case trace, command output, `full_package_gate: deferred_to_T04`. | Any accepted contract not representable without changing authoring behavior; stop and record the conflict. |
| S03B mechanical verification | Haiku | P1 | S03A complete | Same four files; read-only. | No changes. | API redesign, parser correction, or authoring edits. | `gofmt -d`; targeted tests; package compile; scoped Git check. | Raw outputs, exact changed paths, and deferred full-package gate note. | Any failure or diff outside boundary; escalate to Sonnet. |

## Done condition

- Current sequential IDs require the explicit app namespace.
- Current identity is never repaired or normalized into another canonical value.
- Current specs use path-derived `spec:<app>.<segments>` identity.
- H1-adjacent visible metadata is authoritative.
- YAML-front-matter current specs are rejected as valid current sources.
- Invalid but uniquely path-addressable current sources remain available to validation.
- Shared types represent accepted list, get, resolver, validation, diagnostic, source, and conflict contracts.
- Normal response types contain no physical path field.
- T05 through T07 can proceed without editing `types.go`.
- Authoring package compile is preserved without changing authoring source or tests.
- W008 cases C01-C07, C11, C15, and R02-R06 are covered at the parser/model layer.
- Full-package PASS is deferred to T04.

## Verification

Run from repository root:

```powershell
gofmt -d `
  drmcp/src/internal/designrecords/types.go `
  drmcp/src/internal/designrecords/types_test.go `
  drmcp/src/internal/designrecords/parser.go `
  drmcp/src/internal/designrecords/parser_index_test.go
go test ./drmcp/src/internal/designrecords -run 'Test(.*Parse|.*Parser|.*Current.*ID|.*Spec.*Metadata|.*Diagnostic.*JSON)' -count=1
go test ./drmcp/src/internal/designrecords -run '^$' -count=1
```

Do not require full-package PASS for T03 individual acceptance.
Record `full_package_gate: deferred_to_T04`.

## Evidence

Record:

- the frozen shared type inventory for T05 through T07;
- exact identity grammar and no-repair tests;
- current spec metadata and path-derived identity tests;
- invalid-source retention tests;
- path-free projection proof;
- targeted and package-compile outputs;
- `full_package_gate: deferred_to_T04`;
- explicit confirmation that authoring files were not changed.

### Provisional implementation mapping

```yaml
implementation_mapping:
  status: provisional

  contract_refs:
    - DRMCP-REQ-MCP-001
    - DRMCP-ADR-MCP-001
    - DRMCP-WORK-MCP-003
    - DRMCP-WORK-MCP-004
    - DRMCP-WORK-MCP-005
    - DRMCP-WORK-MCP-006
    - DRMCP-WORK-MCP-008

  fixture_cases:
    - C01
    - C02
    - C03
    - C04
    - C05
    - C06
    - C07
    - C11
    - C15
    - R02
    - R03
    - R04
    - R05
    - R06

  implementation:
    - path: drmcp/src/internal/designrecords/types.go
      symbols: []
    - path: drmcp/src/internal/designrecords/parser.go
      symbols: []

  verification:
    - path: drmcp/src/internal/designrecords/types_test.go
      tests: []
    - path: drmcp/src/internal/designrecords/parser_index_test.go
      tests: []

  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```

Populate `symbols` and `tests` with real names before Task closure.
Remove any path that does not contain the final contract-significant implementation or verification.
