# Reference: Authoring guidance source

- **id**: `spec:drmcp.design_records_mcp.schema.authoring_guidance_source`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines how authoring guides are discovered, how guide IDs and titles are derived, and what each authoring tool returns.

## Current contract

Authoring guidance is treated as a separate read-only guidance source, distinct from the Design Records record model.

### Guide source

Guide source directory:

```text
docs/guides/
```

Guide file discovery rule:

```text
docs/guides/*.md
```

### Guide ID and title

Guide ID is the filename stem.

Example:

```text
docs/guides/adr-authoring.md → adr-authoring
```

Guide title is extracted from the first H1. Guide abstract is extracted from the body of the `## Abstract` section.

### Tool response contracts

| tool | returns |
|---|---|
| `list_authoring_guides` | `id` / `title` / `abstract` |
| `get_authoring_guidance` | `id` / `title` / `content` |

Guide source file path is not part of the public response contract. Path is resolved internally from the guide ID as an implementation detail.

Guides are not record kinds. They are not treated as Design Records record IDs, record status, record paths, record headings, record diagnostics, or canonical reference resolver targets.
