from __future__ import annotations

import sys
import tempfile
import unittest
from pathlib import Path


TOOLS_USDM_DIR = Path(__file__).resolve().parents[2]
if str(TOOLS_USDM_DIR) not in sys.path:
    sys.path.insert(0, str(TOOLS_USDM_DIR))

from similarity.usdm_loader import expand_scopes  # noqa: E402


RECORD = """\
# USDM requirement: Hierarchical loader

- **id**: `usdm:sample.requirements.hierarchy`
- **status**: draft
- **date**: 2026-09-30
- **kind**: requirement
- **parent**: root

## Requirements: Hierarchical loader requirements
> source: literal

| id | requirement |
|---|---|
| R001 | Parent requirement. |
| R001-01 | Child requirement. |
| R001-01-02 | Grandchild requirement. |
"""


class HierarchicalLoaderTests(unittest.TestCase):
    def test_record_scope_loads_hierarchical_rows(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "hierarchy.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            result = expand_scopes(
                root,
                ["usdm:sample.requirements.hierarchy"],
                "source",
            )

            self.assertEqual(result.diagnostics, [])
            self.assertEqual(
                [row.requirement_id for row in result.requirements],
                [
                    "usdm:sample.requirements.hierarchy#R001",
                    "usdm:sample.requirements.hierarchy#R001-01",
                    "usdm:sample.requirements.hierarchy#R001-01-02",
                ],
            )

    def test_exact_hierarchical_requirement_scope_resolves(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            path = root / "sample" / "records" / "usdm" / "hierarchy.md"
            path.parent.mkdir(parents=True)
            path.write_text(RECORD, encoding="utf-8")

            result = expand_scopes(
                root,
                ["usdm:sample.requirements.hierarchy#R001-01-02"],
                "source",
            )

            self.assertEqual(result.diagnostics, [])
            self.assertEqual(len(result.requirements), 1)
            self.assertEqual(
                result.requirements[0].detail,
                "Grandchild requirement.",
            )


if __name__ == "__main__":
    unittest.main()
