# DRMCP-TASK-MCP-009-11: Review T05 correction execution graph

- **id**: DRMCP-TASK-MCP-009-11
- **status**: done
- **date**: 2026-06-29
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-10
- **outputs**:
  - DRMCP-TASK-MCP-009-11

## Goal

Independently judge whether the T05 correction graph is executable without hidden design work, overlapping writers, or unowned verification.

## Work

### Read

- `DRMCP-WORK-MCP-009-current-format-read-implementation.md`.
- `DRMCP-TASK-MCP-009-05-implement-current-list-and-exact-retrieval.md`.
- `DRMCP-TASK-MCP-009-08-add-current-fixture-integration-verification.md`.
- `DRMCP-TASK-MCP-009-09-review-and-close-current-read-implementation.md`.
- T10 through T17.
- Exact source and test files named by T13 and T14 only when required to confirm their frozen symbols or test names.

### Previous review results

The first, second, and third independent reviews returned `NEEDS REVISION`.
The next review is limited to F-BLOCK-01 plus preservation of the two closed findings.

| finding | current state | required re-review question |
|---|---|---|
| F-BLOCK-01 | CORRECTED_PENDING_REVIEW | Is T15 diagnostic classification total and ordered, with unconditional blockers evaluated before allowlist comparison and every command result assigned exactly one terminal result? |
| F-MAJ-01 | CLOSED | Is the accepted complete replacement mapping manifest and mechanical T17 replacement contract preserved without regression? |
| F-MIN-01 | CLOSED | Are its accepted exact commands, Git scope, absence checks, and writer boundaries preserved without regression? |

### Review

- Confirm the correction remains inside W009 and `DRMCP-REQ-MCP-001` ownership.
- Confirm T13 and T14 are self-contained executor leaves.
- Confirm every writable path has one writer.
- Confirm dependencies are acyclic.
- Confirm T13 and T14 may execute in parallel without consuming each other's output.
- Confirm model routing matches contract density and ambiguity.
- Confirm focused and aggregate verification ownership are separate.
- Confirm T15 is read-only, requires only correction-owned mandatory checks, and freezes the exact T08-owned test-name allowlist.
- Confirm T15 classification is total and ordered across exactly `PASS`, `DEFERRED_TO_T08`, and `BLOCKED`.
- Confirm compile, package setup, package initialization, panic, unknown, and mechanically unclassifiable failures are unconditional `BLOCKED` results evaluated before allowlist comparison.
- Confirm an allowlisted test panic is `BLOCKED`.
- Confirm an allowlisted but unclassifiable test failure is `BLOCKED`.
- Confirm `DEFERRED_TO_T08` is reachable only when every exact failing test is allowlisted and every unconditional blocker is false.
- Confirm T15 classifies diagnostic test failure by exact allowlist membership without semantic interpretation and blocks mixed and non-allowlisted failures.
- Confirm T15 contains a complete replacement mapping manifest with canonical record IDs, all nine fixture cases, verification items, removals, protected paths, exclusions, and future canonicalization.
- Confirm the mandatory adapter command executes `TestToolsCallSuccess`, `TestToolsListWorkflowKindEnums`, and `TestToolsCallToolErrors`.
- Confirm T16 owns independent B-01, M-01, and M-04 disposition plus complete replacement mapping-manifest acceptance.
- Confirm T17 owns mechanical field-for-field and value-for-value replacement, adds no merge or supplementation judgment, and does not close W009.
- Confirm T08 remains blocked through its existing T05 dependency until T17 closes T05.
- Confirm T09 remains the final W009 review and closure owner.
- Confirm protected T06 and T07 files and responsibilities remain closed.

### Prohibited operations

- Do not modify files.
- Do not correct findings.
- Do not release executor Tasks.
- Do not implement production changes.
- Do not update lifecycle state.
- Do not stage or commit.
- Do not perform repository-wide traversal.

### Output

1. Verdict: `PASS` or `NEEDS REVISION`.
2. Reviewed files.
3. Previous-finding disposition, when applicable.
4. Blocking findings.
5. Major findings.
6. Minor findings.
7. Advisories.
8. Writer and dependency assessment.
9. Model-routing assessment.
10. Verification and closure-owner assessment.
11. Release readiness.

## Done condition

- The graph receives an independent verdict.
- Every material finding has a stable ID and severity.
- The review covers writer uniqueness, dependencies, model routing, verification ownership, and closure ownership.
- No file is changed by the reviewer.

## Verification

- Inspect scoped Design Record diffs for W009, T05, T08, T09, and T10 through T17.
- Confirm each new Task uses the W009 work sequence and source Requirement.
- Confirm T13 and T14 writable paths do not overlap.
- Confirm T15 depends on both executor Tasks.
- Confirm T16 depends on T15 and T17 depends on T16.
- Confirm T08 still depends on T05 and T09 still depends on T08.
- Confirm no implementation prompt or production change is part of the graph-authoring diff.

## Evidence

### First review

- Review verdict: `NEEDS REVISION`.
- Findings: `F-BLOCK-01`, `F-MAJ-01`, and `F-MIN-01`.
- Release denied.
- First correction revision authored.

### Second review

```text
second review verdict: NEEDS REVISION

F-BLOCK-01:
  state: OPEN
  remaining issue:
    T08-owned test-name allowlist is not frozen.

F-MAJ-01:
  state: OPEN
  remaining issues:
    canonical contract_refs are incorrect.
    fixture_cases are missing.
    future_canonicalization is missing.

F-MIN-01:
  state: CLOSED

release denied: true
correction revision: second
re-review required: true
```

- Second correction authoring state: `CORRECTED_PENDING_REVIEW`.
- `F-BLOCK-01` and `F-MAJ-01` remained `OPEN` pending the third independent review.
- `F-MIN-01` remained `CLOSED`; its accepted exact commands, Git scope, absence checks, and writer boundaries were not redesigned.

### Third review

```text
third review verdict: NEEDS REVISION

F-BLOCK-01:
  state: OPEN
  remaining issue:
    diagnostic classification does not assign an allowlisted-test panic
    or equivalent unconditional failure to a unique terminal result.

F-MAJ-01:
  state: CLOSED

F-MIN-01:
  state: CLOSED

release denied: true
correction revision: third
re-review required: true
```

- Third correction authoring state: `CORRECTED_PENDING_REVIEW`.
- `F-BLOCK-01` is `CORRECTED_PENDING_REVIEW`; it must not be recorded as `CLOSED` before the next independent re-review.
- `F-MAJ-01` remains `CLOSED`; the accepted complete replacement mapping manifest was not redesigned.
- `F-MIN-01` remains `CLOSED`; its accepted exact commands, Git scope, absence checks, and writer boundaries were not redesigned.
- T12 remains blocked and `not_started`.
- T13 and T14 remain unreleased.

### Fourth review

```text
fourth review verdict: PASS

F-BLOCK-01:
  state: CLOSED

F-MAJ-01:
  state: CLOSED

F-MIN-01:
  state: CLOSED

new blocking findings: none
new major findings: none
new minor findings: none

release readiness:
  T12 may synchronize release: yes
  T13 ready for release: yes
  T14 ready for release: yes
  T15 mechanically executable: yes
  T16 F-BLOCK-01 gate accepted: yes
  T17 complete replacement synchronization accepted: yes
```

- The fourth limited independent re-review accepted the third correction revision.
- Writer ownership remains unique and T13 and T14 remain safe for parallel execution.
- Dependency order remains T13 and T14, then T15, T16, and T17.
- No production implementation, lifecycle synchronization, staging, or commit occurred during the review.

The reviewer did not use T10's conclusion as proof. The reviewer inspected the graph and exact boundaries directly.
