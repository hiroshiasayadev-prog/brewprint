# Reference: Resolve and validation

- **id**: `spec:product.concepts.traceability.resolve_and_validation`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `spec:product.concepts.traceability`

## What this is

Defines the canonical resolve and validation boundary for the traceability MVP: lookup sources, supported input forms, duplicate detection, declared relation integrity, and validation scope.

## Resolve

> **Drift guard:** PRODUCT owns canonical input semantics and tool behavior descriptions expressed in PRODUCT metadata. DRMCP may expose these as tool API — do not let DRMCP API vocabulary accumulate here.

Resolve is the tool behavior of resolving a canonical semantic ref or artifact ID-as-ref to an actual artifact / spec section. Resolve is not a relation between artifacts.

MVP examples:

```text
spec:product.concepts.traceability -> product/records/spec/concepts/traceability/index.md
spec:product.concepts.traceability.semantic_ref -> product/records/spec/concepts/traceability/semantic-ref.md
spec:product.concepts.project_artifact_model -> product/records/spec/concepts/project-artifact-model/index.md
V01-ADR-088 -> v01/records/adr/V01-ADR-088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
V01-INV-DOCS-003 -> v01/records/investigations/docs/V01-INV-DOCS-003-internal-design-semantic-trace-mvp-necessity.md
```

Per V01-ADR-087, canonical references in investigation `source_refs` and recorded `follow_up_results` are subject to resolve and validation. Per V01-ADR-092, the workflow ID-as-refs additionally permitted in investigation metadata are limited to complete requirement and work item public IDs; task public IDs are excluded. Declared relation integrity validation between workflow artifacts is included in the MVP. Per V01-ADR-088, the MVP does not require resolution of realization relations or external coverage mappings.

## Resolver input

> **Drift guard:** PRODUCT owns canonical input semantics and tool behavior descriptions expressed in PRODUCT metadata. DRMCP may expose these as tool API — do not let DRMCP API vocabulary accumulate here.

The MVP resolver handles at minimum:

- Active canonical `spec:` refs.
- Existing issued legacy public record IDs.
- Canonical app-aware artifact IDs for ADR, investigation, requirement, work item, and task records.
- Legacy `SPEC-*` public IDs only as compatibility inputs for indexed legacy specs.

New and migrated specs use path-derived `spec:` refs. New sequential record IDs follow `spec:product.concepts.namespace_model.artifact_id_grammar`.

Complete requirement, work item, and task public IDs are handled as direct resolver inputs. Workflow inter-relations are validated from metadata ID-as-refs; parent relations are not inferred from ID strings or physical paths.

The following are not required as MVP resolver inputs:

```text
internal-design:...
coverage:...
COV-...
yaml:...
```

`yaml:` is reserve-only. The others are targets whose decision as semantic trace endpoints / mapping mechanisms was deferred per V01-ADR-088.

## Resolver output

The concrete request/response fields and status vocabulary of the resolver are owned by the Design Records MCP spec.

This spec defines only the supported canonical reference boundary, lookup sources, validation responsibility, and the relation/diagnostic scope excluded from the MVP.

## Lookup sources

> **Drift guard:** PRODUCT owns canonical input semantics and tool behavior descriptions expressed in PRODUCT metadata. DRMCP may expose these as tool API — do not let DRMCP API vocabulary accumulate here.

The MVP resolver's lookup sources are:

| lookup source | role |
|---|---|
| spec front matter `semantic_refs` | registers document-level `spec:` refs |
| spec front matter `sections` | maps section-level `spec:` refs to heading text |
| record artifact public IDs indexed by Design Records MCP | Existing legacy public IDs and canonical app-aware artifact IDs; legacy `SPEC-*` only as compatibility input |

The resolver treats refs in the following investigation metadata fields as validation inputs:

- `source_refs`
- `follow_up_results`
- Canonical form of artifact references in `follow_up_candidates`, and `info` diagnostic for unresolved candidates

The workflow ID-as-refs referenceable in investigation metadata are limited to complete requirement and work item public IDs. Complete task public IDs are unsupported in those fields.

Workflow relation validation uses requirement `work_items`, work item `source_requirement` / `tasks`, and task `work_item` / `source_requirement` / `depends_on` as validation inputs.

Investigation metadata and workflow relation fields are the referring side and are not lookup sources that register reference targets.

In the MVP, refs are not inferred from natural language body text. Relation graphs are not built from internal-design front matter or coverage mapping artifacts.

## Section anchor lookup

> **Drift guard:** PRODUCT owns canonical input semantics and tool behavior descriptions expressed in PRODUCT metadata. DRMCP may expose these as tool API — do not let DRMCP API vocabulary accumulate here.

The only target for active resolution of section-level semantic refs is `spec:`, and heading text is mapped via the front matter `sections` mapping.

The resolver checks:

- That `sections` keys conform to `spec:` grammar
- That a Markdown heading matching the `sections` value exists
- That the resolution target is unique within the same file

## Duplicate detection

> **Drift guard:** PRODUCT owns invalid conditions and scope boundary. DRMCP owns diagnostic category names, JSON response shape, and tool response vocabulary — do not let those accumulate here.

MVP duplicate error candidates:

- The same `spec:` document-level ref appears in multiple documents' `semantic_refs`
- The same `spec:` section-level ref appears in multiple `sections` keys
- The same complete public record ID appears in multiple records

Duplicate detection for `COV-*`, `internal-design:`, and `coverage:` is outside MVP scope.

## Unresolved reference and declared relation integrity

> **Drift guard:** PRODUCT owns invalid conditions and scope boundary. DRMCP owns diagnostic category names, JSON response shape, and tool response vocabulary — do not let those accumulate here.

MVP unresolved error cases:

- Investigation metadata `source_refs` points to an unresolvable supported record ID-as-ref or active `spec:` ref
- Investigation metadata recorded `follow_up_results` points to an unresolvable supported record ID-as-ref or active `spec:` ref
- A workflow relation field points to an unresolvable complete requirement, work item, or task public ID

`follow_up_candidates` may point to not-yet-created artifacts. When artifact references are recorded, the canonical form is checked, but being unresolved itself is not an error. Canonical-but-unresolved candidates are returned as `info` diagnostics indicating a planned follow-up artifact that does not yet exist.

Declared workflow relation mismatches treated as integrity errors in the MVP:

- Mismatch between `requirement.work_items` and `work_item.source_requirement`
- Mismatch between `work_item.tasks` and `task.work_item`
- Mismatch between `task.source_requirement` and the parent work item's `source_requirement`

This is a consistency check of relations declared in metadata, not a search for orphan requirements / work items / tasks that are not referenced. Workflow orphan diagnostics are outside MVP scope.

## Reserved and deferred ref handling

`yaml:` is reserve-only; the MVP does not fix whether it may be written or what severity unresolved carries.

`internal-design:` / `coverage:` / `COV-*` were deferred from the MVP contract per V01-ADR-088. If they exist in current drafts/examples, they must not be resolved or validated as new MVP acceptance targets.

## Validation

The MVP handles:

- `spec:` ref grammar / uniqueness / section lookup. Both root document refs (`spec:trace`) and dot-notation refs (`spec:trace.semantic-ref`) are valid in the active `spec:` declaration grammar
- Resolution of complete public record ID-as-refs, including existing legacy IDs and canonical app-aware artifact IDs
- Canonicality and resolution of investigation `source_refs` / recorded `follow_up_results`. Additional workflow ID-as-refs are limited to complete requirement and work item public IDs
- Canonical form of investigation `follow_up_candidates` and `info` diagnostic for canonical-but-unresolved candidates. Additional workflow ID-as-refs are limited to complete requirement and work item public IDs
- Unsupported diagnostic for task public IDs in investigation metadata
- Target kind / resolution / declared bidirectional consistency of workflow relation fields
- Noncanonical diagnostic for physical path references: error when in `source_refs` / `follow_up_results`, `info` when in `follow_up_candidates`

The MVP does not handle:

- Coverage mapping YAML schema
- `maps_to` / `covers` / `validates` relations
- `spec:` → `internal-design:` endpoint constraints
- Internal-design relation declaration / reverse graph
- Workflow orphan diagnostics / progress projection / workflow traversal queries / task dependency cycle detection

## Validation boundary

Not introducing the `validates` relation does not mean resolve/schema validation is not performed. Validation is treated as tool behavior / metadata contract.

## MCP writer contract placeholder

The MVP does not define request/response schemas for writer tools. If semantic ref registration or investigation metadata updates are performed via tools in the future, dry-run diff, confirmation, conflict handling, and format preservation will be defined in a separate spec / ADR.

## Out of scope

- Concrete MCP request/response schema
- Writer tool args
- Internal-design / coverage / YAML endpoint resolution
- Realization relation validation
- Workflow orphan diagnostics / progress projection / traversal query / dependency graph projection
- Fixture/golden validation
- Brewprint DSL YAML schema validation
