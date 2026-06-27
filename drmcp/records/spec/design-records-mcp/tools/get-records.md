# Contract: `get_records`

- **id**: `spec:drmcp.design_records_mcp.tools.get_records`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`get_records` is the sole public exact-retrieval operation for current and configured legacy design records.

The operation does not search, filter, resolve, validate, repair, or infer references.

## Request

```json
{
  "refs": [
    "DRMCP-ADR-MCP-001",
    "spec:drmcp.design_records_mcp.schema.fields",
    "V01-INV-MCP-001"
  ],
  "include_body": false
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `refs` | yes | array of string | One through 20 exact retrieval inputs. |
| `include_body` | no | boolean | Include complete source Markdown in every successful record. Default: `false`. |

Only `refs` and `include_body` are accepted top-level fields.

The request is rejected when any request-shape condition applies:

| condition | result |
|---|---|
| `refs` is missing | Reject the request. |
| `refs` is not an array | Reject the request. |
| `refs` contains zero elements | Reject the request. |
| `refs` contains more than 20 elements | Reject the request. |
| Any `refs` element is not a string | Reject the request. |
| `include_body` is present and is not boolean | Reject the request. |
| Any unsupported top-level field is present | Reject the request. |

An empty string remains a string-shaped item. It is processed as a malformed exact input and is not trimmed, removed, or promoted to a request-shape error.

### Accepted exact input families

| input family | authority | lookup scope |
|---|---|---|
| Current sequential canonical ref | `spec:product.design_records.namespace_model.artifact_id_grammar` | Active index only. |
| Active path-derived `spec:` ref | `spec:product.design_records.spec_format.spec_id_as_ref` | Active index only. |
| Accepted Brewprint legacy sequential ID | `spec:product.brewprint.compatibility.legacy_id_compatibility` | Configured legacy index only. |

The PRODUCT compatibility authority defines the accepted legacy families. This contract does not copy or extend that grammar.

`get_records` applies no current-first fallback sequence. Each exact input is classified once and queried only in the scope assigned to its accepted family.

### Exactness and prohibited repair

Each string is interpreted exactly as supplied.

The operation must not:

- infer an app prefix;
- trim or repair whitespace;
- repair case;
- interpret a physical path;
- perform fuzzy matching;
- complete a partial ID;
- guess between current and legacy families;
- invoke `resolve_reference`;
- use metadata aliases or source paths as lookup keys.

### Ordered deduplication

Duplicate detection uses exact string equality before lookup.

- The first occurrence remains eligible for classification and lookup.
- Later equal strings produce duplicate warning triggers.
- Later duplicate occurrences produce no additional record.
- Successful records follow the request order of their first occurrences.
- Different strings are not merged through normalization or aliasing.

### Partial success

The following per-ref outcomes produce operation-level warning triggers:

- malformed exact input;
- unsupported input family;
- unresolved accepted current or spec ref;
- accepted legacy ref unavailable because no configured legacy index is available;
- unresolved accepted legacy ref;
- duplicate input occurrence.

A failed ref does not discard any retrievable record from another ref.

When every first-occurrence ref fails, the operation returns a normal response with `records: []` and operation warnings.

Warning category names, severity, shared fields, ref association, occurrence details, and source-location representation are owned by `DRMCP-WORK-MCP-006`.

## Response

```json
{
  "records": [
    {
      "ref": "DRMCP-ADR-MCP-001",
      "metadata": {
        "kind": "decision",
        "title": "Design Records MCP contract baseline and realignment",
        "status": "accepted",
        "date": "2026-06-25",
        "depends_on": ["PRODUCT-ADR-SPEC-001"],
        "supersedes": [
          "V01-ADR-076",
          "V01-ADR-077",
          "V01-ADR-087",
          "V01-ADR-088",
          "V01-ADR-090",
          "V01-ADR-092",
          "V01-ADR-093"
        ],
        "migrated_to_spec": null
      },
      "headings": [
        {
          "level": 1,
          "text": "DRMCP-ADR-MCP-001: Design Records MCP contract baseline and realignment"
        },
        {
          "level": 2,
          "text": "Context"
        }
      ]
    }
  ],
  "warnings": []
}
```

Top-level response fields:

| field | required | meaning |
|---|---:|---|
| `records` | yes | Successful records only, in first-occurrence request order. |
| `warnings` | yes | Operation-level warning entries. Empty when no warning trigger occurs. |

No failure placeholder is returned in `records`.

The response does not contain `items`, `retrieval_status`, `record: null`, or per-item diagnostic wrappers.

### Successful record projection

| field | required | meaning |
|---|---:|---|
| `ref` | yes | Canonical public identity of the retrieved record. Legacy results preserve the issued legacy ID. |
| `metadata` | yes | Parsed normalized metadata fields defined by `spec:drmcp.design_records_mcp.schema.fields`, excluding canonical identity because `ref` carries it. |
| `headings` | yes | Real ATX headings in source order, represented as `level` and `text`. Empty when no heading was parsed. |
| `body` | conditional | Complete source Markdown, present only when `include_body` is `true`. |

`metadata` contains `kind` plus every parsed common and kind-specific field available for the record.
Kind-specific fields appear directly in `metadata` under their normalized field names.

For an invalid but addressable record:

- a missing parsed metadata field is omitted;
- a parsed invalid value is returned unchanged;
- a default, repaired value, empty collection, or `null` is not invented for a missing field;
- addressability and retrieval success remain governed by `spec:drmcp.design_records_mcp.schema.record_model`.

When `include_body` is `false`, the `body` field is omitted. The field is not returned as `null` or an empty string.

When `include_body` is `true`, `body` contains the complete source Markdown verbatim. The operation must not format, summarize, normalize, or truncate the body.

### Path and internal-state boundary

Normal retrieval records must not contain:

- physical path;
- source file location;
- source provenance;
- active-index or legacy-index path;
- internal index state;
- resolver trace.

This operation does not authorize path exposure through warnings. Diagnostic, patch, debug, and emergency path exposure remains owned by `DRMCP-WORK-MCP-006`.

## Errors

| code | condition |
|---|---|
| `invalid_request` | The top-level request violates the request-shape rules. |

Malformed, unsupported, unresolved, unavailable legacy, and duplicate string items are not `invalid_request` conditions. They produce warning triggers in a normal partial-success response.

## Boundary

| concern | owner |
|---|---|
| Current-root discovery, source parsing, identity, active-index construction, normalized fields, invalid-source retention, and duplicate conflicts | `DRMCP-WORK-MCP-003` |
| Exact-retrieval request, ordering, partial success, warning placement, body inclusion, and normal record projection | `DRMCP-WORK-MCP-004` |
| Resolver invocation, current-first resolution, configured legacy fallback, fallback order, and current/legacy orchestration | `DRMCP-WORK-MCP-005` |
| Warning entry schema, category names, severity, shared fields, source locations, validation behavior, and exceptional path exposure | `DRMCP-WORK-MCP-006` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.schema.fields` | Parsed normalized metadata vocabulary. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Addressability and invalid-source retention. |
| `spec:drmcp.design_records_mcp.schema.record_source` | Heading and body source availability. |
| `spec:drmcp.design_records_mcp.tools.get_record` | Retired single-record operation. |
| `DRMCP-TASK-MCP-004-03` | Exact batch-retrieval reflection owner. |
