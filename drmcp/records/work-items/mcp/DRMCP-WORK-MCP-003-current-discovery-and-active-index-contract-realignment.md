# DRMCP-WORK-MCP-003: Current discovery and active-index contract realignment

- **id**: DRMCP-WORK-MCP-003
- **status**: done
- **date**: 2026-06-27
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-INV-MCP-002
  - DRMCP-WORK-MCP-001
  - DRMCP-TASK-MCP-001-03
  - PRODUCT-WORK-SPEC-014
  - spec:product.design_records.namespace_model
  - spec:product.design_records.repository_layout
  - spec:product.design_records.repository_layout.record_discovery_paths
  - spec:product.design_records.spec_format
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.responsibility_boundary
  - spec:drmcp.design_records_mcp.schema.overview
  - spec:drmcp.design_records_mcp.schema.discovery
  - spec:drmcp.design_records_mcp.schema.record_source
  - spec:drmcp.design_records_mcp.schema.record_model
  - spec:drmcp.design_records_mcp.schema.fields
  - spec:drmcp.design_records_mcp.schema.id_normalization
  - spec:drmcp.design_records_mcp.schema.metadata_grammar
- **tasks**:
  - DRMCP-TASK-MCP-003-01
  - DRMCP-TASK-MCP-003-02
  - DRMCP-TASK-MCP-003-03
  - DRMCP-TASK-MCP-003-04
  - DRMCP-TASK-MCP-003-05

## Goal

Establish the corrected DRMCP contract for configured current-root discovery, current record parsing, and deterministic active-index construction.

Align current spec discovery with H1-adjacent metadata and path-derived canonical `spec:` refs.
Keep current and optional legacy indexes contractually separate before query, resolver, fixture, or implementation work proceeds.

## Boundary

This Work Item owns:

- configured current-root discovery semantics;
- app namespace and current records-root association;
- supported current record-kind inclusion behavior;
- current spec parsing through H1-adjacent metadata;
- current spec identity derivation from repository placement;
- rejection of YAML front matter as an active spec metadata source;
- deterministic duplicate canonical-identity handling across current roots;
- separation of active-index and optional legacy-index contracts;
- correction of shared record-source, record-model, field, and ID-normalization contracts required by discovery and indexing;
- contract validation and independent review for this boundary.

This Work Item does not own:

- PRODUCT namespace, repository-layout, or spec-format semantics;
- Brewprint accepted legacy-family policy;
- query filters, ordering, ranges, or pagination;
- exact retrieval behavior;
- current-first resolver or legacy fallback lookup;
- validation execution or diagnostic taxonomy outside discovery/index requirements;
- physical-path exposure in tool responses outside the internal index model;
- fixture implementation;
- parser, index, tool, or runtime implementation;
- automated implementation tests;
- authoring transaction behavior;
- migration or rewriting of legacy YAML specs.

`PRODUCT-WORK-SPEC-014` owns removal of `V01-SPEC-*` compatibility authority.
This Work Item consumes that accepted boundary and does not duplicate it.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Source Requirement for current-root discovery and active-index behavior. |
| `DRMCP-ADR-MCP-001` | Governs current-first authority, separate legacy indexing, and implementation sequence. |
| `DRMCP-INV-MCP-002` | Supplies findings F01, F06, F11, F12, and F13 plus affected-file inventory. |
| `PRODUCT-WORK-SPEC-014` | Supplies the accepted removal of `V01-SPEC-*` compatibility authority. |
| `spec:product.design_records.namespace_model` | Canonical app and domain namespace authority consumed by DRMCP. |
| `spec:product.design_records.repository_layout` | Canonical record placement authority consumed by DRMCP. |
| `spec:product.design_records.spec_format` | Canonical current spec document and identity authority consumed by DRMCP. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Replace stale automatic-root and `V01-`-centric scanning assumptions. |
| `spec:drmcp.design_records_mcp.schema.discovery` | Replace YAML-front-matter spec inclusion rules with current-format discovery. |
| `spec:drmcp.design_records_mcp.schema.record_source` | Correct current metadata-source ownership and parsing inputs. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Correct current identity and index representation without defining tool response policy. |
| `spec:drmcp.design_records_mcp.schema.fields` | Correct current common-field and spec-field source descriptions. |
| `spec:drmcp.design_records_mcp.schema.id_normalization` | Replace prefix-stripping and legacy spec-ID assumptions with current app-aware sequential IDs and path-derived spec refs. |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Keep parser grammar DRMCP-owned while pointing semantic requiredness to PRODUCT authorities. |
| DRMCP overview and responsibility boundary | Synchronize active-index scope and ownership pointers. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Baseline and affected-file confirmation | `DRMCP-INV-MCP-002`, `DRMCP-REQ-MCP-001` | Confirm exact files, retained statements, rewritten statements, and downstream exclusions. |
| B. Current-root and index-separation contract | Phase A; accepted compatibility boundary from `PRODUCT-WORK-SPEC-014` | Define configured current roots, app association, active-index scope, and current/legacy separation. |
| C. Current spec parsing and identity contract | Phases A-B | Replace YAML-based active spec discovery with H1-adjacent metadata and path-derived `spec:` identity. |
| D. Record model and duplicate boundary | Phases B-C | Align record source, fields, model, duplicate identity behavior, and source-format rejection. |
| E. Cross-spec synchronization and review | Phases B-D | Synchronize overview pointers, validate changed specs, run independent review, and close findings. |

Query, retrieval, resolver, validation, fixtures, and implementation remain blocked from treating old discovery assumptions as authority.
Those workstreams proceed through their separately owned Work Items.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Confirm the affected-file manifest, finding coverage, retained behavior, and explicit downstream exclusions. | None. |
| T02 | Correct configured current-root discovery, app association, active-index scope, and current/legacy index separation. | T01; accepted `PRODUCT-WORK-SPEC-014` boundary. |
| T03 | Correct current spec metadata parsing and path-derived canonical identity contracts. | T01-T02. |
| T04 | Correct shared record-source, record-model, field, ID-normalization, metadata-grammar, duplicate identity, and invalid-source boundaries. | T02-T03. |
| T05 | Synchronize overview and ownership pointers, validate the changed contract set, run independent review, apply required corrections, and close the Work Item. | T04. |

Each Task must remain within contract authoring and review.
Implementation and fixtures belong to later Work Items under `DRMCP-REQ-MCP-001`.

## Completion Condition

This Work Item is complete when all of the following are true:

- configured current roots are the only source of the active index;
- current-root and app-namespace association is explicit and does not depend on `V01-` examples;
- active discovery covers the corrected supported current record kinds;
- current specs are discovered through H1-adjacent metadata and repository placement;
- current spec canonical identity is path-derived according to PRODUCT authority;
- YAML front matter is not an active current spec metadata source;
- specs without valid current identity are not silently adopted through legacy aliases;
- current and optional legacy indexes are defined as separate structures and operational scopes;
- duplicate canonical identity across current roots deterministically isolates the conflicted identity, preserves all conflicting sources for validation, keeps unrelated records readable, and does not select by filesystem order;
- shared record-source, record-model, fields, ID-normalization, and metadata-grammar contracts point to PRODUCT semantic authorities without duplicating them;
- query, exact retrieval, resolver, validation, fixtures, and implementation remain delegated to their separately owned Work Items;
- all changed specs pass scoped validation;
- independent review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records changed files, validation results, review verdict, and residual limitations.

## Evidence

- `DRMCP-INV-MCP-002`: Findings and affected-file disposition baseline.
- `DRMCP-ADR-MCP-001`: Accepted current-format-first and separate-index direction.
- `DRMCP-REQ-MCP-001`: Source Requirement.
- `PRODUCT-WORK-SPEC-014`: PRODUCT-owned compatibility prerequisite.
- `DRMCP-TASK-MCP-001-03`: Hub lifecycle gate for this Work Item.
- `DRMCP-TASK-MCP-003-01`: Affected-file manifest, finding coverage, retained behavior, rewritten behavior, and downstream exclusions completed on 2026-06-27.
- T01 added `spec:drmcp.design_records_mcp.schema.id_normalization` to the Work Item impact boundary.
- Compatibility gate resolved on 2026-06-27:
  - `PRODUCT-WORK-SPEC-014` is `done`;
  - active PRODUCT compatibility no longer accepts `V01-SPEC-*`;
  - compatibility-only legacy spec identity authority is removed;
  - `DRMCP-TASK-MCP-001-02` accepted the child completion evidence and is `done`.
- Remaining stale DRMCP claims are owned here for `schema.fields` and `schema.metadata_grammar`; `tools.validate_records` remains delegated to `DRMCP-WORK-MCP-006`.
- `DRMCP-TASK-MCP-003-02` completed on 2026-06-27:
  - current roots are explicit `app_namespace` and repository-relative `records_root` pairs;
  - one app maps to one `<app_namespace>/records` root;
  - configured roots are mandatory and invalid roots fail the whole active-index build;
  - valid empty current roots are allowed;
  - auto-discovery and folder-local ignore control are rejected;
  - current and optional legacy roots and indexes remain separate and non-overlapping;
  - duplicate canonical identity is an isolated index-entry conflict rather than a configured-root or repository-wide read failure;
  - conflicted identities are unavailable for normal ID-based retrieval, while conflicting sources remain validation inputs and unrelated records remain readable.
- `DRMCP-TASK-MCP-003-03` contract reflection completed on 2026-06-27:
  - Five normative specs updated: `schema/discovery.md`, `schema/record-source.md`, `schema/metadata-grammar.md`, `schema/fields.md`, `schema/id-normalization.md`.
  - `schema/discovery.md` fully rewritten to replace YAML `design_record` spec inclusion with current spec candidate paths, invalid source behavior, duplicate canonical identity conflict, and no-legacy-alias fallback.
  - `schema/record-source.md` corrected: YAML front matter removed as spec metadata source; H1-adjacent metadata block added as the current spec metadata source; current spec source mapping table added.
  - `schema/metadata-grammar.md` extended: current spec H1-adjacent grammar section added covering block start/continuation/termination, marker syntax, recognized keys, duplicate/unknown/missing key rules, scalar normalization, and field order independence.
  - `schema/fields.md` narrowly corrected: spec `id` source, spec title extraction, spec `status` source, and spec H1 format updated; `V01-SPEC-*` example replaced; `design_record.status` mismatch rule removed.
  - `schema/id-normalization.md` updated: current spec path-derived canonical ref derivation section added; `SPEC-<slug>` bare ID removed; no-repair and no-fallback rules stated.
  - Scoped validator result: `[strict] All 30 file(s) OK.` (2026-06-27).
  - Stale-contract search completed: no unresolved contradictions in T03-owned specs.
  - At this reflection checkpoint, T03 was `in_progress` and ready for review. T03 later closed as recorded below.
- Residual delegated scope recorded at the T03 reflection checkpoint; later T04 and T05 progress is recorded below:
  - T04 then owned stale spec relation-field mappings, shared sequential identity text, `schema/record-model.md`, shared field integration, and common invalid-source behavior.
  - W004 retains list and exact-retrieval response representation for invalid or conflicted sources.
  - W006 retains exact diagnostic identifiers, severity, and source-location response representation.
  - T05 owns overview and responsibility synchronization, final validation, independent review, and W003 closure.
- Independent review of T03 returned NEEDS REVISION. Review corrections applied on 2026-06-27:
  - `schema/id-normalization.md`: root index empty-segment case handled; derivation steps reordered for clarity; root index example added.
  - `schema/fields.md`: `id` section split into current spec (path-derived, no namespace_prefix) and sequential artifact (namespace_prefix-prefixed) paragraphs; title extraction intro scoped to sequential artifacts; spec status vocabulary removed from DRMCP allowed-values table with PRODUCT authority pointer added; Spec H1 section updated with accepted-kind validation and normative read-availability language.
  - `schema/discovery.md`: `may remain available` replaced with normative `remains addressable through that ref`; PRODUCT accepted spec kind set validation added to valid-source conditions; unknown/deferred kind candidates remain addressable when uniquely identifiable.
  - `schema/record-source.md`: `optional contract_class` replaced with `kind-conditional contract_class` including required/prohibited boundary.
  - Validation after corrections: `[strict] All 30 file(s) OK.`
- Second re-review returned NEEDS REVISION. Corrections applied on 2026-06-27: spec status note in `schema/fields.md` corrected to state that PRODUCT does not currently define a complete spec status vocabulary and DRMCP does not apply a spec-specific `invalid_status_for_kind` check; `may remain readable` in `schema/metadata-grammar.md` made normative. Validation: `[strict] All 30 file(s) OK.`
- Final review returned NEEDS REVISION. Corrections applied on 2026-06-27: `schema/metadata-grammar.md` `status` recognized-key row and single-authority paragraph corrected to per-field breakdown reflecting that complete spec status vocabulary is not currently defined by PRODUCT authority; `schema/fields.md` `invalid_status_for_kind` statement scoped to exclude current spec records until PRODUCT defines a complete vocabulary; T03 accepted decisions 15 and 16 narrowed to align. Validation: `[strict] All 30 file(s) OK.` No unresolved contradictions.
- `DRMCP-TASK-MCP-003-03` closure review completed on 2026-06-27:
  - verdict: PASS;
  - no blocking, major, or minor findings remain in the T03-owned contract set;
  - final scoped validation: `[strict] All 30 file(s) OK.`;
  - T03 status changed to `done`;
  - complete current spec status vocabulary remains an explicit PRODUCT authority gap, so DRMCP parses a required non-empty scalar without applying a spec-specific enum check;
  - T04 retains shared record integration and stale `design_record` relation-field cleanup;
  - W004 retains invalid/conflicted read response representation;
  - W006 retains exact diagnostic taxonomy.
- `DRMCP-TASK-MCP-003-04` opened on 2026-06-27:
  - owns shared internal source, parsed-field, canonical-identity, index-entry, validation-input, and conflict-group integration;
  - owns common-field availability boundaries, stale `design_record.*` cleanup, and current sequential artifact identity-model correction;
  - keeps public list/get representation in W004 and exact diagnostics and severity in W006;
  - initial normative outputs are `schema.record_model`, `schema.fields`, `schema.id_normalization`, and `schema.metadata_grammar`;
  - design decisions will be confirmed one at a time before normative spec edits.
- `DRMCP-TASK-MCP-003-04` closed on 2026-06-27:
  - final changed-spec set: `schema.record_model`, `schema.fields`, `schema.id_normalization`, `schema.metadata_grammar`, and `schema.discovery`;
  - final scoped validation: `[strict] All 5 file(s) OK.`;
  - final independent review verdict: `PASS`;
  - no blocking, major, minor, or advisory findings remained;
  - public list/get representation remains delegated to W004;
  - exact diagnostic identifiers, severity, and source-location representation remain delegated to W006.
- `DRMCP-TASK-MCP-003-05` opened on 2026-06-27 for final cross-spec synchronization, validation, independent review, and closure.
- T05 normative synchronization changed:
  - `spec:drmcp.design_records_mcp.overview`;
  - `spec:drmcp.design_records_mcp.responsibility_boundary`;
  - `spec:drmcp.design_records_mcp.schema.overview`.
- `spec:drmcp.design_records_mcp.namespace_scanning` was rechecked against T02 and required no change.
- Initial final review returned `NEEDS REVISION` with one major W004 ownership-boundary finding in `schema.record_source`.
- T05 removed concrete read-operation and response claims from `schema.record_source`; headings and body remain source material only.
- Post-correction validation passed for all 30 DRMCP spec files on 2026-06-27.
- Final independent re-review verdict: `PASS`.
  - Previous Finding 1: `CLOSED`.
  - Blocking findings: none.
  - Major findings: none.
  - Minor findings: none.
  - Advisories: none.
- Final delegated residual scope:
  - W004 owns query, exact retrieval, body inclusion, successful-record behavior, and normal public response representation;
  - W005 owns resolver behavior and legacy fallback lookup;
  - W006 owns validation execution details, diagnostic taxonomy, severity, source-location representation, and exceptional physical-path exposure;
  - later Work Items own fixtures, implementation, and automated tests.
- `DRMCP-TASK-MCP-003-05` is `done`.
- `DRMCP-WORK-MCP-003` closed as `done` on 2026-06-27.
