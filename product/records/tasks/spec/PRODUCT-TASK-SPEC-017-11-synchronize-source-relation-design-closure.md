# PRODUCT-TASK-SPEC-017-11: Synchronize source-relation design closure

- **id**: PRODUCT-TASK-SPEC-017-11
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-017
- **source_requirement**: PRODUCT-REQ-SPEC-006
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-017-08
- **outputs**:
  - PRODUCT-TASK-SPEC-017-11
  - PRODUCT-WORK-SPEC-017

## Goal

Synchronize the accepted integrated-review result into W017 closure Evidence.

Confirm the W017 Work Item completion conditions and synchronize only W017 T11 and W017 Work Item lifecycle state.

Record existing-record migration readiness without executing migration.

## Work

1. Confirm the initial `NEEDS REVISION` verdict in `PRODUCT-TASK-SPEC-017-08`.
2. Confirm `PRODUCT-TASK-SPEC-017-09` is `done`.
3. Confirm `PRODUCT-TASK-SPEC-017-10` reports `PASS` and closes `F-MAJ-01` and `F-MAJ-02`.
4. Read W017 T02 and T04 as read-only decision trace authority without changing their decision status.
5. Confirm the ADR routing outcomes in W017 T05.
6. Confirm `PRODUCT-ADR-SPEC-001`, `PRODUCT-ADR-SPEC-007`, and `PRODUCT-ADR-SPEC-008` as the exact ADR refs.
7. Confirm `PRODUCT-REQ-SPEC-006` and the exact Specification outputs recorded by W017 T07.
8. Record the reviewed decision-to-ADR-to-Requirement-and-Specification trace in this Task's Evidence.
9. Record C-027 through C-030 as downstream DRMCP obligations.
10. Record that existing-record migration semantics are design-complete and migration execution remains separate scope.
11. Confirm every W017 Work Item completion condition.
12. During closure execution, synchronize only this Task and `PRODUCT-WORK-SPEC-017`.

### Acceptance route

| T08 verdict | acceptance gate |
|---|---|
| `PASS` | Accepted T08 Evidence is sufficient. T09 and T10 remain not required. |
| `NEEDS REVISION` | T09 must be `done`. T10 must independently close every required finding. |

The accepted route is `NEEDS REVISION` followed by T09 `done` and T10 `PASS`.
The W017 closure acceptance gate is satisfied.

### Decision status preservation

- W017 T02 D-001 through D-019 remain `decided`.
- W017 T04 C-001 through C-026 and C-031 remain `decided`.
- C-027 through C-030 remain routed downstream DRMCP obligations.
- Downstream work progress does not change decision status.
- This closure Task does not rewrite a completed decision Task.

### ADR handling

- `PRODUCT-ADR-SPEC-001`, `PRODUCT-ADR-SPEC-007`, and `PRODUCT-ADR-SPEC-008` are read-only.
- This Task does not change ADR content or metadata.
- This Task does not update `migrated_to_spec`.
- This Task records ADR-to-Requirement-and-Specification correspondence in its own Evidence.
- Any required ADR metadata synchronization must be routed to a separate Task.

## Done condition

- Accepted T10 finding dispositions exist.
- The W017 decision-to-ADR-to-Requirement-and-Specification trace is recorded in this Task's Evidence.
- C-027 through C-030 are recorded exactly as downstream DRMCP obligations.
- Existing-record migration design readiness is recorded.
- Migration execution remains unperformed.
- `PRODUCT-WORK-SPEC-017` satisfies every completion condition.
- Only this Task and `PRODUCT-WORK-SPEC-017` are closure-owned lifecycle records.
- Completed decision Tasks, ADRs, the Requirement, and Specifications remain unchanged.
- No new design, correction, implementation, or migration execution is performed.

## Verification

- Confirm the accepted route is T08 `NEEDS REVISION`, T09 `done`, and T10 `PASS`.
- Confirm T10 records `F-MAJ-01: CLOSED` and `F-MAJ-02: CLOSED`.
- Confirm W017 T02 and T04 decision entries remain `decided`.
- Confirm W017 T05 routing matches ADR-001, ADR-007, and ADR-008.
- Confirm W017 T07 Requirement and Specification trace matches the accepted ADR decisions.
- Confirm C-027 through C-030 remain downstream DRMCP routes.
- Confirm migration design readiness is not treated as migration execution.
- Confirm the reviewed trace is recorded only in this Task's Evidence.
- Confirm future writable records are exactly this Task and `PRODUCT-WORK-SPEC-017`.
- Confirm no W016 lifecycle record changes.
- Confirm completed decision Tasks, ADRs, the Requirement, and Specifications remain read-only.

## Evidence

### Closure execution state

- Amendment authority: corrected `skills/design-decision-workflow` skill.
- DRMCP operating mode: non-operational under the current agent authoring policy.
- Filesystem fallback reason: DRMCP retrieval and authoring transactions are not operational.
- Historical invalid contract: rewrite completed decision Tasks and transition decision entries to `recorded`.
- Corrected contract: record reviewed trace only in closure-synchronization Task Evidence.
- Closure execution: completed.
- W017 T11 status: `done`.
- `PRODUCT-WORK-SPEC-017` status: `done`.
- Work Item completion conditions: satisfied.
- W017 decision entries remain `decided`.
- Completed decision Tasks: unchanged.
- ADRs: unchanged.
- Requirement: unchanged.
- Specifications: unchanged.
- W016 lifecycle: unchanged.
- Production implementation: not performed.
- Stage or commit: not performed.

### Accepted closure authority

- W017 T08 verdict: `NEEDS REVISION`.
- W017 T09 status: `done`.
- W017 T10 overall verdict: `PASS`.
- `F-MAJ-01`: `CLOSED`.
- `F-MAJ-02`: `CLOSED`.
- W017 closure acceptance gate: satisfied.
- Closure execution: completed.

### Changed records

- `PRODUCT-TASK-SPEC-017-11`
- `PRODUCT-WORK-SPEC-017`

### Read-only trace records

- `PRODUCT-TASK-SPEC-017-02`
- `PRODUCT-TASK-SPEC-017-04`
- `PRODUCT-TASK-SPEC-017-05`
- `PRODUCT-TASK-SPEC-017-06`
- `PRODUCT-TASK-SPEC-017-07`
- `PRODUCT-ADR-SPEC-001`
- `PRODUCT-ADR-SPEC-007`
- `PRODUCT-ADR-SPEC-008`
- `PRODUCT-TASK-SPEC-017-08`
- `PRODUCT-TASK-SPEC-017-09`
- `PRODUCT-TASK-SPEC-017-10`

### Exact canonical trace inputs

Accepted Requirement:

- `PRODUCT-REQ-SPEC-006`

Exact accepted ADR refs:

- `PRODUCT-ADR-SPEC-001`
- `PRODUCT-ADR-SPEC-007`
- `PRODUCT-ADR-SPEC-008`

| ADR ref | Requirement and W017 T07 Specification refs |
|---|---|
| `PRODUCT-ADR-SPEC-001` | `spec:product.design_records.artifact_model.artifact_responsibility_matrix`; `spec:product.design_records.traceability.metadata_schema`; `spec:product.design_records.traceability.resolve_and_validation` |
| `PRODUCT-ADR-SPEC-007` | `spec:product.design_records.authoring_standards.requirement_authoring`; `spec:product.design_records.authoring_standards.work_item_authoring`; `spec:product.design_records.authoring_standards.task_authoring`; `spec:product.design_records.authoring_standards.artifact_boundary`; `spec:product.design_records.artifact_model.artifact_responsibility_matrix`; `spec:product.design_records.traceability`; `spec:product.design_records.traceability.metadata_schema`; `spec:product.design_records.traceability.artifact_refs`; `spec:product.design_records.traceability.resolve_and_validation` |
| `PRODUCT-ADR-SPEC-008` | `spec:product.design_records.authoring_standards.requirement_authoring`; `spec:product.design_records.authoring_standards.work_item_authoring`; `spec:product.design_records.authoring_standards.task_authoring`; `spec:product.design_records.traceability.metadata_schema`; `spec:product.design_records.traceability.resolve_and_validation` |

### Downstream DRMCP obligations

| ID | downstream obligation |
|---|---|
| C-027 | Consume the accepted metadata transition in parser, writer, and migration design and implementation. |
| C-028 | Derive direct Requirement reverse sets separately from transitive traversal in indexing, lookup, and projection. |
| C-029 | Resolve Task owners and normalize Task refs for provenance graph analysis without changing PRODUCT semantics. |
| C-030 | Define diagnostics, response schemas, and user-visible projections for PRODUCT-owned invalid conditions. |

### Migration boundary

- Migration design readiness: satisfied.
- Existing-record migration semantics are design-complete in `PRODUCT-ADR-SPEC-008` and the accepted Specifications.
- Migration command, transaction, rollback, resume, diagnostics, and release gates remain downstream DRMCP scope.
- Migration execution: not performed.
- Migration execution remains outside W017.
- No existing record migration is performed by this Task.
- The historical W017 T07 statement about `migrated_to_spec: null` remains unchanged and is not current closure writer authority.
- W016 T11 remains the separate W016 lifecycle and Evidence closure owner.
