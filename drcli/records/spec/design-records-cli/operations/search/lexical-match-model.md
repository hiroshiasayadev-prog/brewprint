# Contract: Lexical match model

- **id**: `spec:drcli.design_records_cli.operations.search.lexical_match_model`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations.search`
- **contract_class**: `interface`
## What this is

Defines lexical matching modes, searchable target slices, and match-occurrence rules used by `search_records`.

## Non-goals

- Semantic similarity, fuzzy matching, synonym inference, stemming, or morphological analysis.
- Canonical-ref search, completion, repair, or relation traversal.
- Artifact-specific metadata field or H2 heading inventories.
- Complete record or H2 section retrieval.
- Detailed source-defect reporting.

## Request

This contract consumes these `search_records` request fields:

| field | applicability | value |
|---|---|---|
| `query` | literal and pattern mode | Non-empty source-text query. |
| `mode` | always | `literal` or `pattern`; default `literal`. |
| `case_sensitive` | literal mode only | Optional boolean; default `true`. |
| `targets` | always | Unordered selection of supported search targets; omitted or empty selects all targets. |

The logical request shape prohibits `case_sensitive` when `mode` is `pattern`.
The lexical record search contract owns the complete request shape and validation sequence.

## Matching modes

### Literal mode

Literal mode performs substring matching against each selected target source slice.
The query is not interpreted as pattern syntax.

`case_sensitive` controls literal comparison:

| value | comparison |
|---|---|
| `true` | Compare source and query Unicode scalar values exactly. |
| `false` | Compare through locale-independent Unicode simple case folding. |

Literal mode defaults to `case_sensitive: true`.
Case folding affects comparison only.
Returned snippets preserve original source text and source positions.

### Pattern mode

Pattern mode interprets `query` using RE2 syntax and semantics.
Supported constructs and inline flags follow RE2.
DRCLI does not redefine or maintain a separate RE2 flag inventory.

Pattern mode does not accept the `case_sensitive` request field.
The pattern controls case, anchor, dot, and multiline behavior through RE2 syntax.

A pattern is invalid when either condition holds:

- the query is not valid RE2 syntax;
- the pattern can produce a zero-length match.

Both conditions produce request-wide `invalid_pattern`.
Implementation-specific compilation messages, source positions, and engine details are not returned.

## Match occurrences

DRCLI evaluates each target instance independently.
A match cannot cross a target-instance boundary.

Literal and pattern modes use left-to-right, non-overlapping occurrences.
The next search begins at the end of the previous match.

| mode | occurrence selection |
|---|---|
| `literal` | Select the earliest next matching substring. |
| `pattern` | Use RE2 leftmost-first matching for the earliest next occurrence. |

Every reported occurrence has non-zero length.

A query may contain newline characters.
Literal matching may cross source lines inside one target instance.
Pattern matching may cross source lines when RE2 syntax and pattern flags permit it.
Neither mode may cross from one H2 section target instance into another.

DRCLI does not normalize Unicode, whitespace, Markdown, or line endings before matching.
DRCLI does not rewrite source text or infer equivalent words.

## Search targets

The supported search targets are:

```text
title
metadata_source
h2_heading
h2_section_content
```

Artifact Specifications remain authoritative for H1, H1-adjacent metadata, real H2 headings, and section boundaries.
Search applies the shared slices below without duplicating artifact-specific field or heading inventories.

### Title target

One record has at most one `title` target instance.
The slice begins after the H1 line's first `: ` delimiter and ends before the H1 line ending.

The slice excludes:

- the leading `# ` marker;
- the H1 identity, enum, or literal prefix;
- the `: ` delimiter;
- the H1 line ending.

DRCLI does not trim or normalize title text.
The slice matches the title source text used by exact record retrieval.

### Metadata source target

One record has at most one `metadata_source` target instance.
The slice begins immediately after the H1 line ending.
The slice ends immediately before the first real H2 heading line.

The slice preserves:

- field names and field values;
- list markers, emphasis markers, separators, and indentation;
- blank lines;
- source line endings.

Search does not parse the slice into a metadata object.
Malformed metadata remains searchable when both slice boundaries are reliable.

### H2 heading target

Each real H2 heading produces one `h2_heading` target instance.
The slice begins after the leading `## ` marker and ends before the heading line ending.

The slice preserves exact H2 title text.
DRCLI does not trim, normalize, or deduplicate title text.
Duplicate H2 titles produce separate target instances in source order.

### H2 section content target

Each real H2 heading produces one `h2_section_content` target instance.
The slice begins immediately after the selected H2 heading line ending.
The slice ends immediately before the next real H2 heading line or at source end.

The slice excludes the selected H2 heading line.
The slice includes all H3-or-deeper headings and their content.
The slice preserves Markdown markers, whitespace, blank lines, and source line endings.

An empty section produces an empty target instance and no match.
Duplicate H2 titles produce separate section target instances in source order.

## Response

This contract establishes match occurrences and target-extraction warnings for the search response.
Each established occurrence supplies:

- the containing canonical ref;
- the selected target;
- the target-instance source order;
- the match start and end positions within that target instance;
- the exact H2 title for H2 target instances.

The search result contract converts each returnable occurrence into one bounded match entry.
This contract does not return complete target content.

## Target extraction limitations

A nonconforming admitted record remains eligible for search.
Failure to extract one requested target does not remove the record or another extractable target.

| target | unavailable condition |
|---|---|
| `title` | The H1 title slice cannot be determined reliably. |
| `metadata_source` | The H1 or first-real-H2 boundary cannot be determined reliably. |
| `h2_heading` | Real H2 heading instances cannot be projected reliably. |
| `h2_section_content` | Real H2 section boundaries cannot be projected reliably. |

A reliable H2 projection with no real H2 heading produces zero H2 target instances without a warning.

`search_target_unavailable` applies only when record source is available but the requested structural slice cannot be projected reliably.
A source-access or execution failure that prevents reliable search completion produces request-wide `execution_failure`.

For each unavailable requested `record × target` pair, DRCLI:

- skips only that pair;
- continues every other extractable pair;
- emits one top-level `search_target_unavailable` warning;
- includes `canonical_ref` and `target` in warning context;
- does not include defect details or physical source paths.

Multiple defects within one pair still produce one warning.
Warnings use this deterministic order:

```text
canonical_ref ascending simple string order
  -> title
  -> metadata_source
  -> h2_heading
  -> h2_section_content
```

A pair without an extraction limitation produces no warning.
Detailed conformance findings and repair guidance belong to Validation Specifications.

## Errors

| condition | classification | context |
|---|---|---|
| `query` is not valid RE2 syntax in pattern mode | `invalid_pattern` | None. |
| The pattern can produce a zero-length match | `invalid_pattern` | None. |

Both conditions are request-wide errors.
A target-extraction limitation is not an error.
The limitation produces top-level `search_target_unavailable` and preserves every other reliable target.
Shared diagnostic representation belongs to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| Artifact H1, metadata, and H2 source structure | Applicable artifact Specifications. |
| Literal and RE2 matching semantics | This Specification. |
| Target vocabulary and source slices | This Specification. |
| Target-extraction warning trigger and order | This Specification. |
| Request fields and response envelope | `spec:drcli.design_records_cli.operations.search.lexical_record_search`. |
| Match projection, snippets, ordering, and limits | `spec:drcli.design_records_cli.operations.search.search_results_and_limits`. |
| Warning object and canonical context contract | `spec:drcli.design_records_cli.diagnostics`. |
| Detailed source defects and physical paths | Validation Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts` | Defines shared H1, metadata placement, and H2 source shape. |
| `spec:drcli.design_records_cli.artifacts` | Defines metadata source notation preserved by search. |
| `spec:drcli.design_records_cli.operations.search.lexical_record_search` | Supplies the search request and selected records. |
| `spec:drcli.design_records_cli.operations.search.search_results_and_limits` | Projects each established match occurrence. |
| `spec:drcli.design_records_cli.operations.retrieval.exact_record_retrieval` | Defines the corresponding title projection. |
| `spec:drcli.design_records_cli.diagnostics` | Defines `invalid_pattern` and `search_target_unavailable`. |