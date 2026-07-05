# DRMCP-TASK-MCP-016-16: Independently review application-architecture finding closure

- **id**: DRMCP-TASK-MCP-016-16
- **status**: done
- **date**: 2026-07-04
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: review
- **estimate**: 1.0d
- **depends_on**:
  - DRMCP-TASK-MCP-016-15
- **outputs**:
  - DRMCP-TASK-MCP-016-16

## Goal

Independently decide whether T09 findings F-BLK-01 and F-MAJ-01 are closed by the revised authority and canonical Specification state.

## Work

- Establish independence from T12 through T15 decision, routing, and authoring.
- Review T09 F-BLK-01 and F-MAJ-01 from current full text rather than author summaries.
- Review D-018 through D-024 and T13 ADR routing.
- Review DRMCP-REQ-MCP-003 and ADR-001, ADR-007 through ADR-012.
- Review W016 and T10 graph synchronization.
- Review every Specification changed by T15:
  - the five `spec:drmcp.application_architecture` topic files;
  - `spec:drmcp.design_records_mcp.namespace_scanning`;
  - the Schema Overview, discovery, identity, record-source, and Guidance-source Specifications;
  - `list_authoring_guides` and `get_authoring_guidance`.
- Verify F-BLK-01 closure:
  - portable package authority remains consistent with PRODUCT-REQ-SPEC-003 and ADR-001;
  - the package is a configured spec-tree Current Records source under `design_records`;
  - normal current Spec and active-index semantics are reused;
  - no runtime namespace rewrite or package-specific record model remains;
  - ADR-010 through ADR-012 validly supersede ADR-007 through ADR-009.
- Verify F-MAJ-01 closure:
  - canonical current authority uses `## What this is` for Guidance abstract projection;
  - no current assigned Specification retains `## Abstract` extraction;
  - completed T03 and T09 remain historical Evidence rather than being rewritten.
- Verify Guidance list and get are fixed-scope Application aliases over shared record-query orchestration, not public-use-case chaining.
- Verify canonical package refs are public guide IDs and conflicted or unprojectable records are not silently treated as successful catalog entries.
- Decide each named finding as `CLOSED` or `OPEN` with exact Evidence.
- Record only regressions caused or directly exposed by T14 or T15.
- Remain read-only and do not synchronize W016 closure.

## Done condition

- F-BLK-01 has one independent `CLOSED` or `OPEN` disposition.
- F-MAJ-01 has one independent `CLOSED` or `OPEN` disposition.
- Direct cross-artifact consistency effects are reviewed.
- Any new finding is limited to a regression caused or directly exposed by T14 or T15.
- The exact release result for T10 is explicit.

## Verification

- Confirm T12 through T15 are `done` and T16 is independent of their authoring.
- Confirm author Verification, Evidence conclusions, and `PASS` claims are not used as proof.
- Confirm current full text and scoped Git Evidence are inspected directly.
- Confirm DRMCP-REQ-MCP-003 retains the portable-package need while replacing package-specific indexing with normal Current Records treatment.
- Confirm ADR-001 remains accepted and unchanged.
- Confirm ADR-007, ADR-008, and ADR-009 are `superseded` without historical-body rewriting.
- Confirm ADR-010, ADR-011, and ADR-012 are accepted with exact supersession and dependency metadata.
- Confirm five-component, dependency, Current Records, and Guidance alias Specifications agree.
- Confirm spec-tree source mapping derives `spec:design_records` from the package root `index.md`.
- Confirm Guidance list excludes the authoring-standards root and uses the canonical child subtree.
- Confirm list and detail identity is the canonical package Spec ref.
- Confirm first-H1 title, `## What this is` abstract, and verbatim content projection are coherent.
- Confirm no direct PRODUCT-directory source, filename-stem guide identity, separate Guidance Domain, separate Guidance Source, package-specific index, or `## Abstract` projection remains in current assigned artifacts.
- Confirm T09 remains historical `NEEDS REVISION` Evidence.
- Confirm T10 depends on T16 and is not executed by this Task.
- Confirm reviewed artifacts are not edited, staged, or committed.
- Record exact `CLOSED` or `OPEN` disposition for F-BLK-01 and F-MAJ-01.

## Evidence

### Overall verdict

`NEEDS REVISION`

F-BLK-01 and F-MAJ-01 are independently `CLOSED`.
One new Major regression directly exposed by T15 remains open, so T10 is blocked.

### Review precondition

`READY`

- T12, T13, T14, and T15 are `done`.
- T16 was `not_started` before this review and depends on T15.
- T10 depends on T16.
- W016 is `in_progress`.
- T09 is `done` and retains its historical `NEEDS REVISION` verdict and named findings.
- Every named review artifact exists in the current worktree.
- No reviewed artifact has a staged change.

### Reviewer independence

- This review did not participate in T12 through T15 decision, routing, Requirement authoring, ADR authoring, or Specification authoring.
- T12 through T15 self-verification, Evidence conclusions, author reports, prior-session summaries, and stale-authority grep claims were not used as proof.
- Current full text and scoped Git Evidence were inspected directly.
- Reviewed artifacts were not edited.
- Only T16 was written.

### Access mode

DRMCP is non-operational, so Design Records MCP could not be used.
Filesystem fallback was used under the repository startup instruction.

### Reviewed artifacts

- W016 and T09 through T16.
- PRODUCT-REQ-SPEC-003, DRMCP-REQ-MCP-003, and DRMCP-REQ-MCP-005.
- DRMCP-ADR-MCP-001 and ADR-007 through ADR-012.
- All five `spec:drmcp.application_architecture` files.
- `namespace-scanning`, Schema Overview, discovery, ID normalization, record source, and authoring guidance source.
- `list_authoring_guides` and `get_authoring_guidance`.
- Required PRODUCT authoring standards and the design-convergence review gate.
- T03 historical Guidance decisions only to distinguish historical Evidence from current authority.

### Scoped Git Evidence

- Scoped worktree and staged/unstaged/untracked patch inspection covered W016, T09 through T16, DRMCP-REQ-MCP-003, ADR-007 through ADR-012, the five application-architecture Specifications, and every Current Records and Guidance Specification named by this Task.
- Returned patch size was 187,261 bytes from a total of 187,261 bytes. No diff truncation occurred.
- No staged change exists in the reviewed scope.
- Whitespace inspection returned no finding.
- The scope contains tracked modifications and untracked artifacts. Repository-wide cleanliness was not inspected and is not claimed.
- ADR-007 through ADR-012, W016, and T09 through T16 are untracked. Git therefore cannot prove a byte-for-byte status-only historical ADR change. Current full text shows ADR-007 through ADR-009 still contain the superseded six-component, Guidance Domain/Source, and Guidance-outside-snapshot decisions rather than replacement body text.
- T09 current full text still contains the historical `NEEDS REVISION` verdict and F-BLK-01/F-MAJ-01 definitions.

### F-BLK-01 disposition

`CLOSED`

Evidence:

- PRODUCT-REQ-SPEC-003 owns the authoritative `product/records/spec/design-records/` source tree, deterministic copy to `<exe-dir>/design-records/`, fixed `design_records` namespace, canonical ref-prefix rewrite, and physical package root.
- DRMCP-REQ-MCP-003 consumes that package without runtime rewriting or host `product` namespace dependency.
- The package is one configured spec-tree Current Records source with explicit `app_namespace: design_records`.
- The package root `index.md` derives `spec:design_records`. Child paths derive `spec:design_records.*` through the normal current Spec identity rules.
- Records-root and spec-tree sources select different effective Spec roots, then reuse the same discovery, parsing, identity, logical-tree, active-index, retrieval, resolution, and validation semantics.
- Package Specs enter the normal active Current Records index with explicit app association.
- Current authority contains no package-specific parser, logical tree, index, snapshot, resolver, validator, source port, cache, or lifecycle.
- Current Records and Legacy Archive remain separate.
- Mandatory current-source failure blocks a complete active index and is not converted into partial normal data.
- ADR-001 remains `accepted`.
- ADR-007, ADR-008, and ADR-009 are `superseded`; their historical decision bodies remain visibly historical.
- ADR-010 supersedes ADR-007. ADR-011 supersedes ADR-008 and depends on ADR-010 and ADR-012. ADR-012 supersedes ADR-009.
- ADR-010 through ADR-012, DRMCP-REQ-MCP-003, and the current application-architecture Specifications consistently define five components: Composition / Lifecycle, MCP Inbound Adapter, Application Use Cases, Record Domain / Logical Tree, and Infrastructure I/O Adapters.
- Guidance Domain, Guidance Source, and the six-component model remain only in superseded ADRs and historical Tasks.

Remaining issue: none for F-BLK-01.

### F-MAJ-01 disposition

`CLOSED`

Evidence:

- DRMCP-REQ-MCP-003, ADR-011, `dependency-and-responsibility`, `authoring-guidance-source`, and `list_authoring_guides` define `abstract` as the body of `## What this is`.
- The first H1 supplies `title`.
- `get_authoring_guidance` returns the complete Markdown source verbatim as `content`.
- Current assigned Specifications contain no `## Abstract` extraction authority.
- Filename-stem lookup is explicitly rejected rather than used as identity.
- Public guide IDs are exact canonical package Spec refs.
- T03 retains its old direct-directory, filename-derived, and `## Abstract` decisions only as completed historical Evidence.
- T09 retains its historical finding text and verdict.

Remaining issue: none for F-MAJ-01.

### Guidance alias result

- `list_authoring_guides` and `get_authoring_guidance` are operation-specific Application Use Cases over shared internal record-query orchestration.
- They do not call public `list_records` or `get_records` use cases.
- Generic record tools receive no Guidance-specific request field.
- List scope is fixed to app `design_records`, kind `spec`, child subtree `spec:design_records.authoring_standards.*`, excluding the root.
- Detail accepts only one exact canonical ref in that child subtree.
- Basename, filename stem, physical path, title, fuzzy lookup, alias, candidate inference, and reference-resolution fallback are rejected.
- List ordering is canonical-ref ASCII lexical order.
- Normal responses hide physical paths.
- Duplicate identity has no winner.
- In-scope conflict, unreadability, or projection failure does not become partial success or simple not-found.
- `guide_not_found` and `guide_unavailable` remain distinct.

### Cross-artifact consistency

Passed:

- Requirement to replacement ADR authority.
- Replacement ADRs to the five application-architecture views.
- Current Records source registration, effective Spec-root identity, active-index participation, and Guidance operation scope.
- W016 graph to T10 dependency.
- T12 decision ledger to T13 routing.
- T13 routing to T14 authoring.
- T14 authority to the main T15 canonical projections.

Failed:

- `spec:drmcp.design_records_mcp.namespace_scanning` retains a stale runtime-snapshot paragraph that conflicts with ADR-012 and `spec:drmcp.application_architecture.runtime_and_state`.

### New findings

#### F-MAJ-02

- Severity: Major.
- Affected decisions: D-019 and D-023.
- Artifact: `spec:drmcp.design_records_mcp.namespace_scanning`.
- Exact location: `## Runtime snapshot boundary`.
- Observed contradiction: the section says only `list_records`, `get_records`, `resolve_reference`, and `validate_records` rebuild and consume the Current Records snapshot. It then says authoring-guidance runtime architecture is outside the boundary. ADR-012 and `spec:drmcp.application_architecture.runtime_and_state` require every Guidance request to build and consume the normal fresh immutable Current Records snapshot.
- Required outcome: revise the runtime-snapshot boundary to include `list_authoring_guides` and `get_authoring_guidance` in normal Current Records snapshot construction, preserve operation-specific Legacy loading, and remove the stale statement that excludes authoring-guidance runtime architecture.
- New user judgment required: no.
- Required owner type: correction.
- T10 blocking: yes.
- Regression boundary: directly exposed by T15 because T15 amended `namespace-scanning` for portable spec-tree Current Records integration but left the old four-operation W011 runtime boundary active.

No new Blocking finding.
No closure-blocking Minor finding.
No non-blocking Minor or Advisory finding.

### T10 release decision

`BLOCKED`

T10 cannot run while F-MAJ-02 remains open.

### Changed file

- DRMCP-TASK-MCP-016-16

### Stage and commit

- Stage: not performed.
- Commit: not performed.
