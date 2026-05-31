# TASK-DATA-005-03: Create successor split

- **id**: TASK-DATA-005-03
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-DATA-005
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-005-02
- **outputs**:
  - Successor REQ / WORK / TASK split for M15 deferred follow-ups
  - Work item creation recommendations or created artifacts, depending on execution decision

## Goal

Turn the classified deferred items into concrete successor artifact decisions so they no longer remain only in close notes or excluded-scope notes.

## Work

- For each classified item, decide the exact successor artifact action.
- Create or recommend follow-up requirements and work items as appropriate.
- Keep implementation tasks out of this task unless the successor work item explicitly owns them later.
- Ensure successor artifacts separate DATA expressiveness, MCP identity, UC-002 diagnostic debt, and UC-002 cleanup debt.

## Done Condition

- Every deferred item has a concrete successor action.
- Any newly created or recommended artifact has a clear source requirement and boundary.
- No successor split reopens WORK-DATA-002, WORK-DATA-003, or WORK-DATA-004.
- No implementation, fixture, golden, renderer, validator, parser, or MCP tool behavior is changed.

## Verification

- Confirm no item remains only in M15 close notes or excluded-scope notes.
- Confirm successor artifacts do not mix unrelated implementation scopes.
- Confirm the split can be followed without rereading the entire M15 history.

## Evidence

Completed on 2026-06-01.

### Input

TASK-DATA-005-02 classified the deferred M15 follow-up buckets and recommended concrete successor artifacts.

### Created successor requirements

- `REQ-DATA-004`: Tagged union and discriminator payload support.
- `REQ-DATA-005`: DAG asset TypeRef hint render support.
- Existing `REQ-MCP-004` was reused for semantic identity and state machine identity support. No new MCP identity requirement was created.

### Created successor work items

- `WORK-DATA-010`: Implement tagged union and discriminator payload support.
- `WORK-DATA-007`: Implement DAG asset TypeRef hint render support.
- `WORK-DATA-008`: Resolve UC-002 duplicate task QID and unresolved flow task issue.
- `WORK-DATA-009`: Classify remaining UC-002 notes retreat debt.
- Existing `WORK-MCP-004` was reused for semantic identity and state machine identity successor scope. No new MCP identity work item was created.

### Successor split result

| deferred item | successor artifact(s) | source requirement handling |
|---|---|---|
| ADR-073 tagged union / discriminator payload | `REQ-DATA-004`, `WORK-DATA-010` | New DATA requirement created because this is a distinct DATA expressiveness capability. `WORK-DATA-006` was already occupied by helper-shape migration, so the tagged-union successor uses `WORK-DATA-010`. |
| ADR-074 DAG asset TypeRef hint | `REQ-DATA-005`, `WORK-DATA-007` | New DATA requirement created because this is focused DATA render / view support. |
| ADR-078 / ADR-079 / ADR-080 MCP semantic identity / state machine identity | `REQ-MCP-004`, `WORK-MCP-004` | Existing MCP requirement/work item already captured this identity series, so no new identity artifact was created. |
| UC-002 duplicate task QID / unresolved flow task issue | `WORK-DATA-008` | Existing `REQ-DATA-002` used because this is a targeted DATA diagnostic / fixture blocker, not a new feature requirement. |
| Remaining UC-002 notes retreat debt | `WORK-DATA-009` | Existing `REQ-DATA-002` used because this is cleanup planning under the deferred UC-002 helper/model follow-up umbrella. |

### Notes

- New successor work items intentionally have no concrete child tasks yet. Task decomposition is deferred until each successor work item is selected for execution.
- `WORK-DATA-008` and `WORK-DATA-009` point to `REQ-DATA-002`, but `REQ-DATA-002.work_items` synchronization is left to TASK-DATA-005-04.
- `REQ-MCP-007` / `WORK-MCP-007` remain reserved for the separate `list_records` workflow artifact range filter support already captured outside this DATA triage path.
- No existing closed work item was reopened.

### Verification result

- Successor artifact IDs were checked and corrected after collision discovery: `WORK-DATA-006` was already occupied by helper-shape migration, and MCP semantic identity was already captured by `REQ-MCP-004` / `WORK-MCP-004`.
- Successor artifacts separate DATA expressiveness, DATA render/view, MCP identity, UC-002 diagnostic blocker, and UC-002 cleanup planning.
- No implementation, fixture, golden, renderer, validator, parser, MCP tool schema, UC-002 YAML, or render output was changed.
