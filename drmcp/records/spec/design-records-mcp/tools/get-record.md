# Contract: `get_record`

- **id**: `spec:drmcp.design_records_mcp.tools.get_record`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`get_record` retrieves metadata, path, headings, and optionally Markdown body from a single record ID. Enabling body retrieval here reduces the need to return to a filesystem tool after candidate narrowing with `list_records`.

> Source: V01-ADR-077 §get_record の責務

## Request

```json
{
  "id": "V01-ADR-076",
  "include_body": false
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `id` | yes | string | Target record ID |
| `include_body` | no | bool | Whether to return the Markdown raw body. Default: `false` |

## Response

### Response without body

```json
{
  "record": {
    "id": "V01-ADR-076",
    "kind": "decision",
    "title": "Design Records MCP",
    "status": "accepted",
    "path": "v01/records/adr/V01-ADR-076-design-records-mcp.md",
    "decision": {
      "depends_on": ["V01-ADR-050", "V01-ADR-068"],
      "supersedes": [],
      "migrated_to_spec": null
    },
    "headings": [
      { "level": 1, "text": "076: Design Records MCP" },
      { "level": 2, "text": "背景" },
      { "level": 2, "text": "決定" }
    ]
  }
}
```

### Response with body

When `include_body: true`, `record.body` is added with the Markdown file content verbatim.

```json
{
  "record": {
    "id": "V01-ADR-076",
    "kind": "decision",
    "title": "Design Records MCP",
    "status": "accepted",
    "path": "v01/records/adr/V01-ADR-076-design-records-mcp.md",
    "decision": {
      "depends_on": ["V01-ADR-050", "V01-ADR-068"],
      "supersedes": [],
      "migrated_to_spec": null
    },
    "headings": [],
    "body": "# 076: Design Records MCP\n\n- **status**: accepted\n..."
  }
}
```

`body` returns the original file content verbatim. Formatting, summarization, normalization, and truncation are prohibited. Structured metadata and headings are returned as separate fields, not removed from body.

## Errors

| code | condition |
|---|---|
| `record_not_found` | Specified record ID does not exist in the index |
