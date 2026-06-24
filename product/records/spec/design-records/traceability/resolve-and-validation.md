# Reference: Resolve and validation

- **id**: `spec:product.design_records.traceability.resolve_and_validation`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records.traceability`

## What this is

Defines PRODUCT-owned canonical lookup sources and invalid conditions for Design Records traceability. It excludes DRMCP request shape, response shape, diagnostic names, parser behavior, persistence, indexing, UI, tool APIs, and writer behavior.

## Supported canonical inputs

| input class | current rule |
|---|---|
| `spec:` refs | New and migrated specs use path-derived document-level refs from H1-adjacent `id`. |
| Legacy issued record IDs | Complete public legacy IDs may remain compatibility inputs through Brewprint compatibility records. |
| App-aware ADR IDs | Complete public ADR IDs are supported record ID-as-refs. |
| App-aware investigation IDs | Complete public investigation IDs are supported record ID-as-refs. |
| App-aware requirement IDs | Complete public requirement IDs are supported record ID-as-refs. |
| App-aware work item IDs | Complete public work item IDs are supported record ID-as-refs. |
| App-aware task IDs | Complete public task IDs are supported record ID-as-refs and workflow relation targets. |

Physical paths are not canonical inputs for traceability relations. Parent workflow relations are not inferred from ID string structure or file layout.

## Lookup sources

| lookup source | role |
|---|---|
| Spec H1-adjacent `id` | Registers the canonical document-level `spec:` ref for new and migrated specs. |
| Spec path-derived mapping | Confirms the visible `id` matches the canonical ref derived from file path. |
| Record public ID | Registers ADR, investigation, requirement, work item, and task identities. |
| Brewprint compatibility records | Preserve legacy issued IDs and migration compatibility pointers. |
| Investigation metadata | Supplies referring refs in `source_refs`, recorded `follow_up_results`, and artifact refs in `follow_up_candidates`. |
| Workflow metadata | Supplies declared relation values for requirement, work item, and task integrity checks. |

Investigation metadata and workflow relation fields are referring sides. They do not register new reference targets.

Natural-language body text is not a lookup source.

## Spec resolution boundary

Spec resolution uses the path-derived document-level `spec:` ref.

| condition | current rule |
|---|---|
| H1-adjacent `id` matches path-derived ref | Valid canonical spec identity. |
| H1-adjacent `id` differs from path-derived ref | Invalid for new or migrated specs under the accepted spec-format contract. |
| Physical path appears as a relation value | Noncanonical. |
| Section ref appears without accepted visible-table contract | Not active; do not resolve as a current canonical target. |

The old front-matter `semantic_refs` and `sections` lookup model is obsolete for new and migrated specs.

## Investigation validation

| field | condition |
|---|---|
| `source_refs` | Supported canonical refs must resolve. Physical paths are noncanonical. |
| `follow_up_results` | Supported canonical refs must resolve. Physical paths are noncanonical. |
| `follow_up_candidates` | Artifact refs must use canonical form when written. Unresolved not-yet-created candidates are allowed. |

Requirement and work item IDs are supported in these investigation fields. Task IDs are not supported in investigation canonical-reference fields.

DRMCP owns any concrete diagnostic label or response representation for these conditions.

## Workflow relation validation

| condition | invalid when |
|---|---|
| Target existence | A declared requirement, work item, or task relation points to a missing target. |
| Requirement/work item reciprocity | `requirement.work_items` and `work_item.source_requirement` do not agree. |
| Work item/task reciprocity | `work_item.tasks` and `task.work_item` do not agree. |
| Task source requirement | `task.source_requirement` does not match the parent work item's `source_requirement`. |
| Task dependency target | `task.depends_on` points to a missing task. |
| Noncanonical value | A physical path, bare grammar fragment, or unsupported semantic prefix is used as a relation value. |

These checks validate declared relations. They do not search for orphan artifacts, project progress, traversal trees, cycles, or execution order.

## Duplicate identity conditions

| identity class | invalid condition |
|---|---|
| Spec ref | More than one new or migrated spec path has the same canonical `spec:` ref. |
| Record public ID | More than one record carries the same complete public ID. |
| Compatibility ID | A legacy issued ID maps ambiguously through compatibility records. |

Duplicate identity conditions for `coverage:`, `COV-*`, `internal-design:`, `yaml:`, and `fixture:` are outside the current traceability contract.

## Excluded implementation behavior

| excluded concern | owner |
|---|---|
| Request shape | DRMCP. |
| Response shape and status vocabulary | DRMCP. |
| Diagnostic category names and severities | DRMCP. |
| Parser behavior | DRMCP. |
| Persistence and indexing implementation | DRMCP. |
| UI and tool APIs | DRMCP. |
| Writer behavior, dry-run diff, confirmation, conflict handling, format preservation, and permission boundaries | DRMCP app-local specifications. |

## Resolve and validation boundary

Active resolve and validation contracts are defined by the sections above.
No `yaml:`, `internal-design:`, `coverage:`, `COV-*`, or fixture lookup is active.
Workflow orphan diagnostics, progress projection, traversal, cycle detection, and execution-order checks belong to DRMCP app-local specifications.
Historical disposition evidence is recorded in T05.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.traceability.artifact_refs` | Supported reference classes. |
| `spec:product.design_records.traceability.metadata_schema` | Metadata and relation fields. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Spec lookup source and path-derived identity. |
| `spec:product.brewprint.compatibility` | Legacy issued-ID compatibility pointer. |
