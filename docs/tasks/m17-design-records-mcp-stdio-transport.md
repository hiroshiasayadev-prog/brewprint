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

Implementation note: existing `internal/mcp` remains brewprint YAML semantic model / `QueryService`-oriented, so M17 does not directly reuse or generalize it. Design Records MCP uses a separate minimal JSON-RPC / stdio adapter in `internal/designrecordsmcp`; the existing brewprint MCP is only a wire-shape reference, while `internal/designrecords` handlers stay transport independent.

- [x] Check whether existing brewprint MCP stdio code can be reused
- [x] Decide whether to use an MCP Go library or existing local JSON-RPC code
- [x] Decide the adapter boundary for tool registration, request decoding, and response encoding

Done criteria:

- [x] stdio MCP implementation approach is decided
- [x] rationale is recorded in task notes or implementation comments
- [x] responsibility boundary with existing brewprint MCP remains clear

---

## Phase 1: stdio server skeleton

Turn `cmd/design-records-mcp` into a stdio MCP server skeleton.

- [x] Preserve existing repository root resolution behavior
- [x] Preserve existing repository root resolution behavior without building the index during initialize or tools/list
- [x] Respond to MCP initialize request
- [x] List Design Records MCP tools without building the index
- [x] Decide how server mode and summary/debug mode coexist

Done criteria:

- [x] MCP host can start the process as a server
- [x] Startup does not depend on brewprint YAML semantic build
- [x] Summary/debug output does not corrupt stdio MCP protocol

---

## Phase 2: tool registration

Register M16 handlers as MCP tools.

- [x] Register `list_records`
- [x] Register `validate_records`
- [x] Register `get_record`
- [x] Register `suggest_next_record`
- [x] Build the Design Records index for each `tools/call` before invoking a handler
- [x] Do not keep a long-lived Design Records index cache
- [x] Match request schema in `docs/spec/design-records-mcp/tools.md`
- [x] Match response shape in `docs/spec/design-records-mcp/tools.md`

Done criteria:

- [x] MCP host can call `list_records`
- [x] MCP host can call `validate_records`
- [x] MCP host can call `get_record`
- [x] MCP host can call `suggest_next_record`
- [x] each `tools/call` observes the latest docs by rebuilding the index

---

## Phase 3: error mapping

Map M16 `ToolError` to MCP tool errors.

- [x] `invalid_request`
- [x] `record_not_found`
- [x] `unsupported_kind`
- [x] `id_range_requires_decision_kind`

Rules:

- [x] Tool execution errors and validation diagnostics are not mixed
- [x] `validate_records` diagnostics remain normal tool response data
- [x] Request decode/schema errors are machine-readable
- [x] Error response shape matches `docs/spec/design-records-mcp/tools.md`
- [x] Unexpected failures return errors instead of panics

Done criteria:

- [x] invalid request is returned as a tool error
- [x] `record_not_found` is returned as a tool error
- [x] validation diagnostics are not protocol errors

---

## Phase 4: smoke / integration tests

Add minimal stdio MCP smoke coverage.

- [x] Process startup smoke test or script
- [x] initialize / tools list / tool call minimum check
- [x] `list_records` smoke test
- [x] `get_record` smoke test
- [x] `validate_records` smoke test
- [x] `suggest_next_record` smoke test
- [x] repository root argument smoke test
- [x] smoke or unit test confirms index is rebuilt per `tools/call`

Done criteria:

- [x] `go test ./internal/designrecords` passes
- [x] `go test ./...` passes
- [x] stdio MCP smoke check passes
- [x] Windows PowerShell startup works

---

## Phase 5: usage note

Add only minimal usage documentation.

- [x] Point users to `docs/spec/design-records-mcp/tools.md` for tool details
- [x] Do not put build commands or implementation history in `doc-policy.md`
- [x] If `doc-policy.md` is updated, keep it to a tool-spec reference only

Done criteria:

- [x] Tool spec reference is easy to find
- [x] `doc-policy.md` is not expanded into transport implementation notes

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
