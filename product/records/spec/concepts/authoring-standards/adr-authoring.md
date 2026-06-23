# Reference: ADR authoring

- **id**: `spec:product.concepts.authoring_standards.adr_authoring`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `spec:product.concepts.authoring_standards`

## What this is

Authoring rules for Architecture Decision Record artifacts.

This guide defines ADR IDs, paths, file shape, metadata, lifecycle, references, and author-facing inputs.

## Non-goals

- Current DRMCP operating status.
- Concrete DRMCP request, response, or diagnostic schemas.
- Migration plans for existing ADRs.
- Spec authoring rules.
- Cross-artifact ownership beyond ADR-specific writing rules.

## Rules

### ID grammar

| Rule | Level |
|---|---|
| New ADR IDs use `<APP>-ADR-<DOMAIN>-<NNN>`. | MUST |
| `<NNN>` is a three-digit, zero-padded sequence scoped by app namespace, artifact kind, and domain namespace. | MUST |
| Authors may use `<APP>-ADR-<DOMAIN>-new` when requesting automatic sequence allocation. | MAY |
| Existing ADR IDs remain unchanged unless a separate migration decision requires changes. | MUST |

The canonical grammar source is `spec:product.concepts.namespace_model.v2_grammar`.

### File path layout

| Rule | Level |
|---|---|
| New ADRs use `<app>/records/adr/<domain>/<APP>-ADR-<DOMAIN>-<NNN>-<slug>.md`. | MUST |
| The file name prefix matches the public ID through the sequence segment. | MUST |
| Physical paths are repository locations, not canonical references. | MUST |

The discovery pattern is defined by `spec:product.concepts.repository_layout.record_discovery_paths`.

### File shape

| Rule | Level |
|---|---|
| Use exactly one ATX H1 outside fenced code blocks. | MUST |
| Use `# <PUBLIC-ID>: <Title>` for H1. | MUST |
| Place the bullet metadata block immediately after H1. | MUST |
| Start body content at the first H2 after metadata. | MUST |

Body section schema:

| section | required | content |
|---|---|---|
| `## Context` | yes | The problem or constraint requiring a decision. |
| `## Decision` | yes | The adopted decision. |
| `## Rationale` | yes | Why the decision was adopted. |
| `## Rejected alternatives` | no | Decision-relevant alternatives and concise rejection reasons. |
| `## Consequences` | yes | Expected effects, trade-offs, and affected artifacts. |
| `## Evidence` | yes | Supporting records, commits, observations, or implementation evidence. |

All canonical ADR headings use English.

### Metadata schema

ADR metadata uses H1-adjacent bullet fields.

| field | create input | partial update | persisted ADR | meaning |
|---|---|---|---|---|
| `status` | required | optional | required | Current ADR lifecycle state. |
| `date` | required | optional | required | Date when the documented decision became valid. |
| `depends_on` | required; empty allowed | optional | required | ADR IDs required by this decision. |
| `supersedes` | required; empty allowed | optional | required | Older ADR IDs replaced by this decision. |
| `migrated_to_spec` | required; empty allowed | optional | required | Date when the decision was reflected in a spec. |

Rules:

- `date` uses strict `YYYY-MM-DD` format.
- Update `date` only when the decision meaning changes.
- Do not update `date` for editorial corrections.
- `depends_on` and `supersedes` normalize empty values to empty lists.
- `migrated_to_spec` normalizes an empty value to `null`.
- Non-empty `migrated_to_spec` uses strict `YYYY-MM-DD` format.
- ADR bullet metadata does not contain `id`.
- The public ID comes from H1 and the file name.

The parsing grammar is defined by `spec:drmcp.design_records_mcp.schema.metadata_grammar`.

### Status lifecycle

| status | meaning | allowed transition |
|---|---|---|
| `proposed` | The decision remains under discussion. | `accepted` |
| `accepted` | The decision is valid. | `superseded` |
| `superseded` | A newer ADR replaced the decision. | terminal |

Rules:

- A replacing ADR lists the replaced ADR in `supersedes`.
- Do not rewrite an accepted ADR to overturn its decision.
- Create a new ADR and supersede the accepted ADR.
- ADR lifecycle state is carried by `status` and the supersession chain.

### Kind-specific authoring rules

Cross-artifact selection follows `spec:product.concepts.authoring_standards.artifact_boundary`.

ADR-specific rules:

- Record adopted design decisions and their rationale.
- Record concise rejection reasons for decision-relevant alternatives.
- Put detailed comparisons and exploration logs in an investigation.
- Put currently valid contracts in a spec.
- Put implementation steps and progress in work items or tasks.
- Treat observed implementation, fixture, and example shapes as evidence.
- Do not treat observed shapes as normative specifications.
- Separate each observed fact from the resulting design decision.

Canonical ownership remains defined by `spec:product.concepts.project_artifact_model.artifact_responsibility_matrix`.

### Canonical reference policy

| Rule | Level |
|---|---|
| Reference design records by public ID. | MUST |
| Reference specs by active `spec:` semantic ref. | MUST |
| Use physical paths only as supplementary location notes. | MUST |
| Use ADR public IDs in `depends_on` and `supersedes`. | MUST |

## Authoring interface requirements

### Create

The author supplies:

- app namespace;
- domain namespace;
- title;
- exact ID or `new` placeholder;
- all create-required metadata fields;
- ADR body sections.

The author does not supply:

- the resolved sequence when using `new`;
- a generated H1;
- a generated file path.

The body begins with `## Context`. The body excludes H1 and bullet metadata.

### Update

A partial update supplies only changed metadata fields or body sections.

Rules:

- Omitted metadata fields remain unchanged.
- Update `date` only when the decision meaning changes.
- Named section updates use the canonical English H2 headings.
- ADR ID changes are not metadata updates.

Concrete tool contracts belong to DRMCP specs.

## Related specs

| ref | relation |
|---|---|
| `spec:product.concepts.authoring_standards` | Parent Index. |
| `spec:product.concepts.authoring_standards.writing_standard` | Design record prose rules. |
| `spec:product.concepts.authoring_standards.artifact_boundary` | Authoring-time artifact selection boundary. |
| `spec:product.concepts.project_artifact_model.artifact_responsibility_matrix` | Canonical artifact ownership. |
| `spec:product.concepts.namespace_model.v2_grammar` | ADR ID grammar. |
| `spec:product.concepts.repository_layout.record_discovery_paths` | ADR discovery path rules. |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | ADR metadata parsing grammar. |
| `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | Concrete authoring transaction contract. |
| PRODUCT-REQ-SPEC-002 | Source requirement. |
| PRODUCT-WORK-SPEC-011 | Source work item. |
