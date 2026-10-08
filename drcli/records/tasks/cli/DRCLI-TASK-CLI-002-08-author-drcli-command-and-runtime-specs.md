# DRCLI-TASK-CLI-002-08: Author DRCLI command and runtime specs

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: authoring
- **estimate**: 1d-2d
- **depends_on**:
  - DRCLI-TASK-CLI-002-02
- **outputs**:
  - spec:drcli.design_records_cli.cli
  - spec:drcli.design_records_cli.cli.command_surface
  - spec:drcli.design_records_cli.cli.repository_selection_and_traversal
  - spec:drcli.design_records_cli.cli.help_and_usage_diagnostics
  - spec:drcli.design_records_cli.cli.authoring_body_input
  - spec:drcli.design_records_cli.cli.output_and_exit_status
  - spec:drcli.design_records_cli.cli.runtime_and_distribution

## Goal

Project the accepted DRCLI-specific external decisions into canonical CLI/runtime Specification contracts.

## Work

Write only:

- `drcli/records/spec/design-records-cli/cli/**`

Consume the complete T02 decision ledger.

The Specification area must cover:

- command hierarchy and operation mapping;
- repository selection;
- Git-compatible ignore behavior;
- portable Windows/Linux distribution and bundled standards lookup;
- human and machine output modes;
- stdout/stderr and exit-status semantics;
- root and command-specific help;
- actionable invalid-usage diagnostics with valid invocation shape/example and exact help route;
- unknown command/option handling;
- missing and mutually exclusive input handling;
- body input mapping for inline, stdin, and body_cache_id;
- process-independent operational-state behavior;
- one shared installation serving multiple repositories.

Do not rewrite semantic operation contracts owned by T03-T07.
Do not choose implementation technology beyond what T02 explicitly decided.

## Done condition

- Every T02 decision is projected into exactly one appropriate CLI/runtime contract.
- A caller can discover correct usage from the CLI itself without an MCP schema.
- Invalid usage never degrades to an opaque parser-only failure.
- Windows and Linux observable behavior is equivalent.
- CLI/runtime specs do not duplicate operation semantics.

## Verification

- Trace every T02 decision to a section or contract.
- Confirm help is available without repository access.
- Confirm every invalid-usage class names accepted alternatives or a valid usage shape and relevant help route.
- Confirm only `cli/**` changed.

## Evidence

Authoring consumed the complete terminal decision ledger in `DRCLI-TASK-CLI-002-02`.
All T02 ADR routes are `not_required`; no ADR authoring was introduced.

### Decision projection

| T02 decision | canonical projection |
|---|---|
| D-001 | `spec:drcli.design_records_cli.cli.command_surface` — executable and command groups. |
| D-002 | `spec:drcli.design_records_cli.cli.command_surface` — command-to-operation mapping. |
| D-003 | `spec:drcli.design_records_cli.cli.repository_selection_and_traversal` — cwd default and `--repo`. |
| D-004 | `spec:drcli.design_records_cli.cli.repository_selection_and_traversal` — Git-compatible ignore default and `--no-ignore`. |
| D-005 | `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics` — root and scoped help. |
| D-006 | `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics` — actionable invalid-usage classes. |
| D-007 | `spec:drcli.design_records_cli.cli.authoring_body_input` — inline, stdin, and body-cache input mapping. |
| D-008 | `spec:drcli.design_records_cli.cli.output_and_exit_status` — explicit human/JSON mode. |
| D-009 | `spec:drcli.design_records_cli.cli.output_and_exit_status` — stdout/stderr responsibility and JSON framing. |
| D-010 | `spec:drcli.design_records_cli.cli.output_and_exit_status` — statuses 0 through 4. |
| D-011 | `spec:drcli.design_records_cli.cli.runtime_and_distribution` — portable package and bundled standards lookup. |
| D-012 | `spec:drcli.design_records_cli.cli.runtime_and_distribution` — process-independent body-cache and proposal state. |
| D-013 | `spec:drcli.design_records_cli.cli.runtime_and_distribution` — one shared installation for multiple repositories. |
| D-014 | `spec:drcli.design_records_cli.cli.runtime_and_distribution` — Windows/Linux observable parity and UTF-8. |

### Outputs

Created `drcli/records/spec/design-records-cli/cli/index.md` and the six child contracts listed in Task `outputs`.
No semantic operation contract under T03-T07 was rewritten.

### Verification result

- Every T02 decision maps to exactly one child CLI/runtime contract.
- Root, group, and command help require no repository, bundled-standards, or retained-state access.
- Command help must expose accepted operands/options and at least one valid invocation example.
- Every T02 invalid-usage class requires corrective content and an exact scoped help route; parser-only failure text is insufficient.
- Human and JSON modes preserve the same actionable usage information.
- Windows and Linux share logical commands, options, help/diagnostic meaning, JSON contract, body rules, UTF-8 streams, and numeric statuses.
- The authoring boundary contains only `drcli/records/spec/design-records-cli/cli/**` plus this Task record.
- No implementation technology, storage engine, file layout, retention duration, or cleanup mechanism was selected.
- The standalone PRODUCT Task-responsibility validator remains operationally unavailable as already recorded by T02; no semantic validator result was synthesized.
- No independent review was performed in this authoring Task. Integrated review remains a downstream responsibility.
- Blocker: none.
