from __future__ import annotations

import tempfile
import unittest
from pathlib import Path

from tools.usdm.usdm_tools import (
    check_usdm_coverage,
    check_usdm_scope_coverage,
    validate_usdm,
)


RECORD_ONE = """\
# USDM requirement: One

- **id**: `usdm:sample.requirements.one`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: root

## Requirements: spec:sample.one

| id | requirement |
|---|---|
| R001 | Covered requirement. |
| R002 | Uncovered requirement. |
"""

RECORD_TWO = """\
# USDM requirement: Two

- **id**: `usdm:sample.requirements.two`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: root

## Requirements: spec:sample.two

| id | requirement |
|---|---|
| R001 | Covered second requirement. |
"""

RECORD_GAPPED = """\
# USDM requirement: Gapped

- **id**: `usdm:sample.requirements.gapped`
- **status**: draft
- **date**: 2026-07-10
- **kind**: requirement
- **parent**: root

## Requirements: spec:sample.gapped

| id | requirement |
|---|---|
| R001 | First retained requirement. |
| R003 | Third retained requirement after deleting R002. |
"""

SPEC = """\
# Contract: Coverage spec

- **id**: `spec:impl.coverage.alpha`
- **status**: draft
- **date**: 2026-07-10
- **parent**: `spec:impl`
- **contract_class**: `interface`
- **usdm_covers**:
  - `usdm:sample.requirements.one#R001`
  - `usdm:sample.requirements.two#R001`

## What this is

Fixture.
"""


def write_fixture(root: Path) -> None:
    usdm_root = root / "sample" / "records" / "usdm"
    usdm_root.mkdir(parents=True)
    (usdm_root / "one.md").write_text(RECORD_ONE, encoding="utf-8")
    (usdm_root / "two.md").write_text(RECORD_TWO, encoding="utf-8")

    spec_root = root / "impl" / "records" / "spec"
    spec_root.mkdir(parents=True)
    (spec_root / "coverage.md").write_text(SPEC, encoding="utf-8")


class ScopedCoverageTests(unittest.TestCase):
    def test_validate_usdm_allows_row_id_gaps(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            usdm_root = root / "sample" / "records" / "usdm"
            usdm_root.mkdir(parents=True)
            (usdm_root / "gapped.md").write_text(
                RECORD_GAPPED,
                encoding="utf-8",
            )

            response = validate_usdm(root, "sample")

            self.assertTrue(response["ok"])
            self.assertEqual(response["requirements"], 2)
            self.assertEqual(response["diagnostics"], [])

    def test_app_filtered_coverage_counts_cross_app_specs(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root)

            response = check_usdm_coverage(
                repo_root=root,
                app_namespace="sample",
                include_dangling=True,
            )

            self.assertFalse(response["ok"])
            self.assertEqual(response["requirements"], 3)
            self.assertEqual(response["covered"], 2)
            self.assertEqual(
                response["uncovered"],
                ["usdm:sample.requirements.one#R002"],
            )
            self.assertEqual(response["dangling"], [])

    def test_record_scope_coverage_grouping(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root)

            response = check_usdm_scope_coverage(
                ["usdm:sample.requirements.one"],
                repo_root=root,
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["records"], 1)
            self.assertEqual(response["requirements"], 2)
            self.assertEqual(response["covered_requirements"], 1)
            self.assertEqual(response["not_covered_requirements"], 1)
            self.assertEqual(
                response["items"],
                [
                    {
                        "record_id": "usdm:sample.requirements.one",
                        "covered": {"#R001": ["spec:impl.coverage.alpha"]},
                        "not_covered": ["#R002"],
                    }
                ],
            )

    def test_topic_scope_coverage_grouping(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root)

            response = check_usdm_scope_coverage(
                ["usdm:sample.requirements"],
                repo_root=root,
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["records"], 2)
            self.assertEqual(response["requirements"], 3)
            self.assertEqual(
                [item["record_id"] for item in response["items"]],
                ["usdm:sample.requirements.one", "usdm:sample.requirements.two"],
            )

    def test_app_scope_coverage_grouping(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root)

            response = check_usdm_scope_coverage(["usdm:sample"], repo_root=root)

            self.assertTrue(response["ok"])
            self.assertEqual(response["records"], 2)
            self.assertEqual(response["requirements"], 3)
            self.assertEqual(response["covered_requirements"], 2)
            self.assertEqual(response["not_covered_requirements"], 1)

    def test_full_requirement_scope_coverage(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root)

            response = check_usdm_scope_coverage(
                ["usdm:sample.requirements.one#R001"],
                repo_root=root,
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["records"], 1)
            self.assertEqual(response["requirements"], 1)
            self.assertEqual(response["covered_requirements"], 1)
            self.assertEqual(response["not_covered_requirements"], 0)
            self.assertEqual(
                response["items"][0]["covered"],
                {"#R001": ["spec:impl.coverage.alpha"]},
            )

    def test_uncovered_row_compact_listing(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root)

            response = check_usdm_scope_coverage(
                ["usdm:sample.requirements.one#R002"],
                repo_root=root,
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["items"][0]["not_covered"], ["#R002"])
            self.assertNotIn("covered", response["items"][0])

    def test_include_covered_false(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root)

            response = check_usdm_scope_coverage(
                ["usdm:sample.requirements.one"],
                repo_root=root,
                include_covered=False,
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["covered_requirements"], 1)
            self.assertNotIn("covered", response["items"][0])
            self.assertEqual(response["items"][0]["not_covered"], ["#R002"])

    def test_include_not_covered_false(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root)

            response = check_usdm_scope_coverage(
                ["usdm:sample.requirements.one"],
                repo_root=root,
                include_not_covered=False,
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["not_covered_requirements"], 1)
            self.assertEqual(
                response["items"][0]["covered"],
                {"#R001": ["spec:impl.coverage.alpha"]},
            )
            self.assertNotIn("not_covered", response["items"][0])

    def test_duplicate_scopes_do_not_duplicate_rows(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root)

            response = check_usdm_scope_coverage(
                [
                    "usdm:sample.requirements.one",
                    "usdm:sample.requirements.one#R001",
                    "usdm:sample.requirements.one",
                ],
                repo_root=root,
            )

            self.assertTrue(response["ok"])
            self.assertEqual(response["records"], 1)
            self.assertEqual(response["requirements"], 2)

    def test_invalid_and_empty_scope_errors(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root)

            empty = check_usdm_scope_coverage([], repo_root=root)
            invalid = check_usdm_scope_coverage(["spec:not.usdm"], repo_root=root)

            self.assertFalse(empty["ok"])
            self.assertEqual(empty["diagnostics"][0]["category"], "scope_ids")
            self.assertFalse(invalid["ok"])
            self.assertEqual(invalid["diagnostics"][0]["category"], "scope_id")

    def test_public_items_do_not_include_detail_path_or_usdm_id(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root)

            response = check_usdm_scope_coverage(["usdm:sample"], repo_root=root)

            self.assertTrue(response["ok"])
            serialized_items = repr(response["items"])
            self.assertNotIn("detail", serialized_items)
            self.assertNotIn("path", serialized_items)
            self.assertNotIn("usdm_id", serialized_items)


if __name__ == "__main__":
    unittest.main()
