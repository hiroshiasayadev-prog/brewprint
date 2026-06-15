#!/usr/bin/env python3
"""
Temporary PRODUCT-level spec format validator (PRODUCT-WORK-SPEC-006).
Bridge tooling for PRODUCT-WORK-SPEC-005 migration review.

Not a DRMCP reimplementation. Does not patch the current DRMCP codebase.

Usage:
  python validate_spec.py [PATH...] [--strict] [--no-color]

  PATH  One or more spec .md files or directories to scan recursively.
        Defaults to product/records/spec if run from repo root.

  --strict  Treat migration-phase warnings as errors
            (use for new or explicitly migrated specs).

  --no-color  Disable ANSI color output.

Exit codes:
  0 — no errors (warnings may be present)
  1 — one or more errors
"""

import argparse
import re
import sys
from pathlib import Path

# ---------------------------------------------------------------------------
# Constants
# ---------------------------------------------------------------------------

ACCEPTED_KINDS = {"Overview", "Index", "Concept", "Reference", "Contract"}
CONTRACT_CLASSES = {"interface", "format"}

# Section requirement matrix derived from spec:product.concepts.spec_format.document_shape
# Values: "required" | "recommended" | "optional" | "prohibited"
# Key: (section_heading, kind_key)
# kind_key is one of: Overview Index Concept Reference Contract:interface Contract:format
_KINDS_ALL = ("Overview", "Index", "Concept", "Reference", "Contract:interface", "Contract:format")

# fmt: off
SECTION_MATRIX: dict[str, dict[str, str]] = {
    "## What this is": {
        "Overview": "required", "Index": "required", "Concept": "required",
        "Reference": "required", "Contract:interface": "required", "Contract:format": "required",
    },
    "## Current contract": {
        "Overview": "required", "Index": "prohibited", "Concept": "optional",
        "Reference": "optional", "Contract:interface": "optional", "Contract:format": "required",
    },
    "## Non-goals": {
        "Overview": "recommended", "Index": "prohibited", "Concept": "recommended",
        "Reference": "optional", "Contract:interface": "recommended", "Contract:format": "recommended",
    },
    "## Topic map": {
        "Overview": "recommended", "Index": "prohibited", "Concept": "optional",
        "Reference": "optional", "Contract:interface": "optional", "Contract:format": "optional",
    },
    "## Topics": {
        "Overview": "optional", "Index": "required", "Concept": "prohibited",
        "Reference": "prohibited", "Contract:interface": "prohibited", "Contract:format": "prohibited",
    },
    "## Concept model": {
        "Overview": "optional", "Index": "prohibited", "Concept": "required",
        "Reference": "prohibited", "Contract:interface": "optional", "Contract:format": "optional",
    },
    "## Rules": {
        "Overview": "optional", "Index": "prohibited", "Concept": "recommended",
        "Reference": "optional", "Contract:interface": "optional", "Contract:format": "required",
    },
    "## Boundary": {
        "Overview": "optional", "Index": "prohibited", "Concept": "recommended",
        "Reference": "optional", "Contract:interface": "recommended", "Contract:format": "recommended",
    },
    "## Request": {
        "Overview": "prohibited", "Index": "prohibited", "Concept": "prohibited",
        "Reference": "prohibited", "Contract:interface": "required", "Contract:format": "prohibited",
    },
    "## Response": {
        "Overview": "prohibited", "Index": "prohibited", "Concept": "prohibited",
        "Reference": "prohibited", "Contract:interface": "required", "Contract:format": "prohibited",
    },
    "## Errors": {
        "Overview": "optional", "Index": "prohibited", "Concept": "prohibited",
        "Reference": "optional", "Contract:interface": "required", "Contract:format": "optional",
    },
    "## Validation rules": {
        "Overview": "optional", "Index": "prohibited", "Concept": "optional",
        "Reference": "optional", "Contract:interface": "optional", "Contract:format": "required",
    },
    "## Related specs": {
        "Overview": "recommended", "Index": "recommended", "Concept": "recommended",
        "Reference": "recommended", "Contract:interface": "recommended", "Contract:format": "recommended",
    },
}
# fmt: on

TOPICS_REQUIRED_COLUMNS = {"title", "kind", "ref", "summary"}

# Parent grammar: root | - | spec:<segment>(.<segment>)*
_PARENT_RE = re.compile(r'^(root|-|spec:[a-z0-9][a-z0-9_]*(\.[a-z0-9][a-z0-9_]*)*)$')

# H1 format: # <SpecKind>: <Title>
_H1_RE = re.compile(r'^#\s+(\w+):\s+(.+)$')

# H1-adjacent metadata line: - **key**: value
_META_RE = re.compile(r'^-\s+\*\*(\w+)\*\*:\s*(.*)')


# ---------------------------------------------------------------------------
# Path-derived canonical spec ref
# ---------------------------------------------------------------------------

def derive_spec_id(path: Path, repo_root: Path) -> str | None:
    """
    Compute the path-derived canonical spec: ref for a spec file.
    Returns None if the path is not under a recognized spec tree.
    """
    try:
        rel = path.relative_to(repo_root)
    except ValueError:
        return None

    parts = rel.parts  # e.g. ('product', 'records', 'spec', 'concepts', 'spec-format', 'index.md')

    if len(parts) < 3:
        return None

    app = parts[0].lower()

    if len(parts) < 3 or parts[1] != "records" or parts[2] != "spec":
        return None

    sub = list(parts[3:])  # remaining directories + filename

    if not sub:
        return f"spec:{app}"

    filename = sub[-1]
    if filename.lower() == "index.md":
        sub = sub[:-1]
    else:
        sub[-1] = filename[:-3]  # strip .md

    # Normalize hyphens to underscores in each segment
    sub = [s.replace("-", "_") for s in sub]

    if not sub:
        return f"spec:{app}"

    return f"spec:{app}." + ".".join(sub)


# ---------------------------------------------------------------------------
# File parser
# ---------------------------------------------------------------------------

def _classify_lines(lines: list[str]) -> list[str]:
    """
    Classify each line as 'front_matter', 'fenced', or 'content'.
    Parser-aware: YAML front matter and fenced code blocks are not content.
    """
    classified = []
    in_front_matter = False
    in_fenced = False
    fenced_char = None

    for i, line in enumerate(lines):
        stripped = line.strip()

        # YAML front matter starts at line 0 with ---
        if i == 0 and stripped == "---":
            in_front_matter = True
            classified.append("front_matter")
            continue

        if in_front_matter:
            classified.append("front_matter")
            if stripped == "---":
                in_front_matter = False
            continue

        # Fenced code blocks
        if not in_fenced:
            m = re.match(r'^(`{3,}|~{3,})', line)
            if m:
                in_fenced = True
                fenced_char = m.group(1)[0]
                classified.append("fenced")
                continue
        else:
            m = re.match(r'^([`~]{3,})', line)
            if m and m.group(1)[0] == fenced_char:
                in_fenced = False
                fenced_char = None
            classified.append("fenced")
            continue

        classified.append("content")

    return classified


def _parse_yaml_front_matter(lines: list[str], classified: list[str]) -> dict:
    """Extract key-value pairs from YAML front matter (simple flat keys only)."""
    fm = {}
    for i, line in enumerate(lines):
        if classified[i] != "front_matter":
            continue
        m = re.match(r'^(\w+)\s*:\s*(.*)', line)
        if m:
            fm[m.group(1)] = m.group(2).strip()
    return fm


def _parse_h1_adjacent_meta(lines: list[str], classified: list[str], h1_idx: int) -> dict:
    """
    Extract H1-adjacent metadata from lines immediately after H1.
    Stops at the next ## heading or when no more meta lines appear.
    """
    meta = {}
    i = h1_idx + 1
    while i < len(lines):
        if classified[i] != "content":
            i += 1
            continue
        l = lines[i].strip()
        if l.startswith("## "):
            break
        m = _META_RE.match(l)
        if m:
            meta[m.group(1)] = m.group(2).strip()
        # blank lines between meta markers are allowed; non-meta, non-blank stops scan
        elif l and not l.startswith("-"):
            break
        i += 1
    return meta


def _collect_sections(lines: list[str], classified: list[str]) -> set[str]:
    """Collect all real H2 headings (outside front matter and fenced blocks)."""
    sections = set()
    for i, line in enumerate(lines):
        if classified[i] != "content":
            continue
        if line.startswith("## "):
            sections.add(line.rstrip())
    return sections


def _parse_topics_table(lines: list[str], classified: list[str]) -> dict | None:
    """
    Find the ## Topics section and parse the Markdown table.
    Returns {'columns': [...], 'rows': [[...], ...]} or None.
    """
    in_topics = False
    columns: list[str] = []
    rows: list[list[str]] = []

    for i, line in enumerate(lines):
        if classified[i] != "content":
            continue
        stripped = line.strip()

        if not in_topics:
            if stripped == "## Topics":
                in_topics = True
            continue

        if stripped.startswith("## "):
            break  # next section

        if not stripped.startswith("|"):
            continue

        cells = [c.strip() for c in stripped.split("|")]
        cells = [c for c in cells if c != ""]  # drop leading/trailing empty from split

        if not columns:
            columns = [c.lower() for c in cells]
        elif re.match(r'^[\s\-|]+$', stripped):
            continue  # separator row
        else:
            rows.append(cells)

    if not columns:
        return None
    return {"columns": columns, "rows": rows}


def _has_table_in_body(lines: list[str], classified: list[str]) -> bool:
    """Check if any body H2 section contains a Markdown table row."""
    in_h2 = False
    for i, line in enumerate(lines):
        if classified[i] != "content":
            continue
        if line.startswith("# ") and not line.startswith("## "):
            continue  # H1, skip
        if line.startswith("## "):
            in_h2 = True
            continue
        if in_h2 and line.strip().startswith("|") and "|" in line[1:]:
            return True
    return False


def parse_file(path: Path) -> dict:
    text = path.read_text(encoding="utf-8")
    lines = text.splitlines()
    classified = _classify_lines(lines)

    has_front_matter = any(c == "front_matter" for c in classified)
    fm = _parse_yaml_front_matter(lines, classified) if has_front_matter else {}

    # Real ATX H1s
    real_h1s = [
        (i, lines[i].rstrip())
        for i, c in enumerate(classified)
        if c == "content" and re.match(r'^# [^#]', lines[i])
    ]

    meta = {}
    if real_h1s:
        meta = _parse_h1_adjacent_meta(lines, classified, real_h1s[0][0])

    sections = _collect_sections(lines, classified)

    topics_table = None
    if any(s == "## Topics" for s in sections):
        topics_table = _parse_topics_table(lines, classified)

    return {
        "lines": lines,
        "classified": classified,
        "has_front_matter": has_front_matter,
        "front_matter": fm,
        "real_h1s": real_h1s,   # list of (line_idx, text)
        "h1_adjacent_meta": meta,
        "sections": sections,
        "topics_table": topics_table,
        "has_table_in_body": _has_table_in_body(lines, classified),
    }


# ---------------------------------------------------------------------------
# Validator
# ---------------------------------------------------------------------------

def check_spec(path: Path, repo_root: Path, strict: bool = False) -> list[dict]:
    """
    Validate a single spec file. Returns list of diagnostic dicts.
    Each dict: {severity: "ERROR"|"WARN"|"INFO", code: str, message: str}

    Severity convention (matches validation-policy.md):
    - Inventory mode (strict=False): migration-phase issues are WARN
    - Strict mode (strict=True):     migration-phase issues are ERROR
    """
    diags: list[dict] = []

    def emit(severity, code, msg):
        diags.append({"severity": severity, "code": code, "message": msg})

    def error(code, msg):
        emit("ERROR", code, msg)

    def warn(code, msg):
        emit("WARN", code, msg)

    def info(code, msg):
        emit("INFO", code, msg)

    def inv(code, msg):
        # inventory-phase: error in strict mode, warning in inventory mode
        (error if strict else warn)(code, msg)

    try:
        parsed = parse_file(path)
    except Exception as e:
        error("PARSE_FAILED", f"Could not read or parse file: {e}")
        return diags

    # --- YAML front matter ---
    if parsed["has_front_matter"]:
        inv("YAML_FRONT_MATTER",
            "File has YAML front matter. Use H1-adjacent visible markers instead.")
        fm = parsed["front_matter"]
        if "design_record" in "\n".join(parsed["lines"][:20]):
            inv("FM_DESIGN_RECORD", "Front matter appears to contain design_record.kind.")
        for hidden_key in ("depends_on", "semantic_refs", "sections"):
            if hidden_key in fm:
                inv("FM_HIDDEN_REFS", f"Front matter contains hidden '{hidden_key}' metadata.")

    # --- H1 count ---
    h1s = parsed["real_h1s"]

    if len(h1s) == 0:
        error("MISSING_H1", "No real ATX H1 found outside YAML front matter and fenced code blocks.")
        return diags  # cannot proceed without H1

    if len(h1s) > 1:
        h1_texts = [t for _, t in h1s]
        error("MULTIPLE_H1", f"Multiple real ATX H1 headings: {h1_texts}")

    # --- H1 format ---
    h1_text = h1s[0][1]
    h1_match = _H1_RE.match(h1_text)

    if not h1_match:
        inv("H1_FORMAT", f"H1 does not match '# <SpecKind>: <Title>': {h1_text!r}")
        return diags  # cannot determine kind; skip remaining checks

    spec_kind = h1_match.group(1)

    if spec_kind not in ACCEPTED_KINDS:
        error("H1_KIND", f"Spec kind {spec_kind!r} is not in the accepted set {sorted(ACCEPTED_KINDS)}.")
        return diags

    # --- H1-adjacent metadata ---
    meta = parsed["h1_adjacent_meta"]

    for field in ("id", "status", "date", "parent"):
        if field not in meta:
            inv("MISSING_META", f"H1-adjacent '- **{field}**:' marker is missing.")

    # contract_class
    if spec_kind == "Contract":
        if "contract_class" not in meta:
            inv("MISSING_CONTRACT_CLASS",
                "H1-adjacent '- **contract_class**:' is required for Contract specs.")
            contract_class = ""
        else:
            contract_class = meta["contract_class"].strip().strip("`")
            if contract_class not in CONTRACT_CLASSES:
                error("INVALID_CONTRACT_CLASS",
                      f"contract_class {contract_class!r} is not valid; must be one of {sorted(CONTRACT_CLASSES)}.")
                contract_class = ""
    else:
        contract_class = ""
        if "contract_class" in meta:
            error("UNEXPECTED_CONTRACT_CLASS",
                  f"'- **contract_class**:' is prohibited for {spec_kind} specs.")

    # Effective kind key for section matrix
    if spec_kind == "Contract" and contract_class:
        kind_key = f"Contract:{contract_class}"
    else:
        kind_key = spec_kind

    # --- Path-derived ID check ---
    derived_id = derive_spec_id(path, repo_root)
    if derived_id is not None and "id" in meta:
        declared_id = meta["id"].strip().strip("`")
        if declared_id != derived_id:
            inv("ID_MISMATCH",
                f"Declared id {declared_id!r} != path-derived {derived_id!r}.")
    elif derived_id is None:
        info("ID_SKIP", "Path is not under a recognized <app>/records/spec/ tree; skipping ID derivation.")

    # --- Parent grammar check ---
    if "parent" in meta:
        parent_val = meta["parent"].strip().strip("`")
        if not _PARENT_RE.match(parent_val):
            error("PARENT_GRAMMAR",
                  f"Parent value {parent_val!r} does not match allowed grammar "
                  "(root | - | spec:<segment>(.<segment>)*).")

    # --- Section checks ---
    sections = parsed["sections"]

    if kind_key in _KINDS_ALL or spec_kind in ACCEPTED_KINDS:
        for section_heading, requirements in SECTION_MATRIX.items():
            req = requirements.get(kind_key) or requirements.get(spec_kind, "optional")
            present = section_heading in sections

            if req == "required" and not present:
                inv("MISSING_SECTION",
                    f"Required section {section_heading!r} is missing for {spec_kind} "
                    + (f"({contract_class})" if contract_class else "") + " spec.")

            elif req == "prohibited" and present:
                error("PROHIBITED_SECTION",
                      f"Section {section_heading!r} is prohibited for {spec_kind} "
                      + (f"({contract_class})" if contract_class else "") + " spec.")

    # Reference: must have at least one H2 body section containing a Markdown table
    if spec_kind == "Reference" and not parsed["has_table_in_body"]:
        inv("REFERENCE_NO_TABLE",
            "Reference spec has no H2 body section containing a Markdown table.")

    # --- Topics table checks ---
    if spec_kind == "Index" and "## Topics" not in sections:
        inv("MISSING_TOPICS_SECTION", "Index spec is missing the required '## Topics' section.")

    topics = parsed["topics_table"]
    if topics is not None:
        columns_present = set(topics["columns"])

        for col in TOPICS_REQUIRED_COLUMNS:
            if col not in columns_present:
                inv("TOPICS_MISSING_COL",
                    f"'## Topics' table is missing required column {col!r}.")

        if "file" in columns_present and "ref" not in columns_present:
            error("TOPICS_FILE_NOT_REF",
                  "'## Topics' table uses 'file' instead of 'ref' as child target column.")

        # Duplicate ref check
        ref_col_idx = topics["columns"].index("ref") if "ref" in topics["columns"] else None
        if ref_col_idx is not None:
            seen_refs: set[str] = set()
            for row in topics["rows"]:
                if ref_col_idx < len(row):
                    ref_val = row[ref_col_idx].strip().strip("`")
                    if ref_val in seen_refs:
                        error("TOPICS_DUPLICATE_REF",
                              f"Duplicate child ref {ref_val!r} in '## Topics' table.")
                    seen_refs.add(ref_val)

    return diags


# ---------------------------------------------------------------------------
# CLI
# ---------------------------------------------------------------------------

def _color_wrap(text: str, code: str, use_color: bool) -> str:
    return f"\033[{code}m{text}\033[0m" if use_color else text


def main() -> None:
    ap = argparse.ArgumentParser(
        description=__doc__,
        formatter_class=argparse.RawDescriptionHelpFormatter,
    )
    ap.add_argument("paths", nargs="*", metavar="PATH",
                    help="Spec .md file(s) or directories to validate. "
                         "Defaults to product/records/spec.")
    ap.add_argument("--root", default=None, metavar="DIR",
                    help="Repository root (for ID derivation). Auto-detected via .git if omitted.")
    ap.add_argument("--strict", action="store_true",
                    help="Treat inventory-phase warnings as errors.")
    ap.add_argument("--no-color", action="store_true",
                    help="Disable ANSI color output.")
    args = ap.parse_args()

    # Resolve repo root
    if args.root:
        repo_root = Path(args.root).resolve()
    else:
        candidate = Path.cwd()
        repo_root = candidate
        while candidate != candidate.parent:
            if (candidate / ".git").exists():
                repo_root = candidate
                break
            candidate = candidate.parent

    # Collect target files
    target_paths = args.paths or [str(repo_root / "product" / "records" / "spec")]
    spec_files: list[Path] = []
    for p_str in target_paths:
        p = Path(p_str).resolve()
        if p.is_dir():
            spec_files.extend(sorted(p.rglob("*.md")))
        elif p.is_file() and p.suffix == ".md":
            spec_files.append(p)
        else:
            print(f"WARNING: path not found or not a .md file: {p}", file=sys.stderr)

    if not spec_files:
        print("No spec files found.", file=sys.stderr)
        sys.exit(1)

    use_color = not args.no_color and sys.stdout.isatty()

    def col(text, code):
        return _color_wrap(text, code, use_color)

    RED, YELLOW, CYAN, BOLD, DIM = "31", "33", "36", "1", "2"

    total_errors = 0
    total_warns = 0
    total_infos = 0
    files_with_issues = 0

    for spec_file in spec_files:
        try:
            rel_path = spec_file.relative_to(repo_root)
        except ValueError:
            rel_path = spec_file

        diags = check_spec(spec_file, repo_root, strict=args.strict)

        errors = [d for d in diags if d["severity"] == "ERROR"]
        warns  = [d for d in diags if d["severity"] == "WARN"]
        infos  = [d for d in diags if d["severity"] == "INFO"]

        total_errors += len(errors)
        total_warns  += len(warns)
        total_infos  += len(infos)

        if errors or warns:
            files_with_issues += 1

        if not diags:
            continue

        print(col(str(rel_path), BOLD))
        for d in diags:
            sev, code, msg = d["severity"], d["code"], d["message"]
            if sev == "ERROR":
                label = col("  ERROR", RED)
            elif sev == "WARN":
                label = col("  WARN ", YELLOW)
            else:
                label = col("  INFO ", CYAN)
            print(f"{label}  [{code}]  {msg}")
        print()

    # Summary line
    mode_label = "strict" if args.strict else "inventory"
    parts = []
    if total_errors:
        parts.append(col(f"{total_errors} error(s)", RED))
    if total_warns:
        parts.append(col(f"{total_warns} warning(s)", YELLOW))
    if total_infos:
        parts.append(col(f"{total_infos} info(s)", CYAN))

    if parts:
        print(f"[{mode_label}]  {files_with_issues} file(s) with issues: " + ", ".join(parts))
    else:
        print(col(f"[{mode_label}]  All {len(spec_files)} file(s) OK.", CYAN))

    sys.exit(1 if total_errors else 0)


if __name__ == "__main__":
    main()
