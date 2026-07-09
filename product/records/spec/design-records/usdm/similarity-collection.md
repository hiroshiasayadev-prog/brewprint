# Contract: USDM requirement similarity collection

- **id**: `spec:product.design_records.usdm.similarity_collection`
- **status**: draft
- **date**: 2026-07-09
- **parent**: `spec:product.design_records.usdm`
- **contract_class**: `interface`

## What this is

This contract defines a USDM-facing operation that collects semantically similar requirement candidates.

The operation supports human review by grouping similar requirement details. The operation does not judge whether candidates are duplicates.

## Non-goals

- Do not classify candidates as duplicate, overlap, intentional repetition, or distinct.
- Do not merge, suppress, rewrite, or delete requirements.
- Do not replace USDM coverage checks.
- Do not make the vector index a canonical source for requirements.
- Do not require GPU execution as part of the product contract.

## Boundary

The public boundary is a USDM similarity collection operation.

The internal embedding and vector search implementation is independent auxiliary tooling. USDM supplies the requirement IDs and requirement details at the operation boundary.

| component | responsibility |
|---|---|
| Similarity Request Orchestrator | Accept the USDM-facing request, expand source and candidate scopes, collect requirement details, and format the response. |
| Requirement Vector Index | Maintain requirement embedding freshness, call the embedding generator for missing or stale vectors, upsert vectors, and run similarity search. |
| Embedding Generator | Accept text batches and return embedding vectors from the configured Ollama embedding backend. |
| Qdrant vector store | Store requirement embeddings and payloads, and execute cosine similarity search. |

The Embedding Generator receives text batches. It does not receive requirement IDs, document IDs, or USDM scope IDs.

The Requirement Vector Index does not resolve USDM scopes. It receives already-expanded requirement records from the Similarity Request Orchestrator.

## Request

### Operation

| operation | purpose |
|---|---|
| `collect_similar_requirements` | Return semantic similarity candidates for USDM requirements. |

### Request fields

| field | required | default | meaning |
|---|---:|---|---|
| `repo_root` | yes | - | Repository root to inspect. |
| `source_scope_ids` | yes | - | USDM record IDs, USDM requirement IDs, or supported USDM topic IDs to use as source requirements. |
| `candidate_scope_ids` | no | `source_scope_ids` | Scope used to collect candidate requirements. |
| `threshold` | no | `0.86` | Minimum cosine similarity score for returned candidates. |
| `max_candidates_per_requirement` | no | `10` | Maximum number of candidates returned for one source requirement. |
| `exclude_same_document` | no | `false` | When true, candidates from the same USDM record are omitted. |

The operation always excludes the source requirement itself from its candidate list.

When `candidate_scope_ids` is omitted, the candidate set is the expanded `source_scope_ids` set.

## Response

### Response fields

| field | meaning |
|---|---|
| `ok` | False when the operation cannot complete. |
| `source_requirements` | Number of expanded source requirements. |
| `candidate_requirements` | Number of expanded candidate requirements. |
| `threshold` | Effective cosine similarity threshold. |
| `model` | Effective embedding model identity. |
| `dimensions` | Effective embedding vector dimensions. |
| `items` | Source-centric similarity candidate groups. |
| `diagnostics` | Errors or warnings produced during scope expansion, embedding synchronization, or search. |

### Item fields

| field | meaning |
|---|---|
| `source` | Source requirement identity and detail text. |
| `candidates` | Candidate requirements whose score meets or exceeds the effective threshold. |

### Requirement object fields

| field | meaning |
|---|---|
| `requirement_id` | Full USDM requirement ID. |
| `detail` | Requirement detail text used for embedding and review output. |
| `usdm_id` | USDM record ID that owns the requirement. |
| `path` | Repository-relative source path when available. |

### Candidate object fields

| field | meaning |
|---|---|
| `requirement_id` | Full USDM requirement ID for the candidate. |
| `detail` | Candidate requirement detail text. |
| `usdm_id` | USDM record ID that owns the candidate. |
| `path` | Repository-relative source path when available. |
| `score` | Cosine similarity score returned by the vector search. |

The response is source-centric. A pair may appear twice when both requirements are in the source set.

The operation returns only candidates whose score meets or exceeds `threshold`.

## Embedding and vector index contract

| field | value |
|---|---|
| embedding backend | Ollama |
| embedding model | `snowflake-arctic-embed2` |
| vector dimensions | `1024` |
| vector store | Qdrant |
| distance | cosine |

The Requirement Vector Index stores the normalized requirement detail hash with each vector payload.

A stored vector is stale when any stored value differs from the configured value:

| stored value | configured value |
|---|---|
| normalized requirement detail hash | current normalized requirement detail hash |
| embedding model identity | configured embedding model identity |
| vector dimensions | configured vector dimensions |

The Requirement Vector Index regenerates missing or stale embeddings before running similarity search.

The vector store is a search index. The vector store is not the source of truth for requirement content.

## Normalization

The normalized requirement detail text is used only for freshness detection and embedding input.

| input content | handling |
|---|---|
| Markdown table framing | Excluded from detail text. |
| Requirement row ID | Excluded from detail text. |
| Requirement detail text | Included. |
| Code spans and domain identifiers inside the detail | Preserved. |
| Repeated whitespace | Collapsed. |
| Leading and trailing whitespace | Trimmed. |

## Errors

| condition | handling |
|---|---|
| `repo_root` is missing or unreadable | Return `ok: false` and an error diagnostic. |
| `source_scope_ids` is empty | Return `ok: false` and an error diagnostic. |
| A source scope ID cannot be resolved | Return `ok: false` and an error diagnostic. |
| A candidate scope ID cannot be resolved | Return `ok: false` and an error diagnostic. |
| Ollama embedding generation fails | Return `ok: false` and an error diagnostic. |
| Qdrant collection creation or search fails | Return `ok: false` and an error diagnostic. |
| Existing Qdrant collection dimensions or distance do not match the configured values | Return `ok: false` and an error diagnostic. |

Fallback brute-force similarity search is outside the MVP contract.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.usdm` | Parent overview. |
| `spec:product.design_records.usdm.artifact_format` | Defines USDM record and requirement ID format. |
| `spec:product.design_records.usdm.coverage_tools` | Defines existing standalone MVP USDM tools. |
| `spec:product.design_records.usdm.coverage_format` | Defines `usdm_covers` metadata and coverage semantics. |
