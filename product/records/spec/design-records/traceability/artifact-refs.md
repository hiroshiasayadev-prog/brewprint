# Reference: Canonical record kinds and references

- **id**: `spec:product.design_records.traceability.artifact_refs`
- **status**: draft
- **date**: 2026-07-13
- **parent**: `spec:product.design_records.traceability`

## What this is

Defines the Product-owned Design Record kinds and canonical reference forms.

This file is the authority for record kind names and canonical reference forms.
Field-specific allowed targets belong to metadata or authoring contracts.
DRMCP owns parser behavior, resolver requests, resolver responses, diagnostics, indexing, and writer behavior.

## Current record kinds

| record kind | artifact | canonical reference form | authority | rule |
|---|---|---|---|---|
| `spec` | Specification | `spec:<app>.<path_segments>` | `spec:product.design_records.spec_format.spec_id_as_ref` | New and migrated specs use the path-derived H1-adjacent `id`. |
| `decision` | ADR / decision record | `<APP>-ADR-<DOMAIN>-<SEQUENCE>` | `spec:product.design_records.namespace_model.artifact_id_grammar` | The complete public ADR ID in H1 is the canonical ref. |
| `investigation` | Investigation record | `<APP>-INV-<DOMAIN>-<SEQUENCE>` | `spec:product.design_records.namespace_model.artifact_id_grammar` | The complete public investigation ID in H1 is the canonical ref. |
| `requirement` | Requirement record | `<APP>-REQ-<DOMAIN>-<SEQUENCE>` | `spec:product.design_records.namespace_model.artifact_id_grammar` | The complete public requirement ID in H1 is the canonical ref. |
| `work_item` | Work Item record | `<APP>-WORK-<DOMAIN>-<SEQUENCE>` | `spec:product.design_records.namespace_model.artifact_id_grammar` | The complete public Work Item ID in H1 is the canonical ref. |
| `task` | Task record | `<APP>-TASK-<DOMAIN>-<WORK_SEQUENCE>-<TASK_SEQUENCE>` | `spec:product.design_records.namespace_model.artifact_id_grammar` | The complete public Task ID in H1 is the canonical ref. |

For every sequential record kind, the public ID prefix in `# <PUBLIC-ID>: <Title>` is the sole source-internal identity authority.
H1-adjacent metadata and file names do not register or replace sequential identities.
A file-name public-ID prefix is a repository-conformance projection only.

Record kind names are lowercase contract vocabulary.
Public ID literals keep the uppercase artifact-kind segment.
Bare grammar fragments such as `REQ-*`, `WORK-*`, and `TASK-*` are notation for grammar families only.
They are not canonical refs.

## Spec refs

`spec:` refs are canonical document-level identities for new and migrated specs.

Examples:

| path | canonical ref |
|---|---|
| `product/records/spec/design-records/traceability/index.md` | `spec:product.design_records.traceability` |
| `product/records/spec/design-records/traceability/artifact-refs.md` | `spec:product.design_records.traceability.artifact_refs` |
| `product/records/spec/design-records/spec-format/spec-id-as-ref.md` | `spec:product.design_records.spec_format.spec_id_as_ref` |

Spec refs are not registered through hidden front matter.
The path-derived H1-adjacent `id` is the canonical ref.

Section refs are not active in the current contract.
A later visible-table contract is required before section-level refs become canonical.

## Field-use boundary

This file defines available record kinds and canonical reference forms.
It does not define every field-specific allowed target set.

Field-specific rules are owned by the field contract that persists the relation.

| field contract | owner |
|---|---|
| Investigation `source_refs`, `follow_up_results`, and `follow_up_candidates` | Investigation metadata and authoring contracts. |
| Work Item `source_refs` | Workflow relation metadata and Work Item authoring contracts. |
| Work Item `tasks` | Workflow relation metadata and Work Item authoring contracts. |
| Task `work_item` and `depends_on` | Workflow relation metadata and Task authoring contracts. |

Physical paths are repository locations.
They are not canonical refs.
Relations are not inferred from physical paths, file names, parent directories, or ID string structure.

## Compatibility boundary

Compatibility records may preserve legacy issued IDs as accepted inputs.
Compatibility inputs do not define new current record kinds.
Compatibility input handling is owned by `spec:product.brewprint.compatibility` and its child specs.

## Reference boundary

Only the record kinds and reference forms defined in this file are current Product reference forms.
No additional semantic prefixes or realization relations are reserved or adopted.
Historical disposition evidence is recorded in T05.

## Sources

| source | use |
|---|---|
| `spec:product.design_records.spec_format.spec_id_as_ref` | Canonical spec ref derivation. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | App-aware public ID grammar pointer. |
| `spec:product.brewprint.compatibility` | Legacy issued-ID retention pointer. |
| V01-ADR-087 | Investigation reference validation. |
| V01-ADR-088 | Canonical reference foundation and deferred endpoint evidence. |
| V01-ADR-092 | Historical workflow relation identity boundary. |
| PRODUCT-REQ-SPEC-006 | Generic workflow source-relation requirement. |
| PRODUCT-ADR-SPEC-007 | Active reference reuse for Work Item provenance. |
| PRODUCT-ADR-SPEC-008 | Legacy workflow field transition boundary. |
| PRODUCT-ADR-SPEC-019 | H1-only sequential identity authority and nonidentity projection boundary. |
