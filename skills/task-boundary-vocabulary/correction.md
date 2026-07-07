# Boundary vocabulary: `correction` task_type

A `correction` Task owns repairing findings already identified by a prior review or validator. It does not own `synchronization` (propagating results into a parent Work Item's Evidence/lifecycle), `decision`, `authoring` of new canonical scope, or Work Item closure judgment. This file lists confirmed real and realistic phrasings that describe one of those other responsibilities while embedded inside a `correction`-type Task's Work/Done-condition text.

## canonical: synchronization (owned by `synchronization`, not `correction`)

Pattern: a `correction`-type Task (commonly named `*-finalize`) repairs must-fix findings from a preceding review, then in the same Task also propagates the accepted result into the parent Work Item's Evidence — a distinct synchronization action bundled into the correction Task's own Done condition.

Known paraphrase that IS synchronization, confirmed present (same phrase shape) across five independent real Tasks in the same Work Item:

- "Add [migration/phase] evidence entry to PRODUCT-WORK-SPEC-XXX"

(Source: `PRODUCT-TASK-SPEC-005-04` — "Add BPDSL DSL migration evidence entry to PRODUCT-WORK-SPEC-005"; `PRODUCT-TASK-SPEC-005-08` — "Add mcp/ migration evidence entry to PRODUCT-WORK-SPEC-005"; `PRODUCT-TASK-SPEC-005-12` — "Add views/ migration evidence entry to PRODUCT-WORK-SPEC-005"; `PRODUCT-TASK-SPEC-005-16` — "Carry forward the PRODUCT-TASK-SPEC-005-15 deferred relocation candidates list" plus an evidence-entry bullet; `PRODUCT-TASK-SPEC-005-21` — "Add Phase 2 relocation evidence entry". Spot-checked against `PRODUCT-TASK-SPEC-005-04`'s own `## Work` section verbatim on 2026-07-03.)

Why this matters: a `correction`-type Task that also writes parent-Work-Item Evidence is performing synchronization inside a Task whose declared/inferred responsibility is repair, not lifecycle propagation. In this repo's real corpus this compound shape (`review` → `finalize` = correct + synchronize) is the standard closing pattern for a migration sub-batch, not an isolated anomaly — treat the synchronization half as its own responsibility when checking a `correction` Task's boundary, even though the compound shape itself is an accepted, repeated convention here.

## Open questions / not yet confirmed

- The mirror direction — a `synchronization`-declared Task (e.g. `PRODUCT-TASK-SPEC-004-03`, `accept-and-handoff`) embedding a correction action ("Apply all must-fix findings") — was found once. Only one corpus instance exists so far; not promoted to a confirmed pattern. Needs more instances before promotion.
- No confirmed paraphrase examples yet for `decision`, `authoring`, `coordination`, `investigation`, `review`, or decomposition/execution actions disguised inside a `correction` Task besides the synchronization pattern above.
