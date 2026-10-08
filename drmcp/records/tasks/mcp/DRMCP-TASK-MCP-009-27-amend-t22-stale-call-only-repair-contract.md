# DRMCP-TASK-MCP-009-27: Amend T22 stale-call-only repair contract

- **id**: DRMCP-TASK-MCP-009-27
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-26
- **outputs**:
  - DRMCP-WORK-MCP-009
  - DRMCP-TASK-MCP-009-05
  - DRMCP-TASK-MCP-009-15
  - DRMCP-TASK-MCP-009-16
  - DRMCP-TASK-MCP-009-22
  - DRMCP-TASK-MCP-009-23
  - DRMCP-TASK-MCP-009-27
  - DRMCP-TASK-MCP-009-28
  - DRMCP-TASK-MCP-009-29

## Goal

Correct the second T22 amendment after focused execution exposed a resolver-contract contradiction.

Limit T22 to removing retired T05 API calls and proving package compilation without executing T06-owned resolver assertions.

## Work

### Result classification

Classify the supplied T22 execution result as follows.

| observation | disposition |
|---|---|
| The current-root setup and fixture-path migration were applied exactly as released. | Keep the execution evidence as proof that T26 released contract was followed. |
| `gofmt` passed and scoped Git and whitespace checks passed. | Accept as mechanical evidence only. |
| The focused resolver test returned `unresolved` for `spec:non-record.doc`. | Mark T22 blocked. Do not treat the failure as an executor mistake. |
| Current spec parsing rejects YAML front matter and uses path-derived document refs. | Classify the T24-T26 contract as a released-contract gap. |
| T06 removes semantic aliases and section aliases and owns resolver behavior and tests. | Return resolver assertion execution to T06 ownership. |

The second T22 execution is not accepted as completed focused Evidence.

T25 review missed the controlling T06, W003, and W005 contracts. T25 and T26 remain immutable historical records. This amendment supersedes only their T22 execution contract.

### Controlling contract

The accepted current resolver contract requires:

- exact app-aware sequential IDs;
- exact path-derived current `spec:` document refs;
- no front-matter `semantic_refs` alias registry;
- no front-matter `sections` alias registry;
- no section-level resolver target;
- no physical path in normal successful resolver responses.

The named stale test still describes pre-current semantic alias behavior. T05 does not own rewriting or deleting those resolver assertions.

### T22 corrected implementation contract

Keep the T22 writer path unchanged:

```text
drmcp/src/internal/designrecords/resolve_reference_test.go
```

Inside `TestResolveReferenceUsesSemanticRefsFromNonRecordSpec`, keep only the first-execution deletions:

- remove the `ListRecords` call and zero-record assertion;
- remove the `GetRecord` call and expected-failure assertion.

Rollback the second-amendment migration:

1. Change the fixture path from `drmcp/records/spec/non-record.md` back to `records/spec/non-record.md`.
2. Replace the test-local `NewConfig` and `BuildIndex` block with `idx := buildTestIndex(t, root)`.
3. Change the expected document path from `drmcp/records/spec/non-record.md` back to `records/spec/non-record.md`.

Preserve the remaining resolver assertions byte-for-byte apart from `gofmt` effects.

Do not execute the named resolver test in T22.
Do not delete or rewrite the resolver assertions.
T06 retains final ownership of resolver test migration and semantic behavior.

### T22 verification contract

Run from repository root:

```powershell
gofmt -w drmcp/src/internal/designrecords/resolve_reference_test.go

go test ./drmcp/src/internal/designrecords -run '^$' -count=1

git grep -n -E '\b(ListRecords|GetRecord)\b' -- drmcp/src/internal/designrecords/resolve_reference_test.go
```

Verification interpretation:

| command | accepted result |
|---|---|
| `gofmt` | exit `0` |
| compile-only `go test -run '^$'` | exit `0` |
| scoped stale-call grep | exit `1` |

The compile-only command proves that retired symbols are no longer referenced by package tests.
It does not claim resolver semantic behavior or focused resolver-test acceptance.

T15 owns the T05-focused aggregate commands.
T06 and its integration lane own resolver behavioral verification.

### Execution graph amendment

```text
T20 accepted repair release
  -> T21 clean retry remains independently eligible
  -> T22 first stale-call deletion complete
  -> T24-T26 second amendment released current-root migration
  -> T22 second execution followed contract but exposed semantic contradiction
  -> T27 stale-call-only contract amendment
  -> T28 independent amendment review
  -> T29 T22 re-release
  -> T22 rollback migration and run compile-only verification
  -> T23 valid T21 and T22 Evidence synchronization
  -> T15 aggregate verification
  -> T16 independent finding closure
  -> T17 correction closure synchronization
```

T21 retry may run while T28 reviews this amendment.
T22 may execute only after T28 `PASS` and T29 release synchronization.

### Downstream synchronization

Update T23 to require T22 compile-only Evidence and stale-call absence only.

Update T15 and T16 to treat the complete T22 diff as only the two stale assertion-block deletions.

Do not add the resolver test to T15 mandatory commands or the final T05 implementation mapping.

## Done condition

- The focused resolver failure is classified as a released-contract gap.
- T25 review scope omission is recorded without rewriting T25 or T26 history.
- T22 is limited to stale retired-call removal and package-test compilation.
- The erroneous current-root migration is a mandatory rollback.
- Resolver assertion execution and migration remain T06-owned.
- T23, T15, and T16 use the corrected T22 boundary.
- T28 and T29 exist with exact review and release ownership.
- No production source or Go test changes occur during authoring.

## Verification

- Confirm W003 rejects YAML front matter for current specs.
- Confirm W005 and T06 reject semantic and section alias resolution.
- Confirm the second T22 diff followed T24-T26 exactly.
- Confirm the focused failure occurs after index construction and returns `unresolved`.
- Confirm the corrected T22 contract changes only one test file.
- Confirm T28 depends on T27 and T29 depends on T28.
- Confirm T22 depends on T29 before execution.
- Confirm T23 still depends on T21 and T22.
- Inspect only amended Design Records with scoped Git and whitespace checks.

## Evidence

- Second T22 execution result: `BLOCKED`.
- `gofmt` exit: `0`.
- Focused resolver test exit: `1`.
- Failure: `spec:non-record.doc` resolved as `unresolved`.
- Current parser evidence: YAML front matter current specs are rejected.
- Resolver authority evidence: front-matter semantic aliases and section aliases are prohibited.
- Classification: released-contract gap, not executor mistake.
- Corrected minimum: rollback second-amendment migration, preserve stale-call deletions, and use compile-only verification.
- Filesystem authoring was used because DRMCP authoring transactions were unavailable.
- No production source, Go test, fixture, schema, manifest, ADR, Requirement, or Specification changed during T27.

Release history:

```text
T28 independent review: PASS
blocking findings: none
major findings: none
minor findings: none
T29 release synchronization: done
T22 stale-call-only correction: re-released
```
