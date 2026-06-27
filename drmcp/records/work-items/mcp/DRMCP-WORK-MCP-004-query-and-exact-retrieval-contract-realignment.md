# DRMCP-WORK-MCP-004: Query and exact-retrieval contract realignment

- **id**: DRMCP-WORK-MCP-004
- **status**: done
- **date**: 2026-06-27
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-INV-MCP-002
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-006
  - DRMCP-TASK-MCP-001-04
  - spec:product.design_records.namespace_model
  - spec:product.design_records.spec_format
  - spec:product.brewprint.compatibility.legacy_id_compatibility
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.mvp_scope
  - spec:drmcp.design_records_mcp.responsibility_boundary
  - spec:drmcp.design_records_mcp.schema.fields
  - spec:drmcp.design_records_mcp.schema.record_model
  - spec:drmcp.design_records_mcp.tools.overview
  - spec:drmcp.design_records_mcp.tools.list_records
  - spec:drmcp.design_records_mcp.tools.get_record
  - spec:drmcp.design_records_mcp.tools.get_records
  - spec:drmcp.design_records_mcp.tools.suggest_next_record
- **tasks**:
  - DRMCP-TASK-MCP-004-01
  - DRMCP-TASK-MCP-004-02
  - DRMCP-TASK-MCP-004-03
  - DRMCP-TASK-MCP-004-04
  - DRMCP-TASK-MCP-004-05

## Goal

Reflect the accepted query and exact-retrieval baseline into authoritative DRMCP tool contracts.

Provide a compact active-index listing contract and one exact batch-retrieval contract without legacy leakage, fuzzy input repair, or normal physical-path exposure.

## Boundary

This Work Item owns:

- the `list_records` request, ordering, limit, result-shape, and truncation contract;
- the supported sequential record kinds for normal listing;
- the exclusion of specs and legacy archive records from normal listing;
- retirement of `get_record` in favor of one `get_records` exact-retrieval tool;
- retirement of `suggest_next_record` without a public request, response, or compatibility surface;
- accepted exact current sequential refs, active `spec:` refs, and configured legacy sequential IDs as retrieval inputs;
- exact-input behavior without app-prefix inference, case repair, whitespace repair, path interpretation, or fuzzy normalization;
- batch ordering, ordered duplicate removal, partial success, warning placement, and request-size limits;
- request-wide body inclusion and successful-record response behavior;
- normal query and retrieval response boundaries that exclude physical paths;
- synchronization of affected tool catalogs and responsibility summaries;
- scoped validation and independent review for this contract boundary.

This Work Item does not own:

- active-root discovery, current spec parsing, canonical identity derivation, or active-index construction;
- the final normalized record-model field vocabulary owned by `DRMCP-WORK-MCP-003`;
- spec-tree navigation tool design;
- current-first reference resolution or legacy fallback resolution order;
- legacy archive index construction;
- warning and diagnostic code taxonomy owned by the later validation and response-boundary Work Item;
- validation execution semantics;
- fixture authoring;
- parser, index, MCP tool, or runtime implementation;
- implementation tests;
- authoring transaction behavior.

`DRMCP-WORK-MCP-003` supplies the current active-index and normalized-record baseline.
The resolver and validation Work Items consume this Work Item without reopening its accepted tool split.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Source Requirement for query and exact-retrieval behavior. |
| `DRMCP-ADR-MCP-001` | Governs active-only listing, accepted legacy compatibility, and path hiding. |
| `DRMCP-INV-MCP-002` | Supplies stale query, retrieval, and response-contract findings. |
| `DRMCP-WORK-MCP-003` | Supplies active-index scope, current identity, and normalized record-model decisions. |
| `spec:product.design_records.namespace_model` | Supplies app, domain, artifact-kind, and sequential identity semantics. |
| `spec:product.design_records.spec_format` | Supplies active path-derived `spec:` identity semantics. |
| `spec:product.brewprint.compatibility.legacy_id_compatibility` | Supplies accepted Brewprint legacy sequential families. |
| `spec:drmcp.design_records_mcp.overview` | Synchronize navigation summaries for the corrected MVP and tools child specs. |
| `spec:drmcp.design_records_mcp.tools.list_records` | Replace broad and range-oriented listing with the accepted scoped compact query contract. |
| `spec:drmcp.design_records_mcp.tools.get_record` | Retire the single-record tool contract. |
| `spec:drmcp.design_records_mcp.tools.get_records` | Establish the sole exact-retrieval request and response contract. |
| `spec:drmcp.design_records_mcp.tools.suggest_next_record` | Replace the stale active P1 operation contract with an explicit retirement marker. |
| `spec:drmcp.design_records_mcp.tools.overview` | Synchronize the public tool catalog and tool responsibilities. |
| `spec:drmcp.design_records_mcp.mvp_scope` | Remove the retired tool and stale query assumptions from the P0 surface. |
| `spec:drmcp.design_records_mcp.responsibility_boundary` | Synchronize query and retrieval responsibilities without exposing normal paths. |
| Shared record-model specs | Consume W003 output and synchronize retrieval-only headings and optional body behavior without redefining semantic fields. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Accepted-baseline and affected-file confirmation | `DRMCP-INV-MCP-002`, `DRMCP-REQ-MCP-001`, `DRMCP-WORK-MCP-003` | Confirm the accepted query and exact-retrieval baseline, affected specs, and downstream ownership boundaries. |
| B. Compact active listing contract | Phase A | Correct `list_records` scope, inputs, ordering, limits, compact results, and truncation signal. |
| C. Exact batch-retrieval contract | Phases A-B | Retire `get_record` and correct `get_records` exact input, batching, partial success, body, and response behavior. |
| D. Cross-spec synchronization | Phases B-C | Synchronize tool catalogs, MVP scope, responsibility summaries, and shared model pointers. |
| E. Validation and review | Phases B-D | Run scoped validation, independent review, required corrections, and closure. |

Resolver, diagnostic-taxonomy, fixture, and implementation work proceeds through separately owned Work Items.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Confirm the accepted baseline, affected-file manifest, W003 inputs, and explicit T05/T06 exclusions. | W003 contract baseline available. |
| T02 | Reflect the accepted compact active-index `list_records` contract and remove unsupported range and broad-list behavior. | T01. |
| T03 | Retire `get_record` and reflect the accepted `get_records`-only exact retrieval and batch-response contract. | T01-T02. |
| T04 | Synchronize tools overview, MVP scope, responsibility boundary, and shared record-model pointers. | T02-T03. |
| T05 | Validate the changed contract set, run independent review, apply required corrections, and close the Work Item. | T04. |

Each Task remains within contract authoring and review.
Implementation and fixtures belong to later Work Items under `DRMCP-REQ-MCP-001`.

## Completion Condition

This Work Item is complete when all of the following are true:

- `list_records` requires one active app, supported sequential kind, and domain scope;
- normal listing supports decision, investigation, requirement, work-item, and task records only;
- specs and legacy archive records do not appear in normal listing;
- optional status filtering uses the selected kind vocabulary;
- listing order uses canonical ID with `desc` as the default and `asc` as the only alternative;
- listing limit defaults to 20, accepts 1 through 100, and does not support unbounded results;
- listing returns compact `ref`, `title`, `status`, and `date` entries plus `has_more`;
- every addressable record remains discoverable through listing; missing `title`, `status`, or `date` values are returned as `null` with an operation warning trigger; parsed invalid values remain unchanged;
- valid zero-match queries return an empty result without a warning or error;
- `get_record` is removed from the public tool surface;
- `suggest_next_record` is removed from the public tool surface and defines no compatibility alias or active request and response contract;
- `get_records` accepts 1 through 20 exact canonical refs per request;
- exact current sequential refs and active `spec:` refs query the active index only;
- accepted legacy sequential IDs query the legacy index only when legacy roots are configured;
- exact retrieval does not infer prefixes, repair case or whitespace, interpret paths, or perform fuzzy matching;
- response records preserve first-occurrence request order;
- duplicate refs use ordered first-occurrence deduplication and produce warnings;
- unresolved, malformed, unsupported, and unavailable legacy refs produce per-ref warnings without failing retrievable records;
- retrieval responses contain successful records only;
- `include_body` applies to the entire request and defaults to `false`;
- successful records include normalized metadata and headings, with verbatim body only when requested;
- normal listing and retrieval responses do not expose physical paths;
- warning taxonomy and exact diagnostic fields remain delegated to the validation and response-boundary Work Item;
- resolver fallback order remains delegated and does not alter exact retrieval behavior;
- all changed specs pass scoped validation;
- independent review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records changed files, validation results, review verdict, and residual limitations.

## Evidence

- `DRMCP-INV-MCP-002`: Query, exact-retrieval, tool-surface, and response-boundary findings.
- `DRMCP-ADR-MCP-001`: Accepted active/legacy authority and path-exposure direction.
- `DRMCP-REQ-MCP-001`: Source Requirement.
- `DRMCP-WORK-MCP-003`: Upstream current discovery, active-index, and normalized-record contract owner.
- `DRMCP-TASK-MCP-001-04`: Hub lifecycle gate for this Work Item.
- 2026-06-26 accepted baseline: compact active listing, `get_records`-only exact retrieval, ordered partial success, and normal path hiding.
- `DRMCP-TASK-MCP-004-01` opened on 2026-06-27.
  - The authority baseline, candidate-claim classification, and affected-file manifest are recorded.
  - No normative DRMCP spec changed during the baseline inventory.
  - Invalid-but-addressable records remain in normal listing. Missing compact fields use `null` and trigger an operation warning; parsed invalid values remain unchanged.
  - Warning taxonomy, severity, shared fields, and source-location representation remain delegated to W006.
  - No unresolved T01-level design decision remains.
  - Independent baseline review verdict: `PASS`; findings and advisories: none; closure readiness: `ready`.
  - `DRMCP-TASK-MCP-004-01` closed as `done` on 2026-06-27.
- `DRMCP-TASK-MCP-004-02` opened on 2026-06-27.
  - Task scope is limited to the compact active-index `list_records` contract.
  - `tools/list-records.md` was rewritten with required `app_namespace`, `kind`, and `domain`; optional `status`; canonical `ref` ordering; bounded `limit`; compact nullable fields; `has_more`; and top-level warning placement.
  - Exact ID, range, broad, mixed-kind, spec, legacy, arbitrary-ordering, unbounded, full-metadata, physical-path, and path-tie-break behavior was removed.
  - W003 addressability and duplicate-conflict inputs were consumed without reopening discovery or index construction.
  - W005 resolver and legacy-fallback behavior remains excluded.
  - W006 warning taxonomy, severity, shared fields, source-location schema, validation semantics, and exceptional path exposure remain excluded.
  - `schema.fields` and `schema.record_model` required no pointer update because both already delegate public listing representation to W004.
  - Scoped strict validation result: `[strict]  All 1 file(s) OK.`
  - `git diff --check` result: PASS; no whitespace errors reported.
  - Git emitted an LF-to-CRLF working-copy warning for `tools/list-records.md`; the warning is non-blocking.
  - Independent review verdict: `PASS`; blocking, major, minor findings, and advisories: none.
  - Review confirmed validation evidence and LF-to-CRLF warning handling; closure readiness: `ready`.
  - `DRMCP-TASK-MCP-004-02` closed as `done` on 2026-06-27.
- `DRMCP-TASK-MCP-004-03` opened on 2026-06-27.
  - T03 changed-file scope is limited to the T03 Task, W004 workflow synchronization, `tools/get-record.md`, and `tools/get-records.md`.
  - `tools/get-record.md` is now an explicit retirement marker. It defines no public request, response, compatibility alias, or tool-execution error contract.
  - Single-record exact retrieval now uses `get_records` with one `refs` element.
  - `tools/get-records.md` now requires `refs`, accepts 1 through 20 strings, and supports request-wide boolean `include_body` with default `false`.
  - Accepted inputs point to PRODUCT authorities for current sequential refs, active path-derived `spec:` refs, and accepted Brewprint legacy sequential IDs.
  - Exact lookup performs no prefix inference, case or whitespace repair, path interpretation, fuzzy matching, partial completion, family guessing, or resolver invocation.
  - Duplicate detection uses exact string equality. First occurrence wins, successful records preserve first-occurrence request order, and duplicate occurrences trigger top-level warnings.
  - Malformed, unsupported, unresolved, unavailable legacy, and duplicate inputs produce warning triggers without discarding retrievable records.
  - The response now contains successful `records` only plus top-level `warnings`; stale `items`, `retrieval_status`, failure placeholders, `record: null`, and per-item diagnostic wrappers were removed.
  - Successful records contain canonical `ref`, parsed normalized `metadata`, `headings`, and conditional verbatim `body`.
  - Missing parsed metadata fields are omitted. Parsed invalid values remain unchanged. No missing value is fabricated.
  - Normal retrieval records exclude physical paths, source locations, provenance, internal index paths or state, and resolver traces.
  - `schema.fields`, `schema.record_model`, and `schema.record_source` were rechecked. Each already delegates public retrieval projection to W004, so no shared schema change was required.
  - W003 discovery, parsing, identity, normalized-field, addressability, conflict, and invalid-source ownership remains unchanged.
  - W005 resolver, current-first fallback, configured legacy orchestration, and fallback ordering remain excluded.
  - W006 warning schema, category names, severity, shared fields, source-location representation, validation behavior, and exceptional path exposure remain excluded.
  - T04 still owns tools overview, MVP scope, responsibility boundary, and broader pointer synchronization.
  - Scoped stale-claim confirmation completed for the two changed operation specs. No positive stale claim remains; the only removed-token match is a normative prohibition sentence.
  - Scoped strict validator result: `[strict]  All 2 file(s) OK.`
  - `git diff --check` result: PASS; no whitespace errors reported.
  - Git emitted LF-to-CRLF working-copy warnings for `tools/get-record.md` and `tools/get-records.md`; both warnings are non-blocking.
  - Independent review verdict: `PASS`; blocking, major, minor findings, and advisories: none.
  - Review confirmed the validation evidence, ownership boundaries, retirement behavior, request contract, batch semantics, partial-success behavior, and successful-record projection.
  - T03 closure readiness: `ready`.
  - `DRMCP-TASK-MCP-004-03` closed as `done` on 2026-06-27.
- `DRMCP-TASK-MCP-004-04` opened on 2026-06-27.
  - T04 changed-file scope covers the T04 Task, W004 workflow synchronization, the review-required W006 planning-boundary correction, root `overview.md`, `tools/overview.md`, `mvp-scope.md`, and `responsibility-boundary.md`.
  - Root `overview.md` received pointer-only Topics summary corrections for the synchronized MVP and tools child specs.
  - `tools/overview.md` now provides a navigation-first public catalog. It omits invokable `get_record` and `suggest_next_record` entries.
  - The overview identifies `list_records` as compact active-index listing and `get_records` as sole exact retrieval, including single retrieval through one `refs` element.
  - Stale path-bearing common response examples, `ids`, range errors, mixed success/failure wrappers, and duplicated warning authority were removed from the overview.
  - `mvp-scope.md` now classifies the current-format read contract without asserting implementation or release status.
  - The MVP scope removes `get_record`, `suggest_next_record`, YAML-front-matter source claims, broad/range listing, and W004 ownership over resolver or diagnostic behavior.
  - `responsibility-boundary.md` now separates W003 discovery/model, W004 query/retrieval plus normal path hiding, W005 resolver, and W006 diagnostic plus exceptional path exposure ownership.
  - `schema.fields`, `schema.record_model`, and `schema.record_source` were rechecked and already delegate public listing, exact retrieval, headings, body, warning, and path representation correctly. No shared schema file changed.
  - `tools/list-records.md`, `tools/get-record.md`, and `tools/get-records.md` were rechecked without reopening their accepted T02/T03 contracts.
  - W005 resolver behavior, W006 normative diagnostic and validation contracts, fixtures, implementation, automated tests, and authoring transactions remain excluded.
  - Independent review verdict: `NEEDS REVISION`; blocking, minor findings, and advisories: none; major findings: `F-MAJ-01` and `F-MAJ-02`.
  - `F-MAJ-01` identified overlapping normal path-hiding ownership in the W006 planning record. The W006 Boundary, Impact Scope, Task flow, Task Candidates, Completion Condition, Goal, and Evidence were narrowed so W004 is the sole normal list/exact-retrieval response authority and W006 owns only exceptional path exposure.
  - `F-MAJ-02` identified that the recorded diff check preceded Evidence updates.
  - Post-correction seven-file `git diff --check` result: PASS; no whitespace errors reported.
  - That run emitted non-blocking LF-to-CRLF warnings only for root `overview.md`, `tools/overview.md`, `mvp-scope.md`, and `responsibility-boundary.md`; the Task, W004, and W006 planning records emitted no warning.
  - Scoped strict validator result remains `[strict]  All 4 file(s) OK.`; no normative spec changed during review correction.
  - Final seven-file `git diff --check` result after Evidence synchronization: PASS; no whitespace errors or command failure reported.
  - The final run emitted non-blocking LF-to-CRLF warnings only for root `overview.md`, `tools/overview.md`, `mvp-scope.md`, and `responsibility-boundary.md`; the Task, W004, and W006 planning records emitted no warning.
  - Independent re-review verdict: `PASS`.
  - Previous findings `F-MAJ-01` and `F-MAJ-02`: `CLOSED`.
  - Remaining blocking, major, and minor findings: none.
  - Advisories: none affecting closure readiness.
  - T04 closure readiness: `ready`.
  - `DRMCP-TASK-MCP-004-04` closed as `done` on 2026-06-27.
  - Post-closure seven-file `git diff --check` result before final Evidence synchronization: PASS; no whitespace errors or command failure reported.
  - LF-to-CRLF warnings were limited to root `overview.md`, `tools/overview.md`, `mvp-scope.md`, and `responsibility-boundary.md`; the Task, W004, and W006 planning records emitted no warning.
  - This final Evidence synchronization updates the Task and W004 after that run, so one final seven-file `git diff --check` is required.
- `DRMCP-TASK-MCP-004-05` opened on 2026-06-27.
  - T05 and W004 remained `in_progress` through independent review and correction.
  - Final normative changed-spec set after first-review correction: `tools/list-records.md`, `tools/get-record.md`, `tools/get-records.md`, `tools/suggest-next-record.md`, root `overview.md`, `tools/overview.md`, `mvp-scope.md`, and `responsibility-boundary.md`.
  - `schema/fields.md`, `schema/record-model.md`, and `schema/record-source.md` were rechecked and remain unchanged.
  - First independent review verdict: `NEEDS REVISION`; blocking, minor findings, and advisories: none; major finding: `F-MAJ-01`.
  - `F-MAJ-01` found that `tools/suggest-next-record.md` still defined an active P1 operation despite the accepted removal from the current public tool surface.
  - The correction rewrote `tools/suggest-next-record.md` as an explicit retirement marker with no current request, response, compatibility alias, tool-execution error contract, ADR-number suggestion, or V01 path suggestion.
  - The corrected audit confirms compact active listing, `get_records`-only exact retrieval, one-element `refs` single retrieval, retirement of `get_record` and `suggest_next_record`, ordered partial success, successful-record-only responses, current-source H1-adjacent metadata, normal path hiding, and W003-W006 ownership separation.
  - No unresolved design decision remains.
  - Refreshed `git status --short` confirms the final W004 manifest: eight modified normative specs plus untracked T01 through T05, W004, and the T04-corrected W006 planning record.
  - Other dirty or untracked repository files are outside W004 and remain excluded.
  - Previous scoped strict validation result: `[strict]  All 7 file(s) OK.`
  - Final post-correction scoped strict validation result: `[strict]  All 8 file(s) OK.`
  - Final eight-spec `git diff --check` result: no command failure or whitespace error reported.
  - Git emitted non-blocking LF-to-CRLF warnings for all eight normative specs.
  - Refreshed no-index checks covered untracked T01 through T05, W004, and W006.
  - Every no-index check returned the expected difference exit code `1`; no exit code `2` or higher and no whitespace error was reported.
  - Git emitted non-blocking LF-to-CRLF warnings for all seven untracked records.
  - A post-correction no-index recheck covered T05 and W004 before the final Evidence synchronization.
  - Both returned the expected difference exit code `1`, no whitespace error, and LF-to-CRLF warnings only.
  - No command returned exit code `2` or higher. The warnings are non-blocking.
  - The final pre-review post-synchronization no-index check returned exit code `1` for both T05 and W004, with no whitespace error and LF-to-CRLF warnings only. No file changed after that check before re-review.
  - Independent re-review verdict: `PASS`.
  - Previous finding `F-MAJ-01`: `CLOSED`.
  - Remaining blocking, major, and minor findings: none.
  - Advisories: none.
  - Validation evidence assessment: valid.
  - Changed-file integrity evidence assessment: valid.
  - Ownership assessment: valid and non-overlapping.
  - T05 closure readiness: `ready`.
  - W004 closure readiness: `ready`.
  - T05 status changed to `done` on 2026-06-27.
  - W004 status changed to `done` on 2026-06-27.
  - No unresolved W004 decision or residual W004-scope limitation remains.
  - W005 and W006 execution remain outside this closure; W006 remains `not_started` with no child Tasks.
  - A final post-closure no-index check of T05 and W004 must be executed after this synchronization and supplied directly as external closure evidence without another file edit.
