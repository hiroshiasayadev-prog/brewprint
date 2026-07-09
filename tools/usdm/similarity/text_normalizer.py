"""Normalization used for embedding input and freshness detection."""

from __future__ import annotations

import hashlib
import re


REPEATED_WHITESPACE_RE = re.compile(r"\s+")


def normalize_requirement_detail(detail: str) -> str:
    """Collapse repeated whitespace without changing domain identifiers or code spans."""
    return REPEATED_WHITESPACE_RE.sub(" ", detail).strip()


def requirement_detail_hash(detail: str) -> str:
    normalized = normalize_requirement_detail(detail)
    return hashlib.sha256(normalized.encode("utf-8")).hexdigest()
