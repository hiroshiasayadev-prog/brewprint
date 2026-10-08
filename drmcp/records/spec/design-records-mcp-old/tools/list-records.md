# Contract: `list_records`

- **id**: `spec:drmcp.design_records_mcp.tools.list_records`
- **status**: draft
- **date**: 2026-06-28
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`list_records` returns a compact projection of current addressable records from one active app, sequential kind, and domain scope.

The operation narrows candidates before exact retrieval. It is not a broad repository inventory, exact lookup, range query, spec-tree query, resolver, or validation operation.

## Non-goals

- Exact current or legacy record retrieval.
- Spec listing or spec-tree navigation.
- Cross-app, cross-kind, or cross-domain broad listing.
- ID ranges, cursors, totals, or unbounded listing.
- Reference resolution or legacy fallback orchestration.
- Validation execution or diagnostic taxonomy.
- Full metadata, headings, body, source provenance, or physical-path projection.
- Fixture, implementation, or automated-test behavior.

## Request

```json
{
  "app_namespace": "drmcp",
  "kind": "task",
  "domain": "MCP",
  "status": "in_progress",
  "order": "desc",
  "limit": 20
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `app_namespace` | yes | non-empty string | Selects exactly one configured active app namespace. |
| `kind` | yes | string | Selects exactly one supported sequential kind. |
| `domain` | yes | non-empty string | Selects exactly one canonical domain namespace within the app and kind scope. |
| `status` | no | non-empty string | Exact status filter from the selected kind's PRODUCT-owned vocabulary. |
| `order` | no | string | Canonical `ref` order. Default `desc`; alternative `asc`. |
| `limit` | no | integer | Maximum returned record count. Default 20; accepted range 1 through 100. |

Supported `kind` values:

| value | record family |
|---|---|
| `decision` | Current ADR records. |
| `investigation` | Current investigation records. |
| `requirement` | Current requirement records. |
| `work_item` | Current work-item records. |
| `task` | Current task records. |

`spec` is not a supported normal-listing kind. Legacy archive records are not query candidates for this operation.

The operation uses the configured `app_namespace` value and canonical domain namespace value without case repair, path inference, range-derived scope, or app-prefix inference.

A syntactically valid domain selector with no matching addressable records is a valid zero-match query. An `app_namespace` that does not identify a configured current root is an invalid request.

When `status` is present:

- the value must belong to the selected kind's PRODUCT-owned vocabulary;
- matching is exact;
- a record with missing `status` does not match;
- a parsed invalid status does not match any accepted status filter unless PRODUCT authority accepts that exact value.

The operation rejects:

- missing, null, empty, or non-string `app_namespace`, `kind`, or `domain`;
- unsupported or non-string `kind` values;
- empty, non-string, or kind-invalid `status` values;
- `order` values other than `asc` or `desc`;
- non-integer `limit` values or integers outside 1 through 100;
- unsupported request fields, including `id`, `id_range`, `order_by`, `head`, and `tail`.

The request never broadens when a required scope field is missing or invalid.

## Response

Normal response shape:

```json
{
  "records": [
    {
      "ref": "DRMCP-TASK-MCP-004-02",
      "title": "Reflect compact active list-records contract",
      "status": "in_progress",
      "date": "2026-06-27"
    }
  ],
  "has_more": false,
  "warnings": []
}
```

Top-level fields:

| field | required | meaning |
|---|---:|---|
| `records` | yes | Compact results after scope, optional status filtering, ordering, and limit. |
| `has_more` | yes | `true` when at least one additional matching addressable record exists after the returned page. |
| `warnings` | yes | Operation-level warning entries. Empty when no warning trigger fires. W006 owns entry taxonomy and representation. |

Each `records[]` entry has this fixed shape:

| field | required | type | meaning |
|---|---:|---|---|
| `ref` | yes | string | Canonical current record ID-as-ref. |
| `title` | yes | string or null | Parsed H1 title, or `null` when missing. |
| `status` | yes | string or null | Parsed lifecycle status, or `null` when missing. |
| `date` | yes | string or null | Parsed source date, or `null` when missing. |

The operation applies filters before ordering and limit.
It orders records by the canonical `ref` string.
Default order is descending. Ascending is the only alternative.
No secondary physical-path key exists.

`has_more` does not expose the total matching count. The response does not contain a cursor, offset, or repository-wide count.

### Invalid-but-addressable records

Every addressable record remains eligible for normal listing.

| source state | projection |
|---|---|
| All compact fields are available | Return all four fields with parsed values. |
| `title`, `status`, or `date` is missing | Return the missing field as `null`. |
| A compact field has a parsed invalid value | Return the parsed value unchanged. Do not replace, normalize, or repair it. |

Each missing compact field on a returned record triggers one top-level operation warning.
The warning is not nested inside the compact result.
W006 defines the warning category, severity, shared fields, and source-location representation.

A parsed invalid value does not trigger a T02-defined warning category. Validation and diagnostic representation remain W006 responsibilities.

### Duplicate-conflict identities

A canonical identity in W003 duplicate-conflict state has no addressable winner.
`list_records` returns no result for that identity.
The operation does not merge sources, select by traversal order, or use physical path as a tie-break.

Duplicate-conflict diagnostics and source-location representation remain outside this operation contract.

### Zero-match example

A valid request with no matching addressable records returns:

```json
{
  "records": [],
  "has_more": false,
  "warnings": []
}
```

Zero matches are not an error and do not trigger a warning.

## Errors

| code | condition |
|---|---|
| `invalid_request` | Request shape, required scope, supported kind, status vocabulary, ordering, limit, or accepted field set is invalid. |

Invalid requests return no partial listing response.
When a required source-backed warning location cannot be constructed, the operation fails without returning a partial listing response. This contract does not assign a new error identifier for that execution failure.
Exact error representation remains part of the DRMCP response-boundary contracts.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.namespace_model` | App and domain namespace semantics. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Canonical sequential record identity. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Configured current roots and active-index-only scope. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Addressable, invalid-but-addressable, and duplicate-conflict states. |
| `spec:drmcp.design_records_mcp.schema.fields` | Parsed common field vocabulary and missing-value boundary. |
| `spec:drmcp.design_records_mcp.schema.diagnostics` | W006-owned warning and diagnostic representation. |
| `spec:drmcp.design_records_mcp.tools.get_records` | Exact retrieval operation owned by the next W004 Task. |

## Sources

- `DRMCP-TASK-MCP-004-01`: Accepted query and exact-retrieval correction baseline.
- `DRMCP-TASK-MCP-004-02`: Compact active listing contract reflection and evidence.
