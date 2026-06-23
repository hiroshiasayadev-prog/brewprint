# PRODUCT-TASK-SPEC-011-08: Verify spec authoring and cross-guide consistency

- **id**: PRODUCT-TASK-SPEC-011-08
- **status**: done
- **date**: 2026-06-23
- **work_item**: PRODUCT-WORK-SPEC-011
- **source_requirement**: PRODUCT-REQ-SPEC-002
- **estimate**: 0.7d
- **depends_on**:
  - PRODUCT-TASK-SPEC-011-03
  - PRODUCT-TASK-SPEC-011-04
  - PRODUCT-TASK-SPEC-011-05
  - PRODUCT-TASK-SPEC-011-06
  - PRODUCT-TASK-SPEC-011-07
- **outputs**:
  - Spec-authoring coverage result (inline in Evidence)
  - Cross-guide consistency review findings (inline in Evidence)
  - `product/records/work-items/spec/PRODUCT-WORK-SPEC-011-per-artifact-authoring-guide-migration.md`

## Goal

Verify spec-authoring coverage and review all authoring standards for cross-guide consistency before closing PRODUCT-WORK-SPEC-011.

## Work

### Spec-authoring coverage

Review `product/records/guides/spec-authoring.md` against PRODUCT-REQ-SPEC-002 and the current PRODUCT spec-format contract.

Check:

- H1 and metadata rules;
- path-derived canonical spec refs;
- spec kind selection;
- required section rules;
- Topics table rules;
- front matter policy;
- author-facing create inputs;
- English structural headings;
- PRODUCT / DRMCP ownership boundary.

Record either:

- coverage is complete; or
- concrete gaps that require follow-up outside authoring standards.

Do not rewrite or relocate the spec-authoring guide under this task.

### Cross-guide consistency

Review the shared boundary and all per-artifact guides.

Check:

- common section structure;
- English headings and table headers;
- ID grammar and `new` forms;
- sequence scope;
- file path rules;
- metadata terminology and requiredness states;
- status lifecycle terminology;
- semantic date definitions;
- canonical reference policy;
- artifact ownership boundaries;
- author-supplied and generated values;
- exclusion of current DRMCP operating status;
- citation consistency.

Apply only must-fix corrections needed for internal consistency. Record broader follow-up work outside authoring standards.

## Done condition

| item | done when |
|---|---|
| Spec coverage reviewed | Coverage is recorded as complete or concrete external follow-ups are identified. |
| All guides reviewed | Boundary, ADR, requirement, work item, task, and investigation guides are checked. |
| Drift resolved | No ownership, terminology, metadata, lifecycle, ID grammar, heading, or citation drift remains. |
| Index complete | The authoring-standards Index lists the boundary and all per-artifact guides. |
| Work item updated | PRODUCT-WORK-SPEC-011 Evidence records the final review result. |

## Verification

- Confirm every per-artifact guide uses the common structure.
- Confirm all canonical structural headings use English.
- Confirm each guide cites PRODUCT semantic sources and leaves concrete tool contracts to DRMCP.
- Confirm all Index refs resolve to existing files.
- Confirm no guide records current DRMCP operating status or implementation plans.
- Confirm PRODUCT-WORK-SPEC-011 completion conditions are fully evaluated.

## Evidence

### Spec-authoring coverage

`product/records/guides/spec-authoring.md` reviewed against PRODUCT-REQ-SPEC-002 and the current PRODUCT spec-format contract.

Coverage is complete. All required topics are present:

- H1 format and ATX H1 rule: ✅ defined with kind table and example
- H1-adjacent metadata: ✅ defined with marker table and `contract_class` rule
- Path-derived canonical spec refs: ✅ full derivation rules with examples
- Spec kind selection: ✅ Overview / Index / Concept / Reference / Contract with deferred kinds
- Required section matrix by kind: ✅ full table by kind
- `## Topics` table rules: ✅ required columns, `file` column prohibition, example
- Front matter policy: ✅ prohibited for new and migrated specs with item-by-item severity table
- Author-facing create inputs: ✅ implicitly covered by the metadata and body rules; no separate Create section needed (guide form, not Reference spec)
- English structural headings: ✅ all headings in English
- PRODUCT / DRMCP ownership boundary: ✅ guide references `spec:product.concepts.spec_format` and its children; front matter policy notes DRMCP compatibility gap as a separate concern

No rewrite, relocation, or follow-up within authoring standards required.

### Cross-guide consistency

Guides reviewed: artifact-boundary, adr-authoring, requirement-authoring, work-item-authoring, task-authoring, investigation-authoring.

All checks pass after applying F1:

- Common section structure: ✅ all five per-artifact guides use the prescribed section order
- English headings and table headers: ✅ all guides
- ID grammar and `new` forms: ✅ all guides define kind-specific abstract grammar and `new` placeholder
- Sequence scope: ✅ REQ, WORK, INV, and ADR use `app namespace + artifact kind + domain namespace`; TASK is scoped by its parent Work Item
- File path rules: ✅ ADR, INV, REQ, WORK, and TASK use domain subdirectories; ADR uses `<app>/records/adr/<domain>/`
- Metadata terminology: ✅ create input / partial update / persisted consistent across all guides
- Status lifecycle: ✅ kind-appropriate; no shared terms conflict
- Semantic date definitions: ✅ each guide defines distinct decision-validity or scope-change semantics
- Canonical reference policy: ✅ consistent three-rule base across all guides with kind-specific additions
- Artifact ownership boundaries: ✅ all guides cite `artifact_boundary` and `artifact_responsibility_matrix`
- Author-supplied vs generated values: ✅ ADR and INV have no `id` in metadata; REQ/WORK/TASK have `id` generated by MCP; all guides state what the author does and does not supply
- Exclusion of current DRMCP operating status: ✅ all Non-goals sections list this
- Citation consistency: ✅ all guides cite the same eight `spec:` refs plus PRODUCT-REQ-SPEC-002 and PRODUCT-WORK-SPEC-011

Corrections applied:
- `requirement-authoring.md`, `work-item-authoring.md`, and `investigation-authoring.md`: sequence scope made explicit as app + kind + domain.
- `adr-authoring.md`: sequence scope made explicit as app + kind + domain; path changed to `<app>/records/adr/<domain>/...`.
- `v2-grammar.md`: canonical non-task sequence allocation scope added.
- `record-discovery-paths.md`: ADR discovery changed to the domain-subdirectory pattern.

### Index completeness

`authoring-standards/index.md` contains six new entries:
- Artifact boundary (1)
- ADR authoring (2)
- Requirement authoring (3)
- Work-item authoring (4)
- Task authoring (5)
- Investigation authoring (6)

All refs match path-derived canonical forms. ✅

### Work item updated

PRODUCT-WORK-SPEC-011 Evidence updated and status set to `done`.
