# DRMCP-TASK-MCP-026-02: Decide application-architecture relocation topology

- **id**: DRMCP-TASK-MCP-026-02
- **status**: done
- **date**: 2026-07-08
- **work_item**: DRMCP-WORK-MCP-026
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-026-01
- **outputs**:
  - DRMCP-TASK-MCP-026-02

## Goal

Decide the relocation target, compatibility policy, ADR route, and migration execution route for the application-architecture Specification tree.

## Work

Use T01 as the reference inventory.
Record only relocation-topology decisions.
Do not author Specifications or move files in this Task.

Decision ledger:

| id | topic | status | decision summary | reason | downstream owner |
|---|---|---|---|---|---|
| D-001 | Target subtree | decided | Relocate the tree to `spec:drmcp.implementation.application_architecture`. | The tree is implementation-facing architecture authority. The selected ref preserves the existing concept name under the implementation topology. | Migration authoring route. |
| D-002 | Physical target path | decided | Use `drmcp/records/spec/implementation/application-architecture/`. | Path-derived Spec refs map the hyphenated path segment to `application_architecture`. | Migration script. |
| D-003 | `spec:drmcp.implementation` root role | deferred | Do not redesign or broaden the implementation root now. | The relocation can proceed without reworking the current W011-root content. Any root-role cleanup is separate follow-up work. | None in this Work Item unless validation blocks relocation. |
| D-004 | Compatibility policy for old refs | decided | Leave no compatibility stub under `spec:drmcp.application_architecture`. | The old placement reflects missing user intent, not a compatibility contract. Historical workflow records are not rewritten solely to hide the old placement. | Migration authoring route. |
| D-005 | ADR route | decided | Do not create an ADR for this relocation. | The relocation records intended Specification topology. It does not select a new durable architecture design choice. | None. |
| D-006 | Migration execution route | decided | Author a deterministic migration script, review the script and dry-run output, then apply the script after review. | File moves and ref rewrites are coupled. A script makes the mapping reviewable and reduces manual edit risk. | Script implementation, review, and migration application Tasks. |

## Done condition

- Target subtree is decided.
- Compatibility policy is decided.
- ADR route is decided.
- Migration execution route is decided.
- Implementation-root role is either decided or explicitly deferred.
- No Specification file move or canonical authoring is performed.

## Verification

- D-001 through D-006 are terminal.
- T01 inventory remains the source reference set.
- No ADR was created.
- No Specification was edited.
- No file move was performed.

## Evidence

- T01 found no physical-path references under `drmcp/records/spec/`.
- T01 found canonical refs that must be updated after the target subtree is decided.
- User decision on 2026-07-08 selected `spec:drmcp.implementation.application_architecture`.
- User decision on 2026-07-08 deferred the current `spec:drmcp.implementation` root-role cleanup.
- User decision on 2026-07-08 selected no compatibility stub for old refs.
- User decision on 2026-07-08 rejected an ADR for this relocation.
- User question on 2026-07-08 selected the migration-script review route as the preferred execution shape.
