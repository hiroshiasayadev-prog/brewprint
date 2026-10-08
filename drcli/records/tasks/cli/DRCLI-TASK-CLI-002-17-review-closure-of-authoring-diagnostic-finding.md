# DRCLI-TASK-CLI-002-17: Review closure of authoring diagnostic finding

- **status**: done
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: review
- **estimate**: 0.5d
- **depends_on**:
  - DRCLI-TASK-CLI-002-16
- **outputs**:
  - DRCLI-TASK-CLI-002-17

## Goal

Independently determine whether correction T16 closes DRCLI-T10-F-MAJ-001 without changing accepted DRCLI semantics.

## Work

Read only:

- T10 finding DRCLI-T10-F-MAJ-001;
- T16 Evidence;
- the corrected authoring contracts;
- DRCLI diagnostics contracts;
- CLI output/exit contract;
- exact adjacent PRODUCT validation authority needed to verify ownership.

Verify:

- every existing authoring failure class has stable machine-readable identity;
- deterministic diagnostic/result placement exists;
- every negative authoring outcome has deterministic contract_result_failure versus execution_failure treatment;
- JSON mode is complete for authoring failures;
- retryable body preservation/body_cache_id remains coherent;
- PRODUCT validation findings were not redefined;
- direct-write, stale guards, proposal lifecycle, and retry-cache semantics did not change;
- no new unrelated semantic decision was introduced.

Do not edit repository files or close the parent Work Item.

## Done condition

- DRCLI-T10-F-MAJ-001 is explicitly CLOSED or remains OPEN.
- Any new finding must be caused by or directly exposed by T16.
- Closure is accepted only if no required correction remains.

## Verification

- Compare corrected contracts directly against T10's required outcome.
- Check all authoring Error rows and output/exit mappings.
- Confirm review independence and read-only behavior.

## Evidence

Review performed 2026-10-08. **Verdict: NEEDS REVISION. DRCLI-T10-F-MAJ-001: OPEN.** The principal identity/exit-category gap was corrected, but one directly introduced deterministic-context defect prevents full closure.

### Independent review boundary

- Reviewed T10's DRCLI-T10-F-MAJ-001 and T16 correction Evidence directly, not as proof of correctness.
- Read the five corrected authoring child Specs: `create`, `update`, `write`, `retry_cache`, and `deferred_proposal`; the diagnostic catalog and envelope; and CLI output/exit status.
- Cross-checked the authoring overview, CLI body input and usage diagnostics, exact-record validation finding envelope, PRODUCT authoring authority, and PRODUCT traceability/validation ownership. No reviewed Spec, PRODUCT/DRMCP artifact, other Task, or parent Work Item was edited. No Git operations or implementation.
- Independently enumerated all **42** authoring `## Errors` table rows: create **10**, update **11**, write **8**, retry-cache **5**, deferred proposal **8**. Each row names a cataloged `authoring_*` identity and an explicit exit code; **0** rows lack an identity or exit mapping and **0** reference unregistered authoring codes. The catalog contains **12** authoring codes, of which `authoring_cache_preservation_failed` is nested-only.

### Passing closure checks

- Negative authoring results have one top-level `error` with stable code and closed context discriminators. Trustworthy semantic rejections use JSON stdout/`contract_result_failure` status **1**; retained-state unavailability and result-preventing operational faults use JSON stderr/`execution_failure` status **3**; actionable CLI usage faults remain status **2**; unexpected internal faults remain status **4**. No competing failure document is permitted.
- The catalog selects a first primary error by request/check/guard ordering and distinguishes established collision/staleness from unreliable inspection or completion. `write_state`, `modified_refs=[]` for proven nonwrites, optional verified `known_modified_refs`, and `unknown` for indeterminate post-write state prevent false claims of nonmodification.
- An actually retained retry body appears as `error.context.body_cache_id`; failed preservation appears instead as nested `error.context.cache_error` with its own registered code. Neither invents a cache ID or overrides the enclosing primary exit category.
- Candidate validation findings retain the PRODUCT-owned `classification`, `code`, `message`, `contract_ref`, and `subject`; DRCLI adds only candidate grouping, and nonaffected repository diagnostics remain excluded.
- The corrected contracts preserve default direct-write, valid update no-op, optional deferred proposal lifecycle, current-state and collision guards, no force-write bypass, process-independent retry/proposal state, and repository-binding protection. No unrelated design judgment was found in the inspected correction surface.

### Finding: DRCLI-T17-F-MIN-001 (required minor, directly introduced by T16)

- **Artifact/section:** `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog`, `Authoring failure codes and context`, definition of `authoring_candidate_validation_failed.context.affected_findings`; consumed by create, update, write and deferred proposal.
- **Observation:** The new context contract orders its record groups `in candidate affected-set order`, but none of the five authoring child contracts or the authoring overview establishes a deterministic order for that affected **set**. PRODUCT permits coupled changes affecting multiple records. Consequently, two conforming candidate builders can enumerate the same affected records differently and return differently ordered `affected_findings` arrays for the same failure. Per-record PRODUCT finding order is specified; cross-record group order is not. This is a machine-readable context determinism defect, not a PRODUCT validation-code defect.
- **Required outcome:** Define one explicit, stable ordering of affected-record groups, independent of candidate-construction/traversal order, while retaining each record's PRODUCT-defined finding order and every PRODUCT finding identity. Clarify that each group `record_ref` denotes its own affected candidate record if necessary. No new user-owned product judgment required; owner: bounded correction followed by independent re-review.
- **Closure disposition:** Because deterministic machine-readable authoring failure context was an express requirement of DRCLI-T10-F-MAJ-001, that finding remains **OPEN** even though all 42 Error-row code/exit mappings now pass.

### Next gate

T11 closure synchronization is **NOT ELIGIBLE**. Repair DRCLI-T17-F-MIN-001 in the diagnostic catalog and only necessary direct consistency surfaces, then independently re-review the outstanding finding. T17 itself is complete as a read-only finding-closure review; no lifecycle synchronization performed.
