"""Data models for USDM similarity collection."""

from __future__ import annotations

from dataclasses import dataclass


@dataclass(frozen=True)
class RequirementRow:
    requirement_id: str
    detail: str
    usdm_id: str
    path: str

    def as_dict(self) -> dict[str, str]:
        return {
            "requirement_id": self.requirement_id,
            "detail": self.detail,
            "usdm_id": self.usdm_id,
            "path": self.path,
        }


@dataclass(frozen=True)
class SimilarityHit:
    requirement: RequirementRow
    score: float

    def as_dict(self) -> dict[str, str | float]:
        return {**self.requirement.as_dict(), "score": self.score}
