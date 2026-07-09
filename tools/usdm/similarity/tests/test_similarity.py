from __future__ import annotations

import tempfile
import unittest
from pathlib import Path

from tools.usdm.similarity.collector import collect_similar_requirements
from tools.usdm.similarity.config import DEFAULT_VECTOR_DIMENSIONS
from tools.usdm.similarity.models import RequirementRow, SimilarityHit
from tools.usdm.similarity.text_normalizer import (
    normalize_requirement_detail,
    requirement_detail_hash,
)
from tools.usdm.similarity.usdm_loader import expand_scopes
from tools.usdm.similarity.vector_index import RequirementVectorIndex


RECORD = """\
# USDM requirement: Example

- **id**: `usdm:sample.requirements.example`
- **status**: draft
- **date**: 2026-07-09
- **kind**: requirement
- **parent**: root

## What this is

Fixture.

## Requirements: spec:sample.example

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
            self.assertEqual(
                candidates[0]["requirement_id"],
                "usdm:sample.requirements.example#R002",
            )
            self.assertNotIn("classification", candidates[0])

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


if __name__ == "__main__":
    unittest.main()
