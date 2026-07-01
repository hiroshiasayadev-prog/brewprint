# PRODUCT-TASK-SPEC-017-04: Run Specification-conflict decision loop

- **id**: PRODUCT-TASK-SPEC-017-04
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-017
- **source_requirement**: PRODUCT-REQ-SPEC-006
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-017-03
- **outputs**:
  - PRODUCT-TASK-SPEC-017-04

## Goal

Persist one disposition at a time for material Specification and migration conflicts found by T03.

Finish with every W017-owned conflict decided, deferred, or blocked.

## Work

- Initialize a conflict decision ledger from T03.
- Ask exactly one conflict disposition per user turn.
- Check each answer against REQ-006, T02 decisions, and W016 ownership.
- Persist the answer and verify the scoped diff before advancing.
- Route W016-owned Task-type questions to W016 without deciding them.

## Done condition

- Every W017-owned conflict has a durable disposition.
- Every shared conflict has an explicit writer-order or dependency disposition.
- Every W016-owned conflict remains routed to W016.
- ADR routing and Specification synchronization can proceed without hidden conflict decisions.

## Verification

- Confirm at most one conflict decision is `in_discussion`.
- Confirm every explicit answer is persisted before cursor advancement.
- Confirm no ADR, Specification, or migration target file changed.
- Confirm no W016-owned Task-type decision was adopted in W017.

## Evidence

### Loop state

Loop status: decision_complete

Current conflict: none

Inventory source: PRODUCT-TASK-SPEC-017-03

User questions asked: 0

T03 found no materially different disposition requiring user judgment.
Accepted T02 authority therefore determines every W017-owned conflict disposition.

### Conflict decision ledger

| ID | Topic | Status | Authority | Decision summary | ADR route | Canonical target |
|---|---|---|---|---|---|---|
| C-001 | REQ-006 Task `source_refs` wording | decided | D-001, D-010 | Correct REQ-006 so only Work Items persist `source_refs`. Tasks persist no source field and expose provenance through `work_item`. | candidate | `PRODUCT-REQ-SPEC-006` |
| C-002 | Persisted Requirement `work_items` | decided | D-011, D-012 | Remove `work_items` from the canonical persisted Requirement contract. Define the direct Requirement reverse relation as a derived view. | candidate | `spec:product.design_records.authoring_standards.requirement_authoring` |
| C-003 | Requirement reverse-list migration guard | decided | D-015 | Permit removal of Requirement `work_items` only after exact unordered, duplicate-free equality with the derived direct reverse set. A mismatch blocks migration. | candidate | `spec:product.design_records.authoring_standards.requirement_authoring` |
| C-004 | Work Item provenance field shape | decided | D-001, D-006 | Replace required scalar `source_requirement` with required non-empty `source_refs`. Accept every active canonical reference class defined by `artifact_refs`. | candidate | `spec:product.design_records.authoring_standards.work_item_authoring` |
| C-005 | Legacy reciprocal provenance invariants | decided | D-010, D-011, D-012 | Remove Requirement reciprocity and Task source-Requirement equality invariants. Preserve only Work Item `tasks` and Task `work_item` membership consistency. | candidate | `spec:product.design_records.authoring_standards.work_item_authoring` |
| C-006 | Work Item `source_refs` minimum and reference classes | decided | D-001, D-006 | Require at least one canonical ref. Permit active spec refs, supported record public IDs, and resolver-supported legacy issued IDs. | candidate | `spec:product.design_records.authoring_standards.work_item_authoring` |
| C-007 | Direct and material source selection | decided | D-007 | List every direct material source. Omit incidental context and transitively reachable ancestors unless independently material. | not_required | `spec:product.design_records.authoring_standards.work_item_authoring` |
| C-008 | Task-created Work Item provenance | decided | D-008, D-009 | A Work Item created or decomposed from a Task cites the exact Task ID. Do not automatically copy the Task owner or its upstream sources. | candidate | `spec:product.design_records.authoring_standards.work_item_authoring` |
| C-009 | Work Item legacy-field conversion | decided | D-013, D-016 | Convert `source_requirement` into a one-element `source_refs` list and remove the old field in the same record update. Infer no additional source. | candidate | `spec:product.design_records.authoring_standards.work_item_authoring` |
| C-010 | Task source metadata removal | decided | D-001, D-010, D-014 | Remove Task `source_requirement`. Do not add Task `source_refs`. Preserve `work_item` as the provenance traversal edge. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| C-011 | Task source-matching invariant | decided | D-010 | Remove the requirement that a Task repeat or match the Work Item source. Keep Task provenance reachable only through `work_item`. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |
| C-012 | Task removal-only migration | decided | D-014, D-016 | Remove Task `source_requirement` atomically without replacement. Preserve `work_item` and all W16 Task-type metadata. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| C-013 | Shared Work Item writer section | decided | D-007, D-008, D-009, D-019 | Apply W17 provenance wording after completed W16 T07 content. Preserve the accepted parent coordination boundary unchanged. | not_required | `spec:product.design_records.authoring_standards.work_item_authoring` |
| C-014 | Shared Task writer section | decided | D-010, D-014, D-019 | Make relation-only edits around the accepted W16 Task contract. Preserve `task_type`, single responsibility, section alignment, and adjacent responsibility boundaries. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |
| C-015 | Requirement-only Work Item purpose wording | decided | D-007, D-008, D-009 | Generalize Work Item provenance to direct material canonical sources, including Tasks. Preserve Work Item ownership of its resolution flow and Task graph. | candidate | `spec:product.design_records.authoring_standards.artifact_boundary` |
| C-016 | Shared artifact-boundary wording | decided | D-019 | Preserve the W16 decision checkpoint, ADR routing, and Specification synchronization flow. Change only Work Item provenance wording. | not_required | `spec:product.design_records.authoring_standards.artifact_boundary` |
| C-017 | Requirement-only responsibility-matrix wording | decided | D-007, D-008, D-009 | Generalize the Work Item row from Requirement-only provenance to direct material source provenance and Task-originated decomposition. | candidate | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` |
| C-018 | PRODUCT and DRMCP source-relation ownership | decided | D-018 | PRODUCT owns persisted provenance semantics and invalid conditions. DRMCP owns indexing, traversal, owner resolution, diagnostics, response schemas, and projections. | not_required | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` |
| C-019 | Shared responsibility-matrix wording | decided | D-019 | Preserve W16 Task, ADR, and Specification design-state ownership. Add source-relation ownership without changing those boundaries. | not_required | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` |
| C-020 | Active workflow relation set | decided | D-001, D-010, D-011, D-012 | Replace legacy provenance and reverse fields with Work Item `source_refs`. Preserve Work Item `tasks`, Task `work_item`, and Task `depends_on`. | candidate | `spec:product.design_records.traceability` |
| C-021 | Direct reverse relation and transitive traversal | decided | D-011, D-012, D-017, D-018 | Define the direct Requirement reverse set separately from transitive graph traversal. Define provenance-cycle invalidity while leaving algorithms to DRMCP. | candidate | `spec:product.design_records.traceability` |
| C-022 | Trace metadata relation table | decided | D-001, D-010, D-011, D-012 | Replace legacy reciprocal fields and mismatch rules with forward Work Item provenance and persisted Task membership relations. | candidate | `spec:product.design_records.traceability.metadata_schema` |
| C-023 | Work Item source-set validity | decided | D-001, D-002, D-003, D-004, D-005, D-011, D-012 | Define non-empty unordered set semantics. Reject duplicates, self-reference, unresolved refs, and noncanonical refs. Define the derived direct reverse set as unordered and duplicate-free. | candidate | `spec:product.design_records.traceability.metadata_schema` |
| C-024 | Staged migration and mixed metadata | decided | D-013, D-014, D-015, D-016 | Allow repository-wide staged migration. Require each record to switch atomically and prohibit same-record old and new provenance fields. | candidate | `spec:product.design_records.traceability.metadata_schema` |
| C-025 | Workflow relation identities | decided | D-001, D-010, D-012 | Replace Requirement `work_items` and both `source_requirement` relation identities. Keep the existing active canonical ref classes unchanged. | candidate | `spec:product.design_records.traceability.artifact_refs` |
| C-026 | Source-relation validation and cycles | decided | D-003, D-004, D-005, D-017 | Replace reciprocity and Task source-Requirement checks with source-ref existence, canonical identity, duplicate, self-reference, and semantic provenance-cycle invalidity. | candidate | `spec:product.design_records.traceability.resolve_and_validation` |
| C-031 | Workflow bootstrap migration exclusion | decided | BOOTSTRAP-001 | Exclude W16 and W17 workflow bootstrap Tasks from existing-record migration action. Treat the exclusion as workflow scope, not a generic canonical metadata exception. | not_required | Migration execution scope and Evidence |

### Downstream DRMCP routing

| ID | Status | PRODUCT result | Downstream owner |
|---|---|---|---|
| C-027 | routed | DRMCP must consume the accepted field transition after PRODUCT synchronization. | Metadata parser, writer, and migration implementation. |
| C-028 | routed | DRMCP must derive direct Requirement reverse sets separately from transitive traversal. | Indexing, reverse lookup, traversal, and projection. |
| C-029 | routed | Task refs normalize to owning Work Items for cycle semantics. PRODUCT does not select the owner-resolution mechanism. | Owner resolution and provenance graph analysis. |
| C-030 | routed | PRODUCT defines invalid conditions and ownership only. | Diagnostic vocabulary, response schemas, and user-visible projections. |

Current DRMCP Specification or implementation state was not inspected.
T04 adopts no concrete parser, diagnostic, index, traversal, or tool contract.

### Shared-writer and W16 preservation

- W16 T07 remains the completed first writer.
- W17 T07 remains the next shared Specification writer.
- Integrated independent review remains after W17 synchronization.
- No W16-owned Task-type decision is reopened.
- No W16 Task-type, responsibility, review-independence, or design-state wording is weakened.

### Recorded scope boundary

- C-001 through C-026 and C-031 are `decided` from accepted authority.
- C-027 through C-030 are routed to downstream DRMCP work.
- No conflict is `open`, `in_discussion`, `blocked`, or `deferred`.
- No additional user decision was required.
- T05 owns final ADR routing classification.
- T06 owns any required ADR authoring.
- T07 owns canonical Specification synchronization after ADR routing.
- No current Task explicitly owns editing `PRODUCT-REQ-SPEC-006` for C-001.
- C-001 writer ownership requires a coordination amendment or an explicit T07 boundary extension before Requirement correction begins.
- The writer gap requires no new design decision and does not block T05 ADR routing.

### Closure verification

- T03 is `done`.
- Every W17-owned conflict has one durable disposition.
- Every shared conflict preserves the W16-first writer order.
- Every downstream DRMCP impact remains outside PRODUCT implementation scope.
- No ADR, Requirement, Specification, migration target, source, test, or fixture file changed.
- T04 is complete and ready for ADR routing.
- Requirement correction writer ownership remains an explicit downstream coordination gap.
- T04 does not amend the Task graph or assign T07 an undeclared Requirement write boundary.
- Bootstrap authority remains `PRODUCT-TASK-SPEC-016-01` Evidence, `BOOTSTRAP-001`.
- This Task remains authored under the current workflow bootstrap `source_requirement` and `work_item` metadata contract.
- No migration action is owned or performed by this workflow Task.
