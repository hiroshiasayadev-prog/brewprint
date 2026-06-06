# WORK-MCP-023: Fix propose_record_update diff output and no-op detection

- **id**: WORK-MCP-023
- **status**: done
- **date**: 2026-06-07
- **source_requirement**: REQ-MCP-027
- **impact_refs**:
  - SPEC-design-records-mcp-tools
  - ADR-093
  - REQ-MCP-020
- **tasks**:
  - TASK-MCP-023-01
  - TASK-MCP-023-02
  - TASK-MCP-023-03
  - TASK-MCP-023-04

## Goal

Make `propose_record_update` review-safe by returning accurate git-style unified diffs for update proposals and by detecting no-op updates before retaining a write proposal.

The primary failure mode is an existing-file metadata update producing a whole-file addition diff, especially when the requested field value is already unchanged.

## Boundary

In scope:

- Define and implement accurate diff generation for existing-file update proposals.
- Ensure update diffs compare the persisted current file content against the proposed content.
- Emit git-style unified diffs with bounded context for modify proposals.
- Detect no-op updates after applying update operation semantics.
- Prevent retained write proposal creation for no-op update requests.
- Preserve existing accept-time guards and validation behavior.

Out of scope:

- Changing the semantics of `metadata_block_replace`, `metadata_fields_replace`, or `named_section_replace`.
- Automatically accepting update proposals.
- Replacing repository validation or accept-time guards.
- DATA-domain record semantics.
- Exact byte preservation guarantees outside the authoring update path.

## Impact Scope

- `SPEC-design-records-mcp-tools`: document update proposal diff format and no-op response behavior.
- `ADR-093`: confirm whether the authoring transaction model needs a compatibility note or no change.
- `internal/designrecords`: inspect and update authoring proposal diff generation and no-op detection.
- `internal/designrecordsmcp`: expose any no-op response fields or diagnostics through MCP JSON responses if needed.
- Tests: add regression coverage for metadata and named-section update diffs plus no-op behavior.

## Task flow
TASK-MCP-023-01 then TASK-MCP-023-02 then TASK-MCP-023-03 then TASK-MCP-023-04

## Task Candidates

- `TASK-MCP-023-01`: Investigate current update proposal diff and no-op behavior.
- `TASK-MCP-023-02`: Update the MCP tools spec for accurate update diffs and no-op responses.
- `TASK-MCP-023-03`: Implement accurate update diffs, no-op detection, and regression tests.
- `TASK-MCP-023-04`: Run runtime smoke and close synchronization.

## Completion Condition

- Existing-file `propose_record_update` proposals return git-style unified diffs comparing persisted current content to proposed content.
- Metadata-only changes show changed lines plus bounded context, not whole-file addition output.
- Whole-file `+` output remains valid only for actual create/add proposals.
- No-op updates are detected after operation semantics are applied.
- No-op updates do not create retained write proposals.
- No-op responses identify the request as no-op through clear response fields or diagnostics.
- Existing proposal accept guards remain unchanged.
- Tests and runtime smoke evidence are recorded.
- `REQ-MCP-027`, this work item, and all related tasks are synchronized to their final statuses when complete.

## Evidence

Verdict: PASS.

Close evidence is recorded in `TASK-MCP-023-04`.

Completion summary:

- `SPEC-design-records-mcp-tools` was updated by `TASK-MCP-023-02`.
- Accurate update diffs, no-op detection, and regression tests were implemented by `TASK-MCP-023-03`.
- Runtime smoke passed in `TASK-MCP-023-04`.
- `TASK-MCP-023-01`, `TASK-MCP-023-02`, `TASK-MCP-023-03`, and `TASK-MCP-023-04` are synchronized to `done`.
- `REQ-MCP-027` is synchronized to `accepted`.
- `REQ-MCP-029` was not touched.
