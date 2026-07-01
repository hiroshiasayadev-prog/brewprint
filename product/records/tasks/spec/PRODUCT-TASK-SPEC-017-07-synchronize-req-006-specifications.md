# PRODUCT-TASK-SPEC-017-07: Synchronize REQ-006 Specifications

- **id**: PRODUCT-TASK-SPEC-017-07
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-017
- **source_requirement**: PRODUCT-REQ-SPEC-006
- **estimate**: 0.75d
- **depends_on**:
  - PRODUCT-TASK-SPEC-016-07
  - PRODUCT-TASK-SPEC-017-05
  - PRODUCT-TASK-SPEC-017-06
- **outputs**:
  - spec:product.design_records.authoring_standards.requirement_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.artifact_boundary
  - spec:product.design_records.artifact_model.artifact_responsibility_matrix
  - spec:product.design_records.traceability
  - spec:product.design_records.traceability.metadata_schema
  - spec:product.design_records.traceability.artifact_refs
  - spec:product.design_records.traceability.resolve_and_validation

## Goal

Write the accepted REQ-006 source-relation contract into canonical PRODUCT Specifications.

Complete the final shared provenance writer phase after W016 Task-contract synchronization.

## Work

- Wait for `PRODUCT-TASK-SPEC-016-07` to complete.
- Apply decided T02 and T04 semantics to the exact T03 target sections.
- Reflect required ADR choices as current normative rules.
- Replace Requirement-only provenance and persisted Requirement reverse membership in canonical Specifications.
- Preserve the accepted REQ-005 Task-type and single-responsibility contract.
- Record exact changed refs and sections for integrated review.

## Done condition

- Canonical Specifications contain the complete accepted REQ-006 contract.
- No stale contradictory source-relation or reciprocal Requirement wording remains in W017 scope.
- Shared Task-authoring content preserves the accepted W016 result.
- Exact changed refs and sections are recorded for T08.

## Verification

- Trace every decided T02 and T04 item to current Specification text.
- Trace every required ADR to the corresponding normative rule.
- Confirm `PRODUCT-TASK-SPEC-016-07` completed before shared files changed.
- Confirm no independent review, migration execution, or finding correction is performed by this Task.

## Evidence

### Authoring mode and boundary

- DRMCP is non-operational under the current agent authoring policy.
- Filesystem fallback was used for the nine canonical Specifications and this Task.
- `PRODUCT-TASK-SPEC-016-07` was `done` before shared files changed.
- The writable boundary remained the declared ten files.
- `PRODUCT-REQ-SPEC-006` and `PRODUCT-ADR-SPEC-001`, `PRODUCT-ADR-SPEC-007`, and `PRODUCT-ADR-SPEC-008` were read-only authorities.
- The bootstrap workflow records remain outside migration execution scope.
- No canonical metadata exception was added for bootstrap records.

### Changed Specifications

| Specification ref | changed sections |
|---|---|
| `spec:product.design_records.authoring_standards.requirement_authoring` | Metadata schema; Status lifecycle; Kind-specific authoring rules; Canonical reference policy; Related specs. |
| `spec:product.design_records.authoring_standards.work_item_authoring` | File shape; Metadata schema; Status lifecycle; Kind-specific authoring rules; Canonical reference policy; Update; Related specs. |
| `spec:product.design_records.authoring_standards.task_authoring` | Metadata schema; General Task rules; Canonical reference policy; Create; Update; Related specs. |
| `spec:product.design_records.authoring_standards.artifact_boundary` | Covered artifact kinds; Distinguishing adjacent artifacts; Related specs. |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | Artifact responsibility matrix; Source-relation ownership boundary; Related specs. |
| `spec:product.design_records.traceability` | Current contract; Non-goals; Traceability model; Sources. |
| `spec:product.design_records.traceability.metadata_schema` | Workflow relation metadata; Migration state; Metadata boundary; Related specs. |
| `spec:product.design_records.traceability.artifact_refs` | Workflow relation identity; Sources. |
| `spec:product.design_records.traceability.resolve_and_validation` | Workflow relation validation; Excluded implementation behavior; Resolve and validation boundary; Related specs. |

All changed Specifications use `date: 2026-07-01`.

### D-001 through D-019 normative trace

| decision | normative Specification location |
|---|---|
| D-001 | `work_item_authoring` Metadata schema defines required non-empty `source_refs`; `task_authoring` Metadata schema and General Task rules prohibit Task source fields. |
| D-002 | `work_item_authoring` Metadata schema and `traceability.metadata_schema` define unordered-set semantics. |
| D-003 | `work_item_authoring` Metadata schema, `traceability.metadata_schema`, and `resolve_and_validation` reject duplicates. |
| D-004 | `traceability.metadata_schema` and `resolve_and_validation` reject Work Item self-reference. |
| D-005 | `traceability.metadata_schema` and `resolve_and_validation` require resolvable canonical refs. |
| D-006 | `work_item_authoring` Canonical reference policy and `artifact_refs` reuse every active canonical reference class. |
| D-007 | `work_item_authoring` Kind-specific authoring rules define direct material source selection and transitive omission. |
| D-008 | `work_item_authoring` Parent coordination boundary and `artifact_boundary` require exact Task provenance for Task-created Work Items. |
| D-009 | `work_item_authoring` Parent coordination boundary prohibits automatic owner Work Item and upstream-source copying. |
| D-010 | `task_authoring` General Task rules define provenance only through `work_item`. |
| D-011 | `requirement_authoring`, `traceability`, and `traceability.metadata_schema` define a direct-only Requirement reverse relation. |
| D-012 | `requirement_authoring` and `traceability.metadata_schema` define the reverse relation as unordered and duplicate-free. |
| D-013 | `work_item_authoring` Migration rules and `traceability.metadata_schema` define one-element Work Item conversion with no inferred source. |
| D-014 | `task_authoring` Migration rules and `traceability.metadata_schema` define removal-only Task migration. |
| D-015 | `requirement_authoring` Migration rule and `traceability.metadata_schema` require exact reverse-set equality and block mismatches. |
| D-016 | `work_item_authoring` Migration rules and `traceability.metadata_schema` allow staged migration with atomic per-record transitions. |
| D-017 | `traceability` and `resolve_and_validation` define semantic provenance-cycle invalidity and Task-to-owner normalization semantics. |
| D-018 | `artifact_responsibility_matrix`, `traceability.metadata_schema`, and `resolve_and_validation` apply the ownership boundary already accepted by `PRODUCT-ADR-SPEC-001`. |
| D-019 | Shared `work_item_authoring`, `task_authoring`, `artifact_boundary`, and responsibility-matrix content preserves the completed W16 writer result. |

### C-002 through C-026 disposition trace

| conflict | disposition in current Specifications |
|---|---|
| C-002 | Removed persisted Requirement `work_items` ownership from `requirement_authoring`. |
| C-003 | Added the exact-match Requirement migration guard. |
| C-004 | Replaced Work Item `source_requirement` with required non-empty `source_refs`. |
| C-005 | Removed Requirement reciprocity and Task source-matching invariants; retained Work Item/Task ownership consistency. |
| C-006 | Defined Work Item `source_refs` minimum cardinality and active canonical ref boundary. |
| C-007 | Added direct material source-selection rules. |
| C-008 | Added exact Task provenance for Task-created Work Items. |
| C-009 | Added atomic one-element Work Item migration with no inferred sources. |
| C-010 | Removed Task `source_requirement` from metadata and authoring interfaces. |
| C-011 | Removed the Task source-Requirement matching invariant. |
| C-012 | Added removal-only Task migration and prohibited Task `source_refs`. |
| C-013 | Preserved the W16 parent coordination boundary while adding provenance rules. |
| C-014 | Preserved W16 `task_type`, taxonomy, cohesion, and adjacent responsibility boundaries. |
| C-015 | Generalized Work Item purpose from Requirement-only resolution to direct material sources. |
| C-016 | Preserved the complete W16 decision workflow projection. |
| C-017 | Generalized the Work Item responsibility-matrix row and added Task-originated decomposition. |
| C-018 | Added PRODUCT semantic and DRMCP mechanism ownership split. |
| C-019 | Preserved Task, ADR, and Specification canonical design-state ownership. |
| C-020 | Replaced the active legacy workflow relation set in `traceability`. |
| C-021 | Separated the direct Requirement reverse relation from transitive traversal and defined semantic cycle invalidity. |
| C-022 | Replaced the legacy reciprocal relation table in `traceability.metadata_schema`. |
| C-023 | Added non-empty unordered-set semantics and source validity conditions. |
| C-024 | Added staged migration and same-record old/new exclusion. |
| C-025 | Replaced legacy workflow relation identities without changing active canonical reference classes. |
| C-026 | Replaced reciprocity and Task source checks with source-ref and semantic-cycle validation. |

C-001 was already resolved by the corrected accepted Requirement and was not edited by this Task.
C-027 through C-030 remain downstream DRMCP routes.
C-031 remains workflow execution scope only and was not generalized into Specifications.

### ADR projection

| ADR | Specification projection |
|---|---|
| `PRODUCT-ADR-SPEC-007` | Work Item-only persisted provenance, Task provenance through `work_item`, direct Requirement reverse derivation, active canonical ref reuse, direct material source selection, Task-created Work Item provenance, and source validity. |
| `PRODUCT-ADR-SPEC-008` | Staged repository migration, atomic record transition, Work Item one-element conversion, Task removal-only transition, and exact-match Requirement transition. |
| `PRODUCT-ADR-SPEC-001` | D-018 PRODUCT semantic ownership and DRMCP app-local mechanism boundary. |

Specifications reference the ADRs without copying detailed rationale.
ADR files remain unchanged with `migrated_to_spec: null` pending closure synchronization.

### W16 preservation result

- Required scalar `task_type` remains unchanged.
- The closed nine-value Task taxonomy remains unchanged.
- Single-responsibility and common-section alignment rules remain unchanged.
- Review versus verification remains unchanged.
- Correction versus finding closure remains unchanged.
- Coordination versus synchronization remains unchanged.
- The implementation judgment stop boundary remains unchanged.
- The parent coordination boundary remains intact.
- Decision checkpoint and conditional ADR routing remain intact.
- Task, ADR, and Specification canonical design-state ownership remains intact.

### Downstream DRMCP boundary

PRODUCT Specifications now define persisted provenance semantics, canonical relation meaning, invalid conditions, and migration semantics.

DRMCP still owns:

- parser and writer implementation;
- indexing and direct reverse lookup;
- transitive traversal;
- Task-owner resolution mechanics;
- cycle-analysis algorithms;
- diagnostics and severities;
- response schemas;
- user-visible projections;
- migration commands, transactions, rollback, resume, and release gates.

No concrete DRMCP contract or implementation was added.

### T08 review handoff

The exact integrated review file set is:

- `product/records/spec/design-records/authoring-standards/requirement-authoring.md`
- `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
- `product/records/spec/design-records/authoring-standards/task-authoring.md`
- `product/records/spec/design-records/authoring-standards/artifact-boundary.md`
- `product/records/spec/design-records/artifact-model/artifact-responsibility-matrix.md`
- `product/records/spec/design-records/traceability/index.md`
- `product/records/spec/design-records/traceability/metadata-schema.md`
- `product/records/spec/design-records/traceability/artifact-refs.md`
- `product/records/spec/design-records/traceability/resolve-and-validation.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-016-07-synchronize-req-005-specifications.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-017-02-run-req-006-interactive-decision-loop.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-017-03-investigate-specification-and-migration-conflicts.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-017-04-run-specification-conflict-decision-loop.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-017-05-classify-adr-routing.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-017-06-author-required-adrs.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-017-07-synchronize-req-006-specifications.md`
- `product/records/requirements/spec/PRODUCT-REQ-SPEC-005-typed-single-responsibility-task-contract.md`
- `product/records/requirements/spec/PRODUCT-REQ-SPEC-006-generic-workflow-source-relations.md`
- `product/records/adr/spec/PRODUCT-ADR-SPEC-001-product-spec-semantic-ownership-boundary.md`
- `product/records/adr/spec/PRODUCT-ADR-SPEC-004-define-closed-typed-task-responsibility-taxonomy.md`
- `product/records/adr/spec/PRODUCT-ADR-SPEC-005-enforce-single-responsibility-and-independent-task-completion-boundaries.md`
- `product/records/adr/spec/PRODUCT-ADR-SPEC-006-separate-decision-workflow-checkpoints-from-canonical-design-state.md`
- `product/records/adr/spec/PRODUCT-ADR-SPEC-007-use-work-item-source-refs-as-canonical-workflow-provenance.md`
- `product/records/adr/spec/PRODUCT-ADR-SPEC-008-migrate-legacy-workflow-source-relations-through-atomic-record-transitions.md`

### Lifecycle and execution boundary

- T02 and T04 decision states remain `decided`.
- No existing record migration was executed.
- No migration mismatch was repaired.
- No independent review was performed.
- No finding was opened, corrected, or closed.
- No lifecycle closure was performed.
- No stage or commit was performed.
