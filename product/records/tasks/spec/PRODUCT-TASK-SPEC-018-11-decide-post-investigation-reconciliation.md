# PRODUCT-TASK-SPEC-018-11: Decide post-Investigation reconciliation

- **id**: PRODUCT-TASK-SPEC-018-11
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-10
- **outputs**:
  - PRODUCT-TASK-SPEC-018-11

## Goal

Produce one bounded reconciliation decision ledger for every judgment candidate produced by `PRODUCT-INV-SPEC-006`.

## Work

- Classify every investigated mismatch as `consistent_refinement`, `stale_representation`, `semantic_conflict`, or `workflow_graph_drift`.
- Decide Requirement preservation, amendment, split, replacement, or follow-up.
- Decide whether W018 continues or splits.
- Decide prior-decision preservation, reconsideration, or supersession.
- Fix canonical target changes and every required graph amendment.
- Fix shared-writer policy and whether additional authoring is required.
- Preserve completed T01 and T02 decisions as historical inputs.

This Task must not:

- create or edit Tasks;
- change dependencies;
- assign writers in the graph;
- author ADRs, Specifications, Requirements, Work Items, or skills;
- perform review;
- perform synchronization.

## Done condition

- Every Investigation judgment candidate is `decided`, `deferred`, or validly `blocked`.
- Every mismatch has one final classification.
- Requirement and Work Item identity dispositions are fixed.
- Every required graph change is stated but not performed.
- Every canonical authoring target is known or blocked.
- Every shared-writer disposition is fixed.
- No completed T01 or T02 decision is substantively rewritten.

## Verification

- Confirm every `PRODUCT-INV-SPEC-006` judgment candidate has one terminal state.
- Confirm every mismatch has exactly one accepted classification or an exact blocker.
- Confirm Requirement, Work Item, graph, canonical target, and shared-writer dispositions are explicit.
- Confirm graph changes are described without being applied.
- Confirm no artifact authoring, review, synchronization, or implementation work occurred.
- Confirm T01 and T02 remain unchanged.

## Evidence

- Required input: concluded `PRODUCT-INV-SPEC-006` from `PRODUCT-TASK-SPEC-018-10`.
- Historical decision inputs: `PRODUCT-TASK-SPEC-018-01` and `PRODUCT-TASK-SPEC-018-02`.
- Downstream ADR-routing revalidation owner: `PRODUCT-TASK-SPEC-018-12`.
- No reconciliation decision was executed by T09.

### Loop state

- Status: `completed`.
- Current candidate: none.
- Candidate source: `PRODUCT-INV-SPEC-006`, `## Follow-up judgment candidates`.
- Persistence rule: Record each accepted answer before advancing the cursor.

### Judgment inventory

| candidate | topic | status | depends on | decision summary | canonical or graph target |
|---|---|---|---|---|---|
| J-001 | Canonical coordination contract. | `decided` | none | Add `work_item_decomposition` as a new Task type for parent-to-child Work Item decomposition. Keep `coordination` for Task-graph, dependency, blocker, owner, writer-order, review-order, and release-route changes. | `PRODUCT-ADR-SPEC-004`; `PRODUCT-ADR-SPEC-005`; `task_authoring`; graph-coordination workflow. |
| J-002 | Sufficiency of T03 through T06 after the late mandatory Investigation. | `decided` | J-001 | Preserve T03 through T06 as completed historical Evidence. Use one bounded authoring Task for the exact artifacts that require change. | Existing ADR, skill, activation, and Specification outputs. |
| J-003 | W018 continuation or split. | `decided` | J-001 | Continue W018 and complete the Task-type taxonomy repair inside its existing review boundary. Create no additional Work Item. | `PRODUCT-WORK-SPEC-018`. |
| J-004 | Owner and boundary for graph changes identified by T11 or T12. | `decided` | J-001 through J-003; J-005 through J-007 | Use the existing pre-review coordination route to materialize one bounded authoring Task. Keep post-review coordination conditional on the T07 verdict. | `PRODUCT-TASK-SPEC-018-13`; `PRODUCT-TASK-SPEC-018-14`; conditional post-review route. |
| J-005 | Closure-synchronization materialization timing. | `decided` | none | Materialize reserved T08 only after the accepted review route is known: either T07 returns `PASS`, or every required finding is independently `CLOSED`. | Reserved T08 or an alternative synchronization Task. |
| J-006 | Resume T07 or create a new integrated review Task. | `decided` | J-002 | Resume T07 after T12 and every required later writer complete. Preserve the prerequisite failure as historical Evidence and use T07 for the first integrated semantic verdict. | Integrated review route and release dependencies. |
| J-007 | Treatment of observed Requirement reciprocal metadata. | `decided` | none | Perform no migration or follow-up Work Item. The observed metadata is outside W018 and creates no workflow obligation. | none |

### J-001 decision

- Status: `decided`.
- Final classification: `semantic_conflict`.
- Decision: Add `work_item_decomposition` to the closed Task-type taxonomy.
- Responsibility: `work_item_decomposition` owns one bounded parent-to-child Work Item decomposition outcome.
- Completion judgment: Required child Work Items exist with distinguishable, non-overlapping responsibilities and parent-level routing.
- Coordination boundary: `coordination` owns one bounded existing workflow-graph change, including Task creation or splitting, dependencies, blockers, owners, writer order, review order, and release routing.
- Preservation: Existing child Work Item decomposition semantics remain valid but move from `coordination` to `work_item_decomposition`.
- Required canonical changes: Amend `PRODUCT-ADR-SPEC-004`, reconcile `PRODUCT-ADR-SPEC-005`, update `task_authoring`, and update affected workflow guidance.
- Required graph change: A later coordination owner must materialize the authoring route after T12 completes ADR routing.
- Reason: The two responsibilities have different primary outcomes and completion judgments. Keeping both under one `coordination` contract would preserve the current ambiguity.

### J-002 decision

- Status: `decided`.
- Final classification: `workflow_graph_drift`.
- Decision: Preserve T03 through T06 as completed historical Evidence.
- Sufficiency disposition: Their unaffected outputs remain valid inputs, but the current combined state is not final because J-001 requires canonical changes.
- Authoring disposition: Create new authoring Tasks only for exact ADR, Specification, skill, or instruction targets identified by T11 and T12.
- Prohibition: Do not reopen, rewrite, or change the completion judgments of T03 through T06.
- Required graph change: A later coordination owner must add the bounded authoring Tasks and serialize them before integrated review.
- Reason: The late Investigation revealed a bounded defect and one new design decision. It did not invalidate every previously authored workflow rule.

### J-005 decision

- Status: `decided`.
- Final classification: `workflow_graph_drift`.
- Decision: Keep `PRODUCT-TASK-SPEC-018-08` reserved and unmaterialized until the accepted review route is known.
- Direct PASS route: Materialize T08 after the integrated review returns `PASS`.
- Finding route: Materialize T08 after every required finding is independently `CLOSED`.
- Contract rule: The materialized synchronization Task must name the exact accepted review Evidence, writable targets, and mechanically derivable updates.
- Rejected route: Do not create branch-specific or speculative synchronization Tasks before review outcome.
- Required graph change: A later coordination owner must materialize T08 at the selected review-route gate.
- Reason: The exact synchronization dependency and writable boundary cannot be determined before the accepted review route exists.

### J-006 decision

- Status: `decided`.
- Final classification: `workflow_graph_drift`.
- Decision: Resume `PRODUCT-TASK-SPEC-018-07` after T12 and every required later writer complete.
- Historical preservation: Keep the existing `NOT READY` prerequisite result and `P-BLK-01` Evidence unchanged.
- Review disposition: T07 will issue the first integrated semantic verdict for the final combined W018 state.
- Required Task amendment: Add every materialized later writer to T07 dependencies and extend the review boundary to their exact outputs.
- New review Task: Not required because T07 never reached a semantic verdict and its primary outcome remains unchanged.
- Reason: D-022 applies after a completed review. T07 stopped at its prerequisite gate and remains an incomplete review Task.

### J-007 decision

- Status: `decided`.
- Final classification: `consistent_refinement`.
- Decision: Perform no migration and create no follow-up Work Item.
- W018 treatment: Ignore the observed reciprocal metadata for this workflow.
- Review treatment: The metadata creates no W018 prerequisite or review target.
- Required graph change: none.
- Reason: The accepted Task-type repair does not require existing-record migration.

### J-003 decision

- Status: `decided`.
- Final classification: `consistent_refinement`.
- Decision: Continue `PRODUCT-WORK-SPEC-018` as the one design-convergence workflow completion boundary.
- W018-owned repair: Keep the `work_item_decomposition` type addition, coordination-contract repair, affected ADR and Specification authoring, workflow-support authoring, and resumed T07 review inside W018.
- Split boundary: none.
- Identity disposition: W018 retains the same Goal, Requirement sources, and final integrated-review completion judgment.
- Reason: The taxonomy repair remains one bounded W018 authoring change.

### J-004 decision

- Status: `decided`.
- Final classification: `workflow_graph_drift`.
- Decision: Use one pre-review coordination owner and one conditional post-review coordination owner.
- Pre-review coordination: Materialize one bounded authoring Task, update W018 and T07, and release T07 behind the authoring dependency.
- Post-review coordination: After T07, materialize T08 for the accepted closure route or exact finding-derived Tasks.
- Shared-writer order: T12 routing, T13 coordination, T14 authoring, T07 integrated review, then verdict-specific coordination.
- Required graph changes: Materialize only T14 and update the existing review route.
- Reason: The Task-type addition, ADR amendments, Specification projection, and workflow-support projection share one accepted outcome.

### Final reconciliation summary

| mismatch | final classification | disposition |
|---|---|---|
| MC-001 coordination contract conflict | `semantic_conflict` | Add `work_item_decomposition`; retain `coordination` for workflow-graph changes. |
| MC-002 late mandatory Investigation | `workflow_graph_drift` | Preserve T03 through T06 and add only bounded successor authoring. |
| MC-003 missing graph-change owner | `workflow_graph_drift` | Use T13 to materialize T14 and release T07 after authoring. |
| MC-004 missing closure Task route | `workflow_graph_drift` | Materialize T08 only after the accepted review route is known. |
| MC-005 observed Requirement reciprocal metadata | `consistent_refinement` | No migration or follow-up work. |
| W018 identity | `consistent_refinement` | Continue W018 as one design-convergence completion boundary. |
| Incomplete T07 review route | `workflow_graph_drift` | Resume T07 after exact dependency and review-boundary repair. |

Canonical authoring targets requiring T12 routing are `PRODUCT-ADR-SPEC-004`, `PRODUCT-ADR-SPEC-005`, `PRODUCT-ADR-SPEC-010`, `task_authoring`, `work_item_authoring`, and affected design-convergence workflow guidance.

All seven Investigation judgment candidates are `decided`.
None is deferred or blocked.
No graph amendment, canonical authoring, review, synchronization, implementation, stage, or commit was performed by T11.
