# Task vocabulary reference: decision / authoring / coordination

General reference of real phrasing used by `decision`, `authoring`, and `coordination` Tasks in this repository's corpus. This is not a boundary-violation list — every phrase here is normal, in-type vocabulary for the responsibility it expresses. See `skills/task-boundary-vocabulary/` for the separate, narrower list of phrases that disguise a *different* task_type's responsibility.

Source: `PRODUCT-TASK-SPEC-025-06` through `-10` Finding logs, corpus-range extraction, 2026-07-03.

## decision

- `Define` — encode an accepted rule or boundary (`PRODUCT-TASK-SPEC-001-01`, `PRODUCT-TASK-SPEC-013-02`)
- `Combine` — merge two prior classifications into one proposed choice (`PRODUCT-TASK-SPEC-004-01`)
- `Recommend whether` — propose a scope choice without applying it (`PRODUCT-TASK-SPEC-004-01`)
- `Select` — choose the accepted option among named candidates (`PRODUCT-TASK-SPEC-013-04`)
- `Fix` — decide and freeze a target, boundary, or policy without performing it (12 corroborating instances: `PRODUCT-TASK-SPEC-013-04`, `-018-11`, `-018-16`, `-019-03`, `-019-07` ×5, `-019-09` ×3, `-022-01`, `-023-01`, `-024-01`; see also `skills/task-boundary-vocabulary/decision.md`)
- `Persist each explicit answer before advancing [the cursor]` — decision-ledger recording (`PRODUCT-TASK-SPEC-016-02`, `-017-02`, `-019-01`)
- `Rule set confirmed` / `Scope decisions resolved` — approve or settle named choices (`PRODUCT-TASK-SPEC-010-01`)
- `State every required graph change without applying it` — record a decision's downstream implication while leaving execution to coordination (`PRODUCT-TASK-SPEC-019-07`)
- `Classify each item as no ADR, new ADR, amendment, or supersession` — ADR-routing decision (`PRODUCT-TASK-SPEC-016-05`, `-017-05`, `-021-08`, `-023-06`)
- `Group only decisions that form one coherent durable choice` — ADR boundary partitioning (`PRODUCT-TASK-SPEC-016-05`, `-017-05`)
- `Produce one implementation-boundary decision ledger` / `Produce one bounded decision ledger for correcting ...` — decision Task self-description pattern (`PRODUCT-TASK-SPEC-021-02`, `-021-06`)
- `Resolve the exact ... gaps identified by [an Investigation]` — decide from Investigation findings (`PRODUCT-TASK-SPEC-023-04`)
- `Reserve [an ID] as the bounded ... output` — output allocation without authoring the record (`PRODUCT-TASK-SPEC-019-04`)
- `Record the current release order without modifying the graph` — decision-ledger recording of a sequencing choice (`PRODUCT-TASK-SPEC-019-21`)

## authoring

- `Create or confirm` — produce a target file when absent, or verify an existing one (`PRODUCT-TASK-SPEC-001-01`)
- `Move`, `Copy`, `Relocate`, `Rename` — physically relocate files as part of canonical restructuring (`PRODUCT-TASK-SPEC-005-01/05/09/13/17/18`, `-007-01`)
- `Restructure [numbered] sections` — rewrite existing content into an accepted contract shape (`PRODUCT-TASK-SPEC-005-06/10`)
- `Trim` / `Replace ... file body with pointer` / `DRMCP file trimmed` — remove duplicated normative content, leave a cross-reference (`PRODUCT-TASK-SPEC-005-17/18/19`)
- `Add drift guards` — insert ownership-boundary notes into a Specification (`PRODUCT-TASK-SPEC-005-19`)
- `Guide published` — create and register a canonical authoring guide (`PRODUCT-TASK-SPEC-011-04/05/06/07`)
- `Boundary published` — create and register a canonical boundary Specification (`PRODUCT-TASK-SPEC-011-02`)
- `Create, amend, or supersede only ADRs identified by [routing]` — produce exactly the routed ADR set (`PRODUCT-TASK-SPEC-016-06`, `-017-06`, `-023-08`)
- `Apply decided ... semantics to the exact ... target sections` / `Reflect required ADR choices as current normative rules` / `Project [accepted decisions/ADR authority] into [current normative rules / a coherent contract]` — write already-decided content into canonical Specification text (`PRODUCT-TASK-SPEC-016-07`, `-017-07`, `-019-16`, `-021-11`, `-023-09`)
- `Point [an instruction file] to the successor` — edit an active instruction pointer (`PRODUCT-TASK-SPEC-018-05`)
- `Register the new topic directly under [a parent]` — add an area to a topic map (`PRODUCT-TASK-SPEC-019-16`)
- `State that [X] owns [Y]` — add an ownership rule to a Specification (`PRODUCT-TASK-SPEC-019-17`)
- `Activate [a namespace] through one bounded ... authoring set` — create/update canonical Specifications so a namespace becomes active (`PRODUCT-TASK-SPEC-021-10`)
- `Amend [an existing ADR] with [a routed clarification]` — modify one existing ADR from a fixed input (`PRODUCT-TASK-SPEC-021-16`)
- `Activate framing in [an instruction file] before [X]` — update an instruction source to make a new workflow mandatory (`PRODUCT-TASK-SPEC-024-02`)

## coordination

- `Select the first unblocked decision for [the next Task]` — choose the next decision-loop cursor, not the substantive outcome (`PRODUCT-TASK-SPEC-016-01`, `-017-01`)
- `Route [X]-owned questions to [X] without deciding them` — transfer decision ownership while preserving the boundary (`PRODUCT-TASK-SPEC-014-01`, `-016-04`, `-017-04`)
- `Amend the incomplete [Work Item] boundary and downstream route` — update an incomplete Work Item's boundary and successor route (`PRODUCT-TASK-SPEC-019-02`)
- `Synchronize the [Work Item] Task flow and Task Candidates with the materialized route` — align duplicate Work Item graph projections (`PRODUCT-TASK-SPEC-019-04`)
- `Replace the abstract ... step with exact ... routing` — replace an abstract route with concrete Task owners and dependencies (`PRODUCT-TASK-SPEC-019-06`)
- `Consume [an accepted ledger] without reopening its decisions` — use a prior accepted judgment as fixed graph-materialization input (`PRODUCT-TASK-SPEC-019-07`, `-019-13`)
- `Serialize [T14 through T17] in canonical authoring order` — set dependencies for deterministic writer sequencing (`PRODUCT-TASK-SPEC-019-13`)
- `Hand the completed route to [T13] without authoring [content]` — make a completed result the fixed input to downstream coordination (`PRODUCT-TASK-SPEC-019-12`)
- `Materialize [a decision/Investigation/coordination] route` — create the required Tasks, dependencies, gates, and release order (`PRODUCT-TASK-SPEC-020-01`, `-021-01`, `-021-07`, `-021-09`, `-023-03`, `-023-05`, `-023-07`)
- `Repair the [Work Item] Task graph so [X] stops before [Y]` — replace an invalid release route (`PRODUCT-TASK-SPEC-021-05`)
- `Release the existing [writer] after [a blocking condition] becomes canonical` — remove a resolved blocker and add a dependency (`PRODUCT-TASK-SPEC-022-03`)
- `Keep [T08] reserved until [T19] closes every required finding` — defer closure-Task materialization until a gate completes (`PRODUCT-TASK-SPEC-018-15`)
- `Route the accepted [T19] result to [T08]` — connect an accepted result to its lifecycle-synchronization owner (`PRODUCT-TASK-SPEC-018-20`)
