"""Source-centric USDM similarity collection orchestration."""

from __future__ import annotations

from pathlib import Path
from typing import Any, Protocol

from .models import RequirementRow, SimilarityHit
from .usdm_loader import diagnostic, expand_scopes


class VectorIndex(Protocol):
    model: str
    dimensions: int

    def collect_candidates(
        self,
        sources: list[RequirementRow],
        candidates: list[RequirementRow],
        threshold: float,
        max_candidates: int,
        exclude_same_document: bool,
    ) -> dict[str, list[SimilarityHit]]: ...


def collect_similar_requirements(
    repo_root: Path,
    source_scope_ids: list[str],
    candidate_scope_ids: list[str] | None,
    threshold: float,
    max_candidates_per_requirement: int,
    exclude_same_document: bool,
    vector_index: VectorIndex,
) -> dict[str, Any]:
    diagnostics: list[dict[str, Any]] = []
    if not repo_root.is_dir():
        diagnostics.append(
            diagnostic(
                "repo_root",
                "repo_root is missing or unreadable.",
                str(repo_root),
            )
        )
    if not source_scope_ids:
        diagnostics.append(
            diagnostic(
                "source_scope_ids",
                "source_scope_ids must contain at least one USDM scope ID.",
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
    if max_candidates_per_requirement < 1:
        diagnostics.append(
            diagnostic(
                "max_candidates_per_requirement",
                "max_candidates_per_requirement must be at least 1.",
                max_candidates_per_requirement,
            )
        )
    if diagnostics:
        return _failure_response(
            threshold, vector_index.model, vector_index.dimensions, diagnostics
        )

    source_expansion = expand_scopes(repo_root, source_scope_ids, "source")
    effective_candidate_scope_ids = (
        source_scope_ids
        if candidate_scope_ids is None
        else candidate_scope_ids
    )
    candidate_expansion = expand_scopes(
        repo_root,
        effective_candidate_scope_ids,
        "candidate",
    )
    diagnostics.extend(source_expansion.diagnostics)
    diagnostics.extend(candidate_expansion.diagnostics)
    if diagnostics:
        return _failure_response(
            threshold,
            vector_index.model,
            vector_index.dimensions,
            diagnostics,
            len(source_expansion.requirements),
            len(candidate_expansion.requirements),
        )

    try:
        hits_by_source = vector_index.collect_candidates(
            sources=source_expansion.requirements,
            candidates=candidate_expansion.requirements,
            threshold=threshold,
            max_candidates=max_candidates_per_requirement,
            exclude_same_document=exclude_same_document,
        )
    except Exception as exc:
        diagnostics.append(
            diagnostic(
                "similarity_collection",
                "Requirement similarity collection failed.",
                f"{type(exc).__name__}: {exc}",
            )
        )
        return _failure_response(
            threshold,
            vector_index.model,
            vector_index.dimensions,
            diagnostics,
            len(source_expansion.requirements),
            len(candidate_expansion.requirements),
        )

    items = [
        {
            "source": source.as_dict(),
            "candidates": [
                hit.as_dict()
                for hit in hits_by_source.get(source.requirement_id, [])
            ],
        }
        for source in source_expansion.requirements
    ]
    return {
        "ok": True,
        "source_requirements": len(source_expansion.requirements),
        "candidate_requirements": len(candidate_expansion.requirements),
        "threshold": threshold,
        "model": vector_index.model,
        "dimensions": vector_index.dimensions,
        "items": items,
        "diagnostics": diagnostics,
    }


def _failure_response(
    threshold: float,
    model: str,
    dimensions: int,
    diagnostics: list[dict[str, Any]],
    source_requirements: int = 0,
    candidate_requirements: int = 0,
) -> dict[str, Any]:
    return {
        "ok": False,
        "source_requirements": source_requirements,
        "candidate_requirements": candidate_requirements,
        "threshold": threshold,
        "model": model,
        "dimensions": dimensions,
        "items": [],
        "diagnostics": diagnostics,
    }
