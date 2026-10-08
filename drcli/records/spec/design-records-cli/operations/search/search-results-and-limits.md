# Contract: Search results and limits

- **id**: `spec:drcli.design_records_cli.operations.search.search_results_and_limits`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.operations.search`
- **contract_class**: `interface`
## What this is

Defines search match entries, bounded snippets, deterministic ordering, and output limits for `search_records`.

## Non-goals

- Defining search scopes, query modes, or target extraction.
- Relevance ranking, scoring, personalization, or pagination.
- Returning complete record or H2 section content as a search projection.
- Returning physical source paths or validation findings.
- Defining transport-wide response-size limits.

## Request

The result contract consumes these optional request fields:

| field | value | minimum | default | maximum |
|---|---|---:|---:|---:|
| `match_limit` | Maximum returned match entries. | `1` | `50` | `200` |
| `snippet_limit` | Maximum snippet size in Unicode scalar values. | `1` | `240` | `2000` |

A caller-specified value is not clamped.
The response reports both effective values as `applied_match_limit` and `applied_snippet_limit`.

The per-record match cap is fixed at `10`.
The caller cannot change the per-record cap.

## Response

A normal response contains:

```text
matches:
  - canonical_ref
    target
    h2_title?
    snippet
    match_span:
      start
      end
    snippet_truncated_before
    snippet_truncated_after
applied_match_limit
applied_snippet_limit
has_additional_matches
```

The operation may also include top-level `warnings` as defined by the lexical search contracts.
A response without an applicable warning omits `warnings`.

### Match entry

Each item in `matches` represents one non-overlapping match occurrence.

| field | presence | value |
|---|---|---|
| `canonical_ref` | always | Canonical ref of the containing uniquely addressable current record. |
| `target` | always | `title`, `metadata_source`, `h2_heading`, or `h2_section_content`. |
| `h2_title` | for `h2_heading` and `h2_section_content` | Exact H2 title text without the leading `## ` marker. |
| `snippet` | always | One preserved continuous source slice containing the complete match. |
| `match_span` | always | Half-open location of the match within `snippet`. |
| `snippet_truncated_before` | always | Whether source text exists before `snippet` in the same target instance. |
| `snippet_truncated_after` | always | Whether source text exists after `snippet` in the same target instance. |

`h2_title` identifies the matched heading for `h2_heading`.
`h2_title` identifies the enclosing heading for `h2_section_content`.
The field is absent for `title` and `metadata_source`.

The match entry does not contain:

- a duplicated `matched_text` field;
- a record title projection;
- an absolute source offset;
- a physical source path.

### Match span

`match_span` uses this logical shape:

```text
start
end
```

Both values count Unicode scalar values from the beginning of `snippet`.
`start` is zero-based and inclusive.
`end` is exclusive.
Every match satisfies:

```text
0 <= start < end <= snippet scalar-value count
```

The source text selected by the half-open span is the complete matched source text.
DRCLI does not insert highlight markers into `snippet`.

### Snippet construction

DRCLI first constructs the smallest line-aligned source slice that contains the complete match.
The slice begins at the start of the first source line touched by the match.
The slice ends at the end of the last source line touched by the match.
A source line ending is included when the line ending belongs to the target instance.

| condition | snippet rule |
|---|---|
| The line-aligned slice fits `snippet_limit`. | Return the complete line-aligned slice. |
| The line-aligned slice exceeds `snippet_limit`, but the match fits. | Return one match-centered continuous slice at most `snippet_limit` scalar values long. |
| The complete match exceeds `snippet_limit`. | Omit the complete match entry. |

For a match-centered slice, DRCLI allocates the remaining capacity within the line-aligned slice.
DRCLI allocates capacity equally before and after the match.
An odd remaining scalar value is allocated after the match.
Capacity unavailable at one line-aligned-slice boundary is reassigned to the other side.
DRCLI does not add unrelated lines merely to fill `snippet_limit`.

The snippet preserves source text exactly.
DRCLI does not summarize, normalize, reformat, or insert omission markers.
Source line endings remain unchanged.
A CRLF line ending counts as two Unicode scalar values.

`snippet_truncated_before` is `true` only when the target instance has source text before the returned snippet.
`snippet_truncated_after` is `true` only when the target instance has source text after the returned snippet.

A bounded snippet may equal the complete target instance when the target instance already fits the snippet rule.
Search does not provide a projection that requests or guarantees complete record or H2 section content.
For `h2_section_content`, the snippet never includes the owning H2 heading line.

### Deterministic ordering

DRCLI orders all established match occurrences by this tuple:

```text
canonical_ref
-> target order
-> target-instance source order
-> match start position
-> match end position
```

`canonical_ref` uses ascending simple string order.
The fixed target order is:

```text
title
metadata_source
h2_heading
h2_section_content
```

A record has at most one title instance and one metadata-source instance.
H2 heading and H2 section instances use source order.
Matches inside one target instance use source position order.

The same ordering applies to literal and pattern mode.
Scope occurrence order, target request order, filesystem order, scan order, and records-root discovery order do not affect result order.

### Limit application

DRCLI applies result selection in this order:

```text
all established matches
  -> deterministic ordering
  -> omit matches that cannot produce a complete bounded snippet
  -> retain at most 10 matches per canonical_ref
  -> retain at most applied_match_limit matches
```

An oversized match does not consume the per-record cap or total match limit.
The per-record cap retains the first ten returnable entries for each canonical ref under deterministic ordering.
The total limit then retains the first entries from the resulting ordered collection.

A result entry is returned completely or omitted completely.
DRCLI does not partially truncate an entry to satisfy a result limit.

### Additional-match indicator

`has_additional_matches` is always present.
The value is `true` when at least one match entry is omitted because of:

- an oversized match under `snippet_limit`;
- the fixed per-record cap;
- the applied total match limit.

The value is `false` when none of those conditions omits a match.
Target-extraction incompleteness is represented by `search_target_unavailable`, not by this indicator.

`matches: []` may coexist with `has_additional_matches: true` when every established match is oversized.
Reaching a limit remains normal success.
The response does not include an omitted-match count, an omission-reason list, an offset, or a continuation token.

### Empty result

A valid search with no returned entry contains:

```text
matches: []
applied_match_limit
applied_snippet_limit
has_additional_matches
```

A valid scope with no searchable record produces this normal response with `has_additional_matches: false`.

## Errors

| condition | classification | context |
|---|---|---|
| `match_limit` is outside `1..200` | `invalid_match_limit` | `minimum`, `maximum`, `default`, `actual`. |
| `snippet_limit` is outside `1..2000` | `invalid_snippet_limit` | `minimum`, `maximum`, `default`, `actual`. |
| Reliable ordering, snippet construction, or result completion fails | `execution_failure` | None. |

Limit validation precedes current-state construction.
A request-wide error returns no normal result fields or warnings.
Shared diagnostic representation and context contracts belong to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| Match occurrence establishment and target instances | `spec:drcli.design_records_cli.operations.search.lexical_match_model`. |
| Match projection, span, snippet construction, ordering, and output limits | This Specification. |
| Search request envelope and validation sequence | `spec:drcli.design_records_cli.operations.search.lexical_record_search`. |
| Complete record and H2 section retrieval | `spec:drcli.design_records_cli.operations.retrieval`. |
| Shared error and warning representation | `spec:drcli.design_records_cli.diagnostics`. |
| Detailed source defects and physical paths | Validation Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.operations.search.lexical_record_search` | Supplies request fields and the normal response envelope. |
| `spec:drcli.design_records_cli.operations.search.lexical_match_model` | Supplies ordered match occurrences and target-instance boundaries. |
| `spec:drcli.design_records_cli.operations.retrieval.h2_section_retrieval` | Owns complete H2 section retrieval. |
| `spec:drcli.design_records_cli.operations.retrieval.exact_record_retrieval` | Owns complete record retrieval. |
| `spec:drcli.design_records_cli.diagnostics` | Defines request-wide error and warning placement. |
| `spec:drcli.design_records_cli.diagnostics` | Defines search limit error codes. |