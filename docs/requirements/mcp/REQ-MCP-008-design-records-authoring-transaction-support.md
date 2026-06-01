# REQ-MCP-008: Design Records MCP authoring transaction support

- **id**: REQ-MCP-008
- **status**: accepted
- **date**: 2026-06-01
- **source_refs**:
  - SPEC-design-records-mcp-tools
  - REQ-MCP-005
  - REQ-MCP-007
- **work_items**:
  - WORK-MCP-008

## Requirement

Design Records MCP needs an authoring transaction capability so AI assistants can create and update design records and workflow artifacts by artifact identity instead of falling back to direct filesystem path editing.

The authoring surface should cover the MVP operations needed for routine dogfooding:

- create requirement, work item, task, and decision records;
- update existing spec records by metadata block or named section replacement when the target record already exists;
- update metadata blocks by full replacement;
- update named Markdown sections by full replacement;
- return previewable diffs before writing;
- accept an explicit proposed write before modifying repository files.

The write contract should be artifact-oriented. Tool inputs should use record kind, ID, domain, section name, and structured authoring fields. Physical file paths may be returned for transparency, but should not be required as primary inputs for authoring operations.

## Evidence

Current Design Records MCP supports record discovery, retrieval, validation, reference resolution, next ADR suggestion, and authoring guidance retrieval, but it does not provide write or update operations for records.

As a result, AI assistants can use Design Records MCP to read and navigate records, but must switch back to filesystem editing when creating requirements, work items, tasks, ADRs, or updating existing spec sections. This creates a split interface:

- read/navigation operations are artifact-oriented;
- write/update operations are path-oriented.

That split weakens dogfooding because assistants are incentivized to inspect and manipulate concrete files even when the user intent is expressed in artifact terms such as `REQ-MCP-new`, `WORK-DATA-010`, or `TASK-DATA-010-01`.

## Required Outcome

A follow-up work item should define and implement a minimal authoring transaction contract with the following expected properties.

### Proposal and accept flow

Authoring writes should use a two-step transaction model:

1. `propose_*` operations prepare the change, return a proposal ID, resolved target ID, diff, validation result, and note, but do not write repository files.
2. `accept_*` operations take a proposal ID and perform the actual write.

The response note for a proposal should make it clear that no file has been written yet and that the caller must accept the proposal ID to apply the diff.

### New ID resolution

Create operations should allow a `new` placeholder in the ID position and resolve it on the MCP side.

Examples:

- `REQ-MCP-new` resolves to the next available `REQ-MCP-NNN`.
- `WORK-DATA-new` resolves to the next available `WORK-DATA-NNN`.
- `ADR-new` resolves to the next available `ADR-NNN`.

Task creation should require enough parent context to avoid ambiguous numbering. One acceptable direction is to require `parent_id` and an ID shape such as `TASK-DATA-010-new`, then resolve the final task sequence under that parent work item.

The `new` placeholder should be valid for create operations only. Update operations should reject `new` IDs.

### Body source and cache recovery

Operations that need large Markdown body input should accept exactly one body source:

- `body`, or
- `body_cache_id`.

Supplying both should be a validation error. Supplying neither should be valid only for operations that can derive the body from a template or structured fields.

If a write or proposal operation receives a large body and fails before the body is persisted into a proposal, the MCP should cache the submitted body and return a `body_cache_id` with retry guidance. The caller can then retry with `body_cache_id` instead of resending or regenerating the full body.

Proposal and body caches should have a retention period of 3 days.

### Validation and repair hints

The authoring flow should distinguish write failure from validation failure.

Responses should explicitly report whether repository files were written:

- `written: false` when no file was modified;
- `written: true` when a file was modified even if subsequent validation failed.

When validation fails, the response should include actionable diagnostics and repair hints, such as using a front matter update operation to fix an invalid field.

### MVP operation shape

The final tool names and schemas should be decided by the work item / spec update, but the MVP should support at least these capabilities:

- propose record create;
- propose record update;
- get proposed write;
- accept proposed write;
- discard proposed write;
- set or replace kind-specific metadata blocks as a whole;
- set or replace a named Markdown section as a whole.

## Explicitly Excluded Scope

This requirement does not require the MVP to support:

- generic filesystem write or edit operations;
- path-first authoring APIs;
- add/remove relation convenience operations;
- partial Markdown AST editing;
- multi-record transactions;
- automatic close cascades across requirement / work item / task;
- spec skeleton creation or `SPEC-new` placeholder create;
- existing spec document structural refactors;
- formatter integration;
- indefinite proposal or body cache retention;
- force-accepting invalid proposals.

## Boundary

This requirement captures the need for Design Records MCP authoring transaction support and the safety properties expected from the MVP. Spec skeleton creation was considered during refinement but is excluded from this MVP because safe spec placement requires separate spec domain tree / placement discovery support.

It does not own the final MCP tool schema, exact response field names, validation diagnostics, implementation details, or the design rationale for choosing the transaction model. Those should be handled by follow-up work item, spec updates, tests, and, if needed, an ADR for the authoring transaction model.

## Close Evidence

2026-06-02: Requirement accepted after WORK-MCP-008 completion.

Implemented by:

- `WORK-MCP-008`

Completion evidence:

- `ADR-093` accepted the Design Records MCP authoring transaction model.
- `SPEC-design-records-mcp-tools` and `SPEC-design-records-mcp-schema` were updated with the public authoring transaction contract.
- `TASK-MCP-008-01` through `TASK-MCP-008-07` are done.
- Authoring transaction tools were implemented and registered:
  - `propose_record_create`
  - `propose_record_update`
  - `get_proposed_write`
  - `accept_proposed_write`
  - `discard_proposed_write`
- Targeted tests passed:
  - `go test ./internal/designrecords ./internal/designrecordsmcp`
  - `go test ./internal/designrecords ./internal/designrecordsmcp -run Authoring`
- Runtime stdio / JSON-RPC smoke passed against a temporary fixture root.
- Real project records were not modified by runtime smoke.
- Design Records validation returned ok for `ADR-093`, `REQ-MCP-008`, `WORK-MCP-008`, and `TASK-MCP-008-07`.

Explicitly deferred:

- Spec skeleton creation / `SPEC-new` support is outside this MVP and captured separately as `REQ-MCP-010`.

