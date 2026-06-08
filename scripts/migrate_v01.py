#!/usr/bin/env python3
"""
V01 migration: docs/ -> v01/records/ + V01- prefix on all artifact IDs.

Steps:
  1. git mv docs v01/records
  2. Rename artifact files (ADR NNN->V01-ADR-NNN, others add V01- prefix)
  3. Update content: ADR H1s, all artifact ID cross-refs, spec design_record.id
"""

import re
import subprocess
import sys
from pathlib import Path

REPO = Path(__file__).resolve().parent.parent

# ---------------------------------------------------------------------------
# Regex patterns
# ---------------------------------------------------------------------------

# Matches standalone artifact IDs in content.
# Negative lookbehind/lookahead: don't match if surrounded by [A-Z0-9-]
# so that already-prefixed "V01-ADR-097" doesn't re-match "ADR-097".
_DOM = r'[A-Z][A-Z0-9]*(?:-[A-Z][A-Z0-9]*)*'  # domain segment(s) like DATA, MCP, SELF-HOSTING

ARTIFACT_ID_RE = re.compile(
    r'(?<![A-Z0-9-])('
    r'ADR-\d{3}'
    r'|(?:INV|REQ|WORK)-' + _DOM + r'-\d{3}'
    r'|TASK-' + _DOM + r'-\d{3}-\d{2}'
    r')(?![A-Z0-9-])',
    re.ASCII
)

# ADR H1: "# 097: title" -> "# V01-ADR-097: title"
ADR_H1_RE = re.compile(r'^(# )(\d{3})(:.*)$', re.MULTILINE)

# SPEC design_record.id in YAML front matter: "  id: SPEC-foo-bar"
SPEC_DR_ID_RE = re.compile(r'^(\s+id:\s+)(SPEC-[a-zA-Z][a-zA-Z0-9-]*)$', re.MULTILINE)

# Filename patterns
ADR_FNAME_RE = re.compile(r'^(\d{3})(-.+\.md)$')
ARTIFACT_FNAME_RE = re.compile(
    r'^((?:INV|REQ|WORK)-' + _DOM + r'-\d{3}'
    r'|TASK-' + _DOM + r'-\d{3}-\d{2}'
    r')(-.+\.md)$',
    re.ASCII
)

# ---------------------------------------------------------------------------
# Helpers
# ---------------------------------------------------------------------------

def update_content(path: Path, is_adr: bool, is_spec: bool) -> bool:
    text = path.read_text(encoding='utf-8')
    out = text
    if is_adr:
        out = ADR_H1_RE.sub(r'\1V01-ADR-\2\3', out)
    if is_spec:
        out = SPEC_DR_ID_RE.sub(r'\1V01-\2', out)
    out = ARTIFACT_ID_RE.sub(r'V01-\1', out)
    if out != text:
        path.write_text(out, encoding='utf-8')
        return True
    return False


def new_fname(name: str) -> str:
    m = ADR_FNAME_RE.match(name)
    if m:
        return f'V01-ADR-{m.group(1)}{m.group(2)}'
    m = ARTIFACT_FNAME_RE.match(name)
    if m:
        return f'V01-{m.group(1)}{m.group(2)}'
    return name


# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------

def main() -> None:
    v01_records = REPO / 'v01' / 'records'

    if v01_records.exists():
        print(f'ERROR: {v01_records} already exists — migration already run?', file=sys.stderr)
        sys.exit(1)

    # Step 1: git mv docs v01/records
    print('Step 1: git mv docs v01/records')
    v01_records.parent.mkdir(parents=True, exist_ok=True)
    subprocess.run(
        ['git', 'mv', 'docs', str(v01_records.relative_to(REPO))],
        cwd=REPO, check=True
    )

    # Step 2: rename artifact files
    print('Step 2: rename artifact files')

    adr_dir = v01_records / 'adr'
    for p in sorted(adr_dir.glob('*.md')):
        n = new_fname(p.name)
        if n != p.name:
            p.rename(p.parent / n)
            print(f'  {p.name} -> {n}')

    for subdir in ['investigations', 'requirements', 'work-items', 'tasks']:
        d = v01_records / subdir
        if not d.exists():
            continue
        for p in sorted(d.rglob('*.md')):
            n = new_fname(p.name)
            if n != p.name:
                p.rename(p.parent / n)
                print(f'  {p.name} -> {n}')

    # Step 3: update content in all .md files
    print('Step 3: update content')
    changed = 0
    adr_dir_resolved = (v01_records / 'adr').resolve()
    spec_dir_resolved = (v01_records / 'spec').resolve()

    for p in sorted(v01_records.rglob('*.md')):
        is_adr = p.resolve().parent == adr_dir_resolved
        is_spec = spec_dir_resolved in p.resolve().parents
        if update_content(p, is_adr=is_adr, is_spec=is_spec):
            changed += 1

    print(f'  updated {changed} files')
    print('Done. Run: git add -A && git status')


if __name__ == '__main__':
    main()
