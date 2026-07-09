"""Qdrant-backed requirement vector index."""

from __future__ import annotations

import json
import uuid
from typing import Any
from urllib.error import HTTPError, URLError
from urllib.parse import quote
from urllib.request import Request, urlopen

from .config import DEFAULT_VECTOR_DIMENSIONS
from .embedding_generator import OllamaEmbeddingGenerator
from .models import RequirementRow, SimilarityHit
from .text_normalizer import normalize_requirement_detail, requirement_detail_hash


POINT_NAMESPACE = uuid.UUID("c40f44cc-a5c2-4c45-8531-63ec2f48fa71")


class VectorIndexError(RuntimeError):
    """Raised when Qdrant synchronization or search fails."""


class QdrantClient:
    def __init__(self, base_url: str, timeout: float = 60.0) -> None:
        self.base_url = base_url.rstrip("/")
        self.timeout = timeout

    def request(
        self,
        method: str,
        path: str,
        body: dict[str, Any] | None = None,
        allow_not_found: bool = False,
    ) -> dict[str, Any] | None:
        data = None if body is None else json.dumps(body).encode("utf-8")
        request = Request(
            f"{self.base_url}{path}",
            data=data,
            headers={"Content-Type": "application/json"},
            method=method,
        )
        try:
            with urlopen(request, timeout=self.timeout) as response:
                payload = json.loads(response.read().decode("utf-8"))
        except HTTPError as exc:
            if allow_not_found and exc.code == 404:
                return None
            detail = exc.read().decode("utf-8", errors="replace")
            raise VectorIndexError(
                f"Qdrant request {method} {path} failed with HTTP "
                f"{exc.code}: {detail}"
            ) from exc
        except (URLError, OSError, json.JSONDecodeError) as exc:
            raise VectorIndexError(
                f"Qdrant request {method} {path} failed: {exc}"
            ) from exc
        if not isinstance(payload, dict):
            raise VectorIndexError("Qdrant returned a non-object JSON response.")
        if payload.get("status") not in (None, "ok"):
            raise VectorIndexError(f"Qdrant returned failure status: {payload}")
        return payload


class RequirementVectorIndex:
    def __init__(
        self,
        qdrant_url: str,
        collection: str,
        embedding_generator: OllamaEmbeddingGenerator,
        model: str,
        dimensions: int = DEFAULT_VECTOR_DIMENSIONS,
        embedding_batch_size: int = 32,
    ) -> None:
        self.qdrant = QdrantClient(qdrant_url)
        self.collection = collection
        self.embedding_generator = embedding_generator
        self.model = model
        self.dimensions = dimensions
        self.embedding_batch_size = embedding_batch_size

    def collect_candidates(
        self,
        sources: list[RequirementRow],
        candidates: list[RequirementRow],
        threshold: float,
        max_candidates: int,
        exclude_same_document: bool,
    ) -> dict[str, list[SimilarityHit]]:
        all_rows = {
            row.requirement_id: row for row in [*sources, *candidates]
        }
        self._ensure_collection()
        self._synchronize(list(all_rows.values()))

        source_point_ids = [point_id(row.requirement_id) for row in sources]
        source_points = self._retrieve(source_point_ids, with_vector=True)
        candidate_by_point_id = {
            point_id(row.requirement_id): row for row in candidates
        }
        results: dict[str, list[SimilarityHit]] = {}

        for source in sources:
            source_point = source_points.get(point_id(source.requirement_id))
            vector = None if source_point is None else source_point.get("vector")
            if not isinstance(vector, list) or len(vector) != self.dimensions:
                raise VectorIndexError(
                    f"Qdrant did not return a valid source vector for "
                    f"{source.requirement_id}."
                )

            allowed_point_ids = [
                candidate_id
                for candidate_id, candidate in candidate_by_point_id.items()
                if candidate.requirement_id != source.requirement_id
                and (
                    not exclude_same_document
                    or candidate.usdm_id != source.usdm_id
                )
            ]
            if not allowed_point_ids:
                results[source.requirement_id] = []
                continue

            hits = self._search(
                vector=vector,
                allowed_point_ids=allowed_point_ids,
                threshold=threshold,
                limit=max_candidates,
            )
            source_hits: list[SimilarityHit] = []
            for hit in hits:
                candidate = candidate_by_point_id.get(str(hit.get("id", "")))
                score = hit.get("score")
                if candidate is None or not isinstance(score, (int, float)):
                    continue
                source_hits.append(
                    SimilarityHit(requirement=candidate, score=float(score))
                )
            results[source.requirement_id] = source_hits
        return results

    def _ensure_collection(self) -> None:
        name = quote(self.collection, safe="")
        response = self.qdrant.request(
            "GET", f"/collections/{name}", allow_not_found=True
        )
        if response is None:
            self.qdrant.request(
                "PUT",
                f"/collections/{name}",
                {
                    "vectors": {
                        "size": self.dimensions,
                        "distance": "Cosine",
                    }
                },
            )
            return

        result = response.get("result")
        vectors = (
            result.get("config", {}).get("params", {}).get("vectors")
            if isinstance(result, dict)
            else None
        )
        if not isinstance(vectors, dict):
            raise VectorIndexError(
                "Existing Qdrant collection has unsupported vector configuration."
            )
        size = vectors.get("size")
        distance = vectors.get("distance")
        if size != self.dimensions or str(distance).casefold() != "cosine":
            raise VectorIndexError(
                "Existing Qdrant collection vector configuration does not match "
                f"dimensions={self.dimensions}, distance=Cosine "
                f"(found dimensions={size}, distance={distance})."
            )

    def _synchronize(self, rows: list[RequirementRow]) -> None:
        points = self._retrieve(
            [point_id(row.requirement_id) for row in rows],
            with_vector=False,
        )
        stale: list[RequirementRow] = []
        for row in rows:
            stored = points.get(point_id(row.requirement_id))
            payload = stored.get("payload") if isinstance(stored, dict) else None
            if not self._payload_is_fresh(payload, row):
                stale.append(row)

        for offset in range(0, len(stale), self.embedding_batch_size):
            batch = stale[offset : offset + self.embedding_batch_size]
            normalized = [
                normalize_requirement_detail(row.detail) for row in batch
            ]
            vectors = self.embedding_generator.embed(normalized)
            qdrant_points = [
                {
                    "id": point_id(row.requirement_id),
                    "vector": vector,
                    "payload": {
                        "requirement_id": row.requirement_id,
                        "detail": row.detail,
                        "usdm_id": row.usdm_id,
                        "path": row.path,
                        "detail_hash": requirement_detail_hash(row.detail),
                        "embedding_model": self.model,
                        "vector_dimensions": self.dimensions,
                    },
                }
                for row, vector in zip(batch, vectors, strict=True)
            ]
            self._upsert(qdrant_points)

    def _payload_is_fresh(
        self,
        payload: Any,
        row: RequirementRow,
    ) -> bool:
        return (
            isinstance(payload, dict)
            and payload.get("requirement_id") == row.requirement_id
            and payload.get("detail_hash") == requirement_detail_hash(row.detail)
            and payload.get("embedding_model") == self.model
            and payload.get("vector_dimensions") == self.dimensions
        )

    def _retrieve(
        self,
        ids: list[str],
        with_vector: bool,
    ) -> dict[str, dict[str, Any]]:
        if not ids:
            return {}
        name = quote(self.collection, safe="")
        found: dict[str, dict[str, Any]] = {}
        for offset in range(0, len(ids), 256):
            response = self.qdrant.request(
                "POST",
                f"/collections/{name}/points",
                {
                    "ids": ids[offset : offset + 256],
                    "with_payload": True,
                    "with_vector": with_vector,
                },
            )
            result = None if response is None else response.get("result")
            if not isinstance(result, list):
                raise VectorIndexError(
                    "Qdrant point retrieval returned an invalid result."
                )
            for point in result:
                if isinstance(point, dict) and "id" in point:
                    found[str(point["id"])] = point
        return found

    def _upsert(self, points: list[dict[str, Any]]) -> None:
        name = quote(self.collection, safe="")
        self.qdrant.request(
            "PUT",
            f"/collections/{name}/points?wait=true",
            {"points": points},
        )

    def _search(
        self,
        vector: list[float],
        allowed_point_ids: list[str],
        threshold: float,
        limit: int,
    ) -> list[dict[str, Any]]:
        name = quote(self.collection, safe="")
        response = self.qdrant.request(
            "POST",
            f"/collections/{name}/points/search",
            {
                "vector": vector,
                "filter": {"must": [{"has_id": allowed_point_ids}]},
                "limit": limit,
                "score_threshold": threshold,
                "with_payload": False,
                "with_vector": False,
            },
        )
        result = None if response is None else response.get("result")
        if not isinstance(result, list):
            raise VectorIndexError("Qdrant search returned an invalid result.")
        return [item for item in result if isinstance(item, dict)]


def point_id(requirement_id: str) -> str:
    return str(uuid.uuid5(POINT_NAMESPACE, requirement_id))
