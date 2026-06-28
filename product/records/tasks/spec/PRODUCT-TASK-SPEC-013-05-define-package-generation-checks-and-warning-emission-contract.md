# PRODUCT-TASK-SPEC-013-05: Define package generation checks and warning-emission contract

- **id**: PRODUCT-TASK-SPEC-013-05
- **status**: done
- **date**: 2026-06-26
- **work_item**: PRODUCT-WORK-SPEC-013
- **source_requirement**: PRODUCT-REQ-SPEC-003
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-013-02
  - PRODUCT-TASK-SPEC-013-03
  - PRODUCT-TASK-SPEC-013-04
- **outputs**:
  - PRODUCT-ADR-SPEC-003
  - PRODUCT-WORK-SPEC-013 operational-contract evidence update

## Goal

Define the operational checks and warning behavior required before T06 implements the package generator.

Separate successful generation, non-blocking warnings, and operational failures.

## Work

- Define pre-publication checks against the temporary package tree.
- Define the minimal post-publication confirmation.
- Define operational success conditions.
- Define whole-tree copy completeness by relative file-path equality.
- Define canonical rewrite integrity and residual source-ref handling.
- Define semantic warning classes and best-effort boundary scanning.
- Define exit codes for success, warnings, and operational failure.
- Define cleanup and rollback behavior.
- Record concurrent generation as unsupported for the first release.
- Define host-registry and working-directory independence boundaries.
- Define the minimum T06 verification cases.
- Record the adopted decision in `PRODUCT-ADR-SPEC-003`.
- Synchronize the parent Work Item relations and Evidence.
- Obtain independent review before closing this Task.

## Done condition

- `PRODUCT-ADR-SPEC-003` has status `accepted`.
- The ADR defines check phases, success, warnings, failures, cleanup, and recovery.
- The ADR defines first-release concurrency and independence boundaries.
- The ADR supplies a minimum T06 verification matrix.
- `PRODUCT-WORK-SPEC-013` lists this Task and the new ADR.
- Independent review reports no blocking or major findings.
- No generator, test, package, or `scripts/verify.bat` implementation is changed.

## Verification

- Confirm the temporary tree is fully checked before destination replacement.
- Confirm post-publication checking is limited to destination availability.
- Confirm source and generated relative file-path sets must match.
- Confirm non-Markdown files are included in copy completeness.
- Confirm only the accepted canonical prefix rewrite may change content.
- Confirm all semantic warning classes remain non-blocking.
- Confirm exit code `0` covers success with or without semantic warnings.
- Confirm exit code `1` represents operational generation failure.
- Confirm rollback failure is operational failure.
- Confirm cleanup-only failure after valid publication remains non-blocking.
- Confirm concurrent generation is unsupported without adding a lock.
- Confirm generator inputs do not include DRMCP or host-registry state.
- Confirm working-directory independence comes from script-location root resolution.
- Confirm T06 receives a bounded verification minimum.
- Confirm no implementation file changed.

## Evidence

### Accepted operational contract

`PRODUCT-ADR-SPEC-003` records the authoritative T05 decision.

| concern | accepted contract |
|---|---|
| check location | Complete copy, rewrite, and warning checks in the temporary tree. |
| publication confirmation | Confirm the destination exists and can be enumerated. |
| copy completeness | Source and generated relative file-path sets match exactly. |
| rewrite integrity | Only the accepted canonical prefix mapping may change content. |
| warning classes | Duplicate, unresolved, external, unrewritten, and source-boundary findings. |
| source-boundary scan | Best-effort explicit rules without LLM inference. |
| successful exit | `0`, with or without semantic warnings. |
| operational failure exit | `1`. |
| cleanup | Preserve a valid destination and retain recovery artifacts after rollback failure. |
| concurrency | Unsupported in the first release; no lock or waiting protocol. |
| host independence | No app registry, DRMCP state, or Brewprint process state. |
| path independence | Resolve repository paths from the command file location. |

### T06 verification handoff

T06 must verify:

1. whole-tree relative-path preservation;
2. exact canonical prefix rewriting without unrelated changes;
3. non-blocking semantic warning emission;
4. operational failure exit behavior;
5. destination preservation or rollback;
6. successful cleanup;
7. absence of DRMCP and app-registry dependencies;
8. script-location repository-root resolution.

Concurrency tests and repository-external process tests are not required for the first release.

### Artifact boundary

T05 does not implement:

- the generator;
- automated tests;
- `scripts/verify.bat` integration;
- the initial generated package;
- consumer diagnostics;
- source semantic corrections;
- warning persistence or lifecycle machinery.

### Independent review

- **verdict**: PASS
- **reviewer result**:
  - no blocking findings;
  - no major findings;
  - no minor findings;
  - no advisories.
- **artifact checks**:
  - `PRODUCT-ADR-SPEC-003` is accepted, depends on `PRODUCT-ADR-SPEC-002`, and has no supersession;
  - this Task depends on T02-T04 and references `PRODUCT-ADR-SPEC-003` as the authoritative decision;
  - `PRODUCT-WORK-SPEC-013` lists this Task and `PRODUCT-ADR-SPEC-003`;
  - `PRODUCT-REQ-SPEC-003` lists `PRODUCT-WORK-SPEC-013`;
  - no generator, test, package, or `scripts/verify.bat` work was pulled into T05.
- **validation note**: `validate_spec.py` emitted only `H1_FORMAT` warnings on the reviewed ADR, Task, and Work Item. The artifact authoring guides require public-ID H1 forms, so the reviewer did not classify those warnings as T05 findings.
- **closure**: The Done condition is satisfied. T06 may implement the accepted contract.
