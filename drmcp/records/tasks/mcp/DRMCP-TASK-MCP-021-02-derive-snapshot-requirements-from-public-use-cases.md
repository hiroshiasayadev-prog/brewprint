# DRMCP-TASK-MCP-021-02: Investigate snapshot requirements from public use cases

- **id**: DRMCP-TASK-MCP-021-02
- **status**: done
- **date**: 2026-07-07
- **work_item**: DRMCP-WORK-MCP-021
- **task_type**: investigation
- **estimate**: 1.0d
- **depends_on**:
  - DRMCP-TASK-MCP-021-01
- **outputs**:
  - DRMCP-TASK-MCP-021-02

## Goal

Investigate the minimum Application-visible requirements for Current Records Snapshot and Legacy exact lookup map from public use-case contracts.

This Task uses the no-separate-INV-record exception. It records investigation evidence directly in this Task and does not create an Investigation record.

This Task prepares evidence for a later decision loop. It does not decide what W021 can require from Domain construction outputs.

## Work

### Decision boundary

This Task reads public use-case contracts and derives minimum request-state requirements.

`Current Records Snapshot` means the fresh immutable request-scoped transfer object for derived Current Records data produced by Current Records Snapshot Assembly.

`Legacy exact lookup map` means the fresh immutable request-scoped exact lookup map for accepted Legacy Archive issued IDs.

This Task investigates:

- which public use cases read Current Records Snapshot;
- which public use cases require Legacy exact lookup map;
- which lookup, projection, validation, relation, body-retrieval, and diagnostic needs appear in those use cases;
- which implementation-side Domain components appear to provide or own those needs;
- which candidate minimum Domain construction outputs Application may need;
- which requirements are already covered by existing Specifications;
- which gaps or ambiguous handoffs need later decision.

This Task does not decide:

- Domain parser internals;
- Domain logical-tree object layout;
- Domain relation-graph object layout;
- resolver or validator internal algorithms;
- public operation request or response contracts;
- concrete Infrastructure adapter fields;
- whether W021 uses new detailed Specifications or narrow synchronization;
- final Application snapshot or map requirements;
- Go signatures, structs, interfaces, functions, packages, algorithms, fixtures, or tests.

### Authority to read

| authority | purpose |
|---|---|
| `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary` | Application public use cases and shared orchestration ownership. |
| `spec:drmcp.design_records_mcp.tools.list_records` | Listing requirements over current addressable records. |
| `spec:drmcp.design_records_mcp.tools.get_records` | Exact retrieval, metadata, headings, body, and legacy retrieval requirements. |
| `spec:drmcp.design_records_mcp.tools.resolve_reference` | Reference classification, current lookup, and legacy fallback requirements. |
| `spec:drmcp.design_records_mcp.tools.validate_records` | Validation subject, relation, diagnostic, and legacy relation lookup requirements. |
| Guidance tool Specifications | Guidance projection needs over Current Records. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Current addressability, invalid-but-addressable sources, and duplicate conflicts. |
| `spec:drmcp.design_records_mcp.schema.record_source` | Metadata, headings, body, and source availability. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Current source scan and Legacy exact lookup-map construction. |
| `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary` | Domain-owned construction, lookup, graph, resolution, and validation outcomes. |
| DRMCP-TASK-MCP-021-01 | Prior W021 boundary decisions and corrected body-retention boundary. |

### Investigation questions

| ID | question | expected evidence |
|---|---|---|
| Q-001 | Which public use cases consume Current Records Snapshot or Legacy exact lookup map? | Public use-case consumption matrix. |
| Q-002 | What record data, lookup behavior, projection inputs, validation inputs, relation inputs, body-retrieval inputs, and diagnostic inputs does each public use case require? | Public use-case requirement matrix. |
| Q-003 | Which implementation-side Domain component appears to own each requirement? | Requirement-to-Domain-component matrix. |
| Q-004 | Which candidate Domain construction outputs must Application receive or be able to query? | Candidate construction-output table. |
| Q-005 | Which requirements are already covered by existing Specifications? | Coverage table with source refs. |
| Q-006 | Which gaps or ambiguous handoffs need a later decision loop? | Unresolved candidate list for T03. |

### Expected output shape

The final investigation result should include these tables:

| table | required contents |
|---|---|
| Public use-case consumption matrix | One row per public use case with Current Records Snapshot and Legacy exact lookup-map usage. |
| Public use-case requirement matrix | Record data, lookup behavior, projection inputs, validation inputs, relation inputs, body-retrieval inputs, and diagnostic inputs. |
| Requirement-to-Domain-component matrix | Required need, likely Domain owner, source evidence, and confidence. |
| Candidate Domain construction-output table | Minimum outputs Application may need, with uncertainty preserved. |
| Coverage and gap matrix | Existing authority, covered requirement, missing or ambiguous requirement, and candidate target. |
| T03 decision candidates | One row per unresolved question to confirm in the decision loop. |

### Investigation result

#### Public use-case consumption matrix

| public use case | reads Current Records Snapshot | reads Legacy exact lookup map | body read needed | diagnostics or source location needed | evidence refs |
|---|---|---|---|---|---|
| `list_records` | yes | no | no | conditional. Missing compact fields need warning source locations. | `spec:drmcp.design_records_mcp.tools.list_records`; `spec:drmcp.design_records_mcp.schema.record_model`; `spec:drmcp.design_records_mcp.namespace_scanning` |
| `get_records` | yes | yes, only for accepted legacy exact inputs | conditional. `include_body: true` needs complete source Markdown. | yes. Partial-success warning triggers need source-backed locations when applicable. | `spec:drmcp.design_records_mcp.tools.get_records`; `spec:drmcp.design_records_mcp.schema.record_source`; `spec:drmcp.design_records_mcp.namespace_scanning` |
| `resolve_reference` | yes | yes, only after current stage has no resolved target and legacy grammar is accepted | no | yes. Unresolved or unsupported outcomes include cause diagnostics. | `spec:drmcp.design_records_mcp.tools.resolve_reference`; `spec:drmcp.design_records_mcp.namespace_scanning` |
| `validate_records` | yes | yes, when legacy fallback is configured and relation targets need accepted legacy lookup | source Markdown read is required for parsing and validation. Public body projection is not required. | yes. Source, record, duplicate, current-relation, and legacy-relation findings require trustworthy locations. | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.design_records_mcp.schema.record_model`; `spec:drmcp.design_records_mcp.namespace_scanning` |
| `list_authoring_guides` | yes | no | yes. Projection needs the `## What this is` section body. | ambiguous. The spec defines catalog unavailability, not a normal diagnostic projection. | `spec:drmcp.design_records_mcp.tools.list_authoring_guides`; `spec:drmcp.application_architecture.runtime_and_state` |
| `get_authoring_guidance` | yes | no | yes. Projection needs complete source Markdown. | ambiguous. The spec defines guide errors, not a normal diagnostic projection. | `spec:drmcp.design_records_mcp.tools.get_authoring_guidance`; `spec:drmcp.application_architecture.runtime_and_state` |

#### Public use-case requirement matrix

| public use case | required record data | lookup behavior | projection inputs | validation inputs | relation inputs | body-retrieval inputs | diagnostic inputs | evidence refs |
|---|---|---|---|---|---|---|---|---|
| `list_records` | Current addressable record `ref`, `title`, `status`, and `date`. Invalid-but-addressable values remain available. Duplicate-conflict identities have no winner. | Exact configured app namespace, kind, domain, optional status, order, and limit over active index only. | Compact result rows, `has_more`, and operation warnings. | none | none | none | Missing compact-field warning triggers and required warning locations. | `spec:drmcp.design_records_mcp.tools.list_records`; `spec:drmcp.design_records_mcp.schema.fields`; `spec:drmcp.design_records_mcp.schema.record_model` |
| `get_records` | Current normalized metadata, headings, optional body, and legacy minimal metadata, headings, and optional body. | Ordered deduplication. Legacy exact input queries only legacy map. Current spec and sequential inputs query only active index. | Successful record wrappers and operation warnings in first-occurrence request order. | none | none | Complete current or legacy Markdown source when `include_body` is true. | Malformed, unsupported, unresolved, unavailable legacy, duplicate input, conflict, and unreadable-source warning triggers. | `spec:drmcp.design_records_mcp.tools.get_records`; `spec:drmcp.design_records_mcp.schema.record_source`; `spec:drmcp.design_records_mcp.namespace_scanning` |
| `resolve_reference` | Current target `ref`, `kind`, `title`, `status`, or legacy target `ref` and `kind`. | Current canonical grammar and active index first. Accepted legacy lookup runs only after no current target resolves. | Status, target object, and cause diagnostics. | none | none | none. Legacy content retrieval is delegated to `get_records`. | Resolver cause diagnostics for malformed, unsupported, current unresolved, current conflict, legacy disabled, legacy unresolved, legacy conflict, and legacy unreadable outcomes. | `spec:drmcp.design_records_mcp.tools.resolve_reference`; `spec:drmcp.design_records_mcp.namespace_scanning` |
| `validate_records` | Configuration state, every retained current source, parse-failed source, identity-less validation-only source, addressable record, and duplicate-conflict group. | Request scope selects repository-wide, app, or exact current ref. Current relation lookup uses exact active-index identity. Legacy relation lookup uses separate exact legacy map. | `ok`, effective scope, selected-subject counts, and diagnostics. | Per-source syntax, metadata, identity, document-shape validation, local validation, relation validation, Topics graph validation, and aggregation. | Declared current canonical relation targets and accepted legacy relation targets. | Source Markdown read and parser output. No public body field. | Source, record, duplicate, current-relation, legacy-relation, conflict-member, conflict-candidate, and unreadable-source locations. | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`; `spec:drmcp.design_records_mcp.schema.record_model` |
| `list_authoring_guides` | Addressable current package Spec identity, first H1 title, and `## What this is` body from `spec:design_records.authoring_standards.*`. | Fixed Current Records scope under `design_records`. No public `list_records` call. No legacy lookup. | Guide `id`, `title`, and `abstract` ordered by canonical ID. | none | none | Source read and extraction of one section body. | Catalog-unavailable condition for conflicted, unreadable, or unprojectable in-scope candidates. | `spec:drmcp.design_records_mcp.tools.list_authoring_guides`; `spec:drmcp.design_records_mcp.schema.record_source` |
| `get_authoring_guidance` | Exact package Spec identity, first H1 title, and complete Markdown source. | Exact fixed-scope current lookup under `spec:design_records.authoring_standards.*`. No public `get_records` call. No legacy lookup. | Guidance `id`, `title`, and complete `content`. | none | none | Complete current Markdown source. | Guide-not-found and guide-unavailable conditions. Normal diagnostic shape is not specified here. | `spec:drmcp.design_records_mcp.tools.get_authoring_guidance`; `spec:drmcp.design_records_mcp.schema.record_source` |

#### Requirement-to-Domain-component matrix

| required need | likely Domain owner or boundary owner | supporting spec evidence | confidence | uncertainty |
|---|---|---|---|---|
| Parse raw Markdown into typed source, typed record, and parse findings. | Record Parser | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`; `spec:drmcp.design_records_mcp.schema.record_source` | high | Detailed parser output fields remain downstream. |
| Preserve current source provenance for retained sources and diagnostics. | Application-owned source contract / source provenance, with Parser and validators consuming location material | `spec:drmcp.implementation.contracts.infrastructure_io_adapters.contract_boundary`; `spec:drmcp.design_records_mcp.schema.record_model` | high | Exact provenance object boundary is not defined. |
| Build current canonical identity lookup and active-index addressability. | Current Records Logical Tree | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`; `spec:drmcp.design_records_mcp.schema.record_model`; `spec:drmcp.design_records_mcp.namespace_scanning` | high | The detailed handoff between active index and logical tree is not yet specified. |
| Preserve invalid-but-addressable current records. | Current Records Logical Tree and Record Parser | `spec:drmcp.design_records_mcp.schema.record_model`; `spec:drmcp.design_records_mcp.schema.fields` | high | Which parsed invalid values remain queryable by Application needs confirmation. |
| Preserve duplicate-conflict groups with no addressable winner. | Current Records Logical Tree | `spec:drmcp.design_records_mcp.schema.record_model`; `spec:drmcp.design_records_mcp.namespace_scanning` | high | Exact conflict-member projection for diagnostics is owned elsewhere. |
| Provide compact listing fields. | Current Records Logical Tree over Parser field outputs | `spec:drmcp.design_records_mcp.tools.list_records`; `spec:drmcp.design_records_mcp.schema.fields` | high | Application ordering and limit policy remain use-case owned. |
| Provide exact current retrieval metadata and headings. | Current Records Logical Tree over Parser source outputs | `spec:drmcp.design_records_mcp.tools.get_records`; `spec:drmcp.design_records_mcp.schema.record_source`; `spec:drmcp.design_records_mcp.schema.fields` | high | Snapshot should expose headings or a query for headings. Exact surface is undecided. |
| Support body retrieval without requiring every body in snapshot memory. | Application-owned source contract / source provenance, supported by Record Parser source availability | DRMCP-TASK-MCP-021-01 D-006; `spec:drmcp.design_records_mcp.tools.get_records`; `spec:drmcp.design_records_mcp.tools.get_authoring_guidance` | medium | T03 must confirm whether Application queries body through source provenance, Domain, or source contract. |
| Execute current-first reference resolution. | Reference Resolution | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`; `spec:drmcp.design_records_mcp.tools.resolve_reference` | high | Detailed resolver input shape and result taxonomy remain downstream. |
| Provide logical-tree lookup results for validation and retrieval. | Current Records Logical Tree | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`; `spec:drmcp.design_records_mcp.tools.validate_records` | high | Use-case direct lookup versus Domain query capability needs confirmation. |
| Provide relation graph capability over current records. | Record Relation Graph | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`; `spec:drmcp.design_records_mcp.tools.validate_records` | high | Detailed graph query set is not decided. |
| Validate one typed source or record locally. | Local Record Validation | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`; `spec:drmcp.design_records_mcp.tools.validate_records` | high | Rule placement and rule IDs remain outside T02. |
| Validate relations, reciprocity, and Topics graph context. | Relation Graph Validation | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`; `spec:drmcp.design_records_mcp.tools.validate_records` | high | Exact Topics graph inputs need later detailed Domain work. |
| Build and expose legacy exact lookup outcomes. | Legacy exact lookup map | `spec:drmcp.design_records_mcp.namespace_scanning`; DRMCP-TASK-MCP-021-01 D-007 | high | Whether Domain queries the map directly or through Reference Resolution / Relation Graph Validation needs confirmation. |
| Distinguish legacy disabled, absent, duplicate, and unreadable outcomes. | Legacy exact lookup map plus Reference Resolution and Relation Graph Validation consumers | `spec:drmcp.design_records_mcp.tools.resolve_reference`; `spec:drmcp.design_records_mcp.tools.validate_records` | high | Public outcome mapping differs between resolver, retrieval, and validation. |
| Project guidance catalog and detail from portable package Specs. | Current Records Logical Tree plus Application Guidance projection | `spec:drmcp.design_records_mcp.tools.list_authoring_guides`; `spec:drmcp.design_records_mcp.tools.get_authoring_guidance`; `spec:drmcp.application_architecture.runtime_and_state` | medium | Section extraction ownership is not explicit. |
| Select operation failure when mandatory state is incomplete or untrustworthy. | Application Use Cases | `spec:drmcp.application_architecture.failure_and_evolution`; `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary` | high | Exact error-code representation remains downstream. |

#### Candidate Domain construction-output table

| candidate output or query surface | Application need | likely producer | consumers | evidence refs | uncertainty |
|---|---|---|---|---|---|
| Current addressable identity lookup | Listing, exact current retrieval, resolver, validation selector, and relation target checks. | Current Records Logical Tree | Application use cases, Reference Resolution, Relation Graph Validation | `spec:drmcp.design_records_mcp.schema.record_model`; `spec:drmcp.design_records_mcp.tools.validate_records` | Exact lookup-result taxonomy is not yet detailed. |
| Current compact field projection | `list_records` needs `ref`, `title`, `status`, and `date`. | Record Parser and Current Records Logical Tree | List Records Use Case | `spec:drmcp.design_records_mcp.tools.list_records`; `spec:drmcp.design_records_mcp.schema.fields` | Whether Application receives fields directly or queries snapshot is undecided. |
| Current metadata projection | `get_records` needs normalized current metadata. | Record Parser and Current Records Logical Tree | Get Records Use Case | `spec:drmcp.design_records_mcp.tools.get_records`; `spec:drmcp.design_records_mcp.schema.fields` | Missing-field omission behavior must be preserved. |
| Current heading projection | `get_records` and Guidance need source headings or H1 title. | Record Parser | Get Records Use Case and Guidance use cases | `spec:drmcp.design_records_mcp.tools.get_records`; `spec:drmcp.design_records_mcp.tools.get_authoring_guidance` | Section-level body extraction surface is not defined. |
| Body retrieval reference or source-read capability | `get_records(include_body: true)` and `get_authoring_guidance` need verbatim Markdown without retaining all bodies in snapshot memory. | Application-owned source contract / source provenance, with Domain source identity support | Get Records Use Case and Get Authoring Guidance Use Case | DRMCP-TASK-MCP-021-01 D-006; `spec:drmcp.design_records_mcp.schema.record_source` | T03 must confirm the minimum output is a source reference, not body content. |
| Retained validation subjects | `validate_records` needs source, parse-failed, identity-less, addressable, and duplicate-conflict subjects. | Record Parser and Current Records Logical Tree | Validate Records Use Case, Local Record Validation, Relation Graph Validation | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.design_records_mcp.schema.record_model` | Exact subject selection object is not specified. |
| Duplicate conflict group representation | Listing excludes conflicted identities. Validation and diagnostics need conflict members. | Current Records Logical Tree | List Records Use Case and Validate Records Use Case | `spec:drmcp.design_records_mcp.schema.record_model`; `spec:drmcp.design_records_mcp.tools.validate_records` | Conflict-member location shape is outside T02. |
| Source-backed location references | Warnings, diagnostics, conflicts, and unreadable-source outcomes need trustworthy locations. | Parser, logical tree, validators, and Application-owned source provenance | Application projection and diagnostics contract | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.application_architecture.failure_and_evolution` | The diagnostics contract owns concrete location fields and was not changed by T02. |
| Relation graph capability | Relation validation and Guidance-related query use need graph queries. | Record Relation Graph | Relation Graph Validation and Guidance-related query use | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`; `spec:drmcp.design_records_mcp.tools.validate_records` | Guidance graph use is only mentioned as related query use. |
| Reference resolution outcome | `resolve_reference` needs current-first and optional legacy fallback outcome. | Reference Resolution | Resolve Reference Use Case and Get Records Use Case when resolving semantic references is not public exact retrieval | `spec:drmcp.design_records_mcp.tools.resolve_reference`; `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary` | `get_records` explicitly does not call resolver. |
| Local validation findings | `validate_records` needs source-local findings. | Local Record Validation | Validate Records Use Case | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary` | Rule identifiers remain downstream. |
| Relation validation findings | `validate_records` needs current and legacy relation findings. | Relation Graph Validation | Validate Records Use Case | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary` | Exact finding aggregation boundary remains Application-owned. |
| Legacy exact lookup map | `get_records`, `resolve_reference`, and `validate_records` need exact legacy lookup. | Legacy exact lookup map | Get Records Use Case, Reference Resolution, Relation Graph Validation | `spec:drmcp.design_records_mcp.namespace_scanning`; `spec:drmcp.design_records_mcp.tools.get_records`; DRMCP-TASK-MCP-021-01 D-007 | Direct Application access versus Domain-mediated access needs T03 decision. |
| Legacy source reference for retrieval | `get_records(include_body: true)` needs complete archived source for successful legacy retrieval. | Legacy exact lookup map plus Legacy source contract | Get Records Use Case | `spec:drmcp.design_records_mcp.tools.get_records`; `spec:drmcp.design_records_mcp.namespace_scanning` | The map may carry a source reference rather than body or path. |

#### Coverage and gap matrix

| requirement | existing spec coverage | status | ambiguous or missing part | T03 decision needed | evidence refs |
|---|---|---|---|---|---|
| Every Read, Validation, or Guidance request builds one fresh immutable Current Records Snapshot. | Runtime and namespace scanning state the request-scoped snapshot rule. | covered | none for lifetime. | no | `spec:drmcp.application_architecture.runtime_and_state`; `spec:drmcp.design_records_mcp.namespace_scanning` |
| Legacy state is separate and loaded only when an operation requires compatibility lookup. | Runtime, namespace scanning, T01 D-004, and T01 D-007 cover separation and optional loading. | covered | Existing specs still use `Legacy Lookup State` wording in places. | yes, for accepted W021 term projection | `spec:drmcp.application_architecture.runtime_and_state`; DRMCP-TASK-MCP-021-01 |
| `list_records` uses only current addressable records. | Public contract and schema specs cover active-index-only listing. | covered | Exact Domain construction-output surface is not detailed. | yes | `spec:drmcp.design_records_mcp.tools.list_records`; `spec:drmcp.design_records_mcp.schema.record_model` |
| `get_records` exact current and legacy retrieval. | Public contract covers classification, projection, partial success, and legacy retrieval. | partially covered | Construction output for body retrieval and source references is not detailed. | yes | `spec:drmcp.design_records_mcp.tools.get_records`; DRMCP-TASK-MCP-021-01 D-006 |
| `resolve_reference` current-first and legacy fallback. | Public contract and Domain boundary cover Reference Resolution ownership. | partially covered | Direct inputs from snapshot and map to Reference Resolution need confirmation. | yes | `spec:drmcp.design_records_mcp.tools.resolve_reference`; `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary` |
| `validate_records` validation subjects and relation lookup. | Public contract covers subjects, pass order, relation lookup, and execution prerequisites. | partially covered | Snapshot construction outputs for validators are not detailed. | yes | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary` |
| Guidance catalog and detail use Current Records. | Guidance specs and runtime state cover fixed current scope and no separate guide index. | partially covered | Section extraction ownership and failure-to-diagnostic boundary are ambiguous. | yes | `spec:drmcp.design_records_mcp.tools.list_authoring_guides`; `spec:drmcp.design_records_mcp.tools.get_authoring_guidance` |
| Body content is not mandatory snapshot memory. | T01 D-006 covers the corrected boundary. Public retrieval specs need source body. | partially covered | The minimal source-reference output is not specified. | yes | DRMCP-TASK-MCP-021-01 D-006; `spec:drmcp.design_records_mcp.schema.record_source` |
| Source-backed diagnostic locations gate normal response trustworthiness. | Public contracts and failure architecture cover required trustworthy locations. | partially covered | Concrete location fields belong to diagnostics work and are not W021-owned. | yes, to define W021 boundary only | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.application_architecture.failure_and_evolution` |
| Domain components own parser, logical tree, graph, resolver, and validation semantics. | Domain boundary covers module responsibility. | partially covered | Detailed Domain object outputs are explicitly downstream. | yes, for Application-facing minimum only | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`; DRMCP-WORK-MCP-021 |
| Infrastructure owns concrete source access behind source contracts. | Infrastructure boundary covers source-family split and provenance preservation. | partially covered | Application-owned source contract fields are not detailed. | yes | `spec:drmcp.implementation.contracts.infrastructure_io_adapters.contract_boundary`; `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary` |
| Application selects execution failure for incomplete mandatory state. | Failure architecture and public contracts cover no partial normal response. | covered | Exact operation error codes remain outside T02. | no for T02; later detailed response specs may need it | `spec:drmcp.application_architecture.failure_and_evolution`; `spec:drmcp.design_records_mcp.tools.validate_records` |

#### T03 decision candidates

| candidate ID | decision candidate | why T03 must confirm it | source evidence | uncertainty to preserve |
|---|---|---|---|---|
| T03-C-001 | Accept or correct the public use-case consumption matrix. | Downstream specs should not author snapshot outputs before the operation consumers are accepted. | Public tool specs; `spec:drmcp.application_architecture.runtime_and_state` | Guidance diagnostic/source-location needs are less explicit than read and validation tools. |
| T03-C-002 | Confirm the Current Records Snapshot minimum requirement set. | W021 needs a baseline for current identity lookup, fields, headings, validation subjects, duplicate groups, relation inputs, and provenance. | `spec:drmcp.design_records_mcp.schema.record_model`; `spec:drmcp.design_records_mcp.tools.validate_records` | Exact Application-visible surface can be direct fields, query capability, or opaque Domain outcomes. |
| T03-C-003 | Confirm the Legacy exact lookup-map minimum requirement set. | Legacy requirements must remain separate from Current Records Snapshot. | `spec:drmcp.design_records_mcp.namespace_scanning`; `spec:drmcp.design_records_mcp.tools.get_records`; `spec:drmcp.design_records_mcp.tools.resolve_reference`; `spec:drmcp.design_records_mcp.tools.validate_records` | Direct Application access versus Domain-mediated access is not resolved. |
| T03-C-004 | Confirm that body retrieval uses source references or source contract access rather than mandatory snapshot body retention. | T01 D-006 prevents heavy body retention, but public operations need verbatim source content. | DRMCP-TASK-MCP-021-01 D-006; `spec:drmcp.design_records_mcp.tools.get_records`; `spec:drmcp.design_records_mcp.tools.get_authoring_guidance` | Exact source-reference shape is not a T03 outcome unless W021 chooses it as a contract gap. |
| T03-C-005 | Confirm Application-to-Domain construction-output requirements without deciding Domain internals. | W021 can require minimum Domain outcomes but cannot define parser, tree, graph, resolver, or validator layout. | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`; DRMCP-WORK-MCP-021 | Detailed Domain specs may later refine or replace candidate output names. |
| T03-C-006 | Confirm diagnostic and source-location boundary for W021. | Public tools fail when required trustworthy locations cannot be built, but concrete diagnostics remain W006-owned. | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.application_architecture.failure_and_evolution` | W021 should define only required provenance handoff, not diagnostic schema. |
| T03-C-007 | Confirm Guidance projection requirements over Current Records. | Guidance specs are public use cases and explicitly avoid a separate guide index. | `spec:drmcp.design_records_mcp.tools.list_authoring_guides`; `spec:drmcp.design_records_mcp.tools.get_authoring_guidance`; `spec:drmcp.application_architecture.runtime_and_state` | Section extraction ownership is not explicit in module-boundary specs. |
| T03-C-008 | Classify downstream route for covered, partially covered, and missing items. | W021 should avoid premature detailed Specification authoring. | DRMCP-WORK-MCP-021; DRMCP-TASK-MCP-021-03 | T03 must decide narrow synchronization, new detailed specs, or Domain follow-up candidates. |

## Done condition

- Public use-case consumption of Current Records Snapshot and Legacy exact lookup map is inventoried.
- Public use-case record requirements are collected from Specification evidence.
- Each requirement is mapped to an implementation-side Domain component where existing evidence supports the mapping.
- Candidate Domain construction outputs are listed without deciding Domain internals.
- Existing-spec coverage and candidate gaps are classified.
- T03 decision candidates are listed with uncertainty preserved.
- No Investigation record, Specification, ADR, implementation Task, production implementation, Go contract-shape probe, fixture, test, or independent review is authored by this Task.

## Verification

- Authority read scope check: pass. Only the user-listed public use-case, schema, runtime, failure, and boundary specs were read, plus the two existing Guidance tool specs found under the tool spec directory.
- No-separate-INV-record check: pass. No Investigation record was created.
- Output-shape check: pass. The six required tables are recorded in this Task body.
- Decision-boundary check: pass. The tables preserve candidate gaps for T03 and do not decide Domain internals.
- Prohibited-work check: pass. No ADR, Specification, implementation Task, production implementation, Go contract-shape probe, fixture, test, independent review, stage, commit, or push was performed by this Task.
- Status check: pass. This Task is marked `done`.
- Scoped diff and whitespace check: pass. `git.inspect_diff` returned `result: pass` for this Task path. LF to CRLF warning is advisory only.

## Evidence

- DRMCP-TASK-MCP-021-01 completed the initial Application shared-orchestration boundary ledger.
- User clarified that W021 should derive snapshot requirements from public use-case needs before authoring detailed Specifications.
- User clarified that Domain detailed specifications should later decide Domain object internals and feed back into snapshot Specifications.
- User clarified that T02 should be an investigation-type Task using the no-separate-INV-record exception.
- User clarified that the next Task after T02 should run the decision loop over the investigation tables.
- This Task is created before any W021 detailed Specification authoring Task.
- Read `spec:drmcp.design_records_mcp.tools.list_records` for compact current-listing requirements.
- Read `spec:drmcp.design_records_mcp.tools.get_records` for exact current and legacy retrieval requirements.
- Read `spec:drmcp.design_records_mcp.tools.resolve_reference` for current-first and legacy-fallback resolver requirements.
- Read `spec:drmcp.design_records_mcp.tools.validate_records` for validation subject, relation, and diagnostic-location requirements.
- Read `spec:drmcp.design_records_mcp.tools.list_authoring_guides` and `spec:drmcp.design_records_mcp.tools.get_authoring_guidance` after finding Guidance tool specs in the public tool directory.
- Read `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary` for Application shared-orchestration ownership.
- Read `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary` for Domain parser, logical tree, graph, resolver, and validator ownership.
- Read `spec:drmcp.implementation.contracts.infrastructure_io_adapters.contract_boundary` for source-family access and provenance boundaries.
- Read `spec:drmcp.application_architecture.runtime_and_state` and `spec:drmcp.application_architecture.failure_and_evolution` for request-state lifetime and trustworthy-result failure rules.
- Read `spec:drmcp.design_records_mcp.schema.record_model`, `spec:drmcp.design_records_mcp.schema.record_source`, `spec:drmcp.design_records_mcp.schema.fields`, and `spec:drmcp.design_records_mcp.namespace_scanning` for current source, active-index, field, body, and legacy lookup-map evidence.
- Public use-case consumption, requirement, Domain mapping, construction-output, coverage/gap, and T03 decision-candidate tables are recorded directly in this Task under the no-separate-INV-record exception.
