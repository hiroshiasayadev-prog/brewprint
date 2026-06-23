# Reference: Spec authoring

- **id**: `spec:product.concepts.authoring_standards.spec_authoring`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `spec:product.concepts.authoring_standards`

## What this is

Authoring rules for spec records.

This guide defines spec identity, file paths, file shape, metadata, kind selection, and author-facing inputs. The normative spec format contract is owned by `spec:product.concepts.spec_format` and its children; this guide provides the author-facing view.

## Non-goals

- Current DRMCP operating status.
- Concrete DRMCP request, response, or diagnostic schemas.
- ADR, requirement, work-item, task, or investigation authoring rules.
- Normative spec format contract — owned by `spec:product.concepts.spec_format`.
- Cross-artifact ownership beyond spec-specific writing rules.

## Rules

### ID grammar

Spec IDs are path-derived canonical refs, not numeric sequences. The public `id` is derived from the file path under `<app>/records/spec/`.

Derivation rules:

| rule | contract |
|---|---|
| Prefix | `<app>/records/spec/` maps to `spec:<app>.` (lowercase app namespace). |
| Separator | Directory separators under `records/spec/` map to `.`. |
| Extension | `.md` is removed. |
| `index.md` | `index` is omitted; the containing directory name is the final segment. |
| Non-index file | File stem is kept as the final segment. |
| Word separator | Hyphens in path segments become underscores in spec IDs. |

| path | canonical spec ref |
|---|---|
| `product/records/spec/concepts/spec-format/index.md` | `spec:product.concepts.spec_format` |
| `product/records/spec/concepts/traceability/semantic-ref.md` | `spec:product.concepts.traceability.semantic_ref` |
| `drmcp/records/spec/design-records-mcp/tools.md` | `spec:drmcp.design_records_mcp.tools` |

The canonical grammar source is `spec:product.concepts.spec_format.spec_id_as_ref`.

### File path layout

| Rule | Level |
|---|---|
| New specs use `<app>/records/spec/<path-to-topic>/<filename>.md`. | MUST |
| Index specs for a topic area use `<path-to-topic>/index.md`. | MUST |
| Non-index specs use a descriptive `<stem>.md`; hyphens in the stem become underscores in the canonical ref. | MUST |
| Physical paths are repository locations, not canonical references. | MUST |

### File shape

| Rule | Level |
|---|---|
| Use exactly one ATX H1 outside YAML front matter and fenced code blocks. | MUST |
| Use `# <SpecKind>: <Title>` for H1. | MUST |
| Place the H1-adjacent metadata block immediately after H1 and before `## What this is`. | MUST |
| `index.md` as a filename does not force `Index` as the kind. A file named `index.md` may be an `Overview` when it contains substantive specification body. | MUST |

Accepted spec kinds:

| kind | use when |
|---|---|
| `Overview` | Entry point for a spec area. Summarizes current contract, scope, and topic map. |
| `Index` | Navigation-first topic table. Avoid substantive specification body beyond topic organization. |
| `Concept` | Concept model, vocabulary, semantic boundary, or cross-app rules. |
| `Reference` | Field, enum, grammar, fixed rule, vocabulary, or catalog. |
| `Contract` | API, tool, request/response, validation, or error behavior. |

Deferred kinds — do not use for new or migrated specs:

| kind | revisit condition |
|---|---|
| `Guide` | When current guides become spec records or authoring procedure needs first-class topic treatment. |
| `Process` | When lifecycle / workflow / transition models need standalone spec representation. |
| `Architecture` | When standalone component / runtime / storage architecture specs appear. |
| `Glossary` | When term-only specs appear. |

Required sections by kind:

| section | Overview | Index | Concept | Reference | Contract `interface` | Contract `format` |
|---|---|---|---|---|---|---|
| `## What this is` | required | required | required | required | required | required |
| `## Current contract` | required | prohibited | optional | optional | optional | required |
| `## Non-goals` | recommended | prohibited | recommended | optional | recommended | recommended |
| `## Topic map` | recommended | prohibited | optional | optional | optional | optional |
| `## Topics` table | optional | required | prohibited | prohibited | prohibited | prohibited |
| `## Concept model` | optional | prohibited | required | prohibited | optional | optional |
| `## Rules` | optional | prohibited | recommended | optional | optional | required |
| `## Boundary` | optional | prohibited | recommended | optional | recommended | recommended |
| body H2 containing at least one Markdown table | prohibited | prohibited | optional | required | optional | optional |
| `## Request` | prohibited | prohibited | prohibited | prohibited | required | prohibited |
| `## Response` | prohibited | prohibited | prohibited | prohibited | required | prohibited |
| `## Errors` | optional | prohibited | prohibited | optional | required | optional |
| `## Validation rules` | optional | prohibited | optional | optional | optional | required |
| `## Related specs` | recommended | recommended | recommended | recommended | recommended | recommended |

`Contract` specs must declare a H1-adjacent `contract_class`. Do not introduce new H1 spec kinds such as `Interface` or `Format` — use `# Contract: ...` plus `contract_class`.

| contract_class | applies to | required H2 sections |
|---|---|---|
| `interface` | API, tool, integration, external-boundary contracts | `## Request`, `## Response`, `## Errors` |
| `format` | Document, table, metadata, validation-shape, file-format contracts | `## Current contract`, `## Rules`, `## Validation rules` |

### Metadata schema

Spec metadata uses H1-adjacent bullet fields.

| field | create input | partial update | persisted spec | meaning |
|---|---|---|---|---|
| `id` | path-derived; not supplied as a metadata field | not updated | required | Path-derived canonical spec ref. Must match the derivation of the file path. |
| `status` | required | optional | required | Current spec lifecycle state. |
| `date` | required | optional | required | Creation date or date of the latest substantive contract change. |
| `parent` | required | optional | required | Parent spec canonical ref, or `root` / `-` for top-level specs. |
| `contract_class` | required for `Contract` kind; prohibited on all other kinds | optional | conditional | `interface` or `format`. Omit for non-Contract kinds. |

Rules:

- `date` uses strict `YYYY-MM-DD` format.
- Update `date` when the spec contract changes meaningfully. Do not update `date` for editorial corrections.
- `parent` uses a canonical `spec:` ref or the literal `root` / `-`. Physical paths and H1 title strings are not valid parent forms.
- `id` is generated from the file path by tooling. Do not supply it as a metadata field in a create request.
- `id` in the persisted spec must match the path-derived ref.

The parsing grammar is defined by `spec:drmcp.design_records_mcp.schema.metadata_grammar`.

### Status lifecycle

The spec status lifecycle is defined by `spec:product.concepts.spec_format`. `draft` is the current documented status for specs under active development.

Rules:

- Update `date` when the spec's normative content changes. Do not treat it as an automatic modification timestamp.
- A spec that is superseded by a redesigned contract should be retired or redirected; the replacement carries a new path-derived identity.

### Kind-specific authoring rules

Cross-artifact selection follows `spec:product.concepts.authoring_standards.artifact_boundary`.

`## What this is` rules:

- State what this spec owns and what distinguishes it from siblings or parent specs.
- Keep it to one or two sentences.
- Do not open with the H1 title repeated as a sentence.
- Do not open with "This document describes..."
- Do not list everything the spec might cover in place of an ownership statement.

`## Topics` table rules:

- Required in `Index` specs. Allowed in `Overview` specs that declare authoritative child topics.
- Required columns: `title`, `kind`, `ref`, `summary`.
- The `ref` value must match the path-derived canonical ref of the child spec.
- Do not use a `file` column. Tooling derives paths from `ref`.

Front matter policy:

- New and migrated specs must not use YAML front matter as a metadata source of truth.
- Any YAML front matter block in a new or migrated spec is an error.
- Use H1-adjacent metadata markers instead.

### Canonical reference policy

| Rule | Level |
|---|---|
| Reference specs by active `spec:` semantic ref. | MUST |
| Reference design records by public ID. | MUST |
| Use physical paths only as supplementary location notes. | MUST |
| Use canonical `spec:` refs in `parent` and in `## Topics` table `ref` columns. | MUST |

## Authoring interface requirements

### Create

The author supplies:

- app namespace;
- directory path under `records/spec/` (determines the canonical ref);
- spec kind;
- title;
- `parent` canonical ref or `root` / `-`;
- `contract_class` if and only if the kind is `Contract`;
- required body sections for the chosen kind.

The author does not supply:

- `id` as a metadata field — derived from the file path;
- a generated H1;
- a generated file path.

The body begins with `## What this is`. The body excludes H1 and bullet metadata.

### Update

A partial update supplies only changed metadata fields or body sections.

Rules:

- Omitted metadata fields remain unchanged.
- Update `date` only when the spec's normative content changes meaningfully.
- Named section updates use canonical English H2 headings.
- `id` is not a metadata update target.
- `parent` may be updated when a spec is reorganized under a different parent; update the corresponding Topics entry in the old and new parents.

Concrete tool contracts belong to DRMCP specs.

## Related specs

| ref | relation |
|---|---|
| `spec:product.concepts.authoring_standards` | Parent Index. |
| `spec:product.concepts.authoring_standards.writing_standard` | Design record prose rules. |
| `spec:product.concepts.authoring_standards.artifact_boundary` | Authoring-time artifact selection boundary. |
| `spec:product.concepts.project_artifact_model.artifact_responsibility_matrix` | Canonical artifact ownership. |
| `spec:product.concepts.spec_format` | Normative spec format contract. |
| `spec:product.concepts.spec_format.spec_id_as_ref` | Path-derived canonical ref rules and parent grammar. |
| `spec:product.concepts.spec_format.document_shape` | H1 format, metadata markers, kind set, section matrix. |
| `spec:product.concepts.spec_format.topics_table` | `## Topics` table contract and validation rules. |
| `spec:product.concepts.spec_format.validation_policy` | Severity rules during migration phase. |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Spec metadata parsing grammar. |
| `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | Concrete authoring transaction contract. |
| PRODUCT-REQ-SPEC-002 | Source requirement. |
