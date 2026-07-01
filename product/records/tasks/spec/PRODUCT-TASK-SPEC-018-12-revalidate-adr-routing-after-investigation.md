# PRODUCT-TASK-SPEC-018-12: Revalidate ADR routing after Investigation

- **id**: PRODUCT-TASK-SPEC-018-12
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-018
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-018-11
- **outputs**:
  - PRODUCT-TASK-SPEC-018-12

## Goal

Produce one complete revalidation of the historical T02 ADR route against `PRODUCT-INV-SPEC-006` and the accepted T11 reconciliation decisions.

## Work

- Preserve `PRODUCT-TASK-SPEC-018-02` as historical routing Evidence.
- Re-evaluate D-001 through D-023 against the concluded Investigation and accepted reconciliation decisions.
- Classify every decision as `required`, `covered`, `not_required`, or `blocked`.
- Revalidate existing ADR coverage and ADR-boundary partitioning.
- Select `create`, `amend`, `reuse`, or `supersede` for every routed boundary.
- Identify exact Specification and skill targets.
- Identify exact required authoring Task boundaries.
- State whether the existing T03 through T06 outputs remain sufficient.

This Task must not:

- edit T02;
- author or amend ADR files;
- author Specifications or skills;
- change the Task graph;
- perform integrated review;
- synchronize lifecycle.

## Done condition

- Exactly 23 decision routing results exist.
- Every required decision has one coherent ADR boundary.
- Every covered decision names accepted current authority.
- Every not-required result has a reason and canonical target.
- Every blocked result has an exact blocker.
- The required post-T12 authoring route is exact.
- No ADR body is authored.

## Verification

- Count exactly 23 routing results.
- Confirm every result uses `required`, `covered`, `not_required`, or `blocked`.
- Confirm every required boundary has a coherent decision set and one disposition.
- Confirm accepted authority, direct targets, blockers, and authoring Task boundaries are exact.
- Confirm the sufficiency judgment for T03 through T06 is explicit.
- Confirm T02 and every ADR, Specification, skill, and Task graph artifact remain unchanged.
- Confirm integrated review and lifecycle synchronization did not run.

## Evidence

### Execution basis

- Historical routing input: `PRODUCT-TASK-SPEC-018-02`.
- Mandatory Investigation input: concluded `PRODUCT-INV-SPEC-006`.
- Reconciliation input: completed `PRODUCT-TASK-SPEC-018-11`.
- Current ADR authority: accepted `PRODUCT-ADR-SPEC-004` through `PRODUCT-ADR-SPEC-006` and `PRODUCT-ADR-SPEC-009` through `PRODUCT-ADR-SPEC-014`.
- Authoring mode: filesystem authoring.
- Fallback reason: `spec:product.design_records.authoring_standards.agent_authoring_policy` states that DRMCP is non-operational.
- T11 decisions are accepted inputs and were not reopened.

For `not_required` rows, disposition `reuse` means that the direct canonical projection remains authoritative and no ADR artifact is created.
The exact ADR ID is therefore `none`.

### Decision routing results

| decision | outcome | ADR boundary | exact ADR ID | disposition | routing reason | exact Specification target | exact skill or originating-artifact target | blocker |
|---|---|---|---|---|---|---|---|---|
| D-001 | `covered` | existing B-001 | `PRODUCT-ADR-SPEC-009` | `reuse` | The accepted ADR owns the end-to-end workflow boundary. | `spec:product.design_records.authoring_standards.work_item_authoring` | `skills/design-convergence-workflow/SKILL.md`; `PRODUCT-WORK-SPEC-018` | none |
| D-002 | `not_required` | none | none | `reuse` | Replacement and removal are bounded repository-transition actions with no independent durable trade-off. | none | `skills/design-convergence-workflow/SKILL.md`; `prompt_chappy.md`; retired `skills/design-decision-workflow/` path | none |
| D-003 | `not_required` | none | none | `reuse` | The successor path is a naming and placement choice. | none | `skills/design-convergence-workflow/`; `prompt_chappy.md` | none |
| D-004 | `covered` | existing B-001 | `PRODUCT-ADR-SPEC-009` | `reuse` | The accepted ADR starts convergence when a design topic is raised. | none | `skills/design-convergence-workflow/SKILL.md` | none |
| D-005 | `covered` | existing B-001 | `PRODUCT-ADR-SPEC-009` | `reuse` | The accepted ADR owns reviewed closure and synchronization as the workflow terminal state. T11 only fixes later graph materialization timing. | `spec:product.design_records.authoring_standards.work_item_authoring` | `skills/design-convergence-workflow/SKILL.md`; `skills/design-convergence-workflow/closure-synchronization.md`; `PRODUCT-WORK-SPEC-018` | none |
| D-006 | `required` | B-008 | `PRODUCT-ADR-SPEC-004` | `amend` | T11 adds `work_item_decomposition` and separates its outcome from workflow-graph coordination. The closed-taxonomy choice remains unchanged. | `spec:product.design_records.authoring_standards.task_authoring`; `spec:product.design_records.authoring_standards.work_item_authoring` | `skills/design-convergence-workflow/SKILL.md`; new `skills/design-convergence-workflow/work-item-decomposition.md`; `skills/design-convergence-workflow/graph-coordination.md`; `skills/design-convergence-workflow/convergence-routing.md` | none |
| D-007 | `covered` | existing authority | `PRODUCT-ADR-SPEC-004` | `reuse` | The accepted Investigation type still owns one formal Investigation for one bounded question. | `spec:product.design_records.authoring_standards.task_authoring` | `skills/design-convergence-workflow/impact-investigation.md`; `PRODUCT-INV-SPEC-006` | none |
| D-008 | `covered` | existing authority | `PRODUCT-ADR-SPEC-004`; `PRODUCT-ADR-SPEC-005` | `reuse` | Decision ownership and the prohibition on graph or canonical authoring remain accepted and compatible. | `spec:product.design_records.authoring_standards.task_authoring` | `skills/design-convergence-workflow/convergence-routing.md`; `PRODUCT-TASK-SPEC-018-11` | none |
| D-009 | `required` | B-009 | `PRODUCT-ADR-SPEC-005` | `amend` | The single-responsibility decision remains valid, but the coordination section must exclude parent-to-child Work Item decomposition and retain Task-graph routing. | `spec:product.design_records.authoring_standards.task_authoring`; `spec:product.design_records.authoring_standards.work_item_authoring` | `skills/design-convergence-workflow/graph-coordination.md`; new `skills/design-convergence-workflow/work-item-decomposition.md`; `skills/design-convergence-workflow/convergence-routing.md` | none |
| D-010 | `covered` | existing B-006 | `PRODUCT-ADR-SPEC-014` | `reuse` | Append-only return-to-decision routing remains accepted and was followed by T09 through T12. | `spec:product.design_records.authoring_standards.task_authoring` | `skills/design-convergence-workflow/SKILL.md`; `skills/design-convergence-workflow/interactive-decision-loop.md` | none |
| D-011 | `required` | B-010 | `PRODUCT-ADR-SPEC-010` | `amend` | Preserve the four mismatch classes while adding the distinct Work Item decomposition phase and narrowing coordination. | none | `skills/design-convergence-workflow/SKILL.md`; `skills/design-convergence-workflow/graph-coordination.md`; `skills/design-convergence-workflow/convergence-routing.md`; new `skills/design-convergence-workflow/work-item-decomposition.md` | none |
| D-012 | `covered` | existing B-003 | `PRODUCT-ADR-SPEC-011` | `reuse` | Requirement identity and follow-up boundaries remain unchanged. | `spec:product.design_records.authoring_standards.requirement_authoring` | `skills/design-convergence-workflow/convergence-routing.md`; `PRODUCT-REQ-SPEC-005`; `PRODUCT-REQ-SPEC-006` | none |
| D-013 | `covered` | existing B-003 | `PRODUCT-ADR-SPEC-011` | `reuse` | W018 keeps one completion identity and creates no additional Work Item. | `spec:product.design_records.authoring_standards.work_item_authoring` | `PRODUCT-WORK-SPEC-018` | none |
| D-014 | `covered` | existing authority | `PRODUCT-ADR-SPEC-005` | `reuse` | Same-type amendment and split-on-distinct-responsibility remain accepted. The new type applies this rule. | `spec:product.design_records.authoring_standards.task_authoring` | `skills/design-convergence-workflow/graph-coordination.md` | none |
| D-015 | `covered` | existing B-004 | `PRODUCT-ADR-SPEC-012` | `reuse` | Deterministic shared-writer serialization remains accepted. | `spec:product.design_records.authoring_standards.work_item_authoring` | `skills/design-convergence-workflow/SKILL.md`; `skills/design-convergence-workflow/graph-coordination.md` | none |
| D-016 | `covered` | existing B-004 | `PRODUCT-ADR-SPEC-012` | `reuse` | One integrated review after the final writer remains accepted. T11 selected resumption of incomplete T07. | `spec:product.design_records.authoring_standards.work_item_authoring` | `skills/design-convergence-workflow/design-review-gate.md`; `PRODUCT-TASK-SPEC-018-07` | none |
| D-017 | `covered` | existing B-006 | `PRODUCT-ADR-SPEC-014` | `reuse` | Finding correction and renewed decision work remain separate return routes. | `spec:product.design_records.authoring_standards.task_authoring` | `skills/design-convergence-workflow/design-review-gate.md` | none |
| D-018 | `covered` | existing authority | `PRODUCT-ADR-SPEC-005` | `reuse` | The accepted verification boundary still makes a separate pre-authoring verification Task conditional. | `spec:product.design_records.authoring_standards.task_authoring` | `skills/design-convergence-workflow/SKILL.md`; `skills/design-convergence-workflow/graph-coordination.md` | none |
| D-019 | `covered` | existing authority | `PRODUCT-ADR-SPEC-006`; `PRODUCT-ADR-SPEC-014` | `reuse` | Completed decision Tasks remain historical checkpoints and receive no downstream writeback. | `spec:product.design_records.authoring_standards.task_authoring`; `spec:product.design_records.authoring_standards.artifact_boundary`; `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `skills/design-convergence-workflow/SKILL.md`; `PRODUCT-TASK-SPEC-018-01`; `PRODUCT-TASK-SPEC-018-02` | none |
| D-020 | `covered` | existing B-005 | `PRODUCT-ADR-SPEC-013` | `reuse` | Finding repair and closure-review Tasks remain delayed until named findings exist. | `spec:product.design_records.authoring_standards.work_item_authoring`; `spec:product.design_records.authoring_standards.task_authoring` | `skills/design-convergence-workflow/design-review-gate.md`; `PRODUCT-WORK-SPEC-018` | none |
| D-021 | `covered` | existing authority | `PRODUCT-ADR-SPEC-006` | `reuse` | The accepted ADR owns separate routing, coherent boundaries, and create, amend, reuse, or supersede disposition. | `spec:product.design_records.authoring_standards.task_authoring`; `spec:product.design_records.authoring_standards.artifact_boundary` | `skills/design-convergence-workflow/adr-routing.md`; `PRODUCT-TASK-SPEC-018-02`; this Task | none |
| D-022 | `covered` | existing B-006 | `PRODUCT-ADR-SPEC-014` | `reuse` | The accepted ADR requires a new review only after a completed semantic review. T07 has no semantic verdict and may resume. | `spec:product.design_records.authoring_standards.work_item_authoring`; `spec:product.design_records.authoring_standards.task_authoring`; `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `skills/design-convergence-workflow/design-review-gate.md`; `PRODUCT-TASK-SPEC-018-07` | none |
| D-023 | `covered` | existing B-006 | `PRODUCT-ADR-SPEC-014` | `reuse` | The closure write boundary remains accepted. T11 fixes only the later materialization gate. | `spec:product.design_records.authoring_standards.task_authoring`; `spec:product.design_records.authoring_standards.artifact_boundary`; `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `skills/design-convergence-workflow/closure-synchronization.md`; reserved `PRODUCT-TASK-SPEC-018-08` | none |

Routing row count: 23.

Outcome counts:

| outcome | count |
|---|---:|
| `required` | 3 |
| `covered` | 18 |
| `not_required` | 2 |
| `blocked` | 0 |

### Required ADR boundaries

#### B-008 — Extend the closed Task taxonomy with Work Item decomposition

- Included decisions: D-006.
- Bounded durable question: Which Task type owns bounded parent-to-child Work Item decomposition after coordination is restricted to workflow-graph change?
- Coherence: The new type, its primary outcome, its completion judgment, and the narrowed coordination row are one closed-taxonomy change.
- Disposition: `amend`.
- Exact ADR ID: `PRODUCT-ADR-SPEC-004`.
- Amendment versus supersession: Amend. The required scalar field, closed taxonomy model, and one-outcome-per-type choice remain unchanged. The amendment adds one type and removes one overloaded outcome from `coordination` without reversing the selected taxonomy architecture.
- Predecessor ADR dependencies: none in ADR metadata. Accepted `PRODUCT-ADR-SPEC-005` is reconciliation authority and is amended after B-008.
- Exact downstream ADR authoring Task boundary: One `authoring` Task updates only `PRODUCT-ADR-SPEC-004` and `PRODUCT-ADR-SPEC-005`. Within that Task, write B-008 before B-009.
- Exact writer order: `PRODUCT-ADR-SPEC-004` amendment, then `PRODUCT-ADR-SPEC-005` amendment.
- Affected Specification targets: `spec:product.design_records.authoring_standards.task_authoring`; `spec:product.design_records.authoring_standards.work_item_authoring`.
- Affected workflow-support targets: `skills/design-convergence-workflow/SKILL.md`; new `skills/design-convergence-workflow/work-item-decomposition.md`; `skills/design-convergence-workflow/graph-coordination.md`; `skills/design-convergence-workflow/convergence-routing.md`.

#### B-009 — Separate workflow-graph coordination from Work Item decomposition

- Included decisions: D-009.
- Bounded durable question: Which graph changes remain owned by `coordination`, and which parent-to-child Work Item changes move to `work_item_decomposition`?
- Coherence: The boundary reconciles the adjacent Task responsibilities without changing the general single-responsibility rule.
- Disposition: `amend`.
- Exact ADR ID: `PRODUCT-ADR-SPEC-005`.
- Amendment versus separate replacement: Amend in a separate ADR boundary from B-008. ADR-005 owns cross-type cohesion and adjacent responsibility boundaries. ADR-004 owns the closed type inventory. The two records can change independently and must not be merged into one omnibus ADR.
- Predecessor ADR dependencies: amended `PRODUCT-ADR-SPEC-004` from B-008.
- Exact downstream ADR authoring Task boundary: The same bounded ADR `authoring` Task consumes B-008 and B-009, writes two existing ADRs, and has one completion judgment that both authorities express the accepted responsibility split.
- Exact writer order: B-008 first, then B-009.
- Affected Specification targets: `spec:product.design_records.authoring_standards.task_authoring`; `spec:product.design_records.authoring_standards.work_item_authoring`.
- Affected workflow-support targets: `skills/design-convergence-workflow/graph-coordination.md`; new `skills/design-convergence-workflow/work-item-decomposition.md`; `skills/design-convergence-workflow/convergence-routing.md`; `skills/design-convergence-workflow/SKILL.md`.

#### B-010 — Add Work Item decomposition to typed convergence phases

- Included decisions: D-011.
- Bounded durable question: How does the workflow route an accepted child Work Item boundary after coordination is restricted to Task-graph change?
- Coherence: The mismatch classes remain unchanged. The responsibility phase list and coordination boundary require one clarification.
- Disposition: `amend`.
- Exact ADR ID: `PRODUCT-ADR-SPEC-010`.
- Amendment versus supersession: Amend. Typed convergence, conditional routing, and mismatch classification remain the selected architecture.
- Exact downstream authoring boundary: T14 updates ADR-010 with ADR-004, ADR-005, Specification, and workflow-support projections under one accepted Task-type addition.

### Existing ADR treatment

| ADR | treatment | reason |
|---|---|---|
| `PRODUCT-ADR-SPEC-004` | `amend` | Add `work_item_decomposition` and narrow the overloaded `coordination` row. The closed-taxonomy architecture remains current. |
| `PRODUCT-ADR-SPEC-005` | `amend` | Preserve single responsibility and graph coordination. Remove parent-to-child decomposition from coordination ownership. |
| `PRODUCT-ADR-SPEC-006` | `reuse` | Checkpoint, canonical ownership, non-writeback, and ADR-routing separation remain accepted. |
| `PRODUCT-ADR-SPEC-009` | `reuse` | End-to-end workflow start and reviewed closure remain unchanged. |
| `PRODUCT-ADR-SPEC-010` | `amend` | Preserve typed convergence and mismatch routing while adding the distinct Work Item decomposition phase. |
| `PRODUCT-ADR-SPEC-011` | `reuse` | Requirement and Work Item identity rules remain current. |
| `PRODUCT-ADR-SPEC-012` | `reuse` | Shared-writer serialization and one final integrated review remain current. |
| `PRODUCT-ADR-SPEC-013` | `reuse` | Finding-driven repair materialization remains current. |
| `PRODUCT-ADR-SPEC-014` | `reuse` | Append-only reconvergence and closure write limits remain current. |

Supersessions: none.
Creates: none.
The `work_item_decomposition` addition and the `coordination` reconciliation use separate ADR boundaries because the closed type inventory and cross-type responsibility rule have distinct durable questions.

### Exact canonical targets

#### ADR targets

- Amend `PRODUCT-ADR-SPEC-004`.
- Amend `PRODUCT-ADR-SPEC-005`.
- Amend `PRODUCT-ADR-SPEC-010`.
- Reuse `PRODUCT-ADR-SPEC-006`, `PRODUCT-ADR-SPEC-009`, and `PRODUCT-ADR-SPEC-011` through `PRODUCT-ADR-SPEC-014` without body changes.

#### Specification targets

- Update `spec:product.design_records.authoring_standards.task_authoring`:
  - add `work_item_decomposition` to the closed value set;
  - define its outcome and completion judgment;
  - restrict `coordination` to Task creation or split, dependency, blocker, owner, writer order, review order, and release routing;
  - define the adjacent boundary between graph coordination and parent-to-child Work Item decomposition.
- Update `spec:product.design_records.authoring_standards.work_item_authoring`:
  - assign bounded parent-to-child Work Item creation and decomposition to `work_item_decomposition`;
  - retain parent overview and source-relation rules;
  - distinguish coordination routing from decomposition authoring.
- Do not re-author `requirement_authoring`, `artifact_boundary`, or `artifact_responsibility_matrix`.

#### Workflow-support targets

- Update `skills/design-convergence-workflow/SKILL.md` with the new owner and companion route.
- Create `skills/design-convergence-workflow/work-item-decomposition.md` as the bounded decomposition companion.
- Update `skills/design-convergence-workflow/graph-coordination.md` to route child Work Item creation to `work_item_decomposition`.
- Update `skills/design-convergence-workflow/convergence-routing.md` so Work Item split or creation routes through the new type after the identity decision.
- Preserve `skills/design-convergence-workflow/decision-ledger.md` as historical duplicate input. Do not rewrite it.
- Do not update `prompt_chappy.md` or the retired skill path.

### Exact downstream Task boundary

#### Unified T14 authoring boundary

- Task type: `authoring`.
- Primary outcome: One coherent projection of the accepted `work_item_decomposition` Task-type split.
- Inputs: T11 J-001, T12 B-008, and T12 B-009.
- Writable ADR targets: exactly `PRODUCT-ADR-SPEC-004`, `PRODUCT-ADR-SPEC-005`, and `PRODUCT-ADR-SPEC-010`.
- Writable Specification targets: exactly `task_authoring` and `work_item_authoring`.
- Writable workflow-support targets: exactly `SKILL.md`, `graph-coordination.md`, `convergence-routing.md`, and new `work-item-decomposition.md`.
- Completion judgment: ADR, Specification, and workflow-support authorities express one consistent responsibility split.
- Dependency: T13.
- Prohibited: Requirement changes, Work Item creation, migration, graph changes, review, synchronization, and implementation work.

### Shared-writer and review sequence

1. T12 completes ADR routing.
2. T13 materializes T14 and releases T07 as `not_started` behind the T14 dependency.
3. T14 amends ADR-004, ADR-005, and ADR-010.
4. T14 updates `task_authoring` and `work_item_authoring`.
5. T14 updates the workflow-support files and creates `work-item-decomposition.md`.
6. T14 completion satisfies the final T07 dependency.
7. T07 reviews the final combined W018 state.
8. Verdict-specific coordination runs only after T07.

Concurrent writes to the same ADR, Specification, skill file, W018 graph section, or T07 contract are prohibited.

### T07 dependency and review-boundary additions

Post-T12 W018 coordination must add these exact later owners to T07:

- `PRODUCT-TASK-SPEC-018-13`;
- `PRODUCT-TASK-SPEC-018-14`.

T07 must review:

- amended ADR-004, ADR-005, and ADR-010;
- updated `task_authoring` and `work_item_authoring`;
- updated `SKILL.md`, `graph-coordination.md`, and `convergence-routing.md`;
- new `work-item-decomposition.md`;
- the final W018 and T07 graph state.

### T03 through T06 sufficiency

- T03 through T06 remain `done` and retain their historical Evidence.
- T03 output `PRODUCT-ADR-SPEC-006` remains sufficient and is reused.
- T04 outputs `PRODUCT-ADR-SPEC-009` through `PRODUCT-ADR-SPEC-014` remain sufficient and are reused.
- T05 activation, old-skill removal, and unaffected workflow rules remain sufficient.
- T06 projections to `requirement_authoring`, `artifact_boundary`, and `artifact_responsibility_matrix` remain sufficient.
- T05 workflow-support files named above and T06 `task_authoring` and `work_item_authoring` require bounded successor authoring.
- No T03 through T06 Task is reopened or rewritten.
- The current combined state requires additional ADR, Specification, and workflow-support authoring before T07 resumes.

### Required graph changes not applied

Post-T12 W018 coordination must later:

- materialize one bounded T14 authoring Task;
- update W018 routing without changing its identity;
- add T13 and T14 to the T07 dependency and review route;
- release T07 as `not_started` with T14 as an unmet dependency;
- let T14 completion satisfy the final pre-review gate.

Post-review coordination is not materialized now.
After T07:

- `PASS` permits coordination to materialize reserved T08 with the exact accepted review dependency and writable boundary;
- `NEEDS REVISION` permits coordination to materialize only finding-derived correction and independent closure-review Tasks;
- T08 remains unmaterialized until the accepted review route exists.

No graph change was applied by T12.

### Completion verification

- Routing rows: 23.
- Allowed outcomes only: yes.
- Required items in one coherent boundary each: yes.
- Covered items name accepted, non-superseded authority: yes.
- Not-required items include reason and exact direct target: yes.
- Blocked items: none.
- ADR dispositions are unique: yes.
- Exact authoring Task boundaries: recorded.
- T03 through T06 sufficiency: recorded.
- T02, ADRs, Specifications, skills, Work Item, and Task graph: not changed by T12.
- Integrated review, synchronization, migration, implementation, stage, and commit: not performed.

### Final routing summary

Result: `PASS`.

Three durable ADR amendments are required.
Eighteen decisions remain covered by accepted authority.
Two decisions remain direct projections without ADR creation.
No decision is blocked.

The next gate is T13 coordination.
T07 remains blocked until the bounded T14 authoring result is present.
