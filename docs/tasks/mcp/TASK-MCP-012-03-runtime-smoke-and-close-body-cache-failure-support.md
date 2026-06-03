# TASK-MCP-012-03: runtime smoke and close body cache failure support

- **id**: TASK-MCP-012-03
- **status**: done
- **date**: 2026-06-02
- **work_item**: WORK-MCP-012
- **source_requirement**: REQ-MCP-015
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-MCP-012-02
- **outputs**:
  - Runtime smoke results for retryable failed-propose body_cache behavior
  - Close evidence for REQ-MCP-015 / WORK-MCP-012

## Goal
Verify the REQ-MCP-015 body-cache failure behavior through the actual Design Records MCP JSON-RPC entrypoint and record close evidence for TASK-MCP-012 / WORK-MCP-012 / REQ-MCP-015.

## Work
- Run the actual MCP server via `go run ./cmd/design-records-mcp --root .` and send JSON-RPC `initialize`, `tools/list`, and representative `tools/call` requests.
- Verify that `propose_record_update` no-match with submitted `body` returns `proposal_created:false` and `body_cache`.
- Verify that `propose_record_create` `fields + body` with stale full-record body returns `proposal_created:false` and `body_cache`.
- Record close evidence and explicitly separate follow-up create-contract work from REQ-MCP-015 close scope.

## Done condition
- Runtime smoke response demonstrates `body_cache` is present for representative failed-propose cases after submitted body receipt.
- Close evidence is recorded on the task, work item, and requirement.
- Follow-up create-contract changes are explicitly excluded from this close.

## Verification
- Confirmed `tools/list` exposed the expected authoring tools and body/body_cache_id schema.
- Confirmed runtime `tools/call` response for update no-match included `proposal_created:false`, `section_selector_no_match`, and `body_cache`.
- Confirmed runtime `tools/call` response for create stale full-record body included `proposal_created:false`, `invalid_request`, and `body_cache`.
- Ran Design Records MCP validation after close updates for `REQ-MCP-015`, `WORK-MCP-012`, and `TASK-MCP-012-01..03`.
## Evidence
Runtime smoke was executed manually through the actual stdio JSON-RPC MCP server:

```powershell
@'
{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"manual-smoke","version":"0.1.0"}}}
{"jsonrpc":"2.0","method":"notifications/initialized","params":{}}
{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}
{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"propose_record_update","arguments":{"id":"REQ-MCP-015","kind":"requirement","body":"## Dummy\n\nsmoke body","update":{"type":"named_section_replace","section_selector":{"heading":"NO_SUCH_SECTION","match":"exact"}}}}}
{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"propose_record_create","arguments":{"kind":"requirement","id":"REQ-MCP-new","domain":"MCP","title":"smoke full record rejection","fields":{"status":"captured","date":"2026-06-03","source_refs":[],"work_items":[]},"body":"# REQ-MCP-999: stale full record\n\n- **id**: REQ-MCP-999\n- **status**: captured\n- **date**: 2026-06-03\n- **source_refs**:\n- **work_items**:\n\n## Requirement\n\nsmoke body"}}}
'@ | go run ./cmd/design-records-mcp --root .
```

Observed runtime results:

- `tools/list` exposed `propose_record_create` and `propose_record_update` with top-level `body` / `body_cache_id` schema.
- `propose_record_update` named-section no-match with submitted `body` returned `proposal_created:false`, `section_selector_no_match`, and `body_cache` with `body_cache_id`, `expires_at`, and `retention_days`.
- `propose_record_create` `fields + body` with stale full-record body returned `proposal_created:false`, `invalid_request` (`fields plus body create body must omit H1; MCP generates the record heading`), and `body_cache` with `body_cache_id`, `expires_at`, and `retention_days`.

Close evidence:

- `TASK-MCP-012-01` reproduced and classified the current failure behavior.
- `TASK-MCP-012-02` added regression coverage and confirmed no implementation change was required for the REQ-MCP-015 body-cache return behavior.
- Runtime smoke confirmed the behavior through the actual `cmd/design-records-mcp` JSON-RPC entrypoint.

Explicit follow-up boundary:

- `fields + body_cache_id` as a valid retry form for `fields + body` create changes the `propose_record_create` input contract and is intentionally not closed inside REQ-MCP-015 / WORK-MCP-012.
- Legacy full-record create deprecation / fields-required create contract tightening is also outside this close and should be tracked by a separate follow-up requirement.
