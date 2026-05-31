# TASK-DATA-005-01: M15 deferred item inventory

- **id**: TASK-DATA-005-01
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-DATA-005
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - M15 deferred item inventory
  - Source note map for successor split

## Goal

Inventory the M15 deferred follow-up items that are still not owned by a successor work item after WORK-DATA-002, WORK-DATA-003, and WORK-DATA-004.

## Work

- Review only the relevant close notes, excluded-scope notes, and task evidence needed to identify deferred items.
- Confirm the current known state of ADR-073, ADR-074, ADR-078, ADR-079, ADR-080, UC-002 duplicate task QID / unresolved flow task issue, and remaining UC-002 notes retreat debt.
- Record which source artifact currently mentions each item.
- Avoid reopening completed DATA work.

## Done Condition

- The deferred item list is complete enough to classify into successor ownership.
- Each item has at least one source note or artifact reference.
- No implementation, fixture, golden, renderer, validator, parser, or MCP tool behavior is changed.

## Verification

- Confirm WORK-DATA-002 and WORK-DATA-003 remain closed.
- Confirm WORK-DATA-004 scope is not changed.
- Confirm this task only inventories and does not decide implementation scope.

## Evidence

Reviewed on 2026-06-01.

### Sources reviewed

- `docs/tasks/m15-data-layer-expressiveness.md` close note: historical M15 close boundary and explicit deferred-outside list.
- `WORK-DATA-001`: new-format work item that owns the actual M15 minimum-expressiveness close boundary.
- `REQ-DATA-002`: helper model / model render follow-up requirement and explicit exclusions.
- `WORK-DATA-002`: task-file helper model minimum close evidence and deferred model-file / UC-002 migration scope.
- `WORK-DATA-003`: model-file helper render boundary close evidence and UC-002 candidate classification boundary.
- `WORK-DATA-004`: task-file private helper signature exposure policy close evidence and remaining excluded debt.
- `TASK-DATA-003-04`: UC-002 model response helper candidate review and unchanged-candidate classification.
- `INV-DATA-002`: UC-002 notes retreat inventory grouped patterns and later-boundary inputs.
- Design Records MCP metadata for `ADR-073`, `ADR-074`, `ADR-078`, `ADR-079`, `ADR-080`, and `INV-DATA-002`.

### Inventory

| deferred item | source notes | current ownership state | inventory result |
|---|---|---|---|
| ADR-073 tagged union / discriminator payload | M15 close note lists ADR-073 outside close. `REQ-DATA-002` excludes tagged union. `WORK-DATA-002`, `WORK-DATA-003`, and `WORK-DATA-004` all exclude ADR-073 implementation. `TASK-DATA-003-04` marks N-003, N-021, N-026, N-027 and related N-001 as tagged-union / kind-specific payload debt. `INV-DATA-002` groups N-001, N-003, N-021, N-026, N-027 as tagged union candidates. | ADR exists as `ADR-073` with status `proposed`; no dedicated successor work item found in the DATA chain. | Still deferred. Needs successor classification in TASK-DATA-005-02. Likely DATA expressiveness successor, but this task does not decide ownership. |
| ADR-074 DAG asset TypeRef hint | M15 close note lists ADR-074 outside close. `WORK-DATA-001` explicitly excludes DAG asset node label TypeRef hint while allowing enum to exist as machine-readable named model. `REQ-DATA-002`, `WORK-DATA-002`, `WORK-DATA-003`, and `WORK-DATA-004` exclude ADR-074. | ADR exists as `ADR-074` with status `proposed`; no dedicated successor work item found in the DATA chain. | Still deferred. Needs successor classification in TASK-DATA-005-02. Likely DATA render/view successor. |
| ADR-078 / ADR-079 / ADR-080 MCP semantic identity / state machine identity | M15 close note lists ADR-078〜080 outside close. `WORK-DATA-001` excludes MCP semantic identity / state machine identity. `REQ-DATA-002`, `WORK-DATA-002`, `WORK-DATA-003`, and `WORK-DATA-004` exclude MCP identity work. `TASK-DATA-003-04` classifies N-032, N-035, N-038, N-039, N-043, N-047, N-048, and N-050 as MCP identity / semantic reference debt. `INV-DATA-002` groups FileID / path-normalized identity, QualifiedID / synthetic ID / file-local ID, and semantic reference semantics as closer to MCP semantic identity / state-machine identity than pure data-layer expressiveness. | ADR-078 is accepted; ADR-079 and ADR-080 are proposed. No dedicated successor work item found in the DATA chain. | Still deferred. Needs successor classification in TASK-DATA-005-02. Likely MCP-domain successor, not direct DATA implementation. |
| UC-002 duplicate task QID / unresolved flow task issue | M15 close note lists the duplicate task QID / unresolved flow task issue outside close. `WORK-DATA-001` impact scope records UC-002 full validate / render as blocked by a pre-existing duplicate task QID issue and outside the work item boundary. `WORK-DATA-001` close outcome repeats that UC-002 full validate / render still has this pre-existing issue outside enum migration and M15 close. | No dedicated successor work item found in the DATA chain. | Still deferred. Needs successor classification in TASK-DATA-005-02. Likely targeted DATA diagnostic / fixture follow-up. |
| Remaining UC-002 notes retreat debt | M15 close note lists remaining UC-002 notes retreat debt outside close. `WORK-DATA-001` excludes UC-002 notes retreat full resolution. `REQ-DATA-002` captures only helper/model render follow-up and excludes tagged union, DAG TypeRef hint, MCP identity, and M15 reopening. `WORK-DATA-002` excludes treating all UC-002 notes retreat debt as one required migration. `WORK-DATA-003` classifies candidates but delegates actual UC-002 YAML migration and excludes tagged union, DAG TypeRef hint, MCP identity, and actual migration. `WORK-DATA-004` excludes UC-002 model response helper-shape migration and remaining notes retreat debt. `TASK-DATA-003-04` leaves enum/value/default/range/support-matrix, recursive/union, MCP identity, request-side, and human explanation/view-renderer notes unchanged. `INV-DATA-002` records remaining patterns including numeric range/default behavior, selector matrices, recursive ObjectRef, usage-site-dependent vocabularies, and untagged union lists as not cleanly covered by existing ADR candidates. | Partially reduced by ADR-067 enum minimum, WORK-DATA-002 task-file helper minimum, WORK-DATA-003 model-file render support, and WORK-DATA-004 signature policy. Full cleanup still has no dedicated successor work item. | Still deferred. Needs successor classification in TASK-DATA-005-02. Should likely be split rather than handled as one broad cleanup item. |

### Source note map

- M15 historical close note is the canonical legacy source that names the five deferred buckets.
- `WORK-DATA-001` is the canonical new-format source for the actual M15 close boundary and confirms the same deferred buckets.
- `REQ-DATA-002` owns the helper/model render follow-up umbrella but explicitly excludes ADR-073, ADR-074, ADR-078〜080, M15 reopening, and completed F1 reclassification.
- `WORK-DATA-002` completed task-file helper minimum and deferred model-file render / UC-002 migration while excluding ADR-073, ADR-074, ADR-078〜080, and full notes retreat cleanup.
- `WORK-DATA-003` completed model-file helper render exposure and candidate classification, then delegated actual UC-002 migration and excluded tagged union, DAG TypeRef hint, MCP identity, and M15 reopening.
- `WORK-DATA-004` completed task-file private helper signature exposure policy and still excluded UC-002 migration plus remaining notes retreat debt.
- `TASK-DATA-003-04` is the most concrete source for UC-002 candidate IDs and distinguishes model-file helper migration candidates, task-file helper wait candidates, unchanged tagged-union candidates, MCP identity debt, and broader notes-retreat debt.
- `INV-DATA-002` remains the broad inventory source for note-retreat patterns and explicitly says no ADR adoption, M15 close boundary, implementation order, or new artifact proposal was made by that investigation.

### Verification result

- WORK-DATA-002, WORK-DATA-003, and WORK-DATA-004 were treated as closed and not reopened.
- This task only inventories deferred items and source notes.
- No implementation, fixture, golden, renderer, validator, parser, MCP tool schema, UC-002 YAML, or render output was changed.
- Classification and successor ownership are intentionally left to TASK-DATA-005-02.
