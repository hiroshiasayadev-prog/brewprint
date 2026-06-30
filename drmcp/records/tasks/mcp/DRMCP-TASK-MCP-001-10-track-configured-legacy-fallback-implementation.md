# DRMCP-TASK-MCP-001-10: Track configured legacy-fallback implementation

- **id**: DRMCP-TASK-MCP-001-10
- **status**: not_started
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d coordination
- **depends_on**:
  - DRMCP-TASK-MCP-001-02
  - DRMCP-TASK-MCP-001-05
  - DRMCP-TASK-MCP-001-08
  - DRMCP-TASK-MCP-001-09
- **outputs**:
  - DRMCP-WORK-MCP-010

## Goal

Track the rebaselined configured legacy archive fallback implementation through reviewed completion.

## Work

- Track `DRMCP-WORK-MCP-010` as the configured legacy-fallback implementation gate.
- Require W010 to consume the completed W012 current-runtime output.
- Require a new W010 execution-graph authoring, independent review, and release sequence after W012 completion.
- Reject the retired W009 extension seam as implementation authority.
- Treat existing W010 Task Candidates as planning only, not released implementation contracts.
- Track W010 through rebaseline, implementation review, and `done`.
- Record the accepted W010 evidence pointer here.

This Task performs lifecycle tracking only.
This Task does not implement production code, tests, or fixtures.

## Done condition

- W012 is reviewed and `done`.
- W010 has an independently reviewed and released replacement execution graph.
- No W010 production Task starts before rebaseline and release.
- W010 implements the accepted optional configured legacy-fallback scope.
- W010 public behavior and fixture contracts remain unchanged.
- W010 has completed independent implementation review with no blocking or major findings.
- W010 is `done`.
- The accepted W010 evidence pointer is recorded here.

## Verification

- Review W010 rebaseline, implementation, test, leakage, and independent-review evidence.
- Confirm the retired W009 extension seam is not used as authority.
- Confirm W010 consumes the completed W012 runtime.
- Confirm this Task contains no direct implementation evidence beyond the W010 evidence pointer.

## Evidence

Selected child Work Item: `DRMCP-WORK-MCP-010`.

- W010 remains `blocked` pending W012 completion and execution-graph rebaseline.
- Existing Task Candidates are not released implementation contracts.
- No W010 production Task may start before rebaseline.
- T10 remains `not_started` until T09 completes and the rebaselined W010 workflow begins.
- The stale T15 release path is not a predecessor.
