# Reference: Record source

- **id**: `spec:drmcp.design_records_mcp.schema.record_source`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines the Markdown source types that Design Records MCP reads, and the metadata source per record kind.

## Current contract

### Record sources

Design Records MCP MVP builds design records from Markdown files. The following sources are read in MVP.

| source | purpose |
|---|---|
| ADR bullet metadata | `status` / `date` / `depends_on` / `supersedes` / `migrated_to_spec` for decision records |
| Investigation bullet metadata | `status` / `date` / `trigger` / `scope` / `non_scope` / `source_refs` / `follow_up_candidates` / optional related metadata / `follow_up_results` for investigation records |
| Requirement bullet metadata | `id` / `status` / `date` / `source_refs` / `work_items` / optional `subdomain` for requirement records |
| Work item bullet metadata | `id` / `status` / `date` / `source_requirement` / `impact_refs` / `tasks` / optional `subdomain` for work item records |
| Task bullet metadata | `id` / `status` / `date` / `work_item` / `source_requirement` / `estimate` / `depends_on` / `outputs` / optional `subdomain` for task records |
| Spec YAML front matter | `scope` / `status` / `design_record` metadata for spec records. Top-level `depends_on` is read as a doc-policy origin path list; it is not used as record dependency. |
| Markdown H1 | `title` extraction |
| File path | Record path / filename ID validation |
| Markdown headings | `headings` field in `get_record` / `get_records` found record response |
| Markdown body | Raw body for `get_record(include_body=true)` / `get_records(include_body=true)` |
| Authoring guide Markdown | Guide ID / title / abstract / content from `docs/guides/*.md` for `list_authoring_guides` / `get_authoring_guidance` |

MVP does not infer dependency or migration state from Markdown body natural language.

> Source: V01-ADR-076 §front matter 方針, V01-ADR-077 §get_record の責務

### Metadata source per record kind

| kind | metadata source |
|---|---|
| `decision` | Existing ADR format maintained; H1-adjacent bullet metadata block read |
| `spec` | Existing YAML front matter read. `design_record` metadata block may be added. MVP does not adopt `design_record` front matter for ADR records — decision records are built from existing ADR H1-adjacent bullet metadata only, without adding YAML front matter or a `design_record` block to ADRs. |
| `investigation` | H1-adjacent bullet metadata per V01-ADR-086 |
| `requirement` | H1-adjacent bullet metadata |
| `work_item` | H1-adjacent bullet metadata |
| `task` | H1-adjacent bullet metadata |

For the exact bullet metadata grammar, see `spec:drmcp.design_records_mcp.schema.metadata_grammar`.

> Source: V01-ADR-076 §front matter 方針, V01-ADR-077 §get_record の責務
