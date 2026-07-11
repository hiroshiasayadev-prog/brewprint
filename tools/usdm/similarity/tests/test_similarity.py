from __future__ import annotations

import sys
import tempfile
import unittest
from pathlib import Path


TOOLS_USDM_DIR = Path(__file__).resolve().parents[2]
if str(TOOLS_USDM_DIR) not in sys.path:
    sys.path.insert(0, str(TOOLS_USDM_DIR))

from similarity.collector import collect_similar_requirements  # noqa: E402
from similarity.config import DEFAULT_VECTOR_DIMENSIONS  # noqa: E402
from similarity.models import RequirementRow, SimilarityHit  # noqa: E402
from similarity.search import search_requirements  # noqa: E402
from similarity.text_normalizer import (  # noqa: E402
    normalize_requirement_detail,
    requirement_detail_hash,
)
from similarity.usdm_loader import expand_scopes  # noqa: E402
from similarity.vector_index import RequirementVectorIndex  # noqa: E402


RECORD = """\
# USDM requirement: Example

- **id**: `usdm:sample.requirements.example`
- **status**: draft
- **date**: 2026-07-09
- **kind**: requirement
- **parent**: root

## What this is

Fixture.

## Requirements: Example requirements
> source: spec:sample.example

| id | requirement | notes |
|---|---|---|
| R001 | Preserve `domain_id` and   collapse whitespace. | |
| R002 | Compare requirement details. | |
"""


class FakeVectorIndex:
    model = "test-model"
    dimensions = DEFAULT_VECTOR_DIMENSIONS

    def collect_candidates(
        self,
        sources: list[RequirementRow],
        candidates: list[RequirementRow],
        threshold: float,
        max_candidates: int,
        exclude_same_document: bool,
    ) -> dict[str, list[SimilarityHit]]:
        result: dict[str, list[SimilarityHit]] = {}
        for source in sources:
            result[source.requirement_id] = [
                SimilarityHit(candidate, 0.91)
                for candidate in candidates
                if candidate.requirement_id != source.requirement_id
            ][:max_candidates]
        return result


class FakeSearchVectorIndex:
    model = "test-model"
    dimensions = DEFAULT_VECTOR_DIMENSIONS

    def search_requirements(
        self,
        query: str,
        candidates: list[RequirementRow],
        threshold: float,
        max_results: int,
    ) -> list[SimilarityHit]:
        scores = [0.72, 0.91, 0.12]
        hits = [
            SimilarityHit(candidate, score)
            for candidate, score in zip(candidates, scores, strict=False)
            if score >= threshold
        ]
        hits.sort(key=lambda hit: -hit.score)
        return hits[:max_results]


class RecordingEmbeddingGenerator:
    def __init__(self, events: list[str]) -> None:
        self.events = events

    def embed(self, texts: list[str]) -> list[list[float]]:
        self.events.append("embed_query")
        return [[0.1] * DEFAULT_VECTOR_DIMENSIONS for _ in texts]


class RecordingRequirementVectorIndex(RequirementVectorIndex):
    def __init__(self) -> None:
        self.events: list[str] = []
        self.collection = "test"
        self.embedding_generator = RecordingEmbeddingGenerator(self.events)
        self.model = "test-model"
        self.dimensions = DEFAULT_VECTOR_DIMENSIONS

    def _ensure_collection(self) -> None:
        self.events.append("ensure_collection")

    def _synchronize(self, rows: list[RequirementRow]) -> None:
        self.events.append("synchronize_candidates")

    def _search(
        self,
        vector: list[float],
        allowed_point_ids: list[str],
        threshold: float,
        limit: int,
    ) -> list[dict[str, object]]:
        self.events.append("search")
        return [{"id": allowed_point_ids[0], "score": 0.88}]


class SimilarityToolTests(unittest.TestCase):
    def test_normalization_and_hash_are_whitespace_stable(self) -> None:
        first = " Preserve  `domain_id`\n value "
        second = "Preserve `domain_id` value"
        self.assertEqual(normalize_requirement_detail(first), second)
        self.assertEqual(
            requirement_detail_hash(first),
            requirement_detail_hash(second),
        )

    def test_expand_record_and_full_requirement_scopes(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            record = expand_scopes(
                root, ["usdm:sample.requirements.example"], "source"
            )
            full = expand_scopes(
                root,
                ["usdm:sample.requirements.example#R002"],
                "source",
            )

            self.assertEqual(record.diagnostics, [])
            self.assertEqual(len(record.requirements), 2)
            self.assertEqual(
                full.requirements[0].detail,
                "Compare requirement details.",
            )

    def test_expand_scopes_loads_literal_requirement_sections(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(
                RECORD.replace(
                    "> source: spec:sample.example",
                    "> source: literal",
                ),
                encoding="utf-8",
            )

            result = expand_scopes(
                root,
                ["usdm:sample.requirements.example"],
                "source",
            )

            self.assertEqual(result.diagnostics, [])
            self.assertEqual(len(result.requirements), 2)

    def test_expand_scopes_does_not_load_missing_source_sections(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(
                RECORD.replace("> source: spec:sample.example\n", ""),
                encoding="utf-8",
            )

            result = expand_scopes(
                root,
                ["usdm:sample.requirements.example"],
                "source",
            )

            self.assertEqual(result.requirements, [])
            self.assertEqual(result.diagnostics[0]["category"], "scope_resolution")

    def test_expand_topic_scope_to_descendant_requirement_records(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            topic = expand_scopes(
                root,
                ["usdm:sample.requirements"],
                "source",
            )

            self.assertEqual(topic.diagnostics, [])
            self.assertEqual(len(topic.requirements), 2)

    def test_expand_app_scope_to_all_app_requirement_records(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            first = root / "sample" / "records" / "usdm" / "example.md"
            second = root / "sample" / "records" / "usdm" / "other.md"
            first.parent.mkdir(parents=True)
            first.write_text(RECORD, encoding="utf-8")
            second.write_text(
                RECORD.replace(
                    "usdm:sample.requirements.example",
                    "usdm:sample.other.example",
                ),
                encoding="utf-8",
            )

            app = expand_scopes(root, ["usdm:sample"], "candidate")

            self.assertEqual(app.diagnostics, [])
            self.assertEqual(len(app.requirements), 4)
            self.assertEqual(
                app.requirements[0].requirement_id,
                "usdm:sample.other.example#R001",
            )

    def test_collector_returns_source_centric_candidates(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            response = collect_similar_requirements(
                repo_root=root,
                source_scope_ids=[
                    "usdm:sample.requirements.example#R001"
                ],
                candidate_scope_ids=[
                    "usdm:sample.requirements.example"
                ],
                threshold=0.86,
                max_candidates_per_requirement=10,
                exclude_same_document=False,
                vector_index=FakeVectorIndex(),
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["source_requirements"], 1)
            self.assertEqual(response["candidate_requirements"], 2)
            candidates = response["items"][0]["candidates"]
            self.assertEqual(len(candidates), 1)
            self.assertEqual(response["returned_hits"], 1)
            self.assertEqual(
                candidates[0]["requirement_id"],
                "usdm:sample.requirements.example#R002",
            )
            self.assertNotIn("usdm_id", response["items"][0]["source"])
            self.assertNotIn("usdm_id", candidates[0])
            self.assertNotIn("classification", candidates[0])

    def test_collector_omits_empty_items_by_default(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            response = collect_similar_requirements(
                repo_root=root,
                source_scope_ids=[
                    "usdm:sample.requirements.example#R001"
                ],
                candidate_scope_ids=[
                    "usdm:sample.requirements.example#R001"
                ],
                threshold=0.86,
                max_candidates_per_requirement=10,
                exclude_same_document=False,
                vector_index=FakeVectorIndex(),
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["returned_hits"], 0)
            self.assertEqual(response["items"], [])

    def test_collector_can_include_empty_items(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            response = collect_similar_requirements(
                repo_root=root,
                source_scope_ids=[
                    "usdm:sample.requirements.example#R001"
                ],
                candidate_scope_ids=[
                    "usdm:sample.requirements.example#R001"
                ],
                threshold=0.86,
                max_candidates_per_requirement=10,
                exclude_same_document=False,
                vector_index=FakeVectorIndex(),
                include_empty_items=True,
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["returned_hits"], 0)
            self.assertEqual(len(response["items"]), 1)
            self.assertEqual(response["items"][0]["candidates"], [])

    def test_collector_can_omit_details(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            response = collect_similar_requirements(
                repo_root=root,
                source_scope_ids=[
                    "usdm:sample.requirements.example#R001"
                ],
                candidate_scope_ids=[
                    "usdm:sample.requirements.example"
                ],
                threshold=0.86,
                max_candidates_per_requirement=10,
                exclude_same_document=False,
                vector_index=FakeVectorIndex(),
                include_details=False,
            )

            self.assertTrue(response["ok"])
            source = response["items"][0]["source"]
            candidate = response["items"][0]["candidates"][0]
            self.assertNotIn("detail", source)
            self.assertNotIn("detail", candidate)

    def test_collector_limits_total_hits_by_score(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            response = collect_similar_requirements(
                repo_root=root,
                source_scope_ids=["usdm:sample.requirements.example"],
                candidate_scope_ids=["usdm:sample.requirements.example"],
                threshold=0.86,
                max_candidates_per_requirement=10,
                exclude_same_document=False,
                vector_index=FakeVectorIndex(),
                max_total_hits=1,
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["returned_hits"], 1)
            total_candidates = sum(
                len(item["candidates"]) for item in response["items"]
            )
            self.assertEqual(total_candidates, 1)

    def test_freshness_checks_hash_model_and_dimensions(self) -> None:
        index = object.__new__(RequirementVectorIndex)
        index.model = "expected-model"
        index.dimensions = DEFAULT_VECTOR_DIMENSIONS
        row = RequirementRow(
            requirement_id="usdm:sample.requirements.example#R001",
            detail="Requirement detail.",
            usdm_id="usdm:sample.requirements.example",
            path="sample/records/usdm/example.md",
        )
        fresh = {
            "requirement_id": row.requirement_id,
            "detail_hash": requirement_detail_hash(row.detail),
            "embedding_model": "expected-model",
            "vector_dimensions": DEFAULT_VECTOR_DIMENSIONS,
        }

        self.assertTrue(index._payload_is_fresh(fresh, row))
        for field, value in (
            ("detail_hash", "stale"),
            ("embedding_model", "other-model"),
            ("vector_dimensions", 384),
        ):
            stale = {**fresh, field: value}
            self.assertFalse(index._payload_is_fresh(stale, row))

    def test_search_returns_query_centric_results(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            response = search_requirements(
                repo_root=root,
                query="  compare   requirement details  ",
                candidate_scope_ids=["usdm:sample.requirements.example"],
                threshold=0.30,
                max_results=20,
                vector_index=FakeSearchVectorIndex(),
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["candidate_requirements"], 2)
            self.assertEqual(response["query"], "compare requirement details")
            self.assertEqual(len(response["results"]), 2)
            self.assertEqual(response["results"][0]["score"], 0.91)
            self.assertEqual(response["results"][1]["score"], 0.72)

    def test_search_rejects_empty_query(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            response = search_requirements(
                repo_root=Path(directory),
                query="   ",
                candidate_scope_ids=["usdm:sample.requirements.example"],
                threshold=0.30,
                max_results=20,
                vector_index=FakeSearchVectorIndex(),
            )

            self.assertFalse(response["ok"])
            self.assertEqual(response["results"], [])
            self.assertEqual(response["diagnostics"][0]["category"], "query")

    def test_search_rejects_empty_candidate_scope(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            response = search_requirements(
                repo_root=Path(directory),
                query="requirement detail",
                candidate_scope_ids=[],
                threshold=0.30,
                max_results=20,
                vector_index=FakeSearchVectorIndex(),
            )

            self.assertFalse(response["ok"])
            self.assertEqual(
                response["diagnostics"][0]["category"],
                "candidate_scope_ids",
            )

    def test_search_rejects_invalid_threshold(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            response = search_requirements(
                repo_root=Path(directory),
                query="requirement detail",
                candidate_scope_ids=["usdm:sample.requirements.example"],
                threshold=1.1,
                max_results=20,
                vector_index=FakeSearchVectorIndex(),
            )

            self.assertFalse(response["ok"])
            self.assertEqual(response["diagnostics"][0]["category"], "threshold")

    def test_search_rejects_invalid_max_results(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            response = search_requirements(
                repo_root=Path(directory),
                query="requirement detail",
                candidate_scope_ids=["usdm:sample.requirements.example"],
                threshold=0.30,
                max_results=0,
                vector_index=FakeSearchVectorIndex(),
            )

            self.assertFalse(response["ok"])
            self.assertEqual(response["diagnostics"][0]["category"], "max_results")

    def test_search_can_omit_details(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            response = search_requirements(
                repo_root=root,
                query="requirement detail",
                candidate_scope_ids=["usdm:sample.requirements.example"],
                threshold=0.30,
                max_results=20,
                vector_index=FakeSearchVectorIndex(),
                include_details=False,
            )

            self.assertTrue(response["ok"])
            self.assertNotIn("detail", response["results"][0])

    def test_search_limits_result_count(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            response = search_requirements(
                repo_root=root,
                query="requirement detail",
                candidate_scope_ids=["usdm:sample.requirements.example"],
                threshold=0.30,
                max_results=1,
                vector_index=FakeSearchVectorIndex(),
            )

            self.assertTrue(response["ok"])
            self.assertEqual(len(response["results"]), 1)
            self.assertEqual(response["results"][0]["score"], 0.91)

    def test_search_public_results_do_not_include_usdm_id(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "example.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            response = search_requirements(
                repo_root=root,
                query="requirement detail",
                candidate_scope_ids=["usdm:sample.requirements.example"],
                threshold=0.30,
                max_results=20,
                vector_index=FakeSearchVectorIndex(),
            )

            self.assertTrue(response["ok"])
            self.assertNotIn("usdm_id", response["results"][0])

    def test_vector_index_synchronizes_candidates_before_query_embedding(
        self,
    ) -> None:
        index = RecordingRequirementVectorIndex()
        row = RequirementRow(
            requirement_id="usdm:sample.requirements.example#R001",
            detail="Requirement detail.",
            usdm_id="usdm:sample.requirements.example",
            path="sample/records/usdm/example.md",
        )

        hits = index.search_requirements(
            query="requirement detail",
            candidates=[row],
            threshold=0.30,
            max_results=20,
        )

        self.assertEqual(len(hits), 1)
        self.assertEqual(
            index.events,
            ["ensure_collection", "synchronize_candidates", "embed_query", "search"],
        )


if __name__ == "__main__":
    unittest.main()
