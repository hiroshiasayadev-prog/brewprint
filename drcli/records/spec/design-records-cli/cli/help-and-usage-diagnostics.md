# Contract: DRCLI help and usage diagnostics

- **id**: `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.cli`
- **contract_class**: `interface`

## What this is

Defines repository-independent CLI help and actionable diagnostics for invalid CLI usage.
A caller must be able to discover valid command usage without MCP schema knowledge.

## Non-goals

- Redefining semantic operation diagnostics after a type-valid logical request exists.
- Fixing operation-specific option spelling that the accepted CLI decisions did not decide.
- Requiring repository, bundled-standards, or retained-state access to render help.

## Request

The following invocations request help and succeed:

| invocation | help scope |
|---|---|
| `drcli` | Root help. |
| `drcli --help` | Root help. |
| `drcli help` | Root help. |
| `drcli <group> --help` | Group help. |
| `drcli <group> <command> --help` | Command help. |
| `drcli help <group> <command>` | Command help. |

Help lookup does not open a selected repository.
Help lookup does not require the bundled standards tree.
Help lookup does not access body caches or retained proposals.

Root help lists every command group and an exact route to more specific help.
Group help lists every valid child command in that group.
Command help exposes every accepted CLI operand and option for that command.
Command help also exposes mutually exclusive inputs, required inputs, and applicable global options.
Each command help includes at least one valid invocation example.
Each command help shows an exact help command that can reproduce that scoped help.

Explicit help remains human-readable help text even when `--format json` appears elsewhere in an invocation context.

## Response

Successful help is written as human-readable text and uses process status `0`.
Help content describes the actual accepted CLI surface of the selected scope.
Help must not require a caller to infer valid usage from semantic request schemas.

An invalid-usage diagnostic contains enough information to correct the invocation in one step.
Human and machine-readable diagnostics carry the same actionable usage information.

## Errors

| invalid-usage class | required diagnostic content |
|---|---|
| Unknown command or group | Identify the supplied token, list valid siblings, and give the exact scoped help command. |
| Unknown option | Identify the option, show the valid option surface for that scope, and give the exact scoped help command. |
| Missing operand or option value | Name the missing input, show an accepted invocation shape, and give the exact scoped help command. |
| Mutually exclusive inputs | Name every conflicting input, show the accepted alternatives, and give the exact scoped help command. |
| Malformed or invalid selector combination | Show accepted selector shapes and give the exact scoped help command. |

Every row above is a CLI `usage_error` and therefore uses process status `2`.
A well-formed selector that resolves to no semantic target is not a CLI usage error.
Such a selector proceeds to the mapped operation and keeps that operation's normal result semantics.

In human output mode, the usage diagnostic is written to stderr as readable text.
In JSON output mode, stderr contains one machine-readable diagnostic document.
The machine-readable document preserves the offending input, accepted alternative or usage shape, and exact help route required by the applicable row.
No parser-only message may replace the required actionable content.

## Boundary

This contract ends once CLI syntax has produced a type-valid logical operation request.
Semantic request validation and operation diagnostics then belong to the mapped operation and `spec:drcli.design_records_cli.diagnostics`.

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.cli.command_surface` | Supplies groups and commands that help must expose. |
| `spec:drcli.design_records_cli.cli.output_and_exit_status` | Owns stderr placement, JSON diagnostic framing, and status `2`. |
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_envelope` | Begins after CLI usage parsing and owns semantic operation diagnostics. |