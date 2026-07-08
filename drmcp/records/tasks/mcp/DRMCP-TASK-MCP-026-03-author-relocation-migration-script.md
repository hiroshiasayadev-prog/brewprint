# DRMCP-TASK-MCP-026-03: Author relocation migration script

- **id**: DRMCP-TASK-MCP-026-03
- **status**: done
- **date**: 2026-07-08
- **work_item**: DRMCP-WORK-MCP-026
- **task_type**: implementation
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-026-02
- **outputs**:
  - `drmcp/scripts/relocate_application_architecture_specs.py`

## Goal

Author a deterministic migration script for relocating the application-architecture Specification tree under `spec:drmcp.implementation.application_architecture`.

## Work

Before authoring, read:

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/writing-standard.md`
3. `product/records/spec/design-records/authoring-standards/task-authoring.md`
4. `DRMCP-TASK-MCP-026-01`
5. `DRMCP-TASK-MCP-026-02`

The script must support a reviewable dry-run mode.
The dry-run output must list every planned file move and text replacement.
The script must fail when an expected source file is missing or a destination file already exists.
The script must not rewrite ADR, Requirement, Work Item, Task, or Investigation records.

## Implementation contract

| target | required change | acceptance criterion | verification |
|---|---|---|---|
| Migration script under `tools/` or another reviewed tooling location | Create a deterministic script that moves the five application-architecture Specification files into `drmcp/records/spec/implementation/application-architecture/` and rewrites spec-scope refs from `spec:drmcp.application_architecture` to `spec:drmcp.implementation.application_architecture`. | Dry-run lists the exact move and replacement plan without modifying files. Apply mode is explicit and fails on missing sources, existing destinations, or unexpected old-ref leftovers inside `drmcp/records/spec/`. | Run dry-run and inspect the produced plan before T04 review. |

## Done condition

- The migration script exists.
- Dry-run mode produces a deterministic plan.
- The plan includes the five file moves and the spec-scope ref rewrites from T01 and T02.
- The script does not perform production implementation work.
- The script does not run apply mode in this Task unless explicitly authorized by a later route.

## Verification

- Migration script was authored at `drmcp/scripts/relocate_application_architecture_specs.py`.
- The script defaults to dry-run mode.
- The script requires explicit `--apply` before writing files.
- Dry-run was executed by the user on 2026-07-08.
- Apply mode was not run.
- Dry-run returned the expected five planned moves.
- Dry-run returned 12 planned in-place rewrite files under `drmcp/records/spec/`.
- Dry-run returned 57 total ref replacements, matching the T01 search result.

## Evidence

Created script:

- `drmcp/scripts/relocate_application_architecture_specs.py`

Executed dry-run command:

```text
python drmcp/scripts/relocate_application_architecture_specs.py --repo-root C:\Users\imved\projects\brewprint
```

Dry-run output summary:

| item | value |
|---|---:|
| planned moves | 5 |
| planned in-place rewrite files | 12 |
| total ref replacements | 57 |

Planned moves:

| source | target | replacements |
|---|---|---:|
| `drmcp/records/spec/application-architecture/index.md` | `drmcp/records/spec/implementation/application-architecture/index.md` | 10 |
| `drmcp/records/spec/application-architecture/application-boundary-and-components.md` | `drmcp/records/spec/implementation/application-architecture/application-boundary-and-components.md` | 7 |
| `drmcp/records/spec/application-architecture/dependency-and-responsibility.md` | `drmcp/records/spec/implementation/application-architecture/dependency-and-responsibility.md` | 10 |
| `drmcp/records/spec/application-architecture/runtime-and-state.md` | `drmcp/records/spec/implementation/application-architecture/runtime-and-state.md` | 8 |
| `drmcp/records/spec/application-architecture/failure-and-evolution.md` | `drmcp/records/spec/implementation/application-architecture/failure-and-evolution.md` | 7 |

Planned in-place rewrites:

| file | replacements |
|---|---:|
| `drmcp/records/spec/design-records-mcp/namespace-scanning.md` | 4 |
| `drmcp/records/spec/implementation/contracts/application-use-cases/contract-boundary.md` | 1 |
| `drmcp/records/spec/implementation/contracts/application-use-cases/index.md` | 1 |
| `drmcp/records/spec/implementation/contracts/composition-lifecycle/contract-boundary.md` | 1 |
| `drmcp/records/spec/implementation/contracts/composition-lifecycle/index.md` | 1 |
| `drmcp/records/spec/implementation/contracts/index.md` | 1 |
| `drmcp/records/spec/implementation/contracts/infrastructure-io-adapters/contract-boundary.md` | 1 |
| `drmcp/records/spec/implementation/contracts/infrastructure-io-adapters/index.md` | 1 |
| `drmcp/records/spec/implementation/contracts/mcp-inbound-adapter/contract-boundary.md` | 1 |
| `drmcp/records/spec/implementation/contracts/mcp-inbound-adapter/index.md` | 1 |
| `drmcp/records/spec/implementation/contracts/record-domain-logical-tree/contract-boundary.md` | 1 |
| `drmcp/records/spec/implementation/contracts/record-domain-logical-tree/index.md` | 1 |

Intended apply command after T04 review pass:

```text
python drmcp/scripts/relocate_application_architecture_specs.py --repo-root C:\Users\imved\projects\brewprint --apply
```

Script behavior:

- Auto-detects the repository root when `--repo-root` is omitted.
- Uses `spec:drmcp.application_architecture` as the old ref.
- Uses `spec:drmcp.implementation.application_architecture` as the new ref.
- Moves only the five expected application-architecture Specification files.
- Writes target files under `drmcp/records/spec/implementation/application-architecture/`.
- Rewrites only Markdown files under `drmcp/records/spec/`.
- Fails when an expected source file is missing.
- Fails when a destination file already exists.
- Fails when the old source directory contains unexpected files.
- Fails after apply when old refs remain under `drmcp/records/spec/`.
- Removes the old source directory after apply when the directory is empty.

T03 is complete.
T04 can review the script and dry-run output before apply mode.
