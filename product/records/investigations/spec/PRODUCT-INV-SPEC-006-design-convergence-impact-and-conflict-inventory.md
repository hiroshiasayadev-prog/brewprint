# PRODUCT-INV-SPEC-006: Design convergence impact and conflict inventory

- **status**: concluded
- **date**: 2026-07-01
- **trigger**: PRODUCT-TASK-SPEC-018-07 prerequisite P-BLK-01
- **scope**: D-001 through D-023 against the current W018 combined design state.
- **non_scope**: Decision adoption, graph amendment, canonical authoring, review, synchronization, implementation, and migration.
- **source_refs**:
  - PRODUCT-WORK-SPEC-018
  - PRODUCT-REQ-SPEC-005
  - PRODUCT-REQ-SPEC-006
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-006
- **follow_up_candidates**:
  - post-Investigation reconciliation decision
  - ADR-routing revalidation
  - conditional graph coordination
  - conditional canonical authoring
- **related_requirements**:
  - PRODUCT-REQ-SPEC-005
  - PRODUCT-REQ-SPEC-006
- **related_work_items**:
  - PRODUCT-WORK-SPEC-018
- **related_adrs**:
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-006
  - PRODUCT-ADR-SPEC-009
  - PRODUCT-ADR-SPEC-010
  - PRODUCT-ADR-SPEC-011
  - PRODUCT-ADR-SPEC-012
  - PRODUCT-ADR-SPEC-013
  - PRODUCT-ADR-SPEC-014
- **related_specs**:
  - spec:product.design_records.authoring_standards.requirement_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.artifact_boundary
  - spec:product.design_records.artifact_model.artifact_responsibility_matrix

## Investigation scope

This Investigation examines one bounded question:

```text
For D-001 through D-023, how does the completed decision set relate to the
current W018 combined state across repository authority, semantic conflicts,
workflow graph, shared writers, existing ADR coverage, and unresolved evidence?
```

The inspected combined state includes:

- the W018 identity, boundary, completion conditions, and Task graph;
- completed decisions D-001 through D-023;
- historical ADR routing and completed authoring outputs;
- the accepted source Requirements and ADRs;
- the successor workflow skill and activation state;
- the five canonical Specification targets;
- the T07 prerequisite failure and the T09 through T12 reconvergence route.

The Investigation proposes relation classes and follow-up questions.
It does not adopt a final mismatch classification or resolution.

## Out of scope

- Reconciliation decisions.
- Final mismatch-class adoption.
- Task creation, Task amendment, dependency changes, or release changes.
- Requirement, Work Item, ADR, Specification, skill, or instruction authoring.
- Integrated review, correction, finding closure, or lifecycle synchronization.
- Production implementation or existing-record migration.
- Inspection of the retired copy under `memory/`.

## Background

T07 returned `NOT READY` at prerequisite `P-BLK-01`.
The review found no formal Investigation owner or Investigation record.

T09 repaired the missing-owner graph.
T09 created T10, T11, and T12 and reserved `PRODUCT-INV-SPEC-006` for T10.

The repair occurred after T03 through T06 had completed ADR, skill, activation, and Specification authoring.
The current graph therefore reconverges after completed authoring rather than before authoring.

## What was investigated

### Direct record boundary

- `PRODUCT-WORK-SPEC-018`.
- T01 through T07 and T09 through T12.
- `PRODUCT-REQ-SPEC-005` and `PRODUCT-REQ-SPEC-006`.
- `PRODUCT-ADR-SPEC-004` through `PRODUCT-ADR-SPEC-006`.
- `PRODUCT-ADR-SPEC-009` through `PRODUCT-ADR-SPEC-014`.
- All nine successor workflow files.
- The five canonical Specification targets named by W018.
- `prompt_chappy.md` and the repository state of `skills/design-decision-workflow/`.

### Questions applied to each decision

- Which artifact currently owns the rule?
- Does the current text refine, contradict, or fail to represent the decision?
- Does an accepted ADR cover the durable choice?
- Does W018 retain one coherent completion boundary?
- Does every required downstream responsibility have an owner?
- Are dependencies, blockers, writer order, review order, and release conditions complete?
- Would future work write an artifact already written by a completed Task?
- Does any missing Evidence block T11, T12, or resumed review?

## Findings

### Decision impact inventory

| decision | current owner | affected artifact | observed state | relation candidate | evidence | required judgment | uncertainty |
|---|---|---|---|---|---|---|---|
| D-001 | ADR, Work Item, skill | `PRODUCT-ADR-SPEC-009`; `PRODUCT-WORK-SPEC-018`; successor `SKILL.md` | All three define one end-to-end workflow from topic intake through reviewed closure. | `consistent_refinement` | `PRODUCT-ADR-SPEC-009`, `## Decision`; W018, `## Goal` and `## Boundary`; `SKILL.md`, `## End-to-end boundary` | none | none |
| D-002 | authoring Task and instruction activation | `prompt_chappy.md`; `skills/design-decision-workflow/` | The prompt activates only the successor. The old repository path is absent. | `consistent_refinement` | T05, `## Evidence`; `prompt_chappy.md`, `### Mandatory design-convergence workflow skill`; scoped old-path existence check | none | The ignored retired copy was intentionally not inspected. |
| D-003 | skill and instruction activation | `skills/design-convergence-workflow/`; `prompt_chappy.md` | The active path and instruction pointer use the decided successor name. | `consistent_refinement` | T01, `### D-003`; T05, `## Evidence`; `prompt_chappy.md`, workflow skill section | none | none |
| D-004 | ADR and skill | `PRODUCT-ADR-SPEC-009`; successor `SKILL.md` | Both start the workflow when a design topic is raised, before inventory. | `consistent_refinement` | `PRODUCT-ADR-SPEC-009`, `## Decision`; `SKILL.md`, `## End-to-end boundary` | none | none |
| D-005 | ADR, Work Item, skill, future synchronization owner | `PRODUCT-ADR-SPEC-009`; W018; reserved `PRODUCT-TASK-SPEC-018-08` | Reviewed closure and synchronization are required, but T08 is not materialized and no concrete post-review materialization owner exists. | `workflow_graph_drift` | W018, `## Task flow`, `## Task Candidates`, and `## Completion Condition`; `closure-synchronization.md`, `## Entry gate` | T11 reconciliation and a later coordination owner | Whether T08 should exist before review or only after the accepted review route. |
| D-006 | ADR, skill, Work Item, Task taxonomy | `PRODUCT-ADR-SPEC-010`; W018; `PRODUCT-ADR-SPEC-004`; `task_authoring` | W018 now contains all required phases, but Investigation and reconciliation occur after completed authoring. The generic `coordination` contract is also narrower than T09 graph coordination. | `semantic_conflict` | T01, `### D-006`; `PRODUCT-ADR-SPEC-010`, `## Decision`; W018, `## Task flow`; `PRODUCT-ADR-SPEC-004`, coordination row; `task-authoring.md`, Task type contract | T11 reconciliation | Whether append-only reconvergence can preserve T03 through T06 without new authoring. |
| D-007 | Investigation Task and record | T10; `PRODUCT-INV-SPEC-006`; impact-investigation skill | A formal Investigation now exists and covers one bounded question, but it was added after T03 through T06. | `workflow_graph_drift` | T07, `### Prerequisite result`; T09, `## Evidence`; T10, `## Goal`; `impact-investigation.md`, `## Investigation Task contract` | T11 reconciliation | Whether the late Investigation changes the validity or sufficiency of earlier outputs. |
| D-008 | reconciliation decision owner | T11; `convergence-routing.md` | T11 owns the exact post-Investigation judgments and prohibits graph or canonical authoring. | `consistent_refinement` | T11, `## Goal`, `## Work`, and `## Done condition`; `convergence-routing.md`, `## Reconciliation decision ownership` | T11 reconciliation | none |
| D-009 | coordination authority | T09; `PRODUCT-ADR-SPEC-005`; `PRODUCT-ADR-SPEC-004`; `task_authoring` | T09 performs valid graph repair under ADR-005 and the skill. ADR-004 and the Task type table still define coordination as child-Work-Item decomposition. | `semantic_conflict` | T09, `## Goal` and `## Work`; `PRODUCT-ADR-SPEC-005`, `### Coordination and synchronization`; `PRODUCT-ADR-SPEC-004`, coordination row; `task-authoring.md`, Task type contract | T11 reconciliation | No explicit amendment or supersession resolves the two coordination definitions. |
| D-010 | ADR, Task Specification, skill | `PRODUCT-ADR-SPEC-014`; `task_authoring`; successor `SKILL.md` | T01 and T02 remain complete. T09 created new T11 and T12 owners instead of reopening them. | `consistent_refinement` | `PRODUCT-ADR-SPEC-014`, `## Decision`; `task-authoring.md`, `#### Task continuation and reconvergence routes`; T09, `## Done condition` | none | none |
| D-011 | ADR, reconciliation skill, T11 | `PRODUCT-ADR-SPEC-010`; `convergence-routing.md`; T11 | All four classes exist. T11 owns final classification after this Investigation. | `consistent_refinement` | `PRODUCT-ADR-SPEC-010`, mismatch table; `convergence-routing.md`, `## Mismatch classes`; T11, `## Work` | T11 reconciliation | none |
| D-012 | ADR and Requirement Specification | `PRODUCT-ADR-SPEC-011`; `requirement_authoring` | The Requirement identity rules match D-012. W018 still uses the same accepted source Requirements. | `consistent_refinement` | `PRODUCT-ADR-SPEC-011`, `## Decision`; `requirement-authoring.md`, `#### Requirement identity during design convergence`; W018 metadata `source_refs` | T11 confirmation of no Requirement identity change | The persisted `work_items` fields on the source Requirements are a separate stale relation issue. |
| D-013 | ADR, Work Item Specification, W018 | `PRODUCT-ADR-SPEC-011`; `work_item_authoring`; W018 | W018 retains the same goal and completion boundary while adding reconvergence Tasks. | `consistent_refinement` | `PRODUCT-ADR-SPEC-011`, `## Decision`; `work-item-authoring.md`, `#### Work Item identity and design-convergence review`; W018, `## Goal` and `## Completion Condition` | T11 confirmation of continuation or split | Whether the coordination-taxonomy repair is independently completable. |
| D-014 | ADR, Task Specification, coordination output | `PRODUCT-ADR-SPEC-005`; `task_authoring`; T09; T07 | T09 created new responsibility Tasks and amended the incomplete T07 release route. Completed T01 through T06 remained unchanged. | `consistent_refinement` | `PRODUCT-ADR-SPEC-005`, `## Decision`; `task-authoring.md`, `#### Task continuation and reconvergence routes`; T09, `## Verification` | none | none |
| D-015 | ADR, skill, future writer graph | `PRODUCT-ADR-SPEC-012`; `graph-coordination.md`; completed T03 through T06 outputs | Completed writers were serialized or wrote separate targets. Any new writer after T11 or T12 would revisit completed outputs and needs explicit ordering. | `consistent_refinement` | `PRODUCT-ADR-SPEC-012`, `## Decision`; `graph-coordination.md`, `## Shared-writer serialization`; T03 through T06 metadata and outputs | T11 shared-writer disposition | The exact future writer set is unknown until T11 and T12 complete. |
| D-016 | ADR, review Task, Work Item | `PRODUCT-ADR-SPEC-012`; T07; W018 | One integrated review Task exists and remains blocked. Its current dependency list cannot name conditional future writers. | `workflow_graph_drift` | `PRODUCT-ADR-SPEC-012`, `## Decision`; T07 metadata `depends_on` and `### Reconvergence Evidence`; W018, `## Task flow` | T11 reconciliation and conditional coordination | Whether T07 can resume or a new integrated review Task is required. |
| D-017 | ADR and review skill | `PRODUCT-ADR-SPEC-014`; `design-review-gate.md`; T07 | T07 recorded a prerequisite failure, not an integrated-review finding. No speculative correction Task exists. | `consistent_refinement` | T07, `### Prerequisite result` and `### Named findings`; `design-review-gate.md`, `## Finding routing` | none | none |
| D-018 | ADR, graph skill, authoring skill | `PRODUCT-ADR-SPEC-005`; `graph-coordination.md`; `design-authoring.md`; T03 through T06 | The current skill requires Investigation and fixed graph ownership before authoring. T03 through T06 completed before those preconditions existed. | `workflow_graph_drift` | T01, `### D-018`; `graph-coordination.md`, `## Pre-authoring release gate`; `design-authoring.md`, `## Preconditions`; W018, `## Task flow` | T11 reconciliation | Whether revalidation alone is enough or new authoring is required. |
| D-019 | ADRs, Task Specification, artifact ownership specs | `PRODUCT-ADR-SPEC-006`; `PRODUCT-ADR-SPEC-014`; `task_authoring`; `artifact_boundary`; responsibility matrix | Completed decisions remain unchanged. Downstream Evidence is owned by downstream Tasks. No `recorded` state exists. | `consistent_refinement` | `PRODUCT-ADR-SPEC-006`, `### Canonical ownership`; `PRODUCT-ADR-SPEC-014`, `## Decision`; `task-authoring.md`, `#### Decision workflow Evidence` | none | none |
| D-020 | ADR, Work Item Specification, W018 | `PRODUCT-ADR-SPEC-013`; `work_item_authoring`; W018 | Correction and finding-closure Tasks remain absent. The conditional route is recorded. | `consistent_refinement` | `PRODUCT-ADR-SPEC-013`, `## Decision`; `work-item-authoring.md`, convergence review rules; W018, `## Task flow` | none | none |
| D-021 | ADR, historical routing Task, revalidation Task | `PRODUCT-ADR-SPEC-006`; T02; T12; `adr-routing.md` | T02 contains a complete historical route, but it predates the mandatory Investigation. T12 now owns revalidation. | `workflow_graph_drift` | `PRODUCT-ADR-SPEC-006`, `### Conditional ADR routing`; T02, `## Evidence`; T12, `## Goal`; `adr-routing.md`, `## Inputs` | T12 after T11 | Whether T03 through T06 remain sufficient after revalidation. |
| D-022 | ADR, Work Item, review Task | `PRODUCT-ADR-SPEC-014`; W018; T07 | W018 plans to resume the incomplete T07. D-022 can also be read to require a new integrated review Task after new post-authoring judgment. | `semantic_conflict` | T01, `### D-022`; `PRODUCT-ADR-SPEC-014`, `## Decision`; W018, `## Task flow`; T07, `### Reconvergence Evidence` | T11 reconciliation | Whether D-022 applies when the original integrated review never reached a semantic verdict. |
| D-023 | ADR, closure skill, Work Item | `PRODUCT-ADR-SPEC-014`; `closure-synchronization.md`; reserved T08 | The closure write boundary is explicit, but no concrete synchronization Task currently owns it. | `workflow_graph_drift` | `PRODUCT-ADR-SPEC-014`, closure boundary; `closure-synchronization.md`, `## Synchronization Task contract`; W018, `## Task Candidates` | T11 reconciliation and conditional coordination | The correct T08 materialization point is not decided. |

Decision rows: 23.

Proposed relation summary:

| relation candidate | count |
|---|---:|
| `consistent_refinement` | 14 |
| `stale_representation` | 0 |
| `semantic_conflict` | 3 |
| `workflow_graph_drift` | 6 |

These counts are proposed Investigation results.
T11 owns final classification.

### Artifact and conflict inventory

#### Affected Requirements

| artifact | observed state | proposed relation | impact |
|---|---|---|---|
| `PRODUCT-REQ-SPEC-005` | Accepted Task responsibility requirement. The file still persists `work_items`. | `stale_representation` candidate against REQ-006 and `requirement_authoring` | The stale reciprocal field does not change the W018 workflow decisions, but it conflicts with current relation ownership. |
| `PRODUCT-REQ-SPEC-006` | Accepted generic relation requirement. The file still persists its own `work_items` field. | `stale_representation` candidate against its Requirement text and `requirement_authoring` | Existing-record migration ownership is outside W018 and is not visible in the required read boundary. |

#### Affected ADRs

| ADR | status | durable choice | W018 alignment | conflict or stale consequence | supersession |
|---|---|---|---|---|---|
| `PRODUCT-ADR-SPEC-004` | accepted | Closed Task taxonomy and type contracts. | Investigation and decision types align. | The coordination row is narrower than W018, ADR-005, the skill, and T09. | none |
| `PRODUCT-ADR-SPEC-005` | accepted | One outcome and completion judgment; broad graph coordination boundary. | Supports T09 and conditional graph repair. | Conflicts with ADR-004's coordination outcome unless an amendment or precedence rule exists. | none |
| `PRODUCT-ADR-SPEC-006` | accepted | Decision checkpoints remain non-canonical; routing is separate from authoring. | Aligns with T01, T02, T11, and T12 preservation. | No stale decision writeback remains. Historical routing still predates the mandatory Investigation. | none |
| `PRODUCT-ADR-SPEC-009` | accepted | End-to-end workflow boundary. | Aligns with W018 and the successor skill. | Closure Task ownership is not materialized. | none |
| `PRODUCT-ADR-SPEC-010` | accepted | Typed convergence phases and mismatch routing. | The current graph contains the phases. | The current sequence inserted mandatory Investigation after completed authoring. | none |
| `PRODUCT-ADR-SPEC-011` | accepted | Requirement and Work Item identity continuity. | W018 still appears to retain one resolution identity. | T11 must confirm whether the coordination-taxonomy conflict creates independent scope. | none |
| `PRODUCT-ADR-SPEC-012` | accepted | Shared-writer serialization and one final integrated review. | T07 is the current review owner. | Future writer order and the review Task route are not yet fixed. | none |
| `PRODUCT-ADR-SPEC-013` | accepted | Finding-specific repair Tasks are delayed until findings exist. | W018 preserves the conditional branch. | No conflict found. | none |
| `PRODUCT-ADR-SPEC-014` | accepted | Append-only reconvergence and closure write limits. | T01 through T06 remain historical; T11 and T12 are new Tasks. | The wording creates an unresolved question about resuming T07 versus creating a new review Task. | none |

All inspected ADRs are accepted.
No inspected ADR declares a superseding ADR.
No supersession was performed by this Investigation.

#### Affected Specifications

| spec | observed state | proposed relation |
|---|---|---|
| `requirement_authoring` | Contains the D-012 identity boundary and prohibits persisted Requirement `work_items`. | W018 projection is consistent. The source Requirements contain stale metadata. |
| `work_item_authoring` | Contains continuation, shared-writer, integrated-review, conditional finding, and append-only rules. | Consistent with W018 identity and the intended review route. |
| `task_authoring` | Contains completed-record and reconvergence rules. Its coordination type row still requires child Work Items. | Material internal conflict with its broader coordination rules and T09. |
| `artifact_boundary` | Separates checkpoints, ADR routing, authoring, review, and closure. | Consistent with D-019, D-021, and D-023. |
| `artifact_responsibility_matrix` | Preserves completed records and limits synchronization. | Consistent with D-019, D-022, and D-023. |

#### W018 identity and completion boundary

W018 still resolves the same design-convergence successor topic.
The Goal and Completion Condition remain coherent.
The added Investigation, reconciliation, and routing revalidation remain inside the same design closure boundary.

A split is not proven by current Evidence.
The coordination-taxonomy conflict may become independently completable if T11 routes it outside W018.

#### Task ownership and dependency state

Current owned responsibilities:

- T10 owns formal Investigation.
- T11 owns reconciliation judgment.
- T12 owns ADR-routing revalidation.
- T07 owns the integrated review and remains blocked.

Current gaps:

- No incomplete coordination Task owns graph changes discovered by T11 or T12.
- T08 remains reserved and unmaterialized.
- T07 cannot yet depend on future conditional writers.

No obsolete existing dependency was identified.
No current dependency cycle was identified.

#### Successor skill projection

The successor skill projects the decided end-to-end workflow and phase boundaries.
The nine active files are present.

The skill uses the broad graph-coordination meaning accepted by ADR-005.
That meaning conflicts with the narrower coordination type outcome in ADR-004 and `task_authoring`.

#### Prompt activation and old skill state

`prompt_chappy.md` references the successor workflow and companions as active authority.
No active reference to `skills/design-decision-workflow/` was found in the prompt.
The old repository path does not exist.

The ignored retired copy under `memory/` was not read.
The retired copy is not treated as active authority.

#### Existing authority coverage

| decisions | current durable authority |
|---|---|
| D-001, D-004, D-005 | `PRODUCT-ADR-SPEC-009` |
| D-006, D-011 | `PRODUCT-ADR-SPEC-010` |
| D-007, D-008 | `PRODUCT-ADR-SPEC-004` and `PRODUCT-ADR-SPEC-005`, subject to the coordination conflict |
| D-009, D-014, D-018 | `PRODUCT-ADR-SPEC-005`, subject to ADR-004 coordination wording |
| D-010, D-017, D-019, D-022, D-023 | `PRODUCT-ADR-SPEC-014` and `PRODUCT-ADR-SPEC-006` where applicable |
| D-012, D-013 | `PRODUCT-ADR-SPEC-011` |
| D-015, D-016 | `PRODUCT-ADR-SPEC-012` |
| D-020 | `PRODUCT-ADR-SPEC-013` |
| D-021 | `PRODUCT-ADR-SPEC-006` |
| D-002, D-003 | Direct skill activation and repository transition Evidence |

### Material conflicts

| ID | conflict | affected decisions | Evidence | downstream effect |
|---|---|---|---|---|
| MC-001 | Coordination has incompatible current contracts. ADR-004 and the Task type table require child Work Item decomposition. ADR-005, W018, T09, and the skill use coordination for general Task-graph repair. | D-006, D-009, D-010, D-017, D-020, D-022 | ADR-004 coordination row; ADR-005 coordination section; `task-authoring.md` Task type table; T09 `## Work`; `graph-coordination.md` | T11 cannot safely accept the current coordination route without a precedence or repair decision. |
| MC-002 | Mandatory Investigation and pre-authoring convergence occurred after T03 through T06 completed. | D-006, D-007, D-018, D-021 | W018 `## Task flow`; `design-authoring.md`, `## Preconditions`; T03 through T06 status; T09 route | T11 must decide whether existing outputs remain sufficient or require new authoring. |
| MC-003 | No incomplete coordination Task owns graph changes or Task materialization produced by T11 or T12. | D-008, D-009, D-015, D-016, D-021, D-022 | T11 and T12 prohibited work; W018 conditional authoring branch; T09 status `done` | Any required writer, dependency, review-route, or Work Item change lacks an execution owner. |
| MC-004 | Closure synchronization is required, but T08 remains unmaterialized and no materialization route is assigned. | D-005, D-023 | W018 `## Task flow` and `## Task Candidates`; closure skill Task contract | A future `PASS` cannot proceed directly to a concrete synchronization Task. |
| MC-005 | The source Requirements persist reciprocal `work_items` fields prohibited by REQ-006 and `requirement_authoring`. | cross-cutting source relation state | REQ-005 and REQ-006 metadata; REQ-006 `## Requirement`; `requirement-authoring.md`, reverse relation rules | The stale relation state requires a named migration or correction owner outside this Investigation. |

Material conflict count: 5.

### Graph-change candidates

| ID | candidate | category | evidence | candidate next owner |
|---|---|---|---|---|
| G-001 | Materialize a new coordination Task after T11 or T12 when graph or authoring changes are required. | missing owner | T11 and T12 prohibit graph edits; T09 is complete. | T11 identifies the need; a new coordination owner performs it. |
| G-002 | Materialize T08 or another exact closure-synchronization Task before the accepted review route needs it. | missing owner | W018 reserves T08 but does not list a concrete Task. | T11 decides timing; coordination materializes the Task. |
| G-003 | Add exact future authoring Tasks when T11 or T12 identifies changed ADR, Requirement, Work Item, Specification, or skill targets. | future authoring Task | W018 conditional branch; T12 Done condition. | conditional coordination |
| G-004 | Amend T07 dependencies and review boundary after every future writer is materialized. | missing dependency and review release condition | T07 depends only on T05, T06, and T12; T07 Evidence requires later writers. | conditional coordination |
| G-005 | Decide whether T07 resumes or a new integrated review Task is required. | review route | D-014 permits incomplete Task amendment; D-022 may require a new review chain. | T11, then coordination when needed |
| G-006 | Repair the canonical coordination contract through an ADR and Specification route when T11 confirms MC-001. | future canonical authoring | ADR-004 and `task_authoring` conflict with ADR-005 and current usage. | T11, T12, then coordination and authoring |
| G-007 | Split a separate Work Item only when the coordination-taxonomy repair or source-relation migration has an independent completion judgment. | Work Item split candidate | ADR-011 and `work-item-authoring` identity rules | T11 reconciliation |

Additional graph observations:

- No obsolete dependency was identified.
- T07 remains blocked by the accepted prerequisite route.
- No completed Task needs substantive modification.
- Any revised work must use new Tasks.

Graph-change candidate count: 7.

### Shared-writer candidates

| ID | exact artifact or section | candidate writer responsibility | semantic dependency | section separation | candidate writer order | concurrent write risk |
|---|---|---|---|---|---|---|
| SW-001 | `PRODUCT-ADR-SPEC-006`, full ADR | Historical T03 authoring and possible future ADR authoring | Future text must preserve the checkpoint and non-writeback choice. | Not safely separable without one ADR-level acceptance judgment. | T03 output, then T12 route, then a new authoring Task. | A concurrent amendment could conceal a reversal or lose historical clarification. |
| SW-002 | `PRODUCT-ADR-SPEC-009` through `PRODUCT-ADR-SPEC-014` | Historical T04 authoring and possible future ADR authoring | Future changes depend on T11 resolution and T12 routing. | Individual ADRs are separable only when T12 assigns independent boundaries. | T04 output, then T12 route, then routed authoring. | Parallel changes could create incompatible ADR boundaries or consequences. |
| SW-003 | Nine successor skill files and `prompt_chappy.md` activation section | Historical T05 authoring and possible future workflow-support authoring | Future skill text must follow repaired ADR and Specification authority. | Files are physically separable, but activation and phase authority share one semantic release. | ADR and Specification repairs first, skill authoring next, activation last. | Concurrent edits could activate mixed old and new authority. |
| SW-004 | Five canonical Specification targets, especially `task_authoring` Task type contract | Historical T06 authoring and possible future Specification authoring | Specification repair depends on T11 and final ADR routing. | Separate files may be split when decisions are independent. The coordination row and adjacent rules require one combined consistency judgment. | Final ADR authority, then Specification writer order, then integrated review. | Concurrent writes could preserve both incompatible coordination definitions. |
| SW-005 | W018 `## Task flow` and `## Task Candidates`; T07 metadata and `### Reconvergence Evidence` | Historical T09 coordination and possible future coordination | Future graph text depends on materialized Tasks and exact writer dependencies. | W018 and T07 must remain mutually consistent. | Materialize Tasks first, then update W018 and T07 in one coordination boundary. | Concurrent graph edits could release review before required writers complete. |

Shared-writer candidate count: 5.

### Uncertainty and missing Evidence

| ID | unknown fact | missing Evidence | impact | named next owner | downstream blocker |
|---|---|---|---|---|---|
| U-001 | Which coordination contract is intended as current durable authority. | No explicit amendment, supersession, or precedence statement reconciles ADR-004 with ADR-005 and current Specifications. | T09 and future graph coordination cannot be classified safely. | T11 reconciliation | yes |
| U-002 | Whether a post-authoring Investigation and revalidation can preserve T03 through T06 as sufficient outputs. | No accepted rule states when late mandatory Investigation may validate already completed authoring. | Determines whether new ADR, Specification, or skill authoring is required. | T11, then T12 | yes |
| U-003 | Whether D-022 requires a new integrated review Task when the existing review Task is blocked and has no semantic verdict. | The decision and ADR do not state this exact incomplete-review case. | Determines T07 reuse and review release structure. | T11 reconciliation | yes |
| U-004 | Whether the source Requirement `work_items` fields are pending an accepted migration outside W018. | The responsible migration Task or Work Item is outside the required read boundary. | Determines whether MC-005 is an active W018 blocker or external stale state. | existing-record migration owner, named by T11 if material | no, unless T11 expands W018 scope |
| U-005 | Which artifacts require another writer after T11 and T12. | Reconciliation and routing results do not yet exist. | T07 dependencies and writer order cannot be finalized. | T11 and T12 | yes |

Uncertainty count: 5.

## Cross-cutting observations

- The current repository activation state is coherent. The successor is active and the old repository path is absent.
- The durable workflow ADR set is accepted and non-superseded.
- The main semantic defect is not missing ADR coverage. The main defect is inconsistent coordination ownership across current authorities.
- The main sequence defect is late Investigation after completed authoring.
- T09 repaired the first missing owner but did not create the owner for graph changes discovered by T11 or T12.
- W018 still appears to have one coherent design-closure identity.
- The current graph can preserve completed T01 through T06 only through new successor Tasks.
- No evidence supports rewriting completed Tasks.

## Follow-up judgment candidates

| candidate ID | affected decision IDs | exact question | materially different options | supporting Evidence | recommended next owner |
|---|---|---|---|---|---|
| J-001 | D-006, D-009, D-010, D-017, D-020, D-022 | Which coordination outcome and completion judgment is canonical? | Preserve ADR-004's child-Work-Item-only contract; amend ADR-004 and `task_authoring` to broad graph coordination; split broad graph coordination into another Task type or boundary. | MC-001; ADR-004; ADR-005; T09; `graph-coordination.md` | T11 reconciliation |
| J-002 | D-006, D-007, D-018, D-021 | Are T03 through T06 still acceptable after the mandatory Investigation occurred late? | Preserve outputs after T11/T12 revalidation; require bounded new authoring; reconsider one or more prior decisions. | MC-002; T03 through T06; design-authoring preconditions | T11 reconciliation, then T12 routing |
| J-003 | D-012, D-013 | Does W018 remain one Work Item after resolving MC-001 and MC-005? | Continue W018; split coordination-taxonomy repair; defer source-relation migration to an external owner. | ADR-011; W018 Goal and Completion Condition; MC-001 and MC-005 | T11 reconciliation |
| J-004 | D-008, D-009, D-015, D-016, D-021, D-022 | What graph owner executes changes identified by T11 or T12? | Create one post-routing coordination Task; create separate coordination Tasks by completion boundary; conclude no graph change is required. | MC-003; G-001 through G-004 | T11 reconciliation |
| J-005 | D-005, D-023 | When and how is closure synchronization materialized? | Materialize T08 before review resumes; materialize after `PASS`; materialize after finding closure on the revision route. | MC-004; W018 Task flow; closure skill | T11 reconciliation |
| J-006 | D-014, D-016, D-022 | Does T07 resume, or must a new integrated review Task evaluate the revised combined state? | Resume the incomplete T07 after dependency repair; preserve T07 as historical prerequisite Evidence and create a new review Task. | D-014; D-022; T07 status and Evidence | T11 reconciliation |
| J-007 | D-012 and cross-cutting source relation state | Does W018 own any correction of Requirement reciprocal metadata? | Treat it as external migration; add bounded W018 authoring only if directly required; split an independent Work Item. | MC-005; REQ-006; `requirement_authoring` | T11 reconciliation or named migration owner |

Follow-up judgment candidate count: 7.

No option in this table is adopted by this Investigation.

## Recommendation

T11 should evaluate J-001 through J-007 before T12 revalidates ADR routing.

Preserving W018 as one Work Item appears preferable while its Goal and Completion Condition remain unchanged.
A separate Work Item is likely appropriate only for independently completable coordination-taxonomy or migration work.

A new post-T11 or post-T12 coordination Task appears likely when any graph or authoring change is accepted.
The existing completed Tasks should remain unchanged.

T12 should evaluate whether the accepted ADR set and T03 through T06 outputs remain sufficient after T11 resolves the conflicts.
Integrated review should remain blocked until the exact writer and review route is persisted.

## Follow-up artifact candidates

- A post-Investigation reconciliation decision ledger in T11.
- A complete ADR-routing revalidation in T12.
- A new coordination Task when T11 or T12 requires graph changes.
- Bounded ADR, Requirement, Work Item, Specification, or skill authoring Tasks when exact targets change.
- A materialized closure-synchronization Task.
- A new integrated review Task when T11 selects that route.
- A separate Work Item when an independently completable taxonomy or migration boundary is confirmed.

## Open questions

- Which current authority owns the generic coordination contract?
- Can late Investigation and revalidation preserve already completed authoring without a new writer?
- Does D-022 require a new review Task for an incomplete blocked review route?
- Who owns the existing Requirement reciprocal-field migration?
- Which exact writer set and dependency chain will T12 require?
