# PRODUCT-TASK-SPEC-022-01: Decide work_item_execution Task contract

- **id**: PRODUCT-TASK-SPEC-022-01
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-022
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - PRODUCT-TASK-SPEC-022-01

## Goal

Fix one bounded contract for representing one already-created child Work Item as one parent-graph execution unit.

## Work

- Decide the Task type name.
- Decide the primary outcome and completion judgment.
- Decide the Task-to-child-Work-Item relation.
- Decide relation cardinality and target constraints.
- Decide child-status effects on Task completion.
- Decide the parent hub completion boundary.
- Decide adjacent Task-type boundaries.
- Decide ADR disposition, canonical targets, and migration.
- Record rejected options and concise reasons.

This Task must not:

- author ADR, Specification, workflow-support, or checklist content;
- change the Task graph after this simplified route is fixed;
- perform review, correction, synchronization, implementation, stage, or commit work.

### Decision ledger

| ID | topic | status | decision | reason |
|---|---|---|---|---|
| D-001 | Task type name | `decided` | Use `work_item_execution`. | The name states that the Task represents child Work Item execution rather than Work Item creation or split. |
| D-002 | Primary outcome | `decided` | One already-created child Work Item is represented as one execution unit in the parent Work Item Task graph. | The parent graph needs one Task-sized unit without duplicating child internals. |
| D-003 | Completion judgment | `decided` | The Task may become `done` only after the referenced child Work Item is `done` and Task Evidence records that status. | Child completion is the only completion boundary owned by this Task. |
| D-004 | Canonical relation | `decided` | Add the scalar Task metadata field `work_item_ref`. | `work_item`, `depends_on`, and `outputs` express different relations and must not be overloaded. |
| D-005 | Relation constraints | `decided` | `work_item_ref` is required only for `work_item_execution`, names exactly one existing Work Item, differs from the parent `work_item`, and uses the same app and domain namespace. | The relation must be unambiguous and must not replace Task ownership by the parent Work Item. |
| D-006 | Reverse relation | `decided` | Add no reverse execution field to the child Work Item. | The relation is discoverable from Tasks and a second writable relation would create duplicate state. |
| D-007 | Child status semantics | `decided` | `done` permits Task completion. `blocked` permits the Task to be blocked. `not_started` and `in_progress` do not satisfy completion. | These are the complete current Work Item lifecycle states. |
| D-008 | Canceled status | `decided` | Add no `canceled` status. | Current Task and Work Item lifecycle authority does not define it, and lifecycle expansion is outside this additive type change. |
| D-009 | Parent hub completion | `decided` | A completed execution Task satisfies only its own Done condition. The parent Work Item still evaluates its full Completion Condition. | Child completion must not automatically close the parent hub. |
| D-010 | Decomposition boundary | `decided` | `work_item_decomposition` creates or splits child Work Items. `work_item_execution` references one already-created child and waits for its completion. | The two types own different outcomes and completion judgments. |
| D-011 | Coordination boundary | `decided` | `coordination` may create or route a `work_item_execution` Task but does not execute child work or own child completion. | Graph change and child completion remain separate responsibilities. |
| D-012 | Synchronization boundary | `decided` | `synchronization` propagates accepted state. It does not own ongoing child execution or select the child relation. | Ongoing execution and final mechanical propagation have different completion judgments. |
| D-013 | Child ownership | `decided` | The execution Task does not duplicate child Tasks, deliverables, procedures, decisions, or review Evidence. | The child Work Item remains the sole owner of its internal graph and work. |
| D-014 | Applicability | `decided` | The type is a general PRODUCT Task type, not a design-convergence-only type. | Hub-to-child execution is a general workflow relation. |
| D-015 | Migration | `decided` | Perform no existing-record migration. | The type is additive and current records need no reinterpretation. |
| D-016 | ADR disposition | `decided` | Amend ADR-004, ADR-005, and ADR-010 as a non-material responsibility refinement. Create no new ADR and supersede none. | The closed typed-responsibility architecture and its rationale remain valid. |
| D-017 | Canonical targets | `decided` | Update Task authoring, Work Item authoring, design-convergence workflow guidance, its activation pointer, and the derived W020 checklist set. | These are the exact authorities and consumers affected by the new type. |

### Rejected options

| option | reason |
|---|---|
| Reuse `work_item_decomposition`. | Creation or split completes before child execution completes. |
| Use `outputs` as the child relation. | `outputs` does not identify the single execution target or its completion semantics. |
| Use `depends_on` as the child relation. | `depends_on` accepts Task IDs, not Work Item IDs. |
| Copy the child Task graph into the parent Task. | Duplicate ownership would become stale and violate child Work Item boundaries. |
| Automatically close the parent Work Item when the child is done. | The parent may own additional Tasks and Completion Conditions. |

## Done condition

- D-001 through D-017 are `decided`.
- The type has one primary outcome and one completion judgment.
- The one-to-one relation and all status effects are explicit.
- Adjacent Task-type boundaries are non-overlapping.
- ADR, canonical-target, and migration dispositions are fixed.
- No canonical artifact is authored.

## Verification

- Confirm every decision item is terminal.
- Confirm the selected relation does not replace the parent `work_item` relation.
- Confirm the contract uses only current lifecycle states.
- Confirm child internals remain child-owned.
- Confirm the amendment route follows the accepted non-material responsibility-refinement rule.
- Confirm no authoring, review, correction, synchronization, implementation, stage, or commit work occurred.

## Evidence

- User decisions on 2026-07-02 fixed the name, dedicated scalar relation, one-child cardinality, and child completion boundary.
- `PRODUCT-TASK-SPEC-018-16` and `PRODUCT-ADR-SPEC-006` permit in-place ADR amendment for responsibility extraction inside an unchanged architecture.
- `PRODUCT-INV-SPEC-006`, `PRODUCT-INV-SPEC-008`, W020 T04 blocker Evidence, ADR-004, ADR-005, ADR-010, Task authoring, and Work Item authoring supplied the bounded impact evidence.
- Result: `PASS`.
