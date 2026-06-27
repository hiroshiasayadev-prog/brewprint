# DRMCP-WORK-MCP-005: Resolver and configured legacy-fallback contract realignment

- **id**: DRMCP-WORK-MCP-005
- **status**: done
- **date**: 2026-06-27
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-INV-MCP-002
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-004
  - DRMCP-TASK-MCP-001-05
  - PRODUCT-WORK-SPEC-014
  - spec:product.brewprint.compatibility.legacy_id_compatibility
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.resolver
  - spec:drmcp.design_records_mcp.schema.discovery
  - spec:drmcp.design_records_mcp.schema.record_model
  - spec:drmcp.design_records_mcp.schema.record_source
  - spec:drmcp.design_records_mcp.schema.fields
  - spec:drmcp.design_records_mcp.schema.metadata_grammar
  - spec:drmcp.design_records_mcp.schema.overview
  - spec:drmcp.design_records_mcp.schema.id_normalization
  - spec:drmcp.design_records_mcp.tools.resolve_reference
  - spec:drmcp.design_records_mcp.tools.get_records
  - spec:drmcp.design_records_mcp.tools.overview
  - spec:drmcp.design_records_mcp.mvp_scope
  - spec:drmcp.design_records_mcp.responsibility_boundary
- **tasks**:
  - DRMCP-TASK-MCP-005-01
  - DRMCP-TASK-MCP-005-02
  - DRMCP-TASK-MCP-005-03
  - DRMCP-TASK-MCP-005-04
  - DRMCP-TASK-MCP-005-05

## Goal

Establish the corrected current-first reference resolver and configured legacy-fallback contract.

Keep current canonical resolution authoritative while allowing exact read-only fallback for approved legacy sequential IDs.

## Boundary

This Work Item owns:

- current canonical grammar evaluation before any legacy grammar check;
- fallback eligibility only after current resolution returns unresolved;
- exact accepted V01 sequential-family recognition;
- explicit `legacy_roots` configuration and disabled fallback without that configuration;
- exact lookup against a separate minimal legacy issued-ID-to-source map;
- filename-derived legacy identity, source readability, duplicate-ID behavior, and source-first exact retrieval without current-model normalization;
- preservation of the issued legacy ID after successful resolution and retrieval;
- rejection of `V01-SPEC-*`, app-prefixless IDs, physical paths, fuzzy repair, legacy YAML semantic-ref aliases, direct `yaml:` inputs, and `fixture:` inputs;
- synchronization of resolver, configuration, and related overview contracts;
- scoped validation and independent review for this contract boundary.

This Work Item does not own:

- current-root discovery, current spec parsing, or active-index construction;
- compact listing behavior or exact-retrieval request, ordering, partial-success, and current-record projection behavior;
- diagnostic taxonomy or diagnostic/warning response-field design;
- repository validation execution or general path-exposure policy;
- fixture authoring;
- resolver, configuration, or legacy-index implementation;
- accepted legacy-family semantics owned by Brewprint compatibility contracts;
- authoring transaction behavior.

`DRMCP-WORK-MCP-003` supplies separate current and legacy index contracts.
`DRMCP-WORK-MCP-004` supplies exact retrieval and resolver boundaries.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Source Requirement for current-first resolution and configured legacy fallback. |
| `DRMCP-ADR-MCP-001` | Governs resolution order, accepted legacy families, and archive isolation. |
| `DRMCP-INV-MCP-002` | Supplies resolver, configuration, and compatibility findings. |
| `PRODUCT-WORK-SPEC-014` | Supplies the accepted removal of `V01-SPEC-*` compatibility authority. |
| `spec:product.brewprint.compatibility.legacy_id_compatibility` | Supplies approved Brewprint legacy sequential families. |
| `DRMCP-WORK-MCP-003` | Supplies current and legacy index separation. |
| `DRMCP-WORK-MCP-004` | Supplies exact retrieval behavior and the retrieval-versus-resolution boundary. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Add explicit `legacy_roots`, legacy issued-ID parser mapping, source-enumeration boundary, and separate lookup-map construction without reopening current-root rules. |
| `spec:drmcp.design_records_mcp.resolver` | Correct current-first and legacy-fallback orchestration. |
| `spec:drmcp.design_records_mcp.schema.discovery` | Recheck current-only discovery, no-legacy-alias behavior, and resolver delegation. No W005 correction is currently planned. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Recheck current-only scope. No legacy normalization correction is planned. |
| `spec:drmcp.design_records_mcp.schema.record_source` | Recheck current-only scope. No legacy source-model correction is planned. |
| `spec:drmcp.design_records_mcp.schema.fields` | Recheck current normalized vocabulary. Minimal legacy retrieval does not extend it. |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Keep current-only with no legacy exception. |
| `spec:drmcp.design_records_mcp.schema.overview` | Retain unchanged because T03 creates no dedicated legacy schema. |
| `spec:drmcp.design_records_mcp.schema.id_normalization` | Recheck current canonical identity and the separate legacy compatibility boundary. No W005 correction is currently planned. |
| `spec:drmcp.design_records_mcp.tools.resolve_reference` | Correct the public resolver operation contract while delegating diagnostic and path representation to W006. |
| `spec:drmcp.design_records_mcp.tools.get_records` | Synchronize only the minimal legacy source projection; preserve W004 request, ordering, partial-success, and common wrapper contracts. |
| Tool catalog, MVP scope, and responsibility boundary | Recheck W005 pointers in T04. Apply pointer-only synchronization only when required by T02 or T03. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Baseline confirmation | `DRMCP-WORK-MCP-003`, `DRMCP-WORK-MCP-004`, `PRODUCT-WORK-SPEC-014` | Confirm affected specs, accepted inputs, and excluded behavior. |
| B. Current-first resolution contract | Phase A | Define current grammar handling, active-index lookup, unresolved-only fallback eligibility, and current successful non-path target projection. |
| C. Configured legacy fallback contract | Phases A-B | Resolve legacy-root validation, filename-derived identity, minimal issued-ID-to-source lookup, duplicate handling, source-first exact retrieval, final status mapping, and minimal successful-target projection. |
| D. Rejection and cross-spec synchronization | Phases B-C | Remove fuzzy, path, bare-ID, YAML-alias, and `V01-SPEC-*` behavior. |
| E. Validation and review | Phases B-D | Validate changed contracts, run independent review, apply corrections, and close. |

Fixtures and implementation proceed through separate Work Items.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Confirm the affected-file manifest, W003/W004 inputs, compatibility authority, and downstream exclusions. | W003 and W004 contract boundaries available. |
| T02 | Correct current-first grammar evaluation and active-index resolution contracts. | T01. |
| T03 | Resolve `legacy_roots`, filename-derived issued identity, minimal exact lookup-map, duplicate-ID, root-failure, source-first exact retrieval, final outcome, and minimal legacy-target details; then reflect only the required contracts. | T01-T02. |
| T04 | Synchronize rejection behavior and related overview, resolver, and configuration pointers. | T02-T03. |
| T05 | Run scoped validation, independent review, required corrections, and closure. | T04. |

## Completion Condition

This Work Item is complete when all of the following are true:

- current canonical grammar and active-index lookup run before legacy grammar checks;
- legacy fallback is disabled without configured `legacy_roots`;
- only approved V01 decision, investigation, requirement, work-item, and task families are accepted;
- `V01-SPEC-*`, app-prefixless IDs, physical paths, fuzzy repair, legacy YAML aliases, direct `yaml:` inputs, and `fixture:` inputs are rejected;
- fallback queries only the separate legacy issued-ID-to-source lookup map;
- the minimal filename-keyed legacy lookup contract is explicit, including source readability and duplicate-ID behavior;
- legacy exact retrieval preserves readable source access without requiring current-model normalization;
- successful legacy resolution preserves the issued legacy ID;
- exact retrieval and resolver behavior remain separate;
- diagnostic taxonomy, fixtures, and implementation remain delegated;
- all changed specs pass scoped validation;
- independent review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records changed files, validation results, review verdict, and residual limitations.

## Evidence

- `DRMCP-ADR-MCP-001`: Accepted current-first and configuration-gated legacy fallback direction.
- `DRMCP-REQ-MCP-001`: Source Requirement.
- `DRMCP-WORK-MCP-003`: Current and legacy index-separation contract owner. The Work Item is `done`.
- `DRMCP-WORK-MCP-004`: Query and exact-retrieval contract owner. The Work Item is `done`.
- `PRODUCT-WORK-SPEC-014`: Brewprint compatibility correction prerequisite. The Work Item is `done`.
- `DRMCP-TASK-MCP-001-05`: Hub lifecycle gate for this Work Item. The Task moved to `in_progress` on 2026-06-27.
- `DRMCP-TASK-MCP-005-01` completed on 2026-06-27.
  - T01 records the authority baseline, current-first order, fallback gate, accepted and rejected input families, claim classification, affected-file candidate manifest, and ownership exclusions.
  - T01 identifies `resolver.md`, `tools/resolve-reference.md`, `namespace-scanning.md`, `schema/record-model.md`, `schema/record-source.md`, and `schema/fields.md` as W005-owned normative candidates.
  - T01 identifies `schema/metadata-grammar.md` as a conditional T03 candidate and allows a dedicated legacy archive schema when separation is cleaner than extending current-only schema documents.
  - T01 identifies `schema/overview.md` as conditional T03 parent synchronization when a dedicated legacy archive schema is created.
  - T01 identifies `schema/discovery.md`, `schema/id-normalization.md`, `tools/get-records.md`, `tools/overview.md`, `mvp-scope.md`, and `responsibility-boundary.md` as recheck candidates.
  - T01 routes diagnostic, validation, source-location, and exceptional path representation to W006.
  - T01 excludes authoring transaction specs, fixtures, implementation, automated tests, and legacy-file migration.
  - T01 changed no normative DRMCP spec.
  - No unresolved T01-level authority, scope, operation-split, or ownership decision remains.
  - Current/legacy grammar overlap follows the already-accepted current-first then unresolved-only legacy fallback order and is not reopened in T03.
  - T03-specific `legacy_roots`, legacy-source, archive-record/index, duplicate-ID, root-failure, and legacy projection questions are explicit and intentionally remain for T03 resolution.
  - `yaml:` and `fixture:` are classified as unsupported direct resolver inputs.
  - Successful non-path target projection is W005-owned; diagnostics and path representation remain W006-owned.
  - Initial independent review returned `NEEDS REVISION` with F-MAJ-01, F-MAJ-02, and F-MIN-01; all were closed by the first correction pass.
  - First independent re-review returned `NEEDS REVISION` with F-MAJ-03 and F-MIN-02.
  - F-MAJ-03 was corrected by adding conditional `schema/overview.md` parent synchronization to T03 scope.
  - F-MIN-02 was corrected by updating this Work Item date to `2026-06-27` after the substantive scope expansion.
  - Independent final re-review returned `PASS` with no blocking, major, minor, or advisory findings.
  - All findings F-MAJ-01, F-MAJ-02, F-MIN-01, F-MAJ-03, and F-MIN-02 are closed.
  - T02 start readiness is `READY`; T03 design-decision readiness is `READY`.
- `DRMCP-TASK-MCP-005-02` completed on 2026-06-27.
  - T02 changed `spec:drmcp.design_records_mcp.resolver` and `spec:drmcp.design_records_mcp.tools.resolve_reference`.
  - Both changed normative specs use date `2026-06-27`.
  - Current sequential input now points to complete app-aware IDs from `spec:product.design_records.namespace_model.artifact_id_grammar`.
  - Current spec input now points to path-derived document-level refs from `spec:product.design_records.spec_format.spec_id_as_ref`.
  - Front-matter `semantic_refs`, front-matter `sections`, section-level spec refs, and `namespace_prefix = V01-` current-grammar assumptions were removed.
  - Current grammar evaluation and active-index lookup run before any legacy fallback eligibility.
  - One resolved current target stops legacy grammar evaluation and legacy-index lookup.
  - The current stage must remain unresolved before legacy eligibility begins, including grammar overlap cases.
  - A current input is never rewritten into a legacy input.
  - `resolve_reference` and `get_records` remain separate and do not call each other.
  - Current successful targets use `target_type`, canonical `ref`, `kind`, `title`, and `status`.
  - `target_type` distinguishes `current_spec` from `current_sequential_record`.
  - Normal successful targets contain no physical path, source location, provenance, section anchor, index path, or resolver trace.
  - W006 retains diagnostic objects, warning schema, category, severity, message, source-location, validation, and exceptional path representation.
  - T03 retains `legacy_roots` entry fields, legacy-root failures, legacy parsing, archive-record/index construction, duplicate issued IDs, normalized legacy fields, configured legacy lookup, issued-ID preservation, legacy target projection, schema-placement decisions, and final status mapping for accepted legacy inputs.
  - First scoped validator execution on 2026-06-27 returned `FAIL` with `MISSING_SECTION` because the interface spec used `## Response outcomes` instead of required `## Response`.
  - The heading was corrected to canonical `## Response` without changing response semantics.
  - Post-correction scoped validator rerun returned `PASS`: `[strict]  All 2 file(s) OK.`
  - Initial independent review verdict: `NEEDS REVISION`; major finding `F-MAJ-01` was corrected on 2026-06-27.
  - The correction limits `unsupported` to values accepted by neither current nor accepted legacy-family grammar.
  - Accepted legacy-family grammar no longer depends on `legacy_roots` configuration or usability.
  - Disabled, unavailable, unresolved, conflicted, and resolved legacy outcomes remain T03-owned.
  - Independent re-review verdict: `PASS`.
  - F-MAJ-01: `CLOSED`.
  - Re-review reported no blocking, major, minor, or advisory findings.
  - The independent reviewer did not rerun the validator because repository-local command execution was unavailable, but confirmed the recorded failure, correction, PASS result, and canonical `## Response` heading by static inspection.
  - T02 status changed to `done` on 2026-06-27.
  - T03 start readiness: `READY`.
- `DRMCP-TASK-MCP-005-03` opened on 2026-06-27.
  - T03 owns configured legacy-root semantics, filename-derived legacy identity, minimal exact lookup-map construction, duplicate issued-ID behavior, source-first exact retrieval, final accepted-legacy outcome mapping, and minimal successful legacy target projection.
  - The initial normalized archive-record/schema direction was withdrawn before normative reflection as unnecessary compatibility scope.
  - T03 status changed to `done` on 2026-06-27 after independent re-review PASS.
  - T03 changed `namespace-scanning.md`, `resolver.md`, `tools/resolve-reference.md`, and `tools/get-records.md`.
  - Initial scoped strict validation passed on 2026-06-27: `[strict]  All 4 file(s) OK.`
  - Initial independent review returned `NEEDS REVISION` with `F-MAJ-01`, `F-MAJ-02`, and `F-MIN-01`.
  - `F-MAJ-01` was corrected by defining legacy-first exact classification for overlapping `get_records` inputs with one lookup scope and no resolver invocation.
  - `F-MAJ-02` was corrected by defining family-specific issued-ID lexical grammar, fixed sequence widths, domain grammar, slug boundary, and shared resolver/retrieval recognition mapping.
  - `F-MIN-01` was corrected by prohibiting filesystem-alias candidates and traversal and by requiring canonical candidate containment.
  - The initial validator PASS predates these corrections.
  - Post-correction scoped strict validation passed on 2026-06-27: `[strict]  All 4 file(s) OK.`
  - Independent re-review returned `PASS`.
  - `F-MAJ-01`, `F-MAJ-02`, and `F-MIN-01` are all `CLOSED`.
  - Re-review reported no blocking, major, minor, advisory, regression, or ownership-boundary findings.
  - T03 closure readiness was `ready`.
- T03 contract correction, validation, finding closure, and independent-review evidence are complete.
- `DRMCP-TASK-MCP-005-04` opened on 2026-06-27.
  - T04 records the rejected-input matrix and preserves the resolver/exact-retrieval operation split.
  - Normative reflection changed `namespace-scanning.md`, `resolver.md`, `tools/resolve-reference.md`, and `tools/get-records.md`.
  - The changes synchronize exact lexical-mapping pointers, resolver `unsupported` behavior, and `get_records` per-item warning triggers.
  - Root overview, tool overview, MVP scope, responsibility boundary, listing, and current schema candidates were rechecked and require no change.
  - Pre-review scoped strict validation passed on 2026-06-28: `[strict]  All 4 file(s) OK.`
  - Initial independent review returned `NEEDS REVISION` with `F-MAJ-01`.
  - `F-MAJ-01` identified that semantic-origin classification of section-level `spec:` strings conflicted with lexical current-grammar classification and unresolved lookup behavior.
  - The correction classifies every current-grammar-matching `spec:` string as current, queries only the active index, and keeps an exact lookup miss unresolved.
  - Front-matter alias registries, section targets, and heading lookup remain prohibited.
  - The pre-review validator PASS predates this correction.
  - Post-correction scoped strict validation passed on 2026-06-28: `[strict]  All 4 file(s) OK.`
  - Independent re-review returned `PASS` with no blocking, major, or minor findings.
  - `F-MAJ-01` is `CLOSED`.
  - Advisories A-01 and A-02 remain non-blocking and require no T04 correction.
  - T04 closure readiness is `READY`.
  - T04 status changed to `done` on 2026-06-28.
- `DRMCP-TASK-MCP-005-05` opened on 2026-06-28.
  - T05 owns final scoped validation, independent final review, finding closure, and W005 closure synchronization.
  - Final normative changed-file scope is `namespace-scanning.md`, `resolver.md`, `tools/resolve-reference.md`, and `tools/get-records.md`.
  - Root overview, tool overview, MVP scope, responsibility boundary, `tools/list-records.md`, current schema specs, diagnostics and validation specs, and authoring transaction specs were rechecked as unchanged candidates.
  - T05 does not add A-01 or A-02 to scope unless final review identifies a new W005 contract contradiction.
  - Final resolver baseline remains current grammar and active-index lookup first, resolved-current stop, unresolved-only accepted legacy evaluation, configured legacy lookup only, `unresolved` for accepted lookup failure, and `unsupported` only when neither grammar accepts the input.
  - Final `get_records` baseline remains accepted-legacy exact classification first, one lookup scope, no resolver invocation, no second lookup, exact ordered deduplication, first-occurrence ordering, partial success, successful-record-only response, top-level warnings, and request-wide `include_body`.
  - Final legacy baseline remains missing-or-empty fallback disablement, mandatory configured roots, no partial root acceptance, filename-derived issued identity, exact case-sensitive lookup, no duplicate winner, no filesystem-alias traversal, no current-model normalization, no dedicated legacy schema, and issued-ID preservation.
  - Current spec classification remains lexical. Every current-grammar-matching `spec:` value queries only the active index; an exact miss remains unresolved and does not trigger alias, section, or heading lookup.
  - W003 current discovery and index construction, W004 exact-retrieval wrapper and projection, W006 diagnostics and validation, and PRODUCT semantic authority remain outside T05 correction ownership.
  - Final four-file scoped validator command was executed from the repository root and reported on 2026-06-28.
  - Recorded scoped validator result before final-review correction: `[strict]  All 4 file(s) OK.`
  - The recorded PASS covered exactly `namespace-scanning.md`, `resolver.md`, `tools/resolve-reference.md`, and `tools/get-records.md`.
  - Independent final review returned `NEEDS REVISION` with minor finding `F-MIN-FINAL-01`.
  - `F-MIN-FINAL-01` found stale `date: 2026-06-27` metadata in `resolver.md`, `tools/resolve-reference.md`, and `tools/get-records.md` after the substantive 2026-06-28 lexical current-spec classification correction.
  - The three affected spec dates were updated to `2026-06-28` without changing contract body text or expanding the final four-file manifest.
  - The previous validator PASS predates this metadata correction and is superseded for closure purposes.
  - The same four-file scoped validator was rerun after the metadata correction on 2026-06-28 and returned `[strict]  All 4 file(s) OK.`
  - Repository-local `git diff --check` reported no whitespace error; only non-blocking line-ending conversion warnings were emitted.
  - Limited independent re-review returned `PASS`.
  - `F-MIN-FINAL-01` is `CLOSED`.
  - The reviewer confirmed that the three date changes were metadata-only, the contract body remained unchanged, and the final four-file normative manifest remained intact.
  - The reviewer accepted the post-correction validator result and the non-blocking LF-to-CRLF warnings from `git diff --check`.
  - Blocking findings: none.
  - Major findings: none.
  - Minor findings: none.
  - Advisories A-01 and A-02 remain non-blocking and unchanged.
  - T05 closure readiness: `READY`.
  - W005 closure readiness: `READY`.
  - `DRMCP-TASK-MCP-001-05` closure readiness: `READY`.
  - T05 status changed to `done` on 2026-06-28.
  - W005 status changed to `done` on 2026-06-28 after accepted limited re-review PASS.
