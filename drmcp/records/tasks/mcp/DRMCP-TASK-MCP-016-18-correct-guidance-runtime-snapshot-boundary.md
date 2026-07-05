# DRMCP-TASK-MCP-016-18: Correct Guidance runtime snapshot boundary

- **id**: DRMCP-TASK-MCP-016-18
- **status**: done
- **date**: 2026-07-05
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: correction
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-016-17
- **outputs**:
  - spec:drmcp.design_records_mcp.namespace_scanning
  - DRMCP-TASK-MCP-016-18

## Goal

Repair T16 finding F-MAJ-02 by aligning the namespace-scanning runtime snapshot boundary with current Guidance lifecycle authority.

## Work

- Update only `spec:drmcp.design_records_mcp.namespace_scanning` and this Task record.
- Include `list_authoring_guides` and `get_authoring_guidance` in normal fresh Current Records snapshot construction.
- Preserve operation-specific Legacy state loading.
- Remove the stale four-operation W011 runtime boundary.
- Remove the statement that excludes authoring-guidance runtime architecture.
- Replace stale lifecycle authority references with current application-architecture authority.
- Preserve authoring-transaction runtime architecture as excluded scope.
- Avoid unrelated wording, source, identity, or legacy behavior changes.

## Done condition

- `## Runtime snapshot boundary` covers Read, Validation, and Guidance requests.
- Guidance uses one fresh immutable Current Records snapshot per invocation.
- Legacy state remains operation-specific.
- The section no longer limits lifecycle behavior to four W011 operations.
- The section no longer places authoring guidance outside the runtime boundary.
- Current lifecycle and dependency Specifications are referenced as authority.

## Verification

- Inspect the complete scoped diff for this Task and `namespace-scanning.md`.
- Confirm no other Specification changed.
- Confirm `list_authoring_guides` and `get_authoring_guidance` are named in the runtime lifecycle.
- Confirm authoring transactions remain outside this Specification boundary.
- Confirm no new design judgment or finding disposition is recorded.
- Confirm scoped whitespace inspection passes.

## Evidence

- The repository-root `prompt_chappy.md` was read first.
- `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational. Design Record authoring used filesystem fallback.
- T16 F-MAJ-02 supplied the exact finding and required outcome.
- The correction changes only the stale runtime snapshot section and its direct authority references.
- `list_authoring_guides` and `get_authoring_guidance` now use the same fresh immutable Current Records snapshot lifecycle as Read and Validation requests.
- Legacy lookup state remains loaded only when the operation requires legacy compatibility.
- Authoring-transaction runtime architecture remains outside this boundary.
- Limited stale-text search found no remaining four-operation W011 wording or Guidance-outside-boundary statement in `namespace-scanning.md`.
- The scoped T18 diff returned 16,270 of 16,270 bytes with no truncation. The namespace file also contains earlier uncommitted W016 changes, so current full text and the exact corrected section were reviewed directly.
- Scoped whitespace inspection passed. LF-to-CRLF notices were advisory only.
- No scoped file is staged.
- The standalone semantic responsibility validator was not executed because no operational invocation tool is available. No validator PASS was synthesized.
- F-MAJ-02 was repaired but not closed by this Task.
- No stage or commit was performed.
