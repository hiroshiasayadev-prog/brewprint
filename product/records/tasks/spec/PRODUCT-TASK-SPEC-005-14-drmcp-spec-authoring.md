# PRODUCT-TASK-SPEC-005-14: DRMCP — author design-records-mcp/ spec files

- **id**: PRODUCT-TASK-SPEC-005-14
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 3d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-13
- **outputs**:
  - `drmcp/records/spec/design-records-mcp/overview.md`
  - `drmcp/records/spec/design-records-mcp/responsibility-boundary.md`
  - `drmcp/records/spec/design-records-mcp/resolver.md`
  - `drmcp/records/spec/design-records-mcp/namespace-scanning.md`
  - `drmcp/records/spec/design-records-mcp/mvp-scope.md`
  - `drmcp/records/spec/design-records-mcp/schema/overview.md`
  - `drmcp/records/spec/design-records-mcp/schema/record-source.md`
  - `drmcp/records/spec/design-records-mcp/schema/metadata-grammar.md`
  - `drmcp/records/spec/design-records-mcp/schema/fields.md`
  - `drmcp/records/spec/design-records-mcp/schema/id-normalization.md`
  - `drmcp/records/spec/design-records-mcp/schema/discovery.md`
  - `drmcp/records/spec/design-records-mcp/schema/record-model.md`
  - `drmcp/records/spec/design-records-mcp/schema/diagnostics.md`
  - `drmcp/records/spec/design-records-mcp/schema/authoring-guidance-source.md`
  - `drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md`
  - `drmcp/records/spec/design-records-mcp/tools/overview.md`
  - `drmcp/records/spec/design-records-mcp/tools/list-records.md`
  - `drmcp/records/spec/design-records-mcp/tools/get-record.md`
  - `drmcp/records/spec/design-records-mcp/tools/get-records.md`
  - `drmcp/records/spec/design-records-mcp/tools/list-authoring-guides.md`
  - `drmcp/records/spec/design-records-mcp/tools/get-authoring-guidance.md`
  - `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`
  - `drmcp/records/spec/design-records-mcp/tools/validate-records.md`
  - `drmcp/records/spec/design-records-mcp/tools/suggest-next-record.md`
  - `drmcp/records/spec/design-records-mcp/tools/authoring-transaction-model.md`
  - `drmcp/records/spec/design-records-mcp/tools/propose-record-create.md`
  - `drmcp/records/spec/design-records-mcp/tools/propose-record-update.md`
  - `drmcp/records/spec/design-records-mcp/tools/get-proposed-write.md`
  - `drmcp/records/spec/design-records-mcp/tools/accept-proposed-write.md`
  - `drmcp/records/spec/design-records-mcp/tools/discard-proposed-write.md`

## Goal

Author all 30 spec files under `drmcp/records/spec/design-records-mcp/` using `drmcp/old/` as the content source. Each file must pass `--strict` validation. All H1 titles, H2 section titles, and table headers must be English (source files are Japanese throughout). Content must be complete — no section placeholders.

`drmcp/old/tools.md` (1944 lines, 13 distinct tool contracts) and `drmcp/old/schema.md` (834 lines) are each split into multiple target files, matching the granularity of BPDSL's `mcp/tools/*.md` split (PRODUCT-TASK-SPEC-005-05).

## Work

| file | kind | source in drmcp/old/ | notes |
|---|---|---|---|
| `overview.md` | Overview | `overview.md` §目的 (trimmed) | Root entry point, `parent: -`. `## Topics` → all 14 children below (responsibility-boundary, resolver, namespace-scanning, mvp-scope, schema.overview, tools.overview). Detailed content from §既存brewprint MCPとの責務境界 / §filesystem tool との責務境界 moves out — overview keeps only a boundary summary table. Do not restate the boundary rule independently; reference `spec:product.concepts.project_artifact_model` boundary table instead (per PRODUCT-TASK-SPEC-004-02 review note). |
| `responsibility-boundary.md` | Reference | `overview.md` §既存brewprint MCPとの責務境界 + §filesystem tool との責務境界 | Both MCP-vs-MCP and MCP-vs-filesystem-tool boundaries. |
| `resolver.md` | Reference | `overview.md` §Resolver responsibility | |
| `namespace-scanning.md` | Reference | `overview.md` §Record scanning と namespace prefix | Includes records_root derivation, kind-別 prefix table, multi-root scan behavior. |
| `mvp-scope.md` | Reference | `overview.md` §MVP tool set + §MVP外 | |
| `schema/overview.md` | Overview | — | New file. `## Topics` → all 9 schema/ siblings. |
| `schema/record-source.md` | Reference | `schema.md` §Record source + §Metadata source | |
| `schema/metadata-grammar.md` | Reference | `schema.md` §ADR/Investigation/Workflow artifact bullet metadata 文法 (incl. embedded V01-ADR-076 / V01-INV-MCP-001 example blocks) | |
| `schema/fields.md` | Reference | `schema.md` §Field definitions + §id/kind/status/depends_on/supersedes/migrated_to_spec + §Title extraction | |
| `schema/id-normalization.md` | Reference | `schema.md` §ID normalization model (Public ID + Bare ID grammar) | |
| `schema/discovery.md` | Reference | `schema.md` §Discovery and index inclusion | |
| `schema/record-model.md` | Reference | `schema.md` §Record model + §Bootstrap metadata (incl. embedded V01-ADR-076 example block) | |
| `schema/diagnostics.md` | Reference | `schema.md` §Diagnostic category | Largest schema sub-file (~135 lines). |
| `schema/authoring-guidance-source.md` | Reference | `schema.md` §Authoring guidance source model | |
| `schema/authoring-transaction-schema.md` | Reference | `schema.md` §Authoring transaction schema concepts (target identity / proposal model / body cache model / metadata block replacement target / section selector model) | |
| `tools/overview.md` | Overview | `tools.md` §Tool set + §Common response conventions + §Error handling + §Authoring write boundary | `## Topics` → all 13 tool files + `authoring-transaction-model.md`. |
| `tools/list-records.md` | Contract (interface) | `tools.md` §`list_records` | `## Request` / `## Response` / `## Errors`. `contract_class: interface`. |
| `tools/get-record.md` | Contract (interface) | `tools.md` §`get_record` | Has both "Response without body" and "Response with body" — keep both under `## Response`. |
| `tools/get-records.md` | Contract (interface) | `tools.md` §`get_records` | |
| `tools/list-authoring-guides.md` | Contract (interface) | `tools.md` §`list_authoring_guides` | |
| `tools/get-authoring-guidance.md` | Contract (interface) | `tools.md` §`get_authoring_guidance` | |
| `tools/resolve-reference.md` | Contract (interface) | `tools.md` §`resolve_reference` | |
| `tools/validate-records.md` | Contract (interface) | `tools.md` §`validate_records` | Includes §Diagnostic categories sub-section — keep as part of `## Response`, cross-ref `schema/diagnostics.md` rather than duplicating the category table. |
| `tools/suggest-next-record.md` | Contract (interface) | `tools.md` §`suggest_next_record` | |
| `tools/authoring-transaction-model.md` | Reference | `tools.md` §Authoring transaction model (common authoring response fields / body source and body cache / proposal validation affected record set / diff_mode) | Shared concepts for the 5 write tools below — not itself a tool contract, so `contract_class` is not `interface`. |
| `tools/propose-record-create.md` | Contract (interface) | `tools.md` §`propose_record_create` | |
| `tools/propose-record-update.md` | Contract (interface) | `tools.md` §`propose_record_update` | Largest tool file (~390 lines). |
| `tools/get-proposed-write.md` | Contract (interface) | `tools.md` §`get_proposed_write` | |
| `tools/accept-proposed-write.md` | Contract (interface) | `tools.md` §`accept_proposed_write` | |
| `tools/discard-proposed-write.md` | Contract (interface) | `tools.md` §`discard_proposed_write` | |

Run `validate_spec.py drmcp/records/spec/design-records-mcp/ --strict` after each file is authored. Target: 0 errors on all 30 files before closing this task.

## Done condition

| item | done when |
|---|---|
| all 30 files exist | All output files listed above are present. |
| strict validation | `validate_spec.py drmcp/records/spec/design-records-mcp/ --strict` exits 0. |
| English titles | All H1/H2 titles and table headers are English. No Japanese remains. |
| content complete | All content from `drmcp/old/` is accounted for — either migrated or explicitly noted as moved/merged elsewhere. |
| YAML front matter removed | No YAML front matter in any output file. |
| `design_record` block removed | The legacy `design_record:` front-matter block (V01-SPEC-* ids) is not carried into any output file — replaced by H1-adjacent `id:` per the accepted format. |

## Verification

- Cross-check each `parent:` ref against the `## Topics` row that declares it as a child.
- Confirm `overview.md` `## Topics` covers all 14 direct children.
- Confirm `schema/overview.md` `## Topics` covers all 9 schema/ siblings.
- Confirm `tools/overview.md` `## Topics` covers all 13 tool contracts + `authoring-transaction-model.md`.
- Confirm all 13 tool files have `contract_class: interface` and `## Request` / `## Response` / `## Errors` sections.
- Confirm `tools/validate-records.md` does not duplicate the diagnostic category table — it should reference `schema/diagnostics.md`.

## Evidence

All 30 output files authored and passing `validate_spec.py drmcp/records/spec/design-records-mcp/ --strict` (exit 0).

Output breakdown:
- Root: `overview.md` — 6 Topics (responsibility-boundary, resolver, namespace-scanning, mvp-scope, schema/overview, tools/overview). Note: the "14 direct children" stated in the Verification section was a miscalculation in this task file; the work table itself listed the 6 correct names.
- Flat reference files: `responsibility-boundary.md`, `resolver.md`, `namespace-scanning.md`, `mvp-scope.md` — all sourced from `drmcp/old/overview.md`.
- Schema: `schema/overview.md` (9 Topics) + 9 Reference files — all sourced from `drmcp/old/schema.md`.
- Tools: `tools/overview.md` (14 Topics: 13 contracts + authoring-transaction-model) + 14 files — all sourced from `drmcp/old/tools.md`.

Deferred relocation candidates flagged per PRODUCT-INV-SPEC-004 (note added in each file's `## What this is`):
- `resolver.md` — namespace_prefix derivation and public/bare ID grammar (PRODUCT-owned semantics)
- `namespace-scanning.md` — namespace_prefix derivation algorithm, multi-root scan, public/bare ID grammar (PRODUCT-owned semantics)
- `schema/id-normalization.md` — public ID and bare ID grammar (PRODUCT-owned semantics)
- `schema/discovery.md` — discovery path pattern conventions (partial; DRMCP kind-filter stays)

Cross-reference note applied: `overview.md §Tool boundary` cross-refs `spec:product.concepts.project_artifact_model` per PRODUCT-TASK-SPEC-004-02 review note (does not independently restate the boundary rule).

Validator fixes applied during authoring: 3 `REFERENCE_NO_TABLE` errors in `resolver.md`, `schema/authoring-guidance-source.md`, `schema/metadata-grammar.md` — resolved by adding a Markdown table to the `## Current contract` section in each file.
