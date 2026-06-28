# DRMCP-TASK-MCP-003-01: Establish discovery and active-index correction baseline

- **id**: DRMCP-TASK-MCP-003-01
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-003
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**: []
- **outputs**: []

## Goal

Establish the exact contract-correction baseline for current discovery and active-index realignment.

Identify the authoritative inputs, affected DRMCP specs, retained behavior, rewritten behavior, and downstream exclusions before normative spec edits begin.

## Work

- Read `DRMCP-ADR-MCP-001`, `DRMCP-REQ-MCP-001`, `DRMCP-INV-MCP-002`, and `PRODUCT-WORK-SPEC-014` as the decision and dependency baseline.
- Read the PRODUCT namespace, repository-layout, and spec-format authorities referenced by `DRMCP-WORK-MCP-003`.
- Inspect each DRMCP spec named in the Work Item impact scope.
- Map findings F01, F06, F11, F12, and F13 to exact files and sections.
- Classify each affected statement as retain, narrow, rewrite, remove, or delegate.
- Confirm the configured current-root, app association, active-index, current-spec source, canonical identity, duplicate identity, and current/legacy separation boundaries.
- Record query, retrieval, resolver, validation, fixture, implementation, and physical-path response behavior as explicit downstream exclusions.
- Confirm which later changes remain blocked by `PRODUCT-WORK-SPEC-014`.

This Task does not edit normative DRMCP specs.
It establishes the reviewed manifest used by later Tasks.

## Done condition

- Every finding owned by this Work Item maps to at least one exact file and section.
- Every impacted DRMCP spec has a documented disposition.
- Retained behavior and rewritten behavior are separated.
- PRODUCT-owned semantics are identified as consumed authorities rather than DRMCP-owned rules.
- The `PRODUCT-WORK-SPEC-014` dependency is limited to the accepted `V01-SPEC-*` compatibility boundary.
- Query, retrieval, resolver, validation, fixtures, implementation, and normal response path exposure remain explicitly excluded.
- No unresolved ownership ambiguity remains for T02 through T05.

## Verification

- Compared the manifest against the Work Item impact scope and completion condition.
- Compared finding coverage against F01, F06, F11, F12, and F13 in `DRMCP-INV-MCP-002`.
- Confirmed that no normative DRMCP spec changed during this Task.
- Reviewed the final baseline for missing files, duplicated ownership, and premature downstream decisions.

## Evidence

### Authority baseline

| concern | consumed authority |
|---|---|
| App namespace semantics | `spec:product.design_records.namespace_model.app_namespaces` |
| Current sequential artifact ID grammar | `spec:product.design_records.namespace_model.artifact_id_grammar` |
| Records-root and kind placement | `spec:product.design_records.repository_layout` |
| App-independent path patterns | `spec:product.design_records.repository_layout.record_discovery_paths` |
| Current spec visible metadata | `spec:product.design_records.spec_format.document_shape` |
| Current spec canonical identity | `spec:product.design_records.spec_format.spec_id_as_ref` |
| DRMCP current-first and separate-index decision | `DRMCP-ADR-MCP-001` |
| Read-baseline required outcome | `DRMCP-REQ-MCP-001` |
| Contract-audit findings | `DRMCP-INV-MCP-002` |
| Brewprint legacy-family correction gate | `PRODUCT-WORK-SPEC-014` |

DRMCP owns configured root loading, parsing mappings, active-index construction, duplicate handling, and diagnostic representation.
DRMCP does not redefine PRODUCT namespace, layout, artifact grammar, or spec-format semantics.

### Finding coverage

| finding | exact affected sections | T02-T05 handling |
|---|---|---|
| F01 | `overview.md` `Record scope` and `Tool boundary`; `responsibility-boundary.md` `Boundary against existing brewprint MCP`; `schema/discovery.md` `Current contract`; `schema/record-source.md` `Record sources` and `Metadata source per record kind`; `schema/record-model.md` `Internal record model` and `Bootstrap records`; `schema/fields.md` `Common fields`, `id`, `status`, dependency fields, and title extraction; `schema/id-normalization.md` `Public ID` and `Bare ID grammar` | Replace YAML-front-matter spec authority with H1-adjacent metadata and path-derived `spec:` identity. |
| F06 | `namespace-scanning.md` all normative sections; `overview.md` `Record scope`; `schema/discovery.md` current prefix-dependent path statement; `schema/id-normalization.md` all normative sections; `schema/fields.md` `id` and title extraction | Replace automatic root discovery and mechanical `namespace_prefix` behavior with configured current roots and explicit app association. Delegate list filters and validation scope. |
| F11 | `schema/metadata-grammar.md` all kind grammars; `schema/record-source.md` non-spec metadata-source rows; `schema/fields.md` workflow ID and field rules | Retain parser syntax where useful. Point semantic requiredness, ID grammar, and lifecycle rules to PRODUCT authorities. |
| F12 | `schema/fields.md` `status`; `schema/record-source.md` spec source rows; `schema/record-model.md` status source | Replace YAML status source and stale spec status vocabulary. Delegate status-invalid diagnostic taxonomy to W006. |
| F13 | `overview.md` `Record scope`; `schema/discovery.md` spec inclusion rules; `schema/record-source.md` spec source rows; `schema/record-model.md` `Bootstrap records` | Remove silent adoption or omission based on legacy `design_record`. Active discovery accepts only current spec format. Legacy spec aliases do not enter either accepted index. |

### Affected-file manifest

| file | disposition | required correction boundary |
|---|---|---|
| `overview.md` | rewrite affected sections | Make the overview navigation-first. Summarize configured active discovery and current spec source without restating detailed schema. |
| `responsibility-boundary.md` | narrow | Correct the DRMCP data-source row. Leave tool response and filesystem-operation boundaries to their owning Work Items. |
| `namespace-scanning.md` | rewrite | Define configured current roots, explicit app association, no repository-wide auto-discovery, and active versus legacy root separation. |
| `schema/overview.md` | narrow | Keep navigation. Remove or correct summaries that imply prefix stripping, YAML spec metadata, or mixed active and legacy indexing. |
| `schema/discovery.md` | rewrite | Define active inclusion over supported current kinds. Use PRODUCT path patterns and current spec source rules. |
| `schema/record-source.md` | rewrite | Define H1-adjacent current metadata sources. Remove YAML front matter as an active spec source. |
| `schema/record-model.md` | rewrite | Correct internal identity and source fields. Remove legacy bootstrap authority. Keep retrieval wrapper details delegated to W004. |
| `schema/fields.md` | rewrite affected sections | Correct current common and spec field sources. Point status and artifact semantics to PRODUCT. Keep final query response shape delegated to W004. |
| `schema/id-normalization.md` | rewrite | Replace `namespace_prefix + bare_id` and `SPEC-*` assumptions with complete current sequential IDs and path-derived `spec:` refs. Keep resolver input behavior delegated to W005. |
| `schema/metadata-grammar.md` | narrow | Retain DRMCP parser delimiters and value parsing. Remove duplicated semantic grammar and requiredness authority. |

`schema/id-normalization.md` was missing from the original Work Item impact list.
The Work Item now includes that spec in `impact_refs`, `Boundary`, `Impact Scope`, and T04 scope.

### Retained behavior

- The active read surface covers decision, spec, investigation, requirement, work item, and task records.
- Current sequential records use H1-adjacent bullet metadata.
- Current specs use visible H1-adjacent metadata.
- Markdown H1 remains the title source.
- DRMCP may retain parser-level metadata block delimitation and scalar/list normalization.
- PRODUCT path patterns remain discovery inputs rather than DRMCP-owned placement rules.
- One active index may span multiple configured current app roots.
- Current cross-app records may coexist in the active index.
- Duplicate canonical identities must fail deterministically.
- Filesystem order must not choose a duplicate winner.

### Rewritten behavior

- Replace repository-wide `*/records/` auto-discovery with configured current roots only.
- Replace mechanical app-directory uppercasing as identity authority with explicit current-root and app association.
- Replace `namespace_prefix + bare_id` construction with complete current app-aware sequential IDs.
- Replace `V01-SPEC-*` and `SPEC-*` spec identity with path-derived active `spec:` refs.
- Replace YAML `design_record` inclusion with current H1-adjacent metadata and repository placement.
- Reject YAML front matter as an active current spec metadata source.
- Do not silently adopt invalid current specs through a legacy alias.
- Keep active and optional legacy indexes as separate structures and operational scopes.
- Keep physical paths internal to index provenance. Normal response exposure remains outside this Work Item.

### Downstream exclusions

| excluded concern | owner |
|---|---|
| Listing filters, ordering, limits, and result shape | `DRMCP-WORK-MCP-004` |
| Exact retrieval and batch partial-result behavior | `DRMCP-WORK-MCP-004` |
| Current-first resolution and accepted legacy fallback lookup | `DRMCP-WORK-MCP-005` |
| Validation execution, severity, diagnostic codes, and normal response path exposure | `DRMCP-WORK-MCP-006` |
| Validation Work Item disposition | `DRMCP-WORK-MCP-007` and `PRODUCT-WORK-SPEC-015` |
| Current and legacy fixtures | `DRMCP-WORK-MCP-008` |
| Parser, index, and read-tool implementation | `DRMCP-WORK-MCP-009` |
| Configured legacy fallback implementation | `DRMCP-WORK-MCP-010` |

### Dependency result

`PRODUCT-WORK-SPEC-014` is still `not_started`.
The current `spec:product.brewprint.compatibility.legacy_id_compatibility` still accepts `V01-SPEC-*` and retains a compatibility-only spec identity note.

T02 must not write the current/legacy separation contract while that active PRODUCT profile remains contradictory.
The remaining W003 flow is blocked only on removal of that obsolete compatibility authority.

### Verification result

- Finding coverage: F01, F06, F11, F12, and F13 all mapped.
- DRMCP affected-file count: 10.
- Added affected file: `schema/id-normalization.md`.
- Normative DRMCP spec changes during T01: none.
- Workflow changes: this Task created; parent Work Item linked; Work Item impact scope corrected.
- Residual ownership ambiguity for T02-T05: none.
