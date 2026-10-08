# Reference: Authoring guidance source

- **id**: `spec:drmcp.design_records_mcp.schema.authoring_guidance_source`
- **status**: draft
- **date**: 2026-07-04
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines the fixed Current Records scope and projection used by the authoring-guidance operations.

## Current contract

Authoring guidance is a public projection over ordinary current Spec records.
It is not a separate source, record kind, parser, index, or logical model.

### Current Records source

The portable standards package is configured as a spec-tree Current Records source with:

```text
app namespace: design_records
root Spec: spec:design_records
```

Package files use normal current Spec discovery, path-derived identity, active-index, retrieval, resolution, and validation behavior.
DRMCP performs no runtime namespace rewrite.

### Guidance scope

Normal list scope is fixed to addressable child Specs below:

```text
spec:design_records.authoring_standards.*
```

The root Spec is excluded from normal Guidance listing:

```text
spec:design_records.authoring_standards
```

Detail lookup accepts one exact canonical ref inside the child subtree.
Filename stems, basenames, physical paths, titles, aliases, fuzzy values, and inferred candidates are not Guidance identities.

### Guide projection

| field | projection rule |
|---|---|
| `id` | Canonical package Spec ref. |
| `title` | First H1 text. |
| `abstract` | Body of the `## What this is` section. |
| `content` | Complete Markdown source verbatim. |

Record Domain / Logical Tree owns canonical identity and parsed source structure.
The Guidance Application Use Cases own fixed scope, canonical-ref ordering, response shape, and operation-specific errors.

### Tool response contracts

| tool | returns |
|---|---|
| `list_authoring_guides` | `id` / `title` / `abstract` |
| `get_authoring_guidance` | `id` / `title` / `content` |

Normal responses do not expose physical source paths.
Current invalid-source, duplicate-identity, and addressability behavior remains owned by the normal Current Records contracts.
