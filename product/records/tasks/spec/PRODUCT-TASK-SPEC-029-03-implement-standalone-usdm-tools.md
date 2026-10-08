# PRODUCT-TASK-SPEC-029-03: Implement standalone USDM tools

- **id**: PRODUCT-TASK-SPEC-029-03
- **status**: done
- **date**: 2026-07-09
- **work_item**: PRODUCT-WORK-SPEC-029
- **task_type**: implementation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-029-01
  - PRODUCT-TASK-SPEC-029-02
- **outputs**:
  - `tools/usdm/`
  - PRODUCT-TASK-SPEC-029-03
  - PRODUCT-WORK-SPEC-029
  - `.gitignore`

## Startup / Required reading

Before implementation, preserve these records and rules:

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/writing-standard.md`
3. `product/records/spec/design-records/authoring-standards/task-authoring.md`
4. `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
5. `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`
6. `product/records/spec/design-records/authoring-standards/usdm-authoring.md`
7. `product/records/spec/design-records/usdm/index.md`
8. `product/records/spec/design-records/usdm/artifact-format.md`
9. `product/records/spec/design-records/usdm/coverage-format.md`
10. `product/records/spec/design-records/usdm/coverage-tools.md`
11. `product/records/requirements/spec/PRODUCT-REQ-SPEC-015-mvp-usdm-requirement-artifacts-and-coverage-checks.md`
12. `product/records/work-items/spec/PRODUCT-WORK-SPEC-029-mvp-usdm-artifacts-and-coverage-tooling.md`
13. `product/records/tasks/spec/PRODUCT-TASK-SPEC-029-01-author-mvp-usdm-artifact-specification.md`
14. `product/records/tasks/spec/PRODUCT-TASK-SPEC-029-02-author-usdm-authoring-standard.md`

## Goal

Implement standalone MVP USDM validation and coverage commands under `tools/usdm/`.

## Boundary

This Task owns only the standalone Python tooling, minimal repository tracking support for `tools/usdm/`, this Task record, and parent Work Item task registration.

This Task does not own real USDM requirement records, smoke fixtures under app records, USDM spec changes, DRMCP or MCP integration, existing-record migration, DRMCP Domain architecture decisions, or PRODUCT-WORK-SPEC-029 closure.

## Work

- Added `tools/usdm/usdm_tools.py` with the `validate_usdm`, `check_usdm_coverage`, and `usdm_covered_by` subcommands.
- Implemented app namespace discovery from `<app>/records/usdm/` and `<app>/records/spec/`.
- Implemented fenced-code-aware H1 and H2 scanning for USDM records and Specification metadata blocks.
- Implemented USDM requirement table validation, full requirement ID collection, coverage checks, dangling coverage checks, and coverage lookup.
- Updated `.gitignore` so `tools/usdm/` is trackable while other ignored tool output remains ignored.
- Registered PRODUCT-TASK-SPEC-029-03 on PRODUCT-WORK-SPEC-029.

## Implementation contract

| target | required change | acceptance criterion | verification |
|---|---|---|---|
| `tools/usdm/usdm_tools.py` | Provide a standalone Python CLI with `validate_usdm`, `check_usdm_coverage`, and `usdm_covered_by`. | Each subcommand prints JSON and exits `0` only when `ok` is true. | `uv run python tools/usdm/usdm_tools.py validate_usdm --repo-root C:\Users\imved\projects\brewprint`; `uv run python tools/usdm/usdm_tools.py check_usdm_coverage --repo-root C:\Users\imved\projects\brewprint`; `python -m py_compile tools\usdm\usdm_tools.py`. |
| USDM validation behavior | Validate MVP USDM H1, metadata, kind, ID grammar, requirement sections, source refs, tables, row IDs, and duplicate full IDs. | Empty current USDM inventory returns `ok: true` with zero records and zero requirements. Invalid discovered records produce error diagnostics. | `validate_usdm` command output recorded below. |
| USDM coverage behavior | Read full USDM requirement IDs and Specification `usdm_covers` metadata. | Coverage reports uncovered IDs, dangling entries, duplicate coverage, malformed coverage, and covering Specification refs. | `check_usdm_coverage` command output recorded below. |
| `.gitignore` | Preserve ignored tool output while unignoring only `tools/usdm/`. | `tools/usdm/usdm_tools.py` appears as a normal untracked file without `git add -f`. | `git status --short -- .gitignore tools/usdm`. |
| PRODUCT records | Create this Task and add it to PRODUCT-WORK-SPEC-029. | Task metadata references PRODUCT-WORK-SPEC-029 and Work Item metadata lists PRODUCT-TASK-SPEC-029-03. | Scoped status and whitespace checks recorded below. |

## Done condition

- The standalone USDM tool CLI exists under `tools/usdm/`.
- The CLI implements the three required subcommands.
- Required validation and coverage checks pass for the current repository state.
- PRODUCT-WORK-SPEC-029 lists PRODUCT-TASK-SPEC-029-03.
- Verification evidence is recorded in this Task.

## Verification

Commands executed:

```text
uv run python tools/usdm/usdm_tools.py validate_usdm --repo-root C:\Users\imved\projects\brewprint
uv run python tools/usdm/usdm_tools.py check_usdm_coverage --repo-root C:\Users\imved\projects\brewprint
uv run python tools/usdm/usdm_tools.py usdm_covered_by --repo-root C:\Users\imved\projects\brewprint --requirement-id usdm:drmcp.example#R001
python -m py_compile tools\usdm\usdm_tools.py
rg -n "responsibility_boundary_validator|validate.*responsibility|semantic" product tools scripts -g "*.py" -g "*.md"
git status --short -- .gitignore tools/usdm product/records/work-items/spec/PRODUCT-WORK-SPEC-029-mvp-usdm-artifacts-and-coverage-tooling.md product/records/tasks/spec/PRODUCT-TASK-SPEC-029-03-implement-standalone-usdm-tools.md
git diff --check -- .gitignore tools/usdm product/records/work-items/spec/PRODUCT-WORK-SPEC-029-mvp-usdm-artifacts-and-coverage-tooling.md product/records/tasks/spec/PRODUCT-TASK-SPEC-029-03-implement-standalone-usdm-tools.md
```

Results:

| command | result |
|---|---|
| `validate_usdm` | exit 0; `ok: true`; `usdm_records: 0`; `requirements: 0`; `diagnostics: []`. |
| `check_usdm_coverage` | exit 0; `ok: true`; `requirements: 0`; `covered: 0`; `uncovered: []`; `dangling: []`; `diagnostics: []`. |
| `usdm_covered_by` | exit 1; `ok: false`; `exists: false`; `covered_by: []`; missing requirement diagnostic for the well-formed absent ID. |
| `python -m py_compile tools\usdm\usdm_tools.py` | exit 0. |
| responsibility validator search | exit 1; no matching standalone validator command found under the scoped paths. |
| scoped git status | `.gitignore` modified; `tools/usdm/`, PRODUCT-WORK-SPEC-029, and this Task are unstaged. |
| scoped whitespace check | exit 0; Git emitted an LF-to-CRLF advisory for `.gitignore`. |

## Evidence

- `tools/usdm/usdm_tools.py` implements the standalone MVP CLI without OCR, AI calls, or network access.
- The current repository has no `<app>/records/usdm/**/*.md` records, so zero-count successful validation and coverage results are expected.
- `.gitignore` changed from ignoring the whole `tools/` directory to ignoring `tools/*` while unignoring `tools/usdm/` and its contents.
- PRODUCT-WORK-SPEC-029 now lists PRODUCT-TASK-SPEC-029-03 in metadata.
- No files were staged, committed, or pushed.
