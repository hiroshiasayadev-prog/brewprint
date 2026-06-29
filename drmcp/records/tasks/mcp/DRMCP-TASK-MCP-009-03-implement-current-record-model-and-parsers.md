# DRMCP-TASK-MCP-009-03: Implement current record model and parsers

- **id**: DRMCP-TASK-MCP-009-03
- **status**: done
- **date**: 2026-06-29
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

**Accepted review results:**

- Six findings closed (F-MAJ-01, F-MAJ-02, F-MIN-01–F-MIN-04)
- Targeted parser/type tests: PASS
- Package compile/build: PASS
- Exact current identity: no case/prefix/whitespace/fuzzy repair
- Current sequential parsers: `parseCurrentADRRecord`, `parseCurrentInvestigationRecord`, `parseCurrentRequirementRecord`, `parseCurrentWorkItemRecord`, `parseCurrentTaskRecord`
- Path-derived spec parser: `parseCurrentSpecRecord` with `deriveSpecRef`
- Invalid source retention via `RecordCandidate`
- Nullable `CurrentListedRecord` missing title/status/date serialize as JSON null
- Structured `Diagnostic.Location` for portable location representation
- Shared API frozen for T05–T07: current list/get/resolve/validate request/response types
- Spec date retained in `Record.Date`
- `full_package_gate: deferred_to_T04`

### Corrective implementation after T04 review

T04 independent implementation review found that root-level `spec/index.md` returned no path-derived ref and could fall back to metadata `id`.
This contradicted the accepted root-index identity `spec:<app_namespace>` and the no-fallback identity rule.

Corrective scope:

- `deriveSpecRef` now derives `spec:<app_namespace>` for root `spec/index.md`;
- `parseCurrentSpecRecord` no longer uses metadata `id` as an identity fallback;
- `TestCurrentSpecRootIndexUsesPathDerivedIdentity` covers matching, mismatched, and missing metadata `id` while preserving the path-derived identity;
- `TestDeriveSpecRef` directly covers the root-index case.

Corrective verification completed on 2026-06-29:

- `gofmt -w` completed for `parser.go` and `parser_index_test.go`;
- root-index and adjacent current-spec parser tests passed with `-count=10`;
- the exact T04 active-index test set passed with `-count=10`;
- package compile passed with the no-test compile command;
- full-package execution failed only in the previously assigned T05, T06, T07, and T08 stale-test groups;
- scoped `git diff --check` passed with LF-to-CRLF conversion warnings only.

The root-index parser finding is corrected and verified. T03 returned to `done`; T04 remains responsible for scoped re-review acceptance.

### Implementation mapping

```yaml
implementation_mapping:
  status: accepted

  contract_refs:
    - DRMCP-REQ-MCP-001
    - DRMCP-ADR-MCP-001
    - DRMCP-WORK-MCP-003
    - DRMCP-WORK-MCP-004
    - DRMCP-WORK-MCP-005
    - DRMCP-WORK-MCP-006
    - DRMCP-WORK-MCP-008

  implementation:
    - path: drmcp/src/internal/designrecords/types.go
      symbols:
        - Record
        - OperationWarning
        - CurrentListRecordsRequest
        - CurrentListedRecord
        - CurrentListRecordsResponse
        - CurrentGetRecordsRequest
        - CurrentGetRecordsRecord
        - CurrentGetRecordsResponse
        - CurrentResolveReferenceRequest
        - CurrentResolvedTarget
        - CurrentResolveReferenceResponse
        - CurrentValidateRecordsRequest
        - ValidationSubjectSummary
        - CurrentValidateRecordsResponse
        - Diagnostic
        - DiagnosticLocation
        - CurrentConflict
    - path: drmcp/src/internal/designrecords/parser.go
      symbols:
        - parseCurrentADRRecord
        - parseCurrentInvestigationRecord
        - parseCurrentRequirementRecord
        - parseCurrentWorkItemRecord
        - parseCurrentTaskRecord
        - parseCurrentSpecRecord
        - deriveSpecRef

  verification:
    - path: drmcp/src/internal/designrecords/types_test.go
      tests:
        - TestCurrentGetRecordsResponseShape
        - TestCurrentValidateRecordsResponseShape
        - TestDiagnosticLocationShape
        - TestCurrentConflictShape
        - TestCurrentListedRecordJSONNullFields
        - TestCurrentListRecordsRequestJSONShape
        - TestCurrentListRecordsResponseJSONShape
    - path: drmcp/src/internal/designrecords/parser_index_test.go
      tests:
        - TestCurrentSpecRootIndexUsesPathDerivedIdentity
        - TestDeriveSpecRef

  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```
