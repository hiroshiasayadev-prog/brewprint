---
scope: docs/spec/views/model-file.md
status: draft
last_updated: 2026-06-03
summary: >
  Model file render minimum. Defines Markdown render output for model YAML files,
  including public main model and file-private helper models for struct / enum / list / dict / tagged_union.
depends_on:
  - docs/adr/070-model-visibility-file-private-helper-model.md
  - docs/adr/072-model-schema-catalog-view.md
  - docs/adr/075-model-file-render.md
---

# Model file render rules

## Scope

Model file render is the human-facing Markdown render for `model/*.yaml` files.

One model YAML file produces one model-file Markdown render.

The render shows the public main model and the file-private helper models defined in the same YAML file. This satisfies the ADR-070 constraint that helper models must not become invisible in human-facing render surfaces.

Model file render is not a curated module-wide catalog. Module / contract-level catalog views are owned by the model catalog view.

## Execution boundary

This spec defines the model-file render minimum for the following model kinds:

- `struct`
- `enum`
- `list`
- `dict`
- `tagged_union` (added in WORK-DATA-010)

## Output format

Model file render uses Markdown tables, not Mermaid diagrams.

The recommended output structure is:

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

If the model file has no file-private helper models, the renderer may omit the `## Private models` section.

## Public model section by kind

### Struct

Struct models render their fields as a Markdown table.

```markdown
### Fields

| field | type | note |
|---|---|---|
| `{field name}` | `{TypeRef}` | `{note summary}` |
```

### Enum

Enum models render their values as a Markdown table.

```markdown
### Values

| value | note |
|---|---|
| `{enum value}` | `{optional note if available}` |
```

If enum values do not have per-value notes, the `note` column may be empty.

### List

List models render their element TypeRef.

```markdown
### Element

| property | value |
|---|---|
| element | `{TypeRef}` |
```

### Dict

Dict models render their value TypeRef.

```markdown
### Value

| property | value |
|---|---|
| value | `{TypeRef}` |
```

### Tagged union

Tagged union models render the discriminator and each variant.

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

Each variant is rendered as a level-4 heading with its tag. Variants with no payload fields (`fields: []`) render `No payload fields.` instead of a table.

## Private models table

The `## Private models` table lists file-private helper models defined in the same model YAML file.

| column | meaning |
|---|---|
| `model` | helper model local ID |
| `kind` | helper model kind |
| `shape` | compact kind-specific shape summary |
| `note` | helper model note summary |

The `shape` column should stay compact. For struct helper models, it may list fields as ``field: TypeRef`` pairs. For enum helper models, it may list enum values. For list / dict helper models, it may show `element: TypeRef` or `value: TypeRef`. For tagged union helper models, it may show `discriminator: X` followed by `tag: T` entries.

The table does not define signature exposure policy. Rules for helper models appearing in task `params` or `returns` are tracked separately by REQ-DATA-003.

## Output placement

Model file render output placement follows the render output placement rules in `docs/spec/project-layout.md`.

The initial filename convention is:

```text
model-{model-id}.md
```

The exact group / directory placement is owned by `docs/spec/project-layout.md`.

## Excluded scope

The following are not part of this model-file render minimum:

- model catalog render implementation
- DAG asset TypeRef hint
- MCP helper model exposure / semantic identity
- UC-002 model response helper migration beyond tagged union migration
