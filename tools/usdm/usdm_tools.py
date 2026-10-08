#!/usr/bin/env python3
"""Standalone MVP USDM validation and coverage tools."""

from __future__ import annotations

import argparse
import json
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import Any

try:
    from tools.usdm.similarity.usdm_loader import expand_scopes
except ModuleNotFoundError:  # pragma: no cover - CLI/server import shape
    from similarity.usdm_loader import expand_scopes


IGNORED_APP_DIRS = {
    ".git",
    ".agents",
    ".claude",
    ".codex",
    ".discarded",
    "__pycache__",
    "bin",
    "memory",
}

USDM_RECORD_ID_RE = re.compile(r"^usdm:([a-z0-9_]+)\.[a-z0-9_]+(?:\.[a-z0-9_]+)*$")
ROW_ID_RE = re.compile(r"^R\d{3}(?:-\d{2})*$")
FULL_REQUIREMENT_ID_RE = re.compile(
    r"^usdm:([a-z0-9_]+)\.[a-z0-9_]+(?:\.[a-z0-9_]+)*#R\d{3}(?:-\d{2})*$"
)
COMPACT_COVERAGE_ID_RE = re.compile(
    r"^(?P<record>usdm:[a-z0-9_]+\.[a-z0-9_]+(?:\.[a-z0-9_]+)*)#(?P<rows>.+)$"
)
ROW_REF_RE = re.compile(r"^#?(?P<row>R\d{3}(?:-\d{2})*)$")
ROW_RANGE_RE = re.compile(r"^(?P<start>#?R\d{3})-(?P<end>#?R\d{3})$")
USDM_APP_SCOPE_ID_RE = re.compile(r"^usdm:([a-z0-9_]+)$")
USDM_H1_RE = re.compile(r"^# USDM (?P<kind>index|requirement): (?P<title>.+?)\s*$")
ATX_H1_RE = re.compile(r"^#(?!#)\s+(.+?)\s*$")
ATX_H2_RE = re.compile(r"^##(?!#)\s+(.+?)\s*$")
SPEC_REF_RE = re.compile(
    r"^spec:[a-z0-9][a-z0-9_]*(?:\.[a-z0-9][a-z0-9_]*)*$"
)
REQ_H2_RE = re.compile(r"^## Requirements: (?P<title>\S(?:.*?\S)?)\s*$")
SOURCE_FIELD_RE = re.compile(r"^> source: (?P<source>.+?)\s*$")
SOURCE_FIELD_PREFIX_RE = re.compile(r"^> source:")
METADATA_SCALAR_RE = re.compile(r"^- \*\*(?P<key>[^*]+)\*\*:\s*(?P<value>.*?)\s*$")
METADATA_LIST_ITEM_RE = re.compile(r"^\s{2,}-\s+(?P<value>.*?)\s*$")
TABLE_SEPARATOR_RE = re.compile(r"^\s*:?-{3,}:?\s*$")


@dataclass(frozen=True)
class Heading:
    level: int
    title: str
    line_index: int
    raw: str


@dataclass(frozen=True)
class UsdmScan:
    diagnostics: list[dict[str, Any]]
    requirement_ids: set[str]
    usdm_records: int


@dataclass(frozen=True)
class CoverageEntry:
    requirement_id: str
    spec_ref: str
    path: str


@dataclass(frozen=True)
class CoverageScan:
    diagnostics: list[dict[str, Any]]
    entries: list[CoverageEntry]


@dataclass(frozen=True)
class CoverageEvaluation:
    direct_covered: set[str]
    derived_covered: set[str]
    blocking_uncovered: set[str]
    refinement_warnings: set[str]
    direct_covered_ancestor: dict[str, str]
    coverage_by_requirement: dict[str, set[str]]


def diagnostic(
    category: str,
    path: str,
    message: str,
    value: Any = "",
    severity: str = "error",
) -> dict[str, Any]:
    return {
        "category": category,
        "severity": severity,
        "path": path,
        "message": message,
        "value": value,
    }


def normalize_metadata_value(value: str) -> str:
    value = value.strip()
    if len(value) >= 2 and value.startswith("`") and value.endswith("`"):
        return value[1:-1].strip()
    return value


def relpath(path: Path, repo_root: Path) -> str:
    try:
        return path.relative_to(repo_root).as_posix()
    except ValueError:
        return path.as_posix()


def read_markdown(path: Path, repo_root: Path) -> tuple[list[str], dict[str, Any] | None]:
    try:
        return path.read_text(encoding="utf-8").splitlines(), None
    except OSError as exc:
        return [], diagnostic(
            "source_read",
            relpath(path, repo_root),
            "Source file cannot be read as UTF-8 text.",
            str(exc),
        )
    except UnicodeDecodeError as exc:
        return [], diagnostic(
            "source_read",
            relpath(path, repo_root),
            "Source file cannot be read as UTF-8 text.",
            str(exc),
        )


def iter_real_headings(lines: list[str]) -> list[Heading]:
    headings: list[Heading] = []
    in_fence = False
    fence_marker = ""

    for index, line in enumerate(lines):
        stripped = line.lstrip()
        if stripped.startswith("```") or stripped.startswith("~~~"):
            marker = stripped[:3]
            if not in_fence:
                in_fence = True
                fence_marker = marker
            elif marker == fence_marker:
                in_fence = False
                fence_marker = ""
            continue

        if in_fence:
            continue

        match = re.match(r"^(#{1,6})(?!#)\s+(.+?)\s*$", line)
        if match:
            headings.append(
                Heading(
                    level=len(match.group(1)),
                    title=match.group(2).strip(),
                    line_index=index,
                    raw=line,
                )
            )

    return headings


def fenced_line_mask(lines: list[str]) -> list[bool]:
    mask: list[bool] = []
    in_fence = False
    fence_marker = ""

    for line in lines:
        stripped = line.lstrip()
        is_fence_marker = stripped.startswith("```") or stripped.startswith("~~~")
        mask.append(in_fence)
        if is_fence_marker:
            marker = stripped[:3]
            if not in_fence:
                in_fence = True
                fence_marker = marker
            elif marker == fence_marker:
                in_fence = False
                fence_marker = ""

    return mask


def metadata_block_after_h1(lines: list[str], h1: Heading, headings: list[Heading]) -> list[str]:
    next_h2 = next(
        (heading.line_index for heading in headings if heading.level == 2 and heading.line_index > h1.line_index),
        len(lines),
    )
    return lines[h1.line_index + 1 : next_h2]


def parse_metadata(block: list[str]) -> dict[str, str | list[str]]:
    metadata: dict[str, str | list[str]] = {}
    current_list_key: str | None = None

    for line in block:
        scalar_match = METADATA_SCALAR_RE.match(line)
        if scalar_match:
            key = scalar_match.group("key").strip()
            value = normalize_metadata_value(scalar_match.group("value"))
            metadata[key] = value
            current_list_key = key if value == "" else None
            continue

        list_match = METADATA_LIST_ITEM_RE.match(line)
        if list_match and current_list_key:
            current = metadata.get(current_list_key)
            if not isinstance(current, list):
                current = []
                metadata[current_list_key] = current
            current.append(normalize_metadata_value(list_match.group("value")))

    return metadata


def metadata_scalar(metadata: dict[str, str | list[str]], key: str) -> str:
    value = metadata.get(key, "")
    if isinstance(value, list):
        return ""
    return value


def metadata_list(metadata: dict[str, str | list[str]], key: str) -> list[str]:
    value = metadata.get(key)
    if isinstance(value, list):
        return value
    if isinstance(value, str) and value:
        return [value]
    return []


def discover_app_namespaces(repo_root: Path, app_namespace: str | None) -> list[str]:
    if app_namespace:
        app_dir = repo_root / app_namespace
        return [app_namespace] if is_discoverable_app(app_dir) else []

    apps: list[str] = []
    for child in sorted(repo_root.iterdir(), key=lambda item: item.name):
        if not child.is_dir():
            continue
        if child.name in IGNORED_APP_DIRS or child.name.startswith("."):
            continue
        if is_discoverable_app(child):
            apps.append(child.name)
    return apps


def is_discoverable_app(app_dir: Path) -> bool:
    return (app_dir / "records" / "usdm").is_dir() or (app_dir / "records" / "spec").is_dir()


def markdown_files(root: Path) -> list[Path]:
    if not root.is_dir():
        return []
    return sorted(path for path in root.rglob("*.md") if path.is_file())


def validate_markdown_table(
    section_lines: list[str],
    path: str,
) -> tuple[list[str], list[dict[str, Any]]]:
    diagnostics: list[dict[str, Any]] = []
    row_ids: list[str] = []

    table_start: int | None = None
    for index in range(0, max(0, len(section_lines) - 1)):
        current = section_lines[index].strip()
        following = section_lines[index + 1].strip()
        if current.startswith("|") and current.endswith("|") and following.startswith("|") and following.endswith("|"):
            table_start = index
            break

    if table_start is None:
        diagnostics.append(
            diagnostic("requirement_table", path, "Requirement section must contain a Markdown table.")
        )
        return row_ids, diagnostics

    table_lines: list[str] = []
    for line in section_lines[table_start:]:
        stripped = line.strip()
        if not (stripped.startswith("|") and stripped.endswith("|")):
            break
        table_lines.append(stripped)

    if len(table_lines) < 2:
        diagnostics.append(
            diagnostic("requirement_table", path, "Requirement table must include a header and separator.")
        )
        return row_ids, diagnostics

    headers = [cell.strip().lower() for cell in table_lines[0].strip("|").split("|")]
    separator_cells = [cell.strip() for cell in table_lines[1].strip("|").split("|")]
    if not separator_cells or not all(TABLE_SEPARATOR_RE.match(cell) for cell in separator_cells):
        diagnostics.append(
            diagnostic("requirement_table", path, "Requirement table must include a Markdown separator row.")
        )

    missing_columns = [column for column in ("id", "requirement") if column not in headers]
    if missing_columns:
        diagnostics.append(
            diagnostic(
                "requirement_table",
                path,
                "Requirement table must include id and requirement columns.",
                ", ".join(missing_columns),
            )
        )
        return row_ids, diagnostics

    id_index = headers.index("id")
    requirement_index = headers.index("requirement")
    for row in table_lines[2:]:
        cells = [cell.strip() for cell in row.strip("|").split("|")]
        if len(cells) <= max(id_index, requirement_index):
            diagnostics.append(
                diagnostic("requirement_table", path, "Requirement table row has too few cells.", row)
            )
            continue
        row_id = normalize_metadata_value(cells[id_index])
        requirement = normalize_metadata_value(cells[requirement_index])
        if not row_id:
            diagnostics.append(diagnostic("row_id", path, "Requirement row is missing an id.", row))
            continue
        if not requirement:
            diagnostics.append(diagnostic("requirement_table", path, "Requirement row is missing requirement text.", row_id))
        row_ids.append(row_id)

    return row_ids, diagnostics


def scan_usdm(repo_root: Path, app_namespace: str | None) -> UsdmScan:
    diagnostics: list[dict[str, Any]] = []
    requirement_ids: set[str] = set()
    usdm_records = 0

    for app in discover_app_namespaces(repo_root, app_namespace):
        usdm_root = repo_root / app / "records" / "usdm"
        for path in markdown_files(usdm_root):
            usdm_records += 1
            path_display = relpath(path, repo_root)
            lines, read_error = read_markdown(path, repo_root)
            if read_error:
                diagnostics.append(read_error)
                continue

            headings = iter_real_headings(lines)
            h1s = [heading for heading in headings if heading.level == 1]
            if len(h1s) != 1:
                diagnostics.append(
                    diagnostic("h1", path_display, "USDM record must contain exactly one real ATX H1.", len(h1s))
                )
                continue

            h1 = h1s[0]
            h1_match = USDM_H1_RE.match(h1.raw)
            if not h1_match:
                diagnostics.append(
                    diagnostic("h1", path_display, "H1 must match '# USDM <kind>: <Title>'.", h1.raw)
                )
                continue

            h1_kind = h1_match.group("kind")
            metadata = parse_metadata(metadata_block_after_h1(lines, h1, headings))
            required_fields = ["id", "status", "date", "kind", "parent"]
            for field in required_fields:
                if not metadata_scalar(metadata, field):
                    diagnostics.append(
                        diagnostic("metadata", path_display, f"Missing required metadata field '{field}'.")
                    )

            metadata_kind = metadata_scalar(metadata, "kind")
            if metadata_kind and metadata_kind != h1_kind:
                diagnostics.append(
                    diagnostic("metadata", path_display, "Metadata kind must equal H1 USDM kind.", metadata_kind)
                )

            record_id = metadata_scalar(metadata, "id")
            record_id_match = USDM_RECORD_ID_RE.match(record_id)
            if not record_id_match or record_id_match.group(1) != app:
                diagnostics.append(
                    diagnostic(
                        "metadata_id",
                        path_display,
                        "Metadata id must match 'usdm:<app_namespace>.<path.to.topic>'.",
                        record_id,
                    )
                )
                record_id = ""

            requirement_candidates = [
                heading
                for heading in headings
                if heading.level == 2
                and (heading.title == "Requirements:" or heading.title.startswith("Requirements: "))
            ]
            requirement_headings: list[Heading] = []
            for heading in requirement_candidates:
                if REQ_H2_RE.match(heading.raw):
                    requirement_headings.append(heading)
                else:
                    diagnostics.append(
                        diagnostic(
                            "requirements_title",
                            path_display,
                            "Requirements section title must be non-empty.",
                            heading.raw,
                        )
                    )

            if h1_kind == "index" and requirement_candidates:
                diagnostics.append(
                    diagnostic(
                        "requirements_section",
                        path_display,
                        "USDM index records must not contain Requirements sections.",
                    )
                )
            if h1_kind == "requirement" and not requirement_headings:
                diagnostics.append(
                    diagnostic(
                        "requirements_section",
                        path_display,
                        "USDM requirement records must contain one or more titled Requirements sections.",
                    )
                )

            row_ids: list[str] = []
            if h1_kind == "requirement":
                section_titles: set[str] = set()
                for heading in requirement_headings:
                    req_match = REQ_H2_RE.match(heading.raw)
                    assert req_match is not None
                    section_title = req_match.group("title").strip()
                    if section_title in section_titles:
                        diagnostics.append(
                            diagnostic(
                                "requirements_title",
                                path_display,
                                "Requirements section titles must be unique within one USDM record.",
                                section_title,
                            )
                        )
                    if section_title == "literal" or SPEC_REF_RE.fullmatch(section_title):
                        diagnostics.append(
                            diagnostic(
                                "requirements_title",
                                path_display,
                                "Requirements section title must identify content and must not be a source value.",
                                section_title,
                            )
                        )
                    section_titles.add(section_title)

                    next_h2 = next(
                        (
                            later.line_index
                            for later in headings
                            if later.level == 2 and later.line_index > heading.line_index
                        ),
                        len(lines),
                    )
                    section_lines = lines[heading.line_index + 1 : next_h2]
                    source = ""
                    if not section_lines:
                        diagnostics.append(
                            diagnostic(
                                "source",
                                path_display,
                                "Requirement section must have an immediate '> source: <source>' field.",
                                section_title,
                            )
                        )
                    else:
                        source_match = SOURCE_FIELD_RE.fullmatch(section_lines[0])
                        if source_match is None:
                            diagnostics.append(
                                diagnostic(
                                    "source",
                                    path_display,
                                    "Requirement section must have an immediate '> source: <source>' field.",
                                    section_title,
                                )
                            )
                        else:
                            source = source_match.group("source").strip()
                            if source != "literal" and SPEC_REF_RE.fullmatch(source) is None:
                                diagnostics.append(
                                    diagnostic(
                                        "source",
                                        path_display,
                                        "Requirement section source must be 'literal' or a canonical 'spec:' ref.",
                                        source,
                                    )
                                )

                    additional_source_fields = [
                        line
                        for line in section_lines[1:]
                        if SOURCE_FIELD_PREFIX_RE.match(line)
                    ]
                    if additional_source_fields:
                        diagnostics.append(
                            diagnostic(
                                "source",
                                path_display,
                                "Requirement section must contain exactly one source field.",
                                section_title,
                            )
                        )

                    section_row_ids, section_diagnostics = validate_markdown_table(section_lines, path_display)
                    row_ids.extend(section_row_ids)
                    diagnostics.extend(section_diagnostics)

                if record_id:
                    valid_row_ids = [row_id for row_id in row_ids if ROW_ID_RE.fullmatch(row_id)]
                    valid_row_id_set = set(valid_row_ids)
                    for row_id in row_ids:
                        full_id = f"{record_id}#{row_id}"
                        if not ROW_ID_RE.fullmatch(row_id):
                            diagnostics.append(
                                diagnostic(
                                    "row_id",
                                    path_display,
                                    "Requirement row ID must match RNNN(-NN)*.",
                                    row_id,
                                )
                            )
                            continue
                        if "-" in row_id:
                            parent_row_id = row_id.rsplit("-", 1)[0]
                            if parent_row_id not in valid_row_id_set:
                                diagnostics.append(
                                    diagnostic(
                                        "row_parent",
                                        path_display,
                                        "Non-top-level requirement row must have its immediate parent in the same USDM requirement record.",
                                        full_id,
                                    )
                                )
                        if full_id in requirement_ids:
                            diagnostics.append(
                                diagnostic("duplicate_requirement_id", path_display, "Full requirement ID is duplicated.", full_id)
                            )
                        requirement_ids.add(full_id)

    return UsdmScan(diagnostics=diagnostics, requirement_ids=requirement_ids, usdm_records=usdm_records)


def expand_coverage_value(value: str, path: str) -> tuple[list[str], list[dict[str, Any]]]:
    if FULL_REQUIREMENT_ID_RE.match(value):
        return [value], []

    match = COMPACT_COVERAGE_ID_RE.match(value)
    if match is None:
        return [], [
            diagnostic(
                "usdm_covers",
                path,
                "usdm_covers item must be a full USDM requirement ID or compact row list.",
                value,
            )
        ]

    record_id = match.group("record")
    if not USDM_RECORD_ID_RE.match(record_id):
        return [], [
            diagnostic(
                "usdm_covers",
                path,
                "Compact usdm_covers item must start with a USDM record ID.",
                value,
            )
        ]

    requirement_ids: list[str] = []
    diagnostics: list[dict[str, Any]] = []
    for token in (part.strip() for part in match.group("rows").split(",")):
        if not token:
            diagnostics.append(
                diagnostic(
                    "usdm_covers",
                    path,
                    "Compact usdm_covers row list contains an empty row token.",
                    value,
                )
            )
            continue

        range_match = ROW_RANGE_RE.match(token)
        if range_match:
            start_number = int(range_match.group("start").lstrip("#")[1:])
            end_number = int(range_match.group("end").lstrip("#")[1:])
            if start_number > end_number:
                diagnostics.append(
                    diagnostic(
                        "usdm_covers",
                        path,
                        "Compact usdm_covers row range must be ascending.",
                        token,
                    )
                )
                continue
            requirement_ids.extend(
                f"{record_id}#R{number:03d}"
                for number in range(start_number, end_number + 1)
            )
            continue

        row_match = ROW_REF_RE.match(token)
        if row_match:
            requirement_ids.append(f"{record_id}#{row_match.group('row')}")
            continue

        diagnostics.append(
            diagnostic(
                "usdm_covers",
                path,
                "Compact usdm_covers row token must be a hierarchical row ID or a top-level RNNN-RNNN range.",
                token,
            )
        )

    return requirement_ids, diagnostics


def scan_coverage(repo_root: Path, app_namespace: str | None) -> CoverageScan:
    diagnostics: list[dict[str, Any]] = []
    entries: list[CoverageEntry] = []

    for app in discover_app_namespaces(repo_root, app_namespace):
        spec_root = repo_root / app / "records" / "spec"
        for path in markdown_files(spec_root):
            path_display = relpath(path, repo_root)
            lines, read_error = read_markdown(path, repo_root)
            if read_error:
                diagnostics.append(read_error)
                continue

            headings = iter_real_headings(lines)
            h1s = [heading for heading in headings if heading.level == 1]
            if not h1s:
                diagnostics.append(
                    diagnostic("spec_metadata", path_display, "Specification file has no H1 for metadata inspection.")
                )
                continue

            h1 = h1s[0]
            next_h2 = next(
                (
                    heading.line_index
                    for heading in headings
                    if heading.level == 2 and heading.line_index > h1.line_index
                ),
                len(lines),
            )
            metadata = parse_metadata(lines[h1.line_index + 1 : next_h2])
            fenced_lines = fenced_line_mask(lines)
            for index, line in enumerate(lines):
                if h1.line_index < index < next_h2 or fenced_lines[index]:
                    continue
                match = METADATA_SCALAR_RE.match(line)
                if match and match.group("key").strip() == "usdm_covers":
                    diagnostics.append(
                        diagnostic(
                            "usdm_covers_location",
                            path_display,
                            "usdm_covers must appear in H1-adjacent metadata for MVP tools.",
                            f"line {index + 1}",
                        )
                    )

            spec_ref = metadata_scalar(metadata, "id") or path_display
            values = metadata_list(metadata, "usdm_covers")
            if not values:
                continue

            seen: set[str] = set()
            for value in values:
                requirement_ids, value_diagnostics = expand_coverage_value(value, path_display)
                diagnostics.extend(value_diagnostics)
                for requirement_id in requirement_ids:
                    if requirement_id in seen:
                        diagnostics.append(
                            diagnostic(
                                "duplicate_coverage",
                                path_display,
                                "Duplicate usdm_covers item inside one Specification.",
                                requirement_id,
                            )
                        )
                        continue
                    seen.add(requirement_id)
                    entries.append(
                        CoverageEntry(
                            requirement_id=requirement_id,
                            spec_ref=spec_ref,
                            path=path_display,
                        )
                    )

    return CoverageScan(diagnostics=diagnostics, entries=entries)


def has_errors(diagnostics: list[dict[str, Any]]) -> bool:
    return any(item.get("severity") == "error" for item in diagnostics)


def requirement_parent_id(requirement_id: str) -> str | None:
    record_id, row_id = requirement_id.rsplit("#", 1)
    if "-" not in row_id:
        return None
    return f"{record_id}#{row_id.rsplit('-', 1)[0]}"


def evaluate_coverage(
    requirement_ids: set[str],
    coverage_entries: list[CoverageEntry],
) -> CoverageEvaluation:
    coverage_by_requirement: dict[str, set[str]] = {
        requirement_id: set() for requirement_id in requirement_ids
    }
    for entry in coverage_entries:
        if entry.requirement_id in coverage_by_requirement:
            coverage_by_requirement[entry.requirement_id].add(entry.spec_ref)

    direct_covered = {
        requirement_id
        for requirement_id, covering_refs in coverage_by_requirement.items()
        if covering_refs
    }
    children: dict[str, list[str]] = {requirement_id: [] for requirement_id in requirement_ids}
    for requirement_id in requirement_ids:
        parent_id = requirement_parent_id(requirement_id)
        if parent_id in children:
            children[parent_id].append(requirement_id)
    for child_ids in children.values():
        child_ids.sort()

    derived_covered: set[str] = set()
    effective_cache: dict[str, bool] = {}

    def effective_covered(requirement_id: str) -> bool:
        if requirement_id in effective_cache:
            return effective_cache[requirement_id]
        if requirement_id in direct_covered:
            effective_cache[requirement_id] = True
            return True

        direct_children = children[requirement_id]
        if direct_children and all(effective_covered(child_id) for child_id in direct_children):
            derived_covered.add(requirement_id)
            effective_cache[requirement_id] = True
            return True

        effective_cache[requirement_id] = False
        return False

    for requirement_id in sorted(requirement_ids):
        effective_covered(requirement_id)

    blocking_uncovered: set[str] = set()
    refinement_warnings: set[str] = set()
    direct_covered_ancestor: dict[str, str] = {}
    for requirement_id in sorted(requirement_ids):
        if effective_cache[requirement_id]:
            continue

        ancestor_id = requirement_parent_id(requirement_id)
        nearest_direct_ancestor: str | None = None
        while ancestor_id is not None:
            if ancestor_id in direct_covered:
                nearest_direct_ancestor = ancestor_id
                break
            ancestor_id = requirement_parent_id(ancestor_id)

        if nearest_direct_ancestor is None:
            blocking_uncovered.add(requirement_id)
        else:
            refinement_warnings.add(requirement_id)
            direct_covered_ancestor[requirement_id] = nearest_direct_ancestor

    return CoverageEvaluation(
        direct_covered=direct_covered,
        derived_covered=derived_covered,
        blocking_uncovered=blocking_uncovered,
        refinement_warnings=refinement_warnings,
        direct_covered_ancestor=direct_covered_ancestor,
        coverage_by_requirement=coverage_by_requirement,
    )


def validate_usdm(repo_root: Path, app_namespace: str | None) -> dict[str, Any]:
    scan = scan_usdm(repo_root, app_namespace)
    return {
        "ok": not has_errors(scan.diagnostics),
        "usdm_records": scan.usdm_records,
        "requirements": len(scan.requirement_ids),
        "diagnostics": scan.diagnostics,
    }


def check_usdm_coverage(repo_root: Path, app_namespace: str | None, include_dangling: bool) -> dict[str, Any]:
    usdm_scan = scan_usdm(repo_root, app_namespace)
    all_usdm_scan = scan_usdm(repo_root, None)
    coverage_scan = scan_coverage(repo_root, None)
    requirement_ids = usdm_scan.requirement_ids
    all_requirement_ids = all_usdm_scan.requirement_ids
    evaluation = evaluate_coverage(all_requirement_ids, coverage_scan.entries)

    direct_covered_ids = evaluation.direct_covered & requirement_ids
    derived_covered_ids = evaluation.derived_covered & requirement_ids
    covered_ids = direct_covered_ids | derived_covered_ids
    uncovered = sorted(evaluation.blocking_uncovered & requirement_ids)
    refinement_warnings = sorted(evaluation.refinement_warnings & requirement_ids)
    dangling_entries = [
        {
            "requirement_id": entry.requirement_id,
            "spec_ref": entry.spec_ref,
            "path": entry.path,
        }
        for entry in coverage_scan.entries
        if entry.requirement_id not in all_requirement_ids
        and (
            app_namespace is None
            or entry.requirement_id.startswith(f"usdm:{app_namespace}.")
        )
    ]

    diagnostics = [*usdm_scan.diagnostics, *coverage_scan.diagnostics]
    for entry in dangling_entries:
        diagnostics.append(
            diagnostic(
                "dangling",
                entry["path"],
                "usdm_covers item points to a missing USDM requirement ID.",
                entry["requirement_id"],
            )
        )

    return {
        "ok": not has_errors(diagnostics) and not uncovered and not dangling_entries,
        "requirements": len(requirement_ids),
        "covered": len(covered_ids),
        "direct_covered": len(direct_covered_ids),
        "derived_covered": len(derived_covered_ids),
        "uncovered": uncovered,
        "refinement_warnings": refinement_warnings,
        "dangling": dangling_entries if include_dangling else [],
        "diagnostics": diagnostics,
    }


def usdm_covered_by(repo_root: Path, requirement_id: str) -> dict[str, Any]:
    if not FULL_REQUIREMENT_ID_RE.match(requirement_id):
        return {
            "ok": False,
            "requirement_id": requirement_id,
            "exists": False,
            "coverage_state": None,
            "covered_by": [],
            "direct_covered_ancestor": None,
            "diagnostics": [
                diagnostic(
                    "requirement_id",
                    "",
                    "Requested requirement ID must be a full USDM requirement ID.",
                    requirement_id,
                )
            ],
        }

    app_namespace = requirement_id.removeprefix("usdm:").split(".", 1)[0]
    usdm_scan = scan_usdm(repo_root, app_namespace)
    coverage_scan = scan_coverage(repo_root, None)
    exists = requirement_id in usdm_scan.requirement_ids
    diagnostics = [*usdm_scan.diagnostics, *coverage_scan.diagnostics]
    if not exists:
        diagnostics.append(
            diagnostic("missing_requirement", "", "Requested requirement ID was not found.", requirement_id)
        )
        return {
            "ok": False,
            "requirement_id": requirement_id,
            "exists": False,
            "coverage_state": None,
            "covered_by": [],
            "direct_covered_ancestor": None,
            "diagnostics": diagnostics,
        }

    if has_errors(diagnostics):
        return {
            "ok": False,
            "requirement_id": requirement_id,
            "exists": True,
            "coverage_state": None,
            "covered_by": [],
            "direct_covered_ancestor": None,
            "diagnostics": diagnostics,
        }

    evaluation = evaluate_coverage(usdm_scan.requirement_ids, coverage_scan.entries)
    if requirement_id in evaluation.direct_covered:
        coverage_state = "direct"
        covered_by = sorted(evaluation.coverage_by_requirement[requirement_id])
        direct_covered_ancestor = None
    elif requirement_id in evaluation.derived_covered:
        coverage_state = "derived"
        covered_by = []
        direct_covered_ancestor = None
    elif requirement_id in evaluation.refinement_warnings:
        coverage_state = "warning"
        covered_by = []
        direct_covered_ancestor = evaluation.direct_covered_ancestor[requirement_id]
    else:
        coverage_state = "uncovered"
        covered_by = []
        direct_covered_ancestor = None

    return {
        "ok": True,
        "requirement_id": requirement_id,
        "exists": True,
        "coverage_state": coverage_state,
        "covered_by": covered_by,
        "direct_covered_ancestor": direct_covered_ancestor,
        "diagnostics": diagnostics,
    }


def row_id_from_requirement(requirement_id: str) -> str:
    return f"#{requirement_id.rsplit('#', 1)[1]}"


def row_sort_key(row_id: str) -> tuple[int, ...]:
    normalized = row_id.removeprefix("#")
    if not ROW_ID_RE.fullmatch(normalized):
        return (sys.maxsize,)
    top_level, *child_segments = normalized.removeprefix("R").split("-")
    return tuple(int(segment) for segment in (top_level, *child_segments))


def requirement_matches_scope(requirement_id: str, scope_id: str) -> bool:
    if FULL_REQUIREMENT_ID_RE.match(scope_id):
        return requirement_id == scope_id
    if USDM_APP_SCOPE_ID_RE.match(scope_id):
        return requirement_id.startswith(f"{scope_id}.")
    if USDM_RECORD_ID_RE.match(scope_id):
        record_id = requirement_id.split("#", 1)[0]
        return record_id == scope_id or record_id.startswith(f"{scope_id}.")
    return False


def check_usdm_scope_coverage(
    scope_ids: list[str],
    repo_root: Path | str | None = None,
    include_covered: bool = True,
    include_not_covered: bool = True,
    include_warnings: bool = True,
    include_empty_records: bool = False,
) -> dict[str, Any]:
    empty_counts = {
        "records": 0,
        "requirements": 0,
        "covered_requirements": 0,
        "direct_covered_requirements": 0,
        "derived_covered_requirements": 0,
        "not_covered_requirements": 0,
        "refinement_warning_requirements": 0,
    }
    if not repo_root:
        return {
            "ok": False,
            "scope_ids": scope_ids,
            **empty_counts,
            "items": [],
            "diagnostics": [
                diagnostic("repo_root", "", "repo_root is missing or unreadable.")
            ],
        }

    root = Path(repo_root).resolve()
    root_error = repo_root_error(root)
    if root_error:
        return {
            "ok": False,
            "scope_ids": scope_ids,
            **empty_counts,
            "items": [],
            "diagnostics": root_error["diagnostics"],
        }

    if not scope_ids:
        return {
            "ok": False,
            "scope_ids": scope_ids,
            **empty_counts,
            "items": [],
            "diagnostics": [
                diagnostic(
                    "scope_ids",
                    "",
                    "scope_ids must include at least one USDM app, topic, record, or requirement ID.",
                )
            ],
        }

    expansion = expand_scopes(root, scope_ids, "coverage")
    selected_requirements = expansion.requirements
    selected_ids = {row.requirement_id for row in selected_requirements}
    selected_records = sorted({row.usdm_id for row in selected_requirements})

    coverage_scan = scan_coverage(root, None)
    all_usdm_scan = scan_usdm(root, None)
    all_requirements = all_usdm_scan.requirement_ids
    evaluation = evaluate_coverage(all_requirements, coverage_scan.entries)

    direct_covered_ids = evaluation.direct_covered & selected_ids
    derived_covered_ids = evaluation.derived_covered & selected_ids
    covered_ids = direct_covered_ids | derived_covered_ids
    not_covered_ids = evaluation.blocking_uncovered & selected_ids
    warning_ids = evaluation.refinement_warnings & selected_ids

    selected_paths = {row.path for row in selected_requirements}
    selected_app_scopes = {
        match.group(1)
        for scope_id in scope_ids
        if (match := USDM_APP_SCOPE_ID_RE.match(scope_id)) is not None
    }
    diagnostics = list(expansion.diagnostics)
    diagnostics.extend(
        item
        for item in all_usdm_scan.diagnostics
        if item.get("path") in selected_paths
        or any(
            item.get("path", "").startswith(f"{app}/records/usdm/")
            for app in selected_app_scopes
        )
    )
    diagnostics.extend(coverage_scan.diagnostics)
    for entry in coverage_scan.entries:
        if entry.requirement_id in all_requirements:
            continue
        if not any(
            requirement_matches_scope(entry.requirement_id, scope_id)
            for scope_id in scope_ids
        ):
            continue
        diagnostics.append(
            diagnostic(
                "dangling",
                entry.path,
                "usdm_covers item points to a missing USDM requirement ID.",
                entry.requirement_id,
            )
        )

    items: list[dict[str, Any]] = []
    for record_id in selected_records:
        record_rows = [
            row
            for row in selected_requirements
            if row.usdm_id == record_id
        ]
        item: dict[str, Any] = {"record_id": record_id}
        covered: dict[str, list[str]] = {}
        derived_covered: list[str] = []
        not_covered: list[str] = []
        refinement_warnings: list[str] = []

        for row in sorted(
            record_rows,
            key=lambda row: row_sort_key(row_id_from_requirement(row.requirement_id)),
        ):
            requirement_id = row.requirement_id
            row_id = row_id_from_requirement(requirement_id)
            if requirement_id in direct_covered_ids:
                covered[row_id] = sorted(evaluation.coverage_by_requirement[requirement_id])
            elif requirement_id in derived_covered_ids:
                derived_covered.append(row_id)
            elif requirement_id in warning_ids:
                refinement_warnings.append(row_id)
            elif requirement_id in not_covered_ids:
                not_covered.append(row_id)

        if include_covered and covered:
            item["covered"] = covered
        if include_covered and derived_covered:
            item["derived_covered"] = derived_covered
        if include_not_covered and not_covered:
            item["not_covered"] = not_covered
        if include_warnings and refinement_warnings:
            item["refinement_warnings"] = refinement_warnings

        visible_fields = set(item) - {"record_id"}
        if not visible_fields and include_empty_records:
            if include_covered:
                item["covered"] = {}
                item["derived_covered"] = []
            if include_not_covered:
                item["not_covered"] = []
            if include_warnings:
                item["refinement_warnings"] = []
            visible_fields = set(item) - {"record_id"}
        if visible_fields or include_empty_records:
            items.append(item)

    return {
        "ok": not has_errors(diagnostics) and not not_covered_ids,
        "scope_ids": scope_ids,
        "records": len(selected_records),
        "requirements": len(selected_ids),
        "covered_requirements": len(covered_ids),
        "direct_covered_requirements": len(direct_covered_ids),
        "derived_covered_requirements": len(derived_covered_ids),
        "not_covered_requirements": len(not_covered_ids),
        "refinement_warning_requirements": len(warning_ids),
        "items": sorted(items, key=lambda item: item["record_id"]),
        "diagnostics": diagnostics,
    }


def existing_repo_root(value: str) -> Path:
    return Path(value).resolve()


def repo_root_error(repo_root: Path | None) -> dict[str, Any] | None:
    if repo_root is None or not repo_root.is_dir():
        return {
            "ok": False,
            "diagnostics": [
                diagnostic(
                    "repo_root",
                    "",
                    "repo_root is missing or unreadable.",
                    "" if repo_root is None else str(repo_root),
                )
            ],
        }
    return None


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="Standalone MVP USDM tools.")
    subparsers = parser.add_subparsers(dest="command", required=True)

    validate_parser = subparsers.add_parser("validate_usdm")
    validate_parser.add_argument("--repo-root", type=existing_repo_root)
    validate_parser.add_argument("--app-namespace")

    coverage_parser = subparsers.add_parser("check_usdm_coverage")
    coverage_parser.add_argument("--repo-root", type=existing_repo_root)
    coverage_parser.add_argument("--app-namespace")
    coverage_parser.add_argument("--include-dangling", dest="include_dangling", action="store_true", default=True)
    coverage_parser.add_argument("--no-include-dangling", dest="include_dangling", action="store_false")

    covered_by_parser = subparsers.add_parser("usdm_covered_by")
    covered_by_parser.add_argument("--repo-root", type=existing_repo_root)
    covered_by_parser.add_argument("--requirement-id", required=True)

    scope_coverage_parser = subparsers.add_parser("check_usdm_scope_coverage")
    scope_coverage_parser.add_argument("--repo-root", type=existing_repo_root)
    scope_coverage_parser.add_argument(
        "--scope-id",
        dest="scope_ids",
        action="append",
        required=True,
    )
    scope_coverage_parser.add_argument(
        "--include-covered",
        dest="include_covered",
        action="store_true",
        default=True,
    )
    scope_coverage_parser.add_argument(
        "--no-include-covered",
        dest="include_covered",
        action="store_false",
    )
    scope_coverage_parser.add_argument(
        "--include-not-covered",
        dest="include_not_covered",
        action="store_true",
        default=True,
    )
    scope_coverage_parser.add_argument(
        "--no-include-not-covered",
        dest="include_not_covered",
        action="store_false",
    )
    scope_coverage_parser.add_argument(
        "--include-warnings",
        dest="include_warnings",
        action="store_true",
        default=True,
    )
    scope_coverage_parser.add_argument(
        "--no-include-warnings",
        dest="include_warnings",
        action="store_false",
    )
    scope_coverage_parser.add_argument("--include-empty-records", action="store_true", default=False)

    return parser


def main(argv: list[str] | None = None) -> int:
    parser = build_parser()
    args = parser.parse_args(argv)
    root_error = repo_root_error(args.repo_root)
    if root_error:
        print(json.dumps(root_error, indent=2, sort_keys=True))
        return 1

    try:
        if args.command == "validate_usdm":
            response = validate_usdm(args.repo_root, args.app_namespace)
        elif args.command == "check_usdm_coverage":
            response = check_usdm_coverage(args.repo_root, args.app_namespace, args.include_dangling)
        elif args.command == "usdm_covered_by":
            response = usdm_covered_by(args.repo_root, args.requirement_id)
        elif args.command == "check_usdm_scope_coverage":
            response = check_usdm_scope_coverage(
                scope_ids=args.scope_ids,
                repo_root=args.repo_root,
                include_covered=args.include_covered,
                include_not_covered=args.include_not_covered,
                include_warnings=args.include_warnings,
                include_empty_records=args.include_empty_records,
            )
        else:
            parser.error(f"unknown command: {args.command}")
    except Exception as exc:  # pragma: no cover - last-resort CLI failure envelope
        response = {
            "ok": False,
            "diagnostics": [
                diagnostic("execution", "", "USDM tool execution failed.", f"{type(exc).__name__}: {exc}")
            ],
        }

    print(json.dumps(response, indent=2, sort_keys=True))
    return 0 if response.get("ok") is True else 1


if __name__ == "__main__":
    sys.exit(main())
