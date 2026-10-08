from __future__ import annotations

import sys
import tempfile
import unittest
from pathlib import Path


TOOLS_USDM_DIR = Path(__file__).resolve().parents[1]
if str(TOOLS_USDM_DIR) not in sys.path:
    sys.path.insert(0, str(TOOLS_USDM_DIR))

from usdm_tools import (  # noqa: E402
    check_usdm_coverage,
    check_usdm_scope_coverage,
    expand_coverage_value,
    usdm_covered_by,
    validate_usdm,
)


RECORD_ID = "usdm:sample.requirements.hierarchy"


def make_record(rows: list[str]) -> str:
    table_rows = "\n".join(f"| {row_id} | Requirement {row_id}. |" for row_id in rows)
    return f"""\
# USDM requirement: Hierarchy

- **id**: `{RECORD_ID}`
- **status**: draft
- **date**: 2026-09-30
- **kind**: requirement
- **parent**: root

## Requirements: Hierarchical requirements
> source: literal

| id | requirement |
|---|---|
{table_rows}
"""


def make_spec(spec_id: str, covers: list[str]) -> str:
    cover_lines = "\n".join(f"  - `{value}`" for value in covers)
    return f"""\
# Contract: Coverage fixture

- **id**: `{spec_id}`
- **status**: draft
- **date**: 2026-09-30
- **parent**: `spec:impl`
- **contract_class**: interface
- **usdm_covers**:
{cover_lines}

## What this is

Fixture.
"""


def write_fixture(
    root: Path,
    rows: list[str],
    specs: list[tuple[str, list[str]]],
) -> None:
    usdm_root = root / "sample" / "records" / "usdm"
    usdm_root.mkdir(parents=True)
    (usdm_root / "hierarchy.md").write_text(make_record(rows), encoding="utf-8")

    spec_root = root / "impl" / "records" / "spec"
    spec_root.mkdir(parents=True)
    for index, (spec_id, covers) in enumerate(specs, start=1):
        (spec_root / f"coverage-{index}.md").write_text(
            make_spec(spec_id, covers),
            encoding="utf-8",
        )


class HierarchicalValidationTests(unittest.TestCase):
    def test_child_and_grandchild_rows_validate(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001", "R001-01", "R001-01-01", "R001-02"],
                [],
            )

            response = validate_usdm(root, "sample")

            self.assertTrue(response["ok"])
            self.assertEqual(response["requirements"], 4)
            self.assertEqual(response["diagnostics"], [])

    def test_orphan_grandchild_is_rejected(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root, ["R001", "R001-01-01"], [])

            response = validate_usdm(root, "sample")

            self.assertFalse(response["ok"])
            self.assertEqual(
                [item["category"] for item in response["diagnostics"]],
                ["row_parent"],
            )
            self.assertEqual(
                response["diagnostics"][0]["value"],
                f"{RECORD_ID}#R001-01-01",
            )


class HierarchicalCoverageParsingTests(unittest.TestCase):
    def test_hierarchical_compact_comma_list(self) -> None:
        requirement_ids, diagnostics = expand_coverage_value(
            f"{RECORD_ID}#R001-01,R001-02",
            "fixture.md",
        )

        self.assertEqual(diagnostics, [])
        self.assertEqual(
            requirement_ids,
            [f"{RECORD_ID}#R001-01", f"{RECORD_ID}#R001-02"],
        )

    def test_top_level_range_compatibility(self) -> None:
        requirement_ids, diagnostics = expand_coverage_value(
            f"{RECORD_ID}#R001-R003",
            "fixture.md",
        )

        self.assertEqual(diagnostics, [])
        self.assertEqual(
            requirement_ids,
            [
                f"{RECORD_ID}#R001",
                f"{RECORD_ID}#R002",
                f"{RECORD_ID}#R003",
            ],
        )

    def test_hierarchical_range_is_rejected(self) -> None:
        requirement_ids, diagnostics = expand_coverage_value(
            f"{RECORD_ID}#R001-01-R001-03",
            "fixture.md",
        )

        self.assertEqual(requirement_ids, [])
        self.assertEqual(len(diagnostics), 1)
        self.assertEqual(diagnostics[0]["category"], "usdm_covers")


class EffectiveCoverageTests(unittest.TestCase):
    def test_direct_parent_keeps_uncovered_descendants_as_warnings(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001", "R001-01", "R001-01-01", "R001-02", "R002"],
                [
                    (
                        "spec:impl.coverage.parent",
                        [f"{RECORD_ID}#R001", f"{RECORD_ID}#R002"],
                    )
                ],
            )

            response = check_usdm_coverage(root, "sample", True)

            self.assertTrue(response["ok"])
            self.assertEqual(response["requirements"], 5)
            self.assertEqual(response["covered"], 2)
            self.assertEqual(response["direct_covered"], 2)
            self.assertEqual(response["derived_covered"], 0)
            self.assertEqual(response["uncovered"], [])
            self.assertEqual(
                response["refinement_warnings"],
                [
                    f"{RECORD_ID}#R001-01",
                    f"{RECORD_ID}#R001-01-01",
                    f"{RECORD_ID}#R001-02",
                ],
            )

    def test_complete_direct_children_derive_parent_coverage(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001", "R001-01", "R001-02"],
                [
                    (
                        "spec:impl.coverage.children",
                        [f"{RECORD_ID}#R001-01", f"{RECORD_ID}#R001-02"],
                    )
                ],
            )

            response = check_usdm_coverage(root, "sample", True)

            self.assertTrue(response["ok"])
            self.assertEqual(response["covered"], 3)
            self.assertEqual(response["direct_covered"], 2)
            self.assertEqual(response["derived_covered"], 1)
            self.assertEqual(response["uncovered"], [])
            self.assertEqual(response["refinement_warnings"], [])

    def test_missing_child_branch_is_blocking_without_direct_ancestor(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001", "R001-01", "R001-02"],
                [
                    (
                        "spec:impl.coverage.partial",
                        [f"{RECORD_ID}#R001-01"],
                    )
                ],
            )

            response = check_usdm_coverage(root, "sample", True)

            self.assertFalse(response["ok"])
            self.assertEqual(response["direct_covered"], 1)
            self.assertEqual(response["derived_covered"], 0)
            self.assertEqual(
                response["uncovered"],
                [f"{RECORD_ID}#R001", f"{RECORD_ID}#R001-02"],
            )
            self.assertEqual(response["refinement_warnings"], [])

    def test_nested_warning_reports_nearest_direct_covered_ancestor(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001", "R001-01", "R001-01-01"],
                [
                    (
                        "spec:impl.coverage.nested",
                        [f"{RECORD_ID}#R001", f"{RECORD_ID}#R001-01"],
                    )
                ],
            )

            response = usdm_covered_by(root, f"{RECORD_ID}#R001-01-01")

            self.assertTrue(response["ok"])
            self.assertEqual(response["coverage_state"], "warning")
            self.assertEqual(response["covered_by"], [])
            self.assertEqual(
                response["direct_covered_ancestor"],
                f"{RECORD_ID}#R001-01",
            )

    def test_cross_app_full_id_is_direct_coverage(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001", "R001-01"],
                [
                    (
                        "spec:impl.coverage.cross_app",
                        [f"{RECORD_ID}#R001-01"],
                    )
                ],
            )

            response = usdm_covered_by(root, f"{RECORD_ID}#R001-01")

            self.assertTrue(response["ok"])
            self.assertEqual(response["coverage_state"], "direct")
            self.assertEqual(
                response["covered_by"],
                ["spec:impl.coverage.cross_app"],
            )


class CoveredByStateTests(unittest.TestCase):
    def test_all_coverage_states_and_missing_or_malformed_null_behavior(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001", "R001-01", "R002", "R002-01", "R002-02", "R003"],
                [
                    (
                        "spec:impl.coverage.states",
                        [
                            f"{RECORD_ID}#R001",
                            f"{RECORD_ID}#R002-01",
                            f"{RECORD_ID}#R002-02",
                        ],
                    )
                ],
            )

            direct = usdm_covered_by(root, f"{RECORD_ID}#R001")
            warning = usdm_covered_by(root, f"{RECORD_ID}#R001-01")
            derived = usdm_covered_by(root, f"{RECORD_ID}#R002")
            uncovered = usdm_covered_by(root, f"{RECORD_ID}#R003")
            missing = usdm_covered_by(root, f"{RECORD_ID}#R999")
            malformed = usdm_covered_by(root, f"{RECORD_ID}#R01")

            self.assertEqual(direct["coverage_state"], "direct")
            self.assertEqual(direct["covered_by"], ["spec:impl.coverage.states"])
            self.assertIsNone(direct["direct_covered_ancestor"])

            self.assertEqual(warning["coverage_state"], "warning")
            self.assertEqual(warning["covered_by"], [])
            self.assertEqual(
                warning["direct_covered_ancestor"],
                f"{RECORD_ID}#R001",
            )

            self.assertEqual(derived["coverage_state"], "derived")
            self.assertEqual(derived["covered_by"], [])
            self.assertIsNone(derived["direct_covered_ancestor"])

            self.assertEqual(uncovered["coverage_state"], "uncovered")
            self.assertEqual(uncovered["covered_by"], [])
            self.assertIsNone(uncovered["direct_covered_ancestor"])

            for response in (missing, malformed):
                self.assertFalse(response["ok"])
                self.assertFalse(response["exists"])
                self.assertIsNone(response["coverage_state"])
                self.assertEqual(response["covered_by"], [])
                self.assertIsNone(response["direct_covered_ancestor"])

    def test_evaluation_error_returns_null_state(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(root, ["R001"], [])
            spec_root = root / "impl" / "records" / "spec"
            (spec_root / "invalid.md").write_text(
                make_spec("spec:impl.coverage.invalid", ["not-a-usdm-id"]),
                encoding="utf-8",
            )

            response = usdm_covered_by(root, f"{RECORD_ID}#R001")

            self.assertFalse(response["ok"])
            self.assertTrue(response["exists"])
            self.assertIsNone(response["coverage_state"])
            self.assertEqual(response["covered_by"], [])
            self.assertIsNone(response["direct_covered_ancestor"])
            self.assertIn(
                "usdm_covers",
                {item["category"] for item in response["diagnostics"]},
            )


class ScopeCoverageTests(unittest.TestCase):
    def test_scope_response_projects_all_effective_categories(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001", "R001-01", "R002", "R002-01", "R002-02", "R003"],
                [
                    (
                        "spec:impl.coverage.scope",
                        [
                            f"{RECORD_ID}#R001",
                            f"{RECORD_ID}#R002-01",
                            f"{RECORD_ID}#R002-02",
                        ],
                    )
                ],
            )

            response = check_usdm_scope_coverage([RECORD_ID], repo_root=root)

            self.assertFalse(response["ok"])
            self.assertEqual(response["requirements"], 6)
            self.assertEqual(response["covered_requirements"], 4)
            self.assertEqual(response["direct_covered_requirements"], 3)
            self.assertEqual(response["derived_covered_requirements"], 1)
            self.assertEqual(response["not_covered_requirements"], 1)
            self.assertEqual(response["refinement_warning_requirements"], 1)
            self.assertEqual(
                response["items"],
                [
                    {
                        "record_id": RECORD_ID,
                        "covered": {
                            "#R001": ["spec:impl.coverage.scope"],
                            "#R002-01": ["spec:impl.coverage.scope"],
                            "#R002-02": ["spec:impl.coverage.scope"],
                        },
                        "derived_covered": ["#R002"],
                        "not_covered": ["#R003"],
                        "refinement_warnings": ["#R001-01"],
                    }
                ],
            )

    def test_scope_filters_and_include_empty_records_are_additive(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001", "R001-01"],
                [("spec:impl.coverage.scope_filter", [f"{RECORD_ID}#R001"])],
            )

            warning_only = check_usdm_scope_coverage(
                [RECORD_ID],
                repo_root=root,
                include_covered=False,
                include_not_covered=False,
                include_warnings=True,
            )
            hidden_warning = check_usdm_scope_coverage(
                [RECORD_ID],
                repo_root=root,
                include_covered=False,
                include_not_covered=True,
                include_warnings=False,
                include_empty_records=True,
            )

            self.assertTrue(warning_only["ok"])
            self.assertEqual(
                warning_only["items"],
                [
                    {
                        "record_id": RECORD_ID,
                        "refinement_warnings": ["#R001-01"],
                    }
                ],
            )
            self.assertEqual(
                hidden_warning["items"],
                [{"record_id": RECORD_ID, "not_covered": []}],
            )
            self.assertEqual(hidden_warning["covered_requirements"], 1)
            self.assertEqual(hidden_warning["refinement_warning_requirements"], 1)

    def test_hierarchical_row_sorting_is_numeric_and_deterministic(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001", "R001-10", "R001-02"],
                [("spec:impl.coverage.sort", [f"{RECORD_ID}#R001"])],
            )

            response = check_usdm_scope_coverage([RECORD_ID], repo_root=root)

            self.assertTrue(response["ok"])
            self.assertEqual(
                response["items"][0]["refinement_warnings"],
                ["#R001-02", "#R001-10"],
            )

    def test_scope_propagates_malformed_coverage_diagnostic(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001"],
                [
                    (
                        "spec:impl.coverage.malformed",
                        [f"{RECORD_ID}#R001", "not-a-usdm-id"],
                    )
                ],
            )

            response = check_usdm_scope_coverage([RECORD_ID], repo_root=root)

            self.assertFalse(response["ok"])
            self.assertEqual(response["covered_requirements"], 1)
            self.assertIn(
                "usdm_covers",
                {item["category"] for item in response["diagnostics"]},
            )
            self.assertTrue(
                any(item["severity"] == "error" for item in response["diagnostics"])
            )

    def test_scope_propagates_orphan_hierarchy_diagnostic(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            write_fixture(
                root,
                ["R001-01"],
                [("spec:impl.coverage.orphan", [f"{RECORD_ID}#R001-01"])],
            )

            response = check_usdm_scope_coverage([RECORD_ID], repo_root=root)

            self.assertFalse(response["ok"])
            self.assertEqual(response["covered_requirements"], 1)
            row_parent = [
                item
                for item in response["diagnostics"]
                if item["category"] == "row_parent"
            ]
            self.assertEqual(len(row_parent), 1)
            self.assertEqual(row_parent[0]["severity"], "error")
            self.assertEqual(row_parent[0]["value"], f"{RECORD_ID}#R001-01")


if __name__ == "__main__":
    unittest.main()
