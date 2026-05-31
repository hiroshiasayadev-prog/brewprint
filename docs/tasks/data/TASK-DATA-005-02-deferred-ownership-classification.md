# TASK-DATA-005-02: Deferred ownership classification

- **id**: TASK-DATA-005-02
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-DATA-005
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-005-01
- **outputs**:
  - DATA / MCP / cleanup ownership classification
  - Successor artifact recommendation per deferred item

## Goal

Classify each inventoried M15 deferred item into the correct successor ownership bucket without implementing any item.

## Work

- Classify deferred items as DATA expressiveness, MCP identity / public contract, UC-002 diagnostic debt, UC-002 cleanup debt, or no-action / obsolete.
- Decide whether each item needs a new requirement, a new work item under an existing requirement, a task under an existing work item, or no further action.
- Identify cross-domain items that should not remain under DATA.
- Preserve current completed work item boundaries.

## Done Condition

- Each inventory item has a proposed successor ownership and next action.
- DATA and MCP responsibilities are not mixed accidentally.
- No existing closed work item is reopened.
- No implementation, fixture, golden, renderer, validator, parser, or MCP tool behavior is changed.

## Verification

- Confirm each classification maps to a workflow artifact layer boundary.
- Confirm ADR-073 / ADR-074 implementation is not pulled into this task.
- Confirm ADR-078 / ADR-079 / ADR-080 are not treated as DATA-only implementation work unless explicitly justified.

## Evidence

Reviewed on 2026-06-01.

### Input

TASK-DATA-005-01 completed the inventory and confirmed five still-deferred buckets:

- ADR-073 tagged union / discriminator payload.
- ADR-074 DAG asset TypeRef hint.
- ADR-078 / ADR-079 / ADR-080 MCP semantic identity / state machine identity.
- UC-002 duplicate task QID / unresolved flow task issue.
- Remaining UC-002 notes retreat debt.

This task classifies ownership and next action only. It does not create the successor REQ / WORK artifacts.

### Classification

| deferred item | ownership bucket | recommended next action | rationale |
|---|---|---|---|
| ADR-073 tagged union / discriminator payload | DATA expressiveness | Create a new DATA requirement and successor DATA work item. Candidate: `REQ-DATA-004` + next available `WORK-DATA-*`. | ADR-073 is a type/model expressiveness capability, not merely helper-model render scope. It is too large and semantically distinct to remain hidden under `REQ-DATA-002`. |
| ADR-074 DAG asset TypeRef hint | DATA render / view support | Create a new DATA requirement and successor DATA work item. Candidate: `REQ-DATA-005` + `WORK-DATA-007`. | ADR-074 is render/view behavior for existing TypeRef information. It should not be mixed with tagged union implementation or UC-002 cleanup. |
| ADR-078 / ADR-079 / ADR-080 MCP semantic identity / state machine identity | MCP public contract / semantic identity | Use existing `REQ-MCP-004` / `WORK-MCP-004` if confirmed; otherwise create a successor MCP requirement/work item. | The identity series affects MCP object identity, synthetic IDs, state machine identity, transition ID policy, and scenario reference semantics. It is not direct DATA implementation. |
| UC-002 duplicate task QID / unresolved flow task issue | DATA diagnostic / fixture blocker | Create a successor DATA work item under existing `REQ-DATA-002`. Candidate: `WORK-DATA-008`. | This is a pre-existing validation / render blocker found during UC-002 verification, not a new feature requirement. A separate requirement would be excessive unless later investigation reveals a broader contract gap. |
| Remaining UC-002 notes retreat debt | DATA cleanup planning / multi-bucket debt | Create a successor DATA work item under existing `REQ-DATA-002`. Candidate: `WORK-DATA-009`. | The remaining debt is too broad for direct implementation. It should first classify enum-like leftovers, numeric/default behavior, selector matrices, recursive / union cases, request-side containers, and human explanation notes into smaller successors. |

### Recommended successor split

Do not create these in this task; pass them to TASK-DATA-005-03.

- `REQ-DATA-004`: Tagged union and discriminator payload support.
- next available `WORK-DATA-*`: Implement tagged union and discriminator payload support.
- `REQ-DATA-005`: DAG asset TypeRef hint render support.
- `WORK-DATA-007`: Implement DAG asset TypeRef hint render support.
- existing `REQ-MCP-004` / `WORK-MCP-004`, if confirmed: MCP semantic identity and state machine identity support.
- `WORK-DATA-008`: Resolve UC-002 duplicate task QID and unresolved flow task issue.
- `WORK-DATA-009`: Classify remaining UC-002 notes retreat debt.

### Sequencing recommendation

Recommended order for successor execution:

1. Resolve UC-002 duplicate task QID / unresolved flow task issue.
2. Tagged union / discriminator payload support.
3. DAG asset TypeRef hint render support.
4. MCP semantic identity / state machine identity support.
5. Remaining UC-002 notes retreat cleanup planning.

The duplicate task QID / unresolved flow task issue is first because it is a validation / fixture blocker rather than an expressive feature. The remaining items can then be split without confusing a broken UC-002 baseline with future capability work.

### Verification result

- Every TASK-DATA-005-01 inventory item has an ownership bucket and recommended next action.
- DATA and MCP ownership are separated.
- No successor REQ / WORK / TASK artifact was created by this task.
- WORK-DATA-002, WORK-DATA-003, and WORK-DATA-004 were not reopened.
- No implementation, fixture, golden, renderer, validator, parser, MCP tool schema, UC-002 YAML, or render output was changed.
