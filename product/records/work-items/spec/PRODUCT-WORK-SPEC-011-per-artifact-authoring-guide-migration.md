# PRODUCT-WORK-SPEC-011: Per-artifact authoring guide migration to authoring-standards

- **id**: PRODUCT-WORK-SPEC-011
- **status**: done
- **date**: 2026-06-22
- **source_requirement**: PRODUCT-REQ-SPEC-002
- **impact_refs**:
  - spec:product.concepts.authoring_standards
  - spec:product.concepts.project_artifact_model.artifact_responsibility_matrix
  - spec:product.concepts.namespace_model.artifact_id_grammar
- **tasks**:
  - PRODUCT-TASK-SPEC-011-01
  - PRODUCT-TASK-SPEC-011-02
  - PRODUCT-TASK-SPEC-011-03
  - PRODUCT-TASK-SPEC-011-04
  - PRODUCT-TASK-SPEC-011-05
  - PRODUCT-TASK-SPEC-011-06
  - PRODUCT-TASK-SPEC-011-07
  - PRODUCT-TASK-SPEC-011-08

## Goal

Publish a lightweight authoring-time artifact boundary standard and English per-artifact authoring guides for ADR, requirement, work item, task, and investigation.

The authoring boundary standard covers only DRMCP-managed design records:

- ADR
- spec
- investigation
- requirement
- work item
- task

It provides the minimum artifact-selection and ownership guidance needed before authoring. Canonical artifact ownership remains defined by `spec:product.concepts.project_artifact_model.artifact_responsibility_matrix`.

Each per-artifact guide defines how authors should write that artifact. The guides do not document current DRMCP implementation status, migration gaps, follow-up work, or redesign plans.

Verify that the existing spec-authoring guide at `product/records/guides/spec-authoring.md` satisfies PRODUCT-REQ-SPEC-002 for spec authoring.

Correct the v2 ADR grammar recording omission identified during scope review.

Closes the per-artifact authoring scope of PRODUCT-REQ-SPEC-002.

## Decisions

| topic | decision | reason |
|---|---|---|
| guide kind | Use `Reference` for the authoring boundary and all per-artifact guides. | These specs define fixed authoring rules, field semantics, lifecycle rules, and author-facing input requirements. They do not define an executable procedure. |
| canonical artifact ownership | `spec:product.concepts.project_artifact_model.artifact_responsibility_matrix` remains the canonical ownership source. | The project artifact model covers the complete brewprint artifact system, including BPDSL and implementation-facing artifacts. |
| authoring-time boundary projection | Add `spec:product.concepts.authoring_standards.artifact_boundary` as a lightweight projection for DRMCP-managed design records. | Authors need a small artifact-selection reference without loading the complete project artifact model. |
| projection authority | `artifact_boundary` summarizes the canonical responsibility matrix. It does not redefine canonical ownership. Conflicts are resolved in favor of the project artifact model. | This permits operational duplication without creating two equal sources of truth. |
| per-artifact boundary | Each per-artifact guide cites `artifact_boundary` and records only kind-specific writing boundaries and commonly confused adjacent artifacts. | Full ownership matrices must not be copied into every guide. |
| authoring standard scope | Authoring standards define how artifacts should be written. | Current implementation status, migration state, identified gaps, follow-up requirements, and redesign plans belong to other artifacts. |
| common guide shape | All per-artifact guides use the same section structure. | A stable structure reduces cross-guide drift and makes review mechanical. |
| heading language | All structural headings prescribed by the guides, including H1 labels and required section names, use English. | The documentation language policy applies to document structure as well as prose. Japanese legacy headings must not remain canonical forms. |
| canonical metadata sources | Cite PRODUCT semantic ownership specs and DRMCP metadata grammar where applicable. Authoring guides define author-facing requiredness and field meaning. | Parser recognition and authoring requirements are different responsibilities. |
| create and update requirements | Guides distinguish create input, partial update input, and required persisted record state where these differ. | A field may be required in the completed record without being required in every update request. |
| semantic dates | Each guide defines the semantic meaning of its `date` field. Dates are not treated as automatic modification timestamps unless explicitly specified. | Editorial changes and semantic changes must not be conflated. |
| ID grammar | Each guide defines its kind-specific abstract grammar using `<APP>`, `<DOMAIN>`, and sequence placeholders. | Authoring standards describe the reusable brewprint model, not the V01 repository instance. |
| legacy examples | V01 identifiers appear only in compatibility or migration notes when necessary. | Legacy identifiers must not appear as the primary authoring form. |
| authoring interface requirements | Guides define author-facing inputs and generated values abstractly. They do not reproduce the complete DRMCP JSON interface contract. | Authoring standards own what authors provide. DRMCP specs own concrete request, response, and diagnostic contracts. |
| `new` placeholder | Supported guides may define `<APP>-<KIND>-<DOMAIN>-new` as the author-facing unresolved ID form. | Authors should provide app and domain identity without manually allocating a sequence. |
| DRMCP operating status | Do not record whether DRMCP is currently operational in authoring standards. | Operational state changes independently from normative authoring rules. |

## Common per-artifact guide shape

Each per-artifact guide uses the following structure:

```markdown
## What this is
## Non-goals
## Rules

### ID grammar
### File path layout
### File shape
### Metadata schema
### Status lifecycle
### Kind-specific authoring rules
### Canonical reference policy

## Authoring interface requirements
## Related specs
```

Rules:

- `Metadata schema` distinguishes field meaning, create requirements, update requirements, and persisted-record requirements when necessary.
- `Kind-specific authoring rules` contains only rules unique to the artifact kind.
- H1 labels and required section headings defined by each guide use English.
- Japanese legacy headings may appear only in migration notes, never as canonical forms.
- Cross-artifact ownership is referenced from `artifact_boundary`.
- Current tool limitations, migration notes, follow-up candidates, and implementation work are excluded.
- Concrete examples use current abstract or v2 forms rather than V01 forms.
- Complete DRMCP request and response schemas are not duplicated.
- Guide headings, prescribed artifact headings, table headers, and normative heading examples use English.
- Legacy non-English headings may appear only in compatibility or migration notes.

## Artifact boundary standard

Create `spec:product.concepts.authoring_standards.artifact_boundary`.

The standard covers:

| artifact | authoring-time purpose |
|---|---|
| ADR | Record an adopted design decision and its rationale. |
| spec | Record the currently valid specification or contract. |
| investigation | Record research, evidence, uncertainty, and options before a decision. |
| requirement | Record a stable need, gap, or requested outcome. |
| work item | Record the complete resolution flow for a requirement. |
| task | Record concrete short-term work and its verification. |

The standard must:

- identify the canonical responsibility matrix;
- state that it is a lightweight authoring-time projection;
- state that the canonical source wins when wording conflicts;
- cover artifact selection before content is written;
- distinguish commonly confused adjacent artifacts;
- exclude BPDSL YAML, render artifacts, internal design, target implementation, and other non-DRMCP artifacts;
- avoid metadata, file-shape, and status rules owned by per-artifact guides.

## ADR-specific decisions

The ADR guide must include the following rules.

### Date meaning

`date` records when the documented decision became valid.

- It is required when creating an ADR.
- It is optional in a partial update request.
- It remains required in the persisted ADR.
- Update it when the decision meaning changes.
- Do not update it for spelling, formatting, reference repair, or other editorial corrections.
- An accepted decision that is overturned should normally be superseded by another ADR rather than rewritten.

### Required ADR body headings

The ADR guide defines these English body headings:

- `## Context`
- `## Decision`
- `## Rationale`
- `## Rejected alternatives`
- `## Consequences`
- `## Evidence`

`Rejected alternatives` may be optional when no alternatives were considered. The guide must not prescribe Japanese headings as the current form.

### Evidence and specification boundary

- Observed implementation, fixture, or example shapes are evidence.
- Observed shapes are not normative specifications.
- Separate the observed fact from the resulting design decision.
- Current contract text belongs to a spec.
- Detailed comparison matrices and exploration logs belong to an investigation.
- ADR may record decision-relevant alternatives and concise rejection reasons.

### Author-facing ID form

```text
<APP>-ADR-<DOMAIN>-<NNN>
<APP>-ADR-<DOMAIN>-new
```

The author supplies:

- app namespace;
- domain namespace;
- title;
- author-controlled metadata;
- content sections.

The author does not supply:

- the resolved sequence when using `new`;
- a generated H1;
- a generated file path.

## Boundary

In scope:

- Correct `spec:product.concepts.namespace_model.v2_grammar` for new ADR identifiers.
- Create the lightweight DRMCP authoring boundary standard.
- Create five per-artifact Reference specs:
  - `adr-authoring`
  - `requirement-authoring`
  - `work-item-authoring`
  - `task-authoring`
  - `investigation-authoring`
- Apply the common guide shape across all five guides.
- Define all guide headings and prescribed artifact body headings in English.
- Update `authoring-standards/index.md`.
- Verify spec-authoring coverage.
- Perform a cross-guide consistency review.

Out of scope:

- Changes to canonical ownership in the project artifact model.
- BPDSL or implementation-artifact authoring guidance.
- Migration of `v01/records/guides/artifact-boundary.md`.
- Rewrite or relocation of `product/records/guides/spec-authoring.md`.
- DRMCP implementation changes.
- DRMCP interface-contract updates.
- Requirements or work items derived from identified DRMCP gaps.
- Migration of existing workflow records.
- Bulk rewrite of existing records.
- Operational status documentation.
- Follow-up planning inside authoring standards.

## Impact Scope

| ref | impact |
|---|---|
| `spec:product.concepts.authoring_standards` | Add one shared boundary standard and five per-artifact guides. |
| `spec:product.concepts.project_artifact_model.artifact_responsibility_matrix` | Referenced as canonical ownership source; no modification required unless review finds an actual canonical defect. |
| `spec:product.concepts.namespace_model.v2_grammar` | Correct new ADR grammar. |
| `PRODUCT-REQ-SPEC-002` | Satisfy ADR, requirement, work item, task, and investigation authoring outcomes; verify spec coverage. |

## Task flow

1. Correct the v2 ADR grammar. Review gate.
2. Create the lightweight authoring boundary standard and define the common per-artifact guide shape. Update the authoring-standards index. Review gate.
3. Revise and finalize `adr-authoring.md` as the canonical per-artifact example. Review gate.
4. Create `requirement-authoring.md`. Update the index. Review gate.
5. Create `work-item-authoring.md`. Update the index. Review gate.
6. Create `task-authoring.md`. Update the index. Review gate.
7. Create `investigation-authoring.md`. Update the index. Review gate.
8. Verify spec-authoring coverage and perform cross-guide consistency review.

## Task Candidates

| task | scope |
|---|---|
| PRODUCT-TASK-SPEC-011-01 | Correct v2 ADR grammar. |
| PRODUCT-TASK-SPEC-011-02 | Create authoring-time artifact boundary and common guide shape. |
| PRODUCT-TASK-SPEC-011-03 | Revise and finalize ADR authoring guide. |
| PRODUCT-TASK-SPEC-011-04 | Create requirement authoring guide. |
| PRODUCT-TASK-SPEC-011-05 | Create work-item authoring guide. |
| PRODUCT-TASK-SPEC-011-06 | Create task authoring guide. |
| PRODUCT-TASK-SPEC-011-07 | Create investigation authoring guide. |
| PRODUCT-TASK-SPEC-011-08 | Verify spec-authoring coverage and perform cross-guide review. |

## Completion Condition

- `spec:product.concepts.namespace_model.v2_grammar` defines the new ADR form as `<APP>-ADR-<DOMAIN>-<NNN>`.
- `spec:product.concepts.authoring_standards.artifact_boundary` is published as a lightweight authoring-time projection.
- The artifact boundary cites the canonical project artifact responsibility matrix.
- The artifact boundary states precedence when the projection and canonical source conflict.
- Five per-artifact Reference specs are published.
- All five guides use the common section structure.
- All H1 labels and required section headings defined by the guides use English.
- No Japanese legacy heading remains a canonical heading form.
- Each guide defines author-facing rules without recording implementation status, follow-up work, or migration plans.
- Each guide distinguishes create, update, and persisted-state requirements where necessary.
- Each guide uses abstract v2 identifiers as primary forms.
- The ADR guide prescribes `Context`, `Decision`, `Rationale`, `Rejected alternatives`, `Consequences`, and `Evidence` as its body headings.
- The ADR guide defines `date` as the semantic decision-validity date.
- The ADR guide distinguishes evidence, decision, investigation content, and current specification content.
- The authoring-standards index contains six new entries: one boundary standard and five per-artifact guides.
- Cross-guide review reports no ownership, terminology, metadata, lifecycle, ID grammar, or citation drift.
- Spec-authoring coverage is recorded as complete or its gaps are recorded outside the authoring standards.

## Evidence

All completion conditions satisfied as of PRODUCT-TASK-SPEC-011-08.

**Namespace model:**
- `spec:product.concepts.namespace_model.v2_grammar` corrected for new ADR identifiers (`<APP>-ADR-<DOMAIN>-<NNN>`) in T01.

**Authoring boundary standard:**
- `spec:product.concepts.authoring_standards.artifact_boundary` published as a Reference spec in T02.
- Cites `spec:product.concepts.project_artifact_model.artifact_responsibility_matrix` as the canonical ownership source.
- States precedence when the lightweight projection conflicts with the canonical source.
- Covers all six DRMCP-managed design record kinds.
- Excludes BPDSL YAML, render artifacts, internal design, and target implementation.

**Per-artifact guides:**
- `adr-authoring.md`: finalized in T03. Defines `<APP>-ADR-<DOMAIN>-<NNN>`, Context/Decision/Rationale/Rejected alternatives/Consequences/Evidence sections in English, `date` as decision-validity date, evidence and specification ownership boundary.
- `requirement-authoring.md`: created in T04. Defines `<APP>-REQ-<DOMAIN>-<NNN>` with app + kind + domain sequence scope, five body sections, `id` generated, `source_refs` and `work_items` list fields.
- `work-item-authoring.md`: created in T05. Defines `<APP>-WORK-<DOMAIN>-<NNN>` with app + kind + domain sequence scope, seven always-present body sections, `id` generated, `source_requirement` correction-only, persisted workflow relation invariants.
- `task-authoring.md`: created in T06. Defines `<APP>-TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>`, five always-present sections all done-gated, `id` and `work_item` generated, task granularity guidance.
- `investigation-authoring.md`: created in T07. Defines `<APP>-INV-<DOMAIN>-<NNN>` with app + kind + domain sequence scope, ten always-present body sections, no `id` in bullet metadata (ADR pattern), investigation-specific metadata (trigger, scope, non_scope, follow_up_candidates, follow_up_results, supersedes as list), conclusion-readiness rule.

**Common structure and heading language:**
- All five guides use the common section structure from PRODUCT-WORK-SPEC-011 § Common per-artifact guide shape.
- All H1 labels and required section headings use English. No Japanese legacy heading remains canonical.

**Author-facing inputs and DRMCP boundary:**
- Each guide distinguishes create input, partial update, and persisted-state requirements.
- Each guide uses abstract v2 identifiers as primary forms.
- No guide records current DRMCP operating status, implementation state, migration gaps, or follow-up work.
- Concrete DRMCP request/response/diagnostic schemas excluded from all guides.

**Authoring-standards index:**
- Six new entries added (artifact boundary + five per-artifact guides). All refs resolve to existing files.

**Cross-guide consistency (T08):**
- Spec-authoring coverage at `product/records/guides/spec-authoring.md` confirmed complete.
- Sequence allocation scope aligned across REQ, WORK, INV, and ADR as app namespace + artifact kind + domain namespace.
- ADR placement aligned with other domain-scoped sequential artifacts: `<app>/records/adr/<domain>/...`.
- `v2-grammar.md` and `record-discovery-paths.md` updated as canonical sources.
- No remaining ownership, terminology, metadata, lifecycle, ID grammar, path, heading, or citation drift.
