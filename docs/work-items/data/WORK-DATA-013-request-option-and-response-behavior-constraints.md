# WORK-DATA-013: Define request option and response behavior constraints

- **id**: WORK-DATA-013
- **status**: blocked
- **date**: 2026-06-01
- **source_requirement**: REQ-DATA-006
- **impact_refs**:
  - REQ-DATA-006
  - REQ-DATA-002
  - INV-DATA-002
  - WORK-DATA-009
  - TASK-DATA-009-03
  - TASK-DATA-009-04
- **tasks**:
  - TASK-DATA-013-01
  - TASK-DATA-013-02
  - TASK-DATA-013-03
  - TASK-DATA-013-04

## Goal

Define the follow-up path for request option and response behavior constraints captured by `REQ-DATA-006`.

This work item owns the `numeric / default behavior` bucket: N-011, N-017, N-022, N-024, N-025, and N-028.

## Boundary

### Included

- Decide the contract boundary for numeric ranges, defaults, omitted values, unknown values, fallback behavior, and cross-response grouping.
- Identify required spec, diagnostic, YAML, and fixture follow-up before any implementation work begins.
- Split later task artifacts only when this work item is selected for execution.

### Excluded

- Direct UC-002 YAML migration in this capture.
- Fixture / golden regeneration in this capture.
- Parser, renderer, validator, MCP, or other implementation changes in this capture.
- Tagged union / discriminator payload support.
- DAG asset TypeRef hint support.
- MCP identity / semantic reference support.
- Request-side generic container cleanup, enum/literal cleanup, selector support matrix, recursive structure, or untagged-union successor scope.
- Reopening M15, WORK-DATA-001, WORK-DATA-002, WORK-DATA-003, WORK-DATA-004, WORK-DATA-005, WORK-DATA-006, WORK-DATA-007, WORK-DATA-008, WORK-DATA-009, or WORK-DATA-010.

## Impact Scope

| layer | current state | handling in this work item |
|---|---|---|
| source requirement | REQ-DATA-006 captured | Owns request option and response behavior constraints |
| source planning | WORK-DATA-009 done | Consume the selected successor bucket only |
| candidate bucket | numeric / default behavior | Own the bucket as a future contract decision track |

## Task Flow

No task artifacts are created at initial capture time.

Expected later split:

```mermaid
flowchart TD
  T1["Contract boundary decision"]
  T2["Spec and diagnostic alignment"]
  T3["Implementation and fixture tasks if selected later"]
  T4["Verification and close"]
  T1 --> T2 --> T3 --> T4
```

## Completion Condition

This work item can be marked `done` when request option and response behavior constraints are accepted, specified, implemented and verified if selected, or explicitly closed as no-action without mixing in unrelated DATA or MCP identity work.
