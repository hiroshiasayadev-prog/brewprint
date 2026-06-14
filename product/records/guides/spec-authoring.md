# Spec Authoring Guide

Use this guide to write new or migrated spec files under app namespace records trees:
`product/records/spec/`, `drmcp/records/spec/`, `bpdsl/records/spec/`.

Source specs: `spec:product.concepts.spec_format` and its child specs.

---

## H1 format

Each spec file must have exactly one real ATX H1 outside YAML front matter and fenced code blocks.

```markdown
# <SpecKind>: <Title>
```

| file role | H1 example |
|---|---|
| area overview | `# Overview: Design Records MCP` |
| topic index | `# Index: DRMCP tool contracts` |
| concept spec | `# Concept: Spec format` |
| reference spec | `# Reference: Record metadata fields` |
| contract spec | `# Contract: validate_records tool` |

Rule: `index.md` as a filename does not force `Index` as the kind. A file named `index.md` may be an `Overview` when it contains substantive specification body.

---

## H1-adjacent metadata

Place this block immediately after H1 and before `## What this is`.

```markdown
- **id**: `spec:<app>.<path.derived.ref>`
- **status**: draft
- **date**: YYYY-MM-DD
- **parent**: `spec:<parent.ref>` | `root` | `-`
```

| marker | required | notes |
|---|---:|---|
| `id` | yes, all kinds | Path-derived canonical spec ref. See [Spec ID-as-ref](#spec-id-as-ref). |
| `status` | yes, all kinds | Spec lifecycle status. |
| `date` | yes, all kinds | Creation date or latest contract update date. |
| `parent` | yes, all kinds | `root` or `-` for top-level; canonical `spec:` ref for child specs. |
| `contract_class` | yes, `Contract` only | `interface` or `format`. Prohibited on all other kinds. |

---

## Spec ID-as-ref

The H1-adjacent `id` is the canonical spec ref. Do not add a separate `ref` marker.

Derivation rules:

| rule | contract |
|---|---|
| Prefix | `<app>/records/spec/` maps to `spec:<app>.` (lowercase app namespace). |
| Separator | Directory separators under `records/spec/` map to `.`. |
| Extension | `.md` is removed. |
| `index.md` | `index` is omitted. The containing directory name is the final segment. |
| Non-index file | File stem is kept as the final segment. |
| Word separator | Hyphens in path segments become underscores in spec IDs. |

Examples:

| path | canonical spec ref |
|---|---|
| `product/records/spec/concepts/spec-format/index.md` | `spec:product.concepts.spec_format` |
| `product/records/spec/concepts/traceability/index.md` | `spec:product.concepts.traceability` |
| `product/records/spec/concepts/traceability/semantic-ref.md` | `spec:product.concepts.traceability.semantic_ref` |
| `drmcp/records/spec/design-records-mcp/tools.md` | `spec:drmcp.design_records_mcp.tools` |

Parent grammar:

| parent form | allowed |
|---|---|
| `root` | yes |
| `-` | yes (compact root marker) |
| canonical `spec:` ref | yes |
| physical path | no |
| filename | no |
| H1 title string | no |

---

## Spec kind selection

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

`Contract` specs must declare a H1-adjacent `contract_class`. Do not introduce new H1 spec kinds such as `Interface` or `Format` — use `# Contract: ...` plus `contract_class`.

| contract_class | applies to | required H2 sections |
|---|---|---|
| `interface` | API, tool, integration, external-boundary contracts | `## Request`, `## Response`, `## Errors` |
| `format` | Document, table, metadata, validation-shape, file-format contracts | `## Current contract`, `## Rules`, `## Validation rules` |

---

## Required sections

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

---

## `## What this is`

State what this spec owns and what distinguishes it from siblings or parent specs.
Keep it to one or two sentences.

Anti-patterns:

| anti-pattern | problem |
|---|---|
| Opening with the H1 title repeated as a sentence | Redundant. Start with what it owns. |
| "This document describes..." | Throat-clearing. Remove it. |
| Listing everything the spec might cover | Replaces ownership statement with a scope dump. |

Good example:

```markdown
## What this is

This spec defines the visible document shape required for new or migrated spec records:
accepted spec kinds, H1 format, H1-adjacent metadata, and required section matrix.
```

Anti-pattern:

```markdown
## What this is

This document describes the spec format. It covers many topics related to how specs should be written.
```

---

## `## Topics` table

Required in `Index` specs. Allowed in `Overview` specs that declare authoritative child topics.

Required columns:

| column | required | meaning |
|---|---:|---|
| `title` | yes | Human-readable child topic title. |
| `kind` | yes | Child spec kind. Must be one of the accepted kinds. |
| `ref` | yes | Canonical `spec:` ref. Must match the path-derived canonical ref of the child spec. |
| `summary` | yes | Short navigation summary. |

Rule: Do not use `file` as a column. Tooling derives paths from `ref`; the table does not carry file paths.

Example:

```markdown
## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Spec document shape | Contract | `spec:product.concepts.spec_format.document_shape` | H1 format, metadata, kind set, and required section matrix. |
| Spec ID-as-ref | Concept | `spec:product.concepts.spec_format.spec_id_as_ref` | Path-derived canonical refs and parent grammar. |
| Topics table | Contract | `spec:product.concepts.spec_format.topics_table` | `## Topics` table columns and parent declaration rules. |
```

---

## Front matter policy

New and migrated specs must not use YAML front matter as a metadata source of truth.

| item | new / migrated spec | existing unmigrated spec |
|---|---|---|
| Any YAML front matter block | Prohibited — error | Warning until migration |
| `depends_on` in front matter | Prohibited — error | Warning until migration |
| `semantic_refs` / `sections` in front matter | Prohibited — error | Warning until migration |
| `design_record.kind` in front matter | Prohibited — error | Warning until migration |

Use H1-adjacent metadata markers instead. If the current DRMCP requires front matter for spec discovery, treat that as a DRMCP compatibility gap and resolve it in follow-up indexing work.

---

## Source records

| ref | role |
|---|---|
| `spec:product.concepts.spec_format` | Parent index for the spec-format contract. |
| `spec:product.concepts.spec_format.document_shape` | H1 format, metadata markers, kind set, required section matrix. |
| `spec:product.concepts.spec_format.spec_id_as_ref` | ID derivation rules and parent grammar. |
| `spec:product.concepts.spec_format.topics_table` | `## Topics` table contract, columns, and validation rules. |
| `spec:product.concepts.spec_format.validation_policy` | Severity rules during migration phase. |
| `spec:product.concepts.spec_format.follow_up_boundary` | Ownership boundary for follow-up work. |
| PRODUCT-WORK-SPEC-003 | Source work item for this guide. |
