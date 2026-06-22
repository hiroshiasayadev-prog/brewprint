# Overview: Tools

- **id**: `spec:drmcp.design_records_mcp.tools.overview`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Defines the full MCP tool set for Design Records MCP: read/navigation tools, authoring guidance tools, and authoring transaction tools, along with shared response conventions and error handling.

## Current contract

### Tool set

Read/navigation and guidance tools (P0 unless noted):

| tool | priority | purpose |
|---|---|---|
| `list_records` | P0 | Return structured record index |
| `get_record` | P0 | Retrieve metadata / path / headings / raw body from a single record ID |
| `get_records` | P0 | Retrieve detail representation for multiple explicitly specified record IDs in one batch |
| `validate_records` | P0 | Validate basic record metadata integrity and canonical reference resolution |
| `resolve_reference` | P0 | Resolve a canonical semantic/artifact reference to a document / section / record |
| `list_authoring_guides` | P0 | Return project authoring guide catalog by guide ID |
| `get_authoring_guidance` | P0 | Retrieve authoring guidance Markdown by guide ID |
| `suggest_next_record` | P1 | Suggest the next ADR ID and recommended path |

Authoring transaction tools (MVP, P0). Proposal creation does not modify repository files. Only `accept_proposed_write` may write repository files.

| tool | priority | purpose |
|---|---|---|
| `propose_record_create` | P0 | Create an artifact-oriented write proposal; return diff and validation result |
| `propose_record_update` | P0 | Create a metadata block or named section replacement proposal; return diff and validation result |
| `get_proposed_write` | P0 | Retrieve a retained proposal by proposal ID |
| `accept_proposed_write` | P0 | Accept a proposal and attempt to write to repository files |
| `discard_proposed_write` | P0 | Discard a proposal; prevents future acceptance |

> Source: V01-ADR-077 §P0, §P1, V01-ADR-090 §決定, V01-ADR-093 §決定

### Common response conventions

Tools that return records use separate common fields and kind-specific detail objects. Old flat metadata fields and the kind-specific detail object must not coexist in the same response. After the spec update, the implementation and tests migrate to this response shape as a single cut.

Decision example:

```json
{
  "id": "V01-ADR-076",
  "kind": "decision",
  "title": "Design Records MCP",
  "status": "accepted",
  "path": "v01/records/adr/V01-ADR-076-design-records-mcp.md",
  "decision": {
    "depends_on": ["V01-ADR-050", "V01-ADR-068"],
    "supersedes": [],
    "migrated_to_spec": null
  }
}
```

Investigation example:

```json
{
  "id": "V01-INV-MCP-001",
  "kind": "investigation",
  "title": "Design Records MCP investigation support",
  "status": "concluded",
  "path": "v01/records/investigations/mcp/V01-INV-MCP-001-design-records-mcp-investigation-support.md",
  "investigation": {
    "trigger": "V01-ADR-087",
    "scope": "investigation MCP integration",
    "non_scope": "writer tools",
    "source_refs": ["V01-ADR-086", "V01-ADR-087"],
    "follow_up_candidates": ["V01-ADR-088"]
  }
}
```

Workflow artifact example:

```json
{
  "id": "V01-WORK-MCP-003",
  "kind": "work_item",
  "title": "Workflow artifact MCP support の最小 public contract を判断・実現する",
  "status": "in_progress",
  "path": "v01/records/work-items/mcp/V01-WORK-MCP-003-workflow-artifact-mcp-support.md",
  "work_item": {
    "source_requirement": "V01-REQ-MCP-003",
    "impact_refs": ["V01-ADR-092", "V01-SPEC-design-records-mcp-tools"],
    "tasks": ["V01-TASK-MCP-003-01", "V01-TASK-MCP-003-02", "V01-TASK-MCP-003-03"]
  }
}
```

`title` is extracted from H1. `path` is relative to the repository root (the current working directory or the explicitly specified root path at startup). MVP does not format, summarize, or normalize Markdown body in responses.

Authoring guidance tool responses are not design record responses. Guide source path is not exposed in the public response contract; it is resolved internally from the guide ID.

> Source: V01-ADR-077 §list_records の責務, §get_record の責務

### Error handling

Minimum tool error codes:

| code | meaning |
|---|---|
| `record_not_found` | Single record ID specified in `get_record` does not exist. In `get_records` this is an item-level diagnostic, not a tool error. |
| `guide_not_found` | Guide ID specified in `get_authoring_guidance` does not exist |
| `invalid_request` | Request schema or field value is invalid (e.g. unknown `kind` in `list_records`, missing/empty/non-array `ids` in `get_records`, missing `fields` in `propose_record_create`, `fields.id` mismatch, `new` placeholder with `fields.id`, domain/ID domain mismatch) |
| `unsupported_kind` | Tool given a kind it does not support (e.g. `kind: spec` in `suggest_next_record`) |
| `invalid_id_range` | `id_range` endpoint is malformed, unsupported family, mixed family, mixed domain, mixed task work sequence, or mismatched with specified `kind` |
| `id_range_requires_decision_kind` | Legacy error code from before V01-REQ-MCP-007. Use `invalid_id_range` in new implementations. |
| `proposal_not_found` | Requested proposal ID does not exist |
| `proposal_expired` | Requested proposal ID is past expiry |
| `proposal_discarded` | Proposal is already discarded; cannot be accepted |
| `proposal_already_accepted` | Proposal is already accepted; cannot be re-applied |
| `proposal_stale` | Proposal base state and current target state do not match |
| `target_changed` | Target record kind / path / identity differs from proposal creation time |
| `id_collision` | Create proposal's resolved ID was claimed before acceptance |
| `required_follow_up_not_satisfied` | Required workflow reciprocal metadata or other follow-up is incomplete |
| `invalid_body_source` | Body source rule violation (both `body` and `body_cache_id` supplied, or required source missing) |
| `body_cache_not_found` | Requested body cache ID does not exist |
| `body_cache_expired` | Requested body cache ID is past expiry |
| `proposal_preparation_failed` | Proposal preparation failed before proposal persistence |
| `section_selector_no_match` | Named section selector matched no sections in the target record |
| `section_selector_ambiguous` | Named section selector matched multiple sections; single target not resolvable |

`get_record` returns a machine-readable error for a non-existent record ID:

```json
{
  "error": {
    "code": "record_not_found",
    "message": "record V01-ADR-999 was not found"
  }
}
```

`get_record` `record_not_found` is a tool execution error, not a `validate_records` diagnostic category. `get_records` missing IDs are returned as item-level `record_not_found` diagnostics in the batch response; the batch tool execution itself succeeds.

Authoring diagnostics may appear in normal authoring responses (rather than tool execution errors) when the tool can return `written: false`, proposal state, validation result, or retry guidance. Invalid request shape may still be returned as a tool execution error.

### Authoring write boundary

The authoring transaction MVP intentionally excludes:

- Generic filesystem write tools
- Path-first authoring APIs
- Immediate write tools such as `create_record` or `update_record`
- `set_evidence` convenience operation
- `add_record_metadata` or add/remove relation convenience operations
- `migrate_record_to_spec`
- Partial Markdown AST editing
- General-purpose multi-record atomic transactions with rollback semantics
- Spec skeleton creation and `SPEC-new` placeholder create
- Arbitrary unrelated record bundling in one proposal
- Automatic close cascades across requirement / work item / task
- Automatic rollback after accepted post-write validation failure
- Formatter integration
- Indefinite proposal or body cache retention
- Force-accepting invalid proposals

Workflow artifact create proposals may include required reciprocal metadata updates in the same proposal when needed to keep relation validation valid. That allowance is limited to required reciprocal updates and does not create a general-purpose multi-record atomic transaction.

Existing read / navigation / guidance tools keep their read-only behavior and request/response semantics. `suggest_next_record` remains read-only and does not create files; authoring creates use `propose_record_create`.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| `list_records` | Contract | `spec:drmcp.design_records_mcp.tools.list_records` | Return structured record index with kind/status/id-range/order filters. |
| `get_record` | Contract | `spec:drmcp.design_records_mcp.tools.get_record` | Retrieve metadata, path, headings, and optionally raw body for a single record ID. |
| `get_records` | Contract | `spec:drmcp.design_records_mcp.tools.get_records` | Batch retrieval of multiple explicitly specified record IDs. |
| `list_authoring_guides` | Contract | `spec:drmcp.design_records_mcp.tools.list_authoring_guides` | Return authoring guide catalog with ID, title, and abstract. |
| `get_authoring_guidance` | Contract | `spec:drmcp.design_records_mcp.tools.get_authoring_guidance` | Retrieve full authoring guide Markdown content by guide ID. |
| `resolve_reference` | Contract | `spec:drmcp.design_records_mcp.tools.resolve_reference` | Resolve a canonical spec: or record ID reference to its target. |
| `validate_records` | Contract | `spec:drmcp.design_records_mcp.tools.validate_records` | Validate record metadata integrity and canonical references; returns diagnostic list. |
| `suggest_next_record` | Contract | `spec:drmcp.design_records_mcp.tools.suggest_next_record` | Suggest the next ADR number and recommended filename path. |
| Authoring transaction model | Reference | `spec:drmcp.design_records_mcp.tools.authoring_transaction_model` | Shared concepts for the 5 write tools: proposal lifecycle, body cache, diff_mode, affected record set. |
| `propose_record_create` | Contract | `spec:drmcp.design_records_mcp.tools.propose_record_create` | Create a retained proposal for a new record. |
| `propose_record_update` | Contract | `spec:drmcp.design_records_mcp.tools.propose_record_update` | Create a retained proposal for a metadata or section update. |
| `get_proposed_write` | Contract | `spec:drmcp.design_records_mcp.tools.get_proposed_write` | Retrieve a retained proposal by ID. |
| `accept_proposed_write` | Contract | `spec:drmcp.design_records_mcp.tools.accept_proposed_write` | Accept a proposal and write to repository files. |
| `discard_proposed_write` | Contract | `spec:drmcp.design_records_mcp.tools.discard_proposed_write` | Discard a proposal; prevents future acceptance. |
