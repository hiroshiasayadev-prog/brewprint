# Contract: `get_records`

- **id**: `spec:drmcp.design_records_mcp.tools.get_records`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`get_records` retrieves the same record representation as `get_record` for multiple explicitly specified record IDs in a single batch call. It does not perform candidate search, filter, or range query (those are `list_records`), canonical reference resolution (`resolve_reference`), or index integrity validation (`validate_records`).

> Source: V01-ADR-090 §1–§3

## Request

```json
{
  "ids": [
    "V01-ADR-077",
    "V01-SPEC-design-records-mcp-tools",
    "V01-INV-DOCS-001",
    "V01-ADR-077",
    "V01-INV-DOCS-999"
  ],
  "include_body": false
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `ids` | yes | non-empty array of string | Exact record ID lookup keys. Input order is preserved. |
| `include_body` | no | bool | Whether to include Markdown raw body in each found record. Default: `false` |

Missing `ids`, empty array, non-array value, or non-string elements produce an `invalid_request` tool error.

`ids[]` items are evaluated as exact lookup keys against the record index only. Whitespace trimming, case normalization, canonical reference resolution, and input kind classification are not performed. Therefore strings such as `spec:trace`, physical paths, `adr-077`, ` V01-ADR-077 `, or grammar-invalid workflow IDs that do not match an indexed record ID produce an item-level `not_found`, not a tool error or `unsupported`.

`get_records` does not accept `kind` / `status` / `id_range` / `limit` as request fields, and does not accept per-item `include_body`.

## Response

```json
{
  "items": [
    {
      "id": "V01-ADR-077",
      "retrieval_status": "found",
      "record": {
        "id": "V01-ADR-077",
        "kind": "decision",
        "title": "Design Records MCP MVP boundary and tool prioritization",
        "status": "accepted",
        "path": "v01/records/adr/V01-ADR-077-design-records-mcp-mvp-boundary-and-tool-prioritization.md",
        "decision": {
          "depends_on": ["V01-ADR-076"],
          "supersedes": [],
          "migrated_to_spec": null
        },
        "headings": []
      },
      "diagnostics": []
    },
    {
      "id": "V01-INV-DOCS-999",
      "retrieval_status": "not_found",
      "record": null,
      "diagnostics": [
        {
          "category": "record_not_found",
          "severity": "error",
          "requested_id": "V01-INV-DOCS-999",
          "message": "record V01-INV-DOCS-999 was not found"
        }
      ]
    }
  ],
  "diagnostics": [
    {
      "category": "duplicate_requested_id_ignored",
      "severity": "info",
      "requested_id": "V01-ADR-077",
      "first_index": 0,
      "duplicate_indexes": [3],
      "message": "duplicate requested record ID was ignored after its first occurrence"
    }
  ]
}
```

Top-level response fields:

| field | required | meaning |
|---|---:|---|
| `items` | yes | Retrieval item list in first-occurrence-of-deduplicated-IDs order |
| `diagnostics` | yes | Request-level diagnostic list; empty list on success |

Retrieval item fields:

| field | required | meaning |
|---|---:|---|
| `id` | yes | The lookup key supplied in the request |
| `retrieval_status` | yes | `found` / `not_found` |
| `record` | yes | For `found`: same representation as `get_record.record`. For `not_found`: `null` |
| `diagnostics` | yes | Item-level diagnostic list; empty list for `found` items |

The collection is named `items`, not `records`, because missing items are included in the same collection.

When all IDs are not found, the response is still a normal response (not a tool error); each first-occurrence input returns a `retrieval_status: "not_found"` item.

When the same ID appears multiple times in `ids`, only the first occurrence is included in `items`. For each duplicated ID, one `duplicate_requested_id_ignored` top-level diagnostic is returned. `first_index` and `duplicate_indexes` are zero-based indexes into the request `ids` array.

When `include_body: true`, each `found` item's `record.body` contains the original Markdown in full. Formatting, summarization, normalization, and truncation are prohibited, matching `get_record` behavior. No public numeric limit on response total length or body size is defined.

> Source: V01-ADR-090 §4–§7

## Errors

| code | condition |
|---|---|
| `invalid_request` | `ids` is missing, empty, non-array, or contains non-string elements |
