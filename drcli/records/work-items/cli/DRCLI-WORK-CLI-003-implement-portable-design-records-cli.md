# DRCLI-WORK-CLI-003: Implement portable Design Records CLI

- **status**: in_progress
- **date**: 2026-10-08
- **source_refs**:
  - DRCLI-REQ-CLI-001
  - spec:drcli.design_records_cli
  - DRCLI-WORK-CLI-002
- **impact_refs**: []
- **tasks**:
  - DRCLI-TASK-CLI-003-01
  - DRCLI-TASK-CLI-003-02
  - DRCLI-TASK-CLI-003-03
  - DRCLI-TASK-CLI-003-04
  - DRCLI-TASK-CLI-003-05
  - DRCLI-TASK-CLI-003-06
  - DRCLI-TASK-CLI-003-07
  - DRCLI-TASK-CLI-003-08
  - DRCLI-TASK-CLI-003-09
  - DRCLI-TASK-CLI-003-10
  - DRCLI-TASK-CLI-003-11
  - DRCLI-TASK-CLI-003-12
  - DRCLI-TASK-CLI-003-13
  - DRCLI-TASK-CLI-003-14
  - DRCLI-TASK-CLI-003-15
  - DRCLI-TASK-CLI-003-16
  - DRCLI-TASK-CLI-003-17
  - DRCLI-TASK-CLI-003-18
  - DRCLI-TASK-CLI-003-19
  - DRCLI-TASK-CLI-003-20

## Goal

Deliver a verified, copyable Windows/Linux DRCLI implementing every observable contract under `spec:drcli.design_records_cli`.

## Boundary

Own new DRCLI production source, local tests, fixtures, build and packaging scripts, bundled portable standards, portability verification, independent implementation review, and lifecycle closure.

Do not change accepted DRCLI Specification behavior, completed DRCLI-WORK-CLI-002 or its historical Tasks/review verdicts, DRMCP, or PRODUCT Design Records semantics.
Do not stage or commit files. Do not require a daemon, MCP server, per-repository install, or proposal-first authoring.

## Impact Scope

| target | impact |
|---|---|
| `drcli/` excluding `drcli/records/` | New production implementation, tests, fixtures, and distribution assets after graph release. |
| `drcli/records/work-items/cli/DRCLI-WORK-CLI-003-*.md` | Canonical Work Item graph, acceptance gates and closure Evidence. |
| `drcli/records/tasks/cli/DRCLI-TASK-CLI-003-*.md` | Exact executor contracts and independent gate Evidence. |
| `spec:drcli.design_records_cli` | Read-only canonical observable behavior. |

## Task flow

### Exact release and execution order

Graph IDs and `depends_on` are canonical in each Task; the table is the Work Item routing view. Every implementation Task is held until T03 is DONE after an independent T02 PASS; no marker or queue entry overrides repository Task status.

```text
T01 -> T02 [independent PASS] -> T03 [release]
T03 -> T04 -> {T05,T06,T13}
{T05,T06} -> {T07,T08,T09,T10,T11,T12}
{T04,T05,T06,T11,T13} -> T14
{T04,T05,T06,T07,T08,T09,T10,T11,T12,T14} -> T15
{T12,T13,T15} -> T16 -> T17 -> T18 [Windows+Linux PASS]
T18 -> T19 [independent PASS] -> T20 [synchronize]
```

### Complete dependency and owner map

| Task | Requires DONE / accepted state | Next consumers | Single responsibility |
|---|---|---|---|
| T01 | - | T02 | Graph materialization and Work Item graph |
| T02 | T01 | T03 | Independent graph verdict; no edits |
| T03 | T02 PASS | T04 | Review release and Work Item release Evidence |
| T04 | T03 | T05, T06, T13 | Go module, private model and artifact rules |
| T05 | T04 | T07,T08,T09,T10,T11,T12,T14,T15 | Current state, discovery and Git ignore |
| T06 | T04 | T07,T08,T09,T10,T11,T12,T14,T15 | Closed diagnostic envelope |
| T07 | T05,T06 | T15 | Scope/list/tree semantics |
| T08 | T05,T06 | T15 | Batch read and H2 sections |
| T09 | T05,T06 | T15 | RE2/literal search and snippets |
| T10 | T05,T06 | T15 | Current/V01 reference resolution |
| T11 | T04,T05,T06 | T14,T15 | Canonical finding projection and validation |
| T12 | T04,T05,T06 | T15,T16 | Portable authoring guidance projection |
| T13 | T04 | T14,T16 | Opaque state and proposal persistence |
| T14 | T04,T05,T06,T11,T13 | T15 | Direct/deferred authoring and write guards |
| T15 | T04,T05,T06,T07,T08,T09,T10,T11,T12,T14 | T16 | CLI parser, dispatch, help, streams |
| T16 | T12,T13,T15 | T17 | Copyable Windows/Linux packaging |
| T17 | T16 | T18 | End-to-end test suite and fixtures |
| T18 | T17 | T19 | Read-only full OS acceptance gate |
| T19 | T18 | T20 if PASS | Independent code review and named findings |
| T20 | T19 PASS and CLOSED findings | parent DONE | Lifecycle-only synchronization |

### Serialized writers and exact paths

| Task | Exclusive mutable output partition | Focused proof |
|---|---|---|
| T04 | `drcli/go.mod`, `drcli/internal/model/**` | `go -C drcli test ./internal/model -count=1` |
| T05 | `drcli/internal/current/**` | `go -C drcli test ./internal/current -count=1` |
| T06 | `drcli/internal/diagnostic/**` | `go -C drcli test ./internal/diagnostic -count=1` |
| T07 | `drcli/internal/discovery/**` | `go -C drcli test ./internal/discovery -count=1` |
| T08 | `drcli/internal/retrieval/**` | `go -C drcli test ./internal/retrieval -count=1` |
| T09 | `drcli/internal/search/**` | `go -C drcli test ./internal/search -count=1` |
| T10 | `drcli/internal/reference/**` | `go -C drcli test ./internal/reference -count=1` |
| T11 | `drcli/internal/validation/**` | `go -C drcli test ./internal/validation -count=1` |
| T12 | `drcli/internal/guidance/**` | `go -C drcli test ./internal/guidance -count=1` |
| T13 | `drcli/internal/state/**` | `go -C drcli test ./internal/state -count=1` |
| T14 | `drcli/internal/authoring/**` | `go -C drcli test ./internal/authoring -count=1` |
| T15 | `drcli/internal/cli/**`, `drcli/cmd/drcli/**` | `go -C drcli test ./internal/cli -count=1` |
| T16 | `drcli/scripts/**`, `drcli/tests/package/**`, `drcli/README.md`, `drcli/dist/**` | `go -C drcli test ./tests/package -count=1` |
| T17 | `drcli/tests/integration/**`, `drcli/tests/fixtures/**` | `go -C drcli test ./tests/integration -run '^TestEndToEnd' -count=1` |
| T18 | Task T18 Evidence only (read-only test executor) | `go -C drcli test ./... -count=1` plus native OS evidence |
| T19 | Task T19 Evidence only (independent, no fixes) | Independent 82-contract review |
| T20 | T20 Evidence and parent Work Item completion only | Compare verified T18 PASS, T19 PASS and finding closure |

Shared Work Item writers are serialized T01 -> T03 -> T20. T02/T19 own independent verdicts only; T18 owns objective acceptance only. No concurrent writer owns an overlapping path. Neither code review nor verification writes production. Implementers use Task-local read/change boundaries; source specs, historical review, PRODUCT, DRMCP remain read-only.

### Implementation-local private API freeze

- Go 1.22 standalone nested module `github.com/hiroshiasayadev-prog/brewprint/drcli` at `drcli/go.mod`; no root `go.mod`, `scripts/verify.bat`, `bin/`, DRMCP or PRODUCT modifications. Stdlib implementation unless later explicitly reviewed dependency scope change; root dependencies are not silently inherited.
- `model` exports artifact rules, identity, Markdown source slices and records. `current` imports `model` and owns repository snapshot, index, nodes and ignore. `diagnostic` owns catalog and placements independently. Operation packages import `model/current/diagnostic`, never another sibling operation; `validation` owns canonical finding projection.
- `state` owns opaque process-independent caches/proposals only; `authoring` consumes `model/current/validation/diagnostic/state` and owns writes; `cli` wires concrete commands and streams. Packaging owns copied standards sourced from the existing read-only `bin/design-records` tree and places them beside executables.
- Private exported symbols, exact `Change`/`Test` files, negative cases and Stop/Output protocol are frozen in T04-T17. A newly discovered missing contract or out-of-scope symbol stops the executor and requires new coordination, not speculative implementation.
- Initial release T03 may enqueue T04 only. After T04, T05/T06/T13 may run in parallel; after their dependencies, distinct operation-package leaves may run in parallel. Shared parent writers remain serialized.
- Model routing: Sonnet/Codex for all semantic, atomic-write, security, cross-process and integration outcomes; Haiku only for already-frozen mechanical fixture transcription if explicitly reassigned without scope change.
- If T02 or T19 returns NEEDS REVISION, a new finding-specific coordination Task first fixes routing; create correction and independent closure-review Tasks only for formally named findings. Do not reserve them now; keep T03/T20 unreleased until valid PASS.

### Accepted Specification trace (82 current files)

Each row names one **primary** implementation owner and the focused test target. Consumer dependencies may read other specs without acquiring second primary ownership. All 82 current `.md` sources (including topic indexes) are accounted for; no `drcli/records/spec/` files are to be changed.

| accepted spec ref | primary Task | focused check |
|---|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.h1_adjacent_metadata` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.base.definitions.identity_declaration` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.base.definitions` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.base.definitions.record_structure` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.base.definitions.source` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.base` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.base.templates.h1_adjacent_metadata` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.base.templates.identity_and_structure` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.base.templates.identity_and_structure.sequential` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.base.templates.identity_and_structure.tree` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.base.templates` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.base.templates.source` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.decision.h1_adjacent_metadata` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.decision.identity_and_structure` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.decision` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.decision.source` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.investigation.h1_adjacent_metadata` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.investigation.identity_and_structure` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.investigation` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.investigation.source` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.requirement.h1_adjacent_metadata` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.requirement.identity_and_structure` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.requirement` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.requirement.source` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.spec.h1_adjacent_metadata` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.spec.identity_and_structure` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.spec` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.spec.source` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.task.h1_adjacent_metadata` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.task.identity_and_structure` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.task` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.task.source` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.work_item.h1_adjacent_metadata` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.work_item.identity_and_structure` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.work_item` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.artifacts.work_item.source` | T04 | `go -C drcli test ./internal/model -count=1` |
| `spec:drcli.design_records_cli.cli.authoring_body_input` | T15 | `go -C drcli test ./internal/cli -count=1` |
| `spec:drcli.design_records_cli.cli.command_surface` | T15 | `go -C drcli test ./internal/cli -count=1` |
| `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics` | T15 | `go -C drcli test ./internal/cli -count=1` |
| `spec:drcli.design_records_cli.cli` | T15 | `go -C drcli test ./internal/cli -count=1` |
| `spec:drcli.design_records_cli.cli.output_and_exit_status` | T15 | `go -C drcli test ./internal/cli -count=1` |
| `spec:drcli.design_records_cli.cli.repository_selection_and_traversal` | T05 | `go -C drcli test ./internal/current -count=1` |
| `spec:drcli.design_records_cli.cli.runtime_and_distribution` | T16 | `go -C drcli test ./tests/package -count=1` |
| `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission` | T05 | `go -C drcli test ./internal/current -count=1` |
| `spec:drcli.design_records_cli.current_record_model.current_record_scopes` | T05 | `go -C drcli test ./internal/current -count=1` |
| `spec:drcli.design_records_cli.current_record_model.current_source_corpus` | T05 | `go -C drcli test ./internal/current -count=1` |
| `spec:drcli.design_records_cli.current_record_model.identity_conflict_and_addressability` | T05 | `go -C drcli test ./internal/current -count=1` |
| `spec:drcli.design_records_cli.current_record_model` | T05 | `go -C drcli test ./internal/current -count=1` |
| `spec:drcli.design_records_cli.diagnostics` | T06 | `go -C drcli test ./internal/diagnostic -count=1` |
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog` | T06 | `go -C drcli test ./internal/diagnostic -count=1` |
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_envelope` | T06 | `go -C drcli test ./internal/diagnostic -count=1` |
| `spec:drcli.design_records_cli` | T17 | `go -C drcli test ./tests/integration -count=1` |
| `spec:drcli.design_records_cli.operations.authoring.create` | T14 | `go -C drcli test ./internal/authoring -count=1` |
| `spec:drcli.design_records_cli.operations.authoring.deferred_proposal` | T14 | `go -C drcli test ./internal/authoring -count=1` |
| `spec:drcli.design_records_cli.operations.authoring` | T14 | `go -C drcli test ./internal/authoring -count=1` |
| `spec:drcli.design_records_cli.operations.authoring.retry_cache` | T14 | `go -C drcli test ./internal/authoring -count=1` |
| `spec:drcli.design_records_cli.operations.authoring.update` | T14 | `go -C drcli test ./internal/authoring -count=1` |
| `spec:drcli.design_records_cli.operations.authoring.write` | T14 | `go -C drcli test ./internal/authoring -count=1` |
| `spec:drcli.design_records_cli.operations.discovery_and_listing` | T07 | `go -C drcli test ./internal/discovery -count=1` |
| `spec:drcli.design_records_cli.operations.discovery_and_listing.record_scope_discovery` | T07 | `go -C drcli test ./internal/discovery -count=1` |
| `spec:drcli.design_records_cli.operations.discovery_and_listing.sequential_record_listing` | T07 | `go -C drcli test ./internal/discovery -count=1` |
| `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_child_listing` | T07 | `go -C drcli test ./internal/discovery -count=1` |
| `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_overview` | T07 | `go -C drcli test ./internal/discovery -count=1` |
| `spec:drcli.design_records_cli.operations.guidance.exact_read` | T12 | `go -C drcli test ./internal/guidance -count=1` |
| `spec:drcli.design_records_cli.operations.guidance` | T12 | `go -C drcli test ./internal/guidance -count=1` |
| `spec:drcli.design_records_cli.operations.guidance.list` | T12 | `go -C drcli test ./internal/guidance -count=1` |
| `spec:drcli.design_records_cli.operations` | T17 | `go -C drcli test ./tests/integration -count=1` |
| `spec:drcli.design_records_cli.operations.reference_resolution` | T10 | `go -C drcli test ./internal/reference -count=1` |
| `spec:drcli.design_records_cli.operations.retrieval.batch_retrieval_results` | T08 | `go -C drcli test ./internal/retrieval -count=1` |
| `spec:drcli.design_records_cli.operations.retrieval.exact_record_retrieval` | T08 | `go -C drcli test ./internal/retrieval -count=1` |
| `spec:drcli.design_records_cli.operations.retrieval.h2_section_retrieval` | T08 | `go -C drcli test ./internal/retrieval -count=1` |
| `spec:drcli.design_records_cli.operations.retrieval` | T08 | `go -C drcli test ./internal/retrieval -count=1` |
| `spec:drcli.design_records_cli.operations.retrieval.retrieval_output_limits` | T08 | `go -C drcli test ./internal/retrieval -count=1` |
| `spec:drcli.design_records_cli.operations.search` | T09 | `go -C drcli test ./internal/search -count=1` |
| `spec:drcli.design_records_cli.operations.search.lexical_match_model` | T09 | `go -C drcli test ./internal/search -count=1` |
| `spec:drcli.design_records_cli.operations.search.lexical_record_search` | T09 | `go -C drcli test ./internal/search -count=1` |
| `spec:drcli.design_records_cli.operations.search.search_results_and_limits` | T09 | `go -C drcli test ./internal/search -count=1` |
| `spec:drcli.design_records_cli.operations.validation.exact_record_validation` | T11 | `go -C drcli test ./internal/validation -count=1` |
| `spec:drcli.design_records_cli.operations.validation` | T11 | `go -C drcli test ./internal/validation -count=1` |
| `spec:drcli.design_records_cli.operations.validation.scope_validation` | T11 | `go -C drcli test ./internal/validation -count=1` |
| `spec:drcli.design_records_cli.operations.validation.validation_output_limits` | T11 | `go -C drcli test ./internal/validation -count=1` |


## Task Candidates

These are materialized Task records (not proposals); the exact `depends_on` and `outputs` metadata of each Task is authoritative.

| Task ID | task type | sole outcome | predecessor Task numbers |
|---|---|---|---|
| `DRCLI-TASK-CLI-003-01` | coordination | Graph, writers, spec trace, executor cards | - |
| `DRCLI-TASK-CLI-003-02` | review | Independent graph readiness verdict | 01 |
| `DRCLI-TASK-CLI-003-03` | synchronization | Graph PASS and initial wave release | 02 |
| `DRCLI-TASK-CLI-003-04` | implementation | Artifact grammar/markdown and nested Go module | 03 |
| `DRCLI-TASK-CLI-003-05` | implementation | Current repo/root/ignore/admission/scope | 04 |
| `DRCLI-TASK-CLI-003-06` | implementation | Diagnostic catalog and envelope | 04 |
| `DRCLI-TASK-CLI-003-07` | implementation | Scope/list/tree | 05,06 |
| `DRCLI-TASK-CLI-003-08` | implementation | Batch retrieval and H2 sections | 05,06 |
| `DRCLI-TASK-CLI-003-09` | implementation | Literal/RE2 search | 05,06 |
| `DRCLI-TASK-CLI-003-10` | implementation | Current/V01 ref resolution | 05,06 |
| `DRCLI-TASK-CLI-003-11` | implementation | Canonical validation findings, scopes and exact refs | 04,05,06 |
| `DRCLI-TASK-CLI-003-12` | implementation | Portable authoring guidance | 04,05,06 |
| `DRCLI-TASK-CLI-003-13` | implementation | Process-independent body and proposal state | 04 |
| `DRCLI-TASK-CLI-003-14` | implementation | Create/update/guarded write/proposal actions | 04,05,06,11,13 |
| `DRCLI-TASK-CLI-003-15` | implementation | CLI command/usage/help/body/UTF-8 streams | 04,05,06,07,08,09,10,11,12,14 |
| `DRCLI-TASK-CLI-003-16` | implementation | Windows/Linux copyable distributions | 12,13,15 |
| `DRCLI-TASK-CLI-003-17` | implementation | End-to-end fixtures and subprocess harness | 16 |
| `DRCLI-TASK-CLI-003-18` | verification | Objective Windows+Linux integration PASS/FAIL/BLOCKED | 17 |
| `DRCLI-TASK-CLI-003-19` | review | Independent complete code PASS/NEEDS REVISION | 18 |
| `DRCLI-TASK-CLI-003-20` | synchronization | Accepted release/finding closure and Work Item DONE | 19 |

Correction and independent finding-closure review Tasks are **not** allocated before a named T02 or T19 finding. A finding-driven change to the frozen graph requires new coordination and its own independent review, without rewriting a completed Task.


## Completion Condition

- Every accepted DRCLI Specification leaf has traceable production behavior and focused verification.
- Portable command hierarchy, help without repository access, actionable human/JSON usage errors, cwd/`--repo`, Git-compatible ignore/`--no-ignore`, source-root discovery, namespace identity, admission, conflicts, and scopes work.
- Listing, trees, exact and section retrieval, lexical search, reference resolution, broad/exact validation, guidance list/exact read, direct create/update, optional proposals, and inline/stdin/cache body sources meet the accepted contracts.
- Retained caches and proposals survive process and session termination, preserve correct repository binding, and report trustworthy negative results, ordering, streams, write state, and numeric exit outcomes.
- Windows and Linux copyable distributions bundle standards and run without separately installed runtimes; portable tests include independent-process persistence and cross-platform parity.
- Exact writers and task dependencies were independently reviewed and released before production implementation. Focused tests, integration gate, independent code review, and required finding closure all pass.
- Final Work Item Evidence and Task lifecycle are synchronized without modifying completed design history, DRMCP, or PRODUCT.

## Evidence

### Implementation baseline — 2026-10-08

- `drcli/` currently contains only `records/`; no DRCLI production source, tests, local module/build config, or distribution scripts were observed.
- Root `go.mod` (Go 1.22) exists as adjacent repository infrastructure, but does not select DRCLI's implementation technology or license reuse of DRMCP internals.
- Completed `DRCLI-WORK-CLI-002` and 82-file `spec:drcli.design_records_cli` tree are read-only design evidence; both required historical findings are CLOSED after T20 PASS.
- Existing DRCLI design records are untracked in Git. Preserve all existing files; no stage or commit.
- DRMCP authoring MCP remains non-operational per agent authoring policy; use scoped filesystem authoring after reading PRODUCT writing/Work Item/Task standards.
- T01 owns the remaining implementation-local architecture, runnable test/build selection, exact Task graph, and unique write boundaries; no external observable design decision is currently identified.

### T01 coordination materialization — 2026-10-08 (validation blocked)

- T04-T20 have been authored as non-overlapping executor/gate Tasks; T01 status is `blocked` by unavailable standalone responsibility validation, not `done`. Independent T02 graph review and T03 graph release remain not started.
- Scoped Windows check after removal of a partial duplicate: 20 unique Task records, 42 dependency edges with no cycle or unresolved ID, 59 distinct planned output targets without duplicate writer, parent-to-Task relation exact, 82/82 accepted spec refs matched uniquely, and no production Go files.
- Full primary Specification trace and focused command per record are in `## Task flow`; code/API layouts and source/test/fixture/packager writers are fixed in T04-T17. The verification gate is T18, independent code review is T19, final synchronization T20. No finding-specific corrections or independent closure review Tasks are speculatively reserved.
- The PRODUCT-required standalone TRV responsibility validator has no available executable or connected action in this Windows environment (TRV repository namespace contains records only); per-criterion result and overall compliant disposition therefore remain **unavailable**, not PASS. Before T01 DONE and independent T02, it must evaluate each newly authored Task and final T01 Evidence; any violations require explicit human disposition.
- No Linux watcher marker may assert T01 DONE while this mandatory validation gate is unavailable.
