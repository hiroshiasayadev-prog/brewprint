# Task vocabulary reference: graph / lifecycle / execution

General reference of real phrasing used by `synchronization`, lifecycle-closure, and execution/materialization actions in this repository's corpus. Not a boundary-violation list — see `skills/task-boundary-vocabulary/` for that.

Source: `PRODUCT-TASK-SPEC-025-06` through `-10` Finding logs, corpus-range extraction, 2026-07-03.

- `Mark completed tasks done with evidence` / `Add evidence to [Work Item]` — propagate accepted completion into Task lifecycle state and Evidence (`PRODUCT-TASK-SPEC-001-04`)
- `Add [migration/phase] evidence entry to [Work Item]` — propagate a completed batch into parent Work Item Evidence (repeated identical shape across `PRODUCT-TASK-SPEC-005-04/08/12/16/21`; see `skills/task-boundary-vocabulary/correction.md`)
- `Mark [Work Item] status: done` / `update [Work Item] for closure` / `evaluate closing [Work Item]` — close the parent Work Item after accepted artifacts and handoff state are present (`PRODUCT-TASK-SPEC-004-03`, `-007-06`, `-005-21`)
- `WORK-XXX updated` / `Work item closed` / `Work item updated` — synchronize parent Work Item Evidence and lifecycle state (`PRODUCT-TASK-SPEC-009-05`, `-010-04`, `-011-08`)
- `Traceability complete` — persist the complete Task relation list on the parent Work Item (`PRODUCT-TASK-SPEC-010-04`)
- `mechanical reference synchronization` — update refs, parents, Topics rows, and links without semantic redesign (`PRODUCT-TASK-SPEC-012-10`)
- `Synchronize the parent Work Item relations and Evidence` — propagate an accepted contract into parent relations and evidence (`PRODUCT-TASK-SPEC-013-05`)
- `Consolidate accepted [T06/T07] ... evidence` / `Prepare producer handoff evidence for [downstream]` — aggregate and project accepted evidence to a downstream owner (`PRODUCT-TASK-SPEC-013-08`)
- `Update [a coordination Task] with accepted child completion evidence and close [it]` — propagate child completion and close an upstream coordination Task (`PRODUCT-TASK-SPEC-014-02`)
- `Synchronize the hub [T07] lifecycle only when its recorded gate is satisfied` — propagate lifecycle state only after a predefined gate passes (`PRODUCT-TASK-SPEC-015-01`)
- `After review acceptance, close [Work Item] and this Task` — advance both lifecycle records after independent acceptance (`PRODUCT-TASK-SPEC-015-03`)
- `During closure execution, synchronize only this Task and [Work Item]` — update only the exact closure-owned lifecycle records (`PRODUCT-TASK-SPEC-016-11`, `-017-11`)
- `Change [T07] from blocked to not_started` / `Set [Work Item] to in_progress` — apply an accepted decision to Task or Work Item lifecycle state (`PRODUCT-TASK-SPEC-018-13`, `-019-22`)
- `Record [T19] as the accepted finding-closure review` — propagate accepted review authority into closure Evidence (`PRODUCT-TASK-SPEC-018-08`)
- `Treat [a prerequisite gap] as a mechanically necessary missing-owner repair` — classify a prerequisite gap as a graph repair rather than a review finding (`PRODUCT-TASK-SPEC-018-09`)
- `Accept [T18] directly when its verdict is PASS` — use a review verdict as closure authority without a new review verdict (`PRODUCT-TASK-SPEC-019-19`)
- `Record only coarse downstream routing in [Work Item]` — update a parent Work Item with child identities and routing only (`PRODUCT-TASK-SPEC-019-10`)
- `Record the external [gate] for [later writers/review]` — persist a cross-Work-Item dependency gate for later consumers (`PRODUCT-TASK-SPEC-019-11`)
- `Materialize the [checklist/decision/coordination] contract` — create the required Tasks, dependencies, gates, and release order (`PRODUCT-TASK-SPEC-020-01`, `-021-01/07/09`, `-023-03/05/07`)
- `Create the independent [child Work Item] selected by [T06]` — materialize one decided child Work Item boundary without its internal Task graph (`PRODUCT-TASK-SPEC-021-13`; a `work_item_decomposition` example)
- `Propagate the accepted [review/handoff] result into lifecycle, Evidence, relation, and closure state` — mechanically update Work Item state from accepted prior results (`PRODUCT-TASK-SPEC-021-14`)
- `Record verification Evidence and directly close [Work Item] when every Done condition is satisfied` — mechanically close a Work Item inside an authoring Task under the accepted simple-closure exception (`PRODUCT-TASK-SPEC-024-02`)
