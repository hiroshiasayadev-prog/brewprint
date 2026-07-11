# Brewprint USDM MCP

MCP wrapper for the standalone MVP USDM tools in `tools/usdm/usdm_tools.py`.

## Tools

- `validate_usdm`
- `check_usdm_coverage`
- `check_usdm_scope_coverage`
- `usdm_covered_by`
- `collect_similar_requirements`
- `search_requirements`

## Requirement section format

USDM requirement sections use a human-readable title and an immediate source field:

```markdown
## Requirements: Application discovery behavior
> source: spec:product.design_records.repository_layout.record_discovery_paths

| id | requirement |
|---|---|
| R001 | ... |
```

Use `> source: literal` when the rows are direct upstream requirements and no corresponding Specification is their source. Titles must be non-empty, unique within one USDM record, and must not be `literal` or a canonical `spec:` ref.

`validate_usdm`, scoped coverage, similarity collection, and search use this section shape.

## Start

```powershell
$env:USDM_SIMILARITY_QDRANT_URL = "http://192.168.11.22:6333"
$env:USDM_SIMILARITY_OLLAMA_URL = "http://192.168.11.22:11434"
./tools/usdm-mcp/start.ps1 -Root C:\Users\imved\projects\brewprint -Port 8184
```

Optional similarity configuration:

- `USDM_SIMILARITY_QDRANT_URL` defaults to `http://localhost:6333`.
- `USDM_SIMILARITY_OLLAMA_URL` defaults to `http://localhost:11434`.
- `USDM_SIMILARITY_COLLECTION` defaults to
  `usdm_requirement_embeddings`.
- `USDM_SIMILARITY_EMBEDDING_MODEL` defaults to the model configured in
  `tools/usdm/similarity/config.py`.

Equivalent `mcp-proxy` shape:

```bat
start "usdm-mcp" cmd /k "mcp-proxy --host=127.0.0.1 --port=8184 --cwd=C:\Users\imved\projects\brewprint --env USDM_MCP_ROOT C:\Users\imved\projects\brewprint -- uv run --project C:\Users\imved\projects\brewprint\tools\usdm-mcp python C:\Users\imved\projects\brewprint\tools\usdm-mcp\server.py"
```

The server restricts `repo_root` arguments to paths inside `USDM_MCP_ROOT`.

## Scoped coverage request

Example MCP `tools/call` payload:

```json
{
  "name": "check_usdm_scope_coverage",
  "arguments": {
    "scope_ids": [
      "usdm:product.design_records.namespace_and_identity"
    ],
    "include_covered": true,
    "include_not_covered": true,
    "include_empty_records": false
  }
}
```

The tool expands USDM app, topic, record, or full requirement IDs and returns a
compact record-grouped coverage report. Covered rows map compact row IDs such
as `#R001` to covering Specification refs. Uncovered rows are listed as compact
row IDs only. Requirement text, full requirement detail objects, paths, and
internal `usdm_id` fields are omitted from normal items.

## Similarity request

Example MCP `tools/call` payload:

```json
{
  "name": "collect_similar_requirements",
  "arguments": {
    "source_scope_ids": [
      "usdm:product.design_records.spec_document_format.spec_id_as_ref"
    ],
    "threshold": 0.86,
    "max_candidates_per_requirement": 10,
    "max_total_hits": 100,
    "include_empty_items": false,
    "include_details": true,
    "exclude_same_document": false
  }
}
```

The tool uses the standalone implementation in `tools/usdm/similarity` and
returns its response without classifying or rewriting requirements. By default,
it omits source items with no returned candidates, includes requirement detail
text, and caps the response to 100 total candidate hits. Qdrant is used only as
a similarity index; repository USDM records remain the source of truth.

## Search request

Example MCP `tools/call` payload:

```json
{
  "name": "search_requirements",
  "arguments": {
    "query": "path derived canonical spec ref",
    "candidate_scope_ids": [
      "usdm:product.design_records"
    ],
    "threshold": 0.30,
    "max_results": 20,
    "include_details": true
  }
}
```

The tool searches normalized USDM requirement detail text only. It synchronizes
candidate embeddings before embedding the query, then searches Qdrant within
the expanded candidate set. It does not classify, merge, rewrite, suppress, or
delete requirements, and public result objects omit internal `usdm_id` values.
