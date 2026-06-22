# Reference: Record model

- **id**: `spec:drmcp.design_records_mcp.schema.record_model`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines the internal record model and bootstrap record set for Design Records MCP.

## Current contract

### Internal record model

The internal record model holds at least the following information:

| field | source | meaning |
|---|---|---|
| `id` | decision/investigation: H1; spec: `design_record.id` | Record ID |
| `kind` | source kind; spec: `design_record.kind` | Record kind |
| `title` | H1 | Human-readable title |
| `status` | decision/investigation: bullet metadata; spec: top-level front matter | Record status |
| `path` | Filesystem | Markdown file path |
| `decision` | ADR bullet metadata | Decision-specific detail object |
| `spec` | spec `design_record` metadata | Spec-specific detail object |
| `investigation` | Investigation bullet metadata | Investigation-specific detail object |
| `requirement` | Requirement bullet metadata | Requirement-specific detail object |
| `work_item` | Work item bullet metadata | Work-item-specific detail object |
| `task` | Task bullet metadata | Task-specific detail object |
| `headings` | Markdown parse | Heading list for `get_record` |
| `body` | Markdown file | Raw body; only fetched when requested by `get_record` or `get_records` |

### get_records response item

The `get_records` response item wrapper is not the record model itself. For a `retrieval_status: "found"` item, `record` returns the record model above. For a `retrieval_status: "not_found"` item, `record` is `null`. The retrieval-state field `retrieval_status` must not be conflated with the record lifecycle `status`.

### Heading and body rules

`headings` covers ATX headings only. Lines beginning with `#` inside YAML front matter and fenced code blocks are not treated as headings. Setext headings are not handled in MVP.

`body` is included only in the found record response of `get_record(include_body=true)` or `get_records(include_body=true)`. Body is returned verbatim without formatting, summarization, normalization, or truncation.

> Source: V01-ADR-077 §list_records の責務, V01-ADR-077 §get_record の責務

### Bootstrap records

A small set of representative records is designated as bootstrap targets for MVP validation. ADR records use existing bullet metadata without adding YAML front matter. Spec records use `design_record` metadata in YAML front matter.

Initial bootstrap candidates:

- V01-ADR-050
- V01-ADR-067 through V01-ADR-077
- V01-ADR-086 through V01-ADR-088
- `docs/investigations/docs/INV-DOCS-001-investigation-artifact-format-and-lifecycle.md`
- `docs/spec/design-records-mcp/**`

Records outside this set are not bulk-tagged. Tags are added incrementally when touched during consistency reviews, new ADR authoring, or related spec updates.

> Source: V01-ADR-076 §bootstrap方針
