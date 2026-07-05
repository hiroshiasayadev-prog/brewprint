# Contract: `get_authoring_guidance`

- **id**: `spec:drmcp.design_records_mcp.tools.get_authoring_guidance`
- **status**: draft
- **date**: 2026-07-04
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`get_authoring_guidance` returns one exact authoring-standard child Spec from the portable `design_records` Current Records scope as a Guidance projection.

The operation reuses shared exact-retrieval orchestration. It does not call the public `get_records` use case and does not maintain a separate guide lookup map.

## Request

```json
{
  "id": "spec:design_records.authoring_standards.adr_authoring"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `id` | yes | string | Exact canonical package Spec ref inside the Guidance child subtree. |

Accepted IDs match this fixed scope:

```text
spec:design_records.authoring_standards.*
```

The root `spec:design_records.authoring_standards` is not a normal Guidance detail target.

The input is evaluated exactly as supplied.
Whitespace trimming, case normalization, basename lookup, filename-stem lookup, physical-path lookup, title lookup, aliases, fuzzy lookup, inferred candidates, and reference-resolution fallback are not performed.

## Response

Response shape:

```json
{
  "id": "spec:design_records.authoring_standards.adr_authoring",
  "title": "Reference: ADR authoring",
  "content": "<complete Markdown source returned verbatim>"
}
```

The `content` placeholder documents the response field only. An implementation returns the complete source string and does not return the placeholder or a truncated example.

| field | required | meaning |
|---|---:|---|
| `id` | yes | Exact canonical package Spec ref. |
| `title` | yes | First H1 text from the current Spec source. |
| `content` | yes | Complete Markdown source verbatim. |

Formatting, summarization, normalization, and truncation are prohibited.
Normal responses MUST NOT include physical source paths, record metadata, or internal index state.

## Errors

| code | condition |
|---|---|
| `guide_not_found` | The exact canonical ref is absent or outside the accepted Guidance child subtree. |
| `guide_unavailable` | The in-scope canonical identity is conflicted, unreadable, or cannot supply the required title or complete content projection. |
| `invalid_request` | The request shape or `id` type is invalid. |

The operation does not select a duplicate winner or map an in-scope conflicted identity to `guide_not_found`.
