# Reference: Metadata grammar

- **id**: `spec:drmcp.design_records_mcp.schema.metadata_grammar`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines how DRMCP parses H1-adjacent visible metadata for current design records.

## Current contract

### Shared placement rule

Current records use visible metadata immediately after the real ATX H1.

- DRMCP counts H1 headings outside fenced code blocks.
- Zero or more blank lines may appear between H1 and the first metadata marker.
- The first non-blank line after H1 must be a recognized metadata marker for the record kind.
- DRMCP does not search later body content for a replacement metadata block.
- Body headings and prose are not metadata.
- YAML front matter is not a current metadata source.

The visible marker form is:

```text
- **<key>**: <value>
```

- The line begins at column 1 with `- `.
- The key is enclosed by `**` and is case-sensitive.
- An ASCII colon immediately follows the closing `**`.
- Leading and trailing whitespace in a scalar value is trimmed.
- Alternate bullet markers, unbolded keys, case repair, and malformed markers are not normalized.

### Scalar normalization

A scalar value may be written as plain text or as one complete inline-code span.

```text
- **status**: accepted
- **parent**: `spec:exampleapp.search`
```

When the complete trimmed value is enclosed by one matching backtick pair, DRMCP removes only that pair. Partial wrapping, multiple spans, and unmatched backticks are invalid scalar forms.

Dates use strict `YYYY-MM-DD` format when the owning PRODUCT authority defines a date field.

### List normalization

A declared list field may use either:

```text
- **field**: VALUE-1, VALUE-2
```

or:

```text
- **field**:
  - VALUE-1
  - VALUE-2
```

- Inline list items are comma-separated and individually trimmed.
- Child list items must be directly indented below their parent marker.
- An explicitly empty list field normalizes to an empty list when the owning PRODUCT rule allows an empty list.
- Empty child items are invalid.
- Scalar fields do not accept child list items.

### Duplicate and unknown markers

A recognized field may appear at most once in one metadata block. Duplicate occurrences do not use first-wins or last-wins behavior.

Only fields declared for the record kind are mapped into the current record model. Unknown-field validity is kind-specific; this shared rule does not introduce one rejection policy for every sequential kind. Exact diagnostics and severity are owned by `DRMCP-WORK-MCP-006`.

### Complete canonical artifact IDs

Sequential artifact identity and relation values use complete canonical app-aware artifact IDs.

```text
ADR / investigation / requirement / work item:
<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<SEQUENCE>

Task:
<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>
```

DRMCP validates the complete value directly. It does not strip or reattach a runtime namespace prefix.

Bare fragments such as `REQ-*`, `WORK-*`, or `TASK-*` may be used in explanatory grammar text, but they are not canonical current metadata values.

### ADR metadata

ADR metadata contains no `id` field. Identity comes from the complete canonical ID in H1.

| field | form | required |
|---|---|---:|
| `status` | scalar | yes |
| `date` | scalar date | yes |
| `depends_on` | list of complete canonical artifact IDs | yes; empty allowed |
| `supersedes` | list of complete canonical ADR IDs | yes; empty allowed |
| `migrated_to_spec` | scalar date or empty | yes; empty normalizes to `null` |

### Investigation metadata

Investigation metadata contains no `id` field. Identity comes from the complete canonical ID in H1.

Required fields:

| field | form |
|---|---|
| `status` | scalar |
| `date` | scalar date |
| `trigger` | scalar |
| `scope` | scalar |
| `non_scope` | scalar |
| `source_refs` | list |
| `follow_up_candidates` | list |

Optional fields:

| field | form |
|---|---|
| `supersedes` | list of complete canonical investigation IDs |
| `related_requirements` | list of complete canonical requirement IDs |
| `related_work_items` | list of complete canonical work-item IDs |
| `related_adrs` | list of complete canonical ADR IDs |
| `related_specs` | list of active canonical `spec:` refs |
| `related_internal_design` | auxiliary scalar or list as defined by the owning PRODUCT rule |
| `related_coverage` | auxiliary scalar or list as defined by the owning PRODUCT rule |
| `follow_up_results` | list of canonical artifact IDs or active canonical `spec:` refs |

Investigation metadata does not use task IDs in `source_refs`, `follow_up_candidates`, or `follow_up_results` under current PRODUCT authoring authority.

### Requirement metadata

| field | form | required |
|---|---|---:|
| `id` | complete canonical requirement ID | yes |
| `status` | scalar | yes |
| `date` | scalar date | yes |
| `source_refs` | list of canonical artifact IDs or active canonical `spec:` refs | yes; empty allowed |
| `work_items` | list of complete canonical work-item IDs | yes; empty allowed |
| `subdomain` | scalar | no |

The metadata `id`, H1 ID, and filename ID prefix must agree exactly for a valid requirement. H1 remains the canonical identity authority; missing, malformed, or mismatched metadata `id` does not remove addressability when H1 contains one valid unique canonical ID.

### Work-item metadata

| field | form | required |
|---|---|---:|
| `id` | complete canonical work-item ID | yes |
| `status` | scalar | yes |
| `date` | scalar date | yes |
| `source_requirement` | complete canonical requirement ID | yes |
| `impact_refs` | list of canonical artifact IDs or active canonical `spec:` refs | yes; empty allowed |
| `tasks` | list of complete canonical task IDs | yes; empty allowed |
| `subdomain` | scalar | no |

The metadata `id`, H1 ID, and filename ID prefix must agree exactly for a valid work item. H1 remains the canonical identity authority; missing, malformed, or mismatched metadata `id` does not remove addressability when H1 contains one valid unique canonical ID.

### Task metadata

| field | form | required |
|---|---|---:|
| `id` | complete canonical task ID | yes |
| `status` | scalar | yes |
| `date` | scalar date | yes |
| `work_item` | complete canonical work-item ID | yes |
| `source_requirement` | complete canonical requirement ID | yes |
| `estimate` | scalar | yes |
| `depends_on` | list of complete canonical task IDs | yes; empty allowed |
| `outputs` | list of canonical artifact IDs, active canonical `spec:` refs, or declared output identifiers | yes; empty allowed |
| `subdomain` | scalar | no |

The metadata `id`, H1 ID, and filename ID prefix must agree exactly for a valid task. H1 remains the canonical identity authority; missing, malformed, or mismatched metadata `id` does not remove addressability when H1 contains one valid unique canonical ID.
The task ID domain and work-sequence segments must agree with the parent Work Item ID.

### Current spec metadata

Current spec metadata is a contiguous scalar-only marker block after H1.

- The first blank or non-marker line after the block starts ends the block.
- Indented child-list items are not accepted.
- DRMCP does not resume metadata parsing later in the file.
- YAML front matter is invalid and is never used as fallback metadata.

Recognized fields:

| field | required | authority |
|---|---:|---|
| `id` | yes | Path-derived canonical `spec:` ref consistency value. |
| `status` | yes | Required non-empty scalar; PRODUCT does not currently define a complete vocabulary. |
| `date` | yes | Scalar date under PRODUCT spec-format authority. |
| `parent` | yes | `root`, `-`, or active canonical parent `spec:` ref. |
| `contract_class` | only for `Contract` | `interface` or `format`; prohibited on non-`Contract` specs. |

The path-derived canonical spec ref is authoritative. A mismatched metadata `id` is invalid but does not replace the path-derived identity or become an alias.

Any current spec metadata marker outside the recognized field set is a current source-format violation.

- DRMCP does not ignore an unknown current spec marker.
- DRMCP does not retain it as extension metadata.
- DRMCP does not repair it into another field name.
- The uniquely path-addressable source remains addressable and fails validation.

Current specs do not define metadata fields named `depends_on`, `supersedes`, or `migrated_to_spec`. Those names are therefore unknown current spec markers and source-format violations. DRMCP does not read them from YAML or any obsolete nested metadata object.

### Missing and invalid metadata

A valid current record contains every field required by its PRODUCT authority.

When one canonical ID is still determinable, missing or invalid metadata does not remove addressability. DRMCP retains values it parsed and does not invent replacements.

When no canonical ID is determinable, the source remains validation-only and must retain its source path for repair diagnostics. Duplicate canonical identity creates no winner.

Shared behavior is defined by:

- `spec:drmcp.design_records_mcp.schema.record_model`;
- `spec:drmcp.design_records_mcp.schema.discovery`;
- `spec:drmcp.design_records_mcp.schema.id_normalization`.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Current sequential artifact ID grammar. |
| `spec:product.design_records.spec_format.document_shape` | Current spec H1 and metadata requiredness. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Current spec identity and parent grammar. |
| `spec:drmcp.design_records_mcp.schema.fields` | Parsed common and kind-specific field vocabulary. |
| `spec:drmcp.design_records_mcp.schema.id_normalization` | Canonical identity mapping. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Invalid-source retention and active-index behavior. |

## Sources

- `DRMCP-TASK-MCP-003-03`: Current spec metadata decisions.
- `DRMCP-TASK-MCP-003-04`: Shared field and current identity decisions.
- PRODUCT authoring standards for ADR, investigation, requirement, work item, task, and spec records.
