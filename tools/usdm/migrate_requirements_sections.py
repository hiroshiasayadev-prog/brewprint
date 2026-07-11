from __future__ import annotations

import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path


LEGACY_REQUIREMENTS_HEADING_RE = re.compile(
    r"^## Requirements: (?P<source>spec:[a-z0-9][a-z0-9_.]*)\s*$"
)
REQUIREMENTS_HEADING_RE = re.compile(r"^## Requirements: (?P<title>.+?)\s*$")
SPEC_H1_RE = re.compile(
    r"^# (?P<kind>Overview|Index|Concept|Reference|Contract): (?P<title>.+?)\s*$"
)
SPEC_ID_RE = re.compile(r"^- \*\*id\*\*: `(?P<ref>spec:[a-z0-9][a-z0-9_.]*)`\s*$")
NEW_SOURCE_RE = re.compile(r"^> source: (?:literal|spec:[a-z0-9][a-z0-9_.]*)\s*$")


@dataclass(frozen=True)
class SpecTitle:
    ref: str
    title: str
    path: Path


@dataclass(frozen=True)
class PlannedChange:
    path: Path
    updated: str
    section_count: int


class MigrationError(RuntimeError):
    pass


def canonical_spec_ref(spec_path: Path, repository_root: Path) -> str:
    try:
        relative = spec_path.relative_to(repository_root)
    except ValueError as exc:
        raise MigrationError(f"spec path is outside repository root: {spec_path}") from exc

    parts = relative.parts
    try:
        records_index = parts.index("records")
    except ValueError as exc:
        raise MigrationError(f"spec path does not contain records/: {spec_path}") from exc

    if records_index + 1 >= len(parts) or parts[records_index + 1] != "spec":
        raise MigrationError(f"spec path is not under records/spec/: {spec_path}")

    app_namespace = parts[0].lower()
    topic_parts = list(parts[records_index + 2 :])
    if not topic_parts:
        raise MigrationError(f"spec path has no topic segment: {spec_path}")

    final = Path(topic_parts[-1])
    if final.suffix.lower() != ".md":
        raise MigrationError(f"spec path is not Markdown: {spec_path}")

    stem = final.stem
    topic_parts[-1] = stem
    if stem == "index":
        topic_parts.pop()

    normalized = [part.replace("-", "_") for part in topic_parts]
    suffix = ".".join(normalized)
    return f"spec:{app_namespace}" + (f".{suffix}" if suffix else "")


def read_spec_title(spec_path: Path, repository_root: Path) -> SpecTitle:
    text = spec_path.read_text(encoding="utf-8")
    lines = text.splitlines()

    if not lines:
        raise MigrationError(f"empty spec file: {spec_path}")

    h1_match = SPEC_H1_RE.fullmatch(lines[0])
    if not h1_match:
        raise MigrationError(
            f"spec H1 does not match an accepted '# <SpecKind>: <Title>' form: "
            f"{spec_path}: {lines[0]!r}"
        )

    metadata_ref: str | None = None
    for line in lines[1:]:
        if line.startswith("## "):
            break
        match = SPEC_ID_RE.fullmatch(line)
        if match:
            metadata_ref = match.group("ref")
            break

    if metadata_ref is None:
        raise MigrationError(f"spec H1-adjacent id is missing: {spec_path}")

    derived_ref = canonical_spec_ref(spec_path, repository_root)
    if metadata_ref != derived_ref:
        raise MigrationError(
            f"spec id does not match path-derived ref: {spec_path}: "
            f"metadata={metadata_ref!r}, derived={derived_ref!r}"
        )

    return SpecTitle(ref=metadata_ref, title=h1_match.group("title"), path=spec_path)


def collect_legacy_refs(target_root: Path) -> set[str]:
    refs: set[str] = set()
    for path in sorted(target_root.rglob("*.md")):
        for line in path.read_text(encoding="utf-8").splitlines():
            match = LEGACY_REQUIREMENTS_HEADING_RE.fullmatch(line)
            if match:
                refs.add(match.group("source"))
    return refs


def build_spec_index(
    *,
    repository_root: Path,
    required_refs: set[str],
) -> dict[str, SpecTitle]:
    index: dict[str, SpecTitle] = {}

    for spec_root in sorted(repository_root.glob("*/records/spec")):
        if not spec_root.is_dir():
            continue
        for spec_path in sorted(spec_root.rglob("*.md")):
            text = spec_path.read_text(encoding="utf-8")
            metadata_ref: str | None = None
            for line in text.splitlines()[1:]:
                if line.startswith("## "):
                    break
                match = SPEC_ID_RE.fullmatch(line)
                if match:
                    metadata_ref = match.group("ref")
                    break

            if metadata_ref not in required_refs:
                continue

            spec = read_spec_title(spec_path, repository_root)
            previous = index.get(spec.ref)
            if previous is not None:
                raise MigrationError(
                    f"duplicate canonical spec ref {spec.ref!r}: "
                    f"{previous.path} and {spec.path}"
                )
            index[spec.ref] = spec

    missing_refs = sorted(required_refs - index.keys())
    if missing_refs:
        raise MigrationError(
            "unresolved spec ref(s): " + ", ".join(missing_refs)
        )

    return index


def migrate_usdm_text(
    *,
    path: Path,
    text: str,
    spec_index: dict[str, SpecTitle],
) -> tuple[str, int]:
    lines = text.splitlines(keepends=True)
    existing_titles: set[str] = set()

    for line in lines:
        line_body = line.rstrip("\r\n")
        heading_match = REQUIREMENTS_HEADING_RE.fullmatch(line_body)
        if not heading_match:
            continue
        if LEGACY_REQUIREMENTS_HEADING_RE.fullmatch(line_body):
            continue
        title = heading_match.group("title")
        if title in existing_titles:
            raise MigrationError(f"duplicate Requirements title in {path}: {title!r}")
        existing_titles.add(title)

    output: list[str] = []
    migrated_count = 0

    for index, line in enumerate(lines):
        line_body = line.rstrip("\r\n")
        line_ending = line[len(line_body) :]
        match = LEGACY_REQUIREMENTS_HEADING_RE.fullmatch(line_body)
        if not match:
            output.append(line)
            continue

        source_ref = match.group("source")
        spec = spec_index[source_ref]

        if spec.title in existing_titles:
            raise MigrationError(
                f"duplicate Requirements title in {path}: {spec.title!r}"
            )
        existing_titles.add(spec.title)

        if index + 1 < len(lines):
            next_line = lines[index + 1].rstrip("\r\n")
            if NEW_SOURCE_RE.fullmatch(next_line):
                raise MigrationError(
                    f"existing source field follows legacy heading in {path}: "
                    f"line {index + 2}"
                )

        heading_ending = line_ending or "\n"
        output.append(f"## Requirements: {spec.title}{heading_ending}")
        output.append(f"> source: {source_ref}{line_ending}")
        migrated_count += 1

    return "".join(output), migrated_count


def plan_changes(
    *,
    repository_root: Path,
    target_root: Path,
) -> list[PlannedChange]:
    required_refs = collect_legacy_refs(target_root)
    spec_index = build_spec_index(
        repository_root=repository_root,
        required_refs=required_refs,
    )
    changes: list[PlannedChange] = []

    for path in sorted(target_root.rglob("*.md")):
        original = path.read_text(encoding="utf-8")
        updated, section_count = migrate_usdm_text(
            path=path,
            text=original,
            spec_index=spec_index,
        )
        if section_count == 0:
            continue
        changes.append(
            PlannedChange(
                path=path,
                updated=updated,
                section_count=section_count,
            )
        )

    return changes


def parse_args(argv: list[str]) -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description=(
            "Migrate legacy '## Requirements: spec:...' sections to "
            "'## Requirements: <Spec H1 title>' plus '> source: <spec ref>'."
        )
    )
    parser.add_argument(
        "--repo-root",
        type=Path,
        default=Path(__file__).resolve().parents[2],
        help="Repository root. Defaults to the root containing tools/usdm/.",
    )
    parser.add_argument(
        "--target",
        type=Path,
        default=Path("product/records/usdm/design-records"),
        help="USDM directory relative to --repo-root unless absolute.",
    )
    parser.add_argument(
        "--write",
        action="store_true",
        help="Apply changes. Without this flag, the command performs a dry run.",
    )
    return parser.parse_args(argv)


def main(argv: list[str] | None = None) -> int:
    args = parse_args(sys.argv[1:] if argv is None else argv)
    repository_root = args.repo_root.resolve()
    target_root = args.target
    if not target_root.is_absolute():
        target_root = repository_root / target_root
    target_root = target_root.resolve()

    if not target_root.is_dir():
        print(f"error: target directory does not exist: {target_root}", file=sys.stderr)
        return 2

    try:
        changes = plan_changes(
            repository_root=repository_root,
            target_root=target_root,
        )
    except (OSError, UnicodeError, MigrationError) as exc:
        print(f"error: {exc}", file=sys.stderr)
        return 1

    section_total = sum(change.section_count for change in changes)
    mode = "write" if args.write else "dry-run"
    print(f"{mode}: {len(changes)} file(s), {section_total} requirement section(s)")

    for change in changes:
        relative = change.path.relative_to(repository_root)
        print(f"{relative}: {change.section_count} section(s)")

    if args.write:
        for change in changes:
            change.path.write_bytes(change.updated.encode("utf-8"))
        print("migration applied")
    elif changes:
        print("no files written; rerun with --write to apply")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
