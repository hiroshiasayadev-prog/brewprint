"""USDM requirement scope discovery and expansion."""

from __future__ import annotations

import re
from dataclasses import dataclass
from pathlib import Path
from typing import Any

from .models import RequirementRow
from .text_normalizer import normalize_requirement_detail


USDM_APP_SCOPE_ID_RE = re.compile(r"^usdm:(?P<app>[a-z0-9_]+)$")
USDM_RECORD_ID_RE = re.compile(
    r"^usdm:(?P<app>[a-z0-9_]+)\.[a-z0-9_]+(?:\.[a-z0-9_]+)*$"
)
FULL_REQUIREMENT_ID_RE = re.compile(
    r"^(?P<record>usdm:(?P<app>[a-z0-9_]+)\.[a-z0-9_]+"
    r"(?:\.[a-z0-9_]+)*)#(?P<row>R\d{3})$"
)
USDM_REQUIREMENT_H1_RE = re.compile(r"^# USDM requirement: .+?\s*$")
REQ_H2_RE = re.compile(r"^## Requirements: spec:.+?\s*$")
METADATA_ID_RE = re.compile(r"^- \*\*id\*\*:\s*(?P<value>.*?)\s*$")
TABLE_SEPARATOR_RE = re.compile(r"^\s*:?-{3,}:?\s*$")


@dataclass(frozen=True)
class ScopeExpansion:
    requirements: list[RequirementRow]
    diagnostics: list[dict[str, Any]]


def diagnostic(
    category: str,
    message: str,
    value: Any = "",
    path: str = "",
    severity: str = "error",
) -> dict[str, Any]:
    return {
        "category": category,
        "severity": severity,
        "path": path,
        "message": message,
        "value": value,
    }


def expand_scopes(
    repo_root: Path,
    scope_ids: list[str],
    scope_role: str,
) -> ScopeExpansion:
    """Resolve app, record, topic, and requirement IDs into rows."""
    requested_by_app: dict[str, list[str]] = {}
    diagnostics: list[dict[str, Any]] = []

    for scope_id in scope_ids:
        match = (
            FULL_REQUIREMENT_ID_RE.match(scope_id)
            or USDM_RECORD_ID_RE.match(scope_id)
            or USDM_APP_SCOPE_ID_RE.match(scope_id)
        )
        if match is None:
            diagnostics.append(
                diagnostic(
                    "scope_id",
                    f"{scope_role} scope ID is not a USDM app, record, topic, or requirement ID.",
                    scope_id,
                )
            )
            continue
        requested_by_app.setdefault(match.group("app"), []).append(scope_id)

    selected: dict[str, RequirementRow] = {}
    for app, app_scope_ids in requested_by_app.items():
        rows = _load_app_requirements(repo_root, app)
        rows_by_id = {row.requirement_id: row for row in rows}
        rows_by_record: dict[str, list[RequirementRow]] = {}
        for row in rows:
            rows_by_record.setdefault(row.usdm_id, []).append(row)

        for scope_id in app_scope_ids:
            full_match = FULL_REQUIREMENT_ID_RE.match(scope_id)
            if full_match:
                row = rows_by_id.get(scope_id)
                if row is None:
                    diagnostics.append(
                        diagnostic(
                            "scope_resolution",
                            f"{scope_role} requirement scope could not be resolved.",
                            scope_id,
                        )
                    )
                else:
                    selected[row.requirement_id] = row
                continue

            if USDM_APP_SCOPE_ID_RE.match(scope_id):
                matching_rows = rows
            else:
                matching_rows = rows_by_record.get(scope_id)
            if matching_rows is None:
                topic_prefix = f"{scope_id}."
                matching_rows = [
                    row for row in rows if row.usdm_id.startswith(topic_prefix)
                ]
            if not matching_rows:
                diagnostics.append(
                    diagnostic(
                        "scope_resolution",
                        f"{scope_role} USDM app, record, or topic scope could not be "
                        "resolved to requirement rows.",
                        scope_id,
                    )
                )
                continue
            for row in matching_rows:
                selected[row.requirement_id] = row

    return ScopeExpansion(
        requirements=sorted(selected.values(), key=lambda item: item.requirement_id),
        diagnostics=diagnostics,
    )


def _load_app_requirements(repo_root: Path, app: str) -> list[RequirementRow]:
    usdm_root = repo_root / app / "records" / "usdm"
    if not usdm_root.is_dir():
        return []

    rows: list[RequirementRow] = []
    for path in sorted(item for item in usdm_root.rglob("*.md") if item.is_file()):
        rows.extend(_parse_requirement_record(path, repo_root, app))
    return rows


def _parse_requirement_record(
    path: Path,
    repo_root: Path,
    app: str,
) -> list[RequirementRow]:
    try:
        lines = path.read_text(encoding="utf-8").splitlines()
    except (OSError, UnicodeDecodeError):
        return []

    visible = _visible_line_mask(lines)
    h1_indexes = [
        index
        for index, line in enumerate(lines)
        if visible[index] and line.startswith("# ") and not line.startswith("## ")
    ]
    if len(h1_indexes) != 1:
        return []
    h1_index = h1_indexes[0]
    if not USDM_REQUIREMENT_H1_RE.match(lines[h1_index]):
        return []

    record_id = ""
    first_h2 = next(
        (
            index
            for index in range(h1_index + 1, len(lines))
            if visible[index] and lines[index].startswith("## ")
        ),
        len(lines),
    )
    for line in lines[h1_index + 1 : first_h2]:
        match = METADATA_ID_RE.match(line)
        if match:
            record_id = _unquote_code(match.group("value"))
            break
    record_match = USDM_RECORD_ID_RE.match(record_id)
    if record_match is None or record_match.group("app") != app:
        return []

    relative_path = path.relative_to(repo_root).as_posix()
    rows: list[RequirementRow] = []
    h2_indexes = [
        index
        for index, line in enumerate(lines)
        if visible[index] and line.startswith("## ")
    ]
    for position, h2_index in enumerate(h2_indexes):
        if not REQ_H2_RE.match(lines[h2_index]):
            continue
        section_end = (
            h2_indexes[position + 1] if position + 1 < len(h2_indexes) else len(lines)
        )
        rows.extend(
            _parse_requirement_table(
                lines[h2_index + 1 : section_end],
                record_id,
                relative_path,
            )
        )
    return rows


def _parse_requirement_table(
    lines: list[str],
    record_id: str,
    relative_path: str,
) -> list[RequirementRow]:
    for index in range(max(0, len(lines) - 1)):
        header = lines[index].strip()
        separator = lines[index + 1].strip()
        if not (_is_table_row(header) and _is_table_row(separator)):
            continue
        headers = [cell.strip().lower() for cell in _split_table_row(header)]
        separators = _split_table_row(separator)
        if (
            "id" not in headers
            or "requirement" not in headers
            or len(headers) != len(separators)
            or not all(TABLE_SEPARATOR_RE.match(cell) for cell in separators)
        ):
            continue

        id_index = headers.index("id")
        requirement_index = headers.index("requirement")
        result: list[RequirementRow] = []
        for row_line in lines[index + 2 :]:
            if not _is_table_row(row_line.strip()):
                break
            cells = _split_table_row(row_line.strip())
            if len(cells) <= max(id_index, requirement_index):
                continue
            row_id = _unquote_code(cells[id_index].strip())
            detail = normalize_requirement_detail(cells[requirement_index])
            full_id = f"{record_id}#{row_id}"
            if FULL_REQUIREMENT_ID_RE.match(full_id) and detail:
                result.append(
                    RequirementRow(
                        requirement_id=full_id,
                        detail=detail,
                        usdm_id=record_id,
                        path=relative_path,
                    )
                )
        return result
    return []


def _visible_line_mask(lines: list[str]) -> list[bool]:
    result: list[bool] = []
    fence_marker = ""
    for line in lines:
        stripped = line.lstrip()
        marker = stripped[:3]
        if not fence_marker and marker in {"```", "~~~"}:
            fence_marker = marker
            result.append(False)
            continue
        if fence_marker:
            result.append(False)
            if stripped.startswith(fence_marker):
                fence_marker = ""
            continue
        result.append(True)
    return result


def _is_table_row(line: str) -> bool:
    return line.startswith("|") and line.endswith("|")


def _split_table_row(line: str) -> list[str]:
    """Split a Markdown table row while preserving escaped pipes and code spans."""
    cells: list[str] = []
    current: list[str] = []
    escaped = False
    in_code = False
    for char in line[1:-1]:
        if escaped:
            current.extend(("\\", char))
            escaped = False
        elif char == "\\":
            escaped = True
        elif char == "`":
            in_code = not in_code
            current.append(char)
        elif char == "|" and not in_code:
            cells.append("".join(current).strip())
            current = []
        else:
            current.append(char)
    if escaped:
        current.append("\\")
    cells.append("".join(current).strip())
    return cells


def _unquote_code(value: str) -> str:
    value = value.strip()
    if len(value) >= 2 and value.startswith("`") and value.endswith("`"):
        return value[1:-1].strip()
    return value
