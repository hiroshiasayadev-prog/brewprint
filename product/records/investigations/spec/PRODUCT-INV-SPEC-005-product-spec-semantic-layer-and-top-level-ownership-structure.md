# PRODUCT-INV-SPEC-005: Product spec semantic-layer and top-level ownership structure

- **status**: concluded
- **date**: 2026-06-24
- **trigger**: Discovery that the current PRODUCT spec hierarchy lacks a normative semantic-layer and ownership boundary, causing Brewprint-specific and app-specific material to be inserted into supposedly generic contract specs.
- **scope**: Investigate the target top-level structure under `product/records/spec/`, the responsibility of each root overview, decomposition rules inside each top-level area, and a staged classification model that keeps future review scope bounded.
- **non_scope**: Excludes implementation of the restructuring, file moves, ref migration, contract rewriting, DRMCP/BPDSL integration design, and complete normalization of the BPDSL specification hierarchy.
- **source_refs**:
  - PRODUCT-WORK-NAMESPACE-001
  - PRODUCT-INV-SPEC-004
  - spec:product.design_records.namespace_model
  - spec:product.brewprint.namespaces
  - spec:product.brewprint.compatibility
  - spec:product.design_records.artifact_model
  - spec:product.bpdsl
  - spec:product.brewprint
  - spec:product.design_records.spec_format
- **follow_up_candidates**:
  - Future requirement or work item to implement the accepted PRODUCT spec ownership boundary and directory structure.

## Investigation scope

This investigation covers the target top-level semantic structure under `product/records/spec/`.
It must identify the responsibility of each top-level area, the required content of each root overview, and the decomposition rules for child topics inside each area.
It prioritizes restoring a reviewable PRODUCT spec boundary.
It does not require complete normalization of the BPDSL specification hierarchy.

The investigation must use a semantic-layer model broader than the binary PRODUCT-versus-DRMCP classification used by `PRODUCT-INV-SPEC-004`.
At minimum, it must distinguish:

- Design Records semantics.
- PRODUCT-side BPDSL-related staging or profile material, if justified.
- Canonical BPDSL app-local semantics.
- Brewprint-specific profile and current-state information.
- Legacy compatibility and migration policy.
- DRMCP operational contracts.
- BPDSL operational contracts.
- Deferred future cross-system integration.

The classification must not assume that PRODUCT-side BPDSL material is a permanent canonical layer.
Most mature BPDSL-owned semantics already belong under, or are expected to belong under, `bpdsl/records/spec/`.
A possible `product/records/spec/bpdsl/` area is only a temporary quarantine/staging target for separating BPDSL-related material currently mixed into PRODUCT Design Records concept specs.
The immediate concern is isolating BPDSL-specific content so Design Records cleanup can proceed without deciding final BPDSL ownership.

The investigation must answer these questions:

1. What are the required top-level semantic areas under `product/records/spec/`?
2. Should the current `concepts/` area be replaced or narrowed to a Design Records-specific area?
3. Which current specifications belong to Design Records semantics?
4. Which current specifications contain BPDSL-related material, and what disposition should each item receive?
5. Which material belongs to Brewprint-specific profile, inventory, migration, or compatibility specifications?
6. Which material belongs in app-local specs under `drmcp/records/spec/` or `bpdsl/records/spec/`?
7. What content belongs in the root `product/records/spec/index.md`?
8. What must each top-level area overview define?
9. How should each top-level area be decomposed internally?
10. How should the existing mixed files be classified at section level before any move?
11. How should semantic changes be separated from mechanical ref migration so future reviews remain small and reviewable?
12. Should Design Records and BPDSL remain completely independent for now, with their relationship deferred until a concrete integration requirement exists?
13. Is a PRODUCT-side BPDSL area needed at all?
14. If a PRODUCT-side BPDSL area is needed, should it be a temporary staging/profile area or a durable PRODUCT-owned semantic area?
15. What minimal content may be placed in a PRODUCT-side BPDSL area without duplicating `bpdsl/records/spec/`?
16. Which content must move directly to `bpdsl/records/spec/`?
17. How should the restructuring avoid turning PRODUCT into the canonical owner of BPDSL internal semantics?

Classification must proceed in two stages:

1. Classify all current PRODUCT specs at file level.
2. Perform section-level classification only for top-level or area root overview/index files, `namespace-model/**`, `project-artifact-model/**`, and files classified as `mixed` during file-level classification.

The investigation must not require exhaustive section-level classification of files that have a clear single semantic owner.

For each top-level area overview, the investigation must evaluate:

- Owned semantic layer.
- Prohibited content.
- Dependency direction.
- Placement decision rules.
- Child topic map.
- Rules for project-specific examples.
- Rules preventing concrete registry data from being duplicated across contract specs.

Expected evidence:

- File-level classification for all current PRODUCT specs.
- Section-level classification only for root, known mixed areas, and files classified as mixed.
- A proposed top-level PRODUCT spec tree.
- A responsibility and prohibition table for each proposed area.
- The required content of each root overview.
- A dependency-direction model.
- A current-content-to-target-area map for mixed content.
- A focused list of BPDSL-specific content currently mixed into PRODUCT specs.
- A disposition for each BPDSL-specific item: temporarily isolate under PRODUCT-side BPDSL staging, move directly to `bpdsl/records/spec/` only when already plainly app-local or duplicated, or extract unadopted integration statements to follow-up artifacts.
- Reviewability rules separating semantic changes from mechanical ref synchronization.

## Out of scope

This investigation does not implement any restructuring.
It does not move files, migrate refs, rewrite contracts, update existing specs, update work items, update requirements, update ADRs, or create follow-up artifacts.

The investigation also does not design DRMCP/BPDSL integration behavior.
It may classify content that appears to belong in DRMCP or BPDSL app-local specs, but it must not define those apps' parser, scanning, authoring, UI, projection, or implementation contracts.
It does not attempt to complete the final BPDSL ownership model, redesign all BPDSL specifications, or migrate all BPDSL specifications.

## Background

`PRODUCT-WORK-NAMESPACE-001` was closed after restructuring the namespace-model specs.
Subsequent review found that the resulting files still mix several semantic layers:

- Generic Design Records semantics.
- Brewprint-specific app and domain definitions.
- V01 legacy compatibility and migration policy.
- DRMCP parser, scanning, authoring, UI, or projection behavior.
- BPDSL implementation-model concepts and Design Records-to-implementation flow.

The change touched approximately 40 files and became impractical to review because semantic redesign and downstream reference synchronization were combined.

The repository already separates `product/records/spec/concepts/` from `product/records/spec/brewprint/`.
However, their ownership boundary is not normatively defined.

The current `concepts/` tree is mostly about Design Records, but `project-artifact-model/` also contains BPDSL-related concepts such as Brewprint DSL definitions, DSL-to-source flow, render artifacts, internal design, and source and target implementation.
Therefore, a future top-level structure probably needs distinct target areas for at least Design Records semantics and Brewprint-specific profile or compatibility information.
It may also need a limited PRODUCT-side BPDSL-related quarantine/staging area, but this area must not duplicate, supersede, or judge final ownership against `bpdsl/records/spec/`.

This record does not assume final names or exact directory layout before investigation.

## What was investigated

The investigation inspected all 40 Markdown files currently under `product/records/spec/**`.
Each file was classified at file level by semantic layer, mixed status, recommended disposition, and reason.

The mandatory section-level review covered:

- The absent root `product/records/spec/index.md` as a placement gap.
- Current area roots and overview/index files: `product/records/spec/brewprint/index.md`, `product/records/spec/brewprint/layout/index.md`, `product/records/spec/concepts/authoring-standards/index.md`, `product/records/spec/concepts/namespace-model/index.md`, `product/records/spec/concepts/project-artifact-model/index.md`, `product/records/spec/concepts/repository-layout/index.md`, `product/records/spec/concepts/spec-format/index.md`, `product/records/spec/concepts/spec-format/overview.md`, and `product/records/spec/concepts/traceability/index.md`.
- Every file under `product/records/spec/concepts/namespace-model/**`.
- Every file under `product/records/spec/concepts/project-artifact-model/**`.
- Additional files classified as mixed during file-level review.

App-local context was inspected without redesigning either app hierarchy:

- `drmcp/records/spec/**` current tree.
- `drmcp/records/spec/design-records-mcp/overview.md`.
- `drmcp/records/spec/design-records-mcp/namespace-scanning.md`.
- `drmcp/records/spec/design-records-mcp/schema/id-normalization.md`.
- `drmcp/records/spec/design-records-mcp/resolver.md`.
- `drmcp/records/spec/design-records-mcp/schema/overview.md`.
- `bpdsl/records/spec/**` current tree.
- `bpdsl/records/spec/overview.md`.
- `bpdsl/records/spec/dsl/overview.md`.
- `bpdsl/records/spec/dsl/project-layout.md`.
- `bpdsl/records/spec/views/overview.md`.
- `bpdsl/records/spec/mcp/overview.md`.

Historical context was inspected through Design Records MCP:

- `PRODUCT-WORK-NAMESPACE-001`.
- `PRODUCT-INV-SPEC-004`.

`PRODUCT-INV-SPEC-004` was treated as evidence that the earlier binary PRODUCT/DRMCP model was insufficient.
The current classification uses the broader semantic-layer model defined in this investigation.

## Findings

### Finding 1: current PRODUCT spec tree has no placement router

There is no `product/records/spec/index.md`.
The current top-level roots are `product/records/spec/concepts/` and `product/records/spec/brewprint/`.
`concepts/` is too broad to tell authors whether a file owns generic Design Records semantics, Brewprint registry/profile facts, Brewprint compatibility/migration policy, BPDSL staging, or unadopted future mechanisms that should not remain in current specs.

### Finding 2: file-level classification

| current ref/path | primary semantic layer | mixed? | recommended disposition | reason |
|---|---|---:|---|---|
| `product/records/spec/brewprint/index.md` | Brewprint profile/current state | no | keep under Brewprint profile area | Navigation root for project-specific current-state references. |
| `product/records/spec/brewprint/layout/index.md` | Brewprint profile/current state | no | keep under Brewprint profile area | Inventory of current directories, explicitly non-normative. |
| `product/records/spec/concepts/authoring-standards/adr-authoring.md` | Design Records semantics | no | move under Design Records area | Defines ADR authoring grammar, shape, metadata, lifecycle, and boundaries. |
| `product/records/spec/concepts/authoring-standards/agent-authoring-policy.md` | Design Records semantics | no | move under Design Records area | Governs agent behavior for design record authoring; DRMCP-dependent parts are placeholders. |
| `product/records/spec/concepts/authoring-standards/artifact-boundary.md` | Design Records semantics | no | move under Design Records area | Defines artifact-selection boundary for design record authoring. |
| `product/records/spec/concepts/authoring-standards/index.md` | Design Records semantics | no | move under Design Records area | Area index for authoring standards. |
| `product/records/spec/concepts/authoring-standards/investigation-authoring.md` | Design Records semantics | no | move under Design Records area | Defines investigation record authoring rules and metadata. |
| `product/records/spec/concepts/authoring-standards/requirement-authoring.md` | Design Records semantics | no | move under Design Records area | Defines requirement record authoring rules and metadata. |
| `product/records/spec/concepts/authoring-standards/spec-authoring.md` | Design Records semantics | no | move under Design Records area | Defines spec record authoring rules and path-derived identity. |
| `product/records/spec/concepts/authoring-standards/task-authoring.md` | Design Records semantics | no | move under Design Records area | Defines task record authoring rules and metadata. |
| `product/records/spec/concepts/authoring-standards/work-item-authoring.md` | Design Records semantics | no | move under Design Records area | Defines work-item authoring rules and metadata. |
| `product/records/spec/concepts/authoring-standards/writing-standard.md` | Design Records semantics | no | move under Design Records area | Prose and AI-output rules for design records. |
| `product/records/spec/concepts/namespace-model/app-namespaces.md` | Brewprint profile/current state | yes | split | It mixes generic app-namespace concepts with concrete DRMCP/BPDSL/PRODUCT architecture sketches and current domain assignments. |
| `product/records/spec/concepts/namespace-model/artifact-id-grammar.md` | Design Records semantics | no | move under Design Records area | Cross-app public ID grammar for design record and workflow artifacts. |
| `product/records/spec/concepts/namespace-model/domain-catalog.md` | Brewprint profile/current state | yes | move to Brewprint profile with generic links kept elsewhere | Concrete current domain catalog changes when app/domain registry changes. |
| `product/records/spec/concepts/namespace-model/existing-artifacts.md` | compatibility/migration | yes | split | It combines V01 ownership policy, effective current attribution, and UI/MCP projection language. |
| `product/records/spec/concepts/namespace-model/index.md` | Design Records semantics | yes | split | It combines namespace concepts, current registry, future layout, compatibility context, and placement routing. |
| `product/records/spec/concepts/namespace-model/legacy-id-compatibility.md` | compatibility/migration | no | move under Brewprint compatibility | V01 issued-ID retention and compatibility policy belongs with Brewprint historical compatibility, not a top-level PRODUCT area. |
| `product/records/spec/concepts/namespace-model/subdomain-model.md` | Design Records semantics | yes | split | Generic subdomain model is reusable, but the DRMCP MCP example is project/app-specific. |
| `product/records/spec/concepts/project-artifact-model/artifact-responsibility-matrix.md` | Design Records semantics | yes | split | Design record artifact responsibilities are mixed with BPDSL, render, source, target implementation, and impl-note layers. |
| `product/records/spec/concepts/project-artifact-model/change-and-investigation-flow.md` | Design Records semantics | yes | split | Investigation/ADR/REQ/WORK/TASK flow is generic; YAML/internal-design tracking is a deferred integration edge. |
| `product/records/spec/concepts/project-artifact-model/design-flow.md` | PRODUCT-side BPDSL staging/profile | yes | temporarily isolate under PRODUCT-side BPDSL staging | It defines Design Records-to-DSL-to-source flow before BPDSL migration can determine final ownership. |
| `product/records/spec/concepts/project-artifact-model/index.md` | Design Records semantics | yes | split | Overview mixes artifact responsibility, DSL/source/render model, Brewprint current state, and deferred trace mechanisms. |
| `product/records/spec/concepts/project-artifact-model/traceability-boundary.md` | Design Records semantics | yes | split | Traceability boundary is generic, but the diagram includes DRMCP implementation role and deferred internal/YAML endpoints. |
| `product/records/spec/concepts/repository-layout/index.md` | Design Records semantics | yes | split | Records placement is generic, but `dsl/` and `src/` bootstrap rules are BPDSL staging / implementation-model material. |
| `product/records/spec/concepts/repository-layout/record-discovery-paths.md` | Design Records semantics | yes | split | Generic path patterns belong in PRODUCT, but `namespace_prefix` derivation is currently pointed at DRMCP. |
| `product/records/spec/concepts/spec-format/document-shape.md` | Design Records semantics | no | move under Design Records area | Accepted document shape contract for spec records. |
| `product/records/spec/concepts/spec-format/follow-up-boundary.md` | Design Records semantics | no | move under Design Records area | Separates PRODUCT spec-format stabilization from DRMCP implementation follow-up. |
| `product/records/spec/concepts/spec-format/index.md` | Design Records semantics | no | move under Design Records area | Navigation root for spec-format contracts. |
| `product/records/spec/concepts/spec-format/overview.md` | Design Records semantics | no | move under Design Records area | Overview for visible Markdown spec-format contract. |
| `product/records/spec/concepts/spec-format/spec-id-as-ref.md` | Design Records semantics | no | move under Design Records area | Path-derived `spec:` identity model. |
| `product/records/spec/concepts/spec-format/topics-table.md` | Design Records semantics | no | move under Design Records area | Topics table and parent declaration contract. |
| `product/records/spec/concepts/spec-format/validation-policy.md` | Design Records semantics | no | move under Design Records area | Validation ownership policy; DRMCP rows are follow-up ownership references, not parser implementation. |
| `product/records/spec/concepts/traceability/artifact-refs.md` | Design Records semantics | yes | split | Active Design Records refs are mixed with reserved `yaml:` and deferred realization endpoints. |
| `product/records/spec/concepts/traceability/coverage-mapping.md` | deferred integration | yes | extract to follow-up investigation/requirement or remove after evidence transfer | It mainly records deferred realization/coverage mechanisms rather than active trace contract. Current specs should not preserve a catalog of unadopted mechanisms. |
| `product/records/spec/concepts/traceability/index.md` | Design Records semantics | yes | split | Active reference foundation is mixed with reserved BPDSL YAML and deferred relation mechanisms. |
| `product/records/spec/concepts/traceability/metadata-schema.md` | Design Records semantics | yes | split and reconcile | Trace metadata model references spec front matter, which conflicts with accepted visible spec-format policy for migrated specs. |
| `product/records/spec/concepts/traceability/out-of-scope.md` | deferred integration | yes | extract to follow-up investigation/requirement or remove after evidence transfer | It mainly catalogs future endpoints, relation artifacts, YAML trace, fixture trace, and MCP writer candidates rather than active current contracts. |
| `product/records/spec/concepts/traceability/resolve-and-validation.md` | Design Records semantics | yes | split | Canonical reference semantics are generic, but resolver behavior and MCP language need a strict tool-contract boundary. |
| `product/records/spec/concepts/traceability/semantic-ref.md` | Design Records semantics | no | move under Design Records area | Defines active `spec:` grammar and stability rules; deferred endpoint mentions are boundary notes. |

### Finding 3: mixed-file section classification

| current file | section | semantic layer | target owner/area | action | reason |
|---|---|---|---|---|---|
| `product/records/spec/index.md` | absent root overview | Design Records semantics / router | `product/records/spec/index.md` | create in later work | Top-level routing is missing. |
| `product/records/spec/brewprint/index.md` | `## What this is` | Brewprint profile/current state | Brewprint profile | keep | Current-state navigation root. |
| `product/records/spec/brewprint/index.md` | `## Topics` | Brewprint profile/current state | Brewprint profile | keep | Points to project inventory. |
| `product/records/spec/brewprint/layout/index.md` | `## What this is` | Brewprint profile/current state | Brewprint profile | keep | Declares inventory-only role. |
| `product/records/spec/brewprint/layout/index.md` | `## Current app namespace layout` | Brewprint profile/current state | Brewprint profile | keep | Current repository facts. |
| `product/records/spec/brewprint/layout/index.md` | `## Other current repository areas` | Brewprint profile/current state | Brewprint profile | keep | Current repository facts outside app namespaces. |
| `product/records/spec/brewprint/layout/index.md` | `## Maintenance rule` | Brewprint profile/current state | Brewprint profile | keep | Update rule for observed inventory. |
| `product/records/spec/brewprint/layout/index.md` | `## Related specs` | Brewprint profile/current state | Brewprint profile | keep | Points to normative layout and namespace owners. |
| `product/records/spec/concepts/authoring-standards/index.md` | `## What this is` | Design Records semantics | Design Records area | move | Authoring standards belong with design-record semantics. |
| `product/records/spec/concepts/authoring-standards/index.md` | `## Topics` | Design Records semantics | Design Records area | move | Child topics are record-authoring contracts. |
| `product/records/spec/concepts/namespace-model/index.md` | `## What this is` | Design Records semantics | Design Records area | rewrite | Keep generic namespace model; remove Brewprint registry framing. |
| `product/records/spec/concepts/namespace-model/index.md` | `## Current contract` | Brewprint profile/current state | Brewprint profile | split | Active namespace list changes with current registry. |
| `product/records/spec/concepts/namespace-model/index.md` | `## Boundary` | mixed | Design Records root plus Brewprint profile and migration areas | split | Boundary currently combines generic, registry, layout, and tool concerns. |
| `product/records/spec/concepts/namespace-model/index.md` | `## Current placement and future layout` | compatibility/migration / deferred integration | Brewprint compatibility or follow-up work | extract or defer destination | Future registry placement is not a stable namespace contract yet. |
| `product/records/spec/concepts/namespace-model/index.md` | `## App namespace and domain namespace` | Design Records semantics | Design Records area | move | Generic concepts remain valid if app names change. |
| `product/records/spec/concepts/namespace-model/index.md` | `## Topics` | mixed | root placement router plus child areas | split | Current child list crosses generic, profile, and compatibility owners. |
| `product/records/spec/concepts/namespace-model/index.md` | `## Sources` | compatibility/migration | Brewprint compatibility notes | keep or trim | Historical sources are useful but should not define current placement. |
| `product/records/spec/concepts/namespace-model/app-namespaces.md` | `## What this is` | Brewprint profile/current state | Brewprint profile | rewrite | It describes the concrete three active namespaces. |
| `product/records/spec/concepts/namespace-model/app-namespaces.md` | `## App namespace definitions` | Brewprint profile/current state | Brewprint profile | move | Concrete app list changes when registry changes. |
| `product/records/spec/concepts/namespace-model/app-namespaces.md` | `## DRMCP` | DRMCP operational contract / Brewprint profile | DRMCP spec plus Brewprint profile summary | split | Architecture diagram and tool list are DRMCP-owned; current assignment can be profiled. |
| `product/records/spec/concepts/namespace-model/app-namespaces.md` | `## BPDSL` | canonical BPDSL app-local semantics / Brewprint profile | Brewprint profile plus temporary BPDSL staging or app-local pointer | split | DSL type/render/resolution details are expected to be app-local; current namespace assignment can be profiled without deciding final BPDSL migration. |
| `product/records/spec/concepts/namespace-model/app-namespaces.md` | `## PRODUCT` | Brewprint profile/current state | Brewprint profile | move | Defines current PRODUCT namespace nature and domains. |
| `product/records/spec/concepts/namespace-model/artifact-id-grammar.md` | `## What this is` | Design Records semantics | Design Records area | move | Cross-app artifact ID grammar. |
| `product/records/spec/concepts/namespace-model/artifact-id-grammar.md` | `## Grammar` | Design Records semantics | Design Records area | move | Stable public ID grammar. |
| `product/records/spec/concepts/namespace-model/artifact-id-grammar.md` | `## Sequence format` | Design Records semantics | Design Records area | move | Stable allocation rule. |
| `product/records/spec/concepts/namespace-model/artifact-id-grammar.md` | `## Canonical reference` | Design Records semantics | Design Records area | move | ID-as-ref contract. |
| `product/records/spec/concepts/namespace-model/artifact-id-grammar.md` | `## Related specs` | Design Records semantics | Design Records area | update later | Refs need mechanical sync after move. |
| `product/records/spec/concepts/namespace-model/domain-catalog.md` | `## What this is` | Brewprint profile/current state | Brewprint profile | move | Catalogs current assignments. |
| `product/records/spec/concepts/namespace-model/domain-catalog.md` | `## Canonical domain namespaces` | Brewprint profile/current state | Brewprint profile | move | Concrete app/domain registry changes over time. |
| `product/records/spec/concepts/namespace-model/domain-catalog.md` | `## Existing prefixes outside the canonical catalog` | compatibility/migration | Brewprint compatibility | split | Existing legacy prefixes are migration/compatibility facts. |
| `product/records/spec/concepts/namespace-model/existing-artifacts.md` | `## What this is` | compatibility/migration | Brewprint compatibility | split | Combines ownership decision, attribution, and projection. |
| `product/records/spec/concepts/namespace-model/existing-artifacts.md` | `## Historical ownership decision` | compatibility/migration | Brewprint compatibility | move | V01 decision and issued-ID retention. |
| `product/records/spec/concepts/namespace-model/existing-artifacts.md` | `## Effective attribution` | Brewprint profile/current state / DRMCP operational contract | Brewprint profile, with DRMCP projection in DRMCP if needed | split | Attribution map is project profile; UI/MCP projection behavior is tool-owned. |
| `product/records/spec/concepts/namespace-model/existing-artifacts.md` | `## New-artifact ownership` | Design Records semantics | Design Records area | move | General authoring rule for new records. |
| `product/records/spec/concepts/namespace-model/legacy-id-compatibility.md` | `## What this is` | compatibility/migration | Brewprint compatibility | move | V01 compatibility role. |
| `product/records/spec/concepts/namespace-model/legacy-id-compatibility.md` | `## Accepted legacy families` | compatibility/migration | Brewprint compatibility | move | Lists issued legacy families. |
| `product/records/spec/concepts/namespace-model/legacy-id-compatibility.md` | `## Retention policy` | compatibility/migration | Brewprint compatibility | move | Retention and resolvability rule. |
| `product/records/spec/concepts/namespace-model/legacy-id-compatibility.md` | `## Spec identity note` | compatibility/migration / Design Records semantics | Brewprint compatibility with pointer to spec-format | keep pointer only | Spec identity itself is owned by spec-format. |
| `product/records/spec/concepts/namespace-model/subdomain-model.md` | `## What this is` | Design Records semantics | Design Records area | move | Generic grouping model. |
| `product/records/spec/concepts/namespace-model/subdomain-model.md` | `## Subdomain model` | Design Records semantics | Design Records area | move | Generic subdomain rule. |
| `product/records/spec/concepts/namespace-model/subdomain-model.md` | `## Definition and representation` | Design Records semantics | Design Records area | move | Metadata representation rule. |
| `product/records/spec/concepts/namespace-model/subdomain-model.md` | `## Write-time advisory` | Design Records semantics | Design Records area | move | Authoring guidance. |
| `product/records/spec/concepts/namespace-model/subdomain-model.md` | `## DRMCP MCP domain example` | DRMCP operational contract / example | examples or DRMCP spec | split | App-specific example should not be normative in generic model. |
| `product/records/spec/concepts/project-artifact-model/index.md` | `## What this is` | Design Records semantics | Design Records area | rewrite | Keep artifact responsibility scope; remove implementation-model ownership claims. |
| `product/records/spec/concepts/project-artifact-model/index.md` | `## Current contract` | Design Records semantics | Design Records area | split | Generic contract should not own BPDSL internals. |
| `product/records/spec/concepts/project-artifact-model/index.md` | `## Boundary` | mixed | Design Records area plus temporary BPDSL staging and follow-up artifacts | split | Current boundary crosses artifact model, DSL/source, and future trace mechanisms. |
| `product/records/spec/concepts/project-artifact-model/index.md` | `## Artifact classes` | mixed | Design Records area plus BPDSL app-local specs | split | DSL, source, render, and target implementation are not all Design Records semantics. |
| `product/records/spec/concepts/project-artifact-model/index.md` | `## Source of truth and documentation ownership` | mixed | Design Records area plus app-local specs | split | DRMCP and BPDSL ownership rows should be pointers, not PRODUCT-owned contracts. |
| `product/records/spec/concepts/project-artifact-model/index.md` | `## Navigation` | mixed | root router | split | Navigation currently bridges app-local specs and unadopted future mechanism material. |
| `product/records/spec/concepts/project-artifact-model/index.md` | `## MVP scope and future extensions` | deferred integration | follow-up investigation/requirement or removal after evidence transfer | extract | Relation, internal-design, YAML, and coverage candidates are not current contracts. |
| `product/records/spec/concepts/project-artifact-model/index.md` | `## Topics` | mixed | Design Records area plus BPDSL staging | split | Topic list includes mixed child files. |
| `product/records/spec/concepts/project-artifact-model/artifact-responsibility-matrix.md` | `## What this is` | Design Records semantics | Design Records area | rewrite | Keep design record artifact responsibilities. |
| `product/records/spec/concepts/project-artifact-model/artifact-responsibility-matrix.md` | `## Artifact responsibility matrix` | mixed | Design Records area plus temporary BPDSL staging; future mechanisms to follow-up artifacts | split | Matrix crosses ADR/spec/workflow with DSL/source/render/target layers. |
| `product/records/spec/concepts/project-artifact-model/change-and-investigation-flow.md` | `## What this is` | Design Records semantics | Design Records area | move | Change-flow scope is design-record oriented. |
| `product/records/spec/concepts/project-artifact-model/change-and-investigation-flow.md` | `## Change and investigation flow` | mixed | Design Records area plus follow-up extraction for unadopted mechanisms | split | Diagram includes internal-design and YAML update paths. |
| `product/records/spec/concepts/project-artifact-model/change-and-investigation-flow.md` | `## Artifact ownership in the flow` | Design Records semantics | Design Records area | move | Defines ADR/INV/REQ/WORK/TASK roles. |
| `product/records/spec/concepts/project-artifact-model/change-and-investigation-flow.md` | `## Rules` | Design Records semantics | Design Records area | move | Generic workflow rules. |
| `product/records/spec/concepts/project-artifact-model/change-and-investigation-flow.md` | `## Sources` | compatibility/migration | Design Records area | keep | Historical rationale pointer. |
| `product/records/spec/concepts/project-artifact-model/design-flow.md` | `## What this is` | PRODUCT-side BPDSL staging/profile | temporary BPDSL staging | rewrite | Defines target DSL flow before BPDSL migration can determine final ownership. |
| `product/records/spec/concepts/project-artifact-model/design-flow.md` | `## Design artifact flow` | PRODUCT-side BPDSL staging/profile / deferred integration | temporary BPDSL staging plus follow-up extraction for integration claims | split | Spec-to-DSL-to-source flow is not yet an accepted integration contract. |
| `product/records/spec/concepts/project-artifact-model/design-flow.md` | `## Source of truth roles` | mixed | Design Records area plus BPDSL app-local specs | split | Spec role is generic; DSL/source/render roles are BPDSL/implementation model. |
| `product/records/spec/concepts/project-artifact-model/design-flow.md` | `## Rules` | mixed | Design Records area plus BPDSL staging | split | Bootstrap implementation rule may remain PRODUCT-side until app-local ownership is accepted. |
| `product/records/spec/concepts/project-artifact-model/design-flow.md` | `## Related specs` | mixed | owner-specific pointers | update later | Refs need sync only after decision. |
| `product/records/spec/concepts/project-artifact-model/design-flow.md` | `## Sources` | compatibility/migration | staged note | keep | Historical rationale only. |
| `product/records/spec/concepts/project-artifact-model/traceability-boundary.md` | `## What this is` | Design Records semantics | Design Records area | move | Traceability as artifact mechanism. |
| `product/records/spec/concepts/project-artifact-model/traceability-boundary.md` | `## Traceability and tool boundary` | mixed | Design Records area plus DRMCP pointers | split | Diagram includes Design Records MCP implementation role. |
| `product/records/spec/concepts/project-artifact-model/traceability-boundary.md` | `## MVP scope` | mixed | Design Records area plus follow-up extraction for deferred endpoints | split | Active ID-as-ref rules and deferred internal/YAML endpoints are mixed. |
| `product/records/spec/concepts/project-artifact-model/traceability-boundary.md` | `## Sources` | compatibility/migration | Design Records area | keep | Historical rationale pointer. |
| `product/records/spec/concepts/repository-layout/index.md` | `## What this is` | mixed | Design Records area plus BPDSL staging | split | `records/` is generic; `dsl/` and `src/` are implementation-model concerns. |
| `product/records/spec/concepts/repository-layout/index.md` | `## Current contract` | mixed | Design Records area plus BPDSL staging | split | App namespace layout rule crosses records and implementation source. |
| `product/records/spec/concepts/repository-layout/index.md` | `## Rules` | mixed | Design Records area plus BPDSL staging | split | Records placement and DSL/source bootstrap should be separate. |
| `product/records/spec/concepts/repository-layout/index.md` | `## Boundary` | mixed | Design Records area plus BPDSL staging | split | Boundary already excludes BPDSL language/generator contracts. |
| `product/records/spec/concepts/repository-layout/index.md` | `## Topics` | Design Records semantics | Design Records area | move | Record discovery paths topic. |
| `product/records/spec/concepts/repository-layout/index.md` | `## Related specs` | mixed | owner-specific pointers | update later | Cross-owner links must remain pointers. |
| `product/records/spec/concepts/repository-layout/record-discovery-paths.md` | `## What this is` | Design Records semantics / DRMCP operational contract | Design Records area plus DRMCP spec | split | Path patterns are generic; DRMCP filter behavior stays app-local. |
| `product/records/spec/concepts/repository-layout/record-discovery-paths.md` | `## Current contract` | mixed | Design Records area plus DRMCP spec | split | References `namespace_prefix` derivation in DRMCP. |
| `product/records/spec/concepts/spec-format/index.md` | `## What this is` | Design Records semantics | Design Records area | move | Navigation root for spec-format. |
| `product/records/spec/concepts/spec-format/index.md` | `## Topics` | Design Records semantics | Design Records area | move | Child topic routing. |
| `product/records/spec/concepts/spec-format/index.md` | `## Related specs` | Design Records semantics | Design Records area | update later | Cross-owner reference remains pointer-only. |
| `product/records/spec/concepts/spec-format/overview.md` | `## What this is` | Design Records semantics | Design Records area | move | Entry-level spec-format framing. |
| `product/records/spec/concepts/spec-format/overview.md` | `## Current contract` | Design Records semantics | Design Records area | move | Visible Markdown spec-format policy. |
| `product/records/spec/concepts/spec-format/overview.md` | `## Non-goals` | Design Records semantics | Design Records area | move | Ownership boundary against DRMCP implementation. |
| `product/records/spec/concepts/spec-format/overview.md` | `## Topic map` | Design Records semantics | Design Records area | move | Navigation map. |
| `product/records/spec/concepts/spec-format/overview.md` | `## Related specs` | Design Records semantics | Design Records area | update later | Mechanical refs later. |
| `product/records/spec/concepts/traceability/index.md` | `## What this is` | Design Records semantics | Design Records area | rewrite | Active reference foundation belongs with Design Records semantics. |
| `product/records/spec/concepts/traceability/index.md` | `## Current contract` | Design Records semantics | Design Records area | move | Active resolve/validation scope. |
| `product/records/spec/concepts/traceability/index.md` | `## Purpose` | Design Records semantics | Design Records area | move | Reference foundation purpose. |
| `product/records/spec/concepts/traceability/index.md` | `## MVP scope` | mixed | Design Records area plus follow-up extraction for unadopted mechanisms | split | Active `spec:`/ID-as-ref scope and `yaml:` reserve are mixed. |
| `product/records/spec/concepts/traceability/index.md` | `## Out of MVP scope` | deferred integration | follow-up investigation/requirement or removal after evidence transfer | extract | Future endpoints and relation mechanisms are not current spec contracts. |
| `product/records/spec/concepts/traceability/index.md` | `## Terms` | mixed | Design Records area plus BPDSL staging pointer | split | `brewprint DSL YAML` term should not create PRODUCT ownership of BPDSL YAML semantics. |
| `product/records/spec/concepts/traceability/index.md` | `## Topics` | mixed | Design Records area plus follow-up extraction for deferred material | split | Child topics include active and deferred material. |
| `product/records/spec/concepts/traceability/index.md` | `## Source of truth boundary` | Design Records semantics / DRMCP operational contract | Design Records area plus DRMCP pointer | keep with pointer | Product owns semantic model; DRMCP implements. |
| `product/records/spec/concepts/traceability/artifact-refs.md` | `## Active prefixes` | Design Records semantics | Design Records area | move | Active `spec:` identity. |
| `product/records/spec/concepts/traceability/artifact-refs.md` | `## Reserved prefixes` | PRODUCT-side BPDSL staging/profile | temporary BPDSL staging or follow-up extraction | split | `yaml:` reserve touches BPDSL without defining it. |
| `product/records/spec/concepts/traceability/artifact-refs.md` | `## Deferred prefixes` | deferred integration | follow-up investigation/requirement or removal after evidence transfer | extract | Internal-design and coverage endpoint decisions are deferred. |
| `product/records/spec/concepts/traceability/artifact-refs.md` | `## ID-as-ref` | Design Records semantics / compatibility | Design Records area plus Brewprint compatibility | split | New grammar and legacy compatibility are distinct owners. |
| `product/records/spec/concepts/traceability/artifact-refs.md` | `## Design record ID-as-ref boundary` | Design Records semantics | Design Records area | move | Active design record ref boundary. |
| `product/records/spec/concepts/traceability/artifact-refs.md` | `## Workflow ID-as-ref boundary` | Design Records semantics | Design Records area | move | Active workflow relation identity. |
| `product/records/spec/concepts/traceability/artifact-refs.md` | `## COV-*` / `## Relation endpoint boundary` / `## Scope-out prefixes` | deferred integration | follow-up investigation/requirement or removal after evidence transfer | extract | Future endpoints, relation identities, and fixture scope are not current spec contracts. |
| `product/records/spec/concepts/traceability/coverage-mapping.md` | all H2 sections | deferred integration | follow-up investigation/requirement or removal after evidence transfer | extract | File mainly preserves future realization and coverage triggers rather than current contracts. |
| `product/records/spec/concepts/traceability/metadata-schema.md` | `## Trace metadata` / `## Common fields` / `## semantic_refs` / `## sections` | Design Records semantics with stale format assumption | Design Records area plus spec-format reconciliation | rewrite | Front matter assumptions conflict with accepted visible spec-format policy. |
| `product/records/spec/concepts/traceability/metadata-schema.md` | internal design / coverage / YAML / workflow sections | deferred integration | follow-up investigation/requirement or removal after evidence transfer | split | Future endpoints and relations are not active trace contract. |
| `product/records/spec/concepts/traceability/out-of-scope.md` | all H2 sections | deferred integration | follow-up investigation/requirement or removal after evidence transfer | extract | Catalogs excluded future mechanisms and triggers rather than current contracts. |
| `product/records/spec/concepts/traceability/resolve-and-validation.md` | `## Resolve` / `## Resolver input` / `## Lookup sources` / `## Section anchor lookup` / `## Duplicate detection` / `## Unresolved reference and declared relation integrity` / `## Validation` | Design Records semantics | Design Records area | move with wording cleanup | Keep canonical inputs and invalid conditions. |
| `product/records/spec/concepts/traceability/resolve-and-validation.md` | `## Resolver output` / `## MCP writer contract placeholder` | DRMCP operational contract | DRMCP spec or pointer-only note | split | Concrete request/response and writer schemas are tool-owned. |

### Finding 4: proposed top-level PRODUCT spec tree

The preferred candidate tree is:

```text
product/records/spec/
  index.md

  design-records/
    index.md
    authoring-standards/
    namespace-model/
    repository-layout/
    spec-format/
    traceability/
    artifact-model/

  brewprint/
    index.md
    layout/
    namespaces/
    compatibility/

  bpdsl/
    index.md
    ...
```

The exact child layout remains a follow-up decision.
The key ownership correction is that compatibility is not a top-level PRODUCT spec area.
Brewprint V01 compatibility, historical attribution, and migration state belong under `product/records/spec/brewprint/compatibility/`.

There should be no top-level `deferred-integration/` spec area.
Unadopted future material such as inactive `yaml:` relations, internal-design endpoints, coverage artifacts, unrealized Design Records-to-BPDSL relations, future relation lifecycle, and reintroduction triggers should be extracted to an appropriate follow-up artifact or removed after evidence transfer.
Current spec files should contain current contracts, not a catalog of unadopted future mechanisms.

`product/records/spec/bpdsl/` should be a temporary quarantine/staging area.
Its purpose is to isolate BPDSL-related material currently mixed into PRODUCT Design Records specifications so Design Records cleanup can proceed without expanding into a full BPDSL migration.
Placement there is a priority-based scope decision, not a judgment that PRODUCT is the canonical owner of BPDSL semantics.
The final BPDSL specification hierarchy, canonical ownership of BPDSL-related PRODUCT content, boundary with `bpdsl/records/spec/**`, and Design Records-to-BPDSL relationship remain unresolved until BPDSL migration or a later explicit integration requirement.

### Finding 5: root overview contract

`product/records/spec/index.md` should act as a placement router.
It should not summarize child contracts in detail.

Minimum root responsibilities:

| root responsibility | required treatment |
|---|---|
| top-level semantic areas | Define `design-records/`, `brewprint/`, and temporary `bpdsl/`. |
| responsibility of each area | State the owned semantic layer and canonical child overview for each area. |
| prohibited content for each area | State what must move to another area or app-local owner. |
| dependency direction | State that generic Design Records semantics do not normatively depend on DRMCP or BPDSL. |
| placement decision rules | Include tests based on whether content changes with app registry, Brewprint V01 compatibility, DRMCP behavior, BPDSL behavior, or unadopted future integration. |
| examples | Require examples in generic contracts to be non-normative and app-neutral where possible. |
| registry duplication | Prohibit copying current app/domain registry tables into generic contracts. |
| topic/navigation table | Provide a ref-first table pointing to top-level area overviews. |

### Finding 6: area overview contracts

| area | owns | must not contain | may depend on | overview responsibilities |
|---|---|---|---|---|
| `design-records/` | App-independent Design Records semantics: artifact identity, record authoring, spec format, repository `records/` placement, traceability reference model, and design record artifact responsibilities. | Concrete DRMCP parser/tool/UI behavior; canonical BPDSL DSL/schema/render/runtime semantics; Brewprint current app/domain registry tables; Brewprint V01 compatibility/migration state except pointers. | Accepted ADRs, Design Records specs, Brewprint compatibility pointers, app-local specs as implementations. | Define Design Records semantic boundary, child topic map, app-neutral examples policy, and cross-owner pointer policy. |
| `brewprint/` | Brewprint-specific profile and current state: current repository layout, current namespaces, and Brewprint compatibility/migration history. Children include `layout/`, `namespaces/`, and `compatibility/`. | Generic Design Records rules; DRMCP tool contracts; canonical BPDSL language or render contracts; new BPDSL migration judgments. | Design Records contracts, app-local app specs, accepted compatibility/migration ADRs. | Record current project facts and maintenance rules; label all facts as inventory/profile/compatibility, not generic contract. |
| `bpdsl/` | Temporary isolation area for BPDSL-related material removed from mixed PRODUCT Design Records specs. | Canonical BPDSL semantics; BPDSL schema, type, resolver, render, generation, runtime, or MCP contracts; new Design Records-to-BPDSL integration design; unrelated new BPDSL specs. | `bpdsl/records/spec/**` and Design Records contracts as context only. | Declare temporary status, allowed preservation scope, prohibited work, review boundary, migration trigger, and exit condition. |

### Finding 7: temporary BPDSL root overview contract

A future `product/records/spec/bpdsl/index.md` must make the temporary status explicit.

| contract topic | required content |
|---|---|
| Purpose | Temporarily isolate BPDSL-related material removed from Design Records specifications, preserve existing content and context until BPDSL migration can evaluate it, and keep the current Design Records restructuring reviewable. |
| Ownership status | State that this area is not the canonical BPDSL specification hierarchy, placement here does not establish PRODUCT ownership of BPDSL semantics, canonical ownership remains unresolved until BPDSL migration, and `bpdsl/records/spec/**` remains the expected final owner for BPDSL-internal contracts subject to later migration review. |
| Allowed content | Existing BPDSL-related content separated from mixed PRODUCT specs; existing DSL, source, render, implementation-flow, or BPDSL artifact descriptions whose final disposition cannot yet be judged; minimal context required to preserve meaning. |
| Prohibited current work | Redesigning BPDSL semantics; normalizing the BPDSL specification hierarchy; creating new BPDSL schema, type, resolver, render, generation, or MCP contracts; designing a new Design Records-to-BPDSL integration contract; treating staged content as reviewed or accepted BPDSL truth; adding unrelated new BPDSL specifications. |
| Review boundary | Current restructuring reviews only whether BPDSL-related content was removed from Design Records contracts, whether moved content was preserved without accidental semantic change, and whether the temporary area avoids claiming canonical authority. It does not review BPDSL correctness, duplication with app-local BPDSL specs, final directory structure, final semantic ownership, or final integration model. |
| Exit condition | During BPDSL migration, every staged item must be reclassified into one of: move to `bpdsl/records/spec/**`; redefine as genuine PRODUCT-owned cross-cutting policy; move to another appropriate artifact; delete as obsolete or duplicate. After that review, the temporary area must be removed or explicitly redefined with an accepted permanent responsibility. |

This mechanism defers the decision to the known BPDSL migration event.
It is not indefinite postponement.

### Finding 8: dependency-direction model

| dependency question | answer |
|---|---|
| May Design Records semantics normatively depend on DRMCP? | No. DRMCP may implement, validate, expose, or index Design Records contracts, but it must not be the normative source for generic Design Records semantics. |
| May Design Records semantics normatively depend on BPDSL? | No. Design Records semantics may reserve or defer relation points to BPDSL, but the generic contract must remain valid if BPDSL names or internals change. |
| May Brewprint profile instantiate Design Records contracts? | Yes. Brewprint profile may list current namespaces, directories, assignments, and examples that instantiate generic contracts. |
| May DRMCP implement/reference Design Records contracts? | Yes. DRMCP should reference Product-owned Design Records contracts for identity, traceability, format, and placement semantics, then define concrete tool behavior locally. |
| Does BPDSL remain independent for now? | Yes. `bpdsl/records/spec/**` is already organized into DSL, MCP, and Views. PRODUCT should not absorb those contracts. |
| Should Design Records-to-BPDSL integration be deferred? | Yes. Relation semantics between Design Records and BPDSL should remain deferred until an operational requirement justifies endpoint identity, direction, validation, and owner. |
| Does temporary placement under `product/records/spec/bpdsl/` create a Design Records dependency on BPDSL? | No. Temporary placement isolates mixed content; it does not make Design Records semantics depend on BPDSL. |
| Does temporary placement create a PRODUCT ownership claim over BPDSL internals? | No. It preserves content for later review and creates no canonical ownership claim. |
| Who owns final BPDSL dependency and integration decisions? | BPDSL migration or a later explicit integration requirement. |

### Finding 9: BPDSL-specific content disposition

| BPDSL-related content currently in PRODUCT specs | disposition | reason |
|---|---|---|
| `project-artifact-model/design-flow.md` target `dsl/ -> generated src/` flow | Temporarily isolate under `product/records/spec/bpdsl/`; preserve without semantic redesign; require reclassification during BPDSL migration. | Current work lacks enough evidence to decide final ownership or integration semantics. |
| `project-artifact-model/design-flow.md` DSL/source/render/target implementation source-of-truth roles | Temporarily isolate under `product/records/spec/bpdsl/`; preserve without semantic redesign; require reclassification during BPDSL migration. | These roles may overlap app-local BPDSL specs, but the current investigation does not validate final BPDSL correctness. |
| `project-artifact-model/artifact-responsibility-matrix.md` rows for brewprint DSL definition, source implementation, render, target implementation | Temporarily isolate BPDSL-related rows under `product/records/spec/bpdsl/`; keep Design Records rows in `design-records/artifact-model/`. | Design record responsibilities can be cleaned now; BPDSL-related rows require migration review. |
| `repository-layout/index.md` `dsl/` and `src/` directory concerns | Temporarily isolate BPDSL-related implementation-flow content under `product/records/spec/bpdsl/` unless a narrow records-layout rule remains genuinely cross-cutting. | Cross-app repository layout may need a bootstrap rule, but BPDSL internals and implementation flow are not validated here. |
| `traceability/artifact-refs.md` and `traceability/index.md` reserve-only `yaml:` endpoint | Extract to follow-up investigation/requirement or preserve as evidence outside current specs. | A reserved prefix does not justify active PRODUCT ownership of BPDSL YAML refs. |
| `traceability/coverage-mapping.md` `yaml:` activation and realization chain triggers | Extract to follow-up investigation/requirement or remove after evidence transfer. | These are future integration triggers, not active contracts. |
| `namespace-model/app-namespaces.md` BPDSL architecture sketch | Keep only Brewprint profile summary; temporarily isolate BPDSL-specific detail unless already duplicated by app-local BPDSL specs. | Type system, resolution, rendering, and self-hosting semantics are expected to be BPDSL-owned but are not fully reclassified here. |
| `brewprint/layout/index.md` note that `bpdsl/` lacks app-local DSL/source placement | remain under PRODUCT profile | This is a current repository inventory fact, not canonical BPDSL semantics. |
| Any concrete BPDSL schema, validation, resolution, render, generation, or MCP behavior already clearly duplicated by an app-local BPDSL contract | move directly to `bpdsl/records/spec/**` or replace with a pointer during BPDSL migration | Direct movement is appropriate only when ownership is already plain and duplication is clear. |
| Premature Design Records-to-BPDSL realization relation | extract to follow-up investigation/requirement or remove after evidence transfer | Endpoint identity, direction, owner, and validation are not accepted yet. |

### Finding 10: reviewability model

The next implementation must be split so semantic decisions and mechanical churn do not create another 40-file semantic diff.

| stage | purpose | allowed breadth | review gate |
|---|---|---|---|
| 1. ownership/target-tree decision | Accept or revise top-level areas and boundaries. | Small. One ADR or requirement/work item, plus this investigation as source. | Must approve placement tests before file edits. |
| 2. small semantic rewrites | Remove mixed semantics inside source files before moves. | Small batches by area, preferably 3 to 7 files. | Review for responsibility leakage and stale assumptions. |
| 3. file relocation | Move files after semantic content is clean enough to have a single owner. | May touch many paths mechanically, but no semantic rewrites. | Review path changes separately from content changes. |
| 4. app-local handoff | Move or restate DRMCP/BPDSL-owned contracts in app-local specs. | Small batches per app and owner. | App-local owner reviews contract completeness. |
| 5. mechanical ref synchronization | Update refs, parent markers, topic tables, and downstream links. | May touch many files. | Parser-aware validation and scoped diff review. |
| 6. validation and stale-ref cleanup | Run validation and fix stale refs caused by accepted moves. | May touch many files if diagnostics are mechanical. | Diagnostics must be attributed to this work, not pre-existing noise. |

Recommended review limits:

- Do not combine stages 2, 3, and 5 in one change.
- Do not mix BPDSL app-local handoff with generic Design Records rewrites.
- Keep a per-stage evidence section listing exact files touched and validation commands.
- Treat broad ref synchronization as mechanical only after the target tree is already accepted.

## Cross-cutting observations

`concepts/` is not a semantic owner.
It hides Design Records semantics, Brewprint profile/current state, Brewprint compatibility/migration, PRODUCT-side BPDSL staging, and unadopted future mechanisms.

The current structure overuses "PRODUCT-owned" as a shortcut.
PRODUCT can own a governance or staging record without owning the canonical semantics inside the record.
Temporary BPDSL placement is the clearest example: it preserves mixed content for migration review, but does not validate or adopt BPDSL semantics.

`PRODUCT-WORK-NAMESPACE-001` successfully separated some DRMCP parser/scanning content, but it did not solve the broader semantic-layer problem.
`PRODUCT-INV-SPEC-004` is useful historical evidence, but its PRODUCT/DRMCP split is too coarse for the current corpus.

The traceability files contain a stale-format risk.
Several traceability sections still describe spec front matter as an MVP lookup source, while accepted spec-format policy prohibits YAML front matter for new or migrated specs.
That conflict should be handled as a semantic rewrite before any relocation.

App-local specs already exist for both app owners.
DRMCP has local specs for namespace scanning, ID normalization, resolver, schema, and tools.
BPDSL has local specs for DSL, MCP, and render views.
PRODUCT should use pointers to those specs instead of restating app-owned behavior.

Current specs should contain current contracts.
Inactive `yaml:` relations, internal-design endpoints, coverage artifacts, unrealized Design Records-to-BPDSL relations, relation lifecycle ideas, and reintroduction triggers should not be preserved as a dedicated spec area merely because they exist in current files.
Those statements should be transferred to follow-up investigations, requirements, ADR candidates, work items, or removed when obsolete or duplicated.

## Follow-up judgment candidates

| candidate judgment | why it is needed | likely artifact type |
|---|---|---|
| Adopt or revise the proposed top-level PRODUCT spec tree. | File moves and ref updates should not start until placement rules are accepted. | ADR or requirement plus work item. |
| Accept `product/records/spec/bpdsl/` as temporary isolation. | The immediate decision is a scoped quarantine mechanism with migration trigger and exit condition, not a permanent semantic owner. | ADR or requirement. |
| Decide how to reconcile traceability front matter assumptions with accepted spec-format policy. | Current traceability metadata text conflicts with visible metadata direction for migrated specs. | Requirement/work item or spec rewrite task. |
| Confirm Brewprint compatibility placement. | V01 ID retention, historical attribution, and migration state should move under `brewprint/compatibility/`, not remain inside generic namespace docs. | ADR or requirement. |
| Define BPDSL migration judgment. | Later BPDSL migration must classify staged content, decide canonical app-local ownership, delete or redefine the temporary area, and handle any actual Design Records-to-BPDSL integration requirement. | BPDSL migration requirement/work item or ADR. |
| Extract unadopted future mechanisms. | Inactive endpoints, relation candidates, and reintroduction triggers need follow-up artifacts or removal, not a dedicated current spec area. | Investigation, requirement, ADR candidate, or work item. |

## Recommendation

The current PRODUCT spec structure needs revision before any implementation or file movement.
The structure is reviewable only after `concepts/` is replaced or narrowed into explicit semantic areas.

The likely target is:

```text
product/records/spec/
  index.md

  design-records/
    index.md
    authoring-standards/
    namespace-model/
    repository-layout/
    spec-format/
    traceability/
    artifact-model/

  brewprint/
    index.md
    layout/
    namespaces/
    compatibility/

  bpdsl/
    index.md
    ...
```

`design-records/` should become the main owner for app-independent Design Records semantics.
`brewprint/` should own current-state/profile material and include `compatibility/` for Brewprint V01 compatibility, historical attribution, and migration state.
`bpdsl/` should be accepted only as a temporary PRODUCT-side quarantine/staging area with an explicit BPDSL migration trigger and exit condition.
No top-level `compatibility/` area is recommended.
No `deferred-integration/` spec area is recommended.
Future Design Records-to-BPDSL/internal-design/coverage/YAML relation material should move to follow-up artifacts or be removed after evidence transfer.

The next work should accept the ownership model first.
Then it should rewrite mixed content in small semantic batches before any broad relocation or ref synchronization.

## Follow-up artifact candidates

| candidate | responsibility |
|---|---|
| Ownership-boundary ADR | Accept or revise the top-level areas and dependency-direction model. |
| PRODUCT requirement for spec ownership cleanup | State the need to restore a reviewable PRODUCT spec boundary. |
| PRODUCT work item for staged restructuring | Coordinate semantic rewrites, relocations, app-local handoff, mechanical refs, and validation as separate stages. |
| Narrow task for root router creation | Create `product/records/spec/index.md` after the top-level areas are accepted. |
| Narrow task for namespace/profile split | Separate generic namespace model from current Brewprint app/domain registry and Brewprint compatibility policy. |
| Narrow task for Brewprint compatibility placement | Move V01 ID retention, historical attribution, and migration state under `brewprint/compatibility/`. |
| Narrow task for temporary BPDSL isolation | Move BPDSL-related implementation-model material out of generic Design Records specs without semantic redesign. |
| Narrow task for traceability reconciliation | Align traceability metadata language with accepted spec-format policy and extract inactive endpoint/relation candidates from current specs. |
| App-local handoff candidates | Move or restate clearly duplicated DRMCP/BPDSL-owned content in `drmcp/records/spec/**` and `bpdsl/records/spec/**` only when ownership is already plain or during BPDSL migration. |
| Later BPDSL migration candidate | Reclassify every staged BPDSL item, decide canonical app-local ownership, and remove or explicitly redefine the temporary PRODUCT-side BPDSL area. |

Do not allocate exact IDs for these artifacts until the next accepted requirement or work item scopes them.

## Open questions

| question | current answer |
|---|---|
| Is `design-records/` the final directory name? | It is the clearest candidate, but an ADR should accept the name before moves. |
| Where does compatibility belong? | Under `product/records/spec/brewprint/compatibility/`, because V01 compatibility, historical attribution, and migration state are Brewprint-specific history/profile concerns. |
| Should unadopted integration material become a spec area? | No. Extract it to follow-up investigation/requirement/ADR/work items, preserve historical evidence outside current specs when needed, or remove it when obsolete or duplicated. |
| Should `product/records/spec/bpdsl/` exist? | Yes only as temporary quarantine/staging with a BPDSL migration trigger and exit condition. |
| How should examples be handled in generic contracts? | Use app-neutral examples where possible; otherwise label examples as non-normative and keep project-specific facts in Brewprint profile. |
| What validates the future tree? | Existing validation may be noisy; validation scope should be declared per restructuring stage. |
| Can this investigation decide final BPDSL ownership? | No. BPDSL migration is not complete, so final BPDSL hierarchy, canonical ownership, and integration semantics remain later decisions. |
