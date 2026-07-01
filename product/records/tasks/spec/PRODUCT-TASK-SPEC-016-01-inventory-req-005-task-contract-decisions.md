# PRODUCT-TASK-SPEC-016-01: Inventory REQ-005 Task-contract decisions

- **id**: PRODUCT-TASK-SPEC-016-01
- **status**: done
- **date**: 2026-06-30
- **work_item**: PRODUCT-WORK-SPEC-016
- **source_requirement**: PRODUCT-REQ-SPEC-005
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - PRODUCT-TASK-SPEC-016-01

## Goal

Create the complete known REQ-005 decision inventory and dependency order.

Separate accepted constraints, user decisions, and W017-owned conflict candidates before interactive questioning begins.

## Work

- Read the accepted REQ-005 and directly related current PRODUCT Specifications.
- Resolve facts already fixed by accepted authority without asking the user.
- Identify unresolved Task-type, responsibility, completion, overlap, and scope decisions.
- Record dependency edges, ADR routes, and likely canonical targets.
- Route source-relation questions to W017 instead of deciding them in W016.
- Select the first unblocked decision for T02.

## Done condition

- Every currently known REQ-005 decision has a stable local ID.
- Accepted constraints are separated from user decisions.
- Dependencies and blocked conflict candidates are explicit.
- Existing Task migration is outside W016 and requires no migration action.
- One first unblocked decision is selected for T02.

## Verification

- Compare the inventory against every Required Outcome and Explicitly Excluded Scope item in PRODUCT-REQ-SPEC-005.
- Confirm W017-owned source-relation semantics are not decided by W016.
- Confirm no accepted REQ-005 constraint is restated as an open user decision.
- Confirm no more than one decision is selected for discussion.

## Evidence

### Bootstrap disposition

`BOOTSTRAP-001` was decided by the user on 2026-06-30.

- Use the current canonical Task metadata contract.
- Use `source_requirement` and `work_item`.
- Do not add `type`, `task_type`, `primary_type`, or `source_refs` before Specification acceptance.
- Do not require migration of these workflow Tasks.
- Treat the bootstrap choice as workflow authoring policy, not as the canonical REQ-005 product contract.

### Accepted authority constraints

The following points are already fixed by PRODUCT-REQ-SPEC-005 and are not user decisions in T02.

| constraint | accepted result |
|---|---|
| Type cardinality | Every Task declares exactly one primary Task type. |
| Responsibility cardinality | Every Task owns one primary responsibility. |
| Completion cardinality | Every Task closes through one completion judgment. |
| Responsibility test | Owned outcome and completion judgment determine responsibility. Changed-file count does not. |
| Authoring and review | One Task cannot author an artifact and independently review that authoring. |
| Correction and finding closure | One Task cannot correct findings and independently close those findings. |
| Implementation readiness | An implementation Task cannot contain unresolved design decisions. |
| Coordination | A coordination Task cannot produce child Task or child Work Item deliverables. |
| Synchronization | A synchronization Task cannot introduce design, implementation, review, or correction work. |
| Multi-file work | Multiple files are allowed when all changes serve one primary outcome and verification boundary. |
| Membership relation | Task `work_item` remains the owning Work Item relation. |
| Migration execution | Existing Task migration sequencing is outside REQ-005 and W016. |

### Decision inventory

| ID | Topic | Status | Depends on | Decision summary | ADR route | Canonical target |
|---|---|---|---|---|---|---|
| D-001 | Primary Task type field name | in_discussion | — | Select the required persisted field name. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |
| D-002 | Allowed primary Task type closed set | open | D-001 | Select the complete closed set of allowed values. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-003 | Investigation type contract | blocked | D-002 | Define its owned outcome, completion judgment, and prohibited overlaps if selected. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-004 | Decision type contract | blocked | D-002 | Define its owned outcome, completion judgment, and prohibited overlaps if selected. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-005 | Authoring type contract | blocked | D-002 | Define its owned outcome, completion judgment, and prohibited overlaps if selected. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-006 | Implementation type contract | blocked | D-002 | Define its owned outcome, completion judgment, and prohibited overlaps if selected. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-007 | Review type contract | blocked | D-002 | Define its owned outcome, completion judgment, and prohibited overlaps if selected. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-008 | Correction type contract | blocked | D-002 | Define its owned outcome, completion judgment, and prohibited overlaps if selected. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-009 | Verification type contract | blocked | D-002 | Define its owned outcome, completion judgment, and prohibited overlaps if selected. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-010 | Coordination type contract | blocked | D-002 | Define its owned outcome, completion judgment, and prohibited overlaps if selected. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-011 | Synchronization type contract | blocked | D-002 | Define its owned outcome, completion judgment, and prohibited overlaps if selected. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-012 | Decision and authoring boundary | blocked | D-004, D-005 | Define where choice confirmation ends and canonical artifact writing begins. | candidate | `spec:product.design_records.authoring_standards.artifact_boundary` |
| D-013 | Review and verification boundary | blocked | D-007, D-009 | Define independent judgment versus completion-evidence confirmation. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-014 | Correction and finding-closure review boundary | blocked | D-007, D-008 | Define correction completion versus independent finding disposition. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-015 | Coordination and synchronization boundary | blocked | D-010, D-011 | Define graph routing versus lifecycle and Evidence state updates. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-016 | Judgment allowed inside implementation Tasks | blocked | D-006 | Define which local choices remain implementation detail and which choices block implementation. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-017 | Type alignment for Goal, Work, Done condition, and Verification | blocked | D-002 through D-011 | Define the conformance test across the four Task sections. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-018 | Multi-file single-responsibility test | blocked | D-002 | Define the exact cohesion conditions for one Task changing several files. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |

### Scope dispositions and conflict candidates

| topic | disposition |
|---|---|
| Existing Task migration | Outside W016. No migration sequencing, execution, or marker is required. |
| Workflow Task bootstrap migration | Not required by the user decision for BOOTSTRAP-001. |
| `source_requirement` to `source_refs` | W017-owned. T03 may identify a shared Specification conflict, but W016 does not decide the relation contract. |
| Requirement reciprocal `work_items` | W017-owned. Not part of the REQ-005 decision loop. |
| Shared Specification writer order | Fixed by W016 and W017: W016 writer, then W017 writer, then integrated review. |

### Scoped verification result

The inventory covers every unresolved area required by the W016 startup contract.

No repository-wide traversal was used.
The first unblocked decision is D-001.
