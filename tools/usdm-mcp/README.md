# Brewprint USDM MCP

MCP wrapper for the standalone MVP USDM tools in `tools/usdm/usdm_tools.py`.

## Tools

- `validate_usdm`
- `check_usdm_coverage`
- `usdm_covered_by`
- `collect_similar_requirements`

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
    "exclude_same_document": false
  }
}
```

The tool uses the standalone implementation in `tools/usdm/similarity` and
returns its response without classifying or rewriting requirements. Qdrant is
used only as a similarity index; repository USDM records remain the source of
truth.
