#!/usr/bin/env python3
"""Migrate DRMCP PuppyDSL identifiers and source paths to Go-mappable names.

The migration is a dry run unless ``--write`` is supplied.

Default scope:

- PuppyDSL source files beneath ``drmcp/puppydsl``;
- source directories and ``.puppy.yaml`` file stems beneath that root;
- ``function:`` references in ordinary YAML files beneath
  ``drmcp/puppydsl/manifests``.

The ``manifests`` subtree is never treated as PuppyDSL source and its paths are
never renamed. Only exact ``function:`` references to migrated PuppyDSL
FunctionIDs are updated there.

The migration performs two independent naming changes:

1. hyphens in PuppyDSL identifiers and application-owned source path segments
   become underscores;
2. DRMCP application runtime declarations move from ``lib/runtime`` and the
   app-owned ``runtime.*`` / ``mcp.*`` function namespaces to
   ``lib/drmcp_runtime`` and ``drmcp_runtime.*``.

Generic runtime calls such as ``runtime.throw`` that already satisfy the new
identifier grammar remain unchanged. Natural-language comments, ``detail``
values, and quoted string literals are left unchanged.
"""

from __future__ import annotations

import argparse
import os
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import Sequence

DEFAULT_ROOT = Path("drmcp/puppydsl")
MANIFEST_DIRECTORY_NAME = "manifests"
SOURCE_SUFFIX = ".puppy.yaml"
APP_RUNTIME_RELATIVE = Path("lib/runtime")
APP_RUNTIME_TARGET_NAME = "drmcp_runtime"

TYPE_ID_RE = re.compile(r"^[A-Z][A-Za-z0-9]*$")
LOWER_ID_RE = re.compile(r"^[a-z][a-z0-9]*(?:_[a-z0-9]+)*$")
INTERNAL_FUNCTION_SEGMENT_RE = re.compile(r"^_[a-z][a-z0-9]*(?:_[a-z0-9]+)*$")
# Declaration keys are exactly one mapping level beneath the root. Requiring
# the first captured character to be non-whitespace prevents nested fields such
# as ``    enum:`` from being misread as declarations.
DECLARATION_KEY_RE = re.compile(
    r"^ {2}(?P<key>[^\s:#](?:[^:]*[^\s:])?):\s*$"
)
ROOT_RE = re.compile(r"^(?P<root>types|functions|contexts|events|modules):\s*$")
DETAIL_BLOCK_RE = re.compile(r"^(?P<indent>\s*)detail:\s*[>|]")
MANIFEST_FUNCTION_RE = re.compile(
    r"^(?P<prefix>\s*function:\s*)(?P<quote>['\"]?)(?P<value>[^'\"\s]+)(?P=quote)(?P<suffix>\s*)$"
)

# Remaining unquoted source tokens with hyphens are rejected after migration.
HYPHENATED_IDENTIFIER_RE = re.compile(
    r"(?<![A-Za-z0-9_])_?[a-z][a-z0-9_]*(?:-[a-z0-9_]+)+(?![A-Za-z0-9_])"
)


@dataclass(frozen=True)
class Declaration:
    source: Path
    root: str
    old_id: str
    new_id: str


@dataclass(frozen=True)
class ContentChange:
    path: Path
    original: str
    updated: str
    replacements: int
    category: str


@dataclass(frozen=True)
class Rename:
    source: Path
    destination: Path
    category: str


def infer_repo_root() -> Path:
    """Infer the repository root from ``tools/puppydsl/<script>``."""
    return Path(__file__).resolve().parents[2]


def parse_args(argv: Sequence[str]) -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Migrate DRMCP PuppyDSL identifiers and source paths to snake_case.",
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
        type=Path,
        default=DEFAULT_ROOT,
        help="Repository-relative or absolute DRMCP PuppyDSL root.",
    )
    parser.add_argument(
        "--write",
        action="store_true",
        help="Apply the migration. Without this flag, print a dry-run plan only.",
    )
    return parser.parse_args(argv)


def is_within(path: Path, parent: Path) -> bool:
    try:
        path.relative_to(parent)
    except ValueError:
        return False
    return True


def resolve_root(repo_root: Path, supplied_root: Path) -> Path:
    candidate = supplied_root if supplied_root.is_absolute() else repo_root / supplied_root
    candidate = candidate.resolve()
    if not is_within(candidate, repo_root):
        raise ValueError(f"migration root escapes repository root: {candidate}")
    if not candidate.is_dir():
        raise ValueError(f"migration root is not a directory: {candidate}")
    return candidate


def is_manifest_path(path: Path, root: Path) -> bool:
    relative = path.relative_to(root)
    return MANIFEST_DIRECTORY_NAME in relative.parts


def discover_source_files(root: Path) -> list[Path]:
    files = [
        path.resolve()
        for path in root.rglob(f"*{SOURCE_SUFFIX}")
        if path.is_file() and not is_manifest_path(path, root)
    ]
    files.sort(key=lambda path: path.relative_to(root).as_posix().casefold())
    return files


def split_generic_suffix(identifier: str) -> tuple[str, str]:
    marker = identifier.find("<")
    if marker == -1:
        return identifier, ""
    return identifier[:marker], identifier[marker:]


def snake_segment(segment: str) -> str:
    return segment.replace("-", "_")


def transform_function_id(identifier: str, source: Path, root: Path) -> str:
    base, generic_suffix = split_generic_suffix(identifier)
    segments = [snake_segment(segment) for segment in base.split(".")]
    migrated = ".".join(segments)

    runtime_root = (root / APP_RUNTIME_RELATIVE).resolve()
    if is_within(source, runtime_root):
        if migrated.startswith("runtime."):
            migrated = f"drmcp_runtime.{migrated[len('runtime.') :]}"
        elif migrated.startswith("mcp."):
            migrated = f"drmcp_runtime.{migrated}"

    return migrated + generic_suffix


def transform_declaration_id(root_name: str, identifier: str, source: Path, root: Path) -> str:
    if root_name == "types":
        if "-" in identifier:
            raise ValueError(
                f"type ID requires an explicit PascalCase migration rather than '-' -> '_': "
                f"{source}: {identifier}"
            )
        return identifier
    if root_name == "functions":
        return transform_function_id(identifier, source, root)
    return identifier.replace("-", "_")


def validate_lower_identifier(identifier: str, *, allow_internal_terminal: bool) -> None:
    base, generic_suffix = split_generic_suffix(identifier)
    if generic_suffix:
        if not generic_suffix.startswith("<") or not generic_suffix.endswith(">"):
            raise ValueError(f"malformed generic suffix: {identifier}")

    segments = base.split(".")
    if not segments or any(not segment for segment in segments):
        raise ValueError(f"malformed qualified identifier: {identifier}")

    for index, segment in enumerate(segments):
        is_terminal = index == len(segments) - 1
        if allow_internal_terminal and is_terminal and segment.startswith("_"):
            valid = INTERNAL_FUNCTION_SEGMENT_RE.fullmatch(segment) is not None
        else:
            valid = LOWER_ID_RE.fullmatch(segment) is not None
        if not valid:
            raise ValueError(f"identifier violates lower-snake grammar: {identifier}")


def validate_declaration(declaration: Declaration) -> None:
    base, _ = split_generic_suffix(declaration.new_id)
    if declaration.root == "types":
        if TYPE_ID_RE.fullmatch(base) is None:
            raise ValueError(
                f"type ID violates PascalCase grammar: {declaration.source}: {declaration.new_id}"
            )
        return

    validate_lower_identifier(
        declaration.new_id,
        allow_internal_terminal=declaration.root == "functions",
    )


def collect_declarations(source_files: Sequence[Path], root: Path) -> list[Declaration]:
    declarations: list[Declaration] = []

    for source in source_files:
        lines = source.read_text(encoding="utf-8").splitlines()
        root_name: str | None = None
        for line in lines:
            root_match = ROOT_RE.fullmatch(line)
            if root_match:
                root_name = root_match.group("root")
                continue
            if root_name is None:
                continue

            key_match = DECLARATION_KEY_RE.fullmatch(line)
            if key_match is None:
                continue

            old_id = key_match.group("key")
            new_id = transform_declaration_id(root_name, old_id, source, root)
            declaration = Declaration(source, root_name, old_id, new_id)
            validate_declaration(declaration)
            declarations.append(declaration)

    def registry_name(root_name: str) -> str:
        return "callables" if root_name in {"functions", "events"} else root_name

    seen_old: dict[tuple[str, str], Declaration] = {}
    seen_new: dict[tuple[str, str], Declaration] = {}
    for declaration in declarations:
        old_key = (registry_name(declaration.root), declaration.old_id)
        new_key = (registry_name(declaration.root), declaration.new_id)

        old_previous = seen_old.get(old_key)
        if old_previous is not None:
            raise ValueError(
                "duplicate declaration ID in migration input: "
                f"{declaration.old_id} ({old_previous.source}, {declaration.source})"
            )
        seen_old[old_key] = declaration

        new_previous = seen_new.get(new_key)
        if new_previous is not None:
            raise ValueError(
                "identifier migration collision: "
                f"{new_previous.old_id} and {declaration.old_id} -> {declaration.new_id}"
            )
        seen_new[new_key] = declaration

    return declarations


def split_code_and_comment(line: str) -> tuple[str, str]:
    quote: str | None = None
    escaped = False
    for index, character in enumerate(line):
        if escaped:
            escaped = False
            continue
        if quote == '"' and character == "\\":
            escaped = True
            continue
        if quote is not None:
            if character == quote:
                quote = None
            continue
        if character in {"'", '"'}:
            quote = character
            continue
        if character == "#":
            return line[:index], line[index:]
    return line, ""


def transform_unquoted(text: str, transform) -> tuple[str, int]:
    output: list[str] = []
    buffer: list[str] = []
    quote: str | None = None
    escaped = False
    replacements = 0

    def flush_unquoted() -> None:
        nonlocal replacements
        if not buffer:
            return
        original = "".join(buffer)
        updated, count = transform(original)
        output.append(updated)
        replacements += count
        buffer.clear()

    for character in text:
        if quote is None:
            if character in {"'", '"'}:
                flush_unquoted()
                quote = character
                output.append(character)
            else:
                buffer.append(character)
            continue

        output.append(character)
        if escaped:
            escaped = False
        elif quote == '"' and character == "\\":
            escaped = True
        elif character == quote:
            quote = None

    flush_unquoted()
    return "".join(output), replacements


def replace_identifier_tokens(text: str, identifier_map: dict[str, str]) -> tuple[str, int]:
    replacements = 0
    updated = text

    for old_id in sorted(identifier_map, key=len, reverse=True):
        new_id = identifier_map[old_id]
        if old_id == new_id:
            continue
        pattern = re.compile(
            rf"(?<![A-Za-z0-9_]){re.escape(old_id)}(?![A-Za-z0-9_-])"
        )
        updated, count = pattern.subn(new_id, updated)
        replacements += count

    updated, count = HYPHENATED_IDENTIFIER_RE.subn(
        lambda match: match.group(0).replace("-", "_"),
        updated,
    )
    replacements += count
    return updated, replacements


def transform_source_text(
    source: Path,
    original: str,
    declarations: Sequence[Declaration],
    identifier_map: dict[str, str],
) -> tuple[str, int]:
    declaration_map = {
        declaration.old_id: declaration.new_id
        for declaration in declarations
        if declaration.source == source
    }

    lines = original.splitlines(keepends=True)
    output: list[str] = []
    replacement_count = 0
    detail_block_indent: int | None = None

    for line in lines:
        bare_line = line.rstrip("\r\n")
        newline = line[len(bare_line) :]
        indentation = len(bare_line) - len(bare_line.lstrip(" "))

        if detail_block_indent is not None:
            if bare_line.strip() and indentation <= detail_block_indent:
                detail_block_indent = None
            else:
                output.append(line)
                continue

        detail_block_match = DETAIL_BLOCK_RE.match(bare_line)
        if detail_block_match:
            detail_block_indent = len(detail_block_match.group("indent"))
            output.append(line)
            continue

        stripped = bare_line.lstrip()
        if not stripped or stripped.startswith("#") or stripped.startswith("detail:"):
            output.append(line)
            continue

        declaration_key_match = DECLARATION_KEY_RE.fullmatch(bare_line)
        if declaration_key_match:
            old_id = declaration_key_match.group("key")
            new_id = declaration_map.get(old_id)
            if new_id is not None and new_id != old_id:
                updated_line = bare_line.replace(old_id, new_id, 1)
                output.append(updated_line + newline)
                replacement_count += 1
                continue

        code, comment = split_code_and_comment(bare_line)
        code_without_literals = unquoted_text(code)
        if (
            ("/" in code_without_literals or "\\" in code_without_literals)
            and HYPHENATED_IDENTIFIER_RE.search(code_without_literals)
        ):
            raise ValueError(
                "refusing to rewrite ambiguous unquoted path-like application data: "
                f"{source}: {bare_line.strip()}"
            )

        updated_code, count = transform_unquoted(
            code,
            lambda value: replace_identifier_tokens(value, identifier_map),
        )
        output.append(updated_code + comment + newline)
        replacement_count += count

    return "".join(output), replacement_count


def unquoted_text(text: str) -> str:
    output: list[str] = []
    quote: str | None = None
    escaped = False

    for character in text:
        if quote is None:
            if character in {"'", '"'}:
                quote = character
            else:
                output.append(character)
            continue

        if escaped:
            escaped = False
        elif quote == '"' and character == "\\":
            escaped = True
        elif character == quote:
            quote = None

    return "".join(output)


def remaining_hyphenated_code(text: str) -> list[tuple[int, str]]:
    findings: list[tuple[int, str]] = []
    detail_block_indent: int | None = None

    for line_number, bare_line in enumerate(text.splitlines(), start=1):
        indentation = len(bare_line) - len(bare_line.lstrip(" "))
        if detail_block_indent is not None:
            if bare_line.strip() and indentation <= detail_block_indent:
                detail_block_indent = None
            else:
                continue

        detail_block_match = DETAIL_BLOCK_RE.match(bare_line)
        if detail_block_match:
            detail_block_indent = len(detail_block_match.group("indent"))
            continue

        stripped = bare_line.lstrip()
        if not stripped or stripped.startswith("#") or stripped.startswith("detail:"):
            continue

        code, _ = split_code_and_comment(bare_line)
        match = HYPHENATED_IDENTIFIER_RE.search(unquoted_text(code))
        if match:
            findings.append((line_number, match.group(0)))

    return findings


def build_content_changes(
    source_files: Sequence[Path],
    declarations: Sequence[Declaration],
    root: Path,
) -> list[ContentChange]:
    identifier_map: dict[str, str] = {}
    for declaration in declarations:
        if declaration.old_id == declaration.new_id:
            continue
        previous = identifier_map.get(declaration.old_id)
        if previous is not None and previous != declaration.new_id:
            raise ValueError(
                "one textual declaration ID has conflicting migration targets: "
                f"{declaration.old_id} -> {previous}, {declaration.new_id}"
            )
        identifier_map[declaration.old_id] = declaration.new_id
    changes: list[ContentChange] = []

    for source in source_files:
        original = source.read_text(encoding="utf-8")
        updated, count = transform_source_text(source, original, declarations, identifier_map)
        findings = remaining_hyphenated_code(updated)
        if findings:
            formatted = ", ".join(f"line {line}: {token}" for line, token in findings)
            raise ValueError(f"unmigrated PuppyDSL identifier token in {source}: {formatted}")
        if updated != original:
            changes.append(ContentChange(source, original, updated, count, "puppy-source"))

    manifest_root = root / MANIFEST_DIRECTORY_NAME
    if manifest_root.is_dir():
        for manifest in sorted(manifest_root.rglob("*.yaml")):
            if not manifest.is_file():
                continue
            original = manifest.read_text(encoding="utf-8")
            output: list[str] = []
            count = 0
            for line in original.splitlines(keepends=True):
                bare_line = line.rstrip("\r\n")
                newline = line[len(bare_line) :]
                match = MANIFEST_FUNCTION_RE.fullmatch(bare_line)
                if match is None:
                    output.append(line)
                    continue
                old_id = match.group("value")
                new_id = identifier_map.get(old_id, old_id)
                if new_id != old_id:
                    count += 1
                output.append(
                    f"{match.group('prefix')}{match.group('quote')}{new_id}"
                    f"{match.group('quote')}{match.group('suffix')}{newline}"
                )
            updated = "".join(output)
            if updated != original:
                changes.append(ContentChange(manifest, original, updated, count, "manifest-reference"))

    changes.sort(key=lambda change: change.path.relative_to(root).as_posix().casefold())
    return changes


def snake_path_name(name: str) -> str:
    return name.replace("-", "_")


def build_rename_plan(root: Path, source_files: Sequence[Path]) -> list[Rename]:
    renames: list[Rename] = []

    for source in source_files:
        target_name = snake_path_name(source.name)
        target_stem = target_name[: -len(SOURCE_SUFFIX)]
        if LOWER_ID_RE.fullmatch(target_stem) is None:
            raise ValueError(
                f"source-file stem violates lower-snake grammar after migration: {source}"
            )
        if target_name == source.name:
            continue
        destination = source.with_name(target_name)
        renames.append(Rename(source, destination, "source-file"))

    directories = [
        path.resolve()
        for path in root.rglob("*")
        if path.is_dir() and not is_manifest_path(path, root)
    ]
    directories.sort(
        key=lambda path: (len(path.relative_to(root).parts), path.as_posix().casefold()),
        reverse=True,
    )

    runtime_source = (root / APP_RUNTIME_RELATIVE).resolve()
    for directory in directories:
        target_name = directory.name
        category = "source-directory"
        if directory == runtime_source:
            target_name = APP_RUNTIME_TARGET_NAME
            category = "drmcp-runtime-directory"
        elif "-" in target_name:
            target_name = snake_path_name(target_name)
        else:
            if LOWER_ID_RE.fullmatch(target_name) is None:
                raise ValueError(
                    "source-directory segment violates lower-snake grammar: "
                    f"{directory}"
                )
            continue

        if LOWER_ID_RE.fullmatch(target_name) is None:
            raise ValueError(
                "source-directory segment violates lower-snake grammar after migration: "
                f"{directory} -> {target_name}"
            )
        renames.append(Rename(directory, directory.with_name(target_name), category))

    validate_rename_plan(renames)
    return renames


def validate_rename_plan(renames: Sequence[Rename]) -> None:
    destinations: dict[str, Path] = {}
    sources = {os.path.normcase(os.path.abspath(rename.source)) for rename in renames}

    for rename in renames:
        normalized_destination = os.path.normcase(os.path.abspath(rename.destination))
        previous = destinations.get(normalized_destination)
        if previous is not None:
            raise ValueError(
                "multiple paths map to one migration destination: "
                f"{previous} and {rename.source} -> {rename.destination}"
            )
        destinations[normalized_destination] = rename.source

        if rename.destination.exists() and normalized_destination not in sources:
            raise ValueError(f"migration destination already exists: {rename.destination}")


def display_path(path: Path, repo_root: Path) -> str:
    return path.relative_to(repo_root).as_posix()


def apply_changes(changes: Sequence[ContentChange], renames: Sequence[Rename]) -> None:
    applied_renames: list[Rename] = []
    try:
        for change in changes:
            change.path.write_text(change.updated, encoding="utf-8", newline="")

        # File renames precede directory renames. Directory renames are already
        # deepest-first so child paths remain addressable during application.
        for rename in renames:
            rename.source.rename(rename.destination)
            applied_renames.append(rename)
    except OSError as error:
        rollback_errors: list[str] = []

        for rename in reversed(applied_renames):
            try:
                rename.destination.rename(rename.source)
            except OSError as rollback_error:
                rollback_errors.append(
                    f"rename rollback {rename.destination} -> {rename.source}: {rollback_error}"
                )

        for change in changes:
            try:
                change.path.write_text(change.original, encoding="utf-8", newline="")
            except OSError as rollback_error:
                rollback_errors.append(
                    f"content rollback {change.path}: {rollback_error}"
                )

        message = f"migration failed: {error}"
        if rollback_errors:
            message += "\nrollback also failed:\n  " + "\n  ".join(rollback_errors)
        raise RuntimeError(message) from error


def main(argv: Sequence[str] | None = None) -> int:
    args = parse_args(argv if argv is not None else sys.argv[1:])

    try:
        repo_root = args.repo_root.resolve()
        if not repo_root.is_dir():
            raise ValueError(f"repository root is not a directory: {repo_root}")
        root = resolve_root(repo_root, args.root)
        runtime_source = root / APP_RUNTIME_RELATIVE
        if not runtime_source.is_dir():
            raise ValueError(
                "expected one-time DRMCP runtime migration source is missing: "
                f"{runtime_source}"
            )
        source_files = discover_source_files(root)
        declarations = collect_declarations(source_files, root)
        content_changes = build_content_changes(source_files, declarations, root)
        renames = build_rename_plan(root, source_files)
    except (OSError, UnicodeError, ValueError) as error:
        print(f"error: {error}", file=sys.stderr)
        return 2

    mode = "WRITE" if args.write else "DRY RUN"
    print(f"mode: {mode}")
    print(f"root: {display_path(root, repo_root)}")
    print(f"source_files: {len(source_files)}")
    print(f"declarations: {len(declarations)}")

    for declaration in declarations:
        if declaration.old_id == declaration.new_id:
            continue
        print(
            "identifier: "
            f"{declaration.old_id} -> {declaration.new_id} "
            f"({display_path(declaration.source, repo_root)})"
        )

    for change in content_changes:
        print(
            f"content[{change.category}]: {display_path(change.path, repo_root)} "
            f"replacements={change.replacements}"
        )

    for rename in renames:
        print(
            f"rename[{rename.category}]: {display_path(rename.source, repo_root)} "
            f"-> {display_path(rename.destination, repo_root)}"
        )

    print(
        "summary: "
        f"identifiers_changed={sum(d.old_id != d.new_id for d in declarations)} "
        f"content_files_changed={len(content_changes)} "
        f"paths_renamed={len(renames)}"
    )

    if not args.write:
        print("dry run only; pass --write to apply")
        return 0

    try:
        apply_changes(content_changes, renames)
    except RuntimeError as error:
        print(f"error: {error}", file=sys.stderr)
        return 1

    print("migration applied")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
