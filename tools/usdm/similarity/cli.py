#!/usr/bin/env python3
"""CLI for standalone USDM requirement similarity collection."""

from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

if __package__ in (None, ""):
    sys.path.insert(0, str(Path(__file__).resolve().parents[3]))
    from tools.usdm.similarity.collector import collect_similar_requirements
    from tools.usdm.similarity.config import (
        DEFAULT_EMBEDDING_MODEL,
        DEFAULT_VECTOR_DIMENSIONS,
    )
    from tools.usdm.similarity.embedding_generator import (
        OllamaEmbeddingGenerator,
    )
    from tools.usdm.similarity.vector_index import RequirementVectorIndex
else:
    from .collector import collect_similar_requirements
    from .config import DEFAULT_EMBEDDING_MODEL, DEFAULT_VECTOR_DIMENSIONS
    from .embedding_generator import OllamaEmbeddingGenerator
    from .vector_index import RequirementVectorIndex


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Standalone USDM requirement similarity tools."
    )
    subparsers = parser.add_subparsers(dest="command", required=True)
    collect_parser = subparsers.add_parser(
        "collect",
        help="Collect semantically similar USDM requirement candidates.",
    )
    collect_parser.add_argument("--repo-root", type=Path, required=True)
    collect_parser.add_argument(
        "--source-scope-id",
        action="append",
        required=True,
        dest="source_scope_ids",
    )
    collect_parser.add_argument(
        "--candidate-scope-id",
        action="append",
        dest="candidate_scope_ids",
    )
    collect_parser.add_argument("--threshold", type=float, default=0.86)
    collect_parser.add_argument("--max-candidates", type=int, default=10)
    collect_parser.add_argument(
        "--include-empty-items",
        action="store_true",
        default=False,
    )
    collect_parser.add_argument(
        "--include-details",
        action=argparse.BooleanOptionalAction,
        default=True,
    )
    collect_parser.add_argument("--max-total-hits", type=int, default=100)
    collect_parser.add_argument(
        "--exclude-same-document",
        action="store_true",
        default=False,
    )
    collect_parser.add_argument(
        "--qdrant-url",
        default="http://localhost:6333",
    )
    collect_parser.add_argument(
        "--ollama-url",
        default="http://localhost:11434",
    )
    collect_parser.add_argument(
        "--embedding-model",
        default=DEFAULT_EMBEDDING_MODEL,
    )
    collect_parser.add_argument(
        "--collection",
        default="usdm_requirement_embeddings",
    )
    return parser


def main(argv: list[str] | None = None) -> int:
    args = build_parser().parse_args(argv)
    embedding_generator = OllamaEmbeddingGenerator(
        base_url=args.ollama_url,
        model=args.embedding_model,
        dimensions=DEFAULT_VECTOR_DIMENSIONS,
    )
    vector_index = RequirementVectorIndex(
        qdrant_url=args.qdrant_url,
        collection=args.collection,
        embedding_generator=embedding_generator,
        model=args.embedding_model,
        dimensions=DEFAULT_VECTOR_DIMENSIONS,
    )
    response = collect_similar_requirements(
        repo_root=args.repo_root.resolve(),
        source_scope_ids=args.source_scope_ids,
        candidate_scope_ids=args.candidate_scope_ids,
        threshold=args.threshold,
        max_candidates_per_requirement=args.max_candidates,
        exclude_same_document=args.exclude_same_document,
        vector_index=vector_index,
        include_empty_items=args.include_empty_items,
        include_details=args.include_details,
        max_total_hits=args.max_total_hits,
    )
    print(json.dumps(response, indent=2, sort_keys=True))
    return 0 if response["ok"] else 1


if __name__ == "__main__":
    sys.exit(main())
