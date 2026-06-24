# Reference: Investigation authoring

- **id**: `spec:product.design_records.authoring_standards.investigation_authoring`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `spec:product.design_records.authoring_standards`

## What this is

Authoring rules for Investigation artifacts.

This guide defines investigation IDs, paths, file shape, metadata, lifecycle, references, and author-facing inputs.

## Non-goals

- Current DRMCP operating status.
- Concrete DRMCP request, response, or diagnostic schemas.
- ADR, requirement, work-item, or task authoring rules.
- Adoption status of follow-up candidates — investigation records findings only, not what was accepted.
- Cross-artifact ownership beyond investigation-specific writing rules.

## Rules

### ID grammar

| Rule | Level |
|---|---|
| New investigation IDs use `<APP>-INV-<DOMAIN>-<NNN>`. | MUST |
| `<NNN>` is a three-digit, zero-padded sequence scoped by app namespace, artifact kind, and domain namespace. | MUST |
| Authors may use `<APP>-INV-<DOMAIN>-new` when requesting automatic sequence allocation. | MAY |
| Existing investigation IDs remain unchanged unless a separate migration decision requires changes. | MUST |

The canonical grammar source is `spec:product.design_records.namespace_model.artifact_id_grammar`.

### File path layout

| Rule | Level |
|---|---|
| New investigations use `<app>/records/investigations/<domain>/<APP>-INV-<DOMAIN>-<NNN>-<slug>.md`. | MUST |
| The file name prefix matches the public ID through the sequence segment. | MUST |
| Physical paths are repository locations, not canonical references. | MUST |

The discovery pattern is defined by `spec:product.design_records.repository_layout.record_discovery_paths`.

### File shape

| Rule | Level |
|---|---|
| Use exactly one ATX H1 outside fenced code blocks. | MUST |
| Use `# <PUBLIC-ID>: <Title>` for H1. | MUST |
| Place the bullet metadata block immediately after H1. | MUST |
| Start body content at the first H2 after metadata. | MUST |

Body section schema:

| section | heading presence | substantive content | content |
|---|---|---|---|
| `## Investigation scope` | always | required when `concluded` | What this investigation covers. |
| `## Out of scope` | always | TBD allowed | What is explicitly excluded from this investigation. |
| `## Background` | always | TBD allowed | Context and motivation for the investigation. |
| `## What was investigated` | always | TBD allowed | The specific items, questions, or areas examined. |
| `## Findings` | always | required when `concluded` | Factual results organized by investigation item. |
| `## Cross-cutting observations` | always | TBD allowed | Patterns or facts observed across multiple investigation items. |
| `## Follow-up judgment candidates` | always | TBD allowed | Topics or decisions to hand off for follow-up judgment. |
| `## Recommendation` | always | TBD allowed | The investigator's view on next steps. Not a design decision. |
| `## Follow-up artifact candidates` | always | TBD allowed | Proposed follow-up artifacts or artifact updates. |
| `## Open questions` | always | TBD allowed | Unresolved points that remain after concluding the investigation. |

All canonical investigation headings use English.

All canonical sections must be present when creating the record, to preserve a stable document shape. Sections not yet written may contain `TBD`.

### Metadata schema

Investigation metadata uses H1-adjacent bullet fields. Unlike requirement, work-item, and task records, investigation records do not carry an explicit `id` field in bullet metadata; the public ID appears in H1 and the file name only.

Required metadata:

| field | create input | partial update | persisted investigation | meaning |
|---|---|---|---|---|
| `status` | required | optional | required | Current investigation lifecycle state. |
| `date` | required | optional | required | Date the investigation was opened or scope last changed. |
| `trigger` | required | optional | required | The artifact, event, or question that prompted this investigation. |
| `scope` | required | optional | required | Short statement of what this investigation covers. Mirrors `## Investigation scope`. |
| `non_scope` | required | optional | required | Short statement of what is explicitly excluded. Mirrors `## Out of scope`. |
| `source_refs` | required; empty list allowed | optional | required | Canonical artifact IDs or active semantic refs that motivated this investigation. |
| `follow_up_candidates` | required; empty list allowed | optional | required | Proposed follow-up artifact IDs, semantic refs, or human-readable candidates. |

Optional metadata:

| field | create input | partial update | persisted investigation | meaning |
|---|---|---|---|---|
| `supersedes` | optional | optional; correction only | optional | Investigation IDs this record supersedes. |
| `related_requirements` | optional | optional | optional | Related requirement IDs. |
| `related_work_items` | optional | optional | optional | Related work item IDs. |
| `related_adrs` | optional | optional | optional | Related ADR IDs. |
| `related_specs` | optional | optional | optional | Related spec semantic refs. |
| `related_internal_design` | optional | optional | optional | Parser-recognized optional auxiliary identifier. Canonical resolution and validation semantics are not defined by this guide. |
| `related_coverage` | optional | optional | optional | Parser-recognized optional auxiliary identifier. Canonical resolution and validation semantics are not defined by this guide. |
| `follow_up_results` | optional | optional | optional | Canonical refs for artifacts actually created or updated as follow-up results. |

Rules:

- `date` uses strict `YYYY-MM-DD` format.
- Update `date` when investigation scope changes meaningfully. Do not update `date` for editorial corrections or appending findings.
- `scope` and `non_scope` are short inline summaries (one or two sentences); they are not duplications of the body sections, which carry the full narrative.
- `source_refs` normalizes empty values to an empty list. `source_refs` does not accept empty list items.
- `follow_up_candidates` must be present; use an empty list when no candidate exists.
- `follow_up_candidates` does not accept empty list items.
- Change `supersedes` only to correct investigation lineage metadata that was omitted or recorded incorrectly.
- Optional metadata fields are omitted entirely when not applicable; absent optional fields are not treated as validation errors.
- `follow_up_results` is recorded when follow-up artifacts are known; it may be supplied at create for migrated or already-concluded investigations.

The parsing grammar is defined by `spec:drmcp.design_records_mcp.schema.metadata_grammar`.

### Status lifecycle

| status | meaning |
|---|---|
| `investigating` | Investigation is in progress. |
| `concluded` | Investigation results are complete and ready for follow-up judgment. |
| `superseded` | This investigation has been replaced by a later investigation or artifact. |

Rules:

- `concluded` does not mean follow-up candidates have been accepted, adopted, or acted on. Follow-up is tracked in the artifacts that result from it.
- When `status` is `superseded`, the superseding investigation should record this investigation's ID in its `supersedes` field.
- Do not use `proposed` as an investigation status; it conflicts with ADR status terminology.

Investigation conclusion-readiness rule:

- An investigation must not be marked `concluded` unless `## Investigation scope` contains substantive content.
- An investigation must not be marked `concluded` unless `## Findings` contains substantive content.
- `TBD` is placeholder content, not substantive content. `TBD` does not satisfy either condition.

### Kind-specific authoring rules

Cross-artifact selection follows `spec:product.design_records.authoring_standards.artifact_boundary`.

Decision boundary:

- An investigation captures findings, evidence, hypotheses, options, and uncertainties. It does not capture decisions.
- `## Recommendation` may express the investigator's view on next steps. Frame it as "appears preferable" or "is likely appropriate", not as a decided direction.
- Acting on a recommendation requires a subsequent ADR, spec, investigation, requirement, or work item, depending on artifact responsibility — not an investigation update.
- Do not write current specification text in an investigation body; use a spec.
- Do not write design decisions in an investigation body; use an ADR.

Follow-up tracking:

- `follow_up_candidates` at create lists proposed follow-up artifact IDs or artifact types to create or update next.
- Populate `follow_up_results` with canonical refs for artifacts actually created or updated as follow-up results. This may be supplied at create for concluded investigations, or updated when follow-up artifacts are known.
- `follow_up_candidates` that do not appear in `follow_up_results` were not acted on. This is valid.

Split rule:

- When investigation scope expands to cover a distinct domain or question, consider creating a separate investigation.
- Minor additions to the original scope may remain in the same record.

Supersession lineage:

- A superseding investigation records superseded investigation IDs in `supersedes`.
- The superseded investigation sets `status` to `superseded`.
- `## Background` may explain the reason and history but does not establish canonical supersession lineage.

### Canonical reference policy

| Rule | Level |
|---|---|
| Reference design records by public ID. | MUST |
| Reference specs by active `spec:` semantic ref. | MUST |
| Use physical paths only as supplementary location notes. | MUST |
| Do not use `TASK-*` IDs in `source_refs`, `follow_up_candidates`, or `follow_up_results`. | MUST |

`TASK-*` IDs are not used in investigation metadata canonical references. Task relations are recorded in task artifacts, not in investigation metadata.

## Authoring interface requirements

### Create

The author supplies:

- app namespace;
- domain namespace;
- title;
- exact ID or `<APP>-INV-<DOMAIN>-new` placeholder;
- all required metadata fields;
- investigation body sections.

The author does not supply:

- `id` as a metadata field — investigation records do not carry `id` in bullet metadata;
- the resolved sequence when using `new`;
- a generated H1;
- a generated file path.

The body begins with `## Investigation scope`. The body excludes H1 and bullet metadata. Include all canonical sections; use `TBD` for sections not yet written.

### Update

A partial update supplies only changed metadata fields or body sections.

Rules:

- Omitted metadata fields remain unchanged.
- Update `date` only when the investigation scope changes meaningfully.
- Named section updates use the canonical English H2 headings.
- Change `supersedes` only to correct investigation lineage metadata that was omitted or recorded incorrectly.

Concrete tool contracts belong to DRMCP specs.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.authoring_standards` | Parent Index. |
| `spec:product.design_records.authoring_standards.writing_standard` | Design record prose rules. |
| `spec:product.design_records.authoring_standards.artifact_boundary` | Authoring-time artifact selection boundary. |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | Canonical artifact ownership. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Investigation ID grammar. |
| `spec:product.design_records.repository_layout.record_discovery_paths` | Investigation discovery path rules. |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Investigation metadata parsing grammar. |
| `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | Concrete authoring transaction contract. |
| PRODUCT-REQ-SPEC-002 | Source requirement. |
| PRODUCT-WORK-SPEC-011 | Source work item. |
