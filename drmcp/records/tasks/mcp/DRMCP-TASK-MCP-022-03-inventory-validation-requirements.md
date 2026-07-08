# DRMCP-TASK-MCP-022-03: Inventory validation requirements

- **id**: DRMCP-TASK-MCP-022-03
- **status**: done
- **date**: 2026-07-07
- **work_item**: DRMCP-WORK-MCP-022
- **task_type**: investigation
- **estimate**: 1.0d
- **depends_on**:
  - DRMCP-TASK-MCP-022-01
- **outputs**:
  - DRMCP-TASK-MCP-022-03

## Startup / Required reading

Before starting the inventory, read these files and preserve their accepted semantics:

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/writing-standard.md`
3. `skills/design-convergence-workflow/SKILL.md`
4. `skills/design-convergence-workflow/design-authoring.md`
5. `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-022-domain-structural-provisional-contracts.md`
6. `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-022-01-decide-domain-structural-contract-baseline-from-w021.md`
7. `drmcp/records/spec/implementation/contracts/application-use-cases/contract-boundary.md`
8. `drmcp/records/spec/implementation/contracts/record-domain-logical-tree/contract-boundary.md`
9. `drmcp/records/spec/design-records-mcp/schema/record-model.md`
10. `drmcp/records/spec/design-records-mcp/schema/record-source.md`

Then perform a bounded validation-requirement search under these areas only unless the Task records an explicit reason to expand scope:

- `product/records/spec/design-records/`
- `drmcp/records/spec/design-records-mcp/`
- `drmcp/records/spec/implementation/contracts/`
- `drmcp/records/spec/application-architecture/`

Do not rely on remembered validation rules. Do not assign response shapes in this Task.

## Goal

Inventory product and DRMCP validation requirements before deciding validation use-case contracts.

## Investigation questions

- What validation rules exist today?
- Which rules are local record rules, relation rules, graph rules, workflow rules, source/parse rules, or operation-scope rules?
- Which rules require stable validation requirement IDs before downstream decisions?
- Which rules are Domain rule-violation judgments versus Application diagnostic presentation choices?
- Which validation rules are missing or ambiguous in existing specs?

## Expected output

Create a table with at least:

| candidate validation ID | source spec | rule summary | rule class | Domain judgment? | Application presentation? | certainty | notes |
|---|---|---|---|---|---|---|---|

Stable IDs may be provisional but must be deterministic and documented.

## Investigation scope

Bounded search stayed inside the requested scope:

- `product/records/spec/design-records/`
- `drmcp/records/spec/design-records-mcp/`
- `drmcp/records/spec/implementation/contracts/`
- `drmcp/records/spec/application-architecture/`

No scope expansion was needed.

The operational DRMCP authoring tool surface was not available in this session. Per `prompt_chappy.md`, filesystem fallback was used after attempting tool discovery. This Task updates only this investigation record.

## Provisional stable ID scheme

Use `VR-<family>-<NNN>`.

| family | meaning |
|---|---|
| `SRC` | Source discovery, parsing, metadata grammar, source retention, or source-location rule. |
| `ID` | Canonical identity, ID grammar, ID agreement, or duplicate identity rule. |
| `LOC` | Local record, metadata, lifecycle, document-shape, or section rule. |
| `REL` | Declared relation, lookup target, reciprocity, or relation-value rule. |
| `GRP` | Multi-record graph rule. |
| `WFL` | Workflow authoring or workflow-state rule. |
| `OPS` | Operation request, selection, orchestration, execution-prerequisite, or affected-set rule. |
| `PRS` | Presentation, projection, diagnostic envelope, ordering, or duplicate-suppression concern. |
| `AMB` | Ambiguous or missing rule requiring later decision. |

Numbers are assigned in bounded-source reading order within each family. The numbers are provisional. Do not reuse an ID within this Task.

## Validation requirement inventory

| candidate validation ID | source spec | rule summary | rule class | Domain judgment? | Application presentation? | certainty | notes |
|---|---|---|---|---|---|---|---|
| VR-SRC-001 | `spec:drmcp.design_records_mcp.namespace_scanning` | DRMCP configuration declares at least one current source; app namespaces and canonical current roots are unique; each entry selects one source form. | source / parse rule | No; configuration trustworthiness boundary is outside Domain judgment. | Yes; Application or lifecycle chooses startup/execution failure. | high | Domain snapshot construction depends on this state but does not own config validation. |
| VR-SRC-002 | `spec:drmcp.design_records_mcp.namespace_scanning` | Current and Legacy roots must be readable, canonical, repository-contained, unique, and non-overlapping; invalid mandatory roots are not omitted. | source / parse rule | No for root validation; yes only after sources enter Domain as typed inputs. | Yes; failure projection is Application/lifecycle concern. | high | Treat as startup or execution prerequisite, not normal record invalidity. |
| VR-SRC-003 | `spec:drmcp.design_records_mcp.schema.discovery` | Current Spec candidates are regular Markdown files under the effective Spec tree; non-Markdown and symlinked or alias-escaped sources are excluded. | source / parse rule | No; source enumeration is Application/Infrastructure-side input formation. | Yes; only if exposed as operation failure or diagnostic. | high | Candidate inclusion differs from semantic invalidity. |
| VR-SRC-004 | `spec:product.design_records.spec_format.validation_policy`; `spec:drmcp.design_records_mcp.schema.discovery` | Parser-aware heading validation ignores YAML front matter and fenced code blocks when counting real H1 headings. | source / parse rule | Yes; parse finding is Domain-owned once raw Markdown enters parser. | Yes; projection only. | high | Applies to missing/multiple H1 and H1 shape checks. |
| VR-SRC-005 | `spec:drmcp.design_records_mcp.schema.metadata_grammar`; `spec:drmcp.design_records_mcp.schema.record_source` | H1-adjacent visible metadata must immediately follow the real H1 using `- **key**: value`; body prose, later headings, and YAML front matter are not metadata sources. | source / parse rule | Yes. | Yes; projection only. | high | This is distinct from kind-specific requiredness. |
| VR-SRC-006 | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Scalar/list normalization is strict: malformed inline code, invalid child list shape, empty child items, duplicate recognized markers, and invalid scalar/list shape are invalid. | source / parse rule | Yes. | Yes; projection only. | high | Unknown marker validity remains kind-specific; see VR-AMB-002. |
| VR-SRC-007 | `spec:drmcp.design_records_mcp.schema.discovery`; `spec:drmcp.design_records_mcp.schema.record_model` | Invalid but identity-determinable sources remain addressable; identity-less sources remain validation-only; invalid/conflicting sources retain provenance for repair diagnostics. | source / parse rule | Yes for addressability semantics inside Domain snapshot. | Yes for source-location projection. | high | Important input to exact-ref validation in T04. |
| VR-SRC-008 | `spec:drmcp.design_records_mcp.schema.diagnostics` | Source-backed diagnostics require a portable location when one repairable source is the target; if required location cannot be constructed, normal validation output must not be emitted. | operation-scope rule | No; trustworthy-output failure selection is Application-owned. | Yes. | high | Include as operation prerequisite, not semantic rule violation. |
| VR-ID-001 | `spec:product.design_records.namespace_model.artifact_id_grammar`; `spec:drmcp.design_records_mcp.schema.id_normalization` | Sequential ADR, investigation, requirement, work item, and task IDs use complete app-aware public ID grammar; bare `REQ-*`, `WORK-*`, and `TASK-*` fragments are not canonical refs. | local record rule | Yes. | Yes; projection only. | high | Complete ID validation is Domain semantic judgment. |
| VR-ID-002 | `spec:drmcp.design_records_mcp.schema.id_normalization`; Product authoring standards | For sequential artifacts, H1 is canonical identity authority; metadata `id` and filename prefixes are consistency values and must agree where required. | local record rule | Yes. | Yes; projection only. | high | ADR and investigation metadata currently have no `id`. |
| VR-ID-003 | `spec:product.design_records.spec_format.spec_id_as_ref`; `spec:drmcp.design_records_mcp.schema.id_normalization` | Current Spec identity is path-derived from configured app namespace and topic path; H1-adjacent `id` must match exactly and never becomes alias or fallback identity. | local record rule | Yes. | Yes; projection only. | high | Moving or renaming a Spec changes canonical ref unless later compatibility says otherwise. |
| VR-ID-004 | `spec:product.design_records.traceability.resolve_and_validation`; `spec:drmcp.design_records_mcp.schema.discovery`; `spec:drmcp.design_records_mcp.schema.record_model` | Duplicate canonical current identity creates no winner; all conflicting sources remain validation inputs and unrelated identities remain available. | graph rule | Yes. | Yes; conflict member projection only. | high | Applies to sequential record public IDs and path-derived Spec refs. |
| VR-LOC-001 | `spec:drmcp.design_records_mcp.schema.fields`; `spec:drmcp.design_records_mcp.schema.record_model` | Valid current records supply common `id`, `kind`, `title`, `status`, and `date`; invalid records may miss fields or retain invalid parsed values; DRMCP does not synthesize replacements. | local record rule | Yes. | Yes; missing/null projection is Application-owned. | high | Compact-list missing-field warnings are presentation/usecase-specific. |
| VR-LOC-002 | `spec:product.design_records.authoring_standards.adr_authoring`; `spec:drmcp.design_records_mcp.schema.metadata_grammar` | ADR records require accepted ADR ID/H1/file shape, required metadata fields, strict date, lifecycle values, and required body sections. | local record rule | Yes. | Yes; projection only. | high | Supersedes and depends_on also create relation checks. |
| VR-LOC-003 | `spec:product.design_records.authoring_standards.investigation_authoring`; `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Investigation records require accepted ID/H1/file shape, metadata fields, lifecycle values, canonical body sections, and concluded-state substantive findings/scope. | local record rule | Yes. | Yes; projection only. | high | Relation fields are split into VR-REL-002. |
| VR-LOC-004 | `spec:product.design_records.authoring_standards.requirement_authoring`; `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Requirement records require ID consistency, status/date/source refs, accepted-state substantive `Requirement` and `Required Outcome`, and duplicate-free direct Work Item reverse relation semantics. | local record rule | Yes. | Yes; projection only. | high | Reverse relation check also appears as VR-REL-004. |
| VR-LOC-005 | `spec:product.design_records.authoring_standards.work_item_authoring`; `spec:product.design_records.traceability.metadata_schema` | Work Item records require ID consistency, status/date, non-empty `source_refs`, impact/tasks field normalization, and lifecycle-gated substantive sections. | local record rule | Yes. | Yes; projection only. | medium | Existing DRMCP metadata grammar still mentions `source_requirement`; see VR-AMB-003. |
| VR-LOC-006 | `spec:product.design_records.authoring_standards.task_authoring`; `spec:product.design_records.namespace_model.artifact_id_grammar` | Task records require task ID grammar, metadata consistency, task type vocabulary, work item ownership, dependency/output normalization, and done/cancelled gated sections. | local record rule | Yes. | Yes; projection only. | medium | Existing DRMCP metadata grammar still mentions `source_requirement`; see VR-AMB-003. |
| VR-LOC-007 | `spec:product.design_records.spec_format.document_shape`; `spec:product.design_records.spec_format.validation_policy` | Current Specs require accepted Spec kind, exactly one real H1, visible metadata, required sections by kind, and `contract_class` only for Contract specs. | local record rule | Yes. | Yes; severity/projection only. | high | Migration phase affects severity, not whether the condition exists. |
| VR-LOC-008 | `spec:product.design_records.spec_format.topics_table` | Index or Overview specs with topics require a `## Topics` table with `title`, `kind`, `ref`, and `summary`; row title/kind/ref/summary must satisfy the table contract. | local record rule | Yes for table shape. | Yes; projection only. | high | Child ref resolution and parent graph checks are VR-GRP-002. |
| VR-LOC-009 | `spec:product.design_records.spec_format.validation_policy`; `spec:drmcp.design_records_mcp.schema.metadata_grammar` | YAML front matter and obsolete hidden fields are invalid for new/migrated current Specs; YAML never supplies fallback metadata. | source / parse rule | Yes. | Yes; migration severity/projection only. | high | Existing/unmigrated warning behavior is a severity mapping input. |
| VR-REL-001 | `spec:product.design_records.traceability.artifact_refs`; `spec:product.design_records.traceability.semantic_ref` | Active relation values must use supported canonical ref classes; physical paths, bare grammar fragments, inactive section refs, and unsupported semantic prefixes are noncanonical. | relation rule | Yes. | Yes; projection only. | high | Unsupported relation value should be invalid value, not lookup-unavailable. |
| VR-REL-002 | `spec:product.design_records.traceability.metadata_schema`; `spec:product.design_records.traceability.resolve_and_validation` | Investigation `source_refs` and `follow_up_results` must use supported canonical refs and resolve; `follow_up_candidates` must use canonical form when written but may remain unresolved if not yet created. | relation rule | Yes. | Yes; projection only. | high | Task IDs are not supported in investigation canonical-reference fields. |
| VR-REL-003 | `spec:product.design_records.traceability.metadata_schema`; `spec:product.design_records.traceability.resolve_and_validation` | Work Item `source_refs` must be non-empty, active canonical refs, duplicate-free, not self-referential, and resolvable. | relation rule | Yes. | Yes; projection only. | high | Ordering has no semantic meaning. |
| VR-REL-004 | `spec:product.design_records.traceability.metadata_schema`; `spec:product.design_records.authoring_standards.requirement_authoring` | Requirement direct reverse Work Item relation is derived from Work Item `source_refs`; mismatch blocks migration and must not be silently repaired. | relation rule | Yes. | Yes; projection only. | high | Derived set is unordered and duplicate-free. |
| VR-REL-005 | `spec:product.design_records.traceability.resolve_and_validation`; `spec:drmcp.design_records_mcp.tools.validate_records` | Work Item `tasks` and Task `work_item` must agree; Task `depends_on` must point to existing Task targets. | relation rule | Yes. | Yes; projection only. | high | Exact target lookup uses current snapshot and optional Legacy state where accepted. |
| VR-REL-006 | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.design_records_mcp.resolver` | Current relation lookup uses exact current canonical identity across all configured current roots; same-app restriction is prohibited; duplicate conflict does not satisfy target existence. | relation rule | Yes. | Yes; projection only. | high | Public `resolve_reference` is not called by validation. |
| VR-REL-007 | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.design_records_mcp.namespace_scanning` | Accepted legacy relation targets are lookup targets only; disabled fallback, unresolved target, duplicate legacy conflict, unreadable source, and unsupported value remain distinct lookup outcomes. | relation rule | Yes for relation invalidity; lookup state classification feeds Domain outcome. | Yes; projection only. | high | Legacy archive records are never repository-validation subjects. |
| VR-GRP-001 | `spec:product.design_records.traceability.resolve_and_validation` | Semantic provenance cycle is invalid; a Work Item source ref to a Task normalizes to the Task owner Work Item for provenance-cycle analysis. | graph rule | Yes. | Yes; projection only. | high | Algorithm remains DRMCP-local, but invalid state is PRODUCT-owned. |
| VR-GRP-002 | `spec:product.design_records.spec_format.topics_table`; `spec:product.design_records.spec_format.validation_policy` | Topic child refs must resolve, child parent marker must agree with exactly one authoritative parent declaration, duplicate parent declarations are invalid. | graph rule | Yes. | Yes; projection only. | high | Cross-record Spec tree rule. |
| VR-GRP-003 | `spec:product.design_records.spec_format.validation_policy` | Topic cycle is identified but deferred until graph validation contract. | ambiguous / needs decision | Yes eventually, but exact rule is not fully specified. | Yes later. | medium | Keep as missing graph-rule detail before implementation-ready design. |
| VR-WFL-001 | `spec:product.design_records.authoring_standards.task_authoring`; `skills/design-convergence-workflow/design-authoring.md` | Task type controls required sections and prohibits mixing investigation, decision, authoring, implementation, review, coordination, and verification responsibilities in one task. | workflow rule | Yes when validating task record semantics. | Yes; projection only. | high | Workflow semantic rule, not Go implementation planning. |
| VR-WFL-002 | Product authoring standards for requirement/work item/task/investigation | Lifecycle-gated states require substantive sections; `TBD` or empty content is invalid where the state requires completed evidence. | workflow rule | Yes. | Yes; projection only. | high | Applies differently per record kind and lifecycle state. |
| VR-WFL-003 | `spec:product.design_records.authoring_standards.task_authoring` | `work_item_ref` is required only for `work_item_execution` tasks and prohibited otherwise; it must target exactly one existing Work Item under the allowed namespace/domain constraints. | workflow rule | Yes. | Yes; projection only. | high | This is separate from normal `work_item` ownership metadata. |
| VR-WFL-004 | `spec:product.design_records.traceability.metadata_schema` | Migration must atomically replace old workflow fields, must not persist both old and new fields where prohibited, and must not silently repair mismatched derived relations. | workflow rule | Yes. | Yes; projection only. | medium | Exact current migration phase may affect severity. |
| VR-OPS-001 | `spec:drmcp.design_records_mcp.tools.validate_records` | `validate_records` request accepts empty repository-wide scope, one `app_namespace`, or one exact current canonical `ref`; `app_namespace` and `ref` are mutually exclusive; physical paths, legacy IDs, filters, and unknown fields are unsupported. | operation-scope rule | No; request validity and selector policy are Application-owned. | Yes. | high | T04 response mapping should separate request error from validation diagnostic. |
| VR-OPS-002 | `spec:drmcp.design_records_mcp.tools.validate_records` | Validation subjects are selected by request scope: all retained sources/records/conflict groups, one app subset, or one exact current source/record/conflict group; relation targets are lookup-only and not recursively added. | operation-scope rule | No for subject selection; yes for findings emitted by Domain validators. | Yes. | high | Domain should not own public selector policy. |
| VR-OPS-003 | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.application_architecture.runtime_and_state` | Validation builds one fresh immutable snapshot; mandatory current roots, mandatory configured legacy roots, active index, legacy lookup state, and required locations must be trustworthy or no normal validation response is returned. | operation-scope rule | No; failure selection is Application-owned. | Yes. | high | Domain owns semantic snapshot content after assembly. |
| VR-OPS-004 | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary` | Validation pass order is source/local first, current relation, legacy relation, Topics graph, then aggregation/projection; validators perform no I/O, no rescans, no public MCP tool calls, and no response formatting. | operation-scope rule | Partial: Domain owns validator findings; Application owns orchestration/projection. | Yes. | high | T01 refines ownership toward Domain judgment and Application presentation. |
| VR-OPS-005 | `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema`; `spec:drmcp.design_records_mcp.tools.authoring_transaction_model` | Proposal-local validation is limited to the proposed target plus related records modified by the same proposal; unrelated repository diagnostics must not affect proposal-local `validation.ok`. | operation-scope rule | No for affected-set selection; yes for rule findings inside the candidate set. | Yes. | high | This is the closest authoring behavior to "run validate_records on a narrowed candidate affected set". |
| VR-OPS-006 | `spec:drmcp.design_records_mcp.tools.list_records` | `list_records` requires one configured `app_namespace`, supported sequential `kind`, canonical `domain`, optional exact status, accepted order, and limit; invalid scope or unsupported fields reject the request without broadening. | operation-scope rule | No; public query validity and broadening policy are Application-owned. | Yes. | high | Read/query operation rule, not repository-validation finding. |
| VR-OPS-007 | `spec:drmcp.design_records_mcp.tools.list_records`; `spec:drmcp.design_records_mcp.schema.diagnostics` | Addressable invalid records remain listable; missing compact `title` / `status` / `date` projects `null` and triggers operation warnings; duplicate-conflict identities have no list winner. | operation-scope / presentation rule | Domain owns addressability and duplicate-conflict state; Application owns compact projection and warnings. | Yes. | high | Important because this warning behavior is not `validate_records` output. |
| VR-OPS-008 | `spec:drmcp.design_records_mcp.tools.get_records` | `get_records` validates request shape, classifies exact string inputs once, performs ordered exact deduplication, and treats malformed/unsupported/unresolved/unavailable/duplicate item outcomes as normal partial-success warnings when locations are constructible. | operation-scope rule | Domain owns lookup states and conflicts; Application owns exact-retrieval classification order and partial-success projection. | Yes. | high | This is read-usecase response policy, not `validate_records(ref)`. |
| VR-OPS-009 | `spec:drmcp.design_records_mcp.tools.resolve_reference`; `spec:drmcp.design_records_mcp.resolver` | `resolve_reference` accepts one exact string, runs current-first then accepted legacy fallback, and maps semantic outcomes to only `resolved`, `unresolved`, or `unsupported` plus cause diagnostics. | operation-scope rule | Domain owns resolution states; Application owns public status mapping. | Yes. | high | Validation must not call the public resolver, but T04 may need consistent mapping vocabulary. |
| VR-OPS-010 | `spec:drmcp.design_records_mcp.schema.authoring_guidance_source`; `spec:drmcp.design_records_mcp.tools.list_authoring_guides`; `spec:drmcp.design_records_mcp.tools.get_authoring_guidance` | Guidance operations use fixed `design_records` current Spec scope, exact canonical refs, no path/title/fuzzy lookup, and fail rather than omit unprojectable in-scope guides. | operation-scope rule | Domain owns normal Current Records identity/addressability; Application owns fixed Guidance scope and projection availability. | Yes. | high | Guidance is read-like, not a separate parser/index/model. |
| VR-OPS-011 | `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema`; `spec:drmcp.design_records_mcp.tools.authoring_transaction_model`; `spec:drmcp.design_records_mcp.tools.propose_record_create` | Authoring create request validation covers supported kind, exact/new ID forms, domain and parent consistency, required `fields`, section-only body/body-cache exclusivity, `fields.id` consistency, and required reciprocal update mode behavior. | operation-scope rule | No for request validity and target resolution policy; yes for proposal-local record-rule findings after candidate state exists. | Yes. | high | `fields.id` mismatch and placeholder-time `fields.id` are invalid requests, not record diagnostics. |
| VR-OPS-012 | `spec:drmcp.design_records_mcp.tools.propose_record_update`; `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | Authoring update validates exact existing ID, one of `update` or `operations`, body-source placement, metadata operation constraints, exact section selector resolution, conflict detection, final-state validation, and no-op behavior. | operation-scope rule | No for request/selector/proposal creation policy; yes for final candidate record findings. | Yes. | high | Section selector failures create no proposal and are not repository-validation findings. |
| VR-OPS-013 | `spec:drmcp.design_records_mcp.tools.accept_proposed_write`; `spec:drmcp.design_records_mcp.tools.authoring_transaction_model` | Accept validates proposal lifecycle, target staleness, ID collision, target continuity, required follow-up satisfaction, pre-write error diagnostics, affected paths, and source-backed diagnostic locations before writing. | operation-scope rule | No for write eligibility; yes for pre/post validation findings when scoped to affected records. | Yes. | high | `written:false` must not be conflated with `validation.ok:false`; post-write failure after actual write still reports `written:true`. |
| VR-PRS-001 | `spec:drmcp.design_records_mcp.schema.diagnostics` | Diagnostic envelope, category vocabulary, severity vocabulary, association shape, location exposure, deterministic ordering, and duplicate suppression are shared projection concerns. | presentation-only concern | No; representation does not decide violation. | Yes. | high | T04 must decide response mapping without redefining semantic rules. |
| VR-PRS-002 | `spec:drmcp.design_records_mcp.schema.diagnostics`; `spec:drmcp.design_records_mcp.tools.validate_records` | `validate_records.ok` is derived from returned diagnostic severity; warnings/info do not make `ok` false. | presentation-only concern | No. | Yes. | medium | Existing contract says this, but T04 is the explicit response-mapping task. |
| VR-AMB-001 | `spec:product.design_records.spec_format.validation_policy` | Migration-sensitive severity is specified for several Spec rules, but the stable mapping between migration phase, Domain finding, Application severity, and T04 response fields is not yet decided. | ambiguous / needs decision | Domain decides violation; severity/presentation needs Application mapping. | Yes. | high | Keep as T04 input, not a T03 decision. |
| VR-AMB-002 | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Unknown metadata marker validity is kind-specific, but one complete kind-by-kind unknown-field rule table is not centralized in the bounded specs. | ambiguous / needs decision | Yes when unknown marker is a source-format violation. | Yes. | medium | Current Spec unknown markers are invalid; sequential unknown marker policy needs confirmation from PRODUCT authoring authorities. |
| VR-AMB-003 | `spec:product.design_records.traceability.metadata_schema`; `spec:drmcp.design_records_mcp.schema.metadata_grammar`; Product work item/task authoring standards | Bounded specs disagree or are stale around workflow fields: PRODUCT traceability uses Work Item `source_refs` and no Task source field, while DRMCP metadata grammar still names `source_requirement` for Work Item and Task. | ambiguous / needs decision | Yes after authority is selected. | Yes. | high | This is the most important inventory conflict before response mapping. |
| VR-AMB-004 | `spec:drmcp.design_records_mcp.schema.diagnostics`; `DRMCP-TASK-MCP-022-01` | Older diagnostics text says Application aggregates findings and projects diagnostics; T01 now fixes Domain-owned rule-violation judgment and Application-owned presentation. The split is directionally clear but not yet mapped to validation response surfaces. | ambiguous / needs decision | Yes for judgment. | Yes for response projection. | high | T04 should consume this split. |
| VR-AMB-005 | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary` | Downstream detailed contracts still need exact rule IDs, graph queries, resolver outcomes, and validation placement. | ambiguous / needs decision | Yes. | Yes later. | high | This Task supplies provisional IDs only; implementation-ready rule contracts remain future work. |

## Ambiguous / missing validation rules

| item | status | downstream need |
|---|---|---|
| Workflow source-field authority conflict | Product traceability and current DRMCP metadata grammar disagree on `source_refs` versus `source_requirement`, especially for Work Item and Task. | Decide authoritative current field set before implementing validation or response mapping. |
| Unknown sequential metadata markers | Current Spec unknown markers are invalid, but a complete sequential-kind unknown-marker policy is not centralized. | Decide whether each sequential kind rejects, ignores, or preserves unknown markers. |
| Migration-sensitive severity | Product validation policy defines warnings/errors by migration phase, but T04 still needs to map Domain finding to public severity and `ok`. | Preserve violation judgment separately from presentation severity. |
| Topic cycle rule | Topic cycle is deferred until graph validation contract. | Define exact cycle semantics and whether it blocks validation. |
| Rule-ID granularity | Existing diagnostics categories are cause categories, not durable per-rule IDs. | T04 should use this provisional inventory or refine it before public mapping. |
| Request error versus validation diagnostic | `validate_records` has request/selection failures and normal validation findings. | T04 must keep operation errors separate from Domain validation findings. |
| Legacy relation validation | Lookup states are distinct, but rule-level outcomes and public mapping are not finalized under T01's Domain/Application split. | T04 should map disabled/unresolved/conflict/unreadable without changing semantic ownership. |
| Read/query operation warnings | `list_records`, `get_records`, and `resolve_reference` define warning/diagnostic outcomes that are adjacent to validation but not repository-validation findings. | Keep read-usecase warning mapping separate from `validate_records` response mapping unless T04 explicitly includes shared diagnostic projection. |
| Guidance operation availability | Guidance tools reuse Current Records state but have fixed-scope projection failures such as `guidance_catalog_unavailable` and `guide_unavailable`. | Treat as Application operation policy over Domain addressability, not a separate validation domain. |
| Authoring active/deferred boundary | Authoring transaction specs define proposal-local validation and write eligibility, while application architecture marks authoring internals as deferred. | Do not let T04 accidentally make authoring internals implementation-ready; record only the validation-like constraints as scoped inputs. |

## Evidence

Read and used the required startup files, plus bounded-scope validation authorities under:

- `product/records/spec/design-records/authoring-standards/`
- `product/records/spec/design-records/spec-format/`
- `product/records/spec/design-records/traceability/`
- `product/records/spec/design-records/namespace-model/`
- `product/records/spec/design-records/repository-layout/`
- `drmcp/records/spec/design-records-mcp/schema/`
- `drmcp/records/spec/design-records-mcp/tools/`
  - `list-records.md`
  - `get-records.md`
  - `resolve-reference.md`
  - `list-authoring-guides.md`
  - `get-authoring-guidance.md`
  - `authoring-transaction-model.md`
  - `propose-record-create.md`
  - `propose-record-update.md`
  - `accept-proposed-write.md`
- `drmcp/records/spec/design-records-mcp/namespace-scanning.md`
- `drmcp/records/spec/design-records-mcp/resolver.md`
- `drmcp/records/spec/application-architecture/`
  - `application-boundary-and-components.md`
  - `failure-and-evolution.md`
- `drmcp/records/spec/implementation/contracts/`
  - `composition-lifecycle/contract-boundary.md`
  - `infrastructure-io-adapters/contract-boundary.md`
  - `mcp-inbound-adapter/contract-boundary.md`

Search terms included validation, invalid, required, duplicate, canonical, diagnostic, resolve, missing, empty, normalize, and scope. Scope was not expanded.

## Boundary

This Task does not decide validation response contracts, public diagnostic wording, Go types, tests, fixtures, implementation plans, or canonical spec edits.

## Done condition

- Validation requirements are inventoried with provisional stable IDs.
- Gaps and ambiguity are explicitly listed.
- DRMCP-TASK-MCP-022-04 can decide validation usecase response mapping without guessing the rule inventory.

## Verification

- Bounded search stayed within the requested search scope.
- No validation response contract decision was made.
- No public diagnostic wording decision was made.
- No canonical spec authoring was performed.
- No ADR was created.
- No implementation planning, Go type, interface, method, test, or fixture decision was made.
- Stage, commit, and push were not performed.
