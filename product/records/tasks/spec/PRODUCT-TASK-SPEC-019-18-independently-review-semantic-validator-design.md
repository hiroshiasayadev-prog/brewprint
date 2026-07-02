# PRODUCT-TASK-SPEC-019-18: Independently review semantic validator design

- **id**: PRODUCT-TASK-SPEC-019-18
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: review
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-14
  - PRODUCT-TASK-SPEC-019-15
  - PRODUCT-TASK-SPEC-019-16
  - PRODUCT-TASK-SPEC-019-17
- **outputs**:
  - PRODUCT-TASK-SPEC-019-18

## Goal

Independently review the final combined W019 semantic validator design and issue one integrated verdict.

## Work

- Establish reviewer independence from T14 through T17 authoring.
- Review W019 and its complete Task graph.
- Trace T01, T03, T07, and T09 decisions through PRODUCT-INV-SPEC-007, T12 routing, and final canonical projection.
- Review the bounded PRODUCT-ADR-SPEC-001 root-ownership amendment and PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017.
- Verify reuse of PRODUCT-ADR-SPEC-005, PRODUCT-ADR-SPEC-006, PRODUCT-ADR-SPEC-009, PRODUCT-ADR-SPEC-011, PRODUCT-ADR-SPEC-012, and PRODUCT-ADR-SPEC-014.
- Review PRODUCT-REQ-SPEC-005, unchanged PRODUCT-REQ-SPEC-007, and the T15 no-amendment disposition.
- Review `spec:product.responsibility_boundary_validator`, its parent registration, and the final `task_authoring` usage rule.
- Review PRODUCT-WORK-SPEC-020 and PRODUCT-WORK-SPEC-021 only for the accepted downstream boundaries and release order.
- Verify that current DRMCP artifacts, checklist content, and implementation remain outside W019.
- Record one `PASS`, `NEEDS REVISION`, `NOT READY`, or `BLOCKED` verdict and a complete named finding set.

This Task must not:

- edit any reviewed artifact;
- repair findings;
- amend the graph;
- author decisions, ADRs, Requirements, Specifications, checklists, or implementation;
- synchronize lifecycle;
- stage or commit changes.

## Done condition

- One independent verdict covers the final combined W019 state.
- Every terminal decision has a complete decision-to-Investigation-to-routing-to-canonical trace.
- Every material finding names severity, affected decisions, artifact and section, required outcome, judgment requirement, and owner type.
- Work Item integrity, ADR boundaries, Requirement continuity, canonical ownership, downstream boundaries, and shared-writer preservation are reviewed.
- The exact next gate is recorded.

## Verification

- Confirm all review prerequisites are complete before semantic review begins.
- Confirm the reviewer did not author T14 through T17 or modify reviewed artifacts.
- Confirm the exact review boundary is available and scoped Git inspection is complete and non-truncated.
- Confirm checklist authoring and implementation were not required for W019 design review.
- Confirm stage and commit were not performed.

## Evidence

### Verdict

`NEEDS REVISION`.

The standalone validator contract, canonical Specification, Task-authoring usage relation, and downstream Work Item split are semantically coherent.
Three Major trace and graph defects prevent W019 design closure.

### Reviewer independence

- This session did not author PRODUCT-TASK-SPEC-019-14 through PRODUCT-TASK-SPEC-019-17.
- This session did not author or modify the reviewed ADRs, Requirements, Specifications, Investigation, W019 graph, W020, or W021.
- Current full text and scoped Git Evidence were inspected directly.
- T14 through T17 Evidence, author reports, and prior-session summaries were not accepted as proof.
- This session changed only PRODUCT-TASK-SPEC-019-18.
- No finding repair, decision change, graph change, closure synchronization, checklist authoring, implementation, stage, or commit was performed.

### Reviewed artifacts

Workflow and graph:

- PRODUCT-WORK-SPEC-019.
- PRODUCT-TASK-SPEC-019-01 through PRODUCT-TASK-SPEC-019-22.
- PRODUCT-TASK-SPEC-018-19, limited to the external shared-writer prerequisite.

Investigation and Requirements:

- PRODUCT-INV-SPEC-007.
- PRODUCT-REQ-SPEC-005.
- PRODUCT-REQ-SPEC-007.
- `spec:product.design_records.authoring_standards.requirement_authoring`, limited to Requirement identity and design-decision ownership.

ADRs:

- PRODUCT-ADR-SPEC-001.
- PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017.
- Reused PRODUCT-ADR-SPEC-005, PRODUCT-ADR-SPEC-006, PRODUCT-ADR-SPEC-009, PRODUCT-ADR-SPEC-011, PRODUCT-ADR-SPEC-012, and PRODUCT-ADR-SPEC-014.

Canonical Specifications:

- `spec:product`.
- `spec:product.design_records`.
- `spec:product.responsibility_boundary_validator`.
- `spec:product.design_records.authoring_standards.task_authoring`.
- Spec-format identity, document-shape, and Topics-table authority.

Downstream boundaries:

- PRODUCT-WORK-SPEC-020 and its materialized authoring, review, and closure route.
- PRODUCT-WORK-SPEC-021.

### Decision-to-final-state trace result

| source | final projection result |
|---|---|
| T01 D-001 through D-005, D-007, D-009, and D-010 | PASS. Investigated by PRODUCT-INV-SPEC-007, reconciled by T07, routed through T12 B-001, recorded in PRODUCT-ADR-SPEC-015, and projected into `spec:product.responsibility_boundary_validator`. |
| T01 D-008 | PASS. Directly projected as required Task section references, optional excerpts, and no required line numbers. |
| T01 D-011 | PASS. Directly projected without required checklist revision IDs or stable criterion IDs. |
| T03 D-001 and T07 R-001 through R-003 | PASS for current semantics. PRODUCT-ADR-SPEC-016 and the validator Specification preserve the temporary standalone boundary, current DRMCP separation, direct PRODUCT ownership, and historical T01 treatment. ADR-001 routing remains defective under F-MAJ-02. |
| T07 R-004 | PASS. The shared `task_authoring` writer followed PRODUCT-TASK-SPEC-018-19 and preserved the accepted W018 Task-type and decomposition semantics. |
| T01 D-012 and T07 R-005 | PASS. PRODUCT-ADR-SPEC-017, the validator Specification, and `task_authoring` define the two invocation points, one shared validator contract, caller-owned enforcement, and human-owned exception disposition. |
| T07 R-006 | NEEDS REVISION. The final no-amendment Requirement boundary is substantively defensible, but it contradicts the completed R-006 and T12 amendment route without a successor decision. See F-MAJ-01. |
| T09 D-001 | PASS. W020 owns exact checklist authoring and review. W021 owns standalone implementation. W019 absorbs neither deliverable. |
| T21 D-001 and T22 | PASS for the substantive release boundary. W020 decision, Investigation, and authoring may proceed early; W020 review waits for accepted W019 review; W021 waits for W019 closure and accepted W020 review. Integrated-review ordering remains defective under F-MAJ-03. |

Additional integrated results:

- PRODUCT-REQ-SPEC-005 typed single-responsibility semantics are preserved.
- PRODUCT-REQ-SPEC-007 still states the stable semantic-validation need and does not duplicate downstream workflow policy.
- PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017 form coherent, independently supersedable boundaries.
- Reused ADRs are accepted, non-superseded, and materially cover the cited workflow and responsibility rules.
- PRODUCT-ADR-SPEC-001 adds the validator area without reversing existing PRODUCT ownership architecture.
- `spec:product.responsibility_boundary_validator` has the correct path-derived ref, parent, and exactly one parent Topics registration.
- `spec:product.design_records` neither registers nor owns the validator.
- `task_authoring` owns only the narrow usage relation and does not duplicate validator behavior.
- Current DRMCP integration is not required.
- Completed decision Tasks were preserved as historical records.

### Blocking findings

None.

### Major findings

#### F-MAJ-01

- Finding ID: F-MAJ-01
- Severity: Major
- Affected decisions: T07 R-006; T12 B-003 and canonical authoring handoff; T15 no-amendment disposition.
- Affected artifact and section: PRODUCT-TASK-SPEC-019-07 `### R-006`; PRODUCT-TASK-SPEC-019-12 `### B-003` and `### Canonical authoring handoff`; PRODUCT-TASK-SPEC-019-15 `## Goal`, `## Work`, and `### No-amendment disposition`.
- Observed contradiction or omission: T07 explicitly decided to amend PRODUCT-REQ-SPEC-007, and T12 routed that amendment into canonical authoring. T15 instead selected no amendment inside an `authoring` Task. The no-amendment result is consistent with the current Requirement boundary and Requirement-authoring authority, but no successor decision or amended routing record supersedes the completed amendment decision.
- Required outcome: Preserve T07 and T12 as historical checkpoints. Use coordination to materialize a successor decision that records the current no-amendment disposition and its authority, then route the revised combined state to independent review. Do not mechanically amend PRODUCT-REQ-SPEC-007 merely to satisfy the stale route.
- New user judgment required: no
- Required owner type: coordination

#### F-MAJ-02

- Finding ID: F-MAJ-02
- Severity: Major
- Affected decisions: T07 R-002; T12 B-002 routing result; T13 authoring route; T16 PRODUCT-root alignment.
- Affected artifact and section: PRODUCT-TASK-SPEC-019-12 `### Routing result`, `### B-002`, and `### Existing ADR reuse summary`; PRODUCT-TASK-SPEC-019-13 `## Work`; PRODUCT-TASK-SPEC-019-16 `## Work`; PRODUCT-ADR-SPEC-001 amended body and metadata.
- Observed contradiction or omission: T12 records zero existing ADR amendments and routes B-002 to a new ADR. T16 later amends accepted PRODUCT-ADR-SPEC-001, but no ADR-routing decision authorizes that amendment. The amendment itself is bounded, preserves the existing PRODUCT ownership architecture, and does not conceal a material reversal; the defect is the missing amendment route.
- Required outcome: Preserve the current bounded ADR-001 content. Use coordination to materialize an ADR-routing decision that explicitly classifies and authorizes the ADR-001 amendment under the accepted amendment-versus-supersession policy, then independently review the resulting trace.
- New user judgment required: no
- Required owner type: coordination

#### F-MAJ-03

- Finding ID: F-MAJ-03
- Severity: Major
- Affected decisions: T21 D-001; T22 cross-Work-Item graph projection; PRODUCT-ADR-SPEC-012 final-writer review rule.
- Affected artifact and section: PRODUCT-WORK-SPEC-019 `## Task flow`, `## Task Candidates`, and final Evidence; PRODUCT-TASK-SPEC-019-18 metadata `depends_on` and pre-review Evidence; PRODUCT-TASK-SPEC-019-22 outputs and completed graph write.
- Observed contradiction or omission: T22 changes W019, T19, W020, and the downstream release graph that T18 must review, but T18 depends only on T14 through T17 and still identifies T17 as the final W019 writer. The current files were available and reviewed in this session, but the persisted graph does not guarantee that integrated review follows the final graph and release-boundary writer.
- Required outcome: Use coordination to make the accepted integrated review follow T22 and every final W019 writer. Because this T18 is already complete, the route must preserve it as historical Evidence and materialize a new integrated review after the graph correction rather than rewriting this verdict.
- New user judgment required: no
- Required owner type: coordination

### Minor findings

None.

### Advisories

None.

### Scoped Git inspection

- Scoped worktree inspection: `pass`.
- Scoped diff inspection: `pass`.
- The full scoped patch was returned without truncation.
- Scoped changes include tracked unstaged and untracked files.
- Staged scoped changes: none.
- Whitespace result: `pass`; no whitespace findings.
- LF-to-CRLF conversion warnings are advisory only.
- Untracked `git diff --no-index` exit code `1` was treated as the expected file-difference result, not a finding.
- Repository-wide cleanliness is not claimed.
- Stage and commit were not performed.

### Implementation-planning readiness

`NOT READY`.

W021 remains blocked until W019 reaches accepted design closure and W020 reaches accepted checklist review.
W020 decision, Investigation, and checklist authoring may continue under T21, but W020 integrated review must not accept the checklist set before the revised W019 integrated review route is accepted.

### Exact next gate

```text
coordination
  -> materialize exact successor routes for F-MAJ-01 through F-MAJ-03
  -> successor Requirement-disposition decision
  -> ADR-001 amendment routing decision
  -> graph correction for final-writer review order
  -> new integrated independent W019 review
```

PRODUCT-TASK-SPEC-019-19 is not released.
No correction, decision, coordination, or review Task was materialized by this session.

### Filesystem fallback

DRMCP is non-operational under `spec:product.design_records.authoring_standards.agent_authoring_policy`.
Filesystem fallback was therefore used for Design Record reads and this T18 update.
