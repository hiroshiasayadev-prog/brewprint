# PRODUCT-TASK-SPEC-016-07: Synchronize REQ-005 Specifications

- **id**: PRODUCT-TASK-SPEC-016-07
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-016
- **source_requirement**: PRODUCT-REQ-SPEC-005
- **estimate**: 0.75d
- **depends_on**:
  - PRODUCT-TASK-SPEC-016-05
  - PRODUCT-TASK-SPEC-016-06
- **outputs**:
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - spec:product.design_records.authoring_standards.artifact_boundary
  - spec:product.design_records.artifact_model.artifact_responsibility_matrix

## Goal

Write the accepted REQ-005 Task contract into canonical PRODUCT Specifications.

Complete the W016 shared-file writer phase without changing W017-owned source-relation semantics.

## Work

- Apply decided T02 and T04 semantics to the exact T03 target sections.
- Reflect required ADR choices as current normative rules.
- Remove stale Task-contract wording within W016 scope.
- Preserve current provenance fields until W017 writes the accepted REQ-006 relation contract.
- Record exact changed refs and sections for integrated review.

## Done condition

- Canonical Specifications contain the complete accepted REQ-005 contract.
- No stale contradictory Task-type or responsibility wording remains in W016 scope.
- W017-owned source-relation semantics remain unchanged.
- Exact changed refs and sections are recorded for downstream review.

## Verification

- Trace every decided T02 and T04 item to current Specification text.
- Trace every required ADR to the corresponding normative rule.
- Confirm the writer order remains W016, then W017, then integrated review.
- Confirm no independent review or finding correction is performed by this Task.

## Evidence

### Authoring mode

- DRMCP is non-operational under the current agent authoring policy.
- Filesystem authoring was used for the exact T03 target set.
- T03 required no additional canonical target.
- The writable boundary remained the four Specifications and this Task.

### Changed Specifications

| Specification ref | changed sections |
|---|---|
| `spec:product.design_records.authoring_standards.task_authoring` | Metadata schema; File shape; Status lifecycle; Kind-specific authoring rules; Create; Update; Related specs. |
| `spec:product.design_records.authoring_standards.work_item_authoring` | Kind-specific authoring rules; Related specs. |
| `spec:product.design_records.authoring_standards.artifact_boundary` | Distinguishing adjacent artifacts; Decision workflow projection; Related specs. |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | Artifact responsibility matrix; Canonical design-state boundary; Related specs. |

All four changed Specifications use `date: 2026-07-01`.

### D-001 through D-018 trace

| decision | normative Specification location |
|---|---|
| D-001 | `task_authoring` Metadata schema and Create define required scalar `task_type`; Update permits changes. |
| D-002 | `task_authoring` Metadata schema defines the closed nine-value set. |
| D-003 | `task_authoring` Task type contract row for `investigation`. |
| D-004 | `task_authoring` Task type contract row for `decision` and Decision workflow Evidence. |
| D-005 | `task_authoring` Task type contract row for `authoring`. |
| D-006 | `task_authoring` Task type contract row for `implementation` and Implementation contract. |
| D-007 | `task_authoring` Task type contract row for `review`. |
| D-008 | `task_authoring` Task type contract row for `correction`. |
| D-009 | `task_authoring` Task type contract row for `verification`. |
| D-010 | `task_authoring` Task type contract row for `coordination`; `work_item_authoring` Parent coordination boundary. |
| D-011 | `task_authoring` Task type contract row for `synchronization`. |
| D-012 | `task_authoring` decision-versus-authoring boundary and Decision workflow Evidence; `artifact_boundary` Decision workflow projection. |
| D-013 | `task_authoring` review-versus-verification boundary. |
| D-014 | `task_authoring` correction-versus-finding-closure boundary. |
| D-015 | `task_authoring` coordination-versus-synchronization boundary; `work_item_authoring` Parent coordination boundary. |
| D-016 | `task_authoring` implementation-detail-versus-contract-decision boundary. |
| D-017 | `task_authoring` Common section alignment and conditional Implementation contract rules. |
| D-018 | `task_authoring` Single responsibility and multi-file cohesion rules. |

### C-001 through C-014 trace

| conflict | normative Specification location |
|---|---|
| C-001 | `task_authoring` Metadata schema, Create, and Update. |
| C-002 | `task_authoring` canonical Task type contract table. |
| C-003 | `task_authoring` mandatory Single responsibility rules. |
| C-004 | `task_authoring` File shape and Implementation contract placement, columns, and `TBD` rules. |
| C-005 | `task_authoring` Common section alignment. |
| C-006 | `task_authoring` decision-versus-authoring boundary; `artifact_boundary` conditional ADR routing flow. |
| C-007 | `task_authoring` Decision workflow Evidence and durable-rationale boundary. |
| C-008 | `task_authoring` review-versus-verification boundary and separate-gate criteria. |
| C-009 | `task_authoring` correction-versus-finding-closure boundary. |
| C-010 | `task_authoring` coordination-versus-synchronization stop conditions. |
| C-011 | `work_item_authoring` Parent coordination boundary. |
| C-012 | `task_authoring` implementation judgment stop boundary. |
| C-013 | `artifact_responsibility_matrix` Task workflow-state and canonical design-state boundary. |
| C-014 | `artifact_boundary` Task checkpoint, ADR routing, and Specification synchronization flow. |

C-001, C-002, and C-004 were projected directly from T02 and T04.
No additional ADR was created for those bounded Specification details.

### ADR trace

| ADR | normative Specification projection |
|---|---|
| `PRODUCT-ADR-SPEC-004` | `task_authoring` closed taxonomy and type contracts; `work_item_authoring` coordination-facing boundary. |
| `PRODUCT-ADR-SPEC-005` | `task_authoring` mandatory cohesion and independent completion boundaries; `work_item_authoring` parent-child boundary. |
| `PRODUCT-ADR-SPEC-006` | `task_authoring` decision workflow Evidence; `artifact_boundary` conditional routing; `artifact_responsibility_matrix` canonical ownership. |

The Specifications reference the public ADR IDs without copying detailed ADR rationale.
The ADR files remain unchanged and retain `migrated_to_spec: null` pending closure synchronization.

### W017 boundary preservation

- `source_requirement` remains the Task and Work Item provenance field.
- Task `work_item` membership remains unchanged.
- Work Item `tasks` membership remains unchanged.
- Requirement `work_items`, generic `source_refs`, derived reverse relations, relation validation, and relation migration were not changed.
- No traceability Specification or DRMCP Specification was changed.
- Writer order remains W016 T07, then W017 T07, then integrated independent review.

### Stale wording removal

- Replaced advisory responsibility splitting with mandatory outcome-based cohesion.
- Removed the implication that every decision immediately requires an ADR.
- Removed the absolute prohibition on concise Task decision-workflow rationale.
- Replaced the unconditional statement that Tasks cannot own decision state.
- Replaced the artifact selector that routed every completed decision directly to an ADR.

### Lifecycle and verification boundary

- No existing Task migration was performed.
- No parser, validator, diagnostic, index, or tool-projection design was added.
- No independent review was performed.
- No finding was opened, corrected, or closed.
- T02 and T04 decision states remain `decided`.
- PRODUCT-WORK-SPEC-016 remains `in_progress`.
- T07 is complete with four exact canonical Specification outputs.
