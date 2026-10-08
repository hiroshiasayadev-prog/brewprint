#!/usr/bin/env python3
"""Rename PuppyDSL source files from ``*.yaml`` to ``*.puppy.yaml``.

The migration is a dry run unless ``--write`` is supplied. By default it scans:

- ``drmcp/puppydsl`` while excluding every ``manifests`` subtree;
- ``validation/runtime/examples``.

Ordinary YAML manifests are intentionally left unchanged.
"""

from __future__ import annotations

import argparse
import os
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import Sequence

DEFAULT_ROOTS = (
    Path("drmcp/puppydsl"),
    Path("validation/runtime/examples"),
)
MANIFEST_DIRECTORY_NAME = "manifests"
SOURCE_SUFFIX = ".yaml"
PUPPY_SOURCE_SUFFIX = ".puppy.yaml"


@dataclass(frozen=True)
class Rename:
    source: Path
    destination: Path


def infer_repo_root() -> Path:
    """Infer the repository root from ``tools/puppydsl/<script>``."""
    return Path(__file__).resolve().parents[2]


def parse_args(argv: Sequence[str]) -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Rename PuppyDSL *.yaml files to *.puppy.yaml.",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )
    parser.add_argument(
        "--repo-root",
        type=Path,
        default=infer_repo_root(),
        help="Brewprint repository root.",
    )
    parser.add_argument(
        "--root",
        action="append",
        type=Path,
        dest="roots",
        help=(
            "Repository-relative or absolute directory to scan. Repeat for multiple "
            "roots. When omitted, the built-in PuppyDSL roots are used."
        ),
    )
    parser.add_argument(
        "--write",
        action="store_true",
        help="Apply the renames. Without this flag, print a dry-run plan only.",
    )
    return parser.parse_args(argv)


def is_within(path: Path, parent: Path) -> bool:
    try:
        path.relative_to(parent)
    except ValueError:
        return False
    return True


def resolve_scan_roots(repo_root: Path, supplied_roots: Sequence[Path] | None) -> list[Path]:
    raw_roots = supplied_roots if supplied_roots else DEFAULT_ROOTS
    resolved: list[Path] = []

    for raw_root in raw_roots:
        candidate = raw_root if raw_root.is_absolute() else repo_root / raw_root
        candidate = candidate.resolve()

        if not is_within(candidate, repo_root):
            raise ValueError(f"scan root escapes repository root: {candidate}")
        if not candidate.exists():
            raise ValueError(f"scan root does not exist: {candidate}")
        if not candidate.is_dir():
            raise ValueError(f"scan root is not a directory: {candidate}")

        if candidate not in resolved:
            resolved.append(candidate)

    return resolved


def is_manifest_path(path: Path, repo_root: Path) -> bool:
    relative_parts = path.relative_to(repo_root).parts[:-1]
    return MANIFEST_DIRECTORY_NAME in relative_parts


def destination_for(source: Path) -> Path:
    stem = source.name[: -len(SOURCE_SUFFIX)]
    return source.with_name(f"{stem}{PUPPY_SOURCE_SUFFIX}")


def build_plan(repo_root: Path, scan_roots: Sequence[Path]) -> tuple[list[Rename], int, int]:
    sources: set[Path] = set()
    skipped_manifests = 0
    skipped_already_renamed = 0

    for scan_root in scan_roots:
        for source in scan_root.rglob(f"*{SOURCE_SUFFIX}"):
            if not source.is_file():
                continue
            if not source.name.endswith(SOURCE_SUFFIX):
                continue
            if source.name.endswith(PUPPY_SOURCE_SUFFIX):
                skipped_already_renamed += 1
                continue
            if is_manifest_path(source, repo_root):
                skipped_manifests += 1
                continue
            if source.is_symlink():
                raise ValueError(f"refusing to rename symbolic link: {source}")

            resolved_source = source.resolve()
            if not is_within(resolved_source, repo_root):
                raise ValueError(f"source escapes repository root: {source}")
            sources.add(resolved_source)

    plan = [Rename(source, destination_for(source)) for source in sources]
    plan.sort(key=lambda item: str(item.source.relative_to(repo_root)).casefold())
    validate_plan(plan)
    return plan, skipped_manifests, skipped_already_renamed


def validate_plan(plan: Sequence[Rename]) -> None:
    destinations: dict[str, Path] = {}

    for item in plan:
        normalized = os.path.normcase(os.path.abspath(item.destination))
        previous = destinations.get(normalized)
        if previous is not None:
            raise ValueError(
                "multiple source files map to the same destination: "
                f"{previous} and {item.source} -> {item.destination}"
            )
        destinations[normalized] = item.source

        if item.destination.exists():
            raise ValueError(f"destination already exists: {item.destination}")


def apply_plan(plan: Sequence[Rename]) -> None:
    completed: list[Rename] = []

    try:
        for item in plan:
            item.source.rename(item.destination)
            completed.append(item)
    except OSError as error:
        rollback_errors: list[str] = []
        for item in reversed(completed):
            try:
                item.destination.rename(item.source)
            except OSError as rollback_error:
                rollback_errors.append(
                    f"{item.destination} -> {item.source}: {rollback_error}"
                )

        message = f"rename failed: {error}"
        if rollback_errors:
            message += "\nrollback also failed:\n  " + "\n  ".join(rollback_errors)
        raise RuntimeError(message) from error


def display_path(path: Path, repo_root: Path) -> str:
    return path.relative_to(repo_root).as_posix()


def main(argv: Sequence[str] | None = None) -> int:
    args = parse_args(argv if argv is not None else sys.argv[1:])

    try:
        repo_root = args.repo_root.resolve()
        if not repo_root.is_dir():
            raise ValueError(f"repository root is not a directory: {repo_root}")

        scan_roots = resolve_scan_roots(repo_root, args.roots)
        plan, skipped_manifests, skipped_already_renamed = build_plan(
            repo_root,
            scan_roots,
        )
    except ValueError as error:
        print(f"error: {error}", file=sys.stderr)
        return 2

    mode = "WRITE" if args.write else "DRY RUN"
    print(f"mode: {mode}")
    for root in scan_roots:
        print(f"root: {display_path(root, repo_root)}")

    for item in plan:
        print(
            f"rename: {display_path(item.source, repo_root)} "
            f"-> {display_path(item.destination, repo_root)}"
        )

    print(
        "summary: "
        f"rename={len(plan)} "
        f"manifest_skipped={skipped_manifests} "
        f"already_puppy_skipped={skipped_already_renamed}"
    )

    if not args.write:
        print("dry run only; pass --write to apply")
        return 0

    try:
        apply_plan(plan)
    except RuntimeError as error:
        print(f"error: {error}", file=sys.stderr)
        return 1

    print(f"renamed {len(plan)} file(s)")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
