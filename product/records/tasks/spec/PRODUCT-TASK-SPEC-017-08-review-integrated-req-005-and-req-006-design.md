# PRODUCT-TASK-SPEC-017-08: Review integrated REQ-005 and REQ-006 design

- **id**: PRODUCT-TASK-SPEC-017-08
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-017
- **source_requirement**: PRODUCT-REQ-SPEC-006
- **estimate**: 0.75d
- **depends_on**:
  - PRODUCT-TASK-SPEC-016-07
  - PRODUCT-TASK-SPEC-017-07
- **outputs**:
  - PRODUCT-TASK-SPEC-017-08

## Goal

Independently judge the final integrated REQ-005 and REQ-006 design state.

Produce one cross-requirement verdict and complete finding set without correcting artifacts.

## Work

- Read the review-gate workflow authority for this phase.
- Trace W016 and W017 decisions through required ADRs and final Specifications.
- Check Task-type preservation, source-relation semantics, migration contract, stale text, and writer order.
- Check both Work Item completion conditions against the integrated state.
- Record `PASS` or `NEEDS REVISION` with classified findings.

## Done condition

- One independent verdict covers the final integrated Specification state.
- Every finding has severity, exact target, evidence, and required correction boundary.
- The reviewer did not author either Specification writer's changes.
- T09 and T10 have an explicit required or not-required disposition.

## Verification

- Confirm the reviewer did not author W016-07 or W017-07 changes.
- Confirm both Specification writers completed before review.
- Confirm every decided item traces to an ADR when required and to current Specification text.
- Confirm no finding is corrected inside this Task.

## Evidence

### Review execution state

- Verdict: `NEEDS REVISION`.
- Review date: 2026-07-01.
- DRMCP is non-operational under the current agent authoring policy. Filesystem fallback was used.
- This reviewer did not author `PRODUCT-TASK-SPEC-016-07` or `PRODUCT-TASK-SPEC-017-07` changes.
- No reviewed Requirement, Work Item, Task, ADR, or Specification was changed.
- No finding correction, lifecycle closure, migration execution, production implementation, stage, or commit was performed.
- Bootstrap authority remains `PRODUCT-TASK-SPEC-016-01` Evidence, `BOOTSTRAP-001`.
- The bootstrap workflow records remain under the pre-REQ-005/006 metadata shape and are not treated as a canonical exception.

### Reviewed artifact set

Work Items and Requirements:

- `PRODUCT-WORK-SPEC-016`
- `PRODUCT-WORK-SPEC-017`
- `PRODUCT-REQ-SPEC-005`
- `PRODUCT-REQ-SPEC-006`

Workflow records:

- `PRODUCT-TASK-SPEC-016-01` through `PRODUCT-TASK-SPEC-016-11`
- `PRODUCT-TASK-SPEC-017-01` through `PRODUCT-TASK-SPEC-017-11`

ADRs:

- `PRODUCT-ADR-SPEC-001`
- `PRODUCT-ADR-SPEC-004`
- `PRODUCT-ADR-SPEC-005`
- `PRODUCT-ADR-SPEC-006`
- `PRODUCT-ADR-SPEC-007`
- `PRODUCT-ADR-SPEC-008`

Final integrated Specifications:

- `spec:product.design_records.authoring_standards.requirement_authoring`
- `spec:product.design_records.authoring_standards.work_item_authoring`
- `spec:product.design_records.authoring_standards.task_authoring`
- `spec:product.design_records.authoring_standards.artifact_boundary`
- `spec:product.design_records.artifact_model.artifact_responsibility_matrix`
- `spec:product.design_records.traceability`
- `spec:product.design_records.traceability.metadata_schema`
- `spec:product.design_records.traceability.artifact_refs`
- `spec:product.design_records.traceability.resolve_and_validation`

### Review prerequisites

Result: `READY`.

- W016 T02 and T04 have no `open`, `in_discussion`, or unresolved `blocked` decision.
- W017 T02 and T04 are decision-complete and have no unresolved material decision.
- ADR routing covers every decided item.
- Required ADRs exist and are `accepted`.
- `PRODUCT-TASK-SPEC-016-07` and `PRODUCT-TASK-SPEC-017-07` are `done`.
- The exact nine-Specification review set is recorded.
- Shared writer order was W016 first and W017 second.
- Reviewer independence is satisfied.

### Decision-to-ADR-to-Specification trace result

REQ-005 result: `PASS`.

- D-001 through D-011 trace through `PRODUCT-ADR-SPEC-004` to the required scalar `task_type`, closed nine-value taxonomy, and type-specific outcome and completion contracts in `task_authoring`.
- D-013 through D-018 and the associated conflict dispositions trace through `PRODUCT-ADR-SPEC-005` to mandatory single responsibility, section alignment, independent completion boundaries, parent coordination, and implementation stop conditions.
- D-012 and the decision-workflow conflict dispositions trace through `PRODUCT-ADR-SPEC-006` to the Task checkpoint, conditional ADR routing, authoring boundary, and canonical design-state ownership rules.
- C-001, C-002, and C-004 are direct Specification projections and require no separate ADR.
- The implementation-only `## Implementation contract`, review-versus-verification, correction-versus-finding-closure, coordination-versus-synchronization, implementation-detail-versus-contract-decision, and decision-versus-authoring boundaries are present.

REQ-006 result: `PASS` for the Requirement, ADR, and final Specification state.

- D-001 through D-012 and D-017 trace through `PRODUCT-ADR-SPEC-007` to Work Item-only persisted `source_refs`, Task provenance through `work_item`, direct material source selection, direct Requirement reverse derivation, canonical ref reuse, invalid source states, and semantic provenance-cycle rules.
- D-013 through D-016 trace through `PRODUCT-ADR-SPEC-008` to staged repository migration and atomic Work Item, Task, and Requirement transitions.
- D-018 is covered by `PRODUCT-ADR-SPEC-001` and is projected as the PRODUCT semantic versus DRMCP mechanism ownership boundary.
- D-019 requires no ADR and is reflected by preserved W016 content after the W017 writer.
- C-001 through C-026 are represented by the corrected Requirement or current Specifications. C-027 through C-030 remain explicit downstream DRMCP routes. C-031 remains workflow-only migration scope.

Cross-requirement preservation result: `PASS`.

- W017 did not remove or weaken the W016 taxonomy, Task-type contracts, mandatory cohesion, common-section alignment, Evidence ownership, or adjacent responsibility boundaries.
- Work Item provenance generalization preserves the parent coordination boundary.
- Artifact-boundary and responsibility-matrix changes preserve Task workflow state, ADR durable rationale, and Specification current-state ownership.
- Traceability relation semantics do not redefine review, verification, decision, authoring, correction, coordination, or synchronization behavior.

Stale canonical Specification result: `PASS`.

- The nine Specifications do not retain persisted Requirement `work_items`, canonical Work Item `source_requirement`, canonical Task `source_requirement`, Task `source_refs`, Requirement reciprocity, or Task source-Requirement matching as current-state rules.
- Dependency cycles and semantic provenance cycles remain distinct.
- PRODUCT-owned semantics and DRMCP-owned mechanism details remain separated.
- The bootstrap exclusion was not generalized into canonical metadata semantics.

ADR integrity result: `PASS`.

- ADR-004 through ADR-008 materially agree with the final Specifications.
- ADR granularity is coherent and does not replace current normative Specification text.
- The partial V01 relation replacement is explicit without falsely superseding unrelated historical decisions.
- `migrated_to_spec: null` on ADR-007 and ADR-008 is expected before closure synchronization and is not a finding.

### Work Item completion-condition result

`PRODUCT-WORK-SPEC-016`:

- The canonical REQ-005 design deliverables are present and preserved by W017.
- Independent acceptance, conditional finding closure, decision `recorded` state, and final lifecycle closure correctly remain pending.
- Final closure is not ready because F-MAJ-02 leaves the integrated review and closure routing inconsistent.

`PRODUCT-WORK-SPEC-017`:

- The final Requirement, ADR, and Specification design state contains the accepted REQ-006 semantics.
- Independent acceptance, conditional finding closure, decision `recorded` state, and final lifecycle closure correctly remain pending.
- The Work Item completion contract cannot be satisfied as written because F-MAJ-01 retains a rejected Task `source_refs` model.
- Final closure is also blocked by the routing defect in F-MAJ-02.

### Blocking findings

None.

### Major findings

#### F-MAJ-01

Finding ID: F-MAJ-01

Severity: major

Affected decisions: W017 D-001, D-010, D-014; C-010, C-011, C-012

Affected artifact and section: `PRODUCT-WORK-SPEC-017`, `## Impact Scope` Task-authoring row and `## Completion Condition`

Observed contradiction or omission: The Work Item still says Task provenance is migrated to generic Task `source_refs`, and its completion condition says Work Item and Task provenance use one generic `source_refs` contract. The accepted design permits `source_refs` only on Work Items. Tasks persist no source field and expose provenance through `work_item`.

Evidence: `PRODUCT-REQ-SPEC-006` Requirement and Boundary; `PRODUCT-ADR-SPEC-007` Task provenance; `PRODUCT-ADR-SPEC-008` Task transition; `task_authoring` Metadata schema, General Task rules, and Migration rules.

Required correction outcome: Replace the stale Work Item statements with Work Item-only required non-empty `source_refs`, Task provenance through `work_item`, and removal-only Task migration. Preserve the bootstrap old-shape statements only as scoped workflow Evidence, not as canonical product behavior.

New user decision required: no

#### F-MAJ-02

Finding ID: F-MAJ-02

Severity: major

Affected decisions: W016 D-007, D-008, D-014; W017 D-019; accepted integrated-review routing

Affected artifact and section: `PRODUCT-WORK-SPEC-016` `## Task flow` and `## Task Candidates`; `PRODUCT-TASK-SPEC-016-08`; `PRODUCT-TASK-SPEC-016-11` metadata; `PRODUCT-WORK-SPEC-017` `## Task flow` and `## Task Candidates`; `PRODUCT-TASK-SPEC-017-09` `## Work` and `outputs`; `PRODUCT-TASK-SPEC-017-11` metadata

Observed contradiction or omission: W017 declares one cross-requirement integrated review owner, but W016 retains a separate final integrated review as a closure dependency. The PASS route requires T09 and T10 to remain not required with no no-op records, while W017 T11 depends only on T10, so the direct PASS-to-closure path is unreachable. When the integrated review finds Work Item or Task-graph defects, W017 T09 authorizes correction only of ADR or Specification findings and has no output targets for those workflow artifacts.

Evidence: `PRODUCT-WORK-SPEC-017` states that W017 owns the cross-requirement integrated review; T08 PASS disposition routes directly to T11 without writing T09 or T10; `PRODUCT-TASK-SPEC-017-11` declares only T10 as a dependency; W016 T08 and T11 still require a separate W016 review path; `PRODUCT-TASK-SPEC-017-09` limits its correction wording to ADRs and Specifications.

Required correction outcome: Establish one unambiguous integrated-review evidence owner consumed by both Work Item closure paths, make both PASS and NEEDS REVISION branches executable without synthetic no-op Tasks, and extend the named correction boundary to Work Item or Task-graph records when the review finding targets those records. Preserve independent correction closure and do not combine correction with re-review.

New user decision required: no

### Minor findings

None.

### Advisories

- `artifact_boundary` correctly defines conditional ADR routing, but its compact `Required flow` code block can be read as a single linear path. Making the ADR-required and ADR-not-required branches visually explicit would improve clarity; no design ambiguity results from the surrounding normative text.
- The focused Specification validator was not executed because no repo-local command execution tool is available in this session. This does not block the semantic review.

### Implementation-planning readiness

Result: `NOT READY`.

The PRODUCT semantic design is sufficiently complete for later DRMCP design work. No product behavior, ownership, metadata shape, source-selection, reverse-relation, migration-failure, cycle, compatibility, or PRODUCT-versus-DRMCP judgment remains hidden in the nine Specifications.

Implementation Task authoring must wait until F-MAJ-01 and F-MAJ-02 are corrected, independently re-reviewed, and lifecycle closure is synchronized.

Concrete parser, index, diagnostic, request and response, migration command, transaction, rollback, resume, and release-gate design remains correctly routed to downstream DRMCP work.

### T09 and T10 disposition

- T09: `required`
- T10: `required after T09`
- T09 must receive an amended exact correction boundary covering the named Work Item and Task-graph targets before modifying them.

### Exact next gate

`PRODUCT-TASK-SPEC-017-09`

### Verification result

- This review changed only `PRODUCT-TASK-SPEC-017-08`.
- T08 status `done` and verdict `NEEDS REVISION` are consistent.
- Every major finding has an exact target, evidence, correction outcome, and user-decision disposition.
- No blocking or minor finding is recorded.
- T09 and T10 are required and the next gate is T09.
- No finding was corrected and no lifecycle closure was performed.
- Scoped Git whitespace inspection passed. LF-to-CRLF warnings are not whitespace failures.
- The reviewed writer scope contained the expected nine modified Specifications and the W017 T07 record. No repository-wide clean-state claim is made.
