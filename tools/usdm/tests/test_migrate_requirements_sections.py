from __future__ import annotations

from pathlib import Path
import sys

import pytest


TOOLS_USDM_DIR = Path(__file__).resolve().parents[1]
if str(TOOLS_USDM_DIR) not in sys.path:
    sys.path.insert(0, str(TOOLS_USDM_DIR))

from migrate_requirements_sections import (  # noqa: E402
    MigrationError,
    SpecTitle,
    migrate_usdm_text,
)


def test_migrate_legacy_requirement_heading() -> None:
    source_ref = "spec:product.design_records.spec_format.document_shape"
    original = (
        "# USDM requirement: Document shape\n\n"
        "## Requirements: " + source_ref + "\n\n"
        "| id | requirement |\n"
        "|---|---|\n"
        "| R001 | Preserve this row. |\n"
    )

    updated, count = migrate_usdm_text(
        path=Path("document-shape.md"),
        text=original,
        spec_index={
            source_ref: SpecTitle(
                ref=source_ref,
                title="Spec document shape",
                path=Path("product/records/spec/design-records/spec-format/document-shape.md"),
            )
        },
    )

    assert count == 1
    assert (
        "## Requirements: Spec document shape\n"
        f"> source: {source_ref}\n\n"
    ) in updated
    assert "| R001 | Preserve this row. |" in updated


def test_reject_duplicate_resulting_title() -> None:
    first_ref = "spec:product.first"
    second_ref = "spec:product.second"
    original = (
        f"## Requirements: {first_ref}\n\n"
        "| id | requirement |\n|---|---|\n| R001 | First. |\n\n"
        f"## Requirements: {second_ref}\n\n"
        "| id | requirement |\n|---|---|\n| R002 | Second. |\n"
    )
    shared_title = "Shared title"

    with pytest.raises(MigrationError, match="duplicate Requirements title"):
        migrate_usdm_text(
            path=Path("duplicate.md"),
            text=original,
            spec_index={
                first_ref: SpecTitle(first_ref, shared_title, Path("first.md")),
                second_ref: SpecTitle(second_ref, shared_title, Path("second.md")),
            },
        )


def test_reject_existing_source_field_after_legacy_heading() -> None:
    source_ref = "spec:product.example"
    original = (
        f"## Requirements: {source_ref}\n"
        f"> source: {source_ref}\n"
    )

    with pytest.raises(MigrationError, match="existing source field"):
        migrate_usdm_text(
            path=Path("partial.md"),
            text=original,
            spec_index={
                source_ref: SpecTitle(source_ref, "Example", Path("example.md"))
            },
        )


def test_preserve_no_final_newline() -> None:
    source_ref = "spec:product.example"
    original = f"## Requirements: {source_ref}"

    updated, count = migrate_usdm_text(
        path=Path("no-final-newline.md"),
        text=original,
        spec_index={
            source_ref: SpecTitle(source_ref, "Example", Path("example.md"))
        },
    )

    assert count == 1
    assert updated == (
        "## Requirements: Example\n"
        f"> source: {source_ref}"
    )
