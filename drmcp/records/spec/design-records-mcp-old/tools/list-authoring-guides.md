# Contract: `list_authoring_guides`

- **id**: `spec:drmcp.design_records_mcp.tools.list_authoring_guides`
- **status**: draft
- **date**: 2026-07-04
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`list_authoring_guides` returns authoring-standard child Specs from the portable `design_records` Current Records scope as a compact Guidance projection.

The operation is a fixed-scope Application Use Case over shared record-query orchestration. It does not call the public `list_records` use case and does not create a separate guide index.

## Request

```json
{}
```

No request fields. Implementations may reject requests containing unknown fields with `invalid_request`.

## Response

The operation applies this fixed Current Records scope:

| scope | value |
|---|---|
| app namespace | `design_records` |
| record kind | `spec` |
| included canonical refs | `spec:design_records.authoring_standards.*` |
| excluded root | `spec:design_records.authoring_standards` |

Every addressable current Spec in that child subtree is a catalog candidate.
A candidate is projectable only when its canonical identity, first H1, `## What this is` body, and source read are available.
The operation does not broaden to another app, kind, topic, physical path, or legacy source.

```json
{
  "guides": [
    {
      "id": "spec:design_records.authoring_standards.adr_authoring",
      "title": "Reference: ADR authoring",
      "abstract": "Authoring rules for Architecture Decision Record artifacts.\n\nThis guide defines ADR IDs, paths, file shape, metadata, lifecycle, references, and author-facing inputs."
    }
  ]
}
```

`guides[]` fields:

| field | required | meaning |
|---|---:|---|
| `id` | yes | Canonical package Spec ref. |
| `title` | yes | First H1 text from the current Spec source. |
| `abstract` | yes | Body of the `## What this is` section. |

Normal responses MUST NOT include physical source paths.

`guides[]` is ordered by canonical `id` in ASCII lexical order.
The operation uses shared current-record query primitives without exposing generic query fields in its request.

## Errors

| code | condition |
|---|---|
| `invalid_request` | The request is not the accepted empty-object shape. |
| `guidance_catalog_unavailable` | An in-scope canonical identity is conflicted, unreadable, or cannot supply the required title or abstract projection. |

A valid request with no in-scope addressable child Specs returns `guides: []`.
The operation does not omit an unprojectable in-scope candidate and return a misleading partial catalog.
Current source invalidity, duplicate identity, and source-read state remain owned by the normal Current Records model. The operation does not select a duplicate winner or fall back to physical-path enumeration.
