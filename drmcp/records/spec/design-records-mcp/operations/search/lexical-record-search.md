# Contract: Lexical record search

- **id**: `spec:drmcp.design_records_mcp.operations.search.lexical_record_search`
- **status**: draft
- **date**: 2026-07-14
- **parent**: `spec:drmcp.design_records_mcp.operations.search`
- **contract_class**: `interface`
- **usdm_covers**:
  - usdm:drmcp.mcp_capabilities.search.lexical_record_search#R001-R007

## What this is

Defines the `search_records` operation for scoped lexical search over uniquely addressable current records.

## Non-goals

- Defining current scope availability or record addressability.
- Defining artifact-specific source formats.
- Returning complete record or H2 section content.
- Returning validation findings or physical source paths.
- Defining shared diagnostic objects or code-specific context contracts.

## Request

The request accepts these fields:

| field | requirement | value | default |
|---|---|---|---|
| `query` | required | Non-empty lexical query string. | - |
| `mode` | optional | `literal` or `pattern`. | `literal` |
| `case_sensitive` | optional for literal mode; prohibited for pattern mode | Boolean literal-comparison policy. | `true` in literal mode |
| `scopes` | optional | List of sequential or tree scope objects. | Repository-wide search |
| `targets` | optional | List containing supported search-target values. | All supported targets |
| `match_limit` | optional | Integer from `1` through `200`. | `50` |
| `snippet_limit` | optional | Integer from `1` through `2000`, measured in Unicode scalar values. | `240` |

Input-schema violations are rejected before operation result construction.
Schema violations include:

- missing `query`;
- unknown request fields;
- wrong value types;
- unknown `mode`, `scope_type`, or target values;
- malformed discriminator-owned scope object shapes;
- missing scope-object fields;
- scope-type-specific fields on another scope type;
- `case_sensitive` supplied with `mode: pattern`.

An empty `query` string is schema-valid and produces operation-level `empty_query`.

### Scope objects

A sequential scope uses this logical shape:

```text
scope_type: sequential
app_namespace
artifact_kind
domain_namespace
```

A tree scope uses this logical shape:

```text
scope_type: tree
subtree_root_ref
```

`subtree_root_ref` is the exact canonical ref of the selected current tree node.
The ref supplies the artifact kind, app namespace, and subtree position.
The tree scope does not repeat those values in separate fields.

DRMCP does not trim, repair, normalize, complete, or infer any scope selector.

### Scope collection

The following requests select the repository current-state boundary:

- `scopes` omitted;
- `scopes: []`.

Repository-wide search includes uniquely addressable current records across all configured app namespaces.

A non-empty `scopes` collection selects the union of all supplied current-record scopes.
Duplicate scope objects have the same effect as one occurrence.
Overlapping scopes include each uniquely addressable current record at most once.
Scope occurrence order does not affect match ordering.

DRMCP validates non-empty scopes in request occurrence order.
One invalid or unavailable scope produces one request-wide error.
The operation returns no partial search result from the remaining scopes.

A valid scope with no addressable record contributes an empty set.
A valid final union with no searchable record produces normal empty success.

### Target selection

The supported target values are:

```text
title
metadata_source
h2_heading
h2_section_content
```

`targets` is an unordered selection set.
Duplicate values have the same effect as one occurrence.
The following requests select all supported targets:

- `targets` omitted;
- `targets: []`.

Target input order does not affect match ordering.
Target source slices and extraction outcomes belong to the lexical match model.

### Request validation order

The operation evaluates request-wide outcomes in this order:

1. `match_limit` bounds.
2. `snippet_limit` bounds.
3. Non-empty `query`.
4. RE2 pattern validity when `mode` is `pattern`.
5. Pattern zero-length-match prohibition when `mode` is `pattern`.
6. Configured current-state construction.
7. Scope semantic validity and availability in request occurrence order.
8. Requested target extraction.
9. Match evaluation.
10. Ordering, snippet construction, and output-limit application.
11. Reliable operation completion.

Input-schema validation precedes this order.
The first established request-wide error is returned.
Current-state-independent validation precedes configured current-state construction.

A target-extraction limitation produces a warning and does not stop the order.
The operation continues with every reliably extractable requested target.

## Response

A normal response uses this logical shape:

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
warnings?
```

The response fields have these owners:

| field | owner |
|---|---|
| `matches` and each match entry | `spec:drmcp.design_records_mcp.operations.search.search_results_and_limits`. |
| `applied_match_limit` | `spec:drmcp.design_records_mcp.operations.search.search_results_and_limits`. |
| `applied_snippet_limit` | `spec:drmcp.design_records_mcp.operations.search.search_results_and_limits`. |
| `has_additional_matches` | `spec:drmcp.design_records_mcp.operations.search.search_results_and_limits`. |
| `warnings` | This Specification and the lexical match model define the trigger; Diagnostics define representation. |

A valid search with no returned match contains `matches: []`.
Normal responses do not echo the query, mode, scopes, or targets.
Normal responses do not expose physical source paths.

## Warnings

The operation permits top-level `search_target_unavailable` warnings.
The lexical match model defines the trigger and deterministic order.

A warning may coexist with matches, applied limits, and `has_additional_matches`.
A response without applicable warnings omits `warnings`.

## Errors

| condition | classification | operation-specific context |
|---|---|---|
| `match_limit` is outside `1..200` | `invalid_match_limit` | `minimum`, `maximum`, `default`, `actual`. |
| `snippet_limit` is outside `1..2000` | `invalid_snippet_limit` | `minimum`, `maximum`, `default`, `actual`. |
| `query` is empty | `empty_query` | None. |
| Pattern syntax is not valid RE2 syntax | `invalid_pattern` | None. |
| A pattern can produce a zero-length match | `invalid_pattern` | None. |
| Sequential app namespace is unavailable | `invalid_selector` | `app_namespace`, `available_app_namespaces`. |
| Sequential artifact kind is unknown, non-sequential, or unavailable within the selected app | `invalid_selector` | `artifact_kind`, `available_sequential_artifact_kinds`. |
| Sequential domain is unavailable within the selected app and artifact kind | `invalid_selector` | `domain_namespace`, `available_domain_namespaces`. |
| `subtree_root_ref` is malformed or identifies a non-tree artifact kind | `invalid_selector` | `subtree_root_ref`, `accepted_tree_record_kinds`. |
| A valid tree `subtree_root_ref` identifies no available current tree node | `node_not_found` | `subtree_root_ref`. |
| Configured current state cannot be constructed reliably | `configuration_failure` | None. |
| The operation cannot produce a reliable result under this contract | `execution_failure` | None. |

A request-wide error returns exactly one top-level `error`.
The response contains no normal result fields or top-level warnings.
Shared diagnostic objects, canonical codes, and context contracts belong to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| Repository, sequential-domain, and subtree scope availability | `spec:drmcp.design_records_mcp.current_record_model.current_record_scopes`. |
| Repository-wide unique addressability | `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability`. |
| Continued eligibility of nonconforming admitted records | `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission`. |
| Request fields, scope collection, validation order, response envelope, and request-wide outcomes | This Specification. |
| Matching semantics and target extraction | `spec:drmcp.design_records_mcp.operations.search.lexical_match_model`. |
| Match entries, snippets, ordering, and limits | `spec:drmcp.design_records_mcp.operations.search.search_results_and_limits`. |
| Shared diagnostics | `spec:drmcp.design_records_mcp.diagnostics`. |
| Detailed source defects and physical paths | Validation Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.current_record_model.current_record_scopes` | Supplies repository-wide, sequential-domain, and subtree scopes. |
| `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability` | Supplies uniquely addressable current records. |
| `spec:drmcp.design_records_mcp.operations.search.lexical_match_model` | Defines target source slices and lexical occurrence rules. |
| `spec:drmcp.design_records_mcp.operations.search.search_results_and_limits` | Defines the normal result projection. |
| `spec:drmcp.design_records_mcp.operations.discovery_and_listing.record_scope_discovery` | Discovers available sequential and tree scope selectors. |
| `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_envelope` | Defines top-level error and warning placement. |
| `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_catalog` | Defines codes and context fields used by this operation. |
