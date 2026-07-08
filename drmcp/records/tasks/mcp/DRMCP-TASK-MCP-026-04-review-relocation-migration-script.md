# DRMCP-TASK-MCP-026-04: Review relocation migration script

- **id**: DRMCP-TASK-MCP-026-04
- **status**: done
- **date**: 2026-07-08
- **work_item**: DRMCP-WORK-MCP-026
- **task_type**: review
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-026-03
- **outputs**:
  - DRMCP-TASK-MCP-026-04

## Goal

Independently review the relocation migration script and dry-run output before any apply-mode migration runs.

## Work

Review:

- the script created by T03;
- the dry-run output from T03;
- the T01 reference inventory;
- the T02 topology decisions.

Check that the planned migration:

- moves only the five application-architecture Specification files;
- creates `drmcp/records/spec/implementation/application-architecture/` as the target tree;
- rewrites old canonical refs only inside `drmcp/records/spec/`;
- uses `spec:drmcp.implementation.application_architecture` as the new root ref;
- leaves no compatibility stub under the old tree;
- does not edit ADR, Requirement, Work Item, Task, or Investigation records;
- does not redesign `spec:drmcp.implementation` root content.

## Done condition

- Review returns `PASS` or `NEEDS REVISION`.
- Any finding names exact script behavior, planned move, planned rewrite, or missing safeguard.
- Apply-mode migration remains blocked unless the review returns `PASS`.

## Verification

| check | result |
|---|---|
| Script move scope | PASS: The script moves only the five expected application-architecture Specification files. |
| Target path | PASS: All move targets are under `drmcp/records/spec/implementation/application-architecture/`. |
| Canonical ref rewrite | PASS: The script rewrites `spec:drmcp.application_architecture` to `spec:drmcp.implementation.application_architecture`. |
| Rewrite scope | PASS: In-place rewrites scan only Markdown files under `drmcp/records/spec/`. |
| Non-spec record exclusion | PASS: ADR, Requirement, Work Item, Task, and Investigation trees are outside the rewrite scan. |
| Compatibility stub policy | PASS: Apply mode removes the old source files and removes the old source directory when empty. |
| Pre-apply safeguards | PASS: The script fails on missing source directory, missing expected source file, existing destination file, unexpected source file, and non-empty target directory. |
| Post-apply safeguard | PASS: Apply mode fails when old refs remain under `drmcp/records/spec/`. |
| Implementation root role | PASS: The dry-run does not rewrite `drmcp/records/spec/implementation/index.md`. |
| Architecture semantics | PASS: The script only moves files and replaces canonical ref prefixes. |
| W018 module-contract semantics | PASS: Planned contract rewrites are limited to canonical ref replacement. |
| T01 coverage | PASS: The dry-run reports 57 replacements across the same scoped spec reference set. |

## Evidence

### Verdict

PASS.

T05 apply is allowed after this review.

### Review basis

Reviewed artifacts:

- `drmcp/scripts/relocate_application_architecture_specs.py`
- T03 dry-run output recorded in `DRMCP-TASK-MCP-026-03`
- `DRMCP-TASK-MCP-026-01`
- `DRMCP-TASK-MCP-026-02`
- `DRMCP-TASK-MCP-026-03`
- `DRMCP-WORK-MCP-026`

Reviewer independence:

- T04 performed independent review only.
- T04 did not repair the script.
- T04 did not edit Specification records.
- T04 did not run apply mode.

### Script review result

| review item | result |
|---|---|
| Move list | The hard-coded `MOVED_FILES` tuple contains exactly the five expected Specification files. |
| Source tree | Source paths are derived from `drmcp/records/spec/application-architecture/`. |
| Target tree | Target paths are derived from `drmcp/records/spec/implementation/application-architecture/`. |
| Ref mapping | `OLD_REF` and `NEW_REF` match the T02 decision. |
| Dry-run behavior | Default mode prints the move and rewrite plan without writing files. |
| Apply gate | File writes occur only when `--apply` is present. |
| Rewrite scan | `build_plan` scans `(root / SPEC_ROOT).rglob("*.md")` and skips moved source files. |
| Precondition failures | `assert_preconditions` rejects missing source tree, missing source files, existing destinations, unexpected source files, and non-empty target directory. |
| Leftover check | `find_old_ref_leftovers` scans Markdown under `drmcp/records/spec/` after apply. |
| Old tree removal | Apply mode unlinks moved source files and removes the old source directory when empty. |

### Dry-run review result

The dry-run output lists five planned moves.
Each move targets `drmcp/records/spec/implementation/application-architecture/`.
The planned in-place rewrites list only Markdown files under `drmcp/records/spec/`.
The planned in-place rewrites include `drmcp/records/spec/design-records-mcp/namespace-scanning.md`.
The planned in-place rewrites include the expected `drmcp/records/spec/implementation/contracts/**` files.
The total replacement count is 57.
The replacement count matches the T01 scoped search result for `spec:drmcp.application_architecture`.

### Findings

None.

### Boundary confirmation

No apply mode was run.
No file move was performed.
No Specification was edited.
No ADR was created.
No stage, commit, or push was performed.
