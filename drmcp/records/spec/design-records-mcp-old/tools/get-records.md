# Contract: `get_records`

- **id**: `spec:drmcp.design_records_mcp.tools.get_records`
- **status**: draft
- **date**: 2026-06-28
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
| Current path-derived document-level `spec:` grammar | `spec:product.design_records.spec_format.spec_id_as_ref` | Active index only. |
| Accepted Brewprint legacy sequential ID | `spec:product.brewprint.compatibility.legacy_id_compatibility` | Configured legacy index only. |

The PRODUCT compatibility authority defines the accepted legacy families. This contract does not extend that family set.
The exact legacy issued-ID lexical mapping is defined by `spec:drmcp.design_records_mcp.namespace_scanning`.

### Exact input classification order

`get_records` applies no resolver-style current-first fallback sequence.
Each first-occurrence input is classified exactly once in this order:

1. If the input matches one exact issued-ID grammar from `spec:drmcp.design_records_mcp.namespace_scanning` for a PRODUCT-accepted legacy family, classify it as legacy and query only the configured legacy lookup map.
2. Otherwise, if the input matches the active path-derived `spec:` grammar, classify it as a current spec and query only the active index.
3. Otherwise, if the input matches the current sequential canonical grammar, classify it as current sequential and query only the active index.
4. Otherwise, classify it as malformed or unsupported under the W004/W006 response boundary.

Legacy-family precedence is an exact-retrieval classification rule, not reference fallback.
It ensures that an overlapping value such as `V01-REQ-MCP-001` is classified as legacy and never queried against the active index by `get_records`.
The operation does not invoke `resolve_reference`, run a second lookup after failure, or guess between scopes.

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

### Rejected string items

The following table applies after request-shape validation.
Duplicate handling still uses the exact ordered-deduplication contract below.
Every listed value remains a string item and does not become an `invalid_request` condition.

| input class | item behavior | lookup behavior |
|---|---|---|
| `V01-SPEC-*` | Malformed or unsupported item warning trigger. | No active-index or legacy lookup. |
| App-prefixless sequential ID | Malformed or unsupported item warning trigger. | No app, domain, or prefix inference. |
| Physical path | Malformed or unsupported item warning trigger. | No path-to-ref conversion or filesystem lookup. |
| Fuzzy or partial reference | Malformed or unsupported item warning trigger. | No matching, completion, or second lookup. |
| Legacy YAML-only alias spelling that fails current canonical grammar | Malformed or unsupported item warning trigger. | No front-matter alias lookup or repair. |
| Direct `yaml:` input | Malformed or unsupported item warning trigger. | No YAML lookup surface is active. |
| `fixture:` input | Malformed or unsupported item warning trigger. | No fixture lookup surface is active. |
| `internal-design:` input | Malformed or unsupported item warning trigger. | No internal-design lookup surface is active. |
| `coverage:` input | Malformed or unsupported item warning trigger. | No coverage lookup surface is active. |
| `COV-*` input | Malformed or unsupported item warning trigger. | No coverage-ID lookup surface is active. |
| Section-like `spec:` value that does not match current spec grammar | Malformed or unsupported item warning trigger. | No repair, section-target lookup, heading lookup, or alias lookup. |
| Metadata-only alias spelling that fails current canonical grammar | Malformed or unsupported item warning trigger. | Referring metadata does not register a lookup alias. |
| Value requiring whitespace, case, prefix, domain, or sequence repair | Malformed or unsupported item warning trigger. | No trimming, repair, completion, or second lookup. |
| Empty string | Malformed item warning trigger. | No lookup. |

Current spec classification is lexical and does not use semantic origin.
A supplied `spec:` string that matches current spec grammar is classified as a current spec input and queries only the active index, even when an earlier document used the same string as a section alias.
When that exact lookup finds no addressable target, the item produces the unresolved-current warning trigger already defined by the partial-success contract.
The operation does not consult front-matter `semantic_refs`, front-matter `sections`, heading data, or alias registries and does not perform section-target or heading lookup.

The table does not define warning category names, severity, messages, shared fields, ref association, or source-location representation.
Those concerns remain owned by `DRMCP-WORK-MCP-006`.

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
| `metadata` | yes | Current normalized metadata, or the minimal legacy metadata defined below, excluding canonical identity because `ref` carries it. |
| `headings` | yes | Real ATX headings in source order, represented as `level` and `text`. Empty when no heading was parsed. |
| `body` | conditional | Complete source Markdown, present only when `include_body` is `true`. |

For a current record, `metadata` contains `kind` plus every parsed common and kind-specific field available under `spec:drmcp.design_records_mcp.schema.fields`.
Kind-specific current fields appear directly in `metadata` under their normalized field names.

For an invalid but addressable current record:

- a missing parsed metadata field is omitted;
- a parsed invalid value is returned unchanged;
- a default, repaired value, empty collection, or `null` is not invented for a missing field;
- addressability and retrieval success remain governed by `spec:drmcp.design_records_mcp.schema.record_model`.

When `include_body` is `false`, the `body` field is omitted. The field is not returned as `null` or an empty string.

When `include_body` is `true`, `body` contains the complete source Markdown verbatim. The operation must not format, summarize, normalize, or truncate the body.

### Legacy exact retrieval

An accepted legacy input queries only the separate configured legacy lookup map.
The issued legacy ID is derived from the source filename and compared exactly and case-sensitively.
`get_records` does not invoke current-first resolver orchestration, query the active index for that input, or normalize the archived source into the current record model.

Legacy retrieval succeeds when exactly one indexed source exists and the source file can be read.
Incomplete or malformed optional Markdown structure does not prevent source retrieval.

A successful legacy record uses the common record wrapper with these rules:

- `ref` preserves the exact issued legacy ID;
- `metadata.kind` is required and is derived from the accepted legacy ID family;
- `metadata.title` and `metadata.status` are included only when they can be parsed;
- no current kind-specific metadata field is required for legacy compatibility;
- `headings` contains parsed real ATX headings, or an empty array when none can be parsed;
- `body` contains the complete source Markdown verbatim only when `include_body` is `true`.

DRMCP does not repair, default, normalize, or validate archived metadata as a retrieval precondition.
A missing H1, malformed bullet metadata, or mismatch between source content and the filename-derived ID does not block retrieval of a readable unique source and does not create an alias.

A legacy input produces no successful record when:

- no legacy lookup map is configured;
- the exact issued ID is absent;
- duplicate sources prevent one source from being selected; or
- the indexed source cannot be read.

These conditions remain normal partial-success warning triggers.
Warning categories, severity, messages, and source-location representation remain owned by `DRMCP-WORK-MCP-006`.

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

When a required warning, conflict-member, conflict-candidate, or unreadable-source location cannot be constructed, the operation fails without returning a partial normal response. This contract does not assign a new error identifier for that execution failure.

Malformed, unsupported, unresolved, unavailable legacy, and duplicate string items are not `invalid_request` conditions. They produce warning triggers in a normal partial-success response only when every required source-backed location can be constructed.

## Boundary

| concern | owner |
|---|---|
| Current-root discovery, source parsing, identity, active-index construction, normalized fields, invalid-source retention, and duplicate conflicts | `DRMCP-WORK-MCP-003` |
| Exact-retrieval request, ordering, partial success, warning placement, body inclusion, and normal record projection | `DRMCP-WORK-MCP-004` |
| Configured legacy roots, filename-derived legacy identity, separate lookup-map construction, duplicate legacy IDs, and legacy source readability | `DRMCP-WORK-MCP-005` |
| Warning entry schema, category names, severity, shared fields, source locations, validation behavior, and exceptional path exposure | `DRMCP-WORK-MCP-006` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.schema.fields` | Parsed normalized metadata vocabulary for current records. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Current-record addressability and invalid-source retention. |
| `spec:drmcp.design_records_mcp.schema.record_source` | Current heading and body source availability. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Configured legacy roots and exact legacy lookup-map construction. |
| `spec:drmcp.design_records_mcp.tools.get_record` | Retired single-record operation. |
| `DRMCP-TASK-MCP-004-03` | Exact batch-retrieval request and common response reflection owner. |
| `DRMCP-TASK-MCP-005-03` | Minimal legacy source retrieval reflection owner. |
| `DRMCP-TASK-MCP-005-04` | Rejected string-item and cross-spec pointer synchronization owner. |
