# DRMCP-TASK-MCP-003-04: Define shared record model and invalid-source integration contract

- **id**: DRMCP-TASK-MCP-003-04
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-003
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-003-02
  - DRMCP-TASK-MCP-003-03
- **outputs**:
  - spec:drmcp.design_records_mcp.schema.record_model
  - spec:drmcp.design_records_mcp.schema.fields
  - spec:drmcp.design_records_mcp.schema.id_normalization
  - spec:drmcp.design_records_mcp.schema.metadata_grammar
  - spec:drmcp.design_records_mcp.schema.discovery

## Goal

Define the shared DRMCP rules for what current source is addressable by canonical ID and what source remains available only to validation.

Align current fields and identity parsing with PRODUCT authority. Remove stale YAML `design_record` and V01-only assumptions without defining public response, diagnostic taxonomy, resolver, fixture, or implementation behavior.

## Work

- Define normal index behavior for a valid source with one canonical ID.
- Keep a source addressable when one canonical ID is known but H1, metadata, or document content is invalid.
- Keep a source validation-only when no canonical ID can be determined.
- Create no winner when multiple sources produce the same canonical ID.
- Retain repository-relative source paths for validation and repair diagnostics.
- Define common-field presence and invalid or missing field handling.
- Correct current spec field mapping and remove stale `design_record.*` source claims.
- Correct shared current sequential-artifact identity descriptions against PRODUCT authority.
- Consolidate metadata grammar only where required by the shared model.
- Keep public list and get response representation delegated to `DRMCP-WORK-MCP-004`.
- Keep exact diagnostic identifiers and severity delegated to `DRMCP-WORK-MCP-006`.

This Task does not define query behavior, resolver behavior, fixtures, implementation, tests, or authoring transactions.

## Done condition

- A valid source with one canonical ID creates a normal addressable record.
- A source with one canonical ID remains addressable when other content is invalid, without fabricated fields.
- A source without a determinable canonical ID remains validation-only.
- Duplicate canonical identity creates no winner.
- Invalid and conflicting sources retain repository-relative paths for validation and repair diagnostics.
- Common and kind-specific fields follow current PRODUCT authority.
- Current spec fields no longer derive from YAML `design_record.*`.
- Shared identity text does not treat current spec refs or current sequential artifact IDs as V01 `namespace_prefix` constructions.
- T03 discovery and identity decisions remain unchanged.
- `DRMCP-WORK-MCP-004` retains public response representation ownership.
- `DRMCP-WORK-MCP-006` retains exact diagnostic and severity ownership.
- Query, resolver, fixture, implementation, and test concerns remain excluded.
- Normative changes receive scoped static verification and independent review.

## Verification

- Compare the final shared model against `DRMCP-TASK-MCP-003-02` and `DRMCP-TASK-MCP-003-03` accepted contracts.
- Compare current sequential artifact identity text against `spec:product.design_records.namespace_model.artifact_id_grammar`.
- Compare current spec field mapping against `spec:product.design_records.spec_format.document_shape` and `spec:product.design_records.spec_format.spec_id_as_ref`.
- Confirm no active YAML `design_record.*`, `V01-SPEC-*`, or V01-only current identity authority remains in the changed shared specs.
- Confirm public response, diagnostics, resolver, fixture, implementation, and test contracts are not introduced.
- Run scoped static validation after normative edits.
- Obtain independent review and close all blocking and major findings before marking this Task done.

## Evidence

- Upstream Tasks `DRMCP-TASK-MCP-003-02` and `DRMCP-TASK-MCP-003-03` are `done`.
- T04 opened on 2026-06-27 with initial outputs limited to `schema.record_model`, `schema.fields`, `schema.id_normalization`, and `schema.metadata_grammar`.
- `schema.discovery` was added to outputs after independent review identified a direct contradiction between sequential candidate inclusion and the shared invalid-source contract.
- `schema.record_source` remains outside outputs because no direct contradiction required a normative change.
- Closed T02 and T03 root, discovery, identity, invalid-source, duplicate-conflict, and spec-status decisions are consumed without reopening them.
- Initial shared-model contradiction baseline:
  - `schema.record_model` conflates source parsing, indexed record data, and retrieval response content.
  - `schema.record_model` and `schema.fields` require valid `id`, `kind`, `title`, and `status` even for retained invalid sources.
  - Current spec detail and relation fields still cite obsolete YAML `design_record.*` sources.
  - Sequential artifact identity text still models current IDs through a V01-style `namespace_prefix` instead of the PRODUCT app-aware public ID grammar.
  - Cross-kind invalid-source, conflict-group, and provenance integration is not defined.
- Accepted decision 1: retain an invalid current source when DRMCP can determine one unique canonical ID.
  - Preserve only values that were actually parsed from the source.
  - Do not fabricate missing `title`, `status`, `date`, or kind-specific values.
  - A source without one determinable canonical ID remains validation-only and is reported through source provenance.
  - Duplicate canonical identity still creates no winner entry under T02 and T03.
  - T04 does not define a larger invalid-state taxonomy.
- Authority-derived decision 2: current spec fields follow PRODUCT spec-format authority.
  - Common record data uses canonical `id`, record `kind = spec`, H1 title when available, `status`, and `date`.
  - Spec-specific data uses H1 `spec_kind`, metadata `parent`, and kind-conditional `contract_class`.
  - Current specs do not fabricate `depends_on`, `supersedes`, or `migrated_to_spec`; obsolete `design_record.*` mappings are removed.
  - Public missing-field representation remains delegated to W004.
- Authority-derived decision 3: current sequential artifact identity uses the complete canonical app-aware artifact ID directly.
  - General grammar for ADR, investigation, requirement, and work item: `<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<SEQUENCE>`.
  - General grammar for task: `<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>`.
  - ADR and investigation H1 IDs and filename prefixes are compared as complete canonical IDs.
  - Requirement, work-item, and task metadata IDs, H1 IDs, and filename prefixes are compared as complete canonical IDs.
  - DRMCP does not strip or reattach a runtime `namespace_prefix` for current records.
  - Brewprint-specific IDs are examples only and do not define the generic DRMCP contract.
  - Existing issued legacy V01 IDs remain a separate configured-legacy concern and do not define the current shared model.
- Normative reflection completed on 2026-06-27:
  - `schema.record_model` now states the four concrete index and validation outcomes and removes retrieval-response content from the internal model.
  - `schema.fields` now includes common `date`, current spec `spec_kind` / `parent` / conditional `contract_class`, and removes obsolete spec dependency and migration fields.
  - `schema.id_normalization` now validates complete app-aware canonical IDs directly and uses generic grammar before clearly marked synthetic examples.
  - `schema.metadata_grammar` now maps current PRODUCT metadata fields and complete canonical relation IDs without V01 prefix stripping or YAML fallback.
  - Changed normative specs use `<APP_NAMESPACE>`, `<ARTIFACT_KIND>`, `<DOMAIN_NAMESPACE>`, and sequence placeholders as the contract. Concrete IDs appear only in explicitly labeled synthetic examples.
- Static reread found no current-rule occurrence of V01 IDs, runtime `namespace_prefix` reconstruction, or YAML `design_record.*` mapping in the initial four changed specs.
- Initial scoped validation executed against `schema.record_model`, `schema.fields`, `schema.id_normalization`, and `schema.metadata_grammar` with `--strict --no-color`: `[strict]  All 4 file(s) OK.`
- Independent review verdict: `NEEDS REVISION`; no blocking findings, three major findings, and one minor finding.
  - M1: sequential identity authority was ambiguous for requirement, work-item, and task records, and `may remain addressable` weakened the accepted contract.
  - M2: `schema.discovery` incorrectly required valid metadata for sequential candidate inclusion while also denying metadata-gated inclusion.
  - M3: current spec unknown-key rejection from T03 was not stated normatively.
  - m1: validator evidence was not yet recorded.
- Corrections applied after review:
  - H1 is now the canonical identity authority for every current sequential artifact.
  - Requirement, work-item, and task metadata `id` values and all filename ID prefixes are consistency checks only; they are never fallback identity sources.
  - A valid unique H1 ID keeps a sequential source addressable despite missing, malformed, or mismatched metadata.
  - A sequential source without a valid H1 ID remains validation-only even if metadata or filename resembles an ID.
  - `schema.discovery` now treats PRODUCT path matching as candidate inclusion and moves H1 and metadata validity to addressability and validation behavior.
  - Current spec unknown metadata markers are explicitly source-format violations and are neither ignored, retained as extensions, nor repaired.
  - The initial 4-file validator result is recorded above.
- First revalidation after the initial review executed against all five changed specs with `--strict --no-color`: `[strict]  All 5 file(s) OK.`
- First re-review verdict: `NEEDS REVISION`.
  - M1: `CLOSED`.
  - M2: remained `OPEN` only because the ADR candidate condition allowed an unspecified configured compatibility path instead of the PRODUCT-defined flat ADR compatibility pattern.
  - M3: `CLOSED`.
  - m1: `CLOSED`.
  - No new blocking, major, minor, or advisory findings were reported beyond the remaining M2 point.
- Final M2 correction applied:
  - `schema.discovery` now limits decision candidates to the PRODUCT ADR discovery path, including only the PRODUCT-defined flat ADR compatibility pattern.
  - A generic configured compatibility path is no longer a current candidate-inclusion authority.
- Final scoped validation after the final M2 correction executed against all five changed specs with `--strict --no-color`: `[strict]  All 5 file(s) OK.`
- Final independent review verdict: `PASS`.
  - M1: `CLOSED`.
  - M2: `CLOSED`.
  - M3: `CLOSED`.
  - m1: `PARTIALLY CLOSED` only until the final validator result was synchronized; this synchronization is now complete.
  - New blocking findings: none.
  - New major findings: none.
  - New minor findings: none.
  - Advisories: none.
- Closure assessment: T04 closure readiness `ready`; final five-file validator result accepted.
- Task closed as `done` on 2026-06-27 after final validator and review evidence synchronization.
