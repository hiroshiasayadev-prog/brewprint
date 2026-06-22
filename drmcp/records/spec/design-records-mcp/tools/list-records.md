# Contract: `list_records`

- **id**: `spec:drmcp.design_records_mcp.tools.list_records`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`list_records` returns a structured record index for decision / spec / investigation / requirement / work_item / task records. Its purpose is to narrow down candidate records before reading Markdown body. It is not a simple filesystem listing — it returns normalized record metadata and H1 titles extracted from bullet metadata or YAML front matter.

> Source: V01-ADR-077 §list_records の責務

## Request

MVP request schema:

```json
{
  "kind": "decision",
  "status": "accepted",
  "id": "V01-ADR-076",
  "id_range": {
    "from": "V01-ADR-067",
    "to": "V01-ADR-077"
  },
  "order_by": "id",
  "order": "asc",
  "limit": 20
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `kind` | no | string | Filter by `decision` / `spec` / `investigation` / `requirement` / `work_item` / `task` |
| `status` | no | string | Filter by status value |
| `id` | no | string | Filter by exact ID |
| `id_range` | no | object | Filter by ID range (inclusive both ends) |
| `order_by` | no | string | MVP: `id` only |
| `order` | no | string | `asc` / `desc` |
| `limit` | no | integer | Maximum result count |

`order_by` supports `id` only in MVP. `head` / `tail` are not supported.

`id_range`: either `from` or `to` may be omitted for a one-sided range. Endpoints are specified as public IDs (with namespace_prefix).

`id_range` is applicable to these ID families:

| family | effective `kind` | endpoint form (MVP) | ordering |
|---|---|---|---|
| decision | `decision` | `V01-ADR-NNN` | numeric comparison of bare `NNN` |
| requirement | `requirement` | `V01-REQ-<DOMAIN>-NNN` | numeric comparison of bare `NNN` within the same `<DOMAIN>` |
| work item | `work_item` | `V01-WORK-<DOMAIN>-NNN` | numeric comparison of bare `NNN` within the same `<DOMAIN>` |
| task | `task` | `V01-TASK-<DOMAIN>-NNN-MM` | numeric comparison of task sequence `MM` within the same `<DOMAIN>` and work sequence `NNN` |

When `kind` is omitted and `id_range` is specified, the effective `kind` is derived from the endpoint family. When `kind` is specified, the endpoint family must match.

Workflow artifact range is limited to the same family and same domain. Task range is additionally limited to the same work sequence; a range spanning multiple work items (e.g. `V01-TASK-MCP-006-01` to `V01-TASK-MCP-007-05`) is a request error.

One-sided workflow range is evaluated within the family / domain / work sequence scope of the specified endpoint. For example, `kind: work_item` with `id_range.from: V01-WORK-DATA-004` targets `WORK-DATA-*` with sequence ≥ 004.

`SPEC-*` / `INV-*` range, mixed family, mixed domain, mixed task work sequence, malformed endpoints, and `kind`/endpoint family mismatch are request errors. Silent fallback to lexical ordering or broad listing is prohibited.

## Response

```json
{
  "records": [
    {
      "id": "V01-ADR-076",
      "kind": "decision",
      "title": "Design Records MCP",
      "status": "accepted",
      "path": "v01/records/adr/V01-ADR-076-design-records-mcp.md",
      "decision": {
        "depends_on": ["V01-ADR-050", "V01-ADR-068"],
        "supersedes": [],
        "migrated_to_spec": null
      }
    }
  ]
}
```

`records[]` is ordered by `order_by` / `order`. For mixed-kind results with `order_by: id`, canonical `id` ASCII lexical ordering is used. When duplicate canonical IDs exist, tie-breaking is by path ASCII lexical order; a `duplicate_id` diagnostic is returned separately.

When `id_range` is specified, range membership follows the ID family rules above. Response ordering follows `order_by` / `order` as normal.

## Errors

| code | condition |
|---|---|
| `invalid_request` | Unknown `kind` value or other invalid request field |
| `invalid_id_range` | Malformed, unsupported, mixed-family, mixed-domain, mixed-task-work-sequence, or kind-mismatched endpoint |
