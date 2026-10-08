# Contract: Artifact source template

- **id**: `spec:drcli.design_records_cli.artifacts.base.templates.source`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.artifacts.base.templates`
- **contract_class**: `format`

## What this is

Provides the template for an artifact-specific `source.md` Specification.

The artifact-specific Specification declares its complete H1 prefix grammar or values, H2 headings, their presence conditions, and any separate body-format rules.
The H1 prefix declaration uses exactly one of `identity`, `enum`, or `literal` and states the actual value information directly rather than replacing it with a Specification reference.
For a sequential artifact using an `identity` H1 prefix, include the `## Identity authority` section shown below and omit it for other structures or prefix types.
Use `always`, `optional`, `recommended`, or `prohibited` when one condition applies to the artifact kind.
Use an artifact-specific condition expression when presence varies by category.
Separate category clauses with `<br/>`, and write each clause as `<category>: <condition>;`.
`format reference` is used only when another Specification owns a heading-specific body structure or format.

## Current contract

This Specification defines the DRCLI-local projection and declarations for this topic.
Where this Specification cites PRODUCT authority, the cited PRODUCT Specification remains authoritative for PRODUCT-owned semantics.

## Rules

- DRCLI-local declarations must preserve cited PRODUCT-owned semantics.
- A DRCLI-local declaration must not create an independent replacement for cited PRODUCT authority.

## Validation rules

- A DRCLI interpretation that conflicts with cited PRODUCT authority is nonconforming.
- Validation of PRODUCT-owned meaning uses the cited PRODUCT authority.

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

<SEQUENTIAL_IDENTITY_AUTHORITY_SECTION_IF_APPLICABLE>

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
| `spec:drcli.design_records_cli.artifacts.base.definitions.source` | Shared source-document rules. |
| `<PRODUCT_AUTHORITY_REF>` | Product authority consumed by this artifact Specification. |
````

For a sequential artifact using an `identity` H1 prefix, replace `<SEQUENTIAL_IDENTITY_AUTHORITY_SECTION_IF_APPLICABLE>` with:

```markdown
## Identity authority

The complete H1 prefix is the sole source-internal identity authority.
<ARTIFACT_NAME> metadata and the file name do not supply, repair, complete, normalize, infer, or replace the identity.
The file-name public-ID prefix is validated separately as a conformance projection.
```

For other structures or prefix types, remove the placeholder.

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
