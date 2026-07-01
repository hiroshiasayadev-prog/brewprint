# PRODUCT-ADR-SPEC-014: Preserve completed workflow records through append-only reconvergence

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**:
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-006
  - PRODUCT-ADR-SPEC-009
  - PRODUCT-ADR-SPEC-013
- **supersedes**: []
- **migrated_to_spec**: null

## Context

Authoring, review, or closure may expose missing work or a new design judgment after earlier Tasks are complete.

Reopening completed Tasks would rewrite historical outcomes and blur which Task owned the later judgment.
Closure synchronization can also become a hidden correction or graph-amendment phase when its write boundary is unclear.

The workflow needs an append-only return route and an exact closure ownership boundary.

## Decision

Treat completed decision, authoring, and review Tasks as historical records.
Do not change their outcome, Evidence, verdict, or finding set to represent later progress.

When any workflow phase discovers an unresolved design choice:

- stop downstream progression;
- return to an active incomplete decision Task when it already owns the choice and the graph remains valid;
- use coordination to create a new decision Task when the earlier decision Task is complete;
- use coordination first when ownership, dependencies, blockers, writer order, review order, release conditions, or Work Item boundaries must change.

Route integrated-review findings as follows:

| finding condition | route |
|---|---|
| The accepted decision is clear and only its projection is defective. | Use `correction`, then independent finding-closure review. |
| Resolution requires a new choice or exposes conflict among accepted decisions. | Create a new `decision` Task. |
| The required route or owner is missing. | Use `coordination` before correction or decision work. |

When a new design judgment appears after completed review:

- create a new decision Task;
- create new authoring Tasks for affected canonical artifacts;
- create a new integrated review Task for the revised combined state;
- retain completed Tasks as historical Evidence and inputs.

A closure-synchronization Task may write only:

- its own status and Evidence;
- the parent Work Item status and closure Evidence;
- exact mechanically derivable lifecycle or relation targets named by its contract.

Closure synchronization must not:

- rewrite completed decision entries;
- alter authoring Task Evidence or outputs;
- alter review verdicts or finding sets;
- create or amend the Task graph;
- introduce new design judgment;
- author or correct canonical content.

Closure must stop when it discovers missing work, graph change, or unresolved judgment.

## Rationale

Append-only workflow history preserves who decided, authored, reviewed, and closed each state.

New Tasks make later judgment explicit without falsifying earlier completion Evidence.
The closure write boundary also prevents synchronization from absorbing correction, coordination, or decision responsibilities.

Distinct return routes preserve review independence and Task completion authority.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Reopen and rewrite the completed decision Task. | The later judgment would overwrite the original completion record. |
| Update completed authoring or review Evidence after later work. | Historical execution and verdict ownership would become unreliable. |
| Let correction resolve new design choices. | Correction must consume accepted decisions rather than create them. |
| Let closure repair missing content or graph state. | Closure owns mechanical propagation, not judgment or substantive deliverables. |
| Start a new decision without updating the graph. | Missing dependencies, owners, or release gates would remain implicit. |

## Consequences

- Completed decision statuses remain `decided`; no downstream `recorded` state exists.
- Later semantic revisions require new decision Tasks.
- Post-review design changes require new authoring and integrated review Tasks.
- Correction Tasks cannot close their own findings.
- Closure synchronization has an explicit writable boundary and stop rule.
- Work Item history remains append-only across reconvergence cycles.

## Evidence

- `PRODUCT-TASK-SPEC-018-01`: D-010, D-017, D-019, D-022, and D-023.
- `PRODUCT-TASK-SPEC-018-02`: B-006 routing boundary.
- `PRODUCT-ADR-SPEC-006`: completed-decision checkpoint and non-writeback boundary.
- `PRODUCT-ADR-SPEC-013`: finding-driven Task materialization.
