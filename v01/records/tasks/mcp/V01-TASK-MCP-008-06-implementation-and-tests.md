# V01-TASK-MCP-008-06: Implementation and tests

- **id**: V01-TASK-MCP-008-06
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-MCP-008
- **source_requirement**: V01-REQ-MCP-008
- **estimate**: 2d-4d
- **depends_on**:
  - V01-TASK-MCP-008-05
- **outputs**:
  - Implemented Design Records MCP authoring transaction tools
  - Proposal / body cache support
  - Authoring transaction tests
  - Updated runtime MCP tool registration

## Goal

Accepted `V01-ADR-093` and reviewed `SPEC-design-records-mcp-tools` / `SPEC-design-records-mcp-schema` に基づき、Design Records MCP authoring transaction MVP を実装し、targeted tests で public contract を検証する。

## Scope

Implement the authoring transaction MVP only.

Included:

- `propose_record_create`
- `propose_record_update`
- `get_proposed_write`
- `accept_proposed_write`
- `discard_proposed_write`
- proposal store with 3 days retention
- body cache with 3 days retention
- artifact-oriented create target resolution for `decision`, `requirement`, `work_item`, and `task`
- update support for existing `decision`, `spec`, `requirement`, `work_item`, and `task` records
- kind-specific metadata block replacement
- named section replacement with single unambiguous ATX section resolution
- accept-time stale / changed-target / ID-collision guards
- `written: true/false` response semantics
- validation / repair guidance after accept

Excluded:

- spec skeleton creation
- `SPEC-new` placeholder create
- investigation creation
- generic filesystem write tools
- path-first authoring APIs
- add/remove relation convenience operations
- partial Markdown AST editing
- general-purpose multi-record atomic transactions with rollback semantics
- automatic close cascades
- automatic rollback after post-write validation failure
- formatter integration
- force-accepting invalid proposals
- V01-REQ-MCP-010 spec domain tree / placement discovery support

## Work

- Read the current instructions and relevant artifacts before implementation:
  - `AGENTS.md`
  - `docs/prompt_chappy.md`
  - `docs/doc-policy.md`
  - `docs/adr/093-design-records-mcp-authoring-transaction-model.md`
  - `docs/spec/design-records-mcp/tools.md`
  - `docs/spec/design-records-mcp/schema.md`
  - `V01-REQ-MCP-008`
  - `V01-WORK-MCP-008`
  - `V01-TASK-MCP-008-01` through `V01-TASK-MCP-008-05`
- Add internal representation for proposals and body cache.
- Add request / response types for the new authoring tools.
- Add proposal creation for supported create targets.
- Add update proposal creation for metadata block replacement and named section replacement.
- Add proposal retrieval, discard, and accept logic.
- Ensure proposal creation does not write repository files.
- Ensure only `accept_proposed_write` writes repository files.
- Implement body source validation:
  - reject both `body` and `body_cache_id`;
  - allow neither only when the operation does not require a body;
  - reject unknown / expired body cache IDs;
  - keep body cache reusable within 3 days.
- Implement `new` placeholder validation:
  - accept literal `new` token only in create sequence position;
  - reject `SPEC-new`;
  - reject malformed `newish` / `newer` style tokens;
  - reject `new` in update operations.
- Implement task parent-aware creation and required reciprocal metadata handling according to the spec.
- Implement metadata block replacement with spec recognized-field scoped update and unknown auxiliary YAML front matter preservation.
- Implement named section replacement using the same ATX heading source rules as record headings.
- Register the new MCP tools.
- Add targeted unit / MCP call tests for the public contract.

## Required tests

At minimum, cover:

- proposal creation does not write files;
- `accept_proposed_write` writes the proposed diff;
- proposal cannot be accepted twice;
- discard prevents later accept;
- expired / unknown proposal diagnostics;
- stale / changed target rejection before writing;
- create ID collision rejection before writing;
- `ADR-new`, `REQ-<DOMAIN>-new`, `WORK-<DOMAIN>-new`, and `TASK-<DOMAIN>-<WORK-SEQUENCE>-new` resolution;
- `SPEC-new` rejection;
- update operation `new` rejection;
- invalid `newish` / `newer` ID rejection;
- task create requires `parent_id` and does not infer parent relation from ID shape alone;
- required reciprocal workflow metadata handling;
- `body` / `body_cache_id` mutual exclusion;
- neither body source allowed only for no-body operations;
- body cache retry and reuse within retention;
- expired / unknown body cache diagnostics;
- metadata block replacement for workflow artifact records;
- scoped spec metadata replacement preserves unknown YAML front matter fields;
- named section replacement succeeds for exactly one ATX section;
- zero-match and multi-match section selectors do not create proposals;
- fenced code block / YAML front matter headings are ignored by section selector;
- accepted write with post-write validation failure returns `written: true` and repair guidance without rollback;
- existing read / navigation / guidance tools continue to pass current tests.

## Done condition

- New authoring transaction tools are implemented and registered.
- Public behavior matches `SPEC-design-records-mcp-tools` and `SPEC-design-records-mcp-schema`.
- MVP exclusions are enforced, especially `SPEC-new` / spec skeleton creation and path-first authoring APIs.
- Targeted tests pass.
- Existing read/navigation/guidance tests still pass.
- Runtime MCP call smoke tests for representative cases are either completed or explicitly deferred to `V01-TASK-MCP-008-07`.
- Evidence records files changed, tests run, and any known follow-up gaps.

## Verification

Suggested commands:

```powershell
go test ./internal/designrecords ./internal/designrecordsmcp
```

Also verify with Design Records MCP where available:

- `validate_records` for `V01-ADR-093`
- `validate_records` for `V01-REQ-MCP-008`
- `validate_records` for `V01-WORK-MCP-008`
- `validate_records` for this task

## Evidence

2026-06-02: Implementation and targeted tests completed.

2026-06-02 pre-commit review revision: Addressed independent review findings for partial multi-file write reporting, missing required metadata diagnostic categories on ADR / workflow metadata replacement, and candidate headings on failed named section selectors. Added targeted regression coverage and reran the required package tests plus `git diff --check`.

### Files changed

- `internal/designrecords/authoring.go`
- `internal/designrecords/authoring_test.go`
- `internal/designrecords/types.go`
- `internal/designrecordsmcp/server.go`
- `internal/designrecordsmcp/tools.go`
- `internal/designrecordsmcp/tools_call.go`
- `internal/designrecordsmcp/tools_call_test.go`
- `internal/designrecordsmcp/jsonrpc_test.go`
- `docs/tasks/mcp/TASK-MCP-008-06-implementation-and-tests.md`

### Implemented tools

Registered and dispatched the MVP authoring transaction tools:

- `propose_record_create`
- `propose_record_update`
- `get_proposed_write`
- `accept_proposed_write`
- `discard_proposed_write`

Proposal creation is no-write. Only `accept_proposed_write` writes repository files after accept-time guards pass.

### Proposal store / body cache

- Added an in-memory server-owned authoring store.
- Proposal retention is 3 days and expired proposals return `proposal_expired`.
- Body cache retention is 3 days.
- Body cache entries are reusable within retention, including after successful proposal creation.
- `body` / `body_cache_id` mutual exclusion is enforced.
- Unknown / expired body cache IDs return diagnostics and do not create proposals.

### Create / update support

Create support:

- `decision`
- `requirement`
- `work_item`
- `task`

Implemented `new` placeholder resolution for:

- `ADR-new`
- `REQ-<DOMAIN>-new`
- `WORK-<DOMAIN>-new`
- `TASK-<DOMAIN>-<WORK-SEQUENCE>-new`

Update support:

- `decision`
- `spec`
- `requirement`
- `work_item`
- `task`

Implemented metadata block replacement and named ATX section replacement. Spec metadata replacement preserves unknown / auxiliary YAML front matter fields while replacing recognized fields.

### Exclusion confirmation

`SPEC-new`, spec skeleton creation, investigation creation, path-first authoring APIs, generic filesystem write tools, add/remove relation helpers, formatter integration, and force-accept are not implemented.

### Tests added / updated

Added direct authoring tests for:

- proposal no-write behavior and accept writes
- accept-twice rejection
- discard preventing accept
- unknown / expired proposal diagnostics
- stale target and create ID collision rejection before writing
- `ADR-new`, `REQ-<DOMAIN>-new`, `WORK-<DOMAIN>-new`, and `TASK-<DOMAIN>-<WORK-SEQUENCE>-new`
- `SPEC-new`, `newish`, `newer`, and update-time `new` rejection
- task `parent_id` requirement and explicit `fields.work_item` relation requirement
- required reciprocal workflow metadata inclusion and required follow-up rejection
- body source exclusion, no-body rejection for section replacement, body cache reuse, unknown / expired body cache diagnostics
- workflow metadata replacement and scoped spec metadata replacement
- named section replacement, zero-match and multi-match rejection, YAML front matter / fenced heading exclusion
- accepted write with post-write validation failure returning `written: true` and repair guidance without rollback
- partial multi-file accept failure reporting `written: true` with `files_written` and retiring the proposal from unsafe retry
- missing required ADR / workflow metadata fields returning `missing_required_metadata`
- zero-match and ambiguous section selector diagnostics returning `candidate_headings`

Updated MCP tests for tool registration and a JSON-RPC `tools/call` authoring transaction smoke path.

### Commands run

- `go test ./internal/designrecords`
  - passed
- `go test ./internal/designrecordsmcp`
  - passed
- `go test ./internal/designrecords ./internal/designrecordsmcp`
  - passed
- `go test ./internal/designrecords ./internal/designrecordsmcp -run Authoring`
  - passed
- `git diff --check -- internal/designrecords internal/designrecordsmcp docs/tasks/mcp/TASK-MCP-008-06-implementation-and-tests.md`
  - passed; Git reported line-ending conversion warnings for touched Go files.

### Design Records MCP validation

- `validate_records(kind: decision, id_range: V01-ADR-093..V01-ADR-093)` returned `ok: true`.
- `validate_records(kind: requirement, id_range: V01-REQ-MCP-008..V01-REQ-MCP-008)` returned `ok: true`.
- `validate_records(kind: work_item, id_range: V01-WORK-MCP-008..V01-WORK-MCP-008)` returned `ok: true`.
- `validate_records(kind: task, id_range: V01-TASK-MCP-008-06..V01-TASK-MCP-008-06)` returned `ok: true`.

### Runtime MCP smoke

Runtime stdio smoke tests are deferred to `V01-TASK-MCP-008-07`.
This task added a JSON-RPC `tools/call` test covering propose, get, accept, and actual file write behavior.

### Pre-existing dirty state

Before implementation, the worktree already had unrelated dirty / untracked files including:

- `docs/adr/073-tagged-union-model.md`
- `docs/spec/design-records-mcp/schema.md`
- `docs/spec/design-records-mcp/tools.md`
- `docs/work-items/data/WORK-DATA-010-tagged-union-and-discriminator-payload-support.md`
- `docs/adr/093-design-records-mcp-authoring-transaction-model.md`
- `docs/requirements/mcp/REQ-MCP-008-design-records-authoring-transaction-support.md`
- `docs/requirements/mcp/REQ-MCP-009-authoring-guidance-canonicalization-and-legacy-cleanup.md`
- `docs/requirements/mcp/REQ-MCP-010-spec-domain-tree-and-placement-discovery-support.md`
- `docs/tasks/data/TASK-DATA-003-01-adr-075-dependency-and-split-review.md`
- `docs/tasks/data/TASK-DATA-010-01-adr-073-acceptance-split-review.md`
- `docs/tasks/data/TASK-DATA-010-02-adr-073-revision-and-spec-diagnostics-alignment.md`
- `docs/tasks/mcp/V01-TASK-MCP-008-01` through `V01-TASK-MCP-008-05`
- `docs/tasks/mcp/TASK-MCP-008-06-implementation-and-tests.md`
- `docs/tasks/mcp/V01-TASK-MCP-009-01` through `V01-TASK-MCP-009-04`
- `docs/uc/003-task-file-helper-model/`
- `docs/work-items/mcp/WORK-MCP-008-design-records-authoring-transaction-support.md`
- `docs/work-items/mcp/WORK-MCP-009-authoring-guidance-canonicalization-and-legacy-cleanup.md`

Only `V01-TASK-MCP-008-06` was edited among the pre-existing docs.

Final `git status --short` also showed unrelated modified guide / README docs that were not part of the initial status snapshot and were not edited by this task:

- `docs/adr-authoring-guide.md`
- `docs/investigations/README.md`
- `docs/requirements/README.md`
- `docs/spec-authoring-guide.md`
- `docs/tasks/README.md`
- `docs/work-items/README.md`

### Known follow-ups

- `V01-TASK-MCP-008-07` should run runtime MCP stdio smoke tests and close evidence.
- `V01-WORK-MCP-008` was not marked done.
- `V01-TASK-MCP-008-07` was not created by this task.
