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
    include_empty_items: bool = False,
    include_details: bool = True,
    max_total_hits: int | None = 100,
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
    if max_total_hits is not None and max_total_hits < 1:
        diagnostics.append(
            diagnostic(
                "max_total_hits",
                "max_total_hits must be at least 1 when provided.",
                max_total_hits,
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

    selected_hits_by_source = _select_hits_by_score(
        source_expansion.requirements,
        hits_by_source,
        max_total_hits,
    )
    items = []
    for source in source_expansion.requirements:
        hits = selected_hits_by_source.get(source.requirement_id, [])
        if not include_empty_items and not hits:
            continue
        items.append(
            {
                "source": _requirement_as_dict(source, include_details),
                "candidates": [
                    _hit_as_dict(hit, include_details) for hit in hits
                ],
            }
        )
    return {
        "ok": True,
        "source_requirements": len(source_expansion.requirements),
        "candidate_requirements": len(candidate_expansion.requirements),
        "threshold": threshold,
        "model": vector_index.model,
        "dimensions": vector_index.dimensions,
        "returned_hits": sum(len(item["candidates"]) for item in items),
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
        "returned_hits": 0,
        "items": [],
        "diagnostics": diagnostics,
    }


def _select_hits_by_score(
    sources: list[RequirementRow],
    hits_by_source: dict[str, list[SimilarityHit]],
    max_total_hits: int | None,
) -> dict[str, list[SimilarityHit]]:
    source_order = {
        source.requirement_id: index for index, source in enumerate(sources)
    }
    ranked = [
        (source_order.get(source_id, len(source_order)), hit_index, source_id, hit)
        for source_id, hits in hits_by_source.items()
        for hit_index, hit in enumerate(hits)
    ]
    ranked.sort(key=lambda item: (-item[3].score, item[0], item[1]))
    if max_total_hits is not None:
        ranked = ranked[:max_total_hits]

    selected: dict[str, list[SimilarityHit]] = {}
    for _, _, source_id, hit in ranked:
        selected.setdefault(source_id, []).append(hit)
    for hits in selected.values():
        hits.sort(key=lambda hit: -hit.score)
    return selected


def _requirement_as_dict(
    row: RequirementRow,
    include_details: bool,
) -> dict[str, str]:
    result = row.as_dict()
    result.pop("usdm_id", None)
    if not include_details:
        result.pop("detail", None)
    return result


def _hit_as_dict(
    hit: SimilarityHit,
    include_details: bool,
) -> dict[str, str | float]:
    result = hit.as_dict()
    result.pop("usdm_id", None)
    if not include_details:
        result.pop("detail", None)
    return result
