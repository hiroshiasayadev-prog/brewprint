# TASK-MCP-008-07: Runtime verification and close evidence

- **id**: TASK-MCP-008-07
- **status**: done
- **date**: 2026-06-02
- **work_item**: WORK-MCP-008
- **source_requirement**: REQ-MCP-008
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-MCP-008-06
- **outputs**:
  - Runtime MCP smoke verification result
  - WORK-MCP-008 close evidence
  - REQ-MCP-008 close evidence

## Goal

`TASK-MCP-008-06` で実装された Design Records MCP authoring transaction MVP を runtime MCP stdio / JSON-RPC 経由で検証し、WORK-MCP-008 / REQ-MCP-008 を close できる evidence を揃える。

## Scope

This task verifies runtime behavior and close readiness. It should not introduce new feature work.

Verify the implemented authoring transaction tools:

- `propose_record_create`
- `propose_record_update`
- `get_proposed_write`
- `accept_proposed_write`
- `discard_proposed_write`

## Work

- Inspect current `git status --short` and preserve unrelated dirty state.
- Read the required project instructions and task context:
  - `AGENTS.md`
  - `docs/prompt_chappy.md`
  - `docs/doc-policy.md`
  - `TASK-MCP-008-06`
  - `ADR-093`
  - `SPEC-design-records-mcp-tools`
  - `SPEC-design-records-mcp-schema`
- Run targeted tests if needed to confirm the implementation still passes.
- Start Design Records MCP through the runtime entrypoint used by existing MCP call tests or manual stdio JSON-RPC verification.
- Verify representative runtime cases.
- Record evidence and any limitations.
- If all checks pass, update close evidence and status for this task.
- Do not mark `WORK-MCP-008` or `REQ-MCP-008` done unless the user explicitly asks for close after reviewing evidence.

## Required runtime cases

At minimum, verify:

- tools/list includes the five new authoring tools;
- `propose_record_create` returns a proposal for a supported create case;
- proposal creation does not write repository files;
- `get_proposed_write` retrieves the proposal;
- `discard_proposed_write` makes later accept fail with `written: false`;
- `accept_proposed_write` writes a proposal in a safe temporary fixture or isolated test environment;
- `SPEC-new` / spec skeleton create is rejected;
- request with both `body` and `body_cache_id` is rejected;
- named section zero-match or ambiguous selector does not write;
- validate_records still passes for `ADR-093`, `REQ-MCP-008`, `WORK-MCP-008`, and this task.

## Safety / fixture expectations

Runtime smoke should avoid modifying real project records unless the test uses a controlled temporary repository or fixture root.

If runtime verification needs to create files, use a temporary copy / fixture root and report that explicitly.

Do not stage or commit.
Do not clean unrelated dirty files.
Do not modify implementation unless a small verification bug fix is absolutely necessary; if modified, report it clearly.

## Done condition

- Runtime MCP smoke checks for the representative cases are recorded.
- Targeted tests pass after runtime smoke.
- `validate_records` for relevant artifacts returns ok.
- Any fixture / temporary repo usage is documented.
- Known limitations or follow-ups are recorded.
- This task is set to `done` only if runtime verification is complete.
- WORK-MCP-008 and REQ-MCP-008 are ready for close review.

## Verification

Suggested commands:

```powershell
go test ./internal/designrecords ./internal/designrecordsmcp
```

Use existing MCP stdio / JSON-RPC runtime approach where possible. Prefer existing test helpers or documented commands over ad hoc scripts.

Design Records MCP validation targets:

- `ADR-093`
- `REQ-MCP-008`
- `WORK-MCP-008`
- `TASK-MCP-008-07`

## Evidence

2026-06-02: Runtime stdio / JSON-RPC verification completed.

### Pre-existing dirty state

Before verification, `git status --short` already showed dirty / untracked files. They were preserved; nothing was staged, cleaned, reverted, or committed.

Summary of pre-existing state:

- Modified docs / specs / guides / READMEs included `docs/TASKS.md`, authoring guides, `docs/spec/design-records-mcp/schema.md`, `docs/spec/design-records-mcp/tools.md`, project concept specs, and `docs/work-items/data/WORK-DATA-010-tagged-union-and-discriminator-payload-support.md`.
- Modified implementation files included `internal/designrecords/types.go`, `internal/designrecordsmcp/jsonrpc_test.go`, `internal/designrecordsmcp/server.go`, `internal/designrecordsmcp/tools.go`, `internal/designrecordsmcp/tools_call.go`, and `internal/designrecordsmcp/tools_call_test.go`.
- Untracked MCP authoring transaction artifacts included `ADR-093`, `REQ-MCP-008`, `REQ-MCP-010`, `TASK-MCP-008-01` through `TASK-MCP-008-07`, `WORK-MCP-008`, and implementation files `internal/designrecords/authoring.go` / `internal/designrecords/authoring_test.go`.
- Additional unrelated untracked MCP-009, DATA task, UC, and work item files were also present.

After runtime smoke and before this evidence edit, `git status --short` matched the same dirty-state summary. The smoke test did not modify real project records.

### Runtime command and fixture

Runtime command used:

```powershell
go run ./cmd/design-records-mcp --root C:\Users\imved\AppData\Local\Temp\brewprint-mcp00807-kf9zjpr_
```

The server was exercised through stdio JSON-RPC lines with `initialize`, `notifications/initialized`, `tools/list`, and `tools/call`.

A temporary fixture root was used:

```text
C:\Users\imved\AppData\Local\Temp\brewprint-mcp00807-kf9zjpr_
```

The fixture contained minimal ADR / requirement / work item / task / spec records. All create, update, discard, and accept smoke writes targeted only this fixture root.

### Tools discovery

`tools/list` returned 13 tools and included all required authoring transaction tools:

- `propose_record_create`
- `propose_record_update`
- `get_proposed_write`
- `accept_proposed_write`
- `discard_proposed_write`

### Runtime cases

`propose_record_create` for fixture task creation:

- Request: `kind: task`, `id: TASK-MCP-001-new`, `parent_id: WORK-MCP-001`.
- Result: `proposal_created: true`.
- Proposal ID was returned.
- Target resolved to `TASK-MCP-001-02`.
- Target path returned: `docs/tasks/mcp/TASK-MCP-001-02-second-task.md`.
- Diff text was returned.
- Note returned: `No repository files have been written. Call accept_proposed_write with this proposal_id to apply the diff.`
- Target file did not exist before accept.

`get_proposed_write`:

- Retrieved the same proposal ID.
- State was `proposed`.
- Target and diff were present.

Discard prevents accept:

- Created a second task proposal.
- `discard_proposed_write` returned `discarded: true`, `state: discarded`, `written: false`.
- Later `accept_proposed_write` for the discarded proposal returned `written: false`, `state: discarded`.
- Diagnostic category: `proposal_discarded`.
- Discarded proposal target file was not created.

Accept writes safely:

- Accepted the first task create proposal against the fixture root.
- Response returned `written: true`, `state: accepted`.
- Files written:
  - `docs/tasks/mcp/TASK-MCP-001-02-second-task.md`
  - `docs/work-items/mcp/WORK-MCP-001-first-work.md`
- The expected task file existed after accept.
- Post-write validation returned `ok: true`.

Representative rejection cases:

- `SPEC-new` / spec skeleton create was rejected; no proposal was created. Diagnostic category: `unsupported_kind`.
- Request with both `body` and `body_cache_id` was rejected; no proposal or body cache was created. Diagnostic category: `invalid_body_source`.
- Named section zero-match was rejected; no proposal was created and no file was written. Diagnostic category: `section_selector_no_match`; a body cache was returned for retry.
- Named section ambiguous selector was rejected after fixture-only duplicate heading setup; no proposal was created and no file was written. Diagnostic category: `section_selector_ambiguous`; a body cache was returned for retry.
- Update operation with `TASK-MCP-001-new` placeholder was rejected; no proposal was created. Diagnostic category: `invalid_request`.

### Real repository safety confirmation

Runtime smoke used only the temporary fixture root. It did not require or perform authoring writes against the real repository.

`git status --short` after the runtime smoke matched the pre-existing dirty-state summary, confirming that real project records were not accidentally modified by the smoke test. This task update is the only intentional repository file edit made during verification.

### Tests

```powershell
go test ./internal/designrecords ./internal/designrecordsmcp
```

Result: passed.

```powershell
go test ./internal/designrecords ./internal/designrecordsmcp -run Authoring
```

Result: passed.

### Design Records validation

Validation was run through Design Records MCP:

- `ADR-093`: `ok: true`, diagnostics `null`.
- `REQ-MCP-008`: `ok: true`, diagnostics `null`.
- `WORK-MCP-008`: `ok: true`, diagnostics `null`.
- `TASK-MCP-008-07`: `ok: true`, diagnostics `null`.

### Implementation changes

None. No implementation files were edited during this runtime verification task.

### Known limitations / follow-ups

- Smoke coverage was representative, not exhaustive; deeper stale-target / collision / expiry behavior remains covered by targeted tests from `TASK-MCP-008-06`.
- `SPEC-new` rejection currently surfaces as `unsupported_kind`; this still satisfies MVP rejection, but callers should not depend on a more specific diagnostic unless the public spec later requires one.
- `WORK-MCP-008` was not marked done.
- `REQ-MCP-008` was not marked accepted / done.

### Close readiness

`WORK-MCP-008` and `REQ-MCP-008` are ready for close review based on accepted ADR coverage, spec reflection, implementation tests, runtime MCP smoke verification, and successful Design Records validation.
