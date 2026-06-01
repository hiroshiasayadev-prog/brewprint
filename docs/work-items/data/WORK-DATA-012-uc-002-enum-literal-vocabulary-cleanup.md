# WORK-DATA-012: Plan UC-002 enum, literal, and vocabulary cleanup

- **id**: WORK-DATA-012
- **status**: not_started
- **date**: 2026-06-01
- **source_requirement**: REQ-DATA-002
- **impact_refs**:
  - REQ-DATA-002
  - INV-DATA-002
  - WORK-DATA-009
  - TASK-DATA-009-03
  - TASK-DATA-009-04
- **tasks**:

## Goal

Plan the follow-up for remaining UC-002 enum-like, literal, and usage-site vocabulary cleanup after the M15 enum minimum and helper-shape migration have closed.

This work item owns N-019, N-030, N-034, N-045, N-046, N-051, and residual N-006 / N-015 / N-023 / N-029 vocabulary or literal notes.

## Boundary

### Included

- Review the enum-like / literal constraint bucket selected by `TASK-DATA-009-03`.
- Decide which candidates need named enums, literals, usage-site vocabulary constraints, or explicit no-action outcomes.
- Preserve the completed M15 enum minimum as closed and treat this as follow-up cleanup only.
- Split later task artifacts only when this work item is selected for execution.

### Excluded

- Direct UC-002 YAML migration in this capture.
- Fixture / golden regeneration in this capture.
- Parser, renderer, validator, MCP, or other implementation changes in this capture.
- Reopening `WORK-DATA-001` or broadening the M15 enum minimum.
- Tagged union / discriminator payload support.
- DAG asset TypeRef hint support.
- MCP identity / semantic reference support.
- Numeric/default behavior, selector support matrix, recursive structure, request-side container, or untagged-union successor scope.
- Reopening M15, WORK-DATA-002, WORK-DATA-003, WORK-DATA-004, WORK-DATA-005, WORK-DATA-006, WORK-DATA-007, WORK-DATA-008, WORK-DATA-009, or WORK-DATA-010.

## Impact Scope

| layer | current state | handling in this work item |
|---|---|---|
| source requirement | REQ-DATA-002 captured | Use existing UC-002 helper/model follow-up umbrella |
| completed enum minimum | WORK-DATA-001 done | Preserve as closed; do not reopen |
| source planning | WORK-DATA-009 done | Consume the selected successor bucket only |
| candidate bucket | enum-like / literal constraint | Own the bucket as a future cleanup planning track |

## Task Flow

No task artifacts are created at initial capture time.

Expected later split:

```mermaid
flowchart TD
  T1["Candidate review and vocabulary boundary"]
  T2["Spec / YAML cleanup scope decision"]
  T3["Implementation and fixture tasks if selected later"]
  T4["Verification and close"]
  T1 --> T2 --> T3 --> T4
```

## Completion Condition

This work item can be marked `done` when the enum-like / literal constraint bucket has a concrete accepted cleanup path, implemented and verified if selected, or explicit no-action outcomes for all candidates without reopening completed DATA work.
