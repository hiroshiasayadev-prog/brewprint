# PRODUCT-ADR-SPEC-005: Enforce single-responsibility and independent Task completion boundaries

- **status**: accepted
- **date**: 2026-07-02
- **depends_on**:
  - PRODUCT-ADR-SPEC-004
- **supersedes**: []
- **migrated_to_spec**: null

## Context

`PRODUCT-ADR-SPEC-004` defines a closed Task-type taxonomy.
A type label alone does not guarantee one responsibility.

The current Task authoring guidance treats splitting as advisory.
It also leaves several completion boundaries ambiguous:

- semantic review versus objective verification;
- finding correction versus independent finding closure;
- graph coordination versus accepted-state synchronization;
- implementation detail versus contract-affecting judgment;
- one multi-file outcome versus several independent deliverables.

These ambiguities allow one Task to own several outcomes or approve its own work.
They also allow implementation and synchronization Tasks to introduce hidden design decisions.

## Decision

Every Task must own exactly one primary outcome, one completion judgment, and one acceptance or verification boundary.

A Task must be split when any of the following differs:

- primary outcome;
- completion judgment;
- acceptance or verification boundary;
- responsible owner;
- release decision;
- primary Task type.

The approximate `0.5d` to `3d` effort range remains advisory.
Effort and changed-file count are not responsibility tests.

### Common section alignment

`## Goal`, `## Work`, `## Done condition`, and `## Verification` must serve the same `task_type`-aligned outcome.

| section | role |
|---|---|
| `## Goal` | Declares the one primary outcome. |
| `## Work` | Contains only actions required to produce that outcome. |
| `## Done condition` | Defines the one completion judgment. |
| `## Verification` | Confirms the Done condition without adding new acceptance requirements. |

A supporting action associated with another Task type may remain in the Task only when the action owns no separate deliverable or completion judgment.

### Review and verification

Classify review and verification by primary outcome, not by method.

- `review` owns an independent semantic verdict and finding set.
- `verification` owns predefined objective check results.
- A review may run commands or tests as supporting Evidence.
- A separate verification Task is required when the acceptance gate itself needs independent ownership, aggregation, reuse, release gating, failure routing, or post-completion evaluation.

### Correction and finding closure

A `correction` Task may repair named findings and record direct self-check Evidence.
It must not close the findings it repairs.

A separate independent `review` Task decides each original finding as `CLOSED` or `OPEN`.
The same person may perform both phases only as separate Tasks or sessions.

An issue found and fixed inside the original Task needs no separate closure Task when no formal finding was opened.

### Coordination, Work Item decomposition, Work Item execution, and synchronization

A `coordination` Task may create or change:

- Task inventory;
- owner assignment;
- dependency structure;
- blockers;
- writer order;
- review order;
- release order;
- next-step routing.

A `work_item_decomposition` Task creates or splits child Work Items after the Work Item identity decision is fixed.
It owns child Work Item identity, responsibility boundaries, and parent-level routing.
It does not own child-internal deliverables or Task-graph coordination.

A `work_item_execution` Task represents exactly one already-created child Work Item as one parent-graph execution unit.
The Task uses `work_item_ref` for the child relation.
The Task may become `done` only after the referenced child Work Item is `done` and the Task records that status as Evidence.
A blocked child may block the Task.
A child in `not_started` or `in_progress` does not satisfy Task completion.
The execution Task does not create, split, redefine, or duplicate the child Work Item or its internal work.

A coordinating parent Work Item may summarize each child ID, purpose, responsibility boundary, coarse inter-child routing, and parent-level completion state.
It must not duplicate child-internal Task graphs, procedures, detailed dependencies, release conditions, or next-step decisions.

A `synchronization` Task only propagates an already accepted result.
Synchronization must stop when a new Task, dependency, owner, release judgment, or propagation choice is required.
The unresolved work returns to `coordination` or `decision`.

### Implementation judgment boundary

An `implementation` Task may choose observable-equivalent local details.
Examples include private names, helper decomposition, internal control flow, equivalent library calls, local data structures, and test setup.

Implementation must stop and return to `decision` when a choice changes or resolves:

- public API or metadata shape;
- responsibility or dependency boundaries;
- acceptance criteria;
- validation semantics;
- diagnostic categories or lifecycle;
- persistence or compatibility behavior;
- security or performance guarantees;
- external dependencies;
- externally observable behavior;
- conflicting Specification meaning.

### Multi-file cohesion

Several files may remain in one Task only when every change is necessary for the same outcome and completion judgment.

Production code, focused tests, and fixtures may share one implementation Task when they share one acceptance boundary.
Several Specifications may share one authoring Task when they express one accepted decision.
One named finding set may span several locations in one correction Task.

Implementation plus independent review, correction plus finding closure, and independently releasable changes must be split.

## Rationale

Responsibility is defined by owned outcomes and completion authority.
File count and elapsed effort are only weak proxies.

Independent judgments require independent Tasks.
This separation prevents an author or correction executor from approving the same work.

Outcome-based cohesion preserves legitimate multi-file changes.
The rule avoids artificial fragmentation while preventing omnibus Tasks.

Explicit stop conditions prevent implementation and synchronization from becoming hidden decision phases.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Split every multi-file change. | Several files may be required for one valid outcome and acceptance boundary. |
| Split only when estimated effort exceeds three days. | Small Tasks may still contain independent responsibilities. |
| Allow correction to close its own findings. | Repair ownership does not provide independent closure judgment. |
| Require a separate verification Task for every command. | Commands may be supporting Evidence for another Task's one completion judgment. |
| Let synchronization repair inconsistent graph state. | Graph repair requires judgment and belongs to coordination or decision work. |
| Let implementation resolve missing contract details. | Contract-affecting choices must remain explicit design decisions. |
| Use `work_item_decomposition` to wait for child completion. | Child creation and child execution have different outcomes and completion judgments. |
| Copy every child Task graph into the parent Work Item. | Duplicate graph state becomes stale and creates competing ownership. |

## Consequences

- `spec:product.design_records.authoring_standards.task_authoring` must define mandatory cohesion and type-aligned section rules.
- The Task authoring Specification must define the review, verification, correction, synchronization, and implementation stop boundaries.
- `spec:product.design_records.authoring_standards.work_item_authoring` must define Work Item decomposition and execution ownership and prohibit parent duplication of child Work Item internals.
- `implementation` Tasks require the separately defined conditional `## Implementation contract` shape.
- Downstream validators may later enforce the accepted boundaries.
- Validator behavior and diagnostics remain outside W016.
- Existing Task migration remains outside W016.

## Evidence

- `PRODUCT-REQ-SPEC-005`: accepted one-responsibility and one-completion requirement.
- `PRODUCT-ADR-SPEC-004`: accepted Task-type taxonomy.
- `PRODUCT-TASK-SPEC-016-02`: D-013 through D-018.
- `PRODUCT-TASK-SPEC-016-04`: C-003, C-005, and C-008 through C-012.
- `PRODUCT-TASK-SPEC-016-05`: ADR routing and this ADR boundary.
- `PRODUCT-TASK-SPEC-022-01`: accepted one-child execution relation and completion boundary.
