# Reference: Artifact refs

- **id**: `spec:product.design_records.traceability.artifact_refs`
- **status**: draft
- **date**: 2026-07-01
- **parent**: `spec:product.design_records.traceability`

## What this is

Defines the active canonical reference classes used by Design Records traceability. It does not reserve future semantic prefixes or define external relation endpoints.

## Active reference classes

| class | active form | rule |
|---|---|---|
| Spec ref | `spec:<app>.<path_segments>` | New and migrated specs use the path-derived H1-adjacent `id` defined by `spec:product.design_records.spec_format.spec_id_as_ref`. |
| ADR ID-as-ref | Complete public ADR ID. | Existing issued IDs remain valid; new IDs use the app-aware public ID grammar. |
| Investigation ID-as-ref | Complete public investigation ID. | Existing issued IDs remain valid; new IDs use the app-aware public ID grammar. |
| Requirement ID-as-ref | Complete public requirement ID. | Supported as a record ID-as-ref and in investigation canonical-reference fields where allowed. |
| Work item ID-as-ref | Complete public work item ID. | Supported as a record ID-as-ref and in investigation canonical-reference fields where allowed. |
| Task ID-as-ref | Complete public task ID. | Supported as a record ID-as-ref and workflow relation target, but not supported in investigation canonical-reference fields. |
| Legacy issued ID | Complete legacy public ID. | Compatibility is owned by `spec:product.brewprint.compatibility` and its child specs. |

Bare grammar fragments such as `REQ-*`, `WORK-*`, and `TASK-*` are not canonical external refs. They are notation for grammar families only.

## Spec refs

`spec:` refs are canonical document-level identities for new and migrated specs.

Examples:

| path | canonical ref |
|---|---|
| `product/records/spec/design-records/traceability/index.md` | `spec:product.design_records.traceability` |
| `product/records/spec/design-records/traceability/artifact-refs.md` | `spec:product.design_records.traceability.artifact_refs` |
| `product/records/spec/design-records/spec-format/spec-id-as-ref.md` | `spec:product.design_records.spec_format.spec_id_as_ref` |

Spec refs are not registered through hidden front matter. The path-derived H1-adjacent `id` is the canonical ref.

Section refs are not active in the current contract. A later visible-table contract is required before section-level refs become canonical.

## Record ID-as-ref

A record ID-as-ref uses the complete public ID of the target record.

| record type | canonical new form | compatibility pointer |
|---|---|---|
| ADR | `<APP>-ADR-<DOMAIN>-<SEQUENCE>` | `spec:product.brewprint.compatibility` |
| investigation | `<APP>-INV-<DOMAIN>-<SEQUENCE>` | `spec:product.brewprint.compatibility` |
| requirement | `<APP>-REQ-<DOMAIN>-<SEQUENCE>` | `spec:product.brewprint.compatibility` |
| work item | `<APP>-WORK-<DOMAIN>-<SEQUENCE>` | `spec:product.brewprint.compatibility` |
| task | `<APP>-TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>` | `spec:product.brewprint.compatibility` |

New and migrated specs use path-derived `spec:` refs rather than new `SPEC-*` public IDs. Legacy indexed `SPEC-*` public IDs may remain compatibility inputs when Brewprint compatibility records preserve them.

## Investigation reference boundary

Investigation canonical references retain this boundary:

| field | allowed canonical refs | resolution rule |
|---|---|---|
| `source_refs` | Active `spec:` refs, ADR IDs, investigation IDs, requirement IDs, and work item IDs. | Recorded refs must resolve. |
| `follow_up_results` | Active `spec:` refs, ADR IDs, investigation IDs, requirement IDs, and work item IDs. | Recorded refs must resolve. |
| `follow_up_candidates` | Canonical artifact refs when an artifact ref is written. | Not-yet-created candidates are not invalid merely because unresolved. |

Task public IDs are supported as direct resolver inputs and workflow relation targets. They are not supported in investigation canonical-reference fields.

Physical paths are not canonical references in these fields.

## Workflow relation identity

Work Item `source_refs` reuses the active reference classes defined in this file.
It does not create a Work Item-only reference taxonomy or a new semantic prefix.

| source field | target identity |
|---|---|
| `work_item.source_refs` | Any active canonical reference class in `## Active reference classes`. |
| `work_item.tasks` | Complete Task public ID. |
| `task.work_item` | Complete Work Item public ID. |
| `task.depends_on` | Complete Task public ID. |

Task records have no source relation field.
Requirement direct reverse membership is derived and is not a persisted identity field.

Relations are not inferred from physical paths, file names, parent directories, or ID string structure.

## Reference boundary

Only the reference classes defined in this file are active.
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
| PRODUCT-ADR-SPEC-007 | Active canonical reference reuse for Work Item provenance. |
| PRODUCT-ADR-SPEC-008 | Legacy workflow field transition boundary. |
