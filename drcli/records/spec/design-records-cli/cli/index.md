# Index: DRCLI CLI and runtime

- **id**: `spec:drcli.design_records_cli.cli`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli`

## What this is

Navigation entry for DRCLI command spelling, invocation behavior, output transport, and portable runtime contracts.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Command surface | Contract | `spec:drcli.design_records_cli.cli.command_surface` | Fixes the logical executable, command groups, and command-to-operation mapping. |
| Repository selection and traversal | Contract | `spec:drcli.design_records_cli.cli.repository_selection_and_traversal` | Fixes per-invocation repository selection and Git-compatible ignore behavior. |
| Help and usage diagnostics | Contract | `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics` | Makes valid invocation discoverable and defines actionable CLI usage failures. |
| Authoring body input | Contract | `spec:drcli.design_records_cli.cli.authoring_body_input` | Maps the three explicit body-source choices into body-capable authoring commands. |
| Output and exit status | Contract | `spec:drcli.design_records_cli.cli.output_and_exit_status` | Fixes human/JSON modes, stream responsibility, and numeric process status. |
| Runtime and distribution | Contract | `spec:drcli.design_records_cli.cli.runtime_and_distribution` | Fixes portable packaging, bundled standards lookup, retained-state behavior, shared installation, and OS parity. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.operations` | Owns semantic operation contracts mapped by the command surface. |
| `spec:drcli.design_records_cli.diagnostics` | Owns transport-independent semantic operation diagnostics. |