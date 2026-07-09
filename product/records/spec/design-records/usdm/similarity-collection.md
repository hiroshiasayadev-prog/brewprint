# Contract: USDM requirement similarity collection

- **id**: `spec:product.design_records.usdm.similarity_collection`
- **status**: draft
- **date**: 2026-07-09
- **parent**: `spec:product.design_records.usdm`
- **contract_class**: `interface`

## What this is

This contract defines USDM-facing operations that use requirement-detail embeddings for review support.

The similarity operation supports human review by grouping similar requirement details. The search operation supports human lookup by searching requirement detail text with a free-text query. Neither operation judges whether candidates are duplicates.

## Non-goals

- Do not classify candidates as duplicate, overlap, intentional repetition, or distinct.
- Do not merge, suppress, rewrite, or delete requirements.
- Do not replace USDM coverage checks.
- Do not make the vector index a canonical source for requirements.
- Do not require GPU execution as part of the product contract.

## Boundary

The public boundary is a USDM requirement embedding operation set.

The internal embedding and vector search implementation is independent auxiliary tooling. USDM supplies the requirement IDs and requirement details at the operation boundary.

| component | responsibility |
|---|---|
| Similarity Request Orchestrator | Accept the USDM-facing request, expand source and candidate scopes, collect requirement details, and format the response. |
| Requirement Search Orchestrator | Accept a free-text query, expand candidate scopes, synchronize candidate embeddings, execute query search, and format the response. |
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
| `search_requirements` | Return requirement detail matches for a free-text query. |

### `collect_similar_requirements` request fields

| field | required | default | meaning |
|---|---:|---|---|
| `repo_root` | yes | - | Repository root to inspect. |
| `source_scope_ids` | yes | - | USDM record IDs, USDM requirement IDs, or supported USDM topic IDs to use as source requirements. |
| `candidate_scope_ids` | no | `source_scope_ids` | Scope used to collect candidate requirements. |
| `threshold` | no | `0.86` | Minimum cosine similarity score for returned candidates. |
| `max_candidates_per_requirement` | no | `10` | Maximum number of candidates considered for one source requirement before response-wide limiting. |
| `exclude_same_document` | no | `false` | When true, candidates from the same USDM record are omitted. |
| `include_empty_items` | no | `false` | When true, include source items with no returned candidates. |
| `include_details` | no | `true` | When true, include requirement detail text in source and candidate objects. |
| `max_total_hits` | no | `100` | Maximum candidate hits returned across all source items. |

The operation always excludes the source requirement itself from its candidate list.

When `candidate_scope_ids` is omitted, the candidate set is the expanded `source_scope_ids` set.

### `search_requirements` request fields

| field | required | default | meaning |
|---|---:|---|---|
| `repo_root` | yes | - | Repository root to inspect. |
| `query` | yes | - | Free-text search query embedded and compared against requirement details. |
| `candidate_scope_ids` | yes | - | USDM record IDs, USDM requirement IDs, or supported USDM topic IDs to search. |
| `threshold` | no | `0.30` | Minimum cosine similarity score for returned results. |
| `max_results` | no | `20` | Maximum number of results returned for the query. |
| `include_details` | no | `true` | When true, include requirement detail text in result objects. |

The search operation searches only normalized requirement detail text. It does not search notes, prose sections, headings, front matter, file names, or paths.

The search operation MUST synchronize candidate requirement embeddings before embedding and searching the query.

## Response

### `collect_similar_requirements` response fields

| field | meaning |
|---|---|
| `ok` | False when the operation cannot complete. |
| `source_requirements` | Number of expanded source requirements. |
| `candidate_requirements` | Number of expanded candidate requirements. |
| `threshold` | Effective cosine similarity threshold. |
| `model` | Effective embedding model identity. |
| `dimensions` | Effective embedding vector dimensions. |
| `returned_hits` | Number of candidate hits returned after response-wide limiting. |
| `items` | Source-centric similarity candidate groups. |
| `diagnostics` | Errors or warnings produced during scope expansion, embedding synchronization, or search. |

### Item fields

| field | meaning |
|---|---|
| `source` | Source requirement identity and optional detail text. |
| `candidates` | Candidate requirements whose score meets or exceeds the effective threshold and response-wide hit limit. |

### Requirement object fields

| field | meaning |
|---|---|
| `requirement_id` | Full USDM requirement ID. |
| `detail` | Requirement detail text used for embedding and review output. Omitted when `include_details` is false. |
| `path` | Repository-relative source path when available. |

### Candidate object fields

| field | meaning |
|---|---|
| `requirement_id` | Full USDM requirement ID for the candidate. |
| `detail` | Candidate requirement detail text. Omitted when `include_details` is false. |
| `path` | Repository-relative source path when available. |
| `score` | Cosine similarity score returned by the vector search. |

The response is source-centric. A pair may appear twice when both requirements are in the source set.

When `include_empty_items` is false, the response omits source items that have no returned candidates.

When `max_total_hits` is provided, the response keeps the highest-scoring candidate hits across all source items until the limit is reached.

The operation returns only candidates whose score meets or exceeds `threshold`.

### `search_requirements` response fields

| field | meaning |
|---|---|
| `ok` | False when the operation cannot complete. |
| `candidate_requirements` | Number of expanded candidate requirements. |
| `query` | Effective free-text query. |
| `threshold` | Effective cosine similarity threshold. |
| `model` | Effective embedding model identity. |
| `dimensions` | Effective embedding vector dimensions. |
| `results` | Requirement matches whose score meets or exceeds the threshold. |
| `diagnostics` | Errors or warnings produced during scope expansion, embedding synchronization, query embedding, or search. |

### Search result object fields

| field | meaning |
|---|---|
| `requirement_id` | Full USDM requirement ID for the result. |
| `detail` | Requirement detail text. Omitted when `include_details` is false. |
| `path` | Repository-relative source path when available. |
| `score` | Cosine similarity score returned by vector search. |

The search response is query-centric and sorted by descending score.

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

The Requirement Vector Index regenerates missing or stale embeddings before running similarity or free-text requirement search.

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
| `source_scope_ids` is empty for similarity collection | Return `ok: false` and an error diagnostic. |
| `query` is empty for requirement search | Return `ok: false` and an error diagnostic. |
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
