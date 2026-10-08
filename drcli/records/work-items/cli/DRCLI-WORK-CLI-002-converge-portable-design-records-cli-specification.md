# DRCLI-WORK-CLI-002: Converge portable Design Records CLI Specification

- **status**: done
- **date**: 2026-10-07
- **source_refs**:
  - DRCLI-REQ-CLI-001
  - DRCLI-TASK-CLI-001-02
- **impact_refs**: []
- **tasks**:
  - DRCLI-TASK-CLI-002-01
  - DRCLI-TASK-CLI-002-02
  - DRCLI-TASK-CLI-002-03
  - DRCLI-TASK-CLI-002-04
  - DRCLI-TASK-CLI-002-05
  - DRCLI-TASK-CLI-002-06
  - DRCLI-TASK-CLI-002-07
  - DRCLI-TASK-CLI-002-08
  - DRCLI-TASK-CLI-002-09
  - DRCLI-TASK-CLI-002-10
  - DRCLI-TASK-CLI-002-11
  - DRCLI-TASK-CLI-002-12
  - DRCLI-TASK-CLI-002-13
  - DRCLI-TASK-CLI-002-14
  - DRCLI-TASK-CLI-002-15
  - DRCLI-TASK-CLI-002-16
  - DRCLI-TASK-CLI-002-17
  - DRCLI-TASK-CLI-002-18
  - DRCLI-TASK-CLI-002-19
  - DRCLI-TASK-CLI-002-20

## Goal

Produce one complete, current, implementation-ready DRCLI Specification by adapting applicable DRMCP contracts to DRCLI-REQ-CLI-001.

The final Specification must define the observable CLI contract without implementing the product.

## Boundary

This Work Item owns:

- the DRCLI Specification tree;
- adaptation of applicable DRMCP current-record, operation, validation, guidance, and authoring contracts;
- DRCLI-specific CLI invocation, help, usage diagnostics, repository selection, ignore behavior, portable deployment, output, exit, and operational-state contracts;
- cross-spec integration;
- one integrated independent design review.

This Work Item does not own:

- production implementation;
- DRMCP modification, migration, replacement, or deprecation;
- PRODUCT Design Records semantic redefinition;
- MCP compatibility;
- UI behavior;
- implementation technology chosen only for convenience.

## Impact Scope

| target | impact |
|---|---|
| `drcli/records/spec/design-records-cli/**` | Create the complete DRCLI normative Specification. |
| DRMCP current Specification | Read-only baseline and design evidence. |
| DRMCP old guidance and authoring contracts | Read-only baseline where current DRMCP operation specs do not yet contain equivalent areas. |
| DRCLI-REQ-CLI-001 | Must remain satisfied without restating PRODUCT-owned semantics. |

## Task flow

```text
DRCLI-TASK-CLI-002-01 graph coordination [done]
  ├─ DRCLI-TASK-CLI-002-02 CLI surface/runtime decision
  ├─ DRCLI-TASK-CLI-002-03 artifact + current-record-model authoring
  ├─ DRCLI-TASK-CLI-002-04 read/navigation operation authoring
  ├─ DRCLI-TASK-CLI-002-05 validation + diagnostics authoring
  ├─ DRCLI-TASK-CLI-002-06 guidance authoring
  └─ DRCLI-TASK-CLI-002-07 authoring/write operation authoring

DRCLI-TASK-CLI-002-02
  -> DRCLI-TASK-CLI-002-08 CLI/runtime contract authoring

T03 + T04 + T05 + T06 + T07 + T08
  -> DRCLI-TASK-CLI-002-09 integration authoring

T09 current-source discovery blocker
  -> DRCLI-TASK-CLI-002-12 route blocker [done]
  -> DRCLI-TASK-CLI-002-13 current-source discovery decision
  -> DRCLI-TASK-CLI-002-14 bounded current-source spec authoring
  -> T09 resume and complete
  -> DRCLI-TASK-CLI-002-10 integrated independent review [NEEDS REVISION]
  -> T15 finding routing [done] -> T16 authoring diagnostics correction [done]
  -> T17 independent review [NEEDS REVISION; original finding OPEN]
  -> T18 finding routing [done] -> T19 deterministic affected-findings correction
  -> T20 independent finding-closure review
  -> T11 closure synchronization only after all required findings CLOSE
```

T03 through T07 are writer-disjoint and may run in parallel with T02.
T08 starts after T02.
T09 starts only after T03 through T08 complete.

Correction and finding-closure review Tasks are created only if T10 records named findings.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| T01 | coordination | Freeze writer ownership, dependencies, review order, and release waves. | none |
| T02 | decision | Fix DRCLI-specific command, help, output, exit, repository, ignore, packaging, and persistent-state observable contracts. | T01 |
| T03 | authoring | Adapt artifact contracts and current-record model from current DRMCP. | T01 |
| T04 | authoring | Adapt discovery/listing/retrieval/search and reference-resolution operation contracts. | T01 |
| T05 | authoring | Adapt validation and diagnostic contracts. | T01 |
| T06 | authoring | Adapt authoring-guidance list/detail contracts. | T01 |
| T07 | authoring | Adapt create/update/body-cache/proposal/write contracts to DRCLI requirements. | T01 |
| T08 | authoring | Write DRCLI-specific CLI/runtime contracts from T02. | T02 |
| T09 | authoring | Integrate indexes, cross-refs, terminology, and final combined Specification state. | T03-T08 |
| T10 | review | Independently review the complete DRCLI Specification. | T09 |
| T11 | synchronization | Synchronize lifecycle and Evidence only after an accepted review result. | T10 |
| T12 | coordination | Route the T09 current-source discovery blocker into exact decision and authoring owners. | T09 blocker |
| T13 | decision | Fix selected-repository current-source discovery and app-namespace resolution. | T12 |
| T14 | authoring | Project T13 into the bounded current-source/CLI Specification surface. | T13 |
| T15 | coordination | Route review finding DRCLI-T10-F-MAJ-001 into correction and independent closure review. | T10 |
| T16 | correction | Add stable machine-readable authoring diagnostic/result identities and exit-category mapping. | T15 |
| T17 | review | Independently verify closure of DRCLI-T10-F-MAJ-001. | T16 |
| T18 | coordination | Route T17-F-MIN-001 into bounded correction and independent review. | T17 |
| T19 | correction | Define deterministic affected-record group ordering without changing PRODUCT finding order. | T18 |
| T20 | review | Independently re-review T17-F-MIN-001 and underlying T10 major closure. | T19 |

### Writer map

| path family | sole writer before integration |
|---|---|
| `drcli/records/spec/design-records-cli/artifacts/**` | T03 |
| `drcli/records/spec/design-records-cli/current-record-model/**` | T03 |
| `drcli/records/spec/design-records-cli/operations/discovery-and-listing/**` | T04 |
| `drcli/records/spec/design-records-cli/operations/retrieval/**` | T04 |
| `drcli/records/spec/design-records-cli/operations/search/**` | T04 |
| `drcli/records/spec/design-records-cli/operations/reference-resolution/**` | T04 |
| `drcli/records/spec/design-records-cli/diagnostics/**` | T05 |
| `drcli/records/spec/design-records-cli/operations/validation/**` | T05 |
| `drcli/records/spec/design-records-cli/operations/guidance/**` | T06 |
| `drcli/records/spec/design-records-cli/operations/authoring/**` | T07 |
| `drcli/records/spec/design-records-cli/cli/**` | T08 |
| root `index.md`, `operations/index.md`, and post-wave consistency normalization | T09 |

## Completion Condition

- The DRCLI Specification covers every Required Outcome in DRCLI-REQ-CLI-001.
- Applicable DRMCP semantics are reused intentionally without creating DRMCP replacement semantics.
- CLI-specific contracts include discoverable help and actionable invalid-usage guidance.
- Windows/Linux, portable deployment, cwd-default repository selection, ignore behavior, guidance, validation, and authoring are fully specified.
- Body input covers inline arguments, stdin, and retry through body_cache_id.
- Process-independent cache/proposal state behavior is specified without requiring one installation per repository.
- No hidden MCP schema dependency remains.
- The final combined Specification is internally consistent; the integrated independent review and all finding-driven independent closure reviews establish that no required finding remains open.
- No production implementation is started.

## Evidence

- DRCLI-REQ-CLI-001 is accepted.
- DRCLI-WORK-CLI-001 completed Requirement framing and selected design convergence.
- T01 fixed disjoint writer boundaries and the execution waves.
- T10 integrated review required authoring diagnostic corrections; T16 addressed the machine identity and exit mapping gap, but T17 found deterministic affected-findings ordering incomplete.
- T18 routed the sole T17 required minor finding through T19 correction and T20 independent finding-closure review. T11 remains ineligible until all required findings close.
- Downstream authoring, review, correction and closure Evidence is recorded in DRCLI-TASK-CLI-002-02 through -20.

### Final design closure — 2026-10-08

- **Result:** DRCLI-WORK-CLI-002 changed from in_progress to done by DRCLI-TASK-CLI-002-11 after verifying every Completion Condition.
- **Design output:** spec:drcli.design_records_cli, with artifacts, current_record_model, diagnostics, operations, and cli topic areas. DRCLI-TASK-CLI-002-09 integrated the 82-file tree against accepted DRCLI-REQ-CLI-001. DRMCP remains a read-only baseline and PRODUCT remains the Design Records semantic authority.
- **Review history:** DRCLI-TASK-CLI-002-10 returned NEEDS REVISION with DRCLI-T10-F-MAJ-001; T16 corrected machine-readable authoring diagnostics. DRCLI-TASK-CLI-002-17 returned NEEDS REVISION with DRCLI-T17-F-MIN-001, so the major remained OPEN. T19 corrected deterministic affected-record group ordering. These historical verdicts were not rewritten.
- **Accepted independent closure:** DRCLI-TASK-CLI-002-20 returned PASS, explicitly CLOSED DRCLI-T10-F-MAJ-001 and DRCLI-T17-F-MIN-001, and reported zero new required findings and zero blockers.
- **Completion Conditions:** all nine PASS, as itemized in DRCLI-TASK-CLI-002-11 Evidence. They cover Requirement ownership; DRMCP/PRODUCT boundary; discoverable help and actionable usage errors; Windows/Linux portability; cwd/--repo and ignore; guidance, validation and authoring; inline/stdin/body_cache_id; process-independent caches/proposals and shared installation; no MCP schema dependence; accepted review closure; and no production implementation.
- **Decision/ADR:** T02 and T13 decisions terminal with ADR route not_required; no unresolved mandatory Investigation, ADR or design choice.
- **Relations:** all twenty DRCLI-TASK-CLI-002-01 through -20 records exist, point to this Work Item, and are done after T11. The Task graph and relations were not changed by closure.
- **Write boundary:** only parent status/Evidence and T11 status/Evidence were synchronized. Specifications, source Requirement, PRODUCT, DRMCP, other Task records and review verdicts were not modified. No Git operation or implementation source change was made.
