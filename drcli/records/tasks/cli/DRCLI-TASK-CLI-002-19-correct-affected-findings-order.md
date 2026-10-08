# DRCLI-TASK-CLI-002-19: Correct affected-findings deterministic ordering

- **status**: done
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: correction
- **estimate**: 0.5d
- **depends_on**:
  - DRCLI-TASK-CLI-002-18
- **outputs**:
  - spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog

## Goal

Close DRCLI-T17-F-MIN-001 by fixing one stable, candidate-construction-independent order for the authoring candidate-validation failure's affected-record groups.

## Work

Read T17 finding, T16 correction, the exact diagnostic catalog section for authoring_candidate_validation_failed, and only the affected authoring contracts/Product finding-order authority necessary for consistency.

Correct the affected_findings context rule to:

- order groups by their individual affected candidate record's canonical record_ref in ascending simple string order, not by candidate affected-set enumeration, traversal, or implementation construction order;
- each group record_ref identifies that same group's affected candidate record, never an unrelated target;
- preserve PRODUCT-prescribed ordering of each record's findings array exactly as already accepted;
- preserve every PRODUCT-owned finding classification, code, message, contract_ref, and subject;
- preserve nonempty group/finding behavior and existing stable authoring diagnostic codes, closed reason variants, JSON envelope/placement, and exit categories.

Change only the diagnostic catalog and any strictly necessary direct cross-reference consistency surface if the existing contracts explicitly contradict the new rule; explain every additional file. No unrelated changes to create/update/write/proposal behavior or validation semantics.

This is a specification correction, not implementation. Do not change DRMCP, PRODUCT, T10/T17 review Evidence, or the parent Work Item. No Git operations.

If the requested fix cannot be made without a new user-owned choice or changing accepted semantics, stop BLOCKED and name the precise issue.

## Done condition

- All conforming implementations emit identical cross-record group ordering for the same affected canonical record refs.
- Per-record PRODUCT finding order and identities are untouched.
- DRCLI-T17-F-MIN-001's required outcome is addressed without reopening the already passed 42 authoring Error-row mappings.

## Verification

- Check multi-record affected sets in different construction orders yield the same group order.
- Check individual record_ref ownership, nonempty arrays, and PRODUCT finding order.
- Check stable authoring catalog codes, JSON location, and exit mappings are preserved.
- Report exact changed files and set this Task to done with substantive Evidence.

## Evidence

- **Input and scope:** Corrected DRCLI-T17-F-MIN-001 from T17's independent review after T18 completed routing. Used T16 correction Evidence, the exact diagnostic catalog, DRCLI authoring/validation and output-envelope contracts, and directly relevant PRODUCT authoring/validation ownership. No new user-owned choice or ADR change was required.
- **Canonical correction:** Updated only `drcli/records/spec/design-records-cli/diagnostics/operation-diagnostic-catalog.md`, within `Authoring failure codes and context`. The `authoring_candidate_validation_failed.context.affected_findings` groups are now ordered by each affected candidate record's own canonical `record_ref` in ascending simple string order, independent of candidate affected-set enumeration, traversal, or candidate construction. Each group's `record_ref` identifies the record owning that group.
- **Preserved finding semantics:** Each `findings` array remains nonempty and in that record's PRODUCT-prescribed canonical order. PRODUCT-owned `classification`, `code`, `message`, `contract_ref`, and `subject` remain unchanged. The affected-record grouping remains nonempty and excludes nonaffected repository findings.
- **Focused verification:** Inspected the corrected catalog and direct consumers. Three distinct construction permutations for a three-record example all produced the same canonical `record_ref` order (a, b, c), with each original per-record findings sequence and nonempty group/finding membership retained. The previous `in candidate affected-set order` rule is absent; no directly contradictory consumer contract required modification.
- **Preserved contracts:** Verified that the diagnostic catalog still registers exactly 12 stable `authoring_*` codes, including nested-only `authoring_cache_preservation_failed`. Read the five authoring operation `Errors` tables and confirmed the existing 42 rows remain (create 10, update 11, write 8, retry-cache 5, deferred proposal 8). Existing single-JSON placement, status 1/2/3/4 mappings, direct-write, cache, proposal, `body_cache_id`, write-state and guard contracts were not edited.
- **Changed files (exactly two):** `drcli/records/spec/design-records-cli/diagnostics/operation-diagnostic-catalog.md` and `drcli/records/tasks/cli/DRCLI-TASK-CLI-002-19-correct-affected-findings-order.md`.
- **Exclusions and gate:** No implementation or Git operations. No DRMCP, PRODUCT, T10/T17 review, parent Work Item, or T20/T11 changes. T19 corrects but does not independently close DRCLI-T17-F-MIN-001 or DRCLI-T10-F-MAJ-001; T20 owns independent closure review. **Blockers:** none.
