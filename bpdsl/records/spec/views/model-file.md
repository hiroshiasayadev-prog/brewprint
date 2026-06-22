# Contract: Model file render rules

- **id**: `spec:bpdsl.views.model_file`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.views.overview`
- **contract_class**: `format`

## What this is

Model file render is the human-facing Markdown render for `model/*.yaml` files. One model YAML file produces one model-file Markdown render, showing the public main model and the file-private helper models defined in the same YAML file. This satisfies the requirement that helper models must not become invisible in human-facing render surfaces (V01-ADR-070).

Model file render is not a curated module-wide catalog — module / contract-level catalog views are a separate, future view kind.

This spec defines the model-file render minimum for model kinds: `struct`, `enum`, `list`, `dict`, `tagged_union`.

> Source: V01-ADR-070, V01-ADR-072, V01-ADR-075

## Current contract

Model file render uses Markdown tables, not Mermaid diagrams. The recommended output structure is:

````markdown
# {public main model id}

## Public model

| property | value |
|---|---|
| kind | {kind} |
| visibility | public |
| source | {source YAML path} |

{kind-specific public model section}

## Private models

File-private helper schemas defined in this model YAML file.
Promote a helper model to a public model file when it needs to be reused from other YAML files.

| model | kind | shape | note |
|---|---|---|---|
| `{helper model id}` | `{kind}` | `{compact shape}` | `{note summary}` |
````

### Public model section by kind

**Struct** — fields render as a Markdown table:

```markdown
### Fields

| field | type | note |
|---|---|---|
| `{field name}` | `{TypeRef}` | `{note summary}` |
```

**Enum** — values render as a Markdown table:

```markdown
### Values

| value | note |
|---|---|
| `{enum value}` | `{optional note if available}` |
```

**List** — element TypeRef renders as a property table:

```markdown
### Element

| property | value |
|---|---|
| element | `{TypeRef}` |
```

**Dict** — value TypeRef renders as a property table:

```markdown
### Value

| property | value |
|---|---|
| value | `{TypeRef}` |
```

**Tagged union** — discriminator and each variant render as nested sections:

```markdown
### Discriminator

| property | value |
|---|---|
| discriminator | `{discriminator field name}` |

### Variants

#### `{variant tag}`

| field | type | note |
|---|---|---|
| `{field name}` | `{TypeRef}` | `{note summary}` |

#### `{payload-less variant tag}`

No payload fields.
```

Each variant is rendered as a level-4 heading with its tag.

### Private models table

| column | meaning |
|---|---|
| `model` | Helper model local ID. |
| `kind` | Helper model kind. |
| `shape` | Compact kind-specific shape summary. |
| `note` | Helper model note summary. |

`shape` should stay compact: struct helper models may list fields as `field: TypeRef` pairs; enum helper models may list enum values; list/dict helper models may show `element: TypeRef` or `value: TypeRef`; tagged union helper models may show `discriminator: X` followed by `tag: T` entries.

The table does not define signature-exposure policy — rules for helper models appearing in task `params` or `returns` are owned by [`spec:bpdsl.dsl.nodes.processing`](../dsl/nodes/processing.md).

## Rules

Model file render output placement follows [`spec:bpdsl.dsl.project_layout`](../dsl/project-layout.md). The filename convention is:

```text
model-{model-id}.md
```

The exact group / directory placement is owned by that spec, not this one.

## Validation rules

- If the model file has no file-private helper models, the renderer may omit the `## Private models` section entirely.
- A variant with no payload fields (`fields: []`) renders `No payload fields.` instead of a field table — it must not render an empty table.
- Enum values without a per-value note may leave the `note` column empty rather than fabricating a value.
- Render input is assumed already semantically valid; this contract does not define model-definition validation diagnostics. Model kind / field / enum / tagged-union validation is owned by [`spec:bpdsl.dsl.nodes.data`](../dsl/nodes/data.md) and [`spec:bpdsl.dsl.diagnostics`](../dsl/diagnostics.md).

## Non-goals

Out of scope for this model-file render minimum:

- Model catalog render implementation (module / contract-level aggregation).
- DAG asset TypeRef hint rendering (owned by [`spec:bpdsl.views.dag`](dag.md)).
- MCP helper-model exposure / semantic identity (owned by the `bpdsl.mcp` namespace).
- UC-002 model-response helper migration beyond tagged-union migration.

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.views.overview` | Parent overview; view kind catalog. |
| `spec:bpdsl.dsl.nodes.data` | `model` node kind definitions this render renders. |
| `spec:bpdsl.dsl.nodes.processing` | Task-file helper model signature-exposure rules, not owned by this render contract. |
| `spec:bpdsl.dsl.project_layout` | Render output placement and filename convention. |
| `spec:bpdsl.views.dag` | Owns DAG asset TypeRef hint rendering and DAG-side private-model exposure. |
