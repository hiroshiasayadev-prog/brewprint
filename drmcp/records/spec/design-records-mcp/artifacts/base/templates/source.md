# Contract: Artifact source template

- **id**: `spec:drmcp.design_records_mcp.artifacts.base.templates.source`
- **status**: draft
- **date**: 2026-07-11
- **parent**: `spec:drmcp.design_records_mcp.artifacts.base.templates`
- **contract_class**: `format`

## What this is

Provides the template for an artifact-specific `source.md` Specification.

The artifact-specific Specification declares its complete H1 prefix grammar or values, H2 headings, their presence conditions, and any separate body-format rules.
The H1 prefix declaration uses exactly one of `identity`, `enum`, or `literal` and states the actual value information directly rather than replacing it with a Specification reference.
Use `always`, `optional`, `recommended`, or `prohibited` when one condition applies to the artifact kind.
Use an artifact-specific condition expression when presence varies by category.
Separate category clauses with `<br/>`, and write each clause as `<category>: <condition>;`.
`format reference` is used only when another Specification owns a heading-specific body structure or format.

## Template

````markdown
# Contract: <ARTIFACT_NAME> source

- **id**: `<SPEC_REF>`
- **status**: draft
- **date**: `<YYYY-MM-DD>`
- **parent**: `<PARENT_SPEC_REF>`
- **contract_class**: `format`

## What this is

Defines the source-document shape for `<ARTIFACT_NAME>` records.

## H1 prefix

<H1_PREFIX_DECLARATION>

## H2 heading policy

- **unlisted headings**: `<allowed-or-prohibited>`

## H2 headings

| heading | condition | format reference |
|---|---|---|
| `## <FIRST_H2_HEADING>` | `always` | `<SPEC_REF_OR_DASH>` |
| `## <HEADING>` | `<always-or-optional-or-recommended-or-prohibited-or-condition>` | `<SPEC_REF_OR_DASH>` |
| `## <CONDITIONAL_HEADING>` | `<CATEGORY>: <CONDITION>;<br/><CATEGORY>: <CONDITION>;` | `<SPEC_REF_OR_DASH>` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.source` | Shared source-document rules. |
| `<PRODUCT_AUTHORITY_REF>` | Product authority consumed by this artifact Specification. |
````

Replace `<H1_PREFIX_DECLARATION>` with exactly one of the following forms.

### Identity prefix

````markdown
- **type**: `identity`

### Value

```text
<IDENTITY_GRAMMAR>
```
````

### Enum prefix

```markdown
- **type**: `enum`

### Values

| value |
|---|
| `<ALLOWED_VALUE>` |
```

### Literal prefix

````markdown
- **type**: `literal`

### Value

```text
<LITERAL_VALUE>
```
````
