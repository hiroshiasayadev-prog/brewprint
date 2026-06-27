# Reference: Record source

- **id**: `spec:drmcp.design_records_mcp.schema.record_source`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines the Markdown source types that Design Records MCP reads, and the metadata source per record kind.

## Current contract

### Record sources

Design Records MCP builds design records from Markdown files. The following sources are read.

| source | purpose |
|---|---|
| ADR bullet metadata | `status` / `date` / `depends_on` / `supersedes` / `migrated_to_spec` for decision records |
| Investigation bullet metadata | `status` / `date` / `trigger` / `scope` / `non_scope` / `source_refs` / `follow_up_candidates` / optional related metadata / `follow_up_results` for investigation records |
| Requirement bullet metadata | `id` / `status` / `date` / `source_refs` / `work_items` / optional `subdomain` for requirement records |
| Work item bullet metadata | `id` / `status` / `date` / `source_requirement` / `impact_refs` / `tasks` / optional `subdomain` for work item records |
| Task bullet metadata | `id` / `status` / `date` / `work_item` / `source_requirement` / `estimate` / `depends_on` / `outputs` / optional `subdomain` for task records |
| Current spec H1-adjacent metadata | `id` / `status` / `date` / `parent` / kind-conditional `contract_class` for current spec records (`Contract` specs: required; non-`Contract` specs: prohibited; applicability owned by `spec:product.design_records.spec_format.document_shape`) |
| Markdown H1 | `title` extraction and `kind` extraction for current specs |
| File path | Repository placement for canonical spec ref derivation; record path / filename ID validation for sequential kinds |
| Markdown headings | Readable source material available for operation-specific projections. |
| Markdown body | Readable source material available for operation-specific projections. |
| Authoring guide Markdown | Guide ID / title / abstract / content from `docs/guides/*.md` for `list_authoring_guides` / `get_authoring_guidance` |

DRMCP does not infer dependency or migration state from Markdown body natural language.

This specification defines source availability only. Public heading inclusion, optional body inclusion, tool selection, and response representation are owned by `DRMCP-WORK-MCP-004`.

### Current spec metadata source

YAML front matter is not a current spec metadata source. A current spec file that contains YAML front matter is an invalid current source. DRMCP does not read, ignore, or use YAML values as a fallback or supplement.

Current spec metadata is sourced entirely from the H1-adjacent visible metadata block, following the grammar defined by `spec:drmcp.design_records_mcp.schema.metadata_grammar`.

Current spec canonical identity is derived from the configured `app_namespace` and repository placement (file path under `<records_root>/spec/`), as defined by `spec:product.design_records.spec_format.spec_id_as_ref` and mapped by `spec:drmcp.design_records_mcp.schema.id_normalization`. YAML, `SPEC-*`, and `V01-SPEC-*` values do not participate in identity derivation.

### Metadata source per record kind

| kind | metadata source |
|---|---|
| `decision` | H1-adjacent bullet metadata block |
| `spec` | H1-adjacent visible metadata block (see current spec metadata source above) |
| `investigation` | H1-adjacent bullet metadata |
| `requirement` | H1-adjacent bullet metadata |
| `work_item` | H1-adjacent bullet metadata |
| `task` | H1-adjacent bullet metadata |

For the exact bullet metadata grammar, see `spec:drmcp.design_records_mcp.schema.metadata_grammar`.

### Current spec source mapping

| source element | content mapped |
|---|---|
| Repository placement (`<records_root>/spec/**/*.md`) | Path-derived canonical `spec:` ref |
| Real ATX H1 (`# <SpecKind>: <Title>`) | Spec kind and human-readable title |
| H1-adjacent metadata block | `id` (consistency value), `status`, `date`, `parent`, `contract_class` (required for `Contract` specs; prohibited for non-`Contract` specs) |
| Markdown headings | Readable spec source material. |
| Markdown body | Readable spec source material. |

The path-derived canonical ref is the authoritative current spec identity. The H1-adjacent `id` is a required consistency value that must exactly match the path-derived canonical ref. It is not an independent identity authority.

This mapping does not define whether headings or body content appear in a public read response. `DRMCP-WORK-MCP-004` owns that contract.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.spec_format.document_shape` | Authority for current spec H1 format and H1-adjacent metadata requirements. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Authority for path-derived canonical spec ref derivation and H1-adjacent `id` as consistency value. |
| `spec:drmcp.design_records_mcp.schema.discovery` | Defines current spec candidate paths, YAML rejection, and invalid source behavior. |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Defines the H1-adjacent metadata block grammar for all record kinds. |
| `spec:drmcp.design_records_mcp.schema.id_normalization` | Defines path-derived canonical ref derivation steps for current specs. |

## Sources

- `DRMCP-TASK-MCP-003-03`: Accepted decisions 7, 8, 10, 12, 13, 15.
- `spec:product.design_records.spec_format.document_shape`: H1 and metadata requirements authority.
- `spec:product.design_records.spec_format.spec_id_as_ref`: Path-derived identity authority.
