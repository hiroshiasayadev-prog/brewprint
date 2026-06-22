# Contract: `get_authoring_guidance`

- **id**: `spec:drmcp.design_records_mcp.tools.get_authoring_guidance`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`get_authoring_guidance` retrieves authoring guidance Markdown content by guide ID. It returns the guide Markdown verbatim; it does not return record metadata, record path, record headings, or record lifecycle status.

## Request

```json
{
  "id": "adr-authoring"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `id` | yes | string | Target guide ID (exact match) |

`id` is evaluated as an exact guide ID lookup key. Whitespace trimming, case normalization, physical path lookup, and record ID resolution are not performed.

## Response

```json
{
  "id": "adr-authoring",
  "title": "ADR Authoring Guide",
  "content": "# ADR Authoring Guide\n\n## Abstract\n\n..."
}
```

| field | required | meaning |
|---|---:|---|
| `id` | yes | Guide ID from the request |
| `title` | yes | First H1 text from the guide file |
| `content` | yes | Full Markdown content of the guide file |

`content` returns the original Markdown verbatim. Formatting, summarization, normalization, and truncation are prohibited.

Guide source file path MUST NOT be included in the response.

## Errors

| code | condition |
|---|---|
| `guide_not_found` | Specified guide ID does not exist |
