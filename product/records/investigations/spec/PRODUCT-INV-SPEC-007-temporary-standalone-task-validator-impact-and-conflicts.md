# PRODUCT-INV-SPEC-007: Temporary standalone Task validator impact and conflicts

- **status**: concluded
- **date**: 2026-07-01
- **trigger**: W019 mandatory impact and conflict Investigation
- **scope**: Impact of the accepted temporary standalone semantic Task validator decisions on PRODUCT authority, canonical targets, implementation handoff, workflow graph, shared writers, and unresolved conflicts.
- **non_scope**: DRMCP integration, concrete runtime or schema design, checklist wording or storage, canonical authoring, implementation, review, and lifecycle synchronization.
- **source_refs**:
  - PRODUCT-WORK-SPEC-019
  - PRODUCT-REQ-SPEC-007
  - PRODUCT-REQ-SPEC-005
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
  - spec:product.design_records.authoring_standards.task_authoring
- **follow_up_candidates**:
  - reconciliation decision for validator ownership and canonical target boundaries
  - graph coordination for downstream Task ownership and writer order
  - ADR routing for the standalone semantic validator decision set
  - new PRODUCT standalone validator Contract Specification

## Investigation scope

This Investigation evaluates T01 D-001 through D-012 and T03 D-001.

The Investigation compares those decisions with current PRODUCT authority.
It identifies canonical target candidates, implementation handoff boundaries, graph changes, shared writers, conflicts, and uncertainty.

## Out of scope

- DRMCP integration or changes to current DRMCP artifacts.
- Concrete model, provider, runtime, MCP schema, checklist wording, or checklist storage.
- ADR routing outcomes or ADR authoring.
- Requirement, Specification, Work Item, or Task graph authoring.
- Temporary-tool implementation, review, correction, closure, stage, or commit.

## Background

`PRODUCT-REQ-SPEC-007` requires read-only semantic validation for one Task record.
The Requirement limits the compliance claim to Task responsibility boundaries.

T01 fixed the criterion, rationale, aggregation, and failure-boundary contract.
T03 later fixed the immediate product as a temporary standalone tool.
T03 also excluded current DRMCP integration from W019.

`PRODUCT-REQ-SPEC-005`, `PRODUCT-ADR-SPEC-004`, `PRODUCT-ADR-SPEC-005`, and the Task authoring Specification define the evaluated Task responsibility model.

## What was investigated

- T01 D-001 through D-012.
- T03 D-001.
- `PRODUCT-REQ-SPEC-007` and `PRODUCT-REQ-SPEC-005`.
- `PRODUCT-ADR-SPEC-004` and `PRODUCT-ADR-SPEC-005`.
- `spec:product.design_records.authoring_standards.task_authoring`.
- Scoped PRODUCT Specification candidates containing the allowed search terms.
- W019 downstream graph, implementation handoff, and writer boundaries.

## Findings

### Decision impact inventory

| decision | affected authority | observed state | relation candidate | evidence | required judgment | uncertainty |
|---|---|---|---|---|---|---|
| T01 D-001 | PRODUCT-REQ-SPEC-007 | The Requirement already requires one overall result and one result per criterion. The decision fixes both as binary machine-readable judgments. | `consistent_refinement` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-001; `PRODUCT-REQ-SPEC-007`, `## Required Outcome` | ADR routing and canonical authoring | none |
| T01 D-002 | PRODUCT-REQ-SPEC-007 | The Requirement already requires Task-local evidence or a concise reason for each criterion judgment. The decision applies the rule to true and false results. | `consistent_refinement` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-002; `PRODUCT-REQ-SPEC-007`, `## Required Outcome` | canonical authoring | none |
| T01 D-003 | PRODUCT-REQ-SPEC-007 | The Requirement limits evidence to the Task record and treats missing Task-local evidence as non-compliance. | `consistent_refinement` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-003; `PRODUCT-REQ-SPEC-007`, `## Requirement` and `## Required Outcome` | canonical authoring | none |
| T01 D-004 | spec:product.design_records.authoring_standards.task_authoring | The Specification defines a required scalar `task_type` and closed type set. It does not define automatic semantic checklist selection. | `consistent_refinement` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-004; `task-authoring.md`, `### Metadata schema` and `#### Task type contract` | reconciliation decision for ownership wording; canonical authoring | Whether current DRMCP validation ownership includes this semantic validator. |
| T01 D-005 | PRODUCT-ADR-SPEC-004; PRODUCT-ADR-SPEC-005 | The ADRs define common single-responsibility rules and type-specific outcomes. No authority defines their composition as an applied checklist. | `consistent_refinement` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-005; `PRODUCT-ADR-SPEC-004`, `## Decision`; `PRODUCT-ADR-SPEC-005`, `## Decision` | ADR routing and canonical authoring | Exact criterion decomposition remains outside W019. |
| T01 D-006 | PRODUCT-REQ-SPEC-007 | The semantic linter shape is compatible with the Requirement. The T01 row assigns orchestration to “the MCP,” while T03 later fixes a temporary standalone tool. | `stale_representation` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-006; `PRODUCT-TASK-SPEC-019-03`, `### Decision inventory`, D-001 and `### Corrected boundary` | none for the historical Task; downstream authoring must use the T03 boundary | The final standalone tool name and app-local placement are unresolved. |
| T01 D-007 | PRODUCT-REQ-SPEC-007 | The Requirement requires an overall result but does not define aggregation. Logical AND gives deterministic aggregation from criterion judgments. | `consistent_refinement` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-007; `PRODUCT-REQ-SPEC-007`, `## Required Outcome` | ADR routing and canonical authoring | none |
| T01 D-008 | PRODUCT-REQ-SPEC-007 | The Requirement requires evidence or a reason. The decision adds concise reasons, Task section references, optional excerpts, and no line numbers. | `consistent_refinement` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-008; `PRODUCT-REQ-SPEC-007`, `## Evidence` and `## Required Outcome` | canonical authoring | Exact response field names remain outside this Investigation. |
| T01 D-009 | PRODUCT-REQ-SPEC-007; spec:product.design_records.authoring_standards.task_authoring | The Requirement treats missing semantic evidence as non-compliance. The Task Specification defines required `task_type` and required sections. The decision separates unreadable or invalid input from evaluable missing content. | `consistent_refinement` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-009; `PRODUCT-REQ-SPEC-007`, `## Required Outcome`; `task-authoring.md`, `### Metadata schema` and `### Status lifecycle` | ADR routing and canonical authoring | Structural parse ownership must be separated from semantic evaluation ownership. |
| T01 D-010 | PRODUCT-REQ-SPEC-007 | The Requirement excludes retry policy and error taxonomy. The decision only separates execution failure from semantic non-compliance and does not fix a concrete taxonomy. | `consistent_refinement` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-010; `PRODUCT-REQ-SPEC-007`, `## Explicitly Excluded Scope` and `## Boundary` | ADR routing and canonical authoring | Exact execution-error contract remains deferred. |
| T01 D-011 | PRODUCT-REQ-SPEC-007 | No current authority requires checklist identity or stable criterion IDs. The decision removes that result burden while retaining criterion judgments and reasons. | `consistent_refinement` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-011; `PRODUCT-REQ-SPEC-007`, `## Required Outcome` | ADR routing confirmation | Whether implementation needs internal checklist identity without exposing it in results. |
| T01 D-012 | PRODUCT-WORK-SPEC-019 | W019 requires use during authoring or before release. Exact calling-workflow owners and release conditions are not materialized after T05. | `workflow_graph_drift` | `PRODUCT-TASK-SPEC-019-01`, `### Decision inventory`, D-012; `PRODUCT-WORK-SPEC-019`, `## Task flow` and `## Task Candidates` | graph coordination | Exact release gate and calling workflow remain unresolved. |
| T03 D-001 | spec:product.design_records.authoring_standards.task_authoring | T03 fixes a standalone semantic validator. The Task Specification broadly assigns downstream validation and tool projections to DRMCP. The current wording does not distinguish structural contract validation from semantic responsibility validation. | `semantic_conflict` | `PRODUCT-TASK-SPEC-019-03`, `### Decision inventory`, D-001 and `### Corrected boundary`; `task-authoring.md`, `### Metadata schema` and final PRODUCT/DRMCP ownership statements | reconciliation decision | The intended scope of “DRMCP owns downstream validation” is not explicit. |

Decision relation candidate counts:

| relation candidate | count |
|---|---:|
| `consistent_refinement` | 10 |
| `stale_representation` | 1 |
| `semantic_conflict` | 1 |
| `workflow_graph_drift` | 1 |

### Requirement impact

| authority | impact |
|---|---|
| PRODUCT-REQ-SPEC-007 | Covers one-Task scope, Task-local evidence, binary overall and criterion results, read-only behavior, and authoring or release use. It does not own exact criterion composition, deterministic aggregation, preconditions, or the standalone product shape. |
| PRODUCT-REQ-SPEC-005 | Defines the evaluated typed single-responsibility contract. It explicitly excludes validator implementation, checklist format, and diagnostics. No Requirement change appears necessary unless reconciliation changes the PRODUCT and DRMCP ownership boundary. |

### ADR coverage and conflict candidates

| ADR | current coverage | standalone-validator coverage | conflict candidate | routing re-evaluation candidates |
|---|---|---|---|---|
| PRODUCT-ADR-SPEC-004 | Covers the closed `task_type` taxonomy and each type's primary outcome. | Does not cover checklist selection, result aggregation, rationale shape, failure separation, or standalone orchestration. | None found in the decision content. Its Consequences broadly reserve downstream validation to DRMCP. | T01 D-004 through D-007, D-009 through D-011, and T03 D-001. |
| PRODUCT-ADR-SPEC-005 | Covers single responsibility, section alignment, independent completion, and stop boundaries. | Does not cover semantic evaluator behavior or the temporary standalone product shape. | No direct contradiction with read-only semantic evaluation. | T01 D-005, D-008, D-009, D-012, and T03 D-001. |

The ADRs cover the evaluated Task semantics.
They do not cover the validator-specific durable choices.
The routing owner must not infer `create`, `amend`, `reuse`, or `supersede` from this Investigation.

### Task-authoring Specification relationship

The Task authoring Specification supplies the canonical input semantics:

- required scalar `task_type`;
- closed type set;
- one primary outcome and completion judgment;
- type-specific prohibited overlaps;
- common section alignment;
- adjacent responsibility boundaries.

The Specification does not define semantic checklist composition or validator results.
A validator contract can therefore refine the authoring workflow without duplicating the Task contract.

Two current wording conflicts affect the route:

1. The Specification broadly assigns downstream validation and tool projections to DRMCP.
2. The `coordination` primary outcome is limited to child Work Item coordination, while later sections allow Task graph changes.

`PRODUCT-TASK-SPEC-018-11` J-001 already owns and decides the semantic disposition for the second conflict.
That decision adds `work_item_decomposition` for parent-to-child Work Item decomposition and keeps `coordination` for Task-graph changes.
W019 must consume that accepted decision instead of reopening it.

### Downstream PRODUCT canonical target candidates

| candidate | role | current state | required judgment |
|---|---|---|---|
| New PRODUCT standalone validator Contract Specification under `spec:product.design_records.authoring_standards` | Own criterion composition, result semantics, aggregation, preconditions, execution-failure separation, and standalone-tool boundary. | No dedicated current Specification was found in the scoped search. | Select exact semantic ref, path, kind, and section boundary. |
| spec:product.design_records.authoring_standards.task_authoring | Retain the evaluated Task contract and add only a narrow authoring-workflow relation if needed. | Already owns Task semantics. | Decide whether any normative cross-reference is required. |
| spec:product.design_records.authoring_standards | Register a new child Specification if the validator contract becomes a sibling topic. | Existing index has no validator topic. | Coordinate index authoring after exact target selection. |
| PRODUCT-REQ-SPEC-007 | Remain the accepted need and compliance-claim boundary. | Already compatible with most decisions. | Amend only if reconciliation changes ownership or scope. |

A dedicated Contract Specification appears preferable to placing tool behavior inside `task_authoring`.
The next decision owner should evaluate that boundary.

### Downstream temporary-tool implementation handoff

The implementation handoff begins after canonical PRODUCT authoring and accepted integrated review.

The handoff should receive:

- one-Task invocation boundary;
- Task-only semantic evidence boundary;
- common plus type-specific criterion composition;
- criterion-level binary judgments and reasons;
- logical-AND aggregation;
- structural precondition boundary;
- execution-failure separation;
- read-only behavior;
- no automatic correction or Task-splitting proposal;
- explicit exclusion of current DRMCP integration.

A separate downstream implementation Work Item appears necessary.
The exact app namespace, source path, runtime, and tool schema remain unresolved.

### Material conflict inventory

| conflict ID | affected decisions | affected artifacts | observed contradiction or absence | relation candidate | evidence | downstream effect | required next owner |
|---|---|---|---|---|---|---|---|
| MC-001 | T01 D-004, D-006, D-007, D-009, D-010, D-012; T03 D-001 | PRODUCT-REQ-SPEC-007; spec:product.design_records.authoring_standards.task_authoring | T03 requires a standalone semantic validator. The Task Specification broadly assigns downstream validation and tool projections to DRMCP. The wording does not separate structural contract validation from semantic Task evaluation. | `semantic_conflict` | `PRODUCT-TASK-SPEC-019-03`, `### Corrected boundary`; `task-authoring.md`, `### Metadata schema` and PRODUCT/DRMCP ownership statements | Canonical authoring cannot safely select the validator owner or target boundary. | reconciliation decision owner |
| MC-002 | T01 D-012 | PRODUCT-WORK-SPEC-019; spec:product.design_records.authoring_standards.task_authoring; PRODUCT-TASK-SPEC-018-11 | W019 requires downstream graph coordination. The Task type table limits `coordination` to parent Work Item and child Work Item boundaries. Later Specification text allows Task, dependency, blocker, and release-order changes. W018 T11 J-001 already decides the semantic split. | `semantic_conflict` and `workflow_graph_drift` | `PRODUCT-WORK-SPEC-019`, `## Task flow` and `## Task Candidates`; `task-authoring.md`, `#### Task type contract`, `#### Adjacent responsibility boundaries`, and `#### Task continuation and reconvergence routes`; `PRODUCT-TASK-SPEC-018-11`, `### J-001 decision` | Canonical wording remains inconsistent until the W018 authoring route completes. W019 may use `coordination` for graph changes but must serialize dependent authoring and review after the W018 repair. | W018 T11 J-001 as accepted decision authority; W018 T12 and downstream authoring for canonical repair; W019 graph coordination consumes the result |

T01 D-006 also contains stale “MCP” wording.
T03 already supplies the corrected standalone boundary.
The completed T01 Task should remain historical Evidence and should not be substantively rewritten.

### Graph-change candidates

| candidate ID | observed need | candidate change | dependency or release effect | uncertainty |
|---|---|---|---|---|
| GC-001 | MC-001 and W019-specific target and release-gate questions require judgment before authoring. | Materialize one reconciliation decision Task for validator ownership, canonical target, and calling-workflow boundaries. | Blocks W019 graph coordination and ADR routing. | The decision may need one or two bounded ledgers. |
| GC-002 | W019 lists conditional downstream Tasks but does not materialize them after T05. | Materialize the exact `coordination` owner, ADR-routing, conditional ADR, Specification authoring, integrated review, and closure Tasks. | Must follow the W019 reconciliation result and consume W018 T11 J-001. | The exact dependency on W018 T12 and its downstream canonical repair remains unmaterialized. |
| GC-003 | D-012 requires use during authoring and before release. | Define the calling-workflow owner and release condition without moving enforcement into the validator. | Must precede implementation handoff. | Whether one workflow integration Task or several owners are needed. |
| GC-004 | Production implementation is outside W019. | Create a separate implementation Work Item after accepted canonical authoring and review. | Must not block W019 design closure unless the Work Item relation is required by closure policy. | Exact app namespace and source owner are unknown. |

No completed Task requires substantive rewrite.
T01 and T03 remain historical decision checkpoints.
W019 does not appear to require a design-topic split.
The implementation handoff remains a separate Work Item candidate.

Direct ADR routing does not appear ready while MC-001 and the W019-specific target and release-gate judgments remain unresolved.
MC-002 does not require another W019 semantic decision.

### Shared-writer candidates

| exact artifact or section | candidate writer responsibilities | semantic dependency | section separation feasibility | candidate writer order | concurrent-write risk |
|---|---|---|---|---|---|
| spec:product.design_records.authoring_standards.task_authoring, ownership statements and coordination contract | W018 authoring repairs the accepted `work_item_decomposition` and `coordination` split; W019 authoring clarifies validator ownership and adds any narrow validator relation. | W018 T11 J-001 and T12 routing must precede W019 projection into the same Specification. | Partial. The topics use different sections, but both affect Task type and downstream ownership meaning. | W018 canonical repair first, then W019 validator projection, then integrated review. | High. Concurrent edits could preserve contradictory ownership or completion boundaries. |
| spec:product.design_records.authoring_standards, `## Topics` | Register the validator Contract Specification; preserve concurrent authoring-standard topic changes. | Exact child semantic ref must exist first. | Yes, by topic row. | Child Specification before index registration. | Medium. Concurrent row insertion can cause stale or duplicate routing. |
| PRODUCT-WORK-SPEC-019, `## Task flow`, `## Task Candidates`, `## Completion Condition`, and `## Evidence` | Graph coordination materializes downstream ownership; closure synchronization later records accepted completion and relations. | Closure depends on the final graph and accepted review route. | Partial. Sections differ, but closure Evidence depends on graph state. | Graph coordination before closure synchronization. | High. Concurrent writes can overwrite task inventory or close an incomplete route. |

The candidate order places integrated review after the last shared writer.
Writer order remains a coordination responsibility.

### Uncertainty and missing Evidence

| ID | uncertainty or missing Evidence | next owner |
|---|---|---|
| U-001 | The exact semantic ref, path, and document kind for the validator Contract Specification are not selected. | reconciliation decision owner |
| U-002 | The Task Specification does not state whether DRMCP validation ownership excludes independent semantic authoring checks. | reconciliation decision owner |
| U-003 | The exact W018 authoring Task and completion point for canonicalizing the already decided coordination split are not yet materialized. | W018 T12 and downstream graph coordination owner |
| U-004 | The exact app namespace, source owner, and Work Item identity for temporary-tool implementation are not defined. | downstream implementation coordination owner |
| U-005 | The calling workflow and release gate that consume D-012 are not materialized. | graph coordination owner |

## Cross-cutting observations

- The accepted Requirement and Task responsibility model are compatible with criterion-level semantic validation.
- The largest authority issue is ownership wording, not criterion semantics.
- T03 corrects the immediate product boundary without reopening T01.
- The current graph intentionally stops at Investigation, but downstream responsibilities now need materialization.
- W018 T11 J-001 already decides the Task-type split: `coordination` owns Task-graph changes and `work_item_decomposition` owns parent-to-child Work Item decomposition.
- The validator contract and the evaluated Task contract should remain separate semantic layers.
- Current DRMCP artifacts are not required as Evidence for this Investigation.

## Follow-up judgment candidates

| ID | candidate judgment | reason |
|---|---|---|
| J-001 | Distinguish structural DRMCP validation ownership from standalone semantic Task evaluation ownership. | MC-001 blocks canonical target selection. |
| J-002 | Select a dedicated validator Contract Specification or an extension of `task_authoring`. | No dedicated canonical target exists. |
| J-003 | Confirm T03 as the terminal product-boundary input and retain T01 D-006 as historical wording. | Completed decision Tasks should not be substantively rewritten. |
| J-004 | Consume W018 T11 J-001 as existing authority and define the dependency and serialization boundary for W019 graph materialization. | MC-002 needs canonical repair and writer ordering, not another W019 semantic decision. |
| J-005 | Define the D-012 calling-workflow and release-gate ownership. | The validator reports compliance but does not enforce workflow policy. |
| J-006 | Route the final decision set across existing and new ADR boundaries. | ADR-004 and ADR-005 do not cover validator-specific durable choices. |

## Recommendation

A reconciliation decision appears preferable before graph coordination or ADR routing.
The decision should resolve MC-001 and the W019-specific canonical-target and release-gate questions without reopening T01 or T03.
It must consume W018 T11 J-001 for MC-002 instead of deciding the Task-type split again.

A dedicated PRODUCT Contract Specification is likely appropriate for validator behavior.
`task_authoring` should remain the canonical source for the evaluated Task contract.

Graph coordination should then materialize the downstream writer, review, closure, and implementation-handoff route.
The graph must serialize W019 authoring after the W018 canonical coordination repair when both write `task_authoring`.

## Follow-up artifact candidates

- One reconciliation decision Task for MC-001 and W019-specific canonical-target and release-gate judgments. Do not reopen the MC-002 Task-type decision.
- One downstream `coordination` Task that consumes W018 T11 J-001 and serializes any shared `task_authoring` writes after the W018 canonical repair.
- One ADR-routing Task covering all terminal decisions.
- Conditional ADR authoring Tasks based on the routing result.
- One new PRODUCT standalone validator Contract Specification.
- A narrow `task_authoring` update only when reconciliation requires an ownership clarification or workflow relation.
- One integrated independent review Task after the final writer.
- One closure synchronization Task after accepted review.
- One separate temporary-tool implementation Work Item after W019 design closure.

## Open questions

1. Does current DRMCP validation ownership cover only structural Design Record contract validation?
2. Which W018 authoring Task will canonicalize the already decided coordination split, and what dependency must W019 retain?
3. What exact PRODUCT semantic ref and path should own the standalone validator Contract Specification?
4. Which workflow owns the pre-release gate that consumes semantic compliance?
5. Which app namespace owns the temporary standalone implementation?
