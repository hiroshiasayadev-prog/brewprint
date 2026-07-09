from __future__ import annotations

import os
import sys
from pathlib import Path
from typing import Any

from mcp.server.fastmcp import FastMCP


def _load_root() -> Path:
    raw_root = os.environ.get("USDM_MCP_ROOT")
    if not raw_root:
        raise RuntimeError("USDM_MCP_ROOT must point to the allowed repository root")

    root = Path(raw_root).expanduser().resolve(strict=True)
    if not root.is_dir():
        raise RuntimeError(f"USDM_MCP_ROOT is not a directory: {root}")
    return root


ROOT = _load_root()
USDM_TOOLS_DIR = ROOT / "tools" / "usdm"
if not USDM_TOOLS_DIR.is_dir():
    raise RuntimeError(f"USDM tools directory not found: {USDM_TOOLS_DIR}")
sys.path.insert(0, str(USDM_TOOLS_DIR))

import usdm_tools  # noqa: E402
from similarity.collector import (  # noqa: E402
    collect_similar_requirements as collect_similar_requirements_standalone,
)
from similarity.config import (  # noqa: E402
    DEFAULT_EMBEDDING_MODEL,
    DEFAULT_VECTOR_DIMENSIONS,
)
from similarity.embedding_generator import OllamaEmbeddingGenerator  # noqa: E402
from similarity.vector_index import RequirementVectorIndex  # noqa: E402


mcp = FastMCP("brewprint-usdm")

DEFAULT_QDRANT_URL = os.environ.get(
    "USDM_SIMILARITY_QDRANT_URL",
    "http://localhost:6333",
)
DEFAULT_OLLAMA_URL = os.environ.get(
    "USDM_SIMILARITY_OLLAMA_URL",
    "http://localhost:11434",
)
DEFAULT_COLLECTION = os.environ.get(
    "USDM_SIMILARITY_COLLECTION",
    "usdm_requirement_embeddings",
)
DEFAULT_MODEL = os.environ.get(
    "USDM_SIMILARITY_EMBEDDING_MODEL",
    DEFAULT_EMBEDDING_MODEL,
)


def _resolve_repo_root(raw_repo_root: str | None) -> Path:
    if not raw_repo_root:
        return ROOT

    candidate = Path(raw_repo_root).expanduser()
    if not candidate.is_absolute():
        candidate = ROOT / candidate

    resolved = candidate.resolve(strict=True)
    try:
        resolved.relative_to(ROOT)
    except ValueError as exc:
        raise ValueError(f"repo_root escapes USDM_MCP_ROOT: {raw_repo_root}") from exc

    if not resolved.is_dir():
        raise ValueError(f"repo_root is not a directory: {raw_repo_root}")

    return resolved


@mcp.tool()
def validate_usdm(
    repo_root: str | None = None,
    app_namespace: str | None = None,
) -> dict[str, Any]:
    """Validate MVP USDM records under a repository root.

    Args:
        repo_root: Repository root relative to or inside USDM_MCP_ROOT. Defaults to
            USDM_MCP_ROOT.
        app_namespace: Optional app namespace filter.
    """
    root = _resolve_repo_root(repo_root)
    return usdm_tools.validate_usdm(root, app_namespace)


@mcp.tool()
def check_usdm_coverage(
    repo_root: str | None = None,
    app_namespace: str | None = None,
    include_dangling: bool = True,
) -> dict[str, Any]:
    """Check USDM requirement coverage from implementation Specifications.

    Args:
        repo_root: Repository root relative to or inside USDM_MCP_ROOT. Defaults to
            USDM_MCP_ROOT.
        app_namespace: Optional app namespace filter.
        include_dangling: Include dangling usdm_covers entries in the response.
    """
    root = _resolve_repo_root(repo_root)
    return usdm_tools.check_usdm_coverage(root, app_namespace, include_dangling)


@mcp.tool()
def usdm_covered_by(
    requirement_id: str,
    repo_root: str | None = None,
) -> dict[str, Any]:
    """List Specification refs that cover one full USDM requirement ID.

    Args:
        requirement_id: Full USDM requirement ID such as usdm:drmcp.topic#R001.
        repo_root: Repository root relative to or inside USDM_MCP_ROOT. Defaults to
            USDM_MCP_ROOT.
    """
    root = _resolve_repo_root(repo_root)
    return usdm_tools.usdm_covered_by(root, requirement_id)


@mcp.tool()
def collect_similar_requirements(
    source_scope_ids: list[str],
    repo_root: str | None = None,
    candidate_scope_ids: list[str] | None = None,
    threshold: float = 0.86,
    max_candidates_per_requirement: int = 10,
    exclude_same_document: bool = False,
    include_empty_items: bool = False,
    include_details: bool = True,
    max_total_hits: int = 100,
    qdrant_url: str | None = None,
    ollama_url: str | None = None,
    embedding_model: str | None = None,
    collection: str = DEFAULT_COLLECTION,
) -> dict[str, Any]:
    """Collect semantic similarity candidates for USDM requirements.

    Args:
        source_scope_ids: USDM record, requirement, or supported topic IDs to
            use as source requirements.
        repo_root: Repository root relative to or inside USDM_MCP_ROOT. Defaults
            to USDM_MCP_ROOT.
        candidate_scope_ids: Optional candidate scope. Defaults to the source
            scope.
        threshold: Minimum cosine similarity score.
        max_candidates_per_requirement: Maximum candidates per source.
        exclude_same_document: Omit candidates from the source USDM record.
        include_empty_items: Include source items with no returned candidates.
        include_details: Include requirement detail text in source and candidate objects.
        max_total_hits: Maximum total candidate hits returned across all sources.
        qdrant_url: Qdrant base URL. Defaults to USDM_SIMILARITY_QDRANT_URL or
            http://localhost:6333.
        ollama_url: Ollama base URL. Defaults to USDM_SIMILARITY_OLLAMA_URL or
            http://localhost:11434.
        embedding_model: Ollama model. Defaults to
            USDM_SIMILARITY_EMBEDDING_MODEL or the standalone model default.
        collection: Qdrant collection. Defaults to
            USDM_SIMILARITY_COLLECTION or usdm_requirement_embeddings.
    """
    root = _resolve_repo_root(repo_root)
    effective_model = embedding_model or DEFAULT_MODEL
    embedding_generator = OllamaEmbeddingGenerator(
        base_url=ollama_url or DEFAULT_OLLAMA_URL,
        model=effective_model,
        dimensions=DEFAULT_VECTOR_DIMENSIONS,
    )
    vector_index = RequirementVectorIndex(
        qdrant_url=qdrant_url or DEFAULT_QDRANT_URL,
        collection=collection,
        embedding_generator=embedding_generator,
        model=effective_model,
        dimensions=DEFAULT_VECTOR_DIMENSIONS,
    )
    return collect_similar_requirements_standalone(
        repo_root=root,
        source_scope_ids=source_scope_ids,
        candidate_scope_ids=candidate_scope_ids,
        threshold=threshold,
        max_candidates_per_requirement=max_candidates_per_requirement,
        exclude_same_document=exclude_same_document,
        vector_index=vector_index,
        include_empty_items=include_empty_items,
        include_details=include_details,
        max_total_hits=max_total_hits,
    )


def main() -> None:
    mcp.run()


if __name__ == "__main__":
    main()
