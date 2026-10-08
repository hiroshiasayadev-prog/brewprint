# DRCLI-TASK-CLI-002-16: Correct authoring machine-readable diagnostic contracts

- **status**: done
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: correction
- **estimate**: 1d
- **depends_on**:
  - DRCLI-TASK-CLI-002-15
- **outputs**:
  - spec:drcli.design_records_cli.operations.authoring.create
  - spec:drcli.design_records_cli.operations.authoring.update
  - spec:drcli.design_records_cli.operations.authoring.write
  - spec:drcli.design_records_cli.operations.authoring.retry_cache
  - spec:drcli.design_records_cli.operations.authoring.deferred_proposal
  - spec:drcli.design_records_cli.diagnostics
  - spec:drcli.design_records_cli.cli.output_and_exit_status

## Goal

Close DRCLI-T10-F-MAJ-001 by making authoring negative outcomes stable and machine-readable without changing accepted authoring semantics.

## Work

Correct only the direct finding surface:

- authoring create Response/Errors;
- authoring update Response/Errors;
- authoring write Response/Errors;
- retry-cache Response/Errors;
- deferred-proposal Response/Errors;
- DRCLI diagnostics catalog/envelope where the shared machine-readable identity belongs;
- CLI output/exit mapping where deterministic process-status classification belongs.

For every existing authoring failure class, define enough stable machine-readable contract to determine:

- stable diagnostic/result identity;
- deterministic placement/envelope;
- required context needed to distinguish the condition;
- whether the invocation produced a trustworthy negative semantic result (contract_result_failure) or failed to produce a trustworthy result (execution_failure);
- interaction with body_cache_id when retryable submitted body was preserved.

Preserve PRODUCT-owned validation findings as PRODUCT-owned.
Do not redefine validation finding codes.
Do not change direct-write, candidate validation, stale-target guard, collision guard, retry-cache lifecycle, proposal lifecycle, or proposal repository-binding semantics.
Do not invent implementation technology, storage layout, retention duration, or new user-facing authoring policy.

If any existing authoring failure condition is semantically ambiguous enough that classification would require new user judgment, stop BLOCKED and name that exact condition instead of choosing.

## Done condition

- DRCLI-T10-F-MAJ-001 is fully projected into canonical Specification.
- Every authoring failure class has stable machine-readable identity and deterministic exit-category treatment.
- JSON mode can represent all authoring negative outcomes without implementation-defined identity.
- Existing positive authoring and validation semantics are unchanged.
- No unrelated Specification area changes.

## Verification

- Trace every Error row in the five authoring child contracts to a stable machine-readable identity and exit category.
- Confirm unknown/expired body-cache IDs, stale write guards, candidate validation failure, invalid authoring request, proposal lifecycle failures, and state/I/O failures are covered.
- Confirm PRODUCT validation findings are referenced, not redefined.
- Confirm only direct finding-related Specification files plus this Task record changed.

## Evidence

- **Correction target:** DRCLI-T10-F-MAJ-001 from T10 independent review, using T09 final integration Evidence and T02 accepted D-008 through D-010, D-012, and D-014. T15 coordination was already `done`; no new user-owned decision was needed.
- **Writable boundary:** corrected exactly five authoring child contracts, the two DRCLI diagnostic contracts (catalog and envelope), this CLI output/exit contract, and this T16 record. DRMCP and the parent Work Item were not edited. No implementation or Git operations were performed.
- **Canonical output files (8):**
  - `drcli/records/spec/design-records-cli/operations/authoring/create.md`
  - `drcli/records/spec/design-records-cli/operations/authoring/update.md`
  - `drcli/records/spec/design-records-cli/operations/authoring/write.md`
  - `drcli/records/spec/design-records-cli/operations/authoring/retry-cache.md`
  - `drcli/records/spec/design-records-cli/operations/authoring/deferred-proposal.md`
  - `drcli/records/spec/design-records-cli/diagnostics/operation-diagnostic-catalog.md`
  - `drcli/records/spec/design-records-cli/diagnostics/operation-diagnostic-envelope.md`
  - `drcli/records/spec/design-records-cli/cli/output-and-exit-status.md`
- **Machine identities:** the diagnostic catalog now registers 12 authoring-specific stable codes, including the nested cache-preservation failure code, with closed reason/guard/state variants, exact context fields, and no implementation-defined code additions. The existing non-authoring catalog is retained.
- **Negative placement:** trustworthy semantic authoring rejections are one top-level `error` JSON result on stdout/status 1; unavailable cached/proposal state and untrustworthy preparation, state I/O, and write-completion failures are one top-level `error` JSON diagnostic on stderr/status 3. CLI shape and body-option errors remain actionable usage diagnostics/status 2; unexpected invariant/software faults remain status 4. There is never a second JSON document for cache preservation failure.
- **Retry and write evidence:** a successfully retained retry body uses `error.context.body_cache_id`; a failed attempt to preserve it uses distinct nested `error.context.cache_error`. A body-cache ID is not invented or used as repository binding. Every primary authoring error reports actual `write_state`, and verified pre-write nonmodification reports empty `modified_refs`; a post-write execution failure cannot falsely claim no modification.
- **PRODUCT boundary:** candidate failures retain the owning canonical validation finding's `classification`, `code`, `message`, `contract_ref`, and `subject`, grouped by affected candidate record. No PRODUCT validation code or conformance rule was redefined. Nonaffected repository diagnostics remain excluded.
- **Preserved authoring behavior:** direct write/default, valid update no-op, optional deferred proposal creation, process-independent opaque cache/proposal IDs, repository binding, retained proposal states, candidate validation, stale/collision guards, and write-state truthfulness remain as previously specified.
- **Error-row check:** reread all five corrected authoring child Specifications and mechanically checked all 42 `## Errors` table rows (create 10; update 11; write 8; retry-cache 5; deferred proposal 8). Each row has an explicit machine-readable authoring code or CLI usage/error delegation, an exit category, and no unregistered `authoring_*` code reference. The eight corrected Specifications have their substantive-change date set to 2026-10-08 and retain required structure.
- **Review handoff:** T16 correction is complete and ready for independent T17 finding-closure review. T16 does **not** independently declare DRCLI-T10-F-MAJ-001 closed, does not change T10 review Evidence, and does not synchronize T11 or the parent Work Item.
- **Blockers:** none.
