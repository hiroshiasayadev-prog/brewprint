# Reference: Trace metadata schema

- **id**: `spec:product.design_records.traceability.metadata_schema`
- **status**: draft
- **date**: 2026-07-01
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

Current persisted workflow relations are:

| source artifact | field | target or meaning |
|---|---|---|
| work item | `source_refs` | Required non-empty unordered set of active canonical refs for direct material sources. |
| work item | `tasks` | Complete Task public IDs owned by the Work Item. |
| task | `work_item` | Complete owning Work Item public ID. |
| task | `depends_on` | Complete prerequisite Task public IDs. |

Task records persist no source field.
Requirement records persist no `work_items` field.

Work Item `source_refs` semantics:

| condition | contract |
|---|---|
| Cardinality | At least one ref is required. |
| Ordering | Order has no semantic meaning. Reordering alone does not change provenance. |
| Reference classes | Every entry uses an active canonical reference class from `spec:product.design_records.traceability.artifact_refs`. |
| Selection | Include every direct material source. Exclude incidental context and merely transitive ancestors unless independently material. |
| Duplicates | Duplicate canonical refs are invalid. Persistence does not silently deduplicate them. |
| Self-reference | A Work Item must not include its own canonical identity. |
| Resolution | Every entry must resolve. Unresolved and unrecognized refs are invalid. |

The direct Requirement reverse relation is derived from Work Item metadata.
It is the unordered, duplicate-free set of Work Items whose `source_refs` directly contain the Requirement ID.
Transitive descendants are excluded from the direct reverse relation.

Work Item `tasks` and Task `work_item` remain the explicit ownership relation.
Task `depends_on` remains the explicit Task dependency relation.

### Migration state

Repository-wide staged migration is allowed.
Each record transitions atomically.

| artifact | transition contract |
|---|---|
| Work Item | Convert `source_requirement` to a one-element `source_refs` list and remove the old field in the same update. Infer no additional source. |
| Task | Remove `source_requirement` without replacement. Preserve `work_item` and `depends_on`. Do not create Task `source_refs`. |
| Requirement | Remove `work_items` only after exact equality with the derived direct reverse set. Compare unordered, duplicate-free sets. A mismatch blocks migration. |

A Work Item must not persist both `source_requirement` and `source_refs`.
Migration must not silently repair relation mismatches.
The repository may contain migrated and unmigrated records during staged migration.

## Metadata boundary

No internal-design canonical metadata contract is current.
No external coverage metadata contract is current.
No BPDSL YAML endpoint metadata contract is current.
No DRMCP writer schema is owned by PRODUCT.
PRODUCT owns the persisted relation meaning, invalid conditions, and migration semantics.
DRMCP owns parsing, indexing, reverse lookup, traversal, Task-owner resolution, diagnostics, response schemas, and projections.
Historical disposition evidence is recorded in T05.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.spec_format.spec_id_as_ref` | Canonical spec identity owner. |
| `spec:product.design_records.spec_format.document_shape` | Visible metadata and section-shape owner. |
| `spec:product.design_records.spec_format.topics_table` | Authoritative child-topic relationship owner. |
| `spec:product.design_records.traceability.artifact_refs` | Supported reference classes. |
| `spec:product.design_records.traceability.resolve_and_validation` | Lookup sources and invalid condition boundary. |
| PRODUCT-REQ-SPEC-006 | Generic workflow source-relation requirement. |
| PRODUCT-ADR-SPEC-007 | Canonical Work Item provenance and direct Requirement reverse relation. |
| PRODUCT-ADR-SPEC-008 | Staged atomic migration contract. |
