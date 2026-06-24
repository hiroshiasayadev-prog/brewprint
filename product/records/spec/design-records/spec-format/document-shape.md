# Contract: Spec document shape

- **id**: `spec:product.design_records.spec_format.document_shape`
- **status**: accepted
- **date**: 2026-06-11
- **parent**: `spec:product.design_records.spec_format`
- **contract_class**: `format`

## What this is

This spec defines the visible document shape required for new or migrated spec records: accepted spec kinds, H1 format, H1-adjacent metadata, and required section matrix.

It preserves the document-shape rules from the former all-in-one `spec-format/index.md` while keeping navigation in the parent Index.

## Current contract

A spec markdown file under an app namespace records tree conforms to this document-shape contract when it uses an accepted spec kind, has exactly one real ATX H1 outside YAML front matter and fenced code blocks, carries the required H1-adjacent metadata block, and satisfies the section expectations for its kind and contract class.

`Contract` remains the only spec kind for contracts. `interface` and `format` are contract classes declared by H1-adjacent metadata, not additional H1 spec kinds.

## Rules

### Spec kind set

#### Initial accepted kinds

| spec kind | status | intent |
|---|---|---|
| `Overview` | accepted | Entry point for a spec area. It may summarize current contract, scope, and topic map. |
| `Index` | accepted | Navigation-first topic table. It should avoid substantive specification body beyond topic organization. |
| `Concept` | accepted | Concept model, vocabulary, semantic boundary, and cross-app rules. |
| `Reference` | accepted | Field, enum, grammar, fixed rule, vocabulary, or catalog reference. |
| `Contract` | accepted | API, tool, request, response, validation, or error behavior contract. |

#### Deferred kinds

| spec kind | status | revisit condition |
|---|---|---|
| `Guide` | deferred | Revisit when current guides become spec records or usage/authoring procedure specs need first-class topic treatment. |
| `Process` | deferred | Revisit when lifecycle / workflow / transition models need standalone spec representation. |
| `Architecture` | deferred | Revisit when standalone component / runtime / storage / dependency architecture specs appear. |
| `Glossary` | deferred | Revisit when term-only specs appear. |

Deferred kinds must not be used for new migrated specs until this spec or a successor explicitly accepts them.

### Contract class

`Contract` specs must declare a H1-adjacent `contract_class` marker.

| contract_class | applies to | required sections |
|---|---|---|
| `interface` | API, tool, integration, and external-boundary contracts. | H2 sections named `Request`, `Response`, and `Errors`. |
| `format` | Document, table, metadata, validation-shape, and file-format contracts. | H2 sections named `Current contract`, `Rules`, and `Validation rules`. `Errors` is optional unless the format contract defines validation diagnostics. |

Do not introduce new H1 spec kinds such as `Interface` or `Format`; use `# Contract: ...` plus `contract_class`.

### H1 format

Each spec markdown file must have exactly one ATX H1 outside YAML front matter and fenced code blocks.

The H1 format is:

```markdown
# <SpecKind>: <Title>
```

Examples:

| file role | H1 example |
|---|---|
| area overview | `# Overview: Design Records MCP` |
| topic index | `# Index: DRMCP tool contracts` |
| concept spec | `# Concept: Spec format` |
| reference spec | `# Reference: Record metadata fields` |
| contract spec | `# Contract: validate_records tool` |

`index.md` as a file name does not force `Index` as the spec kind. A file named `index.md` may be an `Overview` when it contains substantive specification body.

### H1-adjacent metadata

H1-adjacent metadata is the visible metadata block immediately after H1 and before `## What this is`.

| marker | required | applies to | meaning |
|---|---:|---|---|
| `- **id**:` | yes | all spec files | Canonical document-level ID-as-ref. |
| `- **status**:` | yes | all spec files | Spec lifecycle status. |
| `- **date**:` | yes | all spec files | Creation or latest contract update date, depending on lifecycle policy. |
| `- **parent**:` | yes | all spec files | `root`, `-`, or canonical parent `spec:` ref. |
| `- **contract_class**:` | yes, for `Contract` only | Contract specs | Contract class. Allowed values are `interface` and `format`. |

H1-adjacent metadata is intentionally visible because hidden topic metadata is prone to stale drift.

For non-root child specs, `parent` must match exactly one authoritative parent declaration from an `Index` or `Overview` `## Topics` row.

### Required section matrix

| section / marker | Overview | Index | Concept | Reference | Contract `interface` | Contract `format` |
|---|---|---|---|---|---|---|
| H1-adjacent `- **id**:` | required | required | required | required | required | required |
| H1-adjacent `- **status**:` | required | required | required | required | required | required |
| H1-adjacent `- **date**:` | required | required | required | required | required | required |
| H1-adjacent `- **parent**:` | required | required | required | required | required | required |
| H1-adjacent `- **contract_class**:` | prohibited | prohibited | prohibited | prohibited | required | required |
| `## What this is` | required | required | required | required | required | required |
| `## Current contract` | required | prohibited | optional | optional | optional | required |
| `## Non-goals` | recommended | prohibited | recommended | optional | recommended | recommended |
| `## Topic map` | recommended | prohibited | optional | optional | optional | optional |
| `## Topics` table | optional | required | prohibited | prohibited | prohibited | prohibited |
| `## Concept model` | optional | prohibited | required | prohibited | optional | optional |
| `## Rules` | optional | prohibited | recommended | optional | optional | required |
| `## Boundary` | optional | prohibited | recommended | optional | recommended | recommended |
| body H2 containing at least one Markdown table | prohibited | prohibited | optional | required | optional | optional |
| H2 `Request` | prohibited | prohibited | prohibited | prohibited | required | prohibited |
| H2 `Response` | prohibited | prohibited | prohibited | prohibited | required | prohibited |
| `## Errors` | optional | prohibited | prohibited | optional | required | optional |
| `## Validation rules` | optional | prohibited | optional | optional | optional | required |
| `## Related specs` | recommended | recommended | recommended | recommended | recommended | recommended |

## Validation rules

Existing specs that do not yet contain `## What this is` should produce migration warnings, not hard errors, until the migration work explicitly updates them.

Localized aliases such as `## 目的` are not canonical section names for this format. They may remain during migration, but the canonical `## What this is` section is required for new or migrated specs.

`## Topic map` is a human navigation hint. It has no parser-validated internal structure in this contract; authoring examples belong to PRODUCT-WORK-SPEC-003.

## Errors

| condition | severity |
|---|---|
| Missing real ATX H1 | Error for new or migrated specs. |
| Multiple real ATX H1 headings | Error for new or migrated specs. |
| H1 kind outside the accepted kind set | Error. |
| Missing H1-adjacent `id`, `status`, `date`, or `parent` marker | Error for new or migrated specs; warning during inventory. |
| Missing or invalid H1-adjacent `contract_class` on a `Contract` spec | Error for new or migrated specs; warning during inventory. |
| Missing required section by kind | Warning or error based on migration phase. |
| Format contract uses artificial H2 `Request` or H2 `Response` sections as required shape | Error for new or migrated specs. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.spec_format` | Parent Index for this contract. |
| `spec:product.design_records.spec_format.topics_table` | Defines the Index and Overview topic table contract. |
| `spec:product.design_records.spec_format.validation_policy` | Defines severity and parser-aware validation policy. |
