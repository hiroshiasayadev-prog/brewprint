#!/usr/bin/env python3
"""Relocate DRMCP application-architecture specs under implementation.

Default mode is dry-run.
Use --apply to write files.

Scope:
- Move only the five application-architecture specification files.
- Rewrite only Markdown files under drmcp/records/spec/.
- Replace spec:drmcp.application_architecture with
  spec:drmcp.implementation.application_architecture.
- Leave no compatibility stub under the old spec tree.
"""

from __future__ import annotations

import argparse
from dataclasses import dataclass
from pathlib import Path
import sys

OLD_REF = "spec:drmcp.application_architecture"
NEW_REF = "spec:drmcp.implementation.application_architecture"
SPEC_ROOT = Path("drmcp/records/spec")
SOURCE_DIR = SPEC_ROOT / "application-architecture"
TARGET_DIR = SPEC_ROOT / "implementation" / "application-architecture"

MOVED_FILES = (
    "index.md",
    "application-boundary-and-components.md",
    "dependency-and-responsibility.md",
    "runtime-and-state.md",
    "failure-and-evolution.md",
)


@dataclass(frozen=True)
class MovePlan:
    source: Path
    target: Path
    replacements: int


@dataclass(frozen=True)
class RewritePlan:
    path: Path
    replacements: int


def repo_root_from(start: Path) -> Path:
    """Find the repository root from a starting path."""
    for candidate in (start.resolve(), *start.resolve().parents):
        if (candidate / "prompt_chappy.md").is_file() and (candidate / SPEC_ROOT).is_dir():
            return candidate
    raise RuntimeError("could not locate repo root containing prompt_chappy.md and drmcp/records/spec")


def read_utf8(path: Path) -> str:
    return path.read_text(encoding="utf-8")


def write_utf8(path: Path, text: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(text, encoding="utf-8", newline="\n")


def rel(path: Path, root: Path) -> str:
    return path.relative_to(root).as_posix()


def expected_source_paths(root: Path) -> list[Path]:
    return [root / SOURCE_DIR / name for name in MOVED_FILES]


def expected_target_paths(root: Path) -> list[Path]:
    return [root / TARGET_DIR / name for name in MOVED_FILES]


def assert_preconditions(root: Path) -> None:
    source_dir = root / SOURCE_DIR
    target_dir = root / TARGET_DIR

    if not source_dir.is_dir():
        raise RuntimeError(f"missing source directory: {rel(source_dir, root)}")

    for source in expected_source_paths(root):
        if not source.is_file():
            raise RuntimeError(f"missing expected source file: {rel(source, root)}")

    expected_names = set(MOVED_FILES)
    actual_names = {path.name for path in source_dir.iterdir() if path.is_file()}
    extra_names = sorted(actual_names - expected_names)
    if extra_names:
        joined = ", ".join(extra_names)
        raise RuntimeError(f"unexpected file(s) in source directory: {joined}")

    for target in expected_target_paths(root):
        if target.exists():
            raise RuntimeError(f"destination already exists: {rel(target, root)}")

    if target_dir.exists() and any(target_dir.iterdir()):
        raise RuntimeError(f"target directory exists and is not empty: {rel(target_dir, root)}")


def build_plan(root: Path) -> tuple[list[MovePlan], list[RewritePlan]]:
    source_paths = set(expected_source_paths(root))
    move_plan: list[MovePlan] = []
    rewrite_plan: list[RewritePlan] = []

    for source_name in MOVED_FILES:
        source = root / SOURCE_DIR / source_name
        target = root / TARGET_DIR / source_name
        text = read_utf8(source)
        move_plan.append(MovePlan(source=source, target=target, replacements=text.count(OLD_REF)))

    for path in sorted((root / SPEC_ROOT).rglob("*.md")):
        if path in source_paths:
            continue
        text = read_utf8(path)
        count = text.count(OLD_REF)
        if count:
            rewrite_plan.append(RewritePlan(path=path, replacements=count))

    return move_plan, rewrite_plan


def print_plan(root: Path, move_plan: list[MovePlan], rewrite_plan: list[RewritePlan]) -> None:
    total_move_replacements = sum(item.replacements for item in move_plan)
    total_rewrite_replacements = sum(item.replacements for item in rewrite_plan)

    print("mode: dry-run")
    print(f"old_ref: {OLD_REF}")
    print(f"new_ref: {NEW_REF}")
    print(f"source_dir: {SOURCE_DIR.as_posix()}")
    print(f"target_dir: {TARGET_DIR.as_posix()}")
    print("")
    print("planned_moves:")
    for item in move_plan:
        print(
            "- "
            f"{rel(item.source, root)} -> {rel(item.target, root)} "
            f"with {item.replacements} ref replacement(s)"
        )
    print("")
    print("planned_in_place_rewrites:")
    if rewrite_plan:
        for item in rewrite_plan:
            print(f"- {rel(item.path, root)}: {item.replacements} ref replacement(s)")
    else:
        print("- none")
    print("")
    print("summary:")
    print(f"- move_count: {len(move_plan)}")
    print(f"- in_place_rewrite_file_count: {len(rewrite_plan)}")
    print(f"- ref_replacement_count: {total_move_replacements + total_rewrite_replacements}")


def apply_plan(root: Path, move_plan: list[MovePlan], rewrite_plan: list[RewritePlan]) -> None:
    for item in move_plan:
        text = read_utf8(item.source).replace(OLD_REF, NEW_REF)
        write_utf8(item.target, text)

    for item in rewrite_plan:
        text = read_utf8(item.path)
        write_utf8(item.path, text.replace(OLD_REF, NEW_REF))

    for item in move_plan:
        item.source.unlink()

    source_dir = root / SOURCE_DIR
    if source_dir.exists():
        try:
            source_dir.rmdir()
        except OSError as exc:
            raise RuntimeError(f"source directory is not empty after migration: {rel(source_dir, root)}") from exc

    leftovers = find_old_ref_leftovers(root)
    if leftovers:
        formatted = "\n".join(f"- {rel(path, root)}" for path in leftovers)
        raise RuntimeError(f"old ref remains under {SPEC_ROOT.as_posix()}:\n{formatted}")


def find_old_ref_leftovers(root: Path) -> list[Path]:
    leftovers: list[Path] = []
    for path in sorted((root / SPEC_ROOT).rglob("*.md")):
        if OLD_REF in read_utf8(path):
            leftovers.append(path)
    return leftovers


def parse_args(argv: list[str]) -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Relocate DRMCP application-architecture specs under implementation."
    )
    parser.add_argument(
        "--repo-root",
        type=Path,
        default=None,
        help="Repository root. Defaults to auto-detection from the current directory.",
    )
    parser.add_argument(
        "--apply",
        action="store_true",
        help="Apply the migration. Omit this flag for dry-run.",
    )
    return parser.parse_args(argv)


def main(argv: list[str]) -> int:
    args = parse_args(argv)
    root = args.repo_root.resolve() if args.repo_root else repo_root_from(Path.cwd())

    assert_preconditions(root)
    move_plan, rewrite_plan = build_plan(root)

    if not args.apply:
        print_plan(root, move_plan, rewrite_plan)
        return 0

    apply_plan(root, move_plan, rewrite_plan)
    print("mode: apply")
    print("result: applied")
    print(f"moved_files: {len(move_plan)}")
    print(f"rewritten_files: {len(rewrite_plan)}")
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main(sys.argv[1:]))
    except Exception as exc:  # noqa: BLE001 - migration scripts should print one clear failure.
        print(f"error: {exc}", file=sys.stderr)
        raise SystemExit(1)
