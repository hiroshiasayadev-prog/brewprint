# PRODUCT-TASK-SPEC-013-01: Record corrected whole-tree package source boundary

- **id**: PRODUCT-TASK-SPEC-013-01
- **status**: done
- **date**: 2026-06-26
- **work_item**: PRODUCT-WORK-SPEC-013
- **source_requirement**: PRODUCT-REQ-SPEC-003
- **estimate**: 1d
- **depends_on**: []
- **outputs**:
  - Corrected whole-tree package source boundary recorded in this Task Evidence
  - Source-authoring and placement findings retained as audit evidence only

## Goal

Record the corrected PRODUCT source boundary for the first portable Design Records standards package.

The current package boundary is the whole tree `product/records/spec/design-records/`, copied to `<exe-dir>/design-records/` with only canonical spec-ref prefix rewriting from `spec:product.design_records` to `spec:design_records`.

The previous section-level package projection premise is superseded. The detailed inventory below is retained only as source-authoring / placement audit evidence.

## Work

- Treat `product/records/spec/design-records/` as the authoritative source tree for the package.
- Distinguish package boundary from source-authoring findings.
- Record that the package producer does not perform section-level semantic selection, app-local filtering, or prose generalization.
- Treat app-local / wiring / project-tracking content in the source tree as source-authoring findings. The producer emits warnings during generation or check execution and does not own persistent warning evidence or automatic routing. When a human reviewer decides that tracked source correction is required, a source-authoring Requirement is manually created. Warnings and open follow-ups do not block package generation or distribution.
- Retain the previous detailed inventory only as historical audit evidence for source cleanup.
- Do not define generation implementation, tooling placement, execution algorithm, loader implementation, or release-validation implementation in this Task.
- Do not modify PRODUCT semantic specs, package artifacts, implementation, fixtures, or unrelated records.

## Done condition

Task Evidence records the corrected whole-tree package boundary, clearly supersedes the previous projection premise, preserves the previous inventory only as source-authoring audit evidence, and distinguishes source-authoring findings from package-generation instructions.

## Verification

- Confirm the current package source boundary is `product/records/spec/design-records/` as a whole tree.
- Confirm the package-generation rule is copy plus canonical spec-ref prefix rewrite only.
- Confirm the previous `section_projection=33 / excluded=1` model is marked superseded.
- Confirm the previous detailed tables are not package-generation instructions.
- Confirm the previous Independent review `PASS` is scoped to the superseded premise and is not a verdict on this correction.
- Confirm no unrelated file changed.

## Evidence

### Current normative summary

The authoritative source boundary for the first portable package is the whole tree:

```text
product/records/spec/design-records/
```

The bundled destination is:

```text
<exe-dir>/design-records/
```

The package producer copies the whole tree and performs only canonical spec-ref prefix rewriting:

```text
spec:product.design_records -> spec:design_records
spec:product.design_records.<suffix> -> spec:design_records.<suffix>
```

The producer does not perform section-level meaning selection, app-local semantic filtering, package-time prose generalization, package-specific manifest generation, or package-specific wrapper overview creation.

The previous `section_projection=33 / excluded=1` package model is superseded. The table below is not a package projection instruction set. It is retained only as source-authoring / placement audit evidence showing where the current source tree appears to contain app-local, wiring, project-tracking, or implementation-adjacent material; source-authoring cleanup of such material remains desirable and is separately owned, the producer emits warnings for these findings, and warning closure is not required before package generation or distribution.

Package boundary and source-authoring findings are separate:

| topic | current treatment |
| --- | --- |
| Package source boundary | Whole `product/records/spec/design-records/` tree. |
| Package namespace | `design_records`. |
| Package spec root | `spec:design_records`. |
| Guidance root | `spec:design_records.authoring_standards`. |
| Source cleanup | Producer emits warnings during generation or check execution without owning persistent evidence or automatic routing. When a human reviewer determines that tracked correction is required, a source-authoring Requirement is manually created. Warnings do not block generation or distribution. |
| Package producer | Copies and rewrites canonical spec refs only; does not choose semantic subsets. |
| Existing detailed inventory | Historical source-authoring audit evidence only. |
| Previous Independent review `PASS` | Applies only to the now-superseded projection premise and is not a review verdict for this correction. |
| Current Task status | `done`; independent review of the current whole-tree correction passed. |

### Historical section-projection audit: 0. Inclusion inventory

Historical audit table from the superseded section-level package model.
One row per source spec. All 34 specs under `product/records/spec/design-records/`.
Inclusion mode definitions used by the historical table: `whole_spec` = complete source projected unchanged under the old premise; `section_projection` = specific sections/rows only under the old premise; `excluded` = not projected under the old premise.

| source ref | inclusion mode | included contract area | excluded content | reason |
| --- | --- | --- | --- | --- |
| `spec:product.design_records` | `section_projection` | What this is; Current contract; Non-goals; Rules; Boundary; Topics; Topic map | `## Related specs` rows for `spec:product`, `spec:product.brewprint`, `spec:product.bpdsl` | Root overview retained as package entry point; Related specs rows point to host-specific PRODUCT placement router and Brewprint/BPDSL areas excluded per host-independent boundary |
| `spec:product.design_records.namespace_model` | `section_projection` | What this is; Current contract; App namespace and domain namespace; Topics | `## Boundary` rows for `spec:product.brewprint.namespaces` and `spec:product.brewprint.compatibility`; `## Current placement and future layout` section (contains `spec:product.bpdsl.repository_implementation_flow` pointer); body sentences pointing to `spec:product.brewprint.namespaces` and `spec:product.brewprint.compatibility`; `## Related specs` rows for those same two refs | Namespace model concepts intact; Brewprint profile and compatibility facts are host-specific profile data; BPDSL staging pointer is excluded per host-independent boundary |
| `spec:product.design_records.namespace_model.app_namespaces` | `section_projection` | What this is; Current contract; Rules | `## Boundary` rows for `spec:product.brewprint.namespaces.app_namespaces` and `spec:product.brewprint.namespaces.domain_catalog`; `## Related specs` row for `spec:product.brewprint.namespaces` | App namespace concept, domain relationship, and selection rules intact; Brewprint current registry pointers are host-specific profile data |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | `section_projection` | Grammar (REQ/WORK/INV/ADR and TASK forms); Sequence format; Canonical reference | Inline note "> Existing issued ADR IDs are governed by `spec:product.brewprint.compatibility.legacy_id_compatibility`" in `### REQ / WORK / INV / ADR`; `## Related specs` rows for `spec:product.brewprint.compatibility.legacy_id_compatibility` and `spec:product.brewprint.compatibility.existing_artifacts` | Complete ID grammar and canonical-ref policy intact; Brewprint V01 compatibility pointer is informational; inline note and Related specs rows omitted in projection |
| `spec:product.design_records.namespace_model.existing_artifacts` | `section_projection` | What this is; Current contract; Rules | Body sentence "Historical Brewprint ownership decisions… They are recorded in `spec:product.brewprint.compatibility.existing_artifacts`"; `## Boundary` rows for `spec:product.brewprint.namespaces.domain_catalog` and `spec:product.brewprint.compatibility.existing_artifacts`; `## Related specs` row for `spec:product.brewprint.compatibility.existing_artifacts` | New-artifact namespace selection rules intact; Brewprint historical attribution is host-specific compatibility |
| `spec:product.design_records.namespace_model.subdomain_model` | `section_projection` | What this is; Current contract; Definition and representation; Write-time advisory; Boundary | `## Related specs` row for `spec:product.brewprint.compatibility.existing_artifacts` | Subdomain concept, representation, and write-time advisory intact; legacy attribution context pointer excluded |
| `spec:product.design_records.repository_layout` | `section_projection` | What this is (excl. brewprint.layout sentence); Current contract; Rules; Boundary; Topics; Sources | `## BPDSL staging note` section (all 4 rows point to `spec:product.bpdsl.repository_implementation_flow`); body sentence "The current Brewprint repository inventory is documented by `spec:product.brewprint.layout`" in `## What this is`; `## Related specs` rows for `spec:product.brewprint.layout` and `spec:product.bpdsl.repository_implementation_flow` | Kind-first placement contract, record layout rules, and topic-tree placement intact; BPDSL staging note and Brewprint inventory pointer excluded |
| `spec:product.design_records.repository_layout.record_discovery_paths` | `section_projection` | What this is; Current contract; Rules; Related specs; Sources | `## DRMCP boundary` section entirely (all rows; contains `spec:drmcp.design_records_mcp.namespace_scanning` and `spec:drmcp.design_records_mcp.schema.discovery`) | Per-kind path patterns and discovery rules intact; DRMCP implementation boundary section excluded |
| `spec:product.design_records.spec_format` | `section_projection` | What this is; Topics (5 rows, excl. follow_up_boundary row); Related specs (namespace_model and repository_layout rows only) | `## Topics` row for `spec:product.design_records.spec_format.follow_up_boundary`; `## Related specs` row for `spec:drmcp.design_records_mcp.overview` | Navigation entry point for 5 included child specs intact; follow_up_boundary is excluded from package (project-tracking context); DRMCP boundary pointer excluded |
| `spec:product.design_records.spec_format.overview` | `section_projection` | What this is; Current contract; Non-goals; Topic map rows for Document shape, Topics table, Spec ID-as-ref, and Validation policy; Related specs | `## Topic map` row for `spec:product.design_records.spec_format.follow_up_boundary` | Overview framing and included child pointers retained; Follow-up boundary row omitted because `follow_up_boundary` is excluded from the package |
| `spec:product.design_records.spec_format.document_shape` | `section_projection` | Accepted spec kinds; `contract_class`; H1 form; visible metadata shape; required-section matrix; generic validation conditions | `PRODUCT-WORK-SPEC-003`; Brewprint migration-warning context; concrete DRMCP-specific H1 example; wording tied to current migration execution | Generic document-shape contract retained; project execution refs/examples are excluded or generalized for host-independent package projection |
| `spec:product.design_records.spec_format.topics_table` | `section_projection` | Topics table columns; parent declaration rules; duplicate-parent rules; canonical `ref` behavior; generic validation conditions | `PRODUCT-WORK-SPEC-002`; Brewprint migration-phase wording; project-specific follow-up ownership | Generic topic/ref validation contract retained; project follow-up ownership and migration context excluded or generalized |
| `spec:product.design_records.spec_format.spec_id_as_ref` | `section_projection` | Path-derived identity grammar; H1-adjacent ID rules; one-to-one path/ref behavior; parent grammar; generic compatibility boundary | `PRODUCT-WORK-SPEC-002`; `PRODUCT-WORK-SPEC-005`; concrete `product`, `drmcp`, and `bpdsl` repository examples where host-specific; migration, inventory, and transient working-state execution context | Generic spec ID-as-ref contract retained; host repository examples and project migration/follow-up refs excluded or generalized |
| `spec:product.design_records.spec_format.validation_policy` | `section_projection` | Portable validation semantics: parser-aware heading detection; structural validation conditions; severity semantics; document-shape and topic/ref validation rules | Concrete Brewprint Work Item owner column values; current migration-phase tracking; current DRMCP implementation-phase assignments; temporary PRODUCT tooling ownership | Validation semantics retained; Brewprint execution ownership and migration/implementation phase tracking excluded or generalized through projection |
| `spec:product.design_records.spec_format.follow_up_boundary` | `excluded` | N/A | Entire spec | Maps Brewprint PRODUCT-WORK-SPEC-* and DRMCP-WORK-SPEC-* work items to follow-up ownership areas; PRODUCT-internal project-management context; contains no semantic contracts needed by portable authoring consumers |
| `spec:product.design_records.traceability` | `section_projection` | What this is; Current contract; Non-goals; Traceability model; Topics | `## Sources` row for `spec:product.brewprint.compatibility` | Traceability contract and active reference classes intact; Brewprint compatibility source pointer excluded from projection |
| `spec:product.design_records.traceability.semantic_ref` | `section_projection` | Current spec ref model; generic active `spec:` ref class; obsolete semantic-prefix assumptions; Related specs | `## Ref classes` row for Legacy issued ID compatibility treatment | Generic active `spec:` and record ID-as-ref semantics retained; Brewprint legacy-ID compatibility authority excluded |
| `spec:product.design_records.traceability.artifact_refs` | `section_projection` | What this is; Active reference classes (excl. compatibility cell); Spec refs; Record ID-as-ref (excl. compatibility pointer column); Investigation reference boundary; Workflow relation identity; Reference boundary | `## Active reference classes` table: Legacy issued ID row compatibility-owner cell (`spec:product.brewprint.compatibility`); `## Record ID-as-ref` table: compatibility pointer column values for all 5 rows; `## Sources` row for `spec:product.brewprint.compatibility` | All active reference class rules intact; Brewprint compatibility ownership cell and sources row excluded from projection |
| `spec:product.design_records.traceability.metadata_schema` | `section_projection` | Visible spec metadata roles; investigation reference metadata; workflow relation metadata; PRODUCT-owned invalid semantic conditions; generic metadata boundary | T05 evidence reference; Brewprint migration-history wording; concrete DRMCP parser/writer ownership wording where implementation-specific; BPDSL endpoint statement | Generic metadata and relation boundary retained; project history and implementation-owner statements excluded or generalized |
| `spec:product.design_records.traceability.resolve_and_validation` | `section_projection` | What this is; Supported canonical inputs; Lookup sources; Spec resolution boundary; Investigation validation; Workflow relation validation; Duplicate identity conditions; Excluded implementation behavior; Resolve and validation boundary | `## Related specs` row for `spec:product.brewprint.compatibility` | Canonical lookup sources and invalid conditions intact; Brewprint compatibility Related specs pointer excluded |
| `spec:product.design_records.artifact_model` | `section_projection` | What this is; Current contract; Artifact classes (excl. BPDSL staging sentence); Source of truth (excl. brewprint.layout row); Disposition of previous ownership statements; Artifact model boundary (excl. BPDSL sentence); Topics; Sources | `## Source of truth` row for `spec:product.brewprint.layout`; `## Navigation` row for `spec:product.bpdsl`; `## Related specs` row for `spec:product.bpdsl`; body sentences referencing `spec:product.bpdsl` and `spec:product.brewprint.layout` in Artifact classes and Artifact model boundary sections | Artifact roles, source-of-truth boundaries, and workflow roles intact; BPDSL staging and Brewprint inventory pointers excluded |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `section_projection` | What this is; Artifact responsibility matrix (6 Design Records artifact rows: ADR, investigation, requirement, work item, task, spec); Deferred material disposition | `## Extracted implementation responsibilities` section entirely (6 table rows pointing to `spec:product.bpdsl.artifact_responsibilities`); `## Related specs` row for `spec:product.bpdsl.artifact_responsibilities` | Core 6-artifact responsibility matrix intact; BPDSL-staged implementation-flow rows excluded per BPDSL exclusion boundary |
| `spec:product.design_records.artifact_model.change_and_investigation_flow` | `section_projection` | Generic investigation to ADR / requirement / work item / task / spec flow; artifact ownership rules; generic workflow rules | `V01-*` source and evidence refs; BPDSL staging wording; concrete DRMCP ownership wording; Brewprint legacy-ID compatibility statements | Generic workflow semantics retained; historical provenance, staging, and implementation-owner context excluded or generalized |
| `spec:product.design_records.artifact_model.traceability_boundary` | `section_projection` | What this is; Traceability and tool boundary; MVP scope; Traceability contract boundary; Related specs; Sources | `## DRMCP boundary` section rows for `spec:drmcp.design_records_mcp.tools.list_records`, `spec:drmcp.design_records_mcp.tools.resolve_reference`, and `spec:drmcp.design_records_mcp.tools.validate_records` | Traceability vs. tool boundary concept and MVP scope intact; DRMCP tool ownership rows excluded |
| `spec:product.design_records.authoring_standards` | `section_projection` | Navigation structure; child topic refs; generic topic summaries | `brewprint design records` wording; DRMCP-dependent and DRMCP-managed framing in summaries | Navigation and child refs retained; app-specific prose generalized because every child source is projected |
| `spec:product.design_records.authoring_standards.writing_standard` | `section_projection` | Generic prose rules; block type rules; AI-output rules | Brewprint-specific framing in `## What this is`; `PRODUCT-INV-SPEC-003`; `PRODUCT-WORK-SPEC-010`; project-specific provenance refs | Generic writing standard retained; source-investigation/work-item provenance excluded from package projection |
| `spec:product.design_records.authoring_standards.agent_authoring_policy` | `section_projection` | `## Rules > Prose compliance`; `## Rules > Namespace usage` | `## Rules > DRMCP retrieval` (TBD section, no current semantic contract); `## Rules > Authoring transaction preference` (TBD section, no current semantic contract); `## Related specs` row for `spec:drmcp.design_records_mcp.tools.overview` | Active behavioral rules for prose compliance and namespace usage intact; DRMCP-dependent TBD sections excluded |
| `spec:product.design_records.authoring_standards.artifact_boundary` | `section_projection` | Artifact selection rules; covered Design Records artifact kinds; adjacent-artifact decision table; generic non-goal boundary | `DRMCP-managed` wording; concrete BPDSL / non-DRMCP framing where host-application specific; references requiring excluded app-local authority | Generic artifact-selection contract retained, including that non-Design-Records artifacts are outside this contract; app-local framing excluded or generalized |
| `spec:product.design_records.authoring_standards.adr_authoring` | `section_projection` | All sections (Non-goals; Rules; Authoring interface requirements); metadata schema table intact | Body line "The parsing grammar is defined by `spec:drmcp.design_records_mcp.schema.metadata_grammar`" in `## Rules > Metadata schema`; `## Related specs` rows for `spec:drmcp.design_records_mcp.schema.metadata_grammar` and `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | ADR ID grammar, file shape, metadata field table, status lifecycle, and authoring rules complete; DRMCP parsing grammar pointer and transaction schema pointer excluded |
| `spec:product.design_records.authoring_standards.requirement_authoring` | `section_projection` | All sections; metadata schema table intact | Body line "The parsing grammar is defined by `spec:drmcp.design_records_mcp.schema.metadata_grammar`" in `## Rules > Metadata schema`; `## Related specs` rows for DRMCP metadata grammar and DRMCP authoring transaction schema | Requirement file shape, metadata fields, status lifecycle, and authoring rules complete |
| `spec:product.design_records.authoring_standards.work_item_authoring` | `section_projection` | All sections; metadata schema table intact | Body line "The parsing grammar is defined by `spec:drmcp.design_records_mcp.schema.metadata_grammar`" in `## Rules > Metadata schema`; `## Related specs` rows for DRMCP metadata grammar and DRMCP authoring transaction schema | Work-item authoring contract complete |
| `spec:product.design_records.authoring_standards.task_authoring` | `section_projection` | All sections; metadata schema table intact | Body line "The parsing grammar is defined by `spec:drmcp.design_records_mcp.schema.metadata_grammar`" in `## Rules > Metadata schema`; `## Related specs` rows for DRMCP metadata grammar and DRMCP authoring transaction schema | Task authoring contract complete |
| `spec:product.design_records.authoring_standards.investigation_authoring` | `section_projection` | All sections; metadata schema table intact | Body line "The parsing grammar is defined by `spec:drmcp.design_records_mcp.schema.metadata_grammar`" in `## Rules > Metadata schema`; `## Related specs` rows for DRMCP metadata grammar and DRMCP authoring transaction schema | Investigation authoring contract complete |
| `spec:product.design_records.authoring_standards.spec_authoring` | `section_projection` | All sections; metadata schema table intact | Body line "The parsing grammar is defined by `spec:drmcp.design_records_mcp.schema.metadata_grammar`" in `## Rules > Metadata schema`; `## Related specs` rows for DRMCP metadata grammar and DRMCP authoring transaction schema | Spec authoring contract complete |


### Historical section-projection audit: 1. Required semantic-area coverage

Each row records one (required area, authoritative source) pair.
Spec areas that cover multiple required areas appear in multiple rows.

| required area | authoritative source ref | exact section or contract area | inclusion mode | reason |
| --- | --- | --- | --- | --- |
| Artifact-kind responsibilities for authoring selection | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | Artifact responsibility matrix (ADR/investigation/requirement/work item/task/spec rows) | `section_projection` | Canonical Design Records artifact responsibility boundary; Extracted implementation responsibilities section excluded (BPDSL staging) |
| Artifact-kind responsibilities for authoring selection | `spec:product.design_records.authoring_standards.artifact_boundary` | Covered artifact kinds; distinguishing adjacent artifacts; decision rules | `section_projection` | Authoring-time projection used at kind-selection time; references canonical matrix as authoritative source |
| Public ID and canonical-ref grammar | `spec:product.design_records.namespace_model.artifact_id_grammar` | Grammar; sequence format; canonical reference rule | `section_projection` | Defines complete public ID form for workflow artifacts; Brewprint V01 compatibility inline note and Related specs rows excluded |
| Public ID and canonical-ref grammar | `spec:product.design_records.spec_format.spec_id_as_ref` | Path-derived canonical spec ref derivation rules; parent reference grammar | `section_projection` | Defines spec canonical ref (`spec:<app>.<path>`) and parent grammar used by traceability and validation |
| App and domain namespace rules | `spec:product.design_records.namespace_model.app_namespaces` | Current contract; rules; ownership axis | `section_projection` | App namespace concept, domain relationship, and cross-app concern rule; Brewprint registry rows excluded |
| App and domain namespace rules | `spec:product.design_records.namespace_model.existing_artifacts` | Current contract; namespace selection table; rules | `section_projection` | Namespace selection rule for new records; Brewprint attribution rows excluded |
| App and domain namespace rules | `spec:product.design_records.namespace_model.subdomain_model` | Definition and representation; write-time advisory semantics | `section_projection` | Subdomain metadata field representation and write-time advisory rule; Brewprint compatibility Related specs row excluded |
| Repository placement mappings | `spec:product.design_records.repository_layout` | Current contract; rules; kind-first structure table | `section_projection` | Kind-first `records/` layout, domain-scoped subdirectory structure, and spec topic-tree placement; BPDSL staging note and Brewprint inventory pointer excluded |
| Repository placement mappings | `spec:product.design_records.repository_layout.record_discovery_paths` | Current contract; per-kind path patterns; rules | `section_projection` | Per-kind path patterns required by authoring for file placement; DRMCP boundary section excluded |
| Metadata fields and author-facing requiredness | `spec:product.design_records.authoring_standards.adr_authoring` | Metadata schema (field / create input / partial update / persisted ADR matrix) | `section_projection` | ADR metadata fields with author-facing requiredness; DRMCP parsing grammar pointer excluded |
| Metadata fields and author-facing requiredness | `spec:product.design_records.authoring_standards.requirement_authoring` | Metadata schema | `section_projection` | Requirement metadata fields with requiredness; DRMCP pointer excluded |
| Metadata fields and author-facing requiredness | `spec:product.design_records.authoring_standards.work_item_authoring` | Metadata schema | `section_projection` | Work-item metadata fields; DRMCP pointer excluded |
| Metadata fields and author-facing requiredness | `spec:product.design_records.authoring_standards.task_authoring` | Metadata schema | `section_projection` | Task metadata fields; DRMCP pointer excluded |
| Metadata fields and author-facing requiredness | `spec:product.design_records.authoring_standards.investigation_authoring` | Metadata schema (required and optional field distinction) | `section_projection` | Investigation metadata fields with required/optional split; DRMCP pointer excluded |
| Metadata fields and author-facing requiredness | `spec:product.design_records.authoring_standards.spec_authoring` | Metadata schema (field / create input / partial update / persisted spec matrix) | `section_projection` | Spec metadata fields with author-facing requiredness; DRMCP pointer excluded |
| Metadata fields and author-facing requiredness | `spec:product.design_records.traceability.metadata_schema` | Spec metadata boundary; workflow relation metadata table; investigation reference metadata | `section_projection` | Traceability boundary for metadata fields; confirms which fields carry canonical relations |
| Persisted document shape | `spec:product.design_records.spec_format.document_shape` | Current contract; H1-adjacent metadata block; required section matrix | `section_projection` | Canonical spec file shape: H1 format, H1-adjacent metadata block, required section matrix by kind |
| Persisted document shape (workflow artifacts) | `spec:product.design_records.authoring_standards.adr_authoring` | File shape; body section schema | `section_projection` | ADR file shape: H1 `# <PUBLIC-ID>: <Title>`, metadata block, required body sections; DRMCP pointer excluded |
| Persisted document shape (workflow artifacts) | `spec:product.design_records.authoring_standards.requirement_authoring` | File shape; body section schema | `section_projection` | Requirement file shape and H2 sections; DRMCP pointer excluded |
| Persisted document shape (workflow artifacts) | `spec:product.design_records.authoring_standards.work_item_authoring` | File shape; body section schema | `section_projection` | Work-item file shape and H2 section schema; DRMCP pointer excluded |
| Persisted document shape (workflow artifacts) | `spec:product.design_records.authoring_standards.task_authoring` | File shape; body section schema | `section_projection` | Task file shape and H2 sections; DRMCP pointer excluded |
| Persisted document shape (workflow artifacts) | `spec:product.design_records.authoring_standards.investigation_authoring` | File shape; body section schema | `section_projection` | Investigation file shape and H2 sections; DRMCP pointer excluded |
| Canonical H1 and H2 rules | `spec:product.design_records.spec_format.document_shape` | H1 format rule; H1-adjacent metadata; required section matrix | `section_projection` | H1 `# <SpecKind>: <Title>` form and per-kind required H2 section matrix |
| Canonical H1 and H2 rules (workflow artifacts) | `spec:product.design_records.authoring_standards.adr_authoring` | File shape (H1 rule) | `section_projection` | H1 `# <PUBLIC-ID>: <Title>` rule; canonical English H2 section headings; DRMCP pointer excluded |
| Canonical H1 and H2 rules (workflow artifacts) | `spec:product.design_records.authoring_standards.requirement_authoring` | File shape (H1 rule); body section schema | `section_projection` | H1 form and canonical H2 section headings for requirements; DRMCP pointer excluded |
| Canonical H1 and H2 rules (workflow artifacts) | `spec:product.design_records.authoring_standards.work_item_authoring` | File shape; body section schema | `section_projection` | H1 form and canonical H2 headings for work items; DRMCP pointer excluded |
| Canonical H1 and H2 rules (workflow artifacts) | `spec:product.design_records.authoring_standards.task_authoring` | File shape; body section schema | `section_projection` | H1 form and canonical H2 headings for tasks; DRMCP pointer excluded |
| Canonical H1 and H2 rules (workflow artifacts) | `spec:product.design_records.authoring_standards.investigation_authoring` | File shape; body section schema | `section_projection` | H1 form and canonical English H2 headings for investigations; DRMCP pointer excluded |
| Lifecycle gates used by structural validation | `spec:product.design_records.spec_format.validation_policy` | Validation rules table; portable severity semantics; projected document-shape and topic/ref validation rules | `section_projection` | Severity policy retained while concrete Work Item owners and current migration/implementation phase tracking are excluded or generalized |
| Lifecycle gates used by structural validation | `spec:product.design_records.spec_format.document_shape` | Required section matrix; Errors table | `section_projection` | Required section matrix per spec kind; Errors table severity |
| Lifecycle gates used by structural validation | `spec:product.design_records.authoring_standards.adr_authoring` | Status lifecycle; kind-specific authoring rules | `section_projection` | ADR status (`proposed` → `accepted` → `superseded`) and transition rules; DRMCP pointer excluded |
| Lifecycle gates used by structural validation | `spec:product.design_records.authoring_standards.requirement_authoring` | Status lifecycle; kind-specific authoring rules | `section_projection` | Requirement status lifecycle; DRMCP pointer excluded |
| Lifecycle gates used by structural validation | `spec:product.design_records.authoring_standards.work_item_authoring` | Status lifecycle; kind-specific authoring rules | `section_projection` | Work-item status lifecycle; DRMCP pointer excluded |
| Lifecycle gates used by structural validation | `spec:product.design_records.authoring_standards.task_authoring` | Status lifecycle; kind-specific authoring rules | `section_projection` | Task status lifecycle; DRMCP pointer excluded |
| Lifecycle gates used by structural validation | `spec:product.design_records.authoring_standards.investigation_authoring` | Status lifecycle; conclusion-readiness rule | `section_projection` | Investigation status; DRMCP pointer excluded |
| Spec kinds and `contract_class` | `spec:product.design_records.spec_format.document_shape` | Spec kind set; contract class; H1-adjacent `contract_class` marker | `section_projection` | Accepted kind set (Overview, Index, Concept, Reference, Contract); `contract_class` values (interface / format) |
| Spec kinds and `contract_class` | `spec:product.design_records.authoring_standards.spec_authoring` | File shape (accepted spec kinds; deferred kinds); contract_class rules | `section_projection` | Author-facing kind selection table and `contract_class` authoring rules; DRMCP pointer excluded |
| Logical spec selector and topic/leaf placement | `spec:product.design_records.spec_format.spec_id_as_ref` | Path-derived canonical spec ref derivation; path-to-ref mapping rules | `section_projection` | Path → canonical `spec:` ref derivation: how directory tree maps to logical selector |
| Logical spec selector and topic/leaf placement | `spec:product.design_records.spec_format.topics_table` | Current contract; rules; validation rules | `section_projection` | `## Topics` table as authoritative parent declaration; defines leaf vs. topic placement |
| Traceability forms | `spec:product.design_records.traceability.semantic_ref` | Current spec ref model; generic active `spec:` ref class; obsolete semantic-prefix assumptions | `section_projection` | Current `spec:` ref class model retained; Brewprint legacy-ID compatibility treatment excluded |
| Traceability forms | `spec:product.design_records.traceability.artifact_refs` | Active reference classes; record ID-as-ref forms; investigation reference boundary; workflow relation identity | `section_projection` | All active canonical reference classes and workflow relation fields; Brewprint compatibility pointers excluded from rows |
| Traceability forms | `spec:product.design_records.traceability.metadata_schema` | Spec metadata boundary; removed front-matter schemas; workflow relation metadata | `section_projection` | Metadata and relation field traceability boundary |
| Traceability forms | `spec:product.design_records.traceability.resolve_and_validation` | Supported canonical inputs; lookup sources; invalid conditions | `section_projection` | Canonical lookup sources and PRODUCT-owned invalid conditions; Brewprint compatibility Related specs row excluded |
| Authoring prose rules | `spec:product.design_records.authoring_standards.writing_standard` | All sections (spec-side rules; AI output rules) | `section_projection` | Sentence style, block type, and AI output rules required at generation time |
| Authoring prose rules | `spec:product.design_records.authoring_standards.agent_authoring_policy` | `## Rules > Prose compliance`; `## Rules > Namespace usage` | `section_projection` | Active prose compliance and namespace usage rules; DRMCP-dependent TBD sections excluded |

**Indices and overviews required for package structural integrity** (canonical parent relationships and navigation entry points):

| source ref | inclusion mode | reason |
| --- | --- | --- |
| `spec:product.design_records` | `section_projection` | Root overview and package entry point; Topics table (all 6 child areas) intact; Related specs rows for excluded areas omitted |
| `spec:product.design_records.authoring_standards` | `section_projection` | Index; navigation entry point; no excluded refs |
| `spec:product.design_records.namespace_model` | `section_projection` | Overview; namespace model concept and authoritative parent; Boundary and Related specs rows for Brewprint/BPDSL areas omitted |
| `spec:product.design_records.repository_layout` | `section_projection` | Overview; record placement contract; BPDSL staging note section and Brewprint inventory pointer omitted |
| `spec:product.design_records.spec_format` | `section_projection` | Index; navigation entry point; follow_up_boundary Topics row and DRMCP Related specs row omitted |
| `spec:product.design_records.traceability` | `section_projection` | Overview; active traceability contract; Brewprint compatibility Sources row omitted |
| `spec:product.design_records.artifact_model` | `section_projection` | Overview; artifact classes and source-of-truth boundaries; BPDSL staging and Brewprint inventory rows omitted |


### Historical section-projection audit: 2. Transitive dependency closure

#### Historical section-projection audit: 2a. Internal dependency edges

Every row records a unique canonical `spec:` dependency edge where both source and dependency are within the six PRODUCT Design Records areas. Each disposition matches the dependency target's Section 0 inclusion classification, except edges whose referenced content is omitted and whose target is excluded.

| source ref | dependency ref | source location | disposition | reason |
| --- | --- | --- | --- | --- |
| `spec:product.design_records` | `spec:product.design_records.namespace_model` | `## Topics` | `section_projection` | Child topic; both in included set |
| `spec:product.design_records` | `spec:product.design_records.authoring_standards` | `## Topics` | `section_projection` | Child topic; both in included set |
| `spec:product.design_records` | `spec:product.design_records.spec_format` | `## Topics` | `section_projection` | Child topic; both in included set |
| `spec:product.design_records` | `spec:product.design_records.repository_layout` | `## Topics` | `section_projection` | Child topic; both in included set |
| `spec:product.design_records` | `spec:product.design_records.traceability` | `## Topics` | `section_projection` | Child topic; both in included set |
| `spec:product.design_records` | `spec:product.design_records.artifact_model` | `## Topics` | `section_projection` | Child topic; both in included set |
| `spec:product.design_records.namespace_model` | `spec:product.design_records.namespace_model.artifact_id_grammar` | `## Boundary` + `## Topics` | `section_projection` | Boundary owner pointer + child topic; both included |
| `spec:product.design_records.namespace_model` | `spec:product.design_records.namespace_model.existing_artifacts` | `## Boundary` + `## Topics` | `section_projection` | Boundary owner pointer + child topic; both included |
| `spec:product.design_records.namespace_model` | `spec:product.design_records.namespace_model.subdomain_model` | `## Boundary` + `## Topics` | `section_projection` | Boundary owner pointer + child topic; both included |
| `spec:product.design_records.namespace_model` | `spec:product.design_records.namespace_model.app_namespaces` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.namespace_model` | `spec:product.design_records.repository_layout` | `## Boundary` | `section_projection` | Boundary cross-ref; both included |
| `spec:product.design_records.namespace_model.app_namespaces` | `spec:product.design_records.namespace_model` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | `spec:product.design_records.namespace_model` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | `spec:product.design_records.namespace_model.subdomain_model` | `## Related specs` + inline body note | `section_projection` | Sibling ref (subdomains not in ID segments); both included |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | `spec:product.design_records.traceability.artifact_refs` | `## Related specs` | `section_projection` | Cross-area; how canonical IDs function as semantic refs; both included |
| `spec:product.design_records.namespace_model.existing_artifacts` | `spec:product.design_records.namespace_model` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.namespace_model.existing_artifacts` | `spec:product.design_records.namespace_model.artifact_id_grammar` | `## Related specs` + `## Boundary` | `section_projection` | Sibling ref; canonical grammar source; both included |
| `spec:product.design_records.namespace_model.subdomain_model` | `spec:product.design_records.namespace_model` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.repository_layout` | `spec:product.design_records.namespace_model` | `## Related specs` | `section_projection` | Cross-area; namespace semantics used by layout; both included |
| `spec:product.design_records.repository_layout` | `spec:product.design_records.spec_format` | `## Related specs` | `section_projection` | Cross-area; path-derived identity used by format; both included |
| `spec:product.design_records.repository_layout` | `spec:product.design_records.repository_layout.record_discovery_paths` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.repository_layout.record_discovery_paths` | `spec:product.design_records.repository_layout` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.spec_format` | `spec:product.design_records.spec_format.overview` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.spec_format` | `spec:product.design_records.spec_format.document_shape` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.spec_format` | `spec:product.design_records.spec_format.topics_table` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.spec_format` | `spec:product.design_records.spec_format.spec_id_as_ref` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.spec_format` | `spec:product.design_records.spec_format.validation_policy` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.spec_format` | `spec:product.design_records.spec_format.follow_up_boundary` | `## Topics` | `excluded` | Child topic row omitted in projection; follow_up_boundary is excluded from package |
| `spec:product.design_records.spec_format` | `spec:product.design_records.namespace_model` | `## Related specs` | `section_projection` | Cross-area; namespace model used by path-derived IDs; both included |
| `spec:product.design_records.spec_format` | `spec:product.design_records.repository_layout` | `## Related specs` | `section_projection` | Cross-area; repository layout used by format; both included |
| `spec:product.design_records.spec_format.overview` | `spec:product.design_records.spec_format` | `## Related specs` | `section_projection` | Parent index pointer; both included |
| `spec:product.design_records.spec_format.overview` | `spec:product.design_records.namespace_model` | `## Related specs` | `section_projection` | Cross-area; both included |
| `spec:product.design_records.spec_format.overview` | `spec:product.design_records.repository_layout` | `## Related specs` | `section_projection` | Cross-area; both included |
| `spec:product.design_records.spec_format.overview` | `spec:product.design_records.spec_format.document_shape` | `## Topic map` | `section_projection` | Sibling; both included |
| `spec:product.design_records.spec_format.overview` | `spec:product.design_records.spec_format.topics_table` | `## Topic map` | `section_projection` | Sibling; both included |
| `spec:product.design_records.spec_format.overview` | `spec:product.design_records.spec_format.spec_id_as_ref` | `## Topic map` | `section_projection` | Sibling; both included |
| `spec:product.design_records.spec_format.overview` | `spec:product.design_records.spec_format.validation_policy` | `## Topic map` | `section_projection` | Sibling; both included |
| `spec:product.design_records.spec_format.overview` | `spec:product.design_records.spec_format.follow_up_boundary` | `## Topic map` | `excluded` | Topic map row omitted in projection; follow_up_boundary is excluded from package |
| `spec:product.design_records.spec_format.document_shape` | `spec:product.design_records.spec_format` | `## Related specs` | `section_projection` | Parent index pointer; both included |
| `spec:product.design_records.spec_format.document_shape` | `spec:product.design_records.spec_format.topics_table` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.spec_format.topics_table` | `spec:product.design_records.spec_format` | `## Related specs` | `section_projection` | Parent index pointer; both included |
| `spec:product.design_records.spec_format.topics_table` | `spec:product.design_records.spec_format.document_shape` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.spec_format.spec_id_as_ref` | `spec:product.design_records.spec_format` | `## Related specs` | `section_projection` | Parent index pointer; both included |
| `spec:product.design_records.spec_format.spec_id_as_ref` | `spec:product.design_records.spec_format.topics_table` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.spec_format.spec_id_as_ref` | `spec:product.design_records.spec_format.validation_policy` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.spec_format.validation_policy` | `spec:product.design_records.spec_format` | `## Related specs` | `section_projection` | Parent index pointer; both included |
| `spec:product.design_records.spec_format.validation_policy` | `spec:product.design_records.spec_format.document_shape` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.spec_format.validation_policy` | `spec:product.design_records.spec_format.topics_table` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.spec_format.validation_policy` | `spec:product.design_records.spec_format.spec_id_as_ref` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.traceability` | `spec:product.design_records.traceability.semantic_ref` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.traceability` | `spec:product.design_records.traceability.artifact_refs` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.traceability` | `spec:product.design_records.traceability.metadata_schema` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.traceability` | `spec:product.design_records.traceability.resolve_and_validation` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.traceability` | `spec:product.design_records.spec_format.spec_id_as_ref` | `## Traceability model` + `## Sources` | `section_projection` | Cross-area; path-derived canonical spec identity; both included |
| `spec:product.design_records.traceability` | `spec:product.design_records.namespace_model.artifact_id_grammar` | `## Traceability model` | `section_projection` | Cross-area; artifact ID grammar pointer in traceability model; both included |
| `spec:product.design_records.traceability.semantic_ref` | `spec:product.design_records.spec_format.spec_id_as_ref` | `## Related specs` | `section_projection` | Cross-area; canonical spec identity contract; both included |
| `spec:product.design_records.traceability.semantic_ref` | `spec:product.design_records.traceability.artifact_refs` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.traceability.semantic_ref` | `spec:product.design_records.traceability.metadata_schema` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.traceability.artifact_refs` | `spec:product.design_records.spec_format.spec_id_as_ref` | `## Active reference classes` + `## Sources` | `section_projection` | Cross-area; canonical spec ref derivation; both included |
| `spec:product.design_records.traceability.artifact_refs` | `spec:product.design_records.namespace_model.artifact_id_grammar` | `## Sources` | `section_projection` | Cross-area; app-aware public ID grammar; both included |
| `spec:product.design_records.traceability.metadata_schema` | `spec:product.design_records.spec_format.spec_id_as_ref` | `## Related specs` | `section_projection` | Cross-area; canonical spec identity owner; both included |
| `spec:product.design_records.traceability.metadata_schema` | `spec:product.design_records.spec_format.document_shape` | `## Related specs` | `section_projection` | Cross-area; visible metadata and section-shape owner; both included |
| `spec:product.design_records.traceability.metadata_schema` | `spec:product.design_records.spec_format.topics_table` | `## Related specs` | `section_projection` | Cross-area; authoritative child-topic relationship owner; both included |
| `spec:product.design_records.traceability.metadata_schema` | `spec:product.design_records.traceability.artifact_refs` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.traceability.metadata_schema` | `spec:product.design_records.traceability.resolve_and_validation` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.traceability.resolve_and_validation` | `spec:product.design_records.traceability.artifact_refs` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.traceability.resolve_and_validation` | `spec:product.design_records.traceability.metadata_schema` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.traceability.resolve_and_validation` | `spec:product.design_records.spec_format.spec_id_as_ref` | `## Related specs` | `section_projection` | Cross-area; spec lookup source and path-derived identity; both included |
| `spec:product.design_records.artifact_model` | `spec:product.design_records.repository_layout` | `## Source of truth` + `## Navigation` + `## Related specs` | `section_projection` | Cross-area; app-independent record placement; both included |
| `spec:product.design_records.artifact_model` | `spec:product.design_records.traceability` | `## Source of truth` + `## Related specs` | `section_projection` | Cross-area; canonical reference and validation semantics; both included |
| `spec:product.design_records.artifact_model` | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Disposition` + `## Navigation` + `## Topics` | `section_projection` | Child topic + navigation; both included |
| `spec:product.design_records.artifact_model` | `spec:product.design_records.artifact_model.change_and_investigation_flow` | `## Navigation` + `## Topics` | `section_projection` | Child topic + navigation; both included |
| `spec:product.design_records.artifact_model` | `spec:product.design_records.artifact_model.traceability_boundary` | `## Navigation` + `## Topics` | `section_projection` | Child topic + navigation; both included |
| `spec:product.design_records.artifact_model` | `spec:product.design_records.authoring_standards` | `## Navigation` | `section_projection` | Cross-area; authoring guidance entrypoints; both included |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `spec:product.design_records.artifact_model` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.artifact_model.change_and_investigation_flow` | `spec:product.design_records.artifact_model` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.artifact_model.traceability_boundary` | `spec:product.design_records.artifact_model` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.artifact_model.traceability_boundary` | `spec:product.design_records.traceability` | `## Related specs` | `section_projection` | Cross-area; canonical reference and validation semantics; both included |
| `spec:product.design_records.authoring_standards` | `spec:product.design_records.authoring_standards.writing_standard` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.authoring_standards` | `spec:product.design_records.authoring_standards.agent_authoring_policy` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.authoring_standards` | `spec:product.design_records.authoring_standards.artifact_boundary` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.authoring_standards` | `spec:product.design_records.authoring_standards.adr_authoring` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.authoring_standards` | `spec:product.design_records.authoring_standards.requirement_authoring` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.authoring_standards` | `spec:product.design_records.authoring_standards.work_item_authoring` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.authoring_standards` | `spec:product.design_records.authoring_standards.task_authoring` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.authoring_standards` | `spec:product.design_records.authoring_standards.investigation_authoring` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.authoring_standards` | `spec:product.design_records.authoring_standards.spec_authoring` | `## Topics` | `section_projection` | Child topic; both included |
| `spec:product.design_records.authoring_standards.writing_standard` | `spec:product.design_records.authoring_standards` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.authoring_standards.writing_standard` | `spec:product.design_records.spec_format.document_shape` | `## Related specs` | `section_projection` | Cross-area; section and heading format rules; both included |
| `spec:product.design_records.authoring_standards.agent_authoring_policy` | `spec:product.design_records.authoring_standards` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.authoring_standards.agent_authoring_policy` | `spec:product.design_records.authoring_standards.writing_standard` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.authoring_standards.artifact_boundary` | `spec:product.design_records.authoring_standards` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.authoring_standards.artifact_boundary` | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Related specs` | `section_projection` | Cross-area; canonical artifact ownership; both included |
| `spec:product.design_records.authoring_standards.adr_authoring` | `spec:product.design_records.authoring_standards` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.authoring_standards.adr_authoring` | `spec:product.design_records.authoring_standards.writing_standard` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.authoring_standards.adr_authoring` | `spec:product.design_records.authoring_standards.artifact_boundary` | `## Related specs` + `## Rules > Kind-specific` body | `section_projection` | Sibling; cross-artifact selection; both included |
| `spec:product.design_records.authoring_standards.adr_authoring` | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Related specs` + `## Rules > Kind-specific` body | `section_projection` | Cross-area; canonical artifact ownership; both included |
| `spec:product.design_records.authoring_standards.adr_authoring` | `spec:product.design_records.namespace_model.artifact_id_grammar` | `## Related specs` + `## Rules > ID grammar` body | `section_projection` | Cross-area; ADR ID grammar; both included |
| `spec:product.design_records.authoring_standards.adr_authoring` | `spec:product.design_records.repository_layout.record_discovery_paths` | `## Related specs` + `## Rules > File path layout` body | `section_projection` | Cross-area; ADR discovery path rules; both included |
| `spec:product.design_records.authoring_standards.requirement_authoring` | `spec:product.design_records.authoring_standards` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.authoring_standards.requirement_authoring` | `spec:product.design_records.authoring_standards.writing_standard` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.authoring_standards.requirement_authoring` | `spec:product.design_records.authoring_standards.artifact_boundary` | `## Related specs` + body | `section_projection` | Sibling; cross-artifact selection; both included |
| `spec:product.design_records.authoring_standards.requirement_authoring` | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Related specs` + body | `section_projection` | Cross-area; canonical artifact ownership; both included |
| `spec:product.design_records.authoring_standards.requirement_authoring` | `spec:product.design_records.namespace_model.artifact_id_grammar` | `## Related specs` + body | `section_projection` | Cross-area; requirement ID grammar; both included |
| `spec:product.design_records.authoring_standards.requirement_authoring` | `spec:product.design_records.repository_layout.record_discovery_paths` | `## Related specs` + body | `section_projection` | Cross-area; requirement discovery path rules; both included |
| `spec:product.design_records.authoring_standards.work_item_authoring` | `spec:product.design_records.authoring_standards` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.authoring_standards.work_item_authoring` | `spec:product.design_records.authoring_standards.writing_standard` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.authoring_standards.work_item_authoring` | `spec:product.design_records.authoring_standards.artifact_boundary` | `## Related specs` + body | `section_projection` | Sibling; cross-artifact selection; both included |
| `spec:product.design_records.authoring_standards.work_item_authoring` | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Related specs` + body | `section_projection` | Cross-area; canonical artifact ownership; both included |
| `spec:product.design_records.authoring_standards.work_item_authoring` | `spec:product.design_records.namespace_model.artifact_id_grammar` | `## Related specs` + body | `section_projection` | Cross-area; work-item ID grammar; both included |
| `spec:product.design_records.authoring_standards.work_item_authoring` | `spec:product.design_records.repository_layout.record_discovery_paths` | `## Related specs` + body | `section_projection` | Cross-area; work-item discovery path rules; both included |
| `spec:product.design_records.authoring_standards.task_authoring` | `spec:product.design_records.authoring_standards` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.authoring_standards.task_authoring` | `spec:product.design_records.authoring_standards.writing_standard` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.authoring_standards.task_authoring` | `spec:product.design_records.authoring_standards.artifact_boundary` | `## Related specs` + body | `section_projection` | Sibling; cross-artifact selection; both included |
| `spec:product.design_records.authoring_standards.task_authoring` | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Related specs` + body | `section_projection` | Cross-area; canonical artifact ownership; both included |
| `spec:product.design_records.authoring_standards.task_authoring` | `spec:product.design_records.namespace_model.artifact_id_grammar` | `## Related specs` + body | `section_projection` | Cross-area; task ID grammar; both included |
| `spec:product.design_records.authoring_standards.task_authoring` | `spec:product.design_records.repository_layout.record_discovery_paths` | `## Related specs` + body | `section_projection` | Cross-area; task discovery path rules; both included |
| `spec:product.design_records.authoring_standards.investigation_authoring` | `spec:product.design_records.authoring_standards` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.authoring_standards.investigation_authoring` | `spec:product.design_records.authoring_standards.writing_standard` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.authoring_standards.investigation_authoring` | `spec:product.design_records.authoring_standards.artifact_boundary` | `## Related specs` + body | `section_projection` | Sibling; cross-artifact selection; both included |
| `spec:product.design_records.authoring_standards.investigation_authoring` | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Related specs` + body | `section_projection` | Cross-area; canonical artifact ownership; both included |
| `spec:product.design_records.authoring_standards.investigation_authoring` | `spec:product.design_records.namespace_model.artifact_id_grammar` | `## Related specs` + body | `section_projection` | Cross-area; investigation ID grammar; both included |
| `spec:product.design_records.authoring_standards.investigation_authoring` | `spec:product.design_records.repository_layout.record_discovery_paths` | `## Related specs` + body | `section_projection` | Cross-area; investigation discovery path rules; both included |
| `spec:product.design_records.authoring_standards.spec_authoring` | `spec:product.design_records.authoring_standards` | `## Related specs` | `section_projection` | Parent pointer; both included |
| `spec:product.design_records.authoring_standards.spec_authoring` | `spec:product.design_records.authoring_standards.writing_standard` | `## Related specs` | `section_projection` | Sibling; both included |
| `spec:product.design_records.authoring_standards.spec_authoring` | `spec:product.design_records.authoring_standards.artifact_boundary` | `## Related specs` | `section_projection` | Sibling; cross-artifact selection; both included |
| `spec:product.design_records.authoring_standards.spec_authoring` | `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `## Related specs` | `section_projection` | Cross-area; canonical artifact ownership; both included |
| `spec:product.design_records.authoring_standards.spec_authoring` | `spec:product.design_records.spec_format` | `## Related specs` | `section_projection` | Cross-area; normative spec format contract; both included |
| `spec:product.design_records.authoring_standards.spec_authoring` | `spec:product.design_records.spec_format.spec_id_as_ref` | `## Related specs` + body | `section_projection` | Cross-area; path-derived canonical ref rules; both included |
| `spec:product.design_records.authoring_standards.spec_authoring` | `spec:product.design_records.spec_format.document_shape` | `## Related specs` + body | `section_projection` | Cross-area; H1 format, metadata markers, kind set, section matrix; both included |
| `spec:product.design_records.authoring_standards.spec_authoring` | `spec:product.design_records.spec_format.topics_table` | `## Related specs` + body | `section_projection` | Cross-area; Topics table contract and validation rules; both included |
| `spec:product.design_records.authoring_standards.spec_authoring` | `spec:product.design_records.spec_format.validation_policy` | `## Related specs` + body | `section_projection` | Cross-area; severity rules during migration phase; both included |



#### Historical section-projection audit: 2b. External dependency edges

Internal specs in the included set that point to specs outside the six PRODUCT design-records areas. Disposition describes how the dependency target is included in the package; all external targets are excluded.

| source ref | dependency ref | source location | disposition | reason |
| --- | --- | --- | --- | --- |
| `spec:product.design_records` | `spec:product` | `## Related specs` | `excluded` | PRODUCT root placement router; host-PRODUCT-specific; no generic Design Records semantics beyond what is restated in `spec:product.design_records` |
| `spec:product.design_records` | `spec:product.brewprint` | `## Related specs` | `excluded` | Brewprint profile pointer; Brewprint-specific; excluded per host-independent content boundary |
| `spec:product.design_records` | `spec:product.bpdsl` | `## Related specs` | `excluded` | BPDSL temporary staging area; BPDSL runtime rules outside first-package capability boundary |
| `spec:product.design_records.namespace_model` | `spec:product.brewprint.namespaces` | `## Boundary` + body + `## Related specs` | `excluded` | Current Brewprint app/domain registry; host-specific profile, not generic namespace semantics |
| `spec:product.design_records.namespace_model` | `spec:product.brewprint.compatibility` | `## Boundary` + body + `## Related specs` | `excluded` | Brewprint V01 compatibility and migration state; host-specific |
| `spec:product.design_records.namespace_model.app_namespaces` | `spec:product.brewprint.namespaces.app_namespaces` | `## Boundary` | `excluded` | Brewprint current app namespace list; Brewprint registry fact |
| `spec:product.design_records.namespace_model.app_namespaces` | `spec:product.brewprint.namespaces.domain_catalog` | `## Boundary` | `excluded` | Brewprint current domain catalog; host-specific registry fact |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | `spec:product.brewprint.compatibility.legacy_id_compatibility` | `### REQ / WORK / INV / ADR` + `## Related specs` | `excluded` | Brewprint V01 issued-ID compatibility policy; host-specific migration state |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | `spec:product.brewprint.compatibility.existing_artifacts` | `## Related specs` | `excluded` | Attribution policy for existing Brewprint artifacts; host-specific compatibility |
| `spec:product.design_records.namespace_model.existing_artifacts` | `spec:product.brewprint.namespaces.domain_catalog` | `## Boundary` | `excluded` | Brewprint current domain catalog; host-specific registry fact |
| `spec:product.design_records.namespace_model.existing_artifacts` | `spec:product.brewprint.compatibility.existing_artifacts` | Body + `## Boundary` + `## Related specs` | `excluded` | Brewprint V01 historical ownership and effective attribution; host-specific |
| `spec:product.design_records.namespace_model.subdomain_model` | `spec:product.brewprint.compatibility.existing_artifacts` | `## Related specs` | `excluded` | Brewprint V01 attribution context; host-specific |
| `spec:product.design_records.repository_layout` | `spec:product.brewprint.layout` | `## What this is` + `## Related specs` | `excluded` | Current Brewprint repository inventory; host-specific physical-record inventory |
| `spec:product.design_records.repository_layout` | `spec:product.bpdsl.repository_implementation_flow` | `## BPDSL staging note` + `## Related specs` | `excluded` | Temporary BPDSL staging for removed DSL/src implementation-flow material; BPDSL runtime rules |
| `spec:product.design_records.repository_layout.record_discovery_paths` | `spec:drmcp.design_records_mcp.namespace_scanning` | `## DRMCP boundary` | `excluded` | DRMCP namespace-prefix derivation behavior; DRMCP parser/loader implementation |
| `spec:product.design_records.repository_layout.record_discovery_paths` | `spec:drmcp.design_records_mcp.schema.discovery` | `## DRMCP boundary` | `excluded` | DRMCP index inclusion schema; DRMCP implementation detail |
| `spec:product.design_records.spec_format` | `spec:drmcp.design_records_mcp.overview` | `## Related specs` | `excluded` | DRMCP spec discovery and validation scope; DRMCP implementation boundary |
| `spec:product.design_records.traceability` | `spec:product.brewprint.compatibility` | `## Sources` | `excluded` | Legacy issued-ID compatibility pointer; host-specific |
| `spec:product.design_records.traceability.artifact_refs` | `spec:product.brewprint.compatibility` | `## Active reference classes` + `## Record ID-as-ref` + `## Sources` | `excluded` | Legacy ID compatibility ownership pointer; host-specific |
| `spec:product.design_records.traceability.resolve_and_validation` | `spec:product.brewprint.compatibility` | `## Related specs` | `excluded` | Legacy issued-ID compatibility inputs; host-specific |
| `spec:product.design_records.artifact_model` | `spec:product.brewprint.layout` | `## Source of truth` + body + `## Related specs` | `excluded` | Current Brewprint repository inventory pointer; host-specific |
| `spec:product.design_records.artifact_model` | `spec:product.bpdsl` | `## Navigation` + body + `## Related specs` | `excluded` | BPDSL temporary staging pointer; BPDSL runtime rules outside first-package boundary |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | `spec:product.bpdsl.artifact_responsibilities` | `## Extracted implementation responsibilities` + `## Related specs` | `excluded` | Temporary BPDSL staging for extracted BPDSL artifact rows; BPDSL runtime rules |
| `spec:product.design_records.artifact_model.traceability_boundary` | `spec:drmcp.design_records_mcp.tools.list_records` | `## DRMCP boundary` | `excluded` | DRMCP tool implementation; DRMCP parser/loader/index behavior |
| `spec:product.design_records.artifact_model.traceability_boundary` | `spec:drmcp.design_records_mcp.tools.resolve_reference` | `## DRMCP boundary` | `excluded` | DRMCP tool implementation |
| `spec:product.design_records.artifact_model.traceability_boundary` | `spec:drmcp.design_records_mcp.tools.validate_records` | `## DRMCP boundary` | `excluded` | DRMCP tool implementation |
| `spec:product.design_records.authoring_standards.adr_authoring` | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | `## Rules > Metadata schema` + `## Related specs` | `excluded` | DRMCP parsing grammar; DRMCP implementation contract; not a PRODUCT-owned semantic contract |
| `spec:product.design_records.authoring_standards.adr_authoring` | `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | `## Related specs` | `excluded` | DRMCP authoring transaction schema; DRMCP request/response contract |
| `spec:product.design_records.authoring_standards.requirement_authoring` | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | `## Rules > Metadata schema` + `## Related specs` | `excluded` | DRMCP parsing grammar; DRMCP implementation contract; not a PRODUCT-owned semantic contract |
| `spec:product.design_records.authoring_standards.requirement_authoring` | `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | `## Related specs` | `excluded` | DRMCP authoring transaction schema; DRMCP request/response contract |
| `spec:product.design_records.authoring_standards.work_item_authoring` | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | `## Rules > Metadata schema` + `## Related specs` | `excluded` | DRMCP parsing grammar; DRMCP implementation contract; not a PRODUCT-owned semantic contract |
| `spec:product.design_records.authoring_standards.work_item_authoring` | `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | `## Related specs` | `excluded` | DRMCP authoring transaction schema; DRMCP request/response contract |
| `spec:product.design_records.authoring_standards.task_authoring` | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | `## Rules > Metadata schema` + `## Related specs` | `excluded` | DRMCP parsing grammar; DRMCP implementation contract; not a PRODUCT-owned semantic contract |
| `spec:product.design_records.authoring_standards.task_authoring` | `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | `## Related specs` | `excluded` | DRMCP authoring transaction schema; DRMCP request/response contract |
| `spec:product.design_records.authoring_standards.investigation_authoring` | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | `## Rules > Metadata schema` + `## Related specs` | `excluded` | DRMCP parsing grammar; DRMCP implementation contract; not a PRODUCT-owned semantic contract |
| `spec:product.design_records.authoring_standards.investigation_authoring` | `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | `## Related specs` | `excluded` | DRMCP authoring transaction schema; DRMCP request/response contract |
| `spec:product.design_records.authoring_standards.spec_authoring` | `spec:drmcp.design_records_mcp.schema.metadata_grammar` | `## Rules > Metadata schema` + `## Related specs` | `excluded` | DRMCP parsing grammar; DRMCP implementation contract; not a PRODUCT-owned semantic contract |
| `spec:product.design_records.authoring_standards.spec_authoring` | `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | `## Related specs` | `excluded` | DRMCP authoring transaction schema; DRMCP request/response contract |
| `spec:product.design_records.authoring_standards.agent_authoring_policy` | `spec:drmcp.design_records_mcp.tools.overview` | `## Related specs` | `excluded` | DRMCP tool API contract; TBD section excluded from projected content |

### Historical section-projection audit: 3. Explicit exclusions

#### Brewprint namespace and domain registry inventory

| excluded source ref or contract area | exclusion reason |
| --- | --- |
| `spec:product.brewprint.namespaces` | Current Brewprint app and domain namespace profile; owned by `spec:product.brewprint`; contains profile-specific registry facts, not generic rules |
| `spec:product.brewprint.namespaces.app_namespaces` | Brewprint current app namespace list; host-specific registry |
| `spec:product.brewprint.namespaces.domain_catalog` | Brewprint domain catalog; host-specific registry |

#### Brewprint V01 compatibility and migration state

| excluded source ref or contract area | exclusion reason |
| --- | --- |
| `spec:product.brewprint.compatibility` | V01 compatibility policy, issued-ID retention, migration state; project-specific history not needed by portable authoring consumers |
| `spec:product.brewprint.compatibility.legacy_id_compatibility` | V01 issued-ID compatibility policy; Brewprint-specific |
| `spec:product.brewprint.compatibility.existing_artifacts` | Brewprint V01 historical ownership and effective attribution; Brewprint-specific |

#### Current Brewprint physical-record inventory

| excluded source ref or contract area | exclusion reason |
| --- | --- |
| `spec:product.brewprint.layout` | Current Brewprint repository inventory and observed app namespace states; host-specific; PRODUCT-REQ-SPEC-003 §Host-independent content boundary explicitly requires exclusion |

#### Host repository root assumptions

| excluded source ref or contract area | exclusion reason |
| --- | --- |
| `spec:product` | PRODUCT root placement router; app-local router depends on host repository layout; no generic Design Records semantics beyond what is restated in `spec:product.design_records` |
| `spec:product.brewprint` | Brewprint-specific profile area; all content is host-specific instantiation |

#### DRMCP request/response schemas

| excluded source ref or contract area | exclusion reason |
| --- | --- |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | DRMCP parsing grammar for Design Records metadata; DRMCP implementation contract; owned by DRMCP, not PRODUCT |
| `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | DRMCP authoring transaction request/response schema; DRMCP implementation contract |
| `spec:drmcp.design_records_mcp.schema.discovery` | DRMCP index inclusion schema; DRMCP implementation |
| `spec:drmcp.design_records_mcp.overview` | DRMCP spec discovery and validation scope pointer; DRMCP implementation boundary |

#### DRMCP parser, loader, index, and diagnostic behavior

| excluded source ref or contract area | exclusion reason |
| --- | --- |
| `spec:drmcp.design_records_mcp.namespace_scanning` | DRMCP namespace-prefix derivation from records_root; DRMCP implementation |
| `spec:drmcp.design_records_mcp.tools.list_records` | DRMCP indexing tool; DRMCP implementation |
| `spec:drmcp.design_records_mcp.tools.resolve_reference` | DRMCP resolve tool; DRMCP implementation |
| `spec:drmcp.design_records_mcp.tools.validate_records` | DRMCP validation tool; DRMCP implementation |
| `spec:drmcp.design_records_mcp.tools.overview` | DRMCP tool API overview; DRMCP implementation |
| DRMCP-dependent TBD sections in `spec:product.design_records.authoring_standards.agent_authoring_policy` | DRMCP retrieval rules and authoring transaction preference are TBD pending DRMCP operationalization; no current semantic contract to distribute |

#### BPDSL language, generation, rendering, resolver, and runtime rules

| excluded source ref or contract area | exclusion reason |
| --- | --- |
| `spec:product.bpdsl` | Temporary BPDSL staging area; not canonical BPDSL ownership; no normative dependency from Design Records to BPDSL (PRODUCT-ADR-SPEC-001 §Dependency direction) |
| `spec:product.bpdsl.repository_implementation_flow` | Temporary preservation of removed `dsl/`/`src/` implementation-flow material; BPDSL runtime rules |
| `spec:product.bpdsl.artifact_responsibilities` | Temporary preservation of BPDSL artifact responsibility rows extracted from the artifact model; BPDSL runtime rules |

#### PRODUCT-internal project tracking (additional exclusion within the six areas)

| excluded source ref or contract area | exclusion reason |
| --- | --- |
| `spec:product.design_records.spec_format.follow_up_boundary` | Maps internal Brewprint PRODUCT-WORK-SPEC-* and DRMCP-WORK-SPEC-* work items to follow-up ownership areas; PRODUCT-internal project management context; does not carry semantic contracts needed by authoring consumers; Brewprint-specific project tracking |

#### Excluded canonical record refs

These are canonical Design Records record refs removed from projected content. They are not `spec:` dependency edges and are not counted as external `spec:` dependency refs.

| source ref | source location | removed record ref or family | reason |
| --- | --- | --- | --- |
| `spec:product.design_records.spec_format.document_shape` | `## Validation rules` | `PRODUCT-WORK-SPEC-003` | Authoring examples source ownership; project follow-up context excluded from generic document-shape projection |
| `spec:product.design_records.spec_format.topics_table` | `## Validation rules` | `PRODUCT-WORK-SPEC-002` | Cross-parent and row-level exception follow-up ownership excluded from generic topics-table projection |
| `spec:product.design_records.spec_format.spec_id_as_ref` | `## What this is`; `## Ref classes`; `## Boundary` | `PRODUCT-WORK-SPEC-002` | Compatibility follow-up ownership excluded while retaining generic compatibility boundary |
| `spec:product.design_records.spec_format.spec_id_as_ref` | `## Boundary` | `PRODUCT-WORK-SPEC-005` | Migration execution context excluded from generic path-derived identity projection |
| `spec:product.design_records.artifact_model.change_and_investigation_flow` | `## Deferred implementation tracking disposition` | `V01-INV-DOCS-003` | Historical evidence owner excluded from portable workflow projection |
| `spec:product.design_records.artifact_model.change_and_investigation_flow` | `## Deferred implementation tracking disposition` | `V01-ADR-088` | Historical evidence owner excluded from portable workflow projection |
| `spec:product.design_records.artifact_model.change_and_investigation_flow` | `## Sources` | `V01-ADR-081` | Historical provenance source excluded from portable workflow projection |
| `spec:product.design_records.artifact_model.change_and_investigation_flow` | `## Sources` | `V01-ADR-083` | Historical provenance source excluded from portable workflow projection |
| `spec:product.design_records.artifact_model.change_and_investigation_flow` | `## Sources` | `V01-ADR-085` | Historical provenance source excluded from portable workflow projection |
| `spec:product.design_records.artifact_model.change_and_investigation_flow` | `## Sources` | `V01-ADR-086` | Historical provenance source excluded from portable workflow projection |
| `spec:product.design_records.artifact_model.change_and_investigation_flow` | `## Sources` | `V01-ADR-091` | Historical provenance source excluded from portable workflow projection |
| `spec:product.design_records.authoring_standards.writing_standard` | `## Related specs` | `PRODUCT-INV-SPEC-003` | Source investigation provenance excluded from generic writing-standard projection |
| `spec:product.design_records.authoring_standards.writing_standard` | `## Related specs` | `PRODUCT-WORK-SPEC-010` | Source work-item provenance excluded from generic writing-standard projection |

#### Excluded non-edge prose, sections, or rows without canonical refs

These exclusions are not dependency edges because they do not target a canonical `spec:` ref or canonical record ref.

| source ref | source location | excluded prose, section, or row | reason |
| --- | --- | --- | --- |
| `spec:product.design_records.spec_format.validation_policy` | `## Rules` | DRMCP-owned implementation target and temporary PRODUCT tooling paragraph | Current DRMCP implementation-phase and PRODUCT temporary tooling ownership are project execution context, not portable validation semantics |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for `missing real ATX H1` | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for `multiple real ATX H1 headings` | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for `H1 does not match # <SpecKind>: <Title>` | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for missing H1-adjacent metadata | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for missing or invalid `contract_class` | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for H1-adjacent `id` mismatch | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for YAML front matter | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for hidden front matter `depends_on` / source refs | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for hidden front matter topic refs | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for front matter `design_record.kind` | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for missing `## What this is` | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for invalid accepted spec kind | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for missing required section by kind | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for invalid `## Topics` table columns | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for unresolved child `ref` | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for duplicate parent declaration | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for parent grammar violation | Concrete DRMCP Work Item owner excluded; validation condition and severity retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Validation rules` | owner cell for topic cycle | Concrete DRMCP Work Item owner excluded; deferred graph-validation condition retained |
| `spec:product.design_records.spec_format.validation_policy` | `## Errors` | Existing unmigrated spec migration-warning row | Current migration-phase tracking excluded or generalized |
| `spec:product.design_records.spec_format.validation_policy` | `## Errors` | Inventory, migration, or transient working state mismatch row | Current migration/transient-state tracking excluded or generalized |
| `spec:product.design_records.spec_format.validation_policy` | `## Errors` | Temporary PRODUCT tooling patches current DRMCP implementation row | Temporary PRODUCT tooling ownership and current DRMCP implementation context excluded |
| `spec:product.design_records.traceability.semantic_ref` | `## Ref classes` | Legacy issued ID row | Brewprint legacy-ID compatibility treatment excluded; generic active `spec:` semantics retained |
| `spec:product.design_records.authoring_standards.agent_authoring_policy` | `## Rules > DRMCP retrieval` | TBD section | No current PRODUCT-owned semantic contract to distribute |
| `spec:product.design_records.authoring_standards.agent_authoring_policy` | `## Rules > Authoring transaction preference` | TBD section | No current PRODUCT-owned semantic contract to distribute |
| `spec:product.design_records.spec_format.document_shape` | `## Rules` | Concrete DRMCP-specific H1 example row | Host-specific example excluded or generalized; accepted H1 form retained |
| `spec:product.design_records.spec_format.document_shape` | `## Validation rules` | Migration-only wording around localized aliases and canonical `## What this is` enforcement | Current migration execution context excluded or generalized; generic validation condition retained |
| `spec:product.design_records.spec_format.topics_table` | `## Validation rules` | Project-specific follow-up ownership prose after canonical `ref` rules | Follow-up ownership excluded; parent declaration rule retained |
| `spec:product.design_records.spec_format.spec_id_as_ref` | `## Rules` | Concrete `product`, `drmcp`, and `bpdsl` repository path examples where they act as host-specific examples | Host-specific examples excluded or generalized; generic path-to-ref grammar retained |
| `spec:product.design_records.spec_format.spec_id_as_ref` | `## Boundary` | Migration, inventory, transient working-state, and compatibility-exception execution context | Project execution context excluded; generic compatibility boundary retained |
| `spec:product.design_records.traceability.metadata_schema` | `## Metadata boundary` | Brewprint migration-history wording and historical disposition evidence prose | Project history excluded from portable metadata projection |
| `spec:product.design_records.traceability.metadata_schema` | `## Metadata boundary` | Historical `T05` disposition-evidence pointer | `T05` is a local shorthand rather than a canonical Design Records ref; historical project evidence context is excluded from the portable projection. |
| `spec:product.design_records.traceability.metadata_schema` | `## Metadata boundary` | Concrete DRMCP parser/writer ownership wording and BPDSL endpoint statement | Implementation-specific authority excluded; generic metadata boundary retained |
| `spec:product.design_records.artifact_model.change_and_investigation_flow` | `## Deferred implementation tracking disposition` | BPDSL staging row and concrete DRMCP app-local projection wording | Staging and implementation-owner context excluded; generic workflow semantics retained |
| `spec:product.design_records.artifact_model.change_and_investigation_flow` | `## Deferred implementation tracking disposition` | Internal-design and YAML implementation-update disposition rows where they are historical/non-current | Historical and non-current implementation tracking context excluded; generic artifact flow retained |
| `spec:product.design_records.authoring_standards` | `## What this is` | `brewprint design records` wording | App-specific prose generalized for portable authoring standards projection |
| `spec:product.design_records.authoring_standards` | `## Topics` | `brewprint design records`, `DRMCP-dependent`, and `DRMCP-managed` summary wording | App/implementation-specific summaries generalized; child refs remain included |
| `spec:product.design_records.authoring_standards.writing_standard` | `## What this is` | Brewprint-specific framing | Generic prose and AI-output rules retained; app-specific framing generalized |
| `spec:product.design_records.authoring_standards.artifact_boundary` | `## What this is` | `DRMCP-managed` wording | Implementation-specific artifact framing generalized |
| `spec:product.design_records.authoring_standards.artifact_boundary` | `## Not covered` | Concrete BPDSL / non-DRMCP host-application examples where host-specific | Generic fact that non-Design-Records artifacts are outside the contract retained; app-local examples generalized |

### Historical section-projection audit: 4. Closure result

The metrics in this section were produced for the previous section-level package projection premise. They are retained to preserve source-authoring audit evidence and count provenance, but they no longer define the current package boundary or package-generation contract.

| metric | value |
| --- | --- |
| Required semantic areas from PRODUCT-REQ-SPEC-003 | 12 |
| Areas covered | 12 |
| Unique included source specs: `whole_spec` | 0 |
| Unique included source specs: `section_projection` | 33 |
| Unique excluded source specs | 1 |
| Included internal dependency edges targeting `whole_spec` | 0 |
| Included internal dependency edges targeting `section_projection` | 130 |
| Excluded internal dependency edges | 2 |
| External `spec:` dependency edges | 39 |
| Unique external `spec:` dependency refs | 21 |
| Excluded canonical record refs | 13 |
| Excluded prose/section/row items without canonical refs | 40 |
| Unresolved inclusion decisions | 0 |
| Unresolved retained refs | 0 |
| Ownership or dependency contradictions | 0 |
| Package ref closure verified | Historical only: all retained `spec:` refs in projected content resolved to section-projected sources under the old premise; the two excluded-child edges (`follow_up_boundary`) appeared only in Topics/Topic map rows omitted from that historical projection |
| Historical done condition substantively satisfied | Yes under the superseded premise: all 12 required areas had at least one authoritative PRODUCT source with an explicit contract area and inclusion decision; all transitive dependencies from included/projected content were explicitly enumerated and disposed; no prohibited Brewprint, DRMCP, BPDSL, or V01 authority was retained in that historical model |

Reproducibility note: unique source counts derive from Section 0 inclusion inventory; internal dependency edge counts derive from Section 2a table rows by target disposition; external `spec:` edge and unique-ref counts derive from Section 2b table rows and deduplicated dependency refs; excluded canonical record-ref count derives from Section 3 `Excluded canonical record refs`; excluded prose/section/row count derives from the following Section 3 table; unresolved and contradiction counts derive from the explicit zero rows in this Closure result.

### Historical section-projection audit: previous Independent review

- **verdict**: PASS
- **scope**: Superseded section-level package projection premise only.
- **review result**: The projection manifest covered all 12 required semantic areas, classified all 34 source specs exactly once, closed retained internal refs, separated external `spec:` dependencies from canonical record-ref and non-edge exclusions, and recorded no unresolved inclusion decision or ownership contradiction under the previous model.
- **final counts**:
  - source specs: `whole_spec=0`, `section_projection=33`, `excluded=1`
  - internal dependency edges: `whole_spec=0`, `section_projection=130`, `excluded=2`
  - external `spec:` dependencies: `39` edges, `21` unique refs
  - excluded canonical record refs: `13`
  - excluded non-edge prose, sections, or rows: `40`
- **current treatment**: This PASS is not a review verdict for the current whole-tree correction. The current correction review is recorded separately below.

### Independent review of current whole-tree correction

- **verdict**: PASS
- **reviewed contract**:
  - authoritative source boundary is the whole `product/records/spec/design-records/` tree;
  - bundled destination is `<exe-dir>/design-records/`;
  - package generation performs whole-tree copy plus canonical `spec:product.design_records` to `spec:design_records` prefix rewriting only;
  - the producer emits warnings during generation or check execution without owning persistent warning evidence or automatic routing; human review and optional manual source-authoring Requirement creation are the follow-up path when tracked correction is required; producer warnings do not authorize package-time filtering or block generation;
  - the package producer does not perform section selection, semantic filtering, or prose generalization;
  - the previous section-projection audit is retained only as historical source-authoring evidence.
- **finding result**: No blocking, major, or minor findings.
- **closure**: The corrected T01 Done condition is satisfied. `PRODUCT-TASK-SPEC-013-02` may consume the whole-tree boundary as accepted input.
