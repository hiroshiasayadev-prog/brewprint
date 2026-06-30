# DRMCP-TASK-MCP-001-16: Track parser-aware spec validation implementation

- **id**: DRMCP-TASK-MCP-001-16
- **status**: not_started
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: TBD
- **depends_on**:
  - DRMCP-TASK-MCP-001-09
- **outputs**:
  - DRMCP-WORK-SPEC-001

## Goal

Track the retained W-SPEC-001 per-file detector implementation through reviewed completion.

## Work

- Track `DRMCP-WORK-SPEC-001` through Task graph authoring, implementation, review, and `done`.
- Treat W-SPEC-001 as the retained per-file detector owner.
- Require W-SPEC-001 to consume W012 runtime boundaries.
- Prevent per-file validation ownership from moving into W012 or W010.
- Record the accepted child Work Item evidence pointer.

This Task performs lifecycle tracking only.
This Task does not implement production code, tests, or fixtures.

## Done condition

- W-SPEC-001 has an independently reviewed and released execution graph.
- W-SPEC-001 implements its retained per-file detector boundary.
- W-SPEC-001 consumes W012 runtime boundaries.
- Per-file detector ownership remains outside W012 and W010.
- W-SPEC-001 has completed implementation review with no blocking or major findings.
- W-SPEC-001 is `done`.
- The accepted evidence pointer is recorded here.

## Verification

- Review W-SPEC-001 graph, implementation, test, and independent-review evidence.
- Confirm no W-SPEC-001 detector responsibility was absorbed by W012 or W010.
- Confirm this Task contains no direct implementation evidence beyond the child evidence pointer.

## Evidence

Pending W012 reviewed completion and W-SPEC-001 graph authoring.

The stale T15 release path is not a predecessor.
