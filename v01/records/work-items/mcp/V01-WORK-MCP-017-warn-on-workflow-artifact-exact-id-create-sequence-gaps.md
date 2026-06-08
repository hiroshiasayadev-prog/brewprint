# V01-WORK-MCP-017: Warn on workflow artifact exact ID create sequence gaps

- **id**: V01-WORK-MCP-017
- **status**: done
- **date**: 2026-06-03
- **source_requirement**: V01-REQ-MCP-018
- **impact_refs**:
  - SPEC-design-records-mcp-tools
- **tasks**:
  - V01-TASK-MCP-017-01
  - V01-TASK-MCP-017-02
  - V01-TASK-MCP-017-03
  - V01-TASK-MCP-017-04

## Goal

Add authoring feedback for workflow artifact exact ID create requests that may create sequence gaps.

The MCP should continue to allow exact ID create, but it should surface a warning when the requested exact ID appears to skip an available sequence in the relevant workflow artifact family.

## Boundary

This work item owns proposal-time diagnostics / notes for `propose_record_create` exact ID create gap risk.

In scope:

- Exact ID create warning behavior for workflow artifacts
- Sequence-gap detection rules for domain-scoped WORK / REQ families and parent-scoped TASK families where applicable
- Spec and authoring guidance updates for exact ID vs `*-new` placeholder usage
- Regression tests and runtime smoke evidence

Out of scope:

- Rejecting exact ID create solely because it creates a gap
- Auto-filling or repairing existing gaps
- Renaming or migrating already-created workflow artifacts
- Forcing REQ and WORK sequence numbers to match
- Changing server-side `*-new` allocation behavior

## Impact Scope

- `SPEC-design-records-mcp-tools`
- authoring create validation / proposal note generation
- authoring guidance for workflow artifact creation
- regression tests for `propose_record_create`

## Task flow

1. Define and document exact ID gap-warning rules.
2. Update spec / guidance to describe the warning and preferred `*-new` usage.
3. Implement proposal warning diagnostics / notes with regression tests.
4. Run runtime smoke and close the requirement/work item.

## Completion Condition
`propose_record_create` returns a non-blocking warning when a workflow artifact exact ID create may create a sequence gap, while still allowing acceptance of the proposal.

## Evidence

V01-WORK-MCP-017 is closed on 2026-06-03. V01-TASK-MCP-017-01 through V01-TASK-MCP-017-04 are complete, tests passed, and V01-TASK-MCP-017-04 records runtime smoke evidence for exact ID gap warning behavior.

## Progress Boundary

Do not reopen V01-WORK-MCP-014. Treat this as a follow-up authoring feedback improvement for V01-REQ-MCP-018.
