"""Ollama embedding backend."""

from __future__ import annotations

import json
from typing import Any
from urllib.error import HTTPError, URLError
from urllib.request import Request, urlopen

from .config import DEFAULT_VECTOR_DIMENSIONS


class EmbeddingError(RuntimeError):
    """Raised when Ollama cannot produce valid embeddings."""


class OllamaEmbeddingGenerator:
    def __init__(
        self,
        base_url: str,
        model: str,
        dimensions: int = DEFAULT_VECTOR_DIMENSIONS,
        timeout: float = 120.0,
    ) -> None:
        self.base_url = base_url.rstrip("/")
        self.model = model
        self.dimensions = dimensions
        self.timeout = timeout

    def embed(self, texts: list[str]) -> list[list[float]]:
        if not texts:
            return []

        payload = self._post(
            "/api/embed",
            {
                "model": self.model,
                "input": texts,
            },
        )
        embeddings = payload.get("embeddings")
        if not isinstance(embeddings, list) or len(embeddings) != len(texts):
            raise EmbeddingError(
                "Ollama response did not contain one embedding per input text."
            )

        validated: list[list[float]] = []
        for index, vector in enumerate(embeddings):
            if not isinstance(vector, list) or len(vector) != self.dimensions:
                actual = len(vector) if isinstance(vector, list) else "non-list"
                raise EmbeddingError(
                    f"Ollama embedding {index} has dimensions {actual}; "
                    f"expected {self.dimensions}."
                )
            if not all(isinstance(value, (int, float)) for value in vector):
                raise EmbeddingError(
                    f"Ollama embedding {index} contains a non-numeric value."
                )
            validated.append([float(value) for value in vector])
        return validated

    def _post(self, path: str, body: dict[str, Any]) -> dict[str, Any]:
        request = Request(
            f"{self.base_url}{path}",
            data=json.dumps(body).encode("utf-8"),
            headers={"Content-Type": "application/json"},
            method="POST",
        )
        try:
            with urlopen(request, timeout=self.timeout) as response:
                result = json.loads(response.read().decode("utf-8"))
        except HTTPError as exc:
            detail = exc.read().decode("utf-8", errors="replace")
            raise EmbeddingError(
                f"Ollama request failed with HTTP {exc.code}: {detail}"
            ) from exc
        except (URLError, OSError, json.JSONDecodeError) as exc:
            raise EmbeddingError(f"Ollama request failed: {exc}") from exc
        if not isinstance(result, dict):
            raise EmbeddingError("Ollama returned a non-object JSON response.")
        return result
