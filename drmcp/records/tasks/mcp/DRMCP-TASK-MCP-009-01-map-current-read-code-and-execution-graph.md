# DRMCP-TASK-MCP-009-01: Map current read code and execution graph

- **id**: DRMCP-TASK-MCP-009-01
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-001-08
- **outputs**:
  - DRMCP-WORK-MCP-009
  - DRMCP-TASK-MCP-001-09
  - DRMCP-TASK-MCP-009-02
  - DRMCP-TASK-MCP-009-03
  - DRMCP-TASK-MCP-009-04
  - DRMCP-TASK-MCP-009-05
  - DRMCP-TASK-MCP-009-06
  - DRMCP-TASK-MCP-009-07
  - DRMCP-TASK-MCP-009-08
  - DRMCP-TASK-MCP-009-09

## Goal

Map the bounded current implementation and test inventory to accepted W003 through W008 contracts.

Create a Claude Code execution graph with non-overlapping file boundaries, explicit model ownership, dependency gates, commands, evidence, and escalation conditions.

## Work

- Read `prompt_chappy.md` and the Task and Work Item authoring standards first.
- Read W003 through W009, hub T09, and the accepted W008 fixture manifest.
- Inspect only `drmcp/src/internal/designrecords` production files and existing package tests.
- Treat accepted records and fixtures as authority.
- Treat current implementation and existing tests only as migration inventory.
- Identify shared-file hotspots before defining parallel groups.
- Create T02 through T09 without starting production implementation.
- Synchronize W009 and hub T09 to `in_progress` because child planning work has begun.

This Task does not modify Go source, Go tests, fixtures, accepted ADRs, Requirements, or Specifications.

### Contract-to-code correspondence

| accepted concern | current files | observed starting state | implementation owner |
|---|---|---|---|
| Explicit configured current roots and app association | `config.go`, `config_test.go` | Auto-discovers `*/records`, falls back to `v01/records`, and derives uppercase prefixes. This contradicts W003. | T02 |
| Current public data model, normalized metadata, diagnostics, request and response shapes | `types.go`, `types_test.go` | Uses non-app IDs, scalar diagnostic shortcuts, `ids/items`, path-bearing projections, range selectors, and retired-operation types. Shared with authoring code. | T03 |
| Current sequential and spec parsing | `parser.go`, `parser_index_test.go` | Normalizes identity, accepts YAML-front-matter specs, uses `SPEC-*`, and registers semantic aliases. This contradicts W003. | T03 |
| Deterministic discovery, addressability, and duplicate conflict groups | `index.go`, `index_test.go` | Merges roots into one record slice, retains semantic aliases, and does not expose accepted conflict-state indexing. | T04 |
| Public read-operation entrypoint ownership | `tools.go`, `resolver.go`, `validation.go` | Public `ResolveReference` and `ValidateRecords` are exposed from `tools.go`, which would collide with T05 ownership. | T04 mechanical split |
| Compact listing and exact current retrieval | `tools.go`, `list_records_test.go`, `get_records_test.go` | Listing is broad and path-bearing. Retrieval uses `ids`, item wrappers, and failure placeholders. | T05 |
| Retired public operations and obsolete ranges | `tools.go`, `get_record_test.go`, `suggest_next_record_test.go`, `id_range.go` | `get_record`, `suggest_next_record`, and ID-range behavior remain implemented and tested. | T05 |
| Current canonical resolution | `resolver.go`, `resolve_reference_test.go` | Resolves semantic aliases and paths, normalizes IDs, and exposes physical paths. | T06 |
| Current repository validation and shared diagnostics | `validation.go`, `validation_test.go` | Uses kind/range selectors, old category names, scalar paths, semantic aliases, and obsolete spec checks. | T07 |
| Accepted current fixture integration | `testdata/read-baseline/manifest.json` and current arrangements | Accepted fixture bytes exist, but production runtime coverage is not yet connected. | T08 |
| Authoring non-regression boundary | `authoring.go`, `authoring_test.go`, authoring guidance files | These files consume shared record, index, parser, and diagnostic types. They are not W009 change targets. | T03 and final package verification |

### Accepted fixture allocation

| owner | cases |
|---|---|
| T02 | C08, C10, R08, R10 configuration inputs. |
| T03 | C01-C07, C11, C15, R02-R06 parsing and exact-identity inputs. |
| T04 | C08, C12, R07 active-index and duplicate-conflict inputs. |
| T05 | C11, C12, C16, R02-R05, R14, R20 list, retrieval, rejection, unresolved, and path-hiding inputs. |
| T06 | C06, C11, C13, R02-R05, R14 current resolver inputs. |
| T07 | C09, C14, C15, R07, R08, R10 current validation inputs. |
| T08 | End-to-end coverage for C01-C16 and W009-owned R02-R08, R10, R14, R20. |

C17 and R21 through R23 remain W-SPEC-owned.
All L cases and legacy-owned R cases remain W010-owned.

### Claude Code execution slices

| slice | owner model | parallel group | dependency | exact file boundary or inventory gate | allowed changes | prohibited changes | commands | expected evidence | escalation condition |
|---|---|---|---|---|---|---|---|---|---|
| S01A inventory and graph | Sonnet | P0 | W003-W008 accepted | This Task, W009, and hub T09 only. | Design Records only. | Any Go or fixture change. | Scoped record-shape and Git checks. | Correspondence table, fixture allocation, graph, and changed-file manifest. | Any accepted-contract contradiction or pre-existing T009 Task ID. |
| S01B graph verification | Haiku | P0 | S01A complete | Same eleven Design Record paths listed in Evidence. | No file changes. | Contract interpretation or graph redesign. | Recorded lifecycle and whitespace checks. | Exact pass/fail output only. | Any mismatch; return to Sonnet without repairing. |

### Parallel execution graph

```text
T01
├── P1: T02 config foundation
└── P1: T03 shared model and parser foundation
        T02 targeted acceptance
        T03 targeted acceptance
                 │
                 ▼
T04 serial foundation integration
- merge T02/T03
- active index
- public operation ownership split
- shared API freeze
- compile/full-package integration gate
                 │
                 ▼
├── P3: T05 list/get/catalog
├── P3: T06 resolver semantics
└── P3: T07 validation semantics
        T05 + T06 + T07
                 │
                 ▼
T08 fixture integration
                 │
                 ▼
T09 review and closure
```

T02 and T03 have disjoint source boundaries.
T02 and T03 defer full-package PASS to T04.
T04 integrates T02 and T03, freezes shared APIs, and separates public operation ownership before P3 starts.
T04 must not retain stale behavior only to force full-package PASS.
T05, T06, and T07 have disjoint source and test boundaries after T04 freezes `tools.go`, `resolver.go`, and `validation.go` ownership.
T08 cannot begin until T05, T06, and T07 are merged and the P3 integration full-package gate passes.
Parallel slices must use separate branches or worktrees and merge only after their dependency gate is accepted.

## Done condition

- The correspondence table covers every W009-owned production file and existing package test.
- Accepted fixture cases are allocated without taking W010 or W-SPEC ownership.
- T02 through T09 exist with canonical Task sections.
- Every execution slice records model, parallel group, dependency, boundary, allowed and prohibited changes, commands, evidence, and escalation.
- No judgment-bearing implementation slice is assigned to Haiku.
- Haiku slices are mechanical verification or exact-plan execution only.
- Shared-file ownership prevents parallel write conflicts.
- W009 and hub T09 are `in_progress`.
- Production source, tests, and fixtures remain unchanged.
- T02 through T08 share one provisional implementation-mapping Evidence contract.
- The mapping contract remains execution evidence and does not become current-state implementation-design authority.
- Independent review reports no blocking or major finding before this Task changes to `done`.

## Verification

- Compare the graph with W003 through W009 and the W008 manifest.
- Confirm all Task IDs are unused before creation.
- Confirm T02 and T03 file boundaries do not overlap.
- Confirm T02 and T03 defer full-package PASS to T04.
- Confirm T04 owns serial foundation integration, public operation ownership split, shared API freeze, and package compile.
- Confirm T05, T06, and T07 post-T04 file boundaries do not overlap.
- Confirm T08 records the P3 integration full-package PASS gate as a start condition.
- Confirm T05 and T09 use expanded changed-boundary format verification.
- Confirm every downstream Task depends on the owner that freezes its shared input.
- Confirm no legacy implementation or retained spec-validator implementation is assigned to W009.
- Confirm T02 through T08 contain the provisional `implementation_mapping` shape with accepted contract refs, owned fixture cases, implementation paths, verification paths, and pending future canonical refs.
- Confirm no Internal Design or BPDSL canonical ID, schema, artifact, or sidecar manifest is invented.
- Confirm each mapping must use actual contract-significant symbols and Go test functions before its Task closes.
- Run the scoped lifecycle and whitespace checks recorded below.
- Run an independent review before closing T01.

## Evidence

### Bounded inventory

The exact task directory was listed before creation.
No `DRMCP-TASK-MCP-009-*` file existed.
The selected sequence is T01 through T09.

The implementation inventory was limited to:

- `drmcp/src/internal/designrecords/*.go`;
- existing `*_test.go` files in that package;
- `drmcp/src/internal/designrecords/testdata/read-baseline/manifest.json`.

No repository-wide source traversal was used.

### Shared-file gates

- T03 exclusively owns `types.go` during the model migration.
- T02 and T03 individual acceptance requires targeted tests, package compile, scoped format, scoped Git evidence, and `full_package_gate: deferred_to_T04`.
- T04 cannot begin until T02 and T03 reach targeted acceptance.
- T04 owns the merged T02/T03 serial gate, active index, public operation ownership split, shared API freeze, and full-package integration run.
- T04 must make the package compile before closure.
- T04 may record only named T05-T07 stale-operation test failures as expected transitional failures.
- T04 must return unresolved ownership or API mismatch to T03.
- T05, T06, and T07 must not edit `types.go`, `config.go`, `parser.go`, or `index.go`.
- T05 owns `tools.go` list/get behavior after T04 moves resolver and validation public entrypoints out of that file.
- T06 owns `resolver.go` resolver semantics after T04 moves or isolates `ResolveReference`.
- T07 owns `validation.go` validation semantics after T04 moves or isolates `ValidateRecords`.
- A downstream API mismatch is an escalation to the owning Task, not permission to cross the boundary.
- `authoring.go` and authoring tests are protected non-regression inputs, not implementation targets.

### Changed-file manifest

| action | file |
|---|---|
| new | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-01-map-current-read-code-and-execution-graph.md` |
| new | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-02-implement-configured-current-root-loading.md` |
| new | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-03-implement-current-record-model-and-parsers.md` |
| new | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-04-build-deterministic-current-active-index.md` |
| new | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-05-implement-current-list-and-exact-retrieval.md` |
| new | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-06-implement-current-reference-resolution.md` |
| new | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-07-implement-current-validation-and-diagnostics.md` |
| new | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-08-add-current-fixture-integration-verification.md` |
| new | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-009-09-review-and-close-current-read-implementation.md` |
| modify | `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-009-current-format-read-implementation.md` |
| modify | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-09-track-current-read-implementation.md` |

No Go source, Go test, fixture, ADR, Requirement, or Specification is in this boundary.

### Provisional implementation mapping contract

The T02 through T08 `implementation_mapping` blocks are a common W009 Evidence contract.

- Task Evidence records implementation-time traceability only.
- Task Evidence is not the current-state implementation-design source of truth.
- `contract_refs` contains accepted Requirement, ADR, Work Item, or Specification IDs only.
- `fixture_cases` contains only fixture cases owned by that Task.
- `implementation` records production paths and contract-significant Go symbols, not exhaustive private helpers.
- `verification` records test paths and real Go test function names.
- Empty `symbols` and `tests` arrays are planning placeholders only and must be populated before Task closure.
- `internal_design_ref` and `bpdsl_ref` remain `pending` until formal canonical IDs exist.
- No sidecar manifest, Internal Design artifact, BPDSL artifact, schema, or canonical ID is created by W009.
- Future canonicalization may consume this Evidence, but must move authority into the future canonical artifact.

### Ready-to-run lifecycle check

Run from the repository root:

```powershell
@'
from pathlib import Path

root = Path("drmcp/records/tasks/mcp")
tasks = [root / f"DRMCP-TASK-MCP-009-{i:02d}-" for i in range(1, 10)]
found = sorted(root.glob("DRMCP-TASK-MCP-009-*.md"))
assert len(found) == 9, found
for i, path in enumerate(found, 1):
    text = path.read_text(encoding="utf-8")
    assert f"- **id**: DRMCP-TASK-MCP-009-{i:02d}" in text
    for heading in ["## Goal", "## Work", "## Done condition", "## Verification", "## Evidence"]:
        assert text.splitlines().count(heading) == 1, (path, heading)
work = Path("drmcp/records/work-items/mcp/DRMCP-WORK-MCP-009-current-format-read-implementation.md").read_text(encoding="utf-8")
hub = Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-09-track-current-read-implementation.md").read_text(encoding="utf-8")
assert "- **status**: in_progress" in work
assert "- **status**: in_progress" in hub
for i in range(1, 10):
    assert f"  - DRMCP-TASK-MCP-009-{i:02d}" in work
print("task_graph_shape=OK")
'@ | python -
```

Expected output:

```text
task_graph_shape=OK
```

### Command execution state

Production implementation commands are `NOT RUN`.
Go tests are `NOT RUN` because production implementation has not started.
Review-finding correction is limited to the W009, hub T09, and T01-T09 Design Records.

### Independent review acceptance

- Review date: 2026-06-28.
- Verdict: PASS.
- Previous-finding disposition:
  - F-BLOCK-01: CLOSED.
  - F-MAJ-01: CLOSED.
  - F-MIN-01: CLOSED.
  - F-MIN-02: CLOSED.
- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- Closure readiness: accepted.
- T02 and T03 parallel-start readiness: ready.

The metadata date remains 2026-06-28 because this synchronization adds review Evidence without changing Task scope or done conditions.
