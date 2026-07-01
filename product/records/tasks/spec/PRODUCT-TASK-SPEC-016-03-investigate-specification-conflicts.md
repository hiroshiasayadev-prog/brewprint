# PRODUCT-TASK-SPEC-016-03: Investigate Specification conflicts

- **id**: PRODUCT-TASK-SPEC-016-03
- **status**: done
- **date**: 2026-06-30
- **work_item**: PRODUCT-WORK-SPEC-016
- **source_requirement**: PRODUCT-REQ-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-016-02
- **outputs**:
  - PRODUCT-TASK-SPEC-016-03

## Goal

Identify every material conflict between decided REQ-005 semantics and current canonical Specifications.

Produce an exact conflict inventory without selecting dispositions.

## Work

- Compare decided T02 entries with the current Task authoring contract.
- Identify affected sections in Work Item authoring, artifact boundary, responsibility matrix, and traceability only when required.
- Identify shared-file overlap with REQ-006 and W017.
- Separate W016-owned conflicts from W017-owned source-relation conflicts.
- Record writer-order and downstream DRMCP impact candidates.

## Done condition

- Every material Specification conflict has an exact canonical ref and affected section description.
- Every conflict has a declared owner: W016, W017, shared serialization, or downstream DRMCP.
- No conflict disposition or product decision is introduced.
- T04 can ask one conflict decision at a time.

## Verification

- Trace each conflict to a decided T02 entry or accepted Requirement.
- Confirm W017-owned relation semantics remain undecided by W016.
- Confirm no Specification, ADR, source, test, or fixture file changed.

## Evidence

### Investigation scope

- T02 is `done`; D-001 through D-018 are `decided`.
- T02 decisions were treated as authority and were not reopened.
- DRMCP is non-operational under the current agent authoring policy.
- Filesystem access was used for scoped retrieval and authoring.
- No repository-wide traversal was used.
- No Specification, ADR, source, test, fixture, or Task other than T03 was changed.

### Authority read

| authority | use |
|---|---|
| `PRODUCT-REQ-SPEC-005` | Accepted Task responsibility requirement. |
| `PRODUCT-TASK-SPEC-016-02` | Decided D-001 through D-018 contract. |
| `spec:product.design_records.authoring_standards.task_authoring` | Primary canonical Task contract target. |
| `spec:product.design_records.authoring_standards.work_item_authoring` | Coordination and Work Item boundary target. |
| `spec:product.design_records.authoring_standards.artifact_boundary` | Adjacent artifact selection boundary. |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | Canonical artifact ownership boundary. |
| `spec:product.design_records.traceability` and child refs | Current workflow relation contract. |
| `PRODUCT-REQ-SPEC-006` and `PRODUCT-WORK-SPEC-017` | W017 source-relation ownership and writer order. |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Downstream parser contract impact. |

### Conflict inventory

| ID | target Specification | exact section or statement | T02 authority | conflict description | why it matters | candidate resolution boundary | canonical target | ADR candidate | dependency | owner | status |
|---|---|---|---|---|---|---|---|---|---|---|---|
| C-001 | `task_authoring` | `Metadata schema`, `Create`, and `Update` omit `task_type`. | D-001, D-002 | The accepted required field and closed value set cannot be represented by the current Task metadata contract. | A conforming Task cannot persist its primary type. | Add the required scalar field, allowed values, create input, and update rule. Do not define parser diagnostics here. | `spec:product.design_records.authoring_standards.task_authoring` | candidate | T04, then T05-T07 | W016 | conflict |
| C-002 | `task_authoring` | `Kind-specific authoring rules` defines only generic Task rules. | D-003 through D-011 | The current Specification has no owned outcome, completion judgment, or prohibited overlap contract for any accepted Task type. | `task_type` would be only a label without normative semantics. | Add one type-contract matrix or equivalent structured rules for all nine values. | `spec:product.design_records.authoring_standards.task_authoring` | candidate | T04, then T05-T07 | W016 | conflict |
| C-003 | `task_authoring` | `Task granularity`: “Consider splitting” and `SHOULD` guidance. | D-017, D-018; REQ-005 | The current advisory split rule is weaker than exactly one responsibility, outcome, and completion judgment. | Mixed-responsibility Tasks could remain conforming under the current wording. | Replace advisory responsibility wording with a mandatory cohesion rule while retaining effort guidance as advisory. | `spec:product.design_records.authoring_standards.task_authoring` | candidate | T04, then T05-T07 | W016 | conflict |
| C-004 | `task_authoring` | `File shape` body section schema has no `## Implementation contract`. | D-006, D-017 | The accepted implementation-only H2 is absent from the canonical Task shape. | An implementation Task cannot satisfy both contracts without an undocumented extension. | Define conditional heading presence, substantive content, placement, and alignment with the four common sections. | `spec:product.design_records.authoring_standards.task_authoring` | candidate | T04, then T05-T07 | W016 | conflict |
| C-005 | `task_authoring` | `File shape` describes each common H2 independently. | D-017 | No current rule requires Goal, Work, Done condition, and Verification to serve the same primary outcome. | A structurally valid Task may still own multiple completion judgments. | Add a cross-section conformance rule keyed by `task_type`. | `spec:product.design_records.authoring_standards.task_authoring` | candidate | T04, then T05-T07 | W016 | conflict |
| C-006 | `task_authoring` | “When the task produces a design decision, record it in an ADR.” | D-004, D-012 | The current statement implies every decided item requires an ADR and blurs decision versus authoring ownership. | T02 permits conditional ADR routing and prohibits decision Tasks from authoring canonical ADR body content. | Distinguish the decision-ledger checkpoint from conditional downstream ADR authoring. | `spec:product.design_records.authoring_standards.task_authoring` | not_required | T04, then T07 | W016 | conflict |
| C-007 | `task_authoring` | “Do not record design rationale in a task — use an ADR.” | D-004, D-012 | The absolute prohibition conflicts with the accepted decision ledger retaining a concise rationale and canonical target. | The decision loop would lose resumable Evidence before ADR routing. | Distinguish concise workflow rationale from canonical durable rationale. | `spec:product.design_records.authoring_standards.task_authoring` | not_required | T04, then T07 | W016 | conflict |
| C-008 | `task_authoring` | `Task granularity` lists “judgment” and “verification” as independent concerns. | D-013 | The wording does not distinguish supporting command evidence from a separate verification outcome. | Review or implementation Tasks could be split mechanically even when one acceptance boundary remains. | Clarify classification by primary outcome and completion judgment. | `spec:product.design_records.authoring_standards.task_authoring` | candidate | T04, then T05-T07 | W016 | clarification_needed |
| C-009 | `task_authoring` | No rule defines correction completion versus finding closure. | D-008, D-014 | The current Task contract does not prohibit correction and independent finding closure in one Task. | The same writer could repair and close findings. | Add the correction boundary and separate independent review ownership. | `spec:product.design_records.authoring_standards.task_authoring` | candidate | T04, then T05-T07 | W016 | conflict |
| C-010 | `task_authoring` | No rule defines coordination versus synchronization. | D-010, D-011, D-015 | The current Task contract does not separate graph-boundary changes from mechanical accepted-state propagation. | Synchronization could introduce new decomposition or responsibility judgments. | Add separate owned outcomes, stop conditions, and prohibited overlaps. | `spec:product.design_records.authoring_standards.task_authoring` | candidate | T04, then T05-T07 | W016 | conflict |
| C-011 | `work_item_authoring` | `Task flow` shows order and dependencies; `Task Candidates` shows candidate Tasks. | D-010, D-015 | The current wording does not state that a coordinating parent records only child Work Item overview and responsibility boundaries. | Parent and child Work Items could duplicate child Task graphs, procedures, release conditions, or next steps. | Add a non-duplication rule for child Work Item internals. Leave relation representation to W017. | `spec:product.design_records.authoring_standards.work_item_authoring` | candidate | T04; W017 relation decision | W016 | clarification_needed |
| C-012 | `task_authoring` | No rule limits implementation-time judgment. | D-006, D-016 | The current contract does not distinguish observable-equivalent local choices from contract-affecting decisions. | Implementation may silently change accepted boundaries or acceptance criteria. | Add allowed local-choice examples and a mandatory stop-and-return-to-decision boundary. | `spec:product.design_records.authoring_standards.task_authoring` | candidate | T04, then T05-T07 | W016 | conflict |
| C-013 | `artifact_responsibility_matrix` | Task “does not own … design decisions.” | D-004, D-012 | The unqualified wording can be read as prohibiting a decision Task from owning a decision ledger. | The matrix may contradict the accepted Task type while remaining the canonical ownership source. | Clarify that ADR or Specification owns canonical design state while a Task may own decision workflow state. | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | candidate | T04, then T05-T07 | W016 | clarification_needed |
| C-014 | `artifact_boundary` | “Has a decision been made and needs to be recorded? ADR.” | D-004, D-012 | The artifact selector does not distinguish an immediate Task ledger checkpoint from a durable ADR route. | The decision loop could be forced to author an ADR before its assigned authoring phase. | Clarify that the selector addresses canonical durable recording, not workflow-state persistence. | `spec:product.design_records.authoring_standards.artifact_boundary` | candidate | T04, then T05-T07 | W016 | clarification_needed |

### Confirmed compatible areas

| ID | target | compatible result | T02 authority | status |
|---|---|---|---|---|
| K-001 | `task_authoring` | Task already owns concrete closeable work, individual status, completion conditions, and verification Evidence. | D-003 through D-011 | compatible |
| K-002 | `task_authoring` | The current split guidance already uses independent concerns and one completion judgment as granularity signals. | D-017, D-018 | compatible |
| K-003 | `task_authoring` | The Task body is not the canonical source for ADR, Specification, or Investigation content. | D-005, D-012 | compatible |
| K-004 | `task_authoring` | Changed-file count is not a current responsibility criterion. | D-018 | compatible |
| K-005 | `work_item_authoring` and the responsibility matrix | A Work Item owns its own resolution flow and Task graph; Task status remains owned by each Task. | D-010, D-011 | compatible |
| K-006 | W016 and W017 Work Items | Shared Specification writer order is already W016 writer, then W017 writer, then integrated review. | W016 and W017 accepted workflow boundary | compatible |
| K-007 | `metadata_grammar` | Body headings are not metadata, so the metadata grammar alone does not prohibit an implementation-only H2. | D-006 | compatible |

### W017-routed items

| ID | target Specification | exact current contract | reason for W017 ownership | canonical target | dependency | status |
|---|---|---|---|---|---|---|
| R-001 | `requirement_authoring` | Requirement metadata requires reciprocal `work_items`. | REQ-006 removes persisted reverse Work Item membership. | `spec:product.design_records.authoring_standards.requirement_authoring` | W017 decision and synchronization workflow | routed_to_W017 |
| R-002 | `work_item_authoring` | Work Item metadata requires exactly one `source_requirement`. | REQ-006 replaces Requirement-only provenance with generic multi-value `source_refs`. | `spec:product.design_records.authoring_standards.work_item_authoring` | W017 decision and synchronization workflow | routed_to_W017 |
| R-003 | `task_authoring` | Task metadata and invariants require `source_requirement`. | REQ-006 replaces Task provenance with `source_refs` while preserving `work_item`. | `spec:product.design_records.authoring_standards.task_authoring` | W016-07 before W017-07 | routed_to_W017 |
| R-004 | `traceability`, `metadata_schema`, and `artifact_refs` | Workflow relation tables persist `requirement.work_items`, `work_item.source_requirement`, and `task.source_requirement`. | REQ-006 owns forward source refs and derived reverse relation views. | `spec:product.design_records.traceability` and child refs | W017 decision and synchronization workflow | routed_to_W017 |
| R-005 | `metadata_grammar` | DRMCP parses Requirement `work_items` and Work Item or Task `source_requirement`. | The PRODUCT source-relation contract must be accepted by W017 before DRMCP changes. | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | W017 canonical synchronization, then downstream DRMCP work | routed_to_W017 |

### Out-of-scope downstream impacts

| ID | target | impact | owner | dependency | status |
|---|---|---|---|---|---|
| O-001 | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Task metadata has no `task_type` field. A later DRMCP contract change must consume the accepted PRODUCT field and closed set. | downstream DRMCP | W016 canonical Specification and review closure | out_of_scope |
| O-002 | DRMCP validation and diagnostics | Type cardinality, allowed values, section alignment, and type-specific overlap checks require later parser or validator work. | downstream DRMCP | Accepted PRODUCT contract | out_of_scope |
| O-003 | Existing Task records | Migration sequencing and execution are excluded by REQ-005 and W016. | separate future migration scope | Accepted REQ-005 and REQ-006 migration contracts | out_of_scope |
| O-004 | Semantic boundary validator | Model selection, checklist shape, and diagnostic schema are excluded by REQ-005. | PRODUCT-REQ-SPEC-007 workflow | Accepted Task contract | out_of_scope |

### T04 readiness

- T04 can process C-001 through C-014 one item at a time.
- Each item has an exact target, authority, owner, and candidate resolution boundary.
- R-001 through R-005 must not be decided by T04.
- O-001 through O-004 require no W016 conflict decision.
- No conflict disposition or Specification wording was adopted in T03.

### Scoped verification result

- Every W016 conflict traces to T02 or accepted REQ-005 authority.
- Every source-relation issue is routed to W017.
- Shared writer serialization remains W016-07, then W017-07, then integrated review.
- The conflict inventory contains no Specification edits or ADR authoring.
- T03 completion conditions are satisfied.
