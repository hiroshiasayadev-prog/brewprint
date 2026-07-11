# Reference: Artifact source

- **id**: `spec:drmcp.design_records_mcp.artifacts.base.definitions.source`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.base`

## What this is

Defines the shared source-document rules used by artifact-specific source Specifications.

This Specification defines the common H1 shape, metadata placement, and H2 declaration vocabulary.
It does not define artifact-specific headings or heading body formats.

## Current contract

### H1 format

The H1 is the first line of the source document.
Artifact sources use this H1 shape:

```text
# <H1_PREFIX>: <Title>
```

`<H1_PREFIX>` is the complete text between the leading `# ` marker and the first `: ` delimiter.

Each artifact-specific source Specification declares the H1 prefix by using exactly one of these types:

| type | declaration |
|---|---|
| `identity` | Declare the complete artifact identity grammar under `### Value`. |
| `enum` | Declare every allowed complete H1 prefix under `### Values`. |
| `literal` | Declare the one allowed complete H1 prefix under `### Value`. |

The declaration uses this common field:

```markdown
## H1 prefix

- **type**: `<identity-or-enum-or-literal>`
```

An `identity` declaration uses this form:

````markdown
### Value

```text
<IDENTITY_GRAMMAR>
```
````

An `enum` declaration uses this form:

```markdown
### Values

| value |
|---|
| `<ALLOWED_VALUE>` |
```

A `literal` declaration uses this form:

````markdown
### Value

```text
<LITERAL_VALUE>
```
````

The artifact-specific source Specification states the actual grammar, complete allowed values, or literal directly.
A Specification reference does not replace this declaration.
When the type is `identity`, the declared value must match the identity form in the artifact-specific identity Specification.

### H1-adjacent metadata placement

Artifact sources begin with this structure:

```text
# <H1_PREFIX>: <Title>

<H1-adjacent metadata>

## <FIRST_H2_HEADING>
```

The H1-adjacent metadata block appears after H1 and before the first H2 heading.
`<FIRST_H2_HEADING>` is the heading declared by the first row of the H2 heading table.
The first row must use `always`.

Artifact-specific field inventory and value formats belong to the artifact-specific H1-adjacent metadata Specification.
Shared field notation is defined by `spec:drmcp.design_records_mcp.artifacts.base.definitions.h1_adjacent_metadata`.

### H2 declaration

Artifact-specific source Specifications declare H2 headings, their presence conditions, and any separate body-format rules by using this table shape:

| column | content |
|---|---|
| `heading` | Exact H2 heading text, including the `##` marker. |
| `condition` | `always`, `optional`, `recommended`, `prohibited`, or an artifact-specific condition expression. |
| `format reference` | Canonical Specification ref for heading-specific body-format rules, or `-`. |

The `condition` values have these meanings:

| condition | rule |
|---|---|
| `always` | The heading appears exactly once. |
| `optional` | The heading appears zero or one time. |
| `recommended` | The heading should appear zero or one time. |
| `prohibited` | The heading does not appear. |
| `<condition expression>` | The expression assigns one presence condition to each artifact-specific category. |

A condition expression with multiple category clauses uses this form inside the table cell:

```text
<category>: <condition>;<br/><category>: <condition>;
```

Each clause uses one of `always`, `optional`, `recommended`, or `prohibited`.
A semicolon terminates each clause.
`<br/>` separates adjacent clauses.

The table row order represents the preferred source heading order.
DRMCP does not require source headings to follow the table row order.

Each artifact-specific source Specification must declare one unlisted-heading policy:

| policy | rule |
|---|---|
| `allowed` | An H2 heading not listed in the declaration table may appear when the artifact's Product authority permits it. |
| `prohibited` | An H2 heading not listed in the declaration table does not appear. |

`format reference` does not define heading presence.
It names a separate Specification only when that Specification owns a heading-specific body structure or format.
Use `-` when no separate heading-specific body format applies.
Product authority for heading meaning may be referenced once from the artifact-specific source Specification instead of repeated in every row.

## Boundary

| concern | owner |
|---|---|
| Shared H1 shape | This Specification. |
| Shared H1-adjacent metadata placement | This Specification. |
| H2 declaration columns and condition vocabulary | This Specification. |
| H1 prefix declaration types and forms | This Specification. |
| Artifact-specific H1 prefix grammar, allowed values, or literal | Artifact-specific source Specification. |
| Canonical artifact identity represented by an `identity` prefix | Artifact-specific identity Specification. |
| Artifact-specific metadata fields and value formats | Artifact-specific H1-adjacent metadata Specification. |
| Artifact-specific H2 inventory | Artifact-specific source Specification. |
| Artifact-specific unlisted-heading policy | Artifact-specific source Specification. |
| Artifact-specific heading presence conditions | Artifact-specific source Specification. |
| Heading-specific body format | Specification named by `format reference`. |
| Product-owned authoring meaning | Referenced Product Specifications. |
| Parsing, diagnostics, and validation orchestration | Other DRMCP Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration` | Shared identity declaration forms used when `<H1_PREFIX>` contains record identity. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.h1_adjacent_metadata` | Shared H1-adjacent metadata notation. |
| `spec:product.design_records.authoring_standards` | Product authority for artifact authoring meaning. |
