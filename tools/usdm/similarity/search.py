"""Query-centric USDM requirement semantic search orchestration."""

from __future__ import annotations

from pathlib import Path
from typing import Any, Protocol

from .models import RequirementRow, SimilarityHit
from .text_normalizer import normalize_requirement_detail
from .usdm_loader import diagnostic, expand_scopes


class VectorIndex(Protocol):
    model: str
    dimensions: int

    def search_requirements(
        self,
        query: str,
        candidates: list[RequirementRow],
        threshold: float,
        max_results: int,
    ) -> list[SimilarityHit]: ...


def search_requirements(
    repo_root: Path,
    query: str,
    candidate_scope_ids: list[str],
    threshold: float,
    max_results: int,
    vector_index: VectorIndex,
    include_details: bool = True,
) -> dict[str, Any]:
    diagnostics: list[dict[str, Any]] = []
    effective_query = normalize_requirement_detail(query)

    if not repo_root.is_dir():
        diagnostics.append(
            diagnostic(
                "repo_root",
                "repo_root is missing or unreadable.",
                str(repo_root),
            )
        )
    if not effective_query:
        diagnostics.append(
            diagnostic(
                "query",
                "query must contain non-whitespace text.",
            )
        )
    if not candidate_scope_ids:
        diagnostics.append(
            diagnostic(
                "candidate_scope_ids",
                "candidate_scope_ids must contain at least one USDM scope ID.",
            )
        )
    if not -1.0 <= threshold <= 1.0:
        diagnostics.append(
            diagnostic(
                "threshold",
                "threshold must be between -1.0 and 1.0.",
                threshold,
            )
        )
    if max_results < 1:
        diagnostics.append(
            diagnostic(
                "max_results",
                "max_results must be at least 1.",
                max_results,
            )
        )
    if diagnostics:
        return _failure_response(
            effective_query,
            threshold,
            vector_index.model,
            vector_index.dimensions,
            diagnostics,
        )

    candidate_expansion = expand_scopes(repo_root, candidate_scope_ids, "candidate")
    diagnostics.extend(candidate_expansion.diagnostics)
    if diagnostics:
        return _failure_response(
            effective_query,
            threshold,
            vector_index.model,
            vector_index.dimensions,
            diagnostics,
            len(candidate_expansion.requirements),
        )

    try:
        hits = vector_index.search_requirements(
            query=effective_query,
            candidates=candidate_expansion.requirements,
            threshold=threshold,
            max_results=max_results,
        )
    except Exception as exc:
        diagnostics.append(
            diagnostic(
                "requirement_search",
                "Requirement semantic search failed.",
                f"{type(exc).__name__}: {exc}",
            )
        )
        return _failure_response(
            effective_query,
            threshold,
            vector_index.model,
            vector_index.dimensions,
            diagnostics,
            len(candidate_expansion.requirements),
        )

    results = [
        _hit_as_dict(hit, include_details)
        for hit in sorted(hits, key=lambda item: -item.score)
        if hit.score >= threshold
    ][:max_results]
    return {
        "ok": True,
        "candidate_requirements": len(candidate_expansion.requirements),
        "query": effective_query,
        "threshold": threshold,
        "model": vector_index.model,
        "dimensions": vector_index.dimensions,
        "results": results,
        "diagnostics": diagnostics,
    }


def _failure_response(
    query: str,
    threshold: float,
    model: str,
    dimensions: int,
    diagnostics: list[dict[str, Any]],
    candidate_requirements: int = 0,
) -> dict[str, Any]:
    return {
        "ok": False,
        "candidate_requirements": candidate_requirements,
        "query": query,
        "threshold": threshold,
        "model": model,
        "dimensions": dimensions,
        "results": [],
        "diagnostics": diagnostics,
    }


def _hit_as_dict(
    hit: SimilarityHit,
    include_details: bool,
) -> dict[str, str | float]:
    result = hit.as_dict()
    result.pop("usdm_id", None)
    if not include_details:
        result.pop("detail", None)
    return result
