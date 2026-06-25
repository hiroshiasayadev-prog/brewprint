# DRMCP-INV-MCP-002: Design Records MCP contract consistency and realignment audit

- **status**: concluded
- **date**: 2026-06-25
- **trigger**: PRODUCT-WORK-SPEC-012
- **scope**: Audit the current Design Records MCP contracts, follow-up records, and implementation-facing assumptions against the accepted PRODUCT Design Records semantics. Classify contradictions, stale assumptions, missing contracts, and overlapping ownership before DRMCP redesign or reimplementation begins.
- **non_scope**: Editing DRMCP specs or implementation, adopting final design decisions, reopening PRODUCT-WORK-SPEC-012, performing BPDSL migration, or defining UI behavior.
- **source_refs**:
  - PRODUCT-WORK-SPEC-012
  - PRODUCT-ADR-SPEC-001
  - DRMCP-REQ-MCP-001
  - DRMCP-REQ-MCP-002
  - spec:product.design_records.authoring_standards
  - spec:product.design_records.namespace_model
  - spec:product.design_records.repository_layout
  - spec:product.design_records.spec_format
  - spec:product.design_records.traceability
  - spec:product.design_records.artifact_model
  - spec:product.brewprint.compatibility
  - spec:drmcp.design_records_mcp.overview
- **follow_up_candidates**:
  - DRMCP-ADR-MCP-new
  - DRMCP-WORK-MCP-new
- **related_requirements**:
  - DRMCP-REQ-MCP-001
  - DRMCP-REQ-MCP-002
- **related_work_items**:
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002
- **related_adrs**:
  - PRODUCT-ADR-SPEC-001
- **related_specs**:
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.responsibility_boundary
  - spec:drmcp.design_records_mcp.resolver
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.mvp_scope
  - spec:drmcp.design_records_mcp.schema.overview
  - spec:drmcp.design_records_mcp.tools.overview

## Investigation scope

This investigation audited the active Markdown contract files under `drmcp/records/spec/design-records-mcp/`, the current DRMCP planning records listed in the trigger request, and the current DRMCP implementation and tests under `drmcp/src/`.

The audit was performed after the PRODUCT semantic ownership restructuring represented by `PRODUCT-WORK-SPEC-012` and accepted by `PRODUCT-ADR-SPEC-001`. The accepted ownership baseline used here is:

- PRODUCT owns app-independent Design Records semantics, authoring standards, artifact roles, app/domain namespace semantics, ID grammar, repository layout semantics, traceability semantics, and spec format.
- Brewprint owns current profile and compatibility facts.
- DRMCP owns concrete parser, discovery, indexing, normalization, resolver, validation, diagnostics, MCP tool APIs, and authoring transaction behavior.
- BPDSL is outside this investigation.

The audit classifies current statements and implementation evidence. It does not adopt repaired contracts and does not change any spec, requirement, work item, implementation, test, fixture, `v01/`, `CLAUDE.md`, or `prompt_chappy.md` file.

## Out of scope

- Redesigning or fixing DRMCP contracts.
- Editing DRMCP specs, requirements, work items, tasks, ADRs, implementation, tests, or fixtures.
- Reopening `PRODUCT-WORK-SPEC-012` or `PRODUCT-ADR-SPEC-001`.
- Deciding legacy compatibility policy for spec YAML front matter.
- Performing BPDSL analysis.
- Creating follow-up artifacts.

## Background

`PRODUCT-WORK-SPEC-012` closed the PRODUCT semantic-layer restructuring. The resulting PRODUCT specs moved app-independent Design Records semantics into current `product/records/spec/design-records/**` areas and made path-derived `spec:` identity plus H1-adjacent metadata the canonical new/migrated spec format.

The current DRMCP contract tree predates parts of that restructuring. It contains historical read-only MVP scope, later authoring transaction contracts, multi-root namespace additions, legacy `V01-` examples, and spec parsing rules based on YAML front matter and `design_record` metadata.

Two existing DRMCP requirements are still relevant but partial:

- `DRMCP-REQ-MCP-001` covers multi-root and multi-namespace query, resolve, identity, and validation gaps.
- `DRMCP-REQ-MCP-002` covers namespace-aware authoring transaction conformance.

`DRMCP-INV-MCP-001` remains useful evidence for multi-root gaps, but its source references and examples are historical. `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002` capture real validation work, but they were authored before the final PRODUCT spec-format and Topics table contracts stabilized.

Current `spec:product.design_records.spec_format.validation_policy` still names `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002` as owners of per-file and Topics-graph validation work. Any decision to supersede or absorb those Work Items therefore depends on a coordinated update to those PRODUCT owner pointers; this investigation does not edit the PRODUCT spec.

## What was investigated

### Startup and authority sources

Read first:

- `prompt_chappy.md`
- `product/records/spec/design-records/authoring-standards/investigation-authoring.md`
- `product/records/spec/design-records/authoring-standards/artifact-boundary.md`
- this investigation record
- `product/records/adr/spec/PRODUCT-ADR-SPEC-001-product-spec-semantic-ownership-boundary.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-012-product-spec-semantic-layer-restructuring.md`

Planning records read:

- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-002-namespace-aware-authoring-transaction-conformance.md`
- `drmcp/records/investigations/mcp/DRMCP-INV-MCP-001-multi-root-multi-namespace-mcp-tool-contract-investigation.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md`

PRODUCT authority sources inspected:

- `spec:product.design_records.authoring_standards` and per-artifact authoring standards
- `spec:product.design_records.namespace_model`
- `spec:product.design_records.repository_layout`
- `spec:product.design_records.spec_format`
- `spec:product.design_records.traceability`
- `spec:product.design_records.artifact_model`
- `spec:product.brewprint.compatibility`

Implementation and test roots inspected:

- `drmcp/src/cmd/design-records-mcp`
- `drmcp/src/internal/designrecords`
- `drmcp/src/internal/designrecordsmcp`

### Active DRMCP spec inventory

The active DRMCP spec inventory contains 30 Markdown files:

- Root files: 5 (`overview.md`, `responsibility-boundary.md`, `namespace-scanning.md`, `resolver.md`, `mvp-scope.md`).
- Schema overview and schema contracts: `schema/overview.md`, `schema/authoring-guidance-source.md`, `schema/authoring-transaction-schema.md`, `schema/diagnostics.md`, `schema/discovery.md`, `schema/fields.md`, `schema/id-normalization.md`, `schema/metadata-grammar.md`, `schema/record-model.md`, `schema/record-source.md`.
- Tools overview and tool contracts: `tools/overview.md`, `tools/authoring-transaction-model.md`, `tools/list-records.md`, `tools/get-record.md`, `tools/get-records.md`, `tools/list-authoring-guides.md`, `tools/get-authoring-guidance.md`, `tools/resolve-reference.md`, `tools/validate-records.md`, `tools/suggest-next-record.md`, `tools/propose-record-create.md`, `tools/propose-record-update.md`, `tools/get-proposed-write.md`, `tools/accept-proposed-write.md`, `tools/discard-proposed-write.md`.

The inventory grouping is 5 root files, 10 `schema/` files, and 15 `tools/` files. The file disposition summary contains exactly one row for each active file.

### Baseline repository state

Initial command:

```text
git status --short
```

Initial output:

```text
 M CLAUDE.md
 M product/records/requirements/PRODUCT-REQ-SPEC-001-mcp-readable-spec-format-and-topic-tree.md
 M product/records/tasks/spec/PRODUCT-TASK-SPEC-009-04-opus-review.md
 M product/records/work-items/spec/PRODUCT-WORK-SPEC-009-format-only-migration-traceability-and-artifact-model-specs.md
 M prompt_chappy.md
?? drmcp/records/investigations/mcp/DRMCP-INV-MCP-002-design-records-mcp-contract-consistency-and-realignment-audit.md
```

The investigation file was already untracked at baseline. The other modified files were treated as unrelated existing worktree changes.

## Findings

### Comprehensive audit matrix

| ID | file and section | current claim | semantic owner | authority source | implementation evidence | issue class | severity | required disposition | follow-up owner |
|---|---|---|---|---|---|---|---|---|---|
| F01 | `overview.md` `Record scope`; `schema/discovery.md`; `schema/record-source.md`; `schema/record-model.md`; `schema/fields.md` | Spec records are indexed from `design_record.id` / `design_record.kind` in YAML front matter. | PRODUCT for canonical spec identity; DRMCP for parser compatibility behavior. | PRODUCT spec-format requires H1-adjacent `id/status/date/parent`, path-derived `spec:` identity, and no YAML front matter for new/migrated specs. | `parser.go`, `authoring.go`, and tests still parse and validate spec YAML front matter. | contradiction | major | rewrite | baseline ADR, DRMCP spec correction, compatibility decision |
| F02 | `tools/resolve-reference.md` `Supported input forms` | Active `spec:` refs are resolved from front matter `semantic_refs` and `sections`. | PRODUCT owns canonical spec refs; DRMCP owns lookup implementation. | PRODUCT traceability and spec-format use path-derived document-level `spec:` refs; section refs are not active. | Resolver implementation still contains legacy semantic ref lookup behavior. | stale_assumption | major | rewrite | DRMCP-REQ-MCP-001 update and resolver spec correction |
| F03 | `mvp-scope.md` `P0 tools`; `tools/overview.md` `Tool set` | All P0 MVP tools are read-only, while authoring transaction tools are also classified as MVP P0 and `accept_proposed_write` may write files. | DRMCP. | Existing DRMCP specs conflict; user request requires classification, not resolution. | `designrecordsmcp/tools.go` exposes authoring tools; `authoring.go` writes on accept. | contradiction | major | decision_required | baseline ADR |
| F04 | `tools/propose-record-create.md` `What this is`; `DRMCP-REQ-MCP-002` `Artifact-kind support` | MVP create supports decision, requirement, work item, and task; spec and investigation create are outside MVP. Requirement says redesigned contract requires spec and investigation support. | PRODUCT owns artifact authoring requirements; DRMCP owns support matrix and implementation phases. | Per-artifact PRODUCT authoring standards include investigation create; spec authoring uses path-derived identity. | `validateCreateKind` rejects spec and investigation create. | specified_unimplemented | major | split | DRMCP-REQ-MCP-002 update and implementation work |
| F05 | `tools/propose-record-update.md` `Metadata block replacement`; `authoring.go` | Spec metadata update targets recognized fields inside YAML front matter. | PRODUCT for spec metadata shape; DRMCP for update operation behavior. | PRODUCT spec-format makes H1-adjacent metadata canonical and front matter prohibited for new/migrated specs. | `replaceSpecMetadata` requires spec front matter and edits `design_record.*`. | contradiction | major | rewrite | baseline ADR and authoring transaction spec correction |
| F06 | `namespace-scanning.md`; `schema/id-normalization.md`; `tools/list-records.md`; `tools/validate-records.md` | Examples and some request forms are expressed through `namespace_prefix = V01-`; multi-root exists but namespace filters are not consistently specified. | PRODUCT owns namespace semantics; DRMCP owns query/filter behavior. | PRODUCT namespace model defines app-aware IDs; `DRMCP-REQ-MCP-001` requires namespace scope decisions. | `Config` and index support multiple records roots; tests include many unprefixed legacy fixtures. | missing_contract | major | rewrite | DRMCP-REQ-MCP-001 update |
| F07 | `tools/list-records.md` `id_range`; `tools/validate-records.md` `Request` | Range rules reject `SPEC-*` and `INV-*` ranges and do not define mixed namespace behavior as a first-class rule. | DRMCP. | `DRMCP-INV-MCP-001` already identified range and default-scope gaps. | `validation_test.go` covers range errors; implementation has family/domain checks. | missing_contract | minor | narrow | DRMCP spec correction |
| F08 | `tools/get-records.md` `Request` | Exact IDs are looked up against the record index without ref resolution or normalization. | DRMCP. | PRODUCT says public IDs and `spec:` refs are canonical reference forms; exact lookup behavior is a tool API choice. | Implementation performs exact batch retrieval. | no_issue | none | retain | none |
| F09 | `resolver.md`; `tools/resolve-reference.md`; `schema/diagnostics.md` | Unsupported direct resolver inputs return `unsupported`; validation may use different severity for metadata fields. | DRMCP. | PRODUCT owns invalid semantic conditions; DRMCP owns response representation. | Implementation has direct resolver responses and validation diagnostics. | no_issue | none | retain | none |
| F10 | `tools/validate-records.md`; `schema/diagnostics.md` | Validation responses name diagnostic categories for metadata integrity, references, workflow relations, proposal-local validation, and required sections. | DRMCP response representation; PRODUCT semantic invalidity. | PRODUCT traceability and authoring standards define invalid states; DRMCP defines diagnostics. | `validation.go` and tests cover required metadata, workflow relations, status, and done-gated sections. | duplicated_authority | minor | narrow | DRMCP diagnostics spec correction |
| F11 | `schema/metadata-grammar.md` | Workflow artifact ID grammar and metadata rules are stated directly in DRMCP. | PRODUCT owns grammar and authoring metadata semantics; DRMCP owns parsing grammar. | PRODUCT authoring standards define ID grammar, create inputs, and metadata fields. | Parser consumes bullet metadata for workflow records. | duplicated_authority | major | narrow | DRMCP schema correction |
| F12 | `schema/fields.md` `status`; `tools/validate-records.md` `Valid status values` | Spec status values are `confirmed`, `draft`, `wip`; spec status is top-level YAML. | PRODUCT owns spec lifecycle semantics. | PRODUCT spec-format current examples use H1-adjacent `status`; compatibility must be decided. | Validation tests check top-level spec status. | stale_assumption | major | rewrite | compatibility decision and spec validation work |
| F13 | `schema/record-model.md` `Bootstrap records`; `overview.md` `Record scope` | Existing specs without `design_record` block are not indexed and do not emit diagnostics. | Brewprint compatibility facts plus DRMCP discovery behavior. | PRODUCT says legacy front matter may warn during migration but canonical new/migrated specs use visible metadata. | Implementation excludes specs lacking `design_record`. | compatibility_decision_required | major | decision_required | baseline ADR |
| F14 | `tools/suggest-next-record.md` | Tool supports decision-only next-record suggestions using stale `V01-ADR-NNN`, global max-number, and flat `v01/records/adr/` assumptions. | DRMCP tool contract, constrained by PRODUCT namespace and repository-layout semantics. | PRODUCT namespace and repository-layout contracts require app-aware, domain-aware ID and placement behavior; `DRMCP-REQ-MCP-002` points toward `new` placeholders and namespace-aware create support. | Tool is still exposed; tests cover decision-only support. | stale_assumption | major | supersede | baseline ADR or Work Item |
| F15 | `tools/authoring-transaction-model.md`; `schema/authoring-transaction-schema.md` | Proposal lifecycle, body cache, affected-record validation, and diff modes are defined. | DRMCP. | PRODUCT does not own transaction mechanics. | `authoring.go` implements proposal store, body cache, diff modes, and proposal-local validation. | no_issue | none | retain | none |
| F16 | `tools/accept-proposed-write.md` | Only `accept_proposed_write` may write repository files; failed accept checks must not write. | DRMCP. | Tool boundary is DRMCP-owned. | `AcceptProposedWrite` writes files only after checks; tests cover accept flow. | no_issue | none | retain | none |
| F17 | `tools/propose-record-update.md` `Named section replacement`; `schema/authoring-transaction-schema.md` `Section selector model` | Named section selectors use ATX headings, ignore YAML/fenced headings, support exact match and case repair for canonical headings. | DRMCP, with PRODUCT-owned canonical heading names. | PRODUCT authoring standards define canonical sections by artifact kind. | `replaceNamedSection` and tests implement exact match, ambiguity, no-op, and heading-case behavior. | implemented_unspecified | minor | narrow | DRMCP spec correction |
| F18 | `schema/diagnostics.md` `Required narrative section policy` | Required narrative sections for done/accepted gates are listed in DRMCP diagnostics. | PRODUCT owns artifact lifecycle/section semantics; DRMCP owns diagnostic emission. | PRODUCT authoring standards define lifecycle gates and placeholder-content rules. | `validation_test.go` covers work item, task, and requirement required sections. | duplicated_authority | major | narrow | diagnostics and authoring-standard pointer correction |
| F19 | `tools/overview.md`; per-tool files | Tool overview repeats response shape and error vocabulary also defined in schema and individual tool contracts. | DRMCP. | PRODUCT Topics table contract permits navigation rows but not duplicated authority. | Tool schemas are also exposed in `designrecordsmcp/tools.go`. | duplicated_authority | minor | decision_required | baseline ADR and DRMCP spec hierarchy cleanup |
| F20 | `schema/authoring-guidance-source.md`; `tools/list-authoring-guides.md`; `tools/get-authoring-guidance.md` | Authoring guidance is contracted as `docs/guides/*.md`, separate from records, while current authoring standards are PRODUCT spec records and implementation/tests read legacy guides under `records/guides/*.md`. | PRODUCT owns current authoring-standard semantics; DRMCP owns guidance retrieval behavior; Brewprint compatibility may own retained legacy guide files. | `schema/authoring-guidance-source.md` declares `docs/guides/*.md`; current repository has no `docs/` guide source; PRODUCT authoring standards are under `product/records/spec/design-records/authoring-standards/`. | `authoring_guidance.go` reads `<primaryRecordsRoot>/guides/*.md`; tests use `v01/records/guides/*.md`. | stale_assumption | major | decision_required | baseline ADR, guidance-source contract correction |
| F21 | `DRMCP-WORK-SPEC-001` | Parser-aware spec validation is deferred to later DRMCP implementation, but dependencies and scope are tied to older PRODUCT records. | DRMCP work planning; PRODUCT validation policy for invalid conditions. | PRODUCT validation policy still points per-file validation to DRMCP work, but with current H1-adjacent spec format. | No current implementation of new/migrated spec H1-adjacent validation found. | planning_record_stale | major | supersede | coordinating Work Item |
| F22 | `DRMCP-WORK-SPEC-002` | Topics graph validation uses old `title/kind/parent/file/summary` table shape. | PRODUCT owns Topics table shape; DRMCP owns graph validator implementation. | PRODUCT Topics table now uses `title / kind / ref / summary`; `file` and row `parent` are invalid for new/migrated specs. | No current graph validator for new Topics table found. | planning_record_stale | major | supersede | coordinating Work Item |
| F23 | `DRMCP-REQ-MCP-001`; `DRMCP-INV-MCP-001` | Multi-root query, resolver, identity, and validation gaps are recorded against old `tools.md`/`schema.md` references. | DRMCP. | Current split specs partially incorporated the findings but not fully. | Multi-root config exists; namespace filtering and compatibility decisions remain incomplete. | stale_assumption | major | rewrite | requirement update |
| F24 | `DRMCP-REQ-MCP-002` | Namespace-aware authoring support states create/update requirements including exact IDs, `new`, task parent scope, generated placement, and support for spec/investigation. | PRODUCT for semantics; DRMCP for authoring implementation. | PRODUCT authoring standards support this direction. | Implementation supports namespace-prefixed workflow create placeholders, task parent scope, and reciprocal updates; spec/investigation create remains absent. | specified_unimplemented | major | narrow | requirement update and implementation work |
| F25 | Implementation tests | Tests pass while encoding two compatibility tracks: `V01-*`-prefixed legacy forms and unprefixed bare fixture IDs such as `REQ-MCP-001`, `WORK-MCP-001`, and `TASK-MCP-001-01`. | DRMCP tests; PRODUCT compatibility facts if retained. | Brewprint legacy compatibility semantics cover `V01-*` forms, but warning/error and YAML-format policy still need decisions; no current PRODUCT or Brewprint compatibility authorization was found for the unprefixed bare fixture IDs, so that track requires a separate acceptance/rejection decision and must not be treated as equivalent to accepted `V01-*` compatibility. | `go test ./drmcp/src/...` passed; tests include `SPEC-*` YAML fixtures, `V01-*` forms, and unprefixed `REQ/WORK/TASK` fixtures. | compatibility_decision_required | major | decision_required | compatibility fixture work |
| F26 | `overview.md`; `schema/overview.md`; `tools/overview.md` `## Topics` tables | The three overview specs that contain `## Topics` tables use current `title / kind / ref / summary` shape. | PRODUCT spec-format owns Topics table shape; DRMCP owns topic content. | PRODUCT Topics table contract. | File inspection found the three overview Topics tables in current shape. | no_issue | none | retain | none |

### Finding counts

| issue class | count |
|---|---:|
| `no_issue` | 5 |
| `contradiction` | 3 |
| `stale_assumption` | 5 |
| `duplicated_authority` | 4 |
| `missing_contract` | 2 |
| `implemented_unspecified` | 1 |
| `specified_unimplemented` | 2 |
| `planning_record_stale` | 2 |
| `compatibility_decision_required` | 2 |
| `insufficient_evidence` | 0 |

| severity | count |
|---|---:|
| major | 17 |
| minor | 4 |
| none | 5 |

Severity calibration for `duplicated_authority`: `major` means DRMCP restates or conflicts with a PRODUCT-owned normative rule in a way that may drive parser, validator, authoring, or lifecycle behavior; `minor` means DRMCP duplicates context or response-level framing without materially changing semantic behavior.

### Highest-severity contradictions

1. Spec metadata source: DRMCP contracts and implementation still treat specs as YAML-front-matter records, while PRODUCT current spec format makes H1-adjacent visible metadata and path-derived `spec:` refs canonical for new/migrated specs.
2. Phase/MVP boundary: `mvp-scope.md` says P0 tools are read-only and write tools are excluded; `tools/overview.md` makes authoring transaction tools P0 and allows writes through `accept_proposed_write`.
3. Spec update behavior: `propose_record_update` defines spec metadata updates as YAML-front-matter edits, which conflicts with the PRODUCT visible metadata contract unless retained only as a compatibility path.

## File disposition summary

| file | role | overall result | required action class | primary finding IDs |
|---|---|---|---|---|
| `overview.md` | root overview and record/tool boundary | Not safe as written; stale spec metadata source and duplicated scope statements; its own Topics table shape is current. | rewrite | F01, F13, F19, F26 |
| `responsibility-boundary.md` | boundary against brewprint MCP and filesystem tools | Boundary mostly useful but repeats stale spec YAML source. | narrow | F01, F19 |
| `namespace-scanning.md` | namespace prefix and multi-root scan | Multi-root direction valid; legacy `namespace_prefix` derivation needs alignment to PRODUCT namespace model and compatibility rules. | rewrite | F06 |
| `resolver.md` | resolver responsibility split | Good owner split, but depends on stale resolver input model in tool contract. | narrow | F02, F09 |
| `mvp-scope.md` | historical MVP tool scope | Internally superseded or phase-conflicted by authoring transaction specs. | supersede | F03 |
| `schema/overview.md` | schema area entry point | Coherent navigation role. | retain | F26 |
| `schema/authoring-guidance-source.md` | guide source discovery model | Stale source contract: declares `docs/guides/*.md`, while implementation/tests use `records/guides/*.md` and current standards are PRODUCT spec records. | rewrite | F20 |
| `schema/authoring-transaction-schema.md` | proposal/body-cache/selector schema | Transaction mechanics are useful; spec metadata and phase language need narrowing. | narrow | F05, F15, F17 |
| `schema/diagnostics.md` | diagnostic categories and required-section diagnostics | Useful DRMCP response vocabulary, but duplicates PRODUCT-owned invalidity and section semantics. | narrow | F10, F18 |
| `schema/discovery.md` | discovery contract | Stale for spec discovery through YAML front matter. | rewrite | F01, F13 |
| `schema/fields.md` | normalized record fields | Useful field model, but spec fields and workflow grammar duplicate/contradict PRODUCT. | rewrite | F01, F11, F12 |
| `schema/id-normalization.md` | public ID and bare ID normalization | Needs PRODUCT pointer split and namespace-aware compatibility clarification. | narrow | F06, F11 |
| `schema/metadata-grammar.md` | bullet metadata parser grammar | Parser grammar is DRMCP-owned; semantic requiredness must point to PRODUCT. | narrow | F11 |
| `schema/record-model.md` | internal record model and headings | Useful model, but spec bootstrap/source assumptions stale. | rewrite | F01, F13 |
| `schema/record-source.md` | metadata source per kind | Stale for spec source; otherwise useful for ADR/investigation/workflow bullet metadata parsing. | rewrite | F01 |
| `tools/overview.md` | tool set and shared response conventions | Useful navigation and tool list, but phase classification contradicts `mvp-scope.md` and duplicates schemas; the baseline ADR should decide whether it becomes navigation-only, is narrowed, or retains shared normative content. | decision_required | F03, F19, F26 |
| `tools/authoring-transaction-model.md` | shared authoring transaction behavior | Valid DRMCP-owned transaction model. | retain | F15 |
| `tools/list-records.md` | list query contract | Mostly useful; namespace filters and range compatibility need correction. | narrow | F06, F07 |
| `tools/get-record.md` | single retrieval contract | Generally valid after record model correction. | narrow | F01 |
| `tools/get-records.md` | exact batch retrieval contract | Valid exact lookup behavior. | retain | F08 |
| `tools/list-authoring-guides.md` | guide catalog tool | Stale until the source relationship between legacy guide files and PRODUCT authoring-standard specs is defined. | rewrite | F20 |
| `tools/get-authoring-guidance.md` | guide retrieval tool | Stale until retrieval target and response authority are defined for legacy guides versus PRODUCT spec records. | rewrite | F20 |
| `tools/resolve-reference.md` | resolver tool | Stale for `spec:` lookup source and section-level refs. | rewrite | F02, F09 |
| `tools/validate-records.md` | validation tool | Useful response contract, but namespace scope and PRODUCT/DRMCP invalidity split need cleanup. | narrow | F07, F10, F18 |
| `tools/suggest-next-record.md` | legacy next ADR suggestion | Obsolete or historical once `new` placeholders are baseline. | supersede | F14 |
| `tools/propose-record-create.md` | create proposal tool | Current implementation contract conflicts with captured redesigned support matrix. | split | F04, F24 |
| `tools/propose-record-update.md` | update proposal tool | Useful operation model but spec metadata update path is stale. | rewrite | F05, F17 |
| `tools/get-proposed-write.md` | retained proposal retrieval | Valid transaction support. | retain | F15 |
| `tools/accept-proposed-write.md` | proposal accept/write tool | Valid write boundary. | retain | F16 |
| `tools/discard-proposed-write.md` | proposal discard tool | Valid transaction support. | retain | F15 |

## Artifact-kind support matrix

Legend: `contracted` means current active DRMCP specs claim support; `required` means captured requirement/PRODUCT authoring standard requires it; `implemented` means code evidence exists; `missing` means no sufficient current contract or implementation evidence.

| kind | operation | current spec contract | captured requirement | current implementation | result | finding IDs |
|---|---|---|---|---|---|---|
| decision | discovery/indexing | contracted | required for existing ADRs | implemented | align with legacy compatibility | F06 |
| decision | parsing | contracted from H1 and bullet metadata | required | implemented | retain with PRODUCT pointers | F11 |
| decision | retrieval | contracted | required | implemented | retain | F08 |
| decision | validation | contracted | required | implemented | narrow duplicated semantics | F10, F18 |
| decision | reference resolution | contracted | required | implemented | retain with namespace clarification | F09 |
| decision | create proposal | contracted | required | implemented | retain after phase decision | F03 |
| decision | metadata update | contracted | required | implemented | retain with PRODUCT pointer | F11 |
| decision | named-section update | contracted | required where authoring standards define sections | implemented | narrow selector contract | F17 |
| decision | persisted-state validation | contracted | required | implemented | retain/narrow | F10 |
| spec | discovery/indexing | contracted through YAML `design_record` | required through path-derived visible `spec:` identity | implemented for YAML only | contradiction | F01, F13 |
| spec | parsing | contracted through YAML plus H1 title | required through H1-adjacent metadata and path-derived ref | implemented for YAML only | contradiction | F01, F12 |
| spec | retrieval | contracted | required | implemented only for indexed YAML specs | compatibility decision required | F13 |
| spec | validation | contracted for YAML fields/status | required for current format warning/error policy | partially implemented for old format | specified/current-format gap | F12, F21 |
| spec | reference resolution | contracted through front matter semantic refs and `SPEC-*` | required through path-derived document `spec:` refs | legacy behavior present | stale | F02 |
| spec | create proposal | outside MVP | required by redesigned contract without numeric `new` | not implemented | specified unimplemented | F04, F24 |
| spec | metadata update | contracted as YAML edit | required as visible metadata update or compatibility-only mode | implemented as YAML edit | contradiction | F05 |
| spec | named-section update | contracted | required for section updates | implemented generically | partial align | F17 |
| spec | persisted-state validation | contracted for old status/source | required for current format | missing for new/migrated format | missing contract | F21 |
| investigation | discovery/indexing | contracted | required | implemented | mostly align | F11 |
| investigation | parsing | contracted from H1 and bullet metadata | required; no metadata `id` | implemented | retain with PRODUCT pointer | F11 |
| investigation | retrieval | contracted | required | implemented | retain | F08 |
| investigation | validation | contracted | required | implemented for metadata refs | narrow task-ref and canonical-ref semantics | F10 |
| investigation | reference resolution | contracted as `INV-*` record ID | required | implemented | retain with namespace clarification | F09 |
| investigation | create proposal | outside MVP | required by redesigned contract | not implemented | specified unimplemented | F04, F24 |
| investigation | metadata update | not supported by update contract | required for complete authoring support | not implemented | missing contract | F24 |
| investigation | named-section update | not supported by update contract | required for complete authoring support | not implemented | missing contract | F24 |
| investigation | persisted-state validation | partially contracted | required by authoring standard | partial | narrow | F10 |
| requirement | discovery/indexing | contracted | required | implemented | retain with namespace cleanup | F06 |
| requirement | parsing | contracted | required | implemented | narrow duplicated semantics | F11 |
| requirement | retrieval | contracted | required | implemented | retain | F08 |
| requirement | validation | contracted | required | implemented | narrow semantics vs diagnostics | F18 |
| requirement | reference resolution | contracted | required | implemented | retain with namespace cleanup | F09 |
| requirement | create proposal | contracted | required | implemented | phase decision required | F03 |
| requirement | metadata update | contracted | required | implemented | retain with PRODUCT pointer | F11 |
| requirement | named-section update | contracted | required | implemented | narrow selector contract | F17 |
| requirement | persisted-state validation | contracted | required | implemented | retain/narrow | F18 |
| work item | discovery/indexing | contracted | required | implemented | retain with namespace cleanup | F06 |
| work item | parsing | contracted | required | implemented | narrow duplicated semantics | F11 |
| work item | retrieval | contracted | required | implemented | retain | F08 |
| work item | validation | contracted | required | implemented | narrow semantics vs diagnostics | F18 |
| work item | reference resolution | contracted | required | implemented | retain with namespace cleanup | F09 |
| work item | create proposal | contracted | required | implemented with reciprocal update option | phase decision required | F03 |
| work item | metadata update | contracted | required | implemented | retain with PRODUCT pointer | F11 |
| work item | named-section update | contracted | required | implemented | narrow selector contract | F17 |
| work item | persisted-state validation | contracted | required | implemented | retain/narrow | F18 |
| task | discovery/indexing | contracted | required | implemented | retain with namespace cleanup | F06 |
| task | parsing | contracted | required | implemented | narrow duplicated semantics | F11 |
| task | retrieval | contracted | required | implemented | retain | F08 |
| task | validation | contracted | required | implemented | narrow semantics vs diagnostics | F18 |
| task | reference resolution | contracted | required | implemented | retain with namespace cleanup | F09 |
| task | create proposal | contracted | required, sequence scoped by parent work item | implemented | retain with phase decision | F03, F24 |
| task | metadata update | contracted | required | implemented | retain with PRODUCT pointer | F11 |
| task | named-section update | contracted | required | implemented | narrow selector contract | F17 |
| task | persisted-state validation | contracted | required | implemented | retain/narrow | F18 |

## Planning record disposition

| record | disposition | reason | recommended next handling |
|---|---|---|---|
| `DRMCP-REQ-MCP-001` | revise | Still covers real multi-root query, resolver, identity, and validation gaps, but references older unsplit contracts and does not include the post-PRODUCT ownership baseline. | Update after baseline ADR to point at current split specs and resolve namespace filter/default-scope decisions. |
| `DRMCP-REQ-MCP-002` | revise | Direction is mostly aligned with PRODUCT authoring standards, but current implementation and current specs do not meet its full support matrix, especially spec/investigation create and spec metadata shape. | Keep as requirement input, then narrow into phase-specific acceptance criteria. |
| `DRMCP-INV-MCP-001` | conclude as historical | Its findings remain useful evidence for multi-root gaps, but it references old `tools.md`/`schema.md` locations and predates the current PRODUCT semantic boundary. | Cite as historical input; do not use as current contract. |
| `DRMCP-WORK-SPEC-001` | supersede or absorb into new Work Item | Per-file parser-aware validation remains needed, but current dependencies and spec-source assumptions need realignment to H1-adjacent spec metadata and compatibility policy; `spec:product.design_records.spec_format.validation_policy` currently names this Work Item as the owner of per-file validation work. | Absorb into a coordinated DRMCP Work Item only with a matching PRODUCT validation-policy owner-pointer update. |
| `DRMCP-WORK-SPEC-002` | supersede or absorb into new Work Item | Graph validation remains needed, but the Work Item uses stale Topics columns (`parent`/`file`) instead of current `title/kind/ref/summary`; `spec:product.design_records.spec_format.validation_policy` currently names this Work Item as the owner of Topics-graph validation work. | Absorb into a coordinated DRMCP Work Item only with a matching PRODUCT validation-policy owner-pointer update. |

## Implementation drift summary

| area | contract state | implementation state | test evidence | disposition |
|---|---|---|---|---|
| Spec discovery/parsing | Current DRMCP specs say YAML `design_record`; PRODUCT current format says H1-adjacent visible metadata and path-derived `spec:` refs. | Parser and authoring code still require/modify YAML front matter for spec metadata. | Tests include YAML spec fixtures and pass. | Decide compatibility policy, then rewrite specs and tests. |
| Authoring transaction tools | Tool specs define proposal lifecycle and accept-only write path; phase classification conflicts with read-only MVP. | All authoring tools are exposed and implemented. | `go test ./drmcp/src/...` passed. | Keep mechanics, decide phase model by ADR. |
| Workflow create/update | Current specs and code support requirement/work item/task create/update with reciprocal updates. | Implemented, including parent-scoped task `new`. | Authoring and tools-call tests pass. | Retain after namespace/PRODUCT pointer cleanup. |
| Spec/investigation create | Requirement demands future support; current tool spec excludes it. | Not implemented; create validator rejects spec/investigation. | Tests cover unsupported create kinds. | Add to future implementation phase only after contract correction. |
| Multi-root namespace lookup | Specs partially define multi-root; namespace filters/defaults incomplete. | Config/index support multiple records roots; resolver can use known prefixes; tests cover both `V01-*` legacy forms and unprefixed bare fixture IDs. | Tests pass but encode mixed legacy assumptions. | Treat `V01-*` compatibility and unprefixed bare fixture IDs as separate decisions; the latter must not inherit accepted V01 compatibility by implication. |
| Resolver and validation | Specs mix PRODUCT semantic invalidity and DRMCP diagnostics; `spec:` lookup source stale. | Resolver and validation implemented for current legacy model. | Resolve and validation tests pass. | Rewrite resolver spec and keep diagnostics as representation layer. |
| Topics graph validation | PRODUCT current Topics shape is `title | kind | ref | summary`. | No sufficient evidence of a current graph validator for this shape. | No relevant passing test identified for current PRODUCT Topics graph. | Defer to superseding Work Item. |
| Authoring guidance | Schema declares `docs/guides/*.md`; tool contracts expose read-only guidance retrieval; current PRODUCT authoring standards are spec records, not guide files. | Implementation reads `<primaryRecordsRoot>/guides/*.md`. | Tests pass using `v01/records/guides/*.md`, which proves legacy guide retrieval but not alignment with current PRODUCT authoring-standard specs. | Decide whether guide tools are legacy compatibility, PRODUCT spec projections, or both; then rewrite source contracts. |

Implementation command run:

```text
go test ./drmcp/src/...
```

Exit code: `0`.

Output:

```text
ok  	github.com/hiroshiasayadev-prog/brewprint/drmcp/src/cmd/design-records-mcp	7.676s
ok  	github.com/hiroshiasayadev-prog/brewprint/drmcp/src/internal/designrecords	7.239s
ok  	github.com/hiroshiasayadev-prog/brewprint/drmcp/src/internal/designrecordsmcp	1.018s
```

## Verification

Final verification commands:

```text
git status --short
git diff --cached --name-only
git status --short -- v01 drmcp/records/investigations/mcp/DRMCP-INV-MCP-002-design-records-mcp-contract-consistency-and-realignment-audit.md drmcp/records/spec drmcp/records/requirements drmcp/records/work-items drmcp/records/tasks drmcp/src product/records/spec product/records/requirements product/records/work-items product/records/tasks CLAUDE.md prompt_chappy.md
```

Final `git status --short` output:

```text
 M CLAUDE.md
 M product/records/requirements/PRODUCT-REQ-SPEC-001-mcp-readable-spec-format-and-topic-tree.md
 M product/records/tasks/spec/PRODUCT-TASK-SPEC-009-04-opus-review.md
 M product/records/work-items/spec/PRODUCT-WORK-SPEC-009-format-only-migration-traceability-and-artifact-model-specs.md
 M prompt_chappy.md
?? drmcp/records/investigations/mcp/DRMCP-INV-MCP-002-design-records-mcp-contract-consistency-and-realignment-audit.md
```

Verification results:

- Only the allowed investigation file was written during this task.
- The investigation file remains untracked, as it was at baseline.
- No staged changes were created.
- No commit was created.
- `v01/` is unchanged in final scoped status.
- No DRMCP spec, requirement, work item, task, implementation, test, or fixture file was modified.
- No PRODUCT spec file was modified.
- The pre-existing modified `CLAUDE.md`, `prompt_chappy.md`, PRODUCT requirement, PRODUCT task, and PRODUCT work item files were not touched by this task.
- Investigation authoring shape check: exactly one H1, no metadata `id` field, all canonical H2 sections present, status concluded, substantive scope and findings.
- Trailing whitespace check: no trailing whitespace lines found.

## Cross-cutting observations

- The current DRMCP contract set is not safe to implement as written. Passing implementation tests show code consistency with the old contract, not alignment with the accepted PRODUCT semantic boundary.
- The largest risk is treating implementation evidence as semantic authority. PRODUCT owns the spec format, traceability, ID grammar, authoring standards, and artifact responsibilities; DRMCP must implement and diagnose those rules.
- A compatibility decision is required before any broad parser rewrite. The audit found three distinct states that must not be conflated: canonical new/migrated spec behavior, legacy Brewprint compatibility behavior, and obsolete front-matter semantics that should not remain normative.
- The read-only MVP contract and authoring transaction MVP contract cannot both be current under the same phase name.
- Current overview files are useful as navigation entry points, but normative detail is repeated across overview, schema, and per-tool files. Later corrections should leave navigation overviews thinner and make individual contracts or schema references authoritative.
- `DRMCP-REQ-MCP-001` and `DRMCP-REQ-MCP-002` are complementary. They should feed one coordinated baseline rather than compete as separate implementation starts.
- Superseding or absorbing `DRMCP-WORK-SPEC-001` or `DRMCP-WORK-SPEC-002` is not local to DRMCP planning records because the PRODUCT validation policy currently points to those records as validation owners.

## Follow-up judgment candidates

- Decide the current DRMCP phase model: retain "read-only MVP" as a historical phase, rename authoring transaction work as a later phase, or supersede the old MVP label.
- Decide spec compatibility policy: whether YAML-front-matter specs remain indexed as compatibility-only records, whether they emit warnings, and whether new/migrated specs must be indexed solely through path-derived H1-adjacent metadata.
- Decide whether `suggest_next_record` is retained as legacy decision-only convenience, deprecated, or removed after `new` placeholder support is baseline.
- Decide whether authoring guidance tools should retrieve legacy `records/guides/*.md`, current PRODUCT authoring-standard spec records, or an explicit compatibility projection between the two.
- Decide whether spec and investigation create support belongs in the same phase as workflow artifact create/update or in a later phase after parser realignment.
- Decide whether `DRMCP-REQ-MCP-001` and `DRMCP-REQ-MCP-002` should remain separate requirements under one coordinating Work Item or be superseded by more phase-specific requirements.
- Decide how the PRODUCT validation-policy owner pointers are updated if `DRMCP-WORK-SPEC-001` or `DRMCP-WORK-SPEC-002` is superseded or absorbed.

## Recommendation

Do not implement against the current DRMCP contract set as written.

Recommended dependency order:

1. Create one baseline DRMCP ADR that accepts the contract authority tree, phase model, and compatibility policy.
2. Revise `DRMCP-REQ-MCP-001` and `DRMCP-REQ-MCP-002` to match the ADR and current PRODUCT semantic sources.
3. Coordinate any supersession or absorption of `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002` with an update to the PRODUCT validation-policy owner pointers.
4. Run DRMCP contract-spec correction batches: spec source/parsing, query/resolve/validation, authoring transaction support, authoring-guidance source, and tool hierarchy cleanup.
5. Add or refresh compatibility fixtures for legacy YAML-front-matter specs, legacy flat ADRs, `V01-*` legacy IDs, unprefixed bare fixture IDs only if separately accepted, and current H1-adjacent spec records.
6. Create implementation work only after corrected contracts and fixtures are reviewable.
7. Run an independent review before conclusion artifacts or broad implementation changes.

## Follow-up artifact candidates

| candidate | purpose |
|---|---|
| `DRMCP-ADR-MCP-new` | Decide the DRMCP contract baseline, phase naming, authority tree, compatibility policy, and implementation sequence. |
| `DRMCP-WORK-MCP-new` | Coordinate requirement updates, spec correction batches, compatibility fixtures, implementation work, and independent review. |
| `DRMCP-REQ-MCP-001` update or successor | Resolve namespace filter/default-scope, exact lookup, known namespace, cross-namespace validation, and duplicate identity contracts. |
| `DRMCP-REQ-MCP-002` update or successor | Resolve namespace-aware authoring support by artifact kind, exact ID/new behavior, generated placement, and reciprocal update contracts. |
| `DRMCP-WORK-SPEC-001` disposition | Supersede or absorb parser-aware spec-format validation into the coordinated Work Item, coordinated with the PRODUCT validation-policy owner pointer. |
| `DRMCP-WORK-SPEC-002` disposition | Supersede or absorb Topics graph validation into the coordinated Work Item using the current Topics table contract, coordinated with the PRODUCT validation-policy owner pointer. |
| DRMCP compatibility fixture task | Create fixtures that distinguish canonical current spec format from retained legacy compatibility inputs. |
| DRMCP authoring-guidance source task | Resolve whether guidance tools expose legacy guide files, PRODUCT authoring-standard specs, or a compatibility projection. |
| DRMCP implementation task set | Implement only after the ADR and corrected specs are accepted. |
| Independent review task | Review corrected contracts and implementation plan before execution. |

## Open questions

- Is YAML-front-matter spec parsing retained as a compatibility mode, or is it removed after migration?
- If retained, what diagnostics should legacy YAML-front-matter specs receive during inventory, validation, proposal-local validation, and repository-wide validation?
- Is the old read-only MVP phase historical/superseded, or should authoring transaction tools be assigned to a distinct later phase?
- Should spec creation and investigation creation be delivered with workflow artifact authoring, or separated into a later phase?
- Should `suggest_next_record` be deprecated once `new` placeholders and namespace-aware create are accepted?
- Which legacy ID forms remain compatibility inputs but non-canonical outputs, and are unprefixed bare fixture IDs such as `REQ-MCP-001`, `WORK-MCP-001`, and `TASK-MCP-001-01` accepted or rejected separately from `V01-*` compatibility?
- Are legacy `records/guides/*.md` guide files still an MCP surface, or should authoring guidance resolve to PRODUCT authoring-standard spec records?
- Should overview records become navigation-only, be narrowed, or retain shared normative content where per-tool and per-schema contracts already define behavior?
