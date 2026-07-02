# PRODUCT-TASK-SPEC-021-06: Decide PRODUCT-to-app bootstrap boundary

- **id**: PRODUCT-TASK-SPEC-021-06
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-021-05
- **outputs**:
  - PRODUCT-TASK-SPEC-021-06

## Goal

Produce one bounded decision ledger for correcting W021 ownership and bootstrapping the standalone validator app-local design route.

## Work

- Run one interactive decision loop.
- Treat the user correction as a fixed boundary: PRODUCT owns conceptual design only.
- Decide the current W021 disposition without reopening completed T02 or T03.
- Decide the standalone validator app namespace identity and top-level directory.
- Classify T02 and PRODUCT-INV-SPEC-009 content as PRODUCT-fixed semantics or app-local design input.
- Decide the child app-local design Work Item identity and completion boundary.
- Decide the parent tracking and closure boundary.
- Record one concise reason, canonical target, ADR-routing state, and downstream owner for every terminal item.

This Task must not:

- author namespace profile or app-local Specification content;
- create or split a Work Item;
- change the Task graph;
- perform implementation planning or implementation;
- create an executor prompt;
- perform review, synchronization, stage, or commit work.

### Decision inventory

| ID | Topic | Status | Depends on |
|---|---|---|---|
| D-001 | W021 disposition and corrected PRODUCT completion boundary | `decided` | none |
| D-002 | Standalone validator app namespace, formal name, and top-level directory | `decided` | D-001 |
| D-003 | PRODUCT-fixed semantics versus app-local design inputs | `decided` | D-001 |
| D-004 | Child app-local design Work Item identity and completion boundary | `decided` | D-002, D-003 |
| D-005 | Namespace authoring, child creation, parent execution tracking, and closure route | `decided` | D-004 |

### Decision records

#### D-001: W021 disposition and corrected PRODUCT completion boundary

- **Status**: `decided`
- **Selected outcome**: Continue W021 and rescope it from implementation delivery to the PRODUCT-owned conceptual-design Specification boundary.
- **Completion boundary**: W021 becomes `done` after the PRODUCT conceptual-design Specification is authored and an independent integrated review returns `PASS`.
- **Responsibility boundary**: W021 does not own app-local design, implementation planning, implementation, child design-closure tracking, or a `work_item_execution` Task for the child.
- **Reason**: The current W021 implementation outcome crosses the PRODUCT ownership boundary. The reduced outcome preserves Work Item continuity while closing only reviewed PRODUCT conceptual design.
- **Canonical target**: PRODUCT-WORK-SPEC-021 Goal, Boundary, Completion Condition, and the PRODUCT-owned conceptual-design Specification.
- **ADR route**: `candidate`; final routing remains downstream of T06.
- **Downstream owner**: T07 for graph materialization and later PRODUCT authoring, integrated review, and closure owners.

#### D-002: Standalone validator app namespace, formal name, and top-level directory

- **Status**: `decided`
- **Selected outcome**: Use app namespace `TRV`, formal name `Task Responsibility Validator`, and top-level directory `trv/`.
- **Responsibility boundary**: The namespace identifies the validator application. It does not encode the temporary standalone deployment form, MCP transport, model provider, or future DRMCP integration.
- **Reason**: The short stable identity matches existing Brewprint namespace conventions and avoids binding the app identity to replaceable implementation choices.
- **Canonical target**: `spec:product.brewprint.namespaces.app_namespaces`, `spec:product.brewprint.layout`, and the future `trv/records/` namespace profile.
- **ADR route**: `candidate`; final routing remains downstream of T06.
- **Downstream owner**: T07 for namespace authoring and review routing.

#### D-003: PRODUCT-fixed semantics versus app-local design inputs

- **Status**: `decided`
- **Selected outcome**: PRODUCT fixes only cross-app semantic behavior and workflow invocation policy. TRV owns all concrete interface, runtime, transport, packaging, configuration, and implementation design.
- **PRODUCT-fixed semantics**:
  - one Task per evaluation;
  - automatic checklist selection from declared `task_type`;
  - one boolean result and concise reason per criterion;
  - mechanical logical AND for overall compliance;
  - distinct semantic violation, structural failure, and execution failure outcomes;
  - caller-owned human acceptance or rejection;
  - invocation after Task authoring and after final Evidence;
  - exclusion of current DRMCP integration.
- **App-local design inputs**: T02 choices for MCP, path input, checklist loading, prompt composition, Go, Qwen, Ollama, JSON Schema, retry, timeout, environment variables, packaging, names, launcher, tests, build, and deployment remain historical candidate inputs only.
- **Responsibility boundary**: PRODUCT does not make those app-local choices canonical. TRV may preserve, revise, or reject them through its own decision, ADR, Specification, review, and closure workflow.
- **Reason**: The semantic contract is cross-app policy. Concrete delivery and runtime choices belong to the owning application namespace.
- **Canonical target**: PRODUCT conceptual-design Specification for fixed semantics and later TRV-local Requirement, ADR, and Specification artifacts for concrete design.
- **ADR route**: `candidate`; final routing remains downstream of T06.
- **Downstream owner**: T07 for PRODUCT authoring and child TRV Work Item decomposition routing.

#### D-004: Child app-local design Work Item identity and completion boundary

- **Status**: `decided`
- **Selected outcome**: Create `TRV-WORK-SPEC-001` with title `Define Task Responsibility Validator app-local design`.
- **Primary outcome**: Establish the TRV-local Requirement, decisions, ADRs, and Specifications needed to design the application under the PRODUCT-fixed semantic contract.
- **Completion boundary**: The child Work Item becomes `done` after app-local Requirement establishment, required decisions, required ADR and Specification authoring, integrated independent review `PASS`, and lifecycle closure synchronization.
- **Responsibility boundary**: The child Work Item excludes implementation, executor Tasks, and current DRMCP integration. A separate TRV implementation Work Item follows design closure.
- **Parent relation**: W021 does not wait for or track this child Work Item's completion.
- **Reason**: TRV app-local design has an independent owner and completion judgment. Separating it prevents PRODUCT from owning implementation-specific design or execution.
- **Canonical target**: Future `TRV-WORK-SPEC-001` and its TRV-local Requirement, ADR, and Specification graph.
- **ADR route**: `candidate`; final routing remains downstream of T06.
- **Downstream owner**: T07 for decomposition routing after namespace activation.

#### D-005: Namespace authoring, child creation, parent execution tracking, and closure route

- **Status**: `decided`
- **Selected outcome**: T07 materializes ADR routing, any required ADR authoring, TRV namespace profile and repository-layout authoring, PRODUCT conceptual-design Specification authoring, integrated independent review, and W021 closure synchronization.
- **Successor creation route**: Create `TRV-WORK-SPEC-001` only after the TRV namespace is active and the reviewed PRODUCT conceptual-design Specification is available.
- **Parent tracking boundary**: W021 does not create or use a `work_item_execution` Task and does not wait for `TRV-WORK-SPEC-001` completion.
- **Closure boundary**: W021 may close after PRODUCT authoring is complete, integrated independent review returns `PASS`, and closure synchronization records the accepted PRODUCT result.
- **Responsibility boundary**: `TRV-WORK-SPEC-001` is an independent successor Work Item. It owns its own Task graph, design closure, and later implementation handoff.
- **Reason**: PRODUCT closure depends only on reviewed PRODUCT-owned design. App-local design has a separate completion judgment and must not extend the parent lifecycle.
- **Canonical target**: PRODUCT-WORK-SPEC-021 graph and closure state, PRODUCT namespace and layout profile, PRODUCT conceptual-design Specification, and later `TRV-WORK-SPEC-001`.
- **ADR route**: T07 must materialize one explicit ADR-routing owner before canonical authoring. Required ADRs remain conditional on that routing result.
- **Downstream owner**: T07 for graph coordination; later authoring, review, synchronization, and independent TRV Work Item creation owners materialized by T07.

### Current cursor

- Current item: none.
- Loop state: `decision_complete`.
- Ask exactly one unresolved item at a time during execution.

### Fixed inputs

- PRODUCT owns conceptual and cross-app validator semantics.
- PRODUCT does not own app-local implementation Specifications or implementation execution.
- App namespace bootstrap precedes app-local Design Records.
- App-local design closure precedes a separate implementation Work Item.
- T02 and PRODUCT-INV-SPEC-009 remain historical Evidence.
- T02 app-local choices are not automatically canonical app-local Specification state.
- Current DRMCP integration remains excluded.

## Done condition

- D-001 through D-005 are `decided`, `deferred`, or validly `blocked`.
- W021 disposition and PRODUCT completion boundary are explicit.
- The app namespace identity and directory are explicit or validly blocked.
- PRODUCT and app-local authority are partitioned without ambiguity.
- The child Work Item identity and completion boundary are sufficient for decomposition.
- The downstream ADR routing, authoring, review, decomposition, execution-tracking, and closure route is explicit.
- No authoring, decomposition, graph change, or implementation is performed.

## Verification

- Confirm at most one item is `in_discussion` during execution.
- Confirm every terminal item records a selected outcome and reason.
- Confirm completed T02 and T03 are not substantively changed.
- Confirm no PRODUCT decision claims app-local Specification completion.
- Confirm no namespace profile, child Work Item, or implementation artifact is created.

## Evidence

- The T06 interactive decision session started on 2026-07-02 with D-001 as the only `in_discussion` item.
- DRMCP is non-operational, so this session uses filesystem read and write operations under the accepted fallback policy.
- T05 repaired the graph and created this decision owner.
- PRODUCT-INV-SPEC-009 identified the missing app namespace as an executor-readiness blocker.
- The user clarified that app-local implementation Specifications must precede implementation planning.
- For D-001, the user selected W021 continuation with a reduced PRODUCT conceptual-design outcome and no child design-closure tracking.
- The user fixed the D-001 completion boundary at PRODUCT Specification authoring plus independent integrated review `PASS`.
- D-001 is `decided`.
- The user selected `TRV`, `Task Responsibility Validator`, and `trv/` for D-002.
- D-002 is `decided`.
- The user accepted the PRODUCT-fixed semantic and TRV app-local design partition for D-003.
- D-003 is `decided`.
- The user accepted `TRV-WORK-SPEC-001` and its app-local design closure boundary for D-004.
- D-004 is `decided`.
- The user accepted the independent-successor handoff and PRODUCT-only closure route for D-005.
- D-005 is `decided`.
- D-001 through D-005 are terminal, no item remains `in_discussion`, and the loop state is `decision_complete`.
- Verification confirmed explicit W021 disposition, PRODUCT completion boundary, TRV identity and directory, PRODUCT/TRV authority partition, successor Work Item boundary, ADR routing, authoring, review, closure, and no parent execution tracking.
