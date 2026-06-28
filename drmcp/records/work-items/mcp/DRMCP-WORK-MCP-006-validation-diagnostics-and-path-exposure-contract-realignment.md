# DRMCP-WORK-MCP-006: Validation, diagnostics, and path-exposure contract realignment

- **id**: DRMCP-WORK-MCP-006
- **status**: in_progress
- **date**: 2026-06-28
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-INV-MCP-002
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-004
  - DRMCP-WORK-MCP-005
  - DRMCP-TASK-MCP-001-06
  - spec:product.design_records.traceability
  - spec:product.design_records.spec_format
  - spec:drmcp.design_records_mcp.responsibility_boundary
  - spec:drmcp.design_records_mcp.schema.diagnostics
  - spec:drmcp.design_records_mcp.schema.record_model
  - spec:drmcp.design_records_mcp.schema.authoring_transaction_schema
  - spec:drmcp.design_records_mcp.tools.list_records
  - spec:drmcp.design_records_mcp.tools.get_records
  - spec:drmcp.design_records_mcp.tools.resolve_reference
  - spec:drmcp.design_records_mcp.tools.validate_records
  - spec:drmcp.design_records_mcp.tools.propose_record_create
  - spec:drmcp.design_records_mcp.tools.propose_record_update
  - spec:drmcp.design_records_mcp.tools.authoring_transaction_model
  - spec:drmcp.design_records_mcp.tools.accept_proposed_write
- **tasks**:
  - DRMCP-TASK-MCP-006-01
  - DRMCP-TASK-MCP-006-02
  - DRMCP-TASK-MCP-006-03
  - DRMCP-TASK-MCP-006-04

## Goal

Establish the corrected current-repository validation, machine-readable diagnostic, and exceptional physical-path exposure contracts.

Separate PRODUCT-owned semantic invalidity from DRMCP-owned validation execution and response representation.

## Boundary

This Work Item owns:

- current repository validation over W003-retained current sources, addressable records, conflict groups, and active-index state;
- current cross-namespace relation validation;
- configured current-to-legacy relation validation;
- behavior when a referenced legacy target is unavailable because fallback is disabled;
- exclusion of legacy archive records from normal current repository validation;
- mapping semantic invalidity into DRMCP diagnostic representation;
- shared machine-readable diagnostic structure and category placement;
- representation of unsupported, unresolved, duplicate, disabled-fallback, and source-format failures;
- narrow path exposure for source-location diagnostics, explicit patch output, and debug or emergency inspection;
- synchronization of validation, diagnostic, exceptional-path, and responsibility contracts;
- scoped validation and independent review for this contract boundary.

This Work Item does not own:

- canonical identity, relation, traceability, or spec-format semantics;
- discovery conditions or duplicate detection rules owned by `DRMCP-WORK-MCP-003`;
- tool-specific query and retrieval behavior owned by `DRMCP-WORK-MCP-004`;
- normal listing and exact-retrieval response projection, including physical-path hiding, owned by `DRMCP-WORK-MCP-004`;
- resolution order and fallback eligibility owned by `DRMCP-WORK-MCP-005`;
- individual warning-trigger behavior already owned by an operation contract;
- PRODUCT validation-policy owner-pointer updates;
- disposition of `DRMCP-WORK-SPEC-001` or `DRMCP-WORK-SPEC-002`;
- fixture authoring;
- validator, diagnostic, or response implementation;
- authoring transaction behavior.

Diagnostic representation may cite PRODUCT semantic authorities.
Diagnostic representation must not redefine those authorities.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Source Requirement for validation execution, diagnostics, and path exposure. |
| `DRMCP-ADR-MCP-001` | Governs current validation scope, archive exclusion, and path hiding. |
| `DRMCP-INV-MCP-002` | Supplies validation, diagnostic, and response-boundary findings. |
| PRODUCT traceability and spec-format authorities | Supply semantic validity without defining DRMCP response representation. |
| `DRMCP-WORK-MCP-003` | Supplies current source, addressable-record, validation-only source, duplicate-conflict, source-format, provenance, and active-index state. |
| `DRMCP-WORK-MCP-004` | Supplies normal listing, retrieval, and operation warning boundaries. |
| `DRMCP-WORK-MCP-005` | Supplies resolver outcomes and fallback states. |
| `spec:drmcp.design_records_mcp.schema.diagnostics` | Correct common diagnostic representation and category ownership. |
| `spec:drmcp.design_records_mcp.tools.validate_records` | Correct current repository validation behavior. |
| Normal list and retrieval contracts | Supply W004-owned normal physical-path hiding; W006 must not redefine their response projection. |
| DRMCP responsibility boundary | Synchronize semantic-owner and execution-owner statements. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Authority and affected-file confirmation | `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-005` | Confirm semantic authorities, operation inputs, and W007 exclusions. |
| B. Validation execution contract | Phase A | Define current-source, current-record, conflict-state, cross-namespace, and current-to-legacy validation scope without validating legacy archive records as subjects. |
| C. Diagnostic representation contract | Phases A-B | Define common machine-readable representation and semantic-invalidity mapping. |
| D. Exceptional path-exposure contract | Phases A-C | Define source-location diagnostic, patch, debug, and emergency exposure without redefining W004-owned normal response hiding. |
| E. Cross-spec synchronization and review | Phases B-D | Synchronize affected contracts, validate, review, correct, and close. |

Validation Work Item disposition proceeds through a separate cross-owner gate.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Confirm affected specs, semantic authorities, W003-W005 inputs, and the W007 disposition boundary. | W003-W005 boundaries available. |
| T02 | Correct current repository and relation-validation execution contracts. | T01. |
| T03 | Correct machine-readable diagnostic representation and semantic-invalidity mapping. | T01-T02. |
| T04 | Correct source-location diagnostic, patch, debug, and emergency path-exposure exceptions while preserving W004-owned normal response hiding. | T02-T03. |
| T05 | Synchronize contracts, confirm no W007 overlap, validate, review, correct, and close. | T04. |

## Completion Condition

This Work Item is complete when all of the following are true:

- current repository validation consumes W003-retained current sources, addressable records, validation-only sources, conflict groups, and active-index state;
- current cross-namespace relations resolve across configured current roots;
- configured current-to-legacy relations have explicit validation behavior;
- disabled legacy fallback produces an explicit diagnostic outcome;
- legacy archive records are excluded from normal current repository validation;
- semantic invalidity and DRMCP diagnostic representation remain separate;
- required failures have machine-readable diagnostics;
- W004 remains the sole owner of physical-path hiding in normal listing and exact-retrieval responses;
- W006 path exposure is limited to explicit source-location diagnostics, patch output, and debug or emergency surfaces;
- W007 owns DRMCP validation Work Item disposition; `PRODUCT-WORK-SPEC-015` owns PRODUCT validation-policy owner-pointer synchronization;
- fixtures and implementation remain delegated;
- all changed specs pass scoped validation;
- independent review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records changed files, validation results, review verdict, and residual limitations.

## Evidence

- `DRMCP-ADR-MCP-001`: Accepted validation-scope, diagnostic-ownership, and path-hiding direction.
- `DRMCP-REQ-MCP-001`: Source Requirement.
- `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-005`: Upstream contract owners.
- `DRMCP-TASK-MCP-001-06`: Hub lifecycle gate for this Work Item.
- 2026-06-27 planning-record correction: removed overlapping ownership of normal list and exact-retrieval path hiding; W004 remains the operation-response authority and W006 owns only exceptional path exposure.
- `DRMCP-TASK-MCP-006-01` opened on 2026-06-28.
  - Exact task-directory discovery found no existing `DRMCP-TASK-MCP-006-*` record.
  - T01 records the authority matrix, W003-W005 ownership inputs, contradiction inventory, candidate changed-file manifest, recheck-only manifest, T02-T05 split, open design questions, and independent review prompt.
  - T01 corrected the planning boundary from addressable active-index entries only to W003-retained current sources, records, conflict groups, and active-index state.
  - Legacy archive records remain excluded as repository-validation subjects.
  - No normative DRMCP spec changed during T01.
  - T01 is `done`; its baseline is the accepted input to T02 through T05.
- 2026-06-28 independent baseline review verdict: `NEEDS REVISION` with one minor finding, `F-MIN-01`.
  - The finding identified four stale `T07` ownership labels in this Work Item.
  - The four labels now identify `W007` as the DRMCP disposition boundary.
  - The Completion Condition now separates `W007` ownership of DRMCP validation Work Item disposition from `PRODUCT-WORK-SPEC-015` ownership of the PRODUCT validation-policy owner-pointer update.
  - Limited independent re-review verdict: `PASS`; `F-MIN-01` is closed with no remaining findings or advisories.
  - External `git diff --check` passed for the T01 Task, this Work Item, and hub T06; the LF-to-CRLF message for this file is non-blocking.
- T01 closure accepted on 2026-06-28.
  - Authority, upstream ownership inputs, contradiction inventory, candidate manifests, and T02-T05 split are established.
  - No normative DRMCP spec changed during T01.
  - Repository-wide clean status was not asserted.
- `DRMCP-TASK-MCP-006-02` opened on 2026-06-28.
  - Exact task-directory discovery confirmed that `DRMCP-TASK-MCP-006-01` was the only existing W006 child Task and that `DRMCP-TASK-MCP-006-02` did not exist.
  - T02 owns the `validate_records` request, W003-retained validation-input selection, current cross-root relation execution, configured current-to-legacy relation execution, request-versus-diagnostic boundary, repository-validation wrapper, and `ok` execution semantics.
  - D01 through D07 were accepted one at a time on 2026-06-28.
  - Accepted request scopes are repository-wide empty request, one configured `app_namespace`, or one exact current canonical `ref`; `domain`, `kind`, `id_range`, legacy-ID, and physical-path selectors are excluded.
  - Validation subjects are selected directly from W003-retained repository, app, or exact-ref state. Relation targets are looked up across all configured current roots or configured legacy lookup state but are not recursively added as validation subjects.
  - Invalid mandatory current or configured legacy roots prevent trustworthy index construction and remain startup or execution failures; partial index validation is prohibited.
  - Current relation existence uses exact repository-wide active-index lookup. Accepted legacy relations preserve disabled, unresolved, duplicate-conflict, unreadable, resolved, and unsupported distinctions as diagnostic inputs.
  - Malformed scope requests and unresolved exact-ref selectors are request errors; source, record, duplicate, and relation invalidity after successful selection are validation diagnostics.
  - The normal wrapper contains `ok`, effective `scope`, selected-subject `summary`, and one unified `diagnostics` collection. A separate validation `warnings` collection is not introduced.
  - `tools/validate-records.md` was normatively realigned to the accepted T02 contract and now cites PRODUCT semantic authorities instead of copying obsolete V01 semantics.
  - `schema/discovery.md` and `schema/record-model.md` were rechecked and require no T02 pointer change.
  - T03 retains diagnostic envelope, category, severity, ordering, and deduplication ownership. T04 retains source-location and exceptional physical-path exposure ownership.
  - External scoped strict validation passed for `tools/validate-records.md`: `[strict] All 1 file(s) OK.`
  - External `git diff --check` passed for the two tracked T02 files: `tools/validate-records.md` and this Work Item.
  - LF-to-CRLF warnings for those tracked files are non-blocking working-copy conversion notices.
  - Targeted `git status --short` confirmed those two tracked modifications plus the untracked T02 Task.
  - Independent review verdict: `NEEDS REVISION` with one minor finding, `F-MIN-01`.
    - The normal tracked-file `git diff --check` did not inspect the untracked T02 Task.
    - The earlier complete-manifest wording was corrected.
  - `F-MIN-01` correction verification executed externally after the Evidence correction:
    - tracked-file check returned `tracked_exit=0` with no whitespace error;
    - untracked Task `git diff --no-index --check -- NUL <T02 Task>` returned `untracked_exit=1` with no whitespace error;
    - exit code `1` is expected for a new file differing from `NUL`; no exit code `2` or greater occurred;
    - targeted status remained two tracked modifications plus the untracked T02 Task.
  - Limited independent re-review verdict: `PASS`.
    - `F-MIN-01` is closed.
    - No blocking, major, or minor findings remain.
    - Recorded tracked and untracked whitespace evidence is sufficient.
    - No regression was found in the normative contract or changed-file manifest.
  - Review advisory `A-01` is carried into T03: kind-specific required metadata, lifecycle, and done-gated section diagnostics should cite the applicable authoring standards when category and severity behavior is defined.
  - `DRMCP-TASK-MCP-006-02` closed as `done` on 2026-06-28.
  - T02 supplies the accepted execution baseline to T03 and later W006 phases.
  - Task and Work Item records are outside the strict spec-validator scope.
  - No repository-wide clean working tree is inferred.
- `DRMCP-TASK-MCP-006-03` opened on 2026-06-28.
  - Exact Task inventory confirmed `DRMCP-TASK-MCP-006-01` and `DRMCP-TASK-MCP-006-02`; no T03 record existed before creation.
  - T03 owns the shared machine-readable warning and diagnostic envelope, category mapping, severity, subject and occurrence associations, ordering, duplicate suppression, authoring-envelope compatibility, and the T04 handoff slot.
  - PRODUCT remains the semantic authority for metadata, lifecycle, required sections, references, and current spec invalidity.
  - T02 request forms, validation subjects, failure boundaries, relation lookup, response wrapper, and blocking-severity handoff are fixed inputs.
  - D01 through D11 were accepted one at a time on 2026-06-28 and recorded immediately in the T03 Task.
  - The accepted envelope requires `category`, `severity`, and `message`; structured optional associations are `subject`, `field`, `value`, `target`, `occurrence`, `conflict`, and the T04-owned `location` slot.
  - Repository categories now use stable cause-oriented names instead of one category per PRODUCT rule. PRODUCT authorities still determine requiredness, allowed emptiness, lifecycle, gated sections, identity, and relation validity.
  - Severity is `error`, `warning`, or `info`; `validate_records.ok` is false exactly when at least one returned diagnostic has severity `error`.
  - Current identity conflicts and legacy issued-ID lookup conflicts remain separate and aggregate every repairable source or candidate location without selecting a winner.
  - Repository and proposal-local diagnostics now have deterministic ordering and semantic duplicate suppression independent of traversal, map iteration, message wording, or detector path.
  - Proposal-local validation uses the same repository taxonomy and structured associations as `validate_records`; authoring operation categories and trigger semantics remain operation-owned.
  - `schema/diagnostics.md` was rewritten to remove stale resolver response names, `ids` and item-wrapper retrieval assumptions, scalar diagnostic shortcuts, obsolete YAML and section-ref categories, copied PRODUCT rules, and raw-path requiredness.
  - `tools/validate-records.md` now states the accepted error-severity `ok` rule and points to the shared diagnostics contract and T04 location handoff.
  - `tools/resolve-reference.md` now returns top-level `diagnostics` in every normal resolver response while preserving W005 statuses, current-first order, and successful target projection. Cause diagnostics distinguish malformed, unsupported, unresolved, disabled fallback, current conflict, legacy conflict, and unreadable legacy source states.
  - `tools/propose-record-create.md` and `tools/propose-record-update.md` realign diagnostic examples from scalar `record_id`, `field`, `value`, and standalone diagnostic `path` to structured shared associations without changing proposal behavior.
  - `schema/authoring-transaction-schema.md` and `tools/propose-record-update.md` replace retired proposal-local validation category names with `missing_required_field`, conditional `empty_required_field`, and `invalid_field_value`.
  - Final T03 normative changed-spec set is `schema/diagnostics.md`, `tools/validate-records.md`, `tools/resolve-reference.md`, `tools/propose-record-create.md`, `tools/propose-record-update.md`, and `schema/authoring-transaction-schema.md`.
  - `namespace-scanning.md`, `schema/discovery.md`, and `schema/record-model.md` were rechecked and their existing authority and location pointers remain sufficient; no T03 change was required.
  - `tools/list-records.md`, `tools/get-records.md`, `resolver.md`, `schema/record-source.md`, `schema/fields.md`, `schema/metadata-grammar.md`, and `schema/id-normalization.md` remain unchanged after recheck.
  - T02 advisory A-01 is reflected through explicit PRODUCT authoring-standard pointers for ADR, spec, investigation, requirement, work-item, and task semantic invalidity.
  - Static stale-category inspection found no positive remaining use of retired repository-validation categories; the only stale name retained in `schema/diagnostics.md` is inside an explicit removed-behavior list.
  - Pre-review external scoped strict validation passed for the complete six-file T03 normative set: `[strict]  All 6 file(s) OK.`
  - Pre-review external tracked-file whitespace verification returned `tracked_exit=0` for the six changed specs and this Work Item.
  - Pre-review external untracked Task whitespace verification returned `untracked_exit=1`; this is the expected `git diff --no-index --check -- NUL <T03 Task>` result for a new file and no whitespace error was reported.
  - Targeted status confirmed exactly six modified normative specs, this modified Work Item, and the untracked T03 Task.
  - Initial independent T03 review verdict: `NEEDS REVISION`.
    - `F-MAJ-01`: unsupported legacy relation syntax had two possible category mappings because `legacy_unsupported` remained an accepted lookup state beside the `invalid_field_value` rule.
    - `F-MIN-01`: the narrow case-only authoring fallback referenced a removed generic Required narrative section policy instead of current PRODUCT authoring authorities.
  - Review corrections applied on 2026-06-28:
    - unsupported or noncanonical declared relation values accepted by no current or legacy lookup grammar map only to `invalid_field_value`;
    - no lookup runs and no lookup state is attached for that condition;
    - `legacy_unsupported` was removed from the accepted repository lookup-state vocabulary;
    - `relation_target_unavailable` now applies only after accepted current or legacy grammar;
    - the case-only fallback now cites requirement, work-item, and task PRODUCT authoring standards directly without changing the trigger or heading set.
  - Post-correction external scoped strict validation passed for the complete six-file T03 normative set: `[strict]  All 6 file(s) OK.`
  - Post-correction external tracked-file whitespace verification returned `tracked_exit=0` for the six changed specs and this Work Item.
  - Post-correction external untracked Task whitespace verification returned `untracked_exit=1`; this is the expected `git diff --no-index --check -- NUL <T03 Task>` result for a new file and no whitespace error was reported.
  - Limited independent T03 re-review verdict: `PASS`.
    - `F-MAJ-01` is closed.
    - `F-MIN-01` is closed.
    - No blocking, major, minor, or advisory findings remain.
    - Post-correction validator and tracked/untracked whitespace evidence were accepted.
    - No regression was found in T02 execution boundaries, W005 lookup-state separation, D07-D09 determinism, authoring operation triggers, or T04 location ownership.
  - `DRMCP-TASK-MCP-006-03` closed as `done` on 2026-06-28.
  - T03 supplies the accepted diagnostic envelope, category mapping, severity, deterministic ordering, duplicate suppression, authoring-envelope boundary, and source-location handoff to T04 and T05.
- T03 is closed. W006 remains `in_progress`; the next phase is T04 exceptional source-location and physical-path exposure contract definition.
- `DRMCP-TASK-MCP-006-04` opened on 2026-06-28.
  - Exact Task inventory confirmed T01 through T03 and no existing T04 record.
  - T04 owns the concrete `location` object, portable path representation, configured-root identification, Windows handling, stable location identity and sort key, operation-specific exceptional exposure, explicit patch paths, files-written paths, and debug or emergency exposure limits.
  - T03 source-location requiredness, conflict aggregation, ordering, and duplicate-suppression requirements are fixed inputs and are not reopened.
  - W004 normal list and exact-retrieval path hiding and W005 successful resolver target hiding remain fixed non-regression boundaries.
  - Existing current-root identity is `app_namespace` plus repository-relative `records_root`; existing legacy roots expose repository-relative `records_root` without an app namespace or stable configured root ID.
  - Existing current source provenance includes `app_namespace`, record kind, and repository-relative source path.
  - Existing authoring outputs expose repository-relative paths through proposal targets, `diff.files`, unified patch text, and `files_written`; T04 may define representation and exposure without redesigning authoring transaction behavior.
  - No dedicated debug or emergency operation currently exists. T04 defines only the permitted exceptional boundary and does not add a new public tool.
  - The unconditional normative candidate is `schema/diagnostics.md`; source, root, validation, and authoring contracts remain conditional on accepted T04 decisions.
  - Discovery, record-source, normal read, resolver, overview, MVP, and responsibility contracts are recheck-only or T05 synchronization candidates unless a direct contradiction is found.
  - D01 through D07 were presented and accepted one at a time on 2026-06-28 before normative reflection.
  - Current and legacy locations now share `source_scope`, `records_root`, and `path`; current locations additionally require `app_namespace`.
  - Normal exposed paths are repository-relative, `/`-separated, canonically contained, case-preserving, and prohibit absolute, drive, UNC, URI, `.`, `..`, and alias-escape forms.
  - Location identity is `(source_scope, records_root, path)` and the stable sort key is `(source_scope_rank, records_root, path)` with locale-independent UTF-8 bytewise comparison.
  - Source-backed validation, read-warning, conflict, and applicable authoring findings expose direct or typed member/candidate locations under an explicit surface matrix; normal successful list, retrieval, and resolver projections remain path-free.
  - Proposal targets, `diff.files`, unified patch operands, and `files_written` retain scalar fields but use one normalized repository-relative spelling.
  - No current operation exposes an absolute physical path. A future privileged debug or emergency operation requires a separate contract and cannot be a hidden existing-operation flag.
  - Missing required direct, member, or candidate location fails closed. No partial entry, incomplete conflict, opaque token, absolute fallback, partial read response, retained proposal, or pre-write filesystem change is allowed.
  - Normative reflection changed `schema/diagnostics.md`, `tools/validate-records.md`, `tools/list-records.md`, `tools/get-records.md`, `tools/resolve-reference.md`, `schema/authoring-transaction-schema.md`, `tools/authoring-transaction-model.md`, `tools/propose-record-create.md`, and `tools/accept-proposed-write.md`.
  - `tools/list-records.md`, `tools/get-records.md`, and `tools/resolve-reference.md` changed only because D07 exposed a direct response-construction contradiction; no new execution-failure identifier was invented.
  - `schema/record-model.md`, `namespace-scanning.md`, `schema/discovery.md`, `schema/record-source.md`, `resolver.md`, `tools/propose-record-update.md`, and `tools/get-proposed-write.md` were rechecked and require no T04 change.
  - Overview, tool-overview, MVP, responsibility, final manifest, final review, and closure synchronization remain T05 scope.
  - External scoped strict validation passed for all nine changed T04 normative specifications: `[strict]  All 9 file(s) OK.`
  - External tracked-file `git diff --check` covered the nine changed normative specifications and this Work Item and reported no whitespace error.
  - LF-to-CRLF messages for those tracked files are non-blocking working-copy conversion notices.
  - External untracked T04 Task whitespace verification returned exit code `1` with no whitespace error; this is the expected `git diff --no-index --check -- NUL <T04 Task>` result for a new file.
  - Targeted status confirmed exactly nine modified normative specifications, this modified Work Item, and the untracked T04 Task.
  - Initial independent T04 review verdict: `NEEDS REVISION` with one minor finding, `F-MIN-01`.
    - The T04 normative contract, D01 through D07, changed-file body manifest, W003-W005 non-regression, authoring path contract, and recorded verification evidence were assessed as PASS-equivalent.
    - `F-MIN-01` found that T04 Task metadata `outputs` listed only `schema.diagnostics` and W006 despite nine changed normative specifications.
  - `F-MIN-01` correction applied on 2026-06-28:
    - added the eight omitted spec refs for `validate_records`, `list_records`, `get_records`, `resolve_reference`, `authoring_transaction_schema`, `authoring_transaction_model`, `propose_record_create`, and `accept_proposed_write`;
    - T04 Task metadata now matches its normative reflection and changed-file manifest;
    - no normative specification changed.
  - Post-`F-MIN-01` correction whitespace verification was executed externally before the first limited re-review:
    - the tracked W006 Work Item check returned `tracked_exit=0`;
    - the untracked T04 Task check returned `untracked_exit=1`;
    - neither check reported a whitespace error;
    - no exit code `2` or greater occurred;
    - untracked exit code `1` is expected because the new Task differs from `NUL`.
  - First limited independent T04 re-review verdict: `NEEDS REVISION`.
    - `F-MIN-01` is closed.
    - New `F-MIN-02` found that the post-correction whitespace results were not yet recorded in T04 and W006, while W006 still described verification as pending.
  - `F-MIN-02` correction applied on 2026-06-28:
    - recorded and synchronized `tracked_exit=0` and `untracked_exit=1` with no whitespace error and no exit code `2` or greater;
    - removed the stale corrected-verification-pending statement;
    - changed no normative specification.
  - Second limited independent T04 re-review verdict: `NEEDS REVISION`.
    - `F-MIN-02` is closed.
    - New `F-MIN-03` requested final current-file whitespace results to be recorded back into the same Task and W006 files.
  - That final-evidence requirement is self-invalidating because writing the results changes the checked bytes.
  - Final verification boundary:
    - finalize Task and W006 evidence first;
    - run whitespace checks after the last evidence edit;
    - supply raw command output directly to the final limited re-review;
    - do not write the final result back into either checked file;
    - any later edit requires another external final check.
  - No normative specification changed for this closure correction.
  - Pre-closure final external whitespace verification passed with `tracked_exit=0` and `untracked_exit=1`; no whitespace error and no exit code `2` or greater were reported.
  - Final limited independent T04 re-review verdict: `PASS`.
    - `F-MIN-01` is closed.
    - `F-MIN-02` is closed.
    - `F-MIN-03` is closed.
    - No blocking, major, minor, or advisory findings remain.
    - T04 closure readiness is `READY`.
  - `DRMCP-TASK-MCP-006-04` changed to `done` on 2026-06-28.
  - Closure synchronization changes the checked Task and W006 bytes, so one post-closure external whitespace check must run after these final edits.
  - The post-closure result is supplied externally and intentionally not written back into either checked file.
  - W006 remains `in_progress`; the next phase is T05 final cross-spec synchronization, validation, review, correction, and closure.
  - No repository-wide clean working tree is inferred.
