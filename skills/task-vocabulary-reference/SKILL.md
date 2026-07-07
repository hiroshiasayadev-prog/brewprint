# Task vocabulary reference

General reference of real, corpus-observed phrasing used by each task_type cluster in this repository's Task records. This documents normal vocabulary breadth — how each responsibility is actually worded in practice — independent of whether a phrase crosses a type boundary.

This is a separate purpose from `skills/task-boundary-vocabulary/`, which stays narrowly scoped to phrases that disguise a *different* task_type's responsibility (a much smaller, stricter list). Use this skill when you want to know how a responsibility is typically phrased; use `task-boundary-vocabulary` when you're checking whether a phrase might be smuggling in a responsibility outside the Task's declared type.

## Files

- `decision-authoring-coordination.md`
- `review-verification-correction.md`
- `investigation-decision-design.md`
- `graph-lifecycle-execution.md`

## Source and scope

Built from `PRODUCT-TASK-SPEC-025-06` through `-10`'s raw Finding logs (219 entries, corpus-range extraction over `product/records/tasks/spec/`, cutoff 2026-07-03), grouped by the cluster labels those Tasks already used. Entries are deduplicated: identical or near-identical phrasing repeated across multiple source Tasks is listed once with representative citations, not once per occurrence. This intentionally loses frequency-ranking precision in exchange for readability; if exact per-instance counts matter, consult the source Tasks' Finding logs directly.

This reference does not decide canonical terminology and is not a boundary-violation dictionary.
