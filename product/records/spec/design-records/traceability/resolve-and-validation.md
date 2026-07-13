# Reference: Resolve and validation

- **id**: `spec:product.design_records.traceability.resolve_and_validation`
- **status**: draft
- **date**: 2026-07-13
- **parent**: `spec:product.design_records.traceability`

## What this is

Defines PRODUCT-owned canonical lookup sources and invalid conditions for Design Records traceability. It excludes DRMCP request shape, response shape, diagnostic names, parser behavior, persistence, indexing, UI, tool APIs, and writer behavior.

## Supported canonical inputs

| input kind | current rule |
|---|---|
| `spec:` refs | New and migrated specs use path-derived document-level refs from H1-adjacent `id`. |
| Legacy issued record IDs | Complete public legacy IDs may remain compatibility inputs through Brewprint compatibility records. |
| App-aware ADR IDs | Complete public ADR IDs are supported public ID refs. |
| App-aware investigation IDs | Complete public investigation IDs are supported public ID refs. |
| App-aware requirement IDs | Complete public requirement IDs are supported public ID refs. |
| App-aware work item IDs | Complete public work item IDs are supported public ID refs. |
| App-aware task IDs | Complete public Task IDs are supported public ID refs and workflow relation targets. |

Physical paths are not canonical inputs for traceability relations. Parent workflow relations are not inferred from ID string structure or file layout.

## Lookup sources

| lookup source | role |
|---|---|
| Spec H1-adjacent `id` | Registers the canonical document-level `spec:` ref for new and migrated specs. |
| Spec path-derived mapping | Confirms the visible `id` matches the canonical ref derived from file path. |
| Sequential record H1 public ID | Registers ADR, investigation, requirement, work item, and task identities from the complete public ID prefix in H1. |
| Brewprint compatibility records | Preserve legacy issued IDs and migration compatibility pointers. |
| Investigation metadata | Supplies referring refs in `source_refs`, recorded `follow_up_results`, and artifact refs in `follow_up_candidates`. |
| Workflow metadata | Supplies declared relation values for requirement, work item, and task integrity checks. |

Investigation metadata and workflow relation fields are referring sides. They do not register new reference targets.

Natural-language body text is not a lookup source.
Sequential H1-adjacent metadata and file names are not lookup sources.

## Sequential record identity boundary

Sequential identity is determined only from the public ID prefix in `# <PUBLIC-ID>: <Title>`.

| condition | identity result | conformance result |
|---|---|---|
| H1 public ID is missing or malformed for the corpus-selected artifact kind | No current identity is registered. | Identity-determination failure. |
| H1 app namespace differs from the configured app namespace | No current identity is registered. | Identity-determination failure. |
| H1 artifact kind differs from the corpus-selected artifact kind | No current identity is registered. | Identity-determination failure. |
| H1 domain namespace differs from the physical domain directory | No current identity is registered. | Identity-determination failure. |
| File-name public-ID prefix differs from the valid H1 public ID | The H1 identity remains registered. | Nonidentity repository-conformance violation. |
| Prohibited H1-adjacent metadata `id` is present with a valid H1 public ID | The H1 identity remains registered. | Nonidentity metadata-conformance violation. |

Implementations do not repair, complete, normalize, infer, or replace sequential identity from metadata or file names.

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
| Work Item source cardinality | `source_refs` is missing or empty. |
| Source target existence | A Work Item `source_refs` entry does not resolve. |
| Source canonical identity | A Work Item `source_refs` entry uses an unrecognized or noncanonical ref form. |
| Duplicate source | The same canonical ref appears more than once in one Work Item `source_refs` set. |
| Work Item self-reference | A Work Item includes its own canonical identity in `source_refs`. |
| Work Item/Task ownership | `work_item.tasks` and `task.work_item` do not agree. |
| Task dependency target | `task.depends_on` points to a missing Task. |
| Semantic provenance cycle | The normalized Work Item provenance graph contains a cycle. |

For provenance-cycle semantics, a Work Item source ref to a Task is normalized to the Task's owning Work Item.
Work Item `tasks`, Task `work_item`, derived Requirement reverse relations, and Task `depends_on` do not independently create provenance edges.

PRODUCT defines the invalid states and normalization meaning.
DRMCP defines the Task-owner resolution mechanism and cycle-analysis algorithm.

## Duplicate identity conditions

| identity class | invalid condition |
|---|---|
| Spec ref | More than one new or migrated spec path has the same canonical `spec:` ref. |
| Sequential record H1 public ID | More than one record H1 claims the same complete public ID. |
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
| Task-owner resolution mechanism | DRMCP. |
| Direct reverse lookup and transitive traversal | DRMCP. |
| Cycle-analysis algorithm | DRMCP. |
| UI and tool APIs | DRMCP. |
| Writer behavior, dry-run diff, confirmation, conflict handling, format preservation, and permission boundaries | DRMCP app-local specifications. |

## Resolve and validation boundary

Active resolve and validation contracts are defined by the sections above.
No `yaml:`, `internal-design:`, `coverage:`, `COV-*`, or fixture lookup is active.
Workflow orphan diagnostics, progress projection, traversal mechanisms, cycle-analysis algorithms, and execution-order checks belong to DRMCP app-local specifications.
Semantic provenance-cycle invalidity remains a PRODUCT rule.
Historical disposition evidence is recorded in T05.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.traceability.artifact_refs` | Canonical record kinds and reference forms. |
| `spec:product.design_records.traceability.metadata_schema` | Metadata and relation fields. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Spec lookup source and path-derived identity. |
| `spec:product.brewprint.compatibility` | Legacy issued-ID compatibility pointer. |
| PRODUCT-REQ-SPEC-006 | Generic workflow source-relation requirement. |
| PRODUCT-ADR-SPEC-007 | Source-ref validity and semantic provenance-cycle rules. |
| PRODUCT-ADR-SPEC-008 | Atomic migration and mismatch-blocking boundary. |
| PRODUCT-ADR-SPEC-019 | H1-only sequential identity authority and nonidentity validation boundary. |
