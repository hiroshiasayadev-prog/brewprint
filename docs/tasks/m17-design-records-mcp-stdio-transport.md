# Milestone 17: Design Records MCP stdio transport

- **status**: open
- **scope**: cmd/design-records-mcp / internal/designrecords / tests / minimal docs update
- **source**: M16 Design Records MCP MVP / docs/spec/design-records-mcp/**
- **last_updated**: 2026-05-13

---

## Context

M16 completed the Design Records MCP MVP handler layer in `internal/designrecords`.

M17 exposes those handlers through stdio MCP.

Implemented handlers from M16:

- `list_records`
- `validate_records`
- `get_record`
- `suggest_next_record`

`cmd/design-records-mcp/` currently runs independently as an index summary binary. It is not yet a stdio MCP server.

---

## Goal

`cmd/design-records-mcp` can run as a stdio MCP server and expose Design Records MCP tools.

Target tools:

- `list_records`
- `validate_records`
- `get_record`
- `suggest_next_record`

---

## Non-goals

M17 does not include:

- HTTP transport
- ChatGPT connector compatibility
- write tools
- ADR/spec body generation
- git operations
- integration into existing brewprint MCP
- `cmd/brewprint` subcommand integration

---

## Design policy

- Keep `internal/designrecords` handlers transport independent.
- Keep the stdio adapter thin.
- Preserve existing repository root resolution behavior.
- Keep tool errors separate from validation diagnostics.
- Preserve the read-only boundary.
- Match tool schemas in `docs/spec/design-records-mcp/tools.md`.
- Do not let summary/debug output break stdio MCP protocol output.
- Build the Design Records index per `tools/call`; do not keep a long-lived index cache in M17.

---

## Phase 0: transport implementation choice

Choose how to implement the stdio MCP server.

- [ ] Check whether existing brewprint MCP stdio code can be reused
- [ ] Decide whether to use an MCP Go library or existing local JSON-RPC code
- [ ] Decide the adapter boundary for tool registration, request decoding, and response encoding

Done criteria:

- [ ] stdio MCP implementation approach is decided
- [ ] rationale is recorded in task notes or implementation comments
- [ ] responsibility boundary with existing brewprint MCP remains clear

---

## Phase 1: stdio server skeleton

Turn `cmd/design-records-mcp` into a stdio MCP server skeleton.

- [ ] Preserve existing repository root resolution behavior
- [ ] Preserve existing repository root resolution behavior without building the index during initialize or tools/list
- [ ] Respond to MCP initialize request
- [ ] List Design Records MCP tools without building the index
- [ ] Decide how server mode and summary/debug mode coexist

Done criteria:

- [ ] MCP host can start the process as a server
- [ ] Startup does not depend on brewprint YAML semantic build
- [ ] Summary/debug output does not corrupt stdio MCP protocol

---

## Phase 2: tool registration

Register M16 handlers as MCP tools.

- [ ] Register `list_records`
- [ ] Register `validate_records`
- [ ] Register `get_record`
- [ ] Register `suggest_next_record`
- [ ] Build the Design Records index for each `tools/call` before invoking a handler
- [ ] Do not keep a long-lived Design Records index cache
- [ ] Match request schema in `docs/spec/design-records-mcp/tools.md`
- [ ] Match response shape in `docs/spec/design-records-mcp/tools.md`

Done criteria:

- [ ] MCP host can call `list_records`
- [ ] MCP host can call `validate_records`
- [ ] MCP host can call `get_record`
- [ ] MCP host can call `suggest_next_record`
- [ ] each `tools/call` observes the latest docs by rebuilding the index

---

## Phase 3: error mapping

Map M16 `ToolError` to MCP tool errors.

- [ ] `invalid_request`
- [ ] `record_not_found`
- [ ] `unsupported_kind`
- [ ] `id_range_requires_decision_kind`

Rules:

- [ ] Tool execution errors and validation diagnostics are not mixed
- [ ] `validate_records` diagnostics remain normal tool response data
- [ ] Request decode/schema errors are machine-readable
- [ ] Error response shape matches `docs/spec/design-records-mcp/tools.md`
- [ ] Unexpected failures return errors instead of panics

Done criteria:

- [ ] invalid request is returned as a tool error
- [ ] `record_not_found` is returned as a tool error
- [ ] validation diagnostics are not protocol errors

---

## Phase 4: smoke / integration tests

Add minimal stdio MCP smoke coverage.

- [ ] Process startup smoke test or script
- [ ] initialize / tools list / tool call minimum check
- [ ] `list_records` smoke test
- [ ] `get_record` smoke test
- [ ] `validate_records` smoke test
- [ ] `suggest_next_record` smoke test
- [ ] repository root argument smoke test
- [ ] smoke or unit test confirms index is rebuilt per `tools/call`

Done criteria:

- [ ] `go test ./internal/designrecords` passes
- [ ] `go test ./...` passes
- [ ] stdio MCP smoke check passes
- [ ] Windows PowerShell startup works

---

## Phase 5: usage note

Add only minimal usage documentation.

- [ ] Point users to `docs/spec/design-records-mcp/tools.md` for tool details
- [ ] Do not put build commands or implementation history in `doc-policy.md`
- [ ] If `doc-policy.md` is updated, keep it to a tool-spec reference only

Done criteria:

- [ ] Tool spec reference is easy to find
- [ ] `doc-policy.md` is not expanded into transport implementation notes

---

## Done criteria for M17

M17 is done when:

- [ ] `cmd/design-records-mcp` can run as stdio MCP server
- [ ] MCP host can list Design Records MCP tools
- [ ] MCP host can call `list_records`
- [ ] MCP host can call `validate_records`
- [ ] MCP host can call `get_record`
- [ ] MCP host can call `suggest_next_record`
- [ ] tool schemas match `docs/spec/design-records-mcp/tools.md`
- [ ] tool errors are machine-readable
- [ ] validation diagnostics remain normal tool responses
- [ ] each `tools/call` rebuilds the Design Records index instead of using a long-lived cache
- [ ] summary/debug output does not corrupt stdio MCP protocol
- [ ] implementation remains read-only
- [ ] implementation does not depend on existing brewprint YAML semantic build / `ResolvedProject`
- [ ] tests or smoke checks cover stdio MCP behavior

---

## Follow-up candidates

- HTTP transport
- ChatGPT connector compatibility
- shared server mode
- project isolation
- integration with existing brewprint MCP launcher
- `cmd/brewprint` subcommand integration
