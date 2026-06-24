# Reference: Trace metadata schema

- **id**: `spec:product.design_records.traceability.metadata_schema`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records.traceability`

## What this is

Defines the current metadata and relation boundary for Design Records traceability. It points spec metadata to the accepted visible spec-format contracts and keeps investigation and workflow relation semantics separate from DRMCP implementation schemas.

## Spec metadata boundary

New and migrated specs use visible H1-adjacent metadata.

| visible marker | traceability role | owner |
|---|---|---|
| `id` | Canonical document-level `spec:` ref. | `spec:product.design_records.spec_format.spec_id_as_ref`. |
| `parent` | Canonical parent ref for topic placement. | `spec:product.design_records.spec_format.spec_id_as_ref` and `spec:product.design_records.spec_format.topics_table`. |
| `status` | Spec lifecycle marker. | Spec-format and authoring standards. |
| `date` | Creation or latest substantive contract-change date. | Spec-format and authoring standards. |
| `contract_class` | Contract subtype for `Contract` specs. | `spec:product.design_records.spec_format.document_shape`. |

The authoritative child relationship for an Index or Overview is the visible `## Topics` table with `title`, `kind`, `ref`, and `summary` columns.

This traceability spec does not copy the full spec-format contract. It relies on the accepted spec-format child specs for path-derived identity, document shape, topics tables, and validation policy.

## Removed spec front-matter schemas

The previous traceability drafts treated hidden spec front matter as the identity registry.

| removed field | final disposition | reason |
|---|---|---|
| `semantic_refs` | Delete. | Superseded by the path-derived H1-adjacent `id` contract. |
| `sections` | Delete. | No active section-ref contract exists. A visible-table contract is required before section refs become canonical. |
| Front-matter heading mappings | Delete. | Heading text is not current canonical identity and stale heading maps are not accepted. |
| Append-only semantic ref lists independent of path | Delete. | New and migrated spec refs change when paths move or rename unless a later compatibility design says otherwise. |

YAML front matter may still appear in unmigrated legacy material as validation-policy warning evidence. It is not a current metadata source of truth for new or migrated specs.

## Investigation reference metadata

Investigation metadata keeps canonical-reference rules for the accepted fields below.

| field | current traceability rule |
|---|---|
| `source_refs` | Recorded refs must use supported canonical forms and must resolve. |
| `follow_up_results` | Recorded refs must use supported canonical forms and must resolve. |
| `follow_up_candidates` | Artifact refs must use canonical form when written. A not-yet-created candidate is not invalid merely because unresolved. |

Supported workflow record IDs in these investigation fields are requirement and work item IDs. Task IDs are not supported in investigation canonical-reference fields.

The full investigation metadata field set, lifecycle, and authoring form belong to investigation authoring standards. DRMCP owns parser and concrete response behavior.

## Workflow relation metadata

Workflow artifact relations are declared through complete public IDs in metadata fields.

| source artifact | field | target |
|---|---|---|
| requirement | `work_items` | Complete work item public ID. |
| work item | `source_requirement` | Complete requirement public ID. |
| work item | `tasks` | Complete task public ID. |
| task | `work_item` | Complete work item public ID. |
| task | `source_requirement` | Complete requirement public ID. |
| task | `depends_on` | Complete task public ID. |

The current semantic checks are:

| condition | PRODUCT-owned invalid condition |
|---|---|
| Missing target | A declared workflow relation points to a non-existent supported target. |
| Requirement/work item mismatch | `requirement.work_items` and `work_item.source_requirement` disagree. |
| Work item/task mismatch | `work_item.tasks` and `task.work_item` disagree. |
| Source requirement mismatch | A task's `source_requirement` differs from its parent work item's `source_requirement`. |
| Noncanonical relation value | A physical path, bare grammar fragment, or semantic prefix is used where a public ID is required. |

Task dependency existence is in scope. Dependency-cycle detection and execution-order projection are not current PRODUCT traceability semantics.

## Metadata boundary

No internal-design canonical metadata contract is current.
No external coverage metadata contract is current.
No BPDSL YAML endpoint metadata contract is current.
No DRMCP writer schema is owned by PRODUCT.
Historical disposition evidence is recorded in T05.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.spec_format.spec_id_as_ref` | Canonical spec identity owner. |
| `spec:product.design_records.spec_format.document_shape` | Visible metadata and section-shape owner. |
| `spec:product.design_records.spec_format.topics_table` | Authoritative child-topic relationship owner. |
| `spec:product.design_records.traceability.artifact_refs` | Supported reference classes. |
| `spec:product.design_records.traceability.resolve_and_validation` | Lookup sources and invalid condition boundary. |
